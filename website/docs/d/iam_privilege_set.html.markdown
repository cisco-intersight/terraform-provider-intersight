---
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege_set"
sidebar_current: "docs-intersight-data-source-iamPrivilegeSet"
description: |-
A collection of privileges.

---

# Data Source: intersight_iam_privilege_set
A collection of privileges.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the privilege set.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the privilege set.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
