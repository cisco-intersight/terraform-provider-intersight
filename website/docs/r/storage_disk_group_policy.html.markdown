---
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group_policy"
sidebar_current: "docs-intersight-resource-storageDiskGroupPolicy"
description: |-
  A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
---

# Resource: intersight_storage_disk_group_policy
A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
## Argument Reference
The following arguments are supported:
* `dedicated_hot_spares`:(Array)"A collection of disks used as hot spares for this RAID group."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `slot_number`:(int)"The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server."
* `description`:(string)"Description of the policy."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `organization`:(Array with Maximum of one item) -"Relationship to the Organization that owns the Managed Object."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `raid_level`:(string)"The supported RAID level for the disk group."
* `span_groups`:(Array)"Non spanned RAID levels like Raid0, Raid1, Raid5 and Raid6 expect a single group of disks whereas spanned RAID levels need multiple groups of disks with each group representing a span group. Non spanned RAID levels expect one span group and spanned RAID levels accept minimum 2 span groups and up to 8."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `disks`:(Array)"Collection of local disks that are part of this span group. The minimum number of disks needed in a span group varies based on RAID level. Raid0 requires at least one disk, Raid1 and Raid10 requires at least 2 and in multiples of 2, Raid5 Raid50 Raid6 and Raid60 require at least 3 disks in a span group."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `slot_number`:(int)"The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
* `storage_policies`:(Array)"A collection of references to the [storage.StoragePolicy](mo://storage.StoragePolicy) Managed Object.\nWhen this managed object is deleted, the referenced [storage.StoragePolicy](mo://storage.StoragePolicy) MOs unset their reference to this deleted MO."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
* `use_jbods`:(bool)"Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation."
