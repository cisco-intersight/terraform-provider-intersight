package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceCondHclStatusDetail() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCondHclStatusDetailRead,
		Schema: map[string]*schema.Schema{
			"component": {
				Description: "The related component associated to this status detail HclStatusDetail (adapter or storage controller).",
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
			"hardware_status": {
				Description: "The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \"Incompatible-Server-With-Component\" - the server model and component combination is not listed in HCL \"Incompatible-Firmware\" - The server's firmware is not listed for this component's hardware profile \"Incompatible-Component\" - the component's model is not listed in the HCL \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the hardware profile was not evaulated for the component because the server's hw/sw status is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_cimc_version": {
				Description: "The current CIMC version for the server normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_driver_name": {
				Description: "The current driver name of the component we are validating normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_driver_version": {
				Description: "The current driver version of the component we are validating normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_firmware_version": {
				Description: "The current firmware version of the component model normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_model": {
				Description: "The component model we are trying to validate normalized for querying HCL data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hcl_status": {
				Description: "A collection of references to the [cond.HclStatus](mo://cond.HclStatus) Managed Object.When this managed object is deleted, the referenced [cond.HclStatus](mo://cond.HclStatus) MO unsets its reference to this deleted MO.",
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
			"inv_cimc_version": {
				Description: "The current CIMC version for the server as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_driver_name": {
				Description: "The current driver name of the component we are validating as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_driver_version": {
				Description: "The current driver version of the component we are validating as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_firmware_version": {
				Description: "The current firmware version of the component model as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inv_model": {
				Description: "The component model we are trying to validate as received from inventory.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"reason": {
				Description: "The reason for the status. The reason can be one of \"Incompatible-Server-With-Component\" - HCL validation has failed because the server model is not validated with this component \"Incompatible-Processor\" - HCL validation has failed because the processor is not validated with this server \"Incompatible-Os-Info\" - HCL validation has failed because the os vendor and version is not validated with this server \"Incompatible-Component-Model\" - HCL validation has failed because the component model is not validated \"Incompatible-Firmware\" - HCL validation has failed because the component or server firmware version is not validated \"Incompatible-Driver\" - HCL validation has failed because the driver version is not validated \"Incompatible-Firmware-Driver\" - HCL validation has failed because the firmware version and driver version is not validated \"Missing-Os-Driver-Info\" - HCL validation was not performed because we are missing os driver information form the inventory \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Service-Error\" - HCL data service is available but an error occured when making the request or parsing the response \"Unrecognized-Protocol\" - This service does not recognize the reason code in the response from the HCL data service \"Compatible\" - this component's inventory data has \"Validated\" status with the HCL. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_status": {
				Description: "The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \"Incompatible-Firmware\" - the component's firmware is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Driver\" - the component's driver is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Firmware-Driver\" - the component's firmware and driver are not listed under the server's hardware and software profile and the component's hardware profile \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the component's hardware status was not evaluated because the server's hardware or software profile is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \"Unknown\" - we do not have enough information to evaluate against the HCL data \"Validated\" - we have validated this component against the HCL and it has \"Validated\" status \"Not-Validated\" - we have validated this component against the HCL and it has \"Not-Validated\" status. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.",
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
func dataSourceCondHclStatusDetailRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "cond/HclStatusDetails"
	var o models.CondHclStatusDetail
	if v, ok := d.GetOk("hardware_status"); ok {
		x := (v.(string))
		o.HardwareStatus = &x
	}
	if v, ok := d.GetOk("hcl_cimc_version"); ok {
		x := (v.(string))
		o.HclCimcVersion = x
	}
	if v, ok := d.GetOk("hcl_driver_name"); ok {
		x := (v.(string))
		o.HclDriverName = x
	}
	if v, ok := d.GetOk("hcl_driver_version"); ok {
		x := (v.(string))
		o.HclDriverVersion = x
	}
	if v, ok := d.GetOk("hcl_firmware_version"); ok {
		x := (v.(string))
		o.HclFirmwareVersion = x
	}
	if v, ok := d.GetOk("hcl_model"); ok {
		x := (v.(string))
		o.HclModel = x
	}
	if v, ok := d.GetOk("inv_cimc_version"); ok {
		x := (v.(string))
		o.InvCimcVersion = x
	}
	if v, ok := d.GetOk("inv_driver_name"); ok {
		x := (v.(string))
		o.InvDriverName = x
	}
	if v, ok := d.GetOk("inv_driver_version"); ok {
		x := (v.(string))
		o.InvDriverVersion = x
	}
	if v, ok := d.GetOk("inv_firmware_version"); ok {
		x := (v.(string))
		o.InvFirmwareVersion = x
	}
	if v, ok := d.GetOk("inv_model"); ok {
		x := (v.(string))
		o.InvModel = x
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
			var s models.CondHclStatusDetail
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("component", flattenMapInventoryBaseRef(s.Component, d)); err != nil {
				return err
			}
			if err := d.Set("hardware_status", (s.HardwareStatus)); err != nil {
				return err
			}
			if err := d.Set("hcl_cimc_version", (s.HclCimcVersion)); err != nil {
				return err
			}
			if err := d.Set("hcl_driver_name", (s.HclDriverName)); err != nil {
				return err
			}
			if err := d.Set("hcl_driver_version", (s.HclDriverVersion)); err != nil {
				return err
			}
			if err := d.Set("hcl_firmware_version", (s.HclFirmwareVersion)); err != nil {
				return err
			}
			if err := d.Set("hcl_model", (s.HclModel)); err != nil {
				return err
			}

			if err := d.Set("hcl_status", flattenMapCondHclStatusRef(s.HclStatus, d)); err != nil {
				return err
			}
			if err := d.Set("inv_cimc_version", (s.InvCimcVersion)); err != nil {
				return err
			}
			if err := d.Set("inv_driver_name", (s.InvDriverName)); err != nil {
				return err
			}
			if err := d.Set("inv_driver_version", (s.InvDriverVersion)); err != nil {
				return err
			}
			if err := d.Set("inv_firmware_version", (s.InvFirmwareVersion)); err != nil {
				return err
			}
			if err := d.Set("inv_model", (s.InvModel)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("reason", (s.Reason)); err != nil {
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
