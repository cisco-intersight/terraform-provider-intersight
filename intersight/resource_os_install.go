package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func resourceOsInstall() *schema.Resource {
	return &schema.Resource{
		Create: resourceOsInstallCreate,
		Read:   resourceOsInstallRead,
		Delete: resourceOsInstallDelete,
		Schema: map[string]*schema.Schema{
			"additional_parameters": {
				Description: "If the os.ConfigurationFile MO selected is a template that uses additionalplaceholders other than the ones provided in standard os.Answers MO, the valuesfor those additional placeholders are provided here.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
				ForceNew: true,
			},
			"answers": {
				Description: "Answers provided by user for the unattended OS installation.",
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
						"answer_file": {
							Description: "If the source of the answers is a static file, the content of the file is stored as valuein this property.The value is mandatory only when the 'Source' property has been set to 'File'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"hostname": {
							Description: "Hostname to be configured for the server in the OS. Hostname property is requiredwhen the 'Source' property has been set to 'Template'.Please note that this property will be ignored for the answer instance usedinside the os.InstallTemplate MO since this property must have unique value foreach server and is irrelevant in a reusable os.InstallTemplate MO. The value ofthis property specified in os.Install MO will be considered for the installation.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"ip_config_type": {
							Description: "IP configuration type. Values are Static or Dynamic when the Source property is set to Template.In case of static IP configuration, IP address, gateway and other details needto be populated. In case of dynamic the IP configuration is obtained dynamicallyfrom DHCP.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "static",
							ForceNew:    true,
						},
						"ipv4_config": {
							Description: "In case of static IP configuration, IP address, netmask and gateway details areprovided.Please note that this property will be ignored for the answer instance usedinside the os.InstallTemplate MO since this property must have unique value foreach server and is irrelevant in a reusable os.InstallTemplate MO. The value ofthis property specified in os.Install MO will be considered for the installation.",
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
									"gateway": {
										Description: "The IPv4 address of the default gateway.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"ip_address": {
										Description: "The IPv4 Address, represented in the standard dot-decimal notation, e.g. 192.168.1.3.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"netmask": {
										Description: "The IPv4 Netmask, represented in the standard dot-decimal notation, e.g. 255.255.255.0.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
								},
							},
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							ForceNew:   true,
						},
						"is_answer_file_set": {
							Description: "Indicates whether the value of the 'answerFile' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"is_root_password_set": {
							Description: "",
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
						},
						"nameserver": {
							Description: "IP address of the name server to be configured in the OS in case the 'Source'property has been set to 'Template'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"root_password": {
							Description: "Password to be set for root/administrator user in case the 'Source'property has been set to 'Template'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"source": {
							Description: "Answer values can be provided from three sources - Embedded in OS image, static file,or as placeholder values for an answer file template.Source of the answers is given as value, Embedded/File/Template.'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'option indicates that the answers are provided as a file. 'Template' indicates that theplaceholders in the selected os.ConfigurationFile MO are replaced with values providedas os.Answers MO.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Embedded",
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"configuration_file": {
				Description: "If the answers source is selected as 'Template' in 'Answers' property, this relation provides the os.ConfigurationFile instance to be used for this OS install.",
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"description": {
				Description: "User provided description about the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"image": {
				Description: "OS Image to be installed as part of this OS installation.",
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"install_method": {
				Description: "The install method to be used for OS installation - vMedia, iPXE. Only vMedia is supported as of now.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "vMedia",
				ForceNew:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"osdu_image": {
				Description: "Location of the Intersight OS Deployment Utilityimage, if the user has downloaded and available locally, to be used for this OS install configuration. This image is applicable for vMedia install method.Cisco publishes a OS Deployment Utility image that bootstraps and installs the user provided operating system images along with answers for unattended instllation.If this property is empty for vMedia install type, the image hosted in Intersight image repository will be used. Note that in this case, the image will be downloaded from Intersight in every target server every time.",
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"post_install_scripts": {
				Description: "Post Install Scripts to be executed specified in order.",
				Type:        schema.TypeList,
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"server": {
				Description: "This relation provides the target server in which the OS is to beinstalled.",
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"workflow_info": {
				Description: "This relation is populated with the reference of OS install workflowstarted for this request. This workflow info MO shall be used fortracking further status and completion.",
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
		},
	}
}
func resourceOsInstallCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.OsInstall
	if v, ok := d.GetOk("additional_parameters"); ok {
		x := v
		o.AdditionalParameters = &x

	}

	if v, ok := d.GetOk("answers"); ok {
		p := models.OsAnswers{}
		if len(v.([]interface{})) > 0 {
			o := models.OsAnswers{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.OsAnswersAO0P0.OsAnswersAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["answer_file"]; ok {
				{
					x := (v.(string))
					o.AnswerFile = x
				}
			}
			if v, ok := l["hostname"]; ok {
				{
					x := (v.(string))
					o.Hostname = x
				}
			}
			if v, ok := l["ip_config_type"]; ok {
				{
					x := (v.(string))
					o.IPConfigType = &x
				}
			}
			if v, ok := l["ipv4_config"]; ok {
				{
					p := models.CommIPV4Interface{}
					if len(v.([]interface{})) > 0 {
						o := models.CommIPV4Interface{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.CommIPV4InterfaceAO0P0.CommIPV4InterfaceAO0P0 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["gateway"]; ok {
							{
								x := (v.(string))
								o.Gateway = x
							}
						}
						if v, ok := l["ip_address"]; ok {
							{
								x := (v.(string))
								o.IPAddress = x
							}
						}
						if v, ok := l["netmask"]; ok {
							{
								x := (v.(string))
								o.Netmask = x
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
					o.IPV4Config = &x
				}
			}
			if v, ok := l["is_answer_file_set"]; ok {
				{
					x := (v.(bool))
					o.IsAnswerFileSet = &x
				}
			}
			if v, ok := l["is_root_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsRootPasswordSet = &x
				}
			}
			if v, ok := l["nameserver"]; ok {
				{
					x := (v.(string))
					o.Nameserver = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["root_password"]; ok {
				{
					x := (v.(string))
					o.RootPassword = x
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
		o.Answers = &x

	}

	if v, ok := d.GetOk("configuration_file"); ok {
		p := models.OsConfigurationFileRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OsConfigurationFileRef{}
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
		o.ConfigurationFile = &x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("image"); ok {
		p := models.SoftwarerepositoryOperatingSystemFileRef{}
		if len(v.([]interface{})) > 0 {
			o := models.SoftwarerepositoryOperatingSystemFileRef{}
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
		o.Image = &x

	}

	if v, ok := d.GetOk("install_method"); ok {
		x := (v.(string))
		o.InstallMethod = &x

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

	if v, ok := d.GetOk("osdu_image"); ok {
		p := models.FirmwareServerConfigurationUtilityDistributableRef{}
		if len(v.([]interface{})) > 0 {
			o := models.FirmwareServerConfigurationUtilityDistributableRef{}
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
		o.OsduImage = &x

	}

	if v, ok := d.GetOk("post_install_scripts"); ok {
		x := make([]*models.OsPostInstallScriptRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.OsPostInstallScriptRef{}
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
		o.PostInstallScripts = x

	}

	if v, ok := d.GetOk("server"); ok {
		p := models.ComputePhysicalRef{}
		if len(v.([]interface{})) > 0 {
			o := models.ComputePhysicalRef{}
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

	if v, ok := d.GetOk("workflow_info"); ok {
		p := models.WorkflowWorkflowInfoRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowWorkflowInfoRef{}
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
		o.WorkflowInfo = &x

	}

	url := "os/Installs"
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
	return resourceOsInstallRead(d, meta)
}

func resourceOsInstallRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "os/Installs" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.OsInstall
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("additional_parameters", (s.AdditionalParameters)); err != nil {
		return err
	}

	if err := d.Set("answers", flattenMapOsAnswers(s.Answers, d)); err != nil {
		return err
	}

	if err := d.Set("configuration_file", flattenMapOsConfigurationFileRef(s.ConfigurationFile, d)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("image", flattenMapSoftwarerepositoryOperatingSystemFileRef(s.Image, d)); err != nil {
		return err
	}

	if err := d.Set("install_method", (s.InstallMethod)); err != nil {
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

	if err := d.Set("osdu_image", flattenMapFirmwareServerConfigurationUtilityDistributableRef(s.OsduImage, d)); err != nil {
		return err
	}

	if err := d.Set("post_install_scripts", flattenListOsPostInstallScriptRef(s.PostInstallScripts, d)); err != nil {
		return err
	}

	if err := d.Set("server", flattenMapComputePhysicalRef(s.Server, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRef(s.WorkflowInfo, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}

func resourceOsInstallDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	x := d.Get("workflow_info").([]interface{})[0].(map[string]interface{})
	url := "workflow/WorkflowInfos" + "/" + (x["moid"].(string))
	data, err := conn.SendGetRequest(url, []byte(""))
	var dData map[string]interface{}
	err = json.Unmarshal(data, &dData)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
		return err
	}
	if v, ok := dData["Status"]; ok && v == "RUNNING" {
		payload := map[string]string{"Action": "Cancel"}
		payloadJson, _ := json.Marshal(payload)
		_, err = conn.SendUpdateRequest(url, payloadJson)
		if err != nil {
			log.Printf("error occurred while deleting: %s", err.Error())
		}
	}
	return err
}
