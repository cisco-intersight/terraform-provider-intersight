
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_cluster"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-cluster"
description: |-
A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
---

# Data Source: intersight_hyperflex._hxap_cluster
A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `datacenter_name`:(string) Datacenter to which the cluster belongs. 
* `failure_reason`:(string) Reason of the failure when cluster is in failed state. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor. 
* `identity`:(string) The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). 
* `management_ip_address`:(string) Management IP Address of the cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user-provided name for this cluster to facilitate identification. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `status`:(string) Cluster health status as reported by the hypervisor platform. 
* `total_cores`:(int) Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. 
* `version`:(string) Product version of HyperFlex compute cluster. 
