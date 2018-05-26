#!/bin/bash

GOOS=linux go build

docker build -t us.gcr.io/airliftnw-uw/go-apiserver:0.26 .

go clean

docker push us.gcr.io/airliftnw-uw/go-apiserver:0.26

## [WARNING]: ensure YAML spec updated to current version number
## (elevate-apiserver-deployment.yaml)

kubectl apply -f deployment/elevate-apiserver-deployment.yaml --record
kubectl rollout status deployment/go-api-server-deployment