package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOsInstall() *schema.Resource {
	return &schema.Resource{
		Create: resourceOsInstallCreate,
		Read:   resourceOsInstallRead,
		Delete: resourceOsInstallDelete,
		Schema: map[string]*schema.Schema{
			"additional_parameters": {
				Description: "If the os.ConfigurationFile MO selected is a template that uses additional\nplaceholders other than the ones provided in standard os.Answers MO, the values\nfor those additional placeholders are provided here.",
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"is_value_set": {
							Description: "Flag to indicate if value is set. Value will be used to check if any edit.",
							Type:        schema.TypeBool,
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
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
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
												"class_id": {
													Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													ForceNew:    true,
												},
												"object_type": {
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
												"class_id": {
													Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													ForceNew:    true,
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
															"class_id": {
																Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																ForceNew:    true,
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
																		"class_id": {
																			Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			ForceNew:    true,
																		},
																		"label": {
																			Description: "Label for the enum value. A user friendly short string to identify the enum value.",
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
																Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
															"class_id": {
																Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																ForceNew:    true,
															},
															"display_attributes": {
																Description: "List of properties from an Intersight object which can help to identify the object. Typically the set of identity constraints on the object can be listed here to help the user identity the managed object.",
																Type:        schema.TypeList,
																Optional:    true,
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
													Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													ForceNew:    true,
												},
												"secure": {
													Description: "Intersight supports secure properties as task input/output. The values of\nthese properties are encrypted and stored in Intersight.\nThis flag marks the property to be secure when it is set to true.",
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
				Computed:   true,
				ForceNew:   true,
			},
			"answers": {
				Description: "Answers provided by user for the unattended OS installation.",
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
						"answer_file": {
							Description: "If the source of the answers is a static file, the content of the file is stored as value\nin this property.\nThe value is mandatory only when the 'Source' property has been set to 'File'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"hostname": {
							Description: "Hostname to be configured for the server in the OS.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"ip_config_type": {
							Description: "IP configuration type. Values are Static or Dynamic configuration of IP.\nIn case of static IP configuration, IP address, gateway and other details need\nto be populated. In case of dynamic the IP configuration is obtained dynamically\nfrom DHCP.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "static",
							ForceNew:    true,
						},
						"ip_configuration": {
							Description: "In case of static IP configuration, IP address, netmask and gateway details are\nprovided.",
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
									"class_id": {
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
						"is_answer_file_set": {
							Description: "Indicates whether the value of the 'answerFile' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"is_root_password_crypted": {
							Description: "Enable to indicate Root Password provided is encrypted.",
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
						},
						"is_root_password_set": {
							Description: "Indicates whether the value of the 'rootPassword' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"nameserver": {
							Description: "IP address of the name server to be configured in the OS.",
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
						"product_key": {
							Description: "The product key to be used for a specific version of Windows installation.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"root_password": {
							Description: "Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password.\nIntersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password.\nFor more details on encrypting passwords, see Help Center.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"source": {
							Description: "Answer values can be provided from three sources - Embedded in OS image, static file,\nor as placeholder values for an answer file template.\nSource of the answers is given as value, Embedded/File/Template.\n'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'\noption indicates that the answers are provided as a file. 'Template' indicates that the\nplaceholders in the selected os.ConfigurationFile MO are replaced with values provided\nas os.Answers MO.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "None",
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"configuration_file": {
				Description: "If the answers source is selected as 'Template' in 'Answers' property, this relation provides the os.ConfigurationFile instance to be used for this OS install.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
			"description": {
				Description: "User provided description about the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"image": {
				Description: "OS Image to be installed as part of this OS installation.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
			"install_method": {
				Description: "The install method to be used for OS installation - vMedia, iPXE. \nOnly vMedia is supported as of now.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "vMedia",
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
				Description: "The name of the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"operating_system_parameters": {
				Description: "Parameters specific to selected OS.",
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
			"osdu_image": {
				Description: "Location of the Intersight OS Deployment Utilityimage, if the user has downloaded and available locally, to be used for this OS install configuration. This image is applicable for vMedia install method.\nCisco publishes a OS Deployment Utility image that bootstraps and installs the user provided operating system images along with answers for unattended instllation.\nIf this property is empty for vMedia install type, the image hosted in Intersight image repository will be used. Note that in this case, the image will be downloaded from Intersight in every target server every time.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
			"post_install_scripts": {
				Description: "Post Install Scripts to be executed specified in order.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
			"server": {
				Description: "This relation provides the target server in which the OS is to be\ninstalled.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
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
			"workflow_info": {
				Description: "This relation is populated with the reference of OS install workflow\nstarted for this request. This workflow info MO shall be used for\ntracking further status and completion.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
		},
	}
}
func resourceOsInstallCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.OsInstall
	if v, ok := d.GetOk("additional_parameters"); ok {
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
				if v, ok := l["class_id"]; ok {
					{
						x := (v.(string))
						o.ClassID = x
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
							if v, ok := l["class_id"]; ok {
								{
									x := (v.(string))
									o.ClassID = x
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
									if len(v.([]interface{})) > 0 {
										o.Default = &x
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
										if v, ok := l["class_id"]; ok {
											{
												x := (v.(string))
												o.ClassID = x
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
													if v, ok := l["class_id"]; ok {
														{
															x := (v.(string))
															o.ClassID = x
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
																	if v, ok := l["class_id"]; ok {
																		{
																			x := (v.(string))
																			o.ClassID = x
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
												if len(v.([]interface{})) > 0 {
													o.Constraints = &x
												}
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
														if v, ok := l["class_id"]; ok {
															{
																x := (v.(string))
																o.ClassID = x
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
									if len(v.([]interface{})) > 0 {
										o.Properties = &x
									}
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
						if len(v.([]interface{})) > 0 {
							o.Type = &x
						}
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
		o.AdditionalParameters = x

	}

	if v, ok := d.GetOk("answers"); ok {
		p := models.OsAnswers{}
		if len(v.([]interface{})) > 0 {
			o := models.OsAnswers{}
			l := (v.([]interface{})[0]).(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.OsAnswersAO1P1.OsAnswersAO1P1 = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["answer_file"]; ok {
				{
					x := (v.(string))
					o.AnswerFile = x
				}
			}
			if v, ok := l["class_id"]; ok {
				{
					x := (v.(string))
					o.ClassID = x
				}
			}
			if v, ok := l["hostname"]; ok {
				{
					x := (v.(string))
					o.Hostname = x
				}
			}
			if v, ok := l["ip_config_type"]; ok {
				{
					x := (v.(string))
					o.IPConfigType = &x
				}
			}
			if v, ok := l["ip_configuration"]; ok {
				{
					p := models.OsIPConfiguration{}
					if len(v.([]interface{})) > 0 {
						o := models.OsIPConfiguration{}
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
						o.IPConfiguration = &x
					}
				}
			}
			if v, ok := l["is_answer_file_set"]; ok {
				{
					x := (v.(bool))
					o.IsAnswerFileSet = &x
				}
			}
			if v, ok := l["is_root_password_crypted"]; ok {
				{
					x := (v.(bool))
					o.IsRootPasswordCrypted = &x
				}
			}
			if v, ok := l["is_root_password_set"]; ok {
				{
					x := (v.(bool))
					o.IsRootPasswordSet = &x
				}
			}
			if v, ok := l["nameserver"]; ok {
				{
					x := (v.(string))
					o.Nameserver = x
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}
			if v, ok := l["product_key"]; ok {
				{
					x := (v.(string))
					o.ProductKey = x
				}
			}
			if v, ok := l["root_password"]; ok {
				{
					x := (v.(string))
					o.RootPassword = x
				}
			}
			if v, ok := l["source"]; ok {
				{
					x := (v.(string))
					o.Source = &x
				}
			}

			p = o
		}
		x := p
		if len(v.([]interface{})) > 0 {
			o.Answers = &x
		}

	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x

	}

	if v, ok := d.GetOk("configuration_file"); ok {
		p := models.OsConfigurationFileRef{}
		if len(v.([]interface{})) > 0 {
			o := models.OsConfigurationFileRef{}
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
			o.ConfigurationFile = &x
		}

	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x

	}

	if v, ok := d.GetOk("image"); ok {
		p := models.SoftwarerepositoryOperatingSystemFileRef{}
		if len(v.([]interface{})) > 0 {
			o := models.SoftwarerepositoryOperatingSystemFileRef{}
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
			o.Image = &x
		}

	}

	if v, ok := d.GetOk("install_method"); ok {
		x := (v.(string))
		o.InstallMethod = &x

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

	if v, ok := d.GetOk("operating_system_parameters"); ok {
		p := models.OsOperatingSystemParameters{}
		if len(v.([]interface{})) > 0 {
			o := models.OsOperatingSystemParameters{}
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
			o.OperatingSystemParameters = &x
		}

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

	if v, ok := d.GetOk("osdu_image"); ok {
		p := models.FirmwareServerConfigurationUtilityDistributableRef{}
		if len(v.([]interface{})) > 0 {
			o := models.FirmwareServerConfigurationUtilityDistributableRef{}
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
			o.OsduImage = &x
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

	if v, ok := d.GetOk("post_install_scripts"); ok {
		x := make([]*models.OsPostInstallScriptRef, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.OsPostInstallScriptRef{}
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
		o.PostInstallScripts = x

	}

	if v, ok := d.GetOk("server"); ok {
		p := models.ComputePhysicalRef{}
		if len(v.([]interface{})) > 0 {
			o := models.ComputePhysicalRef{}
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
			o.Server = &x
		}

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

	if v, ok := d.GetOk("workflow_info"); ok {
		p := models.WorkflowWorkflowInfoRef{}
		if len(v.([]interface{})) > 0 {
			o := models.WorkflowWorkflowInfoRef{}
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
			o.WorkflowInfo = &x
		}

	}

	url := "os/Installs"
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
	return resourceOsInstallRead(d, meta)
}

func resourceOsInstallRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "os/Installs" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.OsInstall
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
		return err
	}

	if err := d.Set("additional_parameters", flattenListOsPlaceHolder(s.AdditionalParameters, d)); err != nil {
		return err
	}

	if err := d.Set("answers", flattenMapOsAnswers(s.Answers, d)); err != nil {
		return err
	}

	if err := d.Set("class_id", (s.ClassID)); err != nil {
		return err
	}

	if err := d.Set("configuration_file", flattenMapOsConfigurationFileRef(s.ConfigurationFile, d)); err != nil {
		return err
	}

	if err := d.Set("description", (s.Description)); err != nil {
		return err
	}

	if err := d.Set("image", flattenMapSoftwarerepositoryOperatingSystemFileRef(s.Image, d)); err != nil {
		return err
	}

	if err := d.Set("install_method", (s.InstallMethod)); err != nil {
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

	if err := d.Set("operating_system_parameters", flattenMapOsOperatingSystemParameters(s.OperatingSystemParameters, d)); err != nil {
		return err
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRef(s.Organization, d)); err != nil {
		return err
	}

	if err := d.Set("osdu_image", flattenMapFirmwareServerConfigurationUtilityDistributableRef(s.OsduImage, d)); err != nil {
		return err
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
		return err
	}

	if err := d.Set("post_install_scripts", flattenListOsPostInstallScriptRef(s.PostInstallScripts, d)); err != nil {
		return err
	}

	if err := d.Set("server", flattenMapComputePhysicalRef(s.Server, d)); err != nil {
		return err
	}

	if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRef(s.WorkflowInfo, d)); err != nil {
		return err
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}

func resourceOsInstallDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	x := d.Get("workflow_info").([]interface{})[0].(map[string]interface{})
	url := "workflow/WorkflowInfos" + "/" + (x["moid"].(string))
	data, err := conn.SendGetRequest(url, []byte(""))
	var dData map[string]interface{}
	err = json.Unmarshal(data, &dData)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
		return err
	}
	if v, ok := dData["Status"]; ok && v == "RUNNING" {
		payload := map[string]string{"Action": "Cancel"}
		payloadJson, _ := json.Marshal(payload)
		_, err = conn.SendUpdateRequest(url, payloadJson)
		if err != nil {
			log.Printf("error occurred while deleting: %s", err.Error())
		}
	}
	return err
}
