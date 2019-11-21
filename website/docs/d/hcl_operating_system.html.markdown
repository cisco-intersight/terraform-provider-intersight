---
layout: "intersight"
page_title: "Intersight: intersight_hcl_operating_system"
sidebar_current: "docs-intersight-data-source-hclOperatingSystem"
description: |-
Collection used to store operating system details.

---

# Data Source: intersight_hcl_operating_system
Collection used to store operating system details.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `version`:(string)Version of the Operating System.
