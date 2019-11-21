---
layout: "intersight"
page_title: "Intersight: intersight_ipmioverlan_policy"
sidebar_current: "docs-intersight-data-source-ipmioverlanPolicy"
description: |-
IPMI Over LAN Policy

---

# Data Source: intersight_ipmioverlan_policy
IPMI Over LAN Policy

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `enabled`:(bool)State of the IPMI Over LAN service on the endpoint.
* `encryption_key`:(string)The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters.
* `is_encryption_key_set`:(bool)Indicates whether the value of the 'encryptionKey' property has been set.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `privilege`:(string)The highest privilege level that can be assigned to an IPMI session on a server.
