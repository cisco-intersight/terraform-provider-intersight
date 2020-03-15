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
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `ipv4_address`:(string)The IPv4 address of system.
* `ipv6_address`:(string)The IPv6 address of system.
* `mode`:(string)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `time_zone`:(string)The operational timezone of the system, empty indicates no timezone has been set specifically.
