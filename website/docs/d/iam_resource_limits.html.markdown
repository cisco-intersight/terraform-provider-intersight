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
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `per_account_user_limit`:(int)"The maximum number of users allowed in an account. The default value is 200."
