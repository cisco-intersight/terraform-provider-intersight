
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_virtual_machine"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-virtual-machine"
description: |-
The Virtual machine that runs on a Hyperflex Application platform compute host.
---

# Data Source: intersight_hyperflex._hxap_virtual_machine
The Virtual machine that runs on a Hyperflex Application platform compute host.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `age`:(string) Denotes age or life time of the VM in nano seconds. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `failure_reason`:(string) Reason of the failure when VM is in failed state. 
* `graphic_console_url`:(string) Graphical console URL of this VM. 
* `hypervisor_type`:(string) Type of hypervisor where the virtual machine is hosted for example ESXi. 
* `identity`:(string) The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the virtual machine. 
* `network_count`:(int) Number network interfaces associated with the virtual machine. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `power_state`:(string) Power state of the virtual machine. 
* `start_time`:(string) Denotes the VM start timestamp. 
* `status`:(string) Status of virtual machine. 
* `up_time`:(string) Denotes how long this VM has been running in nano seconds. 
* `uuid`:(string) The uuid of this virtual machine. The uuid is internally generated and not user specified. 
