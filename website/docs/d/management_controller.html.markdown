---
layout: "intersight"
page_title: "Intersight: intersight_management_controller"
sidebar_current: "docs-intersight-data-source-managementController"
description: |-
A specialized service processor that monitors the physical state of a server, using sensors and communicating with the system administrator through an independent connection.

---

# Data Source: intersight_management_controller
A specialized service processor that monitors the physical state of a server, using sensors and communicating with the system administrator through an independent connection.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `model`:(string)Model of the endpoint that houses the management controller.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
