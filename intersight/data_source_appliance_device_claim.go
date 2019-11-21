package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceApplianceDeviceClaim() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceDeviceClaimRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "DeviceClaim managed object to Account relationship.",
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
			"device_id": {
				Description: "Device identifier of the endpoint device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hostname": {
				Description: "Hostname or IP address of the endpoint device the user wants to claim.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"message": {
				Description: "Message set by the device claim process.",
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
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Description: "Password to be used to login to the endpoint device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"platform_type": {
				Description: "Platform type of the endpoint device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"request_id": {
				Description: "User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"security_token": {
				Description: "Device security token of the endpoint device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the device claim process.",
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
			"username": {
				Description: "Username to log in to the endpoint device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceApplianceDeviceClaimRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "appliance/DeviceClaims"
	var o models.ApplianceDeviceClaim
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.DeviceID = x
	}
	if v, ok := d.GetOk("hostname"); ok {
		x := (v.(string))
		o.Hostname = x
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.IsPasswordSet = &x
	}
	if v, ok := d.GetOk("message"); ok {
		x := (v.(string))
		o.Message = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.Password = x
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.PlatformType = x
	}
	if v, ok := d.GetOk("request_id"); ok {
		x := (v.(string))
		o.RequestID = x
	}
	if v, ok := d.GetOk("security_token"); ok {
		x := (v.(string))
		o.SecurityToken = x
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = x
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.Username = x
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
			var s models.ApplianceDeviceClaim
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
				return err
			}
			if err := d.Set("device_id", (s.DeviceID)); err != nil {
				return err
			}
			if err := d.Set("hostname", (s.Hostname)); err != nil {
				return err
			}
			if err := d.Set("is_password_set", (s.IsPasswordSet)); err != nil {
				return err
			}
			if err := d.Set("message", (s.Message)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("password", (s.Password)); err != nil {
				return err
			}
			if err := d.Set("platform_type", (s.PlatformType)); err != nil {
				return err
			}
			if err := d.Set("request_id", (s.RequestID)); err != nil {
				return err
			}
			if err := d.Set("security_token", (s.SecurityToken)); err != nil {
				return err
			}
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("username", (s.Username)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
