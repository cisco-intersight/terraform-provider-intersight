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
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `port`:(int)LDAP Server Port for connection establishment.
* `server`:(string)LDAP Server Address, can be IP address or hostname.
