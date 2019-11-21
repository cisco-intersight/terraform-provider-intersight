---
layout: "intersight"
page_title: "Intersight: intersight_pci_link"
sidebar_current: "docs-intersight-data-source-pciLink"
description: |-
PCI Switch Link connected to PCIe Switch.

---

# Data Source: intersight_pci_link
PCI Switch Link connected to PCIe Switch.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `adapter`:(string)It shows the name of the pci device.
* `device_mo_id`:(string)
* `dn`:(string)
* `link_speed`:(string)It shows the upstream link speed for device.
* `link_status`:(string)It shows the upstream link status for device.
* `link_width`:(string)It shows the upstream link width for device.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `pci_slot`:(string)It shows pci slot name for the pci device.
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `slot_status`:(string)It shows the health information for pci device.
* `vendor`:(string)This field identifies the vendor of the given component.
