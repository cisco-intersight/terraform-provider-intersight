package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceStorageFlexUtilPhysicalDrive() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageFlexUtilPhysicalDriveRead,
		Schema: map[string]*schema.Schema{
			"block_size": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"capacity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_mo_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drives_enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"health": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"manufacturer_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oem_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"partition_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pd_status": {
				Type:     schema.TypeString,
				Optional: true,
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
			"physical_drive": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_revision": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"read_error_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"read_error_threshold": {
				Type:     schema.TypeString,
				Optional: true,
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
			"revision": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
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
			"storage_flex_util_controller": {
				Description: "A collection of references to the [storage.FlexUtilController](mo://storage.FlexUtilController) Managed Object.\nWhen this managed object is deleted, the referenced [storage.FlexUtilController](mo://storage.FlexUtilController) MO unsets its reference to this deleted MO.",
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
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"write_enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"write_error_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"write_error_threshold": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func dataSourceStorageFlexUtilPhysicalDriveRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "storage/FlexUtilPhysicalDrives"
	var o models.StorageFlexUtilPhysicalDrive
	if v, ok := d.GetOk("block_size"); ok {
		x := (v.(string))
		o.BlockSize = x
	}
	if v, ok := d.GetOk("capacity"); ok {
		x := (v.(string))
		o.Capacity = x
	}
	if v, ok := d.GetOk("controller"); ok {
		x := (v.(string))
		o.Controller = x
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.DeviceMoID = x
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.Dn = x
	}
	if v, ok := d.GetOk("drives_enabled"); ok {
		x := (v.(string))
		o.DrivesEnabled = x
	}
	if v, ok := d.GetOk("health"); ok {
		x := (v.(string))
		o.Health = x
	}
	if v, ok := d.GetOk("manufacturer_date"); ok {
		x := (v.(string))
		o.ManufacturerDate = x
	}
	if v, ok := d.GetOk("manufacturer_id"); ok {
		x := (v.(string))
		o.ManufacturerID = x
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
	if v, ok := d.GetOk("oem_id"); ok {
		x := (v.(string))
		o.OemID = x
	}
	if v, ok := d.GetOk("partition_count"); ok {
		x := (v.(string))
		o.PartitionCount = x
	}
	if v, ok := d.GetOk("pd_status"); ok {
		x := (v.(string))
		o.PdStatus = x
	}
	if v, ok := d.GetOk("physical_drive"); ok {
		x := (v.(string))
		o.PhysicalDrive = x
	}
	if v, ok := d.GetOk("product_name"); ok {
		x := (v.(string))
		o.ProductName = x
	}
	if v, ok := d.GetOk("product_revision"); ok {
		x := (v.(string))
		o.ProductRevision = x
	}
	if v, ok := d.GetOk("read_error_count"); ok {
		x := (v.(string))
		o.ReadErrorCount = x
	}
	if v, ok := d.GetOk("read_error_threshold"); ok {
		x := (v.(string))
		o.ReadErrorThreshold = x
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
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x
	}
	if v, ok := d.GetOk("write_enabled"); ok {
		x := (v.(string))
		o.WriteEnabled = x
	}
	if v, ok := d.GetOk("write_error_count"); ok {
		x := (v.(string))
		o.WriteErrorCount = x
	}
	if v, ok := d.GetOk("write_error_threshold"); ok {
		x := (v.(string))
		o.WriteErrorThreshold = x
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
			var s models.StorageFlexUtilPhysicalDrive
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("block_size", (s.BlockSize)); err != nil {
				return err
			}
			if err := d.Set("capacity", (s.Capacity)); err != nil {
				return err
			}
			if err := d.Set("controller", (s.Controller)); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoID)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}
			if err := d.Set("drives_enabled", (s.DrivesEnabled)); err != nil {
				return err
			}
			if err := d.Set("health", (s.Health)); err != nil {
				return err
			}
			if err := d.Set("manufacturer_date", (s.ManufacturerDate)); err != nil {
				return err
			}
			if err := d.Set("manufacturer_id", (s.ManufacturerID)); err != nil {
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
			if err := d.Set("oem_id", (s.OemID)); err != nil {
				return err
			}
			if err := d.Set("partition_count", (s.PartitionCount)); err != nil {
				return err
			}
			if err := d.Set("pd_status", (s.PdStatus)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("physical_drive", (s.PhysicalDrive)); err != nil {
				return err
			}
			if err := d.Set("product_name", (s.ProductName)); err != nil {
				return err
			}
			if err := d.Set("product_revision", (s.ProductRevision)); err != nil {
				return err
			}
			if err := d.Set("read_error_count", (s.ReadErrorCount)); err != nil {
				return err
			}
			if err := d.Set("read_error_threshold", (s.ReadErrorThreshold)); err != nil {
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

			if err := d.Set("storage_flex_util_controller", flattenMapStorageFlexUtilControllerRef(s.StorageFlexUtilController, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			if err := d.Set("write_enabled", (s.WriteEnabled)); err != nil {
				return err
			}
			if err := d.Set("write_error_count", (s.WriteErrorCount)); err != nil {
				return err
			}
			if err := d.Set("write_error_threshold", (s.WriteErrorThreshold)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
