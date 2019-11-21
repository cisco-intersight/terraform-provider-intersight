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
* `cpu_capacity`:(float)CPU Capacity = Number of CPU Sockets x Enabled Cores x Speed (GHz).
* `chassis_id`:(string)The id of the chassis that the blade is located in.
* `device_mo_id`:(string)
* `dn`:(string)
* `fault_summary`:(int)
* `firmware`:(string)The firmware version of the Cisco Integrated Management Controller (CIMC) for this server.
* `ipv4_address`:(string)The IPv4 address configured on the management interface of the Integrated Management Controller.
* `memory_speed`:(string)
* `mgmt_ip_address`:(string)The IP address of the management interface on the UCS Fabric Interconnect or Cisco Integrated Management Controller.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC).When this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect.When this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller.
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
* `server_id`:(int)
* `service_profile`:(string)
* `slot_id`:(int)
* `source_object_type`:(string)Specifies the source object type for View MO.
* `total_memory`:(int)
* `uuid`:(string)
* `user_label`:(string)
* `vendor`:(string)This field identifies the vendor of the given component.
