---
layout: "intersight"
page_title: "Intersight: intersight_sdwan_vmanage_account_policy"
sidebar_current: "docs-intersight-data-source-sdwanVmanageAccountPolicy"
description: |-
A policy specifying vManage account details.
---

# Data Source: intersight_sdwan_vmanage_account_policy
A policy specifying vManage account details.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)"Description of the policy."
* `endpoint_address`:(string)"VManage server hostname (FQDN) that the acccount holds information for."
* `is_password_set`:(bool)
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `password`:(string)"Local password for authenticating with the vManage server."
* `port`:(int)"VManage Port number on which the application is running."
* `username`:(string)"Local username for authenticating with the vManage server."
