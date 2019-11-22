package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceTaskWorkflowAction() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskWorkflowActionRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "",
				Type:        schema.TypeList,
				MaxItems:    1,
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
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"action": {
				Description: "Action for test workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"input_params": {
				Description: "Json formatted string input parameters to workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_dynamic": {
				Description: "When true, this workflow type will be triggered as a dynamic workflow, if not it will be treated as a static workflow.",
				Type:        schema.TypeBool,
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
			"wait_on_duplicate": {
				Description: "When true, the workflow will wait for previous one to complete before starting a new one.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"workflow_context": {
				Description: "Json formatted string that has the contents of the workflow context used when starting a workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"workflow_file": {
				Description: "Path to workflow metadata file that will be published and started.",
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
						"contents": {
							Description: "When the type of the download is inline, it will read the file from the contents here.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"path": {
							Description: "The path of the download from the specified source location and if type is S3, then this is the object key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"source": {
							Description: "The source of the download location and if type is S3, then this is the bucket name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"type": {
							Description: "The type of download location is captured in type.",
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
func dataSourceTaskWorkflowActionRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "task/WorkflowActions"
	var o models.TaskWorkflowAction
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.Action = &x
	}
	if v, ok := d.GetOk("input_params"); ok {
		x := (v.(string))
		o.InputParams = x
	}
	if v, ok := d.GetOk("is_dynamic"); ok {
		x := (v.(bool))
		o.IsDynamic = &x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("wait_on_duplicate"); ok {
		x := (v.(bool))
		o.WaitOnDuplicate = &x
	}
	if v, ok := d.GetOk("workflow_context"); ok {
		x := (v.(string))
		o.WorkflowContext = x
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
			var s models.TaskWorkflowAction
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
				return err
			}
			if err := d.Set("action", (s.Action)); err != nil {
				return err
			}
			if err := d.Set("input_params", (s.InputParams)); err != nil {
				return err
			}
			if err := d.Set("is_dynamic", (s.IsDynamic)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("wait_on_duplicate", (s.WaitOnDuplicate)); err != nil {
				return err
			}
			if err := d.Set("workflow_context", (s.WorkflowContext)); err != nil {
				return err
			}

			if err := d.Set("workflow_file", flattenMapTaskFileDownloadInfo(s.WorkflowFile, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
