---
layout: "intersight"
page_title: "Intersight: intersight_compute_physical_summary"
sidebar_current: "docs-intersight-data-source-computePhysicalSummary"
description: |-
Consolidated view of Blades and RackUnits.
---

# Data Source: intersight_compute_physical_summary
Consolidated view of Blades and RackUnits.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_power_state`:(string)
* `asset_tag`:(string)
* `available_memory`:(int)
* `cpu_capacity`:(float)"CPU Capacity = Number of CPU Sockets x Enabled Cores x Speed (GHz)."
* `chassis_id`:(string)"The id of the chassis that the blade is located in."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `fault_summary`:(int)
* `firmware`:(string)"The firmware version of the Cisco Integrated Management Controller (CIMC) for this server."
* `ipv4_address`:(string)"The IPv4 address configured on the management interface of the Integrated Management Controller."
* `memory_speed`:(string)
* `mgmt_ip_address`:(string)"Management address of the server."
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC).\nWhen this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect.\nWhen this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller."
* `num_adaptors`:(int)"Total number of Adaptors available."
* `num_cpu_cores`:(int)"Total number of CPU cores available."
* `num_cpu_cores_enabled`:(int)"Number of CPU cores enabled."
* `num_cpus`:(int)"Total number of CPU's available."
* `num_eth_host_interfaces`:(int)"Number of Ethernet Host Interfaces."
* `num_fc_host_interfaces`:(int)"Number of Fibre channel Host Interfaces."
* `num_threads`:(int)"Number of threads enabled."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `oper_power_state`:(string)
* `oper_state`:(string)
* `operability`:(string)
* `platform_type`:(string)"Platform type of the device."
* `presence`:(string)
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `scaled_mode`:(string)
* `serial`:(string)"This field identifies the serial of the given component."
* `server_id`:(int)"The server id of the Rack server."
* `service_profile`:(string)"The service profile assigned."
* `slot_id`:(int)
* `source_object_type`:(string)"The source object type of this view MO."
* `total_memory`:(int)
* `uuid`:(string)
* `user_label`:(string)
* `vendor`:(string)"This field identifies the vendor of the given component."
