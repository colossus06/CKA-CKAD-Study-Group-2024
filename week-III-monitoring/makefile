# Makefile for setting up Prometheus and Grafana using Helm in Kubernetes

# Download Helm installation script
download-helm:
	curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
	chmod 700 get_helm.sh
	./get_helm.sh

# Display Kubernetes nodes
get-nodes:
	kubectl get nodes

# Display Helm version
helm-version:
	helm version

# Add Prometheus Helm chart repository
add-prometheus-repo:
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update

# Create a Kubernetes namespace for monitoring
create-monitoring-namespace:
	kubectl create namespace monitoring

# Create admin user and password files
create-admin-credentials:
	echo -n 'adminuser' > ./admin-user # change your username
	echo -n 'p@ssword!' > ./admin-password # change your password

# Create a Kubernetes secret for Grafana admin credentials
create-grafana-secret:
	kubectl create secret generic grafana-admin-credentials --from-file=./admin-user --from-file=admin-password -n monitoring

# Display details of Grafana admin credentials secret
describe-grafana-secret:
	kubectl describe secret -n monitoring grafana-admin-credentials

# Decode and display Grafana admin username from the secret
get-grafana-username:
	kubectl get secret -n monitoring grafana-admin-credentials -o jsonpath="{.data.admin-user}" | base64 --decode

# Decode and display Grafana admin password from the secret
get-grafana-password:
	kubectl get secret -n monitoring grafana-admin-credentials -o jsonpath="{.data.admin-password}" | base64 --decode

# Remove admin user and password files
clean-credentials:
	rm admin-user && rm admin-password

# Edit Prometheus Helm chart values
get-chart-values:
	bash github-content.sh

# Install Prometheus and Grafana using Helm
install-prometheus-grafana:
	helm install -n monitoring prometheus prometheus-community/kube-prometheus-stack -f values.yaml

# Display Prometheus and Grafana pods in the monitoring namespace
get-monitoring-pods:
	kubectl get pods -n monitoring
