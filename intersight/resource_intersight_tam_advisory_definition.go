package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTamAdvisoryDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceTamAdvisoryDefinitionCreate,
		Read:   resourceTamAdvisoryDefinitionRead,
		Update: resourceTamAdvisoryDefinitionUpdate,
		Delete: resourceTamAdvisoryDefinitionDelete,
		Schema: map[string]*schema.Schema{
			"api_data_sources": {
				Description: "An array of data sources that are used to provide data for queries used to identify an Intersight alert applicability.",
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"type": {
							Description: "Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "nxos",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"actions": {
				Description: "An array of actions that are to be taken when a given managed object matches the criteria specified for being affected by an alert definition.",
				Type:        schema.TypeList,
				Optional:    true,
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the filter paramter.",
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
										Description: "Value of the filter paramter.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"name": {
							Description: "Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"type": {
							Description: "Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "restApi",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"advisory_details": {
				Description: "Additional details for the advisory definition. For e.g. if the definition corresponds to a security advisory, the details\nregarding CVE ids and CVSS score would be available here.",
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
						"description": {
							Description: "Brief description of details specified for an advisory type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"advisory_id": {
				Description: "Cisco generated identifier for the published security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"date_published": {
				Description: "Date when the security advisory was first published by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"date_updated": {
				Description: "Date when the security advisory was last updated by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Brief description of the advisory details.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"external_url": {
				Description: "A link to an external URL describing security Advisory in more details.",
				Type:        schema.TypeString,
				Optional:    true,
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
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
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
			"recommendation": {
				Description: "Recommended action to resolve the security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"severity": {
				Description: "Severity level of the Intersight Advisory.",
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
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"state": {
				Description: "Current state of the advisory.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ready",
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"type": {
				Description: "The type (field notice, security advisory etc.) of Intersight advisory.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "securityAdvisory",
			},
			"version": {
				Description: "Cisco assigned advisory version after latest revision.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"workaround": {
				Description: "Workarounds available for the advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func resourceTamAdvisoryDefinitionCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.TamAdvisoryDefinition
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
							o.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
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
											o.TamQueryEntryAO1P1.TamQueryEntryAO1P1 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["class_id"]; ok {
									{
										x := (v.(string))
										o.ClassID = x
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
							o.TamActionAO1P1.TamActionAO1P1 = x1.(map[string]interface{})
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
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
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
											o.TamIdentifiersAO1P1.TamIdentifiersAO1P1 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["class_id"]; ok {
									{
										x := (v.(string))
										o.ClassID = x
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
											o.TamQueryEntryAO1P1.TamQueryEntryAO1P1 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["class_id"]; ok {
									{
										x := (v.(string))
										o.ClassID = x
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

	if v, ok := d.GetOk("advisory_details"); ok {
		p := models.TamBaseAdvisoryDetails{}
		if len(v.([]interface{})) > 0 {
			o := models.TamBaseAdvisoryDetails{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.TamBaseAdvisoryDetailsAO1P1.TamBaseAdvisoryDetailsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.Description = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.AdvisoryDetails = &x
		}

	}

	if v, ok := d.GetOk("advisory_id"); ok {
		x := (v.(string))
		o.AdvisoryID = x

	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x

	}

	if v, ok := d.GetOk("date_published"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.DatePublished = x

	}

	if v, ok := d.GetOk("date_updated"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.DateUpdated = x

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
		p := models.OrganizationOrganizationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OrganizationOrganizationRef{}
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
			o.Organization = &x
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

	if v, ok := d.GetOk("recommendation"); ok {
		x := (v.(string))
		o.Recommendation = x

	}

	if v, ok := d.GetOk("severity"); ok {
		p := models.TamSeverity{}
		if len(v.([]interface{})) > 0 {
			o := models.TamSeverity{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AO1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.Severity = &x
		}

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

	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.Type = &x

	}

	if v, ok := d.GetOk("version"); ok {
		x := (v.(string))
		o.Version = x

	}

	if v, ok := d.GetOk("workaround"); ok {
		x := (v.(string))
		o.Workaround = x

	}

	url := "tam/AdvisoryDefinitions"
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
	return resourceTamAdvisoryDefinitionRead(d, meta)
}

func resourceTamAdvisoryDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "tam/AdvisoryDefinitions" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.TamAdvisoryDefinition
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

	if err := d.Set("advisory_details", flattenMapTamBaseAdvisoryDetails(s.AdvisoryDetails, d)); err != nil {
		return err
	}

	if err := d.Set("advisory_id", (s.AdvisoryID)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassID)); err != nil {
		return err
	}

	if err := d.Set("date_published", (s.DatePublished).String()); err != nil {
		return err
	}

	if err := d.Set("date_updated", (s.DateUpdated).String()); err != nil {
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

	if err := d.Set("organization", flattenMapOrganizationOrganizationRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("recommendation", (s.Recommendation)); err != nil {
		return err
	}

	if err := d.Set("severity", flattenMapTamSeverity(s.Severity, d)); err != nil {
		return err
	}

	if err := d.Set("state", (s.State)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("type", (s.Type)); err != nil {
		return err
	}

	if err := d.Set("version", (s.Version)); err != nil {
		return err
	}

	if err := d.Set("workaround", (s.Workaround)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceTamAdvisoryDefinitionUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.TamAdvisoryDefinition
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
							o.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
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
											o.TamQueryEntryAO1P1.TamQueryEntryAO1P1 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["class_id"]; ok {
									{
										x := (v.(string))
										o.ClassID = x
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
							o.TamActionAO1P1.TamActionAO1P1 = x1.(map[string]interface{})
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
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
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
											o.TamIdentifiersAO1P1.TamIdentifiersAO1P1 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["class_id"]; ok {
									{
										x := (v.(string))
										o.ClassID = x
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
											o.TamQueryEntryAO1P1.TamQueryEntryAO1P1 = x1.(map[string]interface{})
										}
									}
								}
								if v, ok := l["class_id"]; ok {
									{
										x := (v.(string))
										o.ClassID = x
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

	if d.HasChange("advisory_details") {
		v := d.Get("advisory_details")
		p := models.TamBaseAdvisoryDetails{}
		if len(v.([]interface{})) > 0 {
			o := models.TamBaseAdvisoryDetails{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.TamBaseAdvisoryDetailsAO1P1.TamBaseAdvisoryDetailsAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.Description = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.AdvisoryDetails = &x
		}
	}

	if d.HasChange("advisory_id") {
		v := d.Get("advisory_id")
		x := (v.(string))
		o.AdvisoryID = x
	}

	if d.HasChange("class_id") {
		v := d.Get("class_id")
		x := (v.(string))
		o.ClassID = x
	}

	if d.HasChange("date_published") {
		v := d.Get("date_published")
		x, _ := strfmt.ParseDateTime(v.(string))
		o.DatePublished = x
	}

	if d.HasChange("date_updated") {
		v := d.Get("date_updated")
		x, _ := strfmt.ParseDateTime(v.(string))
		o.DateUpdated = x
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
		p := models.OrganizationOrganizationRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OrganizationOrganizationRef{}
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
			o.Organization = &x
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

	if d.HasChange("recommendation") {
		v := d.Get("recommendation")
		x := (v.(string))
		o.Recommendation = x
	}

	if d.HasChange("severity") {
		v := d.Get("severity")
		p := models.TamSeverity{}
		if len(v.([]interface{})) > 0 {
			o := models.TamSeverity{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AO1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.Severity = &x
		}
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

	if d.HasChange("type") {
		v := d.Get("type")
		x := (v.(string))
		o.Type = &x
	}

	if d.HasChange("version") {
		v := d.Get("version")
		x := (v.(string))
		o.Version = x
	}

	if d.HasChange("workaround") {
		v := d.Get("workaround")
		x := (v.(string))
		o.Workaround = x
	}

	url := "tam/AdvisoryDefinitions" + "/" + d.Id()
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
	return resourceTamAdvisoryDefinitionRead(d, meta)
}

func resourceTamAdvisoryDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "tam/AdvisoryDefinitions" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
