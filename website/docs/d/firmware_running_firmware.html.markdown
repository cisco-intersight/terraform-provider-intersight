---
layout: "intersight"
page_title: "Intersight: intersight_firmware_running_firmware"
sidebar_current: "docs-intersight-data-source-firmwareRunningFirmware"
description: |-
Running Firmware on an endpoint.

---

# Data Source: intersight_firmware_running_firmware
Running Firmware on an endpoint.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `component`:(string)Kind of the firmware - boot-booloader/system/kernel.
* `device_mo_id`:(string)
* `dn`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `package_version`:(string)Package version which the firmware belongs to.
* `rn`:(string)
* `type`:(string)Type of the firmware.
* `version`:(string)Version of the firmware.
