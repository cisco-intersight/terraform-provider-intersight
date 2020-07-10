
---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_virtual_disk"
sidebar_current: "docs-intersight-data-source-virtualization-virtual-disk"
description: |-
Depicts disk configuration used to be create a virtual disk on a hypervisor datastore.
---

# Data Source: intersight_virtualization._virtual_disk
Depicts disk configuration used to be create a virtual disk on a hypervisor datastore.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `capacity`:(string) Disk capacity to be provided with units example - 10Gi. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `discovered`:(bool) Flag to indicate whether the configuration is created from inventory object. 
* `mode`:(string) File mode of the disk  example - Filesystem, Block. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `source_disk_to_clone`:(string) Source disk from which the content is copied. 
* `source_file_path`:(string) Image path used to import on the created disk. 
