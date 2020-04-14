package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceVirtualizationVmwareDatastore() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVirtualizationVmwareDatastoreRead,
		Schema: map[string]*schema.Schema{
			"accessible": {
				Description: "Shows if this datastore is accessible.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"capacity": {
				Description: "Storage utilization for this datastore.",
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
							Description: "The total capacity of the entity (bytes).",
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
							Description: "Free storage space remaining in the entity (bytes) as a point-in-time snapshot. The available space is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used storage capacity is also reported.",
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
							Description: "Spaced already used by this entity (bytes), as a point-in-time snapshot.",
							Type:        schema.TypeInt,
							Optional:    true,
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
				Description: "Identifies the cluster associated with the datastore entity. Not every datastore is in a cluster and therefore this relation may not exist for all datastores.",
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
			"host_count": {
				Description: "Number of hosts attached to or supported-by this datastore.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"hosts": {
				Description: "The list of hosts that are attached to this host. Used for the new inventory model and will soon replace hostList above.",
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
			"identity": {
				Description: "The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"maintenance_mode": {
				Description: "Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"multiple_host_access": {
				Description: "Indicates if this datastore is connected to multiple hosts.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"name": {
				Description: "Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations.",
				Type:        schema.TypeString,
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
			"status": {
				Description: "Datastore health status, as reported by the hypervisor platform.",
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
			"thin_provisioning_supported": {
				Description: "Indicates if this datastore supports thin provisioning for files.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"type": {
				Description: "A string indicating the type of the datastore (VMFS, NFS, etc).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"url": {
				Description: "The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/').",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"un_committed": {
				Description: "Space uncommitted in this datastore in bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vm_count": {
				Description: "Number of virtual machines relying on (using) this datastore.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}
func dataSourceVirtualizationVmwareDatastoreRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "virtualization/VmwareDatastores"
	var o models.VirtualizationVmwareDatastore
	if v, ok := d.GetOk("accessible"); ok {
		x := (v.(bool))
		o.Accessible = &x
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x
	}
	if v, ok := d.GetOk("host_count"); ok {
		x := int64(v.(int))
		o.HostCount = x
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.Identity = x
	}
	if v, ok := d.GetOk("maintenance_mode"); ok {
		x := (v.(bool))
		o.MaintenanceMode = &x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("multiple_host_access"); ok {
		x := (v.(bool))
		o.MultipleHostAccess = &x
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = &x
	}
	if v, ok := d.GetOk("thin_provisioning_supported"); ok {
		x := (v.(bool))
		o.ThinProvisioningSupported = &x
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.Type = &x
	}
	if v, ok := d.GetOk("url"); ok {
		x := (v.(string))
		o.URL = x
	}
	if v, ok := d.GetOk("un_committed"); ok {
		x := int64(v.(int))
		o.UnCommitted = x
	}
	if v, ok := d.GetOk("vm_count"); ok {
		x := int64(v.(int))
		o.VMCount = x
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
			var s models.VirtualizationVmwareDatastore
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("accessible", (s.Accessible)); err != nil {
				return err
			}

			if err := d.Set("capacity", flattenMapVirtualizationStorageCapacity(s.Capacity, d)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassID)); err != nil {
				return err
			}

			if err := d.Set("cluster", flattenMapVirtualizationVmwareClusterRef(s.Cluster, d)); err != nil {
				return err
			}

			if err := d.Set("datacenter", flattenMapVirtualizationVmwareDatacenterRef(s.Datacenter, d)); err != nil {
				return err
			}
			if err := d.Set("host_count", (s.HostCount)); err != nil {
				return err
			}

			if err := d.Set("hosts", flattenListVirtualizationVmwareHostRef(s.Hosts, d)); err != nil {
				return err
			}
			if err := d.Set("identity", (s.Identity)); err != nil {
				return err
			}
			if err := d.Set("maintenance_mode", (s.MaintenanceMode)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("multiple_host_access", (s.MultipleHostAccess)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("thin_provisioning_supported", (s.ThinProvisioningSupported)); err != nil {
				return err
			}
			if err := d.Set("type", (s.Type)); err != nil {
				return err
			}
			if err := d.Set("url", (s.URL)); err != nil {
				return err
			}
			if err := d.Set("un_committed", (s.UnCommitted)); err != nil {
				return err
			}
			if err := d.Set("vm_count", (s.VMCount)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
