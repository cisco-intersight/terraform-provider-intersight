---
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_policy"
sidebar_current: "docs-intersight-data-source-iamLdapPolicy"
description: |-
LDAP Policy configurations.

---

# Data Source: intersight_iam_ldap_policy
LDAP Policy configurations.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `enable_dns`:(bool)Enables DNS to access LDAP servers.
* `enabled`:(bool)LDAP server performs authentication.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `user_search_precedence`:(string)Search precedence between local user database and LDAP user database.
