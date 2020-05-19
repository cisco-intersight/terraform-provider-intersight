---
layout: "intersight"
page_title: "Intersight: intersight_virtualization._vmware_datastore"
sidebar_current: "docs-intersight-data-source-virtualization.VmwareDatastore"
description: |-
The VMware Datastore entity with its attributes. Each Datastore belongs to a Datacenter and maybe attached to VMs.
---

# Data Source: intersight_virtualization._vmware_datastore
The VMware Datastore entity with its attributes. Each Datastore belongs to a Datacenter and maybe attached to VMs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accessible`:(bool)"Shows if this datastore is accessible."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `host_count`:(int)"Number of hosts attached to or supported-by this datastore."
* `identity`:(string)"The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference)."
* `maintenance_mode`:(bool)"Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `multiple_host_access`:(bool)"Indicates if this datastore is connected to multiple hosts."
* `name`:(string)"Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `status`:(string)"Datastore health status, as reported by the hypervisor platform."
* `thin_provisioning_supported`:(bool)"Indicates if this datastore supports thin provisioning for files."
* `type`:(string)"A string indicating the type of the datastore (VMFS, NFS, etc)."
* `un_committed`:(int)"Space uncommitted in this datastore in bytes."
* `url`:(string)"The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/')."
* `vm_count`:(int)"Number of virtual machines relying on (using) this datastore."
