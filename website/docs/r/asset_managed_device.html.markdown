---
layout: "intersight"
page_title: "Intersight: intersight_asset_managed_device"
sidebar_current: "docs-intersight-resource-assetManagedDevice"
description: |-
  Attributes for Managed Device in Intersight and it maintains the relationship to the Device Connector Manager (Intersight appliance platform type tagged with 'Intersight Assist') Device Registration. It is possible to add Managed Devices within an Account and later associated to DeviceConnectorManager. Once attached to Device Connector Manager, Device Connector for the Managed Device type is started on the OnPrem and status related to it is maintained in the ManagedDeviceStatus MO. If the Device Connector Manager is deleted, Managed Device continues to exist and can be associated to a different Device Connector Manager subsequently.

---

# Resource: intersight_asset_managed_device
Attributes for Managed Device in Intersight and it maintains the relationship to the Device Connector Manager (Intersight appliance platform type tagged with 'Intersight Assist') Device Registration. It is possible to add Managed Devices within an Account and later associated to DeviceConnectorManager. Once attached to Device Connector Manager, Device Connector for the Managed Device type is started on the OnPrem and status related to it is maintained in the ManagedDeviceStatus MO. If the Device Connector Manager is deleted, Managed Device continues to exist and can be associated to a different Device Connector Manager subsequently.

## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) -ManagedDevice to Account MO relationship.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `credential`:(Array with Maximum of one item) -Credentials to manage Managed Device.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `is_password_set`:(bool)
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `password`:(string)Password for the Managed Device.
  + `username`:(string)Username for the Managed Device. Format and restrictions are not enforced here but usually follow the ManagedDevice requirements.
* `device_connector_manager`:(Array with Maximum of one item) -Device Connector Manager (Intersight Appliance Device Connector tagged as 'Intersight Assist') within the asset Device Registration.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `device_type`:(string)Type of the Device such as VMware, Pure Storage.
* `is_enabled`:(bool)Device is Enabled/Disabled.
* `management_address`:(string)Management address of the device. It can be IPv4, IPv6 or FQDN.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name defined by the administrator for easier identification.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `registered_device`:(Array with Maximum of one item) -ManagedDevice once auto claimed within the asset Device Registration.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `status`:(Array with Maximum of one item) -Status of communication releated to Managed Device.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `cloud_port`:(int)Port used for the connection to the Cloud by the Device Connector for the Managed Device.
  + `connection_failure_reason`:(string)Maintains the reason for the failure of connection to the Device in case of connection failure.
  + `connection_status`:(string)Maintains the status of the connection to the Device.
  + `error_code`:(int)Maintains code related to error from Device Connector, if any.
  + `error_reason`:(string)Maintains the reason for the error from Device Connector, if any.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `process_id`:(int)Maintains the Process pid of the the Device Connector for the Managed Device.
  + `server_port`:(int)Port used for receiving requests from Device Connector Manager by the Device Connector for the Managed Device.
  + `state`:(string)Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
