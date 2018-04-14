#!/bin/bash

# Filename: key-gen.sh
# Created: April 10, 2018
# Modified: April 14, 2018
# Author: J. Benjamin Leeds
# License: None
# Purpose: This script generates private keys and their associated certificates for SSL connections
# Reference: https://cloud.google.com/compute/docs/load-balancing/http/ssl-certificates

# Generate new private key w/ OpenSSL using RSA-2048 encryption
openssl genrsa -out example.key 2048

# Generate signed certificate using a Certificate Signing Request (CSR)
# CSR requires information in the form of a Distinguished Name:
    # Country Name: US
    # State Name: Washington
    # Locality: Seattle
    # Organization Name: University of Washington
    # Organizational Unit Name: UW-IT
    # Common Name: 
    # Email Address: 
    # Challenge Password: may not be needed, password included with CSR only

openssl req -new -key example.key -out example.csr

# For test purposes a self-signed certificate can be generated. In prod a CA needs to sign the cert
openssl x509 -req -days 365 -in example.csr -signkey example.key -out example.crt

# Next up: generating an SSL Certificate Resource
# The SSL certificate resource holds an SSL certificate and makes it available to the load balancer

# create the resource: provide name for resource and filepath to cert and private key
# GCP does not validate whether certificate chain is valid, ensure valid chain prior to creation
gcloud compute ssl-certificates create
    [SSL_CERTIFICATE] \
    --certificate [CRT_FILE_PATH] \
    --private-key [KEY_FILE_PATH]
