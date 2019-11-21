---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host_lun"
sidebar_current: "docs-intersight-data-source-storagePureHostLun"
description: |-
A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.

---

# Data Source: intersight_storage_pure_host_lun
A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hlu`:(string)Logical unit number (LUN) by which hosts address specified volume.
* `host_group_name`:(string)Name of the host group associated with LUN.
* `host_name`:(string)Name of the host associated with LUN.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `shared`:(bool)Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection.
* `volume_name`:(string)Name of the storage volume associated with LUN.
