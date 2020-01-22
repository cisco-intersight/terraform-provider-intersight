package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceWorkflowBatchApiExecutor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceWorkflowBatchApiExecutorRead,
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
												},
											},
										},
										Computed: true,
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
												},
											},
										},
										Computed: true,
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
															},
														},
													},
													Computed: true,
												},
											},
										},
										Computed: true,
									},
								},
							},
							Computed: true,
						},
						"skip_on_condition": {
							Description: "The skip expression, if provided, allows the batch API executor to skip theapi execution when the given expression evaluates to true.The expression is given as such a golang template that has to beevaluated to a final content true/false. The expression is an optional and incase not provided, the API will always be executed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"timeout": {
							Description: "The duration in seconds by which the API response is expected from the API target.If the end point does not respond for the API request within this timeoutduration, the task will be marked as failed.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
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
			},
			"name": {
				Description: "Name for the batch API task.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Computed: true,
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
				Computed: true,
			},
		},
	}
}
func dataSourceWorkflowBatchApiExecutorRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/BatchApiExecutors"
	var o models.WorkflowBatchAPIExecutor
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
	if v, ok := d.GetOk("skip_on_condition"); ok {
		x := (v.(string))
		o.SkipOnCondition = x
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
			var s models.WorkflowBatchAPIExecutor
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
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

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
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
			d.SetId(s.Moid)
		}
	}
	return nil
}
