---
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_if"
sidebar_current: "docs-intersight-resource-vnicFcIf"
description: |-
  Virtual Fibre Channel Interface.
---

# Resource: intersight_vnic_fc_if
Virtual Fibre Channel Interface.
## Argument Reference
The following arguments are supported:
* `fc_adapter_policy`:(Array with Maximum of one item) -Relationship to the the Fibre Channel Adapter Policy.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `fc_network_policy`:(Array with Maximum of one item) -Relationship to the Fibre Channel Network Policy.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `fc_qos_policy`:(Array with Maximum of one item) -Relationship to the Fibre Channel QoS Policy.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the virtual fibre channel interface.
* `object_type`:(string)(Computed)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `order`:(int)The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
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
* `persistent_bindings`:(bool)Enables retention of LUN ID associations in memory until they are manually cleared.
* `placement`:(Array with Maximum of one item) -Placement Settings for the virtual interface.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `id`:(string)PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `pci_link`:(int)The PCI Link used as transport for the virtual interface. All VIC adapters have a single PCI link except VIC 1385 which has two.
  + `uplink`:(int)Adapter port on which the virtual interface will be created.
* `san_connectivity_policy`:(Array with Maximum of one item) -Relationship to the SAN Connectivity Policy.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
