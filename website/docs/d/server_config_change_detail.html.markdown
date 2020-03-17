---
layout: "intersight"
page_title: "Intersight: intersight_server_config_change_detail"
sidebar_current: "docs-intersight-data-source-serverConfigChangeDetail"
description: |-
The configuration change details are captured here.
---

# Data Source: intersight_server_config_change_detail
The configuration change details are captured here.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `config_change_flag`:(string)"Config change flag to differentiate Pending-changes and Config-drift."
* `message`:(string)"Detailed description of the config change."
* `mod_status`:(string)"Modification status of the mo in this config change."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
