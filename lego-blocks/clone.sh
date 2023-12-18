#!/bin/bash
git clone https://github.com/colossus06/cka-ckad-study-group-2024.git
cd cka-ckad-study-group-2024/lego-blocks/
chmod +x ./
./ns.sh
./monitoring-stack.sh
sleep 8
./voting-app-automated.sh