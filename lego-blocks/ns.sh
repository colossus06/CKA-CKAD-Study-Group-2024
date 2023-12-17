#!/bin/bash
kubectl create ns monitoring
kubectl config set-context --current --namespace=monitoring
echo "Namespace created"