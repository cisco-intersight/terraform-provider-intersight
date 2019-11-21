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
* `dn`:(string)
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Disk name available in storage array.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `part_number`:(string)Storage disk part number.
* `protocol`:(string)Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe.
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `speed`:(int)Disk speed for read or write operation measured in rpm.
* `status`:(string)Storage disk health status.
* `type`:(string)Storage disk type, it can be SSD, HDD, NVRAM.
* `vendor`:(string)This field identifies the vendor of the given component.
* `version`:(string)Storage disk version number.
