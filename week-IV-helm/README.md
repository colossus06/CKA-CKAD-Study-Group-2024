# Converting the app to deployment and packaging as Helm Chart

Create the deployment manifests

```sh
kubectl create ns monitoring
kubectl apply -f ./k8s-manifests/deployment/db-cm.yaml
kubectl apply -f ./k8s-manifests/deployment/db-secret.yaml
kubectl apply -f ./k8s-manifests/deployment/pv.yaml
kubectl apply -f ./k8s-manifests/deployment/voting-app.yaml
kubectl apply -f ./k8s-manifests/deployment/svc-app.yaml
kubectl apply -f ./k8s-manifests/deployment/redis.yaml
kubectl apply -f ./k8s-manifests/deployment/svc-redis.yaml
kubectl apply -f ./k8s-manifests/deployment/db.yaml
kubectl apply -f ./k8s-manifests/deployment/svc-db.yaml
kubectl apply -f ./k8s-manifests/deployment/worker-app.yaml
kubectl apply -f ./k8s-manifests/deployment/result.yaml
kubectl apply -f ./k8s-manifests/deployment/svc-result.yaml
minikube service voting-service --url -n monitoring
minikube service result-service --url -n monitoring
```


![](assets/20231220004210.png)

![](assets/20231220004236.png)



**Running the Voting app Locally**

Make sure you have helm installed.

```sh
helm list
helm repo add voting-app https://colossus06.github.io/cka-ckad-study-group-2024/
helm repo update
helm search repo voting-app
helm install voting-app voting-app/voting-app-sg 
```

