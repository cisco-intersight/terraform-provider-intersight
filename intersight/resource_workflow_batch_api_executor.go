package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func resourceWorkflowBatchApiExecutor() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowBatchApiExecutorCreate,
		Read:   resourceWorkflowBatchApiExecutorRead,
		Update: resourceWorkflowBatchApiExecutorUpdate,
		Delete: resourceWorkflowBatchApiExecutorDelete,
		Schema: map[string]*schema.Schema{
			"batch": {
				Description: "Intersight Orchestrator supports one or a batch of APIs to be executed as part ofa task execution.The batch cannot be empty.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"body": {
							Description: "The optional request body that is sent as part of this API request.The request body can contain a golang template that can be populated with task inputparameters and previous API output parameters.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"content_type": {
							Description: "Intersight Orchestrator, with the support of response parser specification,can extract the values from API responses and map them to task output parameters.The value extraction is supported for response content types XML and JSON.The type of the content that gets passed as payload and response in thisAPI.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "json",
						},
						"expected_state": {
							Description: "The expected state expression, if provided, allows the Intersight Orchestratorto mark a task status as success/fail by checking the extracted APIvalues against a set of expected values.The expected state is given as a golang template that has to evaluate to afinal content true/false.This property is deprecated in favour of Outcomes property.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "A reference name for this API request within the batch API request.This name shall be used to map the API output parameters to subsequentAPI input parameters within a batch API task.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"outcomes": {
							Description: "All the possible outcomes of this API are captured here. Outcomes propertyis a collection property of type workflow.Outcome objects.The outcomes can be mapped to the message to be shown. The outcomes areevaluated in the order they are given. At the end of the outcomes list,an catchall success/fail outcome can be added with condition as 'true'.This is an optionalproperty and if not specified the task will be marked as success.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
						"response_spec": {
							Description: "The optional grammar specification for parsing the response to extract therequired values.The specification should have extraction specification specified forall the API output parameters.",
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
									"error_parameters": {
										Description: "The list of parameter definitions, if found in a given API/device response,makes the content handlers to treat the response as error response.This is optional parameter.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accept_single_value": {
													Description: "The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"additional_properties": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: SuppressDiffAdditionProps,
												},
												"complex_type": {
													Description: "The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property.",
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
													Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"path": {
													Description: "The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"type": {
													Description: "The type of the parameter. Accepted values are simple, complex,collection.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "simple",
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
									"parameters": {
										Description: "The list of parameter definitions that mark the parameters to beextracted using this grammar specification.",
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accept_single_value": {
													Description: "The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"additional_properties": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: SuppressDiffAdditionProps,
												},
												"complex_type": {
													Description: "The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property.",
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
													Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"path": {
													Description: "The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"type": {
													Description: "The type of the parameter. Accepted values are simple, complex,collection.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "simple",
												},
											},
										},
										Computed:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
									},
									"types": {
										Description: "The collection of complex types definitions used in this grammarspecification.This is required only if any of the parameters provided in this grammaris of complex type.",
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
													Description: "The unique name of this complex type within the grammar specification.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"parameters": {
													Description: "The collection of parameters that are part of this complex type.",
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"accept_single_value": {
																Description: "The flag that allows single values in content to be extracted as asingle element collection in case the parameter is of Collection type.This flag is applicable for parameters of type Collection only.",
																Type:        schema.TypeBool,
																Optional:    true,
															},
															"additional_properties": {
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: SuppressDiffAdditionProps,
															},
															"complex_type": {
																Description: "The name of the complex type definition in case this is a complexparameter. The content.Grammar object must have a complex type, content.ComplexType,defined with the specified name in types collection property.",
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
																Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"path": {
																Description: "The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"type": {
																Description: "The type of the parameter. Accepted values are simple, complex,collection.",
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "simple",
															},
														},
													},
													Computed:   true,
													ConfigMode: schema.SchemaConfigModeAttr,
												},
											},
										},
										Computed:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
									},
								},
							},
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
						},
						"timeout": {
							Description: "The duration in seconds by which the API response is expected from the API target.If the end point does not respond for the API request within this timeoutduration, the task will be marked as failed.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"constraints": {
				Description: "Enter the constraints on when this task should match against the task definition.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"output": {
				Description: "Intersight Orchestrator allows the extraction of required values from APIresponses using the API response grammar. These extracted values can be mappedto task output parameters defined in task definition.The mapping of API output parameters to the task output parameters is providedas JSON in this property.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"skip_on_condition": {
				Description: "The skip expression, if provided, allows the batch API executor to skip thetask execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"task_definition": {
				Description: "The interface task definition for which this batch API is one of the implementation.",
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
				ForceNew:   true,
			},
		},
	}
}
func resourceWorkflowBatchApiExecutorCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.WorkflowBatchAPIExecutor
	if v, ok := d.GetOk("batch"); ok {
		x := make([]*models.WorkflowAPI, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowAPI{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.WorkflowAPIAO0P0.WorkflowAPIAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["body"]; ok {
					{
						x := (v.(string))
						o.Body = x
					}
				}
				if v, ok := l["content_type"]; ok {
					{
						x := (v.(string))
						o.ContentType = &x
					}
				}
				if v, ok := l["expected_state"]; ok {
					{
						x := (v.(string))
						o.ExpectedState = x
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
				if v, ok := l["outcomes"]; ok {
					{
						x := v
						o.Outcomes = &x
					}
				}
				if v, ok := l["response_spec"]; ok {
					{
						p := models.ContentGrammar{}
						if len(v.([]interface{})) > 0 {
							o := models.ContentGrammar{}
							l := (v.([]interface{})[0]).(map[string]interface{})
							if v, ok := l["additional_properties"]; ok {
								{
									x := []byte(v.(string))
									var x1 interface{}
									err := json.Unmarshal(x, &x1)
									if err == nil && x1 != nil {
										o.ContentGrammarAO0P0.ContentGrammarAO0P0 = x1.(map[string]interface{})
									}
								}
							}
							if v, ok := l["error_parameters"]; ok {
								{
									x := make([]*models.ContentParameter, 0)
									switch reflect.TypeOf(v).Kind() {
									case reflect.Slice:
										s := reflect.ValueOf(v)
										for i := 0; i < s.Len(); i++ {
											o := models.ContentParameter{}
											l := s.Index(i).Interface().(map[string]interface{})
											if v, ok := l["accept_single_value"]; ok {
												{
													x := (v.(bool))
													o.AcceptSingleValue = &x
												}
											}
											if v, ok := l["additional_properties"]; ok {
												{
													x := []byte(v.(string))
													var x1 interface{}
													err := json.Unmarshal(x, &x1)
													if err == nil && x1 != nil {
														o.ContentParameterAO0P0.ContentParameterAO0P0 = x1.(map[string]interface{})
													}
												}
											}
											if v, ok := l["complex_type"]; ok {
												{
													x := (v.(string))
													o.ComplexType = x
												}
											}
											if v, ok := l["item_type"]; ok {
												{
													x := (v.(string))
													o.ItemType = &x
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
											if v, ok := l["path"]; ok {
												{
													x := (v.(string))
													o.Path = x
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
									o.ErrorParameters = x
								}
							}
							if v, ok := l["object_type"]; ok {
								{
									x := (v.(string))
									o.ObjectType = x
								}
							}
							if v, ok := l["parameters"]; ok {
								{
									x := make([]*models.ContentParameter, 0)
									switch reflect.TypeOf(v).Kind() {
									case reflect.Slice:
										s := reflect.ValueOf(v)
										for i := 0; i < s.Len(); i++ {
											o := models.ContentParameter{}
											l := s.Index(i).Interface().(map[string]interface{})
											if v, ok := l["accept_single_value"]; ok {
												{
													x := (v.(bool))
													o.AcceptSingleValue = &x
												}
											}
											if v, ok := l["additional_properties"]; ok {
												{
													x := []byte(v.(string))
													var x1 interface{}
													err := json.Unmarshal(x, &x1)
													if err == nil && x1 != nil {
														o.ContentParameterAO0P0.ContentParameterAO0P0 = x1.(map[string]interface{})
													}
												}
											}
											if v, ok := l["complex_type"]; ok {
												{
													x := (v.(string))
													o.ComplexType = x
												}
											}
											if v, ok := l["item_type"]; ok {
												{
													x := (v.(string))
													o.ItemType = &x
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
											if v, ok := l["path"]; ok {
												{
													x := (v.(string))
													o.Path = x
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
									o.Parameters = x
								}
							}
							if v, ok := l["types"]; ok {
								{
									x := make([]*models.ContentComplexType, 0)
									switch reflect.TypeOf(v).Kind() {
									case reflect.Slice:
										s := reflect.ValueOf(v)
										for i := 0; i < s.Len(); i++ {
											o := models.ContentComplexType{}
											l := s.Index(i).Interface().(map[string]interface{})
											if v, ok := l["additional_properties"]; ok {
												{
													x := []byte(v.(string))
													var x1 interface{}
													err := json.Unmarshal(x, &x1)
													if err == nil && x1 != nil {
														o.ContentComplexTypeAO0P0.ContentComplexTypeAO0P0 = x1.(map[string]interface{})
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
											if v, ok := l["parameters"]; ok {
												{
													x := make([]*models.ContentParameter, 0)
													switch reflect.TypeOf(v).Kind() {
													case reflect.Slice:
														s := reflect.ValueOf(v)
														for i := 0; i < s.Len(); i++ {
															o := models.ContentParameter{}
															l := s.Index(i).Interface().(map[string]interface{})
															if v, ok := l["accept_single_value"]; ok {
																{
																	x := (v.(bool))
																	o.AcceptSingleValue = &x
																}
															}
															if v, ok := l["additional_properties"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		o.ContentParameterAO0P0.ContentParameterAO0P0 = x1.(map[string]interface{})
																	}
																}
															}
															if v, ok := l["complex_type"]; ok {
																{
																	x := (v.(string))
																	o.ComplexType = x
																}
															}
															if v, ok := l["item_type"]; ok {
																{
																	x := (v.(string))
																	o.ItemType = &x
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
															if v, ok := l["path"]; ok {
																{
																	x := (v.(string))
																	o.Path = x
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
													o.Parameters = x
												}
											}
											x = append(x, &o)
										}
									}
									o.Types = x
								}
							}

							p = o
						}
						x := p
						o.ResponseSpec = &x
					}
				}
				if v, ok := l["timeout"]; ok {
					{
						x := int64(v.(int))
						o.Timeout = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Batch = x

	}

	if v, ok := d.GetOk("constraints"); ok {
		x := (v.(string))
		o.Constraints = x

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

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

	if v, ok := d.GetOk("output"); ok {
		x := v
		o.Output = &x

	}

	if v, ok := d.GetOk("skip_on_condition"); ok {
		x := (v.(string))
		o.SkipOnCondition = x

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

	if v, ok := d.GetOk("task_definition"); ok {
		p := models.WorkflowTaskDefinitionRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowTaskDefinitionRef{}
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
		o.TaskDefinition = &x

	}

	url := "workflow/BatchApiExecutors"
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
	return resourceWorkflowBatchApiExecutorRead(d, meta)
}

func resourceWorkflowBatchApiExecutorRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/BatchApiExecutors" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.WorkflowBatchAPIExecutor
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("batch", flattenListWorkflowAPI(s.Batch, d)); err != nil {
		return err
	}

	if err := d.Set("constraints", (s.Constraints)); err != nil {
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

	if err := d.Set("output", (s.Output)); err != nil {
		return err
	}

	if err := d.Set("skip_on_condition", (s.SkipOnCondition)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("task_definition", flattenMapWorkflowTaskDefinitionRef(s.TaskDefinition, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceWorkflowBatchApiExecutorUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.WorkflowBatchAPIExecutor
	if d.HasChange("batch") {
		v := d.Get("batch")
		x := make([]*models.WorkflowAPI, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowAPI{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.WorkflowAPIAO0P0.WorkflowAPIAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["body"]; ok {
					{
						x := (v.(string))
						o.Body = x
					}
				}
				if v, ok := l["content_type"]; ok {
					{
						x := (v.(string))
						o.ContentType = &x
					}
				}
				if v, ok := l["expected_state"]; ok {
					{
						x := (v.(string))
						o.ExpectedState = x
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
				if v, ok := l["outcomes"]; ok {
					{
						x := v
						o.Outcomes = &x
					}
				}
				if v, ok := l["response_spec"]; ok {
					{
						p := models.ContentGrammar{}
						if len(v.([]interface{})) > 0 {
							o := models.ContentGrammar{}
							l := (v.([]interface{})[0]).(map[string]interface{})
							if v, ok := l["additional_properties"]; ok {
								{
									x := []byte(v.(string))
									var x1 interface{}
									err := json.Unmarshal(x, &x1)
									if err == nil && x1 != nil {
										o.ContentGrammarAO0P0.ContentGrammarAO0P0 = x1.(map[string]interface{})
									}
								}
							}
							if v, ok := l["error_parameters"]; ok {
								{
									x := make([]*models.ContentParameter, 0)
									switch reflect.TypeOf(v).Kind() {
									case reflect.Slice:
										s := reflect.ValueOf(v)
										for i := 0; i < s.Len(); i++ {
											o := models.ContentParameter{}
											l := s.Index(i).Interface().(map[string]interface{})
											if v, ok := l["accept_single_value"]; ok {
												{
													x := (v.(bool))
													o.AcceptSingleValue = &x
												}
											}
											if v, ok := l["additional_properties"]; ok {
												{
													x := []byte(v.(string))
													var x1 interface{}
													err := json.Unmarshal(x, &x1)
													if err == nil && x1 != nil {
														o.ContentParameterAO0P0.ContentParameterAO0P0 = x1.(map[string]interface{})
													}
												}
											}
											if v, ok := l["complex_type"]; ok {
												{
													x := (v.(string))
													o.ComplexType = x
												}
											}
											if v, ok := l["item_type"]; ok {
												{
													x := (v.(string))
													o.ItemType = &x
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
											if v, ok := l["path"]; ok {
												{
													x := (v.(string))
													o.Path = x
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
									o.ErrorParameters = x
								}
							}
							if v, ok := l["object_type"]; ok {
								{
									x := (v.(string))
									o.ObjectType = x
								}
							}
							if v, ok := l["parameters"]; ok {
								{
									x := make([]*models.ContentParameter, 0)
									switch reflect.TypeOf(v).Kind() {
									case reflect.Slice:
										s := reflect.ValueOf(v)
										for i := 0; i < s.Len(); i++ {
											o := models.ContentParameter{}
											l := s.Index(i).Interface().(map[string]interface{})
											if v, ok := l["accept_single_value"]; ok {
												{
													x := (v.(bool))
													o.AcceptSingleValue = &x
												}
											}
											if v, ok := l["additional_properties"]; ok {
												{
													x := []byte(v.(string))
													var x1 interface{}
													err := json.Unmarshal(x, &x1)
													if err == nil && x1 != nil {
														o.ContentParameterAO0P0.ContentParameterAO0P0 = x1.(map[string]interface{})
													}
												}
											}
											if v, ok := l["complex_type"]; ok {
												{
													x := (v.(string))
													o.ComplexType = x
												}
											}
											if v, ok := l["item_type"]; ok {
												{
													x := (v.(string))
													o.ItemType = &x
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
											if v, ok := l["path"]; ok {
												{
													x := (v.(string))
													o.Path = x
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
									o.Parameters = x
								}
							}
							if v, ok := l["types"]; ok {
								{
									x := make([]*models.ContentComplexType, 0)
									switch reflect.TypeOf(v).Kind() {
									case reflect.Slice:
										s := reflect.ValueOf(v)
										for i := 0; i < s.Len(); i++ {
											o := models.ContentComplexType{}
											l := s.Index(i).Interface().(map[string]interface{})
											if v, ok := l["additional_properties"]; ok {
												{
													x := []byte(v.(string))
													var x1 interface{}
													err := json.Unmarshal(x, &x1)
													if err == nil && x1 != nil {
														o.ContentComplexTypeAO0P0.ContentComplexTypeAO0P0 = x1.(map[string]interface{})
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
											if v, ok := l["parameters"]; ok {
												{
													x := make([]*models.ContentParameter, 0)
													switch reflect.TypeOf(v).Kind() {
													case reflect.Slice:
														s := reflect.ValueOf(v)
														for i := 0; i < s.Len(); i++ {
															o := models.ContentParameter{}
															l := s.Index(i).Interface().(map[string]interface{})
															if v, ok := l["accept_single_value"]; ok {
																{
																	x := (v.(bool))
																	o.AcceptSingleValue = &x
																}
															}
															if v, ok := l["additional_properties"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		o.ContentParameterAO0P0.ContentParameterAO0P0 = x1.(map[string]interface{})
																	}
																}
															}
															if v, ok := l["complex_type"]; ok {
																{
																	x := (v.(string))
																	o.ComplexType = x
																}
															}
															if v, ok := l["item_type"]; ok {
																{
																	x := (v.(string))
																	o.ItemType = &x
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
															if v, ok := l["path"]; ok {
																{
																	x := (v.(string))
																	o.Path = x
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
													o.Parameters = x
												}
											}
											x = append(x, &o)
										}
									}
									o.Types = x
								}
							}

							p = o
						}
						x := p
						o.ResponseSpec = &x
					}
				}
				if v, ok := l["timeout"]; ok {
					{
						x := int64(v.(int))
						o.Timeout = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Batch = x
	}

	if d.HasChange("constraints") {
		v := d.Get("constraints")
		x := (v.(string))
		o.Constraints = x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
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

	if d.HasChange("output") {
		v := d.Get("output")
		x := v
		o.Output = &x
	}

	if d.HasChange("skip_on_condition") {
		v := d.Get("skip_on_condition")
		x := (v.(string))
		o.SkipOnCondition = x
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

	if d.HasChange("task_definition") {
		v := d.Get("task_definition")
		p := models.WorkflowTaskDefinitionRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowTaskDefinitionRef{}
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
		o.TaskDefinition = &x
	}

	url := "workflow/BatchApiExecutors" + "/" + d.Id()
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
	return resourceWorkflowBatchApiExecutorRead(d, meta)
}

func resourceWorkflowBatchApiExecutorDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "workflow/BatchApiExecutors" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
