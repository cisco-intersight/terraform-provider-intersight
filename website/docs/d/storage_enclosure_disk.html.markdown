---
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure_disk"
sidebar_current: "docs-intersight-data-source-storageEnclosureDisk"
description: |-
Physical Disk on the enclosure.
---

# Data Source: intersight_storage_enclosure_disk
Physical Disk on the enclosure.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `block_size`:(string)The block size of the physical disk in bytes.
* `device_mo_id`:(string)
* `disk_id`:(string)This field represents the disk Id in the storage enclosure.
* `disk_state`:(string)This field identifies the current disk configuration applied in the physical disk.
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `health`:(string)
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_blocks`:(string)The number of blocks present on the physical disk.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `pid`:(string)This field identifies the Product ID for physicalDisk.
* `revision`:(string)
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `sas_address1`:(string)This field identifies the SAS address assigned to the disk SAS port-1.
* `sas_address2`:(string)This field identifies the SAS address assigned to the disk SAS port-2.
* `serial`:(string)This field identifies the serial of the given component.
* `size`:(string)The size of the physical disk in MB.
* `vendor`:(string)This field identifies the vendor of the given component.
