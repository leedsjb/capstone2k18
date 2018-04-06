#!/bin/bash
# Use this script to build the docker image and deploy it to the google container registry

docker build -t us.gcr.io/airliftnw-uw/webclient ../ #move to dir containing Dockerfile

# upload conatiner to google container registry (so Kubernetes Engine can download and run it)
gcloud docker -- push gcr.io/airliftnw-uw/webclient

# optional step - run container locally to test image
# docker run --rm -p 8080:8080 gcr.io/airliftnw-uw/webclient
# curl http://localhost:8080