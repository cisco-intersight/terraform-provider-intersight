---
layout: "intersight"
page_title: "Intersight: intersight_vnic._eth_if"
sidebar_current: "docs-intersight-resource-vnic.EthIf"
description: |-
  Virtual Ethernet Interface.
---

# Resource: intersight_vnic._eth_if
Virtual Ethernet Interface.
## Argument Reference
The following arguments are supported:
* `cdn`:(Array with Maximum of one item) -"Consistent Device Naming configuration for the virtual NIC."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `source`:(string)"Source of the CDN. It can either be user specified or be the same as the vNIC name."
  + `value`:(string)"The CDN value entered in case of user defined mode."
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `eth_adapter_policy`:(Array with Maximum of one item) -"A reference to a vnicEthAdapterPolicy resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `eth_network_policy`:(Array with Maximum of one item) -"A reference to a vnicEthNetworkPolicy resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `eth_qos_policy`:(Array with Maximum of one item) -"A reference to a vnicEthQosPolicy resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `lan_connectivity_policy`:(Array with Maximum of one item) -"A reference to a vnicLanConnectivityPolicy resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the virtual ethernet interface."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `order`:(int)"The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two."
* `organization`:(Array with Maximum of one item) -"A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `permission_resources`:(Array)(Computed)"An array of relationships to moBaseMo resources."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `placement`:(Array with Maximum of one item) -"Placement Settings for the virtual interface."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `id`:(string)"PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `pci_link`:(int)"The PCI Link used as transport for the virtual interface. All VIC adapters have a single PCI link except VIC 1385 which has two."
  + `uplink`:(int)"Adapter port on which the virtual interface will be created."
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string)"The string representation of a tag key."
  + `value`:(string)"The string representation of a tag value."
* `usnic_settings`:(Array with Maximum of one item) -"User Space NIC Settings that enable low-latency and higher throughput by bypassing the kernel layer when sending/receiving packets."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `cos`:(int)"Class of Service to be used for traffic on the usNIC."
  + `count`:(int)"Number of usNIC interfaces to be created."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `usnic_adapter_policy`:(string)"Ethernet Adapter policy to be associated with the usNICs."
* `vmq_settings`:(Array with Maximum of one item) -"Virtual Machine Queue Settings for the virtual interface that allow efficient transfer of network traffic to the guest OS."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `enabled`:(bool)"Enables VMQ feature on the virtual interface."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
