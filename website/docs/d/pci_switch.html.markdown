---
layout: "intersight"
page_title: "Intersight: intersight_pci_switch"
sidebar_current: "docs-intersight-data-source-pciSwitch"
description: |-
PCI Switch present in a server connected to two GPUs and one PCIe adapter.

---

# Data Source: intersight_pci_switch
PCI Switch present in a server connected to two GPUs and one PCIe adapter.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_id`:(string)It shows the device id of the switch.
* `device_mo_id`:(string)
* `dn`:(string)
* `health`:(string)It shows the composite health of the switch.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_of_adaptors`:(string)It shows the number of gpus and pci adapters connected the switch.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `pci_address`:(string)It shows shows the PCI address of switch.
* `pci_slot`:(string)It shows the PCI slot name for switch.
* `product_name`:(string)It shows the model information for the switch.
* `product_revision`:(string)It shows the revision for the product.
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `sub_device_id`:(string)It shows the sub device id of the switch.
* `sub_vendor_id`:(string)It shows the sub vendor id of the switch.
* `temperature`:(string)It shows the current temperature of the switch.
* `type`:(string)It shows the type inforamtion of switch.
* `vendor`:(string)This field identifies the vendor of the given component.
* `vendor_id`:(string)It shows the vendor id of the switch.
