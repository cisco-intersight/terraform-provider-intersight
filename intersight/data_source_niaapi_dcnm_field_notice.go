package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiDcnmFieldNotice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiDcnmFieldNoticeRead,
		Schema: map[string]*schema.Schema{
			"bugid": {
				Description: "Bug Id associated with this notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field_notice_desc": {
				Description: "Field notice Description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field_notice_id": {
				Description: "Fieldnotice Id of this notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field_notice_url": {
				Description: "Field notice URL link to the notice webpage.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"headline": {
				Description: "The headline of this field notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hwpid": {
				Description: "Hardware PID for affected models.",
				Type:        schema.TypeString,
				Optional:    true,
			},
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
			"revision_info": {
				Description: "Revision detail infomation .",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"date_published": {
							Description: "The date the revision is made.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"revision_comment": {
							Description: "The changes made in this revision.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"revision_no": {
							Description: "The Revision No. of this revision.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"sw_release": {
				Description: "Software Release number for affected versions.",
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
			"workaround_url": {
				Description: "URL of workaround of this notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceNiaapiDcnmFieldNoticeRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "niaapi/DcnmFieldNotices"
	var o models.NiaapiDcnmFieldNotice
	if v, ok := d.GetOk("bugid"); ok {
		x := (v.(string))
		o.Bugid = x
	}
	if v, ok := d.GetOk("field_notice_desc"); ok {
		x := (v.(string))
		o.FieldNoticeDesc = x
	}
	if v, ok := d.GetOk("field_notice_id"); ok {
		x := (v.(string))
		o.FieldNoticeID = x
	}
	if v, ok := d.GetOk("field_notice_url"); ok {
		x := (v.(string))
		o.FieldNoticeURL = x
	}
	if v, ok := d.GetOk("headline"); ok {
		x := (v.(string))
		o.Headline = x
	}
	if v, ok := d.GetOk("hwpid"); ok {
		x := (v.(string))
		o.Hwpid = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("sw_release"); ok {
		x := (v.(string))
		o.SwRelease = x
	}
	if v, ok := d.GetOk("workaround_url"); ok {
		x := (v.(string))
		o.WorkaroundURL = x
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
			var s models.NiaapiDcnmFieldNotice
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("bugid", (s.Bugid)); err != nil {
				return err
			}
			if err := d.Set("field_notice_desc", (s.FieldNoticeDesc)); err != nil {
				return err
			}
			if err := d.Set("field_notice_id", (s.FieldNoticeID)); err != nil {
				return err
			}
			if err := d.Set("field_notice_url", (s.FieldNoticeURL)); err != nil {
				return err
			}
			if err := d.Set("headline", (s.Headline)); err != nil {
				return err
			}
			if err := d.Set("hwpid", (s.Hwpid)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("revision_info", flattenListNiaapiRevisionInfo(s.RevisionInfo, d)); err != nil {
				return err
			}
			if err := d.Set("sw_release", (s.SwRelease)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("workaround_url", (s.WorkaroundURL)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
