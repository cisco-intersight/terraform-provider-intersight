package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAssetDeviceContractInformation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssetDeviceContractInformationRead,
		Schema: map[string]*schema.Schema{
			"contract": {
				Description: "Contract information for the Cisco support contract purchased for the Cisco device.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"bill_to": {
							Description: "BillTo address of listed for the contract.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
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
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"bill_to_global_ultimate": {
							Description: "BillToGlobalUltimate information listed in the contract.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"id": {
										Description: "ID of the user in BillToGlobal.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user in BillToGlobal.",
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
								},
							},
						},
						"contract_number": {
							Description: "Contract number for the Cisco support contract purchased for the Cisco device.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"line_status": {
							Description: "Contract status as per the Cisco Contract APIx.",
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
					},
				},
			},
			"contract_status": {
				Description: "Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"covered_product_line_end_date": {
				Description: "End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Description: "Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_type": {
				Description: "Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"end_customer": {
				Description: "End customer information for the Cisco support contract purchased for the Cisco device.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"address": {
							Description: "Address as per the information provided by the user.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
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
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Description: "Unique identifier for an end customer. This identifier is allocated by Cisco.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name as per the information provided by the user.",
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
					},
				},
			},
			"end_user_global_ultimate": {
				Description: "EndUserGlobalUltimate information listed in the contract.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"id": {
							Description: "ID of the user in BillToGlobal.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the user in BillToGlobal.",
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
					},
				},
			},
			"is_valid": {
				Description: "Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"item_type": {
				Description: "Item type of this specific Cisco device. example \"Chassis\".",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"maintenance_purchase_order_number": {
				Description: "Maintenance purchase order number for the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"maintenance_sales_order_number": {
				Description: "Maintenance sales order number for the Cisco device.",
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
			"object_type": {
				Description: "The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"platform_type": {
				Description: "The platform type of the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"product": {
				Description: "Product information of the offering under a contract.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"bill_to": {
							Description: "Billing address provided by customer while buying this Cisco product.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
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
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"description": {
							Description: "Short description of the Cisco product that helps identify the product easily. example \"DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"family": {
							Description: "Family that the product belongs to. Example \"UCSB\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"group": {
							Description: "Group that the product belongs to. It is one higher level categorization above family. example \"Switch\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"number": {
							Description: "Product number that identifies the product. example PID. example \"UCS-FI-6248UP-CH2\".",
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
						"ship_to": {
							Description: "Shipping address provided by customer while buying this Cisco product.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"address1": {
										Description: "Address Line one of the address information. example \"PO BOX 641570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"address2": {
										Description: "Address Line two of the address information. example \"Cisco Systems\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"city": {
										Description: "City in which the address resides. example \"San Jose\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Description: "Country in which the address resides. example \"US\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"location": {
										Description: "Location in which the address resides. example \"14852\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the user whose address is being populated.",
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
									"postal_code": {
										Description: "Postal Code in which the address resides. example \"95164-1570\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"state": {
										Description: "State in which the address resides. example \"CA\".",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"sub_type": {
							Description: "Sub type of the product being specified. example \"UCS 6200 SER\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"purchase_order_number": {
				Description: "Purchase order number for the Cisco device. It is a unique number assigned for every purchase.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "Reference to the device connector through which the device is connected.",
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
			"reseller_global_ultimate": {
				Description: "ResellerGlobalUltimate information listed in the contract.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"id": {
							Description: "ID of the user in BillToGlobal.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the user in BillToGlobal.",
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
					},
				},
			},
			"sales_order_number": {
				Description: "Sales order number for the Cisco device. It is a unique number assigned for every sale.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_description": {
				Description: "The type of service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_end_date": {
				Description: "End date for the Cisco service contract that covers this Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_level": {
				Description: "The type of service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_sku": {
				Description: "The SKU of the service contract that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_start_date": {
				Description: "Start date for the Cisco service contract that covers this Cisco device.",
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
			"warranty_end_date": {
				Description: "End date for the warranty that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"warranty_type": {
				Description: "Type of warranty that covers the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}
func dataSourceAssetDeviceContractInformationRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "asset/DeviceContractInformations"
	var o models.AssetDeviceContractInformation
	if v, ok := d.GetOk("contract_status"); ok {
		x := (v.(string))
		o.ContractStatus = x
	}
	if v, ok := d.GetOk("covered_product_line_end_date"); ok {
		x := (v.(string))
		o.CoveredProductLineEndDate = x
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.DeviceID = x
	}
	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.DeviceType = x
	}
	if v, ok := d.GetOk("is_valid"); ok {
		x := (v.(bool))
		o.IsValid = &x
	}
	if v, ok := d.GetOk("item_type"); ok {
		x := (v.(string))
		o.ItemType = x
	}
	if v, ok := d.GetOk("maintenance_purchase_order_number"); ok {
		x := (v.(string))
		o.MaintenancePurchaseOrderNumber = x
	}
	if v, ok := d.GetOk("maintenance_sales_order_number"); ok {
		x := (v.(string))
		o.MaintenanceSalesOrderNumber = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.PlatformType = x
	}
	if v, ok := d.GetOk("purchase_order_number"); ok {
		x := (v.(string))
		o.PurchaseOrderNumber = x
	}
	if v, ok := d.GetOk("sales_order_number"); ok {
		x := (v.(string))
		o.SalesOrderNumber = x
	}
	if v, ok := d.GetOk("service_description"); ok {
		x := (v.(string))
		o.ServiceDescription = x
	}
	if v, ok := d.GetOk("service_end_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ServiceEndDate = x
	}
	if v, ok := d.GetOk("service_level"); ok {
		x := (v.(string))
		o.ServiceLevel = x
	}
	if v, ok := d.GetOk("service_sku"); ok {
		x := (v.(string))
		o.ServiceSku = x
	}
	if v, ok := d.GetOk("service_start_date"); ok {
		x, _ := strfmt.ParseDateTime(v.(string))
		o.ServiceStartDate = x
	}
	if v, ok := d.GetOk("warranty_end_date"); ok {
		x := (v.(string))
		o.WarrantyEndDate = x
	}
	if v, ok := d.GetOk("warranty_type"); ok {
		x := (v.(string))
		o.WarrantyType = x
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
			var s models.AssetDeviceContractInformation
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}

			if err := d.Set("contract", flattenMapAssetContractInformation(s.Contract, d)); err != nil {
				return err
			}
			if err := d.Set("contract_status", (s.ContractStatus)); err != nil {
				return err
			}
			if err := d.Set("covered_product_line_end_date", (s.CoveredProductLineEndDate)); err != nil {
				return err
			}
			if err := d.Set("device_id", (s.DeviceID)); err != nil {
				return err
			}
			if err := d.Set("device_type", (s.DeviceType)); err != nil {
				return err
			}

			if err := d.Set("end_customer", flattenMapAssetCustomerInformation(s.EndCustomer, d)); err != nil {
				return err
			}

			if err := d.Set("end_user_global_ultimate", flattenMapAssetGlobalUltimate(s.EndUserGlobalUltimate, d)); err != nil {
				return err
			}
			if err := d.Set("is_valid", (s.IsValid)); err != nil {
				return err
			}
			if err := d.Set("item_type", (s.ItemType)); err != nil {
				return err
			}
			if err := d.Set("maintenance_purchase_order_number", (s.MaintenancePurchaseOrderNumber)); err != nil {
				return err
			}
			if err := d.Set("maintenance_sales_order_number", (s.MaintenanceSalesOrderNumber)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("platform_type", (s.PlatformType)); err != nil {
				return err
			}

			if err := d.Set("product", flattenMapAssetProductInformation(s.Product, d)); err != nil {
				return err
			}
			if err := d.Set("purchase_order_number", (s.PurchaseOrderNumber)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}

			if err := d.Set("reseller_global_ultimate", flattenMapAssetGlobalUltimate(s.ResellerGlobalUltimate, d)); err != nil {
				return err
			}
			if err := d.Set("sales_order_number", (s.SalesOrderNumber)); err != nil {
				return err
			}
			if err := d.Set("service_description", (s.ServiceDescription)); err != nil {
				return err
			}

			if err := d.Set("service_end_date", (s.ServiceEndDate).String()); err != nil {
				return err
			}
			if err := d.Set("service_level", (s.ServiceLevel)); err != nil {
				return err
			}
			if err := d.Set("service_sku", (s.ServiceSku)); err != nil {
				return err
			}

			if err := d.Set("service_start_date", (s.ServiceStartDate).String()); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("warranty_end_date", (s.WarrantyEndDate)); err != nil {
				return err
			}
			if err := d.Set("warranty_type", (s.WarrantyType)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
