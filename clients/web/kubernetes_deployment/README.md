# Filename: README.sh
# Created: April 15, 2018 
# Modified: May 3, 2018
# Author: Tiffany Chen & J. Benjamin Leeds
# License: None
# Purpose: Deployment instructions for Elevate client-side components on Cloud Container Engine

1. yarn build: build the latest version of the client
2. docker build: Dockerize the new build into an image for containerization
3. docker push gcr: push the image to Google Container Registry
4. update Kubernetes cluster: use kubectl set image (note, development only)