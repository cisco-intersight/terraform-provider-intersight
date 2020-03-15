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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `status`:(string)Status of the account. To activate the Intersight account, claim a device to the account.
