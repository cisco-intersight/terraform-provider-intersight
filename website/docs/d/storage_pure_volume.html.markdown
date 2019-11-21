---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_volume"
sidebar_current: "docs-intersight-data-source-storagePureVolume"
description: |-
A volume entity in PureStorage FlashArray.

---

# Data Source: intersight_storage_pure_volume
A volume entity in PureStorage FlashArray.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created`:(string)Creation time of the volume.
* `description`:(string)Short description about the volume.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `naa_id`:(string)NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
* `name`:(string)Named entitiy of the volume.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `serial`:(string)Serial number of the volume.
* `size`:(int)User provisioned volume size. It is a size exposed to host.
* `source`:(string)Source from which the volume is created. Applicable only if the volume is cloned from other volume or snapshot.
