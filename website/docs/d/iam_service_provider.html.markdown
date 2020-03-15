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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
