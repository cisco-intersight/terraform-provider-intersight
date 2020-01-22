package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAssetDeviceRegistration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssetDeviceRegistrationRead,
		Schema: map[string]*schema.Schema{
			"api_version": {
				Description: "The version of the connector API, describes the capability of the connector's framework.If the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"access_key_id": {
				Description: "An identifier for the credential used by the device connector to authenticate with the Intersight web socket gateway.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"account": {
				Description: "The account to which the device has been claimed.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"app_partition_number": {
				Description: "The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"claimed_by_user": {
				Description: "The user who claimed the device for the account.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"claimed_by_user_name": {
				Description: "The name of the user who claimed the device for the account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"claimed_time": {
				Description: "The date and time at which the device was claimed to this account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_members": {
				Description: "The set of nodes within the devices cluster.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
			"device_claim": {
				Description: "The user who claimed the device for the account.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"device_configuration": {
				Description: "The devices current configuration.",
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
			"device_hostname": {
				Description: "The hostnames of the managed device. There can be multiple hostnames depending on the number of elements managed (i.e. HA clusters).",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"device_ip_address": {
				Description: "The IP Addresses of the managed device. There can be multiple management IPs depending on the number of elements managed (i.e. HA clusters) and in-band/out-of-band connectivity.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"domain_group": {
				Description: "The domain group to which the device has been assigned.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"execution_mode": {
				Description: "Indicates if the platform is an actual device or an emulated device for testing, demos, etc. Permitted values are [Normal, Emulator, ContainerEmulator].",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parent_connection": {
				Description: "The parent device of this device, through which this device connector is connected. If present the device will inherit its ownership through this object and any claim operation will be cascaded from it.e.g. A rack server may have a parent as the Fabric Interconnect cluster to which it is attached.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"parent_signature": {
				Description: "A signature generated by a parent device used to authenticate that a leaf device is connected through the parent.",
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
						"device_id": {
							Description: "The moid of the parent device registration.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"node_id": {
							Description: "The node identity of the parent device, corresponds to the parents ClusterMember.memberIdentity. Used on connect to establish through which device in a cluster the child is connected through.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"signature": {
							Description: "The result of signing the deviceId appended with the timeStamp fields with the devices private key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"pid": {
				Description: "The product Id of the managed device.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"platform_type": {
				Description: "The platform type on which device connector is executing.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"proxy_app": {
				Description: "The name of the app which will proxy the messages to the device connector.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"public_access_key": {
				Description: "The device connector's public key used by the cloud to authenticate a connection from the device connector. The public key is used to verify that the signature a device connector sends on connect has been signed by the connector's private key stored on the device's filesystem. Must be a PEM encoded RSA public key string.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"read_only": {
				Description: "Flag reported by devices to indicate an administrator of the device has disabled management operations of the device connector and only monitoring is permitted.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"security_token": {
				Description: "The devices current security token that can be used by a device administrator to claim the device.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"serial": {
				Description: "A list of serial numbers of the individual device elements (e.g. HA primary/secondary or cluster members) which are exposed as a single unit of management by the device connector.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
				Description: "The vendor of the managed device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceAssetDeviceRegistrationRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "asset/DeviceRegistrations"
	var o models.AssetDeviceRegistration
	if v, ok := d.GetOk("api_version"); ok {
		x := int64(v.(int))
		o.APIVersion = x
	}
	if v, ok := d.GetOk("access_key_id"); ok {
		x := (v.(string))
		o.AccessKeyID = x
	}
	if v, ok := d.GetOk("app_partition_number"); ok {
		x := int64(v.(int))
		o.AppPartitionNumber = x
	}
	if v, ok := d.GetOk("claimed_by_user_name"); ok {
		x := (v.(string))
		o.ClaimedByUserName = x
	}
	if v, ok := d.GetOk("claimed_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ClaimedTime = x
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
	if v, ok := d.GetOk("execution_mode"); ok {
		x := (v.(string))
		o.ExecutionMode = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.PlatformType = x
	}
	if v, ok := d.GetOk("proxy_app"); ok {
		x := (v.(string))
		o.ProxyApp = x
	}
	if v, ok := d.GetOk("public_access_key"); ok {
		x := (v.(string))
		o.PublicAccessKey = x
	}
	if v, ok := d.GetOk("read_only"); ok {
		x := (v.(bool))
		o.ReadOnly = &x
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x
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
			var s models.AssetDeviceRegistration
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("api_version", (s.APIVersion)); err != nil {
				return err
			}
			if err := d.Set("access_key_id", (s.AccessKeyID)); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
				return err
			}
			if err := d.Set("app_partition_number", (s.AppPartitionNumber)); err != nil {
				return err
			}

			if err := d.Set("claimed_by_user", flattenMapIamUserRef(s.ClaimedByUser, d)); err != nil {
				return err
			}
			if err := d.Set("claimed_by_user_name", (s.ClaimedByUserName)); err != nil {
				return err
			}

			if err := d.Set("claimed_time", (s.ClaimedTime).String()); err != nil {
				return err
			}

			if err := d.Set("cluster_members", flattenListAssetClusterMemberRef(s.ClusterMembers, d)); err != nil {
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

			if err := d.Set("device_claim", flattenMapAssetDeviceClaimRef(s.DeviceClaim, d)); err != nil {
				return err
			}

			if err := d.Set("device_configuration", flattenMapAssetDeviceConfigurationRef(s.DeviceConfiguration, d)); err != nil {
				return err
			}
			if err := d.Set("device_external_ip_address", (s.DeviceExternalIPAddress)); err != nil {
				return err
			}
			if err := d.Set("device_hostname", (s.DeviceHostname)); err != nil {
				return err
			}
			if err := d.Set("device_ip_address", (s.DeviceIPAddress)); err != nil {
				return err
			}

			if err := d.Set("domain_group", flattenMapIamDomainGroupRef(s.DomainGroup, d)); err != nil {
				return err
			}
			if err := d.Set("execution_mode", (s.ExecutionMode)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("parent_connection", flattenMapAssetDeviceRegistrationRef(s.ParentConnection, d)); err != nil {
				return err
			}

			if err := d.Set("parent_signature", flattenMapAssetParentConnectionSignature(s.ParentSignature, d)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("pid", (s.Pid)); err != nil {
				return err
			}
			if err := d.Set("platform_type", (s.PlatformType)); err != nil {
				return err
			}
			if err := d.Set("proxy_app", (s.ProxyApp)); err != nil {
				return err
			}
			if err := d.Set("public_access_key", (s.PublicAccessKey)); err != nil {
				return err
			}
			if err := d.Set("read_only", (s.ReadOnly)); err != nil {
				return err
			}

			if err := d.Set("security_token", flattenMapAssetSecurityTokenRef(s.SecurityToken, d)); err != nil {
				return err
			}
			if err := d.Set("serial", (s.Serial)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
