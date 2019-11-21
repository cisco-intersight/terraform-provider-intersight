---
layout: "intersight"
page_title: "Intersight: intersight_inventory_device_info"
sidebar_current: "docs-intersight-data-source-inventoryDeviceInfo"
description: |-
Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.

---

# Data Source: intersight_inventory_device_info
Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `event_counter`:(int)Indicates the index of the last received event.
* `event_counter_enabled`:(bool)Indicates whether the event counter enabled for the device.
* `interval`:(int)The time interval (in hours) between job runs.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
