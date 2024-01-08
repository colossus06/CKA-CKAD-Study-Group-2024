#!/bin/bash
helm pull oci://registry-1.docker.io/elkakimmie/cka-ckad-sg --version 0.1.0
helm install voting-app oci://registry-1.docker.io/elkakimmie/cka-ckad-sg --version 0.1.0
./ns.sh
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack

sleep 20

# kubectl port-forward svc/$(kubectl get svc -n monitoring | grep grafana | awk '{print $1}') 52222:80
# kubectl describe po voting-app-pod | grep -i Node:
# kubectl describe svc voting-service | grep -i NodePort:
# kubectl describe po result-app-result-pod | grep -i Node:
# kubectl describe svc result-service | grep -i NodePort:
