# domains as of Kubernetes v1.29 - Dec 2023

## Application Design and Build (20%)
[Define, build and modify container images](https://kubernetes.io/docs/concepts/containers/images/)
[Choose and use the right workload resource (Deployment, DaemonSet, CronJob, etc.)](https://kubernetes.io/docs/concepts/workloads/_print/)
[Understand multi-container Pod design patterns (e.g. sidecar, init and others)](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/)
[Utilize persistent and ephemeral volumes](https://kubernetes.io/docs/concepts/storage/)

## Application Deployment (20%)
[Use Kubernetes primitives to implement common deployment strategies (e.g. blue/green or canary)](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
[Understand Deployments and how to perform rolling updates](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
Use the Helm package manager to deploy existing packages
[Kustomize](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/)

## Application Observability and Maintenance (15%)
[Understand API deprecations](https://kubernetes.io/docs/reference/using-api/deprecation-policy/)
[Implement probes and health checks](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)
Use built-in CLI tools to monitor Kubernetes applications
Utilize container logs
Debugging in Kubernetes

## Application Environment, Configuration and Security (25%)
Discover and use resources that extend Kubernetes (CRD, Operators)
Understand authentication, authorization and admission control
Understand Requests, limits, quotas
Understand ConfigMaps
Define resource requirements
Create & consume Secrets
Understand ServiceAccounts
Understand Application Security (SecurityContexts, Capabilities, etc.)

## Services and Networking (20%)
Demonstrate basic understanding of NetworkPolicies
Provide and troubleshoot access to applications via services
Use Ingress rules to expose applications