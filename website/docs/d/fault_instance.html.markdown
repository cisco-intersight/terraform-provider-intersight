---
layout: "intersight"
page_title: "Intersight: intersight_fault_instance"
sidebar_current: "docs-intersight-data-source-faultInstance"
description: |-
An endpoint anomaly is represented by this object.

---

# Data Source: intersight_fault_instance
An endpoint anomaly is represented by this object.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `acknowledged`:(string)
* `affected_dn`:(string)
* `affected_mo_id`:(string)
* `affected_mo_type`:(string)
* `ancestor_mo_id`:(string)
* `ancestor_mo_type`:(string)
* `code`:(string)
* `creation_time`:(string)
* `description`:(string)
* `device_mo_id`:(string)
* `dn`:(string)
* `last_transition_time`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `num_occurrences`:(int)
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `original_severity`:(string)
* `previous_severity`:(string)
* `rn`:(string)
* `rule`:(string)
* `severity`:(string)
