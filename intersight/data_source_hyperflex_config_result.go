package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHyperflexConfigResult() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHyperflexConfigResultRead,
		Schema: map[string]*schema.Schema{
			"cluster_profile": {
				Description: "A collection of references to the [hyperflex.ClusterProfile](mo://hyperflex.ClusterProfile) Managed Object.When this managed object is deleted, the referenced [hyperflex.ClusterProfile](mo://hyperflex.ClusterProfile) MO unsets its reference to this deleted MO.",
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
			"config_progress": {
				Description: "The progress percentage of the running configuration or workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"config_stage": {
				Description: "The current running stage of the configuration or workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"config_state": {
				Description: "Indicates overall configuration state for applying the configuration to the end point. Values  -- ok, ok-with-warning, errored.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"duration": {
				Description: "The duration of the running configuration or workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
			"result_entries": {
				Description: "Detailed result entries for both validation & configration. Each result entry can be error/warning/info messages and the context.",
				Type:        schema.TypeList,
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
			"start_time": {
				Description: "The start time of the configuration or workflow.",
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
			"validation_state": {
				Description: "Indicates overall state for logical model validation. Values  -- ok, ok-with-warning, errored.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceHyperflexConfigResultRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "hyperflex/ConfigResults"
	var o models.HyperflexConfigResult
	if v, ok := d.GetOk("config_progress"); ok {
		x := (v.(string))
		o.ConfigProgress = x
	}
	if v, ok := d.GetOk("config_stage"); ok {
		x := (v.(string))
		o.ConfigStage = x
	}
	if v, ok := d.GetOk("config_state"); ok {
		x := (v.(string))
		o.ConfigState = x
	}
	if v, ok := d.GetOk("duration"); ok {
		x := (v.(string))
		o.Duration = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("start_time"); ok {
		x := (v.(string))
		o.StartTime = x
	}
	if v, ok := d.GetOk("validation_state"); ok {
		x := (v.(string))
		o.ValidationState = x
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
			var s models.HyperflexConfigResult
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("cluster_profile", flattenMapHyperflexClusterProfileRef(s.ClusterProfile, d)); err != nil {
				return err
			}
			if err := d.Set("config_progress", (s.ConfigProgress)); err != nil {
				return err
			}
			if err := d.Set("config_stage", (s.ConfigStage)); err != nil {
				return err
			}
			if err := d.Set("config_state", (s.ConfigState)); err != nil {
				return err
			}
			if err := d.Set("duration", (s.Duration)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("result_entries", flattenListHyperflexConfigResultEntryRef(s.ResultEntries, d)); err != nil {
				return err
			}
			if err := d.Set("start_time", (s.StartTime)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("validation_state", (s.ValidationState)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
