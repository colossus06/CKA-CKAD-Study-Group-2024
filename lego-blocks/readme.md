
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

Activate the monitoring stack running:

```sh
./monitoring-stack.sh
```

- **Username:** adminuser
- **Password:** p@ssword!

`kubectl port-forward $(kubectl get po -n monitoring | grep grafana | awk '{print $1}') 5200:3000``

## Week IV

Make sure you have helm installed.

```sh
helm list
helm repo add voting-app https://colossus06.github.io/cka-ckad-study-group-2024/
helm repo update
helm search repo voting-app
helm install [release-name] voting-app/voting-app 
```


k port-forward service/voting-service 30004:30004


## Week V

#todo
CICD