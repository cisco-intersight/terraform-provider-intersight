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
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `service_ticket_receipient`:(string)The email address recipient for support tickets.
