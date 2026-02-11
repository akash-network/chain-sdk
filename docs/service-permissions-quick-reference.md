# Kubernetes Default Service Account Permissions - Quick Reference

## Default Permissions Matrix

| Resource | Default ServiceAccount | With `logs` Permission | With `events` Permission | With Both |
|----------|----------------------|----------------------|------------------------|-----------|
| **Own Pod Logs** | ❌ No Access | ✅ Can Read | ❌ No Access | ✅ Can Read |
| **Other Pod Logs** | ❌ No Access | ❌ No Access | ❌ No Access | ❌ No Access |
| **Deployment Info** | ❌ No Access | ✅ Can Read | ❌ No Access | ✅ Can Read |
| **Events (Own)** | ❌ No Access | ❌ No Access | ✅ Can Read | ✅ Can Read |
| **Events (Other)** | ❌ No Access | ❌ No Access | ❌ No Access | ❌ No Access |
| **List Pods** | ❌ No Access | ✅ Can Read | ❌ No Access | ✅ Can Read |
| **Cluster Admin** | ❌ No Access | ❌ No Access | ❌ No Access | ❌ No Access |

## Key Points

### ❌ Default (No Permissions)
```yaml
services:
  web:
    image: nginx
    # No params.permissions = No access to cluster resources
```
**Result**: Pod can run but cannot read any cluster information

### ✅ With Logs Permission
```yaml
services:
  web:
    image: myapp
    params:
      permissions:
        read:
          - logs
```
**Result**: Pod can read its own logs and deployment info

### ✅ With Events Permission
```yaml
services:
  web:
    image: myapp
    params:
      permissions:
        read:
          - events
```
**Result**: Pod can read events related to its deployment

### ✅ With Both Permissions
```yaml
services:
  web:
    image: myapp
    params:
      permissions:
        read:
          - logs
          - events
```
**Result**: Pod can read both logs and events for its own resources

## Common Misconceptions

| ❌ Misconception | ✅ Reality |
|-----------------|-----------|
| "My pod can see its own logs by default" | Pods have NO access to read logs without explicit permissions |
| "I need permissions to write logs" | Writing to stdout/stderr (normal logging) doesn't need permissions. Reading logs does. |
| "permissions: logs gives access to specific K8s resources like 'pods/log'" | No, you specify the permission type ('logs' or 'events'), not K8s resources |
| "I can read logs from other deployments" | No, permissions are scoped to your own deployment only |
| "Events are just for troubleshooting" | Events can be used by apps for status tracking, deployment monitoring, etc. |

## Do You Need Permissions?

### 🟢 You DON'T need permissions if:
- Your app is a standard web server (nginx, apache, etc.)
- Your app is a database (postgres, mysql, etc.)
- Your app handles requests and returns responses
- Your app just needs to write logs (goes to stdout/stderr)
- Your app doesn't programmatically read logs or events

**Examples**: nginx, wordpress, postgres, redis, api server, static file server

### 🟡 You MIGHT need permissions if:
- Your app has an admin dashboard showing logs
- Your app monitors its own health using logs
- Your app tracks deployment status
- Your app does custom log analysis

**Examples**: monitoring dashboards, developer tools, log analyzers

### 🔴 Contact security team if:
- You want to read logs from other deployments
- You need write access to cluster resources
- You need cluster-admin level permissions
- You're not sure why you need permissions

## Code Example

### Application that needs to read logs:

```python
# Python example using kubernetes client
from kubernetes import client, config

# Load in-cluster config (uses ServiceAccount token)
config.load_incluster_config()

v1 = client.CoreV1Api()

# Read logs from own pod
pod_name = os.environ['HOSTNAME']  # Pod name
namespace = open('/var/run/secrets/kubernetes.io/serviceaccount/namespace').read()

logs = v1.read_namespaced_pod_log(
    name=pod_name,
    namespace=namespace
)

print(logs)
```

**Important**: This code will FAIL unless your SDL includes:
```yaml
params:
  permissions:
    read:
      - logs
```

## Testing

### How to verify your app needs permissions:

1. **Deploy without permissions first**
   ```yaml
   services:
     myapp:
       image: myapp:latest
       # No permissions
   ```

2. **Try to access logs/events from your app**
   - If you get permission errors → Add permissions
   - If it works → You don't need permissions

3. **Add only what you need**
   ```yaml
   services:
     myapp:
       image: myapp:latest
       params:
         permissions:
           read:
             - logs  # Only add if step 2 needed it
   ```

## Security Implications

### Low Risk ✅
- Reading your own pod's logs
- Reading events for your own deployment
- Using permissions only in development/staging

### Medium Risk ⚠️
- Reading logs in production (could expose sensitive data)
- Long-term storage of log/event data
- Sharing log data with external systems

### High Risk ❌
- Requesting permissions you don't use
- Giving all services permissions "just in case"
- Not documenting why permissions are needed

## Troubleshooting

### "Forbidden: pods is forbidden"
**Problem**: Your app is trying to list pods but doesn't have permission

**Solution**: Add `logs` permission:
```yaml
params:
  permissions:
    read:
      - logs
```

### "Forbidden: events is forbidden"
**Problem**: Your app is trying to read events but doesn't have permission

**Solution**: Add `events` permission:
```yaml
params:
  permissions:
    read:
      - events
```

### "User 'system:serviceaccount:default:default' cannot list resource"
**Problem**: Using default ServiceAccount which has no permissions

**Solution**: This is expected! Add permissions to your SDL if your app needs them.

## Further Reading

- [Full Service Permissions Guide](./service-permissions.md)
- [Service Permissions FAQ](./service-permissions-faq.md)
- [Kubernetes RBAC Documentation](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
