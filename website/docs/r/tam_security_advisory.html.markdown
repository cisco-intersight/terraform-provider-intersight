---
layout: "intersight"
page_title: "Intersight: intersight_tam_security_advisory"
sidebar_current: "docs-intersight-resource-tamSecurityAdvisory"
description: |-
  Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
---

# Resource: intersight_tam_security_advisory
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
## Argument Reference
The following arguments are supported:
* `api_data_sources`:(Array)"An array of data sources that are used to provide data for queries used to identify an Intersight alert applicability."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `mo_type`:(string)"Type of Intersight managed object used as data source."
  + `name`:(string)"Name is used to unique identify and refer a given data source in an alert definition."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `queries`:(Array)"Optional set of Queries to filter the output for Api datasource. the queries are executed in the order specified."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `name`:(string)"Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source."
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `priority`:(int)"An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection."
    + `query`:(string)"A SparkSQL query to be used on a given data source."
  + `type`:(string)"Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.)."
* `actions`:(Array)"An array of actions that are to be taken when a given managed object matches the criteria specified for being affected by an alert definition."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `affected_object_type`:(string)"Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove)."
  + `alert_type`:(string)"Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.)."
  + `identifiers`:(Array)"Identifiers represents the filter criteria (property names and values) used to identify an Intersight managed object of type specified in affectedObjectType property. An instance of an alert is then create on (or removed from) the identified managed object."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `name`:(string)"Name of the filter paramter."
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `value`:(string)"Value of the filter paramter."
  + `name`:(string)"Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `operation_type`:(string)"Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied."
  + `queries`:(Array)"Set of SparkSQL queries used determine if a given alert is applicable or not. Refer to https://spark.apache.org/sql/ for more details."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `name`:(string)"Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source."
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `priority`:(int)"An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection."
    + `query`:(string)"A SparkSQL query to be used on a given data source."
  + `type`:(string)"Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type."
* `advisory_id`:(string)"Cisco generated identifier for the published security advisory."
* `base_score`:(float)"CVSS version 3 base score for the security Advisory."
* `cve_ids`:
                (Array of schema.TypeString) -"CVE (https://cve.mitre.org/about/faqs.html) identifiers associated with the published security Advisory."
* `date_published`:(string)"Date when the security advisory was first published by Cisco."
* `date_updated`:(string)"Date when the security advisory was last updated by Cisco."
* `description`:(string)"Brief description of the advisory details."
* `environmental_score`:(float)"CVSS version 3 environmental score for the security Advisory."
* `external_url`:(string)"A link to an external URL describing security Advisory in more details."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"A user defined name for the Intersight Advisory."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `organization`:(Array with Maximum of one item) -"Relationship to the Organization that owns the Managed Object."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `recommendation`:(string)"Recommended action to resolve the security advisory."
* `severity`:(Array with Maximum of one item) -"Severity level of the Intersight Advisory."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
* `state`:(string)"Current state of the advisory. Indicates if the user is interested in getting updates for the advisory."
* `status`:(string)"Cisco assigned status of the published advisory based on whether the investigation is complete or on-going."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
* `temporal_score`:(float)"CVSS version 3 temporal score for the security Advisory."
* `version`:(string)"Cisco assigned advisory version after latest revision."
