#!/bin/bash
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/voting_app.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/svc-app.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/redis.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/svc-redis.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/db.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/svc-db.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/worker.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/voting_app_result.yaml
kubectl apply -f ../week-II-distributed-app/voting-app/k8s-manifests/svc-result.yaml