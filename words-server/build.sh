#!/bin/bash

docker container stop words-server
docker container rm words-server
docker image rm words-server

docker buildx build --allow=network.host -t words-server .

docker run -d -p 8080:8080 --name=words-server words-server