package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTamSecurityAdvisory() *schema.Resource {
	return &schema.Resource{
		Create: resourceTamSecurityAdvisoryCreate,
		Read:   resourceTamSecurityAdvisoryRead,
		Update: resourceTamSecurityAdvisoryUpdate,
		Delete: resourceTamSecurityAdvisoryDelete,
		Schema: map[string]*schema.Schema{
			"api_data_sources": {
				Description: "An array of data sources that are used to provide data for queries used to identify an Intersight alert applicability.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"mo_type": {
							Description: "Type of Intersight managed object used as data source.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Name is used to unique identify and refer a given data source in an alert definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"queries": {
							Description: "Optional set of Queries to filter the output for Api datasource. the queries are executed in the order specified.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Description: "An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"query": {
										Description: "A SparkSQL query to be used on a given data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
						},
						"type": {
							Description: "Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "nxos",
						},
					},
				},
			},
			"actions": {
				Description: "An array of actions that are to be taken when a given managed object matches the criteria specified for being affected by an alert definition.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"affected_object_type": {
							Description: "Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"alert_type": {
							Description: "Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "psirt",
						},
						"identifiers": {
							Description: "Identifiers represents the filter criteria (property names and values) used to identify an Intersight managed object of type specified in affectedObjectType property. An instance of an alert is then create on (or removed from) the identified managed object.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"name": {
										Description: "Name of the filter paramter.",
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
										Description: "Value of the filter paramter.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"operation_type": {
							Description: "Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "create",
						},
						"queries": {
							Description: "Set of SparkSQL queries used determine if a given alert is applicable or not. Refer to https://spark.apache.org/sql/ for more details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Description: "An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"query": {
										Description: "A SparkSQL query to be used on a given data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
						},
						"type": {
							Description: "Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "restApi",
						},
					},
				},
			},
			"cve_id": {
				Description: "CVE (https://cve.mitre.org/about/faqs.html) identifier for the security Advisory.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Brief description of the advisory details.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"external_url": {
				Description: "A link to an external URL describing security Advisory in more details.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "A user defined name for the Intersight Advisory.",
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
			"organization": {
				Description: "Relationship to the Organization that owns the Managed Object.",
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
			"recommendation": {
				Description: "Recommended action to resolve the security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"score": {
				Description: "CVSS score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Description: "Severity level of the Intersight Advisory.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Description: "Current state of the advisory. Indicates if the user is interested in getting updates for the advisory.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "active",
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
func resourceTamSecurityAdvisoryCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.TamSecurityAdvisory
	if v, ok := d.GetOk("api_data_sources"); ok {
		x := make([]*models.TamAPIDataSource, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.TamAPIDataSource{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.TamBaseDataSourceAO0P0.TamBaseDataSourceAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["mo_type"]; ok {
					{
						x := (v.(string))
						o.MoType = x
					}
				}
				if v, ok := l["name"]; ok {
					{
						x := (v.(string))
						o.Name = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["queries"]; ok {
					{
						x := make([]*models.TamQueryEntry, 0)
						switch reflect.TypeOf(v).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(v)
							for i := 0; i < s.Len(); i++ {
								o := models.TamQueryEntry{}
								l := s.Index(i).Interface().(map[string]interface{})
								if v, ok := l["additional_properties"]; ok {
									{
										x := []byte(v.(string))
										var x1 interface{}
										err := json.Unmarshal(x, &x1)
										if err == nil && x1 != nil {
											o.TamQueryEntryAO0P0.TamQueryEntryAO0P0 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["name"]; ok {
									{
										x := (v.(string))
										o.Name = x
									}
								}
								if v, ok := l["object_type"]; ok {
									{
										x := (v.(string))
										o.ObjectType = x
									}
								}
								if v, ok := l["priority"]; ok {
									{
										x := int64(v.(int))
										o.Priority = x
									}
								}
								if v, ok := l["query"]; ok {
									{
										x := (v.(string))
										o.Query = x
									}
								}
								x = append(x, &o)
							}
						}
						o.Queries = x
					}
				}
				if v, ok := l["type"]; ok {
					{
						x := (v.(string))
						o.Type = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.APIDataSources = x

	}

	if v, ok := d.GetOk("actions"); ok {
		x := make([]*models.TamAction, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.TamAction{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.TamActionAO0P0.TamActionAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["affected_object_type"]; ok {
					{
						x := (v.(string))
						o.AffectedObjectType = x
					}
				}
				if v, ok := l["alert_type"]; ok {
					{
						x := (v.(string))
						o.AlertType = &x
					}
				}
				if v, ok := l["identifiers"]; ok {
					{
						x := make([]*models.TamIdentifiers, 0)
						switch reflect.TypeOf(v).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(v)
							for i := 0; i < s.Len(); i++ {
								o := models.TamIdentifiers{}
								l := s.Index(i).Interface().(map[string]interface{})
								if v, ok := l["additional_properties"]; ok {
									{
										x := []byte(v.(string))
										var x1 interface{}
										err := json.Unmarshal(x, &x1)
										if err == nil && x1 != nil {
											o.TamIdentifiersAO0P0.TamIdentifiersAO0P0 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["name"]; ok {
									{
										x := (v.(string))
										o.Name = x
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
						o.Identifiers = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["operation_type"]; ok {
					{
						x := (v.(string))
						o.OperationType = &x
					}
				}
				if v, ok := l["queries"]; ok {
					{
						x := make([]*models.TamQueryEntry, 0)
						switch reflect.TypeOf(v).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(v)
							for i := 0; i < s.Len(); i++ {
								o := models.TamQueryEntry{}
								l := s.Index(i).Interface().(map[string]interface{})
								if v, ok := l["additional_properties"]; ok {
									{
										x := []byte(v.(string))
										var x1 interface{}
										err := json.Unmarshal(x, &x1)
										if err == nil && x1 != nil {
											o.TamQueryEntryAO0P0.TamQueryEntryAO0P0 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["name"]; ok {
									{
										x := (v.(string))
										o.Name = x
									}
								}
								if v, ok := l["object_type"]; ok {
									{
										x := (v.(string))
										o.ObjectType = x
									}
								}
								if v, ok := l["priority"]; ok {
									{
										x := int64(v.(int))
										o.Priority = x
									}
								}
								if v, ok := l["query"]; ok {
									{
										x := (v.(string))
										o.Query = x
									}
								}
								x = append(x, &o)
							}
						}
						o.Queries = x
					}
				}
				if v, ok := l["type"]; ok {
					{
						x := (v.(string))
						o.Type = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.Actions = x

	}

	if v, ok := d.GetOk("cve_id"); ok {
		x := (v.(string))
		o.CveID = x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("external_url"); ok {
		x := (v.(string))
		o.ExternalURL = x

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

	if v, ok := d.GetOk("organization"); ok {
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
		o.Organization = &x

	}

	if v, ok := d.GetOk("recommendation"); ok {
		x := (v.(string))
		o.Recommendation = x

	}

	if v, ok := d.GetOk("score"); ok {
		x := v.(float32)
		o.Score = x

	}

	if v, ok := d.GetOk("severity"); ok {
		x := (v.(string))
		o.Severity = x

	}

	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.State = &x

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

	url := "tam/SecurityAdvisories"
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
	return resourceTamSecurityAdvisoryRead(d, meta)
}

func resourceTamSecurityAdvisoryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "tam/SecurityAdvisories" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.TamSecurityAdvisory
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("api_data_sources", flattenListTamAPIDataSource(s.APIDataSources, d)); err != nil {
		return err
	}

	if err := d.Set("actions", flattenListTamAction(s.Actions, d)); err != nil {
		return err
	}

	if err := d.Set("cve_id", (s.CveID)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("external_url", (s.ExternalURL)); err != nil {
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

	if err := d.Set("organization", flattenMapIamAccountRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("recommendation", (s.Recommendation)); err != nil {
		return err
	}

	if err := d.Set("score", (s.Score)); err != nil {
		return err
	}

	if err := d.Set("severity", (s.Severity)); err != nil {
		return err
	}

	if err := d.Set("state", (s.State)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceTamSecurityAdvisoryUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.TamSecurityAdvisory
	if d.HasChange("api_data_sources") {
		v := d.Get("api_data_sources")
		x := make([]*models.TamAPIDataSource, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.TamAPIDataSource{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.TamBaseDataSourceAO0P0.TamBaseDataSourceAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["mo_type"]; ok {
					{
						x := (v.(string))
						o.MoType = x
					}
				}
				if v, ok := l["name"]; ok {
					{
						x := (v.(string))
						o.Name = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["queries"]; ok {
					{
						x := make([]*models.TamQueryEntry, 0)
						switch reflect.TypeOf(v).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(v)
							for i := 0; i < s.Len(); i++ {
								o := models.TamQueryEntry{}
								l := s.Index(i).Interface().(map[string]interface{})
								if v, ok := l["additional_properties"]; ok {
									{
										x := []byte(v.(string))
										var x1 interface{}
										err := json.Unmarshal(x, &x1)
										if err == nil && x1 != nil {
											o.TamQueryEntryAO0P0.TamQueryEntryAO0P0 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["name"]; ok {
									{
										x := (v.(string))
										o.Name = x
									}
								}
								if v, ok := l["object_type"]; ok {
									{
										x := (v.(string))
										o.ObjectType = x
									}
								}
								if v, ok := l["priority"]; ok {
									{
										x := int64(v.(int))
										o.Priority = x
									}
								}
								if v, ok := l["query"]; ok {
									{
										x := (v.(string))
										o.Query = x
									}
								}
								x = append(x, &o)
							}
						}
						o.Queries = x
					}
				}
				if v, ok := l["type"]; ok {
					{
						x := (v.(string))
						o.Type = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.APIDataSources = x
	}

	if d.HasChange("actions") {
		v := d.Get("actions")
		x := make([]*models.TamAction, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.TamAction{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.TamActionAO0P0.TamActionAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["affected_object_type"]; ok {
					{
						x := (v.(string))
						o.AffectedObjectType = x
					}
				}
				if v, ok := l["alert_type"]; ok {
					{
						x := (v.(string))
						o.AlertType = &x
					}
				}
				if v, ok := l["identifiers"]; ok {
					{
						x := make([]*models.TamIdentifiers, 0)
						switch reflect.TypeOf(v).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(v)
							for i := 0; i < s.Len(); i++ {
								o := models.TamIdentifiers{}
								l := s.Index(i).Interface().(map[string]interface{})
								if v, ok := l["additional_properties"]; ok {
									{
										x := []byte(v.(string))
										var x1 interface{}
										err := json.Unmarshal(x, &x1)
										if err == nil && x1 != nil {
											o.TamIdentifiersAO0P0.TamIdentifiersAO0P0 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["name"]; ok {
									{
										x := (v.(string))
										o.Name = x
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
						o.Identifiers = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["operation_type"]; ok {
					{
						x := (v.(string))
						o.OperationType = &x
					}
				}
				if v, ok := l["queries"]; ok {
					{
						x := make([]*models.TamQueryEntry, 0)
						switch reflect.TypeOf(v).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(v)
							for i := 0; i < s.Len(); i++ {
								o := models.TamQueryEntry{}
								l := s.Index(i).Interface().(map[string]interface{})
								if v, ok := l["additional_properties"]; ok {
									{
										x := []byte(v.(string))
										var x1 interface{}
										err := json.Unmarshal(x, &x1)
										if err == nil && x1 != nil {
											o.TamQueryEntryAO0P0.TamQueryEntryAO0P0 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["name"]; ok {
									{
										x := (v.(string))
										o.Name = x
									}
								}
								if v, ok := l["object_type"]; ok {
									{
										x := (v.(string))
										o.ObjectType = x
									}
								}
								if v, ok := l["priority"]; ok {
									{
										x := int64(v.(int))
										o.Priority = x
									}
								}
								if v, ok := l["query"]; ok {
									{
										x := (v.(string))
										o.Query = x
									}
								}
								x = append(x, &o)
							}
						}
						o.Queries = x
					}
				}
				if v, ok := l["type"]; ok {
					{
						x := (v.(string))
						o.Type = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.Actions = x
	}

	if d.HasChange("cve_id") {
		v := d.Get("cve_id")
		x := (v.(string))
		o.CveID = x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
	}

	if d.HasChange("external_url") {
		v := d.Get("external_url")
		x := (v.(string))
		o.ExternalURL = x
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

	if d.HasChange("organization") {
		v := d.Get("organization")
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
		o.Organization = &x
	}

	if d.HasChange("recommendation") {
		v := d.Get("recommendation")
		x := (v.(string))
		o.Recommendation = x
	}

	if d.HasChange("score") {
		v := d.Get("score")
		x := v.(float32)
		o.Score = x
	}

	if d.HasChange("severity") {
		v := d.Get("severity")
		x := (v.(string))
		o.Severity = x
	}

	if d.HasChange("state") {
		v := d.Get("state")
		x := (v.(string))
		o.State = &x
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

	url := "tam/SecurityAdvisories" + "/" + d.Id()
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
	return resourceTamSecurityAdvisoryRead(d, meta)
}

func resourceTamSecurityAdvisoryDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "tam/SecurityAdvisories" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
