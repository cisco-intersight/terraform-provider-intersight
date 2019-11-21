---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_protection_group"
sidebar_current: "docs-intersight-data-source-storagePureProtectionGroup"
description: |-
Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.

---

# Data Source: intersight_storage_pure_protection_group
Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the protection Group.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `prefix`:(string)Prefix used for all generated snapshots from the protection group.
* `replication_enabled`:(bool)Flag to determine if the replication is enabled. Snapshots are created on target array if the flag is set.
* `size`:(int)Overall size of all snapshots in the protection group, represented in bytes.
* `snapshot_enabled`:(bool)Flag to determine if the snapshot is enabled. Snapshots are created on local array if the flag is set.
* `source`:(string)Name of PureStorage array name on which the protection group is created.
