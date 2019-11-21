---
layout: "intersight"
page_title: "Intersight: intersight_compute_server_setting"
sidebar_current: "docs-intersight-data-source-computeServerSetting"
description: |-
Models the configurable properties of a server in Intersight.

---

# Data Source: intersight_compute_server_setting
Models the configurable properties of a server in Intersight.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_locator_led_state`:(string)User configured state of the locator LED for the server.
* `admin_power_state`:(string)User configured power state of the server.
* `config_state`:(string)The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server.
* `device_mo_id`:(string)
* `dn`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `one_time_boot_device`:(string)The name of the device chosen by user for configuring One-Time Boot device.
* `rn`:(string)
