package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceApplianceBackupPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceBackupPolicyRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "BackupPolicy managed object to Account relationship.",
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
							Description: "The Object Type of the referenced REST resource.",
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
			"backup_time": {
				Description: "The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field.",
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
			"manual_backup": {
				Description: "Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Description: "Password to authenticate the file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.These resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.All logical and physical resources part of an organization will have organization in PermissionResources field.If DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects willhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.All profiles/policies created with in an organization will have the organization as PermissionResources.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
			"schedule": {
				Description: "Schedule to create a backup of the Intersight Appliance. Manualbackup field must be set to 'false' for this schedule to be active.",
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
						"day_of_month": {
							Description: "Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"day_of_week": {
							Description: "Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"month_of_year": {
							Description: "Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"repeat_interval": {
							Description: "Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"time_of_day": {
							Description: "Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"time_zone": {
							Description: "Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"week_of_month": {
							Description: "Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
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
func dataSourceApplianceBackupPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "appliance/BackupPolicies"
	var o models.ApplianceBackupPolicy
	if v, ok := d.GetOk("backup_time"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.BackupTime = x
	}
	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.Filename = x
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.IsPasswordSet = &x
	}
	if v, ok := d.GetOk("manual_backup"); ok {
		x := (v.(bool))
		o.ManualBackup = &x
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
			var s models.ApplianceBackupPolicy
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("account", flattenMapIamAccountRef(s.Account, d)); err != nil {
				return err
			}

			if err := d.Set("backup_time", (s.BackupTime).String()); err != nil {
				return err
			}
			if err := d.Set("filename", (s.Filename)); err != nil {
				return err
			}
			if err := d.Set("is_password_set", (s.IsPasswordSet)); err != nil {
				return err
			}
			if err := d.Set("manual_backup", (s.ManualBackup)); err != nil {
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

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
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

			if err := d.Set("schedule", flattenMapOnpremSchedule(s.Schedule, d)); err != nil {
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
