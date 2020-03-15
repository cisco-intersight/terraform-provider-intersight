---
layout: "intersight"
page_title: "Intersight: intersight_iam_api_key"
sidebar_current: "docs-intersight-data-source-iamApiKey"
description: |-
API keys are used to programatically perform API calls. API keys can be created by passing purpose field, which is a description for the application using API keys. Maximum of 3 API keys per user is allowed. API key will have RSA key pair generated and as part of create request key MOID, private and public key in PEM format are returned. In Intersight only the public key is stored. Client side private key is stored and is used for signature calculation. API key requests are authenticated using RSA SHA 256 signature validations. If the incoming request for authorization has X-Starship-Token, then session based authorization is done, else API key based authorization is performed. Once User, API key and Account are found and signature verification succeeds, then the privilege validations are performed.
---

# Data Source: intersight_iam_api_key
API keys are used to programatically perform API calls. API keys can be created by passing purpose field, which is a description for the application using API keys. Maximum of 3 API keys per user is allowed. API key will have RSA key pair generated and as part of create request key MOID, private and public key in PEM format are returned. In Intersight only the public key is stored. Client side private key is stored and is used for signature calculation. API key requests are authenticated using RSA SHA 256 signature validations. If the incoming request for authorization has X-Starship-Token, then session based authorization is done, else API key based authorization is performed. Once User, API key and Account are found and signature verification succeeds, then the privilege validations are performed.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hash_algorithm`:(string)The cryptographic hash algorithm to calculate the message digest.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `private_key`:(string)Holds the private key for the API key.
* `purpose`:(string)The purpose of the API Key.
* `signing_algorithm`:(string)The signing algorithm used by the client to authenticate API requests to Intersight.The following key generation schemes are supported:1. RSASSA-PSS, as defined in RFC 8017 [RFC8017], Section 8.1,2. ECDSA P-256, as defined in ANSI X9.62-2005 ECDSA and FIPS 186-4,3. Ed25519ph, Ed25519ctx, and Ed25519, as defined in RFC 8032 [RFC8032], Section 5.1.The signing algorithm must be compatible with the key generation specification.
