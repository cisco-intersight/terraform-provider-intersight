---
layout: "intersight"
page_title: "Intersight: intersight_compute_rack_unit"
sidebar_current: "docs-intersight-data-source-computeRackUnit"
description: |-
Rack-mounted server.
---

# Data Source: intersight_compute_rack_unit
Rack-mounted server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_power_state`:(string)"Desired power state of the server."
* `asset_tag`:(string)
* `available_memory`:(int)"The actual amount of memory currently available to the server."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `fault_summary`:(int)
* `memory_speed`:(string)"The memory speed, in megahertz."
* `mgmt_ip_address`:(string)"Management address of the server."
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `num_adaptors`:(int)"Total number of Adaptors available."
* `num_cpu_cores`:(int)"Total number of CPU cores available."
* `num_cpu_cores_enabled`:(int)"Number of CPU cores enabled."
* `num_cpus`:(int)"Total number of CPU's available."
* `num_eth_host_interfaces`:(int)"Number of Ethernet Host Interfaces."
* `num_fc_host_interfaces`:(int)"Number of Fibre channel Host Interfaces."
* `num_threads`:(int)"Number of threads enabled."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `oper_power_state`:(string)"Current power state of the server."
* `oper_state`:(string)
* `operability`:(string)
* `platform_type`:(string)"Platform type of the device."
* `presence`:(string)"Identifies the presence of the server."
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `serial`:(string)"This field identifies the serial of the given component."
* `server_id`:(int)"The server id of the Rack server."
* `service_profile`:(string)"The service profile assigned."
* `total_memory`:(int)"The total amount of memory installed on the server."
* `uuid`:(string)
* `user_label`:(string)
* `vendor`:(string)"This field identifies the vendor of the given component."
