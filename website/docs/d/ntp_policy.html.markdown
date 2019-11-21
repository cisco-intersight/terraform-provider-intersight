---
layout: "intersight"
page_title: "Intersight: intersight_ntp_policy"
sidebar_current: "docs-intersight-data-source-ntpPolicy"
description: |-
Policy to configure the NTP Servers.

---

# Data Source: intersight_ntp_policy
Policy to configure the NTP Servers.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `enabled`:(bool)State of NTP service on the endpoint.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
