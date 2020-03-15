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
				Description: "If the os.ConfigurationFile MO selected is a template that uses additionalplaceholders other than the ones provided in standard os.Answers MO, the valuesfor those additional placeholders are provided here.",
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
							Description: "If the source of the answers is a static file, the content of the file is stored as valuein this property.The value is mandatory only when the 'Source' property has been set to 'File'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"hostname": {
							Description: "Hostname to be configured for the server in the OS.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"ip_config_type": {
							Description: "IP configuration type. Values are Static or Dynamic configuration of IP.In case of static IP configuration, IP address, gateway and other details needto be populated. In case of dynamic the IP configuration is obtained dynamicallyfrom DHCP.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "static",
							ForceNew:    true,
						},
						"ipv4_config": {
							Description: "In case of static IP configuration, IP address, netmask and gateway details areprovided.",
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
									"gateway": {
										Description: "The IPv4 address of the default gateway.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"ip_address": {
										Description: "The IPv4 Address, represented in the standard dot-decimal notation, e.g. 192.168.1.3.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"netmask": {
										Description: "The IPv4 Netmask, represented in the standard dot-decimal notation, e.g. 255.255.255.0.",
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
							Description: "",
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
						},
						"nameserver": {
							Description: "IP address of the name server to be configured in the OS.",
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
						"product_key": {
							Description: "The product key to be used for a specific version of Windows installation.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"root_password": {
							Description: "Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password.Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password.For more details on encrypting passwords, see Help Center.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"source": {
							Description: "Answer values can be provided from three sources - Embedded in OS image, static file,or as placeholder values for an answer file template.Source of the answers is given as value, Embedded/File/Template.'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'option indicates that the answers are provided as a file. 'Template' indicates that theplaceholders in the selected os.ConfigurationFile MO are replaced with values providedas os.Answers MO.",
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
			"install_method": {
				Description: "The install method to be used for OS installation - vMedia, iPXE. Only vMedia is supported as of now.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
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
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"osdu_image": {
				Description: "Location of the Intersight OS Deployment Utilityimage, if the user has downloaded and available locally, to be used for this OS install configuration. This image is applicable for vMedia install method.Cisco publishes a OS Deployment Utility image that bootstraps and installs the user provided operating system images along with answers for unattended instllation.If this property is empty for vMedia install type, the image hosted in Intersight image repository will be used. Note that in this case, the image will be downloaded from Intersight in every target server every time.",
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
			"server": {
				Description: "This relation provides the target server in which the OS is to beinstalled.",
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
			"workflow_info": {
				Description: "This relation is populated with the reference of OS install workflowstarted for this request. This workflow info MO shall be used fortracking further status and completion.",
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
			if v, ok := l["ipv4_config"]; ok {
				{
					p := models.CommIPV4Interface{}
					if len(v.([]interface{})) > 0 {
						o := models.CommIPV4Interface{}
						l := (v.([]interface{})[0]).(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.CommIPV4InterfaceAO1P1.CommIPV4InterfaceAO1P1 = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["gateway"]; ok {
							{
								x := (v.(string))
								o.Gateway = x
							}
						}
						if v, ok := l["ip_address"]; ok {
							{
								x := (v.(string))
								o.IPAddress = x
							}
						}
						if v, ok := l["netmask"]; ok {
							{
								x := (v.(string))
								o.Netmask = x
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
					o.IPV4Config = &x
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
		o.Answers = &x

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
		o.ConfigurationFile = &x

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
		o.Image = &x

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
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.ObjectType = x
				}
			}

			p = o
		}
		x := p
		o.OperatingSystemParameters = &x

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
		o.Organization = &x

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
		o.OsduImage = &x

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
		o.Server = &x

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
		o.WorkflowInfo = &x

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
