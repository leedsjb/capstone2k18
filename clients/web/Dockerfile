# Filename: Dockerfile.sh
# Created: 
# Modified: May 19, 2018
# Author: J. Benjamin Leeds
# License: None
# Purpose: Build container image for webclient server
# Note multi-stage build to load trusted root Certificate Authorities from alpine.

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY go-saml-web-server go-saml-web-server
COPY ./build build
EXPOSE 80
EXPOSE 443
ENTRYPOINT ["/go-saml-web-server"]