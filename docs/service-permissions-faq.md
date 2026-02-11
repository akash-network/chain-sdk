# Answers to Service Permissions Questions

## Question 1: How is it done in Kubernetes by default?

**Default Behavior**: In Kubernetes, services (pods) have **NO access** to logs and events by default.

### Details:

When you deploy a pod in Kubernetes without specifying a ServiceAccount:
- The pod automatically uses the `default` ServiceAccount in its namespace
- The `default` ServiceAccount has **minimal permissions**
- It can only read its own ServiceAccount token information
- It **CANNOT**:
  - Read logs from any pod (including itself)
  - Read events from the cluster
  - List or watch other pods
  - Access deployment information
  - Perform any administrative operations

This is a security best practice called **"principle of least privilege"** - services should only have the minimum permissions they need to function.

### To grant access in standard Kubernetes, you must:

1. Create a ServiceAccount
2. Create a Role with specific permissions
3. Create a RoleBinding to connect the ServiceAccount to the Role
4. Configure your pod to use that ServiceAccount

Example:
```yaml
# ServiceAccount
apiVersion: v1
kind: ServiceAccount
metadata:
  name: log-reader

---
# Role with log reading permissions
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-log-reader
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list"]

---
# RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-logs
subjects:
- kind: ServiceAccount
  name: log-reader
roleRef:
  kind: Role
  name: pod-log-reader
  apiGroup: rbac.authorization.k8s.io

---
# Pod using the ServiceAccount
apiVersion: v1
kind: Pod
metadata:
  name: my-app
spec:
  serviceAccountName: log-reader  # Important!
  containers:
  - name: app
    image: myapp:latest
```

## Question 2: What is the use case for a service to have access to logs and events?

There are several legitimate use cases where applications need to access their own logs and events:

### Use Case 1: Self-Monitoring Applications
**Example**: A web application with an admin dashboard

Your application has a built-in admin panel at `/admin/logs` that displays recent application logs to help developers debug issues without needing `kubectl` access.

**Why needed**:
- Developers can see logs through the app UI
- No need to give developers cluster access
- Faster troubleshooting during incidents

**Permission needed**: `logs`

### Use Case 2: Log Analysis and Aggregation
**Example**: Custom log processing service

An application that:
- Reads its own logs
- Analyzes them for patterns (e.g., error rates, slow queries)
- Sends aggregated metrics to a monitoring system
- Triggers alerts based on log content

**Why needed**:
- Custom log analysis logic specific to your application
- Integration with existing monitoring tools
- Real-time log-based alerting

**Permission needed**: `logs`

### Use Case 3: Deployment Status Tracking
**Example**: Status page or deployment dashboard

An application that displays:
- "Last deployed: 2 hours ago"
- "Current replicas: 3/3 ready"
- "Recent events: Successfully scheduled, Pod started"

**Why needed**:
- User-facing status information
- Automated rollback decisions based on events
- Deployment health monitoring

**Permission needed**: `events`

### Use Case 4: Self-Debugging Tools
**Example**: Development/staging environment with debug tools

A development version of your app that includes:
- Interactive log viewer
- Event timeline
- Container restart history
- Performance metrics derived from logs

**Why needed**:
- Faster development iteration
- Easier debugging in staging environments
- Reduced need for cluster access tools

**Permission needed**: `logs`, `events`

### Use Case 5: Audit and Compliance
**Example**: Application that maintains audit logs

An application that:
- Tracks all access to sensitive data (from application logs)
- Records deployment changes (from events)
- Generates compliance reports
- Maintains an audit trail

**Why needed**:
- Compliance requirements (SOC2, HIPAA, etc.)
- Security auditing
- Change tracking

**Permission needed**: `logs`, `events`

### Use Case 6: Intelligent Error Recovery
**Example**: Application with self-healing capabilities

An application that:
- Monitors its own logs for specific error patterns
- Checks events for restart/crash information
- Takes automatic recovery actions (e.g., clearing cache, reconnecting to DB)
- Reports to external systems when manual intervention is needed

**Why needed**:
- Automated incident response
- Reduced downtime
- Self-healing capabilities

**Permission needed**: `logs`, `events`

## Important Notes

### Security Scope
- Permissions are **scoped to the service's own resources only**
- A service can only read logs from its own pods
- A service can only see events related to its own deployment
- **No cross-deployment or cluster-wide access** is granted

### When NOT to Use Permissions
Most applications **don't need** these permissions:
- ❌ Simple web servers (nginx, apache)
- ❌ Databases (postgres, mysql)
- ❌ Cache services (redis, memcached)
- ❌ Basic microservices that just handle requests

Only add permissions when you have a **specific, documented reason** your application code needs to access logs or events.

### Best Practice
**Start without permissions** and only add them when you:
1. Know exactly why you need them
2. Have implemented the code to use them
3. Can document the use case
4. Have tested without them first

## How Akash Simplifies This

With Akash SDL, instead of creating ServiceAccount, Role, and RoleBinding manually:

**Before** (3 Kubernetes resources):
```yaml
# 30+ lines of YAML
ServiceAccount + Role + RoleBinding
```

**After** (Akash SDL):
```yaml
services:
  myapp:
    image: myapp:latest
    params:
      permissions:
        read:
          - logs
          - events
```

Akash automatically creates and configures all necessary Kubernetes RBAC resources for you.

## References

- Full documentation: [docs/service-permissions.md](../docs/service-permissions.md)
- Kubernetes RBAC: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
- Service Accounts: https://kubernetes.io/docs/concepts/security/service-accounts/
