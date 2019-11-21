---
layout: "intersight"
page_title: "Intersight: intersight_workflow_batch_api_executor"
sidebar_current: "docs-intersight-data-source-workflowBatchApiExecutor"
description: |-
Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.

Batch API associates the list of API requests to be executed as part of single
task execution. Each API request takes the request body and a response parser
specification.

---

# Data Source: intersight_workflow_batch_api_executor
Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.

Batch API associates the list of API requests to be executed as part of single
task execution. Each API request takes the request body and a response parser
specification.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `constraints`:(string)Enter the constraints on when this task should match against the task definition.
* `description`:(string)A detailed description about the batch APIs.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name for the batch API task.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `skip_on_condition`:(string)The skip expression, if provided, allows the batch API executor to skip thetask execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed.
