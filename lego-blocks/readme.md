
# Collection of the blocks

## Week I

6 node cluster

```sh
vagrant init
vagrant up
```

## Week II

Apply the pod manifest files running the following:

`./voting-app-automated.sh`

Get the node ips and pods:

`./voting-app-displayer.sh`

## Week III


**testing with an nginx deployment**


```sh
k create deployment nginx --image=nginx
kubectl expose deploy nginx --type=NodePort --port 80
k port-forward $(kubectl get po | grep nginx | awk '{print $1}') 8080:80
```




**testing the monitoring stack with the 6 node cluster**

Activate the monitoring stack running:

```sh
./monitoring-stack.sh
```

- **Username:** adminuser
- **Password:** p@ssword!

`kubectl port-forward $(kubectl get po -n monitoring | grep grafana | awk '{print $1}') 5200:3000`
`kubectl port-forward  pod/prometheus-prometheus-prometheus-0 5203:9090`


`k port-forward service/voting-service 30004:80 -n default`

## Week IV

Make sure you have helm installed.

```sh
helm list
helm repo add voting-app https://colossus06.github.io/cka-ckad-study-group-2024/
helm repo update
helm search repo voting-app
helm install [release-name] voting-app/voting-app 
```



## Week V

#todo
CICD