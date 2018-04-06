#!/bin/bash
# Create a cluster in gcloud and deploy an application on it

usage="Creates a cluster in gcloud and deploys an application on it from a given image source. 

Ports for the application and LoadBalancer are defaulted to port 80.
    
Example:
    # Create a cluster named 'elevate-cluster' that has '1' node
    # and deploy an application named 'alnw-deployment' on it
    # sourced from the image 'us.gcr.io/airliftnw-uw/webclient'

    ./cluster-deploy.sh elevate-cluster 1 alnw-deployment us.gcr.io/airliftnw-uw/webclient
    
Usage:
    ./cluster-deploy.sh <cluster-name> <num-nodes> <deployment-name> <image-source>"

if [ "$#" -ne 4 ]; then
    echo "$usage"    
    exit 1
fi

clusterName=$1
numNodes=$2
deploymentName=$3
imageSource=$4
# var for ports?

# create cluster(s)
gcloud container clusters create $clusterName --num-nodes=$numNodes

# check cluster's worker VM instances
gcloud compute instances list

# if using pre-existing cluster, i.e. *created cluster through GCP Console*
# run this to retrieve cluster credentials and configure kubectl with them
# gcloud container clusters get-credentials elevate-cluster

# deploy application
kubectl run $deploymentName --image=$imageSource --port 80

# see pod created by deployment
kubectl get pods


kubectl expose deployment $deploymentName --type=LoadBalancer --port 80 --target-port 80

#check
kubectl get service