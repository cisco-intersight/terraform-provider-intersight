package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNetworkElement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkElementRead,
		Schema: map[string]*schema.Schema{
			"admin_inband_interface_state": {
				Description: "The administrative state of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cards": {
				Description: "An array of relationships to equipmentSwitchCard resources.",
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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
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
			"fanmodules": {
				Description: "An array of relationships to equipmentFanModule resources.",
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
			"fault_summary": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inband_ip_address": {
				Description: "The IP address of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_gateway": {
				Description: "The default gateway of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_mask": {
				Description: "The network mask of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_vlan": {
				Description: "The VLAN ID of the network Element inband management interface.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"management_contoller": {
				Description: "A reference to a managementController resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
			"management_entity": {
				Description: "A reference to a managementEntity resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_address": {
				Description: "The IP address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_gateway": {
				Description: "The default gateway of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_mask": {
				Description: "The network mask of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv4_address": {
				Description: "The IPv4 address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv4_gateway": {
				Description: "The default IPv4 gateway of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv4_mask": {
				Description: "The network mask of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv6_address": {
				Description: "The IPv6 address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_ipv6_gateway": {
				Description: "The default IPv6 gateway of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_ipv6_prefix": {
				Description: "The network mask of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_mac": {
				Description: "The MAC address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"psus": {
				Description: "An array of relationships to equipmentPsu resources.",
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
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
			"revision": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
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
			"switch_id": {
				Description: "The Switch Id of the network Element.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"top_system": {
				Description: "A reference to a topSystem resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"ucsm_running_firmware": {
				Description: "A reference to a firmwareRunningFirmware resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceNetworkElementRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewNetworkElement()
	if v, ok := d.GetOk("admin_inband_interface_state"); ok {
		x := (v.(string))
		o.SetAdminInbandInterfaceState(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("fault_summary"); ok {
		x := int64(v.(int))
		o.SetFaultSummary(x)
	}
	if v, ok := d.GetOk("inband_ip_address"); ok {
		x := (v.(string))
		o.SetInbandIpAddress(x)
	}
	if v, ok := d.GetOk("inband_ip_gateway"); ok {
		x := (v.(string))
		o.SetInbandIpGateway(x)
	}
	if v, ok := d.GetOk("inband_ip_mask"); ok {
		x := (v.(string))
		o.SetInbandIpMask(x)
	}
	if v, ok := d.GetOk("inband_vlan"); ok {
		x := int64(v.(int))
		o.SetInbandVlan(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("out_of_band_ip_address"); ok {
		x := (v.(string))
		o.SetOutOfBandIpAddress(x)
	}
	if v, ok := d.GetOk("out_of_band_ip_gateway"); ok {
		x := (v.(string))
		o.SetOutOfBandIpGateway(x)
	}
	if v, ok := d.GetOk("out_of_band_ip_mask"); ok {
		x := (v.(string))
		o.SetOutOfBandIpMask(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv4_address"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv4Address(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv4_gateway"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv4Gateway(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv4_mask"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv4Mask(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv6_address"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv6Address(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv6_gateway"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv6Gateway(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv6_prefix"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv6Prefix(x)
	}
	if v, ok := d.GetOk("out_of_band_mac"); ok {
		x := (v.(string))
		o.SetOutOfBandMac(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("switch_id"); ok {
		x := (v.(string))
		o.SetSwitchId(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.NetworkApi.GetNetworkElementList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewNetworkElement()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("admin_inband_interface_state", (s.AdminInbandInterfaceState)); err != nil {
				return err
			}

			if err := d.Set("cards", flattenListEquipmentSwitchCardRelationship(s.Cards, d)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoId)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}

			if err := d.Set("fanmodules", flattenListEquipmentFanModuleRelationship(s.Fanmodules, d)); err != nil {
				return err
			}
			if err := d.Set("fault_summary", (s.FaultSummary)); err != nil {
				return err
			}
			if err := d.Set("inband_ip_address", (s.InbandIpAddress)); err != nil {
				return err
			}
			if err := d.Set("inband_ip_gateway", (s.InbandIpGateway)); err != nil {
				return err
			}
			if err := d.Set("inband_ip_mask", (s.InbandIpMask)); err != nil {
				return err
			}
			if err := d.Set("inband_vlan", (s.InbandVlan)); err != nil {
				return err
			}

			if err := d.Set("management_contoller", flattenMapManagementControllerRelationship(s.ManagementContoller, d)); err != nil {
				return err
			}

			if err := d.Set("management_entity", flattenMapManagementEntityRelationship(s.ManagementEntity, d)); err != nil {
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
			if err := d.Set("out_of_band_ip_address", (s.OutOfBandIpAddress)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ip_gateway", (s.OutOfBandIpGateway)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ip_mask", (s.OutOfBandIpMask)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ipv4_address", (s.OutOfBandIpv4Address)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ipv4_gateway", (s.OutOfBandIpv4Gateway)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ipv4_mask", (s.OutOfBandIpv4Mask)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ipv6_address", (s.OutOfBandIpv6Address)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ipv6_gateway", (s.OutOfBandIpv6Gateway)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ipv6_prefix", (s.OutOfBandIpv6Prefix)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_mac", (s.OutOfBandMac)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
				return err
			}

			if err := d.Set("psus", flattenListEquipmentPsuRelationship(s.Psus, d)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("revision", (s.Revision)); err != nil {
				return err
			}
			if err := d.Set("rn", (s.Rn)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}
			if err := d.Set("switch_id", (s.SwitchId)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}

			if err := d.Set("top_system", flattenMapTopSystemRelationship(s.TopSystem, d)); err != nil {
				return err
			}

			if err := d.Set("ucsm_running_firmware", flattenMapFirmwareRunningFirmwareRelationship(s.UcsmRunningFirmware, d)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
