#!/bin/bash

GOOS=linux go build

docker build -t us.gcr.io/airliftnw-uw/go-apiserver:0.3 .

go clean

docker push us.gcr.io/airliftnw-uw/go-apiserver:0.3