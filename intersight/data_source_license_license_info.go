package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceLicenseLicenseInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLicenseLicenseInfoRead,
		Schema: map[string]*schema.Schema{
			"account_license_data": {
				Description: "A collection of references to the [license.AccountLicenseData](mo://license.AccountLicenseData) Managed Object.When this managed object is deleted, the referenced [license.AccountLicenseData](mo://license.AccountLicenseData) MO unsets its reference to this deleted MO.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed: true,
			},
			"active_admin": {
				Description: "The license administrative state.Set this property to 'true' to activate the license entitlements.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"days_left": {
				Description: "The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "The date and time when the trial period expires.The value of the 'endTime' property is set when the account enters the TrialPeriod or OutOfCompliance state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"enforce_mode": {
				Description: "The entitlement mode reported by Cisco Smart Software Manager.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"error_desc": {
				Description: "The detailed error message when there is any error related to this licensing entitlement.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"evaluation_period": {
				Description: "The default Trial or Grace period customer is entitled to.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"extra_evaluation": {
				Description: "The number of days the trial Trial or Grace period is extended.The trial or grace period can be extended once.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"license_count": {
				Description: "The total number of devices claimed in the Intersight account.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"license_state": {
				Description: "The license state defined by Intersight.The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_type": {
				Description: "The name of the Intersight license entitlement.For example, this property may be set to 'Essentials'.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"start_time": {
				Description: "The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"trial_admin": {
				Description: "The administrative state of the trial license.When the LicenseState is set to 'NotLicensed', 'trialAdmin' can be set to true to start the trial period,i.e. licenseState is set to be TrialPeriod.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
	}
}
func dataSourceLicenseLicenseInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "license/LicenseInfos"
	var o models.LicenseLicenseInfo
	if v, ok := d.GetOk("active_admin"); ok {
		x := (v.(bool))
		o.ActiveAdmin = &x
	}
	if v, ok := d.GetOk("days_left"); ok {
		x := int64(v.(int))
		o.DaysLeft = x
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndTime = x
	}
	if v, ok := d.GetOk("enforce_mode"); ok {
		x := (v.(string))
		o.EnforceMode = x
	}
	if v, ok := d.GetOk("error_desc"); ok {
		x := (v.(string))
		o.ErrorDesc = x
	}
	if v, ok := d.GetOk("evaluation_period"); ok {
		x := int64(v.(int))
		o.EvaluationPeriod = x
	}
	if v, ok := d.GetOk("extra_evaluation"); ok {
		x := int64(v.(int))
		o.ExtraEvaluation = x
	}
	if v, ok := d.GetOk("license_count"); ok {
		x := int64(v.(int))
		o.LicenseCount = x
	}
	if v, ok := d.GetOk("license_state"); ok {
		x := (v.(string))
		o.LicenseState = x
	}
	if v, ok := d.GetOk("license_type"); ok {
		x := (v.(string))
		o.LicenseType = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.StartTime = x
	}
	if v, ok := d.GetOk("trial_admin"); ok {
		x := (v.(bool))
		o.TrialAdmin = &x
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
			var s models.LicenseLicenseInfo
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("account_license_data", flattenMapLicenseAccountLicenseDataRef(s.AccountLicenseData, d)); err != nil {
				return err
			}
			if err := d.Set("active_admin", (s.ActiveAdmin)); err != nil {
				return err
			}
			if err := d.Set("days_left", (s.DaysLeft)); err != nil {
				return err
			}

			if err := d.Set("end_time", (s.EndTime).String()); err != nil {
				return err
			}
			if err := d.Set("enforce_mode", (s.EnforceMode)); err != nil {
				return err
			}
			if err := d.Set("error_desc", (s.ErrorDesc)); err != nil {
				return err
			}
			if err := d.Set("evaluation_period", (s.EvaluationPeriod)); err != nil {
				return err
			}
			if err := d.Set("extra_evaluation", (s.ExtraEvaluation)); err != nil {
				return err
			}
			if err := d.Set("license_count", (s.LicenseCount)); err != nil {
				return err
			}
			if err := d.Set("license_state", (s.LicenseState)); err != nil {
				return err
			}
			if err := d.Set("license_type", (s.LicenseType)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("start_time", (s.StartTime).String()); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("trial_admin", (s.TrialAdmin)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
