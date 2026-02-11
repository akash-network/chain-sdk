# Service Permissions in Akash

## Overview

This document explains how service-level permissions work in Akash deployments and their relationship to Kubernetes RBAC (Role-Based Access Control).

## Kubernetes Default Behavior

### Default Service Account Permissions

By default in Kubernetes:

- **Every pod gets a service account** - If you don't specify one, the pod uses the `default` service account in its namespace
- **Default service accounts have minimal permissions** - They can only:
  - Access the Kubernetes API server (with limited permissions)
  - Read their own service account information
- **Default service accounts CANNOT:**
  - Read logs from pods
  - Read events from the cluster
  - List or watch other pods
  - Access deployment information
  - Perform any administrative actions

This restrictive default is a security best practice following the **principle of least privilege** - services should only have access to resources they explicitly need.

### Why Services Need Explicit Permissions

In a standard Kubernetes deployment, applications running inside pods cannot access cluster resources (like logs or events) unless:
1. They are explicitly granted RBAC permissions via a Role/RoleBinding
2. A custom ServiceAccount is created with appropriate permissions
3. The pod is configured to use that ServiceAccount

## Use Cases for Service Permissions

### 1. Self-Monitoring and Observability

**Use Case**: A service that monitors its own health and performance

**Example Applications**:
- A web application that displays its own logs in a dashboard
- A monitoring sidecar that collects and exports metrics from the main container
- An application that needs to track its own restart events for debugging

**Why it needs permissions**:
```yaml
permissions:
  read:
    - logs  # Read pod logs to display or analyze
    - events  # Monitor restart events, scheduling issues, etc.
```

**Real-world example**: A Node.js application with an admin panel that shows recent application logs and deployment events to help developers debug issues without needing kubectl access.

### 2. Log Aggregation and Analysis

**Use Case**: A service that aggregates logs for analysis

**Example Applications**:
- Custom log aggregators (similar to Fluentd or Logstash)
- Services that parse and analyze their own logs for security threats
- Applications that need to correlate logs with deployment events

**Why it needs permissions**:
```yaml
permissions:
  read:
    - logs  # Read logs from all containers in the pod
```

**Real-world example**: A Python application that analyzes its own error logs to detect patterns and automatically create incident reports.

### 3. Deployment Status and Health Checking

**Use Case**: Applications that report their deployment status

**Example Applications**:
- Status page applications that show deployment health
- CI/CD tools that monitor their own deployment
- Applications that need to know when they were last restarted

**Why it needs permissions**:
```yaml
permissions:
  read:
    - events  # Read events to track deployment status
```

**Real-world example**: A service dashboard that shows "Last deployed: 2 hours ago" and "Current health: Running" by reading deployment events.

### 4. Debugging and Development Tools

**Use Case**: Development utilities embedded in applications

**Example Applications**:
- Debug panels in staging/development environments
- Interactive troubleshooting tools
- Developer consoles that need access to logs

**Why it needs permissions**:
```yaml
permissions:
  read:
    - logs
    - events
```

**Real-world example**: A staging environment with a built-in debug panel that developers can access to see recent logs and events without needing direct cluster access.

## Permission Types

### `logs` Permission

Grants read access to:
- **pods** - List and get pod information
- **pods/log** - Read container logs from pods
- **deployments** - Read deployment information

This is translated to Kubernetes RBAC rules that allow the service to:
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: logs-reader
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list"]
- apiGroups: ["apps"]
  resources: ["deployments"]
  verbs: ["get", "list"]
```

### `events` Permission

Grants read access to:
- **events** - Read Kubernetes events

This is translated to Kubernetes RBAC rules that allow the service to:
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: events-reader
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs: ["get", "list", "watch"]
```

## SDL Example

Here's how to specify permissions in your Service Definition Language (SDL):

```yaml
version: "2.0"

services:
  web:
    image: myapp:latest
    expose:
      - port: 8080
        as: 80
        to:
          - global: true
    params:
      permissions:
        read:
          - logs      # Allow reading pod logs
          - events    # Allow reading cluster events
```

### Minimal Example (No Permissions)

If your service doesn't need access to logs or events, you can omit the permissions block:

```yaml
version: "2.0"

services:
  web:
    image: nginx:latest
    expose:
      - port: 80
        as: 80
        to:
          - global: true
    # No params block needed - uses default (no extra permissions)
```

## Security Considerations

### Principle of Least Privilege

Only request permissions your service actually needs:

- ❌ **Don't** blindly add all permissions to every service
- ✅ **Do** only request `logs` if your service reads logs
- ✅ **Do** only request `events` if your service monitors events
- ✅ **Do** omit permissions entirely if your service doesn't need them

### Scope Limitation

The permissions granted are **scoped to the service's own namespace and resources**:

- A service can only read its own pod's logs, not logs from other deployments
- A service can only see events related to its own deployment
- Cross-namespace or cluster-wide access is not granted

### Common Pitfalls

1. **Don't request permissions "just in case"** - This violates security best practices
2. **Test without permissions first** - Only add permissions when you encounter permission errors
3. **Document why permissions are needed** - Help reviewers understand the use case

## Migration from Traditional Kubernetes

If you're migrating from traditional Kubernetes deployments where you manually created ServiceAccounts and RBAC rules:

**Before (Kubernetes manifests)**:
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: myapp-sa
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: myapp-role
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: myapp-rolebinding
subjects:
- kind: ServiceAccount
  name: myapp-sa
roleRef:
  kind: Role
  name: myapp-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      serviceAccountName: myapp-sa
      # ... rest of pod spec
```

**After (Akash SDL)**:
```yaml
services:
  myapp:
    image: myapp:latest
    params:
      permissions:
        read:
          - logs
```

The Akash provider handles all the ServiceAccount, Role, and RoleBinding creation automatically.

## Technical Implementation

When you deploy a service with permissions on Akash:

1. **SDL Parsing** - The Akash SDL parser reads the `permissions` block
2. **Manifest Generation** - Permissions are included in the deployment manifest
3. **Provider Processing** - The Akash provider receives the manifest with permissions
4. **Kubernetes Resource Creation** - The provider automatically creates:
   - A ServiceAccount for the deployment
   - A Role with appropriate permissions
   - A RoleBinding connecting the ServiceAccount to the Role
   - Pod specification using the created ServiceAccount
5. **Runtime Access** - Your application can now access the Kubernetes API with granted permissions

## FAQ

### Q: Do I need permissions to run a web server?
**A**: No. Standard web applications don't need any special permissions. Only add permissions if your application specifically needs to read logs or events.

### Q: Can my service access logs from other deployments?
**A**: No. Permissions are scoped to your own deployment's resources only.

### Q: What happens if I don't specify permissions?
**A**: Your service runs with the default Kubernetes service account, which has minimal permissions. This is the recommended and most secure option for most applications.

### Q: Can I request write permissions?
**A**: Currently, only read permissions are supported for logs and events. Write access to cluster resources is not available for security reasons.

### Q: Will this work with any container image?
**A**: Yes, but your application code needs to actually use the Kubernetes API to access logs or events. Just granting permissions doesn't automatically give your application this functionality - you need to implement it in your code.

### Q: How do I access the Kubernetes API from my application?
**A**: Use a Kubernetes client library for your programming language:
- Go: [client-go](https://github.com/kubernetes/client-go)
- Python: [kubernetes-client](https://github.com/kubernetes-client/python)
- JavaScript/Node.js: [@kubernetes/client-node](https://github.com/kubernetes-client/javascript)
- Java: [kubernetes-client](https://github.com/fabric8io/kubernetes-client)

The ServiceAccount token is automatically mounted at `/var/run/secrets/kubernetes.io/serviceaccount/token` in your container.

## Additional Resources

- [Kubernetes RBAC Documentation](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
- [Kubernetes Service Accounts](https://kubernetes.io/docs/concepts/security/service-accounts/)
- [Akash SDL Documentation](https://docs.akash.network/readme/stack-definition-language)
