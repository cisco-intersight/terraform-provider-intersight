---
layout: "intersight"
page_title: "Intersight: intersight_storage_enclosure"
sidebar_current: "docs-intersight-data-source-storageEnclosure"
description: |-
Storage Enclosure for physical disks.
---

# Data Source: intersight_storage_enclosure
Storage Enclosure for physical disks.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chassis_id`:(int)
* `description`:(string)
* `device_mo_id`:(string)
* `dn`:(string)"The Distinguished Name unambiguously identifies an object in the system."
* `enclosure_id`:(int)
* `model`:(string)"This field identifies the model of the given component."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `num_slots`:(int)
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `presence`:(string)
* `revision`:(string)
* `rn`:(string)"The Relative Name uniquely identifies an object within a given context."
* `serial`:(string)"This field identifies the serial of the given component."
* `server_id`:(int)
* `type`:(string)
* `vendor`:(string)"This field identifies the vendor of the given component."
