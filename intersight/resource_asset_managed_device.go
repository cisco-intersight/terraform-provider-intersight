package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func resourceAssetManagedDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssetManagedDeviceCreate,
		Read:   resourceAssetManagedDeviceRead,
		Update: resourceAssetManagedDeviceUpdate,
		Delete: resourceAssetManagedDeviceDelete,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "ManagedDevice to Account MO relationship.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"credential": {
				Description: "Credentials to manage Managed Device.",
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
						"is_password_set": {
							Description: "",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Description: "Password for the Managed Device.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"username": {
							Description: "Username for the Managed Device. Format and restrictions are not enforced here but usually follow the ManagedDevice requirements.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"device_connector_manager": {
				Description: "Device Connector Manager (Intersight Appliance Device Connector tagged as 'Intersight Assist') within the asset Device Registration.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"device_type": {
				Description: "Type of the Device such as VMware, Pure Storage.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "PureStorage",
			},
			"is_enabled": {
				Description: "Device is Enabled/Disabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"management_address": {
				Description: "Management address of the device. It can be IPv4, IPv6 or FQDN.",
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
			"name": {
				Description: "Name defined by the administrator for easier identification.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "ManagedDevice once auto claimed within the asset Device Registration.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"status": {
				Description: "Status of communication releated to Managed Device.",
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
						"cloud_port": {
							Description: "Port used for the connection to the Cloud by the Device Connector for the Managed Device.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"connection_failure_reason": {
							Description: "Maintains the reason for the failure of connection to the Device in case of connection failure.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"connection_status": {
							Description: "Maintains the status of the connection to the Device.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Unknown",
						},
						"error_code": {
							Description: "Maintains code related to error from Device Connector, if any.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"error_reason": {
							Description: "Maintains the reason for the error from Device Connector, if any.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"process_id": {
							Description: "Maintains the Process pid of the the Device Connector for the Managed Device.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"server_port": {
							Description: "Port used for receiving requests from Device Connector Manager by the Device Connector for the Managed Device.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"state": {
							Description: "Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "New",
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
		},
	}
}
func resourceAssetManagedDeviceCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.AssetManagedDevice
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
		o.Account = &x

	}

	if v, ok := d.GetOk("credential"); ok {
		p := models.CommCredential{}
		if len(v.([]interface{})) > 0 {
			o := models.CommCredential{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.CommCredentialAO0P0.CommCredentialAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsPasswordSet = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.Password = x
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.Username = x
				}
			}

			p = o
		}
		x := p
		o.Credential = &x

	}

	if v, ok := d.GetOk("device_connector_manager"); ok {
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
		o.DeviceConnectorManager = &x

	}

	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.DeviceType = &x

	}

	if v, ok := d.GetOkExists("is_enabled"); ok {
		x := v.(bool)
		o.IsEnabled = &x
	}

	if v, ok := d.GetOk("management_address"); ok {
		x := (v.(string))
		o.ManagementAddress = x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

	}

	if v, ok := d.GetOk("registered_device"); ok {
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
		o.RegisteredDevice = &x

	}

	if v, ok := d.GetOk("status"); ok {
		p := models.AssetManagedDeviceStatus{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetManagedDeviceStatus{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AssetManagedDeviceStatusAO0P0.AssetManagedDeviceStatusAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["cloud_port"]; ok {
				{
					x := int64(v.(int))
					o.CloudPort = x
				}
			}
			if v, ok := l["connection_failure_reason"]; ok {
				{
					x := (v.(string))
					o.ConnectionFailureReason = x
				}
			}
			if v, ok := l["connection_status"]; ok {
				{
					x := (v.(string))
					o.ConnectionStatus = &x
				}
			}
			if v, ok := l["error_code"]; ok {
				{
					x := int64(v.(int))
					o.ErrorCode = x
				}
			}
			if v, ok := l["error_reason"]; ok {
				{
					x := (v.(string))
					o.ErrorReason = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["process_id"]; ok {
				{
					x := int64(v.(int))
					o.ProcessID = x
				}
			}
			if v, ok := l["server_port"]; ok {
				{
					x := int64(v.(int))
					o.ServerPort = x
				}
			}
			if v, ok := l["state"]; ok {
				{
					x := (v.(string))
					o.State = &x
				}
			}

			p = o
		}
		x := p
		o.Status = &x

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
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
						}
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

	url := "asset/ManagedDevices"
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
	return resourceAssetManagedDeviceRead(d, meta)
}

func resourceAssetManagedDeviceRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "asset/ManagedDevices" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.AssetManagedDevice
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
		return err
	}

	if err := d.Set("credential", flattenMapCommCredential(s.Credential, d)); err != nil {
		return err
	}

	if err := d.Set("device_connector_manager", flattenMapAssetDeviceRegistrationRef(s.DeviceConnectorManager, d)); err != nil {
		return err
	}

	if err := d.Set("device_type", (s.DeviceType)); err != nil {
		return err
	}

	if err := d.Set("is_enabled", (s.IsEnabled)); err != nil {
		return err
	}

	if err := d.Set("management_address", (s.ManagementAddress)); err != nil {
		return err
	}

	if err := d.Set("moid", (s.Moid)); err != nil {
		return err
	}

	if err := d.Set("name", (s.Name)); err != nil {
		return err
	}

	if err := d.Set("object_type", (s.ObjectType)); err != nil {
		return err
	}

	if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
		return err
	}

	if err := d.Set("status", flattenMapAssetManagedDeviceStatus(s.Status, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceAssetManagedDeviceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.AssetManagedDevice
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
		o.Account = &x
	}

	if d.HasChange("credential") {
		v := d.Get("credential")
		p := models.CommCredential{}
		if len(v.([]interface{})) > 0 {
			o := models.CommCredential{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.CommCredentialAO0P0.CommCredentialAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["is_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsPasswordSet = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["password"]; ok {
				{
					x := (v.(string))
					o.Password = x
				}
			}
			if v, ok := l["username"]; ok {
				{
					x := (v.(string))
					o.Username = x
				}
			}

			p = o
		}
		x := p
		o.Credential = &x
	}

	if d.HasChange("device_connector_manager") {
		v := d.Get("device_connector_manager")
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
		o.DeviceConnectorManager = &x
	}

	if d.HasChange("device_type") {
		v := d.Get("device_type")
		x := (v.(string))
		o.DeviceType = &x
	}

	if d.HasChange("is_enabled") {
		v := d.Get("is_enabled")
		x := (v.(bool))
		o.IsEnabled = &x
	}

	if d.HasChange("management_address") {
		v := d.Get("management_address")
		x := (v.(string))
		o.ManagementAddress = x
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.Moid = x
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.Name = x
	}

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
	}

	if d.HasChange("registered_device") {
		v := d.Get("registered_device")
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
		o.RegisteredDevice = &x
	}

	if d.HasChange("status") {
		v := d.Get("status")
		p := models.AssetManagedDeviceStatus{}
		if len(v.([]interface{})) > 0 {
			o := models.AssetManagedDeviceStatus{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AssetManagedDeviceStatusAO0P0.AssetManagedDeviceStatusAO0P0 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["cloud_port"]; ok {
				{
					x := int64(v.(int))
					o.CloudPort = x
				}
			}
			if v, ok := l["connection_failure_reason"]; ok {
				{
					x := (v.(string))
					o.ConnectionFailureReason = x
				}
			}
			if v, ok := l["connection_status"]; ok {
				{
					x := (v.(string))
					o.ConnectionStatus = &x
				}
			}
			if v, ok := l["error_code"]; ok {
				{
					x := int64(v.(int))
					o.ErrorCode = x
				}
			}
			if v, ok := l["error_reason"]; ok {
				{
					x := (v.(string))
					o.ErrorReason = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["process_id"]; ok {
				{
					x := int64(v.(int))
					o.ProcessID = x
				}
			}
			if v, ok := l["server_port"]; ok {
				{
					x := int64(v.(int))
					o.ServerPort = x
				}
			}
			if v, ok := l["state"]; ok {
				{
					x := (v.(string))
					o.State = &x
				}
			}

			p = o
		}
		x := p
		o.Status = &x
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
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
						}
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

	url := "asset/ManagedDevices" + "/" + d.Id()
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
	return resourceAssetManagedDeviceRead(d, meta)
}

func resourceAssetManagedDeviceDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "asset/ManagedDevices" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
