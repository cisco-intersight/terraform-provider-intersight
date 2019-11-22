package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceWorkflowPendingDynamicWorkflowInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWorkflowPendingDynamicWorkflowInfoRead,
		Schema: map[string]*schema.Schema{
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "A name for the pending dynamic workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pending_services": {
				Description: "The pending services the dynamic workflow is waiting for to return the task list.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"src": {
				Description: "The src is workflow owner service.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "The current status of the PendingDynamicWorkflowInfo.",
				Type:        schema.TypeString,
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
			"wait_on_duplicate": {
				Description: "When set to true workflow engine will wait for a duplicate to finish before starting a new one.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"workflow_action_task_lists": {
				Description: "The task lists returned by services for building dynamic workflows.  There will be an entry for every different workflow action.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Description: "The action of the Dynamic Workflow.",
							Type:        schema.TypeString,
							Optional:    true,
						},
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
					},
				},
				Computed: true,
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
			"workflow_key": {
				Description: "This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceWorkflowPendingDynamicWorkflowInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/PendingDynamicWorkflowInfos"
	var o models.WorkflowPendingDynamicWorkflowInfo
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
	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.Src = x
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = &x
	}
	if v, ok := d.GetOk("wait_on_duplicate"); ok {
		x := (v.(bool))
		o.WaitOnDuplicate = &x
	}
	if v, ok := d.GetOk("workflow_key"); ok {
		x := (v.(string))
		o.WorkflowKey = x
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
			var s models.WorkflowPendingDynamicWorkflowInfo
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
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
			if err := d.Set("pending_services", (s.PendingServices)); err != nil {
				return err
			}
			if err := d.Set("src", (s.Src)); err != nil {
				return err
			}
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("wait_on_duplicate", (s.WaitOnDuplicate)); err != nil {
				return err
			}

			if err := d.Set("workflow_action_task_lists", flattenListWorkflowDynamicWorkflowActionTaskList(s.WorkflowActionTaskLists, d)); err != nil {
				return err
			}

			if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRef(s.WorkflowInfo, d)); err != nil {
				return err
			}
			if err := d.Set("workflow_key", (s.WorkflowKey)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
