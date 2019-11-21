---
layout: "intersight"
page_title: "Intersight: intersight_compute_board"
sidebar_current: "docs-intersight-data-source-computeBoard"
description: |-
Mother board of a server.

---

# Data Source: intersight_compute_board
Mother board of a server.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `board_id`:(int)
* `cpu_type_controller`:(string)
* `device_mo_id`:(string)
* `dn`:(string)
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `oper_power_state`:(string)
* `presence`:(string)
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `vendor`:(string)This field identifies the vendor of the given component.
