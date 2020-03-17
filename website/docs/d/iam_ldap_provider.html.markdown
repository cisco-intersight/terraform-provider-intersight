---
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_provider"
sidebar_current: "docs-intersight-data-source-iamLdapProvider"
description: |-
LDAP Provider or LDAP Server for user authentication.
---

# Data Source: intersight_iam_ldap_provider
LDAP Provider or LDAP Server for user authentication.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `port`:(int)"LDAP Server Port for connection establishment."
* `server`:(string)"LDAP Server Address, can be IP address or hostname."
