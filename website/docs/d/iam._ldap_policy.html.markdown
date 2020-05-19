---
layout: "intersight"
page_title: "Intersight: intersight_iam._ldap_policy"
sidebar_current: "docs-intersight-data-source-iam.LdapPolicy"
description: |-
LDAP Policy configurations.
---

# Data Source: intersight_iam._ldap_policy
LDAP Policy configurations.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `description`:(string)"Description of the policy."
* `enable_dns`:(bool)"Enables DNS to access LDAP servers."
* `enabled`:(bool)"LDAP server performs authentication."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `user_search_precedence`:(string)"Search precedence between local user database and LDAP user database."
