#!/usr/bin/env sh

docker build -t kvmayer/go-bff:latest .
docker push kvmayer/go-bff:latest
kubectl delete -f deploy.yaml
kubectl apply -f deploy.yaml