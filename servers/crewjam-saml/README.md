


Obtain a signed certificate using Certbot and LetsEncrypt:

Domain: crewjam-saml.test.emeloid.co

Manual Method:

sudo certbot certonly --manual --preferred-challenges dns
Add generated DNS entries to verify domain
Cert and key are saved

----------------------------------------------------------
Google DNS Verification Method:

Create Service Account JSON:

gcloud iam service-accounts keys create \
    ./secrets/certbot/google.json \
    --iam-account certbot@airliftnw-uw.iam.gserviceaccount.com

Ensure service account has following permissions:

dns.changes.create
dns.changes.get
dns.managedZones.list
dns.resourceRecordSets.create
dns.resourceRecordSets.delete
dns.resourceRecordSets.list
dns.resourceRecordSets.update


Run Certbot Google-DNS Docker Image

sudo docker run -it --rm --name certbot \
-v "etc/letsencrypt:/etc/letsencrypt" \
-v "var/lib/letsencrypt:/var/lib/letsencrypt" \
certbot/dns-google certonly


----

certbot certonly \
  --dns-google \
  --dns-google-credentials ./secrets/certbot/google.json \
  -d emeloid.co