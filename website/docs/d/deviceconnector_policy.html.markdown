---
layout: "intersight"
page_title: "Intersight: intersight_deviceconnector_policy"
sidebar_current: "docs-intersight-data-source-deviceconnectorPolicy"
description: |-
Policy to control configuration changes allowed from Cisco IMC.
---

# Data Source: intersight_deviceconnector_policy
Policy to control configuration changes allowed from Cisco IMC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)"Description of the policy."
* `lockout_enabled`:(bool)"Enables configuration lockout on the endpoint."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
