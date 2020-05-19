package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceFirmwareUpgradeStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirmwareUpgradeStatusRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"download_error": {
				Description: "The error message from the endpoint during the download.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_percentage": {
				Description: "The percentage of the image downloaded in the endpoint.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_stage": {
				Description: "The image download stages. Example:downloading, flashing.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_status": {
				Description: "The download status of the image in the endpoint.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ep_power_status": {
				Description: "The server power status after the upgrade request is submitted in the endpoint.",
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
			"overall_error": {
				Description: "The reason for the operation failure.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"overall_percentage": {
				Description: "The overall percentage of the operation.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"overallstatus": {
				Description: "The overall status of the operation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pending_type": {
				Description: "Pending reason for the upgrade waiting.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"upgrade": {
				Description: "A reference to a firmwareUpgrade resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		},
	}
}

func dataSourceFirmwareUpgradeStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewFirmwareUpgradeStatus()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("download_error"); ok {
		x := (v.(string))
		o.SetDownloadError(x)
	}
	if v, ok := d.GetOk("download_percentage"); ok {
		x := int64(v.(int))
		o.SetDownloadPercentage(x)
	}
	if v, ok := d.GetOk("download_stage"); ok {
		x := (v.(string))
		o.SetDownloadStage(x)
	}
	if v, ok := d.GetOk("download_status"); ok {
		x := (v.(string))
		o.SetDownloadStatus(x)
	}
	if v, ok := d.GetOk("ep_power_status"); ok {
		x := (v.(string))
		o.SetEpPowerStatus(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("overall_error"); ok {
		x := (v.(string))
		o.SetOverallError(x)
	}
	if v, ok := d.GetOk("overall_percentage"); ok {
		x := int64(v.(int))
		o.SetOverallPercentage(x)
	}
	if v, ok := d.GetOk("overallstatus"); ok {
		x := (v.(string))
		o.SetOverallstatus(x)
	}
	if v, ok := d.GetOk("pending_type"); ok {
		x := (v.(string))
		o.SetPendingType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.FirmwareApi.GetFirmwareUpgradeStatusList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewFirmwareUpgradeStatus()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}
			if err := d.Set("download_error", (s.DownloadError)); err != nil {
				return err
			}
			if err := d.Set("download_percentage", (s.DownloadPercentage)); err != nil {
				return err
			}
			if err := d.Set("download_stage", (s.DownloadStage)); err != nil {
				return err
			}
			if err := d.Set("download_status", (s.DownloadStatus)); err != nil {
				return err
			}
			if err := d.Set("ep_power_status", (s.EpPowerStatus)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("overall_error", (s.OverallError)); err != nil {
				return err
			}
			if err := d.Set("overall_percentage", (s.OverallPercentage)); err != nil {
				return err
			}
			if err := d.Set("overallstatus", (s.Overallstatus)); err != nil {
				return err
			}
			if err := d.Set("pending_type", (s.PendingType)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}

			if err := d.Set("upgrade", flattenMapFirmwareUpgradeRelationship(s.Upgrade, d)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
