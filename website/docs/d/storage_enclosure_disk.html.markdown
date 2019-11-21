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
* `dn`:(string)
* `health`:(string)
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_blocks`:(string)The number of blocks present on the physical disk.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `pid`:(string)This field identifies the Product ID for physicalDisk.
* `revision`:(string)
* `rn`:(string)
* `sas_address1`:(string)This field identifies the SAS address assigned to the disk SAS port-1.
* `sas_address2`:(string)This field identifies the SAS address assigned to the disk SAS port-2.
* `serial`:(string)This field identifies the serial of the given component.
* `size`:(string)The size of the physical disk in MB.
* `vendor`:(string)This field identifies the vendor of the given component.
