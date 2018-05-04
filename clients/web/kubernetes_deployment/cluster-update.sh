#!/bin/bash

# Filename: cluster_update.sh
# Created: 
# Modified: April 14, 2018
# Author: Tiffany Chen
# License: None
# Purpose: Script to update the image running on a cluster

deploymentName=$1

kubectl set image alnw-deployment webclient=us.gcr.io/airliftnw-uw/webclient # syntax: [container-name]=[image-name]