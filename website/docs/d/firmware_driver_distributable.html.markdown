---
layout: "intersight"
page_title: "Intersight: intersight_firmware_driver_distributable"
sidebar_current: "docs-intersight-data-source-firmwareDriverDistributable"
description: |-
A device driver image distributed by Cisco.
---

# Data Source: intersight_firmware_driver_distributable
A device driver image distributed by Cisco.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bundle_type`:(string)The bundle type of the image, as published on cisco.com.
* `description`:(string)User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.
* `download_count`:(int)The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.
* `guid`:(string)The unique identifier for an image in a Cisco repository.
* `import_action`:(string)The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.
* `import_state`:(string)The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.
* `md5sum`:(string)The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.
* `mdfid`:(string)The mdfid of the image provided by cisco.com.
* `model`:(string)The endpoint model for which this firmware image is applicable.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the file. It is populated as part of the image import operation.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `platform_type`:(string)The platform type of the image.
* `recommended_build`:(string)The build which is recommended by Cisco.
* `release_notes_url`:(string)The url for the release notes of this image.
* `sha512sum`:(string)The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.
* `size`:(int)The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.
* `software_advisory_url`:(string)The software advisory, if any, provided by the vendor for this file.
* `software_type_id`:(string)The software type id provided by cisco.com.
* `vendor`:(string)The vendor or publisher of this file.
* `version`:(string)Vendor provided version for the file.
