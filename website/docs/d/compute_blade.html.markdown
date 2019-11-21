---
layout: "intersight"
page_title: "Intersight: intersight_compute_blade"
sidebar_current: "docs-intersight-data-source-computeBlade"
description: |-
Server which is housed in a chassis and shares some of the hardware with other servers in the chassis.

---

# Data Source: intersight_compute_blade
Server which is housed in a chassis and shares some of the hardware with other servers in the chassis.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_power_state`:(string)
* `asset_tag`:(string)
* `available_memory`:(int)
* `chassis_id`:(string)The id of the chassis that the blade is located in.
* `device_mo_id`:(string)
* `dn`:(string)
* `fault_summary`:(int)
* `memory_speed`:(string)
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_adaptors`:(int)
* `num_cpu_cores`:(int)
* `num_cpu_cores_enabled`:(int)
* `num_cpus`:(int)
* `num_eth_host_interfaces`:(int)
* `num_fc_host_interfaces`:(int)
* `num_threads`:(int)
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `oper_power_state`:(string)
* `oper_state`:(string)
* `operability`:(string)
* `platform_type`:(string)
* `presence`:(string)
* `revision`:(string)
* `rn`:(string)
* `scaled_mode`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `service_profile`:(string)
* `slot_id`:(int)
* `total_memory`:(int)
* `uuid`:(string)
* `user_label`:(string)
* `vendor`:(string)This field identifies the vendor of the given component.
