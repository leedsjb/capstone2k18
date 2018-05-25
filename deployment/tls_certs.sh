#!bin/bash

# Filename: tls_certs.sh
# Created: Thursday May 24, 2018
# Modified:
# Author: J. Benjamin Leeds
# License: None
# Purpose: Scripts and commands to manage TLS certs for *.elevate.airliftnw.org

# Create Certificates:
## Uses Certbot docker image with Google DNS Challenge plugin
sudo docker run -it --rm --name certbot \
-v "/Users/benjaminleeds/Documents/Code/go/src/github.com/leedsjb/capstone2k18/deployment/:/etc/letsencrypt" \
-v "/Users/benjaminleeds/Documents/Code/go/src/github.com/leedsjb/capstone2k18/deployment/:/var/lib/letsencrypt" \
certbot/dns-google certonly --manual \
--preferred-challenges dns --server https://acme-v02.api.letsencrypt.org/directory \
-d elevate.airliftnw.org -d api.elevate.airliftnw.org \
-d test.elevate.airliftnw.org -d api.test.elevate.airliftnw.org

## Renewals:
certbot renew

# TODO desired directories to store letsencrypt files locally, has Docker mount issues
-v "/private/etc/letsencrypt:/etc/letsencrypt" \
-v "/private/var/lib/letsencrypt:/var/lib/letsencrypt" \

# flags for setting custom directories to store certbot files
# --config-dir ./.certbot/config --logs-dir ./.certbot/logs --work-dir ./.certbot/work