---
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_permission"
sidebar_current: "docs-intersight-data-source-iamResourcePermission"
description: |-
ResourcePermission represents the permissions defined on a resource like organization.
---

# Data Source: intersight_iam_resource_permission
ResourcePermission represents the permissions defined on a resource like organization.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `target_app`:(string)"Name of the service owning the resource."
