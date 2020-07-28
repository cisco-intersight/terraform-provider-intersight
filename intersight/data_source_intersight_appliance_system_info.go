package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceApplianceSystemInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceSystemInfoRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cloud_conn_status": {
				Description: "Connection state of the Intersight Appliance to the Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"deployment_size": {
				Description: "Current running deployment size for the Intersight Appliance cluster. Eg. small, medium, large etc.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
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
				Type:     schema.TypeList,
				Optional: true,
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
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"time_zone": {
				Description: "Timezone of the Intersight Appliance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
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
	var o = models.NewApplianceSystemInfoWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cloud_conn_status"); ok {
		x := (v.(string))
		o.SetCloudConnStatus(x)
	}
	if v, ok := d.GetOk("deployment_size"); ok {
		x := (v.(string))
		o.SetDeploymentSize(x)
	}
	if v, ok := d.GetOk("hostname"); ok {
		x := (v.(string))
		o.SetHostname(x)
	}
	if v, ok := d.GetOk("init_done"); ok {
		x := (v.(bool))
		o.SetInitDone(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("operational_status"); ok {
		x := (v.(string))
		o.SetOperationalStatus(x)
	}
	if v, ok := d.GetOk("serial_id"); ok {
		x := (v.(string))
		o.SetSerialId(x)
	}
	if v, ok := d.GetOk("time_zone"); ok {
		x := (v.(string))
		o.SetTimeZone(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.ApplianceApi.GetApplianceSystemInfoList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.ApplianceSystemInfoList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to ApplianceSystemInfo: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewApplianceSystemInfoWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("cloud_conn_status", (s.GetCloudConnStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property CloudConnStatus: %+v", err)
			}
			if err := d.Set("deployment_size", (s.GetDeploymentSize())); err != nil {
				return fmt.Errorf("error occurred while setting property DeploymentSize: %+v", err)
			}
			if err := d.Set("hostname", (s.GetHostname())); err != nil {
				return fmt.Errorf("error occurred while setting property Hostname: %+v", err)
			}
			if err := d.Set("init_done", (s.GetInitDone())); err != nil {
				return fmt.Errorf("error occurred while setting property InitDone: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("operational_status", (s.GetOperationalStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property OperationalStatus: %+v", err)
			}
			if err := d.Set("serial_id", (s.GetSerialId())); err != nil {
				return fmt.Errorf("error occurred while setting property SerialId: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("time_zone", (s.GetTimeZone())); err != nil {
				return fmt.Errorf("error occurred while setting property TimeZone: %+v", err)
			}
			if err := d.Set("nr_version", (s.GetVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property Version: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
