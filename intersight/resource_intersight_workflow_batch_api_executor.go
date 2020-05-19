package intersight

import (
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWorkflowBatchApiExecutor() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowBatchApiExecutorCreate,
		Read:   resourceWorkflowBatchApiExecutorRead,
		Update: resourceWorkflowBatchApiExecutorUpdate,
		Delete: resourceWorkflowBatchApiExecutorDelete,
		Schema: map[string]*schema.Schema{
			"batch": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body": {
							Description: "The optional request body that is sent as part of this API request.\nThe request body can contain a golang template that can be populated with task input\nparameters and previous API output parameters.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"content_type": {
							Description: "Intersight Orchestrator, with the support of response parser specification,\ncan extract the values from API responses and map them to task output parameters.\nThe value extraction is supported for response content types XML and JSON.\nThe type of the content that gets passed as payload and response in this\nAPI.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "json",
						},
						"name": {
							Description: "A reference name for this API request within the batch API request.\nThis name shall be used to map the API output parameters to subsequent\nAPI input parameters within a batch API task.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"outcomes": {
							Description: "All the possible outcomes of this API are captured here. Outcomes property\nis a collection property of type workflow.Outcome objects.\nThe outcomes can be mapped to the message to be shown. The outcomes are\nevaluated in the order they are given. At the end of the outcomes list,\nan catchall success/fail outcome can be added with condition as 'true'.\nThis is an optional\nproperty and if not specified the task will be marked as success.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
						"response_spec": {
							Description: "The optional grammar specification for parsing the response to extract the\nrequired values.\nThe specification should have extraction specification specified for\nall the API output parameters.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"error_parameters": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accept_single_value": {
													Description: "The flag that allows single values in content to be extracted as a\nsingle element collection in case the parameter is of Collection type.\nThis flag is applicable for parameters of type Collection only.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"class_id": {
													Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"complex_type": {
													Description: "The name of the complex type definition in case this is a complex\nparameter. The content.Grammar object must have a complex type, content.ComplexType,\ndefined with the specified name in types collection property.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"item_type": {
													Description: "The type of the collection item in case this is a collection parameter.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "simple",
												},
												"name": {
													Description: "The name of the parameter.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"path": {
													Description: "The content specific path information that identifies the parameter\nvalue within the content. The value is usually a XPath or JSONPath or a\nregular expression in case of text content.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"type": {
													Description: "The type of the parameter. Accepted values are simple, complex,\ncollection.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "simple",
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"parameters": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accept_single_value": {
													Description: "The flag that allows single values in content to be extracted as a\nsingle element collection in case the parameter is of Collection type.\nThis flag is applicable for parameters of type Collection only.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"class_id": {
													Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"complex_type": {
													Description: "The name of the complex type definition in case this is a complex\nparameter. The content.Grammar object must have a complex type, content.ComplexType,\ndefined with the specified name in types collection property.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"item_type": {
													Description: "The type of the collection item in case this is a collection parameter.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "simple",
												},
												"name": {
													Description: "The name of the parameter.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"path": {
													Description: "The content specific path information that identifies the parameter\nvalue within the content. The value is usually a XPath or JSONPath or a\nregular expression in case of text content.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"type": {
													Description: "The type of the parameter. Accepted values are simple, complex,\ncollection.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "simple",
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
									"types": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_id": {
													Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"name": {
													Description: "The unique name of this complex type within the grammar specification.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"parameters": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"accept_single_value": {
																Description: "The flag that allows single values in content to be extracted as a\nsingle element collection in case the parameter is of Collection type.\nThis flag is applicable for parameters of type Collection only.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"class_id": {
																Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"complex_type": {
																Description: "The name of the complex type definition in case this is a complex\nparameter. The content.Grammar object must have a complex type, content.ComplexType,\ndefined with the specified name in types collection property.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"item_type": {
																Description: "The type of the collection item in case this is a collection parameter.",
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "simple",
															},
															"name": {
																Description: "The name of the parameter.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"object_type": {
																Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"path": {
																Description: "The content specific path information that identifies the parameter\nvalue within the content. The value is usually a XPath or JSONPath or a\nregular expression in case of text content.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"type": {
																Description: "The type of the parameter. Accepted values are simple, complex,\ncollection.",
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "simple",
															},
														},
													},
													ConfigMode: schema.SchemaConfigModeAttr,
													Computed:   true,
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"skip_on_condition": {
							Description: "The skip expression, if provided, allows the batch API executor to skip the\napi execution when the given expression evaluates to true.\nThe expression is given as such a golang template that has to be\nevaluated to a final content true/false. The expression is an optional and in\ncase not provided, the API will always be executed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"timeout": {
							Description: "The duration in seconds by which the API response is expected from the API target.\nIf the end point does not respond for the API request within this timeout\nduration, the task will be marked as failed.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"constraints": {
				Description: "Enter the constraints on when this task should match against the task definition.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						},
						"target_data_type": {
							Description: "List of property constraints that helps to narrow down task implementations based on target device input.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"description": {
				Description: "A detailed description about the batch APIs.",
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
				Description: "Name for the batch API task.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"outcomes": {
				Description: "All the possible outcomes of this task are captured here. Outcomes property\nis a collection property of type workflow.Outcome objects.\nThe outcomes can be mapped to the message to be shown. The outcomes are\nevaluated in the order they are given. At the end of the outcomes list,\nan catchall success/fail outcome can be added with condition as 'true'.\nThis is an optional\nproperty and if not specified the task will be marked as success.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"output": {
				Description: "Intersight Orchestrator allows the extraction of required values from API\nresponses using the API response grammar. These extracted values can be mapped\nto task output parameters defined in task definition.\nThe mapping of API output parameters to the task output parameters is provided\nas JSON in this property.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
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
			"skip_on_condition": {
				Description: "The skip expression, if provided, allows the batch API executor to skip the\ntask execution when the given expression evaluates to true.\nThe expression is given as such a golang template that has to be\nevaluated to a final content true/false. The expression is an optional and in\ncase not provided, the API will always be executed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"task_definition": {
				Description: "A reference to a workflowTaskDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
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
		},
	}
}

func resourceWorkflowBatchApiExecutorCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowBatchApiExecutor()
	if v, ok := d.GetOk("batch"); ok {
		x := make([]models.WorkflowApi, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowApiWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["body"]; ok {
				{
					x := (v.(string))
					o.SetBody(x)
				}
			}
			o.SetClassId("workflow.Api")
			if v, ok := l["content_type"]; ok {
				{
					x := (v.(string))
					o.SetContentType(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			o.SetObjectType("workflow.Api")
			if v, ok := l["outcomes"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetOutcomes(x)
				}
			}
			if v, ok := l["response_spec"]; ok {
				{
					p := make([]models.ContentGrammar, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewContentGrammarWithDefaults()
						o.SetClassId("content.Grammar")
						if v, ok := l["error_parameters"]; ok {
							{
								x := make([]models.ContentBaseParameter, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewContentBaseParameterWithDefaults()
									l := s[i].(map[string]interface{})
									if v, ok := l["accept_single_value"]; ok {
										{
											x := (v.(bool))
											o.SetAcceptSingleValue(x)
										}
									}
									o.SetClassId("content.BaseParameter")
									if v, ok := l["complex_type"]; ok {
										{
											x := (v.(string))
											o.SetComplexType(x)
										}
									}
									if v, ok := l["item_type"]; ok {
										{
											x := (v.(string))
											o.SetItemType(x)
										}
									}
									if v, ok := l["name"]; ok {
										{
											x := (v.(string))
											o.SetName(x)
										}
									}
									o.SetObjectType("content.BaseParameter")
									if v, ok := l["path"]; ok {
										{
											x := (v.(string))
											o.SetPath(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									x = append(x, *o)
								}
								o.SetErrorParameters(x)
							}
						}
						o.SetObjectType("content.Grammar")
						if v, ok := l["parameters"]; ok {
							{
								x := make([]models.ContentBaseParameter, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewContentBaseParameterWithDefaults()
									l := s[i].(map[string]interface{})
									if v, ok := l["accept_single_value"]; ok {
										{
											x := (v.(bool))
											o.SetAcceptSingleValue(x)
										}
									}
									o.SetClassId("content.BaseParameter")
									if v, ok := l["complex_type"]; ok {
										{
											x := (v.(string))
											o.SetComplexType(x)
										}
									}
									if v, ok := l["item_type"]; ok {
										{
											x := (v.(string))
											o.SetItemType(x)
										}
									}
									if v, ok := l["name"]; ok {
										{
											x := (v.(string))
											o.SetName(x)
										}
									}
									o.SetObjectType("content.BaseParameter")
									if v, ok := l["path"]; ok {
										{
											x := (v.(string))
											o.SetPath(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									x = append(x, *o)
								}
								o.SetParameters(x)
							}
						}
						if v, ok := l["types"]; ok {
							{
								x := make([]models.ContentComplexType, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewContentComplexTypeWithDefaults()
									l := s[i].(map[string]interface{})
									o.SetClassId("content.ComplexType")
									if v, ok := l["name"]; ok {
										{
											x := (v.(string))
											o.SetName(x)
										}
									}
									o.SetObjectType("content.ComplexType")
									if v, ok := l["parameters"]; ok {
										{
											x := make([]models.ContentBaseParameter, 0)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												o := models.NewContentBaseParameterWithDefaults()
												l := s[i].(map[string]interface{})
												if v, ok := l["accept_single_value"]; ok {
													{
														x := (v.(bool))
														o.SetAcceptSingleValue(x)
													}
												}
												o.SetClassId("content.BaseParameter")
												if v, ok := l["complex_type"]; ok {
													{
														x := (v.(string))
														o.SetComplexType(x)
													}
												}
												if v, ok := l["item_type"]; ok {
													{
														x := (v.(string))
														o.SetItemType(x)
													}
												}
												if v, ok := l["name"]; ok {
													{
														x := (v.(string))
														o.SetName(x)
													}
												}
												o.SetObjectType("content.BaseParameter")
												if v, ok := l["path"]; ok {
													{
														x := (v.(string))
														o.SetPath(x)
													}
												}
												if v, ok := l["type"]; ok {
													{
														x := (v.(string))
														o.SetType(x)
													}
												}
												x = append(x, *o)
											}
											o.SetParameters(x)
										}
									}
									x = append(x, *o)
								}
								o.SetTypes(x)
							}
						}
						p = append(p, *o)
					}
					x := p[0]
					o.SetResponseSpec(x)
				}
			}
			if v, ok := l["skip_on_condition"]; ok {
				{
					x := (v.(string))
					o.SetSkipOnCondition(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			x = append(x, *o)
		}
		o.SetBatch(x)
	}

	o.SetClassId("workflow.BatchApiExecutor")

	if v, ok := d.GetOk("constraints"); ok {
		p := make([]models.WorkflowTaskConstraints, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowTaskConstraintsWithDefaults()
			o.SetClassId("workflow.TaskConstraints")
			o.SetObjectType("workflow.TaskConstraints")
			if v, ok := l["target_data_type"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetTargetDataType(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetConstraints(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("workflow.BatchApiExecutor")

	if v, ok := d.GetOk("outcomes"); ok {
		x := v.(map[string]interface{})
		o.SetOutcomes(x)
	}

	if v, ok := d.GetOk("output"); ok {
		x := v.(map[string]interface{})
		o.SetOutput(x)
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("skip_on_condition"); ok {
		x := (v.(string))
		o.SetSkipOnCondition(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOk("task_definition"); ok {
		p := make([]models.WorkflowTaskDefinitionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowTaskDefinitionRelationship())
		}
		x := p[0]
		o.SetTaskDefinition(x)
	}

	r := conn.ApiClient.WorkflowApi.CreateWorkflowBatchApiExecutor(conn.ctx).WorkflowBatchApiExecutor(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Panicf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowBatchApiExecutorRead(d, meta)
}

func resourceWorkflowBatchApiExecutorRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.WorkflowApi.GetWorkflowBatchApiExecutorByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("batch", flattenListWorkflowApi(s.Batch, d)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassId)); err != nil {
		return err
	}

	if err := d.Set("constraints", flattenMapWorkflowTaskConstraints(s.Constraints, d)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
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

	if err := d.Set("outcomes", (s.Outcomes)); err != nil {
		return err
	}

	if err := d.Set("output", (s.Output)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("skip_on_condition", (s.SkipOnCondition)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("task_definition", flattenMapWorkflowTaskDefinitionRelationship(s.TaskDefinition, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceWorkflowBatchApiExecutorUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowBatchApiExecutor()
	if d.HasChange("batch") {
		v := d.Get("batch")
		x := make([]models.WorkflowApi, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowApiWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["body"]; ok {
				{
					x := (v.(string))
					o.SetBody(x)
				}
			}
			o.SetClassId("workflow.Api")
			if v, ok := l["content_type"]; ok {
				{
					x := (v.(string))
					o.SetContentType(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			o.SetObjectType("workflow.Api")
			if v, ok := l["outcomes"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetOutcomes(x)
				}
			}
			if v, ok := l["response_spec"]; ok {
				{
					p := make([]models.ContentGrammar, 0, 1)
					l := (v.([]interface{})[0]).(map[string]interface{})
					{
						o := models.NewContentGrammarWithDefaults()
						o.SetClassId("content.Grammar")
						if v, ok := l["error_parameters"]; ok {
							{
								x := make([]models.ContentBaseParameter, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewContentBaseParameterWithDefaults()
									l := s[i].(map[string]interface{})
									if v, ok := l["accept_single_value"]; ok {
										{
											x := (v.(bool))
											o.SetAcceptSingleValue(x)
										}
									}
									o.SetClassId("content.BaseParameter")
									if v, ok := l["complex_type"]; ok {
										{
											x := (v.(string))
											o.SetComplexType(x)
										}
									}
									if v, ok := l["item_type"]; ok {
										{
											x := (v.(string))
											o.SetItemType(x)
										}
									}
									if v, ok := l["name"]; ok {
										{
											x := (v.(string))
											o.SetName(x)
										}
									}
									o.SetObjectType("content.BaseParameter")
									if v, ok := l["path"]; ok {
										{
											x := (v.(string))
											o.SetPath(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									x = append(x, *o)
								}
								o.SetErrorParameters(x)
							}
						}
						o.SetObjectType("content.Grammar")
						if v, ok := l["parameters"]; ok {
							{
								x := make([]models.ContentBaseParameter, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewContentBaseParameterWithDefaults()
									l := s[i].(map[string]interface{})
									if v, ok := l["accept_single_value"]; ok {
										{
											x := (v.(bool))
											o.SetAcceptSingleValue(x)
										}
									}
									o.SetClassId("content.BaseParameter")
									if v, ok := l["complex_type"]; ok {
										{
											x := (v.(string))
											o.SetComplexType(x)
										}
									}
									if v, ok := l["item_type"]; ok {
										{
											x := (v.(string))
											o.SetItemType(x)
										}
									}
									if v, ok := l["name"]; ok {
										{
											x := (v.(string))
											o.SetName(x)
										}
									}
									o.SetObjectType("content.BaseParameter")
									if v, ok := l["path"]; ok {
										{
											x := (v.(string))
											o.SetPath(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									x = append(x, *o)
								}
								o.SetParameters(x)
							}
						}
						if v, ok := l["types"]; ok {
							{
								x := make([]models.ContentComplexType, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewContentComplexTypeWithDefaults()
									l := s[i].(map[string]interface{})
									o.SetClassId("content.ComplexType")
									if v, ok := l["name"]; ok {
										{
											x := (v.(string))
											o.SetName(x)
										}
									}
									o.SetObjectType("content.ComplexType")
									if v, ok := l["parameters"]; ok {
										{
											x := make([]models.ContentBaseParameter, 0)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												o := models.NewContentBaseParameterWithDefaults()
												l := s[i].(map[string]interface{})
												if v, ok := l["accept_single_value"]; ok {
													{
														x := (v.(bool))
														o.SetAcceptSingleValue(x)
													}
												}
												o.SetClassId("content.BaseParameter")
												if v, ok := l["complex_type"]; ok {
													{
														x := (v.(string))
														o.SetComplexType(x)
													}
												}
												if v, ok := l["item_type"]; ok {
													{
														x := (v.(string))
														o.SetItemType(x)
													}
												}
												if v, ok := l["name"]; ok {
													{
														x := (v.(string))
														o.SetName(x)
													}
												}
												o.SetObjectType("content.BaseParameter")
												if v, ok := l["path"]; ok {
													{
														x := (v.(string))
														o.SetPath(x)
													}
												}
												if v, ok := l["type"]; ok {
													{
														x := (v.(string))
														o.SetType(x)
													}
												}
												x = append(x, *o)
											}
											o.SetParameters(x)
										}
									}
									x = append(x, *o)
								}
								o.SetTypes(x)
							}
						}
						p = append(p, *o)
					}
					x := p[0]
					o.SetResponseSpec(x)
				}
			}
			if v, ok := l["skip_on_condition"]; ok {
				{
					x := (v.(string))
					o.SetSkipOnCondition(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			x = append(x, *o)
		}
		o.SetBatch(x)
	}

	if d.HasChange("constraints") {
		v := d.Get("constraints")
		p := make([]models.WorkflowTaskConstraints, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewWorkflowTaskConstraintsWithDefaults()
			o.SetClassId("workflow.TaskConstraints")
			o.SetObjectType("workflow.TaskConstraints")
			if v, ok := l["target_data_type"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetTargetDataType(x)
				}
			}
			p = append(p, *o)
		}
		x := p[0]
		o.SetConstraints(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if d.HasChange("outcomes") {
		v := d.Get("outcomes")
		x := v.(map[string]interface{})
		o.SetOutcomes(x)
	}

	if d.HasChange("output") {
		v := d.Get("output")
		x := v.(map[string]interface{})
		o.SetOutput(x)
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("mo.BaseMo")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, o.AsMoBaseMoRelationship())
		}
		o.SetPermissionResources(x)
	}

	if d.HasChange("skip_on_condition") {
		v := d.Get("skip_on_condition")
		x := (v.(string))
		o.SetSkipOnCondition(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if d.HasChange("task_definition") {
		v := d.Get("task_definition")
		p := make([]models.WorkflowTaskDefinitionRelationship, 0, 1)
		l := (v.([]interface{})[0]).(map[string]interface{})
		{
			o := models.NewMoMoRefWithDefaults()
			o.SetClassId("mo.MoRef")
			if v, ok := l["link"]; ok {
				{
					x := (v.(string))
					o.SetLink(x)
				}
			}
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			o.SetObjectType("workflow.TaskDefinition")
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, o.AsWorkflowTaskDefinitionRelationship())
		}
		x := p[0]
		o.SetTaskDefinition(x)
	}

	r := conn.ApiClient.WorkflowApi.UpdateWorkflowBatchApiExecutor(conn.ctx, d.Id()).WorkflowBatchApiExecutor(*o)
	result, _, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowBatchApiExecutorRead(d, meta)
}

func resourceWorkflowBatchApiExecutorDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.WorkflowApi.DeleteWorkflowBatchApiExecutor(conn.ctx, d.Id())
	_, err := r.Execute()
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
