---
layout: "intersight"
page_title: "Intersight: intersight_storage_virtual_drive_extension"
sidebar_current: "docs-intersight-data-source-storageVirtualDriveExtension"
description: |-
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.

---

# Data Source: intersight_storage_virtual_drive_extension
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bootable`:(string)
* `container_id`:(int)
* `device_mo_id`:(string)
* `dn`:(string)
* `drive_state`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `oper_device_id`:(string)
* `rn`:(string)
* `uuid`:(string)
* `vendor_uuid`:(string)
* `virtual_drive_dn`:(string)
* `virtual_drive_id`:(string)
