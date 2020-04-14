---
layout: "intersight"
page_title: "Intersight: intersight_asset_cluster_member"
sidebar_current: "docs-intersight-data-source-assetClusterMember"
description: |-
A node within a cluster of device connectors. A Device Registration may contain multiple ClusterMembers with each holding the connection details of the device connector as well as the nodes current leadership within the cluster.
---

# Data Source: intersight_asset_cluster_member
A node within a cluster of device connectors. A Device Registration may contain multiple ClusterMembers with each holding the connection details of the device connector as well as the nodes current leadership within the cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `api_version`:(int)"The version of the connector API, describes the capability of the connector's framework.\nIf the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded."
* `app_partition_number`:(int)"The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `connection_id`:(string)"The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter."
* `connection_reason`:(string)"If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with the cloud."
* `connection_status`:(string)"The status of the persistent connection between the device connector and Intersight."
* `connection_status_last_change_time`:(string)"The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between the cloud and device connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with the cloud."
* `connector_version`:(string)"The version of the device connector running on the managed device."
* `device_external_ip_address`:(string)"The IP Address of the managed device as seen from the cloud at the time of registration.\nThis could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network."
* `leadership`:(string)"The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection."
* `locked_leader`:(bool)"Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true."
* `member_identity`:(string)"The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `parent_cluster_member_identity`:(string)"The member idenity of the cluster member through which this device is connected if applicable."
* `proxy_app`:(string)"The name of the app which will proxy the messages to the device connector."
