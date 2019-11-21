---
layout: "intersight"
page_title: "Intersight: intersight_iam_permission"
sidebar_current: "docs-intersight-data-source-iamPermission"
description: |-
Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.

---

# Data Source: intersight_iam_permission
Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)The informative description about each permission.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the permission which has to be granted to user.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
