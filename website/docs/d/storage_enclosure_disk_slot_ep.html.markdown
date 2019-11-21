---
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure_disk_slot_ep"
sidebar_current: "docs-intersight-data-source-storageEnclosureDiskSlotEp"
description: |-
Physical Disk slots on the enclosure.

---

# Data Source: intersight_storage_enclosure_disk_slot_ep
Physical Disk slots on the enclosure.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)
* `drive_path`:(string)This field identifies the zoning configuration applied to  this enclosure slot.
* `health`:(string)This field identifies the health of the disk inserted in the slot.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `presence`:(string)This field identifies the disk is present in the enclosure slot.
* `rn`:(string)
* `slot`:(string)This field represents the slot Id in the storage enclosure.
