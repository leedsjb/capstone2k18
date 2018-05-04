#!/bin/bash

# Filename: cluster_update.sh
# Created: 
# Modified: April 14, 2018
# Author: Tiffany Chen
# License: None
# Purpose: Script to update the image running on a cluster

# Preconditions: 
#   must have kubectl context set for the cluster you wish to update
#   Use: kubectl config current-context to verify and 
#   gcloud container clusters get-credentials elevate-cluster

deploymentName=$1

# kubectl set image [deployment-name] [container-name]=[image-name]:[image-tag]
kubectl set image alnw-deployment alnw-deployment=us.gcr.io/airliftnw-uw/webclient:latest 
