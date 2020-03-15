---
layout: "intersight"
page_title: "Intersight: intersight_storage_vd_member_ep"
sidebar_current: "docs-intersight-data-source-storageVdMemberEp"
description: |-
Reference to LocalDisk to build up a VirtualDrive.
---

# Data Source: intersight_storage_vd_member_ep
Reference to LocalDisk to build up a VirtualDrive.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `oper_qualifier_reason`:(string)
* `presence`:(string)
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `role`:(string)Role of the disk normal or hot-spare, used by virtual-drive.
* `span_id`:(string)
* `vd_member_ep_id`:(int)It shows local disk slot number as id.
