#!/bin/bash

# Specify the GitHub repository URL
github_url="wget https://raw.githubusercontent.com/techno-tim/launchpad/master/kubernetes/kube-prometheus-stack/values.yml"

# Create a values.yaml file
touch values.yaml

# Specify the output file
output_file="values.yaml"

# Use wget to download the raw content of the GitHub repository
wget $github_url -O $output_file

echo "Content from $github_url has been successfully copied to $output_file"