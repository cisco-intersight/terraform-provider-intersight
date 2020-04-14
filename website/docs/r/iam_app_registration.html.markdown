---
layout: "intersight"
page_title: "Intersight: intersight_iam_app_registration"
sidebar_current: "docs-intersight-resource-iamAppRegistration"
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

# Resource: intersight_iam_app_registration
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
The following arguments are supported:
* `account`:(Array with Maximum of one item) -(Computed)"A collection of references to the [iam.Account](mo://iam.Account) Managed Object.\nWhen this managed object is deleted, the referenced [iam.Account](mo://iam.Account) MO unsets its reference to this deleted MO."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `client_id`:(string)(Computed)"A unique identifier for the OAuth2 client application.\nThe client ID is auto-generated when the AppRegistration object is created."
* `client_name`:(string)"App Registration name specified by user."
* `client_secret`:(string)"The OAuth2 client secret.\nThe value of this property is generated when grantType includes 'client-credentials'.\nOtherwise, no client-secret is generated."
* `client_type`:(string)"The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1."
* `description`:(string)"Description of the application."
* `grant_types`:
                (Array of schema.TypeString) -(Computed)"The set of grant types that OAuth2 clients can use for this application.\nThe grant type is used in the OAuth2 login flow to validate the grant type that has been requested by the client.\nSee https://tools.ietf.org/html/rfc7591#page-9 for more details.\n# It is set automatically when AppRegistration is created since currently we do not provide option for the user."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `oauth_tokens`:(Array)(Computed)"Collection of the OAuth2 tokens. Each OAuth2 token represents valid OAuth session.\nOAuth2 token is created when login over OAuth2 is performed using Authorization Code grant type."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `permission`:(Array with Maximum of one item) -(Computed)"Permission associated with OAuth token issued through Client Credentials flow. Permission of the current session will be used."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `redirect_uris`:
                (Array of schema.TypeString) -"After a user successfully authorizes an application, the OAuth2 authorization server will redirect the user back to the\napplication with either an authorization code or access token in the URL.\nRegistered redirect URLs may contain query string parameters, but must not contain anything in the fragment.\nThe registration server rejects the request if a user tries to register a redirect URL that contains a fragment.\nFor native and mobile apps, Intersight allows a user to register a URL scheme such as myapp:// which can then be used\nin the redirect URL. The authorization server allows arbitrary URL schemes to be registered in order to support\nregistering redirect URLs for native apps.\nRedirect URLs are a critical part of the OAuth flow. After a user successfully authorizes an application,\nthe authorization server will redirect the user back to the application with either an authorization code or access\ntoken in the URL. Because the redirect URL will contain sensitive information, it is critical that the service\ndoesn’t redirect the user to arbitrary locations.\nThe best way to ensure the user will only be directed to appropriate locations is to require the developer to\nregister one or more redirect URLs when they create the application.\nThe redirection endpoint URI MUST be an absolute URI."
* `renew_client_secret`:(bool)"Set value to true to renew the client-secret. Applicable to client_credentials grant type."
* `response_types`:
                (Array of schema.TypeString) -(Computed)"The set of response types that a OAuth2 client can use.\nThis is static list and it is set automatically when AppRegistration is created.\nAccording to RFC, it is used in OAuth2 login flow to check that this AppRegistration supports response type from the request.\nSee https://tools.ietf.org/html/rfc7591#page-9 for more details."
* `revocation_timestamp`:(string)(Computed)"Used to perform revocation for tokens of AppRegistration.\nUpdated only internally is case Revoke property come from UI with value true.\nOn each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the\ncorresponding App Registration."
* `revoke`:(bool)"Used to trigger update the revocationTimestamp value.\nIf UI sent updating request with the Revoke value is true, then update RevocationTimestamp."
* `roles`:(Array)"The set of roles that can be used when a OAuth2 client is accessing this registered application.\nFor example, multiple roles may be defined in your Intersight account, but you want users to login\nwith the 'Read-Only' role when accessing Intersight through a registered application.\nIn that case, the 'roles' property should contain a single element referencing the 'Read-Only' role.\nA user can only assign roles they already have.\nThis relationship is deprecated. Authorization is now performed by passing the 'scope' query parameter\nin the first request of the Authorization Code OAuth2 flow.\nThe value of the 'scope' parameter is a list of scope names separated by comma:\nROLE.Account Administrator, ROLE.<any role name>."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
* `user`:(Array with Maximum of one item) -(Computed)"A collection of references to the [iam.User](mo://iam.User) Managed Object.\nWhen this managed object is deleted, the referenced [iam.User](mo://iam.User) MO unsets its reference to this deleted MO."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
