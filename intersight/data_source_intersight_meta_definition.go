package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMetaDefinition() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMetaDefinitionRead,
		Schema: map[string]*schema.Schema{
			"access_privileges": {
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
						"method": {
							Description: "The type of CRUD operation (create, read, update, delete) for which an access privilege is required.",
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
						"privilege": {
							Description: "The name of the privilege which is required to invoke the specified CRUD method.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"ancestor_classes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_concrete": {
				Description: "Boolean flag to specify whether the meta class is a concrete class or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"meta_type": {
				Description: "Indicates whether the meta class is a complex type or managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The fully-qualified class name of the Managed Object or complex type. For example, \"compute:Blade\" where the Managed Object is \"Blade\" and the package is 'compute'.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"namespace": {
				Description: "The namespace of the meta.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parent_class": {
				Description: "The fully-qualified name of the parent metaclass in the class inheritance hierarchy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"permission_supported": {
				Description: "Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"properties": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_access": {
							Description: "API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "The name of the property.",
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
						"op_security": {
							Description: "The data-at-rest security protection applied to this property when a Managed Object is persisted.\nFor example, Encrypted or Cleartext.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"search_weight": {
							Description: "Enables the property to be searchable from global search. A value of 0 means this property is not globally searchable.",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"rbac_resource": {
				Description: "Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource=yes. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"relationships": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_access": {
							Description: "API access definition for this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"collection": {
							Description: "Specifies whether the relationship is a collection.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"export": {
							Description: "When turned off, the peer MO is not exported when the local MO is exported.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"export_with_peer": {
							Description: "When turned on, the local MO is exported when the peer is exported.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "The name of the relationship.",
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
						"type": {
							Description: "Fully qualified type of the foreign managed object.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"rest_path": {
				Description: "Restful URL path for the meta.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"nr_version": {
				Description: "The version of the service that defines the meta-data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceMetaDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewMetaDefinitionWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("is_concrete"); ok {
		x := (v.(bool))
		o.SetIsConcrete(x)
	}
	if v, ok := d.GetOk("meta_type"); ok {
		x := (v.(string))
		o.SetMetaType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("namespace"); ok {
		x := (v.(string))
		o.SetNamespace(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("parent_class"); ok {
		x := (v.(string))
		o.SetParentClass(x)
	}
	if v, ok := d.GetOk("permission_supported"); ok {
		x := (v.(bool))
		o.SetPermissionSupported(x)
	}
	if v, ok := d.GetOk("rbac_resource"); ok {
		x := (v.(bool))
		o.SetRbacResource(x)
	}
	if v, ok := d.GetOk("rest_path"); ok {
		x := (v.(string))
		o.SetRestPath(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.MetaApi.GetMetaDefinitionList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.MetaDefinitionList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to MetaDefinition: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewMetaDefinitionWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}

			if err := d.Set("access_privileges", flattenListMetaAccessPrivilege(s.AccessPrivileges, d)); err != nil {
				return fmt.Errorf("error occurred while setting property AccessPrivileges: %+v", err)
			}
			if err := d.Set("ancestor_classes", (s.AncestorClasses)); err != nil {
				return fmt.Errorf("error occurred while setting property AncestorClasses: %+v", err)
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("is_concrete", (s.IsConcrete)); err != nil {
				return fmt.Errorf("error occurred while setting property IsConcrete: %+v", err)
			}
			if err := d.Set("meta_type", (s.MetaType)); err != nil {
				return fmt.Errorf("error occurred while setting property MetaType: %+v", err)
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("namespace", (s.Namespace)); err != nil {
				return fmt.Errorf("error occurred while setting property Namespace: %+v", err)
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("parent_class", (s.ParentClass)); err != nil {
				return fmt.Errorf("error occurred while setting property ParentClass: %+v", err)
			}
			if err := d.Set("permission_supported", (s.PermissionSupported)); err != nil {
				return fmt.Errorf("error occurred while setting property PermissionSupported: %+v", err)
			}

			if err := d.Set("properties", flattenListMetaPropDefinition(s.Properties, d)); err != nil {
				return fmt.Errorf("error occurred while setting property Properties: %+v", err)
			}
			if err := d.Set("rbac_resource", (s.RbacResource)); err != nil {
				return fmt.Errorf("error occurred while setting property RbacResource: %+v", err)
			}

			if err := d.Set("relationships", flattenListMetaRelationshipDefinition(s.Relationships, d)); err != nil {
				return fmt.Errorf("error occurred while setting property Relationships: %+v", err)
			}
			if err := d.Set("rest_path", (s.RestPath)); err != nil {
				return fmt.Errorf("error occurred while setting property RestPath: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("nr_version", (s.Version)); err != nil {
				return fmt.Errorf("error occurred while setting property Version: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
