#!/bin/bash
# Use this script to rebuild the local NGINX webclient image

docker stop alnwnginx && docker rm alnwnginx

docker rmi local/airliftnw-uw/webclient

docker build -t local/airliftnw-uw/webclient ../ #move to dir containing Dockerfile

docker run --name alnwnginx -d -p 8080:80 local/airliftnw-uw/webclient
