---
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade_status"
sidebar_current: "docs-intersight-data-source-firmwareUpgradeStatus"
description: |-
Status of the upgrade operation includes the status of download and upgrade stages.
---

# Data Source: intersight_firmware_upgrade_status
Status of the upgrade operation includes the status of download and upgrade stages.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `download_error`:(string)"The error message from the endpoint during the download."
* `download_percentage`:(int)"The percentage of the image downloaded in the endpoint."
* `download_stage`:(string)"The image download stages. Example:downloading, flashing."
* `download_status`:(string)"The download status of the image in the endpoint."
* `ep_power_status`:(string)"The server power status after the upgrade request is submitted in the endpoint."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `overall_error`:(string)"The reason for the operation failure."
* `overall_percentage`:(int)"The overall percentage of the operation."
* `overallstatus`:(string)"The overall status of the operation."
* `pending_type`:(string)"Pending reason for the upgrade waiting."
