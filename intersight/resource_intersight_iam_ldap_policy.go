package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIamLdapPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceIamLdapPolicyCreate,
		Read:   resourceIamLdapPolicyRead,
		Update: resourceIamLdapPolicyUpdate,
		Delete: resourceIamLdapPolicyDelete,
		Schema: map[string]*schema.Schema{
			"appliance_account": {
				Description: "The appliance account to which the appliance LDAP policy belongs.",
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
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"base_properties": {
				Description: "Base settings of LDAP required while configuring LDAP policy.",
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
						"attribute": {
							Description: "Role and locale information of the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"base_dn": {
							Description: "Base Distinguished Name (DN). Starting point from where server will search for users and groups.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"bind_dn": {
							Description: "Distinguished Name (DN) of the user, that is used to authenticate against LDAP servers.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"bind_method": {
							Description: "Authentication method to access LDAP servers.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "LoginCredentials",
						},
						"domain": {
							Description: "The IPv4 domain that all users must be in.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"enable_encryption": {
							Description: "If enabled, the endpoint encrypts all information it sends to the LDAP server.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"enable_group_authorization": {
							Description: "If enabled, user authorization is also done at the group level for LDAP users not in the local user database.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"filter": {
							Description: "Criteria to identify entries in search requests.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"group_attribute": {
							Description: "Groups to which an LDAP entry belongs.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"is_password_set": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"nested_group_search_depth": {
							Description: "Search depth to look for a nested LDAP group in an LDAP group map.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Description: "Password of the user, that is used to authenticate.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"timeout": {
							Description: "LDAP authentication timeout duration, in seconds.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"dns_parameters": {
				Description: "Configuration settings to resolve LDAP servers, when DNS is enabled.",
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
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"search_domain": {
							Description: "Domain name that acts as a source for a DNS query.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"search_forest": {
							Description: "Forest name that acts as a source for a DNS query.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"source": {
							Description: "Source of the domain name used for the DNS SRV request.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Extracted",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_dns": {
				Description: "Enables DNS to access LDAP servers.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enabled": {
				Description: "LDAP server performs authentication.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"groups": {
				Description: "Relationship to collection of LDAP Groups.",
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
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr0_idp": {
				Description: "A collection of references to the [iam.Idp](mo://iam.Idp) Managed Object.\nWhen this managed object is deleted, the referenced [iam.Idp](mo://iam.Idp) MO unsets its reference to this deleted MO.",
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
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "The organization to which the LDAP policy belongs.",
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
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"profiles": {
				Description: "Relationship to the profile object.",
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
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"providers": {
				Description: "Relationship to collection of LDAP Providers.",
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
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"user_search_precedence": {
				Description: "Search precedence between local user database and LDAP user database.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "LocalUserDb",
			},
		},
	}
}
func resourceIamLdapPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.IamLdapPolicy
	if v, ok := d.GetOk("appliance_account"); ok {
		p := models.IamAccountRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamAccountRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.ApplianceAccount = &x

	}

	if v, ok := d.GetOk("base_properties"); ok {
		p := models.IamLdapBaseProperties{}
		if len(v.([]interface{})) > 0 {
			o := models.IamLdapBaseProperties{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["attribute"]; ok {
				{
					x := (v.(string))
					o.Attribute = x
				}
			}
			if v, ok := l["base_dn"]; ok {
				{
					x := (v.(string))
					o.BaseDn = x
				}
			}
			if v, ok := l["bind_dn"]; ok {
				{
					x := (v.(string))
					o.BindDn = x
				}
			}
			if v, ok := l["bind_method"]; ok {
				{
					x := (v.(string))
					o.BindMethod = &x
				}
			}
			if v, ok := l["domain"]; ok {
				{
					x := (v.(string))
					o.Domain = x
				}
			}
			if v, ok := l["enable_encryption"]; ok {
				{
					x := (v.(bool))
					o.EnableEncryption = &x
				}
			}
			if v, ok := l["enable_group_authorization"]; ok {
				{
					x := (v.(bool))
					o.EnableGroupAuthorization = &x
				}
			}
			if v, ok := l["filter"]; ok {
				{
					x := (v.(string))
					o.Filter = x
				}
			}
			if v, ok := l["group_attribute"]; ok {
				{
					x := (v.(string))
					o.GroupAttribute = x
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsPasswordSet = &x
				}
			}
			if v, ok := l["nested_group_search_depth"]; ok {
				{
					x := int64(v.(int))
					o.NestedGroupSearchDepth = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.Password = x
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.Timeout = x
				}
			}

			p = o
		}
		x := p
		o.BaseProperties = &x

	}

	if v, ok := d.GetOk("dns_parameters"); ok {
		p := models.IamLdapDNSParameters{}
		if len(v.([]interface{})) > 0 {
			o := models.IamLdapDNSParameters{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["search_domain"]; ok {
				{
					x := (v.(string))
					o.SearchDomain = x
				}
			}
			if v, ok := l["search_forest"]; ok {
				{
					x := (v.(string))
					o.SearchForest = x
				}
			}
			if v, ok := l["source"]; ok {
				{
					x := (v.(string))
					o.Source = &x
				}
			}

			p = o
		}
		x := p
		o.DNSParameters = &x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOkExists("enable_dns"); ok {
		x := v.(bool)
		o.EnableDNS = &x
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		x := v.(bool)
		o.Enabled = &x
	}

	if v, ok := d.GetOk("groups"); ok {
		x := make([]*models.IamLdapGroupRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.IamLdapGroupRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Groups = x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x

	}

	if v, ok := d.GetOk("nr0_idp"); ok {
		p := models.IamIdpRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamIdpRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Nr0Idp = &x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

	}

	if v, ok := d.GetOk("organization"); ok {
		p := models.OrganizationOrganizationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OrganizationOrganizationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Organization = &x

	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]*models.MoBaseMoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoBaseMoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.PermissionResources = x

	}

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]*models.PolicyAbstractConfigProfileRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.PolicyAbstractConfigProfileRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Profiles = x

	}

	if v, ok := d.GetOk("providers"); ok {
		x := make([]*models.IamLdapProviderRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.IamLdapProviderRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Providers = x

	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO1P1.MoTagAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x

	}

	if v, ok := d.GetOk("user_search_precedence"); ok {
		x := (v.(string))
		o.UserSearchPrecedence = &x

	}

	url := "iam/LdapPolicies"
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceIamLdapPolicyRead(d, meta)
}
func detachIamLdapPolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "iam/LdapPolicies" + "/" + d.Id()
	var o models.IamLdapPolicy

	o.Profiles = []*models.PolicyAbstractConfigProfileRef{}

	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling sol_policy. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendUpdateRequest(url, data)
	if err != nil {
		log.Printf("error in sending update request. error %s", err.Error())
		return err
	}
	var s models.ServerProfile
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling sol_policy. Error: %s", err.Error())
	}

	return err
}

func resourceIamLdapPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "iam/LdapPolicies" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.IamLdapPolicy
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("appliance_account", flattenMapIamAccountRef(s.ApplianceAccount, d)); err != nil {
		return err
	}

	if err := d.Set("base_properties", flattenMapIamLdapBaseProperties(s.BaseProperties, d)); err != nil {
		return err
	}

	if err := d.Set("dns_parameters", flattenMapIamLdapDNSParameters(s.DNSParameters, d)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("enable_dns", (s.EnableDNS)); err != nil {
		return err
	}

	if err := d.Set("enabled", (s.Enabled)); err != nil {
		return err
	}

	if err := d.Set("groups", flattenListIamLdapGroupRef(s.Groups, d)); err != nil {
		return err
	}

	if err := d.Set("moid", (s.Moid)); err != nil {
		return err
	}

	if err := d.Set("name", (s.Name)); err != nil {
		return err
	}

	if err := d.Set("nr0_idp", flattenMapIamIdpRef(s.Nr0Idp, d)); err != nil {
		return err
	}

	if err := d.Set("object_type", (s.ObjectType)); err != nil {
		return err
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRef(s.Profiles, d)); err != nil {
		return err
	}

	if err := d.Set("providers", flattenListIamLdapProviderRef(s.Providers, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("user_search_precedence", (s.UserSearchPrecedence)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceIamLdapPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.IamLdapPolicy
	if d.HasChange("appliance_account") {
		v := d.Get("appliance_account")
		p := models.IamAccountRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamAccountRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.ApplianceAccount = &x
	}

	if d.HasChange("base_properties") {
		v := d.Get("base_properties")
		p := models.IamLdapBaseProperties{}
		if len(v.([]interface{})) > 0 {
			o := models.IamLdapBaseProperties{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["attribute"]; ok {
				{
					x := (v.(string))
					o.Attribute = x
				}
			}
			if v, ok := l["base_dn"]; ok {
				{
					x := (v.(string))
					o.BaseDn = x
				}
			}
			if v, ok := l["bind_dn"]; ok {
				{
					x := (v.(string))
					o.BindDn = x
				}
			}
			if v, ok := l["bind_method"]; ok {
				{
					x := (v.(string))
					o.BindMethod = &x
				}
			}
			if v, ok := l["domain"]; ok {
				{
					x := (v.(string))
					o.Domain = x
				}
			}
			if v, ok := l["enable_encryption"]; ok {
				{
					x := (v.(bool))
					o.EnableEncryption = &x
				}
			}
			if v, ok := l["enable_group_authorization"]; ok {
				{
					x := (v.(bool))
					o.EnableGroupAuthorization = &x
				}
			}
			if v, ok := l["filter"]; ok {
				{
					x := (v.(string))
					o.Filter = x
				}
			}
			if v, ok := l["group_attribute"]; ok {
				{
					x := (v.(string))
					o.GroupAttribute = x
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsPasswordSet = &x
				}
			}
			if v, ok := l["nested_group_search_depth"]; ok {
				{
					x := int64(v.(int))
					o.NestedGroupSearchDepth = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.Password = x
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.Timeout = x
				}
			}

			p = o
		}
		x := p
		o.BaseProperties = &x
	}

	if d.HasChange("dns_parameters") {
		v := d.Get("dns_parameters")
		p := models.IamLdapDNSParameters{}
		if len(v.([]interface{})) > 0 {
			o := models.IamLdapDNSParameters{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["search_domain"]; ok {
				{
					x := (v.(string))
					o.SearchDomain = x
				}
			}
			if v, ok := l["search_forest"]; ok {
				{
					x := (v.(string))
					o.SearchForest = x
				}
			}
			if v, ok := l["source"]; ok {
				{
					x := (v.(string))
					o.Source = &x
				}
			}

			p = o
		}
		x := p
		o.DNSParameters = &x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
	}

	if d.HasChange("enable_dns") {
		v := d.Get("enable_dns")
		x := (v.(bool))
		o.EnableDNS = &x
	}

	if d.HasChange("enabled") {
		v := d.Get("enabled")
		x := (v.(bool))
		o.Enabled = &x
	}

	if d.HasChange("groups") {
		v := d.Get("groups")
		x := make([]*models.IamLdapGroupRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.IamLdapGroupRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Groups = x
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.Moid = x
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.Name = x
	}

	if d.HasChange("nr0_idp") {
		v := d.Get("nr0_idp")
		p := models.IamIdpRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamIdpRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Nr0Idp = &x
	}

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
	}

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := models.OrganizationOrganizationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OrganizationOrganizationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		o.Organization = &x
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]*models.MoBaseMoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoBaseMoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.PermissionResources = x
	}

	if d.HasChange("profiles") {
		v := d.Get("profiles")
		x := make([]*models.PolicyAbstractConfigProfileRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.PolicyAbstractConfigProfileRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Profiles = x
	}

	if d.HasChange("providers") {
		v := d.Get("providers")
		x := make([]*models.IamLdapProviderRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.IamLdapProviderRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Providers = x
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO1P1.MoTagAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x
	}

	if d.HasChange("user_search_precedence") {
		v := d.Get("user_search_precedence")
		x := (v.(string))
		o.UserSearchPrecedence = &x
	}

	url := "iam/LdapPolicies" + "/" + d.Id()
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendUpdateRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceIamLdapPolicyRead(d, meta)
}

func resourceIamLdapPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "iam/LdapPolicies" + "/" + d.Id()
	detachIamLdapPolicyProfiles(d, meta)
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
