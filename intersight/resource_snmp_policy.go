package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSnmpPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpPolicyCreate,
		Read:   resourceSnmpPolicyRead,
		Update: resourceSnmpPolicyUpdate,
		Delete: resourceSnmpPolicyDelete,
		Schema: map[string]*schema.Schema{
			"access_community_string": {
				Description: "The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"community_access": {
				Description: "Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Disabled",
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
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "Relationship to the Organization that owns the Managed Object.",
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
				Computed:   true,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"snmp_port": {
				Description: "Port on which Cisco IMC SNMP agent runs.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"snmp_traps": {
				Description: "List of SNMP traps for the policy.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Default:     "Trap",
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
							Default:     "V3",
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"snmp_users": {
				Description: "List of SNMP users for the policy.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"auth_password": {
							Description: "Authorization password for the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"auth_type": {
							Description: "Authorization protocol for authenticating the user.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "NA",
						},
						"is_auth_password_set": {
							Description: "Indicates whether the value of the 'authPassword' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"is_privacy_password_set": {
							Description: "",
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Default:     "NA",
						},
						"security_level": {
							Description: "Security mechanism used for communication between agent and manager.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "AuthPriv",
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"trap_community": {
				Description: "SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func resourceSnmpPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.SnmpPolicy
	if v, ok := d.GetOk("access_community_string"); ok {
		x := (v.(string))
		o.AccessCommunityString = x

	}

	if v, ok := d.GetOk("community_access"); ok {
		x := (v.(string))
		o.CommunityAccess = &x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOkExists("enabled"); ok {
		x := v.(bool)
		o.Enabled = &x
	}

	if v, ok := d.GetOk("engine_id"); ok {
		x := (v.(string))
		o.EngineID = x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

	}

	if v, ok := d.GetOk("organization"); ok {
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
		o.Organization = &x

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

	if v, ok := d.GetOk("snmp_port"); ok {
		x := int64(v.(int))
		o.SnmpPort = x

	}

	if v, ok := d.GetOk("snmp_traps"); ok {
		x := make([]*models.SnmpTrap, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.SnmpTrap{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.SnmpTrapAO0P0.SnmpTrapAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["destination"]; ok {
					{
						x := (v.(string))
						o.Destination = x
					}
				}
				if v, ok := l["enabled"]; ok {
					{
						x := (v.(bool))
						o.Enabled = &x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["port"]; ok {
					{
						x := int64(v.(int))
						o.Port = x
					}
				}
				if v, ok := l["type"]; ok {
					{
						x := (v.(string))
						o.Type = &x
					}
				}
				if v, ok := l["user"]; ok {
					{
						x := (v.(string))
						o.User = x
					}
				}
				if v, ok := l["version"]; ok {
					{
						x := (v.(string))
						o.Version = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.SnmpTraps = x

	}

	if v, ok := d.GetOk("snmp_users"); ok {
		x := make([]*models.SnmpUser, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.SnmpUser{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.SnmpUserAO0P0.SnmpUserAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["auth_password"]; ok {
					{
						x := (v.(string))
						o.AuthPassword = x
					}
				}
				if v, ok := l["auth_type"]; ok {
					{
						x := (v.(string))
						o.AuthType = &x
					}
				}
				if v, ok := l["is_auth_password_set"]; ok {
					{
						x := (v.(bool))
						o.IsAuthPasswordSet = &x
					}
				}
				if v, ok := l["is_privacy_password_set"]; ok {
					{
						x := (v.(bool))
						o.IsPrivacyPasswordSet = &x
					}
				}
				if v, ok := l["name"]; ok {
					{
						x := (v.(string))
						o.Name = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["privacy_password"]; ok {
					{
						x := (v.(string))
						o.PrivacyPassword = x
					}
				}
				if v, ok := l["privacy_type"]; ok {
					{
						x := (v.(string))
						o.PrivacyType = &x
					}
				}
				if v, ok := l["security_level"]; ok {
					{
						x := (v.(string))
						o.SecurityLevel = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.SnmpUsers = x

	}

	if v, ok := d.GetOk("sys_contact"); ok {
		x := (v.(string))
		o.SysContact = x

	}

	if v, ok := d.GetOk("sys_location"); ok {
		x := (v.(string))
		o.SysLocation = x

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
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
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

	if v, ok := d.GetOk("trap_community"); ok {
		x := (v.(string))
		o.TrapCommunity = x

	}

	url := "snmp/Policies"
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
	return resourceSnmpPolicyRead(d, meta)
}
func detachSnmpPolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "snmp/Policies" + "/" + d.Id()
	var o models.SnmpPolicy

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

func resourceSnmpPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "snmp/Policies" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.SnmpPolicy
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("access_community_string", (s.AccessCommunityString)); err != nil {
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

	if err := d.Set("engine_id", (s.EngineID)); err != nil {
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

	if err := d.Set("organization", flattenMapIamAccountRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRef(s.Profiles, d)); err != nil {
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

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceSnmpPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.SnmpPolicy
	if d.HasChange("access_community_string") {
		v := d.Get("access_community_string")
		x := (v.(string))
		o.AccessCommunityString = x
	}

	if d.HasChange("community_access") {
		v := d.Get("community_access")
		x := (v.(string))
		o.CommunityAccess = &x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
	}

	if d.HasChange("enabled") {
		v := d.Get("enabled")
		x := (v.(bool))
		o.Enabled = &x
	}

	if d.HasChange("engine_id") {
		v := d.Get("engine_id")
		x := (v.(string))
		o.EngineID = x
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

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
	}

	if d.HasChange("organization") {
		v := d.Get("organization")
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
		o.Organization = &x
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

	if d.HasChange("snmp_port") {
		v := d.Get("snmp_port")
		x := int64(v.(int))
		o.SnmpPort = x
	}

	if d.HasChange("snmp_traps") {
		v := d.Get("snmp_traps")
		x := make([]*models.SnmpTrap, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.SnmpTrap{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.SnmpTrapAO0P0.SnmpTrapAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["destination"]; ok {
					{
						x := (v.(string))
						o.Destination = x
					}
				}
				if v, ok := l["enabled"]; ok {
					{
						x := (v.(bool))
						o.Enabled = &x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["port"]; ok {
					{
						x := int64(v.(int))
						o.Port = x
					}
				}
				if v, ok := l["type"]; ok {
					{
						x := (v.(string))
						o.Type = &x
					}
				}
				if v, ok := l["user"]; ok {
					{
						x := (v.(string))
						o.User = x
					}
				}
				if v, ok := l["version"]; ok {
					{
						x := (v.(string))
						o.Version = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.SnmpTraps = x
	}

	if d.HasChange("snmp_users") {
		v := d.Get("snmp_users")
		x := make([]*models.SnmpUser, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.SnmpUser{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.SnmpUserAO0P0.SnmpUserAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["auth_password"]; ok {
					{
						x := (v.(string))
						o.AuthPassword = x
					}
				}
				if v, ok := l["auth_type"]; ok {
					{
						x := (v.(string))
						o.AuthType = &x
					}
				}
				if v, ok := l["is_auth_password_set"]; ok {
					{
						x := (v.(bool))
						o.IsAuthPasswordSet = &x
					}
				}
				if v, ok := l["is_privacy_password_set"]; ok {
					{
						x := (v.(bool))
						o.IsPrivacyPasswordSet = &x
					}
				}
				if v, ok := l["name"]; ok {
					{
						x := (v.(string))
						o.Name = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["privacy_password"]; ok {
					{
						x := (v.(string))
						o.PrivacyPassword = x
					}
				}
				if v, ok := l["privacy_type"]; ok {
					{
						x := (v.(string))
						o.PrivacyType = &x
					}
				}
				if v, ok := l["security_level"]; ok {
					{
						x := (v.(string))
						o.SecurityLevel = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.SnmpUsers = x
	}

	if d.HasChange("sys_contact") {
		v := d.Get("sys_contact")
		x := (v.(string))
		o.SysContact = x
	}

	if d.HasChange("sys_location") {
		v := d.Get("sys_location")
		x := (v.(string))
		o.SysLocation = x
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
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
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

	if d.HasChange("trap_community") {
		v := d.Get("trap_community")
		x := (v.(string))
		o.TrapCommunity = x
	}

	url := "snmp/Policies" + "/" + d.Id()
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
	return resourceSnmpPolicyRead(d, meta)
}

func resourceSnmpPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "snmp/Policies" + "/" + d.Id()
	detachSnmpPolicyProfiles(d, meta)
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
