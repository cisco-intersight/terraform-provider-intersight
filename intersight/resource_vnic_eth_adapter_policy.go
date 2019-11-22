package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
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
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"enabled": {
							Description: "Status of Accelerated Receive Flow Steering on the virtual ethernet interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"completion_queue_settings": {
				Description: "Completion Queue resource settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"count": {
							Description: "The number of completion queue resources to allocate. In general, the number of completion queue resources you should allocate is equal to the number of transmit queue resources plus the number of receive queue resources.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each completion queue.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"coalescing_time": {
							Description: "The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"coalescing_type": {
							Description: "Interrupt Coalescing Type. This can be one of the following:- MIN  — The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE — The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.",
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
							Description: "Preferred driver interrupt mode. This can be one of the following:- MSIx — Message Signaled Interrupts (MSI) with the optional extension. MSI   — MSI only. INTx  — PCI INTx interrupts. MSIx is the recommended option.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "MSIx",
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"enabled": {
							Description: "Status of the Network Virtualization using Generic Routing Encapsulation on the virtual ethernet interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
			"roce_settings": {
				Description: "Settings for RDMA over Converged Ethernet.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each queue.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"tcp_offload_settings": {
				Description: "The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
			"tx_queue_settings": {
				Description: "Transmit Queue resource settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each queue.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"vxlan_settings": {
				Description: "Virtual Extensible LAN Protocol Settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"enabled": {
							Description: "Status of the Virtual Extensible LAN Protocol on the virtual ethernet interface.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
		},
	}
}
func resourceVnicEthAdapterPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.VnicEthAdapterPolicy
	if v, ok := d.GetOkExists("advanced_filter"); ok {
		x := v.(bool)
		o.AdvancedFilter = &x
	}

	if v, ok := d.GetOk("arfs_settings"); ok {
		p := models.VnicArfsSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicArfsSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicArfsSettingsAO0P0.VnicArfsSettingsAO0P0 = x1.(map[string]interface{})
					}
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

			p = o
		}
		x := p
		o.ArfsSettings = &x

	}

	if v, ok := d.GetOk("completion_queue_settings"); ok {
		p := models.VnicCompletionQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicCompletionQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicCompletionQueueSettingsAO0P0.VnicCompletionQueueSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.RingSize = x
				}
			}

			p = o
		}
		x := p
		o.CompletionQueueSettings = &x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("interrupt_settings"); ok {
		p := models.VnicEthInterruptSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicEthInterruptSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicEthInterruptSettingsAO0P0.VnicEthInterruptSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["coalescing_time"]; ok {
				{
					x := int64(v.(int))
					o.CoalescingTime = x
				}
			}
			if v, ok := l["coalescing_type"]; ok {
				{
					x := (v.(string))
					o.CoalescingType = &x
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.Mode = &x
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
		o.InterruptSettings = &x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x

	}

	if v, ok := d.GetOk("nvgre_settings"); ok {
		p := models.VnicNvgreSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicNvgreSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicNvgreSettingsAO0P0.VnicNvgreSettingsAO0P0 = x1.(map[string]interface{})
					}
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

			p = o
		}
		x := p
		o.NvgreSettings = &x

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

	if v, ok := d.GetOk("roce_settings"); ok {
		p := models.VnicRoceSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicRoceSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicRoceSettingsAO0P0.VnicRoceSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.Enabled = &x
				}
			}
			if v, ok := l["memory_regions"]; ok {
				{
					x := int64(v.(int))
					o.MemoryRegions = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["queue_pairs"]; ok {
				{
					x := int64(v.(int))
					o.QueuePairs = x
				}
			}
			if v, ok := l["resource_groups"]; ok {
				{
					x := int64(v.(int))
					o.ResourceGroups = x
				}
			}

			p = o
		}
		x := p
		o.RoceSettings = &x

	}

	if v, ok := d.GetOkExists("rss_settings"); ok {
		x := v.(bool)
		o.RssSettings = &x
	}

	if v, ok := d.GetOk("rx_queue_settings"); ok {
		p := models.VnicEthRxQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicEthRxQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicEthRxQueueSettingsAO0P0.VnicEthRxQueueSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.RingSize = x
				}
			}

			p = o
		}
		x := p
		o.RxQueueSettings = &x

	}

	if v, ok := d.GetOk("tcp_offload_settings"); ok {
		p := models.VnicTCPOffloadSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicTCPOffloadSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicTCPOffloadSettingsAO0P0.VnicTCPOffloadSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["large_receive"]; ok {
				{
					x := (v.(bool))
					o.LargeReceive = &x
				}
			}
			if v, ok := l["large_send"]; ok {
				{
					x := (v.(bool))
					o.LargeSend = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["rx_checksum"]; ok {
				{
					x := (v.(bool))
					o.RxChecksum = &x
				}
			}
			if v, ok := l["tx_checksum"]; ok {
				{
					x := (v.(bool))
					o.TxChecksum = &x
				}
			}

			p = o
		}
		x := p
		o.TCPOffloadSettings = &x

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

	if v, ok := d.GetOk("tx_queue_settings"); ok {
		p := models.VnicEthTxQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicEthTxQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicEthTxQueueSettingsAO0P0.VnicEthTxQueueSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.RingSize = x
				}
			}

			p = o
		}
		x := p
		o.TxQueueSettings = &x

	}

	if v, ok := d.GetOk("vxlan_settings"); ok {
		p := models.VnicVxlanSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicVxlanSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicVxlanSettingsAO0P0.VnicVxlanSettingsAO0P0 = x1.(map[string]interface{})
					}
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

			p = o
		}
		x := p
		o.VxlanSettings = &x

	}

	url := "vnic/EthAdapterPolicies"
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
	return resourceVnicEthAdapterPolicyRead(d, meta)
}

func resourceVnicEthAdapterPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "vnic/EthAdapterPolicies" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.VnicEthAdapterPolicy
	err = s.UnmarshalJSON(body)
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

	if err := d.Set("organization", flattenMapIamAccountRef(s.Organization, d)); err != nil {
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

	if err := d.Set("tcp_offload_settings", flattenMapVnicTCPOffloadSettings(s.TCPOffloadSettings, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("tx_queue_settings", flattenMapVnicEthTxQueueSettings(s.TxQueueSettings, d)); err != nil {
		return err
	}

	if err := d.Set("vxlan_settings", flattenMapVnicVxlanSettings(s.VxlanSettings, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceVnicEthAdapterPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.VnicEthAdapterPolicy
	if d.HasChange("advanced_filter") {
		v := d.Get("advanced_filter")
		x := (v.(bool))
		o.AdvancedFilter = &x
	}

	if d.HasChange("arfs_settings") {
		v := d.Get("arfs_settings")
		p := models.VnicArfsSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicArfsSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicArfsSettingsAO0P0.VnicArfsSettingsAO0P0 = x1.(map[string]interface{})
					}
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

			p = o
		}
		x := p
		o.ArfsSettings = &x
	}

	if d.HasChange("completion_queue_settings") {
		v := d.Get("completion_queue_settings")
		p := models.VnicCompletionQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicCompletionQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicCompletionQueueSettingsAO0P0.VnicCompletionQueueSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.RingSize = x
				}
			}

			p = o
		}
		x := p
		o.CompletionQueueSettings = &x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
	}

	if d.HasChange("interrupt_settings") {
		v := d.Get("interrupt_settings")
		p := models.VnicEthInterruptSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicEthInterruptSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicEthInterruptSettingsAO0P0.VnicEthInterruptSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["coalescing_time"]; ok {
				{
					x := int64(v.(int))
					o.CoalescingTime = x
				}
			}
			if v, ok := l["coalescing_type"]; ok {
				{
					x := (v.(string))
					o.CoalescingType = &x
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.Mode = &x
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
		o.InterruptSettings = &x
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

	if d.HasChange("nvgre_settings") {
		v := d.Get("nvgre_settings")
		p := models.VnicNvgreSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicNvgreSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicNvgreSettingsAO0P0.VnicNvgreSettingsAO0P0 = x1.(map[string]interface{})
					}
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

			p = o
		}
		x := p
		o.NvgreSettings = &x
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

	if d.HasChange("roce_settings") {
		v := d.Get("roce_settings")
		p := models.VnicRoceSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicRoceSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicRoceSettingsAO0P0.VnicRoceSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.Enabled = &x
				}
			}
			if v, ok := l["memory_regions"]; ok {
				{
					x := int64(v.(int))
					o.MemoryRegions = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["queue_pairs"]; ok {
				{
					x := int64(v.(int))
					o.QueuePairs = x
				}
			}
			if v, ok := l["resource_groups"]; ok {
				{
					x := int64(v.(int))
					o.ResourceGroups = x
				}
			}

			p = o
		}
		x := p
		o.RoceSettings = &x
	}

	if d.HasChange("rss_settings") {
		v := d.Get("rss_settings")
		x := (v.(bool))
		o.RssSettings = &x
	}

	if d.HasChange("rx_queue_settings") {
		v := d.Get("rx_queue_settings")
		p := models.VnicEthRxQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicEthRxQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicEthRxQueueSettingsAO0P0.VnicEthRxQueueSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.RingSize = x
				}
			}

			p = o
		}
		x := p
		o.RxQueueSettings = &x
	}

	if d.HasChange("tcp_offload_settings") {
		v := d.Get("tcp_offload_settings")
		p := models.VnicTCPOffloadSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicTCPOffloadSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicTCPOffloadSettingsAO0P0.VnicTCPOffloadSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["large_receive"]; ok {
				{
					x := (v.(bool))
					o.LargeReceive = &x
				}
			}
			if v, ok := l["large_send"]; ok {
				{
					x := (v.(bool))
					o.LargeSend = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["rx_checksum"]; ok {
				{
					x := (v.(bool))
					o.RxChecksum = &x
				}
			}
			if v, ok := l["tx_checksum"]; ok {
				{
					x := (v.(bool))
					o.TxChecksum = &x
				}
			}

			p = o
		}
		x := p
		o.TCPOffloadSettings = &x
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

	if d.HasChange("tx_queue_settings") {
		v := d.Get("tx_queue_settings")
		p := models.VnicEthTxQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicEthTxQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicEthTxQueueSettingsAO0P0.VnicEthTxQueueSettingsAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.Count = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.RingSize = x
				}
			}

			p = o
		}
		x := p
		o.TxQueueSettings = &x
	}

	if d.HasChange("vxlan_settings") {
		v := d.Get("vxlan_settings")
		p := models.VnicVxlanSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicVxlanSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicVxlanSettingsAO0P0.VnicVxlanSettingsAO0P0 = x1.(map[string]interface{})
					}
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

			p = o
		}
		x := p
		o.VxlanSettings = &x
	}

	url := "vnic/EthAdapterPolicies" + "/" + d.Id()
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
	return resourceVnicEthAdapterPolicyRead(d, meta)
}

func resourceVnicEthAdapterPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "vnic/EthAdapterPolicies" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
