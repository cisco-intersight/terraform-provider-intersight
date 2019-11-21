---
layout: "intersight"
page_title: "Intersight: intersight_hcl_hyperflex_software_compatibility_info"
sidebar_current: "docs-intersight-data-source-hclHyperflexSoftwareCompatibilityInfo"
description: |-
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.

---

# Data Source: intersight_hcl_hyperflex_software_compatibility_info
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hxdp_version`:(string)HXDP component software version.
* `hypervisor_type`:(string)Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.
* `hypervisor_version`:(string)Hypervisor component software version.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `server_fw_version`:(string)UCS Server Firmware component software version.
