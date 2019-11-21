---
layout: "intersight"
page_title: "Intersight: intersight_ssh_policy"
sidebar_current: "docs-intersight-data-source-sshPolicy"
description: |-
Secure shell policy on the endpoint.

---

# Data Source: intersight_ssh_policy
Secure shell policy on the endpoint.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `enabled`:(bool)State of SSH service on the endpoint.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `port`:(int)Port used for secure shell access.
* `timeout`:(int)Number of seconds to wait before the system considers a SSH request to have timed out.
