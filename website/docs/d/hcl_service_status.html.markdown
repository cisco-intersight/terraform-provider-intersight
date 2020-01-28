---
layout: "intersight"
page_title: "Intersight: intersight_hcl_service_status"
sidebar_current: "docs-intersight-data-source-hclServiceStatus"
description: |-
Status of the service indicatating if the service is up or under maintenance due to data update. Service will not be able serve any requests when the data is being updated. Collection will have only one document.

---

# Data Source: intersight_hcl_service_status
Status of the service indicatating if the service is up or under maintenance due to data update. Service will not be able serve any requests when the data is being updated. Collection will have only one document.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `exemption_file_version`:(string)Version of the last modified exemption file.
* `identity`:(string)A field to uniquely identify the document with the status.
* `last_hcl_data_modified_time`:(string)The timestamp of the last modified record in the HCL tool database. Used to query and get updated records.
* `last_imported_data_checksum`:(string)Checksum of the data dump used as the base for delta updates.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `status`:(string)Status of the service indicatating if the service is up or under maintenance due to data update.
