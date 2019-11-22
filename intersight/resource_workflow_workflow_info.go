package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWorkflowWorkflowInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowWorkflowInfoCreate,
		Read:   resourceWorkflowWorkflowInfoRead,
		Update: resourceWorkflowWorkflowInfoUpdate,
		Delete: resourceWorkflowWorkflowInfoDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Description: "The action of the workflow such as start, cancel, retry, pause.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Start",
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
				ForceNew: true,
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
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "Relationship to the Organization that owns the Managed Object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
				Description: "Link to the task in the parent workflow which launched this workflow.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"pending_dynamic_workflow_info": {
				Description: "Reference to the PendingDynamicWorkflowInfo that was used to construct this workflow instance.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"progress": {
				Description: "This field indicates percentage of workflow task execution.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
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
				Description: "The array of tags, which allow to add key, value meta-data to managed objects.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"task_infos": {
				Description: "List of task instances that ran as part of this workflow execution.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
				ForceNew: true,
			},
			"workflow_definition": {
				Description: "The workflow definition that was used to launch this workflow execution instance.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
	var o models.WorkflowWorkflowInfo
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.Action = &x

	}

	if v, ok := d.GetOk("cleanup_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.CleanupTime = x

	}

	if v, ok := d.GetOk("end_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndTime = x

	}

	if v, ok := d.GetOk("failed_workflow_cleanup_duration"); ok {
		x := int64(v.(int))
		o.FailedWorkflowCleanupDuration = x

	}

	if v, ok := d.GetOk("input"); ok {
		x := v
		o.Input = &x

	}

	if v, ok := d.GetOk("inst_id"); ok {
		x := (v.(string))
		o.InstID = x

	}

	if v, ok := d.GetOkExists("internal"); ok {
		x := v.(bool)
		o.Internal = &x
	}

	if v, ok := d.GetOk("meta_version"); ok {
		x := int64(v.(int))
		o.MetaVersion = x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

	}

	if v, ok := d.GetOk("organization"); ok {
		p := models.IamAccountRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamAccountRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Organization = &x

	}

	if v, ok := d.GetOk("output"); ok {
		x := v
		o.Output = &x

	}

	if v, ok := d.GetOk("parent_task_info"); ok {
		p := models.WorkflowTaskInfoRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowTaskInfoRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.ParentTaskInfo = &x

	}

	if v, ok := d.GetOk("pending_dynamic_workflow_info"); ok {
		p := models.WorkflowPendingDynamicWorkflowInfoRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowPendingDynamicWorkflowInfoRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.PendingDynamicWorkflowInfo = &x

	}

	if v, ok := d.GetOk("progress"); ok {
		x := v.(float32)
		o.Progress = x

	}

	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.Src = x

	}

	if v, ok := d.GetOk("start_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.StartTime = x

	}

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = x

	}

	if v, ok := d.GetOk("success_workflow_cleanup_duration"); ok {
		x := int64(v.(int))
		o.SuccessWorkflowCleanupDuration = x

	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x

	}

	if v, ok := d.GetOk("task_infos"); ok {
		x := make([]*models.WorkflowTaskInfoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowTaskInfoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.TaskInfos = x

	}

	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.Type = x

	}

	if v, ok := d.GetOk("user_id"); ok {
		x := (v.(string))
		o.UserID = x

	}

	if v, ok := d.GetOk("wait_reason"); ok {
		x := (v.(string))
		o.WaitReason = &x

	}

	if v, ok := d.GetOk("workflow_ctx"); ok {
		x := v
		o.WorkflowCtx = &x

	}

	if v, ok := d.GetOk("workflow_definition"); ok {
		p := models.WorkflowWorkflowDefinitionRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowWorkflowDefinitionRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.WorkflowDefinition = &x

	}

	if v, ok := d.GetOk("workflow_meta_type"); ok {
		x := (v.(string))
		o.WorkflowMetaType = &x

	}

	if v, ok := d.GetOk("workflow_task_count"); ok {
		x := int64(v.(int))
		o.WorkflowTaskCount = x

	}

	url := "workflow/WorkflowInfos"
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceWorkflowWorkflowInfoRead(d, meta)
}

func resourceWorkflowWorkflowInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/WorkflowInfos" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.WorkflowWorkflowInfo
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("action", (s.Action)); err != nil {
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

	if err := d.Set("inst_id", (s.InstID)); err != nil {
		return err
	}

	if err := d.Set("internal", (s.Internal)); err != nil {
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

	if err := d.Set("organization", flattenMapIamAccountRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("output", (s.Output)); err != nil {
		return err
	}

	if err := d.Set("parent_task_info", flattenMapWorkflowTaskInfoRef(s.ParentTaskInfo, d)); err != nil {
		return err
	}

	if err := d.Set("pending_dynamic_workflow_info", flattenMapWorkflowPendingDynamicWorkflowInfoRef(s.PendingDynamicWorkflowInfo, d)); err != nil {
		return err
	}

	if err := d.Set("progress", (s.Progress)); err != nil {
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

	if err := d.Set("task_infos", flattenListWorkflowTaskInfoRef(s.TaskInfos, d)); err != nil {
		return err
	}

	if err := d.Set("type", (s.Type)); err != nil {
		return err
	}

	if err := d.Set("user_id", (s.UserID)); err != nil {
		return err
	}

	if err := d.Set("wait_reason", (s.WaitReason)); err != nil {
		return err
	}

	if err := d.Set("workflow_ctx", (s.WorkflowCtx)); err != nil {
		return err
	}

	if err := d.Set("workflow_definition", flattenMapWorkflowWorkflowDefinitionRef(s.WorkflowDefinition, d)); err != nil {
		return err
	}

	if err := d.Set("workflow_meta_type", (s.WorkflowMetaType)); err != nil {
		return err
	}

	if err := d.Set("workflow_task_count", (s.WorkflowTaskCount)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceWorkflowWorkflowInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.WorkflowWorkflowInfo
	if d.HasChange("action") {
		v := d.Get("action")
		x := (v.(string))
		o.Action = &x
	}

	if d.HasChange("cleanup_time") {
		v := d.Get("cleanup_time")
		x, _ := strfmt.ParseDateTime(v.(string))
		o.CleanupTime = x
	}

	if d.HasChange("end_time") {
		v := d.Get("end_time")
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndTime = x
	}

	if d.HasChange("failed_workflow_cleanup_duration") {
		v := d.Get("failed_workflow_cleanup_duration")
		x := int64(v.(int))
		o.FailedWorkflowCleanupDuration = x
	}

	if d.HasChange("input") {
		v := d.Get("input")
		x := v
		o.Input = &x
	}

	if d.HasChange("inst_id") {
		v := d.Get("inst_id")
		x := (v.(string))
		o.InstID = x
	}

	if d.HasChange("internal") {
		v := d.Get("internal")
		x := (v.(bool))
		o.Internal = &x
	}

	if d.HasChange("meta_version") {
		v := d.Get("meta_version")
		x := int64(v.(int))
		o.MetaVersion = x
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.Moid = x
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.Name = x
	}

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
	}

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := models.IamAccountRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamAccountRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Organization = &x
	}

	if d.HasChange("output") {
		v := d.Get("output")
		x := v
		o.Output = &x
	}

	if d.HasChange("parent_task_info") {
		v := d.Get("parent_task_info")
		p := models.WorkflowTaskInfoRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowTaskInfoRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.ParentTaskInfo = &x
	}

	if d.HasChange("pending_dynamic_workflow_info") {
		v := d.Get("pending_dynamic_workflow_info")
		p := models.WorkflowPendingDynamicWorkflowInfoRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowPendingDynamicWorkflowInfoRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.PendingDynamicWorkflowInfo = &x
	}

	if d.HasChange("progress") {
		v := d.Get("progress")
		x := v.(float32)
		o.Progress = x
	}

	if d.HasChange("src") {
		v := d.Get("src")
		x := (v.(string))
		o.Src = x
	}

	if d.HasChange("start_time") {
		v := d.Get("start_time")
		x, _ := strfmt.ParseDateTime(v.(string))
		o.StartTime = x
	}

	if d.HasChange("status") {
		v := d.Get("status")
		x := (v.(string))
		o.Status = x
	}

	if d.HasChange("success_workflow_cleanup_duration") {
		v := d.Get("success_workflow_cleanup_duration")
		x := int64(v.(int))
		o.SuccessWorkflowCleanupDuration = x
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x
	}

	if d.HasChange("task_infos") {
		v := d.Get("task_infos")
		x := make([]*models.WorkflowTaskInfoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowTaskInfoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.TaskInfos = x
	}

	if d.HasChange("type") {
		v := d.Get("type")
		x := (v.(string))
		o.Type = x
	}

	if d.HasChange("user_id") {
		v := d.Get("user_id")
		x := (v.(string))
		o.UserID = x
	}

	if d.HasChange("wait_reason") {
		v := d.Get("wait_reason")
		x := (v.(string))
		o.WaitReason = &x
	}

	if d.HasChange("workflow_ctx") {
		v := d.Get("workflow_ctx")
		x := v
		o.WorkflowCtx = &x
	}

	if d.HasChange("workflow_definition") {
		v := d.Get("workflow_definition")
		p := models.WorkflowWorkflowDefinitionRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowWorkflowDefinitionRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.WorkflowDefinition = &x
	}

	if d.HasChange("workflow_meta_type") {
		v := d.Get("workflow_meta_type")
		x := (v.(string))
		o.WorkflowMetaType = &x
	}

	if d.HasChange("workflow_task_count") {
		v := d.Get("workflow_task_count")
		x := int64(v.(int))
		o.WorkflowTaskCount = x
	}

	url := "workflow/WorkflowInfos" + "/" + d.Id()
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendUpdateRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceWorkflowWorkflowInfoRead(d, meta)
}

func resourceWorkflowWorkflowInfoDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "workflow/WorkflowInfos" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
