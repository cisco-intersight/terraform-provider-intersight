---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_file_downloader"
sidebar_current: "docs-intersight-data-source-niaapiFileDownloader"
description: |-
Provide a presigned url to download the metadata file from server.
---

# Data Source: intersight_niaapi_file_downloader
Provide a presigned url to download the metadata file from server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `file_name`:(string)"Filename of this Metadata package file, folder will be handled by api."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `presigned_url`:(string)"The presigned URL from server to download this file."
