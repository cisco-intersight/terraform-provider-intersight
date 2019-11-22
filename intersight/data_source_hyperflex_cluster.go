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
				Description: "The alarms that have been raised for this HyperFlex cluster.New alarms are added to this collection, and existing alarms are updated if the severity changes.Deleted alarms are not removed but are cleared by marking them as green.",
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
				Description: "The deployment type of the HyperFlex cluster.The cluster can have one of the following configurations:1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,the deployment type is set as 'NA' (not available).",
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
				Description: "The number of yellow (warning) and red (critical) alarms stored as an aggregate.The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"health": {
				Description: "The health of the HyperFlex cluster.Detailed information concerning the cluster health, which includes cluster operational status, resiliency health status,number of node and disk failues tolerable, and the status of services such as the ZooKeeper ensemble and arbitration service.This relationship is only populated for devices with HyperFlex Data Platform 3.0+. For clusters running an older version, referto the Summary property of the hyperflex/Clusters API.",
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
			"hx_version": {
				Description: "The HyperFlex Data Platform version of this cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hxdp_build_version": {
				Description: "The version and build number of the HyperFlex Data Platform for this cluster.After a cluster upgrade, this version string will be updated on the next inventory cycle to reflectthe newly installed version.",
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
				Description: "The nodes belonging to this HyperFlex cluster.The node object contains inventory information about a specific HyperFlex node, such as host IP address,hypervisor type and version, and operational status.",
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
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"summary": {
				Description: "The summary of HyperFlex cluster health, storage, and number of nodes.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active_nodes": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
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
						"boottime": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"cluster_access_policy": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"compression_savings": {
							Description: "",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
						"data_replication_compliance": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"data_replication_factor": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"deduplication_savings": {
							Description: "",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
						"downtime": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"free_capacity": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"healing_info": {
							Description: "",
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
									"estimated_completion_time_in_seconds": {
										Description: "",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"in_progress": {
										Description: "",
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
									},
									"messages": {
										Description: "",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_size": {
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
									"percent_complete": {
										Description: "",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
								},
							},
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
						"resiliency_details_size": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"resiliency_info": {
							Description: "",
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
									"hdd_failures_tolerable": {
										Description: "",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"messages": {
										Description: "",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"messages_size": {
										Description: "",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"node_failures_tolerable": {
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
									"ssd_failures_tolerable": {
										Description: "",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"space_status": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"state": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"total_capacity": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"total_savings": {
							Description: "",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
						"uptime": {
							Description: "",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"used_capacity": {
							Description: "",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
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

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}

			if err := d.Set("summary", flattenMapHyperflexSummary(s.Summary, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
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
