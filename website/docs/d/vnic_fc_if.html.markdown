---
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_if"
sidebar_current: "docs-intersight-data-source-vnicFcIf"
description: |-
Virtual Fibre Channel Interface.

---

# Data Source: intersight_vnic_fc_if
Virtual Fibre Channel Interface.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the virtual fibre channel interface.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `order`:(int)The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
* `persistent_bindings`:(bool)Enables retention of LUN ID associations in memory until they are manually cleared.
