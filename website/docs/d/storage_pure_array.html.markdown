---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_array"
sidebar_current: "docs-intersight-data-source-storagePureArray"
description: |-
The details of the Pure storage array.
---

# Data Source: intersight_storage_pure_array
The details of the Pure storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Administrator defined name for the device."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `serial`:(string)"This field identifies the serial of the given component."
* `uuid`:(string)"Unique identity of the device."
* `vendor`:(string)"This field identifies the vendor of the given component."
* `version`:(string)"Current running software version of the device."
