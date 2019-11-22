package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceNetworkconfigPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkconfigPolicyCreate,
		Read:   resourceNetworkconfigPolicyRead,
		Update: resourceNetworkconfigPolicyUpdate,
		Delete: resourceNetworkconfigPolicyDelete,
		Schema: map[string]*schema.Schema{
			"alternate_ipv4dns_server": {
				Description: "IP address of the secondary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"alternate_ipv6dns_server": {
				Description: "IP address of the secondary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dynamic_dns_domain": {
				Description: "The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_dynamic_dns": {
				Description: "If enabled, updates the resource records to the DNS from Cisco IMC.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv6": {
				Description: "If enabled, allows to configure IPv6 properties.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv4dns_from_dhcp": {
				Description: "If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv6dns_from_dhcp": {
				Description: "If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.",
				Type:        schema.TypeBool,
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
			"preferred_ipv4dns_server": {
				Description: "IP address of the primary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"preferred_ipv6dns_server": {
				Description: "IP address of the primary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
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
		},
	}
}
func resourceNetworkconfigPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.NetworkconfigPolicy
	if v, ok := d.GetOk("alternate_ipv4dns_server"); ok {
		x := (v.(string))
		o.AlternateIpv4dnsServer = x

	}

	if v, ok := d.GetOk("alternate_ipv6dns_server"); ok {
		x := (v.(string))
		o.AlternateIpv6dnsServer = x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("dynamic_dns_domain"); ok {
		x := (v.(string))
		o.DynamicDNSDomain = x

	}

	if v, ok := d.GetOkExists("enable_dynamic_dns"); ok {
		x := v.(bool)
		o.EnableDynamicDNS = &x
	}

	if v, ok := d.GetOkExists("enable_ipv6"); ok {
		x := v.(bool)
		o.EnableIPV6 = &x
	}

	if v, ok := d.GetOkExists("enable_ipv4dns_from_dhcp"); ok {
		x := v.(bool)
		o.EnableIpv4dnsFromDhcp = &x
	}

	if v, ok := d.GetOkExists("enable_ipv6dns_from_dhcp"); ok {
		x := v.(bool)
		o.EnableIpv6dnsFromDhcp = &x
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

	if v, ok := d.GetOk("preferred_ipv4dns_server"); ok {
		x := (v.(string))
		o.PreferredIpv4dnsServer = x

	}

	if v, ok := d.GetOk("preferred_ipv6dns_server"); ok {
		x := (v.(string))
		o.PreferredIpv6dnsServer = x

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

	url := "networkconfig/Policies"
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
	return resourceNetworkconfigPolicyRead(d, meta)
}
func detachNetworkconfigPolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "networkconfig/Policies" + "/" + d.Id()
	var o models.NetworkconfigPolicy

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

func resourceNetworkconfigPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "networkconfig/Policies" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.NetworkconfigPolicy
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("alternate_ipv4dns_server", (s.AlternateIpv4dnsServer)); err != nil {
		return err
	}

	if err := d.Set("alternate_ipv6dns_server", (s.AlternateIpv6dnsServer)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("dynamic_dns_domain", (s.DynamicDNSDomain)); err != nil {
		return err
	}

	if err := d.Set("enable_dynamic_dns", (s.EnableDynamicDNS)); err != nil {
		return err
	}

	if err := d.Set("enable_ipv6", (s.EnableIPV6)); err != nil {
		return err
	}

	if err := d.Set("enable_ipv4dns_from_dhcp", (s.EnableIpv4dnsFromDhcp)); err != nil {
		return err
	}

	if err := d.Set("enable_ipv6dns_from_dhcp", (s.EnableIpv6dnsFromDhcp)); err != nil {
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

	if err := d.Set("preferred_ipv4dns_server", (s.PreferredIpv4dnsServer)); err != nil {
		return err
	}

	if err := d.Set("preferred_ipv6dns_server", (s.PreferredIpv6dnsServer)); err != nil {
		return err
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRef(s.Profiles, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceNetworkconfigPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.NetworkconfigPolicy
	if d.HasChange("alternate_ipv4dns_server") {
		v := d.Get("alternate_ipv4dns_server")
		x := (v.(string))
		o.AlternateIpv4dnsServer = x
	}

	if d.HasChange("alternate_ipv6dns_server") {
		v := d.Get("alternate_ipv6dns_server")
		x := (v.(string))
		o.AlternateIpv6dnsServer = x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
	}

	if d.HasChange("dynamic_dns_domain") {
		v := d.Get("dynamic_dns_domain")
		x := (v.(string))
		o.DynamicDNSDomain = x
	}

	if d.HasChange("enable_dynamic_dns") {
		v := d.Get("enable_dynamic_dns")
		x := (v.(bool))
		o.EnableDynamicDNS = &x
	}

	if d.HasChange("enable_ipv6") {
		v := d.Get("enable_ipv6")
		x := (v.(bool))
		o.EnableIPV6 = &x
	}

	if d.HasChange("enable_ipv4dns_from_dhcp") {
		v := d.Get("enable_ipv4dns_from_dhcp")
		x := (v.(bool))
		o.EnableIpv4dnsFromDhcp = &x
	}

	if d.HasChange("enable_ipv6dns_from_dhcp") {
		v := d.Get("enable_ipv6dns_from_dhcp")
		x := (v.(bool))
		o.EnableIpv6dnsFromDhcp = &x
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

	if d.HasChange("preferred_ipv4dns_server") {
		v := d.Get("preferred_ipv4dns_server")
		x := (v.(string))
		o.PreferredIpv4dnsServer = x
	}

	if d.HasChange("preferred_ipv6dns_server") {
		v := d.Get("preferred_ipv6dns_server")
		x := (v.(string))
		o.PreferredIpv6dnsServer = x
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

	url := "networkconfig/Policies" + "/" + d.Id()
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
	return resourceNetworkconfigPolicyRead(d, meta)
}

func resourceNetworkconfigPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "networkconfig/Policies" + "/" + d.Id()
	detachNetworkconfigPolicyProfiles(d, meta)
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
