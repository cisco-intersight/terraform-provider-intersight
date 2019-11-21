---
layout: "intersight"
page_title: "Intersight: intersight_cvd_template"
sidebar_current: "docs-intersight-resource-cvdTemplate"
description: |-
  Template represents a CVD definition

---

# Resource: intersight_cvd_template
Template represents a CVD definition

## Argument Reference
The following arguments are supported:
* `deployer`:(string)URL pointing to the S3 location of the workflow JSON which performs the deployment task for this CVD
* `deployer_input`:(Array)A collection of input name-value pairs
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `field_filter`:(string)Input field filter(REST API path with filter which will return the list of applicable values for this field)
  + `field_label`:(string)Input field label(for GUI use)
  + `field_name`:(string)Input field name
  + `field_value`:(string)Input field value
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `description`:
                (Array of schema.TypeString) -A short description for the CVD
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Unique name identifier for the CVD
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
* `upload_location`:(string)S3 directory location where the CVD definition has been uploaded
* `validator`:(string)URL pointing to the S3 location of the workflow JSON which performs the validation task for this CVD
* `validator_input`:(Array)A collection of input name-value pairs
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `field_filter`:(string)Input field filter(REST API path with filter which will return the list of applicable values for this field)
  + `field_label`:(string)Input field label(for GUI use)
  + `field_name`:(string)Input field name
  + `field_value`:(string)Input field value
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `version`:(int)The version or revision number of the CVD definition
