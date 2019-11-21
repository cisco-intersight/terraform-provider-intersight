---
layout: "intersight"
page_title: "Intersight: intersight_equipment_device_summary"
sidebar_current: "docs-intersight-data-source-equipmentDeviceSummary"
description: |-
Aggregation of properties pertaining to different inventory MOs.

---

# Data Source: intersight_equipment_device_summary
Aggregation of properties pertaining to different inventory MOs.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dn`:(string)The distinguished name for the Network Element.
* `model`:(string)The model information of the Network Element.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `serial`:(string)The serial number for the Network Element.
* `source_object_type`:(string)Specifies the source object type for View MO.
