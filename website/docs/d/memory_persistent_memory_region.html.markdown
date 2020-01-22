---
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_region"
sidebar_current: "docs-intersight-data-source-memoryPersistentMemoryRegion"
description: |-
This represents a Persistent Memory Region configured on the persistent memory modules on a server.

---

# Data Source: intersight_memory_persistent_memory_region
This represents a Persistent Memory Region configured on the persistent memory modules on a server.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dimm_locater_ids`:(string)This represents a set of DIMM locator IDs that are included in the Persistent Memory Region.
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `free_capacity`:(string)This represents the free capacity in GB of a Persistent Memory Region.
* `health_state`:(string)This represents the health state of a Persistent Memory Region.
* `interleaved_set_id`:(string)This represents the ID of a Interleaved Set formed for a Persistent Memory Region.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `persistent_memory_type`:(string)This represents the persistent memory type of a Persistent Memory Region.
* `region_id`:(string)This represents the ID of a Persistent Memory Region.
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `socket_id`:(string)This represents the Socket ID of a Persistent Memory Region.
* `socket_memory_id`:(string)This represents the Socket Memory ID of a Persistent Memory Region.
* `total_capacity`:(string)This represents the total capacity in GB of a Persistent Memory Region.
