package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceApplianceSystemInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceSystemInfoRead,
		Schema: map[string]*schema.Schema{
			"cloud_conn_status": {
				Description: "Connection state of the Intersight Appliance to the Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hostname": {
				Description: "Publicly accessible FQDN or IP address of the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"init_done": {
				Description: "Indicates that the setup initialization process has been completed.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
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
			"operational_status": {
				Description: "Operational status of the Intersight Appliance cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial_id": {
				Description: "SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance.",
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
			"time_zone": {
				Description: "Timezone of the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version": {
				Description: "Current software version of the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceApplianceSystemInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "appliance/SystemInfos"
	var o models.ApplianceSystemInfo
	if v, ok := d.GetOk("cloud_conn_status"); ok {
		x := (v.(string))
		o.CloudConnStatus = x
	}
	if v, ok := d.GetOk("hostname"); ok {
		x := (v.(string))
		o.Hostname = x
	}
	if v, ok := d.GetOk("init_done"); ok {
		x := (v.(bool))
		o.InitDone = &x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("operational_status"); ok {
		x := (v.(string))
		o.OperationalStatus = x
	}
	if v, ok := d.GetOk("serial_id"); ok {
		x := (v.(string))
		o.SerialID = x
	}
	if v, ok := d.GetOk("time_zone"); ok {
		x := (v.(string))
		o.TimeZone = &x
	}
	if v, ok := d.GetOk("version"); ok {
		x := (v.(string))
		o.Version = x
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
			var s models.ApplianceSystemInfo
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("cloud_conn_status", (s.CloudConnStatus)); err != nil {
				return err
			}
			if err := d.Set("hostname", (s.Hostname)); err != nil {
				return err
			}
			if err := d.Set("init_done", (s.InitDone)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("operational_status", (s.OperationalStatus)); err != nil {
				return err
			}
			if err := d.Set("serial_id", (s.SerialID)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("time_zone", (s.TimeZone)); err != nil {
				return err
			}
			if err := d.Set("version", (s.Version)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
