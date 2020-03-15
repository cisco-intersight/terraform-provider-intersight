---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_authorization"
sidebar_current: "docs-intersight-data-source-softwarerepositoryAuthorization"
description: |-
Consent that a user has provided to Intersight for contacting an external software repository such as cisco.com, on the user account's behalf.
---

# Data Source: intersight_softwarerepository_authorization
Consent that a user has provided to Intersight for contacting an external software repository such as cisco.com, on the user account's behalf.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_password_set`:(bool)
* `is_user_id_set`:(bool)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `password`:(string)The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
* `repository_type`:(string)The external repository for which this authorization has been provided. The only supported repository today is cisco.com.
* `user_id`:(string)The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
