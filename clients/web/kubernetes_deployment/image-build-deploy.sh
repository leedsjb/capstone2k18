#!/bin/bash

# Filename: image-build-deploy.sh
# Created: 
# Modified: May 3, 2018
# Author: Tiffany Chen & J. Benjamin Leeds
# License: None
# Purpose: Build the docker image and deploy it to the google container registry
# Preconditions: run from kubernetes_deployment directory for Dockerfile filepath to work

docker stop web_container # stop any old running container
# docker rm web_container # remove any old container
docker rmi us.gcr.io/airliftnw-uw/webclient # delete existing image

yarn install # ensure all necessary dependencies are installed
yarn run build # build production version of react web app

docker build -t us.gcr.io/airliftnw-uw/webclient:0.8 ../ #move to dir containing Dockerfile

# upload container to google container registry (so Kubernetes Engine can download and run it)
# `gcloud docker` will not be supported for Docker client versions above 18.03.
# Please use `gcloud auth configure-docker` to configure `docker` to use `gcloud` as a credential
# helper, then use `docker` as you would for non-GCR registries, 
# e.g. `docker pull gcr.io/project-id/my-image`. Add `--verbosity=error` to silence this warning,
# e.g. `gcloud docker --verbosity=error -- pull gcr.io/project-id/my-image`.
# See: https://cloud.google.com/container-registry/docs/support/deprecation-notices#gcloud-docker

# gcloud docker -- push us.gcr.io/airliftnw-uw/webclient # note command deprecated in Docker 18.03+
docker push us.gcr.io/airliftnw-uw/webclient:0.8
kubectl rollout status deployment/alnw-deployment

# optional step - run container locally to test image
# docker run --rm -d -p 80:80 --name web_container us.gcr.io/airliftnw-uw/webclient

# curl http://localhost:8080