---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node_profile"
sidebar_current: "docs-intersight-data-source-hyperflexNodeProfile"
description: |-
A configuration profile per node in the HyperFlex cluster.
It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.

---

# Data Source: intersight_hyperflex_node_profile
A configuration profile per node in the HyperFlex cluster.
It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the profile.
* `hxdp_data_ip`:(string)IP address for storage data network (Controller VM interface).
* `hxdp_mgmt_ip`:(string)IP address for HyperFlex management network.
* `hypervisor_data_ip`:(string)IP address for storage data network (Hypervisor interface).
* `hypervisor_mgmt_ip`:(string)IP address for Hypervisor management network.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete profile.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `type`:(string)Defines the type of the profile. Accepted value is instance.
