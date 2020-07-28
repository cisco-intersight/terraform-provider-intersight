package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiatelemetryNiaFeatureUsage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiatelemetryNiaFeatureUsageRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
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
			"ip_epg_count": {
				Description: "Number of IP based EPGs. This determines the total number of IP End Point Groups across the fabric.",
				Type:        schema.TypeInt,
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
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
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
			"remote_leaf_count": {
				Description: "Number of remote Leafs. This determines if this feature is being enabled or disabled.",
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
			"ssh_over_v6_count": {
				Description: "Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"syslog_over_v6_count": {
				Description: "Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
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
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
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
	var o = models.NewNiatelemetryNiaFeatureUsageWithDefaults()
	if v, ok := d.GetOk("apic_count"); ok {
		x := int64(v.(int))
		o.SetApicCount(x)
	}
	if v, ok := d.GetOk("app_center_count"); ok {
		x := int64(v.(int))
		o.SetAppCenterCount(x)
	}
	if v, ok := d.GetOk("ave"); ok {
		x := (v.(string))
		o.SetAve(x)
	}
	if v, ok := d.GetOk("bd_count"); ok {
		x := int64(v.(int))
		o.SetBdCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("consistency_checker_app"); ok {
		x := (v.(string))
		o.SetConsistencyCheckerApp(x)
	}
	if v, ok := d.GetOk("contract_count"); ok {
		x := int64(v.(int))
		o.SetContractCount(x)
	}
	if v, ok := d.GetOk("dns_count"); ok {
		x := int64(v.(int))
		o.SetDnsCount(x)
	}
	if v, ok := d.GetOk("eigrp_count"); ok {
		x := int64(v.(int))
		o.SetEigrpCount(x)
	}
	if v, ok := d.GetOk("epg_count"); ok {
		x := int64(v.(int))
		o.SetEpgCount(x)
	}
	if v, ok := d.GetOk("hsrp_count"); ok {
		x := int64(v.(int))
		o.SetHsrpCount(x)
	}
	if v, ok := d.GetOk("ibgp_count"); ok {
		x := int64(v.(int))
		o.SetIbgpCount(x)
	}
	if v, ok := d.GetOk("igmp_access_list_count"); ok {
		x := int64(v.(int))
		o.SetIgmpAccessListCount(x)
	}
	if v, ok := d.GetOk("igmp_snoop"); ok {
		x := (v.(string))
		o.SetIgmpSnoop(x)
	}
	if v, ok := d.GetOk("ip_epg_count"); ok {
		x := int64(v.(int))
		o.SetIpEpgCount(x)
	}
	if v, ok := d.GetOk("isis_count"); ok {
		x := int64(v.(int))
		o.SetIsisCount(x)
	}
	if v, ok := d.GetOk("l2_multicast"); ok {
		x := (v.(string))
		o.SetL2Multicast(x)
	}
	if v, ok := d.GetOk("leaf_count"); ok {
		x := int64(v.(int))
		o.SetLeafCount(x)
	}
	if v, ok := d.GetOk("maintenance_mode_count"); ok {
		x := int64(v.(int))
		o.SetMaintenanceModeCount(x)
	}
	if v, ok := d.GetOk("management_over_v6_count"); ok {
		x := int64(v.(int))
		o.SetManagementOverV6Count(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("nir"); ok {
		x := (v.(string))
		o.SetNir(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("opflex_kubernetes_count"); ok {
		x := int64(v.(int))
		o.SetOpflexKubernetesCount(x)
	}
	if v, ok := d.GetOk("ospf_count"); ok {
		x := int64(v.(int))
		o.SetOspfCount(x)
	}
	if v, ok := d.GetOk("poe_count"); ok {
		x := int64(v.(int))
		o.SetPoeCount(x)
	}
	if v, ok := d.GetOk("qin_vni_tunnel_count"); ok {
		x := int64(v.(int))
		o.SetQinVniTunnelCount(x)
	}
	if v, ok := d.GetOk("remote_leaf_count"); ok {
		x := int64(v.(int))
		o.SetRemoteLeafCount(x)
	}
	if v, ok := d.GetOk("scvmm_count"); ok {
		x := int64(v.(int))
		o.SetScvmmCount(x)
	}
	if v, ok := d.GetOk("shared_l3_out_count"); ok {
		x := int64(v.(int))
		o.SetSharedL3OutCount(x)
	}
	if v, ok := d.GetOk("smart_call_home"); ok {
		x := (v.(string))
		o.SetSmartCallHome(x)
	}
	if v, ok := d.GetOk("snmp"); ok {
		x := (v.(string))
		o.SetSnmp(x)
	}
	if v, ok := d.GetOk("spine_count"); ok {
		x := int64(v.(int))
		o.SetSpineCount(x)
	}
	if v, ok := d.GetOk("ssh_over_v6_count"); ok {
		x := int64(v.(int))
		o.SetSshOverV6Count(x)
	}
	if v, ok := d.GetOk("syslog_over_v6_count"); ok {
		x := int64(v.(int))
		o.SetSyslogOverV6Count(x)
	}
	if v, ok := d.GetOk("tenant_count"); ok {
		x := int64(v.(int))
		o.SetTenantCount(x)
	}
	if v, ok := d.GetOk("tier_two_leaf_count"); ok {
		x := int64(v.(int))
		o.SetTierTwoLeafCount(x)
	}
	if v, ok := d.GetOk("twamp"); ok {
		x := (v.(string))
		o.SetTwamp(x)
	}
	if v, ok := d.GetOk("useg"); ok {
		x := (v.(string))
		o.SetUseg(x)
	}
	if v, ok := d.GetOk("vpod_count"); ok {
		x := int64(v.(int))
		o.SetVpodCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaFeatureUsageList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.NiatelemetryNiaFeatureUsageList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to NiatelemetryNiaFeatureUsage: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewNiatelemetryNiaFeatureUsageWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("apic_count", (s.GetApicCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ApicCount: %+v", err)
			}
			if err := d.Set("app_center_count", (s.GetAppCenterCount())); err != nil {
				return fmt.Errorf("error occurred while setting property AppCenterCount: %+v", err)
			}
			if err := d.Set("ave", (s.GetAve())); err != nil {
				return fmt.Errorf("error occurred while setting property Ave: %+v", err)
			}
			if err := d.Set("bd_count", (s.GetBdCount())); err != nil {
				return fmt.Errorf("error occurred while setting property BdCount: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("consistency_checker_app", (s.GetConsistencyCheckerApp())); err != nil {
				return fmt.Errorf("error occurred while setting property ConsistencyCheckerApp: %+v", err)
			}
			if err := d.Set("contract_count", (s.GetContractCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ContractCount: %+v", err)
			}
			if err := d.Set("dns_count", (s.GetDnsCount())); err != nil {
				return fmt.Errorf("error occurred while setting property DnsCount: %+v", err)
			}
			if err := d.Set("eigrp_count", (s.GetEigrpCount())); err != nil {
				return fmt.Errorf("error occurred while setting property EigrpCount: %+v", err)
			}
			if err := d.Set("epg_count", (s.GetEpgCount())); err != nil {
				return fmt.Errorf("error occurred while setting property EpgCount: %+v", err)
			}
			if err := d.Set("hsrp_count", (s.GetHsrpCount())); err != nil {
				return fmt.Errorf("error occurred while setting property HsrpCount: %+v", err)
			}
			if err := d.Set("ibgp_count", (s.GetIbgpCount())); err != nil {
				return fmt.Errorf("error occurred while setting property IbgpCount: %+v", err)
			}
			if err := d.Set("igmp_access_list_count", (s.GetIgmpAccessListCount())); err != nil {
				return fmt.Errorf("error occurred while setting property IgmpAccessListCount: %+v", err)
			}
			if err := d.Set("igmp_snoop", (s.GetIgmpSnoop())); err != nil {
				return fmt.Errorf("error occurred while setting property IgmpSnoop: %+v", err)
			}
			if err := d.Set("ip_epg_count", (s.GetIpEpgCount())); err != nil {
				return fmt.Errorf("error occurred while setting property IpEpgCount: %+v", err)
			}
			if err := d.Set("isis_count", (s.GetIsisCount())); err != nil {
				return fmt.Errorf("error occurred while setting property IsisCount: %+v", err)
			}
			if err := d.Set("l2_multicast", (s.GetL2Multicast())); err != nil {
				return fmt.Errorf("error occurred while setting property L2Multicast: %+v", err)
			}
			if err := d.Set("leaf_count", (s.GetLeafCount())); err != nil {
				return fmt.Errorf("error occurred while setting property LeafCount: %+v", err)
			}
			if err := d.Set("maintenance_mode_count", (s.GetMaintenanceModeCount())); err != nil {
				return fmt.Errorf("error occurred while setting property MaintenanceModeCount: %+v", err)
			}
			if err := d.Set("management_over_v6_count", (s.GetManagementOverV6Count())); err != nil {
				return fmt.Errorf("error occurred while setting property ManagementOverV6Count: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("nir", (s.GetNir())); err != nil {
				return fmt.Errorf("error occurred while setting property Nir: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("opflex_kubernetes_count", (s.GetOpflexKubernetesCount())); err != nil {
				return fmt.Errorf("error occurred while setting property OpflexKubernetesCount: %+v", err)
			}
			if err := d.Set("ospf_count", (s.GetOspfCount())); err != nil {
				return fmt.Errorf("error occurred while setting property OspfCount: %+v", err)
			}
			if err := d.Set("poe_count", (s.GetPoeCount())); err != nil {
				return fmt.Errorf("error occurred while setting property PoeCount: %+v", err)
			}
			if err := d.Set("qin_vni_tunnel_count", (s.GetQinVniTunnelCount())); err != nil {
				return fmt.Errorf("error occurred while setting property QinVniTunnelCount: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}
			if err := d.Set("remote_leaf_count", (s.GetRemoteLeafCount())); err != nil {
				return fmt.Errorf("error occurred while setting property RemoteLeafCount: %+v", err)
			}
			if err := d.Set("scvmm_count", (s.GetScvmmCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ScvmmCount: %+v", err)
			}
			if err := d.Set("shared_l3_out_count", (s.GetSharedL3OutCount())); err != nil {
				return fmt.Errorf("error occurred while setting property SharedL3OutCount: %+v", err)
			}
			if err := d.Set("smart_call_home", (s.GetSmartCallHome())); err != nil {
				return fmt.Errorf("error occurred while setting property SmartCallHome: %+v", err)
			}
			if err := d.Set("snmp", (s.GetSnmp())); err != nil {
				return fmt.Errorf("error occurred while setting property Snmp: %+v", err)
			}
			if err := d.Set("spine_count", (s.GetSpineCount())); err != nil {
				return fmt.Errorf("error occurred while setting property SpineCount: %+v", err)
			}
			if err := d.Set("ssh_over_v6_count", (s.GetSshOverV6Count())); err != nil {
				return fmt.Errorf("error occurred while setting property SshOverV6Count: %+v", err)
			}
			if err := d.Set("syslog_over_v6_count", (s.GetSyslogOverV6Count())); err != nil {
				return fmt.Errorf("error occurred while setting property SyslogOverV6Count: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("tenant_count", (s.GetTenantCount())); err != nil {
				return fmt.Errorf("error occurred while setting property TenantCount: %+v", err)
			}
			if err := d.Set("tier_two_leaf_count", (s.GetTierTwoLeafCount())); err != nil {
				return fmt.Errorf("error occurred while setting property TierTwoLeafCount: %+v", err)
			}
			if err := d.Set("twamp", (s.GetTwamp())); err != nil {
				return fmt.Errorf("error occurred while setting property Twamp: %+v", err)
			}
			if err := d.Set("useg", (s.GetUseg())); err != nil {
				return fmt.Errorf("error occurred while setting property Useg: %+v", err)
			}
			if err := d.Set("vpod_count", (s.GetVpodCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VpodCount: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
