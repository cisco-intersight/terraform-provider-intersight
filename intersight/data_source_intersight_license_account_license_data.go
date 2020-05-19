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

func dataSourceLicenseAccountLicenseData() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLicenseAccountLicenseDataRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"customer_op": {
				Description: "A reference to a licenseCustomerOp resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed: true,
			},
			"default_license_type": {
				Description: "Default license tier set by user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"error_desc": {
				Description: "The detailed error message when there is any error related to license sync of this account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"group": {
				Description: "Account license data group name.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"highest_compliant_license_tier": {
				Description: "The highest license tier which is in compliant of this account.",
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
				Description: "An array of relationships to licenseLicenseInfo resources.",
				Type:        schema.TypeList,
				Optional:    true,
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
				Computed: true,
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
				Description: "A reference to a licenseSmartlicenseToken resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed: true,
			},
			"sync_status": {
				Description: "Current sync status for the account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
	var o = models.NewLicenseAccountLicenseData()
	if v, ok := d.GetOk("account_id"); ok {
		x := (v.(string))
		o.SetAccountId(x)
	}
	if v, ok := d.GetOk("agent_data"); ok {
		x := (v.(string))
		o.SetAgentData(x)
	}
	if v, ok := d.GetOk("auth_expire_time"); ok {
		x := (v.(string))
		o.SetAuthExpireTime(x)
	}
	if v, ok := d.GetOk("auth_initial_time"); ok {
		x := (v.(string))
		o.SetAuthInitialTime(x)
	}
	if v, ok := d.GetOk("auth_next_time"); ok {
		x := (v.(string))
		o.SetAuthNextTime(x)
	}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_license_type"); ok {
		x := (v.(string))
		o.SetDefaultLicenseType(x)
	}
	if v, ok := d.GetOk("error_desc"); ok {
		x := (v.(string))
		o.SetErrorDesc(x)
	}
	if v, ok := d.GetOk("group"); ok {
		x := (v.(string))
		o.SetGroup(x)
	}
	if v, ok := d.GetOk("highest_compliant_license_tier"); ok {
		x := (v.(string))
		o.SetHighestCompliantLicenseTier(x)
	}
	if v, ok := d.GetOk("last_sync"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastSync(x)
	}
	if v, ok := d.GetOk("last_updated_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastUpdatedTime(x)
	}
	if v, ok := d.GetOk("license_state"); ok {
		x := (v.(string))
		o.SetLicenseState(x)
	}
	if v, ok := d.GetOk("license_tech_support_info"); ok {
		x := (v.(string))
		o.SetLicenseTechSupportInfo(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("register_expire_time"); ok {
		x := (v.(string))
		o.SetRegisterExpireTime(x)
	}
	if v, ok := d.GetOk("register_initial_time"); ok {
		x := (v.(string))
		o.SetRegisterInitialTime(x)
	}
	if v, ok := d.GetOk("register_next_time"); ok {
		x := (v.(string))
		o.SetRegisterNextTime(x)
	}
	if v, ok := d.GetOk("registration_status"); ok {
		x := (v.(string))
		o.SetRegistrationStatus(x)
	}
	if v, ok := d.GetOk("renew_failure_string"); ok {
		x := (v.(string))
		o.SetRenewFailureString(x)
	}
	if v, ok := d.GetOk("smart_account"); ok {
		x := (v.(string))
		o.SetSmartAccount(x)
	}
	if v, ok := d.GetOk("sync_status"); ok {
		x := (v.(string))
		o.SetSyncStatus(x)
	}
	if v, ok := d.GetOk("virtual_account"); ok {
		x := (v.(string))
		o.SetVirtualAccount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.LicenseApi.GetLicenseAccountLicenseDataList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewLicenseAccountLicenseData()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRelationship(s.Account, d)); err != nil {
				return err
			}
			if err := d.Set("account_id", (s.AccountId)); err != nil {
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
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}

			if err := d.Set("customer_op", flattenMapLicenseCustomerOpRelationship(s.CustomerOp, d)); err != nil {
				return err
			}
			if err := d.Set("default_license_type", (s.DefaultLicenseType)); err != nil {
				return err
			}
			if err := d.Set("error_desc", (s.ErrorDesc)); err != nil {
				return err
			}
			if err := d.Set("group", (s.Group)); err != nil {
				return err
			}
			if err := d.Set("highest_compliant_license_tier", (s.HighestCompliantLicenseTier)); err != nil {
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

			if err := d.Set("licenseinfos", flattenListLicenseLicenseInfoRelationship(s.Licenseinfos, d)); err != nil {
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

			if err := d.Set("smartlicense_token", flattenMapLicenseSmartlicenseTokenRelationship(s.SmartlicenseToken, d)); err != nil {
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
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
