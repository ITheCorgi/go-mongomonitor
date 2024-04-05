# Mongo Monitoring #

## Getting started
### Connection pool monitor:
There is following metrics:
- mongo_client_connections_in_pool_count
- mongo_client_connection_usage_in_percent

To add alert to check connection pool usage in grafana add following:
```yaml
groups:
- name: connection_pool_usage
  rules:
    - alert: ConnectionPoolHighUsage
      expr: <namespacename_mongo_>client_connection_usage_in_percent > 90
      for: 30s
      severity: warning
      summary: "Connection pool high usage"
      description: "The percentage of connections in use in the MongoDB client connection pool is approaching its limit."```