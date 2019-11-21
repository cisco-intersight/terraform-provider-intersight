---
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_role"
sidebar_current: "docs-intersight-data-source-iamEndPointUserRole"
description: |-
Mapping of endpoint user to endpoint roles.

---

# Data Source: intersight_iam_end_point_user_role
Mapping of endpoint user to endpoint roles.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `change_password`:(bool)Denotes whether password has changed.
* `enabled`:(bool)Enables the user account on the endpoint.
* `is_password_set`:(bool)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `password`:(string)Password for the user
