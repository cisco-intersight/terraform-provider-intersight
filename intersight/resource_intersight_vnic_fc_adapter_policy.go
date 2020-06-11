package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
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
							Computed:    true,
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
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Computed:    true,
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
			"plogi_settings": {
				Description: "Fibre Channel Plogi Settings.",
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
							Computed:    true,
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
							Computed:    true,
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
						"count": {
							Description: "The number of SCSI I/O queue resources the system should allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
			"tx_queue_settings": {
				Description: "Fibre Channel Transmit Queue Settings.",
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
							Computed:    true,
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
	var o models.VnicFcAdapterPolicy
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("error_detection_timeout"); ok {
		x := int64(v.(int))
		o.ErrorDetectionTimeout = x

	}

	if v, ok := d.GetOk("error_recovery_settings"); ok {
		p := models.VnicFcErrorRecoverySettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcErrorRecoverySettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.Enabled = &x
				}
			}
			if v, ok := l["io_retry_count"]; ok {
				{
					x := int64(v.(int))
					o.IoRetryCount = x
				}
			}
			if v, ok := l["io_retry_timeout"]; ok {
				{
					x := int64(v.(int))
					o.IoRetryTimeout = x
				}
			}
			if v, ok := l["link_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.LinkDownTimeout = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["port_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.PortDownTimeout = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ErrorRecoverySettings = &x
		}

	}

	if v, ok := d.GetOk("flogi_settings"); ok {
		p := models.VnicFlogiSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFlogiSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.Retries = x
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.Timeout = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.FlogiSettings = &x
		}

	}

	if v, ok := d.GetOk("interrupt_settings"); ok {
		p := models.VnicFcInterruptSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcInterruptSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.InterruptSettings = &x
		}

	}

	if v, ok := d.GetOk("io_throttle_count"); ok {
		x := int64(v.(int))
		o.IoThrottleCount = x

	}

	if v, ok := d.GetOk("lun_count"); ok {
		x := int64(v.(int))
		o.LunCount = x

	}

	if v, ok := d.GetOk("lun_queue_depth"); ok {
		x := int64(v.(int))
		o.LunQueueDepth = x

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

	if v, ok := d.GetOk("plogi_settings"); ok {
		p := models.VnicPlogiSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicPlogiSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.Retries = x
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.Timeout = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.PlogiSettings = &x
		}

	}

	if v, ok := d.GetOk("resource_allocation_timeout"); ok {
		x := int64(v.(int))
		o.ResourceAllocationTimeout = x

	}

	if v, ok := d.GetOk("rx_queue_settings"); ok {
		p := models.VnicFcQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.RxQueueSettings = &x
		}

	}

	if v, ok := d.GetOk("scsi_queue_settings"); ok {
		p := models.VnicScsiQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicScsiQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.ScsiQueueSettings = &x
		}

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

	if v, ok := d.GetOk("tx_queue_settings"); ok {
		p := models.VnicFcQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.TxQueueSettings = &x
		}

	}

	url := "vnic/FcAdapterPolicies"
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
	return resourceVnicFcAdapterPolicyRead(d, meta)
}

func resourceVnicFcAdapterPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "vnic/FcAdapterPolicies" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.VnicFcAdapterPolicy
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

	if err := d.Set("organization", flattenMapOrganizationOrganizationRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
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
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceVnicFcAdapterPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.VnicFcAdapterPolicy
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

	if d.HasChange("error_detection_timeout") {
		v := d.Get("error_detection_timeout")
		x := int64(v.(int))
		o.ErrorDetectionTimeout = x
	}

	if d.HasChange("error_recovery_settings") {
		v := d.Get("error_recovery_settings")
		p := models.VnicFcErrorRecoverySettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcErrorRecoverySettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.Enabled = &x
				}
			}
			if v, ok := l["io_retry_count"]; ok {
				{
					x := int64(v.(int))
					o.IoRetryCount = x
				}
			}
			if v, ok := l["io_retry_timeout"]; ok {
				{
					x := int64(v.(int))
					o.IoRetryTimeout = x
				}
			}
			if v, ok := l["link_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.LinkDownTimeout = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["port_down_timeout"]; ok {
				{
					x := int64(v.(int))
					o.PortDownTimeout = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ErrorRecoverySettings = &x
		}
	}

	if d.HasChange("flogi_settings") {
		v := d.Get("flogi_settings")
		p := models.VnicFlogiSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFlogiSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.Retries = x
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.Timeout = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.FlogiSettings = &x
		}
	}

	if d.HasChange("interrupt_settings") {
		v := d.Get("interrupt_settings")
		p := models.VnicFcInterruptSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcInterruptSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.InterruptSettings = &x
		}
	}

	if d.HasChange("io_throttle_count") {
		v := d.Get("io_throttle_count")
		x := int64(v.(int))
		o.IoThrottleCount = x
	}

	if d.HasChange("lun_count") {
		v := d.Get("lun_count")
		x := int64(v.(int))
		o.LunCount = x
	}

	if d.HasChange("lun_queue_depth") {
		v := d.Get("lun_queue_depth")
		x := int64(v.(int))
		o.LunQueueDepth = x
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

	if d.HasChange("plogi_settings") {
		v := d.Get("plogi_settings")
		p := models.VnicPlogiSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicPlogiSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["retries"]; ok {
				{
					x := int64(v.(int))
					o.Retries = x
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.Timeout = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.PlogiSettings = &x
		}
	}

	if d.HasChange("resource_allocation_timeout") {
		v := d.Get("resource_allocation_timeout")
		x := int64(v.(int))
		o.ResourceAllocationTimeout = x
	}

	if d.HasChange("rx_queue_settings") {
		v := d.Get("rx_queue_settings")
		p := models.VnicFcQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.RxQueueSettings = &x
		}
	}

	if d.HasChange("scsi_queue_settings") {
		v := d.Get("scsi_queue_settings")
		p := models.VnicScsiQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicScsiQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.ScsiQueueSettings = &x
		}
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

	if d.HasChange("tx_queue_settings") {
		v := d.Get("tx_queue_settings")
		p := models.VnicFcQueueSettings{}
		if len(v.([]interface{})) > 0 {
			o := models.VnicFcQueueSettings{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
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
		if len(v.([]interface{})) > 0 {
			o.TxQueueSettings = &x
		}
	}

	url := "vnic/FcAdapterPolicies" + "/" + d.Id()
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
	return resourceVnicFcAdapterPolicyRead(d, meta)
}

func resourceVnicFcAdapterPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "vnic/FcAdapterPolicies" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
