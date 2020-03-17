package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWorkflowWorkflowDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowWorkflowDefinitionCreate,
		Read:   resourceWorkflowWorkflowDefinitionRead,
		Update: resourceWorkflowWorkflowDefinitionUpdate,
		Delete: resourceWorkflowWorkflowDefinitionDelete,
		Schema: map[string]*schema.Schema{
			"catalog": {
				Description: "The catalog under which the definition is present.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"default_version": {
				Description: "When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"description": {
				Description: "The description for this workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"input_definition": {
				Description: "The schema expected for input parameters for this workflow.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
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
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"override": {
										Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"value": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
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
							Description: "Provide a detailed description of the data type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"required": {
							Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"label": {
				Description: "A user friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"license_entitlement": {
				Description: "License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"max_task_count": {
				Description: "The maximum number of tasks that can be executed on this workflow.",
				Type:        schema.TypeInt,
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
				Description: "The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"output_definition": {
				Description: "The schema expected for output parameters for this workflow.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
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
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"override": {
										Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"value": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
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
							Description: "Provide a detailed description of the data type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"required": {
							Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"output_parameters": {
				Description: "The output mappings for the workflow. The outputs for worflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is '${Source.output.JsonPath}'. 'Source' is the name of the task within the workflow. You can map any task output to a workflow output as long as the types are compatible. Following this is JSON path expression to extract JSON fragment from source's output.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"properties": {
				Description: "Type to capture the properties of a workflow definition. Some of these properties are passed to workflow execution instance.",
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
						"external_meta": {
							Description: "When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities.",
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"retryable": {
							Description: "When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"support_status": {
							Description: "Supported status of the definition.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Supported",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
			"tasks": {
				Description: "The tasks contained inside of the workflow.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"description": {
							Description: "The description of this task instance in the workflow.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier of the workflow task used for UI display.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "The name of the task within the workflow and it must be unique among all WorkflowTasks within a workflow definition. This name serves as the internal unique identifier for the task and is used to pick input and output parameters to feed into other tasks.",
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
			"ui_rendering_data": {
				Description: "This will hold the data needed for workflow to be rendered in the user interface.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"validation_information": {
				Description: "The current validation state and associated information for this workflow.",
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
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"state": {
							Description: "The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"validation_error": {
							Description: "List of all workflow or task validation errors. The validation errors can be for worker task or for control tasks.",
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
									"error_log": {
										Description: "Description of the error.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"field": {
										Description: "When populated this refers to the input or output field within the workflow or task.",
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
									"task_name": {
										Description: "The task name on which the error is found, when empty the error applies to the top level workflow.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"transition_name": {
										Description: "When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"version": {
				Description: "The version of the workflow to support multiple versions.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}
func resourceWorkflowWorkflowDefinitionCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.WorkflowWorkflowDefinition
	if v, ok := d.GetOk("catalog"); ok {
		p := models.WorkflowCatalogRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowCatalogRef{}
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

	if v, ok := d.GetOkExists("default_version"); ok {
		x := v.(bool)
		o.DefaultVersion = &x
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("input_definition"); ok {
		x := make([]*models.WorkflowBaseDataType, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowBaseDataType{}
				l := s.Index(i).Interface().(map[string]interface{})
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
				if v, ok := l["required"]; ok {
					{
						x := (v.(bool))
						o.Required = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.InputDefinition = x

	}

	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.Label = x

	}

	if v, ok := d.GetOk("license_entitlement"); ok {
		x := (v.(string))
		o.LicenseEntitlement = x

	}

	if v, ok := d.GetOk("max_task_count"); ok {
		x := int64(v.(int))
		o.MaxTaskCount = x

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

	if v, ok := d.GetOk("output_definition"); ok {
		x := make([]*models.WorkflowBaseDataType, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowBaseDataType{}
				l := s.Index(i).Interface().(map[string]interface{})
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
				if v, ok := l["required"]; ok {
					{
						x := (v.(bool))
						o.Required = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.OutputDefinition = x

	}

	if v, ok := d.GetOk("output_parameters"); ok {
		x := v
		o.OutputParameters = &x

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

	if v, ok := d.GetOk("properties"); ok {
		p := models.WorkflowWorkflowProperties{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowWorkflowProperties{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.WorkflowWorkflowPropertiesAO1P1.WorkflowWorkflowPropertiesAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["external_meta"]; ok {
				{
					x := (v.(bool))
					o.ExternalMeta = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["retryable"]; ok {
				{
					x := (v.(bool))
					o.Retryable = &x
				}
			}
			if v, ok := l["support_status"]; ok {
				{
					x := (v.(string))
					o.SupportStatus = &x
				}
			}

			p = o
		}
		x := p
		o.Properties = &x

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

	if v, ok := d.GetOk("tasks"); ok {
		x := make([]*models.WorkflowWorkflowTask, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowWorkflowTask{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1 = x1.(map[string]interface{})
						}
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
				x = append(x, &o)
			}
		}
		o.Tasks = x

	}

	if v, ok := d.GetOk("ui_rendering_data"); ok {
		x := v
		o.UIRenderingData = &x

	}

	if v, ok := d.GetOk("validation_information"); ok {
		p := models.WorkflowValidationInformation{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowValidationInformation{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["state"]; ok {
				{
					x := (v.(string))
					o.State = x
				}
			}
			if v, ok := l["validation_error"]; ok {
				{
					x := make([]*models.WorkflowValidationError, 0)
					switch reflect.TypeOf(v).Kind() {
					case reflect.Slice:
						s := reflect.ValueOf(v)
						for i := 0; i < s.Len(); i++ {
							o := models.WorkflowValidationError{}
							l := s.Index(i).Interface().(map[string]interface{})
							if v, ok := l["additional_properties"]; ok {
								{
									x := []byte(v.(string))
									var x1 interface{}
									err := json.Unmarshal(x, &x1)
									if err == nil && x1 != nil {
										o.WorkflowValidationErrorAO1P1.WorkflowValidationErrorAO1P1 = x1.(map[string]interface{})
									}
								}
							}
							if v, ok := l["error_log"]; ok {
								{
									x := (v.(string))
									o.ErrorLog = x
								}
							}
							if v, ok := l["field"]; ok {
								{
									x := (v.(string))
									o.Field = x
								}
							}
							if v, ok := l["object_type"]; ok {
								{
									x := (v.(string))
									o.ObjectType = x
								}
							}
							if v, ok := l["task_name"]; ok {
								{
									x := (v.(string))
									o.TaskName = x
								}
							}
							if v, ok := l["transition_name"]; ok {
								{
									x := (v.(string))
									o.TransitionName = x
								}
							}
							x = append(x, &o)
						}
					}
					o.ValidationError = x
				}
			}

			p = o
		}
		x := p
		o.ValidationInformation = &x

	}

	if v, ok := d.GetOk("version"); ok {
		x := int64(v.(int))
		o.Version = x

	}

	url := "workflow/WorkflowDefinitions"
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
	return resourceWorkflowWorkflowDefinitionRead(d, meta)
}

func resourceWorkflowWorkflowDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "workflow/WorkflowDefinitions" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.WorkflowWorkflowDefinition
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("catalog", flattenMapWorkflowCatalogRef(s.Catalog, d)); err != nil {
		return err
	}

	if err := d.Set("default_version", (s.DefaultVersion)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("input_definition", flattenListWorkflowBaseDataType(s.InputDefinition, d)); err != nil {
		return err
	}

	if err := d.Set("label", (s.Label)); err != nil {
		return err
	}

	if err := d.Set("license_entitlement", (s.LicenseEntitlement)); err != nil {
		return err
	}

	if err := d.Set("max_task_count", (s.MaxTaskCount)); err != nil {
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

	if err := d.Set("output_definition", flattenListWorkflowBaseDataType(s.OutputDefinition, d)); err != nil {
		return err
	}

	if err := d.Set("output_parameters", (s.OutputParameters)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("properties", flattenMapWorkflowWorkflowProperties(s.Properties, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("tasks", flattenListWorkflowWorkflowTask(s.Tasks, d)); err != nil {
		return err
	}

	if err := d.Set("ui_rendering_data", (s.UIRenderingData)); err != nil {
		return err
	}

	if err := d.Set("validation_information", flattenMapWorkflowValidationInformation(s.ValidationInformation, d)); err != nil {
		return err
	}

	if err := d.Set("version", (s.Version)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceWorkflowWorkflowDefinitionUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.WorkflowWorkflowDefinition
	if d.HasChange("catalog") {
		v := d.Get("catalog")
		p := models.WorkflowCatalogRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowCatalogRef{}
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

	if d.HasChange("default_version") {
		v := d.Get("default_version")
		x := (v.(bool))
		o.DefaultVersion = &x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.Description = x
	}

	if d.HasChange("input_definition") {
		v := d.Get("input_definition")
		x := make([]*models.WorkflowBaseDataType, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowBaseDataType{}
				l := s.Index(i).Interface().(map[string]interface{})
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
				if v, ok := l["required"]; ok {
					{
						x := (v.(bool))
						o.Required = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.InputDefinition = x
	}

	if d.HasChange("label") {
		v := d.Get("label")
		x := (v.(string))
		o.Label = x
	}

	if d.HasChange("license_entitlement") {
		v := d.Get("license_entitlement")
		x := (v.(string))
		o.LicenseEntitlement = x
	}

	if d.HasChange("max_task_count") {
		v := d.Get("max_task_count")
		x := int64(v.(int))
		o.MaxTaskCount = x
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

	if d.HasChange("output_definition") {
		v := d.Get("output_definition")
		x := make([]*models.WorkflowBaseDataType, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowBaseDataType{}
				l := s.Index(i).Interface().(map[string]interface{})
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
				if v, ok := l["required"]; ok {
					{
						x := (v.(bool))
						o.Required = &x
					}
				}
				x = append(x, &o)
			}
		}
		o.OutputDefinition = x
	}

	if d.HasChange("output_parameters") {
		v := d.Get("output_parameters")
		x := v
		o.OutputParameters = &x
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

	if d.HasChange("properties") {
		v := d.Get("properties")
		p := models.WorkflowWorkflowProperties{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowWorkflowProperties{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.WorkflowWorkflowPropertiesAO1P1.WorkflowWorkflowPropertiesAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["external_meta"]; ok {
				{
					x := (v.(bool))
					o.ExternalMeta = &x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["retryable"]; ok {
				{
					x := (v.(bool))
					o.Retryable = &x
				}
			}
			if v, ok := l["support_status"]; ok {
				{
					x := (v.(string))
					o.SupportStatus = &x
				}
			}

			p = o
		}
		x := p
		o.Properties = &x
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

	if d.HasChange("tasks") {
		v := d.Get("tasks")
		x := make([]*models.WorkflowWorkflowTask, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.WorkflowWorkflowTask{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1 = x1.(map[string]interface{})
						}
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
				x = append(x, &o)
			}
		}
		o.Tasks = x
	}

	if d.HasChange("ui_rendering_data") {
		v := d.Get("ui_rendering_data")
		x := v
		o.UIRenderingData = &x
	}

	if d.HasChange("validation_information") {
		v := d.Get("validation_information")
		p := models.WorkflowValidationInformation{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowValidationInformation{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["state"]; ok {
				{
					x := (v.(string))
					o.State = x
				}
			}
			if v, ok := l["validation_error"]; ok {
				{
					x := make([]*models.WorkflowValidationError, 0)
					switch reflect.TypeOf(v).Kind() {
					case reflect.Slice:
						s := reflect.ValueOf(v)
						for i := 0; i < s.Len(); i++ {
							o := models.WorkflowValidationError{}
							l := s.Index(i).Interface().(map[string]interface{})
							if v, ok := l["additional_properties"]; ok {
								{
									x := []byte(v.(string))
									var x1 interface{}
									err := json.Unmarshal(x, &x1)
									if err == nil && x1 != nil {
										o.WorkflowValidationErrorAO1P1.WorkflowValidationErrorAO1P1 = x1.(map[string]interface{})
									}
								}
							}
							if v, ok := l["error_log"]; ok {
								{
									x := (v.(string))
									o.ErrorLog = x
								}
							}
							if v, ok := l["field"]; ok {
								{
									x := (v.(string))
									o.Field = x
								}
							}
							if v, ok := l["object_type"]; ok {
								{
									x := (v.(string))
									o.ObjectType = x
								}
							}
							if v, ok := l["task_name"]; ok {
								{
									x := (v.(string))
									o.TaskName = x
								}
							}
							if v, ok := l["transition_name"]; ok {
								{
									x := (v.(string))
									o.TransitionName = x
								}
							}
							x = append(x, &o)
						}
					}
					o.ValidationError = x
				}
			}

			p = o
		}
		x := p
		o.ValidationInformation = &x
	}

	if d.HasChange("version") {
		v := d.Get("version")
		x := int64(v.(int))
		o.Version = x
	}

	url := "workflow/WorkflowDefinitions" + "/" + d.Id()
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
	return resourceWorkflowWorkflowDefinitionRead(d, meta)
}

func resourceWorkflowWorkflowDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "workflow/WorkflowDefinitions" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
