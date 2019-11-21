---
layout: "intersight"
page_title: "Intersight: intersight_iam_qualifier"
sidebar_current: "docs-intersight-data-source-iamQualifier"
description: |-
The qualifier defines how a user qualifies to be part of a user group.

---

# Data Source: intersight_iam_qualifier
The qualifier defines how a user qualifies to be part of a user group.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
