#!/usr/bin/env bash
./devBuild.sh
docker push vincentmvdm/dev-messages-api

export TLSCERT=/tls/fullchain.pem \
export TLSKEY=/tls/privkey.pem \

export DBADDR=dev-mongosvr:27017
export REDISADDR=dev-redissvr:6379
export MESSAGESSVCADDR=dev-messagessvc:80
export SUMMARYSVCADDR=dev-summarysvc:80
export RABBITADDR=dev-rabbitsvr:5672
export SESSIONKEY=hancmarginisexiguitasnoncaperet

docker network create dev-messages

docker rm -f dev-mongosvr
docker rm -f dev-redissvr
docker rm -f dev-messagessvc
docker rm -f dev-summarysvc
docker rm -f dev-rabbitsvr
docker rm -f dev-messages-api

docker pull vincentmvdm/dev-messages-microservice
docker pull vincentmvdm/dev-summary-microservice
docker pull vincentmvdm/dev-messages-api

docker run -d --name dev-mongosvr  \
--network dev-messages \
mongo

docker run -d --name dev-redissvr \
--network dev-messages \
redis 

docker run -d --hostname rabbit --name dev-rabbitsvr \
--network dev-messages \
rabbitmq:3

sleep 10

docker run -d --name dev-messagessvc \
--network dev-messages \
-e DBADDR=$DBADDR \
-e RABBITADDR=$RABBITADDR \
vincentmvdm/dev-messages-microservice

docker run -d --name dev-summarysvc \
--network dev-messages \
vincentmvdm/dev-summary-microservice

docker run -d --name dev-messages-api \
--network dev-messages \
-p 443:443 \
-v $(pwd)/tls:/tls:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e DBADDR=$DBADDR \
-e REDISADDR=$REDISADDR \
-e RABBITADDR=$RABBITADDR \
-e SESSIONKEY=$SESSIONKEY \
-e MESSAGESSVCADDR=$MESSAGESSVCADDR \
-e SUMMARYSVCADDR=$SUMMARYSVCADDR \
vincentmvdm/dev-messages-api