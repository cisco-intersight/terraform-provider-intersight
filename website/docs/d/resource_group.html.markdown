---
layout: "intersight"
page_title: "Intersight: intersight_resource_group"
sidebar_current: "docs-intersight-data-source-resourceGroup"
description: |-
A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.
---

# Data Source: intersight_resource_group
A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of this resource group.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `qualifier`:(string)Qualifier shall be used to specify if we want to organize resources using multiple resource group or single For an account, resource groups can be of only one of the above types. (Both the types are mutually exclusive for an account.).
