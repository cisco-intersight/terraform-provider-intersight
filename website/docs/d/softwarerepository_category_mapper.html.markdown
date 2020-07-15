
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_mapper"
sidebar_current: "docs-intersight-data-source-softwarerepository-category-mapper"
description: |-
Maps a Cisco software repository image category identifier to its applicable hardware models.
---

# Data Source: intersight_softwarerepository._category_mapper
Maps a Cisco software repository image category identifier to its applicable hardware models.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) The category of the model series. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `regex_pattern`:(string) The regex that all images of this category follow. 
