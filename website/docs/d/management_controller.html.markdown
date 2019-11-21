---
layout: "intersight"
page_title: "Intersight: intersight_management_controller"
sidebar_current: "docs-intersight-data-source-managementController"
description: |-
Provides management controls access to an endpoint.

---

# Data Source: intersight_management_controller
Provides management controls access to an endpoint.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)
* `model`:(string)Model of the endpoint that houses the management controller.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rn`:(string)
