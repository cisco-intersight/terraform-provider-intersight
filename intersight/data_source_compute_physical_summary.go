package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceComputePhysicalSummary() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComputePhysicalSummaryRead,
		Schema: map[string]*schema.Schema{
			"admin_power_state": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"asset_tag": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"available_memory": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"cpu_capacity": {
				Description: "CPU Capacity = Number of CPU Sockets x Enabled Cores x Speed (GHz).",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"chassis_id": {
				Description: "The id of the chassis that the blade is located in.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_mo_id": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fault_summary": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"firmware": {
				Description: "The firmware version of the Cisco Integrated Management Controller (CIMC) for this server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_address": {
				Description: "The IPv4 address configured on the management interface of the Integrated Management Controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"kvm_ip_addresses": {
				Description: "",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"address": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"category": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"default_gateway": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"dn": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"http_port": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"https_port": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"kvm_port": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "",
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
						"subnet": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"memory_speed": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mgmt_ip_address": {
				Description: "The IP address of the management interface on the UCS Fabric Interconnect or Cisco Integrated Management Controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"model": {
				Description: "This field identifies the model of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC).When this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect.When this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_adaptors": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_cpu_cores": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_cpu_cores_enabled": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_cpus": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_eth_host_interfaces": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_fc_host_interfaces": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_threads": {
				Description: "",
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
			"oper_power_state": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_state": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"platform_type": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "The Device to which this Managed Object is associated.",
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
			},
			"revision": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rn": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"scaled_mode": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"server_id": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"service_profile": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"slot_id": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"source_object_type": {
				Description: "Specifies the source object type for View MO.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"total_memory": {
				Description: "",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"user_label": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceComputePhysicalSummaryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "compute/PhysicalSummaries"
	var o models.ComputePhysicalSummary
	if v, ok := d.GetOk("admin_power_state"); ok {
		x := (v.(string))
		o.AdminPowerState = x
	}
	if v, ok := d.GetOk("asset_tag"); ok {
		x := (v.(string))
		o.AssetTag = x
	}
	if v, ok := d.GetOk("available_memory"); ok {
		x := int64(v.(int))
		o.AvailableMemory = x
	}
	if v, ok := d.GetOk("cpu_capacity"); ok {
		x := v.(float32)
		o.CPUCapacity = x
	}
	if v, ok := d.GetOk("chassis_id"); ok {
		x := (v.(string))
		o.ChassisID = x
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.DeviceMoID = x
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.Dn = x
	}
	if v, ok := d.GetOk("fault_summary"); ok {
		x := int64(v.(int))
		o.FaultSummary = x
	}
	if v, ok := d.GetOk("firmware"); ok {
		x := (v.(string))
		o.Firmware = x
	}
	if v, ok := d.GetOk("ipv4_address"); ok {
		x := (v.(string))
		o.IPV4Address = x
	}
	if v, ok := d.GetOk("memory_speed"); ok {
		x := (v.(string))
		o.MemorySpeed = x
	}
	if v, ok := d.GetOk("mgmt_ip_address"); ok {
		x := (v.(string))
		o.MgmtIPAddress = x
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.Model = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x
	}
	if v, ok := d.GetOk("num_adaptors"); ok {
		x := int64(v.(int))
		o.NumAdaptors = x
	}
	if v, ok := d.GetOk("num_cpu_cores"); ok {
		x := int64(v.(int))
		o.NumCPUCores = x
	}
	if v, ok := d.GetOk("num_cpu_cores_enabled"); ok {
		x := int64(v.(int))
		o.NumCPUCoresEnabled = x
	}
	if v, ok := d.GetOk("num_cpus"); ok {
		x := int64(v.(int))
		o.NumCpus = x
	}
	if v, ok := d.GetOk("num_eth_host_interfaces"); ok {
		x := int64(v.(int))
		o.NumEthHostInterfaces = x
	}
	if v, ok := d.GetOk("num_fc_host_interfaces"); ok {
		x := int64(v.(int))
		o.NumFcHostInterfaces = x
	}
	if v, ok := d.GetOk("num_threads"); ok {
		x := int64(v.(int))
		o.NumThreads = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("oper_power_state"); ok {
		x := (v.(string))
		o.OperPowerState = x
	}
	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.OperState = x
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.Operability = x
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.PlatformType = x
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.Presence = x
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.Revision = x
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.Rn = x
	}
	if v, ok := d.GetOk("scaled_mode"); ok {
		x := (v.(string))
		o.ScaledMode = x
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.Serial = x
	}
	if v, ok := d.GetOk("server_id"); ok {
		x := int64(v.(int))
		o.ServerID = x
	}
	if v, ok := d.GetOk("service_profile"); ok {
		x := (v.(string))
		o.ServiceProfile = x
	}
	if v, ok := d.GetOk("slot_id"); ok {
		x := int64(v.(int))
		o.SlotID = x
	}
	if v, ok := d.GetOk("source_object_type"); ok {
		x := (v.(string))
		o.SourceObjectType = x
	}
	if v, ok := d.GetOk("total_memory"); ok {
		x := int64(v.(int))
		o.TotalMemory = x
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.UUID = x
	}
	if v, ok := d.GetOk("user_label"); ok {
		x := (v.(string))
		o.UserLabel = x
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x
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
			var s models.ComputePhysicalSummary
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("admin_power_state", (s.AdminPowerState)); err != nil {
				return err
			}
			if err := d.Set("asset_tag", (s.AssetTag)); err != nil {
				return err
			}
			if err := d.Set("available_memory", (s.AvailableMemory)); err != nil {
				return err
			}
			if err := d.Set("cpu_capacity", (s.CPUCapacity)); err != nil {
				return err
			}
			if err := d.Set("chassis_id", (s.ChassisID)); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoID)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}
			if err := d.Set("fault_summary", (s.FaultSummary)); err != nil {
				return err
			}
			if err := d.Set("firmware", (s.Firmware)); err != nil {
				return err
			}
			if err := d.Set("ipv4_address", (s.IPV4Address)); err != nil {
				return err
			}

			if err := d.Set("kvm_ip_addresses", flattenListComputeIPAddress(s.KvmIPAddresses, d)); err != nil {
				return err
			}
			if err := d.Set("memory_speed", (s.MemorySpeed)); err != nil {
				return err
			}
			if err := d.Set("mgmt_ip_address", (s.MgmtIPAddress)); err != nil {
				return err
			}
			if err := d.Set("model", (s.Model)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("num_adaptors", (s.NumAdaptors)); err != nil {
				return err
			}
			if err := d.Set("num_cpu_cores", (s.NumCPUCores)); err != nil {
				return err
			}
			if err := d.Set("num_cpu_cores_enabled", (s.NumCPUCoresEnabled)); err != nil {
				return err
			}
			if err := d.Set("num_cpus", (s.NumCpus)); err != nil {
				return err
			}
			if err := d.Set("num_eth_host_interfaces", (s.NumEthHostInterfaces)); err != nil {
				return err
			}
			if err := d.Set("num_fc_host_interfaces", (s.NumFcHostInterfaces)); err != nil {
				return err
			}
			if err := d.Set("num_threads", (s.NumThreads)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("oper_power_state", (s.OperPowerState)); err != nil {
				return err
			}
			if err := d.Set("oper_state", (s.OperState)); err != nil {
				return err
			}
			if err := d.Set("operability", (s.Operability)); err != nil {
				return err
			}
			if err := d.Set("platform_type", (s.PlatformType)); err != nil {
				return err
			}
			if err := d.Set("presence", (s.Presence)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("revision", (s.Revision)); err != nil {
				return err
			}
			if err := d.Set("rn", (s.Rn)); err != nil {
				return err
			}
			if err := d.Set("scaled_mode", (s.ScaledMode)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}
			if err := d.Set("server_id", (s.ServerID)); err != nil {
				return err
			}
			if err := d.Set("service_profile", (s.ServiceProfile)); err != nil {
				return err
			}
			if err := d.Set("slot_id", (s.SlotID)); err != nil {
				return err
			}
			if err := d.Set("source_object_type", (s.SourceObjectType)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("total_memory", (s.TotalMemory)); err != nil {
				return err
			}
			if err := d.Set("uuid", (s.UUID)); err != nil {
				return err
			}
			if err := d.Set("user_label", (s.UserLabel)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
