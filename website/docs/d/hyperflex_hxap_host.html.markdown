
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_host"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-host"
description: |-
A HyperFlex Application Platform compute host entity that is part of HyperFlex compute cluster and probably runs VMs.
---

# Data Source: intersight_hyperflex._hxap_host
A HyperFlex Application Platform compute host entity that is part of HyperFlex compute cluster and probably runs VMs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `age`:(string) Denotes age or life time of the Host in nano seconds. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cluster_uuid`:(string) The UUID of the cluster to which this Host belongs to. 
* `failure_reason`:(string) Reason of the failure when host is in failed state. 
* `hw_power_state`:(string) Is the host Powered-up or Powered-down. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor. 
* `identity`:(string) The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). 
* `internal_ip_address`:(string) Internal IP Address of the Host. 
* `maintenance_mode`:(bool) Is this host in maintenance mode. Set to true or false. 
* `management_ip_address`:(string) Management IP Address of the Host. 
* `master_role`:(bool) Is the role of this host is master in the cluster? true or false. 
* `model`:(string) Commercial model information about this hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `serial`:(string) Serial number of this host (internally generated). 
* `status`:(string) Host health status, as reported by the hypervisor platform. 
* `up_time`:(string) The uptime of the host, stored as Duration (from w3c). 
* `uuid`:(string) Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. 
* `vendor`:(string) Commercial vendor details of this hardware. 
* `version`:(string) Product version of the Host. 
