package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCondAlarm() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCondAlarmRead,
		Schema: map[string]*schema.Schema{
			"affected_mo_id": {
				Description: "MoId of the affected object from the managed system's point of view.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"affected_mo_type": {
				Description: "Managed system affected object type. For example Adaptor, FI, CIMC.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"affected_object": {
				Description: "A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ancestor_mo_id": {
				Description: "Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ancestor_mo_type": {
				Description: "Parent MO type of the fault from managed system. For example, Blade for adaptor fault.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"code": {
				Description: "A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"creation_time": {
				Description: "The time the alarm was created.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "A longer description of the alarm than the name. The description contains details of the component reporting the issue.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_transition_time": {
				Description: "The time the alarm last had a change in severity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ms_affected_object": {
				Description: "A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"orig_severity": {
				Description: "The original severity when the alarm was first created.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"registered_device": {
				Description: "Relationship to set accountMoid on Alarms. With \"onpeerdelete\" set to \"cascade\", Alarms get deleted when the associated registered device is deleted.",
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
			"severity": {
				Description: "The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared.",
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
		},
	}
}
func dataSourceCondAlarmRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "cond/Alarms"
	var o models.CondAlarm
	if v, ok := d.GetOk("affected_mo_id"); ok {
		x := (v.(string))
		o.AffectedMoID = x
	}
	if v, ok := d.GetOk("affected_mo_type"); ok {
		x := (v.(string))
		o.AffectedMoType = x
	}
	if v, ok := d.GetOk("affected_object"); ok {
		x := (v.(string))
		o.AffectedObject = x
	}
	if v, ok := d.GetOk("ancestor_mo_id"); ok {
		x := (v.(string))
		o.AncestorMoID = x
	}
	if v, ok := d.GetOk("ancestor_mo_type"); ok {
		x := (v.(string))
		o.AncestorMoType = x
	}
	if v, ok := d.GetOk("code"); ok {
		x := (v.(string))
		o.Code = x
	}
	if v, ok := d.GetOk("creation_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.CreationTime = x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("last_transition_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.LastTransitionTime = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("ms_affected_object"); ok {
		x := (v.(string))
		o.MsAffectedObject = x
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("orig_severity"); ok {
		x := (v.(string))
		o.OrigSeverity = &x
	}
	if v, ok := d.GetOk("severity"); ok {
		x := (v.(string))
		o.Severity = &x
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
			var s models.CondAlarm
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("affected_mo_id", (s.AffectedMoID)); err != nil {
				return err
			}
			if err := d.Set("affected_mo_type", (s.AffectedMoType)); err != nil {
				return err
			}
			if err := d.Set("affected_object", (s.AffectedObject)); err != nil {
				return err
			}
			if err := d.Set("ancestor_mo_id", (s.AncestorMoID)); err != nil {
				return err
			}
			if err := d.Set("ancestor_mo_type", (s.AncestorMoType)); err != nil {
				return err
			}
			if err := d.Set("code", (s.Code)); err != nil {
				return err
			}

			if err := d.Set("creation_time", (s.CreationTime).String()); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}

			if err := d.Set("last_transition_time", (s.LastTransitionTime).String()); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("ms_affected_object", (s.MsAffectedObject)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("orig_severity", (s.OrigSeverity)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("severity", (s.Severity)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
