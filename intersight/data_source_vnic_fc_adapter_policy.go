package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceVnicFcAdapterPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVnicFcAdapterPolicyRead,
		Schema: map[string]*schema.Schema{
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				Computed: true,
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
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				Computed: true,
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
						"mode": {
							Description: "The preferred driver interrupt mode. This can be one of the following:- MSIx — Message Signaled Interrupts (MSI) with the optional extension. MSI   — MSI only. INTx  — PCI INTx interrupts. MSIx is the recommended option.",
							Type:        schema.TypeString,
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
				Computed: true,
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
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Computed: true,
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
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				Computed: true,
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
						"count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
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
				Computed: true,
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
						"count": {
							Description: "The number of SCSI I/O queue resources the system should allocate.",
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
							Description: "The number of descriptors in each SCSI I/O queue.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
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
				Computed: true,
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
						"count": {
							Description: "The number of queue resources to allocate.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
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
				Computed: true,
			},
		},
	}
}
func dataSourceVnicFcAdapterPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "vnic/FcAdapterPolicies"
	var o models.VnicFcAdapterPolicy
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("error_detection_timeout"); ok {
		x := int64(v.(int))
		o.ErrorDetectionTimeout = x
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
	if v, ok := d.GetOk("resource_allocation_timeout"); ok {
		x := int64(v.(int))
		o.ResourceAllocationTimeout = x
	}

	data, err := o.MarshalJSON()
	body, err := conn.SendGetRequest(url, data)
	if err != nil {
		return err
	}
	var x = make(map[string]interface{})
	if err = json.Unmarshal(body, &x); err != nil {
		return err
	}
	result := x["Results"]
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s models.VnicFcAdapterPolicy
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
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

			if err := d.Set("organization", flattenMapIamAccountRef(s.Organization, d)); err != nil {
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
			d.SetId(s.Moid)
		}
	}
	return nil
}
