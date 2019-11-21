---
layout: "intersight"
page_title: "Intersight: intersight_iam_idp_reference"
sidebar_current: "docs-intersight-data-source-iamIdpReference"
description: |-
Default Cisco IdP for authentication.

---

# Data Source: intersight_iam_idp_reference
Default Cisco IdP for authentication.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `domain_name`:(string)The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
* `idp_entity_id`:(string)Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `multi_factor_authentication`:(bool)The flag represents if the second factor of authentication is required for Cisco IdP users.
* `name`:(string)Cisco IdP reference in an account.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
