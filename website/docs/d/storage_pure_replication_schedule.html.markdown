---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_replication_schedule"
sidebar_current: "docs-intersight-data-source-storagePureReplicationSchedule"
description: |-
Pure snapshot replication schedule entity.
---

# Data Source: intersight_storage_pure_replication_schedule
Pure snapshot replication schedule entity.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `daily_limit`:(int)"Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage\nFlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired.\nIn case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day."
* `frequency`:(string)"Replication frequency. It is an interval at which replication is set to trigger.\nExamples:\n    PT2H, Snapshot is generated every 2 hours.\n    P30D, Snapshot is scheduled for every 30 days.\n    PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Replication schedule name."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `retention_time`:(string)"Duration to keep the replicated snapshots on the targets.\nReplicated snapshots are deleted from target array once the retention period has elapsed.\nExamples:\nP30D, Snapshots are available for 30 days.\nPT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds."
* `snapshot_expiry_time`:(string)"Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period."
