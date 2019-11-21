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
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
