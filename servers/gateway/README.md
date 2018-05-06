# Filename: README.md
# Created: April 19, 2018
# Modified:
# Author: J. Benjamin Leeds
# License: None

Development private key and unsigned certs can be generated for local testing using: 
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj "/CN=localhost" -keyout privkey.pem -out fullchain.pem