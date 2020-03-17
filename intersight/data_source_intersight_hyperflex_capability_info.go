package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHyperflexCapabilityInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHyperflexCapabilityInfoRead,
		Schema: map[string]*schema.Schema{
			"app_catalog": {
				Description: "A collection of references to the [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) Managed Object.\nWhen this managed object is deleted, the referenced [hyperflex.AppCatalog](mo://hyperflex.AppCatalog) MO unsets its reference to this deleted MO.",
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
				Computed: true,
			},
			"capability_constraints": {
				Description: "Collection of constraints, which when applied together in tandem with an \"AND\" assertion, will correspond to the specified Value. Hence the Value will make sense only iff all the constraints match.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"constraint_name": {
							Description: "Name or key of the applicable compatibility constraint.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"constraint_value": {
							Description: "Value of the applicable compatibility constraint. Could either be a string value or a regex.",
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
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the capability or feature set consisting of a collection of constraint rules and value.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Computed: true,
			},
			"value": {
				Description: "Capability Value which is valid only iff all specified constraints match.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceHyperflexCapabilityInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "hyperflex/CapabilityInfos"
	var o models.HyperflexCapabilityInfo
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
	if v, ok := d.GetOk("value"); ok {
		x := (v.(string))
		o.Value = x
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
			var s models.HyperflexCapabilityInfo
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("app_catalog", flattenMapHyperflexAppCatalogRef(s.AppCatalog, d)); err != nil {
				return err
			}

			if err := d.Set("capability_constraints", flattenListHclConstraint(s.CapabilityConstraints, d)); err != nil {
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

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("value", (s.Value)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
