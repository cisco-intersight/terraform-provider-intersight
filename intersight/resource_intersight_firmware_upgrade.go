package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirmwareUpgrade() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirmwareUpgradeCreate,
		Read:   resourceFirmwareUpgradeRead,
		Delete: resourceFirmwareUpgradeDelete,
		Schema: map[string]*schema.Schema{
			"device": {
				Description: "The device onto which the upgrade is peformed.",
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
							ForceNew:         true,
						},
						"http_server": {
							Description: "HTTP Server option when the image source is a local https server.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
										ForceNew:         true,
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
										Computed:    true,
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
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
				Description: "The image that is used to upgrade the server for direct download upgrade type operation.",
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
							ForceNew:         true,
						},
						"cifs_server": {
							Description: "CIFS file server option for network share upgrade.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
										ForceNew:         true,
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
										Computed:    true,
										ForceNew:    true,
									},
									"remote_file": {
										Description: "Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"remote_ip": {
										Description: "CIFS Server Hostname or IP Address. Example:CIFS-server-hostname or 10.10.8.7.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"remote_share": {
										Description: "Directory where the image is stored. Example:share/subfolder.",
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
						"http_server": {
							Description: "HTTP (for WWW) file server option for network share upgrade.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
										ForceNew:         true,
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
										Computed:    true,
										ForceNew:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							ForceNew:   true,
						},
						"is_password_set": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
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
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
										ForceNew:         true,
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
										Computed:    true,
										ForceNew:    true,
									},
									"remote_file": {
										Description: "Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"remote_ip": {
										Description: "NFS Server Hostname or IP Address. Example:NFS-server-hostname or 10.10.8.7.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"remote_share": {
										Description: "Directory where the image is stored. Example:/share/subfolder.",
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
				Description: "The server on which this upgrade operation is performed.",
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
				Description: "The array of tags, which allow to add key, value meta-data to managed objects.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
							ForceNew:         true,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"upgrade_status": {
				Description: "Captures status of this upgrade information.",
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
	var o models.FirmwareUpgrade
	if v, ok := d.GetOk("device"); ok {
		p := models.AssetDeviceRegistrationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetDeviceRegistrationRef{}
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
		o.Device = &x

	}

	if v, ok := d.GetOk("direct_download"); ok {
		p := models.FirmwareDirectDownload{}
		if len(v.([]interface{})) > 0 {
			o := models.FirmwareDirectDownload{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["http_server"]; ok {
				{
					p := models.FirmwareHTTPServer{}
					if len(v.([]interface{})) > 0 {
						o := models.FirmwareHTTPServer{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["location_link"]; ok {
							{
								x := (v.(string))
								o.LocationLink = x
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.MountOptions = x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}

						p = o
					}
					x := p
					o.HTTPServer = &x
				}
			}
			if v, ok := l["image_source"]; ok {
				{
					x := (v.(string))
					o.ImageSource = &x
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsPasswordSet = &x
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
			if v, ok := l["upgradeoption"]; ok {
				{
					x := (v.(string))
					o.Upgradeoption = &x
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.Username = x
				}
			}

			p = o
		}
		x := p
		o.DirectDownload = &x

	}

	if v, ok := d.GetOk("distributable"); ok {
		p := models.FirmwareDistributableRef{}
		if len(v.([]interface{})) > 0 {
			o := models.FirmwareDistributableRef{}
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
		o.Distributable = &x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("network_share"); ok {
		p := models.FirmwareNetworkShare{}
		if len(v.([]interface{})) > 0 {
			o := models.FirmwareNetworkShare{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["cifs_server"]; ok {
				{
					p := models.FirmwareCifsServer{}
					if len(v.([]interface{})) > 0 {
						o := models.FirmwareCifsServer{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.FirmwareCifsServerAO1P1.FirmwareCifsServerAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.MountOptions = &x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}
						if v, ok := l["remote_file"]; ok {
							{
								x := (v.(string))
								o.RemoteFile = x
							}
						}
						if v, ok := l["remote_ip"]; ok {
							{
								x := (v.(string))
								o.RemoteIP = x
							}
						}
						if v, ok := l["remote_share"]; ok {
							{
								x := (v.(string))
								o.RemoteShare = x
							}
						}

						p = o
					}
					x := p
					o.CifsServer = &x
				}
			}
			if v, ok := l["http_server"]; ok {
				{
					p := models.FirmwareHTTPServer{}
					if len(v.([]interface{})) > 0 {
						o := models.FirmwareHTTPServer{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["location_link"]; ok {
							{
								x := (v.(string))
								o.LocationLink = x
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.MountOptions = x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}

						p = o
					}
					x := p
					o.HTTPServer = &x
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsPasswordSet = &x
				}
			}
			if v, ok := l["map_type"]; ok {
				{
					x := (v.(string))
					o.MapType = &x
				}
			}
			if v, ok := l["nfs_server"]; ok {
				{
					p := models.FirmwareNfsServer{}
					if len(v.([]interface{})) > 0 {
						o := models.FirmwareNfsServer{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.FirmwareNfsServerAO1P1.FirmwareNfsServerAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["mount_options"]; ok {
							{
								x := (v.(string))
								o.MountOptions = x
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.ObjectType = x
							}
						}
						if v, ok := l["remote_file"]; ok {
							{
								x := (v.(string))
								o.RemoteFile = x
							}
						}
						if v, ok := l["remote_ip"]; ok {
							{
								x := (v.(string))
								o.RemoteIP = x
							}
						}
						if v, ok := l["remote_share"]; ok {
							{
								x := (v.(string))
								o.RemoteShare = x
							}
						}

						p = o
					}
					x := p
					o.NfsServer = &x
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
			if v, ok := l["upgradeoption"]; ok {
				{
					x := (v.(string))
					o.Upgradeoption = &x
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.Username = x
				}
			}

			p = o
		}
		x := p
		o.NetworkShare = &x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

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

	if v, ok := d.GetOk("server"); ok {
		p := models.ComputeRackUnitRef{}
		if len(v.([]interface{})) > 0 {
			o := models.ComputeRackUnitRef{}
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
		o.Server = &x

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

	if v, ok := d.GetOk("upgrade_status"); ok {
		p := models.FirmwareUpgradeStatusRef{}
		if len(v.([]interface{})) > 0 {
			o := models.FirmwareUpgradeStatusRef{}
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
		o.UpgradeStatus = &x

	}

	if v, ok := d.GetOk("upgrade_type"); ok {
		x := (v.(string))
		o.UpgradeType = &x

	}

	url := "firmware/Upgrades"
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
	return resourceFirmwareUpgradeRead(d, meta)
}

func resourceFirmwareUpgradeRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "firmware/Upgrades" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.FirmwareUpgrade
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("device", flattenMapAssetDeviceRegistrationRef(s.Device, d)); err != nil {
		return err
	}

	if err := d.Set("direct_download", flattenMapFirmwareDirectDownload(s.DirectDownload, d)); err != nil {
		return err
	}

	if err := d.Set("distributable", flattenMapFirmwareDistributableRef(s.Distributable, d)); err != nil {
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

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("server", flattenMapComputeRackUnitRef(s.Server, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("upgrade_status", flattenMapFirmwareUpgradeStatusRef(s.UpgradeStatus, d)); err != nil {
		return err
	}

	if err := d.Set("upgrade_type", (s.UpgradeType)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}

func resourceFirmwareUpgradeDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "firmware/Upgrades" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
