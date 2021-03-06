
---
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_limits"
sidebar_current: "docs-intersight-data-source-iam-resource-limits"
description: |-
The resource limits used to limit resources such as User accounts.
---

# Data Source: intersight_iam._resource_limits
The resource limits used to limit resources such as User accounts.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `per_account_user_limit`:(int) The maximum number of users allowed in an account. The default value is 200. 
