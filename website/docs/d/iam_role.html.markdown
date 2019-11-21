---
layout: "intersight"
page_title: "Intersight: intersight_iam_role"
sidebar_current: "docs-intersight-data-source-iamRole"
description: |-
A role is a collection of privilege sets that are assigned to a user using a permission object.

---

# Data Source: intersight_iam_role
A role is a collection of privilege sets that are assigned to a user using a permission object.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Informative description about each role.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the role which has to be granted to user.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
