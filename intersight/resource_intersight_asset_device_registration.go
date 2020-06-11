package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAssetDeviceRegistration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssetDeviceRegistrationCreate,
		Read:   resourceAssetDeviceRegistrationRead,
		Update: resourceAssetDeviceRegistrationUpdate,
		Delete: resourceAssetDeviceRegistrationDelete,
		Schema: map[string]*schema.Schema{
			"api_version": {
				Description: "The version of the connector API, describes the capability of the connector's framework.\nIf the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"device_external_ip_address": {
				Description: "The IP Address of the managed device as seen from Intersight at the time of registration.\nThis could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"execution_mode": {
				Description: "Indicates if the platform is an actual device or an emulated device for testing, demos, etc. Permitted values are [Normal, Emulator, ContainerEmulator].",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parent_connection": {
				Description: "The parent device of this device, through which this device connector is connected. If present the device will inherit its ownership through this object and any claim operation will be cascaded from it.\ne.g. A rack server may have a parent as the Fabric Interconnect cluster to which it is attached.",
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
				ConfigMode: schema.SchemaConfigModeAttr,
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
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
				ForceNew:    true,
			},
			"proxy_app": {
				Description: "The name of the app which will proxy the messages to the device connector.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"public_access_key": {
				Description: "The device connector's public key used by Intersight to authenticate a connection from the device connector. The public key is used to verify that the signature a device connector sends on connect has been signed by the connector's private key stored on the device's filesystem. Must be a PEM encoded RSA public key string.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
func resourceAssetDeviceRegistrationCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.AssetDeviceRegistration
	if v, ok := d.GetOk("api_version"); ok {
		x := int64(v.(int))
		o.APIVersion = x

	}

	if v, ok := d.GetOk("access_key_id"); ok {
		x := (v.(string))
		o.AccessKeyID = x

	}

	if v, ok := d.GetOk("account"); ok {
		p := models.IamAccountRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamAccountRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.Account = &x
		}

	}

	if v, ok := d.GetOk("app_partition_number"); ok {
		x := int64(v.(int))
		o.AppPartitionNumber = x

	}

	if v, ok := d.GetOk("claimed_by_user"); ok {
		p := models.IamUserRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamUserRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ClaimedByUser = &x
		}

	}

	if v, ok := d.GetOk("claimed_by_user_name"); ok {
		x := (v.(string))
		o.ClaimedByUserName = x

	}

	if v, ok := d.GetOk("claimed_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ClaimedTime = x

	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x

	}

	if v, ok := d.GetOk("cluster_members"); ok {
		x := make([]*models.AssetClusterMemberRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.AssetClusterMemberRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.ClusterMembers = x

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

	if v, ok := d.GetOk("device_claim"); ok {
		p := models.AssetDeviceClaimRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetDeviceClaimRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.DeviceClaim = &x
		}

	}

	if v, ok := d.GetOk("device_configuration"); ok {
		p := models.AssetDeviceConfigurationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetDeviceConfigurationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.DeviceConfiguration = &x
		}

	}

	if v, ok := d.GetOk("device_external_ip_address"); ok {
		x := (v.(string))
		o.DeviceExternalIPAddress = x

	}

	if v, ok := d.GetOk("device_hostname"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.DeviceHostname = x

	}

	if v, ok := d.GetOk("device_ip_address"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.DeviceIPAddress = x

	}

	if v, ok := d.GetOk("domain_group"); ok {
		p := models.IamDomainGroupRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamDomainGroupRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.DomainGroup = &x
		}

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

	if v, ok := d.GetOk("parent_connection"); ok {
		p := models.AssetDeviceRegistrationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetDeviceRegistrationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ParentConnection = &x
		}

	}

	if v, ok := d.GetOk("parent_signature"); ok {
		p := models.AssetParentConnectionSignature{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetParentConnectionSignature{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["device_id"]; ok {
				{
					x := (v.(string))
					o.DeviceID = x
				}
			}
			if v, ok := l["node_id"]; ok {
				{
					x := (v.(string))
					o.NodeID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["signature"]; ok {
				{
					x := []byte(v.(string))
					o.Signature = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ParentSignature = &x
		}

	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]*models.MoBaseMoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoBaseMoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.PermissionResources = x

	}

	if v, ok := d.GetOk("pid"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.Pid = x

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

	if v, ok := d.GetOkExists("read_only"); ok {
		x := v.(bool)
		o.ReadOnly = &x
	}

	if v, ok := d.GetOk("security_token"); ok {
		p := models.AssetSecurityTokenRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetSecurityTokenRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.SecurityToken = &x
		}

	}

	if v, ok := d.GetOk("serial"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.Serial = x

	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO1P1.MoTagAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x

	}

	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = x

	}

	url := "asset/DeviceRegistrations"
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceAssetDeviceRegistrationRead(d, meta)
}

func resourceAssetDeviceRegistrationRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "asset/DeviceRegistrations" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.AssetDeviceRegistration
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
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

	if err := d.Set("class_id", (s.ClassID)); err != nil {
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

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceAssetDeviceRegistrationUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.AssetDeviceRegistration
	if d.HasChange("api_version") {
		v := d.Get("api_version")
		x := int64(v.(int))
		o.APIVersion = x
	}

	if d.HasChange("access_key_id") {
		v := d.Get("access_key_id")
		x := (v.(string))
		o.AccessKeyID = x
	}

	if d.HasChange("account") {
		v := d.Get("account")
		p := models.IamAccountRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamAccountRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.Account = &x
		}
	}

	if d.HasChange("app_partition_number") {
		v := d.Get("app_partition_number")
		x := int64(v.(int))
		o.AppPartitionNumber = x
	}

	if d.HasChange("claimed_by_user") {
		v := d.Get("claimed_by_user")
		p := models.IamUserRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamUserRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ClaimedByUser = &x
		}
	}

	if d.HasChange("claimed_by_user_name") {
		v := d.Get("claimed_by_user_name")
		x := (v.(string))
		o.ClaimedByUserName = x
	}

	if d.HasChange("claimed_time") {
		v := d.Get("claimed_time")
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ClaimedTime = x
	}

	if d.HasChange("class_id") {
		v := d.Get("class_id")
		x := (v.(string))
		o.ClassID = x
	}

	if d.HasChange("cluster_members") {
		v := d.Get("cluster_members")
		x := make([]*models.AssetClusterMemberRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.AssetClusterMemberRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.ClusterMembers = x
	}

	if d.HasChange("connection_id") {
		v := d.Get("connection_id")
		x := (v.(string))
		o.ConnectionID = x
	}

	if d.HasChange("connection_reason") {
		v := d.Get("connection_reason")
		x := (v.(string))
		o.ConnectionReason = x
	}

	if d.HasChange("connection_status") {
		v := d.Get("connection_status")
		x := (v.(string))
		o.ConnectionStatus = x
	}

	if d.HasChange("connection_status_last_change_time") {
		v := d.Get("connection_status_last_change_time")
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ConnectionStatusLastChangeTime = x
	}

	if d.HasChange("connector_version") {
		v := d.Get("connector_version")
		x := (v.(string))
		o.ConnectorVersion = x
	}

	if d.HasChange("device_claim") {
		v := d.Get("device_claim")
		p := models.AssetDeviceClaimRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetDeviceClaimRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.DeviceClaim = &x
		}
	}

	if d.HasChange("device_configuration") {
		v := d.Get("device_configuration")
		p := models.AssetDeviceConfigurationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetDeviceConfigurationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.DeviceConfiguration = &x
		}
	}

	if d.HasChange("device_external_ip_address") {
		v := d.Get("device_external_ip_address")
		x := (v.(string))
		o.DeviceExternalIPAddress = x
	}

	if d.HasChange("device_hostname") {
		v := d.Get("device_hostname")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.DeviceHostname = x
	}

	if d.HasChange("device_ip_address") {
		v := d.Get("device_ip_address")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.DeviceIPAddress = x
	}

	if d.HasChange("domain_group") {
		v := d.Get("domain_group")
		p := models.IamDomainGroupRef{}
		if len(v.([]interface{})) > 0 {
			o := models.IamDomainGroupRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.DomainGroup = &x
		}
	}

	if d.HasChange("execution_mode") {
		v := d.Get("execution_mode")
		x := (v.(string))
		o.ExecutionMode = x
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.Moid = x
	}

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
	}

	if d.HasChange("parent_connection") {
		v := d.Get("parent_connection")
		p := models.AssetDeviceRegistrationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetDeviceRegistrationRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ParentConnection = &x
		}
	}

	if d.HasChange("parent_signature") {
		v := d.Get("parent_signature")
		p := models.AssetParentConnectionSignature{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetParentConnectionSignature{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["device_id"]; ok {
				{
					x := (v.(string))
					o.DeviceID = x
				}
			}
			if v, ok := l["node_id"]; ok {
				{
					x := (v.(string))
					o.NodeID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["signature"]; ok {
				{
					x := []byte(v.(string))
					o.Signature = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.ParentSignature = &x
		}
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]*models.MoBaseMoRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoBaseMoRef{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["moid"]; ok {
					{
						x := (v.(string))
						o.Moid = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["selector"]; ok {
					{
						x := (v.(string))
						o.Selector = x
					}
				}
				x = append(x, &o)
			}
		}
		o.PermissionResources = x
	}

	if d.HasChange("pid") {
		v := d.Get("pid")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.Pid = x
	}

	if d.HasChange("platform_type") {
		v := d.Get("platform_type")
		x := (v.(string))
		o.PlatformType = x
	}

	if d.HasChange("proxy_app") {
		v := d.Get("proxy_app")
		x := (v.(string))
		o.ProxyApp = x
	}

	if d.HasChange("public_access_key") {
		v := d.Get("public_access_key")
		x := (v.(string))
		o.PublicAccessKey = x
	}

	if d.HasChange("read_only") {
		v := d.Get("read_only")
		x := (v.(bool))
		o.ReadOnly = &x
	}

	if d.HasChange("security_token") {
		v := d.Get("security_token")
		p := models.AssetSecurityTokenRef{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetSecurityTokenRef{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.Moid = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.Selector = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.SecurityToken = &x
		}
	}

	if d.HasChange("serial") {
		v := d.Get("serial")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.Serial = x
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO1P1.MoTagAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x
	}

	if d.HasChange("vendor") {
		v := d.Get("vendor")
		x := (v.(string))
		o.Vendor = x
	}

	url := "asset/DeviceRegistrations" + "/" + d.Id()
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendUpdateRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceAssetDeviceRegistrationRead(d, meta)
}

func resourceAssetDeviceRegistrationDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "asset/DeviceRegistrations" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
