package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceStorageFlexFlashControllerProps() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageFlexFlashControllerPropsRead,
		Schema: map[string]*schema.Schema{
			"cards_manageable": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"configured_mode": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"controller_name": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"controller_status": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_mo_id": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fw_version": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"internal_state": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "This field identifies the model of the given component.",
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
			"operating_mode": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"physical_drive_count": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"product_name": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"registered_device": {
				Description: "The Device to which this Managed Object is associated.",
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
			"revision": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rn": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"startup_fw_version": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"storage_flex_flash_controller": {
				Description: "A collection of references to the [storage.FlexFlashController](mo://storage.FlexFlashController) Managed Object.When this managed object is deleted, the referenced [storage.FlexFlashController](mo://storage.FlexFlashController) MO unsets its reference to this deleted MO.",
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
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"virtual_drive_count": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceStorageFlexFlashControllerPropsRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "storage/FlexFlashControllerProps"
	var o models.StorageFlexFlashControllerProps
	if v, ok := d.GetOk("cards_manageable"); ok {
		x := (v.(string))
		o.CardsManageable = x
	}
	if v, ok := d.GetOk("configured_mode"); ok {
		x := (v.(string))
		o.ConfiguredMode = x
	}
	if v, ok := d.GetOk("controller_name"); ok {
		x := (v.(string))
		o.ControllerName = x
	}
	if v, ok := d.GetOk("controller_status"); ok {
		x := (v.(string))
		o.ControllerStatus = x
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.DeviceMoID = x
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.Dn = x
	}
	if v, ok := d.GetOk("fw_version"); ok {
		x := (v.(string))
		o.FwVersion = x
	}
	if v, ok := d.GetOk("internal_state"); ok {
		x := (v.(string))
		o.InternalState = x
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.Model = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("operating_mode"); ok {
		x := (v.(string))
		o.OperatingMode = x
	}
	if v, ok := d.GetOk("physical_drive_count"); ok {
		x := (v.(string))
		o.PhysicalDriveCount = x
	}
	if v, ok := d.GetOk("product_name"); ok {
		x := (v.(string))
		o.ProductName = x
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.Revision = x
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.Rn = x
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.Serial = x
	}
	if v, ok := d.GetOk("startup_fw_version"); ok {
		x := (v.(string))
		o.StartupFwVersion = x
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x
	}
	if v, ok := d.GetOk("virtual_drive_count"); ok {
		x := (v.(string))
		o.VirtualDriveCount = x
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
			var s models.StorageFlexFlashControllerProps
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("cards_manageable", (s.CardsManageable)); err != nil {
				return err
			}
			if err := d.Set("configured_mode", (s.ConfiguredMode)); err != nil {
				return err
			}
			if err := d.Set("controller_name", (s.ControllerName)); err != nil {
				return err
			}
			if err := d.Set("controller_status", (s.ControllerStatus)); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoID)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}
			if err := d.Set("fw_version", (s.FwVersion)); err != nil {
				return err
			}
			if err := d.Set("internal_state", (s.InternalState)); err != nil {
				return err
			}
			if err := d.Set("model", (s.Model)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("operating_mode", (s.OperatingMode)); err != nil {
				return err
			}
			if err := d.Set("physical_drive_count", (s.PhysicalDriveCount)); err != nil {
				return err
			}
			if err := d.Set("product_name", (s.ProductName)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("revision", (s.Revision)); err != nil {
				return err
			}
			if err := d.Set("rn", (s.Rn)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}
			if err := d.Set("startup_fw_version", (s.StartupFwVersion)); err != nil {
				return err
			}

			if err := d.Set("storage_flex_flash_controller", flattenMapStorageFlexFlashControllerRef(s.StorageFlexFlashController, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			if err := d.Set("virtual_drive_count", (s.VirtualDriveCount)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
