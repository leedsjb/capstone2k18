#!/bin/bash

# Filename: build-deploy.sh
# Created: May 8, 2018
# Modified:
# Author: J. Benjamin Leeds
# License: None
# Purpose: This script builds and pushes test Go lang web server image to Docker Hub for deployment
# on a web server

docker stop leeds/crewjam-saml
docker rm leeds/crewjam-saml
docker rmi leeds/crewjam-saml
GOOS=linux go build
docker build -t leeds/crewjam-saml .
docker push leeds/crewjam-saml

# export environment=local-dev
# docker run -p 8080:80 -p 443:443 -e environment=local-docker-dev -v $(PWD)/tls:/etc/letsencrypt:ro leeds/crewjam-saml

# export environment=local-docker-dev
# docker run -n leeds/crewjam-saml -p 8080:80 -p 443:443 -e environment=local-docker-dev -v $(PWD)/tls:/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co:ro leeds/crewjam-saml

# Prod URL: crewjam-saml.test.elevate.emeloid.co

# LetsEncrypt TLS Certificate Generator

# lets encrypt command
# sudo letsencrypt certonly --standalone -d crewjam-saml.test.elevate.emeloid.co

# cert and privkey locations
#   /etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain.pem
#   /etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey.pem

# run docker container on web server

# docker stop leeds/crewjam-saml
# export environment=do
# docker run --name crewjam-saml -p 443:443 -e environment=do -v /etc/letsencrypt:/etc/letsencrypt:ro leeds/crewjam-saml