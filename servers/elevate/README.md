# Filename: README.md
# Created: April 19, 2018
# Modified:
# Author: J. Benjamin Leeds
# License: None

Development private key and unsigned certs can be generated for local testing using: 
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj "/CN=localhost" -keyout privkey.pem -out fullchain.pem

For manual testing, begin Cloud SQL Proxy using:
./cloud_sql_proxy -instances=<INSTANCE_CONNECTION_NAME>=tcp:3306

<INSTANCE_CONNECTION_NAME> can be found at https://console.cloud.google.com/sql/instances/alnw-dev/overview?project=airliftnw-uw&authuser=1&duration=PT1H

Can then connect using:
mysql -u <USERNAME> -p --host 127.0.0.1