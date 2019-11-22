package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiApicCcoPost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiApicCcoPostRead,
		Schema: map[string]*schema.Schema{
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
			"post_date": {
				Description: "The date when this new release notice is posted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"post_detail": {
				Description: "Detail of this post including the content and the date it was posted.",
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
						"description": {
							Description: "Description of this new verison release post.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"link": {
							Description: "Link of downloading the release file.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"release_note_link": {
							Description: "The link used to provide a gateway for customer to review the release note.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"release_note_link_title": {
							Description: "The link title used to provide a gateway for customer to review the release note.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"software_download_link": {
							Description: "The link used to provide a gateway for customer to download the release.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"software_download_link_title": {
							Description: "The link title used to provide a gateway for customer to download the release.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"title": {
							Description: "Title of the new verison release post.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"version": {
							Description: "Version number Associate with this Post.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"post_type": {
				Description: "The document type of this post.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"postid": {
				Description: "Identificator of this inbox post.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"revision": {
				Description: "Revision number of this notice.",
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
		},
	}
}
func dataSourceNiaapiApicCcoPostRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "niaapi/ApicCcoPosts"
	var o models.NiaapiApicCcoPost
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("post_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.PostDate = x
	}
	if v, ok := d.GetOk("post_type"); ok {
		x := (v.(string))
		o.PostType = x
	}
	if v, ok := d.GetOk("postid"); ok {
		x := (v.(string))
		o.Postid = x
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.Revision = x
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
			var s models.NiaapiApicCcoPost
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("post_date", (s.PostDate).String()); err != nil {
				return err
			}

			if err := d.Set("post_detail", flattenMapNiaapiNewReleaseDetail(s.PostDetail, d)); err != nil {
				return err
			}
			if err := d.Set("post_type", (s.PostType)); err != nil {
				return err
			}
			if err := d.Set("postid", (s.Postid)); err != nil {
				return err
			}
			if err := d.Set("revision", (s.Revision)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
