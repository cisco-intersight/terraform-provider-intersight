
---
layout: "intersight"
page_title: "Intersight: intersight_testcrypt_credential"
sidebar_current: "docs-intersight-data-source-testcrypt-credential"
description: |-
Basic Mo with secure property.
---

# Data Source: intersight_testcrypt._credential
Basic Mo with secure property.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `password`:(string) Password associated with username, to be stored in the database encrypted. 
* `username`:(string) The name of the user whose credentials are being set. 
