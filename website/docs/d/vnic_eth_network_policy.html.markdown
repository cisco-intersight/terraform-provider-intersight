---
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_network_policy"
sidebar_current: "docs-intersight-data-source-vnicEthNetworkPolicy"
description: |-
An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.

---

# Data Source: intersight_vnic_eth_network_policy
An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
