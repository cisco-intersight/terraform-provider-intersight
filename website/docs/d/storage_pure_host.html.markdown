---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host"
sidebar_current: "docs-intersight-data-source-storagePureHost"
description: |-
A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.

---

# Data Source: intersight_storage_pure_host
A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Short description about the host.
* `host_group_name`:(string)Name of host group where the host belongs to. Empty if HostType is set as HostGroup.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the host in storage array.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `os_type`:(string)Operating system running on the host.
* `type`:(string)Type of host entity whether it is a single host or host group (collection of host).
