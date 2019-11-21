---
layout: "intersight"
page_title: "Intersight: intersight_server_profile"
sidebar_current: "docs-intersight-data-source-serverProfile"
description: |-
Server Profile enables server management using policies.

---

# Data Source: intersight_server_profile
Server Profile enables server management using policies.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string)User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
* `description`:(string)Description of the profile.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete profile.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `type`:(string)Defines the type of the profile. Accepted value is instance.
