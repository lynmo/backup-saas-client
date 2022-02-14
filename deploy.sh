#!/bin/bash

kubectl apply -f deploy/crds.yaml
kubectl apply -f deploy/backend-controller-rbac.yaml
kubectl apply -f deploy/deployment.yaml
kubectl apply -f deploy/tenant-operator-deployment.yaml
# kubectl apply -f deploy/tenant-manager-cr.yaml
