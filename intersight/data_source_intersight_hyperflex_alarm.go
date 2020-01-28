package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHyperflexAlarm() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHyperflexAlarmRead,
		Schema: map[string]*schema.Schema{
			"acknowledged": {
				Description: "",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"acknowledged_by": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"acknowledged_time": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"acknowledged_time_as_utc": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster": {
				Description: "A collection of references to the [hyperflex.Cluster](mo://hyperflex.Cluster) Managed Object.When this managed object is deleted, the referenced [hyperflex.Cluster](mo://hyperflex.Cluster) MO unsets its reference to this deleted MO.",
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
			"description": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_data": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_name": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_type": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_uu_id": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"message": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "",
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
			"status": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"triggered_time": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"triggered_time_as_utc": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceHyperflexAlarmRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "hyperflex/Alarms"
	var o models.HyperflexAlarm
	if v, ok := d.GetOk("acknowledged"); ok {
		x := (v.(bool))
		o.Acknowledged = &x
	}
	if v, ok := d.GetOk("acknowledged_by"); ok {
		x := (v.(string))
		o.AcknowledgedBy = x
	}
	if v, ok := d.GetOk("acknowledged_time"); ok {
		x := int64(v.(int))
		o.AcknowledgedTime = x
	}
	if v, ok := d.GetOk("acknowledged_time_as_utc"); ok {
		x := (v.(string))
		o.AcknowledgedTimeAsUtc = x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("entity_data"); ok {
		x := (v.(string))
		o.EntityData = x
	}
	if v, ok := d.GetOk("entity_name"); ok {
		x := (v.(string))
		o.EntityName = x
	}
	if v, ok := d.GetOk("entity_type"); ok {
		x := (v.(string))
		o.EntityType = x
	}
	if v, ok := d.GetOk("entity_uu_id"); ok {
		x := (v.(string))
		o.EntityUuID = x
	}
	if v, ok := d.GetOk("message"); ok {
		x := (v.(string))
		o.Message = x
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
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = x
	}
	if v, ok := d.GetOk("triggered_time"); ok {
		x := int64(v.(int))
		o.TriggeredTime = x
	}
	if v, ok := d.GetOk("triggered_time_as_utc"); ok {
		x := (v.(string))
		o.TriggeredTimeAsUtc = x
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.UUID = x
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
			var s models.HyperflexAlarm
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("acknowledged", (s.Acknowledged)); err != nil {
				return err
			}
			if err := d.Set("acknowledged_by", (s.AcknowledgedBy)); err != nil {
				return err
			}
			if err := d.Set("acknowledged_time", (s.AcknowledgedTime)); err != nil {
				return err
			}
			if err := d.Set("acknowledged_time_as_utc", (s.AcknowledgedTimeAsUtc)); err != nil {
				return err
			}

			if err := d.Set("cluster", flattenMapHyperflexClusterRef(s.Cluster, d)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("entity_data", (s.EntityData)); err != nil {
				return err
			}
			if err := d.Set("entity_name", (s.EntityName)); err != nil {
				return err
			}
			if err := d.Set("entity_type", (s.EntityType)); err != nil {
				return err
			}
			if err := d.Set("entity_uu_id", (s.EntityUuID)); err != nil {
				return err
			}
			if err := d.Set("message", (s.Message)); err != nil {
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
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("triggered_time", (s.TriggeredTime)); err != nil {
				return err
			}
			if err := d.Set("triggered_time_as_utc", (s.TriggeredTimeAsUtc)); err != nil {
				return err
			}
			if err := d.Set("uuid", (s.UUID)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
