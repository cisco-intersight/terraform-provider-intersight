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
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the protection Group."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `prefix`:(string)"Prefix used for all generated snapshots from the protection group."
* `replication_enabled`:(bool)"Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set."
* `size`:(int)"Overall size of all snapshots in the protection group, represented in bytes."
* `snapshot_enabled`:(bool)"Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set."
* `source`:(string)"Name of PureStorage array name on which the protection group is created."
