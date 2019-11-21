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
* `hostname`:(string)Cluster node's FQDN or IP address.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `node_id`:(int)System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `operational_status`:(string)Operational status of the Intersight Appliance node.
