## Setup vim and 

Task 1 - persist Vim settings in .vimrc
set et ts=2 sw=2 nu
syntax on



Task2- How many worker nodes do you have? create /root/test file on node03 without interacting.

```
NAME           STATUS   ROLES           AGE    VERSION
controlplane   Ready    control-plane   7d5h   v1.28.1
node03         Ready    <none>          7d4h   v1.28.1
```
```
ssh node03 'touch /root/test'
```

Setup alias some aliases to save time and 


alias ka='kubectl apply -f'
alias kns='kubectl config set-context --current --namespace'
export do='--dry-run=client -o yaml'
export kill='--grace-period=0 --force'
This will provide you with a temporary nginx pod with a shell where you can test connectivity across pods, services, and ingress
export tester='kubectl run nginx-tester --image=nginx:alpine --restart=Never --rm -it sh'


kubectl create deploy nginx-server --image=nginx --replicas=2 --dry-run=client -o yaml > deploy.yaml



tip: kubectl is already set up as an alias 'k' in the exam environment so you don't have to type kubectl every time. Kubernetes bash auto-completion is already preconfigured in the environment as well.



Using the correct context and namespaces during the exam:

Modify kubeconfig files using subcommands like "kubectl config set current-context my-context