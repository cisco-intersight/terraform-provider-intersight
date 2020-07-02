
---
layout: "intersight"
page_title: "Intersight: intersight_compute_blade_identity"
sidebar_current: "docs-intersight-data-source-compute-blade-identity"
description: |-
Identity object that uniquely represents a blade server object under a DR.
---

# Data Source: intersight_compute._blade_identity
Identity object that uniquely represents a blade server object under a DR.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_action`:(string) Updated by UI/API to trigger specific chassis action type. 
* `admin_action_state`:(string) The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. 
* `chassis_id`:(int) Chassis Identifier of a blade server. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) FI Device registration Mo ID. 
* `identifier`:(int) Numeric Identifier assigned by the management system to the equipment. 
* `lifecycle`:(string) The equipment's lifecycle status. 
* `model`:(string) The vendor provided model name for the equipment. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pending_discovery`:(string) Indicates pending discovery flag. 
* `serial`:(string) The serial number of the equipment. 
* `slot_id`:(int) Chassis slot number of a blade server. 
* `vendor`:(string) The manufacturer of the equipment. 
