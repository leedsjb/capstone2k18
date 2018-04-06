#!/bin/bash
# Delete a cluster and its associated service/image/application

usage="Deletes a cluster and its associated service. 

Note: Uses kubectl delete, which runs asynchronously. Be sure to wait for confirmed delete.
    
Example:
    # Delete the service 'alnw-deployment' running on 'elevate-cluser'
    # Then delete the cluster 'elevate-cluster'

    ./cluster-delete.sh alnw-deployment elevate-cluster
    
Usage:
    ./cluster-delete.sh <deployment-name> <cluster-name>"

if [ "$#" -ne 2 ]; then
    echo "$usage"    
    exit 1
fi

deploymentName=$1
clusterName=$2

# deallocate Cloud Load Balancer created for Service
kubectl delete service $deploymentName

# load balancer is deleted asynchronously - wait for it to be deleted depending on output
# of this command:
gcloud compute forwarding-rules list

# delete container cluster
gcloud container clusters delete $clusterName

