package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiatelemetryNiaInventory() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiatelemetryNiaInventoryRead,
		Schema: map[string]*schema.Schema{
			"cpu": {
				Description: "CPU usage of device being inventoried.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"crash_reset_logs": {
				Description: "Last crash reset reason of device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_name": {
				Description: "Name of device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_type": {
				Description: "Type of device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disk": {
				Description: "Disk Usage of device being inventoried.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"free": {
							Description: "The free disk capacity, currently the type of this field is set to integer.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"name": {
							Description: "Disk Name used to identified the disk usage record.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"total": {
							Description: "The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"used": {
							Description: "The used disk capacity, currently the type of this field is set to integer.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"license_state": {
				Description: "The license of this device.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"log_in_time": {
				Description: "Last log in time device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"log_out_time": {
				Description: "Last log out time of device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"memory": {
				Description: "Memory usage of device being inventoried.",
				Type:        schema.TypeInt,
				Optional:    true,
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
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"record_type": {
				Description: "Type of record DCNM / APIC / SE.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"record_version": {
				Description: "Version of record being pushed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"registered_device": {
				Description: "Relationship to the Device Registration object for this setup.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"serial": {
				Description: "Serial number of device being invetoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_download": {
				Description: "Last software downloaded of device being inventoried.",
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"version": {
				Description: "Software version of device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceNiatelemetryNiaInventoryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "niatelemetry/NiaInventories"
	var o models.NiatelemetryNiaInventory
	if v, ok := d.GetOk("cpu"); ok {
		x := v.(float32)
		o.CPU = x
	}
	if v, ok := d.GetOk("crash_reset_logs"); ok {
		x := (v.(string))
		o.CrashResetLogs = x
	}
	if v, ok := d.GetOk("device_name"); ok {
		x := (v.(string))
		o.DeviceName = x
	}
	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.DeviceType = x
	}
	if v, ok := d.GetOk("log_in_time"); ok {
		x := (v.(string))
		o.LogInTime = x
	}
	if v, ok := d.GetOk("log_out_time"); ok {
		x := (v.(string))
		o.LogOutTime = x
	}
	if v, ok := d.GetOk("memory"); ok {
		x := int64(v.(int))
		o.Memory = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.RecordType = x
	}
	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.RecordVersion = x
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.Serial = x
	}
	if v, ok := d.GetOk("software_download"); ok {
		x := (v.(string))
		o.SoftwareDownload = x
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
			var s models.NiatelemetryNiaInventory
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("cpu", (s.CPU)); err != nil {
				return err
			}
			if err := d.Set("crash_reset_logs", (s.CrashResetLogs)); err != nil {
				return err
			}
			if err := d.Set("device_name", (s.DeviceName)); err != nil {
				return err
			}
			if err := d.Set("device_type", (s.DeviceType)); err != nil {
				return err
			}

			if err := d.Set("disk", flattenMapNiatelemetryDiskinfo(s.Disk, d)); err != nil {
				return err
			}

			if err := d.Set("license_state", flattenMapNiatelemetryNiaLicenseStateRef(s.LicenseState, d)); err != nil {
				return err
			}
			if err := d.Set("log_in_time", (s.LogInTime)); err != nil {
				return err
			}
			if err := d.Set("log_out_time", (s.LogOutTime)); err != nil {
				return err
			}
			if err := d.Set("memory", (s.Memory)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("record_type", (s.RecordType)); err != nil {
				return err
			}
			if err := d.Set("record_version", (s.RecordVersion)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}
			if err := d.Set("software_download", (s.SoftwareDownload)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
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
