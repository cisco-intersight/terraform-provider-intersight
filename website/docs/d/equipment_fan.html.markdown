---
layout: "intersight"
page_title: "Intersight: intersight_equipment_fan"
sidebar_current: "docs-intersight-data-source-equipmentFan"
description: |-
Fan in a Fabric Interconnect / Chassis / RackUnit.

---

# Data Source: intersight_equipment_fan
Fan in a Fabric Interconnect / Chassis / RackUnit.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)This field is to provide description for the fan.
* `device_mo_id`:(string)
* `dn`:(string)
* `fan_id`:(int)This field acts as the identifier for this particular Fan, within the Fabric Interconnect.
* `fan_module_id`:(int)This field is used to identify the Fan Module to which this Fan belongs.
* `model`:(string)This field identifies the model of the given component.
* `module_id`:(int)Fan module Identifier for the fan.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `oper_state`:(string)
* `part_number`:(string)This field identifies the Part Number for this Fan Unit.
* `pid`:(string)This field identifies the Product ID for the fans.
* `presence`:(string)This field is used to indicate this fan unit's presence.
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `sku`:(string)This field identifies the Stockkeeping Unit for this Fan Unit.
* `tray_id`:(int)Tray identifier for the fan module.
* `vendor`:(string)This field identifies the vendor of the given component.
* `vid`:(string)This field identifies the Vendor ID for this Fan Unit.
