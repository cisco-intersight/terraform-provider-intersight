---
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_info"
sidebar_current: "docs-intersight-data-source-iaasUcsdInfo"
description: |-
UCS Director accounts managed by Intersight.
---

# Data Source: intersight_iaas_ucsd_info
UCS Director accounts managed by Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_id`:(string)"Moid of the UCSD device connector's asset.DeviceRegistration."
* `guid`:(string)"Unique ID of UCSD getting registerd with Intersight."
* `host_name`:(string)"The UCSD host name."
* `ip`:(string)"The UCSD IP address."
* `last_backup`:(string)"Last successful backup created for this UCS Director appliance if backup is configured."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `node_type`:(string)"NodeType specifies if UCSD is deployed in Stand-alone or Multi Node."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `product_name`:(string)"The UCSD product name."
* `product_vendor`:(string)"The UCSD product vendor."
* `product_version`:(string)"The UCSD product/platform version."
* `status`:(string)"The UCSD status. Possible values are Active, Inactive, Unknown."
