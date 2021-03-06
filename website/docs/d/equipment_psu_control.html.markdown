
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_psu_control"
sidebar_current: "docs-intersight-data-source-equipment-psu-control"
description: |-
This represents the power states of an equipment.
---

# Data Source: intersight_equipment._psu_control
This represents the power states of an equipment.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cluster_state`:(string) This field identifies the cluster state of the psu redundancy. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `input_power_state`:(string) This field identifies the input power state of the psus. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This field identifies the name of psu control object. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_qualifier`:(string) This field identifies the operational qualifier for the psu redundancy. 
* `oper_state`:(string) This field identifies the operational state of the psu redundancy. 
* `output_power_state`:(string) This field identifies the output power state of the psus. 
* `redundancy`:(string) This field identifies the redundancy state of the psus. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
