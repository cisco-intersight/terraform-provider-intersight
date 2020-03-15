---
layout: "intersight"
page_title: "Intersight: intersight_asset_device_registration"
sidebar_current: "docs-intersight-data-source-assetDeviceRegistration"
description: |-
DeviceRegistration represents a device connector enabled endpoint which has registered with Intersight.
---

# Data Source: intersight_asset_device_registration
DeviceRegistration represents a device connector enabled endpoint which has registered with Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `api_version`:(int)The version of the connector API, describes the capability of the connector's framework.If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.
* `access_key_id`:(string)An identifier for the credential used by the device connector to authenticate with the Intersight web socket gateway.
* `app_partition_number`:(int)The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.
* `claimed_by_user_name`:(string)The name of the user who claimed the device for the account.
* `claimed_time`:(string)The date and time at which the device was claimed to this account.
* `connection_id`:(string)The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.
* `connection_reason`:(string)If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with the cloud.
* `connection_status`:(string)The status of the persistent connection between the device connector and Intersight.
* `connection_status_last_change_time`:(string)The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between the cloud and device connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with the cloud.
* `connector_version`:(string)The version of the device connector running on the managed device.
* `device_external_ip_address`:(string)The IP Address of the managed device as seen from the cloud at the time of registration.This could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.
* `execution_mode`:(string)Indicates if the platform is an actual device or an emulated device for testing, demos, etc. Permitted values are [Normal, Emulator, ContainerEmulator].
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `platform_type`:(string)The platform type on which device connector is executing.
* `proxy_app`:(string)The name of the app which will proxy the messages to the device connector.
* `public_access_key`:(string)The device connector's public key used by the cloud to authenticate a connection from the device connector. The public key is used to verify that the signature a device connector sends on connect has been signed by the connector's private key stored on the device's filesystem. Must be a PEM encoded RSA public key string.
* `read_only`:(bool)Flag reported by devices to indicate an administrator of the device has disabled management operations of the device connector and only monitoring is permitted.
* `vendor`:(string)The vendor of the managed device.
