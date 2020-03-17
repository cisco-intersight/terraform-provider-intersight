---
layout: "intersight"
page_title: "Intersight: intersight_software_solution_distributable"
sidebar_current: "docs-intersight-resource-softwareSolutionDistributable"
description: |-
  A solution image distributed by Cisco.
---

# Resource: intersight_software_solution_distributable
A solution image distributed by Cisco.
## Argument Reference
The following arguments are supported:
* `bundle_type`:(string)(Computed)"The bundle type of the image, as published on cisco.com."
* `catalog`:(Array with Maximum of one item) -"The catalog where this image is present."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `description`:(string)"User provided description about the file. Cisco provided description for image inventoried from a Cisco repository."
* `download_count`:(int)(Computed)"The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache."
* `file_path`:(string)(Computed)"The path of the file in S3/minio bucket."
* `guid`:(string)(Computed)"The unique identifier for an image in a Cisco repository."
* `import_action`:(string)"The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source."
* `import_state`:(string)(Computed)"The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process."
* `md5sum`:(string)"The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository."
* `mdfid`:(string)"The mdfid of the image provided by cisco.com."
* `model`:(string)(Computed)"The endpoint model for which this firmware image is applicable."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"The name of the file. It is populated as part of the image import operation."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `platform_type`:(string)(Computed)"The platform type of the image."
* `recommended_build`:(string)"The build which is recommended by Cisco."
* `release_notes_url`:(string)"The url for the release notes of this image."
* `sha512sum`:(string)"The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository."
* `size`:(int)"The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository."
* `software_advisory_url`:(string)"The software advisory, if any, provided by the vendor for this file."
* `software_type_id`:(string)(Computed)"The software type id provided by cisco.com."
* `solution_name`:(string)"The name of the solution in which the image belongs."
* `source`:(Array with Maximum of one item) -"Location of the file in an external repository."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
* `sub_type`:(string)"The type of the file like OS image, Script etc."
* `supported_models`:
                (Array of schema.TypeString) -"The server models for which this image is applicable."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
* `vendor`:(string)"The vendor or publisher of this file."
* `version`:(string)"Vendor provided version for the file."
