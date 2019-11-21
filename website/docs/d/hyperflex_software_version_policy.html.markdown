---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_software_version_policy"
sidebar_current: "docs-intersight-data-source-hyperflexSoftwareVersionPolicy"
description: |-
A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.

---

# Data Source: intersight_hyperflex_software_version_policy
A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `hxdp_version`:(string)Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `server_firmware_version`:(string)Desired server firmware version to apply on the HyperFlex Cluster.
