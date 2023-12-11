# Monitoring Stack Setup

This guide outlines the steps to set up a monitoring stack using bash scripts and Kubernetes. The monitoring stack includes Prometheus for metrics collection and Grafana for visualization.

## Prerequisites

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- A running Kubernetes cluster

## Installation

1. Clone the monitoring stack repository:

    ```bash
    git clone https://github.com/your-repo/monitoring-stack.git
    cd monitoring-stack
    ```

2. Run the installation script:

    ```bash
    bash install-monitoring-stack.sh
    ```

3. Access Grafana locally using port-forwarding:

    ```bash
    kubectl port-forward -n monitoring <grafana-pod-name> 52222:3000
    ```

    Replace `<grafana-pod-name>` with the actual name of your Grafana pod.

4. Open your browser and navigate to [http://localhost:52222](http://localhost:52222).

## Grafana Dashboard Access

To access the Grafana dashboard, use the following credentials:

- **Username:** admin
- **Password:** p@ssword!

Update the credentials in the `makefile` if needed.

## Stopping the Services

To stop the monitoring stack services, run the following command:

```bash
bash stop-monitoring-stack.sh
