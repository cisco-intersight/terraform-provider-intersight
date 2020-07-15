
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_claim"
sidebar_current: "docs-intersight-resource-appliance-device-claim"
description: |-
  DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.
Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.
---

# Resource: intersight_appliance._device_claim
DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.
Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.
## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) -(Computed) A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_id`:(string)(Computed) Device identifier of the endpoint device. 
* `hostname`:(string) Hostname or IP address of the endpoint device the user wants to claim. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `message`:(string)(Computed) Message set by the device claim process. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `password`:(string) Password to be used to login to the endpoint device. 
* `platform_type`:(string) Platform type of the endpoint device. 
* `request_id`:(string) User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set. 
* `security_token`:(string)(Computed) Device security token of the endpoint device. 
* `status`:(string)(Computed) Status of the device claim process. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `username`:(string) Username to log in to the endpoint device. 
