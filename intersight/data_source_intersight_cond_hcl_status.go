package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCondHclStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCondHclStatusRead,
		Schema: map[string]*schema.Schema{
			"component_status": {
				Description: "The overall status for the components found in the HCL. This will provide the HCL validation status for all the components. It can be one of the following. \"Validated\" - all the components hardware/software profiles are listed in the HCL. \"Not-Listed\" - one or more components hardware/software profiles are not listed in the HCL \"Incomplete\" - the components are not evaluated as the server's software/hardware profiles are not listed in the HCL. \"Not-Evaluated\" - The components are not evaluated against the HCL because it is exempted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"details": {
				Description: "The collection of all the detailed related components for which we performed HCL validation.",
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
			"hardware_status": {
				Description: "The server model, processor and firmware are considered as part of the hardware profile for the server. This will provide the HCL validation status for the hardware profile. For the failure reason see the serverReason property. The status can be one of the following \"Validated\" - The server model, processor and firmware combination is listed in the HCL \"Not-Listed\" - The server model, processor and firmware combination is not listed in the HCL. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_firmware_version": {
				Description: "The current CIMC version for the server normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_model": {
				Description: "The managed object's model to validate normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_os_vendor": {
				Description: "The OS Vendor for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_os_version": {
				Description: "The OS Version for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_processor": {
				Description: "The managed object's processor to validate if applicable normalized for querying HCL data. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_firmware_version": {
				Description: "The current CIMC version for the server as received from inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_model": {
				Description: "The managed object's model to validate as received from the inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_os_vendor": {
				Description: "The OS Vendor for the managed object to validate as received from inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_os_version": {
				Description: "The OS Version for the managed object to validate as received from inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_processor": {
				Description: "The managed object's processor to validate if applicable as received from inventory. It is empty if we are missing this information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"managed_object": {
				Description: "The managed object relationship for this HCLStatus.",
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
			"reason": {
				Description: "The reason for the HCL status. It will be one of the following \"Missing-Os-Info\" - we are missing os information in the inventory from the device connector \"Incompatible-Components\" - we have 1 or more components with \"Not-Validated\" status \"Compatible\" - all the components have \"Validated\" status. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"registered_device": {
				Description: "The Relationship to the registered device. We need this in order to correctly set permissions during device claim.",
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
			"server_reason": {
				Description: "The reason generated by the server's HCL validation. For HCL the evaluation can be seen in three logical stages 1. Evaluate the server's hardware status 2. Evaluate the server's software status 3. Evaluate the server's components (each component has its own hardware/software evaluation) The evaluation does not proceed to the next stage until the previous stage is evaluated. Therefore there can be only one validation reason. \"Incompatible-Server\" - the server model is not listed in the HCL. \"Incompatible-Processor\" - the server model and processor combination is not listed in HCL. \"Incompatible-Firmware\" - the server model, processor and server firmware is not listed in HCL. \"Missing-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Components\" - there is one or more components with \"Not-Validated\" status \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). \"Compatible\" - the server and all its components are validated. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_status": {
				Description: "The OS vendor and version are considered part of the software profile for the server. This will provide the HCL validation status for the software profile. For the failure reason see the serverReason property. The status can be be one of the following \"Validated\" - The os vendor/version is listed in the HCL for the server model, processor and firmware \"Not-Listed\" - The os vendor/version is not listed in the HCL for the server model, processor and firmware \"Incomplete\" - The inventory is missing os vendor/version and HCL validation was not performed. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "The HCL compatibility status of the managed object. The status can be one of the following \"Incomplete\" - there is no enough information to evaluate against the HCL data \"Validated\" - all components have been evaluated against the HCL and they all have \"Validated\" status \"Not-Listed\" - all components have been evaluated against the HCL and one or more have \"Not-Listed\" status. \"Not-Evaluated\" - server is not evaluated against the HCL because it is exempted.",
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
		},
	}
}
func dataSourceCondHclStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "cond/HclStatuses"
	var o models.CondHclStatus
	if v, ok := d.GetOk("component_status"); ok {
		x := (v.(string))
		o.ComponentStatus = &x
	}
	if v, ok := d.GetOk("hardware_status"); ok {
		x := (v.(string))
		o.HardwareStatus = &x
	}
	if v, ok := d.GetOk("hcl_firmware_version"); ok {
		x := (v.(string))
		o.HclFirmwareVersion = x
	}
	if v, ok := d.GetOk("hcl_model"); ok {
		x := (v.(string))
		o.HclModel = x
	}
	if v, ok := d.GetOk("hcl_os_vendor"); ok {
		x := (v.(string))
		o.HclOsVendor = x
	}
	if v, ok := d.GetOk("hcl_os_version"); ok {
		x := (v.(string))
		o.HclOsVersion = x
	}
	if v, ok := d.GetOk("hcl_processor"); ok {
		x := (v.(string))
		o.HclProcessor = x
	}
	if v, ok := d.GetOk("inv_firmware_version"); ok {
		x := (v.(string))
		o.InvFirmwareVersion = x
	}
	if v, ok := d.GetOk("inv_model"); ok {
		x := (v.(string))
		o.InvModel = x
	}
	if v, ok := d.GetOk("inv_os_vendor"); ok {
		x := (v.(string))
		o.InvOsVendor = x
	}
	if v, ok := d.GetOk("inv_os_version"); ok {
		x := (v.(string))
		o.InvOsVersion = x
	}
	if v, ok := d.GetOk("inv_processor"); ok {
		x := (v.(string))
		o.InvProcessor = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("reason"); ok {
		x := (v.(string))
		o.Reason = &x
	}
	if v, ok := d.GetOk("server_reason"); ok {
		x := (v.(string))
		o.ServerReason = &x
	}
	if v, ok := d.GetOk("software_status"); ok {
		x := (v.(string))
		o.SoftwareStatus = &x
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = &x
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
			var s models.CondHclStatus
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("component_status", (s.ComponentStatus)); err != nil {
				return err
			}

			if err := d.Set("details", flattenListCondHclStatusDetailRef(s.Details, d)); err != nil {
				return err
			}
			if err := d.Set("hardware_status", (s.HardwareStatus)); err != nil {
				return err
			}
			if err := d.Set("hcl_firmware_version", (s.HclFirmwareVersion)); err != nil {
				return err
			}
			if err := d.Set("hcl_model", (s.HclModel)); err != nil {
				return err
			}
			if err := d.Set("hcl_os_vendor", (s.HclOsVendor)); err != nil {
				return err
			}
			if err := d.Set("hcl_os_version", (s.HclOsVersion)); err != nil {
				return err
			}
			if err := d.Set("hcl_processor", (s.HclProcessor)); err != nil {
				return err
			}
			if err := d.Set("inv_firmware_version", (s.InvFirmwareVersion)); err != nil {
				return err
			}
			if err := d.Set("inv_model", (s.InvModel)); err != nil {
				return err
			}
			if err := d.Set("inv_os_vendor", (s.InvOsVendor)); err != nil {
				return err
			}
			if err := d.Set("inv_os_version", (s.InvOsVersion)); err != nil {
				return err
			}
			if err := d.Set("inv_processor", (s.InvProcessor)); err != nil {
				return err
			}

			if err := d.Set("managed_object", flattenMapInventoryBaseRef(s.ManagedObject, d)); err != nil {
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
			if err := d.Set("reason", (s.Reason)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("server_reason", (s.ServerReason)); err != nil {
				return err
			}
			if err := d.Set("software_status", (s.SoftwareStatus)); err != nil {
				return err
			}
			if err := d.Set("status", (s.Status)); err != nil {
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
