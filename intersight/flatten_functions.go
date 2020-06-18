package intersight

import (
	"encoding/json"
	"log"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenListAdapterAdapterConfig(p []*models.AdapterAdapterConfig, d *schema.ResourceData) []map[string]interface{} {
	var adapteradapterconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		adapteradapterconfig := make(map[string]interface{})
		if len(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1) != 0 {
			j, err := json.Marshal(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1)
			if err == nil {
				adapteradapterconfig["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		adapteradapterconfig["class_id"] = item.ClassID
		adapteradapterconfig["dce_interface_settings"] = (func(p []*models.AdapterDceInterfaceSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterdceinterfacesettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				adapterdceinterfacesettings := make(map[string]interface{})
				if len(item.AdapterDceInterfaceSettingsAO1P1.AdapterDceInterfaceSettingsAO1P1) != 0 {
					j, err := json.Marshal(item.AdapterDceInterfaceSettingsAO1P1.AdapterDceInterfaceSettingsAO1P1)
					if err == nil {
						adapterdceinterfacesettings["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				adapterdceinterfacesettings["class_id"] = item.ClassID
				adapterdceinterfacesettings["fec_mode"] = item.FecMode
				adapterdceinterfacesettings["interface_id"] = item.InterfaceID
				adapterdceinterfacesettings["object_type"] = item.ObjectType
				adapterdceinterfacesettingss = append(adapterdceinterfacesettingss, adapterdceinterfacesettings)
			}
			return adapterdceinterfacesettingss
		})(item.DceInterfaceSettings, d)
		adapteradapterconfig["eth_settings"] = (func(p *models.AdapterEthSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterethsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			adapterethsettings := make(map[string]interface{})
			delete(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1, "ObjectType")
			if len(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1)
				if err == nil {
					adapterethsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			adapterethsettings["class_id"] = item.ClassID
			adapterethsettings["lldp_enabled"] = item.LldpEnabled
			adapterethsettings["object_type"] = item.ObjectType

			adapterethsettingss = append(adapterethsettingss, adapterethsettings)
			return adapterethsettingss
		})(item.EthSettings, d)
		adapteradapterconfig["fc_settings"] = (func(p *models.AdapterFcSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterfcsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			adapterfcsettings := make(map[string]interface{})
			delete(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1, "ObjectType")
			if len(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1)
				if err == nil {
					adapterfcsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			adapterfcsettings["class_id"] = item.ClassID
			adapterfcsettings["fip_enabled"] = item.FipEnabled
			adapterfcsettings["object_type"] = item.ObjectType

			adapterfcsettingss = append(adapterfcsettingss, adapterfcsettings)
			return adapterfcsettingss
		})(item.FcSettings, d)
		adapteradapterconfig["object_type"] = item.ObjectType
		adapteradapterconfig["port_channel_settings"] = (func(p *models.AdapterPortChannelSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterportchannelsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			adapterportchannelsettings := make(map[string]interface{})
			delete(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1, "ObjectType")
			if len(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1)
				if err == nil {
					adapterportchannelsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			adapterportchannelsettings["class_id"] = item.ClassID
			adapterportchannelsettings["enabled"] = item.Enabled
			adapterportchannelsettings["object_type"] = item.ObjectType

			adapterportchannelsettingss = append(adapterportchannelsettingss, adapterportchannelsettings)
			return adapterportchannelsettingss
		})(item.PortChannelSettings, d)
		adapteradapterconfig["slot_id"] = item.SlotID
		adapteradapterconfigs = append(adapteradapterconfigs, adapteradapterconfig)
	}
	return adapteradapterconfigs
}
func flattenListAdapterExtEthInterfaceRef(p []*models.AdapterExtEthInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var adapterextethinterfacerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		adapterextethinterfaceref := make(map[string]interface{})
		adapterextethinterfaceref["moid"] = item.Moid
		adapterextethinterfaceref["object_type"] = item.ObjectType
		adapterextethinterfaceref["selector"] = item.Selector
		adapterextethinterfacerefs = append(adapterextethinterfacerefs, adapterextethinterfaceref)
	}
	return adapterextethinterfacerefs
}
func flattenListAdapterHostEthInterfaceRef(p []*models.AdapterHostEthInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostethinterfacerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		adapterhostethinterfaceref := make(map[string]interface{})
		adapterhostethinterfaceref["moid"] = item.Moid
		adapterhostethinterfaceref["object_type"] = item.ObjectType
		adapterhostethinterfaceref["selector"] = item.Selector
		adapterhostethinterfacerefs = append(adapterhostethinterfacerefs, adapterhostethinterfaceref)
	}
	return adapterhostethinterfacerefs
}
func flattenListAdapterHostFcInterfaceRef(p []*models.AdapterHostFcInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostfcinterfacerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		adapterhostfcinterfaceref := make(map[string]interface{})
		adapterhostfcinterfaceref["moid"] = item.Moid
		adapterhostfcinterfaceref["object_type"] = item.ObjectType
		adapterhostfcinterfaceref["selector"] = item.Selector
		adapterhostfcinterfacerefs = append(adapterhostfcinterfacerefs, adapterhostfcinterfaceref)
	}
	return adapterhostfcinterfacerefs
}
func flattenListAdapterHostIscsiInterfaceRef(p []*models.AdapterHostIscsiInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostiscsiinterfacerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		adapterhostiscsiinterfaceref := make(map[string]interface{})
		adapterhostiscsiinterfaceref["moid"] = item.Moid
		adapterhostiscsiinterfaceref["object_type"] = item.ObjectType
		adapterhostiscsiinterfaceref["selector"] = item.Selector
		adapterhostiscsiinterfacerefs = append(adapterhostiscsiinterfacerefs, adapterhostiscsiinterfaceref)
	}
	return adapterhostiscsiinterfacerefs
}
func flattenListAdapterUnitRef(p []*models.AdapterUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		adapterunitref := make(map[string]interface{})
		adapterunitref["moid"] = item.Moid
		adapterunitref["object_type"] = item.ObjectType
		adapterunitref["selector"] = item.Selector
		adapterunitrefs = append(adapterunitrefs, adapterunitref)
	}
	return adapterunitrefs
}
func flattenListApplianceDataExportPolicyRef(p []*models.ApplianceDataExportPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		appliancedataexportpolicyref := make(map[string]interface{})
		appliancedataexportpolicyref["moid"] = item.Moid
		appliancedataexportpolicyref["object_type"] = item.ObjectType
		appliancedataexportpolicyref["selector"] = item.Selector
		appliancedataexportpolicyrefs = append(appliancedataexportpolicyrefs, appliancedataexportpolicyref)
	}
	return appliancedataexportpolicyrefs
}
func flattenListApplianceKeyValuePair(p []*models.ApplianceKeyValuePair, d *schema.ResourceData) []map[string]interface{} {
	var appliancekeyvaluepairs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		appliancekeyvaluepair := make(map[string]interface{})
		if len(item.ApplianceKeyValuePairAO1P1.ApplianceKeyValuePairAO1P1) != 0 {
			j, err := json.Marshal(item.ApplianceKeyValuePairAO1P1.ApplianceKeyValuePairAO1P1)
			if err == nil {
				appliancekeyvaluepair["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		appliancekeyvaluepair["class_id"] = item.ClassID
		appliancekeyvaluepair["key"] = item.Key
		appliancekeyvaluepair["object_type"] = item.ObjectType
		appliancekeyvaluepair["value"] = item.Value
		appliancekeyvaluepairs = append(appliancekeyvaluepairs, appliancekeyvaluepair)
	}
	return appliancekeyvaluepairs
}
func flattenListAssetClusterMemberRef(p []*models.AssetClusterMemberRef, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		assetclustermemberref := make(map[string]interface{})
		assetclustermemberref["moid"] = item.Moid
		assetclustermemberref["object_type"] = item.ObjectType
		assetclustermemberref["selector"] = item.Selector
		assetclustermemberrefs = append(assetclustermemberrefs, assetclustermemberref)
	}
	return assetclustermemberrefs
}
func flattenListAssetConnection(p []*models.AssetConnection, d *schema.ResourceData) []map[string]interface{} {
	var assetconnections []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		assetconnection := make(map[string]interface{})
		if len(item.AO1) != 0 {
			j, err := json.Marshal(item.AO1)
			if err == nil {
				assetconnection["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		assetconnection["class_id"] = item.ClassID
		assetconnection["object_type"] = item.ObjectType
		assetconnections = append(assetconnections, assetconnection)
	}
	return assetconnections
}
func flattenListAssetService(p []*models.AssetService, d *schema.ResourceData) []map[string]interface{} {
	var assetservices []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		assetservice := make(map[string]interface{})
		if len(item.AssetServiceAO1P1.AssetServiceAO1P1) != 0 {
			j, err := json.Marshal(item.AssetServiceAO1P1.AssetServiceAO1P1)
			if err == nil {
				assetservice["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		assetservice["class_id"] = item.ClassID
		assetservice["object_type"] = item.ObjectType
		assetservice["options"] = (func(p *models.AssetServiceOptions, d *schema.ResourceData) []map[string]interface{} {
			var assetserviceoptionss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			assetserviceoptions := make(map[string]interface{})
			delete(item.AO1, "ObjectType")
			if len(item.AO1) != 0 {
				j, err := json.Marshal(item.AO1)
				if err == nil {
					assetserviceoptions["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			assetserviceoptions["class_id"] = item.ClassID
			assetserviceoptions["object_type"] = item.ObjectType

			assetserviceoptionss = append(assetserviceoptionss, assetserviceoptions)
			return assetserviceoptionss
		})(item.Options, d)
		assetservice["status"] = item.Status
		assetservice["status_error_reason"] = item.StatusErrorReason
		assetservices = append(assetservices, assetservice)
	}
	return assetservices
}
func flattenListBiosUnitRef(p []*models.BiosUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		biosunitref := make(map[string]interface{})
		biosunitref["moid"] = item.Moid
		biosunitref["object_type"] = item.ObjectType
		biosunitref["selector"] = item.Selector
		biosunitrefs = append(biosunitrefs, biosunitref)
	}
	return biosunitrefs
}
func flattenListBootDeviceBase(p []*models.BootDeviceBase, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebases []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		bootdevicebase := make(map[string]interface{})
		if len(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1) != 0 {
			j, err := json.Marshal(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1)
			if err == nil {
				bootdevicebase["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		bootdevicebase["class_id"] = item.ClassID
		bootdevicebase["enabled"] = item.Enabled
		bootdevicebase["name"] = item.Name
		bootdevicebase["object_type"] = item.ObjectType
		bootdevicebases = append(bootdevicebases, bootdevicebase)
	}
	return bootdevicebases
}
func flattenListComputeBladeRef(p []*models.ComputeBladeRef, d *schema.ResourceData) []map[string]interface{} {
	var computebladerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		computebladeref := make(map[string]interface{})
		computebladeref["moid"] = item.Moid
		computebladeref["object_type"] = item.ObjectType
		computebladeref["selector"] = item.Selector
		computebladerefs = append(computebladerefs, computebladeref)
	}
	return computebladerefs
}
func flattenListComputeIPAddress(p []*models.ComputeIPAddress, d *schema.ResourceData) []map[string]interface{} {
	var computeipaddresss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		computeipaddress := make(map[string]interface{})
		if len(item.ComputeIPAddressAO1P1.ComputeIPAddressAO1P1) != 0 {
			j, err := json.Marshal(item.ComputeIPAddressAO1P1.ComputeIPAddressAO1P1)
			if err == nil {
				computeipaddress["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		computeipaddress["address"] = item.Address
		computeipaddress["category"] = item.Category
		computeipaddress["class_id"] = item.ClassID
		computeipaddress["default_gateway"] = item.DefaultGateway
		computeipaddress["dn"] = item.Dn
		computeipaddress["http_port"] = item.HTTPPort
		computeipaddress["https_port"] = item.HTTPSPort
		computeipaddress["kvm_port"] = item.KvmPort
		computeipaddress["name"] = item.Name
		computeipaddress["object_type"] = item.ObjectType
		computeipaddress["subnet"] = item.Subnet
		computeipaddress["type"] = item.Type
		computeipaddresss = append(computeipaddresss, computeipaddress)
	}
	return computeipaddresss
}
func flattenListComputeRackUnitRef(p []*models.ComputeRackUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		computerackunitref := make(map[string]interface{})
		computerackunitref["moid"] = item.Moid
		computerackunitref["object_type"] = item.ObjectType
		computerackunitref["selector"] = item.Selector
		computerackunitrefs = append(computerackunitrefs, computerackunitref)
	}
	return computerackunitrefs
}
func flattenListCondHclStatusDetailRef(p []*models.CondHclStatusDetailRef, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusdetailrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		condhclstatusdetailref := make(map[string]interface{})
		condhclstatusdetailref["moid"] = item.Moid
		condhclstatusdetailref["object_type"] = item.ObjectType
		condhclstatusdetailref["selector"] = item.Selector
		condhclstatusdetailrefs = append(condhclstatusdetailrefs, condhclstatusdetailref)
	}
	return condhclstatusdetailrefs
}
func flattenListConfigExportCandidateRef(p []*models.ConfigExportCandidateRef, d *schema.ResourceData) []map[string]interface{} {
	var configexportcandidaterefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		configexportcandidateref := make(map[string]interface{})
		configexportcandidateref["moid"] = item.Moid
		configexportcandidateref["object_type"] = item.ObjectType
		configexportcandidateref["selector"] = item.Selector
		configexportcandidaterefs = append(configexportcandidaterefs, configexportcandidateref)
	}
	return configexportcandidaterefs
}
func flattenListConfigExportedItemRef(p []*models.ConfigExportedItemRef, d *schema.ResourceData) []map[string]interface{} {
	var configexporteditemrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		configexporteditemref := make(map[string]interface{})
		configexporteditemref["moid"] = item.Moid
		configexporteditemref["object_type"] = item.ObjectType
		configexporteditemref["selector"] = item.Selector
		configexporteditemrefs = append(configexporteditemrefs, configexporteditemref)
	}
	return configexporteditemrefs
}
func flattenListConfigImportedItemRef(p []*models.ConfigImportedItemRef, d *schema.ResourceData) []map[string]interface{} {
	var configimporteditemrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		configimporteditemref := make(map[string]interface{})
		configimporteditemref["moid"] = item.Moid
		configimporteditemref["object_type"] = item.ObjectType
		configimporteditemref["selector"] = item.Selector
		configimporteditemrefs = append(configimporteditemrefs, configimporteditemref)
	}
	return configimporteditemrefs
}
func flattenListConfigMoRef(p []*models.ConfigMoRef, d *schema.ResourceData) []map[string]interface{} {
	var configmorefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		configmoref := make(map[string]interface{})
		if len(item.ConfigMoRefAO1P1.ConfigMoRefAO1P1) != 0 {
			j, err := json.Marshal(item.ConfigMoRefAO1P1.ConfigMoRefAO1P1)
			if err == nil {
				configmoref["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		configmoref["class_id"] = item.ClassID
		configmoref["moid"] = item.Moid
		configmoref["object_type"] = item.ObjectType
		configmorefs = append(configmorefs, configmoref)
	}
	return configmorefs
}
func flattenListConnectorpackConnectorPackUpdate(p []*models.ConnectorpackConnectorPackUpdate, d *schema.ResourceData) []map[string]interface{} {
	var connectorpackconnectorpackupdates []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		connectorpackconnectorpackupdate := make(map[string]interface{})
		if len(item.ConnectorpackConnectorPackUpdateAO1P1.ConnectorpackConnectorPackUpdateAO1P1) != 0 {
			j, err := json.Marshal(item.ConnectorpackConnectorPackUpdateAO1P1.ConnectorpackConnectorPackUpdateAO1P1)
			if err == nil {
				connectorpackconnectorpackupdate["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		connectorpackconnectorpackupdate["class_id"] = item.ClassID
		connectorpackconnectorpackupdate["current_version"] = item.CurrentVersion
		connectorpackconnectorpackupdate["name"] = item.Name
		connectorpackconnectorpackupdate["new_version"] = item.NewVersion
		connectorpackconnectorpackupdate["object_type"] = item.ObjectType
		connectorpackconnectorpackupdates = append(connectorpackconnectorpackupdates, connectorpackconnectorpackupdate)
	}
	return connectorpackconnectorpackupdates
}
func flattenListEquipmentFanModuleRef(p []*models.EquipmentFanModuleRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentfanmoduleref := make(map[string]interface{})
		equipmentfanmoduleref["moid"] = item.Moid
		equipmentfanmoduleref["object_type"] = item.ObjectType
		equipmentfanmoduleref["selector"] = item.Selector
		equipmentfanmodulerefs = append(equipmentfanmodulerefs, equipmentfanmoduleref)
	}
	return equipmentfanmodulerefs
}
func flattenListEquipmentFanRef(p []*models.EquipmentFanRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentfanref := make(map[string]interface{})
		equipmentfanref["moid"] = item.Moid
		equipmentfanref["object_type"] = item.ObjectType
		equipmentfanref["selector"] = item.Selector
		equipmentfanrefs = append(equipmentfanrefs, equipmentfanref)
	}
	return equipmentfanrefs
}
func flattenListEquipmentIoCardRef(p []*models.EquipmentIoCardRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentiocardref := make(map[string]interface{})
		equipmentiocardref["moid"] = item.Moid
		equipmentiocardref["object_type"] = item.ObjectType
		equipmentiocardref["selector"] = item.Selector
		equipmentiocardrefs = append(equipmentiocardrefs, equipmentiocardref)
	}
	return equipmentiocardrefs
}
func flattenListEquipmentIoExpanderRef(p []*models.EquipmentIoExpanderRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentioexpanderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentioexpanderref := make(map[string]interface{})
		equipmentioexpanderref["moid"] = item.Moid
		equipmentioexpanderref["object_type"] = item.ObjectType
		equipmentioexpanderref["selector"] = item.Selector
		equipmentioexpanderrefs = append(equipmentioexpanderrefs, equipmentioexpanderref)
	}
	return equipmentioexpanderrefs
}
func flattenListEquipmentPsuRef(p []*models.EquipmentPsuRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentpsurefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentpsuref := make(map[string]interface{})
		equipmentpsuref["moid"] = item.Moid
		equipmentpsuref["object_type"] = item.ObjectType
		equipmentpsuref["selector"] = item.Selector
		equipmentpsurefs = append(equipmentpsurefs, equipmentpsuref)
	}
	return equipmentpsurefs
}
func flattenListEquipmentRackEnclosureSlotRef(p []*models.EquipmentRackEnclosureSlotRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentrackenclosureslotref := make(map[string]interface{})
		equipmentrackenclosureslotref["moid"] = item.Moid
		equipmentrackenclosureslotref["object_type"] = item.ObjectType
		equipmentrackenclosureslotref["selector"] = item.Selector
		equipmentrackenclosureslotrefs = append(equipmentrackenclosureslotrefs, equipmentrackenclosureslotref)
	}
	return equipmentrackenclosureslotrefs
}
func flattenListEquipmentSwitchCardRef(p []*models.EquipmentSwitchCardRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentswitchcardref := make(map[string]interface{})
		equipmentswitchcardref["moid"] = item.Moid
		equipmentswitchcardref["object_type"] = item.ObjectType
		equipmentswitchcardref["selector"] = item.Selector
		equipmentswitchcardrefs = append(equipmentswitchcardrefs, equipmentswitchcardref)
	}
	return equipmentswitchcardrefs
}
func flattenListEquipmentSystemIoControllerRef(p []*models.EquipmentSystemIoControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentsystemiocontrollerref := make(map[string]interface{})
		equipmentsystemiocontrollerref["moid"] = item.Moid
		equipmentsystemiocontrollerref["object_type"] = item.ObjectType
		equipmentsystemiocontrollerref["selector"] = item.Selector
		equipmentsystemiocontrollerrefs = append(equipmentsystemiocontrollerrefs, equipmentsystemiocontrollerref)
	}
	return equipmentsystemiocontrollerrefs
}
func flattenListEquipmentTpmRef(p []*models.EquipmentTpmRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmenttpmrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmenttpmref := make(map[string]interface{})
		equipmenttpmref["moid"] = item.Moid
		equipmenttpmref["object_type"] = item.ObjectType
		equipmenttpmref["selector"] = item.Selector
		equipmenttpmrefs = append(equipmenttpmrefs, equipmenttpmref)
	}
	return equipmenttpmrefs
}
func flattenListEtherPhysicalPortRef(p []*models.EtherPhysicalPortRef, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		etherphysicalportref := make(map[string]interface{})
		etherphysicalportref["moid"] = item.Moid
		etherphysicalportref["object_type"] = item.ObjectType
		etherphysicalportref["selector"] = item.Selector
		etherphysicalportrefs = append(etherphysicalportrefs, etherphysicalportref)
	}
	return etherphysicalportrefs
}
func flattenListFcPhysicalPortRef(p []*models.FcPhysicalPortRef, d *schema.ResourceData) []map[string]interface{} {
	var fcphysicalportrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		fcphysicalportref := make(map[string]interface{})
		fcphysicalportref["moid"] = item.Moid
		fcphysicalportref["object_type"] = item.ObjectType
		fcphysicalportref["selector"] = item.Selector
		fcphysicalportrefs = append(fcphysicalportrefs, fcphysicalportref)
	}
	return fcphysicalportrefs
}
func flattenListFirmwareRunningFirmwareRef(p []*models.FirmwareRunningFirmwareRef, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		firmwarerunningfirmwareref := make(map[string]interface{})
		firmwarerunningfirmwareref["moid"] = item.Moid
		firmwarerunningfirmwareref["object_type"] = item.ObjectType
		firmwarerunningfirmwareref["selector"] = item.Selector
		firmwarerunningfirmwarerefs = append(firmwarerunningfirmwarerefs, firmwarerunningfirmwareref)
	}
	return firmwarerunningfirmwarerefs
}
func flattenListForecastDefinitionRef(p []*models.ForecastDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		forecastdefinitionref := make(map[string]interface{})
		forecastdefinitionref["moid"] = item.Moid
		forecastdefinitionref["object_type"] = item.ObjectType
		forecastdefinitionref["selector"] = item.Selector
		forecastdefinitionrefs = append(forecastdefinitionrefs, forecastdefinitionref)
	}
	return forecastdefinitionrefs
}
func flattenListGraphicsCardRef(p []*models.GraphicsCardRef, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		graphicscardref := make(map[string]interface{})
		graphicscardref["moid"] = item.Moid
		graphicscardref["object_type"] = item.ObjectType
		graphicscardref["selector"] = item.Selector
		graphicscardrefs = append(graphicscardrefs, graphicscardref)
	}
	return graphicscardrefs
}
func flattenListGraphicsControllerRef(p []*models.GraphicsControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var graphicscontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		graphicscontrollerref := make(map[string]interface{})
		graphicscontrollerref["moid"] = item.Moid
		graphicscontrollerref["object_type"] = item.ObjectType
		graphicscontrollerref["selector"] = item.Selector
		graphicscontrollerrefs = append(graphicscontrollerrefs, graphicscontrollerref)
	}
	return graphicscontrollerrefs
}
func flattenListHclConstraint(p []*models.HclConstraint, d *schema.ResourceData) []map[string]interface{} {
	var hclconstraints []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hclconstraint := make(map[string]interface{})
		if len(item.HclConstraintAO1P1.HclConstraintAO1P1) != 0 {
			j, err := json.Marshal(item.HclConstraintAO1P1.HclConstraintAO1P1)
			if err == nil {
				hclconstraint["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hclconstraint["class_id"] = item.ClassID
		hclconstraint["constraint_name"] = item.ConstraintName
		hclconstraint["constraint_value"] = item.ConstraintValue
		hclconstraint["object_type"] = item.ObjectType
		hclconstraints = append(hclconstraints, hclconstraint)
	}
	return hclconstraints
}
func flattenListHclHyperflexSoftwareCompatibilityInfoRef(p []*models.HclHyperflexSoftwareCompatibilityInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var hclhyperflexsoftwarecompatibilityinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hclhyperflexsoftwarecompatibilityinforef := make(map[string]interface{})
		hclhyperflexsoftwarecompatibilityinforef["moid"] = item.Moid
		hclhyperflexsoftwarecompatibilityinforef["object_type"] = item.ObjectType
		hclhyperflexsoftwarecompatibilityinforef["selector"] = item.Selector
		hclhyperflexsoftwarecompatibilityinforefs = append(hclhyperflexsoftwarecompatibilityinforefs, hclhyperflexsoftwarecompatibilityinforef)
	}
	return hclhyperflexsoftwarecompatibilityinforefs
}
func flattenListHclOperatingSystemRef(p []*models.HclOperatingSystemRef, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hcloperatingsystemref := make(map[string]interface{})
		hcloperatingsystemref["moid"] = item.Moid
		hcloperatingsystemref["object_type"] = item.ObjectType
		hcloperatingsystemref["selector"] = item.Selector
		hcloperatingsystemrefs = append(hcloperatingsystemrefs, hcloperatingsystemref)
	}
	return hcloperatingsystemrefs
}
func flattenListHyperflexAlarmRef(p []*models.HyperflexAlarmRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexalarmrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexalarmref := make(map[string]interface{})
		hyperflexalarmref["moid"] = item.Moid
		hyperflexalarmref["object_type"] = item.ObjectType
		hyperflexalarmref["selector"] = item.Selector
		hyperflexalarmrefs = append(hyperflexalarmrefs, hyperflexalarmref)
	}
	return hyperflexalarmrefs
}
func flattenListHyperflexCapabilityInfoRef(p []*models.HyperflexCapabilityInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexcapabilityinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexcapabilityinforef := make(map[string]interface{})
		hyperflexcapabilityinforef["moid"] = item.Moid
		hyperflexcapabilityinforef["object_type"] = item.ObjectType
		hyperflexcapabilityinforef["selector"] = item.Selector
		hyperflexcapabilityinforefs = append(hyperflexcapabilityinforefs, hyperflexcapabilityinforef)
	}
	return hyperflexcapabilityinforefs
}
func flattenListHyperflexClusterProfileRef(p []*models.HyperflexClusterProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexclusterprofileref := make(map[string]interface{})
		hyperflexclusterprofileref["moid"] = item.Moid
		hyperflexclusterprofileref["object_type"] = item.ObjectType
		hyperflexclusterprofileref["selector"] = item.Selector
		hyperflexclusterprofilerefs = append(hyperflexclusterprofilerefs, hyperflexclusterprofileref)
	}
	return hyperflexclusterprofilerefs
}
func flattenListHyperflexConfigResultEntryRef(p []*models.HyperflexConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultentryrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexconfigresultentryref := make(map[string]interface{})
		hyperflexconfigresultentryref["moid"] = item.Moid
		hyperflexconfigresultentryref["object_type"] = item.ObjectType
		hyperflexconfigresultentryref["selector"] = item.Selector
		hyperflexconfigresultentryrefs = append(hyperflexconfigresultentryrefs, hyperflexconfigresultentryref)
	}
	return hyperflexconfigresultentryrefs
}
func flattenListHyperflexFeatureLimitEntry(p []*models.HyperflexFeatureLimitEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitentrys []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexfeaturelimitentry := make(map[string]interface{})
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {
			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				hyperflexfeaturelimitentry["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hyperflexfeaturelimitentry["class_id"] = item.ClassID
		hyperflexfeaturelimitentry["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexappsettingconstraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					hyperflexappsettingconstraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			hyperflexappsettingconstraint["class_id"] = item.ClassID
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.Constraint, d)
		hyperflexfeaturelimitentry["name"] = item.Name
		hyperflexfeaturelimitentry["object_type"] = item.ObjectType
		hyperflexfeaturelimitentry["value"] = item.Value
		hyperflexfeaturelimitentrys = append(hyperflexfeaturelimitentrys, hyperflexfeaturelimitentry)
	}
	return hyperflexfeaturelimitentrys
}
func flattenListHyperflexHxZoneResiliencyInfoDt(p []*models.HyperflexHxZoneResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxzoneresiliencyinfodts []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexhxzoneresiliencyinfodt := make(map[string]interface{})
		if len(item.HyperflexHxZoneResiliencyInfoDtAO1P1.HyperflexHxZoneResiliencyInfoDtAO1P1) != 0 {
			j, err := json.Marshal(item.HyperflexHxZoneResiliencyInfoDtAO1P1.HyperflexHxZoneResiliencyInfoDtAO1P1)
			if err == nil {
				hyperflexhxzoneresiliencyinfodt["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hyperflexhxzoneresiliencyinfodt["class_id"] = item.ClassID
		hyperflexhxzoneresiliencyinfodt["name"] = item.Name
		hyperflexhxzoneresiliencyinfodt["object_type"] = item.ObjectType
		hyperflexhxzoneresiliencyinfodt["resiliency_info"] = (func(p *models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexhxresiliencyinfodts []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexhxresiliencyinfodt := make(map[string]interface{})
			delete(item.HyperflexHxResiliencyInfoDtAO1P1.HyperflexHxResiliencyInfoDtAO1P1, "ObjectType")
			if len(item.HyperflexHxResiliencyInfoDtAO1P1.HyperflexHxResiliencyInfoDtAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexHxResiliencyInfoDtAO1P1.HyperflexHxResiliencyInfoDtAO1P1)
				if err == nil {
					hyperflexhxresiliencyinfodt["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			hyperflexhxresiliencyinfodt["class_id"] = item.ClassID
			hyperflexhxresiliencyinfodt["data_replication_factor"] = item.DataReplicationFactor
			hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.HddFailuresTolerable
			hyperflexhxresiliencyinfodt["messages"] = item.Messages
			hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.NodeFailuresTolerable
			hyperflexhxresiliencyinfodt["object_type"] = item.ObjectType
			hyperflexhxresiliencyinfodt["policy_compliance"] = item.PolicyCompliance
			hyperflexhxresiliencyinfodt["resiliency_state"] = item.ResiliencyState
			hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.SsdFailuresTolerable

			hyperflexhxresiliencyinfodts = append(hyperflexhxresiliencyinfodts, hyperflexhxresiliencyinfodt)
			return hyperflexhxresiliencyinfodts
		})(item.ResiliencyInfo, d)
		hyperflexhxzoneresiliencyinfodts = append(hyperflexhxzoneresiliencyinfodts, hyperflexhxzoneresiliencyinfodt)
	}
	return hyperflexhxzoneresiliencyinfodts
}
func flattenListHyperflexHxdpVersionRef(p []*models.HyperflexHxdpVersionRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxdpversionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexhxdpversionref := make(map[string]interface{})
		hyperflexhxdpversionref["moid"] = item.Moid
		hyperflexhxdpversionref["object_type"] = item.ObjectType
		hyperflexhxdpversionref["selector"] = item.Selector
		hyperflexhxdpversionrefs = append(hyperflexhxdpversionrefs, hyperflexhxdpversionref)
	}
	return hyperflexhxdpversionrefs
}
func flattenListHyperflexNamedVlan(p []*models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvlans []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexnamedvlan := make(map[string]interface{})
		if len(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1) != 0 {
			j, err := json.Marshal(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1)
			if err == nil {
				hyperflexnamedvlan["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hyperflexnamedvlan["class_id"] = item.ClassID
		hyperflexnamedvlan["name"] = item.Name
		hyperflexnamedvlan["object_type"] = item.ObjectType
		hyperflexnamedvlan["vlan_id"] = item.VlanID
		hyperflexnamedvlans = append(hyperflexnamedvlans, hyperflexnamedvlan)
	}
	return hyperflexnamedvlans
}
func flattenListHyperflexNodeProfileRef(p []*models.HyperflexNodeProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexnodeprofileref := make(map[string]interface{})
		hyperflexnodeprofileref["moid"] = item.Moid
		hyperflexnodeprofileref["object_type"] = item.ObjectType
		hyperflexnodeprofileref["selector"] = item.Selector
		hyperflexnodeprofilerefs = append(hyperflexnodeprofilerefs, hyperflexnodeprofileref)
	}
	return hyperflexnodeprofilerefs
}
func flattenListHyperflexNodeRef(p []*models.HyperflexNodeRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnoderefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexnoderef := make(map[string]interface{})
		hyperflexnoderef["moid"] = item.Moid
		hyperflexnoderef["object_type"] = item.ObjectType
		hyperflexnoderef["selector"] = item.Selector
		hyperflexnoderefs = append(hyperflexnoderefs, hyperflexnoderef)
	}
	return hyperflexnoderefs
}
func flattenListHyperflexServerFirmwareVersionEntry(p []*models.HyperflexServerFirmwareVersionEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionentrys []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexserverfirmwareversionentry := make(map[string]interface{})
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {
			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				hyperflexserverfirmwareversionentry["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hyperflexserverfirmwareversionentry["class_id"] = item.ClassID
		hyperflexserverfirmwareversionentry["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexappsettingconstraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					hyperflexappsettingconstraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			hyperflexappsettingconstraint["class_id"] = item.ClassID
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.Constraint, d)
		hyperflexserverfirmwareversionentry["label"] = item.Label
		hyperflexserverfirmwareversionentry["name"] = item.Name
		hyperflexserverfirmwareversionentry["object_type"] = item.ObjectType
		hyperflexserverfirmwareversionentry["value"] = item.Value
		hyperflexserverfirmwareversionentrys = append(hyperflexserverfirmwareversionentrys, hyperflexserverfirmwareversionentry)
	}
	return hyperflexserverfirmwareversionentrys
}
func flattenListHyperflexServerModelEntry(p []*models.HyperflexServerModelEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelentrys []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexservermodelentry := make(map[string]interface{})
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {
			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				hyperflexservermodelentry["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hyperflexservermodelentry["class_id"] = item.ClassID
		hyperflexservermodelentry["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexappsettingconstraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					hyperflexappsettingconstraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			hyperflexappsettingconstraint["class_id"] = item.ClassID
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.Constraint, d)
		hyperflexservermodelentry["name"] = item.Name
		hyperflexservermodelentry["object_type"] = item.ObjectType
		hyperflexservermodelentry["value"] = item.Value
		hyperflexservermodelentrys = append(hyperflexservermodelentrys, hyperflexservermodelentry)
	}
	return hyperflexservermodelentrys
}
func flattenListIaasConnectorPackRef(p []*models.IaasConnectorPackRef, d *schema.ResourceData) []map[string]interface{} {
	var iaasconnectorpackrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iaasconnectorpackref := make(map[string]interface{})
		iaasconnectorpackref["moid"] = item.Moid
		iaasconnectorpackref["object_type"] = item.ObjectType
		iaasconnectorpackref["selector"] = item.Selector
		iaasconnectorpackrefs = append(iaasconnectorpackrefs, iaasconnectorpackref)
	}
	return iaasconnectorpackrefs
}
func flattenListIaasDeviceStatusRef(p []*models.IaasDeviceStatusRef, d *schema.ResourceData) []map[string]interface{} {
	var iaasdevicestatusrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iaasdevicestatusref := make(map[string]interface{})
		iaasdevicestatusref["moid"] = item.Moid
		iaasdevicestatusref["object_type"] = item.ObjectType
		iaasdevicestatusref["selector"] = item.Selector
		iaasdevicestatusrefs = append(iaasdevicestatusrefs, iaasdevicestatusref)
	}
	return iaasdevicestatusrefs
}
func flattenListIaasLicenseKeysInfo(p []*models.IaasLicenseKeysInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicensekeysinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iaaslicensekeysinfo := make(map[string]interface{})
		if len(item.IaasLicenseKeysInfoAO1P1.IaasLicenseKeysInfoAO1P1) != 0 {
			j, err := json.Marshal(item.IaasLicenseKeysInfoAO1P1.IaasLicenseKeysInfoAO1P1)
			if err == nil {
				iaaslicensekeysinfo["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		iaaslicensekeysinfo["class_id"] = item.ClassID
		iaaslicensekeysinfo["count"] = item.Count
		iaaslicensekeysinfo["expiration_date"] = item.ExpirationDate
		iaaslicensekeysinfo["license_id"] = item.LicenseID
		iaaslicensekeysinfo["object_type"] = item.ObjectType
		iaaslicensekeysinfo["pid"] = item.Pid
		iaaslicensekeysinfos = append(iaaslicensekeysinfos, iaaslicensekeysinfo)
	}
	return iaaslicensekeysinfos
}
func flattenListIaasLicenseUtilizationInfo(p []*models.IaasLicenseUtilizationInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseutilizationinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iaaslicenseutilizationinfo := make(map[string]interface{})
		iaaslicenseutilizationinfo["actual_used"] = item.ActualUsed
		if len(item.IaasLicenseUtilizationInfoAO1P1.IaasLicenseUtilizationInfoAO1P1) != 0 {
			j, err := json.Marshal(item.IaasLicenseUtilizationInfoAO1P1.IaasLicenseUtilizationInfoAO1P1)
			if err == nil {
				iaaslicenseutilizationinfo["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		iaaslicenseutilizationinfo["class_id"] = item.ClassID
		iaaslicenseutilizationinfo["label"] = item.Label
		iaaslicenseutilizationinfo["licensed_limit"] = item.LicensedLimit
		iaaslicenseutilizationinfo["object_type"] = item.ObjectType
		iaaslicenseutilizationinfo["sku"] = item.Sku
		iaaslicenseutilizationinfos = append(iaaslicenseutilizationinfos, iaaslicenseutilizationinfo)
	}
	return iaaslicenseutilizationinfos
}
func flattenListIaasMostRunTasksRef(p []*models.IaasMostRunTasksRef, d *schema.ResourceData) []map[string]interface{} {
	var iaasmostruntasksrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iaasmostruntasksref := make(map[string]interface{})
		iaasmostruntasksref["moid"] = item.Moid
		iaasmostruntasksref["object_type"] = item.ObjectType
		iaasmostruntasksref["selector"] = item.Selector
		iaasmostruntasksrefs = append(iaasmostruntasksrefs, iaasmostruntasksref)
	}
	return iaasmostruntasksrefs
}
func flattenListIamAccountPermissions(p []*models.IamAccountPermissions, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountpermissionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamaccountpermissions := make(map[string]interface{})
		iamaccountpermissions["account_identifier"] = item.AccountIdentifier
		iamaccountpermissions["account_name"] = item.AccountName
		iamaccountpermissions["account_status"] = item.AccountStatus
		if len(item.IamAccountPermissionsAO1P1.IamAccountPermissionsAO1P1) != 0 {
			j, err := json.Marshal(item.IamAccountPermissionsAO1P1.IamAccountPermissionsAO1P1)
			if err == nil {
				iamaccountpermissions["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		iamaccountpermissions["class_id"] = item.ClassID
		iamaccountpermissions["object_type"] = item.ObjectType
		iamaccountpermissions["permissions"] = (func(p []*models.IamPermissionReference, d *schema.ResourceData) []map[string]interface{} {
			var iampermissionreferences []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				iampermissionreference := make(map[string]interface{})
				if len(item.IamPermissionReferenceAO1P1.IamPermissionReferenceAO1P1) != 0 {
					j, err := json.Marshal(item.IamPermissionReferenceAO1P1.IamPermissionReferenceAO1P1)
					if err == nil {
						iampermissionreference["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				iampermissionreference["class_id"] = item.ClassID
				iampermissionreference["object_type"] = item.ObjectType
				iampermissionreference["permission_identifier"] = item.PermissionIdentifier
				iampermissionreference["permission_name"] = item.PermissionName
				iampermissionreferences = append(iampermissionreferences, iampermissionreference)
			}
			return iampermissionreferences
		})(item.Permissions, d)
		iamaccountpermissionss = append(iamaccountpermissionss, iamaccountpermissions)
	}
	return iamaccountpermissionss
}
func flattenListIamAPIKeyRef(p []*models.IamAPIKeyRef, d *schema.ResourceData) []map[string]interface{} {
	var iamapikeyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamapikeyref := make(map[string]interface{})
		iamapikeyref["moid"] = item.Moid
		iamapikeyref["object_type"] = item.ObjectType
		iamapikeyref["selector"] = item.Selector
		iamapikeyrefs = append(iamapikeyrefs, iamapikeyref)
	}
	return iamapikeyrefs
}
func flattenListIamAppRegistrationRef(p []*models.IamAppRegistrationRef, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamappregistrationref := make(map[string]interface{})
		iamappregistrationref["moid"] = item.Moid
		iamappregistrationref["object_type"] = item.ObjectType
		iamappregistrationref["selector"] = item.Selector
		iamappregistrationrefs = append(iamappregistrationrefs, iamappregistrationref)
	}
	return iamappregistrationrefs
}
func flattenListIamDomainGroupRef(p []*models.IamDomainGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamdomaingroupref := make(map[string]interface{})
		iamdomaingroupref["moid"] = item.Moid
		iamdomaingroupref["object_type"] = item.ObjectType
		iamdomaingroupref["selector"] = item.Selector
		iamdomaingrouprefs = append(iamdomaingrouprefs, iamdomaingroupref)
	}
	return iamdomaingrouprefs
}
func flattenListIamEndPointPrivilegeRef(p []*models.IamEndPointPrivilegeRef, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointprivilegerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamendpointprivilegeref := make(map[string]interface{})
		iamendpointprivilegeref["moid"] = item.Moid
		iamendpointprivilegeref["object_type"] = item.ObjectType
		iamendpointprivilegeref["selector"] = item.Selector
		iamendpointprivilegerefs = append(iamendpointprivilegerefs, iamendpointprivilegeref)
	}
	return iamendpointprivilegerefs
}
func flattenListIamEndPointRoleRef(p []*models.IamEndPointRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointrolerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamendpointroleref := make(map[string]interface{})
		iamendpointroleref["moid"] = item.Moid
		iamendpointroleref["object_type"] = item.ObjectType
		iamendpointroleref["selector"] = item.Selector
		iamendpointrolerefs = append(iamendpointrolerefs, iamendpointroleref)
	}
	return iamendpointrolerefs
}
func flattenListIamEndPointUserRoleRef(p []*models.IamEndPointUserRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrolerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamendpointuserroleref := make(map[string]interface{})
		iamendpointuserroleref["moid"] = item.Moid
		iamendpointuserroleref["object_type"] = item.ObjectType
		iamendpointuserroleref["selector"] = item.Selector
		iamendpointuserrolerefs = append(iamendpointuserrolerefs, iamendpointuserroleref)
	}
	return iamendpointuserrolerefs
}
func flattenListIamFeatureDefinition(p []*models.IamFeatureDefinition, d *schema.ResourceData) []map[string]interface{} {
	var iamfeaturedefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamfeaturedefinition := make(map[string]interface{})
		if len(item.IamFeatureDefinitionAO1P1.IamFeatureDefinitionAO1P1) != 0 {
			j, err := json.Marshal(item.IamFeatureDefinitionAO1P1.IamFeatureDefinitionAO1P1)
			if err == nil {
				iamfeaturedefinition["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		iamfeaturedefinition["class_id"] = item.ClassID
		iamfeaturedefinition["feature"] = item.Feature
		iamfeaturedefinition["object_type"] = item.ObjectType
		iamfeaturedefinitions = append(iamfeaturedefinitions, iamfeaturedefinition)
	}
	return iamfeaturedefinitions
}
func flattenListIamGroupPermissionToRoles(p []*models.IamGroupPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iamgrouppermissiontoroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamgrouppermissiontoroles := make(map[string]interface{})
		if len(item.IamGroupPermissionToRolesAO1P1.IamGroupPermissionToRolesAO1P1) != 0 {
			j, err := json.Marshal(item.IamGroupPermissionToRolesAO1P1.IamGroupPermissionToRolesAO1P1)
			if err == nil {
				iamgrouppermissiontoroles["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		iamgrouppermissiontoroles["class_id"] = item.ClassID
		iamgrouppermissiontoroles["group"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			cmrfcmrf := make(map[string]interface{})
			delete(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1, "ObjectType")
			if len(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1) != 0 {
				j, err := json.Marshal(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1)
				if err == nil {
					cmrfcmrf["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			cmrfcmrf["class_id"] = item.ClassID
			cmrfcmrf["moid"] = item.Moid
			cmrfcmrf["object_type"] = item.ObjectType

			cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			return cmrfcmrfs
		})(item.Group, d)
		iamgrouppermissiontoroles["object_type"] = item.ObjectType
		iamgrouppermissiontoroles["orgs"] = (func(p []*models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				cmrfcmrf := make(map[string]interface{})
				if len(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1) != 0 {
					j, err := json.Marshal(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1)
					if err == nil {
						cmrfcmrf["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				cmrfcmrf["class_id"] = item.ClassID
				cmrfcmrf["moid"] = item.Moid
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			}
			return cmrfcmrfs
		})(item.Orgs, d)
		iamgrouppermissiontoroless = append(iamgrouppermissiontoroless, iamgrouppermissiontoroles)
	}
	return iamgrouppermissiontoroless
}
func flattenListIamIdpRef(p []*models.IamIdpRef, d *schema.ResourceData) []map[string]interface{} {
	var iamidprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamidpref := make(map[string]interface{})
		iamidpref["moid"] = item.Moid
		iamidpref["object_type"] = item.ObjectType
		iamidpref["selector"] = item.Selector
		iamidprefs = append(iamidprefs, iamidpref)
	}
	return iamidprefs
}
func flattenListIamIdpReferenceRef(p []*models.IamIdpReferenceRef, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamidpreferenceref := make(map[string]interface{})
		iamidpreferenceref["moid"] = item.Moid
		iamidpreferenceref["object_type"] = item.ObjectType
		iamidpreferenceref["selector"] = item.Selector
		iamidpreferencerefs = append(iamidpreferencerefs, iamidpreferenceref)
	}
	return iamidpreferencerefs
}
func flattenListIamLdapGroupRef(p []*models.IamLdapGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var iamldapgrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamldapgroupref := make(map[string]interface{})
		iamldapgroupref["moid"] = item.Moid
		iamldapgroupref["object_type"] = item.ObjectType
		iamldapgroupref["selector"] = item.Selector
		iamldapgrouprefs = append(iamldapgrouprefs, iamldapgroupref)
	}
	return iamldapgrouprefs
}
func flattenListIamLdapProviderRef(p []*models.IamLdapProviderRef, d *schema.ResourceData) []map[string]interface{} {
	var iamldapproviderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamldapproviderref := make(map[string]interface{})
		iamldapproviderref["moid"] = item.Moid
		iamldapproviderref["object_type"] = item.ObjectType
		iamldapproviderref["selector"] = item.Selector
		iamldapproviderrefs = append(iamldapproviderrefs, iamldapproviderref)
	}
	return iamldapproviderrefs
}
func flattenListIamOAuthTokenRef(p []*models.IamOAuthTokenRef, d *schema.ResourceData) []map[string]interface{} {
	var iamoauthtokenrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamoauthtokenref := make(map[string]interface{})
		iamoauthtokenref["moid"] = item.Moid
		iamoauthtokenref["object_type"] = item.ObjectType
		iamoauthtokenref["selector"] = item.Selector
		iamoauthtokenrefs = append(iamoauthtokenrefs, iamoauthtokenref)
	}
	return iamoauthtokenrefs
}
func flattenListIamPermissionRef(p []*models.IamPermissionRef, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iampermissionref := make(map[string]interface{})
		iampermissionref["moid"] = item.Moid
		iampermissionref["object_type"] = item.ObjectType
		iampermissionref["selector"] = item.Selector
		iampermissionrefs = append(iampermissionrefs, iampermissionref)
	}
	return iampermissionrefs
}
func flattenListIamPermissionToRoles(p []*models.IamPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iampermissiontoroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iampermissiontoroles := make(map[string]interface{})
		if len(item.IamPermissionToRolesAO1P1.IamPermissionToRolesAO1P1) != 0 {
			j, err := json.Marshal(item.IamPermissionToRolesAO1P1.IamPermissionToRolesAO1P1)
			if err == nil {
				iampermissiontoroles["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		iampermissiontoroles["class_id"] = item.ClassID
		iampermissiontoroles["object_type"] = item.ObjectType
		iampermissiontoroles["permission"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			cmrfcmrf := make(map[string]interface{})
			delete(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1, "ObjectType")
			if len(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1) != 0 {
				j, err := json.Marshal(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1)
				if err == nil {
					cmrfcmrf["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			cmrfcmrf["class_id"] = item.ClassID
			cmrfcmrf["moid"] = item.Moid
			cmrfcmrf["object_type"] = item.ObjectType

			cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			return cmrfcmrfs
		})(item.Permission, d)
		iampermissiontoroles["roles"] = (func(p []*models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				cmrfcmrf := make(map[string]interface{})
				if len(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1) != 0 {
					j, err := json.Marshal(item.CmrfCmRfAO1P1.CmrfCmRfAO1P1)
					if err == nil {
						cmrfcmrf["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				cmrfcmrf["class_id"] = item.ClassID
				cmrfcmrf["moid"] = item.Moid
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			}
			return cmrfcmrfs
		})(item.Roles, d)
		iampermissiontoroless = append(iampermissiontoroless, iampermissiontoroles)
	}
	return iampermissiontoroless
}
func flattenListIamPrivilegeRef(p []*models.IamPrivilegeRef, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamprivilegeref := make(map[string]interface{})
		iamprivilegeref["moid"] = item.Moid
		iamprivilegeref["object_type"] = item.ObjectType
		iamprivilegeref["selector"] = item.Selector
		iamprivilegerefs = append(iamprivilegerefs, iamprivilegeref)
	}
	return iamprivilegerefs
}
func flattenListIamPrivilegeSetRef(p []*models.IamPrivilegeSetRef, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegesetrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamprivilegesetref := make(map[string]interface{})
		iamprivilegesetref["moid"] = item.Moid
		iamprivilegesetref["object_type"] = item.ObjectType
		iamprivilegesetref["selector"] = item.Selector
		iamprivilegesetrefs = append(iamprivilegesetrefs, iamprivilegesetref)
	}
	return iamprivilegesetrefs
}
func flattenListIamResourcePermissionRef(p []*models.IamResourcePermissionRef, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcepermissionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamresourcepermissionref := make(map[string]interface{})
		iamresourcepermissionref["moid"] = item.Moid
		iamresourcepermissionref["object_type"] = item.ObjectType
		iamresourcepermissionref["selector"] = item.Selector
		iamresourcepermissionrefs = append(iamresourcepermissionrefs, iamresourcepermissionref)
	}
	return iamresourcepermissionrefs
}
func flattenListIamResourceRolesRef(p []*models.IamResourceRolesRef, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcerolesrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamresourcerolesref := make(map[string]interface{})
		iamresourcerolesref["moid"] = item.Moid
		iamresourcerolesref["object_type"] = item.ObjectType
		iamresourcerolesref["selector"] = item.Selector
		iamresourcerolesrefs = append(iamresourcerolesrefs, iamresourcerolesref)
	}
	return iamresourcerolesrefs
}
func flattenListIamRoleRef(p []*models.IamRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var iamrolerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamroleref := make(map[string]interface{})
		iamroleref["moid"] = item.Moid
		iamroleref["object_type"] = item.ObjectType
		iamroleref["selector"] = item.Selector
		iamrolerefs = append(iamrolerefs, iamroleref)
	}
	return iamrolerefs
}
func flattenListIamSessionRef(p []*models.IamSessionRef, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamsessionref := make(map[string]interface{})
		iamsessionref["moid"] = item.Moid
		iamsessionref["object_type"] = item.ObjectType
		iamsessionref["selector"] = item.Selector
		iamsessionrefs = append(iamsessionrefs, iamsessionref)
	}
	return iamsessionrefs
}
func flattenListIamUserGroupRef(p []*models.IamUserGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamusergroupref := make(map[string]interface{})
		iamusergroupref["moid"] = item.Moid
		iamusergroupref["object_type"] = item.ObjectType
		iamusergroupref["selector"] = item.Selector
		iamusergrouprefs = append(iamusergrouprefs, iamusergroupref)
	}
	return iamusergrouprefs
}
func flattenListIamUserLoginTimeRef(p []*models.IamUserLoginTimeRef, d *schema.ResourceData) []map[string]interface{} {
	var iamuserlogintimerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamuserlogintimeref := make(map[string]interface{})
		iamuserlogintimeref["moid"] = item.Moid
		iamuserlogintimeref["object_type"] = item.ObjectType
		iamuserlogintimeref["selector"] = item.Selector
		iamuserlogintimerefs = append(iamuserlogintimerefs, iamuserlogintimeref)
	}
	return iamuserlogintimerefs
}
func flattenListIamUserPreferenceRef(p []*models.IamUserPreferenceRef, d *schema.ResourceData) []map[string]interface{} {
	var iamuserpreferencerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamuserpreferenceref := make(map[string]interface{})
		iamuserpreferenceref["moid"] = item.Moid
		iamuserpreferenceref["object_type"] = item.ObjectType
		iamuserpreferenceref["selector"] = item.Selector
		iamuserpreferencerefs = append(iamuserpreferencerefs, iamuserpreferenceref)
	}
	return iamuserpreferencerefs
}
func flattenListIamUserRef(p []*models.IamUserRef, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		iamuserref := make(map[string]interface{})
		iamuserref["moid"] = item.Moid
		iamuserref["object_type"] = item.ObjectType
		iamuserref["selector"] = item.Selector
		iamuserrefs = append(iamuserrefs, iamuserref)
	}
	return iamuserrefs
}
func flattenListInventoryGenericInventoryHolderRef(p []*models.InventoryGenericInventoryHolderRef, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		inventorygenericinventoryholderref := make(map[string]interface{})
		inventorygenericinventoryholderref["moid"] = item.Moid
		inventorygenericinventoryholderref["object_type"] = item.ObjectType
		inventorygenericinventoryholderref["selector"] = item.Selector
		inventorygenericinventoryholderrefs = append(inventorygenericinventoryholderrefs, inventorygenericinventoryholderref)
	}
	return inventorygenericinventoryholderrefs
}
func flattenListInventoryGenericInventoryRef(p []*models.InventoryGenericInventoryRef, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		inventorygenericinventoryref := make(map[string]interface{})
		inventorygenericinventoryref["moid"] = item.Moid
		inventorygenericinventoryref["object_type"] = item.ObjectType
		inventorygenericinventoryref["selector"] = item.Selector
		inventorygenericinventoryrefs = append(inventorygenericinventoryrefs, inventorygenericinventoryref)
	}
	return inventorygenericinventoryrefs
}
func flattenListLicenseLicenseInfoRef(p []*models.LicenseLicenseInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var licenselicenseinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		licenselicenseinforef := make(map[string]interface{})
		licenselicenseinforef["moid"] = item.Moid
		licenselicenseinforef["object_type"] = item.ObjectType
		licenselicenseinforef["selector"] = item.Selector
		licenselicenseinforefs = append(licenselicenseinforefs, licenselicenseinforef)
	}
	return licenselicenseinforefs
}
func flattenListManagementInterfaceRef(p []*models.ManagementInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var managementinterfacerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		managementinterfaceref := make(map[string]interface{})
		managementinterfaceref["moid"] = item.Moid
		managementinterfaceref["object_type"] = item.ObjectType
		managementinterfaceref["selector"] = item.Selector
		managementinterfacerefs = append(managementinterfacerefs, managementinterfaceref)
	}
	return managementinterfacerefs
}
func flattenListMemoryArrayRef(p []*models.MemoryArrayRef, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memoryarrayref := make(map[string]interface{})
		memoryarrayref["moid"] = item.Moid
		memoryarrayref["object_type"] = item.ObjectType
		memoryarrayref["selector"] = item.Selector
		memoryarrayrefs = append(memoryarrayrefs, memoryarrayref)
	}
	return memoryarrayrefs
}
func flattenListMemoryPersistentMemoryGoal(p []*models.MemoryPersistentMemoryGoal, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorygoals []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memorypersistentmemorygoal := make(map[string]interface{})
		if len(item.MemoryPersistentMemoryGoalAO1P1.MemoryPersistentMemoryGoalAO1P1) != 0 {
			j, err := json.Marshal(item.MemoryPersistentMemoryGoalAO1P1.MemoryPersistentMemoryGoalAO1P1)
			if err == nil {
				memorypersistentmemorygoal["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		memorypersistentmemorygoal["class_id"] = item.ClassID
		memorypersistentmemorygoal["memory_mode_percentage"] = item.MemoryModePercentage
		memorypersistentmemorygoal["object_type"] = item.ObjectType
		memorypersistentmemorygoal["persistent_memory_type"] = item.PersistentMemoryType
		memorypersistentmemorygoal["socket_id"] = item.SocketID
		memorypersistentmemorygoals = append(memorypersistentmemorygoals, memorypersistentmemorygoal)
	}
	return memorypersistentmemorygoals
}
func flattenListMemoryPersistentMemoryLogicalNamespace(p []*models.MemoryPersistentMemoryLogicalNamespace, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylogicalnamespaces []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memorypersistentmemorylogicalnamespace := make(map[string]interface{})
		if len(item.MemoryPersistentMemoryLogicalNamespaceAO1P1.MemoryPersistentMemoryLogicalNamespaceAO1P1) != 0 {
			j, err := json.Marshal(item.MemoryPersistentMemoryLogicalNamespaceAO1P1.MemoryPersistentMemoryLogicalNamespaceAO1P1)
			if err == nil {
				memorypersistentmemorylogicalnamespace["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		memorypersistentmemorylogicalnamespace["capacity"] = item.Capacity
		memorypersistentmemorylogicalnamespace["class_id"] = item.ClassID
		memorypersistentmemorylogicalnamespace["mode"] = item.Mode
		memorypersistentmemorylogicalnamespace["name"] = item.Name
		memorypersistentmemorylogicalnamespace["object_type"] = item.ObjectType
		memorypersistentmemorylogicalnamespace["socket_id"] = item.SocketID
		memorypersistentmemorylogicalnamespace["socket_memory_id"] = item.SocketMemoryID
		memorypersistentmemorylogicalnamespaces = append(memorypersistentmemorylogicalnamespaces, memorypersistentmemorylogicalnamespace)
	}
	return memorypersistentmemorylogicalnamespaces
}
func flattenListMemoryPersistentMemoryNamespaceConfigResultRef(p []*models.MemoryPersistentMemoryNamespaceConfigResultRef, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespaceconfigresultrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memorypersistentmemorynamespaceconfigresultref := make(map[string]interface{})
		memorypersistentmemorynamespaceconfigresultref["moid"] = item.Moid
		memorypersistentmemorynamespaceconfigresultref["object_type"] = item.ObjectType
		memorypersistentmemorynamespaceconfigresultref["selector"] = item.Selector
		memorypersistentmemorynamespaceconfigresultrefs = append(memorypersistentmemorynamespaceconfigresultrefs, memorypersistentmemorynamespaceconfigresultref)
	}
	return memorypersistentmemorynamespaceconfigresultrefs
}
func flattenListMemoryPersistentMemoryNamespaceRef(p []*models.MemoryPersistentMemoryNamespaceRef, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespacerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memorypersistentmemorynamespaceref := make(map[string]interface{})
		memorypersistentmemorynamespaceref["moid"] = item.Moid
		memorypersistentmemorynamespaceref["object_type"] = item.ObjectType
		memorypersistentmemorynamespaceref["selector"] = item.Selector
		memorypersistentmemorynamespacerefs = append(memorypersistentmemorynamespacerefs, memorypersistentmemorynamespaceref)
	}
	return memorypersistentmemorynamespacerefs
}
func flattenListMemoryPersistentMemoryRegionRef(p []*models.MemoryPersistentMemoryRegionRef, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memorypersistentmemoryregionref := make(map[string]interface{})
		memorypersistentmemoryregionref["moid"] = item.Moid
		memorypersistentmemoryregionref["object_type"] = item.ObjectType
		memorypersistentmemoryregionref["selector"] = item.Selector
		memorypersistentmemoryregionrefs = append(memorypersistentmemoryregionrefs, memorypersistentmemoryregionref)
	}
	return memorypersistentmemoryregionrefs
}
func flattenListMemoryPersistentMemoryUnitRef(p []*models.MemoryPersistentMemoryUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memorypersistentmemoryunitref := make(map[string]interface{})
		memorypersistentmemoryunitref["moid"] = item.Moid
		memorypersistentmemoryunitref["object_type"] = item.ObjectType
		memorypersistentmemoryunitref["selector"] = item.Selector
		memorypersistentmemoryunitrefs = append(memorypersistentmemoryunitrefs, memorypersistentmemoryunitref)
	}
	return memorypersistentmemoryunitrefs
}
func flattenListMemoryUnitRef(p []*models.MemoryUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var memoryunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memoryunitref := make(map[string]interface{})
		memoryunitref["moid"] = item.Moid
		memoryunitref["object_type"] = item.ObjectType
		memoryunitref["selector"] = item.Selector
		memoryunitrefs = append(memoryunitrefs, memoryunitref)
	}
	return memoryunitrefs
}
func flattenListMetaAccessPrivilege(p []*models.MetaAccessPrivilege, d *schema.ResourceData) []map[string]interface{} {
	var metaaccessprivileges []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		metaaccessprivilege := make(map[string]interface{})
		if len(item.MetaAccessPrivilegeAO1P1.MetaAccessPrivilegeAO1P1) != 0 {
			j, err := json.Marshal(item.MetaAccessPrivilegeAO1P1.MetaAccessPrivilegeAO1P1)
			if err == nil {
				metaaccessprivilege["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		metaaccessprivilege["class_id"] = item.ClassID
		metaaccessprivilege["method"] = item.Method
		metaaccessprivilege["object_type"] = item.ObjectType
		metaaccessprivilege["privilege"] = item.Privilege
		metaaccessprivileges = append(metaaccessprivileges, metaaccessprivilege)
	}
	return metaaccessprivileges
}
func flattenListMetaPropDefinition(p []*models.MetaPropDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metapropdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		metapropdefinition := make(map[string]interface{})
		metapropdefinition["api_access"] = item.APIAccess
		if len(item.MetaPropDefinitionAO1P1.MetaPropDefinitionAO1P1) != 0 {
			j, err := json.Marshal(item.MetaPropDefinitionAO1P1.MetaPropDefinitionAO1P1)
			if err == nil {
				metapropdefinition["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		metapropdefinition["class_id"] = item.ClassID
		metapropdefinition["name"] = item.Name
		metapropdefinition["object_type"] = item.ObjectType
		metapropdefinition["op_security"] = item.OpSecurity
		metapropdefinition["search_weight"] = item.SearchWeight
		metapropdefinitions = append(metapropdefinitions, metapropdefinition)
	}
	return metapropdefinitions
}
func flattenListMetaRelationshipDefinition(p []*models.MetaRelationshipDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metarelationshipdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		metarelationshipdefinition := make(map[string]interface{})
		metarelationshipdefinition["api_access"] = item.APIAccess
		if len(item.MetaRelationshipDefinitionAO1P1.MetaRelationshipDefinitionAO1P1) != 0 {
			j, err := json.Marshal(item.MetaRelationshipDefinitionAO1P1.MetaRelationshipDefinitionAO1P1)
			if err == nil {
				metarelationshipdefinition["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		metarelationshipdefinition["class_id"] = item.ClassID
		metarelationshipdefinition["collection"] = item.Collection
		metarelationshipdefinition["export"] = item.Export
		metarelationshipdefinition["export_with_peer"] = item.ExportWithPeer
		metarelationshipdefinition["name"] = item.Name
		metarelationshipdefinition["object_type"] = item.ObjectType
		metarelationshipdefinition["type"] = item.Type
		metarelationshipdefinitions = append(metarelationshipdefinitions, metarelationshipdefinition)
	}
	return metarelationshipdefinitions
}
func flattenListMoBaseMoRef(p []*models.MoBaseMoRef, d *schema.ResourceData) []map[string]interface{} {
	var mobasemorefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		mobasemoref := make(map[string]interface{})
		mobasemoref["moid"] = item.Moid
		mobasemoref["object_type"] = item.ObjectType
		mobasemoref["selector"] = item.Selector
		mobasemorefs = append(mobasemorefs, mobasemoref)
	}
	return mobasemorefs
}
func flattenListMoTag(p []*models.MoTag, d *schema.ResourceData) []map[string]interface{} {
	var motags []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		motag := make(map[string]interface{})
		if len(item.MoTagAO1P1.MoTagAO1P1) != 0 {
			j, err := json.Marshal(item.MoTagAO1P1.MoTagAO1P1)
			if err == nil {
				motag["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		motag["class_id"] = item.ClassID
		motag["key"] = item.Key
		motag["object_type"] = item.ObjectType
		motag["value"] = item.Value
		motags = append(motags, motag)
	}
	return motags
}
func flattenListNetworkElementRef(p []*models.NetworkElementRef, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		networkelementref := make(map[string]interface{})
		networkelementref["moid"] = item.Moid
		networkelementref["object_type"] = item.ObjectType
		networkelementref["selector"] = item.Selector
		networkelementrefs = append(networkelementrefs, networkelementref)
	}
	return networkelementrefs
}
func flattenListNiaapiDetail(p []*models.NiaapiDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapidetails []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		niaapidetail := make(map[string]interface{})
		if len(item.NiaapiDetailAO1P1.NiaapiDetailAO1P1) != 0 {
			j, err := json.Marshal(item.NiaapiDetailAO1P1.NiaapiDetailAO1P1)
			if err == nil {
				niaapidetail["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		niaapidetail["chksum"] = item.Chksum
		niaapidetail["class_id"] = item.ClassID
		niaapidetail["filename"] = item.Filename
		niaapidetail["name"] = item.Name
		niaapidetail["object_type"] = item.ObjectType
		niaapidetails = append(niaapidetails, niaapidetail)
	}
	return niaapidetails
}
func flattenListNiaapiRevisionInfo(p []*models.NiaapiRevisionInfo, d *schema.ResourceData) []map[string]interface{} {
	var niaapirevisioninfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		niaapirevisioninfo := make(map[string]interface{})
		if len(item.NiaapiRevisionInfoAO1P1.NiaapiRevisionInfoAO1P1) != 0 {
			j, err := json.Marshal(item.NiaapiRevisionInfoAO1P1.NiaapiRevisionInfoAO1P1)
			if err == nil {
				niaapirevisioninfo["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		niaapirevisioninfo["class_id"] = item.ClassID
		niaapirevisioninfo["date_published"] = item.DatePublished
		niaapirevisioninfo["object_type"] = item.ObjectType
		niaapirevisioninfo["revision_comment"] = item.RevisionComment
		niaapirevisioninfo["revision_no"] = item.RevisionNo
		niaapirevisioninfos = append(niaapirevisioninfos, niaapirevisioninfo)
	}
	return niaapirevisioninfos
}
func flattenListOnpremImagePackage(p []*models.OnpremImagePackage, d *schema.ResourceData) []map[string]interface{} {
	var onpremimagepackages []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		onpremimagepackage := make(map[string]interface{})
		if len(item.OnpremImagePackageAO1P1.OnpremImagePackageAO1P1) != 0 {
			j, err := json.Marshal(item.OnpremImagePackageAO1P1.OnpremImagePackageAO1P1)
			if err == nil {
				onpremimagepackage["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		onpremimagepackage["class_id"] = item.ClassID
		onpremimagepackage["file_path"] = item.FilePath
		onpremimagepackage["file_sha"] = item.FileSha
		onpremimagepackage["file_size"] = item.FileSize
		onpremimagepackage["file_time"] = item.FileTime
		onpremimagepackage["filename"] = item.Filename
		onpremimagepackage["name"] = item.Name
		onpremimagepackage["object_type"] = item.ObjectType
		onpremimagepackage["package_type"] = item.PackageType
		onpremimagepackage["version"] = item.Version
		onpremimagepackages = append(onpremimagepackages, onpremimagepackage)
	}
	return onpremimagepackages
}
func flattenListOnpremUpgradeNote(p []*models.OnpremUpgradeNote, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradenotes []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		onpremupgradenote := make(map[string]interface{})
		if len(item.OnpremUpgradeNoteAO1P1.OnpremUpgradeNoteAO1P1) != 0 {
			j, err := json.Marshal(item.OnpremUpgradeNoteAO1P1.OnpremUpgradeNoteAO1P1)
			if err == nil {
				onpremupgradenote["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		onpremupgradenote["class_id"] = item.ClassID
		onpremupgradenote["message"] = item.Message
		onpremupgradenote["object_type"] = item.ObjectType
		onpremupgradenotes = append(onpremupgradenotes, onpremupgradenote)
	}
	return onpremupgradenotes
}
func flattenListOnpremUpgradePhase(p []*models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradephases []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		onpremupgradephase := make(map[string]interface{})
		if len(item.OnpremUpgradePhaseAO1P1.OnpremUpgradePhaseAO1P1) != 0 {
			j, err := json.Marshal(item.OnpremUpgradePhaseAO1P1.OnpremUpgradePhaseAO1P1)
			if err == nil {
				onpremupgradephase["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		onpremupgradephase["class_id"] = item.ClassID
		onpremupgradephase["elapsed_time"] = item.ElapsedTime
		onpremupgradephase["end_time"] = item.EndTime
		onpremupgradephase["failed"] = item.Failed
		onpremupgradephase["message"] = item.Message
		onpremupgradephase["name"] = item.Name
		onpremupgradephase["object_type"] = item.ObjectType
		onpremupgradephase["start_time"] = item.StartTime
		onpremupgradephases = append(onpremupgradephases, onpremupgradephase)
	}
	return onpremupgradephases
}
func flattenListOrganizationOrganizationRef(p []*models.OrganizationOrganizationRef, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		organizationorganizationref := make(map[string]interface{})
		organizationorganizationref["moid"] = item.Moid
		organizationorganizationref["object_type"] = item.ObjectType
		organizationorganizationref["selector"] = item.Selector
		organizationorganizationrefs = append(organizationorganizationrefs, organizationorganizationref)
	}
	return organizationorganizationrefs
}
func flattenListOsConfigurationFileRef(p []*models.OsConfigurationFileRef, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		osconfigurationfileref := make(map[string]interface{})
		osconfigurationfileref["moid"] = item.Moid
		osconfigurationfileref["object_type"] = item.ObjectType
		osconfigurationfileref["selector"] = item.Selector
		osconfigurationfilerefs = append(osconfigurationfilerefs, osconfigurationfileref)
	}
	return osconfigurationfilerefs
}
func flattenListOsDistributionRef(p []*models.OsDistributionRef, d *schema.ResourceData) []map[string]interface{} {
	var osdistributionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		osdistributionref := make(map[string]interface{})
		osdistributionref["moid"] = item.Moid
		osdistributionref["object_type"] = item.ObjectType
		osdistributionref["selector"] = item.Selector
		osdistributionrefs = append(osdistributionrefs, osdistributionref)
	}
	return osdistributionrefs
}
func flattenListOsPlaceHolder(p []*models.OsPlaceHolder, d *schema.ResourceData) []map[string]interface{} {
	var osplaceholders []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		osplaceholder := make(map[string]interface{})
		if len(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1) != 0 {
			j, err := json.Marshal(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1)
			if err == nil {
				osplaceholder["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		osplaceholder["class_id"] = item.ClassID
		osplaceholder["is_value_set"] = item.IsValueSet
		osplaceholder["object_type"] = item.ObjectType
		osplaceholder["type"] = (func(p *models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {
			var workflowprimitivedatatypes []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			workflowprimitivedatatype := make(map[string]interface{})
			delete(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1, "ObjectType")
			if len(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1) != 0 {
				j, err := json.Marshal(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1)
				if err == nil {
					workflowprimitivedatatype["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			workflowprimitivedatatype["class_id"] = item.ClassID
			workflowprimitivedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				workflowdefaultvalue := make(map[string]interface{})
				delete(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1, "ObjectType")
				if len(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1) != 0 {
					j, err := json.Marshal(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1)
					if err == nil {
						workflowdefaultvalue["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				workflowdefaultvalue["class_id"] = item.ClassID
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.Default, d)
			workflowprimitivedatatype["description"] = item.Description
			workflowprimitivedatatype["label"] = item.Label
			workflowprimitivedatatype["name"] = item.Name
			workflowprimitivedatatype["object_type"] = item.ObjectType
			workflowprimitivedatatype["properties"] = (func(p *models.WorkflowPrimitiveDataProperty, d *schema.ResourceData) []map[string]interface{} {
				var workflowprimitivedatapropertys []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				workflowprimitivedataproperty := make(map[string]interface{})
				delete(item.WorkflowPrimitiveDataPropertyAO1P1.WorkflowPrimitiveDataPropertyAO1P1, "ObjectType")
				if len(item.WorkflowPrimitiveDataPropertyAO1P1.WorkflowPrimitiveDataPropertyAO1P1) != 0 {
					j, err := json.Marshal(item.WorkflowPrimitiveDataPropertyAO1P1.WorkflowPrimitiveDataPropertyAO1P1)
					if err == nil {
						workflowprimitivedataproperty["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				workflowprimitivedataproperty["class_id"] = item.ClassID
				workflowprimitivedataproperty["constraints"] = (func(p *models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {
					var workflowconstraintss []map[string]interface{}
					if p == nil {
						return nil
					}
					item := *p
					workflowconstraints := make(map[string]interface{})
					delete(item.WorkflowConstraintsAO1P1.WorkflowConstraintsAO1P1, "ObjectType")
					if len(item.WorkflowConstraintsAO1P1.WorkflowConstraintsAO1P1) != 0 {
						j, err := json.Marshal(item.WorkflowConstraintsAO1P1.WorkflowConstraintsAO1P1)
						if err == nil {
							workflowconstraints["additional_properties"] = string(j)
						} else {
							log.Printf("Error occured while flattening and json parsing: %s", err)
						}
					}
					workflowconstraints["class_id"] = item.ClassID
					workflowconstraints["enum_list"] = (func(p []*models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
						var workflowenumentrys []map[string]interface{}
						if p == nil {
							return nil
						}
						for _, item := range p {
							item := *item
							workflowenumentry := make(map[string]interface{})
							if len(item.WorkflowEnumEntryAO1P1.WorkflowEnumEntryAO1P1) != 0 {
								j, err := json.Marshal(item.WorkflowEnumEntryAO1P1.WorkflowEnumEntryAO1P1)
								if err == nil {
									workflowenumentry["additional_properties"] = string(j)
								} else {
									log.Printf("Error occured while flattening and json parsing: %s", err)
								}
							}
							workflowenumentry["class_id"] = item.ClassID
							workflowenumentry["label"] = item.Label
							workflowenumentry["object_type"] = item.ObjectType
							workflowenumentry["value"] = item.Value
							workflowenumentrys = append(workflowenumentrys, workflowenumentry)
						}
						return workflowenumentrys
					})(item.EnumList, d)
					workflowconstraints["max"] = item.Max
					workflowconstraints["min"] = item.Min
					workflowconstraints["object_type"] = item.ObjectType
					workflowconstraints["regex"] = item.Regex

					workflowconstraintss = append(workflowconstraintss, workflowconstraints)
					return workflowconstraintss
				})(item.Constraints, d)
				workflowprimitivedataproperty["inventory_selector"] = (func(p []*models.WorkflowMoReferenceProperty, d *schema.ResourceData) []map[string]interface{} {
					var workflowmoreferencepropertys []map[string]interface{}
					if p == nil {
						return nil
					}
					for _, item := range p {
						item := *item
						workflowmoreferenceproperty := make(map[string]interface{})
						if len(item.WorkflowMoReferencePropertyAO1P1.WorkflowMoReferencePropertyAO1P1) != 0 {
							j, err := json.Marshal(item.WorkflowMoReferencePropertyAO1P1.WorkflowMoReferencePropertyAO1P1)
							if err == nil {
								workflowmoreferenceproperty["additional_properties"] = string(j)
							} else {
								log.Printf("Error occured while flattening and json parsing: %s", err)
							}
						}
						workflowmoreferenceproperty["class_id"] = item.ClassID
						workflowmoreferenceproperty["display_attributes"] = item.DisplayAttributes
						workflowmoreferenceproperty["object_type"] = item.ObjectType
						workflowmoreferenceproperty["selector"] = item.Selector
						workflowmoreferenceproperty["value_attribute"] = item.ValueAttribute
						workflowmoreferencepropertys = append(workflowmoreferencepropertys, workflowmoreferenceproperty)
					}
					return workflowmoreferencepropertys
				})(item.InventorySelector, d)
				workflowprimitivedataproperty["object_type"] = item.ObjectType
				workflowprimitivedataproperty["secure"] = item.Secure
				workflowprimitivedataproperty["type"] = item.Type

				workflowprimitivedatapropertys = append(workflowprimitivedatapropertys, workflowprimitivedataproperty)
				return workflowprimitivedatapropertys
			})(item.Properties, d)
			workflowprimitivedatatype["required"] = item.Required

			workflowprimitivedatatypes = append(workflowprimitivedatatypes, workflowprimitivedatatype)
			return workflowprimitivedatatypes
		})(item.Type, d)
		osplaceholder["value"] = item.Value
		osplaceholders = append(osplaceholders, osplaceholder)
	}
	return osplaceholders
}
func flattenListOsPostInstallScriptRef(p []*models.OsPostInstallScriptRef, d *schema.ResourceData) []map[string]interface{} {
	var ospostinstallscriptrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		ospostinstallscriptref := make(map[string]interface{})
		ospostinstallscriptref["moid"] = item.Moid
		ospostinstallscriptref["object_type"] = item.ObjectType
		ospostinstallscriptref["selector"] = item.Selector
		ospostinstallscriptrefs = append(ospostinstallscriptrefs, ospostinstallscriptref)
	}
	return ospostinstallscriptrefs
}
func flattenListPciCoprocessorCardRef(p []*models.PciCoprocessorCardRef, d *schema.ResourceData) []map[string]interface{} {
	var pcicoprocessorcardrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pcicoprocessorcardref := make(map[string]interface{})
		pcicoprocessorcardref["moid"] = item.Moid
		pcicoprocessorcardref["object_type"] = item.ObjectType
		pcicoprocessorcardref["selector"] = item.Selector
		pcicoprocessorcardrefs = append(pcicoprocessorcardrefs, pcicoprocessorcardref)
	}
	return pcicoprocessorcardrefs
}
func flattenListPciDeviceRef(p []*models.PciDeviceRef, d *schema.ResourceData) []map[string]interface{} {
	var pcidevicerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pcideviceref := make(map[string]interface{})
		pcideviceref["moid"] = item.Moid
		pcideviceref["object_type"] = item.ObjectType
		pcideviceref["selector"] = item.Selector
		pcidevicerefs = append(pcidevicerefs, pcideviceref)
	}
	return pcidevicerefs
}
func flattenListPciLinkRef(p []*models.PciLinkRef, d *schema.ResourceData) []map[string]interface{} {
	var pcilinkrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pcilinkref := make(map[string]interface{})
		pcilinkref["moid"] = item.Moid
		pcilinkref["object_type"] = item.ObjectType
		pcilinkref["selector"] = item.Selector
		pcilinkrefs = append(pcilinkrefs, pcilinkref)
	}
	return pcilinkrefs
}
func flattenListPciSwitchRef(p []*models.PciSwitchRef, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pciswitchref := make(map[string]interface{})
		pciswitchref["moid"] = item.Moid
		pciswitchref["object_type"] = item.ObjectType
		pciswitchref["selector"] = item.Selector
		pciswitchrefs = append(pciswitchrefs, pciswitchref)
	}
	return pciswitchrefs
}
func flattenListPolicyAbstractConfigProfileRef(p []*models.PolicyAbstractConfigProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractconfigprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		policyabstractconfigprofileref := make(map[string]interface{})
		policyabstractconfigprofileref["moid"] = item.Moid
		policyabstractconfigprofileref["object_type"] = item.ObjectType
		policyabstractconfigprofileref["selector"] = item.Selector
		policyabstractconfigprofilerefs = append(policyabstractconfigprofilerefs, policyabstractconfigprofileref)
	}
	return policyabstractconfigprofilerefs
}
func flattenListPolicyinventoryJobInfo(p []*models.PolicyinventoryJobInfo, d *schema.ResourceData) []map[string]interface{} {
	var policyinventoryjobinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		policyinventoryjobinfo := make(map[string]interface{})
		if len(item.PolicyinventoryJobInfoAO1P1.PolicyinventoryJobInfoAO1P1) != 0 {
			j, err := json.Marshal(item.PolicyinventoryJobInfoAO1P1.PolicyinventoryJobInfoAO1P1)
			if err == nil {
				policyinventoryjobinfo["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		policyinventoryjobinfo["class_id"] = item.ClassID
		policyinventoryjobinfo["execution_status"] = item.ExecutionStatus
		policyinventoryjobinfo["last_scheduled_time"] = item.LastScheduledTime
		policyinventoryjobinfo["object_type"] = item.ObjectType
		policyinventoryjobinfo["policy_id"] = item.PolicyID
		policyinventoryjobinfo["policy_name"] = item.PolicyName
		policyinventoryjobinfos = append(policyinventoryjobinfos, policyinventoryjobinfo)
	}
	return policyinventoryjobinfos
}
func flattenListPortGroupRef(p []*models.PortGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		portgroupref := make(map[string]interface{})
		portgroupref["moid"] = item.Moid
		portgroupref["object_type"] = item.ObjectType
		portgroupref["selector"] = item.Selector
		portgrouprefs = append(portgrouprefs, portgroupref)
	}
	return portgrouprefs
}
func flattenListPortSubGroupRef(p []*models.PortSubGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		portsubgroupref := make(map[string]interface{})
		portsubgroupref["moid"] = item.Moid
		portsubgroupref["object_type"] = item.ObjectType
		portsubgroupref["selector"] = item.Selector
		portsubgrouprefs = append(portsubgrouprefs, portsubgroupref)
	}
	return portsubgrouprefs
}
func flattenListProcessorUnitRef(p []*models.ProcessorUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var processorunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		processorunitref := make(map[string]interface{})
		processorunitref["moid"] = item.Moid
		processorunitref["object_type"] = item.ObjectType
		processorunitref["selector"] = item.Selector
		processorunitrefs = append(processorunitrefs, processorunitref)
	}
	return processorunitrefs
}
func flattenListRecoveryBackupProfileRef(p []*models.RecoveryBackupProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		recoverybackupprofileref := make(map[string]interface{})
		recoverybackupprofileref["moid"] = item.Moid
		recoverybackupprofileref["object_type"] = item.ObjectType
		recoverybackupprofileref["selector"] = item.Selector
		recoverybackupprofilerefs = append(recoverybackupprofilerefs, recoverybackupprofileref)
	}
	return recoverybackupprofilerefs
}
func flattenListRecoveryConfigResultEntryRef(p []*models.RecoveryConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultentryrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		recoveryconfigresultentryref := make(map[string]interface{})
		recoveryconfigresultentryref["moid"] = item.Moid
		recoveryconfigresultentryref["object_type"] = item.ObjectType
		recoveryconfigresultentryref["selector"] = item.Selector
		recoveryconfigresultentryrefs = append(recoveryconfigresultentryrefs, recoveryconfigresultentryref)
	}
	return recoveryconfigresultentryrefs
}
func flattenListResourceGroupRef(p []*models.ResourceGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resourcegroupref := make(map[string]interface{})
		resourcegroupref["moid"] = item.Moid
		resourcegroupref["object_type"] = item.ObjectType
		resourcegroupref["selector"] = item.Selector
		resourcegrouprefs = append(resourcegrouprefs, resourcegroupref)
	}
	return resourcegrouprefs
}
func flattenListResourcePerTypeCombinedSelector(p []*models.ResourcePerTypeCombinedSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourcepertypecombinedselectors []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resourcepertypecombinedselector := make(map[string]interface{})
		if len(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1) != 0 {
			j, err := json.Marshal(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1)
			if err == nil {
				resourcepertypecombinedselector["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		resourcepertypecombinedselector["class_id"] = item.ClassID
		resourcepertypecombinedselector["combined_selector"] = item.CombinedSelector
		resourcepertypecombinedselector["empty_filter"] = item.EmptyFilter
		resourcepertypecombinedselector["object_type"] = item.ObjectType
		resourcepertypecombinedselector["selector_object_type"] = item.SelectorObjectType
		resourcepertypecombinedselectors = append(resourcepertypecombinedselectors, resourcepertypecombinedselector)
	}
	return resourcepertypecombinedselectors
}
func flattenListResourceSelector(p []*models.ResourceSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourceselectors []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resourceselector := make(map[string]interface{})
		if len(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1) != 0 {
			j, err := json.Marshal(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1)
			if err == nil {
				resourceselector["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		resourceselector["class_id"] = item.ClassID
		resourceselector["object_type"] = item.ObjectType
		resourceselector["selector"] = item.Selector
		resourceselectors = append(resourceselectors, resourceselector)
	}
	return resourceselectors
}
func flattenListSdcardPartition(p []*models.SdcardPartition, d *schema.ResourceData) []map[string]interface{} {
	var sdcardpartitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sdcardpartition := make(map[string]interface{})
		if len(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1) != 0 {
			j, err := json.Marshal(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1)
			if err == nil {
				sdcardpartition["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		sdcardpartition["class_id"] = item.ClassID
		sdcardpartition["object_type"] = item.ObjectType
		sdcardpartition["type"] = item.Type
		sdcardpartition["virtual_drives"] = (func(p []*models.SdcardVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
			var sdcardvirtualdrives []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				sdcardvirtualdrive := make(map[string]interface{})
				if len(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1) != 0 {
					j, err := json.Marshal(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1)
					if err == nil {
						sdcardvirtualdrive["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				sdcardvirtualdrive["class_id"] = item.ClassID
				sdcardvirtualdrive["enable"] = item.Enable
				sdcardvirtualdrive["object_type"] = item.ObjectType
				sdcardvirtualdrives = append(sdcardvirtualdrives, sdcardvirtualdrive)
			}
			return sdcardvirtualdrives
		})(item.VirtualDrives, d)
		sdcardpartitions = append(sdcardpartitions, sdcardpartition)
	}
	return sdcardpartitions
}
func flattenListSdwanNetworkConfigurationType(p []*models.SdwanNetworkConfigurationType, d *schema.ResourceData) []map[string]interface{} {
	var sdwannetworkconfigurationtypes []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sdwannetworkconfigurationtype := make(map[string]interface{})
		if len(item.SdwanNetworkConfigurationTypeAO1P1.SdwanNetworkConfigurationTypeAO1P1) != 0 {
			j, err := json.Marshal(item.SdwanNetworkConfigurationTypeAO1P1.SdwanNetworkConfigurationTypeAO1P1)
			if err == nil {
				sdwannetworkconfigurationtype["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		sdwannetworkconfigurationtype["class_id"] = item.ClassID
		sdwannetworkconfigurationtype["network_type"] = item.NetworkType
		sdwannetworkconfigurationtype["object_type"] = item.ObjectType
		sdwannetworkconfigurationtype["port_group"] = item.PortGroup
		sdwannetworkconfigurationtype["vlan"] = item.Vlan
		sdwannetworkconfigurationtypes = append(sdwannetworkconfigurationtypes, sdwannetworkconfigurationtype)
	}
	return sdwannetworkconfigurationtypes
}
func flattenListSdwanProfileRef(p []*models.SdwanProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sdwanprofileref := make(map[string]interface{})
		sdwanprofileref["moid"] = item.Moid
		sdwanprofileref["object_type"] = item.ObjectType
		sdwanprofileref["selector"] = item.Selector
		sdwanprofilerefs = append(sdwanprofilerefs, sdwanprofileref)
	}
	return sdwanprofilerefs
}
func flattenListSdwanRouterNodeRef(p []*models.SdwanRouterNodeRef, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouternoderefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sdwanrouternoderef := make(map[string]interface{})
		sdwanrouternoderef["moid"] = item.Moid
		sdwanrouternoderef["object_type"] = item.ObjectType
		sdwanrouternoderef["selector"] = item.Selector
		sdwanrouternoderefs = append(sdwanrouternoderefs, sdwanrouternoderef)
	}
	return sdwanrouternoderefs
}
func flattenListSdwanTemplateInputsType(p []*models.SdwanTemplateInputsType, d *schema.ResourceData) []map[string]interface{} {
	var sdwantemplateinputstypes []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sdwantemplateinputstype := make(map[string]interface{})
		if len(item.SdwanTemplateInputsTypeAO1P1.SdwanTemplateInputsTypeAO1P1) != 0 {
			j, err := json.Marshal(item.SdwanTemplateInputsTypeAO1P1.SdwanTemplateInputsTypeAO1P1)
			if err == nil {
				sdwantemplateinputstype["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		sdwantemplateinputstype["class_id"] = item.ClassID
		sdwantemplateinputstype["editable"] = item.Editable
		sdwantemplateinputstype["key"] = item.Key
		sdwantemplateinputstype["object_type"] = item.ObjectType
		sdwantemplateinputstype["required"] = item.Required
		sdwantemplateinputstype["template"] = item.Template
		sdwantemplateinputstype["title"] = item.Title
		sdwantemplateinputstype["type"] = item.Type
		sdwantemplateinputstype["value"] = item.Value
		sdwantemplateinputstypes = append(sdwantemplateinputstypes, sdwantemplateinputstype)
	}
	return sdwantemplateinputstypes
}
func flattenListSecurityUnitRef(p []*models.SecurityUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var securityunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		securityunitref := make(map[string]interface{})
		securityunitref["moid"] = item.Moid
		securityunitref["object_type"] = item.ObjectType
		securityunitref["selector"] = item.Selector
		securityunitrefs = append(securityunitrefs, securityunitref)
	}
	return securityunitrefs
}
func flattenListServerConfigChangeDetailRef(p []*models.ServerConfigChangeDetailRef, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigchangedetailrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		serverconfigchangedetailref := make(map[string]interface{})
		serverconfigchangedetailref["moid"] = item.Moid
		serverconfigchangedetailref["object_type"] = item.ObjectType
		serverconfigchangedetailref["selector"] = item.Selector
		serverconfigchangedetailrefs = append(serverconfigchangedetailrefs, serverconfigchangedetailref)
	}
	return serverconfigchangedetailrefs
}
func flattenListServerConfigResultEntryRef(p []*models.ServerConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultentryrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		serverconfigresultentryref := make(map[string]interface{})
		serverconfigresultentryref["moid"] = item.Moid
		serverconfigresultentryref["object_type"] = item.ObjectType
		serverconfigresultentryref["selector"] = item.Selector
		serverconfigresultentryrefs = append(serverconfigresultentryrefs, serverconfigresultentryref)
	}
	return serverconfigresultentryrefs
}
func flattenListSnmpTrap(p []*models.SnmpTrap, d *schema.ResourceData) []map[string]interface{} {
	var snmptraps []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		snmptrap := make(map[string]interface{})
		if len(item.SnmpTrapAO1P1.SnmpTrapAO1P1) != 0 {
			j, err := json.Marshal(item.SnmpTrapAO1P1.SnmpTrapAO1P1)
			if err == nil {
				snmptrap["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		snmptrap["class_id"] = item.ClassID
		snmptrap["destination"] = item.Destination
		snmptrap["enabled"] = item.Enabled
		snmptrap["object_type"] = item.ObjectType
		snmptrap["port"] = item.Port
		snmptrap["type"] = item.Type
		snmptrap["user"] = item.User
		snmptrap["version"] = item.Version
		snmptraps = append(snmptraps, snmptrap)
	}
	return snmptraps
}
func flattenListSnmpUser(p []*models.SnmpUser, d *schema.ResourceData) []map[string]interface{} {
	var snmpusers []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		snmpuser := make(map[string]interface{})
		if len(item.SnmpUserAO1P1.SnmpUserAO1P1) != 0 {
			j, err := json.Marshal(item.SnmpUserAO1P1.SnmpUserAO1P1)
			if err == nil {
				snmpuser["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		auth_password_x := d.Get("snmp_users").([]interface{})
		auth_password_y := auth_password_x[0].(map[string]interface{})
		snmpuser["auth_password"] = auth_password_y["auth_password"]
		snmpuser["auth_type"] = item.AuthType
		snmpuser["class_id"] = item.ClassID
		snmpuser["is_auth_password_set"] = item.IsAuthPasswordSet
		snmpuser["is_privacy_password_set"] = item.IsPrivacyPasswordSet
		snmpuser["name"] = item.Name
		snmpuser["object_type"] = item.ObjectType
		privacy_password_x := d.Get("snmp_users").([]interface{})
		privacy_password_y := privacy_password_x[0].(map[string]interface{})
		snmpuser["privacy_password"] = privacy_password_y["privacy_password"]
		snmpuser["privacy_type"] = item.PrivacyType
		snmpuser["security_level"] = item.SecurityLevel
		snmpusers = append(snmpusers, snmpuser)
	}
	return snmpusers
}
func flattenListStorageBaseInitiator(p []*models.StorageBaseInitiator, d *schema.ResourceData) []map[string]interface{} {
	var storagebaseinitiators []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagebaseinitiator := make(map[string]interface{})
		if len(item.StorageBaseInitiatorAO1P1.StorageBaseInitiatorAO1P1) != 0 {
			j, err := json.Marshal(item.StorageBaseInitiatorAO1P1.StorageBaseInitiatorAO1P1)
			if err == nil {
				storagebaseinitiator["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		storagebaseinitiator["class_id"] = item.ClassID
		storagebaseinitiator["iqn"] = item.Iqn
		storagebaseinitiator["name"] = item.Name
		storagebaseinitiator["object_type"] = item.ObjectType
		storagebaseinitiator["type"] = item.Type
		storagebaseinitiator["wwn"] = item.Wwn
		storagebaseinitiators = append(storagebaseinitiators, storagebaseinitiator)
	}
	return storagebaseinitiators
}
func flattenListStorageControllerRef(p []*models.StorageControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagecontrollerref := make(map[string]interface{})
		storagecontrollerref["moid"] = item.Moid
		storagecontrollerref["object_type"] = item.ObjectType
		storagecontrollerref["selector"] = item.Selector
		storagecontrollerrefs = append(storagecontrollerrefs, storagecontrollerref)
	}
	return storagecontrollerrefs
}
func flattenListStorageDiskGroupPolicyRef(p []*models.StorageDiskGroupPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouppolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagediskgrouppolicyref := make(map[string]interface{})
		storagediskgrouppolicyref["moid"] = item.Moid
		storagediskgrouppolicyref["object_type"] = item.ObjectType
		storagediskgrouppolicyref["selector"] = item.Selector
		storagediskgrouppolicyrefs = append(storagediskgrouppolicyrefs, storagediskgrouppolicyref)
	}
	return storagediskgrouppolicyrefs
}
func flattenListStorageEnclosureDiskRef(p []*models.StorageEnclosureDiskRef, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurediskrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageenclosurediskref := make(map[string]interface{})
		storageenclosurediskref["moid"] = item.Moid
		storageenclosurediskref["object_type"] = item.ObjectType
		storageenclosurediskref["selector"] = item.Selector
		storageenclosurediskrefs = append(storageenclosurediskrefs, storageenclosurediskref)
	}
	return storageenclosurediskrefs
}
func flattenListStorageEnclosureDiskSlotEpRef(p []*models.StorageEnclosureDiskSlotEpRef, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosuredisksloteprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageenclosurediskslotepref := make(map[string]interface{})
		storageenclosurediskslotepref["moid"] = item.Moid
		storageenclosurediskslotepref["object_type"] = item.ObjectType
		storageenclosurediskslotepref["selector"] = item.Selector
		storageenclosuredisksloteprefs = append(storageenclosuredisksloteprefs, storageenclosurediskslotepref)
	}
	return storageenclosuredisksloteprefs
}
func flattenListStorageEnclosureRef(p []*models.StorageEnclosureRef, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageenclosureref := make(map[string]interface{})
		storageenclosureref["moid"] = item.Moid
		storageenclosureref["object_type"] = item.ObjectType
		storageenclosureref["selector"] = item.Selector
		storageenclosurerefs = append(storageenclosurerefs, storageenclosureref)
	}
	return storageenclosurerefs
}
func flattenListStorageFlexFlashControllerPropsRef(p []*models.StorageFlexFlashControllerPropsRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerpropsrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexflashcontrollerpropsref := make(map[string]interface{})
		storageflexflashcontrollerpropsref["moid"] = item.Moid
		storageflexflashcontrollerpropsref["object_type"] = item.ObjectType
		storageflexflashcontrollerpropsref["selector"] = item.Selector
		storageflexflashcontrollerpropsrefs = append(storageflexflashcontrollerpropsrefs, storageflexflashcontrollerpropsref)
	}
	return storageflexflashcontrollerpropsrefs
}
func flattenListStorageFlexFlashControllerRef(p []*models.StorageFlexFlashControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexflashcontrollerref := make(map[string]interface{})
		storageflexflashcontrollerref["moid"] = item.Moid
		storageflexflashcontrollerref["object_type"] = item.ObjectType
		storageflexflashcontrollerref["selector"] = item.Selector
		storageflexflashcontrollerrefs = append(storageflexflashcontrollerrefs, storageflexflashcontrollerref)
	}
	return storageflexflashcontrollerrefs
}
func flattenListStorageFlexFlashPhysicalDriveRef(p []*models.StorageFlexFlashPhysicalDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashphysicaldriverefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexflashphysicaldriveref := make(map[string]interface{})
		storageflexflashphysicaldriveref["moid"] = item.Moid
		storageflexflashphysicaldriveref["object_type"] = item.ObjectType
		storageflexflashphysicaldriveref["selector"] = item.Selector
		storageflexflashphysicaldriverefs = append(storageflexflashphysicaldriverefs, storageflexflashphysicaldriveref)
	}
	return storageflexflashphysicaldriverefs
}
func flattenListStorageFlexFlashVirtualDriveRef(p []*models.StorageFlexFlashVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashvirtualdriverefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexflashvirtualdriveref := make(map[string]interface{})
		storageflexflashvirtualdriveref["moid"] = item.Moid
		storageflexflashvirtualdriveref["object_type"] = item.ObjectType
		storageflexflashvirtualdriveref["selector"] = item.Selector
		storageflexflashvirtualdriverefs = append(storageflexflashvirtualdriverefs, storageflexflashvirtualdriveref)
	}
	return storageflexflashvirtualdriverefs
}
func flattenListStorageFlexUtilControllerRef(p []*models.StorageFlexUtilControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexutilcontrollerref := make(map[string]interface{})
		storageflexutilcontrollerref["moid"] = item.Moid
		storageflexutilcontrollerref["object_type"] = item.ObjectType
		storageflexutilcontrollerref["selector"] = item.Selector
		storageflexutilcontrollerrefs = append(storageflexutilcontrollerrefs, storageflexutilcontrollerref)
	}
	return storageflexutilcontrollerrefs
}
func flattenListStorageFlexUtilPhysicalDriveRef(p []*models.StorageFlexUtilPhysicalDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilphysicaldriverefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexutilphysicaldriveref := make(map[string]interface{})
		storageflexutilphysicaldriveref["moid"] = item.Moid
		storageflexutilphysicaldriveref["object_type"] = item.ObjectType
		storageflexutilphysicaldriveref["selector"] = item.Selector
		storageflexutilphysicaldriverefs = append(storageflexutilphysicaldriverefs, storageflexutilphysicaldriveref)
	}
	return storageflexutilphysicaldriverefs
}
func flattenListStorageFlexUtilVirtualDriveRef(p []*models.StorageFlexUtilVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilvirtualdriverefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexutilvirtualdriveref := make(map[string]interface{})
		storageflexutilvirtualdriveref["moid"] = item.Moid
		storageflexutilvirtualdriveref["object_type"] = item.ObjectType
		storageflexutilvirtualdriveref["selector"] = item.Selector
		storageflexutilvirtualdriverefs = append(storageflexutilvirtualdriverefs, storageflexutilvirtualdriveref)
	}
	return storageflexutilvirtualdriverefs
}
func flattenListStorageLocalDisk(p []*models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
	var storagelocaldisks []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagelocaldisk := make(map[string]interface{})
		if len(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1) != 0 {
			j, err := json.Marshal(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1)
			if err == nil {
				storagelocaldisk["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		storagelocaldisk["class_id"] = item.ClassID
		storagelocaldisk["object_type"] = item.ObjectType
		storagelocaldisk["slot_number"] = item.SlotNumber
		storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
	}
	return storagelocaldisks
}
func flattenListStoragePhysicalDiskExtensionRef(p []*models.StoragePhysicalDiskExtensionRef, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskextensionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagephysicaldiskextensionref := make(map[string]interface{})
		storagephysicaldiskextensionref["moid"] = item.Moid
		storagephysicaldiskextensionref["object_type"] = item.ObjectType
		storagephysicaldiskextensionref["selector"] = item.Selector
		storagephysicaldiskextensionrefs = append(storagephysicaldiskextensionrefs, storagephysicaldiskextensionref)
	}
	return storagephysicaldiskextensionrefs
}
func flattenListStoragePhysicalDiskRef(p []*models.StoragePhysicalDiskRef, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagephysicaldiskref := make(map[string]interface{})
		storagephysicaldiskref["moid"] = item.Moid
		storagephysicaldiskref["object_type"] = item.ObjectType
		storagephysicaldiskref["selector"] = item.Selector
		storagephysicaldiskrefs = append(storagephysicaldiskrefs, storagephysicaldiskref)
	}
	return storagephysicaldiskrefs
}
func flattenListStoragePhysicalDiskUsageRef(p []*models.StoragePhysicalDiskUsageRef, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskusagerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagephysicaldiskusageref := make(map[string]interface{})
		storagephysicaldiskusageref["moid"] = item.Moid
		storagephysicaldiskusageref["object_type"] = item.ObjectType
		storagephysicaldiskusageref["selector"] = item.Selector
		storagephysicaldiskusagerefs = append(storagephysicaldiskusagerefs, storagephysicaldiskusageref)
	}
	return storagephysicaldiskusagerefs
}
func flattenListStoragePureHostGroupRef(p []*models.StoragePureHostGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagepurehostgroupref := make(map[string]interface{})
		storagepurehostgroupref["moid"] = item.Moid
		storagepurehostgroupref["object_type"] = item.ObjectType
		storagepurehostgroupref["selector"] = item.Selector
		storagepurehostgrouprefs = append(storagepurehostgrouprefs, storagepurehostgroupref)
	}
	return storagepurehostgrouprefs
}
func flattenListStoragePureHostRef(p []*models.StoragePureHostRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagepurehostref := make(map[string]interface{})
		storagepurehostref["moid"] = item.Moid
		storagepurehostref["object_type"] = item.ObjectType
		storagepurehostref["selector"] = item.Selector
		storagepurehostrefs = append(storagepurehostrefs, storagepurehostref)
	}
	return storagepurehostrefs
}
func flattenListStoragePureReplicationBlackout(p []*models.StoragePureReplicationBlackout, d *schema.ResourceData) []map[string]interface{} {
	var storagepurereplicationblackouts []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagepurereplicationblackout := make(map[string]interface{})
		if len(item.StorageBaseReplicationBlackoutAO1P1.StorageBaseReplicationBlackoutAO1P1) != 0 {
			j, err := json.Marshal(item.StorageBaseReplicationBlackoutAO1P1.StorageBaseReplicationBlackoutAO1P1)
			if err == nil {
				storagepurereplicationblackout["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		storagepurereplicationblackout["class_id"] = item.ClassID
		storagepurereplicationblackout["object_type"] = item.ObjectType
		storagepurereplicationblackouts = append(storagepurereplicationblackouts, storagepurereplicationblackout)
	}
	return storagepurereplicationblackouts
}
func flattenListStoragePureVolumeRef(p []*models.StoragePureVolumeRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurevolumerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagepurevolumeref := make(map[string]interface{})
		storagepurevolumeref["moid"] = item.Moid
		storagepurevolumeref["object_type"] = item.ObjectType
		storagepurevolumeref["selector"] = item.Selector
		storagepurevolumerefs = append(storagepurevolumerefs, storagepurevolumeref)
	}
	return storagepurevolumerefs
}
func flattenListStorageSasExpanderRef(p []*models.StorageSasExpanderRef, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagesasexpanderref := make(map[string]interface{})
		storagesasexpanderref["moid"] = item.Moid
		storagesasexpanderref["object_type"] = item.ObjectType
		storagesasexpanderref["selector"] = item.Selector
		storagesasexpanderrefs = append(storagesasexpanderrefs, storagesasexpanderref)
	}
	return storagesasexpanderrefs
}
func flattenListStorageSasPortRef(p []*models.StorageSasPortRef, d *schema.ResourceData) []map[string]interface{} {
	var storagesasportrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagesasportref := make(map[string]interface{})
		storagesasportref["moid"] = item.Moid
		storagesasportref["object_type"] = item.ObjectType
		storagesasportref["selector"] = item.Selector
		storagesasportrefs = append(storagesasportrefs, storagesasportref)
	}
	return storagesasportrefs
}
func flattenListStorageSpanGroup(p []*models.StorageSpanGroup, d *schema.ResourceData) []map[string]interface{} {
	var storagespangroups []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagespangroup := make(map[string]interface{})
		if len(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1) != 0 {
			j, err := json.Marshal(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1)
			if err == nil {
				storagespangroup["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		storagespangroup["class_id"] = item.ClassID
		storagespangroup["disks"] = (func(p []*models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
			var storagelocaldisks []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				storagelocaldisk := make(map[string]interface{})
				if len(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1) != 0 {
					j, err := json.Marshal(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1)
					if err == nil {
						storagelocaldisk["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				storagelocaldisk["class_id"] = item.ClassID
				storagelocaldisk["object_type"] = item.ObjectType
				storagelocaldisk["slot_number"] = item.SlotNumber
				storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
			}
			return storagelocaldisks
		})(item.Disks, d)
		storagespangroup["object_type"] = item.ObjectType
		storagespangroups = append(storagespangroups, storagespangroup)
	}
	return storagespangroups
}
func flattenListStorageStoragePolicyRef(p []*models.StorageStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var storagestoragepolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagestoragepolicyref := make(map[string]interface{})
		storagestoragepolicyref["moid"] = item.Moid
		storagestoragepolicyref["object_type"] = item.ObjectType
		storagestoragepolicyref["selector"] = item.Selector
		storagestoragepolicyrefs = append(storagestoragepolicyrefs, storagestoragepolicyref)
	}
	return storagestoragepolicyrefs
}
func flattenListStorageVdMemberEpRef(p []*models.StorageVdMemberEpRef, d *schema.ResourceData) []map[string]interface{} {
	var storagevdmembereprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagevdmemberepref := make(map[string]interface{})
		storagevdmemberepref["moid"] = item.Moid
		storagevdmemberepref["object_type"] = item.ObjectType
		storagevdmemberepref["selector"] = item.Selector
		storagevdmembereprefs = append(storagevdmembereprefs, storagevdmemberepref)
	}
	return storagevdmembereprefs
}
func flattenListStorageVirtualDriveConfig(p []*models.StorageVirtualDriveConfig, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagevirtualdriveconfig := make(map[string]interface{})
		storagevirtualdriveconfig["access_policy"] = item.AccessPolicy
		if len(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1) != 0 {
			j, err := json.Marshal(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1)
			if err == nil {
				storagevirtualdriveconfig["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		storagevirtualdriveconfig["boot_drive"] = item.BootDrive
		storagevirtualdriveconfig["class_id"] = item.ClassID
		storagevirtualdriveconfig["disk_group_name"] = item.DiskGroupName
		storagevirtualdriveconfig["disk_group_policy"] = item.DiskGroupPolicy
		storagevirtualdriveconfig["drive_cache"] = item.DriveCache
		storagevirtualdriveconfig["expand_to_available"] = item.ExpandToAvailable
		storagevirtualdriveconfig["io_policy"] = item.IoPolicy
		storagevirtualdriveconfig["name"] = item.Name
		storagevirtualdriveconfig["object_type"] = item.ObjectType
		storagevirtualdriveconfig["read_policy"] = item.ReadPolicy
		storagevirtualdriveconfig["size"] = item.Size
		storagevirtualdriveconfig["write_policy"] = item.WritePolicy
		storagevirtualdriveconfigs = append(storagevirtualdriveconfigs, storagevirtualdriveconfig)
	}
	return storagevirtualdriveconfigs
}
func flattenListStorageVirtualDriveExtensionRef(p []*models.StorageVirtualDriveExtensionRef, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagevirtualdriveextensionref := make(map[string]interface{})
		storagevirtualdriveextensionref["moid"] = item.Moid
		storagevirtualdriveextensionref["object_type"] = item.ObjectType
		storagevirtualdriveextensionref["selector"] = item.Selector
		storagevirtualdriveextensionrefs = append(storagevirtualdriveextensionrefs, storagevirtualdriveextensionref)
	}
	return storagevirtualdriveextensionrefs
}
func flattenListStorageVirtualDriveRef(p []*models.StorageVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagevirtualdriveref := make(map[string]interface{})
		storagevirtualdriveref["moid"] = item.Moid
		storagevirtualdriveref["object_type"] = item.ObjectType
		storagevirtualdriveref["selector"] = item.Selector
		storagevirtualdriverefs = append(storagevirtualdriverefs, storagevirtualdriveref)
	}
	return storagevirtualdriverefs
}
func flattenListSyslogLocalClientBase(p []*models.SyslogLocalClientBase, d *schema.ResourceData) []map[string]interface{} {
	var sysloglocalclientbases []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sysloglocalclientbase := make(map[string]interface{})
		if len(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1) != 0 {
			j, err := json.Marshal(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1)
			if err == nil {
				sysloglocalclientbase["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		sysloglocalclientbase["class_id"] = item.ClassID
		sysloglocalclientbase["min_severity"] = item.MinSeverity
		sysloglocalclientbase["object_type"] = item.ObjectType
		sysloglocalclientbases = append(sysloglocalclientbases, sysloglocalclientbase)
	}
	return sysloglocalclientbases
}
func flattenListSyslogRemoteClientBase(p []*models.SyslogRemoteClientBase, d *schema.ResourceData) []map[string]interface{} {
	var syslogremoteclientbases []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		syslogremoteclientbase := make(map[string]interface{})
		if len(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1) != 0 {
			j, err := json.Marshal(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1)
			if err == nil {
				syslogremoteclientbase["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		syslogremoteclientbase["class_id"] = item.ClassID
		syslogremoteclientbase["enabled"] = item.Enabled
		syslogremoteclientbase["hostname"] = item.Hostname
		syslogremoteclientbase["min_severity"] = item.MinSeverity
		syslogremoteclientbase["object_type"] = item.ObjectType
		syslogremoteclientbase["port"] = item.Port
		syslogremoteclientbase["protocol"] = item.Protocol
		syslogremoteclientbases = append(syslogremoteclientbases, syslogremoteclientbase)
	}
	return syslogremoteclientbases
}
func flattenListTamAction(p []*models.TamAction, d *schema.ResourceData) []map[string]interface{} {
	var tamactions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		tamaction := make(map[string]interface{})
		if len(item.TamActionAO1P1.TamActionAO1P1) != 0 {
			j, err := json.Marshal(item.TamActionAO1P1.TamActionAO1P1)
			if err == nil {
				tamaction["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		tamaction["affected_object_type"] = item.AffectedObjectType
		tamaction["alert_type"] = item.AlertType
		tamaction["class_id"] = item.ClassID
		tamaction["identifiers"] = (func(p []*models.TamIdentifiers, d *schema.ResourceData) []map[string]interface{} {
			var tamidentifierss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				tamidentifiers := make(map[string]interface{})
				if len(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1) != 0 {
					j, err := json.Marshal(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1)
					if err == nil {
						tamidentifiers["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				tamidentifiers["class_id"] = item.ClassID
				tamidentifiers["name"] = item.Name
				tamidentifiers["object_type"] = item.ObjectType
				tamidentifiers["value"] = item.Value
				tamidentifierss = append(tamidentifierss, tamidentifiers)
			}
			return tamidentifierss
		})(item.Identifiers, d)
		tamaction["name"] = item.Name
		tamaction["object_type"] = item.ObjectType
		tamaction["operation_type"] = item.OperationType
		tamaction["queries"] = (func(p []*models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				tamqueryentry := make(map[string]interface{})
				if len(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1) != 0 {
					j, err := json.Marshal(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1)
					if err == nil {
						tamqueryentry["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				tamqueryentry["class_id"] = item.ClassID
				tamqueryentry["name"] = item.Name
				tamqueryentry["object_type"] = item.ObjectType
				tamqueryentry["priority"] = item.Priority
				tamqueryentry["query"] = item.Query
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.Queries, d)
		tamaction["type"] = item.Type
		tamactions = append(tamactions, tamaction)
	}
	return tamactions
}
func flattenListTamAPIDataSource(p []*models.TamAPIDataSource, d *schema.ResourceData) []map[string]interface{} {
	var tamapidatasources []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		tamapidatasource := make(map[string]interface{})
		if len(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1) != 0 {
			j, err := json.Marshal(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1)
			if err == nil {
				tamapidatasource["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		tamapidatasource["class_id"] = item.ClassID
		tamapidatasource["mo_type"] = item.MoType
		tamapidatasource["name"] = item.Name
		tamapidatasource["object_type"] = item.ObjectType
		tamapidatasource["queries"] = (func(p []*models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				tamqueryentry := make(map[string]interface{})
				if len(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1) != 0 {
					j, err := json.Marshal(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1)
					if err == nil {
						tamqueryentry["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				tamqueryentry["class_id"] = item.ClassID
				tamqueryentry["name"] = item.Name
				tamqueryentry["object_type"] = item.ObjectType
				tamqueryentry["priority"] = item.Priority
				tamqueryentry["query"] = item.Query
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.Queries, d)
		tamapidatasource["type"] = item.Type
		tamapidatasources = append(tamapidatasources, tamapidatasource)
	}
	return tamapidatasources
}
func flattenListUcsdConnectorPack(p []*models.UcsdConnectorPack, d *schema.ResourceData) []map[string]interface{} {
	var ucsdconnectorpacks []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		ucsdconnectorpack := make(map[string]interface{})
		if len(item.UcsdConnectorPackAO1P1.UcsdConnectorPackAO1P1) != 0 {
			j, err := json.Marshal(item.UcsdConnectorPackAO1P1.UcsdConnectorPackAO1P1)
			if err == nil {
				ucsdconnectorpack["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		ucsdconnectorpack["class_id"] = item.ClassID
		ucsdconnectorpack["connector_feature"] = item.ConnectorFeature
		ucsdconnectorpack["dependency_names"] = item.DependencyNames
		ucsdconnectorpack["downloaded_version"] = item.DownloadedVersion
		ucsdconnectorpack["name"] = item.Name
		ucsdconnectorpack["object_type"] = item.ObjectType
		ucsdconnectorpack["services"] = item.Services
		ucsdconnectorpack["state"] = item.State
		ucsdconnectorpack["version"] = item.Version
		ucsdconnectorpacks = append(ucsdconnectorpacks, ucsdconnectorpack)
	}
	return ucsdconnectorpacks
}
func flattenListVirtualizationVmwareDatastoreRef(p []*models.VirtualizationVmwareDatastoreRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatastorerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		virtualizationvmwaredatastoreref := make(map[string]interface{})
		virtualizationvmwaredatastoreref["moid"] = item.Moid
		virtualizationvmwaredatastoreref["object_type"] = item.ObjectType
		virtualizationvmwaredatastoreref["selector"] = item.Selector
		virtualizationvmwaredatastorerefs = append(virtualizationvmwaredatastorerefs, virtualizationvmwaredatastoreref)
	}
	return virtualizationvmwaredatastorerefs
}
func flattenListVirtualizationVmwareHostRef(p []*models.VirtualizationVmwareHostRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		virtualizationvmwarehostref := make(map[string]interface{})
		virtualizationvmwarehostref["moid"] = item.Moid
		virtualizationvmwarehostref["object_type"] = item.ObjectType
		virtualizationvmwarehostref["selector"] = item.Selector
		virtualizationvmwarehostrefs = append(virtualizationvmwarehostrefs, virtualizationvmwarehostref)
	}
	return virtualizationvmwarehostrefs
}
func flattenListVmediaMapping(p []*models.VmediaMapping, d *schema.ResourceData) []map[string]interface{} {
	var vmediamappings []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		vmediamapping := make(map[string]interface{})
		if len(item.VmediaMappingAO1P1.VmediaMappingAO1P1) != 0 {
			j, err := json.Marshal(item.VmediaMappingAO1P1.VmediaMappingAO1P1)
			if err == nil {
				vmediamapping["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		vmediamapping["authentication_protocol"] = item.AuthenticationProtocol
		vmediamapping["class_id"] = item.ClassID
		vmediamapping["device_type"] = item.DeviceType
		vmediamapping["host_name"] = item.HostName
		vmediamapping["is_password_set"] = item.IsPasswordSet
		vmediamapping["mount_options"] = item.MountOptions
		vmediamapping["mount_protocol"] = item.MountProtocol
		vmediamapping["object_type"] = item.ObjectType
		vmediamapping["password"] = item.Password
		vmediamapping["remote_file"] = item.RemoteFile
		vmediamapping["remote_path"] = item.RemotePath
		vmediamapping["username"] = item.Username
		vmediamapping["volume_name"] = item.VolumeName
		vmediamappings = append(vmediamappings, vmediamapping)
	}
	return vmediamappings
}
func flattenListVnicEthIfRef(p []*models.VnicEthIfRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicethifrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		vnicethifref := make(map[string]interface{})
		vnicethifref["moid"] = item.Moid
		vnicethifref["object_type"] = item.ObjectType
		vnicethifref["selector"] = item.Selector
		vnicethifrefs = append(vnicethifrefs, vnicethifref)
	}
	return vnicethifrefs
}
func flattenListVnicFcIfRef(p []*models.VnicFcIfRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcifrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		vnicfcifref := make(map[string]interface{})
		vnicfcifref["moid"] = item.Moid
		vnicfcifref["object_type"] = item.ObjectType
		vnicfcifref["selector"] = item.Selector
		vnicfcifrefs = append(vnicfcifrefs, vnicfcifref)
	}
	return vnicfcifrefs
}
func flattenListWorkflowAPI(p []*models.WorkflowAPI, d *schema.ResourceData) []map[string]interface{} {
	var workflowapis []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowapi := make(map[string]interface{})
		if len(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1) != 0 {
			j, err := json.Marshal(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1)
			if err == nil {
				workflowapi["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		workflowapi["body"] = item.Body
		workflowapi["class_id"] = item.ClassID
		workflowapi["content_type"] = item.ContentType
		workflowapi["name"] = item.Name
		workflowapi["object_type"] = item.ObjectType
		workflowapi["outcomes"] = item.Outcomes
		workflowapi["response_spec"] = (func(p *models.ContentGrammar, d *schema.ResourceData) []map[string]interface{} {
			var contentgrammars []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			contentgrammar := make(map[string]interface{})
			delete(item.ContentGrammarAO1P1.ContentGrammarAO1P1, "ObjectType")
			if len(item.ContentGrammarAO1P1.ContentGrammarAO1P1) != 0 {
				j, err := json.Marshal(item.ContentGrammarAO1P1.ContentGrammarAO1P1)
				if err == nil {
					contentgrammar["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			contentgrammar["class_id"] = item.ClassID
			contentgrammar["error_parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var contentbaseparameters []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					contentbaseparameter := make(map[string]interface{})
					contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
					if len(item.ContentBaseParameterAO1P1.ContentBaseParameterAO1P1) != 0 {
						j, err := json.Marshal(item.ContentBaseParameterAO1P1.ContentBaseParameterAO1P1)
						if err == nil {
							contentbaseparameter["additional_properties"] = string(j)
						} else {
							log.Printf("Error occured while flattening and json parsing: %s", err)
						}
					}
					contentbaseparameter["class_id"] = item.ClassID
					contentbaseparameter["complex_type"] = item.ComplexType
					contentbaseparameter["item_type"] = item.ItemType
					contentbaseparameter["name"] = item.Name
					contentbaseparameter["object_type"] = item.ObjectType
					contentbaseparameter["path"] = item.Path
					contentbaseparameter["type"] = item.Type
					contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
				}
				return contentbaseparameters
			})(item.ErrorParameters, d)
			contentgrammar["object_type"] = item.ObjectType
			contentgrammar["parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var contentbaseparameters []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					contentbaseparameter := make(map[string]interface{})
					contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
					if len(item.ContentBaseParameterAO1P1.ContentBaseParameterAO1P1) != 0 {
						j, err := json.Marshal(item.ContentBaseParameterAO1P1.ContentBaseParameterAO1P1)
						if err == nil {
							contentbaseparameter["additional_properties"] = string(j)
						} else {
							log.Printf("Error occured while flattening and json parsing: %s", err)
						}
					}
					contentbaseparameter["class_id"] = item.ClassID
					contentbaseparameter["complex_type"] = item.ComplexType
					contentbaseparameter["item_type"] = item.ItemType
					contentbaseparameter["name"] = item.Name
					contentbaseparameter["object_type"] = item.ObjectType
					contentbaseparameter["path"] = item.Path
					contentbaseparameter["type"] = item.Type
					contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
				}
				return contentbaseparameters
			})(item.Parameters, d)
			contentgrammar["types"] = (func(p []*models.ContentComplexType, d *schema.ResourceData) []map[string]interface{} {
				var contentcomplextypes []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					contentcomplextype := make(map[string]interface{})
					if len(item.ContentComplexTypeAO1P1.ContentComplexTypeAO1P1) != 0 {
						j, err := json.Marshal(item.ContentComplexTypeAO1P1.ContentComplexTypeAO1P1)
						if err == nil {
							contentcomplextype["additional_properties"] = string(j)
						} else {
							log.Printf("Error occured while flattening and json parsing: %s", err)
						}
					}
					contentcomplextype["class_id"] = item.ClassID
					contentcomplextype["name"] = item.Name
					contentcomplextype["object_type"] = item.ObjectType
					contentcomplextype["parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
						var contentbaseparameters []map[string]interface{}
						if p == nil {
							return nil
						}
						for _, item := range p {
							item := *item
							contentbaseparameter := make(map[string]interface{})
							contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
							if len(item.ContentBaseParameterAO1P1.ContentBaseParameterAO1P1) != 0 {
								j, err := json.Marshal(item.ContentBaseParameterAO1P1.ContentBaseParameterAO1P1)
								if err == nil {
									contentbaseparameter["additional_properties"] = string(j)
								} else {
									log.Printf("Error occured while flattening and json parsing: %s", err)
								}
							}
							contentbaseparameter["class_id"] = item.ClassID
							contentbaseparameter["complex_type"] = item.ComplexType
							contentbaseparameter["item_type"] = item.ItemType
							contentbaseparameter["name"] = item.Name
							contentbaseparameter["object_type"] = item.ObjectType
							contentbaseparameter["path"] = item.Path
							contentbaseparameter["type"] = item.Type
							contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
						}
						return contentbaseparameters
					})(item.Parameters, d)
					contentcomplextypes = append(contentcomplextypes, contentcomplextype)
				}
				return contentcomplextypes
			})(item.Types, d)

			contentgrammars = append(contentgrammars, contentgrammar)
			return contentgrammars
		})(item.ResponseSpec, d)
		workflowapi["skip_on_condition"] = item.SkipOnCondition
		workflowapi["timeout"] = item.Timeout
		workflowapis = append(workflowapis, workflowapi)
	}
	return workflowapis
}
func flattenListWorkflowBaseDataType(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
	var workflowbasedatatypes []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowbasedatatype := make(map[string]interface{})
		if len(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1) != 0 {
			j, err := json.Marshal(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1)
			if err == nil {
				workflowbasedatatype["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		workflowbasedatatype["class_id"] = item.ClassID
		workflowbasedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
			var workflowdefaultvalues []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			workflowdefaultvalue := make(map[string]interface{})
			delete(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1, "ObjectType")
			if len(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1) != 0 {
				j, err := json.Marshal(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1)
				if err == nil {
					workflowdefaultvalue["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			workflowdefaultvalue["class_id"] = item.ClassID
			workflowdefaultvalue["object_type"] = item.ObjectType
			workflowdefaultvalue["override"] = item.Override
			workflowdefaultvalue["value"] = item.Value

			workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
			return workflowdefaultvalues
		})(item.Default, d)
		workflowbasedatatype["description"] = item.Description
		workflowbasedatatype["label"] = item.Label
		workflowbasedatatype["name"] = item.Name
		workflowbasedatatype["object_type"] = item.ObjectType
		workflowbasedatatype["required"] = item.Required
		workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
	}
	return workflowbasedatatypes
}
func flattenListWorkflowDynamicWorkflowActionTaskList(p []*models.WorkflowDynamicWorkflowActionTaskList, d *schema.ResourceData) []map[string]interface{} {
	var workflowdynamicworkflowactiontasklists []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowdynamicworkflowactiontasklist := make(map[string]interface{})
		workflowdynamicworkflowactiontasklist["action"] = item.Action
		if len(item.WorkflowDynamicWorkflowActionTaskListAO1P1.WorkflowDynamicWorkflowActionTaskListAO1P1) != 0 {
			j, err := json.Marshal(item.WorkflowDynamicWorkflowActionTaskListAO1P1.WorkflowDynamicWorkflowActionTaskListAO1P1)
			if err == nil {
				workflowdynamicworkflowactiontasklist["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		workflowdynamicworkflowactiontasklist["class_id"] = item.ClassID
		workflowdynamicworkflowactiontasklist["object_type"] = item.ObjectType
		workflowdynamicworkflowactiontasklist["tasks"] = item.Tasks
		workflowdynamicworkflowactiontasklists = append(workflowdynamicworkflowactiontasklists, workflowdynamicworkflowactiontasklist)
	}
	return workflowdynamicworkflowactiontasklists
}
func flattenListWorkflowMessage(p []*models.WorkflowMessage, d *schema.ResourceData) []map[string]interface{} {
	var workflowmessages []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowmessage := make(map[string]interface{})
		if len(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1) != 0 {
			j, err := json.Marshal(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1)
			if err == nil {
				workflowmessage["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		workflowmessage["class_id"] = item.ClassID
		workflowmessage["message"] = item.Message
		workflowmessage["object_type"] = item.ObjectType
		workflowmessage["severity"] = item.Severity
		workflowmessages = append(workflowmessages, workflowmessage)
	}
	return workflowmessages
}
func flattenListWorkflowTaskDefinitionRef(p []*models.WorkflowTaskDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowtaskdefinitionref := make(map[string]interface{})
		workflowtaskdefinitionref["moid"] = item.Moid
		workflowtaskdefinitionref["object_type"] = item.ObjectType
		workflowtaskdefinitionref["selector"] = item.Selector
		workflowtaskdefinitionrefs = append(workflowtaskdefinitionrefs, workflowtaskdefinitionref)
	}
	return workflowtaskdefinitionrefs
}
func flattenListWorkflowTaskInfoRef(p []*models.WorkflowTaskInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowtaskinforef := make(map[string]interface{})
		workflowtaskinforef["moid"] = item.Moid
		workflowtaskinforef["object_type"] = item.ObjectType
		workflowtaskinforef["selector"] = item.Selector
		workflowtaskinforefs = append(workflowtaskinforefs, workflowtaskinforef)
	}
	return workflowtaskinforefs
}
func flattenListWorkflowTaskRetryInfo(p []*models.WorkflowTaskRetryInfo, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskretryinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowtaskretryinfo := make(map[string]interface{})
		if len(item.WorkflowTaskRetryInfoAO1P1.WorkflowTaskRetryInfoAO1P1) != 0 {
			j, err := json.Marshal(item.WorkflowTaskRetryInfoAO1P1.WorkflowTaskRetryInfoAO1P1)
			if err == nil {
				workflowtaskretryinfo["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		workflowtaskretryinfo["class_id"] = item.ClassID
		workflowtaskretryinfo["object_type"] = item.ObjectType
		workflowtaskretryinfo["status"] = item.Status
		workflowtaskretryinfo["task_inst_id"] = item.TaskInstID
		workflowtaskretryinfos = append(workflowtaskretryinfos, workflowtaskretryinfo)
	}
	return workflowtaskretryinfos
}
func flattenListWorkflowWorkflowInfoRef(p []*models.WorkflowWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowworkflowinforef := make(map[string]interface{})
		workflowworkflowinforef["moid"] = item.Moid
		workflowworkflowinforef["object_type"] = item.ObjectType
		workflowworkflowinforef["selector"] = item.Selector
		workflowworkflowinforefs = append(workflowworkflowinforefs, workflowworkflowinforef)
	}
	return workflowworkflowinforefs
}
func flattenListWorkflowWorkflowTask(p []*models.WorkflowWorkflowTask, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowtasks []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowworkflowtask := make(map[string]interface{})
		if len(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1) != 0 {
			j, err := json.Marshal(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1)
			if err == nil {
				workflowworkflowtask["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		workflowworkflowtask["class_id"] = item.ClassID
		workflowworkflowtask["description"] = item.Description
		workflowworkflowtask["label"] = item.Label
		workflowworkflowtask["name"] = item.Name
		workflowworkflowtask["object_type"] = item.ObjectType
		workflowworkflowtasks = append(workflowworkflowtasks, workflowworkflowtask)
	}
	return workflowworkflowtasks
}
func flattenListX509Certificate(p []*models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var x509certificates []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		x509certificate := make(map[string]interface{})
		if len(item.X509CertificateAO1P1.X509CertificateAO1P1) != 0 {
			j, err := json.Marshal(item.X509CertificateAO1P1.X509CertificateAO1P1)
			if err == nil {
				x509certificate["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		x509certificate["class_id"] = item.ClassID
		x509certificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			pkixdistinguishedname := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					pkixdistinguishedname["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			pkixdistinguishedname["class_id"] = item.ClassID
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.Issuer, d)
		x509certificate["object_type"] = item.ObjectType
		x509certificate["pem_certificate"] = item.PemCertificate
		x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
		x509certificate["signature_algorithm"] = item.SignatureAlgorithm
		x509certificate["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			pkixdistinguishedname := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					pkixdistinguishedname["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			pkixdistinguishedname["class_id"] = item.ClassID
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.Subject, d)
		x509certificates = append(x509certificates, x509certificate)
	}
	return x509certificates
}
func flattenMapAdapterUnitRef(p *models.AdapterUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	adapterunitref := make(map[string]interface{})
	adapterunitref["moid"] = item.Moid
	adapterunitref["object_type"] = item.ObjectType
	adapterunitref["selector"] = item.Selector

	adapterunitrefs = append(adapterunitrefs, adapterunitref)
	return adapterunitrefs
}
func flattenMapApplianceDataExportPolicyRef(p *models.ApplianceDataExportPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	appliancedataexportpolicyref := make(map[string]interface{})
	appliancedataexportpolicyref["moid"] = item.Moid
	appliancedataexportpolicyref["object_type"] = item.ObjectType
	appliancedataexportpolicyref["selector"] = item.Selector

	appliancedataexportpolicyrefs = append(appliancedataexportpolicyrefs, appliancedataexportpolicyref)
	return appliancedataexportpolicyrefs
}
func flattenMapApplianceImageBundleRef(p *models.ApplianceImageBundleRef, d *schema.ResourceData) []map[string]interface{} {
	var applianceimagebundlerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	applianceimagebundleref := make(map[string]interface{})
	applianceimagebundleref["moid"] = item.Moid
	applianceimagebundleref["object_type"] = item.ObjectType
	applianceimagebundleref["selector"] = item.Selector

	applianceimagebundlerefs = append(applianceimagebundlerefs, applianceimagebundleref)
	return applianceimagebundlerefs
}
func flattenMapAssetClusterMemberRef(p *models.AssetClusterMemberRef, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetclustermemberref := make(map[string]interface{})
	assetclustermemberref["moid"] = item.Moid
	assetclustermemberref["object_type"] = item.ObjectType
	assetclustermemberref["selector"] = item.Selector

	assetclustermemberrefs = append(assetclustermemberrefs, assetclustermemberref)
	return assetclustermemberrefs
}
func flattenMapAssetContractInformation(p *models.AssetContractInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetcontractinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetcontractinformation := make(map[string]interface{})
	delete(item.AssetContractInformationAO1P1.AssetContractInformationAO1P1, "ObjectType")
	if len(item.AssetContractInformationAO1P1.AssetContractInformationAO1P1) != 0 {
		j, err := json.Marshal(item.AssetContractInformationAO1P1.AssetContractInformationAO1P1)
		if err == nil {
			assetcontractinformation["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	assetcontractinformation["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetaddressinformation := make(map[string]interface{})
		delete(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1, "ObjectType")
		if len(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1) != 0 {
			j, err := json.Marshal(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1)
			if err == nil {
				assetaddressinformation["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassID
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.BillTo, d)
	assetcontractinformation["bill_to_global_ultimate"] = (func(p *models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {
		var assetglobalultimates []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetglobalultimate := make(map[string]interface{})
		delete(item.AssetGlobalUltimateAO1P1.AssetGlobalUltimateAO1P1, "ObjectType")
		if len(item.AssetGlobalUltimateAO1P1.AssetGlobalUltimateAO1P1) != 0 {
			j, err := json.Marshal(item.AssetGlobalUltimateAO1P1.AssetGlobalUltimateAO1P1)
			if err == nil {
				assetglobalultimate["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		assetglobalultimate["class_id"] = item.ClassID
		assetglobalultimate["id"] = item.ID
		assetglobalultimate["name"] = item.Name
		assetglobalultimate["object_type"] = item.ObjectType

		assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
		return assetglobalultimates
	})(item.BillToGlobalUltimate, d)
	assetcontractinformation["class_id"] = item.ClassID
	assetcontractinformation["contract_number"] = item.ContractNumber
	assetcontractinformation["line_status"] = item.LineStatus
	assetcontractinformation["object_type"] = item.ObjectType

	assetcontractinformations = append(assetcontractinformations, assetcontractinformation)
	return assetcontractinformations
}
func flattenMapAssetCustomerInformation(p *models.AssetCustomerInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetcustomerinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetcustomerinformation := make(map[string]interface{})
	delete(item.AssetCustomerInformationAO1P1.AssetCustomerInformationAO1P1, "ObjectType")
	if len(item.AssetCustomerInformationAO1P1.AssetCustomerInformationAO1P1) != 0 {
		j, err := json.Marshal(item.AssetCustomerInformationAO1P1.AssetCustomerInformationAO1P1)
		if err == nil {
			assetcustomerinformation["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	assetcustomerinformation["address"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetaddressinformation := make(map[string]interface{})
		delete(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1, "ObjectType")
		if len(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1) != 0 {
			j, err := json.Marshal(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1)
			if err == nil {
				assetaddressinformation["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassID
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.Address, d)
	assetcustomerinformation["class_id"] = item.ClassID
	assetcustomerinformation["id"] = item.ID
	assetcustomerinformation["name"] = item.Name
	assetcustomerinformation["object_type"] = item.ObjectType

	assetcustomerinformations = append(assetcustomerinformations, assetcustomerinformation)
	return assetcustomerinformations
}
func flattenMapAssetDeviceClaimRef(p *models.AssetDeviceClaimRef, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceclaimrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetdeviceclaimref := make(map[string]interface{})
	assetdeviceclaimref["moid"] = item.Moid
	assetdeviceclaimref["object_type"] = item.ObjectType
	assetdeviceclaimref["selector"] = item.Selector

	assetdeviceclaimrefs = append(assetdeviceclaimrefs, assetdeviceclaimref)
	return assetdeviceclaimrefs
}
func flattenMapAssetDeviceConfigurationRef(p *models.AssetDeviceConfigurationRef, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconfigurationrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetdeviceconfigurationref := make(map[string]interface{})
	assetdeviceconfigurationref["moid"] = item.Moid
	assetdeviceconfigurationref["object_type"] = item.ObjectType
	assetdeviceconfigurationref["selector"] = item.Selector

	assetdeviceconfigurationrefs = append(assetdeviceconfigurationrefs, assetdeviceconfigurationref)
	return assetdeviceconfigurationrefs
}
func flattenMapAssetDeviceConnectionRef(p *models.AssetDeviceConnectionRef, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconnectionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetdeviceconnectionref := make(map[string]interface{})
	assetdeviceconnectionref["moid"] = item.Moid
	assetdeviceconnectionref["object_type"] = item.ObjectType
	assetdeviceconnectionref["selector"] = item.Selector

	assetdeviceconnectionrefs = append(assetdeviceconnectionrefs, assetdeviceconnectionref)
	return assetdeviceconnectionrefs
}
func flattenMapAssetDeviceRegistrationRef(p *models.AssetDeviceRegistrationRef, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceregistrationrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetdeviceregistrationref := make(map[string]interface{})
	assetdeviceregistrationref["moid"] = item.Moid
	assetdeviceregistrationref["object_type"] = item.ObjectType
	assetdeviceregistrationref["selector"] = item.Selector

	assetdeviceregistrationrefs = append(assetdeviceregistrationrefs, assetdeviceregistrationref)
	return assetdeviceregistrationrefs
}
func flattenMapAssetGlobalUltimate(p *models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {
	var assetglobalultimates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetglobalultimate := make(map[string]interface{})
	delete(item.AssetGlobalUltimateAO1P1.AssetGlobalUltimateAO1P1, "ObjectType")
	if len(item.AssetGlobalUltimateAO1P1.AssetGlobalUltimateAO1P1) != 0 {
		j, err := json.Marshal(item.AssetGlobalUltimateAO1P1.AssetGlobalUltimateAO1P1)
		if err == nil {
			assetglobalultimate["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	assetglobalultimate["class_id"] = item.ClassID
	assetglobalultimate["id"] = item.ID
	assetglobalultimate["name"] = item.Name
	assetglobalultimate["object_type"] = item.ObjectType

	assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
	return assetglobalultimates
}
func flattenMapAssetManagedDeviceStatus(p *models.AssetManagedDeviceStatus, d *schema.ResourceData) []map[string]interface{} {
	var assetmanageddevicestatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetmanageddevicestatus := make(map[string]interface{})
	delete(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1, "ObjectType")
	if len(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1) != 0 {
		j, err := json.Marshal(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1)
		if err == nil {
			assetmanageddevicestatus["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	assetmanageddevicestatus["class_id"] = item.ClassID
	assetmanageddevicestatus["cloud_port"] = item.CloudPort
	assetmanageddevicestatus["connection_failure_reason"] = item.ConnectionFailureReason
	assetmanageddevicestatus["connection_status"] = item.ConnectionStatus
	assetmanageddevicestatus["error_code"] = item.ErrorCode
	assetmanageddevicestatus["error_reason"] = item.ErrorReason
	assetmanageddevicestatus["object_type"] = item.ObjectType
	assetmanageddevicestatus["process_id"] = item.ProcessID
	assetmanageddevicestatus["server_port"] = item.ServerPort
	assetmanageddevicestatus["state"] = item.State

	assetmanageddevicestatuss = append(assetmanageddevicestatuss, assetmanageddevicestatus)
	return assetmanageddevicestatuss
}
func flattenMapAssetParentConnectionSignature(p *models.AssetParentConnectionSignature, d *schema.ResourceData) []map[string]interface{} {
	var assetparentconnectionsignatures []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetparentconnectionsignature := make(map[string]interface{})
	delete(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1, "ObjectType")
	if len(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1) != 0 {
		j, err := json.Marshal(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1)
		if err == nil {
			assetparentconnectionsignature["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	assetparentconnectionsignature["class_id"] = item.ClassID
	assetparentconnectionsignature["device_id"] = item.DeviceID
	assetparentconnectionsignature["node_id"] = item.NodeID
	assetparentconnectionsignature["object_type"] = item.ObjectType
	assetparentconnectionsignature["signature"] = item.Signature

	assetparentconnectionsignatures = append(assetparentconnectionsignatures, assetparentconnectionsignature)
	return assetparentconnectionsignatures
}
func flattenMapAssetProductInformation(p *models.AssetProductInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetproductinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetproductinformation := make(map[string]interface{})
	delete(item.AssetProductInformationAO1P1.AssetProductInformationAO1P1, "ObjectType")
	if len(item.AssetProductInformationAO1P1.AssetProductInformationAO1P1) != 0 {
		j, err := json.Marshal(item.AssetProductInformationAO1P1.AssetProductInformationAO1P1)
		if err == nil {
			assetproductinformation["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	assetproductinformation["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetaddressinformation := make(map[string]interface{})
		delete(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1, "ObjectType")
		if len(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1) != 0 {
			j, err := json.Marshal(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1)
			if err == nil {
				assetaddressinformation["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassID
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.BillTo, d)
	assetproductinformation["class_id"] = item.ClassID
	assetproductinformation["description"] = item.Description
	assetproductinformation["family"] = item.Family
	assetproductinformation["group"] = item.Group
	assetproductinformation["number"] = item.Number
	assetproductinformation["object_type"] = item.ObjectType
	assetproductinformation["ship_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetaddressinformation := make(map[string]interface{})
		delete(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1, "ObjectType")
		if len(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1) != 0 {
			j, err := json.Marshal(item.AssetAddressInformationAO1P1.AssetAddressInformationAO1P1)
			if err == nil {
				assetaddressinformation["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassID
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.ShipTo, d)
	assetproductinformation["sub_type"] = item.SubType

	assetproductinformations = append(assetproductinformations, assetproductinformation)
	return assetproductinformations
}
func flattenMapAssetSecurityTokenRef(p *models.AssetSecurityTokenRef, d *schema.ResourceData) []map[string]interface{} {
	var assetsecuritytokenrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetsecuritytokenref := make(map[string]interface{})
	assetsecuritytokenref["moid"] = item.Moid
	assetsecuritytokenref["object_type"] = item.ObjectType
	assetsecuritytokenref["selector"] = item.Selector

	assetsecuritytokenrefs = append(assetsecuritytokenrefs, assetsecuritytokenref)
	return assetsecuritytokenrefs
}
func flattenMapAssetSudiInfo(p *models.AssetSudiInfo, d *schema.ResourceData) []map[string]interface{} {
	var assetsudiinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetsudiinfo := make(map[string]interface{})
	delete(item.AssetSudiInfoAO1P1.AssetSudiInfoAO1P1, "ObjectType")
	if len(item.AssetSudiInfoAO1P1.AssetSudiInfoAO1P1) != 0 {
		j, err := json.Marshal(item.AssetSudiInfoAO1P1.AssetSudiInfoAO1P1)
		if err == nil {
			assetsudiinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	assetsudiinfo["class_id"] = item.ClassID
	assetsudiinfo["object_type"] = item.ObjectType
	assetsudiinfo["pid"] = item.Pid
	assetsudiinfo["serial_number"] = item.SerialNumber
	assetsudiinfo["signature"] = item.Signature
	assetsudiinfo["status"] = item.Status
	assetsudiinfo["sudi_certificate"] = (func(p *models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
		var x509certificates []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		x509certificate := make(map[string]interface{})
		delete(item.X509CertificateAO1P1.X509CertificateAO1P1, "ObjectType")
		if len(item.X509CertificateAO1P1.X509CertificateAO1P1) != 0 {
			j, err := json.Marshal(item.X509CertificateAO1P1.X509CertificateAO1P1)
			if err == nil {
				x509certificate["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		x509certificate["class_id"] = item.ClassID
		x509certificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			pkixdistinguishedname := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					pkixdistinguishedname["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			pkixdistinguishedname["class_id"] = item.ClassID
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.Issuer, d)
		x509certificate["object_type"] = item.ObjectType
		x509certificate["pem_certificate"] = item.PemCertificate
		x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
		x509certificate["signature_algorithm"] = item.SignatureAlgorithm
		x509certificate["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			pkixdistinguishedname := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					pkixdistinguishedname["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			pkixdistinguishedname["class_id"] = item.ClassID
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.Subject, d)

		x509certificates = append(x509certificates, x509certificate)
		return x509certificates
	})(item.SudiCertificate, d)

	assetsudiinfos = append(assetsudiinfos, assetsudiinfo)
	return assetsudiinfos
}
func flattenMapBiosBootModeRef(p *models.BiosBootModeRef, d *schema.ResourceData) []map[string]interface{} {
	var biosbootmoderefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	biosbootmoderef := make(map[string]interface{})
	biosbootmoderef["moid"] = item.Moid
	biosbootmoderef["object_type"] = item.ObjectType
	biosbootmoderef["selector"] = item.Selector

	biosbootmoderefs = append(biosbootmoderefs, biosbootmoderef)
	return biosbootmoderefs
}
func flattenMapBiosUnitRef(p *models.BiosUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	biosunitref := make(map[string]interface{})
	biosunitref["moid"] = item.Moid
	biosunitref["object_type"] = item.ObjectType
	biosunitref["selector"] = item.Selector

	biosunitrefs = append(biosunitrefs, biosunitref)
	return biosunitrefs
}
func flattenMapBootDeviceBootModeRef(p *models.BootDeviceBootModeRef, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebootmoderefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	bootdevicebootmoderef := make(map[string]interface{})
	bootdevicebootmoderef["moid"] = item.Moid
	bootdevicebootmoderef["object_type"] = item.ObjectType
	bootdevicebootmoderef["selector"] = item.Selector

	bootdevicebootmoderefs = append(bootdevicebootmoderefs, bootdevicebootmoderef)
	return bootdevicebootmoderefs
}
func flattenMapCommCredential(p *models.CommCredential, d *schema.ResourceData) []map[string]interface{} {
	var commcredentials []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	commcredential := make(map[string]interface{})
	delete(item.CommCredentialAO1P1.CommCredentialAO1P1, "ObjectType")
	if len(item.CommCredentialAO1P1.CommCredentialAO1P1) != 0 {
		j, err := json.Marshal(item.CommCredentialAO1P1.CommCredentialAO1P1)
		if err == nil {
			commcredential["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	commcredential["class_id"] = item.ClassID
	commcredential["is_password_set"] = item.IsPasswordSet
	commcredential["object_type"] = item.ObjectType
	commcredential["password"] = item.Password
	commcredential["username"] = item.Username

	commcredentials = append(commcredentials, commcredential)
	return commcredentials
}
func flattenMapCommIPV4Interface(p *models.CommIPV4Interface, d *schema.ResourceData) []map[string]interface{} {
	var commipv4interfaces []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	commipv4interface := make(map[string]interface{})
	delete(item.CommIPV4InterfaceAO1P1.CommIPV4InterfaceAO1P1, "ObjectType")
	if len(item.CommIPV4InterfaceAO1P1.CommIPV4InterfaceAO1P1) != 0 {
		j, err := json.Marshal(item.CommIPV4InterfaceAO1P1.CommIPV4InterfaceAO1P1)
		if err == nil {
			commipv4interface["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	commipv4interface["class_id"] = item.ClassID
	commipv4interface["gateway"] = item.Gateway
	commipv4interface["ip_address"] = item.IPAddress
	commipv4interface["netmask"] = item.Netmask
	commipv4interface["object_type"] = item.ObjectType

	commipv4interfaces = append(commipv4interfaces, commipv4interface)
	return commipv4interfaces
}
func flattenMapComputeBladeRef(p *models.ComputeBladeRef, d *schema.ResourceData) []map[string]interface{} {
	var computebladerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computebladeref := make(map[string]interface{})
	computebladeref["moid"] = item.Moid
	computebladeref["object_type"] = item.ObjectType
	computebladeref["selector"] = item.Selector

	computebladerefs = append(computebladerefs, computebladeref)
	return computebladerefs
}
func flattenMapComputeBoardRef(p *models.ComputeBoardRef, d *schema.ResourceData) []map[string]interface{} {
	var computeboardrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computeboardref := make(map[string]interface{})
	computeboardref["moid"] = item.Moid
	computeboardref["object_type"] = item.ObjectType
	computeboardref["selector"] = item.Selector

	computeboardrefs = append(computeboardrefs, computeboardref)
	return computeboardrefs
}
func flattenMapComputePersistentMemoryOperation(p *models.ComputePersistentMemoryOperation, d *schema.ResourceData) []map[string]interface{} {
	var computepersistentmemoryoperations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computepersistentmemoryoperation := make(map[string]interface{})
	delete(item.ComputePersistentMemoryOperationAO1P1.ComputePersistentMemoryOperationAO1P1, "ObjectType")
	if len(item.ComputePersistentMemoryOperationAO1P1.ComputePersistentMemoryOperationAO1P1) != 0 {
		j, err := json.Marshal(item.ComputePersistentMemoryOperationAO1P1.ComputePersistentMemoryOperationAO1P1)
		if err == nil {
			computepersistentmemoryoperation["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	computepersistentmemoryoperation["admin_action"] = item.AdminAction
	computepersistentmemoryoperation["class_id"] = item.ClassID
	computepersistentmemoryoperation["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	computepersistentmemoryoperation["modules"] = (func(p []*models.ComputePersistentMemoryModule, d *schema.ResourceData) []map[string]interface{} {
		var computepersistentmemorymodules []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			computepersistentmemorymodule := make(map[string]interface{})
			if len(item.ComputePersistentMemoryModuleAO1P1.ComputePersistentMemoryModuleAO1P1) != 0 {
				j, err := json.Marshal(item.ComputePersistentMemoryModuleAO1P1.ComputePersistentMemoryModuleAO1P1)
				if err == nil {
					computepersistentmemorymodule["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			computepersistentmemorymodule["class_id"] = item.ClassID
			computepersistentmemorymodule["object_type"] = item.ObjectType
			computepersistentmemorymodule["socket_id"] = item.SocketID
			computepersistentmemorymodule["socket_memory_id"] = item.SocketMemoryID
			computepersistentmemorymodules = append(computepersistentmemorymodules, computepersistentmemorymodule)
		}
		return computepersistentmemorymodules
	})(item.Modules, d)
	computepersistentmemoryoperation["object_type"] = item.ObjectType
	computepersistentmemoryoperation["secure_passphrase"] = item.SecurePassphrase

	computepersistentmemoryoperations = append(computepersistentmemoryoperations, computepersistentmemoryoperation)
	return computepersistentmemoryoperations
}
func flattenMapComputePhysicalRef(p *models.ComputePhysicalRef, d *schema.ResourceData) []map[string]interface{} {
	var computephysicalrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computephysicalref := make(map[string]interface{})
	computephysicalref["moid"] = item.Moid
	computephysicalref["object_type"] = item.ObjectType
	computephysicalref["selector"] = item.Selector

	computephysicalrefs = append(computephysicalrefs, computephysicalref)
	return computephysicalrefs
}
func flattenMapComputeRackUnitRef(p *models.ComputeRackUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computerackunitref := make(map[string]interface{})
	computerackunitref["moid"] = item.Moid
	computerackunitref["object_type"] = item.ObjectType
	computerackunitref["selector"] = item.Selector

	computerackunitrefs = append(computerackunitrefs, computerackunitref)
	return computerackunitrefs
}
func flattenMapComputeServerConfig(p *models.ComputeServerConfig, d *schema.ResourceData) []map[string]interface{} {
	var computeserverconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computeserverconfig := make(map[string]interface{})
	delete(item.ComputeServerConfigAO1P1.ComputeServerConfigAO1P1, "ObjectType")
	if len(item.ComputeServerConfigAO1P1.ComputeServerConfigAO1P1) != 0 {
		j, err := json.Marshal(item.ComputeServerConfigAO1P1.ComputeServerConfigAO1P1)
		if err == nil {
			computeserverconfig["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	computeserverconfig["asset_tag"] = item.AssetTag
	computeserverconfig["class_id"] = item.ClassID
	computeserverconfig["object_type"] = item.ObjectType
	computeserverconfig["user_label"] = item.UserLabel

	computeserverconfigs = append(computeserverconfigs, computeserverconfig)
	return computeserverconfigs
}
func flattenMapCondHclStatusRef(p *models.CondHclStatusRef, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	condhclstatusref := make(map[string]interface{})
	condhclstatusref["moid"] = item.Moid
	condhclstatusref["object_type"] = item.ObjectType
	condhclstatusref["selector"] = item.Selector

	condhclstatusrefs = append(condhclstatusrefs, condhclstatusref)
	return condhclstatusrefs
}
func flattenMapConfigExportedItemRef(p *models.ConfigExportedItemRef, d *schema.ResourceData) []map[string]interface{} {
	var configexporteditemrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configexporteditemref := make(map[string]interface{})
	configexporteditemref["moid"] = item.Moid
	configexporteditemref["object_type"] = item.ObjectType
	configexporteditemref["selector"] = item.Selector

	configexporteditemrefs = append(configexporteditemrefs, configexporteditemref)
	return configexporteditemrefs
}
func flattenMapConfigExporterRef(p *models.ConfigExporterRef, d *schema.ResourceData) []map[string]interface{} {
	var configexporterrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configexporterref := make(map[string]interface{})
	configexporterref["moid"] = item.Moid
	configexporterref["object_type"] = item.ObjectType
	configexporterref["selector"] = item.Selector

	configexporterrefs = append(configexporterrefs, configexporterref)
	return configexporterrefs
}
func flattenMapConfigImporterRef(p *models.ConfigImporterRef, d *schema.ResourceData) []map[string]interface{} {
	var configimporterrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configimporterref := make(map[string]interface{})
	configimporterref["moid"] = item.Moid
	configimporterref["object_type"] = item.ObjectType
	configimporterref["selector"] = item.Selector

	configimporterrefs = append(configimporterrefs, configimporterref)
	return configimporterrefs
}
func flattenMapConfigMoRef(p *models.ConfigMoRef, d *schema.ResourceData) []map[string]interface{} {
	var configmorefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configmoref := make(map[string]interface{})
	delete(item.ConfigMoRefAO1P1.ConfigMoRefAO1P1, "ObjectType")
	if len(item.ConfigMoRefAO1P1.ConfigMoRefAO1P1) != 0 {
		j, err := json.Marshal(item.ConfigMoRefAO1P1.ConfigMoRefAO1P1)
		if err == nil {
			configmoref["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	configmoref["class_id"] = item.ClassID
	configmoref["moid"] = item.Moid
	configmoref["object_type"] = item.ObjectType

	configmorefs = append(configmorefs, configmoref)
	return configmorefs
}
func flattenMapCryptDecryptRef(p *models.CryptDecryptRef, d *schema.ResourceData) []map[string]interface{} {
	var cryptdecryptrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	cryptdecryptref := make(map[string]interface{})
	cryptdecryptref["moid"] = item.Moid
	cryptdecryptref["object_type"] = item.ObjectType
	cryptdecryptref["selector"] = item.Selector

	cryptdecryptrefs = append(cryptdecryptrefs, cryptdecryptref)
	return cryptdecryptrefs
}
func flattenMapCryptEncryptRef(p *models.CryptEncryptRef, d *schema.ResourceData) []map[string]interface{} {
	var cryptencryptrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	cryptencryptref := make(map[string]interface{})
	cryptencryptref["moid"] = item.Moid
	cryptencryptref["object_type"] = item.ObjectType
	cryptencryptref["selector"] = item.Selector

	cryptencryptrefs = append(cryptencryptrefs, cryptencryptref)
	return cryptencryptrefs
}
func flattenMapEquipmentChassisRef(p *models.EquipmentChassisRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentchassisrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentchassisref := make(map[string]interface{})
	equipmentchassisref["moid"] = item.Moid
	equipmentchassisref["object_type"] = item.ObjectType
	equipmentchassisref["selector"] = item.Selector

	equipmentchassisrefs = append(equipmentchassisrefs, equipmentchassisref)
	return equipmentchassisrefs
}
func flattenMapEquipmentFanModuleRef(p *models.EquipmentFanModuleRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentfanmoduleref := make(map[string]interface{})
	equipmentfanmoduleref["moid"] = item.Moid
	equipmentfanmoduleref["object_type"] = item.ObjectType
	equipmentfanmoduleref["selector"] = item.Selector

	equipmentfanmodulerefs = append(equipmentfanmodulerefs, equipmentfanmoduleref)
	return equipmentfanmodulerefs
}
func flattenMapEquipmentLocatorLedRef(p *models.EquipmentLocatorLedRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentlocatorledrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentlocatorledref := make(map[string]interface{})
	equipmentlocatorledref["moid"] = item.Moid
	equipmentlocatorledref["object_type"] = item.ObjectType
	equipmentlocatorledref["selector"] = item.Selector

	equipmentlocatorledrefs = append(equipmentlocatorledrefs, equipmentlocatorledref)
	return equipmentlocatorledrefs
}
func flattenMapEquipmentRackEnclosureRef(p *models.EquipmentRackEnclosureRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosurerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentrackenclosureref := make(map[string]interface{})
	equipmentrackenclosureref["moid"] = item.Moid
	equipmentrackenclosureref["object_type"] = item.ObjectType
	equipmentrackenclosureref["selector"] = item.Selector

	equipmentrackenclosurerefs = append(equipmentrackenclosurerefs, equipmentrackenclosureref)
	return equipmentrackenclosurerefs
}
func flattenMapEquipmentRackEnclosureSlotRef(p *models.EquipmentRackEnclosureSlotRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentrackenclosureslotref := make(map[string]interface{})
	equipmentrackenclosureslotref["moid"] = item.Moid
	equipmentrackenclosureslotref["object_type"] = item.ObjectType
	equipmentrackenclosureslotref["selector"] = item.Selector

	equipmentrackenclosureslotrefs = append(equipmentrackenclosureslotrefs, equipmentrackenclosureslotref)
	return equipmentrackenclosureslotrefs
}
func flattenMapEquipmentSharedIoModuleRef(p *models.EquipmentSharedIoModuleRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsharediomodulerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentsharediomoduleref := make(map[string]interface{})
	equipmentsharediomoduleref["moid"] = item.Moid
	equipmentsharediomoduleref["object_type"] = item.ObjectType
	equipmentsharediomoduleref["selector"] = item.Selector

	equipmentsharediomodulerefs = append(equipmentsharediomodulerefs, equipmentsharediomoduleref)
	return equipmentsharediomodulerefs
}
func flattenMapEquipmentSwitchCardRef(p *models.EquipmentSwitchCardRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentswitchcardref := make(map[string]interface{})
	equipmentswitchcardref["moid"] = item.Moid
	equipmentswitchcardref["object_type"] = item.ObjectType
	equipmentswitchcardref["selector"] = item.Selector

	equipmentswitchcardrefs = append(equipmentswitchcardrefs, equipmentswitchcardref)
	return equipmentswitchcardrefs
}
func flattenMapEquipmentSystemIoControllerRef(p *models.EquipmentSystemIoControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentsystemiocontrollerref := make(map[string]interface{})
	equipmentsystemiocontrollerref["moid"] = item.Moid
	equipmentsystemiocontrollerref["object_type"] = item.ObjectType
	equipmentsystemiocontrollerref["selector"] = item.Selector

	equipmentsystemiocontrollerrefs = append(equipmentsystemiocontrollerrefs, equipmentsystemiocontrollerref)
	return equipmentsystemiocontrollerrefs
}
func flattenMapFirmwareDirectDownload(p *models.FirmwareDirectDownload, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredirectdownloads []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwaredirectdownload := make(map[string]interface{})
	delete(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1, "ObjectType")
	if len(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1) != 0 {
		j, err := json.Marshal(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1)
		if err == nil {
			firmwaredirectdownload["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	firmwaredirectdownload["class_id"] = item.ClassID
	firmwaredirectdownload["http_server"] = (func(p *models.FirmwareHTTPServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		firmwarehttpserver := make(map[string]interface{})
		delete(item.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1, "ObjectType")
		if len(item.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1) != 0 {
			j, err := json.Marshal(item.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1)
			if err == nil {
				firmwarehttpserver["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		firmwarehttpserver["class_id"] = item.ClassID
		firmwarehttpserver["location_link"] = item.LocationLink
		firmwarehttpserver["mount_options"] = item.MountOptions
		firmwarehttpserver["object_type"] = item.ObjectType

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.HTTPServer, d)
	firmwaredirectdownload["image_source"] = item.ImageSource
	firmwaredirectdownload["is_password_set"] = item.IsPasswordSet
	firmwaredirectdownload["object_type"] = item.ObjectType
	firmwaredirectdownload["password"] = item.Password
	firmwaredirectdownload["upgradeoption"] = item.Upgradeoption
	firmwaredirectdownload["username"] = item.Username

	firmwaredirectdownloads = append(firmwaredirectdownloads, firmwaredirectdownload)
	return firmwaredirectdownloads
}
func flattenMapFirmwareDistributableRef(p *models.FirmwareDistributableRef, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredistributablerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwaredistributableref := make(map[string]interface{})
	firmwaredistributableref["moid"] = item.Moid
	firmwaredistributableref["object_type"] = item.ObjectType
	firmwaredistributableref["selector"] = item.Selector

	firmwaredistributablerefs = append(firmwaredistributablerefs, firmwaredistributableref)
	return firmwaredistributablerefs
}
func flattenMapFirmwareNetworkShare(p *models.FirmwareNetworkShare, d *schema.ResourceData) []map[string]interface{} {
	var firmwarenetworkshares []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwarenetworkshare := make(map[string]interface{})
	delete(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1, "ObjectType")
	if len(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1) != 0 {
		j, err := json.Marshal(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1)
		if err == nil {
			firmwarenetworkshare["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	firmwarenetworkshare["cifs_server"] = (func(p *models.FirmwareCifsServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarecifsservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		firmwarecifsserver := make(map[string]interface{})
		delete(item.FirmwareCifsServerAO1P1.FirmwareCifsServerAO1P1, "ObjectType")
		if len(item.FirmwareCifsServerAO1P1.FirmwareCifsServerAO1P1) != 0 {
			j, err := json.Marshal(item.FirmwareCifsServerAO1P1.FirmwareCifsServerAO1P1)
			if err == nil {
				firmwarecifsserver["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		firmwarecifsserver["class_id"] = item.ClassID
		firmwarecifsserver["file_location"] = item.FileLocation
		firmwarecifsserver["mount_options"] = item.MountOptions
		firmwarecifsserver["object_type"] = item.ObjectType
		firmwarecifsserver["remote_file"] = item.RemoteFile
		firmwarecifsserver["remote_ip"] = item.RemoteIP
		firmwarecifsserver["remote_share"] = item.RemoteShare

		firmwarecifsservers = append(firmwarecifsservers, firmwarecifsserver)
		return firmwarecifsservers
	})(item.CifsServer, d)
	firmwarenetworkshare["class_id"] = item.ClassID
	firmwarenetworkshare["http_server"] = (func(p *models.FirmwareHTTPServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		firmwarehttpserver := make(map[string]interface{})
		delete(item.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1, "ObjectType")
		if len(item.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1) != 0 {
			j, err := json.Marshal(item.FirmwareHTTPServerAO1P1.FirmwareHTTPServerAO1P1)
			if err == nil {
				firmwarehttpserver["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		firmwarehttpserver["class_id"] = item.ClassID
		firmwarehttpserver["location_link"] = item.LocationLink
		firmwarehttpserver["mount_options"] = item.MountOptions
		firmwarehttpserver["object_type"] = item.ObjectType

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.HTTPServer, d)
	firmwarenetworkshare["is_password_set"] = item.IsPasswordSet
	firmwarenetworkshare["map_type"] = item.MapType
	firmwarenetworkshare["nfs_server"] = (func(p *models.FirmwareNfsServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarenfsservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		firmwarenfsserver := make(map[string]interface{})
		delete(item.FirmwareNfsServerAO1P1.FirmwareNfsServerAO1P1, "ObjectType")
		if len(item.FirmwareNfsServerAO1P1.FirmwareNfsServerAO1P1) != 0 {
			j, err := json.Marshal(item.FirmwareNfsServerAO1P1.FirmwareNfsServerAO1P1)
			if err == nil {
				firmwarenfsserver["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		firmwarenfsserver["class_id"] = item.ClassID
		firmwarenfsserver["file_location"] = item.FileLocation
		firmwarenfsserver["mount_options"] = item.MountOptions
		firmwarenfsserver["object_type"] = item.ObjectType
		firmwarenfsserver["remote_file"] = item.RemoteFile
		firmwarenfsserver["remote_ip"] = item.RemoteIP
		firmwarenfsserver["remote_share"] = item.RemoteShare

		firmwarenfsservers = append(firmwarenfsservers, firmwarenfsserver)
		return firmwarenfsservers
	})(item.NfsServer, d)
	firmwarenetworkshare["object_type"] = item.ObjectType
	firmwarenetworkshare["password"] = item.Password
	firmwarenetworkshare["upgradeoption"] = item.Upgradeoption
	firmwarenetworkshare["username"] = item.Username

	firmwarenetworkshares = append(firmwarenetworkshares, firmwarenetworkshare)
	return firmwarenetworkshares
}
func flattenMapFirmwareRunningFirmwareRef(p *models.FirmwareRunningFirmwareRef, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwarerunningfirmwareref := make(map[string]interface{})
	firmwarerunningfirmwareref["moid"] = item.Moid
	firmwarerunningfirmwareref["object_type"] = item.ObjectType
	firmwarerunningfirmwareref["selector"] = item.Selector

	firmwarerunningfirmwarerefs = append(firmwarerunningfirmwarerefs, firmwarerunningfirmwareref)
	return firmwarerunningfirmwarerefs
}
func flattenMapFirmwareServerConfigurationUtilityDistributableRef(p *models.FirmwareServerConfigurationUtilityDistributableRef, d *schema.ResourceData) []map[string]interface{} {
	var firmwareserverconfigurationutilitydistributablerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwareserverconfigurationutilitydistributableref := make(map[string]interface{})
	firmwareserverconfigurationutilitydistributableref["moid"] = item.Moid
	firmwareserverconfigurationutilitydistributableref["object_type"] = item.ObjectType
	firmwareserverconfigurationutilitydistributableref["selector"] = item.Selector

	firmwareserverconfigurationutilitydistributablerefs = append(firmwareserverconfigurationutilitydistributablerefs, firmwareserverconfigurationutilitydistributableref)
	return firmwareserverconfigurationutilitydistributablerefs
}
func flattenMapFirmwareUpgradeRef(p *models.FirmwareUpgradeRef, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgraderefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwareupgraderef := make(map[string]interface{})
	firmwareupgraderef["moid"] = item.Moid
	firmwareupgraderef["object_type"] = item.ObjectType
	firmwareupgraderef["selector"] = item.Selector

	firmwareupgraderefs = append(firmwareupgraderefs, firmwareupgraderef)
	return firmwareupgraderefs
}
func flattenMapFirmwareUpgradeStatusRef(p *models.FirmwareUpgradeStatusRef, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradestatusrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwareupgradestatusref := make(map[string]interface{})
	firmwareupgradestatusref["moid"] = item.Moid
	firmwareupgradestatusref["object_type"] = item.ObjectType
	firmwareupgradestatusref["selector"] = item.Selector

	firmwareupgradestatusrefs = append(firmwareupgradestatusrefs, firmwareupgradestatusref)
	return firmwareupgradestatusrefs
}
func flattenMapForecastCatalogRef(p *models.ForecastCatalogRef, d *schema.ResourceData) []map[string]interface{} {
	var forecastcatalogrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	forecastcatalogref := make(map[string]interface{})
	forecastcatalogref["moid"] = item.Moid
	forecastcatalogref["object_type"] = item.ObjectType
	forecastcatalogref["selector"] = item.Selector

	forecastcatalogrefs = append(forecastcatalogrefs, forecastcatalogref)
	return forecastcatalogrefs
}
func flattenMapForecastDefinitionRef(p *models.ForecastDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	forecastdefinitionref := make(map[string]interface{})
	forecastdefinitionref["moid"] = item.Moid
	forecastdefinitionref["object_type"] = item.ObjectType
	forecastdefinitionref["selector"] = item.Selector

	forecastdefinitionrefs = append(forecastdefinitionrefs, forecastdefinitionref)
	return forecastdefinitionrefs
}
func flattenMapForecastModel(p *models.ForecastModel, d *schema.ResourceData) []map[string]interface{} {
	var forecastmodels []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	forecastmodel := make(map[string]interface{})
	forecastmodel["accuracy"] = item.Accuracy
	delete(item.ForecastModelAO1P1.ForecastModelAO1P1, "ObjectType")
	if len(item.ForecastModelAO1P1.ForecastModelAO1P1) != 0 {
		j, err := json.Marshal(item.ForecastModelAO1P1.ForecastModelAO1P1)
		if err == nil {
			forecastmodel["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	forecastmodel["class_id"] = item.ClassID
	forecastmodel["model_data"] = item.ModelData
	forecastmodel["model_type"] = item.ModelType
	forecastmodel["object_type"] = item.ObjectType

	forecastmodels = append(forecastmodels, forecastmodel)
	return forecastmodels
}
func flattenMapGraphicsCardRef(p *models.GraphicsCardRef, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	graphicscardref := make(map[string]interface{})
	graphicscardref["moid"] = item.Moid
	graphicscardref["object_type"] = item.ObjectType
	graphicscardref["selector"] = item.Selector

	graphicscardrefs = append(graphicscardrefs, graphicscardref)
	return graphicscardrefs
}
func flattenMapHclOperatingSystemRef(p *models.HclOperatingSystemRef, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hcloperatingsystemref := make(map[string]interface{})
	hcloperatingsystemref["moid"] = item.Moid
	hcloperatingsystemref["object_type"] = item.ObjectType
	hcloperatingsystemref["selector"] = item.Selector

	hcloperatingsystemrefs = append(hcloperatingsystemrefs, hcloperatingsystemref)
	return hcloperatingsystemrefs
}
func flattenMapHclOperatingSystemVendorRef(p *models.HclOperatingSystemVendorRef, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemvendorrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hcloperatingsystemvendorref := make(map[string]interface{})
	hcloperatingsystemvendorref["moid"] = item.Moid
	hcloperatingsystemvendorref["object_type"] = item.ObjectType
	hcloperatingsystemvendorref["selector"] = item.Selector

	hcloperatingsystemvendorrefs = append(hcloperatingsystemvendorrefs, hcloperatingsystemvendorref)
	return hcloperatingsystemvendorrefs
}
func flattenMapHyperflexAppCatalogRef(p *models.HyperflexAppCatalogRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexappcatalogrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexappcatalogref := make(map[string]interface{})
	hyperflexappcatalogref["moid"] = item.Moid
	hyperflexappcatalogref["object_type"] = item.ObjectType
	hyperflexappcatalogref["selector"] = item.Selector

	hyperflexappcatalogrefs = append(hyperflexappcatalogrefs, hyperflexappcatalogref)
	return hyperflexappcatalogrefs
}
func flattenMapHyperflexAutoSupportPolicyRef(p *models.HyperflexAutoSupportPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexautosupportpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexautosupportpolicyref := make(map[string]interface{})
	hyperflexautosupportpolicyref["moid"] = item.Moid
	hyperflexautosupportpolicyref["object_type"] = item.ObjectType
	hyperflexautosupportpolicyref["selector"] = item.Selector

	hyperflexautosupportpolicyrefs = append(hyperflexautosupportpolicyrefs, hyperflexautosupportpolicyref)
	return hyperflexautosupportpolicyrefs
}
func flattenMapHyperflexClusterNetworkPolicyRef(p *models.HyperflexClusterNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusternetworkpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexclusternetworkpolicyref := make(map[string]interface{})
	hyperflexclusternetworkpolicyref["moid"] = item.Moid
	hyperflexclusternetworkpolicyref["object_type"] = item.ObjectType
	hyperflexclusternetworkpolicyref["selector"] = item.Selector

	hyperflexclusternetworkpolicyrefs = append(hyperflexclusternetworkpolicyrefs, hyperflexclusternetworkpolicyref)
	return hyperflexclusternetworkpolicyrefs
}
func flattenMapHyperflexClusterProfileRef(p *models.HyperflexClusterProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexclusterprofileref := make(map[string]interface{})
	hyperflexclusterprofileref["moid"] = item.Moid
	hyperflexclusterprofileref["object_type"] = item.ObjectType
	hyperflexclusterprofileref["selector"] = item.Selector

	hyperflexclusterprofilerefs = append(hyperflexclusterprofilerefs, hyperflexclusterprofileref)
	return hyperflexclusterprofilerefs
}
func flattenMapHyperflexClusterRef(p *models.HyperflexClusterRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexclusterref := make(map[string]interface{})
	hyperflexclusterref["moid"] = item.Moid
	hyperflexclusterref["object_type"] = item.ObjectType
	hyperflexclusterref["selector"] = item.Selector

	hyperflexclusterrefs = append(hyperflexclusterrefs, hyperflexclusterref)
	return hyperflexclusterrefs
}
func flattenMapHyperflexClusterStoragePolicyRef(p *models.HyperflexClusterStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterstoragepolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexclusterstoragepolicyref := make(map[string]interface{})
	hyperflexclusterstoragepolicyref["moid"] = item.Moid
	hyperflexclusterstoragepolicyref["object_type"] = item.ObjectType
	hyperflexclusterstoragepolicyref["selector"] = item.Selector

	hyperflexclusterstoragepolicyrefs = append(hyperflexclusterstoragepolicyrefs, hyperflexclusterstoragepolicyref)
	return hyperflexclusterstoragepolicyrefs
}
func flattenMapHyperflexConfigResultRef(p *models.HyperflexConfigResultRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexconfigresultref := make(map[string]interface{})
	hyperflexconfigresultref["moid"] = item.Moid
	hyperflexconfigresultref["object_type"] = item.ObjectType
	hyperflexconfigresultref["selector"] = item.Selector

	hyperflexconfigresultrefs = append(hyperflexconfigresultrefs, hyperflexconfigresultref)
	return hyperflexconfigresultrefs
}
func flattenMapHyperflexExtFcStoragePolicyRef(p *models.HyperflexExtFcStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextfcstoragepolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexextfcstoragepolicyref := make(map[string]interface{})
	hyperflexextfcstoragepolicyref["moid"] = item.Moid
	hyperflexextfcstoragepolicyref["object_type"] = item.ObjectType
	hyperflexextfcstoragepolicyref["selector"] = item.Selector

	hyperflexextfcstoragepolicyrefs = append(hyperflexextfcstoragepolicyrefs, hyperflexextfcstoragepolicyref)
	return hyperflexextfcstoragepolicyrefs
}
func flattenMapHyperflexExtIscsiStoragePolicyRef(p *models.HyperflexExtIscsiStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextiscsistoragepolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexextiscsistoragepolicyref := make(map[string]interface{})
	hyperflexextiscsistoragepolicyref["moid"] = item.Moid
	hyperflexextiscsistoragepolicyref["object_type"] = item.ObjectType
	hyperflexextiscsistoragepolicyref["selector"] = item.Selector

	hyperflexextiscsistoragepolicyrefs = append(hyperflexextiscsistoragepolicyrefs, hyperflexextiscsistoragepolicyref)
	return hyperflexextiscsistoragepolicyrefs
}
func flattenMapHyperflexFeatureLimitExternalRef(p *models.HyperflexFeatureLimitExternalRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitexternalrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexfeaturelimitexternalref := make(map[string]interface{})
	hyperflexfeaturelimitexternalref["moid"] = item.Moid
	hyperflexfeaturelimitexternalref["object_type"] = item.ObjectType
	hyperflexfeaturelimitexternalref["selector"] = item.Selector

	hyperflexfeaturelimitexternalrefs = append(hyperflexfeaturelimitexternalrefs, hyperflexfeaturelimitexternalref)
	return hyperflexfeaturelimitexternalrefs
}
func flattenMapHyperflexFeatureLimitInternalRef(p *models.HyperflexFeatureLimitInternalRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitinternalrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexfeaturelimitinternalref := make(map[string]interface{})
	hyperflexfeaturelimitinternalref["moid"] = item.Moid
	hyperflexfeaturelimitinternalref["object_type"] = item.ObjectType
	hyperflexfeaturelimitinternalref["selector"] = item.Selector

	hyperflexfeaturelimitinternalrefs = append(hyperflexfeaturelimitinternalrefs, hyperflexfeaturelimitinternalref)
	return hyperflexfeaturelimitinternalrefs
}
func flattenMapHyperflexHealthRef(p *models.HyperflexHealthRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhealthrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexhealthref := make(map[string]interface{})
	hyperflexhealthref["moid"] = item.Moid
	hyperflexhealthref["object_type"] = item.ObjectType
	hyperflexhealthref["selector"] = item.Selector

	hyperflexhealthrefs = append(hyperflexhealthrefs, hyperflexhealthref)
	return hyperflexhealthrefs
}
func flattenMapHyperflexHxNetworkAddressDt(p *models.HyperflexHxNetworkAddressDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxnetworkaddressdts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexhxnetworkaddressdt := make(map[string]interface{})
	delete(item.HyperflexHxNetworkAddressDtAO1P1.HyperflexHxNetworkAddressDtAO1P1, "ObjectType")
	if len(item.HyperflexHxNetworkAddressDtAO1P1.HyperflexHxNetworkAddressDtAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexHxNetworkAddressDtAO1P1.HyperflexHxNetworkAddressDtAO1P1)
		if err == nil {
			hyperflexhxnetworkaddressdt["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexhxnetworkaddressdt["address"] = item.Address
	hyperflexhxnetworkaddressdt["class_id"] = item.ClassID
	hyperflexhxnetworkaddressdt["fqdn"] = item.Fqdn
	hyperflexhxnetworkaddressdt["ip"] = item.IP
	hyperflexhxnetworkaddressdt["object_type"] = item.ObjectType

	hyperflexhxnetworkaddressdts = append(hyperflexhxnetworkaddressdts, hyperflexhxnetworkaddressdt)
	return hyperflexhxnetworkaddressdts
}
func flattenMapHyperflexHxResiliencyInfoDt(p *models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxresiliencyinfodts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexhxresiliencyinfodt := make(map[string]interface{})
	delete(item.HyperflexHxResiliencyInfoDtAO1P1.HyperflexHxResiliencyInfoDtAO1P1, "ObjectType")
	if len(item.HyperflexHxResiliencyInfoDtAO1P1.HyperflexHxResiliencyInfoDtAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexHxResiliencyInfoDtAO1P1.HyperflexHxResiliencyInfoDtAO1P1)
		if err == nil {
			hyperflexhxresiliencyinfodt["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexhxresiliencyinfodt["class_id"] = item.ClassID
	hyperflexhxresiliencyinfodt["data_replication_factor"] = item.DataReplicationFactor
	hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.HddFailuresTolerable
	hyperflexhxresiliencyinfodt["messages"] = item.Messages
	hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.NodeFailuresTolerable
	hyperflexhxresiliencyinfodt["object_type"] = item.ObjectType
	hyperflexhxresiliencyinfodt["policy_compliance"] = item.PolicyCompliance
	hyperflexhxresiliencyinfodt["resiliency_state"] = item.ResiliencyState
	hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.SsdFailuresTolerable

	hyperflexhxresiliencyinfodts = append(hyperflexhxresiliencyinfodts, hyperflexhxresiliencyinfodt)
	return hyperflexhxresiliencyinfodts
}
func flattenMapHyperflexHxUuIDDt(p *models.HyperflexHxUuIDDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxuuiddts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexhxuuiddt := make(map[string]interface{})
	delete(item.HyperflexHxUuIDDtAO1P1.HyperflexHxUuIDDtAO1P1, "ObjectType")
	if len(item.HyperflexHxUuIDDtAO1P1.HyperflexHxUuIDDtAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexHxUuIDDtAO1P1.HyperflexHxUuIDDtAO1P1)
		if err == nil {
			hyperflexhxuuiddt["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexhxuuiddt["class_id"] = item.ClassID
	hyperflexhxuuiddt["links"] = (func(p []*models.HyperflexHxLinkDt, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexhxlinkdts []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			hyperflexhxlinkdt := make(map[string]interface{})
			if len(item.HyperflexHxLinkDtAO1P1.HyperflexHxLinkDtAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexHxLinkDtAO1P1.HyperflexHxLinkDtAO1P1)
				if err == nil {
					hyperflexhxlinkdt["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			hyperflexhxlinkdt["class_id"] = item.ClassID
			hyperflexhxlinkdt["comments"] = item.Comments
			hyperflexhxlinkdt["href"] = item.Href
			hyperflexhxlinkdt["method"] = item.Method
			hyperflexhxlinkdt["object_type"] = item.ObjectType
			hyperflexhxlinkdt["rel"] = item.Rel
			hyperflexhxlinkdts = append(hyperflexhxlinkdts, hyperflexhxlinkdt)
		}
		return hyperflexhxlinkdts
	})(item.Links, d)
	hyperflexhxuuiddt["object_type"] = item.ObjectType
	hyperflexhxuuiddt["uuid"] = item.UUID

	hyperflexhxuuiddts = append(hyperflexhxuuiddts, hyperflexhxuuiddt)
	return hyperflexhxuuiddts
}
func flattenMapHyperflexInstallerImageRef(p *models.HyperflexInstallerImageRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexinstallerimagerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexinstallerimageref := make(map[string]interface{})
	hyperflexinstallerimageref["moid"] = item.Moid
	hyperflexinstallerimageref["object_type"] = item.ObjectType
	hyperflexinstallerimageref["selector"] = item.Selector

	hyperflexinstallerimagerefs = append(hyperflexinstallerimagerefs, hyperflexinstallerimageref)
	return hyperflexinstallerimagerefs
}
func flattenMapHyperflexIPAddrRange(p *models.HyperflexIPAddrRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexipaddrranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexipaddrrange := make(map[string]interface{})
	delete(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1, "ObjectType")
	if len(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1)
		if err == nil {
			hyperflexipaddrrange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexipaddrrange["class_id"] = item.ClassID
	hyperflexipaddrrange["end_addr"] = item.EndAddr
	hyperflexipaddrrange["gateway"] = item.Gateway
	hyperflexipaddrrange["netmask"] = item.Netmask
	hyperflexipaddrrange["object_type"] = item.ObjectType
	hyperflexipaddrrange["start_addr"] = item.StartAddr

	hyperflexipaddrranges = append(hyperflexipaddrranges, hyperflexipaddrrange)
	return hyperflexipaddrranges
}
func flattenMapHyperflexLocalCredentialPolicyRef(p *models.HyperflexLocalCredentialPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlocalcredentialpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexlocalcredentialpolicyref := make(map[string]interface{})
	hyperflexlocalcredentialpolicyref["moid"] = item.Moid
	hyperflexlocalcredentialpolicyref["object_type"] = item.ObjectType
	hyperflexlocalcredentialpolicyref["selector"] = item.Selector

	hyperflexlocalcredentialpolicyrefs = append(hyperflexlocalcredentialpolicyrefs, hyperflexlocalcredentialpolicyref)
	return hyperflexlocalcredentialpolicyrefs
}
func flattenMapHyperflexLogicalAvailabilityZone(p *models.HyperflexLogicalAvailabilityZone, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlogicalavailabilityzones []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexlogicalavailabilityzone := make(map[string]interface{})
	delete(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1, "ObjectType")
	if len(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1)
		if err == nil {
			hyperflexlogicalavailabilityzone["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexlogicalavailabilityzone["auto_config"] = item.AutoConfig
	hyperflexlogicalavailabilityzone["class_id"] = item.ClassID
	hyperflexlogicalavailabilityzone["object_type"] = item.ObjectType

	hyperflexlogicalavailabilityzones = append(hyperflexlogicalavailabilityzones, hyperflexlogicalavailabilityzone)
	return hyperflexlogicalavailabilityzones
}
func flattenMapHyperflexMacAddrPrefixRange(p *models.HyperflexMacAddrPrefixRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexmacaddrprefixranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexmacaddrprefixrange := make(map[string]interface{})
	delete(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1, "ObjectType")
	if len(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1)
		if err == nil {
			hyperflexmacaddrprefixrange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexmacaddrprefixrange["class_id"] = item.ClassID
	hyperflexmacaddrprefixrange["end_addr"] = item.EndAddr
	hyperflexmacaddrprefixrange["object_type"] = item.ObjectType
	hyperflexmacaddrprefixrange["start_addr"] = item.StartAddr

	hyperflexmacaddrprefixranges = append(hyperflexmacaddrprefixranges, hyperflexmacaddrprefixrange)
	return hyperflexmacaddrprefixranges
}
func flattenMapHyperflexNamedVlan(p *models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvlans []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexnamedvlan := make(map[string]interface{})
	delete(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1, "ObjectType")
	if len(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1)
		if err == nil {
			hyperflexnamedvlan["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexnamedvlan["class_id"] = item.ClassID
	hyperflexnamedvlan["name"] = item.Name
	hyperflexnamedvlan["object_type"] = item.ObjectType
	hyperflexnamedvlan["vlan_id"] = item.VlanID

	hyperflexnamedvlans = append(hyperflexnamedvlans, hyperflexnamedvlan)
	return hyperflexnamedvlans
}
func flattenMapHyperflexNamedVsan(p *models.HyperflexNamedVsan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvsans []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexnamedvsan := make(map[string]interface{})
	delete(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1, "ObjectType")
	if len(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1)
		if err == nil {
			hyperflexnamedvsan["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexnamedvsan["class_id"] = item.ClassID
	hyperflexnamedvsan["name"] = item.Name
	hyperflexnamedvsan["object_type"] = item.ObjectType
	hyperflexnamedvsan["vsan_id"] = item.VsanID

	hyperflexnamedvsans = append(hyperflexnamedvsans, hyperflexnamedvsan)
	return hyperflexnamedvsans
}
func flattenMapHyperflexNodeConfigPolicyRef(p *models.HyperflexNodeConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeconfigpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexnodeconfigpolicyref := make(map[string]interface{})
	hyperflexnodeconfigpolicyref["moid"] = item.Moid
	hyperflexnodeconfigpolicyref["object_type"] = item.ObjectType
	hyperflexnodeconfigpolicyref["selector"] = item.Selector

	hyperflexnodeconfigpolicyrefs = append(hyperflexnodeconfigpolicyrefs, hyperflexnodeconfigpolicyref)
	return hyperflexnodeconfigpolicyrefs
}
func flattenMapHyperflexProxySettingPolicyRef(p *models.HyperflexProxySettingPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexproxysettingpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexproxysettingpolicyref := make(map[string]interface{})
	hyperflexproxysettingpolicyref["moid"] = item.Moid
	hyperflexproxysettingpolicyref["object_type"] = item.ObjectType
	hyperflexproxysettingpolicyref["selector"] = item.Selector

	hyperflexproxysettingpolicyrefs = append(hyperflexproxysettingpolicyrefs, hyperflexproxysettingpolicyref)
	return hyperflexproxysettingpolicyrefs
}
func flattenMapHyperflexServerFirmwareVersionRef(p *models.HyperflexServerFirmwareVersionRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexserverfirmwareversionref := make(map[string]interface{})
	hyperflexserverfirmwareversionref["moid"] = item.Moid
	hyperflexserverfirmwareversionref["object_type"] = item.ObjectType
	hyperflexserverfirmwareversionref["selector"] = item.Selector

	hyperflexserverfirmwareversionrefs = append(hyperflexserverfirmwareversionrefs, hyperflexserverfirmwareversionref)
	return hyperflexserverfirmwareversionrefs
}
func flattenMapHyperflexServerModelRef(p *models.HyperflexServerModelRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexservermodelref := make(map[string]interface{})
	hyperflexservermodelref["moid"] = item.Moid
	hyperflexservermodelref["object_type"] = item.ObjectType
	hyperflexservermodelref["selector"] = item.Selector

	hyperflexservermodelrefs = append(hyperflexservermodelrefs, hyperflexservermodelref)
	return hyperflexservermodelrefs
}
func flattenMapHyperflexSoftwareVersionPolicyRef(p *models.HyperflexSoftwareVersionPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwareversionpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexsoftwareversionpolicyref := make(map[string]interface{})
	hyperflexsoftwareversionpolicyref["moid"] = item.Moid
	hyperflexsoftwareversionpolicyref["object_type"] = item.ObjectType
	hyperflexsoftwareversionpolicyref["selector"] = item.Selector

	hyperflexsoftwareversionpolicyrefs = append(hyperflexsoftwareversionpolicyrefs, hyperflexsoftwareversionpolicyref)
	return hyperflexsoftwareversionpolicyrefs
}
func flattenMapHyperflexSummary(p *models.HyperflexSummary, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsummarys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexsummary := make(map[string]interface{})
	hyperflexsummary["active_nodes"] = item.ActiveNodes
	delete(item.HyperflexSummaryAO1P1.HyperflexSummaryAO1P1, "ObjectType")
	if len(item.HyperflexSummaryAO1P1.HyperflexSummaryAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexSummaryAO1P1.HyperflexSummaryAO1P1)
		if err == nil {
			hyperflexsummary["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexsummary["address"] = item.Address
	hyperflexsummary["boottime"] = item.Boottime
	hyperflexsummary["class_id"] = item.ClassID
	hyperflexsummary["cluster_access_policy"] = item.ClusterAccessPolicy
	hyperflexsummary["compression_savings"] = item.CompressionSavings
	hyperflexsummary["data_replication_compliance"] = item.DataReplicationCompliance
	hyperflexsummary["data_replication_factor"] = item.DataReplicationFactor
	hyperflexsummary["deduplication_savings"] = item.DeduplicationSavings
	hyperflexsummary["downtime"] = item.Downtime
	hyperflexsummary["free_capacity"] = item.FreeCapacity
	hyperflexsummary["healing_info"] = (func(p *models.HyperflexStPlatformClusterHealingInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterhealinginfos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		hyperflexstplatformclusterhealinginfo := make(map[string]interface{})
		delete(item.HyperflexStPlatformClusterHealingInfoAO1P1.HyperflexStPlatformClusterHealingInfoAO1P1, "ObjectType")
		if len(item.HyperflexStPlatformClusterHealingInfoAO1P1.HyperflexStPlatformClusterHealingInfoAO1P1) != 0 {
			j, err := json.Marshal(item.HyperflexStPlatformClusterHealingInfoAO1P1.HyperflexStPlatformClusterHealingInfoAO1P1)
			if err == nil {
				hyperflexstplatformclusterhealinginfo["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hyperflexstplatformclusterhealinginfo["class_id"] = item.ClassID
		hyperflexstplatformclusterhealinginfo["estimated_completion_time_in_seconds"] = item.EstimatedCompletionTimeInSeconds
		hyperflexstplatformclusterhealinginfo["in_progress"] = item.InProgress
		hyperflexstplatformclusterhealinginfo["messages"] = item.Messages
		hyperflexstplatformclusterhealinginfo["messages_iterator"] = item.MessagesIterator
		hyperflexstplatformclusterhealinginfo["messages_size"] = item.MessagesSize
		hyperflexstplatformclusterhealinginfo["object_type"] = item.ObjectType
		hyperflexstplatformclusterhealinginfo["percent_complete"] = item.PercentComplete

		hyperflexstplatformclusterhealinginfos = append(hyperflexstplatformclusterhealinginfos, hyperflexstplatformclusterhealinginfo)
		return hyperflexstplatformclusterhealinginfos
	})(item.HealingInfo, d)
	hyperflexsummary["name"] = item.Name
	hyperflexsummary["object_type"] = item.ObjectType
	hyperflexsummary["resiliency_details"] = item.ResiliencyDetails
	hyperflexsummary["resiliency_details_size"] = item.ResiliencyDetailsSize
	hyperflexsummary["resiliency_info"] = (func(p *models.HyperflexStPlatformClusterResiliencyInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterresiliencyinfos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		hyperflexstplatformclusterresiliencyinfo := make(map[string]interface{})
		delete(item.HyperflexStPlatformClusterResiliencyInfoAO1P1.HyperflexStPlatformClusterResiliencyInfoAO1P1, "ObjectType")
		if len(item.HyperflexStPlatformClusterResiliencyInfoAO1P1.HyperflexStPlatformClusterResiliencyInfoAO1P1) != 0 {
			j, err := json.Marshal(item.HyperflexStPlatformClusterResiliencyInfoAO1P1.HyperflexStPlatformClusterResiliencyInfoAO1P1)
			if err == nil {
				hyperflexstplatformclusterresiliencyinfo["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		hyperflexstplatformclusterresiliencyinfo["class_id"] = item.ClassID
		hyperflexstplatformclusterresiliencyinfo["hdd_failures_tolerable"] = item.HddFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["messages"] = item.Messages
		hyperflexstplatformclusterresiliencyinfo["messages_iterator"] = item.MessagesIterator
		hyperflexstplatformclusterresiliencyinfo["messages_size"] = item.MessagesSize
		hyperflexstplatformclusterresiliencyinfo["node_failures_tolerable"] = item.NodeFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["object_type"] = item.ObjectType
		hyperflexstplatformclusterresiliencyinfo["ssd_failures_tolerable"] = item.SsdFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["state"] = item.State

		hyperflexstplatformclusterresiliencyinfos = append(hyperflexstplatformclusterresiliencyinfos, hyperflexstplatformclusterresiliencyinfo)
		return hyperflexstplatformclusterresiliencyinfos
	})(item.ResiliencyInfo, d)
	hyperflexsummary["space_status"] = item.SpaceStatus
	hyperflexsummary["state"] = item.State
	hyperflexsummary["total_capacity"] = item.TotalCapacity
	hyperflexsummary["total_savings"] = item.TotalSavings
	hyperflexsummary["uptime"] = item.Uptime
	hyperflexsummary["used_capacity"] = item.UsedCapacity

	hyperflexsummarys = append(hyperflexsummarys, hyperflexsummary)
	return hyperflexsummarys
}
func flattenMapHyperflexSysConfigPolicyRef(p *models.HyperflexSysConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsysconfigpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexsysconfigpolicyref := make(map[string]interface{})
	hyperflexsysconfigpolicyref["moid"] = item.Moid
	hyperflexsysconfigpolicyref["object_type"] = item.ObjectType
	hyperflexsysconfigpolicyref["selector"] = item.Selector

	hyperflexsysconfigpolicyrefs = append(hyperflexsysconfigpolicyrefs, hyperflexsysconfigpolicyref)
	return hyperflexsysconfigpolicyrefs
}
func flattenMapHyperflexUcsmConfigPolicyRef(p *models.HyperflexUcsmConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexucsmconfigpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexucsmconfigpolicyref := make(map[string]interface{})
	hyperflexucsmconfigpolicyref["moid"] = item.Moid
	hyperflexucsmconfigpolicyref["object_type"] = item.ObjectType
	hyperflexucsmconfigpolicyref["selector"] = item.Selector

	hyperflexucsmconfigpolicyrefs = append(hyperflexucsmconfigpolicyrefs, hyperflexucsmconfigpolicyref)
	return hyperflexucsmconfigpolicyrefs
}
func flattenMapHyperflexVcenterConfigPolicyRef(p *models.HyperflexVcenterConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvcenterconfigpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexvcenterconfigpolicyref := make(map[string]interface{})
	hyperflexvcenterconfigpolicyref["moid"] = item.Moid
	hyperflexvcenterconfigpolicyref["object_type"] = item.ObjectType
	hyperflexvcenterconfigpolicyref["selector"] = item.Selector

	hyperflexvcenterconfigpolicyrefs = append(hyperflexvcenterconfigpolicyrefs, hyperflexvcenterconfigpolicyref)
	return hyperflexvcenterconfigpolicyrefs
}
func flattenMapHyperflexWwxnPrefixRange(p *models.HyperflexWwxnPrefixRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexwwxnprefixranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexwwxnprefixrange := make(map[string]interface{})
	delete(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1, "ObjectType")
	if len(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1)
		if err == nil {
			hyperflexwwxnprefixrange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	hyperflexwwxnprefixrange["class_id"] = item.ClassID
	hyperflexwwxnprefixrange["end_addr"] = item.EndAddr
	hyperflexwwxnprefixrange["object_type"] = item.ObjectType
	hyperflexwwxnprefixrange["start_addr"] = item.StartAddr

	hyperflexwwxnprefixranges = append(hyperflexwwxnprefixranges, hyperflexwwxnprefixrange)
	return hyperflexwwxnprefixranges
}
func flattenMapIaasLicenseInfoRef(p *models.IaasLicenseInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iaaslicenseinforef := make(map[string]interface{})
	iaaslicenseinforef["moid"] = item.Moid
	iaaslicenseinforef["object_type"] = item.ObjectType
	iaaslicenseinforef["selector"] = item.Selector

	iaaslicenseinforefs = append(iaaslicenseinforefs, iaaslicenseinforef)
	return iaaslicenseinforefs
}
func flattenMapIaasUcsdInfoRef(p *models.IaasUcsdInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iaasucsdinforef := make(map[string]interface{})
	iaasucsdinforef["moid"] = item.Moid
	iaasucsdinforef["object_type"] = item.ObjectType
	iaasucsdinforef["selector"] = item.Selector

	iaasucsdinforefs = append(iaasucsdinforefs, iaasucsdinforef)
	return iaasucsdinforefs
}
func flattenMapIaasUcsdManagedInfraRef(p *models.IaasUcsdManagedInfraRef, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdmanagedinfrarefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iaasucsdmanagedinfraref := make(map[string]interface{})
	iaasucsdmanagedinfraref["moid"] = item.Moid
	iaasucsdmanagedinfraref["object_type"] = item.ObjectType
	iaasucsdmanagedinfraref["selector"] = item.Selector

	iaasucsdmanagedinfrarefs = append(iaasucsdmanagedinfrarefs, iaasucsdmanagedinfraref)
	return iaasucsdmanagedinfrarefs
}
func flattenMapIamAccountRef(p *models.IamAccountRef, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamaccountref := make(map[string]interface{})
	iamaccountref["moid"] = item.Moid
	iamaccountref["object_type"] = item.ObjectType
	iamaccountref["selector"] = item.Selector

	iamaccountrefs = append(iamaccountrefs, iamaccountref)
	return iamaccountrefs
}
func flattenMapIamAppRegistrationRef(p *models.IamAppRegistrationRef, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamappregistrationref := make(map[string]interface{})
	iamappregistrationref["moid"] = item.Moid
	iamappregistrationref["object_type"] = item.ObjectType
	iamappregistrationref["selector"] = item.Selector

	iamappregistrationrefs = append(iamappregistrationrefs, iamappregistrationref)
	return iamappregistrationrefs
}
func flattenMapIamCertificateRef(p *models.IamCertificateRef, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamcertificateref := make(map[string]interface{})
	iamcertificateref["moid"] = item.Moid
	iamcertificateref["object_type"] = item.ObjectType
	iamcertificateref["selector"] = item.Selector

	iamcertificaterefs = append(iamcertificaterefs, iamcertificateref)
	return iamcertificaterefs
}
func flattenMapIamCertificateRequestRef(p *models.IamCertificateRequestRef, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterequestrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamcertificaterequestref := make(map[string]interface{})
	iamcertificaterequestref["moid"] = item.Moid
	iamcertificaterequestref["object_type"] = item.ObjectType
	iamcertificaterequestref["selector"] = item.Selector

	iamcertificaterequestrefs = append(iamcertificaterequestrefs, iamcertificaterequestref)
	return iamcertificaterequestrefs
}
func flattenMapIamClientMeta(p *models.IamClientMeta, d *schema.ResourceData) []map[string]interface{} {
	var iamclientmetas []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamclientmeta := make(map[string]interface{})
	delete(item.IamClientMetaAO1P1.IamClientMetaAO1P1, "ObjectType")
	if len(item.IamClientMetaAO1P1.IamClientMetaAO1P1) != 0 {
		j, err := json.Marshal(item.IamClientMetaAO1P1.IamClientMetaAO1P1)
		if err == nil {
			iamclientmeta["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	iamclientmeta["class_id"] = item.ClassID
	iamclientmeta["device_model"] = item.DeviceModel
	iamclientmeta["object_type"] = item.ObjectType
	iamclientmeta["user_agent"] = item.UserAgent

	iamclientmetas = append(iamclientmetas, iamclientmeta)
	return iamclientmetas
}
func flattenMapIamDomainGroupRef(p *models.IamDomainGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamdomaingroupref := make(map[string]interface{})
	iamdomaingroupref["moid"] = item.Moid
	iamdomaingroupref["object_type"] = item.ObjectType
	iamdomaingroupref["selector"] = item.Selector

	iamdomaingrouprefs = append(iamdomaingrouprefs, iamdomaingroupref)
	return iamdomaingrouprefs
}
func flattenMapIamEndPointPasswordProperties(p *models.IamEndPointPasswordProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointpasswordpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamendpointpasswordproperties := make(map[string]interface{})
	delete(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1, "ObjectType")
	if len(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1)
		if err == nil {
			iamendpointpasswordproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	iamendpointpasswordproperties["class_id"] = item.ClassID
	iamendpointpasswordproperties["enable_password_expiry"] = item.EnablePasswordExpiry
	iamendpointpasswordproperties["enforce_strong_password"] = item.EnforceStrongPassword
	iamendpointpasswordproperties["force_send_password"] = item.ForceSendPassword
	iamendpointpasswordproperties["grace_period"] = item.GracePeriod
	iamendpointpasswordproperties["notification_period"] = item.NotificationPeriod
	iamendpointpasswordproperties["object_type"] = item.ObjectType
	iamendpointpasswordproperties["password_expiry_duration"] = item.PasswordExpiryDuration
	iamendpointpasswordproperties["password_history"] = item.PasswordHistory

	iamendpointpasswordpropertiess = append(iamendpointpasswordpropertiess, iamendpointpasswordproperties)
	return iamendpointpasswordpropertiess
}
func flattenMapIamEndPointUserPolicyRef(p *models.IamEndPointUserPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamendpointuserpolicyref := make(map[string]interface{})
	iamendpointuserpolicyref["moid"] = item.Moid
	iamendpointuserpolicyref["object_type"] = item.ObjectType
	iamendpointuserpolicyref["selector"] = item.Selector

	iamendpointuserpolicyrefs = append(iamendpointuserpolicyrefs, iamendpointuserpolicyref)
	return iamendpointuserpolicyrefs
}
func flattenMapIamEndPointUserRef(p *models.IamEndPointUserRef, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamendpointuserref := make(map[string]interface{})
	iamendpointuserref["moid"] = item.Moid
	iamendpointuserref["object_type"] = item.ObjectType
	iamendpointuserref["selector"] = item.Selector

	iamendpointuserrefs = append(iamendpointuserrefs, iamendpointuserref)
	return iamendpointuserrefs
}
func flattenMapIamIdpRef(p *models.IamIdpRef, d *schema.ResourceData) []map[string]interface{} {
	var iamidprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamidpref := make(map[string]interface{})
	iamidpref["moid"] = item.Moid
	iamidpref["object_type"] = item.ObjectType
	iamidpref["selector"] = item.Selector

	iamidprefs = append(iamidprefs, iamidpref)
	return iamidprefs
}
func flattenMapIamIdpReferenceRef(p *models.IamIdpReferenceRef, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamidpreferenceref := make(map[string]interface{})
	iamidpreferenceref["moid"] = item.Moid
	iamidpreferenceref["object_type"] = item.ObjectType
	iamidpreferenceref["selector"] = item.Selector

	iamidpreferencerefs = append(iamidpreferencerefs, iamidpreferenceref)
	return iamidpreferencerefs
}
func flattenMapIamLdapBaseProperties(p *models.IamLdapBaseProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamldapbasepropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamldapbaseproperties := make(map[string]interface{})
	delete(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1, "ObjectType")
	if len(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1)
		if err == nil {
			iamldapbaseproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	iamldapbaseproperties["attribute"] = item.Attribute
	iamldapbaseproperties["base_dn"] = item.BaseDn
	iamldapbaseproperties["bind_dn"] = item.BindDn
	iamldapbaseproperties["bind_method"] = item.BindMethod
	iamldapbaseproperties["class_id"] = item.ClassID
	iamldapbaseproperties["domain"] = item.Domain
	iamldapbaseproperties["enable_encryption"] = item.EnableEncryption
	iamldapbaseproperties["enable_group_authorization"] = item.EnableGroupAuthorization
	iamldapbaseproperties["filter"] = item.Filter
	iamldapbaseproperties["group_attribute"] = item.GroupAttribute
	iamldapbaseproperties["is_password_set"] = item.IsPasswordSet
	iamldapbaseproperties["nested_group_search_depth"] = item.NestedGroupSearchDepth
	iamldapbaseproperties["object_type"] = item.ObjectType
	password_x := d.Get("base_properties").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	iamldapbaseproperties["password"] = password_y["password"]
	iamldapbaseproperties["timeout"] = item.Timeout

	iamldapbasepropertiess = append(iamldapbasepropertiess, iamldapbaseproperties)
	return iamldapbasepropertiess
}
func flattenMapIamLdapDNSParameters(p *models.IamLdapDNSParameters, d *schema.ResourceData) []map[string]interface{} {
	var iamldapdnsparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamldapdnsparameters := make(map[string]interface{})
	delete(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1, "ObjectType")
	if len(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1) != 0 {
		j, err := json.Marshal(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1)
		if err == nil {
			iamldapdnsparameters["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	iamldapdnsparameters["class_id"] = item.ClassID
	iamldapdnsparameters["object_type"] = item.ObjectType
	iamldapdnsparameters["search_domain"] = item.SearchDomain
	iamldapdnsparameters["search_forest"] = item.SearchForest
	iamldapdnsparameters["source"] = item.Source

	iamldapdnsparameterss = append(iamldapdnsparameterss, iamldapdnsparameters)
	return iamldapdnsparameterss
}
func flattenMapIamLdapPolicyRef(p *models.IamLdapPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var iamldappolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamldappolicyref := make(map[string]interface{})
	iamldappolicyref["moid"] = item.Moid
	iamldappolicyref["object_type"] = item.ObjectType
	iamldappolicyref["selector"] = item.Selector

	iamldappolicyrefs = append(iamldappolicyrefs, iamldappolicyref)
	return iamldappolicyrefs
}
func flattenMapIamLocalUserPasswordRef(p *models.IamLocalUserPasswordRef, d *schema.ResourceData) []map[string]interface{} {
	var iamlocaluserpasswordrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamlocaluserpasswordref := make(map[string]interface{})
	iamlocaluserpasswordref["moid"] = item.Moid
	iamlocaluserpasswordref["object_type"] = item.ObjectType
	iamlocaluserpasswordref["selector"] = item.Selector

	iamlocaluserpasswordrefs = append(iamlocaluserpasswordrefs, iamlocaluserpasswordref)
	return iamlocaluserpasswordrefs
}
func flattenMapIamPermissionRef(p *models.IamPermissionRef, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iampermissionref := make(map[string]interface{})
	iampermissionref["moid"] = item.Moid
	iampermissionref["object_type"] = item.ObjectType
	iampermissionref["selector"] = item.Selector

	iampermissionrefs = append(iampermissionrefs, iampermissionref)
	return iampermissionrefs
}
func flattenMapIamPrivateKeySpecRef(p *models.IamPrivateKeySpecRef, d *schema.ResourceData) []map[string]interface{} {
	var iamprivatekeyspecrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamprivatekeyspecref := make(map[string]interface{})
	iamprivatekeyspecref["moid"] = item.Moid
	iamprivatekeyspecref["object_type"] = item.ObjectType
	iamprivatekeyspecref["selector"] = item.Selector

	iamprivatekeyspecrefs = append(iamprivatekeyspecrefs, iamprivatekeyspecref)
	return iamprivatekeyspecrefs
}
func flattenMapIamQualifierRef(p *models.IamQualifierRef, d *schema.ResourceData) []map[string]interface{} {
	var iamqualifierrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamqualifierref := make(map[string]interface{})
	iamqualifierref["moid"] = item.Moid
	iamqualifierref["object_type"] = item.ObjectType
	iamqualifierref["selector"] = item.Selector

	iamqualifierrefs = append(iamqualifierrefs, iamqualifierref)
	return iamqualifierrefs
}
func flattenMapIamResourceLimitsRef(p *models.IamResourceLimitsRef, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcelimitsrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamresourcelimitsref := make(map[string]interface{})
	iamresourcelimitsref["moid"] = item.Moid
	iamresourcelimitsref["object_type"] = item.ObjectType
	iamresourcelimitsref["selector"] = item.Selector

	iamresourcelimitsrefs = append(iamresourcelimitsrefs, iamresourcelimitsref)
	return iamresourcelimitsrefs
}
func flattenMapIamSecurityHolderRef(p *models.IamSecurityHolderRef, d *schema.ResourceData) []map[string]interface{} {
	var iamsecurityholderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamsecurityholderref := make(map[string]interface{})
	iamsecurityholderref["moid"] = item.Moid
	iamsecurityholderref["object_type"] = item.ObjectType
	iamsecurityholderref["selector"] = item.Selector

	iamsecurityholderrefs = append(iamsecurityholderrefs, iamsecurityholderref)
	return iamsecurityholderrefs
}
func flattenMapIamServiceProviderRef(p *models.IamServiceProviderRef, d *schema.ResourceData) []map[string]interface{} {
	var iamserviceproviderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamserviceproviderref := make(map[string]interface{})
	iamserviceproviderref["moid"] = item.Moid
	iamserviceproviderref["object_type"] = item.ObjectType
	iamserviceproviderref["selector"] = item.Selector

	iamserviceproviderrefs = append(iamserviceproviderrefs, iamserviceproviderref)
	return iamserviceproviderrefs
}
func flattenMapIamSessionLimitsRef(p *models.IamSessionLimitsRef, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionlimitsrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamsessionlimitsref := make(map[string]interface{})
	iamsessionlimitsref["moid"] = item.Moid
	iamsessionlimitsref["object_type"] = item.ObjectType
	iamsessionlimitsref["selector"] = item.Selector

	iamsessionlimitsrefs = append(iamsessionlimitsrefs, iamsessionlimitsref)
	return iamsessionlimitsrefs
}
func flattenMapIamSessionRef(p *models.IamSessionRef, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamsessionref := make(map[string]interface{})
	iamsessionref["moid"] = item.Moid
	iamsessionref["object_type"] = item.ObjectType
	iamsessionref["selector"] = item.Selector

	iamsessionrefs = append(iamsessionrefs, iamsessionref)
	return iamsessionrefs
}
func flattenMapIamSystemRef(p *models.IamSystemRef, d *schema.ResourceData) []map[string]interface{} {
	var iamsystemrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamsystemref := make(map[string]interface{})
	iamsystemref["moid"] = item.Moid
	iamsystemref["object_type"] = item.ObjectType
	iamsystemref["selector"] = item.Selector

	iamsystemrefs = append(iamsystemrefs, iamsystemref)
	return iamsystemrefs
}
func flattenMapIamUserGroupRef(p *models.IamUserGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamusergroupref := make(map[string]interface{})
	iamusergroupref["moid"] = item.Moid
	iamusergroupref["object_type"] = item.ObjectType
	iamusergroupref["selector"] = item.Selector

	iamusergrouprefs = append(iamusergrouprefs, iamusergroupref)
	return iamusergrouprefs
}
func flattenMapIamUserRef(p *models.IamUserRef, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamuserref := make(map[string]interface{})
	iamuserref["moid"] = item.Moid
	iamuserref["object_type"] = item.ObjectType
	iamuserref["selector"] = item.Selector

	iamuserrefs = append(iamuserrefs, iamuserref)
	return iamuserrefs
}
func flattenMapInfraHardwareInfo(p *models.InfraHardwareInfo, d *schema.ResourceData) []map[string]interface{} {
	var infrahardwareinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	infrahardwareinfo := make(map[string]interface{})
	delete(item.InfraHardwareInfoAO1P1.InfraHardwareInfoAO1P1, "ObjectType")
	if len(item.InfraHardwareInfoAO1P1.InfraHardwareInfoAO1P1) != 0 {
		j, err := json.Marshal(item.InfraHardwareInfoAO1P1.InfraHardwareInfoAO1P1)
		if err == nil {
			infrahardwareinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	infrahardwareinfo["cpu_cores"] = item.CPUCores
	infrahardwareinfo["cpu_speed"] = item.CPUSpeed
	infrahardwareinfo["class_id"] = item.ClassID
	infrahardwareinfo["memory_size"] = item.MemorySize
	infrahardwareinfo["object_type"] = item.ObjectType

	infrahardwareinfos = append(infrahardwareinfos, infrahardwareinfo)
	return infrahardwareinfos
}
func flattenMapInventoryBaseRef(p *models.InventoryBaseRef, d *schema.ResourceData) []map[string]interface{} {
	var inventorybaserefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	inventorybaseref := make(map[string]interface{})
	inventorybaseref["moid"] = item.Moid
	inventorybaseref["object_type"] = item.ObjectType
	inventorybaseref["selector"] = item.Selector

	inventorybaserefs = append(inventorybaserefs, inventorybaseref)
	return inventorybaserefs
}
func flattenMapInventoryGenericInventoryHolderRef(p *models.InventoryGenericInventoryHolderRef, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	inventorygenericinventoryholderref := make(map[string]interface{})
	inventorygenericinventoryholderref["moid"] = item.Moid
	inventorygenericinventoryholderref["object_type"] = item.ObjectType
	inventorygenericinventoryholderref["selector"] = item.Selector

	inventorygenericinventoryholderrefs = append(inventorygenericinventoryholderrefs, inventorygenericinventoryholderref)
	return inventorygenericinventoryholderrefs
}
func flattenMapIwotenantTenantRef(p *models.IwotenantTenantRef, d *schema.ResourceData) []map[string]interface{} {
	var iwotenanttenantrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iwotenanttenantref := make(map[string]interface{})
	iwotenanttenantref["moid"] = item.Moid
	iwotenanttenantref["object_type"] = item.ObjectType
	iwotenanttenantref["selector"] = item.Selector

	iwotenanttenantrefs = append(iwotenanttenantrefs, iwotenanttenantref)
	return iwotenanttenantrefs
}
func flattenMapLicenseAccountLicenseDataRef(p *models.LicenseAccountLicenseDataRef, d *schema.ResourceData) []map[string]interface{} {
	var licenseaccountlicensedatarefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	licenseaccountlicensedataref := make(map[string]interface{})
	licenseaccountlicensedataref["moid"] = item.Moid
	licenseaccountlicensedataref["object_type"] = item.ObjectType
	licenseaccountlicensedataref["selector"] = item.Selector

	licenseaccountlicensedatarefs = append(licenseaccountlicensedatarefs, licenseaccountlicensedataref)
	return licenseaccountlicensedatarefs
}
func flattenMapLicenseCustomerOpRef(p *models.LicenseCustomerOpRef, d *schema.ResourceData) []map[string]interface{} {
	var licensecustomeroprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	licensecustomeropref := make(map[string]interface{})
	licensecustomeropref["moid"] = item.Moid
	licensecustomeropref["object_type"] = item.ObjectType
	licensecustomeropref["selector"] = item.Selector

	licensecustomeroprefs = append(licensecustomeroprefs, licensecustomeropref)
	return licensecustomeroprefs
}
func flattenMapLicenseSmartlicenseTokenRef(p *models.LicenseSmartlicenseTokenRef, d *schema.ResourceData) []map[string]interface{} {
	var licensesmartlicensetokenrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	licensesmartlicensetokenref := make(map[string]interface{})
	licensesmartlicensetokenref["moid"] = item.Moid
	licensesmartlicensetokenref["object_type"] = item.ObjectType
	licensesmartlicensetokenref["selector"] = item.Selector

	licensesmartlicensetokenrefs = append(licensesmartlicensetokenrefs, licensesmartlicensetokenref)
	return licensesmartlicensetokenrefs
}
func flattenMapManagementControllerRef(p *models.ManagementControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var managementcontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	managementcontrollerref := make(map[string]interface{})
	managementcontrollerref["moid"] = item.Moid
	managementcontrollerref["object_type"] = item.ObjectType
	managementcontrollerref["selector"] = item.Selector

	managementcontrollerrefs = append(managementcontrollerrefs, managementcontrollerref)
	return managementcontrollerrefs
}
func flattenMapManagementEntityRef(p *models.ManagementEntityRef, d *schema.ResourceData) []map[string]interface{} {
	var managemententityrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	managemententityref := make(map[string]interface{})
	managemententityref["moid"] = item.Moid
	managemententityref["object_type"] = item.ObjectType
	managemententityref["selector"] = item.Selector

	managemententityrefs = append(managemententityrefs, managemententityref)
	return managemententityrefs
}
func flattenMapMemoryArrayRef(p *models.MemoryArrayRef, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memoryarrayref := make(map[string]interface{})
	memoryarrayref["moid"] = item.Moid
	memoryarrayref["object_type"] = item.ObjectType
	memoryarrayref["selector"] = item.Selector

	memoryarrayrefs = append(memoryarrayrefs, memoryarrayref)
	return memoryarrayrefs
}
func flattenMapMemoryPersistentMemoryConfigResultRef(p *models.MemoryPersistentMemoryConfigResultRef, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigresultrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemoryconfigresultref := make(map[string]interface{})
	memorypersistentmemoryconfigresultref["moid"] = item.Moid
	memorypersistentmemoryconfigresultref["object_type"] = item.ObjectType
	memorypersistentmemoryconfigresultref["selector"] = item.Selector

	memorypersistentmemoryconfigresultrefs = append(memorypersistentmemoryconfigresultrefs, memorypersistentmemoryconfigresultref)
	return memorypersistentmemoryconfigresultrefs
}
func flattenMapMemoryPersistentMemoryConfigurationRef(p *models.MemoryPersistentMemoryConfigurationRef, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigurationrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemoryconfigurationref := make(map[string]interface{})
	memorypersistentmemoryconfigurationref["moid"] = item.Moid
	memorypersistentmemoryconfigurationref["object_type"] = item.ObjectType
	memorypersistentmemoryconfigurationref["selector"] = item.Selector

	memorypersistentmemoryconfigurationrefs = append(memorypersistentmemoryconfigurationrefs, memorypersistentmemoryconfigurationref)
	return memorypersistentmemoryconfigurationrefs
}
func flattenMapMemoryPersistentMemoryLocalSecurity(p *models.MemoryPersistentMemoryLocalSecurity, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylocalsecuritys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemorylocalsecurity := make(map[string]interface{})
	delete(item.MemoryPersistentMemoryLocalSecurityAO1P1.MemoryPersistentMemoryLocalSecurityAO1P1, "ObjectType")
	if len(item.MemoryPersistentMemoryLocalSecurityAO1P1.MemoryPersistentMemoryLocalSecurityAO1P1) != 0 {
		j, err := json.Marshal(item.MemoryPersistentMemoryLocalSecurityAO1P1.MemoryPersistentMemoryLocalSecurityAO1P1)
		if err == nil {
			memorypersistentmemorylocalsecurity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	memorypersistentmemorylocalsecurity["class_id"] = item.ClassID
	memorypersistentmemorylocalsecurity["enabled"] = item.Enabled
	memorypersistentmemorylocalsecurity["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	memorypersistentmemorylocalsecurity["object_type"] = item.ObjectType
	memorypersistentmemorylocalsecurity["secure_passphrase"] = item.SecurePassphrase

	memorypersistentmemorylocalsecuritys = append(memorypersistentmemorylocalsecuritys, memorypersistentmemorylocalsecurity)
	return memorypersistentmemorylocalsecuritys
}
func flattenMapMemoryPersistentMemoryRegionRef(p *models.MemoryPersistentMemoryRegionRef, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemoryregionref := make(map[string]interface{})
	memorypersistentmemoryregionref["moid"] = item.Moid
	memorypersistentmemoryregionref["object_type"] = item.ObjectType
	memorypersistentmemoryregionref["selector"] = item.Selector

	memorypersistentmemoryregionrefs = append(memorypersistentmemoryregionrefs, memorypersistentmemoryregionref)
	return memorypersistentmemoryregionrefs
}
func flattenMapMoBaseMoRef(p *models.MoBaseMoRef, d *schema.ResourceData) []map[string]interface{} {
	var mobasemorefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	mobasemoref := make(map[string]interface{})
	mobasemoref["moid"] = item.Moid
	mobasemoref["object_type"] = item.ObjectType
	mobasemoref["selector"] = item.Selector

	mobasemorefs = append(mobasemorefs, mobasemoref)
	return mobasemorefs
}
func flattenMapNetworkElementRef(p *models.NetworkElementRef, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	networkelementref := make(map[string]interface{})
	networkelementref["moid"] = item.Moid
	networkelementref["object_type"] = item.ObjectType
	networkelementref["selector"] = item.Selector

	networkelementrefs = append(networkelementrefs, networkelementref)
	return networkelementrefs
}
func flattenMapNiaapiNewReleaseDetail(p *models.NiaapiNewReleaseDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapinewreleasedetails []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	niaapinewreleasedetail := make(map[string]interface{})
	delete(item.NiaapiNewReleaseDetailAO1P1.NiaapiNewReleaseDetailAO1P1, "ObjectType")
	if len(item.NiaapiNewReleaseDetailAO1P1.NiaapiNewReleaseDetailAO1P1) != 0 {
		j, err := json.Marshal(item.NiaapiNewReleaseDetailAO1P1.NiaapiNewReleaseDetailAO1P1)
		if err == nil {
			niaapinewreleasedetail["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	niaapinewreleasedetail["class_id"] = item.ClassID
	niaapinewreleasedetail["description"] = item.Description
	niaapinewreleasedetail["link"] = item.Link
	niaapinewreleasedetail["object_type"] = item.ObjectType
	niaapinewreleasedetail["release_note_link"] = item.ReleaseNoteLink
	niaapinewreleasedetail["release_note_link_title"] = item.ReleaseNoteLinkTitle
	niaapinewreleasedetail["software_download_link"] = item.SoftwareDownloadLink
	niaapinewreleasedetail["software_download_link_title"] = item.SoftwareDownloadLinkTitle
	niaapinewreleasedetail["title"] = item.Title
	niaapinewreleasedetail["version"] = item.Version

	niaapinewreleasedetails = append(niaapinewreleasedetails, niaapinewreleasedetail)
	return niaapinewreleasedetails
}
func flattenMapNiaapiVersionRegexPlatform(p *models.NiaapiVersionRegexPlatform, d *schema.ResourceData) []map[string]interface{} {
	var niaapiversionregexplatforms []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	niaapiversionregexplatform := make(map[string]interface{})
	delete(item.NiaapiVersionRegexPlatformAO1P1.NiaapiVersionRegexPlatformAO1P1, "ObjectType")
	if len(item.NiaapiVersionRegexPlatformAO1P1.NiaapiVersionRegexPlatformAO1P1) != 0 {
		j, err := json.Marshal(item.NiaapiVersionRegexPlatformAO1P1.NiaapiVersionRegexPlatformAO1P1)
		if err == nil {
			niaapiversionregexplatform["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	niaapiversionregexplatform["anyllregex"] = item.Anyllregex
	niaapiversionregexplatform["class_id"] = item.ClassID
	niaapiversionregexplatform["currentlltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		niaapisoftwareregex := make(map[string]interface{})
		delete(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1, "ObjectType")
		if len(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1) != 0 {
			j, err := json.Marshal(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1)
			if err == nil {
				niaapisoftwareregex["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		niaapisoftwareregex["class_id"] = item.ClassID
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.Currentlltrain, d)
	niaapiversionregexplatform["latestsltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		niaapisoftwareregex := make(map[string]interface{})
		delete(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1, "ObjectType")
		if len(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1) != 0 {
			j, err := json.Marshal(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1)
			if err == nil {
				niaapisoftwareregex["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		niaapisoftwareregex["class_id"] = item.ClassID
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.Latestsltrain, d)
	niaapiversionregexplatform["object_type"] = item.ObjectType
	niaapiversionregexplatform["sltrain"] = (func(p []*models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			niaapisoftwareregex := make(map[string]interface{})
			if len(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1) != 0 {
				j, err := json.Marshal(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1)
				if err == nil {
					niaapisoftwareregex["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			niaapisoftwareregex["class_id"] = item.ClassID
			niaapisoftwareregex["object_type"] = item.ObjectType
			niaapisoftwareregex["regex"] = item.Regex
			niaapisoftwareregex["software_version"] = item.SoftwareVersion
			niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		}
		return niaapisoftwareregexs
	})(item.Sltrain, d)
	niaapiversionregexplatform["upcominglltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		niaapisoftwareregex := make(map[string]interface{})
		delete(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1, "ObjectType")
		if len(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1) != 0 {
			j, err := json.Marshal(item.NiaapiSoftwareRegexAO1P1.NiaapiSoftwareRegexAO1P1)
			if err == nil {
				niaapisoftwareregex["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		niaapisoftwareregex["class_id"] = item.ClassID
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.Upcominglltrain, d)

	niaapiversionregexplatforms = append(niaapiversionregexplatforms, niaapiversionregexplatform)
	return niaapiversionregexplatforms
}
func flattenMapNiatelemetryDiskinfo(p *models.NiatelemetryDiskinfo, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrydiskinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	niatelemetrydiskinfo := make(map[string]interface{})
	delete(item.NiatelemetryDiskinfoAO1P1.NiatelemetryDiskinfoAO1P1, "ObjectType")
	if len(item.NiatelemetryDiskinfoAO1P1.NiatelemetryDiskinfoAO1P1) != 0 {
		j, err := json.Marshal(item.NiatelemetryDiskinfoAO1P1.NiatelemetryDiskinfoAO1P1)
		if err == nil {
			niatelemetrydiskinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	niatelemetrydiskinfo["class_id"] = item.ClassID
	niatelemetrydiskinfo["free"] = item.Free
	niatelemetrydiskinfo["name"] = item.Name
	niatelemetrydiskinfo["object_type"] = item.ObjectType
	niatelemetrydiskinfo["total"] = item.Total
	niatelemetrydiskinfo["used"] = item.Used

	niatelemetrydiskinfos = append(niatelemetrydiskinfos, niatelemetrydiskinfo)
	return niatelemetrydiskinfos
}
func flattenMapNiatelemetryNiaInventoryRef(p *models.NiatelemetryNiaInventoryRef, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetryniainventoryrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	niatelemetryniainventoryref := make(map[string]interface{})
	niatelemetryniainventoryref["moid"] = item.Moid
	niatelemetryniainventoryref["object_type"] = item.ObjectType
	niatelemetryniainventoryref["selector"] = item.Selector

	niatelemetryniainventoryrefs = append(niatelemetryniainventoryrefs, niatelemetryniainventoryref)
	return niatelemetryniainventoryrefs
}
func flattenMapNiatelemetryNiaLicenseStateRef(p *models.NiatelemetryNiaLicenseStateRef, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynialicensestaterefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	niatelemetrynialicensestateref := make(map[string]interface{})
	niatelemetrynialicensestateref["moid"] = item.Moid
	niatelemetrynialicensestateref["object_type"] = item.ObjectType
	niatelemetrynialicensestateref["selector"] = item.Selector

	niatelemetrynialicensestaterefs = append(niatelemetrynialicensestaterefs, niatelemetrynialicensestateref)
	return niatelemetrynialicensestaterefs
}
func flattenMapOnpremSchedule(p *models.OnpremSchedule, d *schema.ResourceData) []map[string]interface{} {
	var onpremschedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	onpremschedule := make(map[string]interface{})
	delete(item.OnpremScheduleAO1P1.OnpremScheduleAO1P1, "ObjectType")
	if len(item.OnpremScheduleAO1P1.OnpremScheduleAO1P1) != 0 {
		j, err := json.Marshal(item.OnpremScheduleAO1P1.OnpremScheduleAO1P1)
		if err == nil {
			onpremschedule["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	onpremschedule["class_id"] = item.ClassID
	onpremschedule["day_of_month"] = item.DayOfMonth
	onpremschedule["day_of_week"] = item.DayOfWeek
	onpremschedule["month_of_year"] = item.MonthOfYear
	onpremschedule["object_type"] = item.ObjectType
	onpremschedule["repeat_interval"] = item.RepeatInterval
	onpremschedule["time_of_day"] = item.TimeOfDay
	onpremschedule["time_zone"] = item.TimeZone
	onpremschedule["week_of_month"] = item.WeekOfMonth

	onpremschedules = append(onpremschedules, onpremschedule)
	return onpremschedules
}
func flattenMapOnpremUpgradePhase(p *models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradephases []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	onpremupgradephase := make(map[string]interface{})
	delete(item.OnpremUpgradePhaseAO1P1.OnpremUpgradePhaseAO1P1, "ObjectType")
	if len(item.OnpremUpgradePhaseAO1P1.OnpremUpgradePhaseAO1P1) != 0 {
		j, err := json.Marshal(item.OnpremUpgradePhaseAO1P1.OnpremUpgradePhaseAO1P1)
		if err == nil {
			onpremupgradephase["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	onpremupgradephase["class_id"] = item.ClassID
	onpremupgradephase["elapsed_time"] = item.ElapsedTime
	onpremupgradephase["end_time"] = item.EndTime
	onpremupgradephase["failed"] = item.Failed
	onpremupgradephase["message"] = item.Message
	onpremupgradephase["name"] = item.Name
	onpremupgradephase["object_type"] = item.ObjectType
	onpremupgradephase["start_time"] = item.StartTime

	onpremupgradephases = append(onpremupgradephases, onpremupgradephase)
	return onpremupgradephases
}
func flattenMapOrganizationOrganizationRef(p *models.OrganizationOrganizationRef, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	organizationorganizationref := make(map[string]interface{})
	organizationorganizationref["moid"] = item.Moid
	organizationorganizationref["object_type"] = item.ObjectType
	organizationorganizationref["selector"] = item.Selector

	organizationorganizationrefs = append(organizationorganizationrefs, organizationorganizationref)
	return organizationorganizationrefs
}
func flattenMapOsAnswers(p *models.OsAnswers, d *schema.ResourceData) []map[string]interface{} {
	var osanswerss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	osanswers := make(map[string]interface{})
	delete(item.OsAnswersAO1P1.OsAnswersAO1P1, "ObjectType")
	if len(item.OsAnswersAO1P1.OsAnswersAO1P1) != 0 {
		j, err := json.Marshal(item.OsAnswersAO1P1.OsAnswersAO1P1)
		if err == nil {
			osanswers["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	osanswers["answer_file"] = item.AnswerFile
	osanswers["class_id"] = item.ClassID
	osanswers["hostname"] = item.Hostname
	osanswers["ip_config_type"] = item.IPConfigType
	osanswers["ip_configuration"] = (func(p *models.OsIPConfiguration, d *schema.ResourceData) []map[string]interface{} {
		var osipconfigurations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		osipconfiguration := make(map[string]interface{})
		delete(item.AO1, "ObjectType")
		if len(item.AO1) != 0 {
			j, err := json.Marshal(item.AO1)
			if err == nil {
				osipconfiguration["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		osipconfiguration["class_id"] = item.ClassID
		osipconfiguration["object_type"] = item.ObjectType

		osipconfigurations = append(osipconfigurations, osipconfiguration)
		return osipconfigurations
	})(item.IPConfiguration, d)
	osanswers["is_answer_file_set"] = item.IsAnswerFileSet
	osanswers["is_root_password_crypted"] = item.IsRootPasswordCrypted
	osanswers["is_root_password_set"] = item.IsRootPasswordSet
	osanswers["nameserver"] = item.Nameserver
	osanswers["object_type"] = item.ObjectType
	osanswers["product_key"] = item.ProductKey
	osanswers["root_password"] = item.RootPassword
	osanswers["source"] = item.Source

	osanswerss = append(osanswerss, osanswers)
	return osanswerss
}
func flattenMapOsCatalogRef(p *models.OsCatalogRef, d *schema.ResourceData) []map[string]interface{} {
	var oscatalogrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	oscatalogref := make(map[string]interface{})
	oscatalogref["moid"] = item.Moid
	oscatalogref["object_type"] = item.ObjectType
	oscatalogref["selector"] = item.Selector

	oscatalogrefs = append(oscatalogrefs, oscatalogref)
	return oscatalogrefs
}
func flattenMapOsConfigurationFileRef(p *models.OsConfigurationFileRef, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	osconfigurationfileref := make(map[string]interface{})
	osconfigurationfileref["moid"] = item.Moid
	osconfigurationfileref["object_type"] = item.ObjectType
	osconfigurationfileref["selector"] = item.Selector

	osconfigurationfilerefs = append(osconfigurationfilerefs, osconfigurationfileref)
	return osconfigurationfilerefs
}
func flattenMapOsOperatingSystemParameters(p *models.OsOperatingSystemParameters, d *schema.ResourceData) []map[string]interface{} {
	var osoperatingsystemparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	osoperatingsystemparameters := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			osoperatingsystemparameters["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	osoperatingsystemparameters["class_id"] = item.ClassID
	osoperatingsystemparameters["object_type"] = item.ObjectType

	osoperatingsystemparameterss = append(osoperatingsystemparameterss, osoperatingsystemparameters)
	return osoperatingsystemparameterss
}
func flattenMapPciSwitchRef(p *models.PciSwitchRef, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	pciswitchref := make(map[string]interface{})
	pciswitchref["moid"] = item.Moid
	pciswitchref["object_type"] = item.ObjectType
	pciswitchref["selector"] = item.Selector

	pciswitchrefs = append(pciswitchrefs, pciswitchref)
	return pciswitchrefs
}
func flattenMapPkixDistinguishedName(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
	var pkixdistinguishednames []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	pkixdistinguishedname := make(map[string]interface{})
	delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
	if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
		j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
		if err == nil {
			pkixdistinguishedname["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	pkixdistinguishedname["class_id"] = item.ClassID
	pkixdistinguishedname["common_name"] = item.CommonName
	pkixdistinguishedname["country"] = item.Country
	pkixdistinguishedname["locality"] = item.Locality
	pkixdistinguishedname["object_type"] = item.ObjectType
	pkixdistinguishedname["organization"] = item.Organization
	pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
	pkixdistinguishedname["state"] = item.State

	pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
	return pkixdistinguishednames
}
func flattenMapPkixKeyGenerationSpec(p *models.PkixKeyGenerationSpec, d *schema.ResourceData) []map[string]interface{} {
	var pkixkeygenerationspecs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	pkixkeygenerationspec := make(map[string]interface{})
	delete(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1, "ObjectType")
	if len(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1) != 0 {
		j, err := json.Marshal(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1)
		if err == nil {
			pkixkeygenerationspec["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	pkixkeygenerationspec["class_id"] = item.ClassID
	pkixkeygenerationspec["name"] = item.Name
	pkixkeygenerationspec["object_type"] = item.ObjectType

	pkixkeygenerationspecs = append(pkixkeygenerationspecs, pkixkeygenerationspec)
	return pkixkeygenerationspecs
}
func flattenMapPkixSubjectAlternateName(p *models.PkixSubjectAlternateName, d *schema.ResourceData) []map[string]interface{} {
	var pkixsubjectalternatenames []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	pkixsubjectalternatename := make(map[string]interface{})
	delete(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1, "ObjectType")
	if len(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1) != 0 {
		j, err := json.Marshal(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1)
		if err == nil {
			pkixsubjectalternatename["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	pkixsubjectalternatename["class_id"] = item.ClassID
	pkixsubjectalternatename["dns_name"] = item.DNSName
	pkixsubjectalternatename["email_address"] = item.EmailAddress
	pkixsubjectalternatename["ip_address"] = item.IPAddress
	pkixsubjectalternatename["object_type"] = item.ObjectType
	pkixsubjectalternatename["uri"] = item.URI

	pkixsubjectalternatenames = append(pkixsubjectalternatenames, pkixsubjectalternatename)
	return pkixsubjectalternatenames
}
func flattenMapPolicyAbstractProfileRef(p *models.PolicyAbstractProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	policyabstractprofileref := make(map[string]interface{})
	policyabstractprofileref["moid"] = item.Moid
	policyabstractprofileref["object_type"] = item.ObjectType
	policyabstractprofileref["selector"] = item.Selector

	policyabstractprofilerefs = append(policyabstractprofilerefs, policyabstractprofileref)
	return policyabstractprofilerefs
}
func flattenMapPolicyConfigChange(p *models.PolicyConfigChange, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigchanges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	policyconfigchange := make(map[string]interface{})
	delete(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1, "ObjectType")
	if len(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1) != 0 {
		j, err := json.Marshal(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1)
		if err == nil {
			policyconfigchange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	policyconfigchange["changes"] = item.Changes
	policyconfigchange["class_id"] = item.ClassID
	policyconfigchange["disruptions"] = item.Disruptions
	policyconfigchange["object_type"] = item.ObjectType

	policyconfigchanges = append(policyconfigchanges, policyconfigchange)
	return policyconfigchanges
}
func flattenMapPolicyConfigContext(p *models.PolicyConfigContext, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigcontexts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	policyconfigcontext := make(map[string]interface{})
	delete(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1, "ObjectType")
	if len(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1) != 0 {
		j, err := json.Marshal(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1)
		if err == nil {
			policyconfigcontext["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	policyconfigcontext["class_id"] = item.ClassID
	policyconfigcontext["config_state"] = item.ConfigState
	policyconfigcontext["control_action"] = item.ControlAction
	policyconfigcontext["error_state"] = item.ErrorState
	policyconfigcontext["object_type"] = item.ObjectType
	policyconfigcontext["oper_state"] = item.OperState

	policyconfigcontexts = append(policyconfigcontexts, policyconfigcontext)
	return policyconfigcontexts
}
func flattenMapPolicyConfigResultContext(p *models.PolicyConfigResultContext, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigresultcontexts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	policyconfigresultcontext := make(map[string]interface{})
	delete(item.PolicyConfigResultContextAO1P1.PolicyConfigResultContextAO1P1, "ObjectType")
	if len(item.PolicyConfigResultContextAO1P1.PolicyConfigResultContextAO1P1) != 0 {
		j, err := json.Marshal(item.PolicyConfigResultContextAO1P1.PolicyConfigResultContextAO1P1)
		if err == nil {
			policyconfigresultcontext["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	policyconfigresultcontext["class_id"] = item.ClassID
	policyconfigresultcontext["entity_data"] = item.EntityData
	policyconfigresultcontext["entity_moid"] = item.EntityMoid
	policyconfigresultcontext["entity_name"] = item.EntityName
	policyconfigresultcontext["entity_type"] = item.EntityType
	policyconfigresultcontext["object_type"] = item.ObjectType

	policyconfigresultcontexts = append(policyconfigresultcontexts, policyconfigresultcontext)
	return policyconfigresultcontexts
}
func flattenMapPortGroupRef(p *models.PortGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	portgroupref := make(map[string]interface{})
	portgroupref["moid"] = item.Moid
	portgroupref["object_type"] = item.ObjectType
	portgroupref["selector"] = item.Selector

	portgrouprefs = append(portgrouprefs, portgroupref)
	return portgrouprefs
}
func flattenMapPortSubGroupRef(p *models.PortSubGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	portsubgroupref := make(map[string]interface{})
	portsubgroupref["moid"] = item.Moid
	portsubgroupref["object_type"] = item.ObjectType
	portsubgroupref["selector"] = item.Selector

	portsubgrouprefs = append(portsubgrouprefs, portsubgroupref)
	return portsubgrouprefs
}
func flattenMapRecoveryAbstractBackupInfoRef(p *models.RecoveryAbstractBackupInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var recoveryabstractbackupinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoveryabstractbackupinforef := make(map[string]interface{})
	recoveryabstractbackupinforef["moid"] = item.Moid
	recoveryabstractbackupinforef["object_type"] = item.ObjectType
	recoveryabstractbackupinforef["selector"] = item.Selector

	recoveryabstractbackupinforefs = append(recoveryabstractbackupinforefs, recoveryabstractbackupinforef)
	return recoveryabstractbackupinforefs
}
func flattenMapRecoveryAbstractRestoreStatusRef(p *models.RecoveryAbstractRestoreStatusRef, d *schema.ResourceData) []map[string]interface{} {
	var recoveryabstractrestorestatusrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoveryabstractrestorestatusref := make(map[string]interface{})
	recoveryabstractrestorestatusref["moid"] = item.Moid
	recoveryabstractrestorestatusref["object_type"] = item.ObjectType
	recoveryabstractrestorestatusref["selector"] = item.Selector

	recoveryabstractrestorestatusrefs = append(recoveryabstractrestorestatusrefs, recoveryabstractrestorestatusref)
	return recoveryabstractrestorestatusrefs
}
func flattenMapRecoveryBackupConfigPolicyRef(p *models.RecoveryBackupConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupconfigpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoverybackupconfigpolicyref := make(map[string]interface{})
	recoverybackupconfigpolicyref["moid"] = item.Moid
	recoverybackupconfigpolicyref["object_type"] = item.ObjectType
	recoverybackupconfigpolicyref["selector"] = item.Selector

	recoverybackupconfigpolicyrefs = append(recoverybackupconfigpolicyrefs, recoverybackupconfigpolicyref)
	return recoverybackupconfigpolicyrefs
}
func flattenMapRecoveryBackupProfileRef(p *models.RecoveryBackupProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoverybackupprofileref := make(map[string]interface{})
	recoverybackupprofileref["moid"] = item.Moid
	recoverybackupprofileref["object_type"] = item.ObjectType
	recoverybackupprofileref["selector"] = item.Selector

	recoverybackupprofilerefs = append(recoverybackupprofilerefs, recoverybackupprofileref)
	return recoverybackupprofilerefs
}
func flattenMapRecoveryBackupSchedule(p *models.RecoveryBackupSchedule, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupschedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoverybackupschedule := make(map[string]interface{})
	delete(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1, "ObjectType")
	if len(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1) != 0 {
		j, err := json.Marshal(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1)
		if err == nil {
			recoverybackupschedule["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	recoverybackupschedule["class_id"] = item.ClassID
	recoverybackupschedule["execution_time"] = item.ExecutionTime
	recoverybackupschedule["frequency_unit"] = item.FrequencyUnit
	recoverybackupschedule["hours"] = item.Hours
	recoverybackupschedule["object_type"] = item.ObjectType

	recoverybackupschedules = append(recoverybackupschedules, recoverybackupschedule)
	return recoverybackupschedules
}
func flattenMapRecoveryConfigParams(p *models.RecoveryConfigParams, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigparamss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoveryconfigparams := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			recoveryconfigparams["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	recoveryconfigparams["class_id"] = item.ClassID
	recoveryconfigparams["object_type"] = item.ObjectType

	recoveryconfigparamss = append(recoveryconfigparamss, recoveryconfigparams)
	return recoveryconfigparamss
}
func flattenMapRecoveryConfigResultRef(p *models.RecoveryConfigResultRef, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoveryconfigresultref := make(map[string]interface{})
	recoveryconfigresultref["moid"] = item.Moid
	recoveryconfigresultref["object_type"] = item.ObjectType
	recoveryconfigresultref["selector"] = item.Selector

	recoveryconfigresultrefs = append(recoveryconfigresultrefs, recoveryconfigresultref)
	return recoveryconfigresultrefs
}
func flattenMapRecoveryOnDemandBackupRef(p *models.RecoveryOnDemandBackupRef, d *schema.ResourceData) []map[string]interface{} {
	var recoveryondemandbackuprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoveryondemandbackupref := make(map[string]interface{})
	recoveryondemandbackupref["moid"] = item.Moid
	recoveryondemandbackupref["object_type"] = item.ObjectType
	recoveryondemandbackupref["selector"] = item.Selector

	recoveryondemandbackuprefs = append(recoveryondemandbackuprefs, recoveryondemandbackupref)
	return recoveryondemandbackuprefs
}
func flattenMapRecoveryScheduleConfigPolicyRef(p *models.RecoveryScheduleConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var recoveryscheduleconfigpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoveryscheduleconfigpolicyref := make(map[string]interface{})
	recoveryscheduleconfigpolicyref["moid"] = item.Moid
	recoveryscheduleconfigpolicyref["object_type"] = item.ObjectType
	recoveryscheduleconfigpolicyref["selector"] = item.Selector

	recoveryscheduleconfigpolicyrefs = append(recoveryscheduleconfigpolicyrefs, recoveryscheduleconfigpolicyref)
	return recoveryscheduleconfigpolicyrefs
}
func flattenMapResourceGroupRef(p *models.ResourceGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	resourcegroupref := make(map[string]interface{})
	resourcegroupref["moid"] = item.Moid
	resourcegroupref["object_type"] = item.ObjectType
	resourcegroupref["selector"] = item.Selector

	resourcegrouprefs = append(resourcegrouprefs, resourcegroupref)
	return resourcegrouprefs
}
func flattenMapResourceMembershipHolderRef(p *models.ResourceMembershipHolderRef, d *schema.ResourceData) []map[string]interface{} {
	var resourcemembershipholderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	resourcemembershipholderref := make(map[string]interface{})
	resourcemembershipholderref["moid"] = item.Moid
	resourcemembershipholderref["object_type"] = item.ObjectType
	resourcemembershipholderref["selector"] = item.Selector

	resourcemembershipholderrefs = append(resourcemembershipholderrefs, resourcemembershipholderref)
	return resourcemembershipholderrefs
}
func flattenMapSdwanProfileRef(p *models.SdwanProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sdwanprofileref := make(map[string]interface{})
	sdwanprofileref["moid"] = item.Moid
	sdwanprofileref["object_type"] = item.ObjectType
	sdwanprofileref["selector"] = item.Selector

	sdwanprofilerefs = append(sdwanprofilerefs, sdwanprofileref)
	return sdwanprofilerefs
}
func flattenMapSdwanRouterPolicyRef(p *models.SdwanRouterPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouterpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sdwanrouterpolicyref := make(map[string]interface{})
	sdwanrouterpolicyref["moid"] = item.Moid
	sdwanrouterpolicyref["object_type"] = item.ObjectType
	sdwanrouterpolicyref["selector"] = item.Selector

	sdwanrouterpolicyrefs = append(sdwanrouterpolicyrefs, sdwanrouterpolicyref)
	return sdwanrouterpolicyrefs
}
func flattenMapSdwanVmanageAccountPolicyRef(p *models.SdwanVmanageAccountPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var sdwanvmanageaccountpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sdwanvmanageaccountpolicyref := make(map[string]interface{})
	sdwanvmanageaccountpolicyref["moid"] = item.Moid
	sdwanvmanageaccountpolicyref["object_type"] = item.ObjectType
	sdwanvmanageaccountpolicyref["selector"] = item.Selector

	sdwanvmanageaccountpolicyrefs = append(sdwanvmanageaccountpolicyrefs, sdwanvmanageaccountpolicyref)
	return sdwanvmanageaccountpolicyrefs
}
func flattenMapServerConfigResultRef(p *models.ServerConfigResultRef, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	serverconfigresultref := make(map[string]interface{})
	serverconfigresultref["moid"] = item.Moid
	serverconfigresultref["object_type"] = item.ObjectType
	serverconfigresultref["selector"] = item.Selector

	serverconfigresultrefs = append(serverconfigresultrefs, serverconfigresultref)
	return serverconfigresultrefs
}
func flattenMapServerProfileRef(p *models.ServerProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var serverprofilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	serverprofileref := make(map[string]interface{})
	serverprofileref["moid"] = item.Moid
	serverprofileref["object_type"] = item.ObjectType
	serverprofileref["selector"] = item.Selector

	serverprofilerefs = append(serverprofilerefs, serverprofileref)
	return serverprofilerefs
}
func flattenMapSoftwareHyperflexDistributableRef(p *models.SoftwareHyperflexDistributableRef, d *schema.ResourceData) []map[string]interface{} {
	var softwarehyperflexdistributablerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	softwarehyperflexdistributableref := make(map[string]interface{})
	softwarehyperflexdistributableref["moid"] = item.Moid
	softwarehyperflexdistributableref["object_type"] = item.ObjectType
	softwarehyperflexdistributableref["selector"] = item.Selector

	softwarehyperflexdistributablerefs = append(softwarehyperflexdistributablerefs, softwarehyperflexdistributableref)
	return softwarehyperflexdistributablerefs
}
func flattenMapSoftwareSolutionDistributableRef(p *models.SoftwareSolutionDistributableRef, d *schema.ResourceData) []map[string]interface{} {
	var softwaresolutiondistributablerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	softwaresolutiondistributableref := make(map[string]interface{})
	softwaresolutiondistributableref["moid"] = item.Moid
	softwaresolutiondistributableref["object_type"] = item.ObjectType
	softwaresolutiondistributableref["selector"] = item.Selector

	softwaresolutiondistributablerefs = append(softwaresolutiondistributablerefs, softwaresolutiondistributableref)
	return softwaresolutiondistributablerefs
}
func flattenMapSoftwarerepositoryCatalogRef(p *models.SoftwarerepositoryCatalogRef, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositorycatalogrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	softwarerepositorycatalogref := make(map[string]interface{})
	softwarerepositorycatalogref["moid"] = item.Moid
	softwarerepositorycatalogref["object_type"] = item.ObjectType
	softwarerepositorycatalogref["selector"] = item.Selector

	softwarerepositorycatalogrefs = append(softwarerepositorycatalogrefs, softwarerepositorycatalogref)
	return softwarerepositorycatalogrefs
}
func flattenMapSoftwarerepositoryFileServer(p *models.SoftwarerepositoryFileServer, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryfileservers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	softwarerepositoryfileserver := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			softwarerepositoryfileserver["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	softwarerepositoryfileserver["class_id"] = item.ClassID
	softwarerepositoryfileserver["object_type"] = item.ObjectType

	softwarerepositoryfileservers = append(softwarerepositoryfileservers, softwarerepositoryfileserver)
	return softwarerepositoryfileservers
}
func flattenMapSoftwarerepositoryOperatingSystemFileRef(p *models.SoftwarerepositoryOperatingSystemFileRef, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryoperatingsystemfilerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	softwarerepositoryoperatingsystemfileref := make(map[string]interface{})
	softwarerepositoryoperatingsystemfileref["moid"] = item.Moid
	softwarerepositoryoperatingsystemfileref["object_type"] = item.ObjectType
	softwarerepositoryoperatingsystemfileref["selector"] = item.Selector

	softwarerepositoryoperatingsystemfilerefs = append(softwarerepositoryoperatingsystemfilerefs, softwarerepositoryoperatingsystemfileref)
	return softwarerepositoryoperatingsystemfilerefs
}
func flattenMapStorageBaseCapacity(p *models.StorageBaseCapacity, d *schema.ResourceData) []map[string]interface{} {
	var storagebasecapacitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagebasecapacity := make(map[string]interface{})
	delete(item.StorageBaseCapacityAO1P1.StorageBaseCapacityAO1P1, "ObjectType")
	if len(item.StorageBaseCapacityAO1P1.StorageBaseCapacityAO1P1) != 0 {
		j, err := json.Marshal(item.StorageBaseCapacityAO1P1.StorageBaseCapacityAO1P1)
		if err == nil {
			storagebasecapacity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	storagebasecapacity["available"] = item.Available
	storagebasecapacity["class_id"] = item.ClassID
	storagebasecapacity["free"] = item.Free
	storagebasecapacity["object_type"] = item.ObjectType
	storagebasecapacity["total"] = item.Total
	storagebasecapacity["used"] = item.Used

	storagebasecapacitys = append(storagebasecapacitys, storagebasecapacity)
	return storagebasecapacitys
}
func flattenMapStorageControllerRef(p *models.StorageControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagecontrollerref := make(map[string]interface{})
	storagecontrollerref["moid"] = item.Moid
	storagecontrollerref["object_type"] = item.ObjectType
	storagecontrollerref["selector"] = item.Selector

	storagecontrollerrefs = append(storagecontrollerrefs, storagecontrollerref)
	return storagecontrollerrefs
}
func flattenMapStorageEnclosureRef(p *models.StorageEnclosureRef, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageenclosureref := make(map[string]interface{})
	storageenclosureref["moid"] = item.Moid
	storageenclosureref["object_type"] = item.ObjectType
	storageenclosureref["selector"] = item.Selector

	storageenclosurerefs = append(storageenclosurerefs, storageenclosureref)
	return storageenclosurerefs
}
func flattenMapStorageFlexFlashControllerRef(p *models.StorageFlexFlashControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageflexflashcontrollerref := make(map[string]interface{})
	storageflexflashcontrollerref["moid"] = item.Moid
	storageflexflashcontrollerref["object_type"] = item.ObjectType
	storageflexflashcontrollerref["selector"] = item.Selector

	storageflexflashcontrollerrefs = append(storageflexflashcontrollerrefs, storageflexflashcontrollerref)
	return storageflexflashcontrollerrefs
}
func flattenMapStorageFlexUtilControllerRef(p *models.StorageFlexUtilControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageflexutilcontrollerref := make(map[string]interface{})
	storageflexutilcontrollerref["moid"] = item.Moid
	storageflexutilcontrollerref["object_type"] = item.ObjectType
	storageflexutilcontrollerref["selector"] = item.Selector

	storageflexutilcontrollerrefs = append(storageflexutilcontrollerrefs, storageflexutilcontrollerref)
	return storageflexutilcontrollerrefs
}
func flattenMapStoragePhysicalDiskRef(p *models.StoragePhysicalDiskRef, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagephysicaldiskref := make(map[string]interface{})
	storagephysicaldiskref["moid"] = item.Moid
	storagephysicaldiskref["object_type"] = item.ObjectType
	storagephysicaldiskref["selector"] = item.Selector

	storagephysicaldiskrefs = append(storagephysicaldiskrefs, storagephysicaldiskref)
	return storagephysicaldiskrefs
}
func flattenMapStoragePureArrayRef(p *models.StoragePureArrayRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurearrayrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepurearrayref := make(map[string]interface{})
	storagepurearrayref["moid"] = item.Moid
	storagepurearrayref["object_type"] = item.ObjectType
	storagepurearrayref["selector"] = item.Selector

	storagepurearrayrefs = append(storagepurearrayrefs, storagepurearrayref)
	return storagepurearrayrefs
}
func flattenMapStoragePureControllerRef(p *models.StoragePureControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurecontrollerrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepurecontrollerref := make(map[string]interface{})
	storagepurecontrollerref["moid"] = item.Moid
	storagepurecontrollerref["object_type"] = item.ObjectType
	storagepurecontrollerref["selector"] = item.Selector

	storagepurecontrollerrefs = append(storagepurecontrollerrefs, storagepurecontrollerref)
	return storagepurecontrollerrefs
}
func flattenMapStoragePureHostGroupRef(p *models.StoragePureHostGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepurehostgroupref := make(map[string]interface{})
	storagepurehostgroupref["moid"] = item.Moid
	storagepurehostgroupref["object_type"] = item.ObjectType
	storagepurehostgroupref["selector"] = item.Selector

	storagepurehostgrouprefs = append(storagepurehostgrouprefs, storagepurehostgroupref)
	return storagepurehostgrouprefs
}
func flattenMapStoragePureHostRef(p *models.StoragePureHostRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepurehostref := make(map[string]interface{})
	storagepurehostref["moid"] = item.Moid
	storagepurehostref["object_type"] = item.ObjectType
	storagepurehostref["selector"] = item.Selector

	storagepurehostrefs = append(storagepurehostrefs, storagepurehostref)
	return storagepurehostrefs
}
func flattenMapStoragePureHostUtilization(p *models.StoragePureHostUtilization, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostutilizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepurehostutilization := make(map[string]interface{})
	delete(item.StorageBaseCapacityAO1P1.StorageBaseCapacityAO1P1, "ObjectType")
	if len(item.StorageBaseCapacityAO1P1.StorageBaseCapacityAO1P1) != 0 {
		j, err := json.Marshal(item.StorageBaseCapacityAO1P1.StorageBaseCapacityAO1P1)
		if err == nil {
			storagepurehostutilization["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	storagepurehostutilization["available"] = item.Available
	storagepurehostutilization["class_id"] = item.ClassID
	storagepurehostutilization["data_reduction"] = item.DataReduction
	storagepurehostutilization["free"] = item.Free
	storagepurehostutilization["object_type"] = item.ObjectType
	storagepurehostutilization["snapshot"] = item.Snapshot
	storagepurehostutilization["thin_provisioned"] = item.ThinProvisioned
	storagepurehostutilization["total"] = item.Total
	storagepurehostutilization["total_reduction"] = item.TotalReduction
	storagepurehostutilization["used"] = item.Used
	storagepurehostutilization["volume"] = item.Volume

	storagepurehostutilizations = append(storagepurehostutilizations, storagepurehostutilization)
	return storagepurehostutilizations
}
func flattenMapStoragePureProtectionGroupRef(p *models.StoragePureProtectionGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepureprotectiongrouprefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepureprotectiongroupref := make(map[string]interface{})
	storagepureprotectiongroupref["moid"] = item.Moid
	storagepureprotectiongroupref["object_type"] = item.ObjectType
	storagepureprotectiongroupref["selector"] = item.Selector

	storagepureprotectiongrouprefs = append(storagepureprotectiongrouprefs, storagepureprotectiongroupref)
	return storagepureprotectiongrouprefs
}
func flattenMapStoragePureProtectionGroupSnapshotRef(p *models.StoragePureProtectionGroupSnapshotRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepureprotectiongroupsnapshotrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepureprotectiongroupsnapshotref := make(map[string]interface{})
	storagepureprotectiongroupsnapshotref["moid"] = item.Moid
	storagepureprotectiongroupsnapshotref["object_type"] = item.ObjectType
	storagepureprotectiongroupsnapshotref["selector"] = item.Selector

	storagepureprotectiongroupsnapshotrefs = append(storagepureprotectiongroupsnapshotrefs, storagepureprotectiongroupsnapshotref)
	return storagepureprotectiongroupsnapshotrefs
}
func flattenMapStoragePureVolumeRef(p *models.StoragePureVolumeRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepurevolumerefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagepurevolumeref := make(map[string]interface{})
	storagepurevolumeref["moid"] = item.Moid
	storagepurevolumeref["object_type"] = item.ObjectType
	storagepurevolumeref["selector"] = item.Selector

	storagepurevolumerefs = append(storagepurevolumerefs, storagepurevolumeref)
	return storagepurevolumerefs
}
func flattenMapStorageSasExpanderRef(p *models.StorageSasExpanderRef, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagesasexpanderref := make(map[string]interface{})
	storagesasexpanderref["moid"] = item.Moid
	storagesasexpanderref["object_type"] = item.ObjectType
	storagesasexpanderref["selector"] = item.Selector

	storagesasexpanderrefs = append(storagesasexpanderrefs, storagesasexpanderref)
	return storagesasexpanderrefs
}
func flattenMapStorageVirtualDriveExtensionRef(p *models.StorageVirtualDriveExtensionRef, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagevirtualdriveextensionref := make(map[string]interface{})
	storagevirtualdriveextensionref["moid"] = item.Moid
	storagevirtualdriveextensionref["object_type"] = item.ObjectType
	storagevirtualdriveextensionref["selector"] = item.Selector

	storagevirtualdriveextensionrefs = append(storagevirtualdriveextensionrefs, storagevirtualdriveextensionref)
	return storagevirtualdriveextensionrefs
}
func flattenMapStorageVirtualDriveRef(p *models.StorageVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagevirtualdriveref := make(map[string]interface{})
	storagevirtualdriveref["moid"] = item.Moid
	storagevirtualdriveref["object_type"] = item.ObjectType
	storagevirtualdriveref["selector"] = item.Selector

	storagevirtualdriverefs = append(storagevirtualdriverefs, storagevirtualdriveref)
	return storagevirtualdriverefs
}
func flattenMapTamBaseAdvisoryDetails(p *models.TamBaseAdvisoryDetails, d *schema.ResourceData) []map[string]interface{} {
	var tambaseadvisorydetailss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	tambaseadvisorydetails := make(map[string]interface{})
	delete(item.TamBaseAdvisoryDetailsAO1P1.TamBaseAdvisoryDetailsAO1P1, "ObjectType")
	if len(item.TamBaseAdvisoryDetailsAO1P1.TamBaseAdvisoryDetailsAO1P1) != 0 {
		j, err := json.Marshal(item.TamBaseAdvisoryDetailsAO1P1.TamBaseAdvisoryDetailsAO1P1)
		if err == nil {
			tambaseadvisorydetails["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	tambaseadvisorydetails["class_id"] = item.ClassID
	tambaseadvisorydetails["description"] = item.Description
	tambaseadvisorydetails["object_type"] = item.ObjectType

	tambaseadvisorydetailss = append(tambaseadvisorydetailss, tambaseadvisorydetails)
	return tambaseadvisorydetailss
}
func flattenMapTamBaseAdvisoryRef(p *models.TamBaseAdvisoryRef, d *schema.ResourceData) []map[string]interface{} {
	var tambaseadvisoryrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	tambaseadvisoryref := make(map[string]interface{})
	tambaseadvisoryref["moid"] = item.Moid
	tambaseadvisoryref["object_type"] = item.ObjectType
	tambaseadvisoryref["selector"] = item.Selector

	tambaseadvisoryrefs = append(tambaseadvisoryrefs, tambaseadvisoryref)
	return tambaseadvisoryrefs
}
func flattenMapTamSeverity(p *models.TamSeverity, d *schema.ResourceData) []map[string]interface{} {
	var tamseveritys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	tamseverity := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			tamseverity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	tamseverity["class_id"] = item.ClassID
	tamseverity["object_type"] = item.ObjectType

	tamseveritys = append(tamseveritys, tamseverity)
	return tamseveritys
}
func flattenMapTopSystemRef(p *models.TopSystemRef, d *schema.ResourceData) []map[string]interface{} {
	var topsystemrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	topsystemref := make(map[string]interface{})
	topsystemref["moid"] = item.Moid
	topsystemref["object_type"] = item.ObjectType
	topsystemref["selector"] = item.Selector

	topsystemrefs = append(topsystemrefs, topsystemref)
	return topsystemrefs
}
func flattenMapVirtualizationComputeCapacity(p *models.VirtualizationComputeCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcomputecapacitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationcomputecapacity := make(map[string]interface{})
	delete(item.VirtualizationComputeCapacityAO1P1.VirtualizationComputeCapacityAO1P1, "ObjectType")
	if len(item.VirtualizationComputeCapacityAO1P1.VirtualizationComputeCapacityAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationComputeCapacityAO1P1.VirtualizationComputeCapacityAO1P1)
		if err == nil {
			virtualizationcomputecapacity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationcomputecapacity["capacity"] = item.Capacity
	virtualizationcomputecapacity["class_id"] = item.ClassID
	virtualizationcomputecapacity["free"] = item.Free
	virtualizationcomputecapacity["object_type"] = item.ObjectType
	virtualizationcomputecapacity["used"] = item.Used

	virtualizationcomputecapacitys = append(virtualizationcomputecapacitys, virtualizationcomputecapacity)
	return virtualizationcomputecapacitys
}
func flattenMapVirtualizationCPUInfo(p *models.VirtualizationCPUInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcpuinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationcpuinfo := make(map[string]interface{})
	delete(item.VirtualizationCPUInfoAO1P1.VirtualizationCPUInfoAO1P1, "ObjectType")
	if len(item.VirtualizationCPUInfoAO1P1.VirtualizationCPUInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationCPUInfoAO1P1.VirtualizationCPUInfoAO1P1)
		if err == nil {
			virtualizationcpuinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationcpuinfo["class_id"] = item.ClassID
	virtualizationcpuinfo["cores"] = item.Cores
	virtualizationcpuinfo["description"] = item.Description
	virtualizationcpuinfo["object_type"] = item.ObjectType
	virtualizationcpuinfo["sockets"] = item.Sockets
	virtualizationcpuinfo["speed"] = item.Speed
	virtualizationcpuinfo["vendor"] = item.Vendor

	virtualizationcpuinfos = append(virtualizationcpuinfos, virtualizationcpuinfo)
	return virtualizationcpuinfos
}
func flattenMapVirtualizationGuestInfo(p *models.VirtualizationGuestInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationguestinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationguestinfo := make(map[string]interface{})
	delete(item.VirtualizationGuestInfoAO1P1.VirtualizationGuestInfoAO1P1, "ObjectType")
	if len(item.VirtualizationGuestInfoAO1P1.VirtualizationGuestInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationGuestInfoAO1P1.VirtualizationGuestInfoAO1P1)
		if err == nil {
			virtualizationguestinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationguestinfo["class_id"] = item.ClassID
	virtualizationguestinfo["hostname"] = item.Hostname
	virtualizationguestinfo["ip_address"] = item.IPAddress
	virtualizationguestinfo["name"] = item.Name
	virtualizationguestinfo["object_type"] = item.ObjectType
	virtualizationguestinfo["operating_system"] = item.OperatingSystem

	virtualizationguestinfos = append(virtualizationguestinfos, virtualizationguestinfo)
	return virtualizationguestinfos
}
func flattenMapVirtualizationMemoryCapacity(p *models.VirtualizationMemoryCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationmemorycapacitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationmemorycapacity := make(map[string]interface{})
	delete(item.VirtualizationMemoryCapacityAO1P1.VirtualizationMemoryCapacityAO1P1, "ObjectType")
	if len(item.VirtualizationMemoryCapacityAO1P1.VirtualizationMemoryCapacityAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationMemoryCapacityAO1P1.VirtualizationMemoryCapacityAO1P1)
		if err == nil {
			virtualizationmemorycapacity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationmemorycapacity["capacity"] = item.Capacity
	virtualizationmemorycapacity["class_id"] = item.ClassID
	virtualizationmemorycapacity["free"] = item.Free
	virtualizationmemorycapacity["object_type"] = item.ObjectType
	virtualizationmemorycapacity["used"] = item.Used

	virtualizationmemorycapacitys = append(virtualizationmemorycapacitys, virtualizationmemorycapacity)
	return virtualizationmemorycapacitys
}
func flattenMapVirtualizationProductInfo(p *models.VirtualizationProductInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationproductinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationproductinfo := make(map[string]interface{})
	delete(item.VirtualizationProductInfoAO1P1.VirtualizationProductInfoAO1P1, "ObjectType")
	if len(item.VirtualizationProductInfoAO1P1.VirtualizationProductInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationProductInfoAO1P1.VirtualizationProductInfoAO1P1)
		if err == nil {
			virtualizationproductinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationproductinfo["class_id"] = item.ClassID
	virtualizationproductinfo["object_type"] = item.ObjectType
	virtualizationproductinfo["product_name"] = item.ProductName
	virtualizationproductinfo["product_type"] = item.ProductType
	virtualizationproductinfo["product_vendor"] = item.ProductVendor
	virtualizationproductinfo["version"] = item.Version

	virtualizationproductinfos = append(virtualizationproductinfos, virtualizationproductinfo)
	return virtualizationproductinfos
}
func flattenMapVirtualizationStorageCapacity(p *models.VirtualizationStorageCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationstoragecapacitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationstoragecapacity := make(map[string]interface{})
	delete(item.VirtualizationStorageCapacityAO1P1.VirtualizationStorageCapacityAO1P1, "ObjectType")
	if len(item.VirtualizationStorageCapacityAO1P1.VirtualizationStorageCapacityAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationStorageCapacityAO1P1.VirtualizationStorageCapacityAO1P1)
		if err == nil {
			virtualizationstoragecapacity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationstoragecapacity["capacity"] = item.Capacity
	virtualizationstoragecapacity["class_id"] = item.ClassID
	virtualizationstoragecapacity["free"] = item.Free
	virtualizationstoragecapacity["object_type"] = item.ObjectType
	virtualizationstoragecapacity["used"] = item.Used

	virtualizationstoragecapacitys = append(virtualizationstoragecapacitys, virtualizationstoragecapacity)
	return virtualizationstoragecapacitys
}
func flattenMapVirtualizationVmwareClusterRef(p *models.VirtualizationVmwareClusterRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareclusterrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwareclusterref := make(map[string]interface{})
	virtualizationvmwareclusterref["moid"] = item.Moid
	virtualizationvmwareclusterref["object_type"] = item.ObjectType
	virtualizationvmwareclusterref["selector"] = item.Selector

	virtualizationvmwareclusterrefs = append(virtualizationvmwareclusterrefs, virtualizationvmwareclusterref)
	return virtualizationvmwareclusterrefs
}
func flattenMapVirtualizationVmwareDatacenterRef(p *models.VirtualizationVmwareDatacenterRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatacenterrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwaredatacenterref := make(map[string]interface{})
	virtualizationvmwaredatacenterref["moid"] = item.Moid
	virtualizationvmwaredatacenterref["object_type"] = item.ObjectType
	virtualizationvmwaredatacenterref["selector"] = item.Selector

	virtualizationvmwaredatacenterrefs = append(virtualizationvmwaredatacenterrefs, virtualizationvmwaredatacenterref)
	return virtualizationvmwaredatacenterrefs
}
func flattenMapVirtualizationVmwareHostRef(p *models.VirtualizationVmwareHostRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwarehostref := make(map[string]interface{})
	virtualizationvmwarehostref["moid"] = item.Moid
	virtualizationvmwarehostref["object_type"] = item.ObjectType
	virtualizationvmwarehostref["selector"] = item.Selector

	virtualizationvmwarehostrefs = append(virtualizationvmwarehostrefs, virtualizationvmwarehostref)
	return virtualizationvmwarehostrefs
}
func flattenMapVirtualizationVmwareRemoteDisplayInfo(p *models.VirtualizationVmwareRemoteDisplayInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareremotedisplayinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwareremotedisplayinfo := make(map[string]interface{})
	delete(item.VirtualizationVmwareRemoteDisplayInfoAO1P1.VirtualizationVmwareRemoteDisplayInfoAO1P1, "ObjectType")
	if len(item.VirtualizationVmwareRemoteDisplayInfoAO1P1.VirtualizationVmwareRemoteDisplayInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationVmwareRemoteDisplayInfoAO1P1.VirtualizationVmwareRemoteDisplayInfoAO1P1)
		if err == nil {
			virtualizationvmwareremotedisplayinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationvmwareremotedisplayinfo["class_id"] = item.ClassID
	virtualizationvmwareremotedisplayinfo["object_type"] = item.ObjectType
	virtualizationvmwareremotedisplayinfo["remote_display_password"] = item.RemoteDisplayPassword
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_key"] = item.RemoteDisplayVncKey
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_port"] = item.RemoteDisplayVncPort

	virtualizationvmwareremotedisplayinfos = append(virtualizationvmwareremotedisplayinfos, virtualizationvmwareremotedisplayinfo)
	return virtualizationvmwareremotedisplayinfos
}
func flattenMapVirtualizationVmwareResourceConsumption(p *models.VirtualizationVmwareResourceConsumption, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareresourceconsumptions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwareresourceconsumption := make(map[string]interface{})
	delete(item.VirtualizationVmwareResourceConsumptionAO1P1.VirtualizationVmwareResourceConsumptionAO1P1, "ObjectType")
	if len(item.VirtualizationVmwareResourceConsumptionAO1P1.VirtualizationVmwareResourceConsumptionAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationVmwareResourceConsumptionAO1P1.VirtualizationVmwareResourceConsumptionAO1P1)
		if err == nil {
			virtualizationvmwareresourceconsumption["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationvmwareresourceconsumption["cpu_consumed"] = item.CPUConsumed
	virtualizationvmwareresourceconsumption["class_id"] = item.ClassID
	virtualizationvmwareresourceconsumption["memory_consumed"] = item.MemoryConsumed
	virtualizationvmwareresourceconsumption["object_type"] = item.ObjectType

	virtualizationvmwareresourceconsumptions = append(virtualizationvmwareresourceconsumptions, virtualizationvmwareresourceconsumption)
	return virtualizationvmwareresourceconsumptions
}
func flattenMapVirtualizationVmwareVcenterRef(p *models.VirtualizationVmwareVcenterRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevcenterrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwarevcenterref := make(map[string]interface{})
	virtualizationvmwarevcenterref["moid"] = item.Moid
	virtualizationvmwarevcenterref["object_type"] = item.ObjectType
	virtualizationvmwarevcenterref["selector"] = item.Selector

	virtualizationvmwarevcenterrefs = append(virtualizationvmwarevcenterrefs, virtualizationvmwarevcenterref)
	return virtualizationvmwarevcenterrefs
}
func flattenMapVirtualizationVmwareVMCPUShareInfo(p *models.VirtualizationVmwareVMCPUShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmcpushareinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwarevmcpushareinfo := make(map[string]interface{})
	delete(item.VirtualizationVmwareVMCPUShareInfoAO1P1.VirtualizationVmwareVMCPUShareInfoAO1P1, "ObjectType")
	if len(item.VirtualizationVmwareVMCPUShareInfoAO1P1.VirtualizationVmwareVMCPUShareInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationVmwareVMCPUShareInfoAO1P1.VirtualizationVmwareVMCPUShareInfoAO1P1)
		if err == nil {
			virtualizationvmwarevmcpushareinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationvmwarevmcpushareinfo["cpu_limit"] = item.CPULimit
	virtualizationvmwarevmcpushareinfo["cpu_overhead_limit"] = item.CPUOverheadLimit
	virtualizationvmwarevmcpushareinfo["cpu_reservation"] = item.CPUReservation
	virtualizationvmwarevmcpushareinfo["cpu_shares"] = item.CPUShares
	virtualizationvmwarevmcpushareinfo["class_id"] = item.ClassID
	virtualizationvmwarevmcpushareinfo["object_type"] = item.ObjectType

	virtualizationvmwarevmcpushareinfos = append(virtualizationvmwarevmcpushareinfos, virtualizationvmwarevmcpushareinfo)
	return virtualizationvmwarevmcpushareinfos
}
func flattenMapVirtualizationVmwareVMCPUSocketInfo(p *models.VirtualizationVmwareVMCPUSocketInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmcpusocketinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwarevmcpusocketinfo := make(map[string]interface{})
	delete(item.VirtualizationVmwareVMCPUSocketInfoAO1P1.VirtualizationVmwareVMCPUSocketInfoAO1P1, "ObjectType")
	if len(item.VirtualizationVmwareVMCPUSocketInfoAO1P1.VirtualizationVmwareVMCPUSocketInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationVmwareVMCPUSocketInfoAO1P1.VirtualizationVmwareVMCPUSocketInfoAO1P1)
		if err == nil {
			virtualizationvmwarevmcpusocketinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationvmwarevmcpusocketinfo["class_id"] = item.ClassID
	virtualizationvmwarevmcpusocketinfo["cores_per_socket"] = item.CoresPerSocket
	virtualizationvmwarevmcpusocketinfo["num_cpus"] = item.NumCpus
	virtualizationvmwarevmcpusocketinfo["num_sockets"] = item.NumSockets
	virtualizationvmwarevmcpusocketinfo["object_type"] = item.ObjectType

	virtualizationvmwarevmcpusocketinfos = append(virtualizationvmwarevmcpusocketinfos, virtualizationvmwarevmcpusocketinfo)
	return virtualizationvmwarevmcpusocketinfos
}
func flattenMapVirtualizationVmwareVMDiskCommitInfo(p *models.VirtualizationVmwareVMDiskCommitInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmdiskcommitinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwarevmdiskcommitinfo := make(map[string]interface{})
	delete(item.VirtualizationVmwareVMDiskCommitInfoAO1P1.VirtualizationVmwareVMDiskCommitInfoAO1P1, "ObjectType")
	if len(item.VirtualizationVmwareVMDiskCommitInfoAO1P1.VirtualizationVmwareVMDiskCommitInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationVmwareVMDiskCommitInfoAO1P1.VirtualizationVmwareVMDiskCommitInfoAO1P1)
		if err == nil {
			virtualizationvmwarevmdiskcommitinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationvmwarevmdiskcommitinfo["class_id"] = item.ClassID
	virtualizationvmwarevmdiskcommitinfo["committed_disk"] = item.CommittedDisk
	virtualizationvmwarevmdiskcommitinfo["object_type"] = item.ObjectType
	virtualizationvmwarevmdiskcommitinfo["un_committed_disk"] = item.UnCommittedDisk
	virtualizationvmwarevmdiskcommitinfo["unshared_disk"] = item.UnsharedDisk

	virtualizationvmwarevmdiskcommitinfos = append(virtualizationvmwarevmdiskcommitinfos, virtualizationvmwarevmdiskcommitinfo)
	return virtualizationvmwarevmdiskcommitinfos
}
func flattenMapVirtualizationVmwareVMMemoryShareInfo(p *models.VirtualizationVmwareVMMemoryShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmmemoryshareinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmwarevmmemoryshareinfo := make(map[string]interface{})
	delete(item.VirtualizationVmwareVMMemoryShareInfoAO1P1.VirtualizationVmwareVMMemoryShareInfoAO1P1, "ObjectType")
	if len(item.VirtualizationVmwareVMMemoryShareInfoAO1P1.VirtualizationVmwareVMMemoryShareInfoAO1P1) != 0 {
		j, err := json.Marshal(item.VirtualizationVmwareVMMemoryShareInfoAO1P1.VirtualizationVmwareVMMemoryShareInfoAO1P1)
		if err == nil {
			virtualizationvmwarevmmemoryshareinfo["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	virtualizationvmwarevmmemoryshareinfo["class_id"] = item.ClassID
	virtualizationvmwarevmmemoryshareinfo["mem_limit"] = item.MemLimit
	virtualizationvmwarevmmemoryshareinfo["mem_overhead_limit"] = item.MemOverheadLimit
	virtualizationvmwarevmmemoryshareinfo["mem_reservation"] = item.MemReservation
	virtualizationvmwarevmmemoryshareinfo["mem_shares"] = item.MemShares
	virtualizationvmwarevmmemoryshareinfo["object_type"] = item.ObjectType

	virtualizationvmwarevmmemoryshareinfos = append(virtualizationvmwarevmmemoryshareinfos, virtualizationvmwarevmmemoryshareinfo)
	return virtualizationvmwarevmmemoryshareinfos
}
func flattenMapVnicArfsSettings(p *models.VnicArfsSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicarfssettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicarfssettings := make(map[string]interface{})
	delete(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1, "ObjectType")
	if len(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1)
		if err == nil {
			vnicarfssettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicarfssettings["class_id"] = item.ClassID
	vnicarfssettings["enabled"] = item.Enabled
	vnicarfssettings["object_type"] = item.ObjectType

	vnicarfssettingss = append(vnicarfssettingss, vnicarfssettings)
	return vnicarfssettingss
}
func flattenMapVnicCdn(p *models.VnicCdn, d *schema.ResourceData) []map[string]interface{} {
	var vniccdns []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vniccdn := make(map[string]interface{})
	delete(item.VnicCdnAO1P1.VnicCdnAO1P1, "ObjectType")
	if len(item.VnicCdnAO1P1.VnicCdnAO1P1) != 0 {
		j, err := json.Marshal(item.VnicCdnAO1P1.VnicCdnAO1P1)
		if err == nil {
			vniccdn["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vniccdn["class_id"] = item.ClassID
	vniccdn["object_type"] = item.ObjectType
	vniccdn["source"] = item.Source
	vniccdn["value"] = item.Value

	vniccdns = append(vniccdns, vniccdn)
	return vniccdns
}
func flattenMapVnicCompletionQueueSettings(p *models.VnicCompletionQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vniccompletionqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vniccompletionqueuesettings := make(map[string]interface{})
	delete(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1)
		if err == nil {
			vniccompletionqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vniccompletionqueuesettings["class_id"] = item.ClassID
	vniccompletionqueuesettings["count"] = item.Count
	vniccompletionqueuesettings["object_type"] = item.ObjectType
	vniccompletionqueuesettings["ring_size"] = item.RingSize

	vniccompletionqueuesettingss = append(vniccompletionqueuesettingss, vniccompletionqueuesettings)
	return vniccompletionqueuesettingss
}
func flattenMapVnicEthAdapterPolicyRef(p *models.VnicEthAdapterPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicethadapterpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethadapterpolicyref := make(map[string]interface{})
	vnicethadapterpolicyref["moid"] = item.Moid
	vnicethadapterpolicyref["object_type"] = item.ObjectType
	vnicethadapterpolicyref["selector"] = item.Selector

	vnicethadapterpolicyrefs = append(vnicethadapterpolicyrefs, vnicethadapterpolicyref)
	return vnicethadapterpolicyrefs
}
func flattenMapVnicEthInterruptSettings(p *models.VnicEthInterruptSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethinterruptsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethinterruptsettings := make(map[string]interface{})
	delete(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1, "ObjectType")
	if len(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1)
		if err == nil {
			vnicethinterruptsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicethinterruptsettings["class_id"] = item.ClassID
	vnicethinterruptsettings["coalescing_time"] = item.CoalescingTime
	vnicethinterruptsettings["coalescing_type"] = item.CoalescingType
	vnicethinterruptsettings["count"] = item.Count
	vnicethinterruptsettings["mode"] = item.Mode
	vnicethinterruptsettings["object_type"] = item.ObjectType

	vnicethinterruptsettingss = append(vnicethinterruptsettingss, vnicethinterruptsettings)
	return vnicethinterruptsettingss
}
func flattenMapVnicEthNetworkPolicyRef(p *models.VnicEthNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicethnetworkpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethnetworkpolicyref := make(map[string]interface{})
	vnicethnetworkpolicyref["moid"] = item.Moid
	vnicethnetworkpolicyref["object_type"] = item.ObjectType
	vnicethnetworkpolicyref["selector"] = item.Selector

	vnicethnetworkpolicyrefs = append(vnicethnetworkpolicyrefs, vnicethnetworkpolicyref)
	return vnicethnetworkpolicyrefs
}
func flattenMapVnicEthQosPolicyRef(p *models.VnicEthQosPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicethqospolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethqospolicyref := make(map[string]interface{})
	vnicethqospolicyref["moid"] = item.Moid
	vnicethqospolicyref["object_type"] = item.ObjectType
	vnicethqospolicyref["selector"] = item.Selector

	vnicethqospolicyrefs = append(vnicethqospolicyrefs, vnicethqospolicyref)
	return vnicethqospolicyrefs
}
func flattenMapVnicEthRxQueueSettings(p *models.VnicEthRxQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethrxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethrxqueuesettings := make(map[string]interface{})
	delete(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1)
		if err == nil {
			vnicethrxqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicethrxqueuesettings["class_id"] = item.ClassID
	vnicethrxqueuesettings["count"] = item.Count
	vnicethrxqueuesettings["object_type"] = item.ObjectType
	vnicethrxqueuesettings["ring_size"] = item.RingSize

	vnicethrxqueuesettingss = append(vnicethrxqueuesettingss, vnicethrxqueuesettings)
	return vnicethrxqueuesettingss
}
func flattenMapVnicEthTxQueueSettings(p *models.VnicEthTxQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethtxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethtxqueuesettings := make(map[string]interface{})
	delete(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1)
		if err == nil {
			vnicethtxqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicethtxqueuesettings["class_id"] = item.ClassID
	vnicethtxqueuesettings["count"] = item.Count
	vnicethtxqueuesettings["object_type"] = item.ObjectType
	vnicethtxqueuesettings["ring_size"] = item.RingSize

	vnicethtxqueuesettingss = append(vnicethtxqueuesettingss, vnicethtxqueuesettings)
	return vnicethtxqueuesettingss
}
func flattenMapVnicFcAdapterPolicyRef(p *models.VnicFcAdapterPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcadapterpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcadapterpolicyref := make(map[string]interface{})
	vnicfcadapterpolicyref["moid"] = item.Moid
	vnicfcadapterpolicyref["object_type"] = item.ObjectType
	vnicfcadapterpolicyref["selector"] = item.Selector

	vnicfcadapterpolicyrefs = append(vnicfcadapterpolicyrefs, vnicfcadapterpolicyref)
	return vnicfcadapterpolicyrefs
}
func flattenMapVnicFcErrorRecoverySettings(p *models.VnicFcErrorRecoverySettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcerrorrecoverysettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcerrorrecoverysettings := make(map[string]interface{})
	delete(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1, "ObjectType")
	if len(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1)
		if err == nil {
			vnicfcerrorrecoverysettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicfcerrorrecoverysettings["class_id"] = item.ClassID
	vnicfcerrorrecoverysettings["enabled"] = item.Enabled
	vnicfcerrorrecoverysettings["io_retry_count"] = item.IoRetryCount
	vnicfcerrorrecoverysettings["io_retry_timeout"] = item.IoRetryTimeout
	vnicfcerrorrecoverysettings["link_down_timeout"] = item.LinkDownTimeout
	vnicfcerrorrecoverysettings["object_type"] = item.ObjectType
	vnicfcerrorrecoverysettings["port_down_timeout"] = item.PortDownTimeout

	vnicfcerrorrecoverysettingss = append(vnicfcerrorrecoverysettingss, vnicfcerrorrecoverysettings)
	return vnicfcerrorrecoverysettingss
}
func flattenMapVnicFcInterruptSettings(p *models.VnicFcInterruptSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcinterruptsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcinterruptsettings := make(map[string]interface{})
	delete(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1, "ObjectType")
	if len(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1)
		if err == nil {
			vnicfcinterruptsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicfcinterruptsettings["class_id"] = item.ClassID
	vnicfcinterruptsettings["mode"] = item.Mode
	vnicfcinterruptsettings["object_type"] = item.ObjectType

	vnicfcinterruptsettingss = append(vnicfcinterruptsettingss, vnicfcinterruptsettings)
	return vnicfcinterruptsettingss
}
func flattenMapVnicFcNetworkPolicyRef(p *models.VnicFcNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcnetworkpolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcnetworkpolicyref := make(map[string]interface{})
	vnicfcnetworkpolicyref["moid"] = item.Moid
	vnicfcnetworkpolicyref["object_type"] = item.ObjectType
	vnicfcnetworkpolicyref["selector"] = item.Selector

	vnicfcnetworkpolicyrefs = append(vnicfcnetworkpolicyrefs, vnicfcnetworkpolicyref)
	return vnicfcnetworkpolicyrefs
}
func flattenMapVnicFcQosPolicyRef(p *models.VnicFcQosPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqospolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcqospolicyref := make(map[string]interface{})
	vnicfcqospolicyref["moid"] = item.Moid
	vnicfcqospolicyref["object_type"] = item.ObjectType
	vnicfcqospolicyref["selector"] = item.Selector

	vnicfcqospolicyrefs = append(vnicfcqospolicyrefs, vnicfcqospolicyref)
	return vnicfcqospolicyrefs
}
func flattenMapVnicFcQueueSettings(p *models.VnicFcQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcqueuesettings := make(map[string]interface{})
	delete(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1)
		if err == nil {
			vnicfcqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicfcqueuesettings["class_id"] = item.ClassID
	vnicfcqueuesettings["count"] = item.Count
	vnicfcqueuesettings["object_type"] = item.ObjectType
	vnicfcqueuesettings["ring_size"] = item.RingSize

	vnicfcqueuesettingss = append(vnicfcqueuesettingss, vnicfcqueuesettings)
	return vnicfcqueuesettingss
}
func flattenMapVnicFlogiSettings(p *models.VnicFlogiSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicflogisettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicflogisettings := make(map[string]interface{})
	delete(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1, "ObjectType")
	if len(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1)
		if err == nil {
			vnicflogisettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicflogisettings["class_id"] = item.ClassID
	vnicflogisettings["object_type"] = item.ObjectType
	vnicflogisettings["retries"] = item.Retries
	vnicflogisettings["timeout"] = item.Timeout

	vnicflogisettingss = append(vnicflogisettingss, vnicflogisettings)
	return vnicflogisettingss
}
func flattenMapVnicLanConnectivityPolicyRef(p *models.VnicLanConnectivityPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vniclanconnectivitypolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vniclanconnectivitypolicyref := make(map[string]interface{})
	vniclanconnectivitypolicyref["moid"] = item.Moid
	vniclanconnectivitypolicyref["object_type"] = item.ObjectType
	vniclanconnectivitypolicyref["selector"] = item.Selector

	vniclanconnectivitypolicyrefs = append(vniclanconnectivitypolicyrefs, vniclanconnectivitypolicyref)
	return vniclanconnectivitypolicyrefs
}
func flattenMapVnicNvgreSettings(p *models.VnicNvgreSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicnvgresettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicnvgresettings := make(map[string]interface{})
	delete(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1, "ObjectType")
	if len(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1)
		if err == nil {
			vnicnvgresettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicnvgresettings["class_id"] = item.ClassID
	vnicnvgresettings["enabled"] = item.Enabled
	vnicnvgresettings["object_type"] = item.ObjectType

	vnicnvgresettingss = append(vnicnvgresettingss, vnicnvgresettings)
	return vnicnvgresettingss
}
func flattenMapVnicPlacementSettings(p *models.VnicPlacementSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicplacementsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicplacementsettings := make(map[string]interface{})
	delete(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1, "ObjectType")
	if len(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1)
		if err == nil {
			vnicplacementsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicplacementsettings["class_id"] = item.ClassID
	vnicplacementsettings["id"] = item.ID
	vnicplacementsettings["object_type"] = item.ObjectType
	vnicplacementsettings["pci_link"] = item.PciLink
	vnicplacementsettings["uplink"] = item.Uplink

	vnicplacementsettingss = append(vnicplacementsettingss, vnicplacementsettings)
	return vnicplacementsettingss
}
func flattenMapVnicPlogiSettings(p *models.VnicPlogiSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicplogisettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicplogisettings := make(map[string]interface{})
	delete(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1, "ObjectType")
	if len(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1)
		if err == nil {
			vnicplogisettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicplogisettings["class_id"] = item.ClassID
	vnicplogisettings["object_type"] = item.ObjectType
	vnicplogisettings["retries"] = item.Retries
	vnicplogisettings["timeout"] = item.Timeout

	vnicplogisettingss = append(vnicplogisettingss, vnicplogisettings)
	return vnicplogisettingss
}
func flattenMapVnicRoceSettings(p *models.VnicRoceSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicrocesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicrocesettings := make(map[string]interface{})
	delete(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1, "ObjectType")
	if len(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1)
		if err == nil {
			vnicrocesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicrocesettings["class_id"] = item.ClassID
	vnicrocesettings["enabled"] = item.Enabled
	vnicrocesettings["memory_regions"] = item.MemoryRegions
	vnicrocesettings["object_type"] = item.ObjectType
	vnicrocesettings["queue_pairs"] = item.QueuePairs
	vnicrocesettings["resource_groups"] = item.ResourceGroups

	vnicrocesettingss = append(vnicrocesettingss, vnicrocesettings)
	return vnicrocesettingss
}
func flattenMapVnicSanConnectivityPolicyRef(p *models.VnicSanConnectivityPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var vnicsanconnectivitypolicyrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicsanconnectivitypolicyref := make(map[string]interface{})
	vnicsanconnectivitypolicyref["moid"] = item.Moid
	vnicsanconnectivitypolicyref["object_type"] = item.ObjectType
	vnicsanconnectivitypolicyref["selector"] = item.Selector

	vnicsanconnectivitypolicyrefs = append(vnicsanconnectivitypolicyrefs, vnicsanconnectivitypolicyref)
	return vnicsanconnectivitypolicyrefs
}
func flattenMapVnicScsiQueueSettings(p *models.VnicScsiQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicscsiqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicscsiqueuesettings := make(map[string]interface{})
	delete(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1)
		if err == nil {
			vnicscsiqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicscsiqueuesettings["class_id"] = item.ClassID
	vnicscsiqueuesettings["count"] = item.Count
	vnicscsiqueuesettings["object_type"] = item.ObjectType
	vnicscsiqueuesettings["ring_size"] = item.RingSize

	vnicscsiqueuesettingss = append(vnicscsiqueuesettingss, vnicscsiqueuesettings)
	return vnicscsiqueuesettingss
}
func flattenMapVnicTCPOffloadSettings(p *models.VnicTCPOffloadSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnictcpoffloadsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnictcpoffloadsettings := make(map[string]interface{})
	delete(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1, "ObjectType")
	if len(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1)
		if err == nil {
			vnictcpoffloadsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnictcpoffloadsettings["class_id"] = item.ClassID
	vnictcpoffloadsettings["large_receive"] = item.LargeReceive
	vnictcpoffloadsettings["large_send"] = item.LargeSend
	vnictcpoffloadsettings["object_type"] = item.ObjectType
	vnictcpoffloadsettings["rx_checksum"] = item.RxChecksum
	vnictcpoffloadsettings["tx_checksum"] = item.TxChecksum

	vnictcpoffloadsettingss = append(vnictcpoffloadsettingss, vnictcpoffloadsettings)
	return vnictcpoffloadsettingss
}
func flattenMapVnicUsnicSettings(p *models.VnicUsnicSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicusnicsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicusnicsettings := make(map[string]interface{})
	delete(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1, "ObjectType")
	if len(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1)
		if err == nil {
			vnicusnicsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicusnicsettings["class_id"] = item.ClassID
	vnicusnicsettings["cos"] = item.Cos
	vnicusnicsettings["count"] = item.Count
	vnicusnicsettings["object_type"] = item.ObjectType
	vnicusnicsettings["usnic_adapter_policy"] = item.UsnicAdapterPolicy

	vnicusnicsettingss = append(vnicusnicsettingss, vnicusnicsettings)
	return vnicusnicsettingss
}
func flattenMapVnicVlanSettings(p *models.VnicVlanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvlansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicvlansettings := make(map[string]interface{})
	delete(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1, "ObjectType")
	if len(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1)
		if err == nil {
			vnicvlansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicvlansettings["class_id"] = item.ClassID
	vnicvlansettings["default_vlan"] = item.DefaultVlan
	vnicvlansettings["mode"] = item.Mode
	vnicvlansettings["object_type"] = item.ObjectType

	vnicvlansettingss = append(vnicvlansettingss, vnicvlansettings)
	return vnicvlansettingss
}
func flattenMapVnicVmqSettings(p *models.VnicVmqSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvmqsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicvmqsettings := make(map[string]interface{})
	delete(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1, "ObjectType")
	if len(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1)
		if err == nil {
			vnicvmqsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicvmqsettings["class_id"] = item.ClassID
	vnicvmqsettings["enabled"] = item.Enabled
	vnicvmqsettings["object_type"] = item.ObjectType

	vnicvmqsettingss = append(vnicvmqsettingss, vnicvmqsettings)
	return vnicvmqsettingss
}
func flattenMapVnicVsanSettings(p *models.VnicVsanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvsansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicvsansettings := make(map[string]interface{})
	delete(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1, "ObjectType")
	if len(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1)
		if err == nil {
			vnicvsansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicvsansettings["class_id"] = item.ClassID
	vnicvsansettings["id"] = item.ID
	vnicvsansettings["object_type"] = item.ObjectType

	vnicvsansettingss = append(vnicvsansettingss, vnicvsansettings)
	return vnicvsansettingss
}
func flattenMapVnicVxlanSettings(p *models.VnicVxlanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvxlansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicvxlansettings := make(map[string]interface{})
	delete(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1, "ObjectType")
	if len(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1)
		if err == nil {
			vnicvxlansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vnicvxlansettings["class_id"] = item.ClassID
	vnicvxlansettings["enabled"] = item.Enabled
	vnicvxlansettings["object_type"] = item.ObjectType

	vnicvxlansettingss = append(vnicvxlansettingss, vnicvxlansettings)
	return vnicvxlansettingss
}
func flattenMapWorkflowCatalogRef(p *models.WorkflowCatalogRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowcatalogrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowcatalogref := make(map[string]interface{})
	workflowcatalogref["moid"] = item.Moid
	workflowcatalogref["object_type"] = item.ObjectType
	workflowcatalogref["selector"] = item.Selector

	workflowcatalogrefs = append(workflowcatalogrefs, workflowcatalogref)
	return workflowcatalogrefs
}
func flattenMapWorkflowInternalProperties(p *models.WorkflowInternalProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowinternalpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowinternalproperties := make(map[string]interface{})
	delete(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1)
		if err == nil {
			workflowinternalproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	workflowinternalproperties["base_task_type"] = item.BaseTaskType
	workflowinternalproperties["class_id"] = item.ClassID
	workflowinternalproperties["constraints"] = (func(p *models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
		var workflowtaskconstraintss []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		workflowtaskconstraints := make(map[string]interface{})
		delete(item.WorkflowTaskConstraintsAO1P1.WorkflowTaskConstraintsAO1P1, "ObjectType")
		if len(item.WorkflowTaskConstraintsAO1P1.WorkflowTaskConstraintsAO1P1) != 0 {
			j, err := json.Marshal(item.WorkflowTaskConstraintsAO1P1.WorkflowTaskConstraintsAO1P1)
			if err == nil {
				workflowtaskconstraints["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		workflowtaskconstraints["class_id"] = item.ClassID
		workflowtaskconstraints["object_type"] = item.ObjectType
		workflowtaskconstraints["target_data_type"] = item.TargetDataType

		workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
		return workflowtaskconstraintss
	})(item.Constraints, d)
	workflowinternalproperties["internal"] = item.Internal
	workflowinternalproperties["object_type"] = item.ObjectType
	workflowinternalproperties["owner"] = item.Owner

	workflowinternalpropertiess = append(workflowinternalpropertiess, workflowinternalproperties)
	return workflowinternalpropertiess
}
func flattenMapWorkflowPendingDynamicWorkflowInfoRef(p *models.WorkflowPendingDynamicWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowpendingdynamicworkflowinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowpendingdynamicworkflowinforef := make(map[string]interface{})
	workflowpendingdynamicworkflowinforef["moid"] = item.Moid
	workflowpendingdynamicworkflowinforef["object_type"] = item.ObjectType
	workflowpendingdynamicworkflowinforef["selector"] = item.Selector

	workflowpendingdynamicworkflowinforefs = append(workflowpendingdynamicworkflowinforefs, workflowpendingdynamicworkflowinforef)
	return workflowpendingdynamicworkflowinforefs
}
func flattenMapWorkflowProperties(p *models.WorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowproperties := make(map[string]interface{})
	delete(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1)
		if err == nil {
			workflowproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	workflowproperties["class_id"] = item.ClassID
	workflowproperties["external_meta"] = item.ExternalMeta
	workflowproperties["input_definition"] = (func(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			workflowbasedatatype := make(map[string]interface{})
			if len(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1) != 0 {
				j, err := json.Marshal(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1)
				if err == nil {
					workflowbasedatatype["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			workflowbasedatatype["class_id"] = item.ClassID
			workflowbasedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				workflowdefaultvalue := make(map[string]interface{})
				delete(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1, "ObjectType")
				if len(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1) != 0 {
					j, err := json.Marshal(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1)
					if err == nil {
						workflowdefaultvalue["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				workflowdefaultvalue["class_id"] = item.ClassID
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.Default, d)
			workflowbasedatatype["description"] = item.Description
			workflowbasedatatype["label"] = item.Label
			workflowbasedatatype["name"] = item.Name
			workflowbasedatatype["object_type"] = item.ObjectType
			workflowbasedatatype["required"] = item.Required
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.InputDefinition, d)
	workflowproperties["object_type"] = item.ObjectType
	workflowproperties["output_definition"] = (func(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			workflowbasedatatype := make(map[string]interface{})
			if len(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1) != 0 {
				j, err := json.Marshal(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1)
				if err == nil {
					workflowbasedatatype["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			workflowbasedatatype["class_id"] = item.ClassID
			workflowbasedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				workflowdefaultvalue := make(map[string]interface{})
				delete(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1, "ObjectType")
				if len(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1) != 0 {
					j, err := json.Marshal(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1)
					if err == nil {
						workflowdefaultvalue["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				workflowdefaultvalue["class_id"] = item.ClassID
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.Default, d)
			workflowbasedatatype["description"] = item.Description
			workflowbasedatatype["label"] = item.Label
			workflowbasedatatype["name"] = item.Name
			workflowbasedatatype["object_type"] = item.ObjectType
			workflowbasedatatype["required"] = item.Required
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.OutputDefinition, d)
	workflowproperties["retry_count"] = item.RetryCount
	workflowproperties["retry_delay"] = item.RetryDelay
	workflowproperties["retry_policy"] = item.RetryPolicy
	workflowproperties["support_status"] = item.SupportStatus
	workflowproperties["timeout"] = item.Timeout
	workflowproperties["timeout_policy"] = item.TimeoutPolicy

	workflowpropertiess = append(workflowpropertiess, workflowproperties)
	return workflowpropertiess
}
func flattenMapWorkflowTaskConstraints(p *models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskconstraintss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowtaskconstraints := make(map[string]interface{})
	delete(item.WorkflowTaskConstraintsAO1P1.WorkflowTaskConstraintsAO1P1, "ObjectType")
	if len(item.WorkflowTaskConstraintsAO1P1.WorkflowTaskConstraintsAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowTaskConstraintsAO1P1.WorkflowTaskConstraintsAO1P1)
		if err == nil {
			workflowtaskconstraints["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	workflowtaskconstraints["class_id"] = item.ClassID
	workflowtaskconstraints["object_type"] = item.ObjectType
	workflowtaskconstraints["target_data_type"] = item.TargetDataType

	workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
	return workflowtaskconstraintss
}
func flattenMapWorkflowTaskDefinitionRef(p *models.WorkflowTaskDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowtaskdefinitionref := make(map[string]interface{})
	workflowtaskdefinitionref["moid"] = item.Moid
	workflowtaskdefinitionref["object_type"] = item.ObjectType
	workflowtaskdefinitionref["selector"] = item.Selector

	workflowtaskdefinitionrefs = append(workflowtaskdefinitionrefs, workflowtaskdefinitionref)
	return workflowtaskdefinitionrefs
}
func flattenMapWorkflowTaskInfoRef(p *models.WorkflowTaskInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowtaskinforef := make(map[string]interface{})
	workflowtaskinforef["moid"] = item.Moid
	workflowtaskinforef["object_type"] = item.ObjectType
	workflowtaskinforef["selector"] = item.Selector

	workflowtaskinforefs = append(workflowtaskinforefs, workflowtaskinforef)
	return workflowtaskinforefs
}
func flattenMapWorkflowValidationInformation(p *models.WorkflowValidationInformation, d *schema.ResourceData) []map[string]interface{} {
	var workflowvalidationinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowvalidationinformation := make(map[string]interface{})
	delete(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1, "ObjectType")
	if len(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1)
		if err == nil {
			workflowvalidationinformation["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	workflowvalidationinformation["class_id"] = item.ClassID
	workflowvalidationinformation["object_type"] = item.ObjectType
	workflowvalidationinformation["state"] = item.State
	workflowvalidationinformation["validation_error"] = (func(p []*models.WorkflowValidationError, d *schema.ResourceData) []map[string]interface{} {
		var workflowvalidationerrors []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			workflowvalidationerror := make(map[string]interface{})
			if len(item.WorkflowValidationErrorAO1P1.WorkflowValidationErrorAO1P1) != 0 {
				j, err := json.Marshal(item.WorkflowValidationErrorAO1P1.WorkflowValidationErrorAO1P1)
				if err == nil {
					workflowvalidationerror["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			workflowvalidationerror["class_id"] = item.ClassID
			workflowvalidationerror["error_log"] = item.ErrorLog
			workflowvalidationerror["field"] = item.Field
			workflowvalidationerror["object_type"] = item.ObjectType
			workflowvalidationerror["task_name"] = item.TaskName
			workflowvalidationerror["transition_name"] = item.TransitionName
			workflowvalidationerrors = append(workflowvalidationerrors, workflowvalidationerror)
		}
		return workflowvalidationerrors
	})(item.ValidationError, d)

	workflowvalidationinformations = append(workflowvalidationinformations, workflowvalidationinformation)
	return workflowvalidationinformations
}
func flattenMapWorkflowWorkflowDefinitionRef(p *models.WorkflowWorkflowDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowdefinitionrefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowworkflowdefinitionref := make(map[string]interface{})
	workflowworkflowdefinitionref["moid"] = item.Moid
	workflowworkflowdefinitionref["object_type"] = item.ObjectType
	workflowworkflowdefinitionref["selector"] = item.Selector

	workflowworkflowdefinitionrefs = append(workflowworkflowdefinitionrefs, workflowworkflowdefinitionref)
	return workflowworkflowdefinitionrefs
}
func flattenMapWorkflowWorkflowInfoProperties(p *models.WorkflowWorkflowInfoProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinfopropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowworkflowinfoproperties := make(map[string]interface{})
	delete(item.WorkflowWorkflowInfoPropertiesAO1P1.WorkflowWorkflowInfoPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowWorkflowInfoPropertiesAO1P1.WorkflowWorkflowInfoPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowWorkflowInfoPropertiesAO1P1.WorkflowWorkflowInfoPropertiesAO1P1)
		if err == nil {
			workflowworkflowinfoproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	workflowworkflowinfoproperties["class_id"] = item.ClassID
	workflowworkflowinfoproperties["object_type"] = item.ObjectType
	workflowworkflowinfoproperties["retryable"] = item.Retryable

	workflowworkflowinfopropertiess = append(workflowworkflowinfopropertiess, workflowworkflowinfoproperties)
	return workflowworkflowinfopropertiess
}
func flattenMapWorkflowWorkflowInfoRef(p *models.WorkflowWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowworkflowinforef := make(map[string]interface{})
	workflowworkflowinforef["moid"] = item.Moid
	workflowworkflowinforef["object_type"] = item.ObjectType
	workflowworkflowinforef["selector"] = item.Selector

	workflowworkflowinforefs = append(workflowworkflowinforefs, workflowworkflowinforef)
	return workflowworkflowinforefs
}
func flattenMapWorkflowWorkflowProperties(p *models.WorkflowWorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowworkflowproperties := make(map[string]interface{})
	delete(item.WorkflowWorkflowPropertiesAO1P1.WorkflowWorkflowPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowWorkflowPropertiesAO1P1.WorkflowWorkflowPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowWorkflowPropertiesAO1P1.WorkflowWorkflowPropertiesAO1P1)
		if err == nil {
			workflowworkflowproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	workflowworkflowproperties["class_id"] = item.ClassID
	workflowworkflowproperties["external_meta"] = item.ExternalMeta
	workflowworkflowproperties["object_type"] = item.ObjectType
	workflowworkflowproperties["retryable"] = item.Retryable
	workflowworkflowproperties["support_status"] = item.SupportStatus

	workflowworkflowpropertiess = append(workflowworkflowpropertiess, workflowworkflowproperties)
	return workflowworkflowpropertiess
}
func flattenMapX509Certificate(p *models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var x509certificates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	x509certificate := make(map[string]interface{})
	delete(item.X509CertificateAO1P1.X509CertificateAO1P1, "ObjectType")
	if len(item.X509CertificateAO1P1.X509CertificateAO1P1) != 0 {
		j, err := json.Marshal(item.X509CertificateAO1P1.X509CertificateAO1P1)
		if err == nil {
			x509certificate["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	x509certificate["class_id"] = item.ClassID
	x509certificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		pkixdistinguishedname := make(map[string]interface{})
		delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
		if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
			j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
			if err == nil {
				pkixdistinguishedname["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		pkixdistinguishedname["class_id"] = item.ClassID
		pkixdistinguishedname["common_name"] = item.CommonName
		pkixdistinguishedname["country"] = item.Country
		pkixdistinguishedname["locality"] = item.Locality
		pkixdistinguishedname["object_type"] = item.ObjectType
		pkixdistinguishedname["organization"] = item.Organization
		pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
		pkixdistinguishedname["state"] = item.State

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.Issuer, d)
	x509certificate["object_type"] = item.ObjectType
	x509certificate["pem_certificate"] = item.PemCertificate
	x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
	x509certificate["signature_algorithm"] = item.SignatureAlgorithm
	x509certificate["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		pkixdistinguishedname := make(map[string]interface{})
		delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
		if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
			j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
			if err == nil {
				pkixdistinguishedname["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		pkixdistinguishedname["class_id"] = item.ClassID
		pkixdistinguishedname["common_name"] = item.CommonName
		pkixdistinguishedname["country"] = item.Country
		pkixdistinguishedname["locality"] = item.Locality
		pkixdistinguishedname["object_type"] = item.ObjectType
		pkixdistinguishedname["organization"] = item.Organization
		pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
		pkixdistinguishedname["state"] = item.State

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.Subject, d)

	x509certificates = append(x509certificates, x509certificate)
	return x509certificates
}
