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
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rss_settings`:(bool)Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
