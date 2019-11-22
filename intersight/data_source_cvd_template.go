package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCvdTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCvdTemplateRead,
		Schema: map[string]*schema.Schema{
			"deployer": {
				Description: "URL pointing to the S3 location of the workflow JSON which performs the deployment task for this CVD",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"deployer_input": {
				Description: "A collection of input name-value pairs",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"field_filter": {
							Description: "Input field filter(REST API path with filter which will return the list of applicable values for this field)",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"field_label": {
							Description: "Input field label(for GUI use)",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"field_name": {
							Description: "Input field name",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"field_value": {
							Description: "Input field value",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"description": {
				Description: "A short description for the CVD",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Unique name identifier for the CVD",
				Type:        schema.TypeString,
				Optional:    true,
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
			"upload_location": {
				Description: "S3 directory location where the CVD definition has been uploaded",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"validator": {
				Description: "URL pointing to the S3 location of the workflow JSON which performs the validation task for this CVD",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"validator_input": {
				Description: "A collection of input name-value pairs",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"field_filter": {
							Description: "Input field filter(REST API path with filter which will return the list of applicable values for this field)",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"field_label": {
							Description: "Input field label(for GUI use)",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"field_name": {
							Description: "Input field name",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"field_value": {
							Description: "Input field value",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"version": {
				Description: "The version or revision number of the CVD definition",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}
func dataSourceCvdTemplateRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "cvd/Templates"
	var o models.CvdTemplate
	if v, ok := d.GetOk("deployer"); ok {
		x := (v.(string))
		o.Deployer = x
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
	if v, ok := d.GetOk("upload_location"); ok {
		x := (v.(string))
		o.UploadLocation = x
	}
	if v, ok := d.GetOk("validator"); ok {
		x := (v.(string))
		o.Validator = x
	}
	if v, ok := d.GetOk("version"); ok {
		x := int64(v.(int))
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
			var s models.CvdTemplate
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("deployer", (s.Deployer)); err != nil {
				return err
			}

			if err := d.Set("deployer_input", flattenListCvdInputMeta(s.DeployerInput, d)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
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

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("upload_location", (s.UploadLocation)); err != nil {
				return err
			}
			if err := d.Set("validator", (s.Validator)); err != nil {
				return err
			}

			if err := d.Set("validator_input", flattenListCvdInputMeta(s.ValidatorInput, d)); err != nil {
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
