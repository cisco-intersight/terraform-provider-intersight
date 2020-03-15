---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_disk"
sidebar_current: "docs-intersight-data-source-storagePureDisk"
description: |-
Disk entity associated with Pure FlashArray.
---

# Data Source: intersight_storage_pure_disk
Disk entity associated with Pure FlashArray.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Disk name available in storage array.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `part_number`:(string)Storage disk part number.
* `protocol`:(string)Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe.
* `revision`:(string)
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `serial`:(string)This field identifies the serial of the given component.
* `speed`:(int)Disk speed for read or write operation measured in rpm.
* `status`:(string)Storage disk health status.
* `type`:(string)Storage disk type - it can be SSD, HDD, NVRAM.
* `vendor`:(string)This field identifies the vendor of the given component.
* `version`:(string)Storage disk version number.
