---
layout: "intersight"
page_title: "Intersight: intersight_iam._system"
sidebar_current: "docs-intersight-data-source-iam.System"
description: |-
System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
---

# Data Source: intersight_iam._system
System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
