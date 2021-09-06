#!/bin/bash

kubectl apply -f deploy/crds.yaml
kubectl apply -f deploy/deployment.yaml
kubectl apply -f deploy/tenant-operator-deployment.yaml
