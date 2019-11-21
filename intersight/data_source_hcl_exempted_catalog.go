package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceHclExemptedCatalog() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHclExemptedCatalogRead,
		Schema: map[string]*schema.Schema{
			"comments": {
				Description: "Reason for the exemption.",
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
				Description: "A unique descriptive name of the exemption.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"os_vendor": {
				Description: "Vendor of the Operating System.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"os_version": {
				Description: "Version of the Operating system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_name": {
				Description: "Name of the processor supported for the server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"product_models": {
				Description: "Models of the product/adapter.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"product_type": {
				Description: "Type of the product/adapter say PT for Pass Through controllers.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"server_pid": {
				Description: "Three part ID representing the server model as returned by UCSM/CIMC XML APIs.",
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
			"ucs_version": {
				Description: "Version of the UCS software.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version_type": {
				Description: "Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceHclExemptedCatalogRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "hcl/ExemptedCatalogs"
	var o models.HclExemptedCatalog
	if v, ok := d.GetOk("comments"); ok {
		x := (v.(string))
		o.Comments = x
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
	if v, ok := d.GetOk("os_vendor"); ok {
		x := (v.(string))
		o.OsVendor = x
	}
	if v, ok := d.GetOk("os_version"); ok {
		x := (v.(string))
		o.OsVersion = x
	}
	if v, ok := d.GetOk("processor_name"); ok {
		x := (v.(string))
		o.ProcessorName = x
	}
	if v, ok := d.GetOk("product_type"); ok {
		x := (v.(string))
		o.ProductType = x
	}
	if v, ok := d.GetOk("server_pid"); ok {
		x := (v.(string))
		o.ServerPid = x
	}
	if v, ok := d.GetOk("ucs_version"); ok {
		x := (v.(string))
		o.UcsVersion = x
	}
	if v, ok := d.GetOk("version_type"); ok {
		x := (v.(string))
		o.VersionType = x
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
			var s models.HclExemptedCatalog
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("comments", (s.Comments)); err != nil {
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
			if err := d.Set("os_vendor", (s.OsVendor)); err != nil {
				return err
			}
			if err := d.Set("os_version", (s.OsVersion)); err != nil {
				return err
			}
			if err := d.Set("processor_name", (s.ProcessorName)); err != nil {
				return err
			}
			if err := d.Set("product_models", (s.ProductModels)); err != nil {
				return err
			}
			if err := d.Set("product_type", (s.ProductType)); err != nil {
				return err
			}
			if err := d.Set("server_pid", (s.ServerPid)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("ucs_version", (s.UcsVersion)); err != nil {
				return err
			}
			if err := d.Set("version_type", (s.VersionType)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
