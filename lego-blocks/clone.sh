#!/bin/bash
# git clone https://github.com/colossus06/cka-ckad-study-group-2024.git
# cd cka-ckad-study-group-2024/lego-blocks/
# chmod +x ./
kubectl get nodes -owide
kubectl taint nodes controlplane node-role.kubernetes.io/master-
echo "untainted master node"
./ns.sh
./monitoring-stack.sh
sleep 8
./voting-app-automated.sh