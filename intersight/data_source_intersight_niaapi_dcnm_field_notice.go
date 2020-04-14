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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"date_published": {
							Description: "The date the revision is made.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.ClassID = x
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
			if err := d.Set("class_id", (s.ClassID)); err != nil {
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

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
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
