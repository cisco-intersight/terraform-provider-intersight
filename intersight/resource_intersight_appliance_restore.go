package intersight

import (
	"log"
	"reflect"
	"time"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApplianceRestore() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplianceRestoreCreate,
		Read:   resourceApplianceRestoreRead,
		Delete: resourceApplianceRestoreDelete,
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
							ForceNew:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"elapsed_time": {
				Description: "Elapsed time in seconds since the restore process has started.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"end_time": {
				Description: "End date and time of the restore process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"filename": {
				Description: "Backup filename to backup or restore.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}, ForceNew: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"password": {
				Description: "Password for authenticating with the file server.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
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
							ForceNew:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"protocol": {
				Description: "Communication protocol used by the file server (e.g. scp or sftp).",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "scp",
				ForceNew:    true,
			},
			"remote_host": {
				Description: "Hostname of the remote file server.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"remote_path": {
				Description: "File server directory to copy the file.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"remote_port": {
				Description: "Remote TCP port on the file server (e.g. 22 for scp).",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"start_time": {
				Description: "Start date and time of the restore process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"status": {
				Description: "Status of the restore managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							ForceNew:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ForceNew: true,
			},
			"username": {
				Description: "Username to authenticate the fileserver.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceApplianceRestoreCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewApplianceRestore()
	if v, ok := d.GetOk("account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
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
			o.SetObjectType("iam.Account")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsIamAccountRelationship())
		}
		x := p[0]
		o.SetAccount(x)
	}

	o.SetClassId("appliance.Restore")

	if v, ok := d.GetOk("elapsed_time"); ok {
		x := int64(v.(int))
		o.SetElapsedTime(x)
	}

	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}

	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.SetFilename(x)
	}

	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}

	if v, ok := d.GetOk("messages"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.SetMessages(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("appliance.Restore")

	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.SetPassword(x)
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

	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}

	if v, ok := d.GetOk("remote_host"); ok {
		x := (v.(string))
		o.SetRemoteHost(x)
	}

	if v, ok := d.GetOk("remote_path"); ok {
		x := (v.(string))
		o.SetRemotePath(x)
	}

	if v, ok := d.GetOk("remote_port"); ok {
		x := int64(v.(int))
		o.SetRemotePort(x)
	}

	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
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

	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	r := conn.ApiClient.ApplianceApi.CreateApplianceRestore(conn.ctx).ApplianceRestore(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceApplianceRestoreRead(d, meta)
}

func resourceApplianceRestoreRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.ApplianceApi.GetApplianceRestoreByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("account", flattenMapIamAccountRelationship(s.Account, d)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("elapsed_time", (s.ElapsedTime)); err != nil {
		return err
	}

	if err := d.Set("end_time", (s.EndTime).String()); err != nil {
		return err
	}

	if err := d.Set("filename", (s.Filename)); err != nil {
		return err
	}

	if err := d.Set("is_password_set", (s.IsPasswordSet)); err != nil {
		return err
	}

	if err := d.Set("messages", (s.Messages)); err != nil {
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

	if err := d.Set("protocol", (s.Protocol)); err != nil {
		return err
	}

	if err := d.Set("remote_host", (s.RemoteHost)); err != nil {
		return err
	}

	if err := d.Set("remote_path", (s.RemotePath)); err != nil {
		return err
	}

	if err := d.Set("remote_port", (s.RemotePort)); err != nil {
		return err
	}

	if err := d.Set("start_time", (s.StartTime).String()); err != nil {
		return err
	}

	if err := d.Set("status", (s.Status)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("username", (s.Username)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceApplianceRestoreDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.ApplianceApi.DeleteApplianceRestore(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
