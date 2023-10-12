# 🛒 Analytics Ecommerce Dashboard Analytics 

A real-time Analytics Dashboard for an E-commerce store

## DB Migrations
```migrate -path=db/migrations -database 'cockroachdb://root:@localhost:26257/defaultdb?sslmode=disable' up```

## k8s General Commands

### show logs for a pod
```kubectl logs -f <pod_id>```

### port forward
```kubectl port-forward <pod_id> HostPort:containerPort```

### Deployment
```kubectl apply -f <deployment_yaml_file_path>```

### Service
```kubectl apply -f <service_yaml_file_path>```