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

func dataSourceAssetClusterMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssetClusterMemberRead,
		Schema: map[string]*schema.Schema{
			"api_version": {
				Description: "The version of the connector API, describes the capability of the connector's framework.If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.",
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
			"connection_id": {
				Description: "The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_reason": {
				Description: "If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with the cloud.",
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
				Description: "The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between the cloud and device connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with the cloud.",
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
				Description: "A collection of references to the [asset.DeviceRegistration](mo://asset.DeviceRegistration) Managed Object.When this managed object is deleted, the referenced [asset.DeviceRegistration](mo://asset.DeviceRegistration) MO unsets its reference to this deleted MO.",
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
			"device_external_ip_address": {
				Description: "The IP Address of the managed device as seen from the cloud at the time of registration.This could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.",
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
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
												"common_name": {
													Description: "A required component that identifies a person or an object.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"country": {
													Description: "Identifier for the country in which the entity resides.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"locality": {
													Description: "Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"object_type": {
													Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"organization": {
													Description: "Identifier for the organization in which the entity resides.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"organizational_unit": {
													Description: "Identifier for a unit within the organization.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"state": {
													Description: "Identifier for the state or province of the entity.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
											},
										},
									},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
												"common_name": {
													Description: "A required component that identifies a person or an object.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"country": {
													Description: "Identifier for the country in which the entity resides.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"locality": {
													Description: "Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"object_type": {
													Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"organization": {
													Description: "Identifier for the organization in which the entity resides.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"organizational_unit": {
													Description: "Identifier for a unit within the organization.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString}},
												"state": {
													Description: "Identifier for the state or province of the entity.",
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
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
		},
	}
}
func dataSourceAssetClusterMemberRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "asset/ClusterMembers"
	var o models.AssetClusterMember
	if v, ok := d.GetOk("api_version"); ok {
		x := int64(v.(int))
		o.APIVersion = x
	}
	if v, ok := d.GetOk("app_partition_number"); ok {
		x := int64(v.(int))
		o.AppPartitionNumber = x
	}
	if v, ok := d.GetOk("connection_id"); ok {
		x := (v.(string))
		o.ConnectionID = x
	}
	if v, ok := d.GetOk("connection_reason"); ok {
		x := (v.(string))
		o.ConnectionReason = x
	}
	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.ConnectionStatus = x
	}
	if v, ok := d.GetOk("connection_status_last_change_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ConnectionStatusLastChangeTime = x
	}
	if v, ok := d.GetOk("connector_version"); ok {
		x := (v.(string))
		o.ConnectorVersion = x
	}
	if v, ok := d.GetOk("device_external_ip_address"); ok {
		x := (v.(string))
		o.DeviceExternalIPAddress = x
	}
	if v, ok := d.GetOk("leadership"); ok {
		x := (v.(string))
		o.Leadership = x
	}
	if v, ok := d.GetOk("member_identity"); ok {
		x := (v.(string))
		o.MemberIdentity = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("parent_cluster_member_identity"); ok {
		x := (v.(string))
		o.ParentClusterMemberIdentity = x
	}
	if v, ok := d.GetOk("proxy_app"); ok {
		x := (v.(string))
		o.ProxyApp = x
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
			var s models.AssetClusterMember
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("api_version", (s.APIVersion)); err != nil {
				return err
			}
			if err := d.Set("app_partition_number", (s.AppPartitionNumber)); err != nil {
				return err
			}
			if err := d.Set("connection_id", (s.ConnectionID)); err != nil {
				return err
			}
			if err := d.Set("connection_reason", (s.ConnectionReason)); err != nil {
				return err
			}
			if err := d.Set("connection_status", (s.ConnectionStatus)); err != nil {
				return err
			}

			if err := d.Set("connection_status_last_change_time", (s.ConnectionStatusLastChangeTime).String()); err != nil {
				return err
			}
			if err := d.Set("connector_version", (s.ConnectorVersion)); err != nil {
				return err
			}

			if err := d.Set("device", flattenMapAssetDeviceRegistrationRef(s.Device, d)); err != nil {
				return err
			}
			if err := d.Set("device_external_ip_address", (s.DeviceExternalIPAddress)); err != nil {
				return err
			}
			if err := d.Set("leadership", (s.Leadership)); err != nil {
				return err
			}
			if err := d.Set("member_identity", (s.MemberIdentity)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("parent_cluster_member_identity", (s.ParentClusterMemberIdentity)); err != nil {
				return err
			}
			if err := d.Set("proxy_app", (s.ProxyApp)); err != nil {
				return err
			}

			if err := d.Set("sudi", flattenMapAssetSudiInfo(s.Sudi, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
