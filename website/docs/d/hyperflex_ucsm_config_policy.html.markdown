---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ucsm_config_policy"
sidebar_current: "docs-intersight-data-source-hyperflexUcsmConfigPolicy"
description: |-
A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.

---

# Data Source: intersight_hyperflex_ucsm_config_policy
A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `server_firmware_version`:(string)The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc.
