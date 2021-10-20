#!/bin/bash

# This script starts the Google Cloud SQL Proxy and connects to the alnw-elevate-test instance
# Author: J. Benjamin Leeds
# Modified: Wednesday October 20, 2021

# Connection Specification
# Project: alnw-elevate
# Region and Zone: us-west1
# Cloud SQL Instance: elevate

# Reference: https://github.com/GoogleCloudPlatform/cloudsql-proxy

# Install latest version of proxy with:
# go install github.com/GoogleCloudPlatform/cloudsql-proxy/cmd/cloud_sql_proxy@1.26.0
# binary is installed in $GOPATH/bin and can be run by executable name: cloud_sql_proxy

# Install VSCODE MySQL Marketplace Extension: formulahendry.vscode-mysql

# Start the proxy to connect to the database
# ./cloud_sql_proxy -instances=airliftnw-uw:us-west1:alnw-elevate-test=tcp:3306
cloud_sql_proxy -instances=alnw-elevate:us-west1:elevate=tcp:3306

# Connect to the proxy locally with MySQL Extension:
# host: 127.0.0.1
# user: root
# password: 
# port: 3306 (or any locally avail port)