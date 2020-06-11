---
layout: "intersight"
page_title: "Intersight: intersight_storage_physical_disk"
sidebar_current: "docs-intersight-data-source-storagePhysicalDisk"
description: |-
Physical Disk on a server.
---

# Data Source: intersight_storage_physical_disk
Physical Disk on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `block_size`:(string)"The block size of the physical disk."
* `bootable`:(string)
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `configuration_checkpoint`:(string)
* `configuration_state`:(string)
* `device_mo_id`:(string)
* `discovered_path`:(string)
* `disk_id`:(string)
* `disk_state`:(string)"This field identifies the health of the disk."
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `drive_firmware`:(string)
* `drive_state`:(string)"The drive state as reported by the controller."
* `fde_capable`:(string)
* `link_speed`:(string)"The speed of the link between the drive and the controller."
* `link_state`:(string)
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `num_blocks`:(string)"The number of blocks present on the physical disk."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `oper_power_state`:(string)
* `oper_qualifier_reason`:(string)
* `operability`:(string)
* `physical_block_size`:(string)
* `pid`:(string)"This field identifies the Product ID for physicalDisk."
* `presence`:(string)
* `protocol`:(string)
* `raw_size`:(string)
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `secured`:(string)"This field identifies whether the disk is encrypted."
* `serial`:(string)"This field identifies the serial of the given component."
* `size`:(string)
* `thermal`:(string)
* `type`:(string)
* `variant_type`:(string)
* `vendor`:(string)"This field identifies the vendor of the given component."
