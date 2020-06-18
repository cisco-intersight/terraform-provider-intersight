package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSmtpPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSmtpPolicyCreate,
		Read:   resourceSmtpPolicyRead,
		Update: resourceSmtpPolicyUpdate,
		Delete: resourceSmtpPolicyDelete,
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
				Description: "If enabled, controls the state of the SMTP client service on the managed device.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"min_severity": {
				Description: "Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "critical",
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
							Description: "The Object Type of the referenced REST resource.",
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
			"smtp_port": {
				Description: "Port number used by the SMTP server for outgoing SMTP communication.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"smtp_recipients": {
				Description: "List of email addresses that will receive notifications for faults.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"smtp_server": {
				Description: "IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sender_email": {
				Description: "The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field.",
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
		},
	}
}
func resourceSmtpPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.SMTPPolicy
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOkExists("enabled"); ok {
		x := v.(bool)
		o.Enabled = &x
	}

	if v, ok := d.GetOk("min_severity"); ok {
		x := (v.(string))
		o.MinSeverity = &x

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
		if len(v.([]interface{})) > 0 {
			o.Organization = &x
		}

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

	if v, ok := d.GetOk("smtp_port"); ok {
		x := int64(v.(int))
		o.SMTPPort = x

	}

	if v, ok := d.GetOk("smtp_recipients"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.SMTPRecipients = x

	}

	if v, ok := d.GetOk("smtp_server"); ok {
		x := (v.(string))
		o.SMTPServer = x

	}

	if v, ok := d.GetOk("sender_email"); ok {
		x := (v.(string))
		o.SenderEmail = x

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
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
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

	url := "smtp/Policies"
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
	return resourceSmtpPolicyRead(d, meta)
}
func detachSMTPPolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "smtp/Policies" + "/" + d.Id()
	var o models.SMTPPolicy

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

func resourceSmtpPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "smtp/Policies" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.SMTPPolicy
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("class_id", (s.ClassID)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("enabled", (s.Enabled)); err != nil {
		return err
	}

	if err := d.Set("min_severity", (s.MinSeverity)); err != nil {
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

	if err := d.Set("organization", flattenMapOrganizationOrganizationRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRef(s.Profiles, d)); err != nil {
		return err
	}

	if err := d.Set("smtp_port", (s.SMTPPort)); err != nil {
		return err
	}

	if err := d.Set("smtp_recipients", (s.SMTPRecipients)); err != nil {
		return err
	}

	if err := d.Set("smtp_server", (s.SMTPServer)); err != nil {
		return err
	}

	if err := d.Set("sender_email", (s.SenderEmail)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceSmtpPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.SMTPPolicy
	if d.HasChange("class_id") {
		v := d.Get("class_id")
		x := (v.(string))
		o.ClassID = x
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

	if d.HasChange("min_severity") {
		v := d.Get("min_severity")
		x := (v.(string))
		o.MinSeverity = &x
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
		if len(v.([]interface{})) > 0 {
			o.Organization = &x
		}
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

	if d.HasChange("smtp_port") {
		v := d.Get("smtp_port")
		x := int64(v.(int))
		o.SMTPPort = x
	}

	if d.HasChange("smtp_recipients") {
		v := d.Get("smtp_recipients")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.SMTPRecipients = x
	}

	if d.HasChange("smtp_server") {
		v := d.Get("smtp_server")
		x := (v.(string))
		o.SMTPServer = x
	}

	if d.HasChange("sender_email") {
		v := d.Get("sender_email")
		x := (v.(string))
		o.SenderEmail = x
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
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
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

	url := "smtp/Policies" + "/" + d.Id()
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
	return resourceSmtpPolicyRead(d, meta)
}

func resourceSmtpPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "smtp/Policies" + "/" + d.Id()
	detachSMTPPolicyProfiles(d, meta)
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
