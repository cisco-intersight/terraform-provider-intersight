---
layout: "intersight"
page_title: "Intersight: intersight_sdwan_profile"
sidebar_current: "docs-intersight-data-source-sdwanProfile"
description: |-
A profile that specifies configuration settings for a SDWAN router.
---

# Data Source: intersight_sdwan_profile
A profile that specifies configuration settings for a SDWAN router.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string)User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
* `description`:(string)Description of the profile.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete profile.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `type`:(string)Defines the type of the profile. Accepted value is instance.
