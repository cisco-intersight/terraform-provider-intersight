---
layout: "intersight"
page_title: "Intersight: intersight_pci_device"
sidebar_current: "docs-intersight-data-source-pciDevice"
description: |-
PCI device present in a server.

---

# Data Source: intersight_pci_device
PCI device present in a server.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)
* `firmware_version`:(string)It shows the running firmware version.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `pid`:(string)It shows the product identifier.
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `slot_id`:(string)It show PCI slot id of the device.
* `vendor`:(string)This field identifies the vendor of the given component.
