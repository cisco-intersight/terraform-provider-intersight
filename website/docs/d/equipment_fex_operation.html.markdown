
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_fex_operation"
sidebar_current: "docs-intersight-data-source-equipment-fex-operation"
description: |-
Models the configuration states of a FEX in Intersight.
---

# Data Source: intersight_equipment._fex_operation
Models the configuration states of a FEX in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_locator_led_action`:(string) Action performed on the locator LED for a FEX. 
* `admin_locator_led_action_state`:(string) Defines status of action performed on AdminLocatorLedState. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
