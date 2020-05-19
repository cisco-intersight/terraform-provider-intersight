package intersight

import (
	"log"
	"time"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWorkflowWorkflowInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowWorkflowInfoCreate,
		Read:   resourceWorkflowWorkflowInfoRead,
		Update: resourceWorkflowWorkflowInfoUpdate,
		Delete: resourceWorkflowWorkflowInfoDelete,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"action": {
				Description: "The action of the workflow such as start, cancel, retry, pause.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "None",
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cleanup_time": {
				Description: "The time when the workflow info will be removed from database.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "The time when the workflow reached a final state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"failed_workflow_cleanup_duration": {
				Description: "The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database.",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"input": {
				Description: "All the given inputs for the workflow.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"inst_id": {
				Description: "A workflow instance Id which is the unique identified for the workflow execution.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"internal": {
				Description: "Denotes if this workflow is internal and should be hidden from user view of running workflows.",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"last_action": {
				Description: "The last action that was issued on the workflow is saved in this field.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"message": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"message": {
							Description: "An i18n message that can be translated in multiple languages to support internationalization.\ntype: string",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"severity": {
							Description: "The severity of the Task or Workflow message warning/error/info etc.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Info",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"meta_version": {
				Description: "Version of the workflow metadata for which this workflow execution was started.",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "A name of the workflow execution instance.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"output": {
				Description: "All the generated outputs for the workflow.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
				Computed: true,
			},
			"parent_task_info": {
				Description: "A reference to a workflowTaskInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"pending_dynamic_workflow_info": {
				Description: "A reference to a workflowPendingDynamicWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"permission": {
				Description: "A reference to a iamPermission resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"progress": {
				Description: "This field indicates percentage of workflow task execution.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"properties": {
				Description: "Type to capture all the properties for the workflow info passed on from workflow definition.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"retryable": {
							Description: "When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.",
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"retry_from_task_name": {
				Description: "This field is applicable when Retry action is issued for a workflow which is in a final state. When this field is not specified then the workflow will retry from the start of the workflow. When this field is specified then the workflow will be retried from the specified task. The field should carry the task name which is the unique name of the task within the workflow. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"src": {
				Description: "The source microservice name which is the owner for this workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"start_time": {
				Description: "The time when the workflow was started for execution.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"success_workflow_cleanup_duration": {
				Description: "The duration in hours after which the workflow info for successful workflow will be removed from database.",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"task_infos": {
				Description: "An array of relationships to workflowTaskInfo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"trace_id": {
				Description: "The trace id to keep track of workflow execution.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "A type of the workflow (serverconfig, ansible_monitoring).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"user_id": {
				Description: "The user identifier which indicates the user that started this workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"wait_reason": {
				Description: "Denotes the reason workflow is in waiting status.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "None",
				ForceNew:    true,
			},
			"workflow_ctx": {
				Description: "The workflow context which contains initiator and target information.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"workflow_definition": {
				Description: "A reference to a workflowWorkflowDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"workflow_meta_type": {
				Description: "The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "SystemDefined",
			},
			"workflow_task_count": {
				Description: "Total number of workflow tasks in this workflow.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWorkflowWorkflowInfoCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowWorkflowInfo()
	if v, ok := d.GetOk("account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("iam.Account")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsIamAccountRelationship())
		}
		x := p[0]
		o.SetAccount(x)
	}

	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}

	o.SetClassId("workflow.WorkflowInfo")

	if v, ok := d.GetOk("cleanup_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCleanupTime(x)
	}

	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}

	if v, ok := d.GetOk("failed_workflow_cleanup_duration"); ok {
		x := int64(v.(int))
		o.SetFailedWorkflowCleanupDuration(x)
	}

	if v, ok := d.GetOk("input"); ok {
		x := v.(map[string]interface{})
		o.SetInput(x)
	}

	if v, ok := d.GetOk("inst_id"); ok {
		x := (v.(string))
		o.SetInstId(x)
	}

	if v, ok := d.GetOk("internal"); ok {
		x := (v.(bool))
		o.SetInternal(x)
	}

	if v, ok := d.GetOk("last_action"); ok {
		x := (v.(string))
		o.SetLastAction(x)
	}

	if v, ok := d.GetOk("message"); ok {
		x := make([]models.WorkflowMessage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowMessageWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("workflow.Message")
			if v, ok := l["message"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetMessage(x)
				}
			}
			o.SetObjectType("workflow.Message")
			if v, ok := l["severity"]; ok {
				{
					x := (v.(string))
					o.SetSeverity(x)
				}
			}
			x = append(x, *o)
		}
		o.SetMessage(x)
	}

	if v, ok := d.GetOk("meta_version"); ok {
		x := int64(v.(int))
		o.SetMetaVersion(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("workflow.WorkflowInfo")

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("organization.Organization")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsOrganizationOrganizationRelationship())
		}
		x := p[0]
		o.SetOrganization(x)
	}

	if v, ok := d.GetOk("output"); ok {
		x := v.(map[string]interface{})
		o.SetOutput(x)
	}

	if v, ok := d.GetOk("parent_task_info"); ok {
		p := make([]models.WorkflowTaskInfoRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskInfo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowTaskInfoRelationship())
		}
		x := p[0]
		o.SetParentTaskInfo(x)
	}

	if v, ok := d.GetOk("pending_dynamic_workflow_info"); ok {
		p := make([]models.WorkflowPendingDynamicWorkflowInfoRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.PendingDynamicWorkflowInfo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowPendingDynamicWorkflowInfoRelationship())
		}
		x := p[0]
		o.SetPendingDynamicWorkflowInfo(x)
	}

	if v, ok := d.GetOk("permission"); ok {
		p := make([]models.IamPermissionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("iam.Permission")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsIamPermissionRelationship())
		}
		x := p[0]
		o.SetPermission(x)
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("progress"); ok {
		x := v.(float32)
		o.SetProgress(x)
	}

	if v, ok := d.GetOk("properties"); ok {
		p := make([]models.WorkflowWorkflowInfoProperties, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowWorkflowInfoPropertiesWithDefaults()
			o.SetClassId("workflow.WorkflowInfoProperties")
			o.SetObjectType("workflow.WorkflowInfoProperties")
			if v, ok := l["retryable"]; ok {
				{
					x := (v.(bool))
					o.SetRetryable(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetProperties(x)
	}

	if v, ok := d.GetOk("retry_from_task_name"); ok {
		x := (v.(string))
		o.SetRetryFromTaskName(x)
	}

	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.SetSrc(x)
	}

	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	if v, ok := d.GetOk("success_workflow_cleanup_duration"); ok {
		x := int64(v.(int))
		o.SetSuccessWorkflowCleanupDuration(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOk("task_infos"); ok {
		x := make([]models.WorkflowTaskInfoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskInfo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsWorkflowTaskInfoRelationship())
		}
		o.SetTaskInfos(x)
	}

	if v, ok := d.GetOk("trace_id"); ok {
		x := (v.(string))
		o.SetTraceId(x)
	}

	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	if v, ok := d.GetOk("user_id"); ok {
		x := (v.(string))
		o.SetUserId(x)
	}

	if v, ok := d.GetOk("wait_reason"); ok {
		x := (v.(string))
		o.SetWaitReason(x)
	}

	if v, ok := d.GetOk("workflow_ctx"); ok {
		x := v.(map[string]interface{})
		o.SetWorkflowCtx(x)
	}

	if v, ok := d.GetOk("workflow_definition"); ok {
		p := make([]models.WorkflowWorkflowDefinitionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.WorkflowDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowWorkflowDefinitionRelationship())
		}
		x := p[0]
		o.SetWorkflowDefinition(x)
	}

	if v, ok := d.GetOk("workflow_meta_type"); ok {
		x := (v.(string))
		o.SetWorkflowMetaType(x)
	}

	if v, ok := d.GetOk("workflow_task_count"); ok {
		x := int64(v.(int))
		o.SetWorkflowTaskCount(x)
	}

	r := conn.ApiClient.WorkflowApi.CreateWorkflowWorkflowInfo(conn.ctx).WorkflowWorkflowInfo(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowWorkflowInfoRead(d, meta)
}

func resourceWorkflowWorkflowInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowInfoByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("account", flattenMapIamAccountRelationship(s.Account, d)); err != nil {
		return err
	}

	if err := d.Set("action", (s.Action)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("cleanup_time", (s.CleanupTime).String()); err != nil {
		return err
	}

	if err := d.Set("end_time", (s.EndTime).String()); err != nil {
		return err
	}

	if err := d.Set("failed_workflow_cleanup_duration", (s.FailedWorkflowCleanupDuration)); err != nil {
		return err
	}

	if err := d.Set("input", (s.Input)); err != nil {
		return err
	}

	if err := d.Set("inst_id", (s.InstId)); err != nil {
		return err
	}

	if err := d.Set("internal", (s.Internal)); err != nil {
		return err
	}

	if err := d.Set("last_action", (s.LastAction)); err != nil {
		return err
	}

	if err := d.Set("message", flattenListWorkflowMessage(s.Message, d)); err != nil {
		return err
	}

	if err := d.Set("meta_version", (s.MetaVersion)); err != nil {
		return err
	}

	if err := d.Set("moid", (s.Moid)); err != nil {
		return err
	}

	if err := d.Set("name", (s.Name)); err != nil {
		return err
	}

	if err := d.Set("object_type", (s.ObjectType)); err != nil {
		return err
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("output", (s.Output)); err != nil {
		return err
	}

	if err := d.Set("parent_task_info", flattenMapWorkflowTaskInfoRelationship(s.ParentTaskInfo, d)); err != nil {
		return err
	}

	if err := d.Set("pending_dynamic_workflow_info", flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(s.PendingDynamicWorkflowInfo, d)); err != nil {
		return err
	}

	if err := d.Set("permission", flattenMapIamPermissionRelationship(s.Permission, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("progress", (s.Progress)); err != nil {
		return err
	}

	if err := d.Set("properties", flattenMapWorkflowWorkflowInfoProperties(s.Properties, d)); err != nil {
		return err
	}

	if err := d.Set("retry_from_task_name", (s.RetryFromTaskName)); err != nil {
		return err
	}

	if err := d.Set("src", (s.Src)); err != nil {
		return err
	}

	if err := d.Set("start_time", (s.StartTime).String()); err != nil {
		return err
	}

	if err := d.Set("status", (s.Status)); err != nil {
		return err
	}

	if err := d.Set("success_workflow_cleanup_duration", (s.SuccessWorkflowCleanupDuration)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("task_infos", flattenListWorkflowTaskInfoRelationship(s.TaskInfos, d)); err != nil {
		return err
	}

	if err := d.Set("trace_id", (s.TraceId)); err != nil {
		return err
	}

	if err := d.Set("type", (s.Type)); err != nil {
		return err
	}

	if err := d.Set("user_id", (s.UserId)); err != nil {
		return err
	}

	if err := d.Set("wait_reason", (s.WaitReason)); err != nil {
		return err
	}

	if err := d.Set("workflow_ctx", (s.WorkflowCtx)); err != nil {
		return err
	}

	if err := d.Set("workflow_definition", flattenMapWorkflowWorkflowDefinitionRelationship(s.WorkflowDefinition, d)); err != nil {
		return err
	}

	if err := d.Set("workflow_meta_type", (s.WorkflowMetaType)); err != nil {
		return err
	}

	if err := d.Set("workflow_task_count", (s.WorkflowTaskCount)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceWorkflowWorkflowInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowWorkflowInfo()
	if d.HasChange("account") {
		v := d.Get("account")
		p := make([]models.IamAccountRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("iam.Account")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsIamAccountRelationship())
		}
		x := p[0]
		o.SetAccount(x)
	}

	if d.HasChange("action") {
		v := d.Get("action")
		x := (v.(string))
		o.SetAction(x)
	}

	if d.HasChange("cleanup_time") {
		v := d.Get("cleanup_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCleanupTime(x)
	}

	if d.HasChange("end_time") {
		v := d.Get("end_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}

	if d.HasChange("failed_workflow_cleanup_duration") {
		v := d.Get("failed_workflow_cleanup_duration")
		x := int64(v.(int))
		o.SetFailedWorkflowCleanupDuration(x)
	}

	if d.HasChange("input") {
		v := d.Get("input")
		x := v.(map[string]interface{})
		o.SetInput(x)
	}

	if d.HasChange("inst_id") {
		v := d.Get("inst_id")
		x := (v.(string))
		o.SetInstId(x)
	}

	if d.HasChange("internal") {
		v := d.Get("internal")
		x := (v.(bool))
		o.SetInternal(x)
	}

	if d.HasChange("last_action") {
		v := d.Get("last_action")
		x := (v.(string))
		o.SetLastAction(x)
	}

	if d.HasChange("message") {
		v := d.Get("message")
		x := make([]models.WorkflowMessage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowMessageWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("workflow.Message")
			if v, ok := l["message"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetMessage(x)
				}
			}
			o.SetObjectType("workflow.Message")
			if v, ok := l["severity"]; ok {
				{
					x := (v.(string))
					o.SetSeverity(x)
				}
			}
			x = append(x, *o)
		}
		o.SetMessage(x)
	}

	if d.HasChange("meta_version") {
		v := d.Get("meta_version")
		x := int64(v.(int))
		o.SetMetaVersion(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("organization.Organization")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsOrganizationOrganizationRelationship())
		}
		x := p[0]
		o.SetOrganization(x)
	}

	if d.HasChange("output") {
		v := d.Get("output")
		x := v.(map[string]interface{})
		o.SetOutput(x)
	}

	if d.HasChange("parent_task_info") {
		v := d.Get("parent_task_info")
		p := make([]models.WorkflowTaskInfoRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskInfo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowTaskInfoRelationship())
		}
		x := p[0]
		o.SetParentTaskInfo(x)
	}

	if d.HasChange("pending_dynamic_workflow_info") {
		v := d.Get("pending_dynamic_workflow_info")
		p := make([]models.WorkflowPendingDynamicWorkflowInfoRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.PendingDynamicWorkflowInfo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowPendingDynamicWorkflowInfoRelationship())
		}
		x := p[0]
		o.SetPendingDynamicWorkflowInfo(x)
	}

	if d.HasChange("permission") {
		v := d.Get("permission")
		p := make([]models.IamPermissionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("iam.Permission")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsIamPermissionRelationship())
		}
		x := p[0]
		o.SetPermission(x)
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if d.HasChange("progress") {
		v := d.Get("progress")
		x := v.(float32)
		o.SetProgress(x)
	}

	if d.HasChange("properties") {
		v := d.Get("properties")
		p := make([]models.WorkflowWorkflowInfoProperties, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowWorkflowInfoPropertiesWithDefaults()
			o.SetClassId("workflow.WorkflowInfoProperties")
			o.SetObjectType("workflow.WorkflowInfoProperties")
			if v, ok := l["retryable"]; ok {
				{
					x := (v.(bool))
					o.SetRetryable(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetProperties(x)
	}

	if d.HasChange("retry_from_task_name") {
		v := d.Get("retry_from_task_name")
		x := (v.(string))
		o.SetRetryFromTaskName(x)
	}

	if d.HasChange("src") {
		v := d.Get("src")
		x := (v.(string))
		o.SetSrc(x)
	}

	if d.HasChange("start_time") {
		v := d.Get("start_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}

	if d.HasChange("status") {
		v := d.Get("status")
		x := (v.(string))
		o.SetStatus(x)
	}

	if d.HasChange("success_workflow_cleanup_duration") {
		v := d.Get("success_workflow_cleanup_duration")
		x := int64(v.(int))
		o.SetSuccessWorkflowCleanupDuration(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if d.HasChange("task_infos") {
		v := d.Get("task_infos")
		x := make([]models.WorkflowTaskInfoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskInfo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsWorkflowTaskInfoRelationship())
		}
		o.SetTaskInfos(x)
	}

	if d.HasChange("trace_id") {
		v := d.Get("trace_id")
		x := (v.(string))
		o.SetTraceId(x)
	}

	if d.HasChange("type") {
		v := d.Get("type")
		x := (v.(string))
		o.SetType(x)
	}

	if d.HasChange("user_id") {
		v := d.Get("user_id")
		x := (v.(string))
		o.SetUserId(x)
	}

	if d.HasChange("wait_reason") {
		v := d.Get("wait_reason")
		x := (v.(string))
		o.SetWaitReason(x)
	}

	if d.HasChange("workflow_ctx") {
		v := d.Get("workflow_ctx")
		x := v.(map[string]interface{})
		o.SetWorkflowCtx(x)
	}

	if d.HasChange("workflow_definition") {
		v := d.Get("workflow_definition")
		p := make([]models.WorkflowWorkflowDefinitionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.WorkflowDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowWorkflowDefinitionRelationship())
		}
		x := p[0]
		o.SetWorkflowDefinition(x)
	}

	if d.HasChange("workflow_meta_type") {
		v := d.Get("workflow_meta_type")
		x := (v.(string))
		o.SetWorkflowMetaType(x)
	}

	if d.HasChange("workflow_task_count") {
		v := d.Get("workflow_task_count")
		x := int64(v.(int))
		o.SetWorkflowTaskCount(x)
	}

	r := conn.ApiClient.WorkflowApi.UpdateWorkflowWorkflowInfo(conn.ctx, d.Id()).WorkflowWorkflowInfo(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowWorkflowInfoRead(d, meta)
}

func resourceWorkflowWorkflowInfoDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.WorkflowApi.DeleteWorkflowWorkflowInfo(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
