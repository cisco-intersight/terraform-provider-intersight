---
layout: "intersight"
page_title: "Intersight: intersight_inventory_dn_mo_binding"
sidebar_current: "docs-intersight-data-source-inventoryDnMoBinding"
description: |-
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.

---

# Data Source: intersight_inventory_dn_mo_binding
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dn`:(string)Distinguished Name of the UCSM object.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `target_mo_id`:(string)The property targetMoId represents the Binding target's MoId.
* `target_mo_type`:(string)The property targetMoType represents the Binding target's Mo type.
