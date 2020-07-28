package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAssetClusterMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssetClusterMemberRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"api_version": {
				Description: "The version of the connector API, describes the capability of the connector's framework.\nIf the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"app_partition_number": {
				Description: "The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.",
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
			"connection_id": {
				Description: "The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_reason": {
				Description: "If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_status": {
				Description: "The status of the persistent connection between the device connector and Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_status_last_change_time": {
				Description: "The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between Intersight and Device Connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connector_version": {
				Description: "The version of the device connector running on the managed device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device": {
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
			"device_external_ip_address": {
				Description: "The IP Address of the managed device as seen from Intersight at the time of registration.\nThis could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"leadership": {
				Description: "The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"locked_leader": {
				Description: "Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"member_identity": {
				Description: "The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time.",
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
			"parent_cluster_member_identity": {
				Description: "The member idenity of the cluster member through which this device is connected if applicable.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"proxy_app": {
				Description: "The name of the app which will proxy the messages to the device connector.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sudi": {
				Description: "The SUDI status read from the Trust Anchor Module on a device.",
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
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"pid": {
							Description: "The device model (PID) extracted from the X.509 SUDI Leaf Certificate.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"serial_number": {
							Description: "The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"signature": {
							Description: "The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"status": {
							Description: "The validation status of the device.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sudi_certificate": {
							Description: "The X.509 SUDI Leaf Certificate from the Trust Anchor Module. The certificate is serialized in PEM format (Base64 encoded DER certificate).",
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
									"issuer": {
										Description: "The X.509 distinguished name of the issuer of this certificate.",
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
												"common_name": {
													Description: "A required component that identifies a person or an object.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"country": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"locality": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"organization": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"organizational_unit": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"state": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
											},
										},
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"pem_certificate": {
										Description: "The base64 encoded certificate in PEM format.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"sha256_fingerprint": {
										Description: "The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"signature_algorithm": {
										Description: "Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"subject": {
										Description: "The X.509 distinguished name of the subject of this certificate.",
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
												"common_name": {
													Description: "A required component that identifies a person or an object.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"country": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"locality": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"organization": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"organizational_unit": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"state": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
											},
										},
									},
								},
							},
							Computed: true,
						},
					},
				},
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
		},
	}
}

func dataSourceAssetClusterMemberRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewAssetClusterMemberWithDefaults()
	if v, ok := d.GetOk("api_version"); ok {
		x := int64(v.(int))
		o.SetApiVersion(x)
	}
	if v, ok := d.GetOk("app_partition_number"); ok {
		x := int64(v.(int))
		o.SetAppPartitionNumber(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connection_id"); ok {
		x := (v.(string))
		o.SetConnectionId(x)
	}
	if v, ok := d.GetOk("connection_reason"); ok {
		x := (v.(string))
		o.SetConnectionReason(x)
	}
	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.SetConnectionStatus(x)
	}
	if v, ok := d.GetOk("connection_status_last_change_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetConnectionStatusLastChangeTime(x)
	}
	if v, ok := d.GetOk("connector_version"); ok {
		x := (v.(string))
		o.SetConnectorVersion(x)
	}
	if v, ok := d.GetOk("device_external_ip_address"); ok {
		x := (v.(string))
		o.SetDeviceExternalIpAddress(x)
	}
	if v, ok := d.GetOk("leadership"); ok {
		x := (v.(string))
		o.SetLeadership(x)
	}
	if v, ok := d.GetOk("locked_leader"); ok {
		x := (v.(bool))
		o.SetLockedLeader(x)
	}
	if v, ok := d.GetOk("member_identity"); ok {
		x := (v.(string))
		o.SetMemberIdentity(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("parent_cluster_member_identity"); ok {
		x := (v.(string))
		o.SetParentClusterMemberIdentity(x)
	}
	if v, ok := d.GetOk("proxy_app"); ok {
		x := (v.(string))
		o.SetProxyApp(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.AssetApi.GetAssetClusterMemberList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.AssetClusterMemberList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to AssetClusterMember: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewAssetClusterMemberWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("api_version", (s.GetApiVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property ApiVersion: %+v", err)
			}
			if err := d.Set("app_partition_number", (s.GetAppPartitionNumber())); err != nil {
				return fmt.Errorf("error occurred while setting property AppPartitionNumber: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("connection_id", (s.GetConnectionId())); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectionId: %+v", err)
			}
			if err := d.Set("connection_reason", (s.GetConnectionReason())); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectionReason: %+v", err)
			}
			if err := d.Set("connection_status", (s.GetConnectionStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectionStatus: %+v", err)
			}

			if err := d.Set("connection_status_last_change_time", (s.GetConnectionStatusLastChangeTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectionStatusLastChangeTime: %+v", err)
			}
			if err := d.Set("connector_version", (s.GetConnectorVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectorVersion: %+v", err)
			}

			if err := d.Set("device", flattenMapAssetDeviceRegistrationRelationship(s.GetDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Device: %+v", err)
			}
			if err := d.Set("device_external_ip_address", (s.GetDeviceExternalIpAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceExternalIpAddress: %+v", err)
			}
			if err := d.Set("leadership", (s.GetLeadership())); err != nil {
				return fmt.Errorf("error occurred while setting property Leadership: %+v", err)
			}
			if err := d.Set("locked_leader", (s.GetLockedLeader())); err != nil {
				return fmt.Errorf("error occurred while setting property LockedLeader: %+v", err)
			}
			if err := d.Set("member_identity", (s.GetMemberIdentity())); err != nil {
				return fmt.Errorf("error occurred while setting property MemberIdentity: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("parent_cluster_member_identity", (s.GetParentClusterMemberIdentity())); err != nil {
				return fmt.Errorf("error occurred while setting property ParentClusterMemberIdentity: %+v", err)
			}
			if err := d.Set("proxy_app", (s.GetProxyApp())); err != nil {
				return fmt.Errorf("error occurred while setting property ProxyApp: %+v", err)
			}

			if err := d.Set("sudi", flattenMapAssetSudiInfo(s.GetSudi(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Sudi: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
