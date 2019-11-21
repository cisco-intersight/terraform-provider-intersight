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

func dataSourceLicenseAccountLicenseData() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLicenseAccountLicenseDataRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "AccountLicenseData record to Account record relationship.",
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
			"account_id": {
				Description: "Root user's ID of the account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"agent_data": {
				Description: "Agent trusted store data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"auth_expire_time": {
				Description: "Authorization expiration time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"auth_initial_time": {
				Description: "Intial authorization time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"auth_next_time": {
				Description: "Next time for the authorization.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Description: "Account license data category name.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"customer_op": {
				Description: "AccountLicenseData record to CustomerOp record relationship.",
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
			"group": {
				Description: "Account license data group name.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_sync": {
				Description: "Specifies last sync time with SA.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_updated_time": {
				Description: "Record's last update datetime.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_state": {
				Description: "Aggregrated mode for the agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_tech_support_info": {
				Description: "Tech-support info of a smart-agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"licenseinfos": {
				Description: "All LicenceInfo records refercing this AccountLicenseData record.",
				Type:        schema.TypeList,
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
			"register_expire_time": {
				Description: "Registration exipiration time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"register_initial_time": {
				Description: "Initial time of registration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"register_next_time": {
				Description: "Next time for the license registration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registration_status": {
				Description: "Registration status of a smart-agent.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"renew_failure_string": {
				Description: "License renewal failure message.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"smart_account": {
				Description: "Name of the smart account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"smartlicense_token": {
				Description: "AccountLicenseData record to SmartlicenseToken record relationship.",
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
			"sync_status": {
				Description: "Current sync status for the account.",
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
			"virtual_account": {
				Description: "Name of the virtual account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceLicenseAccountLicenseDataRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "license/AccountLicenseData"
	var o models.LicenseAccountLicenseData
	if v, ok := d.GetOk("account_id"); ok {
		x := (v.(string))
		o.AccountID = x
	}
	if v, ok := d.GetOk("agent_data"); ok {
		x := (v.(string))
		o.AgentData = x
	}
	if v, ok := d.GetOk("auth_expire_time"); ok {
		x := (v.(string))
		o.AuthExpireTime = x
	}
	if v, ok := d.GetOk("auth_initial_time"); ok {
		x := (v.(string))
		o.AuthInitialTime = x
	}
	if v, ok := d.GetOk("auth_next_time"); ok {
		x := (v.(string))
		o.AuthNextTime = x
	}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.Category = x
	}
	if v, ok := d.GetOk("group"); ok {
		x := (v.(string))
		o.Group = x
	}
	if v, ok := d.GetOk("last_sync"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.LastSync = x
	}
	if v, ok := d.GetOk("last_updated_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.LastUpdatedTime = x
	}
	if v, ok := d.GetOk("license_state"); ok {
		x := (v.(string))
		o.LicenseState = x
	}
	if v, ok := d.GetOk("license_tech_support_info"); ok {
		x := (v.(string))
		o.LicenseTechSupportInfo = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("register_expire_time"); ok {
		x := (v.(string))
		o.RegisterExpireTime = x
	}
	if v, ok := d.GetOk("register_initial_time"); ok {
		x := (v.(string))
		o.RegisterInitialTime = x
	}
	if v, ok := d.GetOk("register_next_time"); ok {
		x := (v.(string))
		o.RegisterNextTime = x
	}
	if v, ok := d.GetOk("registration_status"); ok {
		x := (v.(string))
		o.RegistrationStatus = x
	}
	if v, ok := d.GetOk("renew_failure_string"); ok {
		x := (v.(string))
		o.RenewFailureString = x
	}
	if v, ok := d.GetOk("smart_account"); ok {
		x := (v.(string))
		o.SmartAccount = x
	}
	if v, ok := d.GetOk("sync_status"); ok {
		x := (v.(string))
		o.SyncStatus = x
	}
	if v, ok := d.GetOk("virtual_account"); ok {
		x := (v.(string))
		o.VirtualAccount = x
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
			var s models.LicenseAccountLicenseData
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
				return err
			}
			if err := d.Set("account_id", (s.AccountID)); err != nil {
				return err
			}
			if err := d.Set("agent_data", (s.AgentData)); err != nil {
				return err
			}
			if err := d.Set("auth_expire_time", (s.AuthExpireTime)); err != nil {
				return err
			}
			if err := d.Set("auth_initial_time", (s.AuthInitialTime)); err != nil {
				return err
			}
			if err := d.Set("auth_next_time", (s.AuthNextTime)); err != nil {
				return err
			}
			if err := d.Set("category", (s.Category)); err != nil {
				return err
			}

			if err := d.Set("customer_op", flattenMapLicenseCustomerOpRef(s.CustomerOp, d)); err != nil {
				return err
			}
			if err := d.Set("group", (s.Group)); err != nil {
				return err
			}

			if err := d.Set("last_sync", (s.LastSync).String()); err != nil {
				return err
			}

			if err := d.Set("last_updated_time", (s.LastUpdatedTime).String()); err != nil {
				return err
			}
			if err := d.Set("license_state", (s.LicenseState)); err != nil {
				return err
			}
			if err := d.Set("license_tech_support_info", (s.LicenseTechSupportInfo)); err != nil {
				return err
			}

			if err := d.Set("licenseinfos", flattenListLicenseLicenseInfoRef(s.Licenseinfos, d)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("register_expire_time", (s.RegisterExpireTime)); err != nil {
				return err
			}
			if err := d.Set("register_initial_time", (s.RegisterInitialTime)); err != nil {
				return err
			}
			if err := d.Set("register_next_time", (s.RegisterNextTime)); err != nil {
				return err
			}
			if err := d.Set("registration_status", (s.RegistrationStatus)); err != nil {
				return err
			}
			if err := d.Set("renew_failure_string", (s.RenewFailureString)); err != nil {
				return err
			}
			if err := d.Set("smart_account", (s.SmartAccount)); err != nil {
				return err
			}

			if err := d.Set("smartlicense_token", flattenMapLicenseSmartlicenseTokenRef(s.SmartlicenseToken, d)); err != nil {
				return err
			}
			if err := d.Set("sync_status", (s.SyncStatus)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("virtual_account", (s.VirtualAccount)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
