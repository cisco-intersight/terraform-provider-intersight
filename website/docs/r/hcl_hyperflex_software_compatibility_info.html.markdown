---
layout: "intersight"
page_title: "Intersight: intersight_hcl_hyperflex_software_compatibility_info"
sidebar_current: "docs-intersight-resource-hclHyperflexSoftwareCompatibilityInfo"
description: |-
  Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.

---

# Resource: intersight_hcl_hyperflex_software_compatibility_info
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.

## Argument Reference
The following arguments are supported:
* `app_catalog`:(Array with Maximum of one item) -A collection of references to the [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) Managed Object.When this managed object is deleted, the referenced [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) MO unsets its reference to this deleted MO.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `constraints`:(Array)Constraint for the matching software compatibility info, in case the match is applicable for certain cases only. For example a combination of (HyperFlex Data Platform, serverFw and hypervisor) versions can be applicable only for a HyperFlex Cluster UPGRADE operation, so a constraint of \"supportedOperations=upgrade\" can be added to the matching row.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `constraint_name`:(string)Name or key of the applicable compatibility constraint.
  + `constraint_value`:(string)Value of the applicable compatibility constraint. Could either be a string value or a regex.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `hxdp_version`:(string)HXDP component software version.
* `hypervisor_type`:(string)Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.
* `hypervisor_version`:(string)Hypervisor component software version.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)(Computed)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `permission_resources`:(Array)(Computed)A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `server_fw_version`:(string)UCS Server Firmware component software version.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
