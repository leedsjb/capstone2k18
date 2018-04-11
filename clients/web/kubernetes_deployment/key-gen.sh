#!/bin/bash

# Filename: key-gen.sh
# Created: April 10, 2018
# Modified: 
# Author: J. Benjamin Leeds
# License: None

# This script generates private keys and their associated certificates for SSL connections
# Reference: https://cloud.google.com/compute/docs/load-balancing/http/ssl-certificates

# Generate new private key w/ OpenSSL using RSA-2048 encryption
openssl genrsa -out example.key 2048

# Generate signed certificate using a Certificate Signing Request (CSR)
# CSR requires information in the form of a Distinguised Name:
    # Country Name: US
    # State Name: Washington
    # Locality: Seattle
    # Organization Name: University of Washington
    # Organizational Unit Name: UW-IT
    # Common Name: 
    # Email Address: 
    # Challenge Password: may not be needed, password including with CSR only

openssl req -new -key example.key -out example.csr

# For test purposed a self-signed certificate can be generated. In prod a CA needs to sign the cert
openssl x509 -req -days 365 -in example.csr -signkey example.key -out example.crt

# Next up: generating an SSL Certificate Resource