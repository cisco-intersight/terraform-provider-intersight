
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_config_change_detail"
sidebar_current: "docs-intersight-data-source-fabric-config-change-detail"
description: |-
This captures details of configuration changes.
---

# Data Source: intersight_fabric._config_change_detail
This captures details of configuration changes.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `config_change_flag`:(string) Config change flag to differentiate Pending-changes and Config-drift. 
* `mod_status`:(string) Modification status of the mo in this config change. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
