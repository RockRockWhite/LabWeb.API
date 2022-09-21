docker rm -f $(docker ps -a)
docker rmi -f $(docker images)
docker volume remove $(docker volume list)