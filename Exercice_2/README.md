# Exercice 2

```shell
docker volume create docker-certs-ca
docker volume create docker-certs-client

docker network create dind_network

docker run --privileged --name docker-daemon -d --network dind_network --network-alias docker -e DOCKER_TLS_CERTDIR=/certs -v docker-certs-ca:/certs/ca -v docker-certs-client:/certs/client docker:dind

docker run --rm --network dind_network -e DOCKER_TLS_CERTDIR=/certs -e DOCKER_HOST=tcp://docker:2376 -v docker-certs-client:/certs/client:ro docker:latest version

docker run -it --rm --network dind_network -e DOCKER_TLS_CERTDIR=/certs -e DOCKER_HOST=tcp://docker:2376 -v docker-certs-client:/certs/client:ro docker:latest sh

docker run --name db -d -e POSTGRES_DB=tprevision -e POSTGRES_PASSWORD=password  postgres:latest
```