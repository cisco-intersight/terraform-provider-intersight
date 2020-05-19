package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirmwareUpgrade() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirmwareUpgradeCreate,
		Read:   resourceFirmwareUpgradeRead,
		Delete: resourceFirmwareUpgradeDelete,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
			"direct_download": {
				Description: "Direct download options in case the upgradeType is direct download based upgrade.",
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
						"http_server": {
							Description: "HTTP Server option when the image source is a local https server.",
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
									"location_link": {
										Description: "HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"mount_options": {
										Description: "Mount option as configured on the HTTP server. Empty if nothing is configured.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							ForceNew:   true,
						},
						"image_source": {
							Description: "Source type referring the image to be downloaded from CCO or from a local https server.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "cisco",
							ForceNew:    true,
						},
						"is_password_set": {
							Description: "Indicates whether the value of the 'password' property has been set.",
							Type:        schema.TypeBool,
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
						"password": {
							Description: "Password as configured on the local https server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"upgradeoption": {
							Description: "Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "sd_upgrade_mount_only",
							ForceNew:    true,
						},
						"username": {
							Description: "Username as configured on the local https server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"distributable": {
				Description: "A reference to a firmwareDistributable resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"network_share": {
				Description: "Network share options in case of the upgradeType is network share based upgrade.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cifs_server": {
							Description: "CIFS file server option for network share upgrade.",
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
									"file_location": {
										Description: "The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"mount_options": {
										Description: "Mount option (Authentication Protocol) as configured on the CIFS Server. Example:ntlmv2.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "none",
										ForceNew:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"remote_file": {
										Description: "Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"remote_ip": {
										Description: "CIFS Server Hostname or IP Address. Example:CIFS-server-hostname or 10.10.8.7.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"remote_share": {
										Description: "Directory where the image is stored. Example:share/subfolder.",
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
						"http_server": {
							Description: "HTTP (for WWW) file server option for network share upgrade.",
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
									"location_link": {
										Description: "HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"mount_options": {
										Description: "Mount option as configured on the HTTP server. Empty if nothing is configured.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							ForceNew:   true,
						},
						"is_password_set": {
							Description: "Indicates whether the value of the 'password' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"map_type": {
							Description: "File server protocols like CIFS, NFS, WWW for HTTP (S) that hosts the image.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "nfs",
							ForceNew:    true,
						},
						"nfs_server": {
							Description: "NFS file server option for network share upgrade.",
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
									"file_location": {
										Description: "The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"mount_options": {
										Description: "Mount option as configured on the NFS Server. Example:nolock.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"remote_file": {
										Description: "Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"remote_ip": {
										Description: "NFS Server Hostname or IP Address. Example:NFS-server-hostname or 10.10.8.7.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"remote_share": {
										Description: "Directory where the image is stored. Example:/share/subfolder.",
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
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"password": {
							Description: "Password as configured on the file server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"upgradeoption": {
							Description: "Option to control the upgrade, e.g., 1) nw_upgrade_mount_only - mount the image from a file server and run upgrade on-next server boot 2) nw_upgrade_full - mount the image and run upgrade immediately.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "nw_upgrade_full",
							ForceNew:    true,
						},
						"username": {
							Description: "Username as configured on the file server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"server": {
				Description: "A reference to a computeRackUnit resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"upgrade_status": {
				Description: "A reference to a firmwareUpgradeStatus resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
			"upgrade_type": {
				Description: "Desired upgrade mode to choose either direct download based upgrade or network share upgrade.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "direct_upgrade",
				ForceNew:    true,
			},
		},
	}
}

func resourceFirmwareUpgradeCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewFirmwareUpgrade()
	o.SetClassId("firmware.Upgrade")

	if v, ok := d.GetOk("device"); ok {
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
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
			o.SetObjectType("asset.DeviceRegistration")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsAssetDeviceRegistrationRelationship())
		}
		x := p[0]
		o.SetDevice(x)
	}

	if v, ok := d.GetOk("direct_download"); ok {
		p := make([]models.FirmwareDirectDownload, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewFirmwareDirectDownloadWithDefaults()
			o.SetClassId("firmware.DirectDownload")
			if v, ok := l["http_server"]; ok {
				{
					p := make([]models.FirmwareHttpServer, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewFirmwareHttpServerWithDefaults()
						o.SetClassId("firmware.HttpServer")
						if v, ok := l["location_link"]; ok {
							{
								x := (v.(string))
								o.SetLocationLink(x)
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.SetMountOptions(x)
							}
						}
						o.SetObjectType("firmware.HttpServer")
						p = append(p, *o)
					}
					x := p[0]
					o.SetHttpServer(x)
				}
			}
			if v, ok := l["image_source"]; ok {
				{
					x := (v.(string))
					o.SetImageSource(x)
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsPasswordSet(x)
				}
			}
			o.SetObjectType("firmware.DirectDownload")
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.SetPassword(x)
				}
			}
			if v, ok := l["upgradeoption"]; ok {
				{
					x := (v.(string))
					o.SetUpgradeoption(x)
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.SetUsername(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetDirectDownload(x)
	}

	if v, ok := d.GetOk("distributable"); ok {
		p := make([]models.FirmwareDistributableRelationship, 0, 1)
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
			o.SetObjectType("firmware.Distributable")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsFirmwareDistributableRelationship())
		}
		x := p[0]
		o.SetDistributable(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("network_share"); ok {
		p := make([]models.FirmwareNetworkShare, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewFirmwareNetworkShareWithDefaults()
			if v, ok := l["cifs_server"]; ok {
				{
					p := make([]models.FirmwareCifsServer, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewFirmwareCifsServerWithDefaults()
						o.SetClassId("firmware.CifsServer")
						if v, ok := l["file_location"]; ok {
							{
								x := (v.(string))
								o.SetFileLocation(x)
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.SetMountOptions(x)
							}
						}
						o.SetObjectType("firmware.CifsServer")
						if v, ok := l["remote_file"]; ok {
							{
								x := (v.(string))
								o.SetRemoteFile(x)
							}
						}
						if v, ok := l["remote_ip"]; ok {
							{
								x := (v.(string))
								o.SetRemoteIp(x)
							}
						}
						if v, ok := l["remote_share"]; ok {
							{
								x := (v.(string))
								o.SetRemoteShare(x)
							}
						}
						p = append(p, *o)
					}
					x := p[0]
					o.SetCifsServer(x)
				}
			}
			o.SetClassId("firmware.NetworkShare")
			if v, ok := l["http_server"]; ok {
				{
					p := make([]models.FirmwareHttpServer, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewFirmwareHttpServerWithDefaults()
						o.SetClassId("firmware.HttpServer")
						if v, ok := l["location_link"]; ok {
							{
								x := (v.(string))
								o.SetLocationLink(x)
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.SetMountOptions(x)
							}
						}
						o.SetObjectType("firmware.HttpServer")
						p = append(p, *o)
					}
					x := p[0]
					o.SetHttpServer(x)
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsPasswordSet(x)
				}
			}
			if v, ok := l["map_type"]; ok {
				{
					x := (v.(string))
					o.SetMapType(x)
				}
			}
			if v, ok := l["nfs_server"]; ok {
				{
					p := make([]models.FirmwareNfsServer, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewFirmwareNfsServerWithDefaults()
						o.SetClassId("firmware.NfsServer")
						if v, ok := l["file_location"]; ok {
							{
								x := (v.(string))
								o.SetFileLocation(x)
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.SetMountOptions(x)
							}
						}
						o.SetObjectType("firmware.NfsServer")
						if v, ok := l["remote_file"]; ok {
							{
								x := (v.(string))
								o.SetRemoteFile(x)
							}
						}
						if v, ok := l["remote_ip"]; ok {
							{
								x := (v.(string))
								o.SetRemoteIp(x)
							}
						}
						if v, ok := l["remote_share"]; ok {
							{
								x := (v.(string))
								o.SetRemoteShare(x)
							}
						}
						p = append(p, *o)
					}
					x := p[0]
					o.SetNfsServer(x)
				}
			}
			o.SetObjectType("firmware.NetworkShare")
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.SetPassword(x)
				}
			}
			if v, ok := l["upgradeoption"]; ok {
				{
					x := (v.(string))
					o.SetUpgradeoption(x)
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.SetUsername(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetNetworkShare(x)
	}

	o.SetObjectType("firmware.Upgrade")

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

	if v, ok := d.GetOk("server"); ok {
		p := make([]models.ComputeRackUnitRelationship, 0, 1)
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
			o.SetObjectType("compute.RackUnit")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsComputeRackUnitRelationship())
		}
		x := p[0]
		o.SetServer(x)
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

	if v, ok := d.GetOk("upgrade_status"); ok {
		p := make([]models.FirmwareUpgradeStatusRelationship, 0, 1)
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
			o.SetObjectType("firmware.UpgradeStatus")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsFirmwareUpgradeStatusRelationship())
		}
		x := p[0]
		o.SetUpgradeStatus(x)
	}

	if v, ok := d.GetOk("upgrade_type"); ok {
		x := (v.(string))
		o.SetUpgradeType(x)
	}

	r := conn.ApiClient.FirmwareApi.CreateFirmwareUpgrade(conn.ctx).FirmwareUpgrade(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceFirmwareUpgradeRead(d, meta)
}

func resourceFirmwareUpgradeRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.FirmwareApi.GetFirmwareUpgradeByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("device", flattenMapAssetDeviceRegistrationRelationship(s.Device, d)); err != nil {
		return err
	}

	if err := d.Set("direct_download", flattenMapFirmwareDirectDownload(s.DirectDownload, d)); err != nil {
		return err
	}

	if err := d.Set("distributable", flattenMapFirmwareDistributableRelationship(s.Distributable, d)); err != nil {
		return err
	}

	if err := d.Set("moid", (s.Moid)); err != nil {
		return err
	}

	if err := d.Set("network_share", flattenMapFirmwareNetworkShare(s.NetworkShare, d)); err != nil {
		return err
	}

	if err := d.Set("object_type", (s.ObjectType)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("server", flattenMapComputeRackUnitRelationship(s.Server, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("upgrade_status", flattenMapFirmwareUpgradeStatusRelationship(s.UpgradeStatus, d)); err != nil {
		return err
	}

	if err := d.Set("upgrade_type", (s.UpgradeType)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceFirmwareUpgradeDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.FirmwareApi.DeleteFirmwareUpgrade(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
