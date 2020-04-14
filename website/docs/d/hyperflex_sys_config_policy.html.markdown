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
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `dns_domain_name`:(string)"The DNS Search Domain Name. This setting applies to HyperFlex Data Platform 3.0 or later only."
* `description`:(string)"Description of the policy."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `timezone`:(string)"The timezone of the HyperFlex cluster's system clock."
