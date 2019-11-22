package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceFirmwareUpgradeStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirmwareUpgradeStatusRead,
		Schema: map[string]*schema.Schema{
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
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"upgrade": {
				Description: "A collection of references to the [firmware.Upgrade](mo://firmware.Upgrade) Managed Object.When this managed object is deleted, the referenced [firmware.Upgrade](mo://firmware.Upgrade) MO unsets its reference to this deleted MO.",
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
		},
	}
}
func dataSourceFirmwareUpgradeStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "firmware/UpgradeStatuses"
	var o models.FirmwareUpgradeStatus
	if v, ok := d.GetOk("download_error"); ok {
		x := (v.(string))
		o.DownloadError = x
	}
	if v, ok := d.GetOk("download_percentage"); ok {
		x := int64(v.(int))
		o.DownloadPercentage = x
	}
	if v, ok := d.GetOk("download_stage"); ok {
		x := (v.(string))
		o.DownloadStage = x
	}
	if v, ok := d.GetOk("download_status"); ok {
		x := (v.(string))
		o.DownloadStatus = x
	}
	if v, ok := d.GetOk("ep_power_status"); ok {
		x := (v.(string))
		o.EpPowerStatus = &x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("overall_error"); ok {
		x := (v.(string))
		o.OverallError = x
	}
	if v, ok := d.GetOk("overall_percentage"); ok {
		x := int64(v.(int))
		o.OverallPercentage = x
	}
	if v, ok := d.GetOk("overallstatus"); ok {
		x := (v.(string))
		o.Overallstatus = &x
	}
	if v, ok := d.GetOk("pending_type"); ok {
		x := (v.(string))
		o.PendingType = &x
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
			var s models.FirmwareUpgradeStatus
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
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

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}

			if err := d.Set("upgrade", flattenMapFirmwareUpgradeRef(s.Upgrade, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
