package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVnicFcAdapterPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVnicFcAdapterPolicyCreate,
		Read:   resourceVnicFcAdapterPolicyRead,
		Update: resourceVnicFcAdapterPolicyUpdate,
		Delete: resourceVnicFcAdapterPolicyDelete,
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
			"error_detection_timeout": {
				Description: "Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"error_recovery_settings": {
				Description: "Fibre Channel Error Recovery Settings.",
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
							Description: "Enables Fibre Channel Error recovery.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"io_retry_count": {
							Description: "The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"io_retry_timeout": {
							Description: "The number of seconds the adapter waits before aborting the pending command and resending the same IO request.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"link_down_timeout": {
							Description: "The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"port_down_timeout": {
							Description: "The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"flogi_settings": {
				Description: "Fibre Channel Flogi Settings.",
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
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"retries": {
							Description: "The number of times that the system tries to log in to the fabric after the first failure.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"timeout": {
							Description: "The number of milliseconds that the system waits before it tries to log in again.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"interrupt_settings": {
				Description: "Interrupt Settings for the virtual fibre channel interface.",
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
						"mode": {
							Description: "The preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.",
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
			"io_throttle_count": {
				Description: "The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"lun_count": {
				Description: "The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"lun_queue_depth": {
				Description: "The number of commands that the HBA can send and receive in a single transmission per LUN.",
				Type:        schema.TypeInt,
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
			"plogi_settings": {
				Description: "Fibre Channel Plogi Settings.",
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
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"retries": {
							Description: "The number of times that the system tries to log in to a port after the first failure.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"timeout": {
							Description: "The number of milliseconds that the system waits before it tries to log in again.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"resource_allocation_timeout": {
				Description: "Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"rx_queue_settings": {
				Description: "Fibre Channel Receive Queue Settings.",
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
							Computed:    true,
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
			"scsi_queue_settings": {
				Description: "SCSI Input/Output Queue Settings.",
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
							Description: "The number of SCSI I/O queue resources the system should allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ring_size": {
							Description: "The number of descriptors in each SCSI I/O queue.",
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
			"tx_queue_settings": {
				Description: "Fibre Channel Transmit Queue Settings.",
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
							Computed:    true,
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
		},
	}
}

func resourceVnicFcAdapterPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewVnicFcAdapterPolicy()
	o.SetClassId("vnic.FcAdapterPolicy")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("error_detection_timeout"); ok {
		x := int64(v.(int))
		o.SetErrorDetectionTimeout(x)
	}

	if v, ok := d.GetOk("error_recovery_settings"); ok {
		p := make([]models.VnicFcErrorRecoverySettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcErrorRecoverySettingsWithDefaults()
			o.SetClassId("vnic.FcErrorRecoverySettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["io_retry_count"]; ok {
				{
					x := int64(v.(int))
					o.SetIoRetryCount(x)
				}
			}
			if v, ok := l["io_retry_timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetIoRetryTimeout(x)
				}
			}
			if v, ok := l["link_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetLinkDownTimeout(x)
				}
			}
			o.SetObjectType("vnic.FcErrorRecoverySettings")
			if v, ok := l["port_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetPortDownTimeout(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetErrorRecoverySettings(x)
	}

	if v, ok := d.GetOk("flogi_settings"); ok {
		p := make([]models.VnicFlogiSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFlogiSettingsWithDefaults()
			o.SetClassId("vnic.FlogiSettings")
			o.SetObjectType("vnic.FlogiSettings")
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.SetRetries(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetFlogiSettings(x)
	}

	if v, ok := d.GetOk("interrupt_settings"); ok {
		p := make([]models.VnicFcInterruptSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcInterruptSettingsWithDefaults()
			o.SetClassId("vnic.FcInterruptSettings")
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.SetMode(x)
				}
			}
			o.SetObjectType("vnic.FcInterruptSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetInterruptSettings(x)
	}

	if v, ok := d.GetOk("io_throttle_count"); ok {
		x := int64(v.(int))
		o.SetIoThrottleCount(x)
	}

	if v, ok := d.GetOk("lun_count"); ok {
		x := int64(v.(int))
		o.SetLunCount(x)
	}

	if v, ok := d.GetOk("lun_queue_depth"); ok {
		x := int64(v.(int))
		o.SetLunQueueDepth(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("vnic.FcAdapterPolicy")

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

	if v, ok := d.GetOk("plogi_settings"); ok {
		p := make([]models.VnicPlogiSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicPlogiSettingsWithDefaults()
			o.SetClassId("vnic.PlogiSettings")
			o.SetObjectType("vnic.PlogiSettings")
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.SetRetries(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetPlogiSettings(x)
	}

	if v, ok := d.GetOk("resource_allocation_timeout"); ok {
		x := int64(v.(int))
		o.SetResourceAllocationTimeout(x)
	}

	if v, ok := d.GetOk("rx_queue_settings"); ok {
		p := make([]models.VnicFcQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcQueueSettingsWithDefaults()
			o.SetClassId("vnic.FcQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.FcQueueSettings")
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

	if v, ok := d.GetOk("scsi_queue_settings"); ok {
		p := make([]models.VnicScsiQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicScsiQueueSettingsWithDefaults()
			o.SetClassId("vnic.ScsiQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.ScsiQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetScsiQueueSettings(x)
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

	if v, ok := d.GetOk("tx_queue_settings"); ok {
		p := make([]models.VnicFcQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcQueueSettingsWithDefaults()
			o.SetClassId("vnic.FcQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.FcQueueSettings")
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

	r := conn.ApiClient.VnicApi.CreateVnicFcAdapterPolicy(conn.ctx).VnicFcAdapterPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceVnicFcAdapterPolicyRead(d, meta)
}

func resourceVnicFcAdapterPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.VnicApi.GetVnicFcAdapterPolicyByMoid(conn.ctx, d.Id())
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

	if err := d.Set("error_detection_timeout", (s.ErrorDetectionTimeout)); err != nil {
		return err
	}

	if err := d.Set("error_recovery_settings", flattenMapVnicFcErrorRecoverySettings(s.ErrorRecoverySettings, d)); err != nil {
		return err
	}

	if err := d.Set("flogi_settings", flattenMapVnicFlogiSettings(s.FlogiSettings, d)); err != nil {
		return err
	}

	if err := d.Set("interrupt_settings", flattenMapVnicFcInterruptSettings(s.InterruptSettings, d)); err != nil {
		return err
	}

	if err := d.Set("io_throttle_count", (s.IoThrottleCount)); err != nil {
		return err
	}

	if err := d.Set("lun_count", (s.LunCount)); err != nil {
		return err
	}

	if err := d.Set("lun_queue_depth", (s.LunQueueDepth)); err != nil {
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

	if err := d.Set("plogi_settings", flattenMapVnicPlogiSettings(s.PlogiSettings, d)); err != nil {
		return err
	}

	if err := d.Set("resource_allocation_timeout", (s.ResourceAllocationTimeout)); err != nil {
		return err
	}

	if err := d.Set("rx_queue_settings", flattenMapVnicFcQueueSettings(s.RxQueueSettings, d)); err != nil {
		return err
	}

	if err := d.Set("scsi_queue_settings", flattenMapVnicScsiQueueSettings(s.ScsiQueueSettings, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("tx_queue_settings", flattenMapVnicFcQueueSettings(s.TxQueueSettings, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceVnicFcAdapterPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewVnicFcAdapterPolicy()

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("error_detection_timeout") {
		v := d.Get("error_detection_timeout")
		x := int64(v.(int))
		o.SetErrorDetectionTimeout(x)
	}

	if d.HasChange("error_recovery_settings") {
		v := d.Get("error_recovery_settings")
		p := make([]models.VnicFcErrorRecoverySettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcErrorRecoverySettingsWithDefaults()
			o.SetClassId("vnic.FcErrorRecoverySettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["io_retry_count"]; ok {
				{
					x := int64(v.(int))
					o.SetIoRetryCount(x)
				}
			}
			if v, ok := l["io_retry_timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetIoRetryTimeout(x)
				}
			}
			if v, ok := l["link_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetLinkDownTimeout(x)
				}
			}
			o.SetObjectType("vnic.FcErrorRecoverySettings")
			if v, ok := l["port_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetPortDownTimeout(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetErrorRecoverySettings(x)
	}

	if d.HasChange("flogi_settings") {
		v := d.Get("flogi_settings")
		p := make([]models.VnicFlogiSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFlogiSettingsWithDefaults()
			o.SetClassId("vnic.FlogiSettings")
			o.SetObjectType("vnic.FlogiSettings")
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.SetRetries(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetFlogiSettings(x)
	}

	if d.HasChange("interrupt_settings") {
		v := d.Get("interrupt_settings")
		p := make([]models.VnicFcInterruptSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcInterruptSettingsWithDefaults()
			o.SetClassId("vnic.FcInterruptSettings")
			if v, ok := l["mode"]; ok {
				{
					x := (v.(string))
					o.SetMode(x)
				}
			}
			o.SetObjectType("vnic.FcInterruptSettings")
			p = append(p, *o)
		}
		x := p[0]
		o.SetInterruptSettings(x)
	}

	if d.HasChange("io_throttle_count") {
		v := d.Get("io_throttle_count")
		x := int64(v.(int))
		o.SetIoThrottleCount(x)
	}

	if d.HasChange("lun_count") {
		v := d.Get("lun_count")
		x := int64(v.(int))
		o.SetLunCount(x)
	}

	if d.HasChange("lun_queue_depth") {
		v := d.Get("lun_queue_depth")
		x := int64(v.(int))
		o.SetLunQueueDepth(x)
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

	if d.HasChange("plogi_settings") {
		v := d.Get("plogi_settings")
		p := make([]models.VnicPlogiSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicPlogiSettingsWithDefaults()
			o.SetClassId("vnic.PlogiSettings")
			o.SetObjectType("vnic.PlogiSettings")
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.SetRetries(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetPlogiSettings(x)
	}

	if d.HasChange("resource_allocation_timeout") {
		v := d.Get("resource_allocation_timeout")
		x := int64(v.(int))
		o.SetResourceAllocationTimeout(x)
	}

	if d.HasChange("rx_queue_settings") {
		v := d.Get("rx_queue_settings")
		p := make([]models.VnicFcQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcQueueSettingsWithDefaults()
			o.SetClassId("vnic.FcQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.FcQueueSettings")
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

	if d.HasChange("scsi_queue_settings") {
		v := d.Get("scsi_queue_settings")
		p := make([]models.VnicScsiQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicScsiQueueSettingsWithDefaults()
			o.SetClassId("vnic.ScsiQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.ScsiQueueSettings")
			if v, ok := l["ring_size"]; ok {
				{
					x := int64(v.(int))
					o.SetRingSize(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetScsiQueueSettings(x)
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

	if d.HasChange("tx_queue_settings") {
		v := d.Get("tx_queue_settings")
		p := make([]models.VnicFcQueueSettings, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewVnicFcQueueSettingsWithDefaults()
			o.SetClassId("vnic.FcQueueSettings")
			if v, ok := l["count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			o.SetObjectType("vnic.FcQueueSettings")
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

	r := conn.ApiClient.VnicApi.UpdateVnicFcAdapterPolicy(conn.ctx, d.Id()).VnicFcAdapterPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceVnicFcAdapterPolicyRead(d, meta)
}

func resourceVnicFcAdapterPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.VnicApi.DeleteVnicFcAdapterPolicy(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
