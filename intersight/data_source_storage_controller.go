package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceStorageController() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageControllerRead,
		Schema: map[string]*schema.Schema{
			"compute_board": {
				Description: "A collection of references to the [compute.Board](mo://compute.Board) Managed Object.When this managed object is deleted, the referenced [compute.Board](mo://compute.Board) MO unsets its reference to this deleted MO.",
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
			"controller_flags": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"controller_id": {
				Description: "It shows the Id of controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"controller_status": {
				Description: "It shows the current status of controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"hw_revision": {
				Description: "It shows the hardware revision of controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"oob_interface_supported": {
				Description: "It shows CIMC support for out-of-band configuration of controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_state": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pci_addr": {
				Description: "It shows the current pci address of controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pci_slot": {
				Description: "It shows the pci slot name for the controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"physical_disk_extensions": {
				Description: "This object is created to indicate a SCSI controller has physical connectivity to specified physical disk.",
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
			"physical_disks": {
				Description: "Physical Disk on a server.",
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
			"presence": {
				Description: "It shows physical presence or absence of the controller on server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"raid_support": {
				Description: "It shows the RAID levels supported by controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rebuild_rate": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"running_firmware": {
				Description: "",
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
			"self_encrypt_enabled": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
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
			"type": {
				Description: "Controller types are SAS, SATA, PCH, NVME.",
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
			"virtual_drive_extensions": {
				Description: "",
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
			"virtual_drives": {
				Description: "Storage physical drives are grouped as Drive Group, a drive group then can be partitioned into virtual drives.",
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
		},
	}
}
func dataSourceStorageControllerRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "storage/Controllers"
	var o models.StorageController
	if v, ok := d.GetOk("controller_flags"); ok {
		x := (v.(string))
		o.ControllerFlags = x
	}
	if v, ok := d.GetOk("controller_id"); ok {
		x := (v.(string))
		o.ControllerID = x
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
	if v, ok := d.GetOk("hw_revision"); ok {
		x := (v.(string))
		o.HwRevision = x
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
	if v, ok := d.GetOk("oob_interface_supported"); ok {
		x := (v.(string))
		o.OobInterfaceSupported = x
	}
	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.OperState = x
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.Operability = x
	}
	if v, ok := d.GetOk("pci_addr"); ok {
		x := (v.(string))
		o.PciAddr = x
	}
	if v, ok := d.GetOk("pci_slot"); ok {
		x := (v.(string))
		o.PciSlot = x
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.Presence = x
	}
	if v, ok := d.GetOk("raid_support"); ok {
		x := (v.(string))
		o.RaidSupport = x
	}
	if v, ok := d.GetOk("rebuild_rate"); ok {
		x := (v.(string))
		o.RebuildRate = x
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.Revision = x
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.Rn = x
	}
	if v, ok := d.GetOk("self_encrypt_enabled"); ok {
		x := (v.(string))
		o.SelfEncryptEnabled = x
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.Serial = x
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.Type = x
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x
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
			var s models.StorageController
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("compute_board", flattenMapComputeBoardRef(s.ComputeBoard, d)); err != nil {
				return err
			}
			if err := d.Set("controller_flags", (s.ControllerFlags)); err != nil {
				return err
			}
			if err := d.Set("controller_id", (s.ControllerID)); err != nil {
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
			if err := d.Set("hw_revision", (s.HwRevision)); err != nil {
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
			if err := d.Set("oob_interface_supported", (s.OobInterfaceSupported)); err != nil {
				return err
			}
			if err := d.Set("oper_state", (s.OperState)); err != nil {
				return err
			}
			if err := d.Set("operability", (s.Operability)); err != nil {
				return err
			}
			if err := d.Set("pci_addr", (s.PciAddr)); err != nil {
				return err
			}
			if err := d.Set("pci_slot", (s.PciSlot)); err != nil {
				return err
			}

			if err := d.Set("physical_disk_extensions", flattenListStoragePhysicalDiskExtensionRef(s.PhysicalDiskExtensions, d)); err != nil {
				return err
			}

			if err := d.Set("physical_disks", flattenListStoragePhysicalDiskRef(s.PhysicalDisks, d)); err != nil {
				return err
			}
			if err := d.Set("presence", (s.Presence)); err != nil {
				return err
			}
			if err := d.Set("raid_support", (s.RaidSupport)); err != nil {
				return err
			}
			if err := d.Set("rebuild_rate", (s.RebuildRate)); err != nil {
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

			if err := d.Set("running_firmware", flattenListFirmwareRunningFirmwareRef(s.RunningFirmware, d)); err != nil {
				return err
			}
			if err := d.Set("self_encrypt_enabled", (s.SelfEncryptEnabled)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("type", (s.Type)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}

			if err := d.Set("virtual_drive_extensions", flattenListStorageVirtualDriveExtensionRef(s.VirtualDriveExtensions, d)); err != nil {
				return err
			}

			if err := d.Set("virtual_drives", flattenListStorageVirtualDriveRef(s.VirtualDrives, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
