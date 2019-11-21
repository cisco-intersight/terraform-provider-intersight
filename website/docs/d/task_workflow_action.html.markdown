---
layout: "intersight"
page_title: "Intersight: intersight_task_workflow_action"
sidebar_current: "docs-intersight-data-source-taskWorkflowAction"
description: |-
Start a test workflow using this object.

---

# Data Source: intersight_task_workflow_action
Start a test workflow using this object.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string)Action for test workflow.
* `input_params`:(string)Json formatted string input parameters to workflow.
* `is_dynamic`:(bool)When true, this workflow type will be triggered as a dynamic workflow, if not it will be treated as a static workflow.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `wait_on_duplicate`:(bool)When true, the workflow will wait for previous one to complete before starting a new one.
* `workflow_context`:(string)Json formatted string that has the contents of the workflow context used when starting a workflow.
