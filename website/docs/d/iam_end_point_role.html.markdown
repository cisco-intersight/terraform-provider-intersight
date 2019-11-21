---
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_role"
sidebar_current: "docs-intersight-data-source-iamEndPointRole"
description: |-
The role defined in the end point which can be assigned to a user.

---

# Data Source: intersight_iam_end_point_role
The role defined in the end point which can be assigned to a user.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the end point role.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `role_type`:(string)User specified tags to roles like as ep-admin or ep-readonly.
* `type`:(string)The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.
