package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIamOAuthToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIamOAuthTokenRead,
		Schema: map[string]*schema.Schema{
			"access_expiration_time": {
				Description: "Expiration time for the JWT token to which it can be used for api calls.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"app_registration": {
				Description: "A collection of references to the [iam.AppRegistration](mo://iam.AppRegistration) Managed Object.When this managed object is deleted, the referenced [iam.AppRegistration](mo://iam.AppRegistration) MO unsets its reference to this deleted MO.",
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
			"client_id": {
				Description: "The identifier of the registered application to which the token belongs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_ip_address": {
				Description: "The user agent IP address from which the auth token is launched.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"client_name": {
				Description: "The name of the registered application to which the token belongs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"expiration_time": {
				Description: "Expiration time for the JWT token to which it can be refreshed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_login_client": {
				Description: "The client address from which last login is initiated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_login_time": {
				Description: "The last login time for user.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission": {
				Description: "Permissions associated with the OAuth session.Permissions provides a way to assign roles to a user or user group to perform operations on object hierarchy.",
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
			"token_id": {
				Description: "Token identifier. Not the Access Token itself.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Description: "A collection of references to the [iam.User](mo://iam.User) Managed Object.When this managed object is deleted, the referenced [iam.User](mo://iam.User) MO unsets its reference to this deleted MO.",
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
			"user_meta": {
				Description: "User Device meta information.",
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
						"device_model": {
							Description: "Parsed device model from raw User-Agent.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"user_agent": {
							Description: "The value of the \"User-Agent\" HTTP header, as sent by the HTTP client when it initiate a session to Intersight. This can be used to identify the client operating system, browser type and browser version.Example - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36It is set when User successfully passed OAuth login flow and receives Access Token.",
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
func dataSourceIamOAuthTokenRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "iam/OAuthTokens"
	var o models.IamOAuthToken
	if v, ok := d.GetOk("access_expiration_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.AccessExpirationTime = x
	}
	if v, ok := d.GetOk("client_id"); ok {
		x := (v.(string))
		o.ClientID = x
	}
	if v, ok := d.GetOk("client_ip_address"); ok {
		x := (v.(string))
		o.ClientIPAddress = x
	}
	if v, ok := d.GetOk("client_name"); ok {
		x := (v.(string))
		o.ClientName = x
	}
	if v, ok := d.GetOk("expiration_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ExpirationTime = x
	}
	if v, ok := d.GetOk("last_login_client"); ok {
		x := (v.(string))
		o.LastLoginClient = x
	}
	if v, ok := d.GetOk("last_login_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.LastLoginTime = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("token_id"); ok {
		x := (v.(string))
		o.TokenID = x
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
			var s models.IamOAuthToken
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("access_expiration_time", (s.AccessExpirationTime).String()); err != nil {
				return err
			}

			if err := d.Set("app_registration", flattenMapIamAppRegistrationRef(s.AppRegistration, d)); err != nil {
				return err
			}
			if err := d.Set("client_id", (s.ClientID)); err != nil {
				return err
			}
			if err := d.Set("client_ip_address", (s.ClientIPAddress)); err != nil {
				return err
			}
			if err := d.Set("client_name", (s.ClientName)); err != nil {
				return err
			}

			if err := d.Set("expiration_time", (s.ExpirationTime).String()); err != nil {
				return err
			}
			if err := d.Set("last_login_client", (s.LastLoginClient)); err != nil {
				return err
			}

			if err := d.Set("last_login_time", (s.LastLoginTime).String()); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("permission", flattenMapIamPermissionRef(s.Permission, d)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("token_id", (s.TokenID)); err != nil {
				return err
			}

			if err := d.Set("user", flattenMapIamUserRef(s.User, d)); err != nil {
				return err
			}

			if err := d.Set("user_meta", flattenMapIamClientMeta(s.UserMeta, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
