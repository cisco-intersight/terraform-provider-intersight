package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceTamSecurityAdvisory() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTamSecurityAdvisoryRead,
		Schema: map[string]*schema.Schema{
			"api_data_sources": {
				Description: "An array of data sources that are used to provide data for queries used to identify an Intersight alert applicability.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"mo_type": {
							Description: "Type of Intersight managed object used as data source.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Name is used to unique identify and refer a given data source in an alert definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"queries": {
							Description: "Optional set of Queries to filter the output for Api datasource. the queries are executed in the order specified.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Description: "An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"query": {
										Description: "A SparkSQL query to be used on a given data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed: true,
						},
						"type": {
							Description: "Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"actions": {
				Description: "An array of actions that are to be taken when a given managed object matches the criteria specified for being affected by an alert definition.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"affected_object_type": {
							Description: "Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"alert_type": {
							Description: "Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"identifiers": {
							Description: "Identifiers represents the filter criteria (property names and values) used to identify an Intersight managed object of type specified in affectedObjectType property. An instance of an alert is then create on (or removed from) the identified managed object.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"name": {
										Description: "Name of the filter paramter.",
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
										Description: "Value of the filter paramter.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed: true,
						},
						"name": {
							Description: "Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"operation_type": {
							Description: "Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"queries": {
							Description: "Set of SparkSQL queries used determine if a given alert is applicable or not. Refer to https://spark.apache.org/sql/ for more details.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Description: "An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"query": {
										Description: "A SparkSQL query to be used on a given data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed: true,
						},
						"type": {
							Description: "Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"advisory_id": {
				Description: "Cisco generated identifier for the published security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"base_score": {
				Description: "CVSS version 3 base score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"cve_ids": {
				Description: "CVE (https://cve.mitre.org/about/faqs.html) identifiers associated with the published security Advisory.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"date_published": {
				Description: "Date when the security advisory was first published by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"date_updated": {
				Description: "Date when the security advisory was last updated by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Brief description of the advisory details.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environmental_score": {
				Description: "CVSS version 3 environmental score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"external_url": {
				Description: "A link to an external URL describing security Advisory in more details.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "A user defined name for the Intersight Advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "Relationship to the Organization that owns the Managed Object.",
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
			"recommendation": {
				Description: "Recommended action to resolve the security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"severity": {
				Description: "Severity level of the Intersight Advisory.",
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
			"state": {
				Description: "Current state of the advisory. Indicates if the user is interested in getting updates for the advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "Cisco assigned status of the published advisory based on whether the investigation is complete or on-going.",
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
			"temporal_score": {
				Description: "CVSS version 3 temporal score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"version": {
				Description: "Cisco assigned advisory version after latest revision.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceTamSecurityAdvisoryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "tam/SecurityAdvisories"
	var o models.TamSecurityAdvisory
	if v, ok := d.GetOk("advisory_id"); ok {
		x := (v.(string))
		o.AdvisoryID = x
	}
	if v, ok := d.GetOk("base_score"); ok {
		x := v.(float64)
		o.BaseScore = x
	}
	if v, ok := d.GetOk("date_published"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.DatePublished = x
	}
	if v, ok := d.GetOk("date_updated"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.DateUpdated = x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("environmental_score"); ok {
		x := v.(float64)
		o.EnvironmentalScore = x
	}
	if v, ok := d.GetOk("external_url"); ok {
		x := (v.(string))
		o.ExternalURL = x
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
	if v, ok := d.GetOk("recommendation"); ok {
		x := (v.(string))
		o.Recommendation = x
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.State = &x
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.Status = &x
	}
	if v, ok := d.GetOk("temporal_score"); ok {
		x := v.(float64)
		o.TemporalScore = x
	}
	if v, ok := d.GetOk("version"); ok {
		x := (v.(string))
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
			var s models.TamSecurityAdvisory
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("api_data_sources", flattenListTamAPIDataSource(s.APIDataSources, d)); err != nil {
				return err
			}

			if err := d.Set("actions", flattenListTamAction(s.Actions, d)); err != nil {
				return err
			}
			if err := d.Set("advisory_id", (s.AdvisoryID)); err != nil {
				return err
			}
			if err := d.Set("base_score", (s.BaseScore)); err != nil {
				return err
			}
			if err := d.Set("cve_ids", (s.CveIds)); err != nil {
				return err
			}

			if err := d.Set("date_published", (s.DatePublished).String()); err != nil {
				return err
			}

			if err := d.Set("date_updated", (s.DateUpdated).String()); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("environmental_score", (s.EnvironmentalScore)); err != nil {
				return err
			}
			if err := d.Set("external_url", (s.ExternalURL)); err != nil {
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

			if err := d.Set("organization", flattenMapOrganizationOrganizationRef(s.Organization, d)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("recommendation", (s.Recommendation)); err != nil {
				return err
			}

			if err := d.Set("severity", flattenMapTamSeverity(s.Severity, d)); err != nil {
				return err
			}
			if err := d.Set("state", (s.State)); err != nil {
				return err
			}
			if err := d.Set("status", (s.Status)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("temporal_score", (s.TemporalScore)); err != nil {
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
