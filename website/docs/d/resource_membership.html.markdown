---
layout: "intersight"
page_title: "Intersight: intersight_resource_membership"
sidebar_current: "docs-intersight-data-source-resourceMembership"
description: |-
ResourceMembership represents a resource's associated groups, organizations and the permissions associated to this resource via organizations.
---

# Data Source: intersight_resource_membership
ResourceMembership represents a resource's associated groups, organizations and the permissions associated to this resource via organizations.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `target_app`:(string)Name of the Service owning the resource.
