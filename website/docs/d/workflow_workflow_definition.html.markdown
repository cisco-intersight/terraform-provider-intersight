---
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_definition"
sidebar_current: "docs-intersight-data-source-workflowWorkflowDefinition"
description: |-
Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.

---

# Data Source: intersight_workflow_workflow_definition
Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `default_version`:(bool)When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version.
* `description`:(string)The description for this workflow.
* `label`:(string)A user friendly short name to identify the workflow.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name for this workflow. You can have multiple version of the workflow with the same name.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `version`:(int)The version of the workflow to support multiple versions.
