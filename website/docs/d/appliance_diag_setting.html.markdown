---
layout: "intersight"
page_title: "Intersight: intersight_appliance_diag_setting"
sidebar_current: "docs-intersight-data-source-applianceDiagSetting"
description: |-
DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
---

# Data Source: intersight_appliance_diag_setting
DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_password_set`:(bool)
* `message`:(string)Status message of the password change operation.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `password`:(string)Password of the Intersight Appliance's OS diagnostic user account.
