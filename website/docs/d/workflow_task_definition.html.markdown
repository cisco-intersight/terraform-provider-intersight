---
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_definition"
sidebar_current: "docs-intersight-data-source-workflowTaskDefinition"
description: |-
Used to define a task which can be included within a workflow. Task definition conveys the intent that we want to achieve with the task. We can have a standalone task definition that is bound to a single implementation for that task, or we can define an TaskDefinition that will serve as the interface task definition which is linked to multiple implementation tasks. Each implemented TaskDefinition will be bound to its own implementation so we can achieve a case where single TaskDefinition has multiple implementations.
---

# Data Source: intersight_workflow_task_definition
Used to define a task which can be included within a workflow. Task definition conveys the intent that we want to achieve with the task. We can have a standalone task definition that is bound to a single implementation for that task, or we can define an TaskDefinition that will serve as the interface task definition which is linked to multiple implementation tasks. Each implemented TaskDefinition will be bound to its own implementation so we can achieve a case where single TaskDefinition has multiple implementations.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `default_version`:(bool)When true this will be the task version that is used when a specific task definition version is not specified. The very first task definition created with a name will be set as the default version, after that user can explicitly set any version of the task definition as the default version.
* `description`:(string)The task definition description to describe what this task will do when executed.
* `label`:(string)A user friendly short name to identify the task definition.
* `license_entitlement`:(string)License entitlement required to run this task. It is determined by license requirement of features.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the task definition. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \"Generic\" if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `secure_prop_access`:(bool)If set to true, the task requires access to secure properties and uses an encyption token associated with a workflow moid to encrypt or decrypt the secure properties.
* `version`:(int)The version of the task definition so we can support multiple versions of a task definition.
