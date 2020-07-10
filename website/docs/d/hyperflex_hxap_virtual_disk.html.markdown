
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_virtual_disk"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-virtual-disk"
description: |-
The Virtual disk created on HyperFlex Application Platform compute cluster.
---

# Data Source: intersight_hyperflex._hxap_virtual_disk
The Virtual disk created on HyperFlex Application Platform compute cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_mode`:(string) Access mode of the virtual disk. 
* `capacity`:(int) Disk capacity represented in bytes. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `mode`:(string) File mode of the disk  example - Filesystem, Block. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `source_file_path`:(string) Source file path associated with virtual machine disk. 
* `source_virtual_disk`:(string) Source disk name from where the clone is done. 
* `uuid`:(string) UUID of the virtual disk. 
