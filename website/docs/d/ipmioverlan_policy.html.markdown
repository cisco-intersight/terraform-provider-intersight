---
layout: "intersight"
page_title: "Intersight: intersight_ipmioverlan_policy"
sidebar_current: "docs-intersight-data-source-ipmioverlanPolicy"
description: |-
Intelligent Platform Management Interface Over LAN Policy.
---

# Data Source: intersight_ipmioverlan_policy
Intelligent Platform Management Interface Over LAN Policy.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)"Description of the policy."
* `enabled`:(bool)"State of the IPMI Over LAN service on the endpoint."
* `encryption_key`:(string)"The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters."
* `is_encryption_key_set`:(bool)"Indicates whether the value of the 'encryptionKey' property has been set."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `privilege`:(string)"The highest privilege level that can be assigned to an IPMI session on a server."
