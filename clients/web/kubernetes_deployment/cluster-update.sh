#!/bin/bash
# Script to update the image running on a cluster

deploymentName=$1


kubectl set image alnw-deployment webclient=nginx:1.9.1