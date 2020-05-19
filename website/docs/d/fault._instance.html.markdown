---
layout: "intersight"
page_title: "Intersight: intersight_fault._instance"
sidebar_current: "docs-intersight-data-source-fault.Instance"
description: |-
An endpoint anomaly is represented by this object.
---

# Data Source: intersight_fault._instance
An endpoint anomaly is represented by this object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `acknowledged`:(string)
* `affected_dn`:(string)
* `affected_mo_id`:(string)"Managed object Id which was affected."
* `affected_mo_type`:(string)"Managed object type which was affected."
* `ancestor_mo_id`:(string)
* `ancestor_mo_type`:(string)
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `code`:(string)
* `creation_time`:(string)
* `description`:(string)"Short summary of the fault found."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `last_transition_time`:(string)
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `num_occurrences`:(int)
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `original_severity`:(string)
* `previous_severity`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `rule`:(string)
* `severity`:(string)"Severity of the fault found."
