#!/bin/bash
kubectl apply -f deploy/cluster-roles.yaml
kubectl apply -f deploy/crds.yaml
kubectl apply -f deploy/deployment.yaml
kubectl apply -f deploy/tenant_operator_deployment.yaml
