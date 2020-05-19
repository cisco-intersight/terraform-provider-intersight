---
layout: "intersight"
page_title: "Intersight: intersight_iam._app_registration"
sidebar_current: "docs-intersight-data-source-iam.AppRegistration"
description: |-
AppRegistration encapsulates the meta-data values of a registered OAuth2 client application, as described in
https://tools.ietf.org/html/rfc7591#section-2.
Registered client applications have a set of metadata values associated with their client identifier
at the Intersight authorization server, including the list of valid redirection URIs or a display name.
The meta-data is used to specify how a client application can retrieve a OAuth2 Access Token and subsequently
invoke Intersight API on behalf of this AppRegistration.
To register an OAuth2 application, the following information must be provided.
1) Application name
2) An icon for the application
3) URL to the application's home page
4) A short description of the application
5) A list of redirect URLs
When an AppRegistration is created, a unique OAuth2 clientId is generated and returned in the HTTP response.
---

# Data Source: intersight_iam._app_registration
AppRegistration encapsulates the meta-data values of a registered OAuth2 client application, as described in
https://tools.ietf.org/html/rfc7591#section-2.
Registered client applications have a set of metadata values associated with their client identifier
at the Intersight authorization server, including the list of valid redirection URIs or a display name.
The meta-data is used to specify how a client application can retrieve a OAuth2 Access Token and subsequently
invoke Intersight API on behalf of this AppRegistration.
To register an OAuth2 application, the following information must be provided.
1) Application name
2) An icon for the application
3) URL to the application's home page
4) A short description of the application
5) A list of redirect URLs
When an AppRegistration is created, a unique OAuth2 clientId is generated and returned in the HTTP response.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `client_id`:(string)"A unique identifier for the OAuth2 client application.\nThe client ID is auto-generated when the AppRegistration object is created."
* `client_name`:(string)"App Registration name specified by user."
* `client_secret`:(string)"The OAuth2 client secret.\nThe value of this property is generated when grantType includes 'client-credentials'.\nOtherwise, no client-secret is generated."
* `client_type`:(string)"The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1."
* `description`:(string)"Description of the application."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `renew_client_secret`:(bool)"Set value to true to renew the client-secret. Applicable to client_credentials grant type."
* `revocation_timestamp`:(string)"Used to perform revocation for tokens of AppRegistration.\nUpdated only internally is case Revoke property come from UI with value true.\nOn each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the\ncorresponding App Registration."
* `revoke`:(bool)"Used to trigger update the revocationTimestamp value.\nIf UI sent updating request with the Revoke value is true, then update RevocationTimestamp."
