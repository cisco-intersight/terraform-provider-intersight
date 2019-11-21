---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_sys_config_policy"
sidebar_current: "docs-intersight-data-source-hyperflexSysConfigPolicy"
description: |-
A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.

---

# Data Source: intersight_hyperflex_sys_config_policy
A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dns_domain_name`:(string)The DNS Search Domain Name. This setting applies to HyperFlex Data Platform 3.0 or later only.
* `description`:(string)Description of the policy.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `timezone`:(string)The timezone of the HyperFlex cluster's system clock.
