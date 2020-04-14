---
layout: "intersight"
page_title: "Intersight: intersight_snmp_policy"
sidebar_current: "docs-intersight-resource-snmpPolicy"
description: |-
  Policy to configure SNMP settings on endpoint.
---

# Resource: intersight_snmp_policy
Policy to configure SNMP settings on endpoint.
## Argument Reference
The following arguments are supported:
* `access_community_string`:(string)"The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long."
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `community_access`:(string)"Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users."
* `description`:(string)"Description of the policy."
* `enabled`:(bool)"State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host."
* `engine_id`:(string)"User-defined unique identification of the static engine."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `organization`:(Array with Maximum of one item) -"Relationship to the Organization that owns the Managed Object."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `profiles`:(Array)"Relationship to the profile object."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `snmp_port`:(int)"Port on which Cisco IMC SNMP agent runs."
* `snmp_traps`:(Array)"List of SNMP traps for the policy."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `destination`:(string)"Address to which the SNMP trap information is sent."
  + `enabled`:(bool)"Enables/disables the trap on the server If enabled, trap is active on the server."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `port`:(int)"Port used by the server to communicate with trap destination. Enter a value between 1-65535."
  + `type`:(string)"Type of trap which decides whether to receive a notification when a trap is received at the destination."
  + `user`:(string)"SNMP user for the trap. Applicable only to SNMPv3."
  + `version`:(string)"SNMP version used for the trap."
* `snmp_users`:(Array)"List of SNMP users for the policy."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `auth_password`:(string)"Authorization password for the user."
  + `auth_type`:(string)"Authorization protocol for authenticating the user."
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `is_auth_password_set`:(bool)(Computed)"Indicates whether the value of the 'authPassword' property has been set."
  + `is_privacy_password_set`:(bool)(Computed)"Indicates whether the value of the 'privacyPassword' property has been set."
  + `name`:(string)"SNMP username. Must have a minimum of 1 and and a maximum of 31 characters."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `privacy_password`:(string)"Privacy password for the user."
  + `privacy_type`:(string)"Privacy protocol for the user."
  + `security_level`:(string)"Security mechanism used for communication between agent and manager."
* `sys_contact`:(string)"Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number."
* `sys_location`:(string)"Location of host on which the SNMP agent (server) runs."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
* `trap_community`:(string)"SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users."
