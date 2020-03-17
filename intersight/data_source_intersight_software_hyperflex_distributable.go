package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSoftwareHyperflexDistributable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSoftwareHyperflexDistributableRead,
		Schema: map[string]*schema.Schema{
			"bundle_type": {
				Description: "The bundle type of the image, as published on cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"catalog": {
				Description: "The catalog where this image is present.",
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"description": {
				Description: "User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_count": {
				Description: "The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"guid": {
				Description: "The unique identifier for an image in a Cisco repository.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"import_action": {
				Description: "The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"import_state": {
				Description: "The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"md5sum": {
				Description: "The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mdfid": {
				Description: "The mdfid of the image provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "The endpoint model for which this firmware image is applicable.",
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
			"name": {
				Description: "The name of the file. It is populated as part of the image import operation.",
				Type:        schema.TypeString,
				Optional:    true,
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
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"platform_type": {
				Description: "The platform type of the image.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"recommended_build": {
				Description: "The build which is recommended by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"release_notes_url": {
				Description: "The url for the release notes of this image.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sha512sum": {
				Description: "The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"size": {
				Description: "The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"software_advisory_url": {
				Description: "The software advisory, if any, provided by the vendor for this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_type_id": {
				Description: "The software type id provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"source": {
				Description: "Location of the file in an external repository.",
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"supported_models": {
				Description: "The server models for which this image is applicable.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
			"vendor": {
				Description: "The vendor or publisher of this file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version": {
				Description: "Vendor provided version for the file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}
func dataSourceSoftwareHyperflexDistributableRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "software/HyperflexDistributables"
	var o models.SoftwareHyperflexDistributable
	if v, ok := d.GetOk("bundle_type"); ok {
		x := (v.(string))
		o.BundleType = x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("download_count"); ok {
		x := int64(v.(int))
		o.DownloadCount = x
	}
	if v, ok := d.GetOk("guid"); ok {
		x := (v.(string))
		o.GUID = x
	}
	if v, ok := d.GetOk("import_action"); ok {
		x := (v.(string))
		o.ImportAction = &x
	}
	if v, ok := d.GetOk("import_state"); ok {
		x := (v.(string))
		o.ImportState = x
	}
	if v, ok := d.GetOk("md5sum"); ok {
		x := (v.(string))
		o.Md5sum = x
	}
	if v, ok := d.GetOk("mdfid"); ok {
		x := (v.(string))
		o.Mdfid = x
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.Model = x
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
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.PlatformType = x
	}
	if v, ok := d.GetOk("recommended_build"); ok {
		x := (v.(string))
		o.RecommendedBuild = x
	}
	if v, ok := d.GetOk("release_notes_url"); ok {
		x := (v.(string))
		o.ReleaseNotesURL = x
	}
	if v, ok := d.GetOk("sha512sum"); ok {
		x := (v.(string))
		o.Sha512sum = x
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.Size = x
	}
	if v, ok := d.GetOk("software_advisory_url"); ok {
		x := (v.(string))
		o.SoftwareAdvisoryURL = x
	}
	if v, ok := d.GetOk("software_type_id"); ok {
		x := (v.(string))
		o.SoftwareTypeID = x
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.Vendor = &x
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
			var s models.SoftwareHyperflexDistributable
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("bundle_type", (s.BundleType)); err != nil {
				return err
			}

			if err := d.Set("catalog", flattenMapSoftwarerepositoryCatalogRef(s.Catalog, d)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("download_count", (s.DownloadCount)); err != nil {
				return err
			}
			if err := d.Set("guid", (s.GUID)); err != nil {
				return err
			}
			if err := d.Set("import_action", (s.ImportAction)); err != nil {
				return err
			}
			if err := d.Set("import_state", (s.ImportState)); err != nil {
				return err
			}
			if err := d.Set("md5sum", (s.Md5sum)); err != nil {
				return err
			}
			if err := d.Set("mdfid", (s.Mdfid)); err != nil {
				return err
			}
			if err := d.Set("model", (s.Model)); err != nil {
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

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("platform_type", (s.PlatformType)); err != nil {
				return err
			}
			if err := d.Set("recommended_build", (s.RecommendedBuild)); err != nil {
				return err
			}
			if err := d.Set("release_notes_url", (s.ReleaseNotesURL)); err != nil {
				return err
			}
			if err := d.Set("sha512sum", (s.Sha512sum)); err != nil {
				return err
			}
			if err := d.Set("size", (s.Size)); err != nil {
				return err
			}
			if err := d.Set("software_advisory_url", (s.SoftwareAdvisoryURL)); err != nil {
				return err
			}
			if err := d.Set("software_type_id", (s.SoftwareTypeID)); err != nil {
				return err
			}

			if err := d.Set("source", flattenMapSoftwarerepositoryFileServer(s.Source, d)); err != nil {
				return err
			}
			if err := d.Set("supported_models", (s.SupportedModels)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("vendor", (s.Vendor)); err != nil {
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
