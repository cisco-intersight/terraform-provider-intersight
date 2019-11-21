---
layout: "intersight"
page_title: "Intersight: intersight_management_entity"
sidebar_current: "docs-intersight-data-source-managementEntity"
description: |-
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.

---

# Data Source: intersight_management_entity
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)
* `entity_id`:(string)Identity of the Fabric Interconnect - A/B.
* `leadership`:(string)Role (Primary / Subordinate) of the Fabric Interconnect.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rn`:(string)
