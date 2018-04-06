#!/bin/bash

# delete cluster - only ever have one
minikube delete

# start new cluster
minikube start --vm-driver=hyperkit

# optional, push docker image up to cloud
# docker tag local/airliftnw-uw/webclient us.gcr.io/airliftnw-uw/local-alnw-webclient
# gcloud docker --push us.gcr.io/airliftnw-uw/local-alnw-webclient

# run locally stored webclient image
# cloud stored is us.gcr...
# kubectl run alnwnginx --image=local/airliftnw-uw/webclient

# delete pre-existing image
kubectl delete deployments/alnwnginx

# run cloud stored webclient image
# us.gcr.io/[name-of-project]/[name-of-image]
kubectl run alnwnginx --image=us.gcr.io/airliftnw-uw/local-alnw-webclient