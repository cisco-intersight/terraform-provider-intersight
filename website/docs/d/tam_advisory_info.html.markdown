---
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_info"
sidebar_current: "docs-intersight-data-source-tamAdvisoryInfo"
description: |-
State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.

---

# Data Source: intersight_tam_advisory_info
State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `state`:(string)Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory.
