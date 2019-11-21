---
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_limits"
sidebar_current: "docs-intersight-data-source-iamResourceLimits"
description: |-
The resource limits used to limit resources such as User accounts.

---

# Data Source: intersight_iam_resource_limits
The resource limits used to limit resources such as User accounts.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `per_account_user_limit`:(int)The maximum number of users allowed in an account. The default value is 200.
