#!/bin/bash
# add functionality to take in an argument to use as cluster name
# ./cluster-deploy.sh elevate-cluster

# create cluster(s)
gcloud container clusters create elevate-cluster --num-nodes=1

# check cluster's worker VM instances
gcloud compute instances list

# if using pre-existing cluster, i.e. *created cluster through GCP Console*
# run this to retrieve cluster credentials and configure kubectl with them
# gcloud container clusters get-credentials elevate-cluster

# deploy application
kubectl run alnw-deployment --image=us.gcr.io/airliftnw-uw/webclient --port 80

# see pod created by deployment
kubectl get pods


kubectl expose deployment alnw-deployment --type=LoadBalancer --port 80 --target-port 80

#check
kubectl get service