#!/bin/bash 

#### Run the docker compose file (Jenkins) #### 
echo " Running Jenkins Container "
docker compose up -d

echo "--------------------------"
#### List the jenkins container ####
echo "Getting the docker container"
echo "--------------------------"

docker ps | grep jenkins_master

echo "--------------------------"
# Get the container ID of the Jenkins container
echo "Getting the Container ID and exec in the contanier"
echo "--------------------------------------------------"
container_id=$(docker ps -qf "name=jenkins_master")

# Check if the container ID is not empty
if [ -n "$container_id" ]; then
    # Execute bash inside the Jenkins container as root
    docker exec -it --user root "$container_id" bash -c '
      curl https://get.docker.com/ > dockerinstall && \
      chmod 777 dockerinstall && \
      ./dockerinstall
      exit
    '
else
    echo "Jenkins container not found. Make sure it is running."
fi

echo "--------------------------"
### Run the docker sock in the ubuntu ####
echo " Running the docker socket"
echo "--------------------------"
sudo chmod 666 /var/run/docker.sock

echo "--------------------------"
#### Get the password #### 
echo " Getting the password"
echo "--------------------------"
container_id=$(docker ps -qf "name=jenkins_master")



# Check if the container ID is not empty
if [ -n "$container_id" ]; then
    # Execute bash inside the Jenkins container as root
    docker exec -it "$container_id" bash -c '
      cat var/jenkins_home/secrets/initialAdminPassword
      exit
    '
else
    echo "Jenkins container not found. Make sure it is running."
fi 

echo "--------------------------"
### Getting the url and ip 
echo "Please head over to http://localhost:8080 to Complete Jenkins Setup"