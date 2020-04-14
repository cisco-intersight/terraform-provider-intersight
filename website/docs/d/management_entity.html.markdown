---
layout: "intersight"
page_title: "Intersight: intersight_management_entity"
sidebar_current: "docs-intersight-data-source-managementEntity"
description: |-
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
---

# Data Source: intersight_management_entity
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `entity_id`:(string)"Identity of the Fabric Interconnect - A/B."
* `leadership`:(string)"Role (Primary / Subordinate) of the Fabric Interconnect."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
