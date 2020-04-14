---
layout: "intersight"
page_title: "Intersight: intersight_storage_sas_port"
sidebar_current: "docs-intersight-data-source-storageSasPort"
description: |-
Sas Port details of the SAS endpoint.
---

# Data Source: intersight_storage_sas_port
Sas Port details of the SAS endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `address`:(string)"The SAS Address assigned to storage port."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `disk_id`:(int)"The disk identifier."
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `end_point_id`:(int)"The end-point Id assigned to storage port."
* `link_description`:(string)"The link description."
* `link_speed`:(string)"The link speed negotiated for communication."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
