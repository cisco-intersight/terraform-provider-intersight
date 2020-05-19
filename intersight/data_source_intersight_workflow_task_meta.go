package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "A description of the task.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"input_keys": {
				Type:     schema.TypeList,
				Optional: true,
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"output_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
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
	var o = models.NewWorkflowTaskMeta()
	if v, ok := d.GetOk("action_type"); ok {
		x := (v.(string))
		o.SetActionType(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("internal"); ok {
		x := (v.(bool))
		o.SetInternal(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("response_timeout_sec"); ok {
		x := int64(v.(int))
		o.SetResponseTimeoutSec(x)
	}
	if v, ok := d.GetOk("retry_count"); ok {
		x := int64(v.(int))
		o.SetRetryCount(x)
	}
	if v, ok := d.GetOk("retry_delay_sec"); ok {
		x := int64(v.(int))
		o.SetRetryDelaySec(x)
	}
	if v, ok := d.GetOk("retry_logic"); ok {
		x := (v.(string))
		o.SetRetryLogic(x)
	}
	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.SetSrc(x)
	}
	if v, ok := d.GetOk("timeout_policy"); ok {
		x := (v.(string))
		o.SetTimeoutPolicy(x)
	}
	if v, ok := d.GetOk("timeout_sec"); ok {
		x := int64(v.(int))
		o.SetTimeoutSec(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.WorkflowApi.GetWorkflowTaskMetaList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewWorkflowTaskMeta()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("action_type", (s.ActionType)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
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

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
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
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
