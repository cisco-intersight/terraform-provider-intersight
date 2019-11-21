---
layout: "intersight"
page_title: "Intersight: intersight_iam_service_provider"
sidebar_current: "docs-intersight-data-source-iamServiceProvider"
description: |-
SAML Service provider related information in Intersight.

---

# Data Source: intersight_iam_service_provider
SAML Service provider related information in Intersight.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `entity_id`:(string)Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider.
* `metadata`:(string)Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the Intersight Service Provider.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
