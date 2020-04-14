package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceVirtualizationVmwareVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVirtualizationVmwareVirtualMachineRead,
		Schema: map[string]*schema.Schema{
			"annotation": {
				Description: "List of annotations provided to this VM by user. Can be long.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_time": {
				Description: "Time when this VM booted up.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_hot_add_enabled": {
				Description: "Indicates if the capability to add CPUs to a running VM is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"cpu_shares": {
				Description: "Shows the relative importance of a VM and its CPU limits.",
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
						"cpu_limit": {
							Description: "Upper limit on CPU allocation (MHz).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_overhead_limit": {
							Description: "Amount of CPU for virtualization overhead.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_reservation": {
							Description: "Guaranteed minimum allocation of CPU resource (MHz).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_shares": {
							Description: "Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation.",
							Type:        schema.TypeInt,
							Optional:    true,
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
					},
				},
				Computed: true,
			},
			"cpu_socket_info": {
				Description: "Details of CPUs/sockets of this VM.",
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
						"cores_per_socket": {
							Description: "The number of core per CPU socket (value may be empty).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"num_cpus": {
							Description: "Number of CPUs allocated to this VM.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"num_sockets": {
							Description: "The number of CPU sockets allocated.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"capacity": {
				Description: "Provisioned CPU and memory information for this virtual machine.",
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
						"cpu_cores": {
							Description: "The number of CPU cores on this hardware platform.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_speed": {
							Description: "Speed of CPU in MHz, as reported by the hardware platform.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"memory_size": {
							Description: "The amount of memory allocated (bytes) to this hardware platform.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster": {
				Description: "If the virtual machine is running on a cluster, it provides associated cluster information.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"config_name": {
				Description: "The configuration name for this VM. This maybe the same as the guest hostname.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connection_state": {
				Description: "Shows if virtual machine is connected to vCenter. Values are Connected, Disconnected, Orphaned, Inaccessible, and Invalid.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"custom_attributes": {
				Description: "User provided meta information associated with the VMs. Can be long.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"dns_server_list": {
				Description: "List of DNS server IPs assigned to this VM.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"dns_suffix_list": {
				Description: "List of DNS suffixes given to this VM.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"datacenter": {
				Description: "Every entity is grouped under the datacenter object and managed as a group.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"datastores": {
				Description: "The list of datastores that are attached to this VM. Used for the new inventory model and will soon replace dataStoreList above.",
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
			},
			"default_power_off_type": {
				Description: "Indicates how the VM will be powered off (soft, hard etc.).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dhcp_enabled": {
				Description: "Shows if DHCP is used for IP/DNS on this VM.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"disk_commit_info": {
				Description: "Information about the virtual machine's disk commits, sharing and limits.",
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
						"committed_disk": {
							Description: "Disk committed in bytes on this virtual machine (disk space used up).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"un_committed_disk": {
							Description: "Total uncommitted disk space that is available for use (in bytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"unshared_disk": {
							Description: "Total unshared disk space (in bytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"folder": {
				Description: "The folder name associated with this VM.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"guest_info": {
				Description: "Guest operating system details running on this machine.",
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
						"hostname": {
							Description: "Name provided to the host OS (example, ubuntu6410, test-gateway, etc.).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ip_address": {
							Description: "Primary IP address of the guest os.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "The name of the guest running on this VM. This may not be the same as the hostname.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"operating_system": {
							Description: "The name of the guest OS running on this VM (Cent OS 4/5/6/7).",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"guest_state": {
				Description: "The state of the guest OS running on this VM. Could be running, not running etc.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"host": {
				Description: "Host on which the Virtual Machine resides.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"hypervisor_type": {
				Description: "Type of hypervisor where the virtual machine is hosted, for example VMware ESXi.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ip_address": {
				Description: "The IP address of the virtual machine. There could be multiple addresses of IPv4 and IPv6 types.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"identity": {
				Description: "The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"instance_uuid": {
				Description: "UUID assigned by vCenter to every VM.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_template": {
				Description: "If true, indicates that the entity refers to a template of a virtual machine and not a real virtual machine.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"mac_address": {
				Description: "Standard MAC address assigned to this VM.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"mem_shares": {
				Description: "Similar to CPU Shares but applicable to memory.",
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
						"mem_limit": {
							Description: "Limit on the memory sharing imposed (in Mbytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mem_overhead_limit": {
							Description: "Limit on memory overhead imposed (in Mbytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mem_reservation": {
							Description: "Similar to CPU reservations (Mbytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mem_shares": {
							Description: "Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"memory_capacity": {
				Description: "The capacity and usage information for memory on this virtual machine.",
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
						"capacity": {
							Description: "The total memory capacity of the entity in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"free": {
							Description: "Free memory (bytes) that is unused and available for allocation, as a point-in-time snapshot. The available memory capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used memory capacity is also reported.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"used": {
							Description: "Memory (bytes) that has been already used up, as a point-in-time snapshot.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"memory_hot_add_enabled": {
				Description: "Adding memory to a running VM.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "User-provided name to identify the virtual machine.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_count": {
				Description: "Indicates how many networks are used by this VM.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			},
			"port_groups": {
				Description: "List of portgroup names allocated to this VM.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"power_state": {
				Description: "Power state of the virtual machine.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_capacity": {
				Description: "The capacity and usage information for CPU power on this virtual machine.",
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
						"capacity": {
							Description: "Total capacity of the entity in MHz.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"free": {
							Description: "Free CPU capacity in MHz, as a point-in-time snapshot. The available CPU capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used CPU capacity is also reported.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"used": {
							Description: "Used CPU capacity of the entity in MHz, as a point-in-time snapshot.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"protected_vm": {
				Description: "Shows if this is a protected VM. VMs can be in protection groups.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"registered_device": {
				Description: "Every inventory object comes from a device endpoint. The identity of that device is captured here so that any entity that needs to send a request to that device can use the inventory object to access it.",
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
				Computed: true,
			},
			"remote_display_info": {
				Description: "Applies only when remoteDisplayvnc is enabled.",
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
						"remote_display_password": {
							Description: "The password used for remote access. It is stored base64 encoded.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"remote_display_vnc_key": {
							Description: "The access key for the remote display, potentially a long string.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"remote_display_vnc_port": {
							Description: "Applies only when remoteDisplayvnc is enabled.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"remote_display_vnc_enabled": {
				Description: "Shows if support for a remote VNC access is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"resource_pool": {
				Description: "Name of the resource pool to which this VM belongs (optional).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_pool_owner": {
				Description: "Who owns the resource pool.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_pool_parent": {
				Description: "The parent of the current resource pool to which this VM belongs.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Computed: true,
			},
			"tool_running_status": {
				Description: "Indicates if guest tools are running on this VM. Could be set to guestToolNotRunning or guestToolsRunning.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tools_version": {
				Description: "The version of the guest tools, usually not specified.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"uuid": {
				Description: "The uuid of this virtual machine. The uuid is internally generated and not user specified.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_disk_count": {
				Description: "Shows the number of disks assigned to this VM.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vm_overall_status": {
				Description: "The operational state of the VM. Could be Available, Provisioned, Maintenance mode, Deleting, etc.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_path": {
				Description: "Example - [datastore3] VCSA-134/VCSA-134.vmx.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_version": {
				Description: "Information about the version of this VM (vmx-09, vmx-11 etc.).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_vnic_count": {
				Description: "How many vnics are present.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vnic_device_config_id": {
				Description: "Information related to the guest info's VNIC virtual device. It is a comma-separated list.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceVirtualizationVmwareVirtualMachineRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "virtualization/VmwareVirtualMachines"
	var o models.VirtualizationVmwareVirtualMachine
	if v, ok := d.GetOk("annotation"); ok {
		x := (v.(string))
		o.Annotation = x
	}
	if v, ok := d.GetOk("boot_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.BootTime = x
	}
	if v, ok := d.GetOk("cpu_hot_add_enabled"); ok {
		x := (v.(bool))
		o.CPUHotAddEnabled = &x
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x
	}
	if v, ok := d.GetOk("config_name"); ok {
		x := (v.(string))
		o.ConfigName = x
	}
	if v, ok := d.GetOk("connection_state"); ok {
		x := (v.(string))
		o.ConnectionState = x
	}
	if v, ok := d.GetOk("default_power_off_type"); ok {
		x := (v.(string))
		o.DefaultPowerOffType = x
	}
	if v, ok := d.GetOk("dhcp_enabled"); ok {
		x := (v.(bool))
		o.DhcpEnabled = &x
	}
	if v, ok := d.GetOk("folder"); ok {
		x := (v.(string))
		o.Folder = x
	}
	if v, ok := d.GetOk("guest_state"); ok {
		x := (v.(string))
		o.GuestState = &x
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.HypervisorType = x
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.Identity = x
	}
	if v, ok := d.GetOk("instance_uuid"); ok {
		x := (v.(string))
		o.InstanceUUID = x
	}
	if v, ok := d.GetOk("is_template"); ok {
		x := (v.(bool))
		o.IsTemplate = &x
	}
	if v, ok := d.GetOk("memory_hot_add_enabled"); ok {
		x := (v.(bool))
		o.MemoryHotAddEnabled = &x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x
	}
	if v, ok := d.GetOk("network_count"); ok {
		x := int64(v.(int))
		o.NetworkCount = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("power_state"); ok {
		x := (v.(string))
		o.PowerState = &x
	}
	if v, ok := d.GetOk("protected_vm"); ok {
		x := (v.(bool))
		o.ProtectedVM = &x
	}
	if v, ok := d.GetOk("remote_display_vnc_enabled"); ok {
		x := (v.(bool))
		o.RemoteDisplayVncEnabled = &x
	}
	if v, ok := d.GetOk("resource_pool"); ok {
		x := (v.(string))
		o.ResourcePool = x
	}
	if v, ok := d.GetOk("resource_pool_owner"); ok {
		x := (v.(string))
		o.ResourcePoolOwner = x
	}
	if v, ok := d.GetOk("resource_pool_parent"); ok {
		x := (v.(string))
		o.ResourcePoolParent = x
	}
	if v, ok := d.GetOk("tool_running_status"); ok {
		x := (v.(string))
		o.ToolRunningStatus = x
	}
	if v, ok := d.GetOk("tools_version"); ok {
		x := (v.(string))
		o.ToolsVersion = x
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.UUID = x
	}
	if v, ok := d.GetOk("vm_disk_count"); ok {
		x := int64(v.(int))
		o.VMDiskCount = x
	}
	if v, ok := d.GetOk("vm_overall_status"); ok {
		x := (v.(string))
		o.VMOverallStatus = x
	}
	if v, ok := d.GetOk("vm_path"); ok {
		x := (v.(string))
		o.VMPath = x
	}
	if v, ok := d.GetOk("vm_version"); ok {
		x := (v.(string))
		o.VMVersion = x
	}
	if v, ok := d.GetOk("vm_vnic_count"); ok {
		x := int64(v.(int))
		o.VMVnicCount = x
	}
	if v, ok := d.GetOk("vnic_device_config_id"); ok {
		x := (v.(string))
		o.VnicDeviceConfigID = x
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
			var s models.VirtualizationVmwareVirtualMachine
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("annotation", (s.Annotation)); err != nil {
				return err
			}

			if err := d.Set("boot_time", (s.BootTime).String()); err != nil {
				return err
			}
			if err := d.Set("cpu_hot_add_enabled", (s.CPUHotAddEnabled)); err != nil {
				return err
			}

			if err := d.Set("cpu_shares", flattenMapVirtualizationVMCPUShareInfo(s.CPUShares, d)); err != nil {
				return err
			}

			if err := d.Set("cpu_socket_info", flattenMapVirtualizationVMCPUSocketInfo(s.CPUSocketInfo, d)); err != nil {
				return err
			}

			if err := d.Set("capacity", flattenMapInfraHardwareInfo(s.Capacity, d)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassID)); err != nil {
				return err
			}

			if err := d.Set("cluster", flattenMapVirtualizationVmwareClusterRef(s.Cluster, d)); err != nil {
				return err
			}
			if err := d.Set("config_name", (s.ConfigName)); err != nil {
				return err
			}
			if err := d.Set("connection_state", (s.ConnectionState)); err != nil {
				return err
			}
			if err := d.Set("custom_attributes", (s.CustomAttributes)); err != nil {
				return err
			}
			if err := d.Set("dns_server_list", (s.DNSServerList)); err != nil {
				return err
			}
			if err := d.Set("dns_suffix_list", (s.DNSSuffixList)); err != nil {
				return err
			}

			if err := d.Set("datacenter", flattenMapVirtualizationVmwareDatacenterRef(s.Datacenter, d)); err != nil {
				return err
			}

			if err := d.Set("datastores", flattenListVirtualizationVmwareDatastoreRef(s.Datastores, d)); err != nil {
				return err
			}
			if err := d.Set("default_power_off_type", (s.DefaultPowerOffType)); err != nil {
				return err
			}
			if err := d.Set("dhcp_enabled", (s.DhcpEnabled)); err != nil {
				return err
			}

			if err := d.Set("disk_commit_info", flattenMapVirtualizationVMDiskCommitInfo(s.DiskCommitInfo, d)); err != nil {
				return err
			}
			if err := d.Set("folder", (s.Folder)); err != nil {
				return err
			}

			if err := d.Set("guest_info", flattenMapVirtualizationGuestInfo(s.GuestInfo, d)); err != nil {
				return err
			}
			if err := d.Set("guest_state", (s.GuestState)); err != nil {
				return err
			}

			if err := d.Set("host", flattenMapVirtualizationVmwareHostRef(s.Host, d)); err != nil {
				return err
			}
			if err := d.Set("hypervisor_type", (s.HypervisorType)); err != nil {
				return err
			}
			if err := d.Set("ip_address", (s.IPAddress)); err != nil {
				return err
			}
			if err := d.Set("identity", (s.Identity)); err != nil {
				return err
			}
			if err := d.Set("instance_uuid", (s.InstanceUUID)); err != nil {
				return err
			}
			if err := d.Set("is_template", (s.IsTemplate)); err != nil {
				return err
			}
			if err := d.Set("mac_address", (s.MacAddress)); err != nil {
				return err
			}

			if err := d.Set("mem_shares", flattenMapVirtualizationVMMemoryShareInfo(s.MemShares, d)); err != nil {
				return err
			}

			if err := d.Set("memory_capacity", flattenMapVirtualizationMemoryCapacity(s.MemoryCapacity, d)); err != nil {
				return err
			}
			if err := d.Set("memory_hot_add_enabled", (s.MemoryHotAddEnabled)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("network_count", (s.NetworkCount)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("port_groups", (s.PortGroups)); err != nil {
				return err
			}
			if err := d.Set("power_state", (s.PowerState)); err != nil {
				return err
			}

			if err := d.Set("processor_capacity", flattenMapVirtualizationComputeCapacity(s.ProcessorCapacity, d)); err != nil {
				return err
			}
			if err := d.Set("protected_vm", (s.ProtectedVM)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}

			if err := d.Set("remote_display_info", flattenMapVirtualizationRemoteDisplayInfo(s.RemoteDisplayInfo, d)); err != nil {
				return err
			}
			if err := d.Set("remote_display_vnc_enabled", (s.RemoteDisplayVncEnabled)); err != nil {
				return err
			}
			if err := d.Set("resource_pool", (s.ResourcePool)); err != nil {
				return err
			}
			if err := d.Set("resource_pool_owner", (s.ResourcePoolOwner)); err != nil {
				return err
			}
			if err := d.Set("resource_pool_parent", (s.ResourcePoolParent)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("tool_running_status", (s.ToolRunningStatus)); err != nil {
				return err
			}
			if err := d.Set("tools_version", (s.ToolsVersion)); err != nil {
				return err
			}
			if err := d.Set("uuid", (s.UUID)); err != nil {
				return err
			}
			if err := d.Set("vm_disk_count", (s.VMDiskCount)); err != nil {
				return err
			}
			if err := d.Set("vm_overall_status", (s.VMOverallStatus)); err != nil {
				return err
			}
			if err := d.Set("vm_path", (s.VMPath)); err != nil {
				return err
			}
			if err := d.Set("vm_version", (s.VMVersion)); err != nil {
				return err
			}
			if err := d.Set("vm_vnic_count", (s.VMVnicCount)); err != nil {
				return err
			}
			if err := d.Set("vnic_device_config_id", (s.VnicDeviceConfigID)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
