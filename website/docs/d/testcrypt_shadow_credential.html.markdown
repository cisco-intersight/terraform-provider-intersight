---
layout: "intersight"
page_title: "Intersight: intersight_testcrypt_shadow_credential"
sidebar_current: "docs-intersight-data-source-testcryptShadowCredential"
description: |-
Shadow Mo to read secure property, only for test.

---

# Data Source: intersight_testcrypt_shadow_credential
Shadow Mo to read secure property, only for test.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `password`:(string)Password associated with username.
* `username`:(string)The username of user, whose password marked as 'secure'.
