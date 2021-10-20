#!/bin/zsh
# zshell script to generate a private key

openssl genrsa -des3 -out privkey1.pem 2048

openssl req -new -sha256 -key privkey1.pem -out server.csr

openssl req -x509 -sha256 -days 365 -key privkey1.pem -in server.csr -out fullchain1.pem