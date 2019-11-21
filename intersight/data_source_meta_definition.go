package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceMetaDefinition() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMetaDefinitionRead,
		Schema: map[string]*schema.Schema{
			"access_privileges": {
				Description: "The list of access privileges that are required to perform CRUD operations on this managed object. If no access privileges are specified, the object is not accessible.",
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
						"method": {
							Description: "The type of CRUD operation (create, read, update, delete) for which an access privilege is required.",
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
						"privilege": {
							Description: "The name of the privilege which is required to invoke the specified CRUD method.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ancestor_classes": {
				Description: "An array of parent metaclasses in the class inheritance hierarchy. The first element in the array is the parent class. The next element is the grand-parent, etc. The last element in the array is the mo.BaseMo class.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"properties": {
				Description: "Meta definition for the properties in the meta class and from all classes in the inheritance hierarchy.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_access": {
							Description: "API access control for the property. Examples are NoAccess, ReadOnly, CreateOnly etc.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"name": {
							Description: "The name of the property.",
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
						"op_security": {
							Description: "The data-at-rest security protection applied to this property when a Managed Object is persisted.For example, Encrypted or Cleartext.",
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
			},
			"relationships": {
				Description: "Meta definition for the relationship in the meta class.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_access": {
							Description: "API access definition for this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"collection": {
							Description: "Specifies whether the relationship is a collection.",
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			},
			"rest_path": {
				Description: "Restful URL path for the meta.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"version": {
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

	url := "meta/Definitions"
	var o models.MetaDefinition
	if v, ok := d.GetOk("is_concrete"); ok {
		x := (v.(bool))
		o.IsConcrete = &x
	}
	if v, ok := d.GetOk("meta_type"); ok {
		x := (v.(string))
		o.MetaType = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.Name = x
	}
	if v, ok := d.GetOk("namespace"); ok {
		x := (v.(string))
		o.Namespace = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("parent_class"); ok {
		x := (v.(string))
		o.ParentClass = x
	}
	if v, ok := d.GetOk("rest_path"); ok {
		x := (v.(string))
		o.RestPath = x
	}
	if v, ok := d.GetOk("version"); ok {
		x := (v.(string))
		o.Version = x
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
			var s models.MetaDefinition
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("access_privileges", flattenListMetaAccessPrivilege(s.AccessPrivileges, d)); err != nil {
				return err
			}
			if err := d.Set("ancestor_classes", (s.AncestorClasses)); err != nil {
				return err
			}
			if err := d.Set("is_concrete", (s.IsConcrete)); err != nil {
				return err
			}
			if err := d.Set("meta_type", (s.MetaType)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("namespace", (s.Namespace)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("parent_class", (s.ParentClass)); err != nil {
				return err
			}

			if err := d.Set("properties", flattenListMetaPropDefinition(s.Properties, d)); err != nil {
				return err
			}

			if err := d.Set("relationships", flattenListMetaRelationshipDefinition(s.Relationships, d)); err != nil {
				return err
			}
			if err := d.Set("rest_path", (s.RestPath)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("version", (s.Version)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
