package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementInterfaceRead,
		Schema: map[string]*schema.Schema{
			"device_mo_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Description: "Default gateway for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"host_name": {
				Description: "Hostname configured for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ip_address": {
				Description: "IP address of the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_address": {
				Description: "IPv4 address of the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_gateway": {
				Description: "IPv4 default gateway for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_mask": {
				Description: "IPv4 Netmask for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv6_address": {
				Description: "IPv6 address of the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv6_gateway": {
				Description: "IPv6 default gateway for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv6_prefix": {
				Description: "IPv6 prefix for the interface.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"mac_address": {
				Description: "MAC address configured for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"management_controller": {
				Description: "A collection of references to the [management.Controller](mo://management.Controller) Managed Object.\nWhen this managed object is deleted, the referenced [management.Controller](mo://management.Controller) MO unsets its reference to this deleted MO.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mask": {
				Description: "Netmask for the interface.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"switch_id": {
				Description: "Switch Id of the interface.",
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
			"uem_conn_status": {
				Description: "Status of UEM connection.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"virtual_host_name": {
				Description: "Virtual hostname configured for the interface in case of clustered environment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceManagementInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "management/Interfaces"
	var o models.ManagementInterface
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.DeviceMoID = x
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.Dn = x
	}
	if v, ok := d.GetOk("gateway"); ok {
		x := (v.(string))
		o.Gateway = x
	}
	if v, ok := d.GetOk("host_name"); ok {
		x := (v.(string))
		o.HostName = x
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.IPAddress = x
	}
	if v, ok := d.GetOk("ipv4_address"); ok {
		x := (v.(string))
		o.IPV4Address = x
	}
	if v, ok := d.GetOk("ipv4_gateway"); ok {
		x := (v.(string))
		o.IPV4Gateway = x
	}
	if v, ok := d.GetOk("ipv4_mask"); ok {
		x := (v.(string))
		o.IPV4Mask = x
	}
	if v, ok := d.GetOk("ipv6_address"); ok {
		x := (v.(string))
		o.IPV6Address = x
	}
	if v, ok := d.GetOk("ipv6_gateway"); ok {
		x := (v.(string))
		o.IPV6Gateway = x
	}
	if v, ok := d.GetOk("ipv6_prefix"); ok {
		x := int64(v.(int))
		o.IPV6Prefix = x
	}
	if v, ok := d.GetOk("mac_address"); ok {
		x := (v.(string))
		o.MacAddress = x
	}
	if v, ok := d.GetOk("mask"); ok {
		x := (v.(string))
		o.Mask = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.Rn = x
	}
	if v, ok := d.GetOk("switch_id"); ok {
		x := (v.(string))
		o.SwitchID = x
	}
	if v, ok := d.GetOk("uem_conn_status"); ok {
		x := (v.(string))
		o.UemConnStatus = x
	}
	if v, ok := d.GetOk("virtual_host_name"); ok {
		x := (v.(string))
		o.VirtualHostName = x
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
			var s models.ManagementInterface
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoID)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}
			if err := d.Set("gateway", (s.Gateway)); err != nil {
				return err
			}
			if err := d.Set("host_name", (s.HostName)); err != nil {
				return err
			}
			if err := d.Set("ip_address", (s.IPAddress)); err != nil {
				return err
			}
			if err := d.Set("ipv4_address", (s.IPV4Address)); err != nil {
				return err
			}
			if err := d.Set("ipv4_gateway", (s.IPV4Gateway)); err != nil {
				return err
			}
			if err := d.Set("ipv4_mask", (s.IPV4Mask)); err != nil {
				return err
			}
			if err := d.Set("ipv6_address", (s.IPV6Address)); err != nil {
				return err
			}
			if err := d.Set("ipv6_gateway", (s.IPV6Gateway)); err != nil {
				return err
			}
			if err := d.Set("ipv6_prefix", (s.IPV6Prefix)); err != nil {
				return err
			}
			if err := d.Set("mac_address", (s.MacAddress)); err != nil {
				return err
			}

			if err := d.Set("management_controller", flattenMapManagementControllerRef(s.ManagementController, d)); err != nil {
				return err
			}
			if err := d.Set("mask", (s.Mask)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
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
			if err := d.Set("rn", (s.Rn)); err != nil {
				return err
			}
			if err := d.Set("switch_id", (s.SwitchID)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("uem_conn_status", (s.UemConnStatus)); err != nil {
				return err
			}
			if err := d.Set("virtual_host_name", (s.VirtualHostName)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
