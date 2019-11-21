---
layout: "intersight"
page_title: "Intersight: intersight_storage_controller"
sidebar_current: "docs-intersight-data-source-storageController"
description: |-
Storage Controller present in a server.

---

# Data Source: intersight_storage_controller
Storage Controller present in a server.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `controller_flags`:(string)
* `controller_id`:(string)It shows the Id of controller.
* `controller_status`:(string)It shows the current status of controller.
* `device_mo_id`:(string)
* `dn`:(string)
* `hw_revision`:(string)It shows the hardware revision of controller.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `oob_interface_supported`:(string)It shows CIMC support for out-of-band configuration of controller.
* `oper_state`:(string)
* `operability`:(string)
* `pci_addr`:(string)It shows the current pci address of controller.
* `pci_slot`:(string)It shows the pci slot name for the controller.
* `presence`:(string)It shows physical presence or absence of the controller on server.
* `raid_support`:(string)It shows the RAID levels supported by controller.
* `rebuild_rate`:(string)
* `revision`:(string)
* `rn`:(string)
* `self_encrypt_enabled`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `type`:(string)Controller types are SAS, SATA, PCH, NVME.
* `vendor`:(string)This field identifies the vendor of the given component.
