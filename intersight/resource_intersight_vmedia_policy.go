package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVmediaPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmediaPolicyCreate,
		Read:   resourceVmediaPolicyRead,
		Update: resourceVmediaPolicyUpdate,
		Delete: resourceVmediaPolicyDelete,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "State of the Virtual Media service on the endpoint.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"encryption": {
				Description: "If enabled, allows encryption of all Virtual Media communications.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"low_power_usb": {
				Description: "If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"mappings": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_protocol": {
							Description: "Type of Authentication protocol when CIFS is used for communication with the remote server.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "none",
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"device_type": {
							Description: "Type of remote Virtual Media device.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "cdd",
						},
						"host_name": {
							Description: "IP address or hostname of the remote server.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"is_password_set": {
							Description: "Indicates whether the value of the 'password' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"mount_options": {
							Description: "Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\\n For NFS, supported options are ro, rw, nolock, noexec, soft, port=VALUE, timeo=VALUE, retry=VALUE.\\n For CIFS, supported options are soft, nounix, noserverino, guest.\\n For CIFS version < 3.0, vers=VALUE is mandatory. e.g. vers=2.0\\n For HTTP/HTTPS, the only supported option is noauto.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"mount_protocol": {
							Description: "Protocol to use to communicate with the remote server.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "nfs",
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"password": {
							Description: "Password associated with the username.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"remote_file": {
							Description: "Name of the remote file. It should be of .img format for HDD Virtual Media mapping and .iso format for CDD Virtual Media mapping.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"remote_path": {
							Description: "URL path to the location of the image on the remote server. The preferred format is '/path'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"username": {
							Description: "Username to log in to the remote server.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_name": {
							Description: "Identity of the image for Virtual Media mapping.",
							Type:        schema.TypeString,
							Optional:    true,
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
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
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
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
		},
	}
}

func resourceVmediaPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewVmediaPolicy()
	o.SetClassId("vmedia.Policy")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(bool))
		o.SetEnabled(x)
	}

	if v, ok := d.GetOk("encryption"); ok {
		x := (v.(bool))
		o.SetEncryption(x)
	}

	if v, ok := d.GetOk("low_power_usb"); ok {
		x := (v.(bool))
		o.SetLowPowerUsb(x)
	}

	if v, ok := d.GetOk("mappings"); ok {
		x := make([]models.VmediaMapping, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewVmediaMappingWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["authentication_protocol"]; ok {
				{
					x := (v.(string))
					o.SetAuthenticationProtocol(x)
				}
			}
			o.SetClassId("vmedia.Mapping")
			if v, ok := l["device_type"]; ok {
				{
					x := (v.(string))
					o.SetDeviceType(x)
				}
			}
			if v, ok := l["host_name"]; ok {
				{
					x := (v.(string))
					o.SetHostName(x)
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsPasswordSet(x)
				}
			}
			if v, ok := l["mount_options"]; ok {
				{
					x := (v.(string))
					o.SetMountOptions(x)
				}
			}
			if v, ok := l["mount_protocol"]; ok {
				{
					x := (v.(string))
					o.SetMountProtocol(x)
				}
			}
			o.SetObjectType("vmedia.Mapping")
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.SetPassword(x)
				}
			}
			if v, ok := l["remote_file"]; ok {
				{
					x := (v.(string))
					o.SetRemoteFile(x)
				}
			}
			if v, ok := l["remote_path"]; ok {
				{
					x := (v.(string))
					o.SetRemotePath(x)
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.SetUsername(x)
				}
			}
			if v, ok := l["volume_name"]; ok {
				{
					x := (v.(string))
					o.SetVolumeName(x)
				}
			}
			x = append(x, *o)
		}
		o.SetMappings(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("vmedia.Policy")

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("organization.Organization")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsOrganizationOrganizationRelationship())
		}
		x := p[0]
		o.SetOrganization(x)
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("policy.AbstractConfigProfile")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsPolicyAbstractConfigProfileRelationship())
		}
		o.SetProfiles(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	r := conn.ApiClient.VmediaApi.CreateVmediaPolicy(conn.ctx).VmediaPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceVmediaPolicyRead(d, meta)
}
func detachVmediaPolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewVmediaPolicy()

	o.Profiles = new([]models.PolicyAbstractConfigProfileRelationship)

	r := conn.ApiClient.VmediaApi.UpdateVmediaPolicy(conn.ctx, d.Id()).VmediaPolicy(*o)
	_, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while creating: %s", err.Error())
	}
	return err
}

func resourceVmediaPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.VmediaApi.GetVmediaPolicyByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("enabled", (s.Enabled)); err != nil {
		return err
	}

	if err := d.Set("encryption", (s.Encryption)); err != nil {
		return err
	}

	if err := d.Set("low_power_usb", (s.LowPowerUsb)); err != nil {
		return err
	}

	if err := d.Set("mappings", flattenListVmediaMapping(s.Mappings, d)); err != nil {
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

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceVmediaPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewVmediaPolicy()

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("enabled") {
		v := d.Get("enabled")
		x := (v.(bool))
		o.SetEnabled(x)
	}

	if d.HasChange("encryption") {
		v := d.Get("encryption")
		x := (v.(bool))
		o.SetEncryption(x)
	}

	if d.HasChange("low_power_usb") {
		v := d.Get("low_power_usb")
		x := (v.(bool))
		o.SetLowPowerUsb(x)
	}

	if d.HasChange("mappings") {
		v := d.Get("mappings")
		x := make([]models.VmediaMapping, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewVmediaMappingWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["authentication_protocol"]; ok {
				{
					x := (v.(string))
					o.SetAuthenticationProtocol(x)
				}
			}
			o.SetClassId("vmedia.Mapping")
			if v, ok := l["device_type"]; ok {
				{
					x := (v.(string))
					o.SetDeviceType(x)
				}
			}
			if v, ok := l["host_name"]; ok {
				{
					x := (v.(string))
					o.SetHostName(x)
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsPasswordSet(x)
				}
			}
			if v, ok := l["mount_options"]; ok {
				{
					x := (v.(string))
					o.SetMountOptions(x)
				}
			}
			if v, ok := l["mount_protocol"]; ok {
				{
					x := (v.(string))
					o.SetMountProtocol(x)
				}
			}
			o.SetObjectType("vmedia.Mapping")
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.SetPassword(x)
				}
			}
			if v, ok := l["remote_file"]; ok {
				{
					x := (v.(string))
					o.SetRemoteFile(x)
				}
			}
			if v, ok := l["remote_path"]; ok {
				{
					x := (v.(string))
					o.SetRemotePath(x)
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.SetUsername(x)
				}
			}
			if v, ok := l["volume_name"]; ok {
				{
					x := (v.(string))
					o.SetVolumeName(x)
				}
			}
			x = append(x, *o)
		}
		o.SetMappings(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("organization.Organization")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsOrganizationOrganizationRelationship())
		}
		x := p[0]
		o.SetOrganization(x)
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if d.HasChange("profiles") {
		v := d.Get("profiles")
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("policy.AbstractConfigProfile")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsPolicyAbstractConfigProfileRelationship())
		}
		o.SetProfiles(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	r := conn.ApiClient.VmediaApi.UpdateVmediaPolicy(conn.ctx, d.Id()).VmediaPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceVmediaPolicyRead(d, meta)
}

func resourceVmediaPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	e := detachVmediaPolicyProfiles(d, meta)
	if e != nil {
		return e
	}

	r := conn.ApiClient.VmediaApi.DeleteVmediaPolicy(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
