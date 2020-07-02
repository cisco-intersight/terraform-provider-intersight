
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_fex_identity"
sidebar_current: "docs-intersight-data-source-equipment-fex-identity"
description: |-
FexIdentity Object conatains basic information of fabric extender. moduleId is uniquely allocated for the combination of vendor, model and serial number of the chassis.
---

# Data Source: intersight_equipment._fex_identity
FexIdentity Object conatains basic information of fabric extender. moduleId is uniquely allocated for the combination of vendor, model and serial number of the chassis.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_action`:(string) Updated by UI/API to trigger specific chassis action type. 
* `admin_action_state`:(string) The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `identifier`:(int) Numeric Identifier assigned by the management system to the equipment. 
* `lifecycle`:(string) The equipment's lifecycle status. 
* `model`:(string) The vendor provided model name for the equipment. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `serial`:(string) The serial number of the equipment. 
* `switch_id`:(int) Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B. 
* `vendor`:(string) The manufacturer of the equipment. 
