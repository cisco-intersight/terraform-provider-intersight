package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCondHclStatusDetail() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCondHclStatusDetailRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"component": {
				Description: "A reference to a inventoryBase resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Description: "A reference to a condHclStatus resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
		},
	}
}

func dataSourceCondHclStatusDetailRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewCondHclStatusDetailWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("hardware_status"); ok {
		x := (v.(string))
		o.SetHardwareStatus(x)
	}
	if v, ok := d.GetOk("hcl_cimc_version"); ok {
		x := (v.(string))
		o.SetHclCimcVersion(x)
	}
	if v, ok := d.GetOk("hcl_driver_name"); ok {
		x := (v.(string))
		o.SetHclDriverName(x)
	}
	if v, ok := d.GetOk("hcl_driver_version"); ok {
		x := (v.(string))
		o.SetHclDriverVersion(x)
	}
	if v, ok := d.GetOk("hcl_firmware_version"); ok {
		x := (v.(string))
		o.SetHclFirmwareVersion(x)
	}
	if v, ok := d.GetOk("hcl_model"); ok {
		x := (v.(string))
		o.SetHclModel(x)
	}
	if v, ok := d.GetOk("inv_cimc_version"); ok {
		x := (v.(string))
		o.SetInvCimcVersion(x)
	}
	if v, ok := d.GetOk("inv_driver_name"); ok {
		x := (v.(string))
		o.SetInvDriverName(x)
	}
	if v, ok := d.GetOk("inv_driver_version"); ok {
		x := (v.(string))
		o.SetInvDriverVersion(x)
	}
	if v, ok := d.GetOk("inv_firmware_version"); ok {
		x := (v.(string))
		o.SetInvFirmwareVersion(x)
	}
	if v, ok := d.GetOk("inv_model"); ok {
		x := (v.(string))
		o.SetInvModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("reason"); ok {
		x := (v.(string))
		o.SetReason(x)
	}
	if v, ok := d.GetOk("software_status"); ok {
		x := (v.(string))
		o.SetSoftwareStatus(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.CondApi.GetCondHclStatusDetailList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.CondHclStatusDetailList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to CondHclStatusDetail: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewCondHclStatusDetailWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("component", flattenMapInventoryBaseRelationship(s.Component, d)); err != nil {
				return fmt.Errorf("error occurred while setting property Component: %+v", err)
			}
			if err := d.Set("hardware_status", (s.HardwareStatus)); err != nil {
				return fmt.Errorf("error occurred while setting property HardwareStatus: %+v", err)
			}
			if err := d.Set("hcl_cimc_version", (s.HclCimcVersion)); err != nil {
				return fmt.Errorf("error occurred while setting property HclCimcVersion: %+v", err)
			}
			if err := d.Set("hcl_driver_name", (s.HclDriverName)); err != nil {
				return fmt.Errorf("error occurred while setting property HclDriverName: %+v", err)
			}
			if err := d.Set("hcl_driver_version", (s.HclDriverVersion)); err != nil {
				return fmt.Errorf("error occurred while setting property HclDriverVersion: %+v", err)
			}
			if err := d.Set("hcl_firmware_version", (s.HclFirmwareVersion)); err != nil {
				return fmt.Errorf("error occurred while setting property HclFirmwareVersion: %+v", err)
			}
			if err := d.Set("hcl_model", (s.HclModel)); err != nil {
				return fmt.Errorf("error occurred while setting property HclModel: %+v", err)
			}

			if err := d.Set("hcl_status", flattenMapCondHclStatusRelationship(s.HclStatus, d)); err != nil {
				return fmt.Errorf("error occurred while setting property HclStatus: %+v", err)
			}
			if err := d.Set("inv_cimc_version", (s.InvCimcVersion)); err != nil {
				return fmt.Errorf("error occurred while setting property InvCimcVersion: %+v", err)
			}
			if err := d.Set("inv_driver_name", (s.InvDriverName)); err != nil {
				return fmt.Errorf("error occurred while setting property InvDriverName: %+v", err)
			}
			if err := d.Set("inv_driver_version", (s.InvDriverVersion)); err != nil {
				return fmt.Errorf("error occurred while setting property InvDriverVersion: %+v", err)
			}
			if err := d.Set("inv_firmware_version", (s.InvFirmwareVersion)); err != nil {
				return fmt.Errorf("error occurred while setting property InvFirmwareVersion: %+v", err)
			}
			if err := d.Set("inv_model", (s.InvModel)); err != nil {
				return fmt.Errorf("error occurred while setting property InvModel: %+v", err)
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("reason", (s.Reason)); err != nil {
				return fmt.Errorf("error occurred while setting property Reason: %+v", err)
			}
			if err := d.Set("software_status", (s.SoftwareStatus)); err != nil {
				return fmt.Errorf("error occurred while setting property SoftwareStatus: %+v", err)
			}
			if err := d.Set("status", (s.Status)); err != nil {
				return fmt.Errorf("error occurred while setting property Status: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
