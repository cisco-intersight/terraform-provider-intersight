---
layout: "intersight"
page_title: "Intersight: intersight_iam._ldap_group"
sidebar_current: "docs-intersight-data-source-iam.LdapGroup"
description: |-
Mapping of LDAP Group to EndPointRoles.
---

# Data Source: intersight_iam._ldap_group
Mapping of LDAP Group to EndPointRoles.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `domain`:(string)"LDAP server domain the Group resides in."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"LDAP Group name in the LDAP server database."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
