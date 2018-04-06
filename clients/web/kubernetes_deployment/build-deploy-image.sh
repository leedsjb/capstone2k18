#!/bin/bash
# Use this script to build the docker image and deploy it to the google container registry
# Run from kubernetes_deployment folder for Dockerfile filepath to work

docker stop web_container # stop any old running container
# docker rm web_container # remove any old container
docker rmi us.gcr.io/airliftnw-uw/webclient # delete existing image

yarn install # ensure all necessary dependencies are installed
yarn run build # build production version of react web app

docker build -t us.gcr.io/airliftnw-uw/webclient ../ #move to dir containing Dockerfile

# upload conatiner to google container registry (so Kubernetes Engine can download and run it)
gcloud docker -- push us.gcr.io/airliftnw-uw/webclient # note command deprecated in Docker 18.03+

# optional step - run container locally to test image
# docker run --rm -d -p 8080:8080 --name web_container us.gcr.io/airliftnw-uw/webclient
# curl http://localhost:8080