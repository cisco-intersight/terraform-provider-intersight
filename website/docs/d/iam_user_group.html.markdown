---
layout: "intersight"
page_title: "Intersight: intersight_iam_user_group"
sidebar_current: "docs-intersight-data-source-iamUserGroup"
description: |-
User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.

---

# Data Source: intersight_iam_user_group
User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the user group which the dynamic user belongs to.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
