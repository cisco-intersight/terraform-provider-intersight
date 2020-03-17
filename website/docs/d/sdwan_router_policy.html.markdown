---
layout: "intersight"
page_title: "Intersight: intersight_sdwan_router_policy"
sidebar_current: "docs-intersight-data-source-sdwanRouterPolicy"
description: |-
A policy specifying SD-WAN router details.
---

# Data Source: intersight_sdwan_router_policy
A policy specifying SD-WAN router details.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `deployment_size`:(string)"Scale of the SD-WAN router virtual machine deployment."
* `description`:(string)"Description of the policy."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `wan_count`:(int)"Number of WAN connections across the SD-WAN site."
* `wan_termination_type`:(string)"Defines if the WAN networks are singly or dually terminated. Dually terminated WANs are configured on all the SD-WAN routers. Singly terminated WANs are configured only on one of the SD-WAN routers."
