package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMemoryPersistentMemoryUnit() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMemoryPersistentMemoryUnitRead,
		Schema: map[string]*schema.Schema{
			"admin_state": {
				Description: "This represents the administrative state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"app_direct_capacity": {
				Description: "This represents the appdirect capacity in GB of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"array_id": {
				Description: "This represents the memory array to which the memory unit belongs to.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"bank": {
				Description: "This represents the memory bank of the memory unit on a server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"capacity": {
				Description: "This represents the memory capacity in MiB of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"clock": {
				Description: "This represents the clock of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"count_status": {
				Description: "This represents the count status of a persistent memory module on a server.",
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
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"firmware_version": {
				Description: "This represents the firmware version of the firware running on a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"form_factor": {
				Description: "This represents the form factor of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"frozen_status": {
				Description: "This represents the frozen status of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_state": {
				Description: "This represents the health state of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"latency": {
				Description: "This represents the latency of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Description: "This represents the location of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"lock_status": {
				Description: "This represents the lock status of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"memory_array": {
				Description: "A collection of references to the [memory.Array](mo://memory.Array) Managed Object.When this managed object is deleted, the referenced [memory.Array](mo://memory.Array) MO unsets its reference to this deleted MO.",
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
							Description: "The Object Type of the referenced REST resource.",
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
			"memory_capacity": {
				Description: "This represents the memory capacity in GB of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"memory_id": {
				Description: "This represents the ID of a persistent memory module on a server.",
				Type:        schema.TypeInt,
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
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_power_state": {
				Description: "This represents the operational power state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_state": {
				Description: "This represents the operational state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "This represents the operability of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"persistent_memory_capacity": {
				Description: "This represents the persistent memory capacity in GB of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "This represents the presence state of the memory unit on a server.",
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
							Description: "The Object Type of the referenced REST resource.",
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
			"reserved_capacity": {
				Description: "This represents the reserved capacity in GB of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"revision": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"security_status": {
				Description: "This represents the security status of a persistent memory module on a server.",
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
			"set": {
				Description: "This represents the set of the memory unit on a server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"socket_id": {
				Description: "This represents the Socket ID of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"socket_memory_id": {
				Description: "This represents the Socket Memory ID of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"speed": {
				Description: "This represents the speed of the memory unit on a server.",
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
			"thermal": {
				Description: "This represents the thremal state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"total_capacity": {
				Description: "This represents the total capacity in GB of a persistent memory module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "This represents the memory type of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uid": {
				Description: "This represents the uid of a persistent memory module on a server.",
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
			"visibility": {
				Description: "This represents the visibility of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"width": {
				Description: "This represents the width of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceMemoryPersistentMemoryUnitRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "memory/PersistentMemoryUnits"
	var o models.MemoryPersistentMemoryUnit
	if v, ok := d.GetOk("admin_state"); ok {
		x := (v.(string))
		o.AdminState = x
	}
	if v, ok := d.GetOk("app_direct_capacity"); ok {
		x := (v.(string))
		o.AppDirectCapacity = x
	}
	if v, ok := d.GetOk("array_id"); ok {
		x := int64(v.(int))
		o.ArrayID = x
	}
	if v, ok := d.GetOk("bank"); ok {
		x := int64(v.(int))
		o.Bank = x
	}
	if v, ok := d.GetOk("capacity"); ok {
		x := (v.(string))
		o.Capacity = x
	}
	if v, ok := d.GetOk("clock"); ok {
		x := (v.(string))
		o.Clock = x
	}
	if v, ok := d.GetOk("count_status"); ok {
		x := (v.(string))
		o.CountStatus = x
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.DeviceMoID = x
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.Dn = x
	}
	if v, ok := d.GetOk("firmware_version"); ok {
		x := (v.(string))
		o.FirmwareVersion = x
	}
	if v, ok := d.GetOk("form_factor"); ok {
		x := (v.(string))
		o.FormFactor = x
	}
	if v, ok := d.GetOk("frozen_status"); ok {
		x := (v.(string))
		o.FrozenStatus = x
	}
	if v, ok := d.GetOk("health_state"); ok {
		x := (v.(string))
		o.HealthState = x
	}
	if v, ok := d.GetOk("latency"); ok {
		x := (v.(string))
		o.Latency = x
	}
	if v, ok := d.GetOk("location"); ok {
		x := (v.(string))
		o.Location = x
	}
	if v, ok := d.GetOk("lock_status"); ok {
		x := (v.(string))
		o.LockStatus = x
	}
	if v, ok := d.GetOk("memory_capacity"); ok {
		x := (v.(string))
		o.MemoryCapacity = x
	}
	if v, ok := d.GetOk("memory_id"); ok {
		x := int64(v.(int))
		o.MemoryID = x
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.Model = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
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
	if v, ok := d.GetOk("persistent_memory_capacity"); ok {
		x := (v.(string))
		o.PersistentMemoryCapacity = x
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.Presence = x
	}
	if v, ok := d.GetOk("reserved_capacity"); ok {
		x := (v.(string))
		o.ReservedCapacity = x
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.Revision = x
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.Rn = x
	}
	if v, ok := d.GetOk("security_status"); ok {
		x := (v.(string))
		o.SecurityStatus = x
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.Serial = x
	}
	if v, ok := d.GetOk("set"); ok {
		x := int64(v.(int))
		o.Set = x
	}
	if v, ok := d.GetOk("socket_id"); ok {
		x := (v.(string))
		o.SocketID = x
	}
	if v, ok := d.GetOk("socket_memory_id"); ok {
		x := (v.(string))
		o.SocketMemoryID = x
	}
	if v, ok := d.GetOk("speed"); ok {
		x := (v.(string))
		o.Speed = x
	}
	if v, ok := d.GetOk("thermal"); ok {
		x := (v.(string))
		o.Thermal = x
	}
	if v, ok := d.GetOk("total_capacity"); ok {
		x := (v.(string))
		o.TotalCapacity = x
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.Type = x
	}
	if v, ok := d.GetOk("uid"); ok {
		x := (v.(string))
		o.UID = x
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x
	}
	if v, ok := d.GetOk("visibility"); ok {
		x := (v.(string))
		o.Visibility = x
	}
	if v, ok := d.GetOk("width"); ok {
		x := (v.(string))
		o.Width = x
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
			var s models.MemoryPersistentMemoryUnit
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("admin_state", (s.AdminState)); err != nil {
				return err
			}
			if err := d.Set("app_direct_capacity", (s.AppDirectCapacity)); err != nil {
				return err
			}
			if err := d.Set("array_id", (s.ArrayID)); err != nil {
				return err
			}
			if err := d.Set("bank", (s.Bank)); err != nil {
				return err
			}
			if err := d.Set("capacity", (s.Capacity)); err != nil {
				return err
			}
			if err := d.Set("clock", (s.Clock)); err != nil {
				return err
			}
			if err := d.Set("count_status", (s.CountStatus)); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoID)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}
			if err := d.Set("firmware_version", (s.FirmwareVersion)); err != nil {
				return err
			}
			if err := d.Set("form_factor", (s.FormFactor)); err != nil {
				return err
			}
			if err := d.Set("frozen_status", (s.FrozenStatus)); err != nil {
				return err
			}
			if err := d.Set("health_state", (s.HealthState)); err != nil {
				return err
			}
			if err := d.Set("latency", (s.Latency)); err != nil {
				return err
			}
			if err := d.Set("location", (s.Location)); err != nil {
				return err
			}
			if err := d.Set("lock_status", (s.LockStatus)); err != nil {
				return err
			}

			if err := d.Set("memory_array", flattenMapMemoryArrayRef(s.MemoryArray, d)); err != nil {
				return err
			}
			if err := d.Set("memory_capacity", (s.MemoryCapacity)); err != nil {
				return err
			}
			if err := d.Set("memory_id", (s.MemoryID)); err != nil {
				return err
			}
			if err := d.Set("model", (s.Model)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
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

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("persistent_memory_capacity", (s.PersistentMemoryCapacity)); err != nil {
				return err
			}
			if err := d.Set("presence", (s.Presence)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("reserved_capacity", (s.ReservedCapacity)); err != nil {
				return err
			}
			if err := d.Set("revision", (s.Revision)); err != nil {
				return err
			}
			if err := d.Set("rn", (s.Rn)); err != nil {
				return err
			}
			if err := d.Set("security_status", (s.SecurityStatus)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}
			if err := d.Set("set", (s.Set)); err != nil {
				return err
			}
			if err := d.Set("socket_id", (s.SocketID)); err != nil {
				return err
			}
			if err := d.Set("socket_memory_id", (s.SocketMemoryID)); err != nil {
				return err
			}
			if err := d.Set("speed", (s.Speed)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("thermal", (s.Thermal)); err != nil {
				return err
			}
			if err := d.Set("total_capacity", (s.TotalCapacity)); err != nil {
				return err
			}
			if err := d.Set("type", (s.Type)); err != nil {
				return err
			}
			if err := d.Set("uid", (s.UID)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			if err := d.Set("visibility", (s.Visibility)); err != nil {
				return err
			}
			if err := d.Set("width", (s.Width)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
