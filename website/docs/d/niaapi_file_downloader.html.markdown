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
* `file_name`:(string)Filename of this Metadata package file, folder will be handled by api.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `presigned_url`:(string)The presigned URL from server to download this file.
