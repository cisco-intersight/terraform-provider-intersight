---
layout: "intersight"
page_title: "Intersight: intersight_os_install"
sidebar_current: "docs-intersight-resource-osInstall"
description: |-
  This MO models the target server, answers and other properties needed for
OS installation. The OS installation can be started in the target server by doing
a POST on this MO.
The requests to this MO starts a OS installation workflow that can be tracked
using workflow engine MO workflow.WorkflowInfo.
---

# Resource: intersight_os_install
This MO models the target server, answers and other properties needed for
OS installation. The OS installation can be started in the target server by doing
a POST on this MO.
The requests to this MO starts a OS installation workflow that can be tracked
using workflow engine MO workflow.WorkflowInfo.
## Argument Reference
The following arguments are supported:
* `additional_parameters`:(Array)"If the os.ConfigurationFile MO selected is a template that uses additional\nplaceholders other than the ones provided in standard os.Answers MO, the values\nfor those additional placeholders are provided here."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `is_value_set`:(bool)"Flag to indicate if value is set. Value will be used to check if any edit."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `type`:(Array with Maximum of one item) -"Definition of place holder."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
    + `default`:(Array with Maximum of one item) -"Default value for the data type. If default value was provided and the input was required the default value will be used as the input."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `override`:(bool)"Override the default value provided for the data type. When true, allow the user to enter value for the data type."
    + `value`:"Default value for the data type. If default value was provided and the input was required the default value will be used as the input."
  + `description`:(string)"Provide a detailed description of the data type."
  + `label`:(string)"Descriptive label for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character."
  + `name`:(string)"Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `properties`:(Array with Maximum of one item) -"Primitive data type properties."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
    + `constraints`:(Array with Maximum of one item) -"Constraints that must be applied to the parameter value supplied for this data type."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
    + `enum_list`:(Array)"When the parameter is a enum then this list of enum entry is used to validate the input belongs to one of items in the list."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
    + `label`:(string)"Label for the enum value. A user friendly short string to identify the enum value."
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `value`:(string)"Enum value for this enum entry. Value will be passed to the workflow as string type for execution."
  + `max`:(float)"Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked."
  + `min`:(float)"Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `regex`:(string)"When the parameter is a string this regular expression is used to ensure the value is valid."
  + `inventory_selector`:(Array)"List of Intersight managed object selectors. The workflow execution user interface show objects from inventory that are matching the selectors to help with selecting inputs."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
    + `display_attributes`:
                (Array of schema.TypeString) -"List of properties from an Intersight object which can help to identify the object. Typically the set of identity constraints on the object can be listed here to help the user identity the managed object."
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
    + `selector`:(string)"Field to hold an Intersight API along with an optional filter to narrow down the search options."
    + `value_attribute`:(string)"A property from the Intersight object, value of which can be used as value for referenced input definition."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `secure`:(bool)"Intersight supports secure properties as task input/output. The values of\nthese properties are encrypted and stored in Intersight.\nThis flag marks the property to be secure when it is set to true."
  + `type`:(string)"Specify the enum type for primitive data type."
  + `required`:(bool)"Specifies whether this parameter is required. The field is applicable for task and workflow."
  + `value`:"Value for placeholder provided by user."
* `answers`:(Array with Maximum of one item) -"Answers provided by user for the unattended OS installation."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `answer_file`:(string)"If the source of the answers is a static file, the content of the file is stored as value\nin this property.\nThe value is mandatory only when the 'Source' property has been set to 'File'."
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `hostname`:(string)"Hostname to be configured for the server in the OS."
  + `ip_config_type`:(string)"IP configuration type. Values are Static or Dynamic configuration of IP.\nIn case of static IP configuration, IP address, gateway and other details need\nto be populated. In case of dynamic the IP configuration is obtained dynamically\nfrom DHCP."
  + `ipv4_config`:(Array with Maximum of one item) -"In case of static IP configuration, IP address, netmask and gateway details are\nprovided."
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
    + `gateway`:(string)"The IPv4 address of the default gateway."
    + `ip_address`:(string)"The IPv4 Address, represented in the standard dot-decimal notation, e.g. 192.168.1.3."
    + `netmask`:(string)"The IPv4 Netmask, represented in the standard dot-decimal notation, e.g. 255.255.255.0."
    + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `is_answer_file_set`:(bool)(Computed)"Indicates whether the value of the 'answerFile' property has been set."
  + `is_root_password_crypted`:(bool)"Enable to indicate Root Password provided is encrypted."
  + `is_root_password_set`:(bool)(Computed)"Indicates whether the value of the 'rootPassword' property has been set."
  + `nameserver`:(string)"IP address of the name server to be configured in the OS."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `product_key`:(string)"The product key to be used for a specific version of Windows installation."
  + `root_password`:(string)"Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password.\nIntersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password.\nFor more details on encrypting passwords, see Help Center."
  + `source`:(string)"Answer values can be provided from three sources - Embedded in OS image, static file,\nor as placeholder values for an answer file template.\nSource of the answers is given as value, Embedded/File/Template.\n'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'\noption indicates that the answers are provided as a file. 'Template' indicates that the\nplaceholders in the selected os.ConfigurationFile MO are replaced with values provided\nas os.Answers MO."
* `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `configuration_file`:(Array with Maximum of one item) -"If the answers source is selected as 'Template' in 'Answers' property, this relation provides the os.ConfigurationFile instance to be used for this OS install."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `description`:(string)"User provided description about the OS install configuration."
* `image`:(Array with Maximum of one item) -"OS Image to be installed as part of this OS installation."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `install_method`:(string)"The install method to be used for OS installation - vMedia, iPXE. \nOnly vMedia is supported as of now."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"The name of the OS install configuration."
* `object_type`:(string)(Computed)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `operating_system_parameters`:(Array with Maximum of one item) -"Parameters specific to selected OS."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
* `organization`:(Array with Maximum of one item) -"Relationship to the Organization that owns the Managed Object."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `osdu_image`:(Array with Maximum of one item) -"Location of the Intersight OS Deployment Utilityimage, if the user has downloaded and available locally, to be used for this OS install configuration. This image is applicable for vMedia install method.\nCisco publishes a OS Deployment Utility image that bootstraps and installs the user provided operating system images along with answers for unattended instllation.\nIf this property is empty for vMedia install type, the image hosted in Intersight image repository will be used. Note that in this case, the image will be downloaded from Intersight in every target server every time."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `permission_resources`:(Array)(Computed)"A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `post_install_scripts`:(Array)"Post Install Scripts to be executed specified in order."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `server`:(Array with Maximum of one item) -"This relation provides the target server in which the OS is to be\ninstalled."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
* `tags`:(Array)"The array of tags, which allow to add key, value meta-data to managed objects."
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
  + `key`:(string)"The string representation of a tag key."
  + `object_type`:(string)"The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types."
  + `value`:(string)"The string representation of a tag value."
* `workflow_info`:(Array with Maximum of one item) -"This relation is populated with the reference of OS install workflow\nstarted for this request. This workflow info MO shall be used for\ntracking further status and completion."
This complex property has following sub-properties:
  + `moid`:(string)(Computed)"The Moid of the referenced REST resource."
  + `object_type`:(string)(Computed)"The Object Type of the referenced REST resource."
  + `selector`:(string)(Computed)"An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'."
