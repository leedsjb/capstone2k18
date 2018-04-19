# Filename: README.md
# Created: April 18, 2018
# Modified: 
# Author: J. Benjamin Leeds
# License: None

Shibboleth / SAML 2.0 Authentication via University of Washington NetID

To use Shibboleth with NGINX you must: 
    1. Rebuild Shibboleth SP with FastCGI Support
    2. Recompile NGINX with nginx-http-shibboleth Module

The FastCGI Authorizer and Responder applications faciliate authentication. 

FastCGI: Open Extension to CGI (http://www.mit.edu/~yandros/doc/specs/fcgi-spec.html)
    -> Contrasts w/ CGI in that FastCGI is long-running, not used for only 1 requests and exited
    -> Does not use env vars for input or stdin, stdout, and stderr
    -> Initial state: web socket to accept connections from NGINX

FastCGI Roles: 
    1. Responder: Receives HTTP Request and generates HTTP Response
    2. Authorizer: Receives HTTP Request and generates authorized / unauthorized decision

Shibboleth User Flow:

1. User tries to access protected resource
    -> Protected resources can be defined in shibboleth2.xml (or elsewhere)

2. SP determines IdP and Issues Authentication Request
    -> We will use fixed IdP: NetID
    -> Originally requested resource preserved using a "rely state" mechanism

3. User Authenticates to IdP
    -> Client redirected to NetID sign-on via GET or POST

4. IdP packages response
    -> Response in form of SAML Assertion
    -> Signed with IdPs key and encrypted with SPs key
    -> Sent to SP's Assertion Consumer Service endpoint (via client)

5. Back to SP
    -> IdP response passed by user's browser to the SP's Assertion Consumer Service (ACS) endpoint
    -> ACS decrypts/decodes IdP response, then: 
    -> Creates new user session and:
    -> Determines where to send browser based on information stored in the relay state
    -> When SP tells browser where to redirect it also sets a cookie to associate client with session
        using a session key (Bearer token?)

Note: It's important ACS endpoint and resource location share a virtual host
Virtual Host: two websites served by the same physical hardware (??), possibility different server 
    blocks in NGINX are themselves virtual hosts

