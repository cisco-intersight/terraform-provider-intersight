---
layout: "intersight"
page_title: "Intersight: intersight_appliance_system_info"
sidebar_current: "docs-intersight-data-source-applianceSystemInfo"
description: |-
The Intersight Appliance's system information. SystemInfo is a singleton managed object
created during the Intersight Appliance setup. The Intersight Appliance updates the
SystemInfo managed object with up to date cluster status information periodically.

---

# Data Source: intersight_appliance_system_info
The Intersight Appliance's system information. SystemInfo is a singleton managed object
created during the Intersight Appliance setup. The Intersight Appliance updates the
SystemInfo managed object with up to date cluster status information periodically.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cloud_conn_status`:(string)Connection state of the Intersight Appliance to the Intersight.
* `hostname`:(string)Publicly accessible FQDN or IP address of the Intersight Appliance.
* `init_done`:(bool)Indicates that the setup initialization process has been completed.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `operational_status`:(string)Operational status of the Intersight Appliance cluster.
* `serial_id`:(string)SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance.
* `time_zone`:(string)Timezone of the Intersight Appliance.
* `version`:(string)Current software version of the Intersight Appliance.
