
---
layout: "intersight"
page_title: "Intersight: intersight_tam_security_advisory"
sidebar_current: "docs-intersight-resource-tam-security-advisory"
description: |-
  Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
---

# Resource: intersight_tam._security_advisory
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
## Argument Reference
The following arguments are supported:
* `actions`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `affected_object_type`:(string) Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove). 
  + `alert_type`:(string) Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.). 
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `identifiers`:(Array)
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `name`:(string) Name of the filter paramter. 
    + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
    + `value`:(string) Value of the filter paramter. 
  + `name`:(string) Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `operation_type`:(string) Operation type for the alert action. An action is used to carry out the process of \ reacting\  to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied. 
  + `queries`:(Array)
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `name`:(string) Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. 
    + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
    + `priority`:(int) An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. 
    + `query`:(string) A SparkSQL query to be used on a given data source. 
  + `type`:(string) Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type. 
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `advisory_id`:(string) Cisco generated identifier for the published security advisory. 
* `api_data_sources`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `mo_type`:(string) Type of Intersight managed object used as data source. 
  + `name`:(string) Name is used to unique identify and refer a given data source in an alert definition. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `queries`:(Array)
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
    + `name`:(string) Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. 
    + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
    + `priority`:(int) An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. 
    + `query`:(string) A SparkSQL query to be used on a given data source. 
  + `type`:(string) Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.). 
* `base_score`:(float) CVSS version 3 base score for the security Advisory. 
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cve_ids`:
                (Array of schema.TypeString) -
* `date_published`:(string) Date when the security advisory was first published by Cisco. 
* `date_updated`:(string) Date when the security advisory was last updated by Cisco. 
* `description`:(string) Brief description of the advisory details. 
* `environmental_score`:(float) CVSS version 3 environmental score for the security Advisory. 
* `external_url`:(string) A link to an external URL describing security Advisory in more details. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user defined name for the Intersight Advisory. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `recommendation`:(string) Recommended action to resolve the security advisory. 
* `severity`:(Array with Maximum of one item) - Severity level of the Intersight Advisory. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
* `state`:(string) Current state of the advisory. 
* `status`:(string) Cisco assigned status of the published advisory based on whether the investigation is complete or on-going. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `temporal_score`:(float) CVSS version 3 temporal score for the security Advisory. 
* `version`:(string) Cisco assigned advisory version after latest revision. 
* `workaround`:(string) Workarounds available for the advisory. 
