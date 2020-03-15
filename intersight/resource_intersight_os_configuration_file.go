package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOsConfigurationFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceOsConfigurationFileCreate,
		Read:   resourceOsConfigurationFileRead,
		Delete: resourceOsConfigurationFileDelete,
		Schema: map[string]*schema.Schema{
			"catalog": {
				Description: "A collection of references to the [os.Catalog](mo://os.Catalog) Managed Object.When this managed object is deleted, the referenced [os.Catalog](mo://os.Catalog) MO unsets its reference to this deleted MO.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"distributions": {
				Description: "This captures the operating system for which this configuration file isdefined.",
				Type:        schema.TypeList,
				Optional:    true,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"file_content": {
				Description: "The content of the entire configuration file is stored as value. The contentcan either be a static file content or a template content.The template is expected to conform to the golang template syntax. The valuesfrom os.Answers properties will be used to populate this template.",
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
				Description: "The name of the OS ConfigurationFile that uniquely identifies the configuration file.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
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
			"placeholders": {
				Description: "This readonly property holds the list of placeholder names used in theconfiguration file content in case it is a template.",
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
						"is_value_set": {
							Description: "Flag to indicate if value is set. Value will be used to check if any edit.",
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"type": {
							Description: "Definition of place holder.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
										ForceNew:         true,
									},
									"default": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"additional_properties": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: SuppressDiffAdditionProps,
													ForceNew:         true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													ForceNew:    true,
												},
												"override": {
													Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
													Type:        schema.TypeBool,
													Optional:    true,
													ForceNew:    true,
												},
												"value": {
													Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
													Type:        schema.TypeMap,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													}, Optional: true,
													ForceNew: true,
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										ForceNew:   true,
									},
									"description": {
										Description: "Provide a detailed description of the data type.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"label": {
										Description: "Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"name": {
										Description: "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"properties": {
										Description: "Primitive data type properties.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"additional_properties": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: SuppressDiffAdditionProps,
													ForceNew:         true,
												},
												"constraints": {
													Description: "Constraints that must be applied to the parameter value supplied for this data type.",
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"additional_properties": {
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: SuppressDiffAdditionProps,
																ForceNew:         true,
															},
															"enum_list": {
																Description: "When the parameter is a enum then this list of enum entry is used to validate the input belongs to one of items in the list.",
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
																		"label": {
																			Description: "Label for the enum value. A user friendly short string to identify the enum value.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																		},
																		"object_type": {
																			Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			ForceNew:    true,
																		},
																		"value": {
																			Description: "Enum value for this enum entry. Value will be passed to the workflow as string type for execution.",
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
															"max": {
																Description: "Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked.",
																Type:        schema.TypeFloat,
																Optional:    true,
																ForceNew:    true,
															},
															"min": {
																Description: "Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked.",
																Type:        schema.TypeFloat,
																Optional:    true,
																ForceNew:    true,
															},
															"object_type": {
																Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																ForceNew:    true,
															},
															"regex": {
																Description: "When the parameter is a string this regular expression is used to ensure the value is valid.",
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
												"inventory_selector": {
													Description: "List of Intersight managed object selectors. The workflow execution user interface show objects from inventory that are matching the selectors to help with selecting inputs.",
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
															"display_attributes": {
																Description: "List of properties from an Intersight object which can help to identify the object. Typically the set of identity constraints on the object can be listed here to help the user identity the managed object.",
																Type:        schema.TypeList,
																Optional:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString}, ForceNew: true,
															},
															"object_type": {
																Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																ForceNew:    true,
															},
															"selector": {
																Description: "Field to hold an Intersight API along with an optional filter to narrow down the search options.",
																Type:        schema.TypeString,
																Optional:    true,
																ForceNew:    true,
															},
															"value_attribute": {
																Description: "A property from the Intersight object, value of which can be used as value for referenced input definition.",
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
												"object_type": {
													Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													ForceNew:    true,
												},
												"secure": {
													Description: "Intersight supports secure properties as task input/output. The values ofthese properties are encrypted and stored in Intersight.This flag marks the property to be secure when it is set to true.",
													Type:        schema.TypeBool,
													Optional:    true,
													ForceNew:    true,
												},
												"type": {
													Description: "Specify the enum type for primitive data type.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "string",
													ForceNew:    true,
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										ForceNew:   true,
									},
									"required": {
										Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
										Type:        schema.TypeBool,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							ForceNew:   true,
						},
						"value": {
							Description: "Value for placeholder provided by user.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
							ForceNew: true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				ForceNew:   true,
			},
			"supported": {
				Description: "An internal property that is used to distinguish between the pre-canned OSconfiguration file entries and user provided entries.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
func resourceOsConfigurationFileCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.OsConfigurationFile
	if v, ok := d.GetOk("catalog"); ok {
		p := models.OsCatalogRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OsCatalogRef{}
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
		o.Catalog = &x

	}

	if v, ok := d.GetOk("distributions"); ok {
		x := make([]*models.HclOperatingSystemRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.HclOperatingSystemRef{}
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
		o.Distributions = x

	}

	if v, ok := d.GetOk("file_content"); ok {
		x := (v.(string))
		o.FileContent = x

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

	if v, ok := d.GetOk("placeholders"); ok {
		x := make([]*models.OsPlaceHolder, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.OsPlaceHolder{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["is_value_set"]; ok {
					{
						x := (v.(bool))
						o.IsValueSet = &x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["type"]; ok {
					{
						p := models.WorkflowPrimitiveDataType{}
						if len(v.([]interface{})) > 0 {
							o := models.WorkflowPrimitiveDataType{}
							l := (v.([]interface{})[0]).(map[string]interface{})
							if v, ok := l["additional_properties"]; ok {
								{
									x := []byte(v.(string))
									var x1 interface{}
									err := json.Unmarshal(x, &x1)
									if err == nil && x1 != nil {
										o.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1 = x1.(map[string]interface{})
									}
								}
							}
							if v, ok := l["default"]; ok {
								{
									p := models.WorkflowDefaultValue{}
									if len(v.([]interface{})) > 0 {
										o := models.WorkflowDefaultValue{}
										l := (v.([]interface{})[0]).(map[string]interface{})
										if v, ok := l["additional_properties"]; ok {
											{
												x := []byte(v.(string))
												var x1 interface{}
												err := json.Unmarshal(x, &x1)
												if err == nil && x1 != nil {
													o.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1 = x1.(map[string]interface{})
												}
											}
										}
										if v, ok := l["object_type"]; ok {
											{
												x := (v.(string))
												o.ObjectType = x
											}
										}
										if v, ok := l["override"]; ok {
											{
												x := (v.(bool))
												o.Override = &x
											}
										}
										if v, ok := l["value"]; ok {
											{
												x := v
												o.Value = &x
											}
										}

										p = o
									}
									x := p
									o.Default = &x
								}
							}
							if v, ok := l["description"]; ok {
								{
									x := (v.(string))
									o.Description = x
								}
							}
							if v, ok := l["label"]; ok {
								{
									x := (v.(string))
									o.Label = x
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
							if v, ok := l["properties"]; ok {
								{
									p := models.WorkflowPrimitiveDataProperty{}
									if len(v.([]interface{})) > 0 {
										o := models.WorkflowPrimitiveDataProperty{}
										l := (v.([]interface{})[0]).(map[string]interface{})
										if v, ok := l["additional_properties"]; ok {
											{
												x := []byte(v.(string))
												var x1 interface{}
												err := json.Unmarshal(x, &x1)
												if err == nil && x1 != nil {
													o.WorkflowPrimitiveDataPropertyAO1P1.WorkflowPrimitiveDataPropertyAO1P1 = x1.(map[string]interface{})
												}
											}
										}
										if v, ok := l["constraints"]; ok {
											{
												p := models.WorkflowConstraints{}
												if len(v.([]interface{})) > 0 {
													o := models.WorkflowConstraints{}
													l := (v.([]interface{})[0]).(map[string]interface{})
													if v, ok := l["additional_properties"]; ok {
														{
															x := []byte(v.(string))
															var x1 interface{}
															err := json.Unmarshal(x, &x1)
															if err == nil && x1 != nil {
																o.WorkflowConstraintsAO1P1.WorkflowConstraintsAO1P1 = x1.(map[string]interface{})
															}
														}
													}
													if v, ok := l["enum_list"]; ok {
														{
															x := make([]*models.WorkflowEnumEntry, 0)
															switch reflect.TypeOf(v).Kind() {
															case reflect.Slice:
																s := reflect.ValueOf(v)
																for i := 0; i < s.Len(); i++ {
																	o := models.WorkflowEnumEntry{}
																	l := s.Index(i).Interface().(map[string]interface{})
																	if v, ok := l["additional_properties"]; ok {
																		{
																			x := []byte(v.(string))
																			var x1 interface{}
																			err := json.Unmarshal(x, &x1)
																			if err == nil && x1 != nil {
																				o.WorkflowEnumEntryAO1P1.WorkflowEnumEntryAO1P1 = x1.(map[string]interface{})
																			}
																		}
																	}
																	if v, ok := l["label"]; ok {
																		{
																			x := (v.(string))
																			o.Label = x
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
															o.EnumList = x
														}
													}
													if v, ok := l["max"]; ok {
														{
															x := v.(float64)
															o.Max = x
														}
													}
													if v, ok := l["min"]; ok {
														{
															x := v.(float64)
															o.Min = x
														}
													}
													if v, ok := l["object_type"]; ok {
														{
															x := (v.(string))
															o.ObjectType = x
														}
													}
													if v, ok := l["regex"]; ok {
														{
															x := (v.(string))
															o.Regex = x
														}
													}

													p = o
												}
												x := p
												o.Constraints = &x
											}
										}
										if v, ok := l["inventory_selector"]; ok {
											{
												x := make([]*models.WorkflowMoReferenceProperty, 0)
												switch reflect.TypeOf(v).Kind() {
												case reflect.Slice:
													s := reflect.ValueOf(v)
													for i := 0; i < s.Len(); i++ {
														o := models.WorkflowMoReferenceProperty{}
														l := s.Index(i).Interface().(map[string]interface{})
														if v, ok := l["additional_properties"]; ok {
															{
																x := []byte(v.(string))
																var x1 interface{}
																err := json.Unmarshal(x, &x1)
																if err == nil && x1 != nil {
																	o.WorkflowMoReferencePropertyAO1P1.WorkflowMoReferencePropertyAO1P1 = x1.(map[string]interface{})
																}
															}
														}
														if v, ok := l["display_attributes"]; ok {
															{
																x := make([]string, 0)
																y := reflect.ValueOf(v)
																for i := 0; i < y.Len(); i++ {
																	x = append(x, y.Index(i).Interface().(string))
																}
																o.DisplayAttributes = x
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
														if v, ok := l["value_attribute"]; ok {
															{
																x := (v.(string))
																o.ValueAttribute = x
															}
														}
														x = append(x, &o)
													}
												}
												o.InventorySelector = x
											}
										}
										if v, ok := l["object_type"]; ok {
											{
												x := (v.(string))
												o.ObjectType = x
											}
										}
										if v, ok := l["secure"]; ok {
											{
												x := (v.(bool))
												o.Secure = &x
											}
										}
										if v, ok := l["type"]; ok {
											{
												x := (v.(string))
												o.Type = &x
											}
										}

										p = o
									}
									x := p
									o.Properties = &x
								}
							}
							if v, ok := l["required"]; ok {
								{
									x := (v.(bool))
									o.Required = &x
								}
							}

							p = o
						}
						x := p
						o.Type = &x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := v
						o.Value = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.Placeholders = x

	}

	if v, ok := d.GetOkExists("supported"); ok {
		x := v.(bool)
		o.Supported = &x
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

	url := "os/ConfigurationFiles"
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
	return resourceOsConfigurationFileRead(d, meta)
}

func resourceOsConfigurationFileRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "os/ConfigurationFiles" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.OsConfigurationFile
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("catalog", flattenMapOsCatalogRef(s.Catalog, d)); err != nil {
		return err
	}

	if err := d.Set("distributions", flattenListHclOperatingSystemRef(s.Distributions, d)); err != nil {
		return err
	}

	if err := d.Set("file_content", (s.FileContent)); err != nil {
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

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("placeholders", flattenListOsPlaceHolder(s.Placeholders, d)); err != nil {
		return err
	}

	if err := d.Set("supported", (s.Supported)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}

func resourceOsConfigurationFileDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "os/ConfigurationFiles" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
