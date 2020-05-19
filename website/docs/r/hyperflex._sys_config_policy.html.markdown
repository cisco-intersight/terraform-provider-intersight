---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex._sys_config_policy"
sidebar_current: "docs-intersight-resource-hyperflex.SysConfigPolicy"
description: |-
  A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.
---

# Resource: intersight_hyperflex._sys_config_policy
A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.
## Argument Reference
The following arguments are supported:
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `cluster_profiles`:(Array)"An array of relationships to hyperflexClusterProfile resources."
This complex property has following sub-properties:
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `link`:(string)"A URL to an instance of the 'mo.MoRef' class."
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `description`:(string)"Description of the policy."
* `dns_domain_name`:(string)"The DNS Search Domain Name. This setting applies to HyperFlex Data Platform 3.0 or later only."
* `dns_servers`:
                (Array of schema.TypeString) -
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `ntp_servers`:
                (Array of schema.TypeString) -
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
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
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string)"The string representation of a tag key."
  + `value`:(string)"The string representation of a tag value."
* `timezone`:(string)"The timezone of the HyperFlex cluster's system clock."
