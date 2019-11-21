---
layout: "intersight"
page_title: "Intersight: intersight_testcrypt_credential"
sidebar_current: "docs-intersight-data-source-testcryptCredential"
description: |-
Basic Mo with secure property.

---

# Data Source: intersight_testcrypt_credential
Basic Mo with secure property.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_password_set`:(bool)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `password`:(string)Password associated with username, to be stored in the database encrypted.
* `username`:(string)The name of the user whose credentials are being set.
