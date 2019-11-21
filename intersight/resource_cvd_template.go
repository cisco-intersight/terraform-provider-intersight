package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func resourceCvdTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCvdTemplateCreate,
		Read:   resourceCvdTemplateRead,
		Update: resourceCvdTemplateUpdate,
		Delete: resourceCvdTemplateDelete,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
				ForceNew:    true,
			},
			"name": {
				Description: "Unique name identifier for the CVD",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
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
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
			},
			"version": {
				Description: "The version or revision number of the CVD definition",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}
func resourceCvdTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.CvdTemplate
	if v, ok := d.GetOk("deployer"); ok {
		x := (v.(string))
		o.Deployer = x

	}

	if v, ok := d.GetOk("deployer_input"); ok {
		x := make([]*models.CvdInputMeta, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.CvdInputMeta{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.CvdInputMetaAO0P0.CvdInputMetaAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["field_filter"]; ok {
					{
						x := (v.(string))
						o.FieldFilter = x
					}
				}
				if v, ok := l["field_label"]; ok {
					{
						x := (v.(string))
						o.FieldLabel = x
					}
				}
				if v, ok := l["field_name"]; ok {
					{
						x := (v.(string))
						o.FieldName = x
					}
				}
				if v, ok := l["field_value"]; ok {
					{
						x := (v.(string))
						o.FieldValue = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				x = append(x, &o)
			}
		}
		o.DeployerInput = x

	}

	if v, ok := d.GetOk("description"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.Description = x

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

	if v, ok := d.GetOk("tags"); ok {
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x

	}

	if v, ok := d.GetOk("upload_location"); ok {
		x := (v.(string))
		o.UploadLocation = x

	}

	if v, ok := d.GetOk("validator"); ok {
		x := (v.(string))
		o.Validator = x

	}

	if v, ok := d.GetOk("validator_input"); ok {
		x := make([]*models.CvdInputMeta, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.CvdInputMeta{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.CvdInputMetaAO0P0.CvdInputMetaAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["field_filter"]; ok {
					{
						x := (v.(string))
						o.FieldFilter = x
					}
				}
				if v, ok := l["field_label"]; ok {
					{
						x := (v.(string))
						o.FieldLabel = x
					}
				}
				if v, ok := l["field_name"]; ok {
					{
						x := (v.(string))
						o.FieldName = x
					}
				}
				if v, ok := l["field_value"]; ok {
					{
						x := (v.(string))
						o.FieldValue = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				x = append(x, &o)
			}
		}
		o.ValidatorInput = x

	}

	if v, ok := d.GetOk("version"); ok {
		x := int64(v.(int))
		o.Version = x

	}

	url := "cvd/Templates"
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceCvdTemplateRead(d, meta)
}

func resourceCvdTemplateRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "cvd/Templates" + "/" + d.Id()

	body, err := conn.SendGetRequest(url, []byte(""))
	if err != nil {
		return err
	}
	var s models.CvdTemplate
	err = s.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model for read Error: %s", err.Error())
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

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.Moid)
	return nil
}
func resourceCvdTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o models.CvdTemplate
	if d.HasChange("deployer") {
		v := d.Get("deployer")
		x := (v.(string))
		o.Deployer = x
	}

	if d.HasChange("deployer_input") {
		v := d.Get("deployer_input")
		x := make([]*models.CvdInputMeta, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.CvdInputMeta{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.CvdInputMetaAO0P0.CvdInputMetaAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["field_filter"]; ok {
					{
						x := (v.(string))
						o.FieldFilter = x
					}
				}
				if v, ok := l["field_label"]; ok {
					{
						x := (v.(string))
						o.FieldLabel = x
					}
				}
				if v, ok := l["field_name"]; ok {
					{
						x := (v.(string))
						o.FieldName = x
					}
				}
				if v, ok := l["field_value"]; ok {
					{
						x := (v.(string))
						o.FieldValue = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				x = append(x, &o)
			}
		}
		o.DeployerInput = x
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.Description = x
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.Moid = x
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.Name = x
	}

	if d.HasChange("object_type") {
		v := d.Get("object_type")
		x := (v.(string))
		o.ObjectType = x
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]*models.MoTag, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.MoTag{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.MoTagAO0P0.MoTagAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["key"]; ok {
					{
						x := (v.(string))
						o.Key = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				if v, ok := l["value"]; ok {
					{
						x := (v.(string))
						o.Value = x
					}
				}
				x = append(x, &o)
			}
		}
		o.Tags = x
	}

	if d.HasChange("upload_location") {
		v := d.Get("upload_location")
		x := (v.(string))
		o.UploadLocation = x
	}

	if d.HasChange("validator") {
		v := d.Get("validator")
		x := (v.(string))
		o.Validator = x
	}

	if d.HasChange("validator_input") {
		v := d.Get("validator_input")
		x := make([]*models.CvdInputMeta, 0)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				o := models.CvdInputMeta{}
				l := s.Index(i).Interface().(map[string]interface{})
				if v, ok := l["additional_properties"]; ok {
					{
						x := []byte(v.(string))
						var x1 interface{}
						err := json.Unmarshal(x, &x1)
						if err == nil && x1 != nil {
							o.CvdInputMetaAO0P0.CvdInputMetaAO0P0 = x1.(map[string]interface{})
						}
					}
				}
				if v, ok := l["field_filter"]; ok {
					{
						x := (v.(string))
						o.FieldFilter = x
					}
				}
				if v, ok := l["field_label"]; ok {
					{
						x := (v.(string))
						o.FieldLabel = x
					}
				}
				if v, ok := l["field_name"]; ok {
					{
						x := (v.(string))
						o.FieldName = x
					}
				}
				if v, ok := l["field_value"]; ok {
					{
						x := (v.(string))
						o.FieldValue = x
					}
				}
				if v, ok := l["object_type"]; ok {
					{
						x := (v.(string))
						o.ObjectType = x
					}
				}
				x = append(x, &o)
			}
		}
		o.ValidatorInput = x
	}

	if d.HasChange("version") {
		v := d.Get("version")
		x := int64(v.(int))
		o.Version = x
	}

	url := "cvd/Templates" + "/" + d.Id()
	data, err := o.MarshalJSON()
	if err != nil {
		log.Printf("error in marshaling model object. Error: %s", err.Error())
		return err
	}

	body, err := conn.SendUpdateRequest(url, data)
	if err != nil {
		return err
	}

	err = o.UnmarshalJSON(body)
	if err != nil {
		log.Printf("error in unmarshaling model object. Error: %s", err.Error())
		return err
	}
	log.Printf("Moid: %s", o.Moid)
	d.SetId(o.Moid)
	return resourceCvdTemplateRead(d, meta)
}

func resourceCvdTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	url := "cvd/Templates" + "/" + d.Id()
	_, err := conn.SendDeleteRequest(url)
	if err != nil {
		log.Printf("error occurred while deleting: %s", err.Error())
	}
	return err
}
