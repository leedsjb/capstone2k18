#!/bin/bash

# This script starts the Google Cloud SQL Proxy and connects to the alnw-elevate-test instance

# Connection Specification
# Project: airliftnw-uw
# Region and Zone: us-west1
# Cloud SQL Instance: alnw-elevate-test

./cloud_sql_proxy -instances=airliftnw-uw:us-west1:alnw-elevate-test=tcp:3306