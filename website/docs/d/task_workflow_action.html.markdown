
---
layout: "intersight"
page_title: "Intersight: intersight_task_workflow_action"
sidebar_current: "docs-intersight-data-source-task-workflow-action"
description: |-
Start a test workflow using this object.
---

# Data Source: intersight_task._workflow_action
Start a test workflow using this object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) Action for test workflow. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `input_params`:(string) Json formatted string input parameters to workflow. 
* `is_dynamic`:(bool) When true, this workflow type will be triggered as a dynamic workflow, if not it will be treated as a static workflow. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `wait_on_duplicate`:(bool) When true, the workflow will wait for previous one to complete before starting a new one. 
* `workflow_context`:(string) Json formatted string that has the contents of the workflow context used when starting a workflow. 
