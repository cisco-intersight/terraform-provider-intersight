package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vvb/terraform-provider-intersight/models"
)

func dataSourceApplianceBackup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceBackupRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "Backup managed object to Account relationship.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed: true,
			},
			"elapsed_time": {
				Description: "Elapsed time in seconds since the backup process has started.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "End date and time of the backup process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"filename": {
				Description: "Backup filename to backup or restore.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"messages": {
				Description: "Messages generated during the backup process.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
			"password": {
				Description: "Password to authenticate the fileserver.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"protocol": {
				Description: "Communication protocol used by the file server (e.g. scp or sftp).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_host": {
				Description: "Hostname of the remote file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_path": {
				Description: "File server directory to copy the file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_port": {
				Description: "Remote TCP port on the file server (e.g. 22 for scp).",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"start_time": {
				Description: "Start date and time of the backup process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the backup managed object.",
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
			"username": {
				Description: "Username to authenticate the fileserver.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceApplianceBackupRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "appliance/Backups"
	var o models.ApplianceBackup
	if v, ok := d.GetOk("elapsed_time"); ok {
		x := int64(v.(int))
		o.ElapsedTime = x
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.EndTime = x
	}
	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.Filename = x
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.IsPasswordSet = &x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.Password = x
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.Protocol = &x
	}
	if v, ok := d.GetOk("remote_host"); ok {
		x := (v.(string))
		o.RemoteHost = x
	}
	if v, ok := d.GetOk("remote_path"); ok {
		x := (v.(string))
		o.RemotePath = x
	}
	if v, ok := d.GetOk("remote_port"); ok {
		x := int64(v.(int))
		o.RemotePort = x
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.StartTime = x
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = x
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.Username = x
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
			var s models.ApplianceBackup
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
				return err
			}
			if err := d.Set("elapsed_time", (s.ElapsedTime)); err != nil {
				return err
			}

			if err := d.Set("end_time", (s.EndTime).String()); err != nil {
				return err
			}
			if err := d.Set("filename", (s.Filename)); err != nil {
				return err
			}
			if err := d.Set("is_password_set", (s.IsPasswordSet)); err != nil {
				return err
			}
			if err := d.Set("messages", (s.Messages)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("password", (s.Password)); err != nil {
				return err
			}
			if err := d.Set("protocol", (s.Protocol)); err != nil {
				return err
			}
			if err := d.Set("remote_host", (s.RemoteHost)); err != nil {
				return err
			}
			if err := d.Set("remote_path", (s.RemotePath)); err != nil {
				return err
			}
			if err := d.Set("remote_port", (s.RemotePort)); err != nil {
				return err
			}

			if err := d.Set("start_time", (s.StartTime).String()); err != nil {
				return err
			}
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("username", (s.Username)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
