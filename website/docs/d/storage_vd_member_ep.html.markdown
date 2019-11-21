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
* `dn`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `oper_qualifier_reason`:(string)
* `presence`:(string)
* `rn`:(string)
* `role`:(string)Role of the disk normal or hot-spare, used by virtual-drive.
* `span_id`:(string)
* `vd_member_ep_id`:(int)It shows local disk slot number as id.
