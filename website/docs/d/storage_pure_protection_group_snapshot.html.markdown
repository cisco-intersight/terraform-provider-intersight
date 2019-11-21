---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_protection_group_snapshot"
sidebar_current: "docs-intersight-data-source-storagePureProtectionGroupSnapshot"
description: |-
Protection group snapshot entity in Pure protection group.

---

# Data Source: intersight_storage_pure_protection_group_snapshot
Protection group snapshot entity in Pure protection group.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created_time`:(string)Protection group snapshot created time.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Protection group snapshot name which represents point in time copy of all members in protection group.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `size`:(int)Snapshot size represented in bytes. It is a cummulative size of all snapshots in a set.
* `source`:(string)Source protection group name on which the snapshot is created.
