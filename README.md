# 🛒 Analytics Ecommerce Dashboard Analytics 

A real-time Analytics Dashboard for an E-commerce store

## Functional Requirements

- CreateTransaction: Ability to simulate a new transaction.

- GetTransaction: Ability to fetch a specific transaction based on its id.

- StreamTransactions: Ability to stream transactions in real-time.

- GetTotalSales: Ability to get the current total sales.

- GetSalesByProduct: Ability to get the current sales by product.

- GetTopCustomers: Ability to get the top 5 customers based on sales.


## How to Run

- `minikube start`


### DB Migrations
- `kubectl apply -f pkgs/database/deploy/deployment.yaml`


- `migrate -path=db/migrations -database 'cockroachdb://root:@localhost:26257/defaultdb?sslmode=disable' up`


- `bazel run //pkgs/customers/server:image`


- `minikube image load bazel/pkgs/customers/server:image`


- `kubectl apply -f pkgs/customers/deploy/deployment.yaml`


- `kubectl apply -f pkgs/customers/deploy/deployment.yaml`


- `kubectl port-forward <pod_id> hostPort:containerPort`

## k8s General Commands

### show logs for a pod
```kubectl logs -f <pod_id>```

### port forward
```kubectl port-forward <pod_id> HostPort:containerPort```

### Deployment
```kubectl apply -f <deployment_yaml_file_path>```

### Service
```kubectl apply -f <service_yaml_file_path>```