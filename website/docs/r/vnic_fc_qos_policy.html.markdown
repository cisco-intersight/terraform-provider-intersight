---
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_qos_policy"
sidebar_current: "docs-intersight-resource-vnicFcQosPolicy"
description: |-
  A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters you can also specify additional controls like burst and rate on the outgoing traffic.
---

# Resource: intersight_vnic_fc_qos_policy
A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters you can also specify additional controls like burst and rate on the outgoing traffic.
## Argument Reference
The following arguments are supported:
* `cos`:(int)Class of Service to be associated to the traffic on the virtual interface.
* `description`:(string)Description of the policy.
* `max_data_field_size`:(int)The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
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
* `rate_limit`:(int)The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
