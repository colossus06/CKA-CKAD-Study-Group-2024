#!/bin/bash
kubectl describe po voting-app-pod | grep -i Node:
kubectl describe svc voting-service | grep -i NodePort:
kubectl describe po result-app-result-pod | grep -i Node:
kubectl describe svc result-service | grep -i NodePort:
