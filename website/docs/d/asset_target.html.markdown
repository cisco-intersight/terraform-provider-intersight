---
layout: "intersight"
page_title: "Intersight: intersight_asset_target"
sidebar_current: "docs-intersight-data-source-assetTarget"
description: |-
Target represents an entity which can be managed by Intersight. This includes physical entities like UCS and HyperFlex servers and software entities like VMware vCenter and Microsoft Azure cloud accounts.
---

# Data Source: intersight_asset_target
Target represents an entity which can be managed by Intersight. This includes physical entities like UCS and HyperFlex servers and software entities like VMware vCenter and Microsoft Azure cloud accounts.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `status`:(string)"Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target."
* `status_error_reason`:(string)"StatusErrorReason provides additional context for the Status."
* `target_type`:(string)"The type of the managed target. For example a UCS Server or Vmware Vcenter target."
