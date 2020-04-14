---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_snapshot_schedule"
sidebar_current: "docs-intersight-data-source-storagePureSnapshotSchedule"
description: |-
PureStorage FlashArray snapshot schedule configuration entity.
---

# Data Source: intersight_storage_pure_snapshot_schedule
PureStorage FlashArray snapshot schedule configuration entity.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `daily_limit`:(int)"Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day."
* `frequency`:(string)"Snapshot frequency. It is an interval at which snapshot is set to trigger on source array.\nExamples:\n    PT2H Snapshot is generated every 2 hours.\n    P4D, Snapshot is scheduled for every 4 days.\n    PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the snapshot schedule."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `retention_time`:(string)"Duration to keep the snapshots on the source array.\nOnce this period expires, system deletes the snapshot automatically from the source array.\nExamples:\nP200D,  200 days.\nPT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds."
* `snapshot_expiry_time`:(string)"Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period."
