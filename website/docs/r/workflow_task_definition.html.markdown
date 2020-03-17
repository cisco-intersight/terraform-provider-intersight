---
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_definition"
sidebar_current: "docs-intersight-resource-workflowTaskDefinition"
description: |-
  Used to define a task which can be included within a workflow. Task definition conveys the intent that we want to achieve with the task. We can have a standalone task definition that is bound to a single implementation for that task, or we can define an TaskDefinition that will serve as the interface task definition which is linked to multiple implementation tasks. Each implemented TaskDefinition will be bound to its own implementation so we can achieve a case where single TaskDefinition has multiple implementations.
---

# Resource: intersight_workflow_task_definition
Used to define a task which can be included within a workflow. Task definition conveys the intent that we want to achieve with the task. We can have a standalone task definition that is bound to a single implementation for that task, or we can define an TaskDefinition that will serve as the interface task definition which is linked to multiple implementation tasks. Each implemented TaskDefinition will be bound to its own implementation so we can achieve a case where single TaskDefinition has multiple implementations.
## Argument Reference
The following arguments are supported:
* `catalog`:(Array with Maximum of one item) -"The catalog under which the definition has been added."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `default_version`:(bool)"When true this will be the task version that is used when a specific task definition version is not specified. The very first task definition created with a name will be set as the default version, after that user can explicitly set any version of the task definition as the default version."
* `description`:(string)"The task definition description to describe what this task will do when executed."
* `implemented_tasks`:(Array)"List of all the implemented task for this TaskDefinition. When this list is populated it implies that this TaskDefinition has multiple implementations."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `interface_task`:(Array with Maximum of one item) -"A collection of references to the [workflow.TaskDefinition](mo://workflow.TaskDefinition) Managed Object.\nWhen this managed object is deleted, the referenced [workflow.TaskDefinition](mo://workflow.TaskDefinition) MO unsets its reference to this deleted MO."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `internal_properties`:(Array with Maximum of one item) -(Computed)"Type to capture all the internal properties for the task definition."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `base_task_type`:(string)(Computed)"This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask."
  + `constraints`:(Array with Maximum of one item) -(Computed)"This field will hold any constraints a concrete task definition will specify in order to limit the environment where the task can execute."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `target_data_type`:"List of property constraints that helps to narrow down task implementations based on target device input."
  + `internal`:(bool)(Computed)"Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `owner`:(string)(Computed)"The service that owns and is responsible for execution of the task."
* `label`:(string)"A user friendly short name to identify the task definition."
* `license_entitlement`:(string)(Computed)"License entitlement required to run this task. It is determined by license requirement of features."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"The name of the task definition. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \"Generic\" if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `properties`:(Array with Maximum of one item) -"Type to capture all the properties for the task definition."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `external_meta`:(bool)"When set to false the task definition can only be used by internal system workflows. When set to true then the task can be included in user defined workflows."
  + `input_definition`:(Array)"The schema expected for input parameters for this task."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `default`:(Array with Maximum of one item) -"Default value for the data type. If default value was provided and the input was required the default value will be used as the input."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `override`:(bool)"Override the default value provided for the data type. When true, allow the user to enter value for the data type."
    + `value`:"Default value for the data type. If default value was provided and the input was required the default value will be used as the input."
  + `description`:(string)"Provide a detailed description of the data type."
  + `label`:(string)"Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character."
  + `name`:(string)"Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `required`:(bool)"Specifies whether this parameter is required. The field is applicable for task and workflow."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `output_definition`:(Array)"The schema expected for output parameters for this task."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `default`:(Array with Maximum of one item) -"Default value for the data type. If default value was provided and the input was required the default value will be used as the input."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `override`:(bool)"Override the default value provided for the data type. When true, allow the user to enter value for the data type."
    + `value`:"Default value for the data type. If default value was provided and the input was required the default value will be used as the input."
  + `description`:(string)"Provide a detailed description of the data type."
  + `label`:(string)"Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character."
  + `name`:(string)"Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `required`:(bool)"Specifies whether this parameter is required. The field is applicable for task and workflow."
  + `retry_count`:(int)"The number of times a task should be tried before marking as failed."
  + `retry_delay`:(int)"The delay in seconds after which the the task is re-tried."
  + `retry_policy`:(string)"The retry policy for the task."
  + `support_status`:(string)"Supported status of the definition."
  + `timeout`:(int)"The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days."
  + `timeout_policy`:(string)"The timeout policy for the task."
* `secure_prop_access`:(bool)"If set to true, the task requires access to secure properties and uses an encyption token associated with a workflow moid to encrypt or decrypt the secure properties."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
* `version`:(int)"The version of the task definition so we can support multiple versions of a task definition."
