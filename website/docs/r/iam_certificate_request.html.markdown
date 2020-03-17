---
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate_request"
sidebar_current: "docs-intersight-resource-iamCertificateRequest"
description: |-
  The information required to generate a certificate signing request (CSR),
which is a block of encoded text that is given to a Certificate Authority when applying for an SSL Certificate.
---

# Resource: intersight_iam_certificate_request
The information required to generate a certificate signing request (CSR),
which is a block of encoded text that is given to a Certificate Authority when applying for an SSL Certificate.
## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) -(Computed)"The account associated with the CertificateRequest."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `certificate`:(Array with Maximum of one item) -"A collection of references to the [iam.Certificate](mo://iam.Certificate) Managed Object.\nWhen this managed object is deleted, the referenced [iam.Certificate](mo://iam.Certificate) MO on the other side of the relationship is deleted."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `email_address`:(string)"User input email address, an optional part of the subject of the certificate request."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the certificate request."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `private_key_spec`:(Array with Maximum of one item) -"A collection of references to the [iam.PrivateKeySpec](mo://iam.PrivateKeySpec) Managed Object.\nWhen this managed object is deleted, the referenced [iam.PrivateKeySpec](mo://iam.PrivateKeySpec) MO on the other side of the relationship is deleted."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `request`:(string)(Computed)"Generated certificate signing request (CSR) in PEM format."
* `self_signed`:(bool)"Whether the user wants the generated CSR to be self-signed by the appliance."
* `subject`:(Array with Maximum of one item) -"The x.509 distinguished name of the subject of the certificate request."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `common_name`:(string)(Computed)"A required component that identifies a person or an object."
  + `country`:
                (Array of schema.TypeString) -(Computed)"Identifier for the country in which the entity resides."
  + `locality`:
                (Array of schema.TypeString) -(Computed)"Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `organization`:
                (Array of schema.TypeString) -(Computed)"Identifier for the organization in which the entity resides."
  + `organizational_unit`:
                (Array of schema.TypeString) -(Computed)"Identifier for a unit within the organization."
  + `state`:
                (Array of schema.TypeString) -(Computed)"Identifier for the state or province of the entity."
* `subject_alternate_name`:(Array with Maximum of one item) -"The x.509 subject alternate name values of the certificate request."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `dns_name`:
                (Array of schema.TypeString) -(Computed)"Alternate DNS names for the host."
  + `email_address`:
                (Array of schema.TypeString) -(Computed)"Alternate email addresses for the host."
  + `ip_address`:
                (Array of schema.TypeString) -(Computed)"Alternate IP addresses for the host."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `uri`:
                (Array of schema.TypeString) -(Computed)"Alternate URIs for the host."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
