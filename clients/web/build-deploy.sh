#!/bin/bash

# Filename: build-deploy.sh
# Created: May 8, 2018
# Modified: Thursday May 24, 2018
# Author: J. Benjamin Leeds
# License: None
# Purpose: This script builds and pushes test Golang web server image to Google Container Registry
#   for deployment on Kubernetes

usage="deploys Golang webserver to Kubernetes. Must provide an image version #.
usage: build-deploy.sh [v#]"

# echo $STR

if [ "$#" -ne 1 ]; then
    echo "$usage"    
    exit 1
fi

version=$1

yarn build

# docker stop crewjam-saml
# docker rm crewjam-saml
docker rmi us.gcr.io/airliftnw-uw/go-webclient:$version
GOOS=linux go build
docker build -t us.gcr.io/airliftnw-uw/go-webclient:$version .

# Stop here for local testing. 
# For Kubernetes deployment only:
go clean
docker push us.gcr.io/airliftnw-uw/go-webclient:$version
docker rmi us.gcr.io/airliftnw-uw/go-webclient:$version
kubectl apply -f ../../deployment/yaml/go-saml-web-server-deployment.yaml --record
kubectl rollout status deployment/go-saml-web-server-deployment

# export environment=local-dev
# docker run -p 8080:80 -p 4430:443 -e environment=local-docker-dev -v $(PWD)/tls:/etc/letsencrypt:ro leeds/crewjam-saml

# export environment=local-docker-dev
# docker run --name crewjam-saml -p 8080:80 -p 4430:443 -e environment=local-docker-dev \
# -v $(PWD)/tls:/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co:ro us.gcr.io/airliftnw-uw/go-webclient:0.20
