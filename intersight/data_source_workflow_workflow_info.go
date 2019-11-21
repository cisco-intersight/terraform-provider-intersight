package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceWorkflowWorkflowInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWorkflowWorkflowInfoRead,
		Schema: map[string]*schema.Schema{
			"action": {
				Description: "The action of the workflow such as start, cancel, retry, pause.",
				Type:        schema.TypeString,
				Optional:    true,
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
			},
			"meta_version": {
				Description: "Version of the workflow metadata for which this workflow execution was started.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "A name of the workflow execution instance.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Computed: true,
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
				Computed: true,
			},
			"workflow_meta_type": {
				Description: "The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance.",
				Type:        schema.TypeString,
				Optional:    true,
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
func dataSourceWorkflowWorkflowInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/WorkflowInfos"
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
	if v, ok := d.GetOk("inst_id"); ok {
		x := (v.(string))
		o.InstID = x
	}
	if v, ok := d.GetOk("internal"); ok {
		x := (v.(bool))
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
	if v, ok := d.GetOk("workflow_meta_type"); ok {
		x := (v.(string))
		o.WorkflowMetaType = &x
	}
	if v, ok := d.GetOk("workflow_task_count"); ok {
		x := int64(v.(int))
		o.WorkflowTaskCount = x
	}

	data, err := o.MarshalJSON()
	body, err := conn.SendGetRequest(url, data)
	if err != nil {
		return err
	}
	var x = make(map[string]interface{})
	if err = json.Unmarshal(body, &x); err != nil {
		return err
	}
	result := x["Results"]
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s models.WorkflowWorkflowInfo
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
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

			if err := d.Set("workflow_definition", flattenMapWorkflowWorkflowDefinitionRef(s.WorkflowDefinition, d)); err != nil {
				return err
			}
			if err := d.Set("workflow_meta_type", (s.WorkflowMetaType)); err != nil {
				return err
			}
			if err := d.Set("workflow_task_count", (s.WorkflowTaskCount)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
