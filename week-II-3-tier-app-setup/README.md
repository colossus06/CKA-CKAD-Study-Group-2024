# 3 Tier App Deployment on Kubernetes


## Mysql Setup

**Step-1** Create configmap and secret

`k create cm config-name --from-literal=MYSQL_DATABASE=db-name`

`k create secret generic secret-name --from-literal=MYSQL_ROOT_PASSWORD=rootpassword`

**Step-2** Inject the configmap and secret into the mysql pod manifest file.


**Step-3** Apply the mysql-pod.yaml file running `k apply -f mysql-pod.yaml`


**Step-4** Expose the mysql pod as a Cluster service


## phpmyadmin Setup


**Step-1** Create configmap and secret

`k create cm config-name --from-literal=PMA_HOST=<mysql-server-address>, --from-literal=PMA_PORT=<mysql-svc-port>`

`k create secret generic secret-name --from-literal=PMA_USER=<user>, --from-literal=PMA_PASSWORD=<root-password>`

**Step-2** Inject the configmap and secret into phphpadmin-pod manifest file.


**Step-3** Apply the phpadmin-pod.yaml file running `k apply -f mysql-pod.yaml`


**Step-4** Expose the phpmyadmin pod as a Nodeport service


Congratulations, now php myadmin and mysql is connected.

Now, navigate to the  phpadmin service port and import the dummy data named `simple_todo.sql`


![](./assets/image%201.png)

## App server Setup


**Step-1** Create a pod for the frontend php app running the following:

`k run php-app --image=elkakimmie/phpwebapp:latest`

**Step-2** Expose the php-app pod as a Nodeport service.


Congratulations!

You've succesfully deployed a 3 tier app on your kubernetes cluster. :tada:


![](./assets/image%202.png)

### Tested Environments

* minikube

![](./assets/image%203.png)