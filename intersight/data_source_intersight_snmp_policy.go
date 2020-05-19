package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSnmpPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSnmpPolicyRead,
		Schema: map[string]*schema.Schema{
			"access_community_string": {
				Description: "The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"community_access": {
				Description: "Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"engine_id": {
				Description: "User-defined unique identification of the static engine.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
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
			"snmp_port": {
				Description: "Port on which Cisco IMC SNMP agent runs.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"snmp_traps": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"destination": {
							Description: "Address to which the SNMP trap information is sent.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"enabled": {
							Description: "Enables/disables the trap on the server If enabled, trap is active on the server.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Description: "Port used by the server to communicate with trap destination. Enter a value between 1-65535.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"type": {
							Description: "Type of trap which decides whether to receive a notification when a trap is received at the destination.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"user": {
							Description: "SNMP user for the trap. Applicable only to SNMPv3.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"version": {
							Description: "SNMP version used for the trap.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"snmp_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_password": {
							Description: "Authorization password for the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"auth_type": {
							Description: "Authorization protocol for authenticating the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"is_auth_password_set": {
							Description: "Indicates whether the value of the 'authPassword' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"is_privacy_password_set": {
							Description: "Indicates whether the value of the 'privacyPassword' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "SNMP username. Must have a minimum of 1 and and a maximum of 31 characters.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"privacy_password": {
							Description: "Privacy password for the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"privacy_type": {
							Description: "Privacy protocol for the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"security_level": {
							Description: "Security mechanism used for communication between agent and manager.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"sys_contact": {
				Description: "Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sys_location": {
				Description: "Location of host on which the SNMP agent (server) runs.",
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
			"trap_community": {
				Description: "SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSnmpPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewSnmpPolicy()
	if v, ok := d.GetOk("access_community_string"); ok {
		x := (v.(string))
		o.SetAccessCommunityString(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("community_access"); ok {
		x := (v.(string))
		o.SetCommunityAccess(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(bool))
		o.SetEnabled(x)
	}
	if v, ok := d.GetOk("engine_id"); ok {
		x := (v.(string))
		o.SetEngineId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("snmp_port"); ok {
		x := int64(v.(int))
		o.SetSnmpPort(x)
	}
	if v, ok := d.GetOk("sys_contact"); ok {
		x := (v.(string))
		o.SetSysContact(x)
	}
	if v, ok := d.GetOk("sys_location"); ok {
		x := (v.(string))
		o.SetSysLocation(x)
	}
	if v, ok := d.GetOk("trap_community"); ok {
		x := (v.(string))
		o.SetTrapCommunity(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.SnmpApi.GetSnmpPolicyList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewSnmpPolicy()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("access_community_string", (s.AccessCommunityString)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}
			if err := d.Set("community_access", (s.CommunityAccess)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("enabled", (s.Enabled)); err != nil {
				return err
			}
			if err := d.Set("engine_id", (s.EngineId)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.Organization, d)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
				return err
			}

			if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.Profiles, d)); err != nil {
				return err
			}
			if err := d.Set("snmp_port", (s.SnmpPort)); err != nil {
				return err
			}

			if err := d.Set("snmp_traps", flattenListSnmpTrap(s.SnmpTraps, d)); err != nil {
				return err
			}

			if err := d.Set("snmp_users", flattenListSnmpUser(s.SnmpUsers, d)); err != nil {
				return err
			}
			if err := d.Set("sys_contact", (s.SysContact)); err != nil {
				return err
			}
			if err := d.Set("sys_location", (s.SysLocation)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("trap_community", (s.TrapCommunity)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
