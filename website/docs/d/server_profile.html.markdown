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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `type`:(string)Defines the type of the profile. Accepted value is instance.
