---
layout: "intersight"
page_title: "Intersight: intersight_tam_security_advisory"
sidebar_current: "docs-intersight-data-source-tamSecurityAdvisory"
description: |-
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
---

# Data Source: intersight_tam_security_advisory
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advisory_id`:(string)Cisco generated identifier for the published security advisory.
* `base_score`:(float)CVSS version 3 base score for the security Advisory.
* `date_published`:(string)Date when the security advisory was first published by Cisco.
* `date_updated`:(string)Date when the security advisory was last updated by Cisco.
* `description`:(string)Brief description of the advisory details.
* `environmental_score`:(float)CVSS version 3 environmental score for the security Advisory.
* `external_url`:(string)A link to an external URL describing security Advisory in more details.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)A user defined name for the Intersight Advisory.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `recommendation`:(string)Recommended action to resolve the security advisory.
* `state`:(string)Current state of the advisory. Indicates if the user is interested in getting updates for the advisory.
* `status`:(string)Cisco assigned status of the published advisory based on whether the investigation is complete or on-going.
* `temporal_score`:(float)CVSS version 3 temporal score for the security Advisory.
* `version`:(string)Cisco assigned advisory version after latest revision.
