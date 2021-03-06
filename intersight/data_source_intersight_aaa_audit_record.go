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

func dataSourceAaaAuditRecord() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAaaAuditRecordRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"email": {
				Description: "The email of the associated user that made the change.  In case the user is later deleted, we still have some reference to the information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"event": {
				Description: "The operation that was performed on this Managed Object.\nThe event is a compound string that includes the CRUD operation such as Create, Modify, Delete, and a string representing the Managed Object type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"inst_id": {
				Description: "The instance id of AuditRecordLocal, which is used to identify if the comming AuditRecordLocal was already processed before.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mo_type": {
				Description: "The object type of the REST resource that was created, modified or deleted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_moid": {
				Description: "The Moid of the REST resource that was created, modified or deleted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sessions": {
				Description: "A reference to a iamSession resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"source_ip": {
				Description: "The source IP of the client.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"timestamp": {
				Description: "The creation time of AuditRecordLocal, which is the time when the affected MO was created/modified/deleted.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"trace_id": {
				Description: "The trace id of the request that was used to create, modify or delete a REST resource.\nA trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purpose\nby the Intersight technical support team.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"user": {
				Description: "A reference to a iamUser resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"user_id_or_email": {
				Description: "The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceAaaAuditRecordRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewAaaAuditRecordWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("email"); ok {
		x := (v.(string))
		o.SetEmail(x)
	}
	if v, ok := d.GetOk("event"); ok {
		x := (v.(string))
		o.SetEvent(x)
	}
	if v, ok := d.GetOk("inst_id"); ok {
		x := (v.(string))
		o.SetInstId(x)
	}
	if v, ok := d.GetOk("mo_type"); ok {
		x := (v.(string))
		o.SetMoType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_moid"); ok {
		x := (v.(string))
		o.SetObjectMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("source_ip"); ok {
		x := (v.(string))
		o.SetSourceIp(x)
	}
	if v, ok := d.GetOk("timestamp"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetTimestamp(x)
	}
	if v, ok := d.GetOk("trace_id"); ok {
		x := (v.(string))
		o.SetTraceId(x)
	}
	if v, ok := d.GetOk("user_id_or_email"); ok {
		x := (v.(string))
		o.SetUserIdOrEmail(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.AaaApi.GetAaaAuditRecordList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.AaaAuditRecordList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to AaaAuditRecord: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewAaaAuditRecordWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}

			if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Account: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("email", (s.GetEmail())); err != nil {
				return fmt.Errorf("error occurred while setting property Email: %+v", err)
			}
			if err := d.Set("event", (s.GetEvent())); err != nil {
				return fmt.Errorf("error occurred while setting property Event: %+v", err)
			}
			if err := d.Set("inst_id", (s.GetInstId())); err != nil {
				return fmt.Errorf("error occurred while setting property InstId: %+v", err)
			}
			if err := d.Set("mo_type", (s.GetMoType())); err != nil {
				return fmt.Errorf("error occurred while setting property MoType: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_moid", (s.GetObjectMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectMoid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("sessions", flattenMapIamSessionRelationship(s.GetSessions(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Sessions: %+v", err)
			}
			if err := d.Set("source_ip", (s.GetSourceIp())); err != nil {
				return fmt.Errorf("error occurred while setting property SourceIp: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}

			if err := d.Set("timestamp", (s.GetTimestamp()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property Timestamp: %+v", err)
			}
			if err := d.Set("trace_id", (s.GetTraceId())); err != nil {
				return fmt.Errorf("error occurred while setting property TraceId: %+v", err)
			}

			if err := d.Set("user", flattenMapIamUserRelationship(s.GetUser(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property User: %+v", err)
			}
			if err := d.Set("user_id_or_email", (s.GetUserIdOrEmail())); err != nil {
				return fmt.Errorf("error occurred while setting property UserIdOrEmail: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
