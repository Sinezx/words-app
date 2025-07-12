#!/bin/bash

docker container stop words-server
docker container rm words-server
docker image rm words-server

# build dsn and replace config.json file's default value
host=`docker inspect postgres | grep '"IPAddress":' | sed -n '1p' | sed 's/[":,]'//g | awk '{print $2}'`
user="postgres"
password="postgres"
dbname="db"
port="5432"
dsn="host=${host} user=${user} password=${password} dbname=${dbname} port=${port} sslmode=disable TimeZone=Asia\/Shanghai"
sed -i "s/host=192.168.2.1 user=postgres password=postgres dbname=db port=5432 sslmode=disable TimeZone=Asia\/Shanghai/${dsn}/g" config.json

# build image
docker buildx build --allow=network.host -t words-server .

# run container
docker run -d -p 8080:8080 --name=words-server words-server