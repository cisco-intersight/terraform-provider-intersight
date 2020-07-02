
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_cached_image"
sidebar_current: "docs-intersight-data-source-softwarerepository-cached-image"
description: |-
The image cached in the customer's datacenter.
---

# Data Source: intersight_softwarerepository._cached_image
The image cached in the customer's datacenter.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in endpoint. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source. 
* `cache_state`:(string) The state  of this file in the endpoint The importState is updated during the cache operation and as part of the cache monitoring process. 
* `cached_time`:(string) The time when the image or file was cached into the FI storage. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `download_progress`:(int) The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. 
* `download_retries`:(int) The number of retries the plugin attempted before succeeding or failing the download. 
* `md5sum`:(string) The MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `original_sha512sum`:(string) The actual sha512sum of the cached image. 
* `path`:(string) The absolute path of the imported file in the endpoint. 
* `used_count`:(int) The number of times this file has been used to copy or upgrade or install actions. Used by the cache monitoring process to determine the files to be evicted from the cache. 
