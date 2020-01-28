---
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_privilege"
sidebar_current: "docs-intersight-data-source-iamEndPointPrivilege"
description: |-
The privilege defined at the end point which can be assigned to a user.

---

# Data Source: intersight_iam_end_point_privilege
The privilege defined at the end point which can be assigned to a user.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)The functionality of this privilege.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the end point privilege.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `type`:(string)The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC.
