---
layout: "intersight"
page_title: "Intersight: intersight_storage_virtual_drive_extension"
sidebar_current: "docs-intersight-data-source-storageVirtualDriveExtension"
description: |-
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
---

# Data Source: intersight_storage_virtual_drive_extension
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bootable`:(string)"It shows virtual drive is bootable."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `container_id`:(int)
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `drive_state`:(string)
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `oper_device_id`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `uuid`:(string)
* `vendor_uuid`:(string)
* `virtual_drive_dn`:(string)
* `virtual_drive_id`:(string)"It shows virtual drive Id."
