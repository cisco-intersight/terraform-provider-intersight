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
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `controller_flags`:(string)
* `controller_id`:(string)"It shows the Id of controller."
* `controller_status`:(string)"It shows the current status of controller."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `hw_revision`:(string)"It shows the hardware revision of controller."
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `oob_interface_supported`:(string)"It shows CIMC support for out-of-band configuration of controller."
* `oper_state`:(string)"It shows the current operational state of controller."
* `operability`:(string)
* `pci_addr`:(string)"It shows the current pci address of controller."
* `pci_slot`:(string)"It shows the pci slot name for the controller."
* `presence`:(string)"It shows physical presence or absence of the controller on server."
* `raid_support`:(string)"It shows the RAID levels supported by controller."
* `rebuild_rate`:(string)
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `self_encrypt_enabled`:(string)
* `serial`:(string)"This field identifies the serial of the given component."
* `type`:(string)"Controller types are SAS, SATA, PCH, NVME."
* `vendor`:(string)"This field identifies the vendor of the given component."
