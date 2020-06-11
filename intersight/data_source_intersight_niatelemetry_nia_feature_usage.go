package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiatelemetryNiaFeatureUsage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiatelemetryNiaFeatureUsageRead,
		Schema: map[string]*schema.Schema{
			"apic_count": {
				Description: "Number of APIC controllers. This determines the value of controllers for the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"app_center_count": {
				Description: "ACI APPs feature usage. This determines the total number of apps installed on the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ave": {
				Description: "AVE feature usage. This determines if ACI virtual edge feature is enabled or disabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"bd_count": {
				Description: "Number of BDs. This determines the total number of Broadcast Domains across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"consistency_checker_app": {
				Description: "Consistency checker application usage. This determines if the fabric has Consistency checker application installed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"contract_count": {
				Description: "Number of contracts. This determines the total number of Contracts configured across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"dns_count": {
				Description: "DNS feature usage. This determines the total number of DNS configurations across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"eigrp_count": {
				Description: "Eigrp feature usage. This determines the total number of EIGRP sessions across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"epg_count": {
				Description: "Number of EPGs. This determines the total number of End Point Groups across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"hsrp_count": {
				Description: "Hsrp feature usage. This determines the total number of HSRP sessions across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ip_epg_count": {
				Description: "Number of IP based EPGs. This determines the total number of IP End Point Groups across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ibgp_count": {
				Description: "Ibgp feature usage. This determines the total number of BGP sessions across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"igmp_access_list_count": {
				Description: "IGMP Access List feature usage. This determines the total number of IGMP access lists configured across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"igmp_snoop": {
				Description: "IGMP Snooping feature usage. This determines if this feature is enabled or disabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"isis_count": {
				Description: "Isis feature usage. TThis determines the total number of ISIS sessions across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"l2_multicast": {
				Description: "L2Multicast feature usage. This determines if this Layer 2 Multicast feature is being enabled / disabled on the fabric.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"leaf_count": {
				Description: "Number of Leafs. This determines the total number of Leaf switches in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"maintenance_mode_count": {
				Description: "Maintenance Mode feature usage. This determines the number of switches that are currently in maintenance mode.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"management_over_v6_count": {
				Description: "Management over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nir": {
				Description: "NIR application usage. This determines if the fabric has NIR application installed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"opflex_kubernetes_count": {
				Description: "Opflex for Kubernetes feature usage. This determines the total number of VMM sessions of type kubernetes.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ospf_count": {
				Description: "Ospf feature usage. This determines the total number of OSPF sessions across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
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
			"poe_count": {
				Description: "POE feature usage. This determines the total number of POE configurations across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"qin_vni_tunnel_count": {
				Description: "QinVniTunnel feature usage. This determines if the qinVniTunnel feature is being used on the fabric and the scale of it.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"registered_device": {
				Description: "Relationship to the Device Registration object for this setup.",
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
			"remote_leaf_count": {
				Description: "Number of remote Leafs. This determines if this feature is being enabled or disabled.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ssh_over_v6_count": {
				Description: "Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"scvmm_count": {
				Description: "SCVMM feature usage. This determines the total number of SCVMM configurations in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"shared_l3_out_count": {
				Description: "SharedL3Out feature usage. This determines the total number of Shared L3 out configured across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"smart_call_home": {
				Description: "Smart callhome feature usage. This determines if this feature is being enabled or disabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snmp": {
				Description: "SNMP feature usage. This determines if this feature is enabled or disabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"spine_count": {
				Description: "Number of Spines. This determines the total number of spine switches in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"syslog_over_v6_count": {
				Description: "Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
				Type:        schema.TypeInt,
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
			"tenant_count": {
				Description: "Number of tenants. This determines the total number of tenants configured across the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"tier_two_leaf_count": {
				Description: "Number of tier 2 Leafs. This determines the total number of tier 2 Leaf switches in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"twamp": {
				Description: "TWAMP feature usage. This determines if this feature is enabled or disabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"useg": {
				Description: "VMM uSegmentation feature usage. This determines if microsegmentation feature is enabled or disabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vpod_count": {
				Description: "Virtual pod feature usage. This determines the total number of virtual POD configurations in the fabrics.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}
func dataSourceNiatelemetryNiaFeatureUsageRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "niatelemetry/NiaFeatureUsages"
	var o models.NiatelemetryNiaFeatureUsage
	if v, ok := d.GetOk("apic_count"); ok {
		x := int64(v.(int))
		o.ApicCount = x
	}
	if v, ok := d.GetOk("app_center_count"); ok {
		x := int64(v.(int))
		o.AppCenterCount = x
	}
	if v, ok := d.GetOk("ave"); ok {
		x := (v.(string))
		o.Ave = x
	}
	if v, ok := d.GetOk("bd_count"); ok {
		x := int64(v.(int))
		o.BdCount = x
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x
	}
	if v, ok := d.GetOk("consistency_checker_app"); ok {
		x := (v.(string))
		o.ConsistencyCheckerApp = x
	}
	if v, ok := d.GetOk("contract_count"); ok {
		x := int64(v.(int))
		o.ContractCount = x
	}
	if v, ok := d.GetOk("dns_count"); ok {
		x := int64(v.(int))
		o.DNSCount = x
	}
	if v, ok := d.GetOk("eigrp_count"); ok {
		x := int64(v.(int))
		o.EigrpCount = x
	}
	if v, ok := d.GetOk("epg_count"); ok {
		x := int64(v.(int))
		o.EpgCount = x
	}
	if v, ok := d.GetOk("hsrp_count"); ok {
		x := int64(v.(int))
		o.HsrpCount = x
	}
	if v, ok := d.GetOk("ip_epg_count"); ok {
		x := int64(v.(int))
		o.IPEpgCount = x
	}
	if v, ok := d.GetOk("ibgp_count"); ok {
		x := int64(v.(int))
		o.IbgpCount = x
	}
	if v, ok := d.GetOk("igmp_access_list_count"); ok {
		x := int64(v.(int))
		o.IgmpAccessListCount = x
	}
	if v, ok := d.GetOk("igmp_snoop"); ok {
		x := (v.(string))
		o.IgmpSnoop = x
	}
	if v, ok := d.GetOk("isis_count"); ok {
		x := int64(v.(int))
		o.IsisCount = x
	}
	if v, ok := d.GetOk("l2_multicast"); ok {
		x := (v.(string))
		o.L2Multicast = x
	}
	if v, ok := d.GetOk("leaf_count"); ok {
		x := int64(v.(int))
		o.LeafCount = x
	}
	if v, ok := d.GetOk("maintenance_mode_count"); ok {
		x := int64(v.(int))
		o.MaintenanceModeCount = x
	}
	if v, ok := d.GetOk("management_over_v6_count"); ok {
		x := int64(v.(int))
		o.ManagementOverV6Count = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("nir"); ok {
		x := (v.(string))
		o.Nir = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("opflex_kubernetes_count"); ok {
		x := int64(v.(int))
		o.OpflexKubernetesCount = x
	}
	if v, ok := d.GetOk("ospf_count"); ok {
		x := int64(v.(int))
		o.OspfCount = x
	}
	if v, ok := d.GetOk("poe_count"); ok {
		x := int64(v.(int))
		o.PoeCount = x
	}
	if v, ok := d.GetOk("qin_vni_tunnel_count"); ok {
		x := int64(v.(int))
		o.QinVniTunnelCount = x
	}
	if v, ok := d.GetOk("remote_leaf_count"); ok {
		x := int64(v.(int))
		o.RemoteLeafCount = x
	}
	if v, ok := d.GetOk("ssh_over_v6_count"); ok {
		x := int64(v.(int))
		o.SSHOverV6Count = x
	}
	if v, ok := d.GetOk("scvmm_count"); ok {
		x := int64(v.(int))
		o.ScvmmCount = x
	}
	if v, ok := d.GetOk("shared_l3_out_count"); ok {
		x := int64(v.(int))
		o.SharedL3OutCount = x
	}
	if v, ok := d.GetOk("smart_call_home"); ok {
		x := (v.(string))
		o.SmartCallHome = x
	}
	if v, ok := d.GetOk("snmp"); ok {
		x := (v.(string))
		o.Snmp = x
	}
	if v, ok := d.GetOk("spine_count"); ok {
		x := int64(v.(int))
		o.SpineCount = x
	}
	if v, ok := d.GetOk("syslog_over_v6_count"); ok {
		x := int64(v.(int))
		o.SyslogOverV6Count = x
	}
	if v, ok := d.GetOk("tenant_count"); ok {
		x := int64(v.(int))
		o.TenantCount = x
	}
	if v, ok := d.GetOk("tier_two_leaf_count"); ok {
		x := int64(v.(int))
		o.TierTwoLeafCount = x
	}
	if v, ok := d.GetOk("twamp"); ok {
		x := (v.(string))
		o.Twamp = x
	}
	if v, ok := d.GetOk("useg"); ok {
		x := (v.(string))
		o.Useg = x
	}
	if v, ok := d.GetOk("vpod_count"); ok {
		x := int64(v.(int))
		o.VpodCount = x
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
			var s models.NiatelemetryNiaFeatureUsage
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("apic_count", (s.ApicCount)); err != nil {
				return err
			}
			if err := d.Set("app_center_count", (s.AppCenterCount)); err != nil {
				return err
			}
			if err := d.Set("ave", (s.Ave)); err != nil {
				return err
			}
			if err := d.Set("bd_count", (s.BdCount)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassID)); err != nil {
				return err
			}
			if err := d.Set("consistency_checker_app", (s.ConsistencyCheckerApp)); err != nil {
				return err
			}
			if err := d.Set("contract_count", (s.ContractCount)); err != nil {
				return err
			}
			if err := d.Set("dns_count", (s.DNSCount)); err != nil {
				return err
			}
			if err := d.Set("eigrp_count", (s.EigrpCount)); err != nil {
				return err
			}
			if err := d.Set("epg_count", (s.EpgCount)); err != nil {
				return err
			}
			if err := d.Set("hsrp_count", (s.HsrpCount)); err != nil {
				return err
			}
			if err := d.Set("ip_epg_count", (s.IPEpgCount)); err != nil {
				return err
			}
			if err := d.Set("ibgp_count", (s.IbgpCount)); err != nil {
				return err
			}
			if err := d.Set("igmp_access_list_count", (s.IgmpAccessListCount)); err != nil {
				return err
			}
			if err := d.Set("igmp_snoop", (s.IgmpSnoop)); err != nil {
				return err
			}
			if err := d.Set("isis_count", (s.IsisCount)); err != nil {
				return err
			}
			if err := d.Set("l2_multicast", (s.L2Multicast)); err != nil {
				return err
			}
			if err := d.Set("leaf_count", (s.LeafCount)); err != nil {
				return err
			}
			if err := d.Set("maintenance_mode_count", (s.MaintenanceModeCount)); err != nil {
				return err
			}
			if err := d.Set("management_over_v6_count", (s.ManagementOverV6Count)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("nir", (s.Nir)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("opflex_kubernetes_count", (s.OpflexKubernetesCount)); err != nil {
				return err
			}
			if err := d.Set("ospf_count", (s.OspfCount)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("poe_count", (s.PoeCount)); err != nil {
				return err
			}
			if err := d.Set("qin_vni_tunnel_count", (s.QinVniTunnelCount)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("remote_leaf_count", (s.RemoteLeafCount)); err != nil {
				return err
			}
			if err := d.Set("ssh_over_v6_count", (s.SSHOverV6Count)); err != nil {
				return err
			}
			if err := d.Set("scvmm_count", (s.ScvmmCount)); err != nil {
				return err
			}
			if err := d.Set("shared_l3_out_count", (s.SharedL3OutCount)); err != nil {
				return err
			}
			if err := d.Set("smart_call_home", (s.SmartCallHome)); err != nil {
				return err
			}
			if err := d.Set("snmp", (s.Snmp)); err != nil {
				return err
			}
			if err := d.Set("spine_count", (s.SpineCount)); err != nil {
				return err
			}
			if err := d.Set("syslog_over_v6_count", (s.SyslogOverV6Count)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("tenant_count", (s.TenantCount)); err != nil {
				return err
			}
			if err := d.Set("tier_two_leaf_count", (s.TierTwoLeafCount)); err != nil {
				return err
			}
			if err := d.Set("twamp", (s.Twamp)); err != nil {
				return err
			}
			if err := d.Set("useg", (s.Useg)); err != nil {
				return err
			}
			if err := d.Set("vpod_count", (s.VpodCount)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
