package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiApicHweol() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiApicHweolRead,
		Schema: map[string]*schema.Schema{
			"affected_pids": {
				Description: "String contains the PID of hardwares affected by this notice, seperated by comma.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date": {
				Description: "When this notice is announced.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date_epoch": {
				Description: "Epoch time of Announcement Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"bulletin_no": {
				Description: "The bulletinno of this hardware end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The description of this hardware end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date": {
				Description: "Date time of end of new services attachment  .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date_epoch": {
				Description: "Epoch time of New service attachment Date .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_routine_failure_analysis_date": {
				Description: "Date time of end of routinefailure analysis.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_routine_failure_analysis_date_epoch": {
				Description: "Epoch time of End of Routine Failure Analysis Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_sale_date": {
				Description: "When this hardware will end sale.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_sale_date_epoch": {
				Description: "Epoch time of End of Sale Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_security_support": {
				Description: "Date time of end of security support .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_security_support_epoch": {
				Description: "Epoch time of End of Security Support Date .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_service_contract_renewal_date": {
				Description: "Date time of end of service contract renew .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_service_contract_renewal_date_epoch": {
				Description: "Epoch time of End of Renewal service contract.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_sw_maintenance_date": {
				Description: "The date of end of maintainance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_sw_maintenance_date_epoch": {
				Description: "Epoch time of End of maintenance Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"hardware_eol_url": {
				Description: "Hardware end of sale URL link to the notice webpage.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"headline": {
				Description: "The title of this hardware end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support": {
				Description: "Date time of end of support .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support_epoch": {
				Description: "Epoch time of last date of support .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"last_ship_date": {
				Description: "Date time of Lastship Date.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_ship_date_epoch": {
				Description: "Epoch time of last ship Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"migration_url": {
				Description: "The URL of this migration notice.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.",
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
		},
	}
}
func dataSourceNiaapiApicHweolRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "niaapi/ApicHweols"
	var o models.NiaapiApicHweol
	if v, ok := d.GetOk("affected_pids"); ok {
		x := (v.(string))
		o.AffectedPids = x
	}
	if v, ok := d.GetOk("announcement_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.AnnouncementDate = x
	}
	if v, ok := d.GetOk("announcement_date_epoch"); ok {
		x := int64(v.(int))
		o.AnnouncementDateEpoch = x
	}
	if v, ok := d.GetOk("bulletin_no"); ok {
		x := (v.(string))
		o.BulletinNo = x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("endof_new_service_attachment_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndofNewServiceAttachmentDate = x
	}
	if v, ok := d.GetOk("endof_new_service_attachment_date_epoch"); ok {
		x := int64(v.(int))
		o.EndofNewServiceAttachmentDateEpoch = x
	}
	if v, ok := d.GetOk("endof_routine_failure_analysis_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndofRoutineFailureAnalysisDate = x
	}
	if v, ok := d.GetOk("endof_routine_failure_analysis_date_epoch"); ok {
		x := int64(v.(int))
		o.EndofRoutineFailureAnalysisDateEpoch = x
	}
	if v, ok := d.GetOk("endof_sale_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndofSaleDate = x
	}
	if v, ok := d.GetOk("endof_sale_date_epoch"); ok {
		x := int64(v.(int))
		o.EndofSaleDateEpoch = x
	}
	if v, ok := d.GetOk("endof_security_support"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndofSecuritySupport = x
	}
	if v, ok := d.GetOk("endof_security_support_epoch"); ok {
		x := int64(v.(int))
		o.EndofSecuritySupportEpoch = x
	}
	if v, ok := d.GetOk("endof_service_contract_renewal_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndofServiceContractRenewalDate = x
	}
	if v, ok := d.GetOk("endof_service_contract_renewal_date_epoch"); ok {
		x := int64(v.(int))
		o.EndofServiceContractRenewalDateEpoch = x
	}
	if v, ok := d.GetOk("endof_sw_maintenance_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndofSwMaintenanceDate = x
	}
	if v, ok := d.GetOk("endof_sw_maintenance_date_epoch"); ok {
		x := int64(v.(int))
		o.EndofSwMaintenanceDateEpoch = x
	}
	if v, ok := d.GetOk("hardware_eol_url"); ok {
		x := (v.(string))
		o.HardwareEolURL = x
	}
	if v, ok := d.GetOk("headline"); ok {
		x := (v.(string))
		o.Headline = x
	}
	if v, ok := d.GetOk("last_dateof_support"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.LastDateofSupport = x
	}
	if v, ok := d.GetOk("last_dateof_support_epoch"); ok {
		x := int64(v.(int))
		o.LastDateofSupportEpoch = x
	}
	if v, ok := d.GetOk("last_ship_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.LastShipDate = x
	}
	if v, ok := d.GetOk("last_ship_date_epoch"); ok {
		x := int64(v.(int))
		o.LastShipDateEpoch = x
	}
	if v, ok := d.GetOk("migration_url"); ok {
		x := (v.(string))
		o.MigrationURL = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
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
			var s models.NiaapiApicHweol
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("affected_pids", (s.AffectedPids)); err != nil {
				return err
			}

			if err := d.Set("announcement_date", (s.AnnouncementDate).String()); err != nil {
				return err
			}
			if err := d.Set("announcement_date_epoch", (s.AnnouncementDateEpoch)); err != nil {
				return err
			}
			if err := d.Set("bulletin_no", (s.BulletinNo)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}

			if err := d.Set("endof_new_service_attachment_date", (s.EndofNewServiceAttachmentDate).String()); err != nil {
				return err
			}
			if err := d.Set("endof_new_service_attachment_date_epoch", (s.EndofNewServiceAttachmentDateEpoch)); err != nil {
				return err
			}

			if err := d.Set("endof_routine_failure_analysis_date", (s.EndofRoutineFailureAnalysisDate).String()); err != nil {
				return err
			}
			if err := d.Set("endof_routine_failure_analysis_date_epoch", (s.EndofRoutineFailureAnalysisDateEpoch)); err != nil {
				return err
			}

			if err := d.Set("endof_sale_date", (s.EndofSaleDate).String()); err != nil {
				return err
			}
			if err := d.Set("endof_sale_date_epoch", (s.EndofSaleDateEpoch)); err != nil {
				return err
			}

			if err := d.Set("endof_security_support", (s.EndofSecuritySupport).String()); err != nil {
				return err
			}
			if err := d.Set("endof_security_support_epoch", (s.EndofSecuritySupportEpoch)); err != nil {
				return err
			}

			if err := d.Set("endof_service_contract_renewal_date", (s.EndofServiceContractRenewalDate).String()); err != nil {
				return err
			}
			if err := d.Set("endof_service_contract_renewal_date_epoch", (s.EndofServiceContractRenewalDateEpoch)); err != nil {
				return err
			}

			if err := d.Set("endof_sw_maintenance_date", (s.EndofSwMaintenanceDate).String()); err != nil {
				return err
			}
			if err := d.Set("endof_sw_maintenance_date_epoch", (s.EndofSwMaintenanceDateEpoch)); err != nil {
				return err
			}
			if err := d.Set("hardware_eol_url", (s.HardwareEolURL)); err != nil {
				return err
			}
			if err := d.Set("headline", (s.Headline)); err != nil {
				return err
			}

			if err := d.Set("last_dateof_support", (s.LastDateofSupport).String()); err != nil {
				return err
			}
			if err := d.Set("last_dateof_support_epoch", (s.LastDateofSupportEpoch)); err != nil {
				return err
			}

			if err := d.Set("last_ship_date", (s.LastShipDate).String()); err != nil {
				return err
			}
			if err := d.Set("last_ship_date_epoch", (s.LastShipDateEpoch)); err != nil {
				return err
			}
			if err := d.Set("migration_url", (s.MigrationURL)); err != nil {
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

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
