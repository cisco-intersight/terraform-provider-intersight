---
layout: "intersight"
page_title: "Intersight: intersight_externalsite_authorization"
sidebar_current: "docs-intersight-data-source-externalsiteAuthorization"
description: |-
An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.

---

# Data Source: intersight_externalsite_authorization
An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_password_set`:(bool)
* `is_user_id_set`:(bool)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `password`:(string)The password of the given username to download the image from external repository like cisco.com.
* `repository_type`:(string)The repository type to which this authorization will be requested. Cisco is the only available repository today.
* `user_id`:(string)The username that has permission to download the image from external repository like cisco.com.
