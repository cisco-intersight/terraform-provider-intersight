package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceStoragePhysicalDisk() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStoragePhysicalDiskRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"block_size": {
				Description: "The block size of the physical disk in bytes.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"bootable": {
				Description: "This field identifies the disk drive as bootable if set to true.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"configuration_checkpoint": {
				Description: "The current configuration checkpoint of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"configuration_state": {
				Description: "The current configuration state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_mo_id": {
				Description: "The database identifier of the registered device of an object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"discovered_path": {
				Description: "The discovered path of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"disk_id": {
				Description: "This field identifies the ID assigned to physical disks.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"disk_state": {
				Description: "This field identifies the health of the disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drive_firmware": {
				Description: "This field identifies the disk firmware running in the disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"drive_state": {
				Description: "The drive state as reported by the controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fde_capable": {
				Description: "Full-Disk Encryption capability parameter of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hot_spare_type": {
				Description: "Type of hotspare configured on the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inventory_device_info": {
				Description: "A reference to a inventoryDeviceInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
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
			"link_speed": {
				Description: "The speed of the link between the drive and the controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"link_state": {
				Description: "The current link state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"locator_led": {
				Description: "A reference to a equipmentLocatorLed resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
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
			"num_blocks": {
				Description: "The number of blocks present on the physical disk.",
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
			"oper_power_state": {
				Description: "Operational power of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_qualifier_reason": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operability": {
				Description: "This field identifies the disk operability of the disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"physical_block_size": {
				Description: "The block size of the installed physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"physical_disk_extensions": {
				Description: "An array of relationships to storagePhysicalDiskExtension resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
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
				Computed: true,
			},
			"pid": {
				Description: "This field identifies the Product ID for physicalDisk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "The presence state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Description: "This field identifies the disk protocol used for communication.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"raw_size": {
				Description: "The raw size of the physical disk in MB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
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
			"revision": {
				Description: "This field identifies the revision of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"running_firmware": {
				Description: "An array of relationships to firmwareRunningFirmware resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
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
			"sas_ports": {
				Description: "An array of relationships to storageSasPort resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
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
			"secured": {
				Description: "This field identifies whether the disk is encrypted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"size": {
				Description: "The size of the physical disk in MB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"storage_controller": {
				Description: "A reference to a storageController resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
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
			"storage_enclosure": {
				Description: "A reference to a storageEnclosure resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
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
			"thermal": {
				Description: "Thermal state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "This field identifies the type of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"variant_type": {
				Description: "The variant type of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceStoragePhysicalDiskRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewStoragePhysicalDiskWithDefaults()
	if v, ok := d.GetOk("block_size"); ok {
		x := (v.(string))
		o.SetBlockSize(x)
	}
	if v, ok := d.GetOk("bootable"); ok {
		x := (v.(string))
		o.SetBootable(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("configuration_checkpoint"); ok {
		x := (v.(string))
		o.SetConfigurationCheckpoint(x)
	}
	if v, ok := d.GetOk("configuration_state"); ok {
		x := (v.(string))
		o.SetConfigurationState(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("discovered_path"); ok {
		x := (v.(string))
		o.SetDiscoveredPath(x)
	}
	if v, ok := d.GetOk("disk_id"); ok {
		x := (v.(string))
		o.SetDiskId(x)
	}
	if v, ok := d.GetOk("disk_state"); ok {
		x := (v.(string))
		o.SetDiskState(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("drive_firmware"); ok {
		x := (v.(string))
		o.SetDriveFirmware(x)
	}
	if v, ok := d.GetOk("drive_state"); ok {
		x := (v.(string))
		o.SetDriveState(x)
	}
	if v, ok := d.GetOk("fde_capable"); ok {
		x := (v.(string))
		o.SetFdeCapable(x)
	}
	if v, ok := d.GetOk("hot_spare_type"); ok {
		x := (v.(string))
		o.SetHotSpareType(x)
	}
	if v, ok := d.GetOk("link_speed"); ok {
		x := (v.(string))
		o.SetLinkSpeed(x)
	}
	if v, ok := d.GetOk("link_state"); ok {
		x := (v.(string))
		o.SetLinkState(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("num_blocks"); ok {
		x := (v.(string))
		o.SetNumBlocks(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oper_power_state"); ok {
		x := (v.(string))
		o.SetOperPowerState(x)
	}
	if v, ok := d.GetOk("oper_qualifier_reason"); ok {
		x := (v.(string))
		o.SetOperQualifierReason(x)
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.SetOperability(x)
	}
	if v, ok := d.GetOk("physical_block_size"); ok {
		x := (v.(string))
		o.SetPhysicalBlockSize(x)
	}
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}
	if v, ok := d.GetOk("raw_size"); ok {
		x := (v.(string))
		o.SetRawSize(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("secured"); ok {
		x := (v.(string))
		o.SetSecured(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := (v.(string))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("thermal"); ok {
		x := (v.(string))
		o.SetThermal(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("variant_type"); ok {
		x := (v.(string))
		o.SetVariantType(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.StorageApi.GetStoragePhysicalDiskList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.StoragePhysicalDiskList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to StoragePhysicalDisk: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewStoragePhysicalDiskWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("block_size", (s.GetBlockSize())); err != nil {
				return fmt.Errorf("error occurred while setting property BlockSize: %+v", err)
			}
			if err := d.Set("bootable", (s.GetBootable())); err != nil {
				return fmt.Errorf("error occurred while setting property Bootable: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("configuration_checkpoint", (s.GetConfigurationCheckpoint())); err != nil {
				return fmt.Errorf("error occurred while setting property ConfigurationCheckpoint: %+v", err)
			}
			if err := d.Set("configuration_state", (s.GetConfigurationState())); err != nil {
				return fmt.Errorf("error occurred while setting property ConfigurationState: %+v", err)
			}
			if err := d.Set("device_mo_id", (s.GetDeviceMoId())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceMoId: %+v", err)
			}
			if err := d.Set("discovered_path", (s.GetDiscoveredPath())); err != nil {
				return fmt.Errorf("error occurred while setting property DiscoveredPath: %+v", err)
			}
			if err := d.Set("disk_id", (s.GetDiskId())); err != nil {
				return fmt.Errorf("error occurred while setting property DiskId: %+v", err)
			}
			if err := d.Set("disk_state", (s.GetDiskState())); err != nil {
				return fmt.Errorf("error occurred while setting property DiskState: %+v", err)
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return fmt.Errorf("error occurred while setting property Dn: %+v", err)
			}
			if err := d.Set("drive_firmware", (s.GetDriveFirmware())); err != nil {
				return fmt.Errorf("error occurred while setting property DriveFirmware: %+v", err)
			}
			if err := d.Set("drive_state", (s.GetDriveState())); err != nil {
				return fmt.Errorf("error occurred while setting property DriveState: %+v", err)
			}
			if err := d.Set("fde_capable", (s.GetFdeCapable())); err != nil {
				return fmt.Errorf("error occurred while setting property FdeCapable: %+v", err)
			}
			if err := d.Set("hot_spare_type", (s.GetHotSpareType())); err != nil {
				return fmt.Errorf("error occurred while setting property HotSpareType: %+v", err)
			}

			if err := d.Set("inventory_device_info", flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property InventoryDeviceInfo: %+v", err)
			}
			if err := d.Set("link_speed", (s.GetLinkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property LinkSpeed: %+v", err)
			}
			if err := d.Set("link_state", (s.GetLinkState())); err != nil {
				return fmt.Errorf("error occurred while setting property LinkState: %+v", err)
			}

			if err := d.Set("locator_led", flattenMapEquipmentLocatorLedRelationship(s.GetLocatorLed(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property LocatorLed: %+v", err)
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return fmt.Errorf("error occurred while setting property Model: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("num_blocks", (s.GetNumBlocks())); err != nil {
				return fmt.Errorf("error occurred while setting property NumBlocks: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("oper_power_state", (s.GetOperPowerState())); err != nil {
				return fmt.Errorf("error occurred while setting property OperPowerState: %+v", err)
			}
			if err := d.Set("oper_qualifier_reason", (s.GetOperQualifierReason())); err != nil {
				return fmt.Errorf("error occurred while setting property OperQualifierReason: %+v", err)
			}
			if err := d.Set("operability", (s.GetOperability())); err != nil {
				return fmt.Errorf("error occurred while setting property Operability: %+v", err)
			}
			if err := d.Set("physical_block_size", (s.GetPhysicalBlockSize())); err != nil {
				return fmt.Errorf("error occurred while setting property PhysicalBlockSize: %+v", err)
			}

			if err := d.Set("physical_disk_extensions", flattenListStoragePhysicalDiskExtensionRelationship(s.GetPhysicalDiskExtensions(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PhysicalDiskExtensions: %+v", err)
			}
			if err := d.Set("pid", (s.GetPid())); err != nil {
				return fmt.Errorf("error occurred while setting property Pid: %+v", err)
			}
			if err := d.Set("presence", (s.GetPresence())); err != nil {
				return fmt.Errorf("error occurred while setting property Presence: %+v", err)
			}
			if err := d.Set("protocol", (s.GetProtocol())); err != nil {
				return fmt.Errorf("error occurred while setting property Protocol: %+v", err)
			}
			if err := d.Set("raw_size", (s.GetRawSize())); err != nil {
				return fmt.Errorf("error occurred while setting property RawSize: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}
			if err := d.Set("revision", (s.GetRevision())); err != nil {
				return fmt.Errorf("error occurred while setting property Revision: %+v", err)
			}
			if err := d.Set("rn", (s.GetRn())); err != nil {
				return fmt.Errorf("error occurred while setting property Rn: %+v", err)
			}

			if err := d.Set("running_firmware", flattenListFirmwareRunningFirmwareRelationship(s.GetRunningFirmware(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RunningFirmware: %+v", err)
			}

			if err := d.Set("sas_ports", flattenListStorageSasPortRelationship(s.GetSasPorts(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property SasPorts: %+v", err)
			}
			if err := d.Set("secured", (s.GetSecured())); err != nil {
				return fmt.Errorf("error occurred while setting property Secured: %+v", err)
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return fmt.Errorf("error occurred while setting property Serial: %+v", err)
			}
			if err := d.Set("size", (s.GetSize())); err != nil {
				return fmt.Errorf("error occurred while setting property Size: %+v", err)
			}

			if err := d.Set("storage_controller", flattenMapStorageControllerRelationship(s.GetStorageController(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property StorageController: %+v", err)
			}

			if err := d.Set("storage_enclosure", flattenMapStorageEnclosureRelationship(s.GetStorageEnclosure(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property StorageEnclosure: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("thermal", (s.GetThermal())); err != nil {
				return fmt.Errorf("error occurred while setting property Thermal: %+v", err)
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return fmt.Errorf("error occurred while setting property Type: %+v", err)
			}
			if err := d.Set("variant_type", (s.GetVariantType())); err != nil {
				return fmt.Errorf("error occurred while setting property VariantType: %+v", err)
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return fmt.Errorf("error occurred while setting property Vendor: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
