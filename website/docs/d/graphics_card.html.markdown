---
layout: "intersight"
page_title: "Intersight: intersight_graphics_card"
sidebar_current: "docs-intersight-data-source-graphicsCard"
description: |-
Graphics Card present in a server.

---

# Data Source: intersight_graphics_card
Graphics Card present in a server.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `card_id`:(int)It shows the id of graphics card.
* `device_id`:(int)It shows the device id of grphics card.
* `device_mo_id`:(string)
* `dn`:(string)
* `expander_slot`:(string)It shows the expander slot inforamtion for the card.
* `firmware_version`:(string)It shows current firmware version of graphics card.
* `mode`:(string)It shows the current mode of graphics card.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_gpus`:(string)It shows number of controllers under each card.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `oper_state`:(string)It shows the current operational state of graphics card.
* `pci_address`:(string)It shows the pci address of graphics card.
* `pci_address_list`:(string)This list contains the pci address of all controllers for corresponding card.
* `pci_slot`:(string)It shows the pci slot name for grapchics card.
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `vendor`:(string)This field identifies the vendor of the given component.
