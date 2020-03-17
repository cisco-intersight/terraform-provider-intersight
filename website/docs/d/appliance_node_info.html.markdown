---
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_info"
sidebar_current: "docs-intersight-data-source-applianceNodeInfo"
description: |-
NodeInfo managed object stores the Intersight Appliance's cluster node information.
NodeInfo managed objects are created during the Intersight Appliance setup. The
Intersight Appliance updates the NodeInfo managed objects with status information
periodically.
---

# Data Source: intersight_appliance_node_info
NodeInfo managed object stores the Intersight Appliance's cluster node information.
NodeInfo managed objects are created during the Intersight Appliance setup. The
Intersight Appliance updates the NodeInfo managed objects with status information
periodically.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hostname`:(string)"Cluster node's FQDN or IP address."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `node_id`:(int)"System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `operational_status`:(string)"Operational status of the Intersight Appliance node."
