---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex._capability_info"
sidebar_current: "docs-intersight-resource-hyperflex.CapabilityInfo"
description: |-
  A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, "minUcsVersion" for HyperFlex version "4.0.1a" corresponds to "3.2.3" or "minHxdpVersion" for HyperFlex Upgrade operation is "4.0.1a" etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
---

# Resource: intersight_hyperflex._capability_info
A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, "minUcsVersion" for HyperFlex version "4.0.1a" corresponds to "3.2.3" or "minHxdpVersion" for HyperFlex Upgrade operation is "4.0.1a" etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
## Argument Reference
The following arguments are supported:
* `app_catalog`:(Array with Maximum of one item) -"A reference to a hyperflexAppCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `capability_constraints`:(Array)
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `constraint_name`:(string)"Name or key of the applicable compatibility constraint."
  + `constraint_value`:(string)"Value of the applicable compatibility constraint. Could either be a string value or a regex."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the capability or feature set consisting of a collection of constraint rules and value."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `permission_resources`:(Array)(Computed)"An array of relationships to moBaseMo resources."
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
* `value`:(string)(Computed)"Capability Value which is valid only iff all specified constraints match."
