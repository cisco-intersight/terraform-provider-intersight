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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
