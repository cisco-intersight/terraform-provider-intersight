---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository._operating_system_file"
sidebar_current: "docs-intersight-data-source-softwarerepository.OperatingSystemFile"
description: |-
An operating system image that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository.
---

# Data Source: intersight_softwarerepository._operating_system_file
An operating system image that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `description`:(string)"User provided description about the file. Cisco provided description for image inventoried from a Cisco repository."
* `download_count`:(int)"The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache."
* `import_action`:(string)"The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source."
* `import_state`:(string)"The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process."
* `md5sum`:(string)"The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"The name of the file. It is populated as part of the image import operation."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `sha512sum`:(string)"The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository."
* `size`:(int)"The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository."
* `software_advisory_url`:(string)"The software advisory, if any, provided by the vendor for this file."
* `vendor`:(string)"The vendor or publisher of this file."
* `version`:(string)"Vendor provided version for the file."
