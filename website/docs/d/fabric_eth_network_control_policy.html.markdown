
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_eth_network_control_policy"
sidebar_current: "docs-intersight-data-source-fabric-eth-network-control-policy"
description: |-
An Ethernet Network Control policy determines the features that are applied on a vethernet that is connected to the vNIC.
---

# Data Source: intersight_fabric._eth_network_control_policy
An Ethernet Network Control policy determines the features that are applied on a vethernet that is connected to the vNIC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cdp_enabled`:(bool) Enables the CDP on an interface. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `forge_mac`:(string) Determines if the MAC forging is allowed or denied on an interface. 
* `mac_register_mode`:(string) Determines the MAC addresses that have to be registered with the switch. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `uplink_fail_action`:(string) Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. 
