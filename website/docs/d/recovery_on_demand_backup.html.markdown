---
layout: "intersight"
page_title: "Intersight: intersight_recovery_on_demand_backup"
sidebar_current: "docs-intersight-data-source-recoveryOnDemandBackup"
description: |-
Handles requests for on demand backup for a given endpoint.
---

# Data Source: intersight_recovery_on_demand_backup
Handles requests for on demand backup for a given endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)"Description of the policy."
* `file_name_prefix`:(string)"The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418."
* `is_password_set`:(bool)
* `location_type`:(string)"Specifies whether the backup will be stored locally or remotely."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `password`:(string)"Backup server password."
* `path`:(string)"The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/"
* `protocol`:(string)"Protocol for transferring the backup image to the network share location."
* `retention_count`:(int)"Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner."
* `user_name`:(string)"Backup server user name."
