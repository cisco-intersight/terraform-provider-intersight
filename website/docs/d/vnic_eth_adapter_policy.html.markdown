---
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_adapter_policy"
sidebar_current: "docs-intersight-data-source-vnicEthAdapterPolicy"
description: |-
An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface you can configure various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings.
---

# Data Source: intersight_vnic_eth_adapter_policy
An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface you can configure various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advanced_filter`:(bool)Enables advanced filtering on the interface.
* `description`:(string)Description of the policy.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `rss_settings`:(bool)Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
