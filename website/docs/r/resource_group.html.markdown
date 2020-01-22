---
layout: "intersight"
page_title: "Intersight: intersight_resource_group"
sidebar_current: "docs-intersight-resource-resourceGroup"
description: |-
  A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.

---

# Resource: intersight_resource_group
A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.

## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) -(Computed)The account to which this resource group belongs to.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of this resource group.
* `object_type`:(string)(Computed)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `organizations`:(Array)A collection of references to the [organization.Organization](mo://organization.Organization) Managed Object.When this managed object is deleted, the referenced [organization.Organization](mo://organization.Organization) MOs unset their reference to this deleted MO.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `per_type_combined_selector`:(Array)(Computed)A single filter expression created by OR'ing the $filter criteria of the 'selectors' of a given object type. Used to efficiently maintain the membership of the group. This filter is maintained per object type in this type.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `combined_selector`:(string)(Computed)A single filter expression created by OR'ing the $filter criteria of the 'selectors'. Used to efficiently maintain the membership of the Group.
  + `empty_filter`:(bool)(Computed)If true, then resources are added using just object type without filter.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector_object_type`:(string)(Computed)The ObjectType on which the selectors are defined. Used to efficiently query resource groups for a given ObjectType.
* `permission_resources`:(Array)(Computed)A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `qualifier`:(string)Qualifier shall be used to specify if we want to organize resources using multiple resource group or single For an account, resource groups can be of only one of the above types. (Both the types are mutually exclusive for an account.).
* `selectors`:(Array)A list of ODATA filters to select resources. The group selectors may be include URLs of individual resources, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host).
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)ODATA filter to select resources. The group selector may include URLs of individual resource, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host).
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
