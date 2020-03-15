---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_profile"
sidebar_current: "docs-intersight-data-source-hyperflexClusterProfile"
description: |-
A profile specifying configuration settings for a HyperFlex cluster.
---

# Data Source: intersight_hyperflex_cluster_profile
A profile specifying configuration settings for a HyperFlex cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string)User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
* `data_ip_address`:(string)The storage data IP address for the HyperFlex cluster.
* `description`:(string)Description of the profile.
* `hypervisor_type`:(string)The hypervisor type for the HyperFlex cluster.
* `mac_address_prefix`:(string)The MAC address prefix in the form of 00:25:B5:XX.
* `mgmt_ip_address`:(string)The management IP address for the HyperFlex cluster.
* `mgmt_platform`:(string)The management platform for the HyperFlex cluster.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete profile.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `replication`:(int)The number of copies of each data block written.
* `type`:(string)Defines the type of the profile. Accepted value is instance.
* `wwxn_prefix`:(string)The WWxN prefix in the form of 20:00:00:25:B5:XX.
