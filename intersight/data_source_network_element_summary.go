package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceNetworkElementSummary() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkElementSummaryRead,
		Schema: map[string]*schema.Schema{
			"admin_inband_interface_state": {
				Description: "",
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
				Description: "Running firmware information.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_address": {
				Description: "IP version 4 address is saved in this property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_address": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_gateway": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_mask": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_vlan": {
				Description: "",
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
			"name": {
				Description: "Name of the ElementSummary object is saved in this property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_ether_ports": {
				Description: "Total number of Ethernet ports.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_ether_ports_configured": {
				Description: "Total number of configured Ethernet ports.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_ether_ports_link_up": {
				Description: "Total number of Ethernet ports which are UP.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_expansion_modules": {
				Description: "Total number of expansion modules.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_fc_ports": {
				Description: "Total number of FC ports.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_fc_ports_configured": {
				Description: "Total number of configured FC ports.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_fc_ports_link_up": {
				Description: "Total number of FC ports which are UP.",
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
			"out_of_band_ip_address": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_gateway": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_mask": {
				Description: "",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_mac": {
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
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"source_object_type": {
				Description: "Specifies the source object type for View MO.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"switch_id": {
				Description: "",
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
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"version": {
				Description: "Version holds the firmware version related information.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceNetworkElementSummaryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "network/ElementSummaries"
	var o models.NetworkElementSummary
	if v, ok := d.GetOk("admin_inband_interface_state"); ok {
		x := (v.(string))
		o.AdminInbandInterfaceState = x
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
	if v, ok := d.GetOk("inband_ip_address"); ok {
		x := (v.(string))
		o.InbandIPAddress = x
	}
	if v, ok := d.GetOk("inband_ip_gateway"); ok {
		x := (v.(string))
		o.InbandIPGateway = x
	}
	if v, ok := d.GetOk("inband_ip_mask"); ok {
		x := (v.(string))
		o.InbandIPMask = x
	}
	if v, ok := d.GetOk("inband_vlan"); ok {
		x := int64(v.(int))
		o.InbandVlan = x
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
	if v, ok := d.GetOk("num_ether_ports"); ok {
		x := int64(v.(int))
		o.NumEtherPorts = x
	}
	if v, ok := d.GetOk("num_ether_ports_configured"); ok {
		x := int64(v.(int))
		o.NumEtherPortsConfigured = x
	}
	if v, ok := d.GetOk("num_ether_ports_link_up"); ok {
		x := int64(v.(int))
		o.NumEtherPortsLinkUp = x
	}
	if v, ok := d.GetOk("num_expansion_modules"); ok {
		x := int64(v.(int))
		o.NumExpansionModules = x
	}
	if v, ok := d.GetOk("num_fc_ports"); ok {
		x := int64(v.(int))
		o.NumFcPorts = x
	}
	if v, ok := d.GetOk("num_fc_ports_configured"); ok {
		x := int64(v.(int))
		o.NumFcPortsConfigured = x
	}
	if v, ok := d.GetOk("num_fc_ports_link_up"); ok {
		x := int64(v.(int))
		o.NumFcPortsLinkUp = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("out_of_band_ip_address"); ok {
		x := (v.(string))
		o.OutOfBandIPAddress = x
	}
	if v, ok := d.GetOk("out_of_band_ip_gateway"); ok {
		x := (v.(string))
		o.OutOfBandIPGateway = x
	}
	if v, ok := d.GetOk("out_of_band_ip_mask"); ok {
		x := (v.(string))
		o.OutOfBandIPMask = x
	}
	if v, ok := d.GetOk("out_of_band_mac"); ok {
		x := (v.(string))
		o.OutOfBandMac = x
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.Revision = x
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.Rn = x
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.Serial = x
	}
	if v, ok := d.GetOk("source_object_type"); ok {
		x := (v.(string))
		o.SourceObjectType = x
	}
	if v, ok := d.GetOk("switch_id"); ok {
		x := (v.(string))
		o.SwitchID = x
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x
	}
	if v, ok := d.GetOk("version"); ok {
		x := (v.(string))
		o.Version = x
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
			var s models.NetworkElementSummary
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("admin_inband_interface_state", (s.AdminInbandInterfaceState)); err != nil {
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
			if err := d.Set("inband_ip_address", (s.InbandIPAddress)); err != nil {
				return err
			}
			if err := d.Set("inband_ip_gateway", (s.InbandIPGateway)); err != nil {
				return err
			}
			if err := d.Set("inband_ip_mask", (s.InbandIPMask)); err != nil {
				return err
			}
			if err := d.Set("inband_vlan", (s.InbandVlan)); err != nil {
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
			if err := d.Set("num_ether_ports", (s.NumEtherPorts)); err != nil {
				return err
			}
			if err := d.Set("num_ether_ports_configured", (s.NumEtherPortsConfigured)); err != nil {
				return err
			}
			if err := d.Set("num_ether_ports_link_up", (s.NumEtherPortsLinkUp)); err != nil {
				return err
			}
			if err := d.Set("num_expansion_modules", (s.NumExpansionModules)); err != nil {
				return err
			}
			if err := d.Set("num_fc_ports", (s.NumFcPorts)); err != nil {
				return err
			}
			if err := d.Set("num_fc_ports_configured", (s.NumFcPortsConfigured)); err != nil {
				return err
			}
			if err := d.Set("num_fc_ports_link_up", (s.NumFcPortsLinkUp)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ip_address", (s.OutOfBandIPAddress)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ip_gateway", (s.OutOfBandIPGateway)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_ip_mask", (s.OutOfBandIPMask)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_mac", (s.OutOfBandMac)); err != nil {
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
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}
			if err := d.Set("source_object_type", (s.SourceObjectType)); err != nil {
				return err
			}
			if err := d.Set("switch_id", (s.SwitchID)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			if err := d.Set("version", (s.Version)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
