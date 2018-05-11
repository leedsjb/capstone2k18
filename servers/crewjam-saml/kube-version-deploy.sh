#!/bin/bash

# Builds and pushes to Google Cloud Registry for Kubernetes Deployments

usage="Builds and pushes a test Go lang web server image to Google Cloud Registry
for deployment on a web server. (us.gcr.io/airliftnw-uw/go-webclient)

Must be used concurrently with:
    kubectl apply -f <crewjam-deployment.yaml>
for changes to take effect.
    
Example:
    # Builds version 0.17 of 'go-webclient' and pushes to GCR
    ./gcr-version-push.sh 0.17
    
Usage:
    ./gcr-version-push.sh <version-number #.##>"

if [ "$#" -ne 1 ]; then
    echo "$usage"    
    exit 1
fi

version=$1

GOOS=linux go build
docker build -t us.gcr.io/airliftnw-uw/go-webclient:$version .
docker push us.gcr.io/airliftnw-uw/go-webclient:$version