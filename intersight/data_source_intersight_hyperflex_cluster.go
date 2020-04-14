package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHyperflexCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHyperflexClusterRead,
		Schema: map[string]*schema.Schema{
			"alarm": {
				Description: "The alarms that have been raised for this HyperFlex cluster.\nNew alarms are added to this collection, and existing alarms are updated if the severity changes.\nDeleted alarms are not removed but are cleared by marking them as green.",
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
			"capacity_runway": {
				Description: "The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%.\nDefault value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_name": {
				Description: "The name of this HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_type": {
				Description: "The storage type of this cluster (All Flash or Hybrid).",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"cluster_uuid": {
				Description: "The unique identifier for this HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"compute_node_count": {
				Description: "The number of compute nodes that belong to this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"converged_node_count": {
				Description: "The number of converged nodes that belong to this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"deployment_type": {
				Description: "The deployment type of the HyperFlex cluster.\nThe cluster can have one of the following configurations:\n1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.\n2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.\n3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.\nIf the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,\nthe deployment type is set as 'NA' (not available).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Description: "The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"flt_aggr": {
				Description: "The number of yellow (warning) and red (critical) alarms stored as an aggregate.\nThe first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"health": {
				Description: "The health of the HyperFlex cluster.\nDetailed information concerning the cluster health, which includes cluster operational status, resiliency health status,\nnumber of node and disk failues tolerable, and the status of services such as the ZooKeeper ensemble and arbitration service.\nThis relationship is only populated for devices with HyperFlex Data Platform 3.0+. For clusters running an older version, refer\nto the Summary property of the hyperflex/Clusters API.",
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
			"hx_version": {
				Description: "The HyperFlex Data Platform version of this cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hxdp_build_version": {
				Description: "The version and build number of the HyperFlex Data Platform for this cluster.\nAfter a cluster upgrade, this version string will be updated on the next inventory cycle to reflect\nthe newly installed version.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hypervisor_type": {
				Description: "The type of hypervisor running on this cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hypervisor_version": {
				Description: "The version of hypervisor running on this cluster.",
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
			"nodes": {
				Description: "The nodes belonging to this HyperFlex cluster.\nThe node object contains inventory information about a specific HyperFlex node, such as host IP address,\nhypervisor type and version, and operational status.",
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
				Description: "The registration that represents this HyperFlex cluster's connection to Intersight.",
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
			"summary": {
				Description: "The summary of HyperFlex cluster health, storage, and number of nodes.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_nodes": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"boottime": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"cluster_access_policy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"compression_savings": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"data_replication_compliance": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"data_replication_factor": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"deduplication_savings": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"downtime": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"free_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"healing_info": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
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
									"estimated_completion_time_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"in_progress": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"messages": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"percent_complete": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"resiliency_details_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"resiliency_info": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
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
									"hdd_failures_tolerable": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"messages": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"node_failures_tolerable": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"ssd_failures_tolerable": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"space_status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"total_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"total_savings": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"uptime": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"used_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
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
			"utilization_percentage": {
				Description: "The storage utilization percentage is computed based on total capacity and current capacity utilization.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"utilization_trend_percentage": {
				Description: "The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"vm_count": {
				Description: "The number of virtual machines present on this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceHyperflexClusterRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "hyperflex/Clusters"
	var o models.HyperflexCluster
	if v, ok := d.GetOk("capacity_runway"); ok {
		x := int64(v.(int))
		o.CapacityRunway = x
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x
	}
	if v, ok := d.GetOk("cluster_name"); ok {
		x := (v.(string))
		o.ClusterName = x
	}
	if v, ok := d.GetOk("cluster_type"); ok {
		x := int64(v.(int))
		o.ClusterType = x
	}
	if v, ok := d.GetOk("cluster_uuid"); ok {
		x := (v.(string))
		o.ClusterUUID = x
	}
	if v, ok := d.GetOk("compute_node_count"); ok {
		x := int64(v.(int))
		o.ComputeNodeCount = x
	}
	if v, ok := d.GetOk("converged_node_count"); ok {
		x := int64(v.(int))
		o.ConvergedNodeCount = x
	}
	if v, ok := d.GetOk("deployment_type"); ok {
		x := (v.(string))
		o.DeploymentType = x
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.DeviceID = x
	}
	if v, ok := d.GetOk("flt_aggr"); ok {
		x := int64(v.(int))
		o.FltAggr = x
	}
	if v, ok := d.GetOk("hx_version"); ok {
		x := (v.(string))
		o.HxVersion = x
	}
	if v, ok := d.GetOk("hxdp_build_version"); ok {
		x := (v.(string))
		o.HxdpBuildVersion = x
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.HypervisorType = x
	}
	if v, ok := d.GetOk("hypervisor_version"); ok {
		x := (v.(string))
		o.HypervisorVersion = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("utilization_percentage"); ok {
		x := v.(float32)
		o.UtilizationPercentage = x
	}
	if v, ok := d.GetOk("utilization_trend_percentage"); ok {
		x := v.(float32)
		o.UtilizationTrendPercentage = x
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
			var s models.HyperflexCluster
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("alarm", flattenListHyperflexAlarmRef(s.Alarm, d)); err != nil {
				return err
			}
			if err := d.Set("capacity_runway", (s.CapacityRunway)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassID)); err != nil {
				return err
			}
			if err := d.Set("cluster_name", (s.ClusterName)); err != nil {
				return err
			}
			if err := d.Set("cluster_type", (s.ClusterType)); err != nil {
				return err
			}
			if err := d.Set("cluster_uuid", (s.ClusterUUID)); err != nil {
				return err
			}
			if err := d.Set("compute_node_count", (s.ComputeNodeCount)); err != nil {
				return err
			}
			if err := d.Set("converged_node_count", (s.ConvergedNodeCount)); err != nil {
				return err
			}
			if err := d.Set("deployment_type", (s.DeploymentType)); err != nil {
				return err
			}
			if err := d.Set("device_id", (s.DeviceID)); err != nil {
				return err
			}
			if err := d.Set("flt_aggr", (s.FltAggr)); err != nil {
				return err
			}

			if err := d.Set("health", flattenMapHyperflexHealthRef(s.Health, d)); err != nil {
				return err
			}
			if err := d.Set("hx_version", (s.HxVersion)); err != nil {
				return err
			}
			if err := d.Set("hxdp_build_version", (s.HxdpBuildVersion)); err != nil {
				return err
			}
			if err := d.Set("hypervisor_type", (s.HypervisorType)); err != nil {
				return err
			}
			if err := d.Set("hypervisor_version", (s.HypervisorVersion)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}

			if err := d.Set("nodes", flattenListHyperflexNodeRef(s.Nodes, d)); err != nil {
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

			if err := d.Set("summary", flattenMapHyperflexSummary(s.Summary, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("utilization_percentage", (s.UtilizationPercentage)); err != nil {
				return err
			}
			if err := d.Set("utilization_trend_percentage", (s.UtilizationTrendPercentage)); err != nil {
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
