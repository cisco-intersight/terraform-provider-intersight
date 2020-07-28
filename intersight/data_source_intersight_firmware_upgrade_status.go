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
			"download_error": {
				Description: "The error message from the endpoint during the download.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_message": {
				Description: "The message from the endpoint during the download.}\ntype: string",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"download_percentage": {
				Description: "The percentage of the image downloaded in the endpoint.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_progress": {
				Description: "The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_retries": {
				Description: "The number of retries the plugin attempted before succeeding or failing the download.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_stage": {
				Description: "The image download stages. Example:downloading, flashing.",
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
			"upgrade": {
				Description: "A reference to a firmwareUpgradeBase resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"workflow": {
				Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
		},
	}
}

func dataSourceFirmwareUpgradeStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewFirmwareUpgradeStatusWithDefaults()
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
	if v, ok := d.GetOk("download_progress"); ok {
		x := int64(v.(int))
		o.SetDownloadProgress(x)
	}
	if v, ok := d.GetOk("download_retries"); ok {
		x := int64(v.(int))
		o.SetDownloadRetries(x)
	}
	if v, ok := d.GetOk("download_stage"); ok {
		x := (v.(string))
		o.SetDownloadStage(x)
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
	res, _, err := conn.ApiClient.FirmwareApi.GetFirmwareUpgradeStatusList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.FirmwareUpgradeStatusList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to FirmwareUpgradeStatus: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewFirmwareUpgradeStatusWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("download_error", (s.GetDownloadError())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadError: %+v", err)
			}
			if err := d.Set("download_percentage", (s.GetDownloadPercentage())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadPercentage: %+v", err)
			}
			if err := d.Set("download_progress", (s.GetDownloadProgress())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadProgress: %+v", err)
			}
			if err := d.Set("download_retries", (s.GetDownloadRetries())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadRetries: %+v", err)
			}
			if err := d.Set("download_stage", (s.GetDownloadStage())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadStage: %+v", err)
			}
			if err := d.Set("ep_power_status", (s.GetEpPowerStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property EpPowerStatus: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("overall_error", (s.GetOverallError())); err != nil {
				return fmt.Errorf("error occurred while setting property OverallError: %+v", err)
			}
			if err := d.Set("overall_percentage", (s.GetOverallPercentage())); err != nil {
				return fmt.Errorf("error occurred while setting property OverallPercentage: %+v", err)
			}
			if err := d.Set("overallstatus", (s.GetOverallstatus())); err != nil {
				return fmt.Errorf("error occurred while setting property Overallstatus: %+v", err)
			}
			if err := d.Set("pending_type", (s.GetPendingType())); err != nil {
				return fmt.Errorf("error occurred while setting property PendingType: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}

			if err := d.Set("upgrade", flattenMapFirmwareUpgradeBaseRelationship(s.GetUpgrade(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Upgrade: %+v", err)
			}

			if err := d.Set("workflow", flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflow(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Workflow: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
