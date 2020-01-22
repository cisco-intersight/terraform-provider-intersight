---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_auto_support_policy"
sidebar_current: "docs-intersight-data-source-hyperflexAutoSupportPolicy"
description: |-
A policy specifying the configuration to automatically generate support tickets with Cisco TAC.

---

# Data Source: intersight_hyperflex_auto_support_policy
A policy specifying the configuration to automatically generate support tickets with Cisco TAC.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(bool)Enable or disable Auto Support.
* `description`:(string)Description of the policy.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `service_ticket_receipient`:(string)The email address recipient for support tickets.
