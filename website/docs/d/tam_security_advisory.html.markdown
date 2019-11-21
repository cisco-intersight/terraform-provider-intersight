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
* `cve_id`:(string)CVE (https://cve.mitre.org/about/faqs.html) identifier for the security Advisory.
* `description`:(string)Brief description of the advisory details.
* `external_url`:(string)A link to an external URL describing security Advisory in more details.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)A user defined name for the Intersight Advisory.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `recommendation`:(string)Recommended action to resolve the security advisory.
* `score`:(float)CVSS score for the security Advisory.
* `severity`:(string)Severity level of the Intersight Advisory.
* `state`:(string)Current state of the advisory. Indicates if the user is interested in getting updates for the advisory.
