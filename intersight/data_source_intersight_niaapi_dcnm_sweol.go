package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiDcnmSweol() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiDcnmSweolRead,
		Schema: map[string]*schema.Schema{
			"affected_versions": {
				Description: "String contains the Release versions affected by this notice, seperated by comma.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date": {
				Description: "Date time of this notice Announced.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date_epoch": {
				Description: "Epoch time of this notice Announced.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"bulletin_no": {
				Description: "The bulletinno of this software release end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The description of this software release end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date": {
				Description: "Date time of End of New service attachment .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date_epoch": {
				Description: "Epoch time of End of New service attachment .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_service_contract_renewal_date": {
				Description: "Date time of End of Renewal of service contract .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_service_contract_renewal_date_epoch": {
				Description: "Epoch time of End of Renewal of service contract .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_sw_maintenance_date": {
				Description: "Date time of End of Maintenance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_sw_maintenance_date_epoch": {
				Description: "Epoch time of End of Maintenance.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"headline": {
				Description: "The title of this software release end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support": {
				Description: "Date time of Last day of Support .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support_epoch": {
				Description: "Epoch time of Last day of Support .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"last_ship_date": {
				Description: "Date time of Lastship Date.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_ship_date_epoch": {
				Description: "Epoch time of Lastship Date.",
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
			"software_eol_url": {
				Description: "Software end of life notice URL link to the notice webpage.",
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
func dataSourceNiaapiDcnmSweolRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "niaapi/DcnmSweols"
	var o models.NiaapiDcnmSweol
	if v, ok := d.GetOk("affected_versions"); ok {
		x := (v.(string))
		o.AffectedVersions = x
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
	if v, ok := d.GetOk("software_eol_url"); ok {
		x := (v.(string))
		o.SoftwareEolURL = x
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
			var s models.NiaapiDcnmSweol
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("affected_versions", (s.AffectedVersions)); err != nil {
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
			if err := d.Set("software_eol_url", (s.SoftwareEolURL)); err != nil {
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
