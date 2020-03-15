---
layout: "intersight"
page_title: "Intersight: intersight_asset_device_registration"
sidebar_current: "docs-intersight-resource-assetDeviceRegistration"
description: |-
  DeviceRegistration represents a device connector enabled endpoint which has registered with Intersight.
---

# Resource: intersight_asset_device_registration
DeviceRegistration represents a device connector enabled endpoint which has registered with Intersight.
## Argument Reference
The following arguments are supported:
* `api_version`:(int)(Computed)The version of the connector API, describes the capability of the connector's framework.If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.
* `access_key_id`:(string)An identifier for the credential used by the device connector to authenticate with the Intersight web socket gateway.
* `account`:(Array with Maximum of one item) -(Computed)The account to which the device has been claimed.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `app_partition_number`:(int)(Computed)The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.
* `claimed_by_user`:(Array with Maximum of one item) -(Computed)The user who claimed the device for the account.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `claimed_by_user_name`:(string)(Computed)The name of the user who claimed the device for the account.
* `claimed_time`:(string)(Computed)The date and time at which the device was claimed to this account.
* `cluster_members`:(Array)(Computed)The set of nodes within the devices cluster.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `connection_id`:(string)(Computed)The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.
* `connection_reason`:(string)(Computed)If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with the cloud.
* `connection_status`:(string)(Computed)The status of the persistent connection between the device connector and Intersight.
* `connection_status_last_change_time`:(string)(Computed)The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between the cloud and device connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with the cloud.
* `connector_version`:(string)(Computed)The version of the device connector running on the managed device.
* `device_claim`:(Array with Maximum of one item) -(Computed)The user who claimed the device for the account.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `device_configuration`:(Array with Maximum of one item) -(Computed)The devices current configuration.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `device_external_ip_address`:(string)(Computed)The IP Address of the managed device as seen from the cloud at the time of registration.This could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.
* `device_hostname`:
                (Array of schema.TypeString) -(Computed)The hostnames of the managed device. There can be multiple hostnames depending on the number of elements managed (i.e. HA clusters).
* `device_ip_address`:
                (Array of schema.TypeString) -(Computed)The IP Addresses of the managed device. There can be multiple management IPs depending on the number of elements managed (i.e. HA clusters) and in-band/out-of-band connectivity.
* `domain_group`:(Array with Maximum of one item) -(Computed)The domain group to which the device has been assigned.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `execution_mode`:(string)Indicates if the platform is an actual device or an emulated device for testing, demos, etc. Permitted values are [Normal, Emulator, ContainerEmulator].
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)(Computed)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `parent_connection`:(Array with Maximum of one item) -(Computed)The parent device of this device, through which this device connector is connected. If present the device will inherit its ownership through this object and any claim operation will be cascaded from it.e.g. A rack server may have a parent as the Fabric Interconnect cluster to which it is attached.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `parent_signature`:(Array with Maximum of one item) -A signature generated by a parent device used to authenticate that a leaf device is connected through the parent.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `device_id`:(string)The moid of the parent device registration.
  + `node_id`:(string)The node identity of the parent device, corresponds to the parents ClusterMember.memberIdentity. Used on connect to establish through which device in a cluster the child is connected through.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `signature`:(string)The result of signing the deviceId appended with the timeStamp fields with the devices private key.
* `permission_resources`:(Array)(Computed)A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `pid`:
                (Array of schema.TypeString) -(Computed)The product Id of the managed device.
* `platform_type`:(string)The platform type on which device connector is executing.
* `proxy_app`:(string)(Computed)The name of the app which will proxy the messages to the device connector.
* `public_access_key`:(string)(Computed)The device connector's public key used by the cloud to authenticate a connection from the device connector. The public key is used to verify that the signature a device connector sends on connect has been signed by the connector's private key stored on the device's filesystem. Must be a PEM encoded RSA public key string.
* `read_only`:(bool)(Computed)Flag reported by devices to indicate an administrator of the device has disabled management operations of the device connector and only monitoring is permitted.
* `security_token`:(Array with Maximum of one item) -(Computed)The devices current security token that can be used by a device administrator to claim the device.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)(Computed)The Object Type of the referenced REST resource.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `serial`:
                (Array of schema.TypeString) -(Computed)A list of serial numbers of the individual device elements (e.g. HA primary/secondary or cluster members) which are exposed as a single unit of management by the device connector.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
* `vendor`:(string)(Computed)The vendor of the managed device.
