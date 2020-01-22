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
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `fault_summary`:(int)
* `memory_speed`:(string)
* `mgmt_ip_address`:(string)Management address of the server.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_adaptors`:(int)Total number of Adaptors available.
* `num_cpu_cores`:(int)
* `num_cpu_cores_enabled`:(int)Number of CPU cores enabled.
* `num_cpus`:(int)Total number of CPU's available.
* `num_eth_host_interfaces`:(int)Number of Ethernet Host Interfaces.
* `num_fc_host_interfaces`:(int)
* `num_threads`:(int)Number of threads enabled.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `oper_power_state`:(string)
* `oper_state`:(string)
* `operability`:(string)
* `platform_type`:(string)Platform type of the device.
* `presence`:(string)
* `revision`:(string)
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `scaled_mode`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `service_profile`:(string)
* `slot_id`:(int)
* `total_memory`:(int)
* `uuid`:(string)
* `user_label`:(string)
* `vendor`:(string)This field identifies the vendor of the given component.
