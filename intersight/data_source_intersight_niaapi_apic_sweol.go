package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiApicSweol() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiApicSweolRead,
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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
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
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
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
			"software_eol_url": {
				Description: "Software end of life notice URL link to the notice webpage.",
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

func dataSourceNiaapiApicSweolRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewNiaapiApicSweol()
	if v, ok := d.GetOk("affected_versions"); ok {
		x := (v.(string))
		o.SetAffectedVersions(x)
	}
	if v, ok := d.GetOk("announcement_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetAnnouncementDate(x)
	}
	if v, ok := d.GetOk("announcement_date_epoch"); ok {
		x := int64(v.(int))
		o.SetAnnouncementDateEpoch(x)
	}
	if v, ok := d.GetOk("bulletin_no"); ok {
		x := (v.(string))
		o.SetBulletinNo(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("endof_new_service_attachment_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofNewServiceAttachmentDate(x)
	}
	if v, ok := d.GetOk("endof_new_service_attachment_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofNewServiceAttachmentDateEpoch(x)
	}
	if v, ok := d.GetOk("endof_service_contract_renewal_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofServiceContractRenewalDate(x)
	}
	if v, ok := d.GetOk("endof_service_contract_renewal_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofServiceContractRenewalDateEpoch(x)
	}
	if v, ok := d.GetOk("endof_sw_maintenance_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofSwMaintenanceDate(x)
	}
	if v, ok := d.GetOk("endof_sw_maintenance_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofSwMaintenanceDateEpoch(x)
	}
	if v, ok := d.GetOk("headline"); ok {
		x := (v.(string))
		o.SetHeadline(x)
	}
	if v, ok := d.GetOk("last_dateof_support"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastDateofSupport(x)
	}
	if v, ok := d.GetOk("last_dateof_support_epoch"); ok {
		x := int64(v.(int))
		o.SetLastDateofSupportEpoch(x)
	}
	if v, ok := d.GetOk("last_ship_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastShipDate(x)
	}
	if v, ok := d.GetOk("last_ship_date_epoch"); ok {
		x := int64(v.(int))
		o.SetLastShipDateEpoch(x)
	}
	if v, ok := d.GetOk("migration_url"); ok {
		x := (v.(string))
		o.SetMigrationUrl(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("software_eol_url"); ok {
		x := (v.(string))
		o.SetSoftwareEolUrl(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.NiaapiApi.GetNiaapiApicSweolList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewNiaapiApicSweol()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
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
			if err := d.Set("class_id", (s.ClassId)); err != nil {
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
			if err := d.Set("migration_url", (s.MigrationUrl)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("software_eol_url", (s.SoftwareEolUrl)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
