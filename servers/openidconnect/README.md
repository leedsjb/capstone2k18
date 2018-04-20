




JWT Authentication Steps

1. User visits Elevate protected resource or clicks sign-in. Elevate detects user does not have an
    active session

2. Elevate redirects user to UW IdP server. Included in redirect is a return-to URL (possibly as a 
    query string parameter) 
    e.g.: idp.u.washington.edu?return-to="https://app.elevate.airliftnw.org/home"

3. User authenticates to UW IdP successfully. UW IdP packages JWT with return-to query param as 
    part of the JWT token

4. UW IdP ships JWT token to elevate.airliftnw.org/access/jwt?jwt={JWT w/ return-to embedded}

5. Elevate creates session for user based on JWT payload and stores session token in users cookies
    or local storage 


JWT Anatomy:

    JWT Header
    JWT Claims Set
    JWS Signature
