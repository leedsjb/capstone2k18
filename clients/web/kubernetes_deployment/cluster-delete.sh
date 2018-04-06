#!/bin/bash

# deallocate Cloud Load Balancer created for Service
kubectl delete service alnw-deployment

# load balancer is deleted asynchronously - wait for it to be deleted depending on output
# of this command:
gcloud compute forwarding-rules list

# delete container cluster
gcloud container clusters delete elevate-cluster

