---
layout: "intersight"
page_title: "Intersight: intersight_inventory_generic_inventory"
sidebar_current: "docs-intersight-data-source-inventoryGenericInventory"
description: |-
Any inventory which is represented as a key / value pair. Example - moInvKv in UCSM representing OS tools running on ESX.

---

# Data Source: intersight_inventory_generic_inventory
Any inventory which is represented as a key / value pair. Example - moInvKv in UCSM representing OS tools running on ESX.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)
* `key`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rn`:(string)
* `type`:(string)
* `value`:(string)
