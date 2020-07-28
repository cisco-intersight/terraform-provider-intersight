package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiApicCcoPost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiApicCcoPostRead,
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
						"nr_version": {
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
		},
	}
}

func dataSourceNiaapiApicCcoPostRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewNiaapiApicCcoPostWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("post_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetPostDate(x)
	}
	if v, ok := d.GetOk("post_type"); ok {
		x := (v.(string))
		o.SetPostType(x)
	}
	if v, ok := d.GetOk("postid"); ok {
		x := (v.(string))
		o.SetPostid(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.NiaapiApi.GetNiaapiApicCcoPostList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.NiaapiApicCcoPostList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to NiaapiApicCcoPost: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewNiaapiApicCcoPostWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("post_date", (s.GetPostDate()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property PostDate: %+v", err)
			}

			if err := d.Set("post_detail", flattenMapNiaapiNewReleaseDetail(s.GetPostDetail(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PostDetail: %+v", err)
			}
			if err := d.Set("post_type", (s.GetPostType())); err != nil {
				return fmt.Errorf("error occurred while setting property PostType: %+v", err)
			}
			if err := d.Set("postid", (s.GetPostid())); err != nil {
				return fmt.Errorf("error occurred while setting property Postid: %+v", err)
			}
			if err := d.Set("revision", (s.GetRevision())); err != nil {
				return fmt.Errorf("error occurred while setting property Revision: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
