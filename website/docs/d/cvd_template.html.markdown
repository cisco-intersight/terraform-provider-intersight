---
layout: "intersight"
page_title: "Intersight: intersight_cvd_template"
sidebar_current: "docs-intersight-data-source-cvdTemplate"
description: |-
Template represents a CVD definition

---

# Data Source: intersight_cvd_template
Template represents a CVD definition

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `deployer`:(string)URL pointing to the S3 location of the workflow JSON which performs the deployment task for this CVD
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Unique name identifier for the CVD
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `upload_location`:(string)S3 directory location where the CVD definition has been uploaded
* `validator`:(string)URL pointing to the S3 location of the workflow JSON which performs the validation task for this CVD
* `version`:(int)The version or revision number of the CVD definition
