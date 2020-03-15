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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
