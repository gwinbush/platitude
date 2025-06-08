#!/bin/bash
set -e

echo "Building microservices with Bazel..."


# Deploy to Kubernetes
echo "Deploying to Kubernetes..."
kubectl apply -f k8s/python-api.yaml
kubectl apply -f k8s/go-api.yaml

# Wait for deployments
echo "Waiting for deployments..."
kubectl rollout status deployment/python-api
kubectl rollout status deployment/go-api

echo "Deployment complete!"
echo ""
echo "Services:"
kubectl get services
echo ""
echo "Pods:"
kubectl get pods
echo ""

# Print port forwarding information
echo "Access your services:"
echo "Python API: curl http://localhost:8081/python-api/platitude"
echo "Go API: curl http://localhost:8082/go-api/platitude"
echo ""

# Start port forwarding
kubectl port-forward svc/python-api 8081:8081 &
kubectl port-forward svc/go-api 8082:8082
