FROM mysql:latest 
COPY setup.sql /docker-entrypoint-initdb.d/ 

#docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=colo -e MYSQL_DATABASE=ckad_crud -e MYSQL_USER=colo -e MYSQL_PASSWORD=colo -v mysql-data:/var/lib/mysql -v /setup.sql:/docker-entrypoint-initdb.d/setup.sql -d my-custom-mysql:latest
# export MYSQL_ROOT_PASSWORD=colo
# export MYSQL_DATABASE=ckad_crud
# export MYSQL_USER=colo
# export MYSQL_PASSWORD=colo
# echo $MYSQL_ROOT_PASSWORD, $MYSQL_DATABASE, $MYSQL_USER, $MYSQL_PASSWORD

# volume will be on /var/lib/docker/volumes

# docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -e MYSQL_DATABASE=$MYSQL_DATABASE -e MYSQL_USER=$MYSQL_USER -e MYSQL_PASSWORD=$MYSQL_PASSWORD -v mysql-data:/var/lib/mysql -v ./setup.sql:/docker-entrypoint-initdb.d/setup.sql -d my-custom-mysql:latest
