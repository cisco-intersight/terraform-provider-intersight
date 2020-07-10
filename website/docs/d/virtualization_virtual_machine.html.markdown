
---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_virtual_machine"
sidebar_current: "docs-intersight-data-source-virtualization-virtual-machine"
description: |-
Depicts operations to control the life cycle of a VM on a Hypervisor.
---

# Data Source: intersight_virtualization._virtual_machine
Depicts operations to control the life cycle of a VM on a Hypervisor.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc). 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cpu`:(int) Number of vCPUs allocated to virtual machine. 
* `discovered`:(bool) Flag to indicate whether the configuration is created from inventory object. 
* `guest_os`:(string) Guest operating system running on virtual machine. 
* `memory`:(int) Virtual machine memory defined in mega bytes. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Virtual machine name contains only letters, numbers, allowed special character and must be unique. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `power_state`:(string) Expected power state of virtual machine (PowerOn, PowerOff, Restart). 
