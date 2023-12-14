# Application Design and Build


The first domain in the exam is Application Design and Build having a 20% weight. Before moving on to the next domain make sure that you're fluent with the following. 


- [x] storage and definition in pods and multiple containers
- [x] batch, periodic workloads

## Achivements

We aim to achieve some operations on pod object such as the following:

* creating/running (`k create <name>`, `k run <name>`, `k apply -f`)
* creating objects from a recursive dir (`k apply -f <folder>/ -R`)
* switching between contexts/set context (`k config set-context`)
* editing/updating (`k edit pod`,`k patch pod`, `k apply -f`)
* deleting (`k delete pod`)
* listing available api versions (`k api-versions`)
* short names for some of the resource(`k api-resources`)
* getting info about an object(`k describe po ckad | grep -i <search-term>`)
* finding info about a command(`k run --help`)
* finding info about the fields of a resource(`k explain po.kind`, `k explain po.status`)
* starting a temporary pod(add --rm and -it)


![](assets/assets/20231211171458.png)


Let's get started with the core skills.

### Set the context and the namespace.

A "Context" is a combination of a cluster, user, and namespace. It is a way to specify the cluster you want to interact with, the user or authentication credentials you want to use, and the default namespace for that user. Setting a context is important because it helps you manage multiple Kubernetes clusters and switch between them easily.

**Use Cases:**  Multi-cluster Management, User and Authentication, isolate and organize resources, simplifies command-line operations.

**Hint:** Don't change the namespace, change the context as per the question. If there is a namespace add -n `namespace` to lessen the risk of error. Set context before each question using the kubectl config command, and switch between contexts using kubectl config use-context. 


## Tasks

**Duration: 3 mins/each**

**1-)** You have only one namespace called `tossi`.  Create a new namespace called `semvar`. 

Set the context to 'ckad', set the user field on the ckad context to admin, specify the newly created namespace simultaneously. Switch to context 'tossi'. 

<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash
k config set-context ckad --ns=semvar --user=admin

k config use-context tossi
```

</details>
</span>

<br>

**2-)** Find out the short name for namespaces

<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash

k api-resources | grep -i namespace
```

![Alt text](assets/image.png)
</details>
</span>

<br>

**3-)** Create a pod named `sleepy-c` using the `busybox` image, make the container sleep for `3600` seconds and display real-time updates in your terminal.
Once you see the pod status running, delete the pod object immediately.

<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash

k run sleepy-c --image=busybox --command \
-- sleep 3600; \
k get po --watch

k delete po sleepy-c --grace-period=0 --force 
#or
k delete po sleepy-c --now
```
![Alt text](assets/image-2.png)

![Alt text](assets/image-3.png)

</details>
</span>

<br>

> **Note**
> Graceful delete may take up to 30 seconds. Use it when you realize you've made a mistake and need a short window to promptly correct your actions.
`--now` is same as `--grace-period=1`,  you can use it for immediate shutdown.


**4-)** Autocompletion and alias `k` ias already preconfigured for the CKAD. 

Let's try something different. Enable autocompletion for kubectl commands in your Bash shell permanently.



<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash

echo "source <(kubectl completion bash)" >> ~/.bashrc

echo "source <(complete -o default -F __start_kubectl k" >> ~/.bashrc

k completion -h # get more info on auto completion

alias k="kubetcl" >> ~/.bashrc

# let's say you made a mistake and added a trailind space
# delete multiple lines from .bashrc on cli

sed -i -e '/source < (kubectl completion bash)/d' \
-e '/source < (k completion bash)/d' ~/.bashrc

# or you can replace in the place

sed -i -e 's/source < (kubectl completion bash)/source <(kubectl completion bash)/g' \
-e 's/source < (k completion bash)/source <(k completion bash)/g' ~/.bashrc


```
![Alt text](assets/image-4.png)

</details>
</span>

<br>


**5-)** Preview the changes that would occur if you run a pod named `sleepy-t` using the `busybox` image, making it sleep for 120/s.  Don't make any server-side changes.

Create an environment variable in a bash shell that can achieve the upper result and persist across different shell sessions.

Nope, don't create a manifest file.

<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash

--dry-run: Preview with 👇
=client: no real changes to the actual resources

# get persistence
echo 'export do="--dry-run=client -o yaml"' >> ~/.bashrc
source ~/.bashrc

# verification
k run sleepy-t --image=busybox $do

```

![Alt text](assets/image-5.png)

</details>
</span>

<br>






**6-)**
Create a Pod named ckad running image ngnix in a container, in the default namespace, and expose port 80, watch for changes.

<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash
k run ckad --image=nginx --port=80
k get po -w
```

![](assets/assets/20231211165852.png)

![](assets/assets/20231211165929.png)


</details>
</span>

<br>


**7-)**
Create a pod named ckad with the ubuntu/apache2 image, a restart policy of Never, expose port 80. Generate the YAML representation of the pod definition with yaml manifestation. Expose the ckad pod using service type of Nodeport accessible on 8080. 

<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash
kubectl run ckad --image=ubuntu/apache2 --restart=Never --port=80 --dry-run=client -oyaml> apache.yaml
kubectl expose pod ckad --type=NodePort --port=80 --target-port=8080 --name=ckad-svc
```


![](assets/assets/20231211165852.png)

![](assets/assets/20231211165929.png)


</details>
</span>

<br>

**8-)**
Create a namespace called `mynamespace` , create an nginx pod named mynginx, name the container to mynginx-container in the mynamespace ns, expose port 80 declaratively.
Get the IP address of the Pod, create an alias for a temporary pod named temp. Temp pod shoould connect to the pod on the exposed port. Download the content to the standard output.

<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash
touch mynamespace.yaml
vim mynamespace.yaml
k apply -f mynamespace.yaml
touch mynginx.yaml
k apply -f mynginx.yaml

k config set-context --current --namespace=mynamespace
export alias temp="kubectl run temp --image=busybox --restart=Never --rm -it"
$temp -- wget -o- <pod-ip>:<exposed-port>
k logs mygninx

```
![](assets/20231212172825.png)

after setting my namespace:

![](assets/20231212172933.png)

![](assets/20231212165841.png)


get the logs:

![](assets/20231212173056.png)


</details>
</span>

<br>




