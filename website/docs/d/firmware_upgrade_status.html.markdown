
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade_status"
sidebar_current: "docs-intersight-data-source-firmware-upgrade-status"
description: |-
The status for the upgrade operation to include the status for the download and upgrade stages.
---

# Data Source: intersight_firmware._upgrade_status
The status for the upgrade operation to include the status for the download and upgrade stages.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `download_error`:(string) The error message from the endpoint during the download. 
* `download_percentage`:(int) The percentage of the image downloaded in the endpoint. 
* `download_progress`:(int) The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. 
* `download_retries`:(int) The number of retries the plugin attempted before succeeding or failing the download. 
* `download_stage`:(string) The image download stages. Example:downloading, flashing. 
* `ep_power_status`:(string) The server power status after the upgrade request is submitted in the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `overall_error`:(string) The reason for the operation failure. 
* `overall_percentage`:(int) The overall percentage of the operation. 
* `overallstatus`:(string) The overall status of the operation. 
* `pending_type`:(string) Pending reason for the upgrade waiting. 
