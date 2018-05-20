#!/bin/bash

# Filename: build-deploy.sh
# Created: May 8, 2018
# Modified: Saturday May 19, 2018
# Author: J. Benjamin Leeds
# License: None
# Purpose: This script builds and pushes test Golang web server image to Google Container Registry
#   for deployment on Kubernetes

docker stop crewjam-saml
docker rm crewjam-saml
docker rmi us.gcr.io/airliftnw-uw/go-webclient:0.21
GOOS=linux go build
docker build -t us.gcr.io/airliftnw-uw/go-webclient:0.21 .

# Stop here for local testing. 
# For Kubernetes deployment only:
go clean
docker push us.gcr.io/airliftnw-uw/go-webclient:0.21
docker rmi us.gcr.io/airliftnw-uw/go-webclient:0.21
kubectl apply -f ../../clients/web/kubernetes_deployment/yaml/crewjam-deployment.yaml --record
kubectl rollout status deployment/go-saml-web-server-deployment

# export environment=local-dev
# docker run -p 8080:80 -p 4430:443 -e environment=local-docker-dev -v $(PWD)/tls:/etc/letsencrypt:ro leeds/crewjam-saml

# export environment=local-docker-dev
# docker run --name crewjam-saml -p 8080:80 -p 4430:443 -e environment=local-docker-dev \
# -v $(PWD)/tls:/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co:ro us.gcr.io/airliftnw-uw/go-webclient:0.20

# LetsEncrypt TLS Certificate Generator

# lets encrypt command
# sudo letsencrypt certonly --standalone -d crewjam-saml.test.elevate.emeloid.co

# cert and privkey locations
#   /etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain.pem
#   /etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey.pem



