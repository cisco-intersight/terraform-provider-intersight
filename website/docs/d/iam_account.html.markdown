---
layout: "intersight"
page_title: "Intersight: intersight_iam_account"
sidebar_current: "docs-intersight-data-source-iamAccount"
description: |-
The Intersight Account used to access Intersight.

---

# Data Source: intersight_iam_account
The Intersight Account used to access Intersight.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the Intersight account. By default, name is same as the MoID of the account.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `status`:(string)Status of the account. To activate the Intersight account, claim a device to the account.
