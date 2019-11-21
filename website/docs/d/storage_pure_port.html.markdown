---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_port"
sidebar_current: "docs-intersight-data-source-storagePurePort"
description: |-
Port entity in Pure FlashArray.

---

# Data Source: intersight_storage_pure_port
Port entity in Pure FlashArray.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `failover`:(string)Name of the port to which this port has failed over.
* `iqn`:(string)ISCSI qualified name applicable for ethernet port. eg 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the physical port available in storage array.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `portal`:(string)Ip address of iSCSI portal configured on the port.
* `speed`:(int)Operational speed of physical port measured in Gbps.
* `status`:(string)Status of storage array port.
* `type`:(string)Port type, it can be a any of the following category FC, FCoE and iSCSI.
* `wwn`:(string)World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexa decimal formatted with colon. e.g '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'.
* `wwnn`:(string)World wide node name of FC port. Represented in 64 bit digit hexa formatted with colon eg '51:4f:0c:50:14:1f:af:01'.
* `wwpn`:(string)World wide port name of FC port. Represented in 64 bit digit hexa formatted with colon eg '51:4f:0c:51:14:1f:af:01'.
