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
#   gcloud container clusters get-credentials elevate-cluster --zone us-west1-b --project airliftnw-uw

deploymentName=$1

# note: it is important new image for new deployment has a unique [image-tag] 
# to facilitate rollback if needed

# kubectl set image deployment/[deployment-name] [container-name]=[image-name]:[image-tag]
# --record flag includes command causing deployment revision when executing: kubectl rollout history
kubectl set image deployment/alnw-deployment alnw-deployment=us.gcr.io/airliftnw-uw/webclient:0.3 --record

<< COMMENT

For more details visit: 
https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#updating-a-deployment

kubectl describe deployments -> displays current deployment details incuding image details

kubectl rollout status deployment/alnw-deployment -> display rollout details

kubectl get deployments -> summarize all deployments

kubectl get rs -> display status of all replica sets

kubectl get pods -> display status of individual pods
    pod name format: [deployment-name]-[unique-deployment-id]-[pod-id]
    e.g.: alnw-deployment-7fdc6c66f8-xbfz6

View rollout history summary:
kubectl rollout history deployment/alnw-deployment -> display rollout revision history

View rollout history detail:
kubectl rollout history deployment/alnw-deployment --revision=3

Rollback deployment: 
kubectl rollout undo deployment/alnw-deployment --to-revision=2

Print Deployment YAML
kubectl get deployment alnw-deployment -o yaml

COMMENT