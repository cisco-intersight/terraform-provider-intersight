---
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure"
sidebar_current: "docs-intersight-data-source-storageEnclosure"
description: |-
Storage Enclosure for physical disks.

---

# Data Source: intersight_storage_enclosure
Storage Enclosure for physical disks.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chassis_id`:(int)
* `description`:(string)
* `device_mo_id`:(string)
* `dn`:(string)
* `enclosure_id`:(int)
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_slots`:(int)
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `presence`:(string)
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `server_id`:(int)
* `type`:(string)
* `vendor`:(string)This field identifies the vendor of the given component.
