# Application Environment, Configuration and Security (25%)

* Discover and use resources that extend Kubernetes (CRD, Operators)
* Understand authentication, authorization and admission control
* Understand Requests, limits, quotas
* Understand ConfigMaps
* Define resource requirements
* Create & consume Secrets
* Understand ServiceAccounts
* Understand Application Security (SecurityContexts, Capabilities, etc.)

</br>

## Tasks
**Duration: 20 mins**

**1-)** Create a namespace named `tampa`, create a pod named `limited` using image bitnami/nginx that has a request of 25m CPU and 30 MiB of memory, and a limit of 300m CPU and 30MiB of memory.


<span style="color:green;">
<details closed>
  <summary>
  Answer
  </summary>

```bash
k create ns tampa
k run limited --image=bitnami/nginx -n tampa --dry-run=client -oyaml>pod.yaml
```

and add the resource requests and limits as shown 

```sh
    resources:
      requests:
        memory: "30Mi"
        cpu: "25m"
      limits:
        memory: "300Mi"
        cpu: "300m"
```


![](assets/20231204200812.png)

</details>
</span>

<br>


## **under constuction 🚧**
