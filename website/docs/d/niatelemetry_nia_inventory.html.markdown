---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_inventory"
sidebar_current: "docs-intersight-data-source-niatelemetryNiaInventory"
description: |-

---

# Data Source: intersight_niatelemetry_nia_inventory

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cpu`:(float)"CPU usage of device being inventoried."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `crash_reset_logs`:(string)"Last crash reset reason of device being inventoried."
* `device_name`:(string)"Name of device being inventoried."
* `device_type`:(string)"Type of device being inventoried."
* `log_in_time`:(string)"Last log in time device being inventoried."
* `log_out_time`:(string)"Last log out time of device being inventoried."
* `memory`:(int)"Memory usage of device being inventoried."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `record_type`:(string)"Type of record DCNM / APIC / SE."
* `record_version`:(string)"Version of record being pushed."
* `serial`:(string)"Serial number of device being invetoried."
* `software_download`:(string)"Last software downloaded of device being inventoried."
* `version`:(string)"Software version of device being inventoried."
