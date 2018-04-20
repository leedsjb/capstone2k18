#!/usr/bin/env bash
./build.sh
docker push vincentmvdm/messages-api

export TLSCERT=/etc/letsencrypt/live/api.messages.vincentmvdm.com/fullchain.pem 
export TLSKEY=/etc/letsencrypt/live/api.messages.vincentmvdm.com/privkey.pem 

export DBADDR=mongosvr:27017
export REDISADDR=redissvr:6379
export MESSAGESSVCADDR=messagessvc:80
export SUMMARYSVCADDR=summarysvc:80
export RABBITADDR=rabbitsvr:5672
export SESSIONKEY=hancmarginisexiguitasnoncaperet

ssh -oStrictHostKeyChecking=no root@165.227.17.79 'bash -s' << EOF
docker network create messages

docker rm -f mongosvr
docker rm -f redissvr
docker rm -f messagessvc
docker rm -f summarysvc
docker rm -f rabbitsvr
docker rm -f messages-api

docker pull vincentmvdm/messages-microservice
docker pull vincentmvdm/summary-microservice
docker pull vincentmvdm/messages-api

docker run -d --name mongosvr  \
--network messages \
mongo

docker run -d --name redissvr \
--network messages \
redis

docker run -d --hostname rabbit --name rabbitsvr \
--network messages \
rabbitmq:3

docker run -d --name messagessvc \
--network messages \
-e DBADDR=$DBADDR \
-e RABBITADDR=$RABBITADDR \
vincentmvdm/messages-microservice

docker run -d --name summarysvc \
--network messages \
vincentmvdm/summary-microservice

docker run -d --name messages-api \
--network messages \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e DBADDR=$DBADDR \
-e REDISADDR=$REDISADDR \
-e RABBITADDR=$RABBITADDR \
-e SESSIONKEY=$SESSIONKEY \
-e MESSAGESSVCADDR=$MESSAGESSVCADDR \
-e SUMMARYSVCADDR=$SUMMARYSVCADDR \
vincentmvdm/messages-api
exit
EOF

