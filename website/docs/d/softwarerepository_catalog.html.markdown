---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_catalog"
sidebar_current: "docs-intersight-data-source-softwarerepositoryCatalog"
description: |-
A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
---

# Data Source: intersight_softwarerepository_catalog
A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the catalog. The names are populated and predefined during MO creation.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
