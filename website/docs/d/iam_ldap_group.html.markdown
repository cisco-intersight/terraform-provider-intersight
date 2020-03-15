---
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_group"
sidebar_current: "docs-intersight-data-source-iamLdapGroup"
description: |-
Mapping of LDAP Group to EndPointRoles.
---

# Data Source: intersight_iam_ldap_group
Mapping of LDAP Group to EndPointRoles.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `domain`:(string)LDAP server domain the Group resides in.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)LDAP Group name in the LDAP server database.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
