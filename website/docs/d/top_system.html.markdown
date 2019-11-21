---
layout: "intersight"
page_title: "Intersight: intersight_top_system"
sidebar_current: "docs-intersight-data-source-topSystem"
description: |-
Root container for all UCSM / CIMC MOs.

---

# Data Source: intersight_top_system
Root container for all UCSM / CIMC MOs.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)
* `ipv4_address`:(string)The IPv4 address of system.
* `ipv6_address`:(string)The IPv6 address of system.
* `mode`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rn`:(string)
* `time_zone`:(string)The operational timezone of the system, empty indicates no timezone has been set specifically.
