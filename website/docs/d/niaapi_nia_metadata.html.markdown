---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nia_metadata"
sidebar_current: "docs-intersight-data-source-niaapiNiaMetadata"
description: |-
Contains the latest Metadata available for download from server.
---

# Data Source: intersight_niaapi_nia_metadata
Contains the latest Metadata available for download from server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `date`:(string)"The date when this package is generated."
* `metadata_chksum`:(string)"Chksum used to check the integrity of the Metadata file downloaded."
* `metadata_filename`:(string)"The Filename of this Metadata package."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `version`:(int)"The version number of the Metadata package."
