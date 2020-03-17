---
layout: "intersight"
page_title: "Intersight: intersight_iam_security_holder"
sidebar_current: "docs-intersight-data-source-iamSecurityHolder"
description: |-
Holder for organization aggregated permissions and global account permissions.
User configures permissions for entire account or subset of organizations and specifies associated roles with each organization.
Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
---

# Data Source: intersight_iam_security_holder
Holder for organization aggregated permissions and global account permissions.
User configures permissions for entire account or subset of organizations and specifies associated roles with each organization.
Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
