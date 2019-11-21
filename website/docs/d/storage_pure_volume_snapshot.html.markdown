---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_volume_snapshot"
sidebar_current: "docs-intersight-data-source-storagePureVolumeSnapshot"
description: |-
Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.

---

# Data Source: intersight_storage_pure_volume_snapshot
Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created_time`:(string)Exact date and time on which snapshot is created.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the snapshot which represents point in time copy of volume.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `protection_group_name`:(string)Name of the protection group to which the snapshot belongs. Value is emptry, if the snapshot is created directly on volume.
* `serial`:(string)Unique serial number of the snapshot allocated by the storage array.
* `size`:(int)Snapshot size represented in bytes.
* `source`:(string)Source object on which the snapshot is created. It is a name of the originating volume.
