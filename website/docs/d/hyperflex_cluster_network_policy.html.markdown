---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_network_policy"
sidebar_current: "docs-intersight-data-source-hyperflexClusterNetworkPolicy"
description: |-
A policy specifying VLANs for management, VM migration, and VM traffic.

---

# Data Source: intersight_hyperflex_cluster_network_policy
A policy specifying VLANs for management, VM migration, and VM traffic.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `jumbo_frame`:(bool)Enable or disable jumbo frames.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `uplink_speed`:(string)Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G'. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only.
