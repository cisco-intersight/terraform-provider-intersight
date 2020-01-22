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
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `package_version`:(string)Package version which the firmware belongs to.
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `type`:(string)Type of the firmware.
* `version`:(string)Version of the firmware.
