---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_host"
sidebar_current: "docs-intersight-data-source-virtualizationVmwareHost"
description: |-
The VMware Host entity with its attributes. Every Host belongs to a Datacenter and probably runs VMs.
---

# Data Source: intersight_virtualization_vmware_host
The VMware Host entity with its attributes. Every Host belongs to a Datacenter and probably runs VMs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `boot_time`:(string)"It is the time when this host booted up."
* `connection_state`:(string)"Is host connected to vCenter. Values are connected, notconnected."
* `hw_power_state`:(string)"Is the host powered-up or powered-down."
* `hypervisor_type`:(string)"Identifies the broad type of the underlying hypervisor."
* `identity`:(string)"The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference)."
* `maintenance_mode`:(bool)"Is this host in maintenance mode. Set to true or false."
* `model`:(string)"Commercial model information about this hardware."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations."
* `network_adapter_count`:(int)"It is the count of all network adapters attached to this host."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `serial`:(string)"Serial number of this host (internally generated)."
* `status`:(string)"Host health status, as reported by the hypervisor platform."
* `storage_adapter_count`:(int)"It is the count of all storage adapters attached to this host."
* `uuid`:(string)"Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated."
* `up_time`:(string)"The uptime of the host, stored as Duration (from w3c)."
* `vcenter_host_id`:(string)"The identity of this host within vCenter (optional)."
* `vendor`:(string)"Commercial vendor details of this hardware."
