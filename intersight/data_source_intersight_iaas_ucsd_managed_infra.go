package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIaasUcsdManagedInfra() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIaasUcsdManagedInfraRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"advanced_catalog_count": {
				Description: "Total advanced catalogs in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"bm_catalog_count": {
				Description: "Total bare metal catalogs in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"container_catalog_count": {
				Description: "Total service container catalogs in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"esxi_host_count": {
				Description: "Total ESXi hosts in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"external_group_count": {
				Description: "Total external (Ldap) groups in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"guid": {
				Description: "A reference to a iaasUcsdInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"hyperv_host_count": {
				Description: "Total HyperV hosts in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"local_group_count": {
				Description: "Total local groups in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
			"standard_catalog_count": {
				Description: "Total standard catalogs in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
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
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"user_count": {
				Description: "Total user accounts in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"vdc_count": {
				Description: "Total virtual datacenters in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"vm_count": {
				Description: "Total Virtual machines in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceIaasUcsdManagedInfraRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewIaasUcsdManagedInfraWithDefaults()
	if v, ok := d.GetOk("advanced_catalog_count"); ok {
		x := int64(v.(int))
		o.SetAdvancedCatalogCount(x)
	}
	if v, ok := d.GetOk("bm_catalog_count"); ok {
		x := int64(v.(int))
		o.SetBmCatalogCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("container_catalog_count"); ok {
		x := int64(v.(int))
		o.SetContainerCatalogCount(x)
	}
	if v, ok := d.GetOk("esxi_host_count"); ok {
		x := int64(v.(int))
		o.SetEsxiHostCount(x)
	}
	if v, ok := d.GetOk("external_group_count"); ok {
		x := int64(v.(int))
		o.SetExternalGroupCount(x)
	}
	if v, ok := d.GetOk("hyperv_host_count"); ok {
		x := int64(v.(int))
		o.SetHypervHostCount(x)
	}
	if v, ok := d.GetOk("local_group_count"); ok {
		x := int64(v.(int))
		o.SetLocalGroupCount(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("standard_catalog_count"); ok {
		x := int64(v.(int))
		o.SetStandardCatalogCount(x)
	}
	if v, ok := d.GetOk("user_count"); ok {
		x := int64(v.(int))
		o.SetUserCount(x)
	}
	if v, ok := d.GetOk("vdc_count"); ok {
		x := int64(v.(int))
		o.SetVdcCount(x)
	}
	if v, ok := d.GetOk("vm_count"); ok {
		x := int64(v.(int))
		o.SetVmCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.IaasApi.GetIaasUcsdManagedInfraList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.IaasUcsdManagedInfraList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to IaasUcsdManagedInfra: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewIaasUcsdManagedInfraWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("advanced_catalog_count", (s.GetAdvancedCatalogCount())); err != nil {
				return fmt.Errorf("error occurred while setting property AdvancedCatalogCount: %+v", err)
			}
			if err := d.Set("bm_catalog_count", (s.GetBmCatalogCount())); err != nil {
				return fmt.Errorf("error occurred while setting property BmCatalogCount: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("container_catalog_count", (s.GetContainerCatalogCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ContainerCatalogCount: %+v", err)
			}
			if err := d.Set("esxi_host_count", (s.GetEsxiHostCount())); err != nil {
				return fmt.Errorf("error occurred while setting property EsxiHostCount: %+v", err)
			}
			if err := d.Set("external_group_count", (s.GetExternalGroupCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ExternalGroupCount: %+v", err)
			}

			if err := d.Set("guid", flattenMapIaasUcsdInfoRelationship(s.GetGuid(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Guid: %+v", err)
			}
			if err := d.Set("hyperv_host_count", (s.GetHypervHostCount())); err != nil {
				return fmt.Errorf("error occurred while setting property HypervHostCount: %+v", err)
			}
			if err := d.Set("local_group_count", (s.GetLocalGroupCount())); err != nil {
				return fmt.Errorf("error occurred while setting property LocalGroupCount: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("standard_catalog_count", (s.GetStandardCatalogCount())); err != nil {
				return fmt.Errorf("error occurred while setting property StandardCatalogCount: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("user_count", (s.GetUserCount())); err != nil {
				return fmt.Errorf("error occurred while setting property UserCount: %+v", err)
			}
			if err := d.Set("vdc_count", (s.GetVdcCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VdcCount: %+v", err)
			}
			if err := d.Set("vm_count", (s.GetVmCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VmCount: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
