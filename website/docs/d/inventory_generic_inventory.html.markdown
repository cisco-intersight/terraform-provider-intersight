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
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `key`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `type`:(string)
* `value`:(string)
