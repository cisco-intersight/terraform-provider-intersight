---
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate"
sidebar_current: "docs-intersight-resource-iamCertificate"
description: |-
  Holds a certificate, signed by a CAcert.

---

# Resource: intersight_iam_certificate
Holds a certificate, signed by a CAcert.

## Argument Reference
The following arguments are supported:
* `certificate`:(Array with Maximum of one item) -User-input pem-encoded certificate, signed by a CAcert.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `issuer`:(Array with Maximum of one item) -(Computed)The X.509 distinguished name of the issuer of this certificate.
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `common_name`:(string)(Computed)A required component that identifies a person or an object.
    + `country`:
                (Array of schema.TypeString) -(Computed)Identifier for the country in which the entity resides.
    + `locality`:
                (Array of schema.TypeString) -(Computed)Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.
    + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
    + `organization`:
                (Array of schema.TypeString) -(Computed)Identifier for the organization in which the entity resides.
    + `organizational_unit`:
                (Array of schema.TypeString) -(Computed)Identifier for a unit within the organization.
    + `state`:
                (Array of schema.TypeString) -(Computed)Identifier for the state or province of the entity.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `pem_certificate`:(string)The base64 encoded certificate in PEM format.
  + `sha256_fingerprint`:(string)(Computed)The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.
  + `signature_algorithm`:(string)(Computed)Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).
  + `subject`:(Array with Maximum of one item) -(Computed)The X.509 distinguished name of the subject of this certificate.
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `common_name`:(string)(Computed)A required component that identifies a person or an object.
    + `country`:
                (Array of schema.TypeString) -(Computed)Identifier for the country in which the entity resides.
    + `locality`:
                (Array of schema.TypeString) -(Computed)Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.
    + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
    + `organization`:
                (Array of schema.TypeString) -(Computed)Identifier for the organization in which the entity resides.
    + `organizational_unit`:
                (Array of schema.TypeString) -(Computed)Identifier for a unit within the organization.
    + `state`:
                (Array of schema.TypeString) -(Computed)Identifier for the state or province of the entity.
* `certificate_request`:(Array with Maximum of one item) -The certificate signing request associated with this certificate.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)(Computed)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `permission_resources`:(Array)(Computed)A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `status`:(string)(Computed)Status of the certificate.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
