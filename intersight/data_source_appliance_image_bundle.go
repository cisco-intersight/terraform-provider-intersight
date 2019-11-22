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

func dataSourceApplianceImageBundle() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceImageBundleRead,
		Schema: map[string]*schema.Schema{
			"ansible_packages": {
				Description: "Collection of the Intersight Appliance's system installation packages.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"auto_upgrade": {
				Description: "Indicates that the software upgrade was automatically initiated by the Intersight Appliance.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"dc_packages": {
				Description: "Collection of the Intersight Appliance's device connector packages.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"debug_packages": {
				Description: "Collection of the Intersight Appliance's developer debug packages. Optional, and not installed by default.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"description": {
				Description: "Short description of the software upgrade bundle.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"endpoint_packages": {
				Description: "Collection of the Intersight Appliance's endpoint UI packages such as Cisco UCSM, Cisco HyperFlex etc.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fingerprint": {
				Description: "Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"has_error": {
				Description: "Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"infra_packages": {
				Description: "Collection of the Intersight Appliance's infrastructure service packages such as database.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"init_packages": {
				Description: "Collection of the Intersight Appliance's initialization service packages.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
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
			"name": {
				Description: "Name of the software upgrade bundle.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"notes": {
				Description: "Detailed description of the software upgrade bundle.",
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
			"priority": {
				Description: "Software upgrade manifest's upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"release_time": {
				Description: "Software upgrade manifest's release date and time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_packages": {
				Description: "Collection of the Intersight Appliance's micro-services pakages.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status_message": {
				Description: "Status message set during the manifest processing.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"system_packages": {
				Description: "Collection of the Intersight Appliance's system packages such as DNS etc.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
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
			"ui_packages": {
				Description: "Collection of the Intersight Appliance's UI packages of the micro-services.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
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
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"upgrade_end_time": {
				Description: "End date of the software upgrade process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"upgrade_grace_period": {
				Description: "Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"upgrade_start_time": {
				Description: "Start date of the software upgrade process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"version": {
				Description: "Software upgrade manifest's version.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceApplianceImageBundleRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "appliance/ImageBundles"
	var o models.ApplianceImageBundle
	if v, ok := d.GetOk("auto_upgrade"); ok {
		x := (v.(bool))
		o.AutoUpgrade = &x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("fingerprint"); ok {
		x := (v.(string))
		o.Fingerprint = x
	}
	if v, ok := d.GetOk("has_error"); ok {
		x := (v.(bool))
		o.HasError = &x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x
	}
	if v, ok := d.GetOk("notes"); ok {
		x := (v.(string))
		o.Notes = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("priority"); ok {
		x := (v.(string))
		o.Priority = x
	}
	if v, ok := d.GetOk("release_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ReleaseTime = x
	}
	if v, ok := d.GetOk("status_message"); ok {
		x := (v.(string))
		o.StatusMessage = x
	}
	if v, ok := d.GetOk("upgrade_end_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.UpgradeEndTime = x
	}
	if v, ok := d.GetOk("upgrade_grace_period"); ok {
		x := int64(v.(int))
		o.UpgradeGracePeriod = x
	}
	if v, ok := d.GetOk("upgrade_start_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.UpgradeStartTime = x
	}
	if v, ok := d.GetOk("version"); ok {
		x := (v.(string))
		o.Version = x
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
			var s models.ApplianceImageBundle
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("ansible_packages", flattenListOnpremImagePackage(s.AnsiblePackages, d)); err != nil {
				return err
			}
			if err := d.Set("auto_upgrade", (s.AutoUpgrade)); err != nil {
				return err
			}

			if err := d.Set("dc_packages", flattenListOnpremImagePackage(s.DcPackages, d)); err != nil {
				return err
			}

			if err := d.Set("debug_packages", flattenListOnpremImagePackage(s.DebugPackages, d)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}

			if err := d.Set("endpoint_packages", flattenListOnpremImagePackage(s.EndpointPackages, d)); err != nil {
				return err
			}
			if err := d.Set("fingerprint", (s.Fingerprint)); err != nil {
				return err
			}
			if err := d.Set("has_error", (s.HasError)); err != nil {
				return err
			}

			if err := d.Set("infra_packages", flattenListOnpremImagePackage(s.InfraPackages, d)); err != nil {
				return err
			}

			if err := d.Set("init_packages", flattenListOnpremImagePackage(s.InitPackages, d)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("notes", (s.Notes)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("priority", (s.Priority)); err != nil {
				return err
			}

			if err := d.Set("release_time", (s.ReleaseTime).String()); err != nil {
				return err
			}

			if err := d.Set("service_packages", flattenListOnpremImagePackage(s.ServicePackages, d)); err != nil {
				return err
			}
			if err := d.Set("status_message", (s.StatusMessage)); err != nil {
				return err
			}

			if err := d.Set("system_packages", flattenListOnpremImagePackage(s.SystemPackages, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}

			if err := d.Set("ui_packages", flattenListOnpremImagePackage(s.UIPackages, d)); err != nil {
				return err
			}

			if err := d.Set("upgrade_end_time", (s.UpgradeEndTime).String()); err != nil {
				return err
			}
			if err := d.Set("upgrade_grace_period", (s.UpgradeGracePeriod)); err != nil {
				return err
			}

			if err := d.Set("upgrade_start_time", (s.UpgradeStartTime).String()); err != nil {
				return err
			}
			if err := d.Set("version", (s.Version)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
