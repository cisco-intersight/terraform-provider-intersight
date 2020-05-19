---
layout: "intersight"
page_title: "Intersight: intersight_organization._organization"
sidebar_current: "docs-intersight-resource-organization.Organization"
description: |-
  Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.
---

# Resource: intersight_organization._organization
Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.
## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) -(Computed)"A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `description`:(string)"The informative description about the usage of this organization."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"The name of the organization. There can be multiple organizations under an account."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `permission_resources`:(Array)(Computed)"An array of relationships to moBaseMo resources."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `resource_groups`:(Array)"An array of relationships to resourceGroup resources."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string)"The string representation of a tag key."
  + `value`:(string)"The string representation of a tag value."
