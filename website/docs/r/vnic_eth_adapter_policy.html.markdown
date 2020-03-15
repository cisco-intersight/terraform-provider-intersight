---
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_adapter_policy"
sidebar_current: "docs-intersight-resource-vnicEthAdapterPolicy"
description: |-
  An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface you can configure various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings.
---

# Resource: intersight_vnic_eth_adapter_policy
An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface you can configure various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings.
## Argument Reference
The following arguments are supported:
* `advanced_filter`:(bool)Enables advanced filtering on the interface.
* `arfs_settings`:(Array with Maximum of one item) -Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `enabled`:(bool)Status of Accelerated Receive Flow Steering on the virtual ethernet interface.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `completion_queue_settings`:(Array with Maximum of one item) -Completion Queue resource settings.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `count`:(int)The number of completion queue resources to allocate. In general, the number of completion queue resources you should allocate is equal to the number of transmit queue resources plus the number of receive queue resources.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `ring_size`:(int)(Computed)The number of descriptors in each completion queue.
* `description`:(string)Description of the policy.
* `interrupt_settings`:(Array with Maximum of one item) -Interrupt Settings for the virtual ethernet interface.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `coalescing_time`:(int)The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field.
  + `coalescing_type`:(string)Interrupt Coalescing Type. This can be one of the following:- MIN  - The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.
  + `count`:(int)The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources.
  + `mode`:(string)Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `nvgre_settings`:(Array with Maximum of one item) -Network Virtualization using Generic Routing Encapsulation Settings.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `enabled`:(bool)Status of the Network Virtualization using Generic Routing Encapsulation on the virtual ethernet interface.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `object_type`:(string)(Computed)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `organization`:(Array with Maximum of one item) -Relationship to the Organization that owns the Managed Object.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `permission_resources`:(Array)(Computed)A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `roce_settings`:(Array with Maximum of one item) -Settings for RDMA over Converged Ethernet.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `enabled`:(bool)If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface.
  + `memory_regions`:(int)The number of memory regions per adapter. Recommended value = integer power of 2.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `queue_pairs`:(int)The number of queue pairs per adapter. Recommended value = integer power of 2.
  + `resource_groups`:(int)The number of resource groups per adapter. Recommended value = be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance.
* `rss_settings`:(bool)Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
* `rx_queue_settings`:(Array with Maximum of one item) -Receive Queue resouce settings.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `count`:(int)The number of queue resources to allocate.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `ring_size`:(int)The number of descriptors in each queue.
* `tcp_offload_settings`:(Array with Maximum of one item) -The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `large_receive`:(bool)Enables the reassembly of segmented packets in hardware before sending them to the CPU.
  + `large_send`:(bool)Enables the CPU to send large packets to the hardware for segmentation.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `rx_checksum`:(bool)When enabled, the CPU sends all packet checksums to the hardware for validation.
  + `tx_checksum`:(bool)When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
* `tx_queue_settings`:(Array with Maximum of one item) -Transmit Queue resource settings.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `count`:(int)The number of queue resources to allocate.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `ring_size`:(int)The number of descriptors in each queue.
* `vxlan_settings`:(Array with Maximum of one item) -Virtual Extensible LAN Protocol Settings.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `enabled`:(bool)Status of the Virtual Extensible LAN Protocol on the virtual ethernet interface.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
