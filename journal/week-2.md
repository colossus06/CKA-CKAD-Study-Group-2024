# CKAD Study Journal

## Week II

This week i finally started studying the CKAD curriculum's first domain and deployed 2 sample multi-tier apps using

* docker compose
* minikube
* kubeadm

## Adding 3 more vms to the week I cluster

Had difficulty in adding 3 more vms to the existing vagrant configuration. Kept getting the following error:

`INFO warden: Calling OUT action: #<Proc:0x000001a01bb591b0 C:/Program Files/Vagrant/embedded/gems/gems/vagrant-2.4.0/lib/vagrant/action/warden.rb:116 (lambda)>
 INFO warden: Calling OUT action: #<Vagrant::Action::Builtin::Provision:0x000001a01b21a8d8>`

![](assets/20231208041444.png)

## Fixing the ip of the kworker3

Got the Host key verification failed. error, fixed it by clearing the .ssh/known_hosts file

![](assets/20231208110945.png)


![](assets/20231208111100.png)


My kworker3 had a different ip than the one in the /ect/hosts file

`Environment="KUBELET_EXTRA_ARGS=--node-ip=172.16.16.103"`

`k port-forward service/voting-service 80:voting-service
kubectl port-forward pod/voting-app-pod 80:30004`

![](assets/20231208121127.png)


Added the node ip and run the following:

```
systemctl daemon-reload
systemctl restart kubelet
```

![](assets/20231208121208.png)



Voila! this solved the issue :)


![](assets/20231208121306.png)


Oh, there comes another issue:

`connection refused IPv6 dial tcp6: address localhost: no suitable address found`

![](assets/20231208121557.png)

Let's try displaying using the node ip before diving in debugging the issue. Time is precious right? ⌛

After identifying which node my pod is running on:

![](assets/20231208121842.png)

The nodejs app is also running on kworker1:

![](assets/20231208122013.png)

## Key Learnings

* Remember that you don't necessarily need to use `kubectl port-forward` if you are using a `NodePort` service, as NodePort exposes the service on a static port on each node in the cluster.

* Vagrant snapshots helps me manage my time efficiently. 
* Working with time constraints is a great challenge, and it often requires prioritization, efficient and focused planning. Just a simple sticky note has proven to be more efficient than a great dotted notebook.