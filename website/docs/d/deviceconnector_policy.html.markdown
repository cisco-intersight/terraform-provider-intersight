---
layout: "intersight"
page_title: "Intersight: intersight_deviceconnector_policy"
sidebar_current: "docs-intersight-data-source-deviceconnectorPolicy"
description: |-
Device Connector Policy.

---

# Data Source: intersight_deviceconnector_policy
Device Connector Policy.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `lockout_enabled`:(bool)Enables configuration lockout on the endpoint.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
