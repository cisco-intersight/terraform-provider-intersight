package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIamTrustPoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceIamTrustPointCreate,
		Read:   resourceIamTrustPointRead,
		Delete: resourceIamTrustPointDelete,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "The account associated with the Trustpoint.",
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"certificates": {
				Description: "The collection of certificates in X509 certificate format.\nThis was obtained by parsing the chain property which holds the base 64 encoded chain of certificates.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
							ForceNew:         true,
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
										ForceNew:         true,
									},
									"common_name": {
										Description: "A required component that identifies a person or an object.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"country": {
										Description: "Identifier for the country in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"locality": {
										Description: "Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"organization": {
										Description: "Identifier for the organization in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"organizational_unit": {
										Description: "Identifier for a unit within the organization.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"state": {
										Description: "Identifier for the state or province of the entity.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							ForceNew:   true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"pem_certificate": {
							Description: "The base64 encoded certificate in PEM format.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"sha256_fingerprint": {
							Description: "The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"signature_algorithm": {
							Description: "Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
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
										ForceNew:         true,
									},
									"common_name": {
										Description: "A required component that identifies a person or an object.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"country": {
										Description: "Identifier for the country in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"locality": {
										Description: "Identifier for the place where the entry resides. The locality can be a city, county, township, or other geographic region.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"organization": {
										Description: "Identifier for the organization in which the entity resides.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"organizational_unit": {
										Description: "Identifier for a unit within the organization.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
									"state": {
										Description: "Identifier for the state or province of the entity.",
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString}, ForceNew: true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							ForceNew:   true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"chain": {
				Description: "The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format.",
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
				ForceNew:    true,
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
							ForceNew:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
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
							ForceNew:         true,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
		},
	}
}
func resourceIamTrustPointCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.IamTrustPoint
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

	if v, ok := d.GetOk("certificates"); ok {
		x := make([]*models.X509Certificate, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.X509Certificate{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.X509CertificateAO1P1.X509CertificateAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["issuer"]; ok {
					{
						p := models.PkixDistinguishedName{}
						if len(v.([]interface{})) > 0 {
							o := models.PkixDistinguishedName{}
							l := (v.([]interface{})[0]).(map[string]interface{})
							if v, ok := l["additional_properties"]; ok {
								{
									x := []byte(v.(string))
									var x1 interface{}
									err := json.Unmarshal(x, &x1)
									if err == nil && x1 != nil {
										o.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1 = x1.(map[string]interface{})
									}
								}
							}
							if v, ok := l["common_name"]; ok {
								{
									x := (v.(string))
									o.CommonName = x
								}
							}
							if v, ok := l["country"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.Country = x
								}
							}
							if v, ok := l["locality"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.Locality = x
								}
							}
							if v, ok := l["object_type"]; ok {
								{
									x := (v.(string))
									o.ObjectType = x
								}
							}
							if v, ok := l["organization"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.Organization = x
								}
							}
							if v, ok := l["organizational_unit"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.OrganizationalUnit = x
								}
							}
							if v, ok := l["state"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.State = x
								}
							}

							p = o
						}
						x := p
						o.Issuer = &x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["pem_certificate"]; ok {
					{
						x := (v.(string))
						o.PemCertificate = x
					}
				}
				if v, ok := l["sha256_fingerprint"]; ok {
					{
						x := (v.(string))
						o.Sha256Fingerprint = x
					}
				}
				if v, ok := l["signature_algorithm"]; ok {
					{
						x := (v.(string))
						o.SignatureAlgorithm = x
					}
				}
				if v, ok := l["subject"]; ok {
					{
						p := models.PkixDistinguishedName{}
						if len(v.([]interface{})) > 0 {
							o := models.PkixDistinguishedName{}
							l := (v.([]interface{})[0]).(map[string]interface{})
							if v, ok := l["additional_properties"]; ok {
								{
									x := []byte(v.(string))
									var x1 interface{}
									err := json.Unmarshal(x, &x1)
									if err == nil && x1 != nil {
										o.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1 = x1.(map[string]interface{})
									}
								}
							}
							if v, ok := l["common_name"]; ok {
								{
									x := (v.(string))
									o.CommonName = x
								}
							}
							if v, ok := l["country"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.Country = x
								}
							}
							if v, ok := l["locality"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.Locality = x
								}
							}
							if v, ok := l["object_type"]; ok {
								{
									x := (v.(string))
									o.ObjectType = x
								}
							}
							if v, ok := l["organization"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.Organization = x
								}
							}
							if v, ok := l["organizational_unit"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.OrganizationalUnit = x
								}
							}
							if v, ok := l["state"]; ok {
								{
									x := make([]string, 0)
									y := reflect.ValueOf(v)
									for i := 0; i < y.Len(); i++ {
										x = append(x, y.Index(i).Interface().(string))
									}
									o.State = x
								}
							}

							p = o
						}
						x := p
						o.Subject = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.Certificates = x

	}

	if v, ok := d.GetOk("chain"); ok {
		x := (v.(string))
		o.Chain = x

	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x

	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x

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

	url := "iam/TrustPoints"
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
	return resourceIamTrustPointRead(d, meta)
}

func resourceIamTrustPointRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "iam/TrustPoints" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.IamTrustPoint
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
		return err
	}

	if err := d.Set("certificates", flattenListX509Certificate(s.Certificates, d)); err != nil {
		return err
	}

	if err := d.Set("chain", (s.Chain)); err != nil {
		return err
	}

	if err := d.Set("moid", (s.Moid)); err != nil {
		return err
	}

	if err := d.Set("object_type", (s.ObjectType)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}

func resourceIamTrustPointDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "iam/TrustPoints" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
