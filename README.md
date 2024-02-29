# Authentication Microservice

## Logic
1. POST to /login the login name (UPN/email)
1. Lookup authenitcation strategy for user (Microsoft OAuth | Google OAuth | SAML 2.0)
1. Start authentication flow with appropriate IdP
    1. Store OAuth token if applicable
1. Generate JWT with user information, permissions, expiration time
1. Return JWT for use with subsequent requests
1. Expose endpoints for services to check authentication

## Routes
Below are the routes exposed by the service.

POST **/saml/acs**

GET **/saml/metadata**

POST /login

GET /logout

GET /user/\[attributes]