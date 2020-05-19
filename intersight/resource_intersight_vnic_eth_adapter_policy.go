package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVnicEthAdapterPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVnicEthAdapterPolicyCreate,
		Read:   resourceVnicEthAdapterPolicyRead,
		Update: resourceVnicEthAdapterPolicyUpdate,
		Delete: resourceVnicEthAdapterPolicyDelete,
		Schema: map[string]*schema.Schema{
			"advanced_filter": {
				Description: "Enables advanced filtering on the interface.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"arfs_settings": {
				Description: "Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency.",
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
						"enabled": {
							Description: "Status of Accelerated Receive Flow Steering on the virtual ethernet interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"completion_queue_settings": {
				Description: "Completion Queue resource settings.",
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
						"count": {
							Description: "The number of completion queue resources to allocate. In general, the number of completion queue resources you should allocate is equal to the number of transmit queue resources plus the number of receive queue resources.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each completion queue.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
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
			"interrupt_settings": {
				Description: "Interrupt Settings for the virtual ethernet interface.",
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
						"coalescing_time": {
							Description: "The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"coalescing_type": {
							Description: "Interrupt Coalescing Type. This can be one of the following:- MIN  - The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MIN",
						},
						"count": {
							Description: "The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mode": {
							Description: "Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MSIx",
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"nvgre_settings": {
				Description: "Network Virtualization using Generic Routing Encapsulation Settings.",
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
						"enabled": {
							Description: "Status of the Network Virtualization using Generic Routing Encapsulation on the virtual ethernet interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
			"roce_settings": {
				Description: "Settings for RDMA over Converged Ethernet.",
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
						"enabled": {
							Description: "If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"memory_regions": {
							Description: "The number of memory regions per adapter. Recommended value = integer power of 2.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"queue_pairs": {
							Description: "The number of queue pairs per adapter. Recommended value = integer power of 2.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"resource_groups": {
							Description: "The number of resource groups per adapter. Recommended value = be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"rss_settings": {
				Description: "Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"rx_queue_settings": {
				Description: "Receive Queue resouce settings.",
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
						"count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each queue.",
							Type:        schema.TypeInt,
							Optional:    true,
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
			"tcp_offload_settings": {
				Description: "The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not.",
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
						"large_receive": {
							Description: "Enables the reassembly of segmented packets in hardware before sending them to the CPU.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"large_send": {
							Description: "Enables the CPU to send large packets to the hardware for segmentation.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"rx_checksum": {
							Description: "When enabled, the CPU sends all packet checksums to the hardware for validation.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tx_checksum": {
							Description: "When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"tx_queue_settings": {
				Description: "Transmit Queue resource settings.",
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
						"count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each queue.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"vxlan_settings": {
				Description: "Virtual Extensible LAN Protocol Settings.",
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
						"enabled": {
							Description: "Status of the Virtual Extensible LAN Protocol on the virtual ethernet interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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

func resourceVnicEthAdapterPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewVnicEthAdapterPolicy()
	if v, ok := d.GetOk("advanced_filter"); ok {
		x := (v.(bool))
		o.SetAdvancedFilter(x)
	}

	if v, ok := d.GetOk("arfs_settings"); ok {
		p := make([]models.VnicArfsSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicArfsSettingsWithDefaults()
			o.SetClassId("vnic.ArfsSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			o.SetObjectType("vnic.ArfsSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetArfsSettings(x)
	}

	o.SetClassId("vnic.EthAdapterPolicy")

	if v, ok := d.GetOk("completion_queue_settings"); ok {
		p := make([]models.VnicCompletionQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicCompletionQueueSettingsWithDefaults()
			o.SetClassId("vnic.CompletionQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.CompletionQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetCompletionQueueSettings(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("interrupt_settings"); ok {
		p := make([]models.VnicEthInterruptSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicEthInterruptSettingsWithDefaults()
			o.SetClassId("vnic.EthInterruptSettings")
			if v, ok := l["coalescing_time"]; ok {
				{
					x := int64(v.(int))
					o.SetCoalescingTime(x)
				}
			}
			if v, ok := l["coalescing_type"]; ok {
				{
					x := (v.(string))
					o.SetCoalescingType(x)
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.SetMode(x)
				}
			}
			o.SetObjectType("vnic.EthInterruptSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetInterruptSettings(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	if v, ok := d.GetOk("nvgre_settings"); ok {
		p := make([]models.VnicNvgreSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicNvgreSettingsWithDefaults()
			o.SetClassId("vnic.NvgreSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			o.SetObjectType("vnic.NvgreSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetNvgreSettings(x)
	}

	o.SetObjectType("vnic.EthAdapterPolicy")

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

	if v, ok := d.GetOk("roce_settings"); ok {
		p := make([]models.VnicRoceSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicRoceSettingsWithDefaults()
			o.SetClassId("vnic.RoceSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["memory_regions"]; ok {
				{
					x := int64(v.(int))
					o.SetMemoryRegions(x)
				}
			}
			o.SetObjectType("vnic.RoceSettings")
			if v, ok := l["queue_pairs"]; ok {
				{
					x := int64(v.(int))
					o.SetQueuePairs(x)
				}
			}
			if v, ok := l["resource_groups"]; ok {
				{
					x := int64(v.(int))
					o.SetResourceGroups(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetRoceSettings(x)
	}

	if v, ok := d.GetOk("rss_settings"); ok {
		x := (v.(bool))
		o.SetRssSettings(x)
	}

	if v, ok := d.GetOk("rx_queue_settings"); ok {
		p := make([]models.VnicEthRxQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicEthRxQueueSettingsWithDefaults()
			o.SetClassId("vnic.EthRxQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.EthRxQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetRxQueueSettings(x)
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

	if v, ok := d.GetOk("tcp_offload_settings"); ok {
		p := make([]models.VnicTcpOffloadSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicTcpOffloadSettingsWithDefaults()
			o.SetClassId("vnic.TcpOffloadSettings")
			if v, ok := l["large_receive"]; ok {
				{
					x := (v.(bool))
					o.SetLargeReceive(x)
				}
			}
			if v, ok := l["large_send"]; ok {
				{
					x := (v.(bool))
					o.SetLargeSend(x)
				}
			}
			o.SetObjectType("vnic.TcpOffloadSettings")
			if v, ok := l["rx_checksum"]; ok {
				{
					x := (v.(bool))
					o.SetRxChecksum(x)
				}
			}
			if v, ok := l["tx_checksum"]; ok {
				{
					x := (v.(bool))
					o.SetTxChecksum(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetTcpOffloadSettings(x)
	}

	if v, ok := d.GetOk("tx_queue_settings"); ok {
		p := make([]models.VnicEthTxQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicEthTxQueueSettingsWithDefaults()
			o.SetClassId("vnic.EthTxQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.EthTxQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetTxQueueSettings(x)
	}

	if v, ok := d.GetOk("vxlan_settings"); ok {
		p := make([]models.VnicVxlanSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicVxlanSettingsWithDefaults()
			o.SetClassId("vnic.VxlanSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			o.SetObjectType("vnic.VxlanSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetVxlanSettings(x)
	}

	r := conn.ApiClient.VnicApi.CreateVnicEthAdapterPolicy(conn.ctx).VnicEthAdapterPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceVnicEthAdapterPolicyRead(d, meta)
}

func resourceVnicEthAdapterPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.VnicApi.GetVnicEthAdapterPolicyByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("advanced_filter", (s.AdvancedFilter)); err != nil {
		return err
	}

	if err := d.Set("arfs_settings", flattenMapVnicArfsSettings(s.ArfsSettings, d)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("completion_queue_settings", flattenMapVnicCompletionQueueSettings(s.CompletionQueueSettings, d)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("interrupt_settings", flattenMapVnicEthInterruptSettings(s.InterruptSettings, d)); err != nil {
		return err
	}

	if err := d.Set("moid", (s.Moid)); err != nil {
		return err
	}

	if err := d.Set("name", (s.Name)); err != nil {
		return err
	}

	if err := d.Set("nvgre_settings", flattenMapVnicNvgreSettings(s.NvgreSettings, d)); err != nil {
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

	if err := d.Set("roce_settings", flattenMapVnicRoceSettings(s.RoceSettings, d)); err != nil {
		return err
	}

	if err := d.Set("rss_settings", (s.RssSettings)); err != nil {
		return err
	}

	if err := d.Set("rx_queue_settings", flattenMapVnicEthRxQueueSettings(s.RxQueueSettings, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("tcp_offload_settings", flattenMapVnicTcpOffloadSettings(s.TcpOffloadSettings, d)); err != nil {
		return err
	}

	if err := d.Set("tx_queue_settings", flattenMapVnicEthTxQueueSettings(s.TxQueueSettings, d)); err != nil {
		return err
	}

	if err := d.Set("vxlan_settings", flattenMapVnicVxlanSettings(s.VxlanSettings, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceVnicEthAdapterPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewVnicEthAdapterPolicy()
	if d.HasChange("advanced_filter") {
		v := d.Get("advanced_filter")
		x := (v.(bool))
		o.SetAdvancedFilter(x)
	}

	if d.HasChange("arfs_settings") {
		v := d.Get("arfs_settings")
		p := make([]models.VnicArfsSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicArfsSettingsWithDefaults()
			o.SetClassId("vnic.ArfsSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			o.SetObjectType("vnic.ArfsSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetArfsSettings(x)
	}

	if d.HasChange("completion_queue_settings") {
		v := d.Get("completion_queue_settings")
		p := make([]models.VnicCompletionQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicCompletionQueueSettingsWithDefaults()
			o.SetClassId("vnic.CompletionQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.CompletionQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetCompletionQueueSettings(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("interrupt_settings") {
		v := d.Get("interrupt_settings")
		p := make([]models.VnicEthInterruptSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicEthInterruptSettingsWithDefaults()
			o.SetClassId("vnic.EthInterruptSettings")
			if v, ok := l["coalescing_time"]; ok {
				{
					x := int64(v.(int))
					o.SetCoalescingTime(x)
				}
			}
			if v, ok := l["coalescing_type"]; ok {
				{
					x := (v.(string))
					o.SetCoalescingType(x)
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.SetMode(x)
				}
			}
			o.SetObjectType("vnic.EthInterruptSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetInterruptSettings(x)
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

	if d.HasChange("nvgre_settings") {
		v := d.Get("nvgre_settings")
		p := make([]models.VnicNvgreSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicNvgreSettingsWithDefaults()
			o.SetClassId("vnic.NvgreSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			o.SetObjectType("vnic.NvgreSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetNvgreSettings(x)
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

	if d.HasChange("roce_settings") {
		v := d.Get("roce_settings")
		p := make([]models.VnicRoceSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicRoceSettingsWithDefaults()
			o.SetClassId("vnic.RoceSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["memory_regions"]; ok {
				{
					x := int64(v.(int))
					o.SetMemoryRegions(x)
				}
			}
			o.SetObjectType("vnic.RoceSettings")
			if v, ok := l["queue_pairs"]; ok {
				{
					x := int64(v.(int))
					o.SetQueuePairs(x)
				}
			}
			if v, ok := l["resource_groups"]; ok {
				{
					x := int64(v.(int))
					o.SetResourceGroups(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetRoceSettings(x)
	}

	if d.HasChange("rss_settings") {
		v := d.Get("rss_settings")
		x := (v.(bool))
		o.SetRssSettings(x)
	}

	if d.HasChange("rx_queue_settings") {
		v := d.Get("rx_queue_settings")
		p := make([]models.VnicEthRxQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicEthRxQueueSettingsWithDefaults()
			o.SetClassId("vnic.EthRxQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.EthRxQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetRxQueueSettings(x)
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

	if d.HasChange("tcp_offload_settings") {
		v := d.Get("tcp_offload_settings")
		p := make([]models.VnicTcpOffloadSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicTcpOffloadSettingsWithDefaults()
			o.SetClassId("vnic.TcpOffloadSettings")
			if v, ok := l["large_receive"]; ok {
				{
					x := (v.(bool))
					o.SetLargeReceive(x)
				}
			}
			if v, ok := l["large_send"]; ok {
				{
					x := (v.(bool))
					o.SetLargeSend(x)
				}
			}
			o.SetObjectType("vnic.TcpOffloadSettings")
			if v, ok := l["rx_checksum"]; ok {
				{
					x := (v.(bool))
					o.SetRxChecksum(x)
				}
			}
			if v, ok := l["tx_checksum"]; ok {
				{
					x := (v.(bool))
					o.SetTxChecksum(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetTcpOffloadSettings(x)
	}

	if d.HasChange("tx_queue_settings") {
		v := d.Get("tx_queue_settings")
		p := make([]models.VnicEthTxQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicEthTxQueueSettingsWithDefaults()
			o.SetClassId("vnic.EthTxQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.EthTxQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetTxQueueSettings(x)
	}

	if d.HasChange("vxlan_settings") {
		v := d.Get("vxlan_settings")
		p := make([]models.VnicVxlanSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicVxlanSettingsWithDefaults()
			o.SetClassId("vnic.VxlanSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			o.SetObjectType("vnic.VxlanSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetVxlanSettings(x)
	}

	r := conn.ApiClient.VnicApi.UpdateVnicEthAdapterPolicy(conn.ctx, d.Id()).VnicEthAdapterPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceVnicEthAdapterPolicyRead(d, meta)
}

func resourceVnicEthAdapterPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.VnicApi.DeleteVnicEthAdapterPolicy(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
