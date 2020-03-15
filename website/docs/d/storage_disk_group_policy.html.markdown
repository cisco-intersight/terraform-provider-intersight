---
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group_policy"
sidebar_current: "docs-intersight-data-source-storageDiskGroupPolicy"
description: |-
A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
---

# Data Source: intersight_storage_disk_group_policy
A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `raid_level`:(string)The supported RAID level for the disk group.
* `use_jbods`:(bool)Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation.
