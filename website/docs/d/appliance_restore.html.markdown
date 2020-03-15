---
layout: "intersight"
page_title: "Intersight: intersight_appliance_restore"
sidebar_current: "docs-intersight-data-source-applianceRestore"
description: |-
Restore tracks requests to restore the Intersight Appliance. There will be only
one Restore managed object with a 'Started' state at any time. All other Restore
managed objects will be in terminal states.
---

# Data Source: intersight_appliance_restore
Restore tracks requests to restore the Intersight Appliance. There will be only
one Restore managed object with a 'Started' state at any time. All other Restore
managed objects will be in terminal states.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `elapsed_time`:(int)Elapsed time in seconds since the restore process has started.
* `end_time`:(string)End date and time of the restore process.
* `filename`:(string)Backup filename to backup or restore.
* `is_password_set`:(bool)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `password`:(string)Password for authenticating with the file server.
* `protocol`:(string)Communication protocol used by the file server (e.g. scp or sftp).
* `remote_host`:(string)Hostname of the remote file server.
* `remote_path`:(string)File server directory to copy the file.
* `remote_port`:(int)Remote TCP port on the file server (e.g. 22 for scp).
* `start_time`:(string)Start date and time of the restore process.
* `status`:(string)Status of the restore managed object.
* `username`:(string)Username to authenticate the fileserver.
