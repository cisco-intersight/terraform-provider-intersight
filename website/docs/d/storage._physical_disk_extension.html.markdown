---
layout: "intersight"
page_title: "Intersight: intersight_storage._physical_disk_extension"
sidebar_current: "docs-intersight-data-source-storage.PhysicalDiskExtension"
description: |-
Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
---

# Data Source: intersight_storage._physical_disk_extension
Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bootable`:(string)"It shows whether disk is bootable or not."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `disk_dn`:(string)"It shows the Physical drive Dn."
* `disk_id`:(int)"It shows storage Enclosure slotId."
* `disk_state`:(string)"It shows the current drive state of disk."
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `health`:(string)"It shows the current drive state of disk."
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `serial`:(string)"This field identifies the serial of the given component."
* `vendor`:(string)"This field identifies the vendor of the given component."
