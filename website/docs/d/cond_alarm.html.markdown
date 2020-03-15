---
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm"
sidebar_current: "docs-intersight-data-source-condAlarm"
description: |-
A state-full entity representing a found problem. Alarms can be reported by the managed system itself or can be determined by Intersight.
---

# Data Source: intersight_cond_alarm
A state-full entity representing a found problem. Alarms can be reported by the managed system itself or can be determined by Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `affected_mo_id`:(string)MoId of the affected object from the managed system's point of view.
* `affected_mo_type`:(string)Managed system affected object type. For example Adaptor, FI, CIMC.
* `affected_object`:(string)A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject.
* `ancestor_mo_id`:(string)Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault.
* `ancestor_mo_type`:(string)Parent MO type of the fault from managed system. For example, Blade for adaptor fault.
* `code`:(string)A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code.
* `creation_time`:(string)The time the alarm was created.
* `description`:(string)A longer description of the alarm than the name. The description contains details of the component reporting the issue.
* `last_transition_time`:(string)The time the alarm last had a change in severity.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `ms_affected_object`:(string)A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn.
* `name`:(string)Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `orig_severity`:(string)The original severity when the alarm was first created.
* `severity`:(string)The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared.
