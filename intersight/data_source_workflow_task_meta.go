package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceWorkflowTaskMeta() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWorkflowTaskMetaRead,
		Schema: map[string]*schema.Schema{
			"action_type": {
				Description: "A task execution type to indicate if it is a system task.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "A description of the task.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"input_keys": {
				Description: "Input keys for the task which specifies parameters the task can take in as inputs.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"internal": {
				Description: "Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "A task name that should be unique in Conductor DB.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"output_keys": {
				Description: "Output keys for the task which specifies parameters the task will output at the end of execution.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"response_timeout_sec": {
				Description: "The worker respnose timeout value.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"retry_count": {
				Description: "A number of reties for this task.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"retry_delay_sec": {
				Description: "The time on which the retry will be delayed.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"retry_logic": {
				Description: "A logic which defines the way to handle retry (FIXED, EXPONENTIAL_BACKOFF).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"src": {
				Description: "A service owns the task metadata.",
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
			"timeout_policy": {
				Description: "A policy which defines the way to handle timeout (RETRY, TIME_OUT_WF, ALERT_ONLY).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"timeout_sec": {
				Description: "A timeout value for the task in seconds.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}
func dataSourceWorkflowTaskMetaRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/TaskMeta"
	var o models.WorkflowTaskMeta
	if v, ok := d.GetOk("action_type"); ok {
		x := (v.(string))
		o.ActionType = x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("internal"); ok {
		x := (v.(bool))
		o.Internal = &x
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
	if v, ok := d.GetOk("response_timeout_sec"); ok {
		x := int64(v.(int))
		o.ResponseTimeoutSec = x
	}
	if v, ok := d.GetOk("retry_count"); ok {
		x := int64(v.(int))
		o.RetryCount = x
	}
	if v, ok := d.GetOk("retry_delay_sec"); ok {
		x := int64(v.(int))
		o.RetryDelaySec = x
	}
	if v, ok := d.GetOk("retry_logic"); ok {
		x := (v.(string))
		o.RetryLogic = x
	}
	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.Src = x
	}
	if v, ok := d.GetOk("timeout_policy"); ok {
		x := (v.(string))
		o.TimeoutPolicy = x
	}
	if v, ok := d.GetOk("timeout_sec"); ok {
		x := int64(v.(int))
		o.TimeoutSec = x
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
			var s models.WorkflowTaskMeta
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("action_type", (s.ActionType)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("input_keys", (s.InputKeys)); err != nil {
				return err
			}
			if err := d.Set("internal", (s.Internal)); err != nil {
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
			if err := d.Set("output_keys", (s.OutputKeys)); err != nil {
				return err
			}
			if err := d.Set("response_timeout_sec", (s.ResponseTimeoutSec)); err != nil {
				return err
			}
			if err := d.Set("retry_count", (s.RetryCount)); err != nil {
				return err
			}
			if err := d.Set("retry_delay_sec", (s.RetryDelaySec)); err != nil {
				return err
			}
			if err := d.Set("retry_logic", (s.RetryLogic)); err != nil {
				return err
			}
			if err := d.Set("src", (s.Src)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("timeout_policy", (s.TimeoutPolicy)); err != nil {
				return err
			}
			if err := d.Set("timeout_sec", (s.TimeoutSec)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
