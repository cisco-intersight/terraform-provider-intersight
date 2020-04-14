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
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `cloud_conn_status`:(string)"Connection state of the Intersight Appliance to the Intersight."
* `deployment_size`:(string)"Current running deployment size for the Intersight Appliance cluster. Eg. small, medium, large etc."
* `hostname`:(string)"Publicly accessible FQDN or IP address of the Intersight Appliance."
* `init_done`:(bool)"Indicates that the setup initialization process has been completed."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `operational_status`:(string)"Operational status of the Intersight Appliance cluster."
* `serial_id`:(string)"SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance."
* `time_zone`:(string)"Timezone of the Intersight Appliance."
* `version`:(string)"Current software version of the Intersight Appliance."
