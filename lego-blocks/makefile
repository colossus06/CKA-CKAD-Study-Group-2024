# Makefile for setting up Prometheus and Grafana using Helm in Kubernetes
.PHONY: all

.PHONY: stop

stop: remove-admin-creds stop-monitoring-stack clean-credentials
# Define the default target, 'all', which depends on all other targets
all: create-ns get-nodes helm-version add-prometheus-repo create-admin-credentials create-grafana-secret describe-grafana-secret get-grafana-username get-grafana-password get-chart-values install-prometheus-grafana get-monitoring-pods

create-ns:
	./ns.sh

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
# create-monitoring-namespace:
# 	kubectl create namespace monitoring

create-admin-credentials:
	echo -n 'adminuser' > ./admin-user # change your username
	echo -n 'p@ssword!' > ./admin-password # change your password

# Create a Kubernetes secret for Grafana admin credentials
create-grafana-secret:
	kubectl create secret generic grafana-admin-credentials \
		--from-file=admin-user=./admin-user \
		--from-file=admin-password=./admin-password \
		-n monitoring


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


# Edit Prometheus Helm chart values
get-chart-values:
	bash github-content.sh

# Install Prometheus and Grafana using Helm
install-prometheus-grafana:
	helm install -n monitoring prometheus prometheus-community/kube-prometheus-stack -f values.yaml

# Display Prometheus and Grafana pods in the monitoring namespace
get-monitoring-pods:
	kubectl get pods -n monitoring

stop-monitoring-stack:
	helm uninstall -n monitoring prometheus
	echo "uninstalled prometheus, grafana ++++++++"


clean-credentials:
	# rm admin-user && rm admin-password
	kubectl delete secret grafana-admin-credentials prometheus-admission
	echo "secret, creds deleted++++++++"

remove-admin-creds:
	rm admin-password admin-user
