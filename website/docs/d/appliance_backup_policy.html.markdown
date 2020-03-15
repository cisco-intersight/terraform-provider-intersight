---
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup_policy"
sidebar_current: "docs-intersight-data-source-applianceBackupPolicy"
description: |-
BackupPolicy stores the Intersight Appliance's backup policy. There will be only
one BackupPolicy managed object in the Intersight Appliance. Default backup policy
managed object is created during the Intersight Appliance setup, and it is configured
in the manual backup mode.
---

# Data Source: intersight_appliance_backup_policy
BackupPolicy stores the Intersight Appliance's backup policy. There will be only
one BackupPolicy managed object in the Intersight Appliance. Default backup policy
managed object is created during the Intersight Appliance setup, and it is configured
in the manual backup mode.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `backup_time`:(string)The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field.
* `filename`:(string)Backup filename to backup or restore.
* `is_password_set`:(bool)
* `manual_backup`:(bool)Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `password`:(string)Password to authenticate the file server.
* `protocol`:(string)Communication protocol used by the file server (e.g. scp or sftp).
* `remote_host`:(string)Hostname of the remote file server.
* `remote_path`:(string)File server directory to copy the file.
* `remote_port`:(int)Remote TCP port on the file server (e.g. 22 for scp).
* `username`:(string)Username to authenticate the fileserver.
