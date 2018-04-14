#!/bin/bash

# Filename: GCP_container_registry_private_image_access.sh
# Created: 
# Modified: April 14, 2018
# Author: J. Benjamin Leeds
# License: None

# This script allows one to access a Private GCP Cloud Repository Docker image
# from a non GCP Kubernetes environment (e.g. minikube)
# Steps obtained from: http://docs.heptio.com/content/private-registries/pr-gcr.html
# Script must be run by user with sufficient permissions: gcloud.projects.add-iam-policy-binding

# create a GCP service account; format of account is email address
SA_EMAIL=$(gcloud iam service-accounts --format='value(email)' create k8s-gcr-auth-ro)

# create the json key file and associate it with the service account
gcloud iam service-accounts keys create k8s-gcr-auth-ro.json --iam-account=$SA_EMAIL

# get the project id
PROJECT=$(gcloud config list core/project --format='value(core.project)')

# add the IAM policy binding for the defined project and service account
gcloud projects add-iam-policy-binding $PROJECT --member serviceAccount:$SA_EMAIL --role roles/storage.objectViewer

SECRETNAME=k8s-local-access-secret

kubectl create secret docker-registry $SECRETNAME \
  --docker-server=https://us.gcr.io \
  --docker-username=_json_key \
  --docker-email=airliftnorthwestcloud@gmail.com \
  --docker-password="$(cat k8s-gcr-auth-ro.json)"

kubectl patch serviceaccount default \
  -p "{\"imagePullSecrets\": [{\"name\": \"$SECRETNAME\"}]}"