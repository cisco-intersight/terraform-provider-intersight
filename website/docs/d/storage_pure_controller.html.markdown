---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_controller"
sidebar_current: "docs-intersight-data-source-storagePureController"
description: |-
A storage controller entity in Pure FlashArray.
---

# Data Source: intersight_storage_pure_controller
A storage controller entity in Pure FlashArray.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Storage array controller name.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `operational_mode`:(string)Controller running mode, Primary or Secondary.
* `revision`:(string)
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `serial`:(string)This field identifies the serial of the given component.
* `status`:(string)Status of the storage controller.
* `vendor`:(string)This field identifies the vendor of the given component.
* `version`:(string)Software version running on a storage controller.
