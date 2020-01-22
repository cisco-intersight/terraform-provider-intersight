package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceWorkflowTaskInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWorkflowTaskInfoRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Description: "The task description and this is the description that was added when the task was included into the workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "The time stamp when the task reached a final state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"failure_reason": {
				Description: "Description of the reason why the task failed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inst_id": {
				Description: "The current running task instance Id.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"internal": {
				Description: "Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"label": {
				Description: "User friendly short label to describe this task instance in the workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"message": {
				Description: "Collection of Task execution messages with severity.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"message": {
							Description: "An i18n message that can be translated in multiple languages to support internationalization.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"severity": {
							Description: "The severity of the Task or Workflow message warning/error/info etc.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Task definition name which specifies the task type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.",
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
							Description: "The Object Type of the referenced REST resource.",
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
			"ref_name": {
				Description: "The task reference name to ensure its unique inside a workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"retry_count": {
				Description: "A counter for number of retries.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"start_time": {
				Description: "The time stamp when the task started execution.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "The status of the task and this will specify if the task is running or has reached a final state.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sub_workflow_info": {
				Description: "A collection of references to the [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) Managed Object.When this managed object is deleted, the referenced [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) MO unsets its reference to this deleted MO.",
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
							Description: "The Object Type of the referenced REST resource.",
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
			"task_definition": {
				Description: "The task definition that was used to launch this task execution instance.",
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
							Description: "The Object Type of the referenced REST resource.",
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
			"task_inst_id_list": {
				Description: "The list keeps a list of retried task instances.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Description: "Status of the retried task.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"task_inst_id": {
							Description: "Retry instance will get a unique instance id.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"workflow_info": {
				Description: "A collection of references to the [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) Managed Object.When this managed object is deleted, the referenced [workflow.WorkflowInfo](mo://workflow.WorkflowInfo) MO unsets its reference to this deleted MO.",
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
							Description: "The Object Type of the referenced REST resource.",
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
		},
	}
}
func dataSourceWorkflowTaskInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/TaskInfos"
	var o models.WorkflowTaskInfo
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndTime = x
	}
	if v, ok := d.GetOk("failure_reason"); ok {
		x := (v.(string))
		o.FailureReason = x
	}
	if v, ok := d.GetOk("inst_id"); ok {
		x := (v.(string))
		o.InstID = x
	}
	if v, ok := d.GetOk("internal"); ok {
		x := (v.(bool))
		o.Internal = &x
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.Label = x
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
	if v, ok := d.GetOk("ref_name"); ok {
		x := (v.(string))
		o.RefName = x
	}
	if v, ok := d.GetOk("retry_count"); ok {
		x := int64(v.(int))
		o.RetryCount = x
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.StartTime = x
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = x
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
			var s models.WorkflowTaskInfo
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}

			if err := d.Set("end_time", (s.EndTime).String()); err != nil {
				return err
			}
			if err := d.Set("failure_reason", (s.FailureReason)); err != nil {
				return err
			}
			if err := d.Set("inst_id", (s.InstID)); err != nil {
				return err
			}
			if err := d.Set("internal", (s.Internal)); err != nil {
				return err
			}
			if err := d.Set("label", (s.Label)); err != nil {
				return err
			}

			if err := d.Set("message", flattenListWorkflowMessage(s.Message, d)); err != nil {
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

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("ref_name", (s.RefName)); err != nil {
				return err
			}
			if err := d.Set("retry_count", (s.RetryCount)); err != nil {
				return err
			}

			if err := d.Set("start_time", (s.StartTime).String()); err != nil {
				return err
			}
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("sub_workflow_info", flattenMapWorkflowWorkflowInfoRef(s.SubWorkflowInfo, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}

			if err := d.Set("task_definition", flattenMapWorkflowTaskDefinitionRef(s.TaskDefinition, d)); err != nil {
				return err
			}

			if err := d.Set("task_inst_id_list", flattenListWorkflowTaskRetryInfo(s.TaskInstIDList, d)); err != nil {
				return err
			}

			if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRef(s.WorkflowInfo, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
