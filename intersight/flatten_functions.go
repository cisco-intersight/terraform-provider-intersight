package intersight

import (
	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenListAdapterAdapterConfig(p *[]models.AdapterAdapterConfig, d *schema.ResourceData) []map[string]interface{} {
	var adapteradapterconfigs []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		adapteradapterconfig := make(map[string]interface{})
		adapteradapterconfig["class_id"] = item.ClassId
		adapteradapterconfig["dce_interface_settings"] = (func(p *[]models.AdapterDceInterfaceSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterdceinterfacesettingss []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				adapterdceinterfacesettings := make(map[string]interface{})
				adapterdceinterfacesettings["class_id"] = item.ClassId
				adapterdceinterfacesettings["fec_mode"] = item.FecMode
				adapterdceinterfacesettings["interface_id"] = item.InterfaceId
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
			adapterethsettings["class_id"] = item.ClassId
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
			adapterfcsettings["class_id"] = item.ClassId
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
			adapterportchannelsettings["class_id"] = item.ClassId
			adapterportchannelsettings["enabled"] = item.Enabled
			adapterportchannelsettings["object_type"] = item.ObjectType

			adapterportchannelsettingss = append(adapterportchannelsettingss, adapterportchannelsettings)
			return adapterportchannelsettingss
		})(item.PortChannelSettings, d)
		adapteradapterconfig["slot_id"] = item.SlotId
		adapteradapterconfigs = append(adapteradapterconfigs, adapteradapterconfig)
	}
	return adapteradapterconfigs
}
func flattenListAdapterExtEthInterfaceRelationship(p *[]models.AdapterExtEthInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterextethinterfacerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.AdapterExtEthInterfaceRelationshipInterface.(*models.MoMoRef)
		adapterextethinterfacerelationship := make(map[string]interface{})
		adapterextethinterfacerelationship["class_id"] = item.ClassId
		adapterextethinterfacerelationship["link"] = item.Link
		adapterextethinterfacerelationship["moid"] = item.Moid
		adapterextethinterfacerelationship["object_type"] = item.ObjectType
		adapterextethinterfacerelationship["selector"] = item.Selector
		adapterextethinterfacerelationships = append(adapterextethinterfacerelationships, adapterextethinterfacerelationship)
	}
	return adapterextethinterfacerelationships
}
func flattenListAdapterHostEthInterfaceRelationship(p *[]models.AdapterHostEthInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostethinterfacerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.AdapterHostEthInterfaceRelationshipInterface.(*models.MoMoRef)
		adapterhostethinterfacerelationship := make(map[string]interface{})
		adapterhostethinterfacerelationship["class_id"] = item.ClassId
		adapterhostethinterfacerelationship["link"] = item.Link
		adapterhostethinterfacerelationship["moid"] = item.Moid
		adapterhostethinterfacerelationship["object_type"] = item.ObjectType
		adapterhostethinterfacerelationship["selector"] = item.Selector
		adapterhostethinterfacerelationships = append(adapterhostethinterfacerelationships, adapterhostethinterfacerelationship)
	}
	return adapterhostethinterfacerelationships
}
func flattenListAdapterHostFcInterfaceRelationship(p *[]models.AdapterHostFcInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostfcinterfacerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.AdapterHostFcInterfaceRelationshipInterface.(*models.MoMoRef)
		adapterhostfcinterfacerelationship := make(map[string]interface{})
		adapterhostfcinterfacerelationship["class_id"] = item.ClassId
		adapterhostfcinterfacerelationship["link"] = item.Link
		adapterhostfcinterfacerelationship["moid"] = item.Moid
		adapterhostfcinterfacerelationship["object_type"] = item.ObjectType
		adapterhostfcinterfacerelationship["selector"] = item.Selector
		adapterhostfcinterfacerelationships = append(adapterhostfcinterfacerelationships, adapterhostfcinterfacerelationship)
	}
	return adapterhostfcinterfacerelationships
}
func flattenListAdapterHostIscsiInterfaceRelationship(p *[]models.AdapterHostIscsiInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostiscsiinterfacerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.AdapterHostIscsiInterfaceRelationshipInterface.(*models.MoMoRef)
		adapterhostiscsiinterfacerelationship := make(map[string]interface{})
		adapterhostiscsiinterfacerelationship["class_id"] = item.ClassId
		adapterhostiscsiinterfacerelationship["link"] = item.Link
		adapterhostiscsiinterfacerelationship["moid"] = item.Moid
		adapterhostiscsiinterfacerelationship["object_type"] = item.ObjectType
		adapterhostiscsiinterfacerelationship["selector"] = item.Selector
		adapterhostiscsiinterfacerelationships = append(adapterhostiscsiinterfacerelationships, adapterhostiscsiinterfacerelationship)
	}
	return adapterhostiscsiinterfacerelationships
}
func flattenListAdapterUnitRelationship(p *[]models.AdapterUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.AdapterUnitRelationshipInterface.(*models.MoMoRef)
		adapterunitrelationship := make(map[string]interface{})
		adapterunitrelationship["class_id"] = item.ClassId
		adapterunitrelationship["link"] = item.Link
		adapterunitrelationship["moid"] = item.Moid
		adapterunitrelationship["object_type"] = item.ObjectType
		adapterunitrelationship["selector"] = item.Selector
		adapterunitrelationships = append(adapterunitrelationships, adapterunitrelationship)
	}
	return adapterunitrelationships
}
func flattenListApplianceDataExportPolicyRelationship(p *[]models.ApplianceDataExportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ApplianceDataExportPolicyRelationshipInterface.(*models.MoMoRef)
		appliancedataexportpolicyrelationship := make(map[string]interface{})
		appliancedataexportpolicyrelationship["class_id"] = item.ClassId
		appliancedataexportpolicyrelationship["link"] = item.Link
		appliancedataexportpolicyrelationship["moid"] = item.Moid
		appliancedataexportpolicyrelationship["object_type"] = item.ObjectType
		appliancedataexportpolicyrelationship["selector"] = item.Selector
		appliancedataexportpolicyrelationships = append(appliancedataexportpolicyrelationships, appliancedataexportpolicyrelationship)
	}
	return appliancedataexportpolicyrelationships
}
func flattenListApplianceKeyValuePair(p *[]models.ApplianceKeyValuePair, d *schema.ResourceData) []map[string]interface{} {
	var appliancekeyvaluepairs []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		appliancekeyvaluepair := make(map[string]interface{})
		appliancekeyvaluepair["class_id"] = item.ClassId
		appliancekeyvaluepair["key"] = item.Key
		appliancekeyvaluepair["object_type"] = item.ObjectType
		appliancekeyvaluepair["value"] = item.Value
		appliancekeyvaluepairs = append(appliancekeyvaluepairs, appliancekeyvaluepair)
	}
	return appliancekeyvaluepairs
}
func flattenListAssetClusterMemberRelationship(p *[]models.AssetClusterMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.AssetClusterMemberRelationshipInterface.(*models.MoMoRef)
		assetclustermemberrelationship := make(map[string]interface{})
		assetclustermemberrelationship["class_id"] = item.ClassId
		assetclustermemberrelationship["link"] = item.Link
		assetclustermemberrelationship["moid"] = item.Moid
		assetclustermemberrelationship["object_type"] = item.ObjectType
		assetclustermemberrelationship["selector"] = item.Selector
		assetclustermemberrelationships = append(assetclustermemberrelationships, assetclustermemberrelationship)
	}
	return assetclustermemberrelationships
}
func flattenListBiosUnitRelationship(p *[]models.BiosUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.BiosUnitRelationshipInterface.(*models.MoMoRef)
		biosunitrelationship := make(map[string]interface{})
		biosunitrelationship["class_id"] = item.ClassId
		biosunitrelationship["link"] = item.Link
		biosunitrelationship["moid"] = item.Moid
		biosunitrelationship["object_type"] = item.ObjectType
		biosunitrelationship["selector"] = item.Selector
		biosunitrelationships = append(biosunitrelationships, biosunitrelationship)
	}
	return biosunitrelationships
}
func flattenListBootDeviceBase(p *[]models.BootDeviceBase, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebases []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		bootdevicebase := make(map[string]interface{})
		bootdevicebase["class_id"] = item.ClassId
		bootdevicebase["enabled"] = item.Enabled
		bootdevicebase["name"] = item.Name
		bootdevicebase["object_type"] = item.ObjectType
		bootdevicebases = append(bootdevicebases, bootdevicebase)
	}
	return bootdevicebases
}
func flattenListComputeBladeRelationship(p *[]models.ComputeBladeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computebladerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ComputeBladeRelationshipInterface.(*models.MoMoRef)
		computebladerelationship := make(map[string]interface{})
		computebladerelationship["class_id"] = item.ClassId
		computebladerelationship["link"] = item.Link
		computebladerelationship["moid"] = item.Moid
		computebladerelationship["object_type"] = item.ObjectType
		computebladerelationship["selector"] = item.Selector
		computebladerelationships = append(computebladerelationships, computebladerelationship)
	}
	return computebladerelationships
}
func flattenListComputeIpAddress(p *[]models.ComputeIpAddress, d *schema.ResourceData) []map[string]interface{} {
	var computeipaddresss []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		computeipaddress := make(map[string]interface{})
		computeipaddress["address"] = item.Address
		computeipaddress["category"] = item.Category
		computeipaddress["class_id"] = item.ClassId
		computeipaddress["default_gateway"] = item.DefaultGateway
		computeipaddress["dn"] = item.Dn
		computeipaddress["http_port"] = item.HttpPort
		computeipaddress["https_port"] = item.HttpsPort
		computeipaddress["kvm_port"] = item.KvmPort
		computeipaddress["name"] = item.Name
		computeipaddress["object_type"] = item.ObjectType
		computeipaddress["subnet"] = item.Subnet
		computeipaddress["type"] = item.Type
		computeipaddresss = append(computeipaddresss, computeipaddress)
	}
	return computeipaddresss
}
func flattenListComputeRackUnitRelationship(p *[]models.ComputeRackUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ComputeRackUnitRelationshipInterface.(*models.MoMoRef)
		computerackunitrelationship := make(map[string]interface{})
		computerackunitrelationship["class_id"] = item.ClassId
		computerackunitrelationship["link"] = item.Link
		computerackunitrelationship["moid"] = item.Moid
		computerackunitrelationship["object_type"] = item.ObjectType
		computerackunitrelationship["selector"] = item.Selector
		computerackunitrelationships = append(computerackunitrelationships, computerackunitrelationship)
	}
	return computerackunitrelationships
}
func flattenListCondHclStatusDetailRelationship(p *[]models.CondHclStatusDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusdetailrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.CondHclStatusDetailRelationshipInterface.(*models.MoMoRef)
		condhclstatusdetailrelationship := make(map[string]interface{})
		condhclstatusdetailrelationship["class_id"] = item.ClassId
		condhclstatusdetailrelationship["link"] = item.Link
		condhclstatusdetailrelationship["moid"] = item.Moid
		condhclstatusdetailrelationship["object_type"] = item.ObjectType
		condhclstatusdetailrelationship["selector"] = item.Selector
		condhclstatusdetailrelationships = append(condhclstatusdetailrelationships, condhclstatusdetailrelationship)
	}
	return condhclstatusdetailrelationships
}
func flattenListEquipmentFanRelationship(p *[]models.EquipmentFanRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentFanRelationshipInterface.(*models.MoMoRef)
		equipmentfanrelationship := make(map[string]interface{})
		equipmentfanrelationship["class_id"] = item.ClassId
		equipmentfanrelationship["link"] = item.Link
		equipmentfanrelationship["moid"] = item.Moid
		equipmentfanrelationship["object_type"] = item.ObjectType
		equipmentfanrelationship["selector"] = item.Selector
		equipmentfanrelationships = append(equipmentfanrelationships, equipmentfanrelationship)
	}
	return equipmentfanrelationships
}
func flattenListEquipmentFanModuleRelationship(p *[]models.EquipmentFanModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentFanModuleRelationshipInterface.(*models.MoMoRef)
		equipmentfanmodulerelationship := make(map[string]interface{})
		equipmentfanmodulerelationship["class_id"] = item.ClassId
		equipmentfanmodulerelationship["link"] = item.Link
		equipmentfanmodulerelationship["moid"] = item.Moid
		equipmentfanmodulerelationship["object_type"] = item.ObjectType
		equipmentfanmodulerelationship["selector"] = item.Selector
		equipmentfanmodulerelationships = append(equipmentfanmodulerelationships, equipmentfanmodulerelationship)
	}
	return equipmentfanmodulerelationships
}
func flattenListEquipmentIoCardRelationship(p *[]models.EquipmentIoCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentIoCardRelationshipInterface.(*models.MoMoRef)
		equipmentiocardrelationship := make(map[string]interface{})
		equipmentiocardrelationship["class_id"] = item.ClassId
		equipmentiocardrelationship["link"] = item.Link
		equipmentiocardrelationship["moid"] = item.Moid
		equipmentiocardrelationship["object_type"] = item.ObjectType
		equipmentiocardrelationship["selector"] = item.Selector
		equipmentiocardrelationships = append(equipmentiocardrelationships, equipmentiocardrelationship)
	}
	return equipmentiocardrelationships
}
func flattenListEquipmentIoExpanderRelationship(p *[]models.EquipmentIoExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentioexpanderrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentIoExpanderRelationshipInterface.(*models.MoMoRef)
		equipmentioexpanderrelationship := make(map[string]interface{})
		equipmentioexpanderrelationship["class_id"] = item.ClassId
		equipmentioexpanderrelationship["link"] = item.Link
		equipmentioexpanderrelationship["moid"] = item.Moid
		equipmentioexpanderrelationship["object_type"] = item.ObjectType
		equipmentioexpanderrelationship["selector"] = item.Selector
		equipmentioexpanderrelationships = append(equipmentioexpanderrelationships, equipmentioexpanderrelationship)
	}
	return equipmentioexpanderrelationships
}
func flattenListEquipmentPsuRelationship(p *[]models.EquipmentPsuRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentpsurelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentPsuRelationshipInterface.(*models.MoMoRef)
		equipmentpsurelationship := make(map[string]interface{})
		equipmentpsurelationship["class_id"] = item.ClassId
		equipmentpsurelationship["link"] = item.Link
		equipmentpsurelationship["moid"] = item.Moid
		equipmentpsurelationship["object_type"] = item.ObjectType
		equipmentpsurelationship["selector"] = item.Selector
		equipmentpsurelationships = append(equipmentpsurelationships, equipmentpsurelationship)
	}
	return equipmentpsurelationships
}
func flattenListEquipmentRackEnclosureSlotRelationship(p *[]models.EquipmentRackEnclosureSlotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentRackEnclosureSlotRelationshipInterface.(*models.MoMoRef)
		equipmentrackenclosureslotrelationship := make(map[string]interface{})
		equipmentrackenclosureslotrelationship["class_id"] = item.ClassId
		equipmentrackenclosureslotrelationship["link"] = item.Link
		equipmentrackenclosureslotrelationship["moid"] = item.Moid
		equipmentrackenclosureslotrelationship["object_type"] = item.ObjectType
		equipmentrackenclosureslotrelationship["selector"] = item.Selector
		equipmentrackenclosureslotrelationships = append(equipmentrackenclosureslotrelationships, equipmentrackenclosureslotrelationship)
	}
	return equipmentrackenclosureslotrelationships
}
func flattenListEquipmentSwitchCardRelationship(p *[]models.EquipmentSwitchCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentSwitchCardRelationshipInterface.(*models.MoMoRef)
		equipmentswitchcardrelationship := make(map[string]interface{})
		equipmentswitchcardrelationship["class_id"] = item.ClassId
		equipmentswitchcardrelationship["link"] = item.Link
		equipmentswitchcardrelationship["moid"] = item.Moid
		equipmentswitchcardrelationship["object_type"] = item.ObjectType
		equipmentswitchcardrelationship["selector"] = item.Selector
		equipmentswitchcardrelationships = append(equipmentswitchcardrelationships, equipmentswitchcardrelationship)
	}
	return equipmentswitchcardrelationships
}
func flattenListEquipmentSystemIoControllerRelationship(p *[]models.EquipmentSystemIoControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentSystemIoControllerRelationshipInterface.(*models.MoMoRef)
		equipmentsystemiocontrollerrelationship := make(map[string]interface{})
		equipmentsystemiocontrollerrelationship["class_id"] = item.ClassId
		equipmentsystemiocontrollerrelationship["link"] = item.Link
		equipmentsystemiocontrollerrelationship["moid"] = item.Moid
		equipmentsystemiocontrollerrelationship["object_type"] = item.ObjectType
		equipmentsystemiocontrollerrelationship["selector"] = item.Selector
		equipmentsystemiocontrollerrelationships = append(equipmentsystemiocontrollerrelationships, equipmentsystemiocontrollerrelationship)
	}
	return equipmentsystemiocontrollerrelationships
}
func flattenListEquipmentTpmRelationship(p *[]models.EquipmentTpmRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmenttpmrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EquipmentTpmRelationshipInterface.(*models.MoMoRef)
		equipmenttpmrelationship := make(map[string]interface{})
		equipmenttpmrelationship["class_id"] = item.ClassId
		equipmenttpmrelationship["link"] = item.Link
		equipmenttpmrelationship["moid"] = item.Moid
		equipmenttpmrelationship["object_type"] = item.ObjectType
		equipmenttpmrelationship["selector"] = item.Selector
		equipmenttpmrelationships = append(equipmenttpmrelationships, equipmenttpmrelationship)
	}
	return equipmenttpmrelationships
}
func flattenListEtherPhysicalPortRelationship(p *[]models.EtherPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.EtherPhysicalPortRelationshipInterface.(*models.MoMoRef)
		etherphysicalportrelationship := make(map[string]interface{})
		etherphysicalportrelationship["class_id"] = item.ClassId
		etherphysicalportrelationship["link"] = item.Link
		etherphysicalportrelationship["moid"] = item.Moid
		etherphysicalportrelationship["object_type"] = item.ObjectType
		etherphysicalportrelationship["selector"] = item.Selector
		etherphysicalportrelationships = append(etherphysicalportrelationships, etherphysicalportrelationship)
	}
	return etherphysicalportrelationships
}
func flattenListFcPhysicalPortRelationship(p *[]models.FcPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcphysicalportrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.FcPhysicalPortRelationshipInterface.(*models.MoMoRef)
		fcphysicalportrelationship := make(map[string]interface{})
		fcphysicalportrelationship["class_id"] = item.ClassId
		fcphysicalportrelationship["link"] = item.Link
		fcphysicalportrelationship["moid"] = item.Moid
		fcphysicalportrelationship["object_type"] = item.ObjectType
		fcphysicalportrelationship["selector"] = item.Selector
		fcphysicalportrelationships = append(fcphysicalportrelationships, fcphysicalportrelationship)
	}
	return fcphysicalportrelationships
}
func flattenListFirmwareRunningFirmwareRelationship(p *[]models.FirmwareRunningFirmwareRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.FirmwareRunningFirmwareRelationshipInterface.(*models.MoMoRef)
		firmwarerunningfirmwarerelationship := make(map[string]interface{})
		firmwarerunningfirmwarerelationship["class_id"] = item.ClassId
		firmwarerunningfirmwarerelationship["link"] = item.Link
		firmwarerunningfirmwarerelationship["moid"] = item.Moid
		firmwarerunningfirmwarerelationship["object_type"] = item.ObjectType
		firmwarerunningfirmwarerelationship["selector"] = item.Selector
		firmwarerunningfirmwarerelationships = append(firmwarerunningfirmwarerelationships, firmwarerunningfirmwarerelationship)
	}
	return firmwarerunningfirmwarerelationships
}
func flattenListForecastDefinitionRelationship(p *[]models.ForecastDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ForecastDefinitionRelationshipInterface.(*models.MoMoRef)
		forecastdefinitionrelationship := make(map[string]interface{})
		forecastdefinitionrelationship["class_id"] = item.ClassId
		forecastdefinitionrelationship["link"] = item.Link
		forecastdefinitionrelationship["moid"] = item.Moid
		forecastdefinitionrelationship["object_type"] = item.ObjectType
		forecastdefinitionrelationship["selector"] = item.Selector
		forecastdefinitionrelationships = append(forecastdefinitionrelationships, forecastdefinitionrelationship)
	}
	return forecastdefinitionrelationships
}
func flattenListGraphicsCardRelationship(p *[]models.GraphicsCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.GraphicsCardRelationshipInterface.(*models.MoMoRef)
		graphicscardrelationship := make(map[string]interface{})
		graphicscardrelationship["class_id"] = item.ClassId
		graphicscardrelationship["link"] = item.Link
		graphicscardrelationship["moid"] = item.Moid
		graphicscardrelationship["object_type"] = item.ObjectType
		graphicscardrelationship["selector"] = item.Selector
		graphicscardrelationships = append(graphicscardrelationships, graphicscardrelationship)
	}
	return graphicscardrelationships
}
func flattenListGraphicsControllerRelationship(p *[]models.GraphicsControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscontrollerrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.GraphicsControllerRelationshipInterface.(*models.MoMoRef)
		graphicscontrollerrelationship := make(map[string]interface{})
		graphicscontrollerrelationship["class_id"] = item.ClassId
		graphicscontrollerrelationship["link"] = item.Link
		graphicscontrollerrelationship["moid"] = item.Moid
		graphicscontrollerrelationship["object_type"] = item.ObjectType
		graphicscontrollerrelationship["selector"] = item.Selector
		graphicscontrollerrelationships = append(graphicscontrollerrelationships, graphicscontrollerrelationship)
	}
	return graphicscontrollerrelationships
}
func flattenListHclConstraint(p *[]models.HclConstraint, d *schema.ResourceData) []map[string]interface{} {
	var hclconstraints []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		hclconstraint := make(map[string]interface{})
		hclconstraint["class_id"] = item.ClassId
		hclconstraint["constraint_name"] = item.ConstraintName
		hclconstraint["constraint_value"] = item.ConstraintValue
		hclconstraint["object_type"] = item.ObjectType
		hclconstraints = append(hclconstraints, hclconstraint)
	}
	return hclconstraints
}
func flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(p *[]models.HclHyperflexSoftwareCompatibilityInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hclhyperflexsoftwarecompatibilityinforelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HclHyperflexSoftwareCompatibilityInfoRelationshipInterface.(*models.MoMoRef)
		hclhyperflexsoftwarecompatibilityinforelationship := make(map[string]interface{})
		hclhyperflexsoftwarecompatibilityinforelationship["class_id"] = item.ClassId
		hclhyperflexsoftwarecompatibilityinforelationship["link"] = item.Link
		hclhyperflexsoftwarecompatibilityinforelationship["moid"] = item.Moid
		hclhyperflexsoftwarecompatibilityinforelationship["object_type"] = item.ObjectType
		hclhyperflexsoftwarecompatibilityinforelationship["selector"] = item.Selector
		hclhyperflexsoftwarecompatibilityinforelationships = append(hclhyperflexsoftwarecompatibilityinforelationships, hclhyperflexsoftwarecompatibilityinforelationship)
	}
	return hclhyperflexsoftwarecompatibilityinforelationships
}
func flattenListHclOperatingSystemRelationship(p *[]models.HclOperatingSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HclOperatingSystemRelationshipInterface.(*models.MoMoRef)
		hcloperatingsystemrelationship := make(map[string]interface{})
		hcloperatingsystemrelationship["class_id"] = item.ClassId
		hcloperatingsystemrelationship["link"] = item.Link
		hcloperatingsystemrelationship["moid"] = item.Moid
		hcloperatingsystemrelationship["object_type"] = item.ObjectType
		hcloperatingsystemrelationship["selector"] = item.Selector
		hcloperatingsystemrelationships = append(hcloperatingsystemrelationships, hcloperatingsystemrelationship)
	}
	return hcloperatingsystemrelationships
}
func flattenListHyperflexAlarmRelationship(p *[]models.HyperflexAlarmRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexalarmrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HyperflexAlarmRelationshipInterface.(*models.MoMoRef)
		hyperflexalarmrelationship := make(map[string]interface{})
		hyperflexalarmrelationship["class_id"] = item.ClassId
		hyperflexalarmrelationship["link"] = item.Link
		hyperflexalarmrelationship["moid"] = item.Moid
		hyperflexalarmrelationship["object_type"] = item.ObjectType
		hyperflexalarmrelationship["selector"] = item.Selector
		hyperflexalarmrelationships = append(hyperflexalarmrelationships, hyperflexalarmrelationship)
	}
	return hyperflexalarmrelationships
}
func flattenListHyperflexCapabilityInfoRelationship(p *[]models.HyperflexCapabilityInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexcapabilityinforelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HyperflexCapabilityInfoRelationshipInterface.(*models.MoMoRef)
		hyperflexcapabilityinforelationship := make(map[string]interface{})
		hyperflexcapabilityinforelationship["class_id"] = item.ClassId
		hyperflexcapabilityinforelationship["link"] = item.Link
		hyperflexcapabilityinforelationship["moid"] = item.Moid
		hyperflexcapabilityinforelationship["object_type"] = item.ObjectType
		hyperflexcapabilityinforelationship["selector"] = item.Selector
		hyperflexcapabilityinforelationships = append(hyperflexcapabilityinforelationships, hyperflexcapabilityinforelationship)
	}
	return hyperflexcapabilityinforelationships
}
func flattenListHyperflexClusterProfileRelationship(p *[]models.HyperflexClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HyperflexClusterProfileRelationshipInterface.(*models.MoMoRef)
		hyperflexclusterprofilerelationship := make(map[string]interface{})
		hyperflexclusterprofilerelationship["class_id"] = item.ClassId
		hyperflexclusterprofilerelationship["link"] = item.Link
		hyperflexclusterprofilerelationship["moid"] = item.Moid
		hyperflexclusterprofilerelationship["object_type"] = item.ObjectType
		hyperflexclusterprofilerelationship["selector"] = item.Selector
		hyperflexclusterprofilerelationships = append(hyperflexclusterprofilerelationships, hyperflexclusterprofilerelationship)
	}
	return hyperflexclusterprofilerelationships
}
func flattenListHyperflexConfigResultEntryRelationship(p *[]models.HyperflexConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultentryrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HyperflexConfigResultEntryRelationshipInterface.(*models.MoMoRef)
		hyperflexconfigresultentryrelationship := make(map[string]interface{})
		hyperflexconfigresultentryrelationship["class_id"] = item.ClassId
		hyperflexconfigresultentryrelationship["link"] = item.Link
		hyperflexconfigresultentryrelationship["moid"] = item.Moid
		hyperflexconfigresultentryrelationship["object_type"] = item.ObjectType
		hyperflexconfigresultentryrelationship["selector"] = item.Selector
		hyperflexconfigresultentryrelationships = append(hyperflexconfigresultentryrelationships, hyperflexconfigresultentryrelationship)
	}
	return hyperflexconfigresultentryrelationships
}
func flattenListHyperflexFeatureLimitEntry(p *[]models.HyperflexFeatureLimitEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitentrys []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		hyperflexfeaturelimitentry := make(map[string]interface{})
		hyperflexfeaturelimitentry["class_id"] = item.ClassId
		hyperflexfeaturelimitentry["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["class_id"] = item.ClassId
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
func flattenListHyperflexHxZoneResiliencyInfoDt(p *[]models.HyperflexHxZoneResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxzoneresiliencyinfodts []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		hyperflexhxzoneresiliencyinfodt := make(map[string]interface{})
		hyperflexhxzoneresiliencyinfodt["class_id"] = item.ClassId
		hyperflexhxzoneresiliencyinfodt["name"] = item.Name
		hyperflexhxzoneresiliencyinfodt["object_type"] = item.ObjectType
		hyperflexhxzoneresiliencyinfodt["resiliency_info"] = (func(p *models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexhxresiliencyinfodts []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexhxresiliencyinfodt := make(map[string]interface{})
			hyperflexhxresiliencyinfodt["class_id"] = item.ClassId
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
func flattenListHyperflexHxdpVersionRelationship(p *[]models.HyperflexHxdpVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxdpversionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HyperflexHxdpVersionRelationshipInterface.(*models.MoMoRef)
		hyperflexhxdpversionrelationship := make(map[string]interface{})
		hyperflexhxdpversionrelationship["class_id"] = item.ClassId
		hyperflexhxdpversionrelationship["link"] = item.Link
		hyperflexhxdpversionrelationship["moid"] = item.Moid
		hyperflexhxdpversionrelationship["object_type"] = item.ObjectType
		hyperflexhxdpversionrelationship["selector"] = item.Selector
		hyperflexhxdpversionrelationships = append(hyperflexhxdpversionrelationships, hyperflexhxdpversionrelationship)
	}
	return hyperflexhxdpversionrelationships
}
func flattenListHyperflexNamedVlan(p *[]models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvlans []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		hyperflexnamedvlan := make(map[string]interface{})
		hyperflexnamedvlan["class_id"] = item.ClassId
		hyperflexnamedvlan["name"] = item.Name
		hyperflexnamedvlan["object_type"] = item.ObjectType
		hyperflexnamedvlan["vlan_id"] = item.VlanId
		hyperflexnamedvlans = append(hyperflexnamedvlans, hyperflexnamedvlan)
	}
	return hyperflexnamedvlans
}
func flattenListHyperflexNodeRelationship(p *[]models.HyperflexNodeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnoderelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HyperflexNodeRelationshipInterface.(*models.MoMoRef)
		hyperflexnoderelationship := make(map[string]interface{})
		hyperflexnoderelationship["class_id"] = item.ClassId
		hyperflexnoderelationship["link"] = item.Link
		hyperflexnoderelationship["moid"] = item.Moid
		hyperflexnoderelationship["object_type"] = item.ObjectType
		hyperflexnoderelationship["selector"] = item.Selector
		hyperflexnoderelationships = append(hyperflexnoderelationships, hyperflexnoderelationship)
	}
	return hyperflexnoderelationships
}
func flattenListHyperflexNodeProfileRelationship(p *[]models.HyperflexNodeProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeprofilerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.HyperflexNodeProfileRelationshipInterface.(*models.MoMoRef)
		hyperflexnodeprofilerelationship := make(map[string]interface{})
		hyperflexnodeprofilerelationship["class_id"] = item.ClassId
		hyperflexnodeprofilerelationship["link"] = item.Link
		hyperflexnodeprofilerelationship["moid"] = item.Moid
		hyperflexnodeprofilerelationship["object_type"] = item.ObjectType
		hyperflexnodeprofilerelationship["selector"] = item.Selector
		hyperflexnodeprofilerelationships = append(hyperflexnodeprofilerelationships, hyperflexnodeprofilerelationship)
	}
	return hyperflexnodeprofilerelationships
}
func flattenListHyperflexServerFirmwareVersionEntry(p *[]models.HyperflexServerFirmwareVersionEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionentrys []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		hyperflexserverfirmwareversionentry := make(map[string]interface{})
		hyperflexserverfirmwareversionentry["class_id"] = item.ClassId
		hyperflexserverfirmwareversionentry["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["class_id"] = item.ClassId
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
func flattenListHyperflexServerModelEntry(p *[]models.HyperflexServerModelEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelentrys []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		hyperflexservermodelentry := make(map[string]interface{})
		hyperflexservermodelentry["class_id"] = item.ClassId
		hyperflexservermodelentry["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["class_id"] = item.ClassId
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
func flattenListIaasConnectorPackRelationship(p *[]models.IaasConnectorPackRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasconnectorpackrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IaasConnectorPackRelationshipInterface.(*models.MoMoRef)
		iaasconnectorpackrelationship := make(map[string]interface{})
		iaasconnectorpackrelationship["class_id"] = item.ClassId
		iaasconnectorpackrelationship["link"] = item.Link
		iaasconnectorpackrelationship["moid"] = item.Moid
		iaasconnectorpackrelationship["object_type"] = item.ObjectType
		iaasconnectorpackrelationship["selector"] = item.Selector
		iaasconnectorpackrelationships = append(iaasconnectorpackrelationships, iaasconnectorpackrelationship)
	}
	return iaasconnectorpackrelationships
}
func flattenListIaasDeviceStatusRelationship(p *[]models.IaasDeviceStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasdevicestatusrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IaasDeviceStatusRelationshipInterface.(*models.MoMoRef)
		iaasdevicestatusrelationship := make(map[string]interface{})
		iaasdevicestatusrelationship["class_id"] = item.ClassId
		iaasdevicestatusrelationship["link"] = item.Link
		iaasdevicestatusrelationship["moid"] = item.Moid
		iaasdevicestatusrelationship["object_type"] = item.ObjectType
		iaasdevicestatusrelationship["selector"] = item.Selector
		iaasdevicestatusrelationships = append(iaasdevicestatusrelationships, iaasdevicestatusrelationship)
	}
	return iaasdevicestatusrelationships
}
func flattenListIaasLicenseKeysInfo(p *[]models.IaasLicenseKeysInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicensekeysinfos []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		iaaslicensekeysinfo := make(map[string]interface{})
		iaaslicensekeysinfo["class_id"] = item.ClassId
		iaaslicensekeysinfo["count"] = item.Count
		iaaslicensekeysinfo["expiration_date"] = item.ExpirationDate
		iaaslicensekeysinfo["license_id"] = item.LicenseId
		iaaslicensekeysinfo["object_type"] = item.ObjectType
		iaaslicensekeysinfo["pid"] = item.Pid
		iaaslicensekeysinfos = append(iaaslicensekeysinfos, iaaslicensekeysinfo)
	}
	return iaaslicensekeysinfos
}
func flattenListIaasLicenseUtilizationInfo(p *[]models.IaasLicenseUtilizationInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseutilizationinfos []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		iaaslicenseutilizationinfo := make(map[string]interface{})
		iaaslicenseutilizationinfo["actual_used"] = item.ActualUsed
		iaaslicenseutilizationinfo["class_id"] = item.ClassId
		iaaslicenseutilizationinfo["label"] = item.Label
		iaaslicenseutilizationinfo["licensed_limit"] = item.LicensedLimit
		iaaslicenseutilizationinfo["object_type"] = item.ObjectType
		iaaslicenseutilizationinfo["sku"] = item.Sku
		iaaslicenseutilizationinfos = append(iaaslicenseutilizationinfos, iaaslicenseutilizationinfo)
	}
	return iaaslicenseutilizationinfos
}
func flattenListIaasMostRunTasksRelationship(p *[]models.IaasMostRunTasksRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasmostruntasksrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IaasMostRunTasksRelationshipInterface.(*models.MoMoRef)
		iaasmostruntasksrelationship := make(map[string]interface{})
		iaasmostruntasksrelationship["class_id"] = item.ClassId
		iaasmostruntasksrelationship["link"] = item.Link
		iaasmostruntasksrelationship["moid"] = item.Moid
		iaasmostruntasksrelationship["object_type"] = item.ObjectType
		iaasmostruntasksrelationship["selector"] = item.Selector
		iaasmostruntasksrelationships = append(iaasmostruntasksrelationships, iaasmostruntasksrelationship)
	}
	return iaasmostruntasksrelationships
}
func flattenListIamAccountPermissions(p *[]models.IamAccountPermissions, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountpermissionss []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		iamaccountpermissions := make(map[string]interface{})
		iamaccountpermissions["account_identifier"] = item.AccountIdentifier
		iamaccountpermissions["account_name"] = item.AccountName
		iamaccountpermissions["account_status"] = item.AccountStatus
		iamaccountpermissions["class_id"] = item.ClassId
		iamaccountpermissions["object_type"] = item.ObjectType
		iamaccountpermissions["permissions"] = (func(p *[]models.IamPermissionReference, d *schema.ResourceData) []map[string]interface{} {
			var iampermissionreferences []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				iampermissionreference := make(map[string]interface{})
				iampermissionreference["class_id"] = item.ClassId
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
func flattenListIamApiKeyRelationship(p *[]models.IamApiKeyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamapikeyrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamApiKeyRelationshipInterface.(*models.MoMoRef)
		iamapikeyrelationship := make(map[string]interface{})
		iamapikeyrelationship["class_id"] = item.ClassId
		iamapikeyrelationship["link"] = item.Link
		iamapikeyrelationship["moid"] = item.Moid
		iamapikeyrelationship["object_type"] = item.ObjectType
		iamapikeyrelationship["selector"] = item.Selector
		iamapikeyrelationships = append(iamapikeyrelationships, iamapikeyrelationship)
	}
	return iamapikeyrelationships
}
func flattenListIamAppRegistrationRelationship(p *[]models.IamAppRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamAppRegistrationRelationshipInterface.(*models.MoMoRef)
		iamappregistrationrelationship := make(map[string]interface{})
		iamappregistrationrelationship["class_id"] = item.ClassId
		iamappregistrationrelationship["link"] = item.Link
		iamappregistrationrelationship["moid"] = item.Moid
		iamappregistrationrelationship["object_type"] = item.ObjectType
		iamappregistrationrelationship["selector"] = item.Selector
		iamappregistrationrelationships = append(iamappregistrationrelationships, iamappregistrationrelationship)
	}
	return iamappregistrationrelationships
}
func flattenListIamDomainGroupRelationship(p *[]models.IamDomainGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamDomainGroupRelationshipInterface.(*models.MoMoRef)
		iamdomaingrouprelationship := make(map[string]interface{})
		iamdomaingrouprelationship["class_id"] = item.ClassId
		iamdomaingrouprelationship["link"] = item.Link
		iamdomaingrouprelationship["moid"] = item.Moid
		iamdomaingrouprelationship["object_type"] = item.ObjectType
		iamdomaingrouprelationship["selector"] = item.Selector
		iamdomaingrouprelationships = append(iamdomaingrouprelationships, iamdomaingrouprelationship)
	}
	return iamdomaingrouprelationships
}
func flattenListIamEndPointPrivilegeRelationship(p *[]models.IamEndPointPrivilegeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointprivilegerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamEndPointPrivilegeRelationshipInterface.(*models.MoMoRef)
		iamendpointprivilegerelationship := make(map[string]interface{})
		iamendpointprivilegerelationship["class_id"] = item.ClassId
		iamendpointprivilegerelationship["link"] = item.Link
		iamendpointprivilegerelationship["moid"] = item.Moid
		iamendpointprivilegerelationship["object_type"] = item.ObjectType
		iamendpointprivilegerelationship["selector"] = item.Selector
		iamendpointprivilegerelationships = append(iamendpointprivilegerelationships, iamendpointprivilegerelationship)
	}
	return iamendpointprivilegerelationships
}
func flattenListIamEndPointRoleRelationship(p *[]models.IamEndPointRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointrolerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamEndPointRoleRelationshipInterface.(*models.MoMoRef)
		iamendpointrolerelationship := make(map[string]interface{})
		iamendpointrolerelationship["class_id"] = item.ClassId
		iamendpointrolerelationship["link"] = item.Link
		iamendpointrolerelationship["moid"] = item.Moid
		iamendpointrolerelationship["object_type"] = item.ObjectType
		iamendpointrolerelationship["selector"] = item.Selector
		iamendpointrolerelationships = append(iamendpointrolerelationships, iamendpointrolerelationship)
	}
	return iamendpointrolerelationships
}
func flattenListIamEndPointUserRoleRelationship(p *[]models.IamEndPointUserRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrolerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamEndPointUserRoleRelationshipInterface.(*models.MoMoRef)
		iamendpointuserrolerelationship := make(map[string]interface{})
		iamendpointuserrolerelationship["class_id"] = item.ClassId
		iamendpointuserrolerelationship["link"] = item.Link
		iamendpointuserrolerelationship["moid"] = item.Moid
		iamendpointuserrolerelationship["object_type"] = item.ObjectType
		iamendpointuserrolerelationship["selector"] = item.Selector
		iamendpointuserrolerelationships = append(iamendpointuserrolerelationships, iamendpointuserrolerelationship)
	}
	return iamendpointuserrolerelationships
}
func flattenListIamGroupPermissionToRoles(p *[]models.IamGroupPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iamgrouppermissiontoroless []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		iamgrouppermissiontoroles := make(map[string]interface{})
		iamgrouppermissiontoroles["class_id"] = item.ClassId
		iamgrouppermissiontoroles["group"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			cmrfcmrf := make(map[string]interface{})
			cmrfcmrf["class_id"] = item.ClassId
			cmrfcmrf["moid"] = item.Moid
			cmrfcmrf["object_type"] = item.ObjectType
			cmrfcmrf["object_type"] = item.ObjectType

			cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			return cmrfcmrfs
		})(item.Group, d)
		iamgrouppermissiontoroles["object_type"] = item.ObjectType
		iamgrouppermissiontoroles["orgs"] = (func(p *[]models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				cmrfcmrf := make(map[string]interface{})
				cmrfcmrf["class_id"] = item.ClassId
				cmrfcmrf["moid"] = item.Moid
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			}
			return cmrfcmrfs
		})(item.Orgs, d)
		iamgrouppermissiontoroless = append(iamgrouppermissiontoroless, iamgrouppermissiontoroles)
	}
	return iamgrouppermissiontoroless
}
func flattenListIamIdpRelationship(p *[]models.IamIdpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamIdpRelationshipInterface.(*models.MoMoRef)
		iamidprelationship := make(map[string]interface{})
		iamidprelationship["class_id"] = item.ClassId
		iamidprelationship["link"] = item.Link
		iamidprelationship["moid"] = item.Moid
		iamidprelationship["object_type"] = item.ObjectType
		iamidprelationship["selector"] = item.Selector
		iamidprelationships = append(iamidprelationships, iamidprelationship)
	}
	return iamidprelationships
}
func flattenListIamIdpReferenceRelationship(p *[]models.IamIdpReferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamIdpReferenceRelationshipInterface.(*models.MoMoRef)
		iamidpreferencerelationship := make(map[string]interface{})
		iamidpreferencerelationship["class_id"] = item.ClassId
		iamidpreferencerelationship["link"] = item.Link
		iamidpreferencerelationship["moid"] = item.Moid
		iamidpreferencerelationship["object_type"] = item.ObjectType
		iamidpreferencerelationship["selector"] = item.Selector
		iamidpreferencerelationships = append(iamidpreferencerelationships, iamidpreferencerelationship)
	}
	return iamidpreferencerelationships
}
func flattenListIamLdapGroupRelationship(p *[]models.IamLdapGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldapgrouprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamLdapGroupRelationshipInterface.(*models.MoMoRef)
		iamldapgrouprelationship := make(map[string]interface{})
		iamldapgrouprelationship["class_id"] = item.ClassId
		iamldapgrouprelationship["link"] = item.Link
		iamldapgrouprelationship["moid"] = item.Moid
		iamldapgrouprelationship["object_type"] = item.ObjectType
		iamldapgrouprelationship["selector"] = item.Selector
		iamldapgrouprelationships = append(iamldapgrouprelationships, iamldapgrouprelationship)
	}
	return iamldapgrouprelationships
}
func flattenListIamLdapProviderRelationship(p *[]models.IamLdapProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldapproviderrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamLdapProviderRelationshipInterface.(*models.MoMoRef)
		iamldapproviderrelationship := make(map[string]interface{})
		iamldapproviderrelationship["class_id"] = item.ClassId
		iamldapproviderrelationship["link"] = item.Link
		iamldapproviderrelationship["moid"] = item.Moid
		iamldapproviderrelationship["object_type"] = item.ObjectType
		iamldapproviderrelationship["selector"] = item.Selector
		iamldapproviderrelationships = append(iamldapproviderrelationships, iamldapproviderrelationship)
	}
	return iamldapproviderrelationships
}
func flattenListIamOAuthTokenRelationship(p *[]models.IamOAuthTokenRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamoauthtokenrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamOAuthTokenRelationshipInterface.(*models.MoMoRef)
		iamoauthtokenrelationship := make(map[string]interface{})
		iamoauthtokenrelationship["class_id"] = item.ClassId
		iamoauthtokenrelationship["link"] = item.Link
		iamoauthtokenrelationship["moid"] = item.Moid
		iamoauthtokenrelationship["object_type"] = item.ObjectType
		iamoauthtokenrelationship["selector"] = item.Selector
		iamoauthtokenrelationships = append(iamoauthtokenrelationships, iamoauthtokenrelationship)
	}
	return iamoauthtokenrelationships
}
func flattenListIamPermissionRelationship(p *[]models.IamPermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamPermissionRelationshipInterface.(*models.MoMoRef)
		iampermissionrelationship := make(map[string]interface{})
		iampermissionrelationship["class_id"] = item.ClassId
		iampermissionrelationship["link"] = item.Link
		iampermissionrelationship["moid"] = item.Moid
		iampermissionrelationship["object_type"] = item.ObjectType
		iampermissionrelationship["selector"] = item.Selector
		iampermissionrelationships = append(iampermissionrelationships, iampermissionrelationship)
	}
	return iampermissionrelationships
}
func flattenListIamPermissionToRoles(p *[]models.IamPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iampermissiontoroless []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		iampermissiontoroles := make(map[string]interface{})
		iampermissiontoroles["class_id"] = item.ClassId
		iampermissiontoroles["object_type"] = item.ObjectType
		iampermissiontoroles["permission"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			cmrfcmrf := make(map[string]interface{})
			cmrfcmrf["class_id"] = item.ClassId
			cmrfcmrf["moid"] = item.Moid
			cmrfcmrf["object_type"] = item.ObjectType
			cmrfcmrf["object_type"] = item.ObjectType

			cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			return cmrfcmrfs
		})(item.Permission, d)
		iampermissiontoroles["roles"] = (func(p *[]models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				cmrfcmrf := make(map[string]interface{})
				cmrfcmrf["class_id"] = item.ClassId
				cmrfcmrf["moid"] = item.Moid
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			}
			return cmrfcmrfs
		})(item.Roles, d)
		iampermissiontoroless = append(iampermissiontoroless, iampermissiontoroles)
	}
	return iampermissiontoroless
}
func flattenListIamPrivilegeRelationship(p *[]models.IamPrivilegeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamPrivilegeRelationshipInterface.(*models.MoMoRef)
		iamprivilegerelationship := make(map[string]interface{})
		iamprivilegerelationship["class_id"] = item.ClassId
		iamprivilegerelationship["link"] = item.Link
		iamprivilegerelationship["moid"] = item.Moid
		iamprivilegerelationship["object_type"] = item.ObjectType
		iamprivilegerelationship["selector"] = item.Selector
		iamprivilegerelationships = append(iamprivilegerelationships, iamprivilegerelationship)
	}
	return iamprivilegerelationships
}
func flattenListIamPrivilegeSetRelationship(p *[]models.IamPrivilegeSetRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegesetrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamPrivilegeSetRelationshipInterface.(*models.MoMoRef)
		iamprivilegesetrelationship := make(map[string]interface{})
		iamprivilegesetrelationship["class_id"] = item.ClassId
		iamprivilegesetrelationship["link"] = item.Link
		iamprivilegesetrelationship["moid"] = item.Moid
		iamprivilegesetrelationship["object_type"] = item.ObjectType
		iamprivilegesetrelationship["selector"] = item.Selector
		iamprivilegesetrelationships = append(iamprivilegesetrelationships, iamprivilegesetrelationship)
	}
	return iamprivilegesetrelationships
}
func flattenListIamResourcePermissionRelationship(p *[]models.IamResourcePermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcepermissionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamResourcePermissionRelationshipInterface.(*models.MoMoRef)
		iamresourcepermissionrelationship := make(map[string]interface{})
		iamresourcepermissionrelationship["class_id"] = item.ClassId
		iamresourcepermissionrelationship["link"] = item.Link
		iamresourcepermissionrelationship["moid"] = item.Moid
		iamresourcepermissionrelationship["object_type"] = item.ObjectType
		iamresourcepermissionrelationship["selector"] = item.Selector
		iamresourcepermissionrelationships = append(iamresourcepermissionrelationships, iamresourcepermissionrelationship)
	}
	return iamresourcepermissionrelationships
}
func flattenListIamResourceRolesRelationship(p *[]models.IamResourceRolesRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcerolesrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamResourceRolesRelationshipInterface.(*models.MoMoRef)
		iamresourcerolesrelationship := make(map[string]interface{})
		iamresourcerolesrelationship["class_id"] = item.ClassId
		iamresourcerolesrelationship["link"] = item.Link
		iamresourcerolesrelationship["moid"] = item.Moid
		iamresourcerolesrelationship["object_type"] = item.ObjectType
		iamresourcerolesrelationship["selector"] = item.Selector
		iamresourcerolesrelationships = append(iamresourcerolesrelationships, iamresourcerolesrelationship)
	}
	return iamresourcerolesrelationships
}
func flattenListIamRoleRelationship(p *[]models.IamRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamrolerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamRoleRelationshipInterface.(*models.MoMoRef)
		iamrolerelationship := make(map[string]interface{})
		iamrolerelationship["class_id"] = item.ClassId
		iamrolerelationship["link"] = item.Link
		iamrolerelationship["moid"] = item.Moid
		iamrolerelationship["object_type"] = item.ObjectType
		iamrolerelationship["selector"] = item.Selector
		iamrolerelationships = append(iamrolerelationships, iamrolerelationship)
	}
	return iamrolerelationships
}
func flattenListIamSessionRelationship(p *[]models.IamSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamSessionRelationshipInterface.(*models.MoMoRef)
		iamsessionrelationship := make(map[string]interface{})
		iamsessionrelationship["class_id"] = item.ClassId
		iamsessionrelationship["link"] = item.Link
		iamsessionrelationship["moid"] = item.Moid
		iamsessionrelationship["object_type"] = item.ObjectType
		iamsessionrelationship["selector"] = item.Selector
		iamsessionrelationships = append(iamsessionrelationships, iamsessionrelationship)
	}
	return iamsessionrelationships
}
func flattenListIamUserRelationship(p *[]models.IamUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamUserRelationshipInterface.(*models.MoMoRef)
		iamuserrelationship := make(map[string]interface{})
		iamuserrelationship["class_id"] = item.ClassId
		iamuserrelationship["link"] = item.Link
		iamuserrelationship["moid"] = item.Moid
		iamuserrelationship["object_type"] = item.ObjectType
		iamuserrelationship["selector"] = item.Selector
		iamuserrelationships = append(iamuserrelationships, iamuserrelationship)
	}
	return iamuserrelationships
}
func flattenListIamUserGroupRelationship(p *[]models.IamUserGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamUserGroupRelationshipInterface.(*models.MoMoRef)
		iamusergrouprelationship := make(map[string]interface{})
		iamusergrouprelationship["class_id"] = item.ClassId
		iamusergrouprelationship["link"] = item.Link
		iamusergrouprelationship["moid"] = item.Moid
		iamusergrouprelationship["object_type"] = item.ObjectType
		iamusergrouprelationship["selector"] = item.Selector
		iamusergrouprelationships = append(iamusergrouprelationships, iamusergrouprelationship)
	}
	return iamusergrouprelationships
}
func flattenListIamUserPreferenceRelationship(p *[]models.IamUserPreferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserpreferencerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.IamUserPreferenceRelationshipInterface.(*models.MoMoRef)
		iamuserpreferencerelationship := make(map[string]interface{})
		iamuserpreferencerelationship["class_id"] = item.ClassId
		iamuserpreferencerelationship["link"] = item.Link
		iamuserpreferencerelationship["moid"] = item.Moid
		iamuserpreferencerelationship["object_type"] = item.ObjectType
		iamuserpreferencerelationship["selector"] = item.Selector
		iamuserpreferencerelationships = append(iamuserpreferencerelationships, iamuserpreferencerelationship)
	}
	return iamuserpreferencerelationships
}
func flattenListInventoryGenericInventoryRelationship(p *[]models.InventoryGenericInventoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.InventoryGenericInventoryRelationshipInterface.(*models.MoMoRef)
		inventorygenericinventoryrelationship := make(map[string]interface{})
		inventorygenericinventoryrelationship["class_id"] = item.ClassId
		inventorygenericinventoryrelationship["link"] = item.Link
		inventorygenericinventoryrelationship["moid"] = item.Moid
		inventorygenericinventoryrelationship["object_type"] = item.ObjectType
		inventorygenericinventoryrelationship["selector"] = item.Selector
		inventorygenericinventoryrelationships = append(inventorygenericinventoryrelationships, inventorygenericinventoryrelationship)
	}
	return inventorygenericinventoryrelationships
}
func flattenListInventoryGenericInventoryHolderRelationship(p *[]models.InventoryGenericInventoryHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.InventoryGenericInventoryHolderRelationshipInterface.(*models.MoMoRef)
		inventorygenericinventoryholderrelationship := make(map[string]interface{})
		inventorygenericinventoryholderrelationship["class_id"] = item.ClassId
		inventorygenericinventoryholderrelationship["link"] = item.Link
		inventorygenericinventoryholderrelationship["moid"] = item.Moid
		inventorygenericinventoryholderrelationship["object_type"] = item.ObjectType
		inventorygenericinventoryholderrelationship["selector"] = item.Selector
		inventorygenericinventoryholderrelationships = append(inventorygenericinventoryholderrelationships, inventorygenericinventoryholderrelationship)
	}
	return inventorygenericinventoryholderrelationships
}
func flattenListLicenseLicenseInfoRelationship(p *[]models.LicenseLicenseInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenselicenseinforelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.LicenseLicenseInfoRelationshipInterface.(*models.MoMoRef)
		licenselicenseinforelationship := make(map[string]interface{})
		licenselicenseinforelationship["class_id"] = item.ClassId
		licenselicenseinforelationship["link"] = item.Link
		licenselicenseinforelationship["moid"] = item.Moid
		licenselicenseinforelationship["object_type"] = item.ObjectType
		licenselicenseinforelationship["selector"] = item.Selector
		licenselicenseinforelationships = append(licenselicenseinforelationships, licenselicenseinforelationship)
	}
	return licenselicenseinforelationships
}
func flattenListManagementInterfaceRelationship(p *[]models.ManagementInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managementinterfacerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ManagementInterfaceRelationshipInterface.(*models.MoMoRef)
		managementinterfacerelationship := make(map[string]interface{})
		managementinterfacerelationship["class_id"] = item.ClassId
		managementinterfacerelationship["link"] = item.Link
		managementinterfacerelationship["moid"] = item.Moid
		managementinterfacerelationship["object_type"] = item.ObjectType
		managementinterfacerelationship["selector"] = item.Selector
		managementinterfacerelationships = append(managementinterfacerelationships, managementinterfacerelationship)
	}
	return managementinterfacerelationships
}
func flattenListMemoryArrayRelationship(p *[]models.MemoryArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.MemoryArrayRelationshipInterface.(*models.MoMoRef)
		memoryarrayrelationship := make(map[string]interface{})
		memoryarrayrelationship["class_id"] = item.ClassId
		memoryarrayrelationship["link"] = item.Link
		memoryarrayrelationship["moid"] = item.Moid
		memoryarrayrelationship["object_type"] = item.ObjectType
		memoryarrayrelationship["selector"] = item.Selector
		memoryarrayrelationships = append(memoryarrayrelationships, memoryarrayrelationship)
	}
	return memoryarrayrelationships
}
func flattenListMemoryPersistentMemoryGoal(p *[]models.MemoryPersistentMemoryGoal, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorygoals []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		memorypersistentmemorygoal := make(map[string]interface{})
		memorypersistentmemorygoal["class_id"] = item.ClassId
		memorypersistentmemorygoal["memory_mode_percentage"] = item.MemoryModePercentage
		memorypersistentmemorygoal["object_type"] = item.ObjectType
		memorypersistentmemorygoal["persistent_memory_type"] = item.PersistentMemoryType
		memorypersistentmemorygoal["socket_id"] = item.SocketId
		memorypersistentmemorygoals = append(memorypersistentmemorygoals, memorypersistentmemorygoal)
	}
	return memorypersistentmemorygoals
}
func flattenListMemoryPersistentMemoryLogicalNamespace(p *[]models.MemoryPersistentMemoryLogicalNamespace, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylogicalnamespaces []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		memorypersistentmemorylogicalnamespace := make(map[string]interface{})
		memorypersistentmemorylogicalnamespace["capacity"] = item.Capacity
		memorypersistentmemorylogicalnamespace["class_id"] = item.ClassId
		memorypersistentmemorylogicalnamespace["mode"] = item.Mode
		memorypersistentmemorylogicalnamespace["name"] = item.Name
		memorypersistentmemorylogicalnamespace["object_type"] = item.ObjectType
		memorypersistentmemorylogicalnamespace["socket_id"] = item.SocketId
		memorypersistentmemorylogicalnamespace["socket_memory_id"] = item.SocketMemoryId
		memorypersistentmemorylogicalnamespaces = append(memorypersistentmemorylogicalnamespaces, memorypersistentmemorylogicalnamespace)
	}
	return memorypersistentmemorylogicalnamespaces
}
func flattenListMemoryPersistentMemoryNamespaceRelationship(p *[]models.MemoryPersistentMemoryNamespaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespacerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.MemoryPersistentMemoryNamespaceRelationshipInterface.(*models.MoMoRef)
		memorypersistentmemorynamespacerelationship := make(map[string]interface{})
		memorypersistentmemorynamespacerelationship["class_id"] = item.ClassId
		memorypersistentmemorynamespacerelationship["link"] = item.Link
		memorypersistentmemorynamespacerelationship["moid"] = item.Moid
		memorypersistentmemorynamespacerelationship["object_type"] = item.ObjectType
		memorypersistentmemorynamespacerelationship["selector"] = item.Selector
		memorypersistentmemorynamespacerelationships = append(memorypersistentmemorynamespacerelationships, memorypersistentmemorynamespacerelationship)
	}
	return memorypersistentmemorynamespacerelationships
}
func flattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(p *[]models.MemoryPersistentMemoryNamespaceConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespaceconfigresultrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.MemoryPersistentMemoryNamespaceConfigResultRelationshipInterface.(*models.MoMoRef)
		memorypersistentmemorynamespaceconfigresultrelationship := make(map[string]interface{})
		memorypersistentmemorynamespaceconfigresultrelationship["class_id"] = item.ClassId
		memorypersistentmemorynamespaceconfigresultrelationship["link"] = item.Link
		memorypersistentmemorynamespaceconfigresultrelationship["moid"] = item.Moid
		memorypersistentmemorynamespaceconfigresultrelationship["object_type"] = item.ObjectType
		memorypersistentmemorynamespaceconfigresultrelationship["selector"] = item.Selector
		memorypersistentmemorynamespaceconfigresultrelationships = append(memorypersistentmemorynamespaceconfigresultrelationships, memorypersistentmemorynamespaceconfigresultrelationship)
	}
	return memorypersistentmemorynamespaceconfigresultrelationships
}
func flattenListMemoryPersistentMemoryRegionRelationship(p *[]models.MemoryPersistentMemoryRegionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.MemoryPersistentMemoryRegionRelationshipInterface.(*models.MoMoRef)
		memorypersistentmemoryregionrelationship := make(map[string]interface{})
		memorypersistentmemoryregionrelationship["class_id"] = item.ClassId
		memorypersistentmemoryregionrelationship["link"] = item.Link
		memorypersistentmemoryregionrelationship["moid"] = item.Moid
		memorypersistentmemoryregionrelationship["object_type"] = item.ObjectType
		memorypersistentmemoryregionrelationship["selector"] = item.Selector
		memorypersistentmemoryregionrelationships = append(memorypersistentmemoryregionrelationships, memorypersistentmemoryregionrelationship)
	}
	return memorypersistentmemoryregionrelationships
}
func flattenListMemoryPersistentMemoryUnitRelationship(p *[]models.MemoryPersistentMemoryUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryunitrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.MemoryPersistentMemoryUnitRelationshipInterface.(*models.MoMoRef)
		memorypersistentmemoryunitrelationship := make(map[string]interface{})
		memorypersistentmemoryunitrelationship["class_id"] = item.ClassId
		memorypersistentmemoryunitrelationship["link"] = item.Link
		memorypersistentmemoryunitrelationship["moid"] = item.Moid
		memorypersistentmemoryunitrelationship["object_type"] = item.ObjectType
		memorypersistentmemoryunitrelationship["selector"] = item.Selector
		memorypersistentmemoryunitrelationships = append(memorypersistentmemoryunitrelationships, memorypersistentmemoryunitrelationship)
	}
	return memorypersistentmemoryunitrelationships
}
func flattenListMemoryUnitRelationship(p *[]models.MemoryUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryunitrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.MemoryUnitRelationshipInterface.(*models.MoMoRef)
		memoryunitrelationship := make(map[string]interface{})
		memoryunitrelationship["class_id"] = item.ClassId
		memoryunitrelationship["link"] = item.Link
		memoryunitrelationship["moid"] = item.Moid
		memoryunitrelationship["object_type"] = item.ObjectType
		memoryunitrelationship["selector"] = item.Selector
		memoryunitrelationships = append(memoryunitrelationships, memoryunitrelationship)
	}
	return memoryunitrelationships
}
func flattenListMetaAccessPrivilege(p *[]models.MetaAccessPrivilege, d *schema.ResourceData) []map[string]interface{} {
	var metaaccessprivileges []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		metaaccessprivilege := make(map[string]interface{})
		metaaccessprivilege["class_id"] = item.ClassId
		metaaccessprivilege["method"] = item.Method
		metaaccessprivilege["object_type"] = item.ObjectType
		metaaccessprivilege["privilege"] = item.Privilege
		metaaccessprivileges = append(metaaccessprivileges, metaaccessprivilege)
	}
	return metaaccessprivileges
}
func flattenListMetaPropDefinition(p *[]models.MetaPropDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metapropdefinitions []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		metapropdefinition := make(map[string]interface{})
		metapropdefinition["api_access"] = item.ApiAccess
		metapropdefinition["class_id"] = item.ClassId
		metapropdefinition["name"] = item.Name
		metapropdefinition["object_type"] = item.ObjectType
		metapropdefinition["op_security"] = item.OpSecurity
		metapropdefinition["search_weight"] = item.SearchWeight
		metapropdefinitions = append(metapropdefinitions, metapropdefinition)
	}
	return metapropdefinitions
}
func flattenListMetaRelationshipDefinition(p *[]models.MetaRelationshipDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metarelationshipdefinitions []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		metarelationshipdefinition := make(map[string]interface{})
		metarelationshipdefinition["api_access"] = item.ApiAccess
		metarelationshipdefinition["class_id"] = item.ClassId
		metarelationshipdefinition["collection"] = item.Collection
		metarelationshipdefinition["name"] = item.Name
		metarelationshipdefinition["object_type"] = item.ObjectType
		metarelationshipdefinition["type"] = item.Type
		metarelationshipdefinitions = append(metarelationshipdefinitions, metarelationshipdefinition)
	}
	return metarelationshipdefinitions
}
func flattenListMoBaseMoRelationship(p *[]models.MoBaseMoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var mobasemorelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.MoBaseMoRelationshipInterface.(*models.MoBaseMo)
		mobasemorelationship := make(map[string]interface{})
		mobasemorelationships = append(mobasemorelationships, mobasemorelationship)
	}
	return mobasemorelationships
}
func flattenListMoTag(p *[]models.MoTag, d *schema.ResourceData) []map[string]interface{} {
	var motags []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		motag := make(map[string]interface{})
		motag["key"] = item.Key
		motag["value"] = item.Value
		motags = append(motags, motag)
	}
	return motags
}
func flattenListNetworkElementRelationship(p *[]models.NetworkElementRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.NetworkElementRelationshipInterface.(*models.MoMoRef)
		networkelementrelationship := make(map[string]interface{})
		networkelementrelationship["class_id"] = item.ClassId
		networkelementrelationship["link"] = item.Link
		networkelementrelationship["moid"] = item.Moid
		networkelementrelationship["object_type"] = item.ObjectType
		networkelementrelationship["selector"] = item.Selector
		networkelementrelationships = append(networkelementrelationships, networkelementrelationship)
	}
	return networkelementrelationships
}
func flattenListNiaapiDetail(p *[]models.NiaapiDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapidetails []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		niaapidetail := make(map[string]interface{})
		niaapidetail["chksum"] = item.Chksum
		niaapidetail["class_id"] = item.ClassId
		niaapidetail["filename"] = item.Filename
		niaapidetail["name"] = item.Name
		niaapidetail["object_type"] = item.ObjectType
		niaapidetails = append(niaapidetails, niaapidetail)
	}
	return niaapidetails
}
func flattenListNiaapiRevisionInfo(p *[]models.NiaapiRevisionInfo, d *schema.ResourceData) []map[string]interface{} {
	var niaapirevisioninfos []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		niaapirevisioninfo := make(map[string]interface{})
		niaapirevisioninfo["class_id"] = item.ClassId
		niaapirevisioninfo["date_published"] = item.DatePublished
		niaapirevisioninfo["object_type"] = item.ObjectType
		niaapirevisioninfo["revision_comment"] = item.RevisionComment
		niaapirevisioninfo["revision_no"] = item.RevisionNo
		niaapirevisioninfos = append(niaapirevisioninfos, niaapirevisioninfo)
	}
	return niaapirevisioninfos
}
func flattenListOnpremImagePackage(p *[]models.OnpremImagePackage, d *schema.ResourceData) []map[string]interface{} {
	var onpremimagepackages []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		onpremimagepackage := make(map[string]interface{})
		onpremimagepackage["class_id"] = item.ClassId
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
func flattenListOnpremUpgradeNote(p *[]models.OnpremUpgradeNote, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradenotes []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		onpremupgradenote := make(map[string]interface{})
		onpremupgradenote["class_id"] = item.ClassId
		onpremupgradenote["message"] = item.Message
		onpremupgradenote["object_type"] = item.ObjectType
		onpremupgradenotes = append(onpremupgradenotes, onpremupgradenote)
	}
	return onpremupgradenotes
}
func flattenListOnpremUpgradePhase(p *[]models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradephases []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		onpremupgradephase := make(map[string]interface{})
		onpremupgradephase["class_id"] = item.ClassId
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
func flattenListOrganizationOrganizationRelationship(p *[]models.OrganizationOrganizationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.OrganizationOrganizationRelationshipInterface.(*models.MoMoRef)
		organizationorganizationrelationship := make(map[string]interface{})
		organizationorganizationrelationship["class_id"] = item.ClassId
		organizationorganizationrelationship["link"] = item.Link
		organizationorganizationrelationship["moid"] = item.Moid
		organizationorganizationrelationship["object_type"] = item.ObjectType
		organizationorganizationrelationship["selector"] = item.Selector
		organizationorganizationrelationships = append(organizationorganizationrelationships, organizationorganizationrelationship)
	}
	return organizationorganizationrelationships
}
func flattenListOsConfigurationFileRelationship(p *[]models.OsConfigurationFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.OsConfigurationFileRelationshipInterface.(*models.MoMoRef)
		osconfigurationfilerelationship := make(map[string]interface{})
		osconfigurationfilerelationship["class_id"] = item.ClassId
		osconfigurationfilerelationship["link"] = item.Link
		osconfigurationfilerelationship["moid"] = item.Moid
		osconfigurationfilerelationship["object_type"] = item.ObjectType
		osconfigurationfilerelationship["selector"] = item.Selector
		osconfigurationfilerelationships = append(osconfigurationfilerelationships, osconfigurationfilerelationship)
	}
	return osconfigurationfilerelationships
}
func flattenListOsPlaceHolder(p *[]models.OsPlaceHolder, d *schema.ResourceData) []map[string]interface{} {
	var osplaceholders []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		osplaceholder := make(map[string]interface{})
		osplaceholder["class_id"] = item.ClassId
		osplaceholder["is_value_set"] = item.IsValueSet
		osplaceholder["object_type"] = item.ObjectType
		osplaceholder["type"] = (func(p *models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {
			var workflowprimitivedatatypes []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			workflowprimitivedatatype := make(map[string]interface{})
			workflowprimitivedatatype["class_id"] = item.ClassId
			workflowprimitivedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["class_id"] = item.ClassId
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
				workflowprimitivedataproperty["class_id"] = item.ClassId
				workflowprimitivedataproperty["constraints"] = (func(p *models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {
					var workflowconstraintss []map[string]interface{}
					if p == nil {
						return nil
					}
					item := *p
					workflowconstraints := make(map[string]interface{})
					workflowconstraints["class_id"] = item.ClassId
					workflowconstraints["enum_list"] = (func(p *[]models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
						var workflowenumentrys []map[string]interface{}
						if *p == nil {
							return nil
						}
						for _, item := range *p {
							workflowenumentry := make(map[string]interface{})
							workflowenumentry["class_id"] = item.ClassId
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
				workflowprimitivedataproperty["inventory_selector"] = (func(p *[]models.WorkflowMoReferenceProperty, d *schema.ResourceData) []map[string]interface{} {
					var workflowmoreferencepropertys []map[string]interface{}
					if *p == nil {
						return nil
					}
					for _, item := range *p {
						workflowmoreferenceproperty := make(map[string]interface{})
						workflowmoreferenceproperty["class_id"] = item.ClassId
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
func flattenListPciCoprocessorCardRelationship(p *[]models.PciCoprocessorCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcicoprocessorcardrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.PciCoprocessorCardRelationshipInterface.(*models.MoMoRef)
		pcicoprocessorcardrelationship := make(map[string]interface{})
		pcicoprocessorcardrelationship["class_id"] = item.ClassId
		pcicoprocessorcardrelationship["link"] = item.Link
		pcicoprocessorcardrelationship["moid"] = item.Moid
		pcicoprocessorcardrelationship["object_type"] = item.ObjectType
		pcicoprocessorcardrelationship["selector"] = item.Selector
		pcicoprocessorcardrelationships = append(pcicoprocessorcardrelationships, pcicoprocessorcardrelationship)
	}
	return pcicoprocessorcardrelationships
}
func flattenListPciDeviceRelationship(p *[]models.PciDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcidevicerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.PciDeviceRelationshipInterface.(*models.MoMoRef)
		pcidevicerelationship := make(map[string]interface{})
		pcidevicerelationship["class_id"] = item.ClassId
		pcidevicerelationship["link"] = item.Link
		pcidevicerelationship["moid"] = item.Moid
		pcidevicerelationship["object_type"] = item.ObjectType
		pcidevicerelationship["selector"] = item.Selector
		pcidevicerelationships = append(pcidevicerelationships, pcidevicerelationship)
	}
	return pcidevicerelationships
}
func flattenListPciLinkRelationship(p *[]models.PciLinkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcilinkrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.PciLinkRelationshipInterface.(*models.MoMoRef)
		pcilinkrelationship := make(map[string]interface{})
		pcilinkrelationship["class_id"] = item.ClassId
		pcilinkrelationship["link"] = item.Link
		pcilinkrelationship["moid"] = item.Moid
		pcilinkrelationship["object_type"] = item.ObjectType
		pcilinkrelationship["selector"] = item.Selector
		pcilinkrelationships = append(pcilinkrelationships, pcilinkrelationship)
	}
	return pcilinkrelationships
}
func flattenListPciSwitchRelationship(p *[]models.PciSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.PciSwitchRelationshipInterface.(*models.MoMoRef)
		pciswitchrelationship := make(map[string]interface{})
		pciswitchrelationship["class_id"] = item.ClassId
		pciswitchrelationship["link"] = item.Link
		pciswitchrelationship["moid"] = item.Moid
		pciswitchrelationship["object_type"] = item.ObjectType
		pciswitchrelationship["selector"] = item.Selector
		pciswitchrelationships = append(pciswitchrelationships, pciswitchrelationship)
	}
	return pciswitchrelationships
}
func flattenListPolicyAbstractConfigProfileRelationship(p *[]models.PolicyAbstractConfigProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractconfigprofilerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.PolicyAbstractConfigProfileRelationshipInterface.(*models.MoMoRef)
		policyabstractconfigprofilerelationship := make(map[string]interface{})
		policyabstractconfigprofilerelationship["class_id"] = item.ClassId
		policyabstractconfigprofilerelationship["link"] = item.Link
		policyabstractconfigprofilerelationship["moid"] = item.Moid
		policyabstractconfigprofilerelationship["object_type"] = item.ObjectType
		policyabstractconfigprofilerelationship["selector"] = item.Selector
		policyabstractconfigprofilerelationships = append(policyabstractconfigprofilerelationships, policyabstractconfigprofilerelationship)
	}
	return policyabstractconfigprofilerelationships
}
func flattenListPolicyinventoryJobInfo(p *[]models.PolicyinventoryJobInfo, d *schema.ResourceData) []map[string]interface{} {
	var policyinventoryjobinfos []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		policyinventoryjobinfo := make(map[string]interface{})
		policyinventoryjobinfo["class_id"] = item.ClassId
		policyinventoryjobinfo["execution_status"] = item.ExecutionStatus
		policyinventoryjobinfo["last_scheduled_time"] = item.LastScheduledTime
		policyinventoryjobinfo["object_type"] = item.ObjectType
		policyinventoryjobinfo["policy_id"] = item.PolicyId
		policyinventoryjobinfo["policy_name"] = item.PolicyName
		policyinventoryjobinfos = append(policyinventoryjobinfos, policyinventoryjobinfo)
	}
	return policyinventoryjobinfos
}
func flattenListPortGroupRelationship(p *[]models.PortGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.PortGroupRelationshipInterface.(*models.MoMoRef)
		portgrouprelationship := make(map[string]interface{})
		portgrouprelationship["class_id"] = item.ClassId
		portgrouprelationship["link"] = item.Link
		portgrouprelationship["moid"] = item.Moid
		portgrouprelationship["object_type"] = item.ObjectType
		portgrouprelationship["selector"] = item.Selector
		portgrouprelationships = append(portgrouprelationships, portgrouprelationship)
	}
	return portgrouprelationships
}
func flattenListPortSubGroupRelationship(p *[]models.PortSubGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.PortSubGroupRelationshipInterface.(*models.MoMoRef)
		portsubgrouprelationship := make(map[string]interface{})
		portsubgrouprelationship["class_id"] = item.ClassId
		portsubgrouprelationship["link"] = item.Link
		portsubgrouprelationship["moid"] = item.Moid
		portsubgrouprelationship["object_type"] = item.ObjectType
		portsubgrouprelationship["selector"] = item.Selector
		portsubgrouprelationships = append(portsubgrouprelationships, portsubgrouprelationship)
	}
	return portsubgrouprelationships
}
func flattenListProcessorUnitRelationship(p *[]models.ProcessorUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var processorunitrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ProcessorUnitRelationshipInterface.(*models.MoMoRef)
		processorunitrelationship := make(map[string]interface{})
		processorunitrelationship["class_id"] = item.ClassId
		processorunitrelationship["link"] = item.Link
		processorunitrelationship["moid"] = item.Moid
		processorunitrelationship["object_type"] = item.ObjectType
		processorunitrelationship["selector"] = item.Selector
		processorunitrelationships = append(processorunitrelationships, processorunitrelationship)
	}
	return processorunitrelationships
}
func flattenListRecoveryBackupProfileRelationship(p *[]models.RecoveryBackupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.RecoveryBackupProfileRelationshipInterface.(*models.MoMoRef)
		recoverybackupprofilerelationship := make(map[string]interface{})
		recoverybackupprofilerelationship["class_id"] = item.ClassId
		recoverybackupprofilerelationship["link"] = item.Link
		recoverybackupprofilerelationship["moid"] = item.Moid
		recoverybackupprofilerelationship["object_type"] = item.ObjectType
		recoverybackupprofilerelationship["selector"] = item.Selector
		recoverybackupprofilerelationships = append(recoverybackupprofilerelationships, recoverybackupprofilerelationship)
	}
	return recoverybackupprofilerelationships
}
func flattenListRecoveryConfigResultEntryRelationship(p *[]models.RecoveryConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultentryrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.RecoveryConfigResultEntryRelationshipInterface.(*models.MoMoRef)
		recoveryconfigresultentryrelationship := make(map[string]interface{})
		recoveryconfigresultentryrelationship["class_id"] = item.ClassId
		recoveryconfigresultentryrelationship["link"] = item.Link
		recoveryconfigresultentryrelationship["moid"] = item.Moid
		recoveryconfigresultentryrelationship["object_type"] = item.ObjectType
		recoveryconfigresultentryrelationship["selector"] = item.Selector
		recoveryconfigresultentryrelationships = append(recoveryconfigresultentryrelationships, recoveryconfigresultentryrelationship)
	}
	return recoveryconfigresultentryrelationships
}
func flattenListResourceGroupRelationship(p *[]models.ResourceGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ResourceGroupRelationshipInterface.(*models.MoMoRef)
		resourcegrouprelationship := make(map[string]interface{})
		resourcegrouprelationship["class_id"] = item.ClassId
		resourcegrouprelationship["link"] = item.Link
		resourcegrouprelationship["moid"] = item.Moid
		resourcegrouprelationship["object_type"] = item.ObjectType
		resourcegrouprelationship["selector"] = item.Selector
		resourcegrouprelationships = append(resourcegrouprelationships, resourcegrouprelationship)
	}
	return resourcegrouprelationships
}
func flattenListResourceMembershipRelationship(p *[]models.ResourceMembershipRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcemembershiprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ResourceMembershipRelationshipInterface.(*models.MoMoRef)
		resourcemembershiprelationship := make(map[string]interface{})
		resourcemembershiprelationship["class_id"] = item.ClassId
		resourcemembershiprelationship["link"] = item.Link
		resourcemembershiprelationship["moid"] = item.Moid
		resourcemembershiprelationship["object_type"] = item.ObjectType
		resourcemembershiprelationship["selector"] = item.Selector
		resourcemembershiprelationships = append(resourcemembershiprelationships, resourcemembershiprelationship)
	}
	return resourcemembershiprelationships
}
func flattenListResourcePerTypeCombinedSelector(p *[]models.ResourcePerTypeCombinedSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourcepertypecombinedselectors []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		resourcepertypecombinedselector := make(map[string]interface{})
		resourcepertypecombinedselector["class_id"] = item.ClassId
		resourcepertypecombinedselector["combined_selector"] = item.CombinedSelector
		resourcepertypecombinedselector["empty_filter"] = item.EmptyFilter
		resourcepertypecombinedselector["object_type"] = item.ObjectType
		resourcepertypecombinedselector["selector_object_type"] = item.SelectorObjectType
		resourcepertypecombinedselectors = append(resourcepertypecombinedselectors, resourcepertypecombinedselector)
	}
	return resourcepertypecombinedselectors
}
func flattenListResourceSelector(p *[]models.ResourceSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourceselectors []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		resourceselector := make(map[string]interface{})
		resourceselector["class_id"] = item.ClassId
		resourceselector["object_type"] = item.ObjectType
		resourceselector["selector"] = item.Selector
		resourceselectors = append(resourceselectors, resourceselector)
	}
	return resourceselectors
}
func flattenListSdcardPartition(p *[]models.SdcardPartition, d *schema.ResourceData) []map[string]interface{} {
	var sdcardpartitions []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		sdcardpartition := make(map[string]interface{})
		sdcardpartition["class_id"] = item.ClassId
		sdcardpartition["object_type"] = item.ObjectType
		sdcardpartition["type"] = item.Type
		sdcardpartition["virtual_drives"] = (func(p *[]models.SdcardVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
			var sdcardvirtualdrives []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				sdcardvirtualdrive := make(map[string]interface{})
				sdcardvirtualdrive["class_id"] = item.ClassId
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
func flattenListSdwanNetworkConfigurationType(p *[]models.SdwanNetworkConfigurationType, d *schema.ResourceData) []map[string]interface{} {
	var sdwannetworkconfigurationtypes []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		sdwannetworkconfigurationtype := make(map[string]interface{})
		sdwannetworkconfigurationtype["class_id"] = item.ClassId
		sdwannetworkconfigurationtype["network_type"] = item.NetworkType
		sdwannetworkconfigurationtype["object_type"] = item.ObjectType
		sdwannetworkconfigurationtype["port_group"] = item.PortGroup
		sdwannetworkconfigurationtype["vlan"] = item.Vlan
		sdwannetworkconfigurationtypes = append(sdwannetworkconfigurationtypes, sdwannetworkconfigurationtype)
	}
	return sdwannetworkconfigurationtypes
}
func flattenListSdwanProfileRelationship(p *[]models.SdwanProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.SdwanProfileRelationshipInterface.(*models.MoMoRef)
		sdwanprofilerelationship := make(map[string]interface{})
		sdwanprofilerelationship["class_id"] = item.ClassId
		sdwanprofilerelationship["link"] = item.Link
		sdwanprofilerelationship["moid"] = item.Moid
		sdwanprofilerelationship["object_type"] = item.ObjectType
		sdwanprofilerelationship["selector"] = item.Selector
		sdwanprofilerelationships = append(sdwanprofilerelationships, sdwanprofilerelationship)
	}
	return sdwanprofilerelationships
}
func flattenListSdwanRouterNodeRelationship(p *[]models.SdwanRouterNodeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouternoderelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.SdwanRouterNodeRelationshipInterface.(*models.MoMoRef)
		sdwanrouternoderelationship := make(map[string]interface{})
		sdwanrouternoderelationship["class_id"] = item.ClassId
		sdwanrouternoderelationship["link"] = item.Link
		sdwanrouternoderelationship["moid"] = item.Moid
		sdwanrouternoderelationship["object_type"] = item.ObjectType
		sdwanrouternoderelationship["selector"] = item.Selector
		sdwanrouternoderelationships = append(sdwanrouternoderelationships, sdwanrouternoderelationship)
	}
	return sdwanrouternoderelationships
}
func flattenListSdwanTemplateInputsType(p *[]models.SdwanTemplateInputsType, d *schema.ResourceData) []map[string]interface{} {
	var sdwantemplateinputstypes []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		sdwantemplateinputstype := make(map[string]interface{})
		sdwantemplateinputstype["class_id"] = item.ClassId
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
func flattenListSecurityUnitRelationship(p *[]models.SecurityUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var securityunitrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.SecurityUnitRelationshipInterface.(*models.MoMoRef)
		securityunitrelationship := make(map[string]interface{})
		securityunitrelationship["class_id"] = item.ClassId
		securityunitrelationship["link"] = item.Link
		securityunitrelationship["moid"] = item.Moid
		securityunitrelationship["object_type"] = item.ObjectType
		securityunitrelationship["selector"] = item.Selector
		securityunitrelationships = append(securityunitrelationships, securityunitrelationship)
	}
	return securityunitrelationships
}
func flattenListServerConfigChangeDetailRelationship(p *[]models.ServerConfigChangeDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigchangedetailrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ServerConfigChangeDetailRelationshipInterface.(*models.MoMoRef)
		serverconfigchangedetailrelationship := make(map[string]interface{})
		serverconfigchangedetailrelationship["class_id"] = item.ClassId
		serverconfigchangedetailrelationship["link"] = item.Link
		serverconfigchangedetailrelationship["moid"] = item.Moid
		serverconfigchangedetailrelationship["object_type"] = item.ObjectType
		serverconfigchangedetailrelationship["selector"] = item.Selector
		serverconfigchangedetailrelationships = append(serverconfigchangedetailrelationships, serverconfigchangedetailrelationship)
	}
	return serverconfigchangedetailrelationships
}
func flattenListServerConfigResultEntryRelationship(p *[]models.ServerConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultentryrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.ServerConfigResultEntryRelationshipInterface.(*models.MoMoRef)
		serverconfigresultentryrelationship := make(map[string]interface{})
		serverconfigresultentryrelationship["class_id"] = item.ClassId
		serverconfigresultentryrelationship["link"] = item.Link
		serverconfigresultentryrelationship["moid"] = item.Moid
		serverconfigresultentryrelationship["object_type"] = item.ObjectType
		serverconfigresultentryrelationship["selector"] = item.Selector
		serverconfigresultentryrelationships = append(serverconfigresultentryrelationships, serverconfigresultentryrelationship)
	}
	return serverconfigresultentryrelationships
}
func flattenListSnmpTrap(p *[]models.SnmpTrap, d *schema.ResourceData) []map[string]interface{} {
	var snmptraps []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		snmptrap := make(map[string]interface{})
		snmptrap["class_id"] = item.ClassId
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
func flattenListSnmpUser(p *[]models.SnmpUser, d *schema.ResourceData) []map[string]interface{} {
	var snmpusers []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		snmpuser := make(map[string]interface{})
		auth_password_x := d.Get("snmp_users").([]interface{})
		auth_password_y := auth_password_x[0].(map[string]interface{})
		snmpuser["auth_password"] = auth_password_y["auth_password"]
		snmpuser["auth_type"] = item.AuthType
		snmpuser["class_id"] = item.ClassId
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
func flattenListStorageControllerRelationship(p *[]models.StorageControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageControllerRelationshipInterface.(*models.MoMoRef)
		storagecontrollerrelationship := make(map[string]interface{})
		storagecontrollerrelationship["class_id"] = item.ClassId
		storagecontrollerrelationship["link"] = item.Link
		storagecontrollerrelationship["moid"] = item.Moid
		storagecontrollerrelationship["object_type"] = item.ObjectType
		storagecontrollerrelationship["selector"] = item.Selector
		storagecontrollerrelationships = append(storagecontrollerrelationships, storagecontrollerrelationship)
	}
	return storagecontrollerrelationships
}
func flattenListStorageDiskGroupPolicyRelationship(p *[]models.StorageDiskGroupPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouppolicyrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageDiskGroupPolicyRelationshipInterface.(*models.MoMoRef)
		storagediskgrouppolicyrelationship := make(map[string]interface{})
		storagediskgrouppolicyrelationship["class_id"] = item.ClassId
		storagediskgrouppolicyrelationship["link"] = item.Link
		storagediskgrouppolicyrelationship["moid"] = item.Moid
		storagediskgrouppolicyrelationship["object_type"] = item.ObjectType
		storagediskgrouppolicyrelationship["selector"] = item.Selector
		storagediskgrouppolicyrelationships = append(storagediskgrouppolicyrelationships, storagediskgrouppolicyrelationship)
	}
	return storagediskgrouppolicyrelationships
}
func flattenListStorageEnclosureRelationship(p *[]models.StorageEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageEnclosureRelationshipInterface.(*models.MoMoRef)
		storageenclosurerelationship := make(map[string]interface{})
		storageenclosurerelationship["class_id"] = item.ClassId
		storageenclosurerelationship["link"] = item.Link
		storageenclosurerelationship["moid"] = item.Moid
		storageenclosurerelationship["object_type"] = item.ObjectType
		storageenclosurerelationship["selector"] = item.Selector
		storageenclosurerelationships = append(storageenclosurerelationships, storageenclosurerelationship)
	}
	return storageenclosurerelationships
}
func flattenListStorageEnclosureDiskRelationship(p *[]models.StorageEnclosureDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurediskrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageEnclosureDiskRelationshipInterface.(*models.MoMoRef)
		storageenclosurediskrelationship := make(map[string]interface{})
		storageenclosurediskrelationship["class_id"] = item.ClassId
		storageenclosurediskrelationship["link"] = item.Link
		storageenclosurediskrelationship["moid"] = item.Moid
		storageenclosurediskrelationship["object_type"] = item.ObjectType
		storageenclosurediskrelationship["selector"] = item.Selector
		storageenclosurediskrelationships = append(storageenclosurediskrelationships, storageenclosurediskrelationship)
	}
	return storageenclosurediskrelationships
}
func flattenListStorageEnclosureDiskSlotEpRelationship(p *[]models.StorageEnclosureDiskSlotEpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosuredisksloteprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageEnclosureDiskSlotEpRelationshipInterface.(*models.MoMoRef)
		storageenclosuredisksloteprelationship := make(map[string]interface{})
		storageenclosuredisksloteprelationship["class_id"] = item.ClassId
		storageenclosuredisksloteprelationship["link"] = item.Link
		storageenclosuredisksloteprelationship["moid"] = item.Moid
		storageenclosuredisksloteprelationship["object_type"] = item.ObjectType
		storageenclosuredisksloteprelationship["selector"] = item.Selector
		storageenclosuredisksloteprelationships = append(storageenclosuredisksloteprelationships, storageenclosuredisksloteprelationship)
	}
	return storageenclosuredisksloteprelationships
}
func flattenListStorageFlexFlashControllerRelationship(p *[]models.StorageFlexFlashControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageFlexFlashControllerRelationshipInterface.(*models.MoMoRef)
		storageflexflashcontrollerrelationship := make(map[string]interface{})
		storageflexflashcontrollerrelationship["class_id"] = item.ClassId
		storageflexflashcontrollerrelationship["link"] = item.Link
		storageflexflashcontrollerrelationship["moid"] = item.Moid
		storageflexflashcontrollerrelationship["object_type"] = item.ObjectType
		storageflexflashcontrollerrelationship["selector"] = item.Selector
		storageflexflashcontrollerrelationships = append(storageflexflashcontrollerrelationships, storageflexflashcontrollerrelationship)
	}
	return storageflexflashcontrollerrelationships
}
func flattenListStorageFlexFlashControllerPropsRelationship(p *[]models.StorageFlexFlashControllerPropsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerpropsrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageFlexFlashControllerPropsRelationshipInterface.(*models.MoMoRef)
		storageflexflashcontrollerpropsrelationship := make(map[string]interface{})
		storageflexflashcontrollerpropsrelationship["class_id"] = item.ClassId
		storageflexflashcontrollerpropsrelationship["link"] = item.Link
		storageflexflashcontrollerpropsrelationship["moid"] = item.Moid
		storageflexflashcontrollerpropsrelationship["object_type"] = item.ObjectType
		storageflexflashcontrollerpropsrelationship["selector"] = item.Selector
		storageflexflashcontrollerpropsrelationships = append(storageflexflashcontrollerpropsrelationships, storageflexflashcontrollerpropsrelationship)
	}
	return storageflexflashcontrollerpropsrelationships
}
func flattenListStorageFlexFlashPhysicalDriveRelationship(p *[]models.StorageFlexFlashPhysicalDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashphysicaldriverelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageFlexFlashPhysicalDriveRelationshipInterface.(*models.MoMoRef)
		storageflexflashphysicaldriverelationship := make(map[string]interface{})
		storageflexflashphysicaldriverelationship["class_id"] = item.ClassId
		storageflexflashphysicaldriverelationship["link"] = item.Link
		storageflexflashphysicaldriverelationship["moid"] = item.Moid
		storageflexflashphysicaldriverelationship["object_type"] = item.ObjectType
		storageflexflashphysicaldriverelationship["selector"] = item.Selector
		storageflexflashphysicaldriverelationships = append(storageflexflashphysicaldriverelationships, storageflexflashphysicaldriverelationship)
	}
	return storageflexflashphysicaldriverelationships
}
func flattenListStorageFlexFlashVirtualDriveRelationship(p *[]models.StorageFlexFlashVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashvirtualdriverelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageFlexFlashVirtualDriveRelationshipInterface.(*models.MoMoRef)
		storageflexflashvirtualdriverelationship := make(map[string]interface{})
		storageflexflashvirtualdriverelationship["class_id"] = item.ClassId
		storageflexflashvirtualdriverelationship["link"] = item.Link
		storageflexflashvirtualdriverelationship["moid"] = item.Moid
		storageflexflashvirtualdriverelationship["object_type"] = item.ObjectType
		storageflexflashvirtualdriverelationship["selector"] = item.Selector
		storageflexflashvirtualdriverelationships = append(storageflexflashvirtualdriverelationships, storageflexflashvirtualdriverelationship)
	}
	return storageflexflashvirtualdriverelationships
}
func flattenListStorageFlexUtilControllerRelationship(p *[]models.StorageFlexUtilControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageFlexUtilControllerRelationshipInterface.(*models.MoMoRef)
		storageflexutilcontrollerrelationship := make(map[string]interface{})
		storageflexutilcontrollerrelationship["class_id"] = item.ClassId
		storageflexutilcontrollerrelationship["link"] = item.Link
		storageflexutilcontrollerrelationship["moid"] = item.Moid
		storageflexutilcontrollerrelationship["object_type"] = item.ObjectType
		storageflexutilcontrollerrelationship["selector"] = item.Selector
		storageflexutilcontrollerrelationships = append(storageflexutilcontrollerrelationships, storageflexutilcontrollerrelationship)
	}
	return storageflexutilcontrollerrelationships
}
func flattenListStorageFlexUtilPhysicalDriveRelationship(p *[]models.StorageFlexUtilPhysicalDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilphysicaldriverelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageFlexUtilPhysicalDriveRelationshipInterface.(*models.MoMoRef)
		storageflexutilphysicaldriverelationship := make(map[string]interface{})
		storageflexutilphysicaldriverelationship["class_id"] = item.ClassId
		storageflexutilphysicaldriverelationship["link"] = item.Link
		storageflexutilphysicaldriverelationship["moid"] = item.Moid
		storageflexutilphysicaldriverelationship["object_type"] = item.ObjectType
		storageflexutilphysicaldriverelationship["selector"] = item.Selector
		storageflexutilphysicaldriverelationships = append(storageflexutilphysicaldriverelationships, storageflexutilphysicaldriverelationship)
	}
	return storageflexutilphysicaldriverelationships
}
func flattenListStorageFlexUtilVirtualDriveRelationship(p *[]models.StorageFlexUtilVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilvirtualdriverelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageFlexUtilVirtualDriveRelationshipInterface.(*models.MoMoRef)
		storageflexutilvirtualdriverelationship := make(map[string]interface{})
		storageflexutilvirtualdriverelationship["class_id"] = item.ClassId
		storageflexutilvirtualdriverelationship["link"] = item.Link
		storageflexutilvirtualdriverelationship["moid"] = item.Moid
		storageflexutilvirtualdriverelationship["object_type"] = item.ObjectType
		storageflexutilvirtualdriverelationship["selector"] = item.Selector
		storageflexutilvirtualdriverelationships = append(storageflexutilvirtualdriverelationships, storageflexutilvirtualdriverelationship)
	}
	return storageflexutilvirtualdriverelationships
}
func flattenListStorageInitiator(p *[]models.StorageInitiator, d *schema.ResourceData) []map[string]interface{} {
	var storageinitiators []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		storageinitiator := make(map[string]interface{})
		storageinitiator["class_id"] = item.ClassId
		storageinitiator["iqn"] = item.Iqn
		storageinitiator["name"] = item.Name
		storageinitiator["object_type"] = item.ObjectType
		storageinitiator["type"] = item.Type
		storageinitiator["wwn"] = item.Wwn
		storageinitiators = append(storageinitiators, storageinitiator)
	}
	return storageinitiators
}
func flattenListStorageLocalDisk(p *[]models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
	var storagelocaldisks []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		storagelocaldisk := make(map[string]interface{})
		storagelocaldisk["class_id"] = item.ClassId
		storagelocaldisk["object_type"] = item.ObjectType
		storagelocaldisk["slot_number"] = item.SlotNumber
		storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
	}
	return storagelocaldisks
}
func flattenListStoragePhysicalDiskRelationship(p *[]models.StoragePhysicalDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StoragePhysicalDiskRelationshipInterface.(*models.MoMoRef)
		storagephysicaldiskrelationship := make(map[string]interface{})
		storagephysicaldiskrelationship["class_id"] = item.ClassId
		storagephysicaldiskrelationship["link"] = item.Link
		storagephysicaldiskrelationship["moid"] = item.Moid
		storagephysicaldiskrelationship["object_type"] = item.ObjectType
		storagephysicaldiskrelationship["selector"] = item.Selector
		storagephysicaldiskrelationships = append(storagephysicaldiskrelationships, storagephysicaldiskrelationship)
	}
	return storagephysicaldiskrelationships
}
func flattenListStoragePhysicalDiskExtensionRelationship(p *[]models.StoragePhysicalDiskExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskextensionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StoragePhysicalDiskExtensionRelationshipInterface.(*models.MoMoRef)
		storagephysicaldiskextensionrelationship := make(map[string]interface{})
		storagephysicaldiskextensionrelationship["class_id"] = item.ClassId
		storagephysicaldiskextensionrelationship["link"] = item.Link
		storagephysicaldiskextensionrelationship["moid"] = item.Moid
		storagephysicaldiskextensionrelationship["object_type"] = item.ObjectType
		storagephysicaldiskextensionrelationship["selector"] = item.Selector
		storagephysicaldiskextensionrelationships = append(storagephysicaldiskextensionrelationships, storagephysicaldiskextensionrelationship)
	}
	return storagephysicaldiskextensionrelationships
}
func flattenListStoragePhysicalDiskUsageRelationship(p *[]models.StoragePhysicalDiskUsageRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskusagerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StoragePhysicalDiskUsageRelationshipInterface.(*models.MoMoRef)
		storagephysicaldiskusagerelationship := make(map[string]interface{})
		storagephysicaldiskusagerelationship["class_id"] = item.ClassId
		storagephysicaldiskusagerelationship["link"] = item.Link
		storagephysicaldiskusagerelationship["moid"] = item.Moid
		storagephysicaldiskusagerelationship["object_type"] = item.ObjectType
		storagephysicaldiskusagerelationship["selector"] = item.Selector
		storagephysicaldiskusagerelationships = append(storagephysicaldiskusagerelationships, storagephysicaldiskusagerelationship)
	}
	return storagephysicaldiskusagerelationships
}
func flattenListStoragePureHostRelationship(p *[]models.StoragePureHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StoragePureHostRelationshipInterface.(*models.MoMoRef)
		storagepurehostrelationship := make(map[string]interface{})
		storagepurehostrelationship["class_id"] = item.ClassId
		storagepurehostrelationship["link"] = item.Link
		storagepurehostrelationship["moid"] = item.Moid
		storagepurehostrelationship["object_type"] = item.ObjectType
		storagepurehostrelationship["selector"] = item.Selector
		storagepurehostrelationships = append(storagepurehostrelationships, storagepurehostrelationship)
	}
	return storagepurehostrelationships
}
func flattenListStoragePureHostGroupRelationship(p *[]models.StoragePureHostGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StoragePureHostGroupRelationshipInterface.(*models.MoMoRef)
		storagepurehostgrouprelationship := make(map[string]interface{})
		storagepurehostgrouprelationship["class_id"] = item.ClassId
		storagepurehostgrouprelationship["link"] = item.Link
		storagepurehostgrouprelationship["moid"] = item.Moid
		storagepurehostgrouprelationship["object_type"] = item.ObjectType
		storagepurehostgrouprelationship["selector"] = item.Selector
		storagepurehostgrouprelationships = append(storagepurehostgrouprelationships, storagepurehostgrouprelationship)
	}
	return storagepurehostgrouprelationships
}
func flattenListStoragePureReplicationBlackout(p *[]models.StoragePureReplicationBlackout, d *schema.ResourceData) []map[string]interface{} {
	var storagepurereplicationblackouts []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		storagepurereplicationblackout := make(map[string]interface{})
		storagepurereplicationblackout["class_id"] = item.ClassId
		storagepurereplicationblackout["end"] = item.End
		storagepurereplicationblackout["object_type"] = item.ObjectType
		storagepurereplicationblackout["start"] = item.Start
		storagepurereplicationblackouts = append(storagepurereplicationblackouts, storagepurereplicationblackout)
	}
	return storagepurereplicationblackouts
}
func flattenListStoragePureVolumeRelationship(p *[]models.StoragePureVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurevolumerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StoragePureVolumeRelationshipInterface.(*models.MoMoRef)
		storagepurevolumerelationship := make(map[string]interface{})
		storagepurevolumerelationship["class_id"] = item.ClassId
		storagepurevolumerelationship["link"] = item.Link
		storagepurevolumerelationship["moid"] = item.Moid
		storagepurevolumerelationship["object_type"] = item.ObjectType
		storagepurevolumerelationship["selector"] = item.Selector
		storagepurevolumerelationships = append(storagepurevolumerelationships, storagepurevolumerelationship)
	}
	return storagepurevolumerelationships
}
func flattenListStorageSasExpanderRelationship(p *[]models.StorageSasExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageSasExpanderRelationshipInterface.(*models.MoMoRef)
		storagesasexpanderrelationship := make(map[string]interface{})
		storagesasexpanderrelationship["class_id"] = item.ClassId
		storagesasexpanderrelationship["link"] = item.Link
		storagesasexpanderrelationship["moid"] = item.Moid
		storagesasexpanderrelationship["object_type"] = item.ObjectType
		storagesasexpanderrelationship["selector"] = item.Selector
		storagesasexpanderrelationships = append(storagesasexpanderrelationships, storagesasexpanderrelationship)
	}
	return storagesasexpanderrelationships
}
func flattenListStorageSasPortRelationship(p *[]models.StorageSasPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasportrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageSasPortRelationshipInterface.(*models.MoMoRef)
		storagesasportrelationship := make(map[string]interface{})
		storagesasportrelationship["class_id"] = item.ClassId
		storagesasportrelationship["link"] = item.Link
		storagesasportrelationship["moid"] = item.Moid
		storagesasportrelationship["object_type"] = item.ObjectType
		storagesasportrelationship["selector"] = item.Selector
		storagesasportrelationships = append(storagesasportrelationships, storagesasportrelationship)
	}
	return storagesasportrelationships
}
func flattenListStorageSpanGroup(p *[]models.StorageSpanGroup, d *schema.ResourceData) []map[string]interface{} {
	var storagespangroups []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		storagespangroup := make(map[string]interface{})
		storagespangroup["class_id"] = item.ClassId
		storagespangroup["disks"] = (func(p *[]models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
			var storagelocaldisks []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				storagelocaldisk := make(map[string]interface{})
				storagelocaldisk["class_id"] = item.ClassId
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
func flattenListStorageStoragePolicyRelationship(p *[]models.StorageStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagestoragepolicyrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageStoragePolicyRelationshipInterface.(*models.MoMoRef)
		storagestoragepolicyrelationship := make(map[string]interface{})
		storagestoragepolicyrelationship["class_id"] = item.ClassId
		storagestoragepolicyrelationship["link"] = item.Link
		storagestoragepolicyrelationship["moid"] = item.Moid
		storagestoragepolicyrelationship["object_type"] = item.ObjectType
		storagestoragepolicyrelationship["selector"] = item.Selector
		storagestoragepolicyrelationships = append(storagestoragepolicyrelationships, storagestoragepolicyrelationship)
	}
	return storagestoragepolicyrelationships
}
func flattenListStorageVdMemberEpRelationship(p *[]models.StorageVdMemberEpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevdmembereprelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageVdMemberEpRelationshipInterface.(*models.MoMoRef)
		storagevdmembereprelationship := make(map[string]interface{})
		storagevdmembereprelationship["class_id"] = item.ClassId
		storagevdmembereprelationship["link"] = item.Link
		storagevdmembereprelationship["moid"] = item.Moid
		storagevdmembereprelationship["object_type"] = item.ObjectType
		storagevdmembereprelationship["selector"] = item.Selector
		storagevdmembereprelationships = append(storagevdmembereprelationships, storagevdmembereprelationship)
	}
	return storagevdmembereprelationships
}
func flattenListStorageVirtualDriveRelationship(p *[]models.StorageVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageVirtualDriveRelationshipInterface.(*models.MoMoRef)
		storagevirtualdriverelationship := make(map[string]interface{})
		storagevirtualdriverelationship["class_id"] = item.ClassId
		storagevirtualdriverelationship["link"] = item.Link
		storagevirtualdriverelationship["moid"] = item.Moid
		storagevirtualdriverelationship["object_type"] = item.ObjectType
		storagevirtualdriverelationship["selector"] = item.Selector
		storagevirtualdriverelationships = append(storagevirtualdriverelationships, storagevirtualdriverelationship)
	}
	return storagevirtualdriverelationships
}
func flattenListStorageVirtualDriveConfig(p *[]models.StorageVirtualDriveConfig, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveconfigs []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		storagevirtualdriveconfig := make(map[string]interface{})
		storagevirtualdriveconfig["access_policy"] = item.AccessPolicy
		storagevirtualdriveconfig["boot_drive"] = item.BootDrive
		storagevirtualdriveconfig["class_id"] = item.ClassId
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
func flattenListStorageVirtualDriveExtensionRelationship(p *[]models.StorageVirtualDriveExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.StorageVirtualDriveExtensionRelationshipInterface.(*models.MoMoRef)
		storagevirtualdriveextensionrelationship := make(map[string]interface{})
		storagevirtualdriveextensionrelationship["class_id"] = item.ClassId
		storagevirtualdriveextensionrelationship["link"] = item.Link
		storagevirtualdriveextensionrelationship["moid"] = item.Moid
		storagevirtualdriveextensionrelationship["object_type"] = item.ObjectType
		storagevirtualdriveextensionrelationship["selector"] = item.Selector
		storagevirtualdriveextensionrelationships = append(storagevirtualdriveextensionrelationships, storagevirtualdriveextensionrelationship)
	}
	return storagevirtualdriveextensionrelationships
}
func flattenListSyslogLocalClientBase(p *[]models.SyslogLocalClientBase, d *schema.ResourceData) []map[string]interface{} {
	var sysloglocalclientbases []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		sysloglocalclientbase := make(map[string]interface{})
		sysloglocalclientbase["class_id"] = item.ClassId
		sysloglocalclientbase["min_severity"] = item.MinSeverity
		sysloglocalclientbase["object_type"] = item.ObjectType
		sysloglocalclientbases = append(sysloglocalclientbases, sysloglocalclientbase)
	}
	return sysloglocalclientbases
}
func flattenListSyslogRemoteClientBase(p *[]models.SyslogRemoteClientBase, d *schema.ResourceData) []map[string]interface{} {
	var syslogremoteclientbases []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		syslogremoteclientbase := make(map[string]interface{})
		syslogremoteclientbase["class_id"] = item.ClassId
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
func flattenListTamAction(p *[]models.TamAction, d *schema.ResourceData) []map[string]interface{} {
	var tamactions []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		tamaction := make(map[string]interface{})
		tamaction["affected_object_type"] = item.AffectedObjectType
		tamaction["alert_type"] = item.AlertType
		tamaction["class_id"] = item.ClassId
		tamaction["identifiers"] = (func(p *[]models.TamIdentifiers, d *schema.ResourceData) []map[string]interface{} {
			var tamidentifierss []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				tamidentifiers := make(map[string]interface{})
				tamidentifiers["class_id"] = item.ClassId
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
		tamaction["queries"] = (func(p *[]models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["class_id"] = item.ClassId
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
func flattenListTamApiDataSource(p *[]models.TamApiDataSource, d *schema.ResourceData) []map[string]interface{} {
	var tamapidatasources []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		tamapidatasource := make(map[string]interface{})
		tamapidatasource["class_id"] = item.ClassId
		tamapidatasource["mo_type"] = item.MoType
		tamapidatasource["name"] = item.Name
		tamapidatasource["object_type"] = item.ObjectType
		tamapidatasource["queries"] = (func(p *[]models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if *p == nil {
				return nil
			}
			for _, item := range *p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["class_id"] = item.ClassId
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
func flattenListUcsdConnectorPack(p *[]models.UcsdConnectorPack, d *schema.ResourceData) []map[string]interface{} {
	var ucsdconnectorpacks []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		ucsdconnectorpack := make(map[string]interface{})
		ucsdconnectorpack["class_id"] = item.ClassId
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
func flattenListVirtualizationVmwareDatastoreRelationship(p *[]models.VirtualizationVmwareDatastoreRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatastorerelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.VirtualizationVmwareDatastoreRelationshipInterface.(*models.MoMoRef)
		virtualizationvmwaredatastorerelationship := make(map[string]interface{})
		virtualizationvmwaredatastorerelationship["class_id"] = item.ClassId
		virtualizationvmwaredatastorerelationship["link"] = item.Link
		virtualizationvmwaredatastorerelationship["moid"] = item.Moid
		virtualizationvmwaredatastorerelationship["object_type"] = item.ObjectType
		virtualizationvmwaredatastorerelationship["selector"] = item.Selector
		virtualizationvmwaredatastorerelationships = append(virtualizationvmwaredatastorerelationships, virtualizationvmwaredatastorerelationship)
	}
	return virtualizationvmwaredatastorerelationships
}
func flattenListVirtualizationVmwareHostRelationship(p *[]models.VirtualizationVmwareHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.VirtualizationVmwareHostRelationshipInterface.(*models.MoMoRef)
		virtualizationvmwarehostrelationship := make(map[string]interface{})
		virtualizationvmwarehostrelationship["class_id"] = item.ClassId
		virtualizationvmwarehostrelationship["link"] = item.Link
		virtualizationvmwarehostrelationship["moid"] = item.Moid
		virtualizationvmwarehostrelationship["object_type"] = item.ObjectType
		virtualizationvmwarehostrelationship["selector"] = item.Selector
		virtualizationvmwarehostrelationships = append(virtualizationvmwarehostrelationships, virtualizationvmwarehostrelationship)
	}
	return virtualizationvmwarehostrelationships
}
func flattenListVmediaMapping(p *[]models.VmediaMapping, d *schema.ResourceData) []map[string]interface{} {
	var vmediamappings []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		vmediamapping := make(map[string]interface{})
		vmediamapping["authentication_protocol"] = item.AuthenticationProtocol
		vmediamapping["class_id"] = item.ClassId
		vmediamapping["device_type"] = item.DeviceType
		vmediamapping["host_name"] = item.HostName
		vmediamapping["is_password_set"] = item.IsPasswordSet
		vmediamapping["mount_options"] = item.MountOptions
		vmediamapping["mount_protocol"] = item.MountProtocol
		vmediamapping["object_type"] = item.ObjectType
		password_x := d.Get("mappings").([]interface{})
		password_y := password_x[0].(map[string]interface{})
		vmediamapping["password"] = password_y["password"]
		vmediamapping["remote_file"] = item.RemoteFile
		vmediamapping["remote_path"] = item.RemotePath
		vmediamapping["username"] = item.Username
		vmediamapping["volume_name"] = item.VolumeName
		vmediamappings = append(vmediamappings, vmediamapping)
	}
	return vmediamappings
}
func flattenListVnicEthIfRelationship(p *[]models.VnicEthIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethifrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.VnicEthIfRelationshipInterface.(*models.MoMoRef)
		vnicethifrelationship := make(map[string]interface{})
		vnicethifrelationship["class_id"] = item.ClassId
		vnicethifrelationship["link"] = item.Link
		vnicethifrelationship["moid"] = item.Moid
		vnicethifrelationship["object_type"] = item.ObjectType
		vnicethifrelationship["selector"] = item.Selector
		vnicethifrelationships = append(vnicethifrelationships, vnicethifrelationship)
	}
	return vnicethifrelationships
}
func flattenListVnicFcIfRelationship(p *[]models.VnicFcIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcifrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.VnicFcIfRelationshipInterface.(*models.MoMoRef)
		vnicfcifrelationship := make(map[string]interface{})
		vnicfcifrelationship["class_id"] = item.ClassId
		vnicfcifrelationship["link"] = item.Link
		vnicfcifrelationship["moid"] = item.Moid
		vnicfcifrelationship["object_type"] = item.ObjectType
		vnicfcifrelationship["selector"] = item.Selector
		vnicfcifrelationships = append(vnicfcifrelationships, vnicfcifrelationship)
	}
	return vnicfcifrelationships
}
func flattenListWorkflowApi(p *[]models.WorkflowApi, d *schema.ResourceData) []map[string]interface{} {
	var workflowapis []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		workflowapi := make(map[string]interface{})
		workflowapi["body"] = item.Body
		workflowapi["class_id"] = item.ClassId
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
			contentgrammar["class_id"] = item.ClassId
			contentgrammar["error_parameters"] = (func(p *[]models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var contentbaseparameters []map[string]interface{}
				if *p == nil {
					return nil
				}
				for _, item := range *p {
					contentbaseparameter := make(map[string]interface{})
					contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
					contentbaseparameter["class_id"] = item.ClassId
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
			contentgrammar["parameters"] = (func(p *[]models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var contentbaseparameters []map[string]interface{}
				if *p == nil {
					return nil
				}
				for _, item := range *p {
					contentbaseparameter := make(map[string]interface{})
					contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
					contentbaseparameter["class_id"] = item.ClassId
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
			contentgrammar["types"] = (func(p *[]models.ContentComplexType, d *schema.ResourceData) []map[string]interface{} {
				var contentcomplextypes []map[string]interface{}
				if *p == nil {
					return nil
				}
				for _, item := range *p {
					contentcomplextype := make(map[string]interface{})
					contentcomplextype["class_id"] = item.ClassId
					contentcomplextype["name"] = item.Name
					contentcomplextype["object_type"] = item.ObjectType
					contentcomplextype["parameters"] = (func(p *[]models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
						var contentbaseparameters []map[string]interface{}
						if *p == nil {
							return nil
						}
						for _, item := range *p {
							contentbaseparameter := make(map[string]interface{})
							contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
							contentbaseparameter["class_id"] = item.ClassId
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
func flattenListWorkflowBaseDataType(p *[]models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
	var workflowbasedatatypes []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		workflowbasedatatype := make(map[string]interface{})
		workflowbasedatatype["class_id"] = item.ClassId
		workflowbasedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
			var workflowdefaultvalues []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			workflowdefaultvalue := make(map[string]interface{})
			workflowdefaultvalue["class_id"] = item.ClassId
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
func flattenListWorkflowDynamicWorkflowActionTaskList(p *[]models.WorkflowDynamicWorkflowActionTaskList, d *schema.ResourceData) []map[string]interface{} {
	var workflowdynamicworkflowactiontasklists []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		workflowdynamicworkflowactiontasklist := make(map[string]interface{})
		workflowdynamicworkflowactiontasklist["action"] = item.Action
		workflowdynamicworkflowactiontasklist["class_id"] = item.ClassId
		workflowdynamicworkflowactiontasklist["object_type"] = item.ObjectType
		workflowdynamicworkflowactiontasklist["tasks"] = item.Tasks
		workflowdynamicworkflowactiontasklists = append(workflowdynamicworkflowactiontasklists, workflowdynamicworkflowactiontasklist)
	}
	return workflowdynamicworkflowactiontasklists
}
func flattenListWorkflowMessage(p *[]models.WorkflowMessage, d *schema.ResourceData) []map[string]interface{} {
	var workflowmessages []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		workflowmessage := make(map[string]interface{})
		workflowmessage["class_id"] = item.ClassId
		workflowmessage["message"] = item.Message
		workflowmessage["object_type"] = item.ObjectType
		workflowmessage["severity"] = item.Severity
		workflowmessages = append(workflowmessages, workflowmessage)
	}
	return workflowmessages
}
func flattenListWorkflowTaskDefinitionRelationship(p *[]models.WorkflowTaskDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.WorkflowTaskDefinitionRelationshipInterface.(*models.MoMoRef)
		workflowtaskdefinitionrelationship := make(map[string]interface{})
		workflowtaskdefinitionrelationship["class_id"] = item.ClassId
		workflowtaskdefinitionrelationship["link"] = item.Link
		workflowtaskdefinitionrelationship["moid"] = item.Moid
		workflowtaskdefinitionrelationship["object_type"] = item.ObjectType
		workflowtaskdefinitionrelationship["selector"] = item.Selector
		workflowtaskdefinitionrelationships = append(workflowtaskdefinitionrelationships, workflowtaskdefinitionrelationship)
	}
	return workflowtaskdefinitionrelationships
}
func flattenListWorkflowTaskInfoRelationship(p *[]models.WorkflowTaskInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.WorkflowTaskInfoRelationshipInterface.(*models.MoMoRef)
		workflowtaskinforelationship := make(map[string]interface{})
		workflowtaskinforelationship["class_id"] = item.ClassId
		workflowtaskinforelationship["link"] = item.Link
		workflowtaskinforelationship["moid"] = item.Moid
		workflowtaskinforelationship["object_type"] = item.ObjectType
		workflowtaskinforelationship["selector"] = item.Selector
		workflowtaskinforelationships = append(workflowtaskinforelationships, workflowtaskinforelationship)
	}
	return workflowtaskinforelationships
}
func flattenListWorkflowTaskRetryInfo(p *[]models.WorkflowTaskRetryInfo, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskretryinfos []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		workflowtaskretryinfo := make(map[string]interface{})
		workflowtaskretryinfo["class_id"] = item.ClassId
		workflowtaskretryinfo["object_type"] = item.ObjectType
		workflowtaskretryinfo["status"] = item.Status
		workflowtaskretryinfo["task_inst_id"] = item.TaskInstId
		workflowtaskretryinfos = append(workflowtaskretryinfos, workflowtaskretryinfo)
	}
	return workflowtaskretryinfos
}
func flattenListWorkflowWorkflowInfoRelationship(p *[]models.WorkflowWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforelationships []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		item := item.WorkflowWorkflowInfoRelationshipInterface.(*models.MoMoRef)
		workflowworkflowinforelationship := make(map[string]interface{})
		workflowworkflowinforelationship["class_id"] = item.ClassId
		workflowworkflowinforelationship["link"] = item.Link
		workflowworkflowinforelationship["moid"] = item.Moid
		workflowworkflowinforelationship["object_type"] = item.ObjectType
		workflowworkflowinforelationship["selector"] = item.Selector
		workflowworkflowinforelationships = append(workflowworkflowinforelationships, workflowworkflowinforelationship)
	}
	return workflowworkflowinforelationships
}
func flattenListWorkflowWorkflowTask(p *[]models.WorkflowWorkflowTask, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowtasks []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		workflowworkflowtask := make(map[string]interface{})
		workflowworkflowtask["class_id"] = item.ClassId
		workflowworkflowtask["description"] = item.Description
		workflowworkflowtask["label"] = item.Label
		workflowworkflowtask["name"] = item.Name
		workflowworkflowtask["object_type"] = item.ObjectType
		workflowworkflowtasks = append(workflowworkflowtasks, workflowworkflowtask)
	}
	return workflowworkflowtasks
}
func flattenListX509Certificate(p *[]models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var x509certificates []map[string]interface{}
	if *p == nil {
		return nil
	}
	for _, item := range *p {
		x509certificate := make(map[string]interface{})
		x509certificate["class_id"] = item.ClassId
		x509certificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["class_id"] = item.ClassId
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
			pkixdistinguishedname["class_id"] = item.ClassId
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
func flattenMapAdapterUnitRelationship(p *models.AdapterUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.AdapterUnitRelationshipInterface.(*models.MoMoRef)
	adapterunitrelationship := make(map[string]interface{})
	adapterunitrelationship["class_id"] = item.ClassId
	adapterunitrelationship["link"] = item.Link
	adapterunitrelationship["moid"] = item.Moid
	adapterunitrelationship["object_type"] = item.ObjectType
	adapterunitrelationship["selector"] = item.Selector

	adapterunitrelationships = append(adapterunitrelationships, adapterunitrelationship)
	return adapterunitrelationships
}
func flattenMapApplianceDataExportPolicyRelationship(p *models.ApplianceDataExportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ApplianceDataExportPolicyRelationshipInterface.(*models.MoMoRef)
	appliancedataexportpolicyrelationship := make(map[string]interface{})
	appliancedataexportpolicyrelationship["class_id"] = item.ClassId
	appliancedataexportpolicyrelationship["link"] = item.Link
	appliancedataexportpolicyrelationship["moid"] = item.Moid
	appliancedataexportpolicyrelationship["object_type"] = item.ObjectType
	appliancedataexportpolicyrelationship["selector"] = item.Selector

	appliancedataexportpolicyrelationships = append(appliancedataexportpolicyrelationships, appliancedataexportpolicyrelationship)
	return appliancedataexportpolicyrelationships
}
func flattenMapApplianceImageBundleRelationship(p *models.ApplianceImageBundleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var applianceimagebundlerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ApplianceImageBundleRelationshipInterface.(*models.MoMoRef)
	applianceimagebundlerelationship := make(map[string]interface{})
	applianceimagebundlerelationship["class_id"] = item.ClassId
	applianceimagebundlerelationship["link"] = item.Link
	applianceimagebundlerelationship["moid"] = item.Moid
	applianceimagebundlerelationship["object_type"] = item.ObjectType
	applianceimagebundlerelationship["selector"] = item.Selector

	applianceimagebundlerelationships = append(applianceimagebundlerelationships, applianceimagebundlerelationship)
	return applianceimagebundlerelationships
}
func flattenMapAssetClusterMemberRelationship(p *models.AssetClusterMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.AssetClusterMemberRelationshipInterface.(*models.MoMoRef)
	assetclustermemberrelationship := make(map[string]interface{})
	assetclustermemberrelationship["class_id"] = item.ClassId
	assetclustermemberrelationship["link"] = item.Link
	assetclustermemberrelationship["moid"] = item.Moid
	assetclustermemberrelationship["object_type"] = item.ObjectType
	assetclustermemberrelationship["selector"] = item.Selector

	assetclustermemberrelationships = append(assetclustermemberrelationships, assetclustermemberrelationship)
	return assetclustermemberrelationships
}
func flattenMapAssetContractInformation(p *models.AssetContractInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetcontractinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetcontractinformation := make(map[string]interface{})
	assetcontractinformation["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
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
		assetglobalultimate["class_id"] = item.ClassId
		assetglobalultimate["id"] = item.Id
		assetglobalultimate["name"] = item.Name
		assetglobalultimate["object_type"] = item.ObjectType

		assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
		return assetglobalultimates
	})(item.BillToGlobalUltimate, d)
	assetcontractinformation["class_id"] = item.ClassId
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
	assetcustomerinformation["address"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.Address, d)
	assetcustomerinformation["class_id"] = item.ClassId
	assetcustomerinformation["id"] = item.Id
	assetcustomerinformation["name"] = item.Name
	assetcustomerinformation["object_type"] = item.ObjectType

	assetcustomerinformations = append(assetcustomerinformations, assetcustomerinformation)
	return assetcustomerinformations
}
func flattenMapAssetDeviceClaimRelationship(p *models.AssetDeviceClaimRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceclaimrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.AssetDeviceClaimRelationshipInterface.(*models.MoMoRef)
	assetdeviceclaimrelationship := make(map[string]interface{})
	assetdeviceclaimrelationship["class_id"] = item.ClassId
	assetdeviceclaimrelationship["link"] = item.Link
	assetdeviceclaimrelationship["moid"] = item.Moid
	assetdeviceclaimrelationship["object_type"] = item.ObjectType
	assetdeviceclaimrelationship["selector"] = item.Selector

	assetdeviceclaimrelationships = append(assetdeviceclaimrelationships, assetdeviceclaimrelationship)
	return assetdeviceclaimrelationships
}
func flattenMapAssetDeviceConfigurationRelationship(p *models.AssetDeviceConfigurationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconfigurationrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.AssetDeviceConfigurationRelationshipInterface.(*models.MoMoRef)
	assetdeviceconfigurationrelationship := make(map[string]interface{})
	assetdeviceconfigurationrelationship["class_id"] = item.ClassId
	assetdeviceconfigurationrelationship["link"] = item.Link
	assetdeviceconfigurationrelationship["moid"] = item.Moid
	assetdeviceconfigurationrelationship["object_type"] = item.ObjectType
	assetdeviceconfigurationrelationship["selector"] = item.Selector

	assetdeviceconfigurationrelationships = append(assetdeviceconfigurationrelationships, assetdeviceconfigurationrelationship)
	return assetdeviceconfigurationrelationships
}
func flattenMapAssetDeviceConnectionRelationship(p *models.AssetDeviceConnectionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconnectionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.AssetDeviceConnectionRelationshipInterface.(*models.MoMoRef)
	assetdeviceconnectionrelationship := make(map[string]interface{})
	assetdeviceconnectionrelationship["class_id"] = item.ClassId
	assetdeviceconnectionrelationship["link"] = item.Link
	assetdeviceconnectionrelationship["moid"] = item.Moid
	assetdeviceconnectionrelationship["object_type"] = item.ObjectType
	assetdeviceconnectionrelationship["selector"] = item.Selector

	assetdeviceconnectionrelationships = append(assetdeviceconnectionrelationships, assetdeviceconnectionrelationship)
	return assetdeviceconnectionrelationships
}
func flattenMapAssetDeviceRegistrationRelationship(p *models.AssetDeviceRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceregistrationrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.AssetDeviceRegistrationRelationshipInterface.(*models.MoMoRef)
	assetdeviceregistrationrelationship := make(map[string]interface{})
	assetdeviceregistrationrelationship["class_id"] = item.ClassId
	assetdeviceregistrationrelationship["link"] = item.Link
	assetdeviceregistrationrelationship["moid"] = item.Moid
	assetdeviceregistrationrelationship["object_type"] = item.ObjectType
	assetdeviceregistrationrelationship["selector"] = item.Selector

	assetdeviceregistrationrelationships = append(assetdeviceregistrationrelationships, assetdeviceregistrationrelationship)
	return assetdeviceregistrationrelationships
}
func flattenMapAssetGlobalUltimate(p *models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {
	var assetglobalultimates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetglobalultimate := make(map[string]interface{})
	assetglobalultimate["class_id"] = item.ClassId
	assetglobalultimate["id"] = item.Id
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
	assetmanageddevicestatus["class_id"] = item.ClassId
	assetmanageddevicestatus["cloud_port"] = item.CloudPort
	assetmanageddevicestatus["connection_failure_reason"] = item.ConnectionFailureReason
	assetmanageddevicestatus["connection_status"] = item.ConnectionStatus
	assetmanageddevicestatus["error_code"] = item.ErrorCode
	assetmanageddevicestatus["error_reason"] = item.ErrorReason
	assetmanageddevicestatus["object_type"] = item.ObjectType
	assetmanageddevicestatus["process_id"] = item.ProcessId
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
	assetparentconnectionsignature["class_id"] = item.ClassId
	assetparentconnectionsignature["device_id"] = item.DeviceId
	assetparentconnectionsignature["node_id"] = item.NodeId
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
	assetproductinformation["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.BillTo, d)
	assetproductinformation["class_id"] = item.ClassId
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
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
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
func flattenMapAssetSudiInfo(p *models.AssetSudiInfo, d *schema.ResourceData) []map[string]interface{} {
	var assetsudiinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assetsudiinfo := make(map[string]interface{})
	assetsudiinfo["class_id"] = item.ClassId
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
		x509certificate["class_id"] = item.ClassId
		x509certificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["class_id"] = item.ClassId
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
			pkixdistinguishedname["class_id"] = item.ClassId
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
func flattenMapBiosBootModeRelationship(p *models.BiosBootModeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosbootmoderelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.BiosBootModeRelationshipInterface.(*models.MoMoRef)
	biosbootmoderelationship := make(map[string]interface{})
	biosbootmoderelationship["class_id"] = item.ClassId
	biosbootmoderelationship["link"] = item.Link
	biosbootmoderelationship["moid"] = item.Moid
	biosbootmoderelationship["object_type"] = item.ObjectType
	biosbootmoderelationship["selector"] = item.Selector

	biosbootmoderelationships = append(biosbootmoderelationships, biosbootmoderelationship)
	return biosbootmoderelationships
}
func flattenMapBiosUnitRelationship(p *models.BiosUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.BiosUnitRelationshipInterface.(*models.MoMoRef)
	biosunitrelationship := make(map[string]interface{})
	biosunitrelationship["class_id"] = item.ClassId
	biosunitrelationship["link"] = item.Link
	biosunitrelationship["moid"] = item.Moid
	biosunitrelationship["object_type"] = item.ObjectType
	biosunitrelationship["selector"] = item.Selector

	biosunitrelationships = append(biosunitrelationships, biosunitrelationship)
	return biosunitrelationships
}
func flattenMapBootDeviceBootModeRelationship(p *models.BootDeviceBootModeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebootmoderelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.BootDeviceBootModeRelationshipInterface.(*models.MoMoRef)
	bootdevicebootmoderelationship := make(map[string]interface{})
	bootdevicebootmoderelationship["class_id"] = item.ClassId
	bootdevicebootmoderelationship["link"] = item.Link
	bootdevicebootmoderelationship["moid"] = item.Moid
	bootdevicebootmoderelationship["object_type"] = item.ObjectType
	bootdevicebootmoderelationship["selector"] = item.Selector

	bootdevicebootmoderelationships = append(bootdevicebootmoderelationships, bootdevicebootmoderelationship)
	return bootdevicebootmoderelationships
}
func flattenMapCommCredential(p *models.CommCredential, d *schema.ResourceData) []map[string]interface{} {
	var commcredentials []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	commcredential := make(map[string]interface{})
	commcredential["class_id"] = item.ClassId
	commcredential["is_password_set"] = item.IsPasswordSet
	commcredential["object_type"] = item.ObjectType
	password_x := d.Get("credential").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	commcredential["password"] = password_y["password"]
	commcredential["username"] = item.Username

	commcredentials = append(commcredentials, commcredential)
	return commcredentials
}
func flattenMapCommIpV4Interface(p *models.CommIpV4Interface, d *schema.ResourceData) []map[string]interface{} {
	var commipv4interfaces []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	commipv4interface := make(map[string]interface{})
	commipv4interface["class_id"] = item.ClassId
	commipv4interface["gateway"] = item.Gateway
	commipv4interface["ip_address"] = item.IpAddress
	commipv4interface["netmask"] = item.Netmask
	commipv4interface["object_type"] = item.ObjectType

	commipv4interfaces = append(commipv4interfaces, commipv4interface)
	return commipv4interfaces
}
func flattenMapComputeBladeRelationship(p *models.ComputeBladeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computebladerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ComputeBladeRelationshipInterface.(*models.MoMoRef)
	computebladerelationship := make(map[string]interface{})
	computebladerelationship["class_id"] = item.ClassId
	computebladerelationship["link"] = item.Link
	computebladerelationship["moid"] = item.Moid
	computebladerelationship["object_type"] = item.ObjectType
	computebladerelationship["selector"] = item.Selector

	computebladerelationships = append(computebladerelationships, computebladerelationship)
	return computebladerelationships
}
func flattenMapComputeBoardRelationship(p *models.ComputeBoardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computeboardrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ComputeBoardRelationshipInterface.(*models.MoMoRef)
	computeboardrelationship := make(map[string]interface{})
	computeboardrelationship["class_id"] = item.ClassId
	computeboardrelationship["link"] = item.Link
	computeboardrelationship["moid"] = item.Moid
	computeboardrelationship["object_type"] = item.ObjectType
	computeboardrelationship["selector"] = item.Selector

	computeboardrelationships = append(computeboardrelationships, computeboardrelationship)
	return computeboardrelationships
}
func flattenMapComputePersistentMemoryOperation(p *models.ComputePersistentMemoryOperation, d *schema.ResourceData) []map[string]interface{} {
	var computepersistentmemoryoperations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computepersistentmemoryoperation := make(map[string]interface{})
	computepersistentmemoryoperation["admin_action"] = item.AdminAction
	computepersistentmemoryoperation["class_id"] = item.ClassId
	computepersistentmemoryoperation["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	computepersistentmemoryoperation["modules"] = (func(p *[]models.ComputePersistentMemoryModule, d *schema.ResourceData) []map[string]interface{} {
		var computepersistentmemorymodules []map[string]interface{}
		if *p == nil {
			return nil
		}
		for _, item := range *p {
			computepersistentmemorymodule := make(map[string]interface{})
			computepersistentmemorymodule["class_id"] = item.ClassId
			computepersistentmemorymodule["object_type"] = item.ObjectType
			computepersistentmemorymodule["socket_id"] = item.SocketId
			computepersistentmemorymodule["socket_memory_id"] = item.SocketMemoryId
			computepersistentmemorymodules = append(computepersistentmemorymodules, computepersistentmemorymodule)
		}
		return computepersistentmemorymodules
	})(item.Modules, d)
	computepersistentmemoryoperation["object_type"] = item.ObjectType
	secure_passphrase_x := d.Get("persistent_memory_operation").([]interface{})
	secure_passphrase_y := secure_passphrase_x[0].(map[string]interface{})
	computepersistentmemoryoperation["secure_passphrase"] = secure_passphrase_y["secure_passphrase"]

	computepersistentmemoryoperations = append(computepersistentmemoryoperations, computepersistentmemoryoperation)
	return computepersistentmemoryoperations
}
func flattenMapComputePhysicalRelationship(p *models.ComputePhysicalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computephysicalrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ComputePhysicalRelationshipInterface.(*models.MoMoRef)
	computephysicalrelationship := make(map[string]interface{})
	computephysicalrelationship["class_id"] = item.ClassId
	computephysicalrelationship["link"] = item.Link
	computephysicalrelationship["moid"] = item.Moid
	computephysicalrelationship["object_type"] = item.ObjectType
	computephysicalrelationship["selector"] = item.Selector

	computephysicalrelationships = append(computephysicalrelationships, computephysicalrelationship)
	return computephysicalrelationships
}
func flattenMapComputeRackUnitRelationship(p *models.ComputeRackUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ComputeRackUnitRelationshipInterface.(*models.MoMoRef)
	computerackunitrelationship := make(map[string]interface{})
	computerackunitrelationship["class_id"] = item.ClassId
	computerackunitrelationship["link"] = item.Link
	computerackunitrelationship["moid"] = item.Moid
	computerackunitrelationship["object_type"] = item.ObjectType
	computerackunitrelationship["selector"] = item.Selector

	computerackunitrelationships = append(computerackunitrelationships, computerackunitrelationship)
	return computerackunitrelationships
}
func flattenMapComputeServerConfig(p *models.ComputeServerConfig, d *schema.ResourceData) []map[string]interface{} {
	var computeserverconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computeserverconfig := make(map[string]interface{})
	computeserverconfig["asset_tag"] = item.AssetTag
	computeserverconfig["class_id"] = item.ClassId
	computeserverconfig["object_type"] = item.ObjectType
	computeserverconfig["user_label"] = item.UserLabel

	computeserverconfigs = append(computeserverconfigs, computeserverconfig)
	return computeserverconfigs
}
func flattenMapCondHclStatusRelationship(p *models.CondHclStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.CondHclStatusRelationshipInterface.(*models.MoMoRef)
	condhclstatusrelationship := make(map[string]interface{})
	condhclstatusrelationship["class_id"] = item.ClassId
	condhclstatusrelationship["link"] = item.Link
	condhclstatusrelationship["moid"] = item.Moid
	condhclstatusrelationship["object_type"] = item.ObjectType
	condhclstatusrelationship["selector"] = item.Selector

	condhclstatusrelationships = append(condhclstatusrelationships, condhclstatusrelationship)
	return condhclstatusrelationships
}
func flattenMapEquipmentChassisRelationship(p *models.EquipmentChassisRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentchassisrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentChassisRelationshipInterface.(*models.MoMoRef)
	equipmentchassisrelationship := make(map[string]interface{})
	equipmentchassisrelationship["class_id"] = item.ClassId
	equipmentchassisrelationship["link"] = item.Link
	equipmentchassisrelationship["moid"] = item.Moid
	equipmentchassisrelationship["object_type"] = item.ObjectType
	equipmentchassisrelationship["selector"] = item.Selector

	equipmentchassisrelationships = append(equipmentchassisrelationships, equipmentchassisrelationship)
	return equipmentchassisrelationships
}
func flattenMapEquipmentFanModuleRelationship(p *models.EquipmentFanModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentFanModuleRelationshipInterface.(*models.MoMoRef)
	equipmentfanmodulerelationship := make(map[string]interface{})
	equipmentfanmodulerelationship["class_id"] = item.ClassId
	equipmentfanmodulerelationship["link"] = item.Link
	equipmentfanmodulerelationship["moid"] = item.Moid
	equipmentfanmodulerelationship["object_type"] = item.ObjectType
	equipmentfanmodulerelationship["selector"] = item.Selector

	equipmentfanmodulerelationships = append(equipmentfanmodulerelationships, equipmentfanmodulerelationship)
	return equipmentfanmodulerelationships
}
func flattenMapEquipmentLocatorLedRelationship(p *models.EquipmentLocatorLedRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentlocatorledrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentLocatorLedRelationshipInterface.(*models.MoMoRef)
	equipmentlocatorledrelationship := make(map[string]interface{})
	equipmentlocatorledrelationship["class_id"] = item.ClassId
	equipmentlocatorledrelationship["link"] = item.Link
	equipmentlocatorledrelationship["moid"] = item.Moid
	equipmentlocatorledrelationship["object_type"] = item.ObjectType
	equipmentlocatorledrelationship["selector"] = item.Selector

	equipmentlocatorledrelationships = append(equipmentlocatorledrelationships, equipmentlocatorledrelationship)
	return equipmentlocatorledrelationships
}
func flattenMapEquipmentRackEnclosureRelationship(p *models.EquipmentRackEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosurerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentRackEnclosureRelationshipInterface.(*models.MoMoRef)
	equipmentrackenclosurerelationship := make(map[string]interface{})
	equipmentrackenclosurerelationship["class_id"] = item.ClassId
	equipmentrackenclosurerelationship["link"] = item.Link
	equipmentrackenclosurerelationship["moid"] = item.Moid
	equipmentrackenclosurerelationship["object_type"] = item.ObjectType
	equipmentrackenclosurerelationship["selector"] = item.Selector

	equipmentrackenclosurerelationships = append(equipmentrackenclosurerelationships, equipmentrackenclosurerelationship)
	return equipmentrackenclosurerelationships
}
func flattenMapEquipmentRackEnclosureSlotRelationship(p *models.EquipmentRackEnclosureSlotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentRackEnclosureSlotRelationshipInterface.(*models.MoMoRef)
	equipmentrackenclosureslotrelationship := make(map[string]interface{})
	equipmentrackenclosureslotrelationship["class_id"] = item.ClassId
	equipmentrackenclosureslotrelationship["link"] = item.Link
	equipmentrackenclosureslotrelationship["moid"] = item.Moid
	equipmentrackenclosureslotrelationship["object_type"] = item.ObjectType
	equipmentrackenclosureslotrelationship["selector"] = item.Selector

	equipmentrackenclosureslotrelationships = append(equipmentrackenclosureslotrelationships, equipmentrackenclosureslotrelationship)
	return equipmentrackenclosureslotrelationships
}
func flattenMapEquipmentSharedIoModuleRelationship(p *models.EquipmentSharedIoModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsharediomodulerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentSharedIoModuleRelationshipInterface.(*models.MoMoRef)
	equipmentsharediomodulerelationship := make(map[string]interface{})
	equipmentsharediomodulerelationship["class_id"] = item.ClassId
	equipmentsharediomodulerelationship["link"] = item.Link
	equipmentsharediomodulerelationship["moid"] = item.Moid
	equipmentsharediomodulerelationship["object_type"] = item.ObjectType
	equipmentsharediomodulerelationship["selector"] = item.Selector

	equipmentsharediomodulerelationships = append(equipmentsharediomodulerelationships, equipmentsharediomodulerelationship)
	return equipmentsharediomodulerelationships
}
func flattenMapEquipmentSwitchCardRelationship(p *models.EquipmentSwitchCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentSwitchCardRelationshipInterface.(*models.MoMoRef)
	equipmentswitchcardrelationship := make(map[string]interface{})
	equipmentswitchcardrelationship["class_id"] = item.ClassId
	equipmentswitchcardrelationship["link"] = item.Link
	equipmentswitchcardrelationship["moid"] = item.Moid
	equipmentswitchcardrelationship["object_type"] = item.ObjectType
	equipmentswitchcardrelationship["selector"] = item.Selector

	equipmentswitchcardrelationships = append(equipmentswitchcardrelationships, equipmentswitchcardrelationship)
	return equipmentswitchcardrelationships
}
func flattenMapEquipmentSystemIoControllerRelationship(p *models.EquipmentSystemIoControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.EquipmentSystemIoControllerRelationshipInterface.(*models.MoMoRef)
	equipmentsystemiocontrollerrelationship := make(map[string]interface{})
	equipmentsystemiocontrollerrelationship["class_id"] = item.ClassId
	equipmentsystemiocontrollerrelationship["link"] = item.Link
	equipmentsystemiocontrollerrelationship["moid"] = item.Moid
	equipmentsystemiocontrollerrelationship["object_type"] = item.ObjectType
	equipmentsystemiocontrollerrelationship["selector"] = item.Selector

	equipmentsystemiocontrollerrelationships = append(equipmentsystemiocontrollerrelationships, equipmentsystemiocontrollerrelationship)
	return equipmentsystemiocontrollerrelationships
}
func flattenMapFirmwareDirectDownload(p *models.FirmwareDirectDownload, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredirectdownloads []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwaredirectdownload := make(map[string]interface{})
	firmwaredirectdownload["class_id"] = item.ClassId
	http_server_x := d.Get("direct_download").([]interface{})
	http_server_y := http_server_x[0].(map[string]interface{})
	firmwaredirectdownload["http_server"] = http_server_y["http_server"]
	firmwaredirectdownload["image_source"] = item.ImageSource
	firmwaredirectdownload["is_password_set"] = item.IsPasswordSet
	firmwaredirectdownload["object_type"] = item.ObjectType
	password_x := d.Get("direct_download").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	firmwaredirectdownload["password"] = password_y["password"]
	firmwaredirectdownload["upgradeoption"] = item.Upgradeoption
	username_x := d.Get("direct_download").([]interface{})
	username_y := username_x[0].(map[string]interface{})
	firmwaredirectdownload["username"] = username_y["username"]

	firmwaredirectdownloads = append(firmwaredirectdownloads, firmwaredirectdownload)
	return firmwaredirectdownloads
}
func flattenMapFirmwareDistributableRelationship(p *models.FirmwareDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredistributablerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.FirmwareDistributableRelationshipInterface.(*models.MoMoRef)
	firmwaredistributablerelationship := make(map[string]interface{})
	firmwaredistributablerelationship["class_id"] = item.ClassId
	firmwaredistributablerelationship["link"] = item.Link
	firmwaredistributablerelationship["moid"] = item.Moid
	firmwaredistributablerelationship["object_type"] = item.ObjectType
	firmwaredistributablerelationship["selector"] = item.Selector

	firmwaredistributablerelationships = append(firmwaredistributablerelationships, firmwaredistributablerelationship)
	return firmwaredistributablerelationships
}
func flattenMapFirmwareNetworkShare(p *models.FirmwareNetworkShare, d *schema.ResourceData) []map[string]interface{} {
	var firmwarenetworkshares []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	firmwarenetworkshare := make(map[string]interface{})
	cifs_server_x := d.Get("network_share").([]interface{})
	cifs_server_y := cifs_server_x[0].(map[string]interface{})
	firmwarenetworkshare["cifs_server"] = cifs_server_y["cifs_server"]
	firmwarenetworkshare["class_id"] = item.ClassId
	http_server_x := d.Get("network_share").([]interface{})
	http_server_y := http_server_x[0].(map[string]interface{})
	firmwarenetworkshare["http_server"] = http_server_y["http_server"]
	firmwarenetworkshare["is_password_set"] = item.IsPasswordSet
	firmwarenetworkshare["map_type"] = item.MapType
	nfs_server_x := d.Get("network_share").([]interface{})
	nfs_server_y := nfs_server_x[0].(map[string]interface{})
	firmwarenetworkshare["nfs_server"] = nfs_server_y["nfs_server"]
	firmwarenetworkshare["object_type"] = item.ObjectType
	password_x := d.Get("network_share").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	firmwarenetworkshare["password"] = password_y["password"]
	firmwarenetworkshare["upgradeoption"] = item.Upgradeoption
	username_x := d.Get("network_share").([]interface{})
	username_y := username_x[0].(map[string]interface{})
	firmwarenetworkshare["username"] = username_y["username"]

	firmwarenetworkshares = append(firmwarenetworkshares, firmwarenetworkshare)
	return firmwarenetworkshares
}
func flattenMapFirmwareRunningFirmwareRelationship(p *models.FirmwareRunningFirmwareRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.FirmwareRunningFirmwareRelationshipInterface.(*models.MoMoRef)
	firmwarerunningfirmwarerelationship := make(map[string]interface{})
	firmwarerunningfirmwarerelationship["class_id"] = item.ClassId
	firmwarerunningfirmwarerelationship["link"] = item.Link
	firmwarerunningfirmwarerelationship["moid"] = item.Moid
	firmwarerunningfirmwarerelationship["object_type"] = item.ObjectType
	firmwarerunningfirmwarerelationship["selector"] = item.Selector

	firmwarerunningfirmwarerelationships = append(firmwarerunningfirmwarerelationships, firmwarerunningfirmwarerelationship)
	return firmwarerunningfirmwarerelationships
}
func flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(p *models.FirmwareServerConfigurationUtilityDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareserverconfigurationutilitydistributablerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.FirmwareServerConfigurationUtilityDistributableRelationshipInterface.(*models.MoMoRef)
	firmwareserverconfigurationutilitydistributablerelationship := make(map[string]interface{})
	firmwareserverconfigurationutilitydistributablerelationship["class_id"] = item.ClassId
	firmwareserverconfigurationutilitydistributablerelationship["link"] = item.Link
	firmwareserverconfigurationutilitydistributablerelationship["moid"] = item.Moid
	firmwareserverconfigurationutilitydistributablerelationship["object_type"] = item.ObjectType
	firmwareserverconfigurationutilitydistributablerelationship["selector"] = item.Selector

	firmwareserverconfigurationutilitydistributablerelationships = append(firmwareserverconfigurationutilitydistributablerelationships, firmwareserverconfigurationutilitydistributablerelationship)
	return firmwareserverconfigurationutilitydistributablerelationships
}
func flattenMapFirmwareUpgradeRelationship(p *models.FirmwareUpgradeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgraderelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.FirmwareUpgradeRelationshipInterface.(*models.MoMoRef)
	firmwareupgraderelationship := make(map[string]interface{})
	firmwareupgraderelationship["class_id"] = item.ClassId
	firmwareupgraderelationship["link"] = item.Link
	firmwareupgraderelationship["moid"] = item.Moid
	firmwareupgraderelationship["object_type"] = item.ObjectType
	firmwareupgraderelationship["selector"] = item.Selector

	firmwareupgraderelationships = append(firmwareupgraderelationships, firmwareupgraderelationship)
	return firmwareupgraderelationships
}
func flattenMapFirmwareUpgradeStatusRelationship(p *models.FirmwareUpgradeStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradestatusrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.FirmwareUpgradeStatusRelationshipInterface.(*models.MoMoRef)
	firmwareupgradestatusrelationship := make(map[string]interface{})
	firmwareupgradestatusrelationship["class_id"] = item.ClassId
	firmwareupgradestatusrelationship["link"] = item.Link
	firmwareupgradestatusrelationship["moid"] = item.Moid
	firmwareupgradestatusrelationship["object_type"] = item.ObjectType
	firmwareupgradestatusrelationship["selector"] = item.Selector

	firmwareupgradestatusrelationships = append(firmwareupgradestatusrelationships, firmwareupgradestatusrelationship)
	return firmwareupgradestatusrelationships
}
func flattenMapForecastCatalogRelationship(p *models.ForecastCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastcatalogrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ForecastCatalogRelationshipInterface.(*models.MoMoRef)
	forecastcatalogrelationship := make(map[string]interface{})
	forecastcatalogrelationship["class_id"] = item.ClassId
	forecastcatalogrelationship["link"] = item.Link
	forecastcatalogrelationship["moid"] = item.Moid
	forecastcatalogrelationship["object_type"] = item.ObjectType
	forecastcatalogrelationship["selector"] = item.Selector

	forecastcatalogrelationships = append(forecastcatalogrelationships, forecastcatalogrelationship)
	return forecastcatalogrelationships
}
func flattenMapForecastDefinitionRelationship(p *models.ForecastDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ForecastDefinitionRelationshipInterface.(*models.MoMoRef)
	forecastdefinitionrelationship := make(map[string]interface{})
	forecastdefinitionrelationship["class_id"] = item.ClassId
	forecastdefinitionrelationship["link"] = item.Link
	forecastdefinitionrelationship["moid"] = item.Moid
	forecastdefinitionrelationship["object_type"] = item.ObjectType
	forecastdefinitionrelationship["selector"] = item.Selector

	forecastdefinitionrelationships = append(forecastdefinitionrelationships, forecastdefinitionrelationship)
	return forecastdefinitionrelationships
}
func flattenMapForecastModel(p *models.ForecastModel, d *schema.ResourceData) []map[string]interface{} {
	var forecastmodels []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	forecastmodel := make(map[string]interface{})
	forecastmodel["accuracy"] = item.Accuracy
	forecastmodel["class_id"] = item.ClassId
	forecastmodel["model_data"] = item.ModelData
	forecastmodel["model_type"] = item.ModelType
	forecastmodel["object_type"] = item.ObjectType

	forecastmodels = append(forecastmodels, forecastmodel)
	return forecastmodels
}
func flattenMapGraphicsCardRelationship(p *models.GraphicsCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.GraphicsCardRelationshipInterface.(*models.MoMoRef)
	graphicscardrelationship := make(map[string]interface{})
	graphicscardrelationship["class_id"] = item.ClassId
	graphicscardrelationship["link"] = item.Link
	graphicscardrelationship["moid"] = item.Moid
	graphicscardrelationship["object_type"] = item.ObjectType
	graphicscardrelationship["selector"] = item.Selector

	graphicscardrelationships = append(graphicscardrelationships, graphicscardrelationship)
	return graphicscardrelationships
}
func flattenMapHclOperatingSystemVendorRelationship(p *models.HclOperatingSystemVendorRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemvendorrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HclOperatingSystemVendorRelationshipInterface.(*models.MoMoRef)
	hcloperatingsystemvendorrelationship := make(map[string]interface{})
	hcloperatingsystemvendorrelationship["class_id"] = item.ClassId
	hcloperatingsystemvendorrelationship["link"] = item.Link
	hcloperatingsystemvendorrelationship["moid"] = item.Moid
	hcloperatingsystemvendorrelationship["object_type"] = item.ObjectType
	hcloperatingsystemvendorrelationship["selector"] = item.Selector

	hcloperatingsystemvendorrelationships = append(hcloperatingsystemvendorrelationships, hcloperatingsystemvendorrelationship)
	return hcloperatingsystemvendorrelationships
}
func flattenMapHyperflexAppCatalogRelationship(p *models.HyperflexAppCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexappcatalogrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexAppCatalogRelationshipInterface.(*models.MoMoRef)
	hyperflexappcatalogrelationship := make(map[string]interface{})
	hyperflexappcatalogrelationship["class_id"] = item.ClassId
	hyperflexappcatalogrelationship["link"] = item.Link
	hyperflexappcatalogrelationship["moid"] = item.Moid
	hyperflexappcatalogrelationship["object_type"] = item.ObjectType
	hyperflexappcatalogrelationship["selector"] = item.Selector

	hyperflexappcatalogrelationships = append(hyperflexappcatalogrelationships, hyperflexappcatalogrelationship)
	return hyperflexappcatalogrelationships
}
func flattenMapHyperflexAutoSupportPolicyRelationship(p *models.HyperflexAutoSupportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexautosupportpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexAutoSupportPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexautosupportpolicyrelationship := make(map[string]interface{})
	hyperflexautosupportpolicyrelationship["class_id"] = item.ClassId
	hyperflexautosupportpolicyrelationship["link"] = item.Link
	hyperflexautosupportpolicyrelationship["moid"] = item.Moid
	hyperflexautosupportpolicyrelationship["object_type"] = item.ObjectType
	hyperflexautosupportpolicyrelationship["selector"] = item.Selector

	hyperflexautosupportpolicyrelationships = append(hyperflexautosupportpolicyrelationships, hyperflexautosupportpolicyrelationship)
	return hyperflexautosupportpolicyrelationships
}
func flattenMapHyperflexClusterRelationship(p *models.HyperflexClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexClusterRelationshipInterface.(*models.MoMoRef)
	hyperflexclusterrelationship := make(map[string]interface{})
	hyperflexclusterrelationship["class_id"] = item.ClassId
	hyperflexclusterrelationship["link"] = item.Link
	hyperflexclusterrelationship["moid"] = item.Moid
	hyperflexclusterrelationship["object_type"] = item.ObjectType
	hyperflexclusterrelationship["selector"] = item.Selector

	hyperflexclusterrelationships = append(hyperflexclusterrelationships, hyperflexclusterrelationship)
	return hyperflexclusterrelationships
}
func flattenMapHyperflexClusterNetworkPolicyRelationship(p *models.HyperflexClusterNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusternetworkpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexClusterNetworkPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexclusternetworkpolicyrelationship := make(map[string]interface{})
	hyperflexclusternetworkpolicyrelationship["class_id"] = item.ClassId
	hyperflexclusternetworkpolicyrelationship["link"] = item.Link
	hyperflexclusternetworkpolicyrelationship["moid"] = item.Moid
	hyperflexclusternetworkpolicyrelationship["object_type"] = item.ObjectType
	hyperflexclusternetworkpolicyrelationship["selector"] = item.Selector

	hyperflexclusternetworkpolicyrelationships = append(hyperflexclusternetworkpolicyrelationships, hyperflexclusternetworkpolicyrelationship)
	return hyperflexclusternetworkpolicyrelationships
}
func flattenMapHyperflexClusterProfileRelationship(p *models.HyperflexClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexClusterProfileRelationshipInterface.(*models.MoMoRef)
	hyperflexclusterprofilerelationship := make(map[string]interface{})
	hyperflexclusterprofilerelationship["class_id"] = item.ClassId
	hyperflexclusterprofilerelationship["link"] = item.Link
	hyperflexclusterprofilerelationship["moid"] = item.Moid
	hyperflexclusterprofilerelationship["object_type"] = item.ObjectType
	hyperflexclusterprofilerelationship["selector"] = item.Selector

	hyperflexclusterprofilerelationships = append(hyperflexclusterprofilerelationships, hyperflexclusterprofilerelationship)
	return hyperflexclusterprofilerelationships
}
func flattenMapHyperflexClusterStoragePolicyRelationship(p *models.HyperflexClusterStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterstoragepolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexClusterStoragePolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexclusterstoragepolicyrelationship := make(map[string]interface{})
	hyperflexclusterstoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexclusterstoragepolicyrelationship["link"] = item.Link
	hyperflexclusterstoragepolicyrelationship["moid"] = item.Moid
	hyperflexclusterstoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexclusterstoragepolicyrelationship["selector"] = item.Selector

	hyperflexclusterstoragepolicyrelationships = append(hyperflexclusterstoragepolicyrelationships, hyperflexclusterstoragepolicyrelationship)
	return hyperflexclusterstoragepolicyrelationships
}
func flattenMapHyperflexConfigResultRelationship(p *models.HyperflexConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexConfigResultRelationshipInterface.(*models.MoMoRef)
	hyperflexconfigresultrelationship := make(map[string]interface{})
	hyperflexconfigresultrelationship["class_id"] = item.ClassId
	hyperflexconfigresultrelationship["link"] = item.Link
	hyperflexconfigresultrelationship["moid"] = item.Moid
	hyperflexconfigresultrelationship["object_type"] = item.ObjectType
	hyperflexconfigresultrelationship["selector"] = item.Selector

	hyperflexconfigresultrelationships = append(hyperflexconfigresultrelationships, hyperflexconfigresultrelationship)
	return hyperflexconfigresultrelationships
}
func flattenMapHyperflexExtFcStoragePolicyRelationship(p *models.HyperflexExtFcStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextfcstoragepolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexExtFcStoragePolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexextfcstoragepolicyrelationship := make(map[string]interface{})
	hyperflexextfcstoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexextfcstoragepolicyrelationship["link"] = item.Link
	hyperflexextfcstoragepolicyrelationship["moid"] = item.Moid
	hyperflexextfcstoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexextfcstoragepolicyrelationship["selector"] = item.Selector

	hyperflexextfcstoragepolicyrelationships = append(hyperflexextfcstoragepolicyrelationships, hyperflexextfcstoragepolicyrelationship)
	return hyperflexextfcstoragepolicyrelationships
}
func flattenMapHyperflexExtIscsiStoragePolicyRelationship(p *models.HyperflexExtIscsiStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextiscsistoragepolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexExtIscsiStoragePolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexextiscsistoragepolicyrelationship := make(map[string]interface{})
	hyperflexextiscsistoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexextiscsistoragepolicyrelationship["link"] = item.Link
	hyperflexextiscsistoragepolicyrelationship["moid"] = item.Moid
	hyperflexextiscsistoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexextiscsistoragepolicyrelationship["selector"] = item.Selector

	hyperflexextiscsistoragepolicyrelationships = append(hyperflexextiscsistoragepolicyrelationships, hyperflexextiscsistoragepolicyrelationship)
	return hyperflexextiscsistoragepolicyrelationships
}
func flattenMapHyperflexFeatureLimitExternalRelationship(p *models.HyperflexFeatureLimitExternalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitexternalrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexFeatureLimitExternalRelationshipInterface.(*models.MoMoRef)
	hyperflexfeaturelimitexternalrelationship := make(map[string]interface{})
	hyperflexfeaturelimitexternalrelationship["class_id"] = item.ClassId
	hyperflexfeaturelimitexternalrelationship["link"] = item.Link
	hyperflexfeaturelimitexternalrelationship["moid"] = item.Moid
	hyperflexfeaturelimitexternalrelationship["object_type"] = item.ObjectType
	hyperflexfeaturelimitexternalrelationship["selector"] = item.Selector

	hyperflexfeaturelimitexternalrelationships = append(hyperflexfeaturelimitexternalrelationships, hyperflexfeaturelimitexternalrelationship)
	return hyperflexfeaturelimitexternalrelationships
}
func flattenMapHyperflexFeatureLimitInternalRelationship(p *models.HyperflexFeatureLimitInternalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitinternalrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexFeatureLimitInternalRelationshipInterface.(*models.MoMoRef)
	hyperflexfeaturelimitinternalrelationship := make(map[string]interface{})
	hyperflexfeaturelimitinternalrelationship["class_id"] = item.ClassId
	hyperflexfeaturelimitinternalrelationship["link"] = item.Link
	hyperflexfeaturelimitinternalrelationship["moid"] = item.Moid
	hyperflexfeaturelimitinternalrelationship["object_type"] = item.ObjectType
	hyperflexfeaturelimitinternalrelationship["selector"] = item.Selector

	hyperflexfeaturelimitinternalrelationships = append(hyperflexfeaturelimitinternalrelationships, hyperflexfeaturelimitinternalrelationship)
	return hyperflexfeaturelimitinternalrelationships
}
func flattenMapHyperflexHealthRelationship(p *models.HyperflexHealthRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhealthrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexHealthRelationshipInterface.(*models.MoMoRef)
	hyperflexhealthrelationship := make(map[string]interface{})
	hyperflexhealthrelationship["class_id"] = item.ClassId
	hyperflexhealthrelationship["link"] = item.Link
	hyperflexhealthrelationship["moid"] = item.Moid
	hyperflexhealthrelationship["object_type"] = item.ObjectType
	hyperflexhealthrelationship["selector"] = item.Selector

	hyperflexhealthrelationships = append(hyperflexhealthrelationships, hyperflexhealthrelationship)
	return hyperflexhealthrelationships
}
func flattenMapHyperflexHxNetworkAddressDt(p *models.HyperflexHxNetworkAddressDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxnetworkaddressdts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexhxnetworkaddressdt := make(map[string]interface{})
	hyperflexhxnetworkaddressdt["address"] = item.Address
	hyperflexhxnetworkaddressdt["class_id"] = item.ClassId
	hyperflexhxnetworkaddressdt["fqdn"] = item.Fqdn
	hyperflexhxnetworkaddressdt["ip"] = item.Ip
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
	hyperflexhxresiliencyinfodt["class_id"] = item.ClassId
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
func flattenMapHyperflexHxUuIdDt(p *models.HyperflexHxUuIdDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxuuiddts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexhxuuiddt := make(map[string]interface{})
	hyperflexhxuuiddt["class_id"] = item.ClassId
	hyperflexhxuuiddt["links"] = (func(p *[]models.HyperflexHxLinkDt, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexhxlinkdts []map[string]interface{}
		if *p == nil {
			return nil
		}
		for _, item := range *p {
			hyperflexhxlinkdt := make(map[string]interface{})
			hyperflexhxlinkdt["class_id"] = item.ClassId
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
	hyperflexhxuuiddt["uuid"] = item.Uuid

	hyperflexhxuuiddts = append(hyperflexhxuuiddts, hyperflexhxuuiddt)
	return hyperflexhxuuiddts
}
func flattenMapHyperflexIpAddrRange(p *models.HyperflexIpAddrRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexipaddrranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexipaddrrange := make(map[string]interface{})
	hyperflexipaddrrange["class_id"] = item.ClassId
	hyperflexipaddrrange["end_addr"] = item.EndAddr
	hyperflexipaddrrange["gateway"] = item.Gateway
	hyperflexipaddrrange["netmask"] = item.Netmask
	hyperflexipaddrrange["object_type"] = item.ObjectType
	hyperflexipaddrrange["start_addr"] = item.StartAddr

	hyperflexipaddrranges = append(hyperflexipaddrranges, hyperflexipaddrrange)
	return hyperflexipaddrranges
}
func flattenMapHyperflexLocalCredentialPolicyRelationship(p *models.HyperflexLocalCredentialPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlocalcredentialpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexLocalCredentialPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexlocalcredentialpolicyrelationship := make(map[string]interface{})
	hyperflexlocalcredentialpolicyrelationship["class_id"] = item.ClassId
	hyperflexlocalcredentialpolicyrelationship["link"] = item.Link
	hyperflexlocalcredentialpolicyrelationship["moid"] = item.Moid
	hyperflexlocalcredentialpolicyrelationship["object_type"] = item.ObjectType
	hyperflexlocalcredentialpolicyrelationship["selector"] = item.Selector

	hyperflexlocalcredentialpolicyrelationships = append(hyperflexlocalcredentialpolicyrelationships, hyperflexlocalcredentialpolicyrelationship)
	return hyperflexlocalcredentialpolicyrelationships
}
func flattenMapHyperflexLogicalAvailabilityZone(p *models.HyperflexLogicalAvailabilityZone, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlogicalavailabilityzones []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexlogicalavailabilityzone := make(map[string]interface{})
	hyperflexlogicalavailabilityzone["auto_config"] = item.AutoConfig
	hyperflexlogicalavailabilityzone["class_id"] = item.ClassId
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
	hyperflexmacaddrprefixrange["class_id"] = item.ClassId
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
	hyperflexnamedvlan["class_id"] = item.ClassId
	hyperflexnamedvlan["name"] = item.Name
	hyperflexnamedvlan["object_type"] = item.ObjectType
	hyperflexnamedvlan["vlan_id"] = item.VlanId

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
	hyperflexnamedvsan["class_id"] = item.ClassId
	hyperflexnamedvsan["name"] = item.Name
	hyperflexnamedvsan["object_type"] = item.ObjectType
	hyperflexnamedvsan["vsan_id"] = item.VsanId

	hyperflexnamedvsans = append(hyperflexnamedvsans, hyperflexnamedvsan)
	return hyperflexnamedvsans
}
func flattenMapHyperflexNodeConfigPolicyRelationship(p *models.HyperflexNodeConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeconfigpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexNodeConfigPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexnodeconfigpolicyrelationship := make(map[string]interface{})
	hyperflexnodeconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexnodeconfigpolicyrelationship["link"] = item.Link
	hyperflexnodeconfigpolicyrelationship["moid"] = item.Moid
	hyperflexnodeconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexnodeconfigpolicyrelationship["selector"] = item.Selector

	hyperflexnodeconfigpolicyrelationships = append(hyperflexnodeconfigpolicyrelationships, hyperflexnodeconfigpolicyrelationship)
	return hyperflexnodeconfigpolicyrelationships
}
func flattenMapHyperflexProxySettingPolicyRelationship(p *models.HyperflexProxySettingPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexproxysettingpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexProxySettingPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexproxysettingpolicyrelationship := make(map[string]interface{})
	hyperflexproxysettingpolicyrelationship["class_id"] = item.ClassId
	hyperflexproxysettingpolicyrelationship["link"] = item.Link
	hyperflexproxysettingpolicyrelationship["moid"] = item.Moid
	hyperflexproxysettingpolicyrelationship["object_type"] = item.ObjectType
	hyperflexproxysettingpolicyrelationship["selector"] = item.Selector

	hyperflexproxysettingpolicyrelationships = append(hyperflexproxysettingpolicyrelationships, hyperflexproxysettingpolicyrelationship)
	return hyperflexproxysettingpolicyrelationships
}
func flattenMapHyperflexServerFirmwareVersionRelationship(p *models.HyperflexServerFirmwareVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexServerFirmwareVersionRelationshipInterface.(*models.MoMoRef)
	hyperflexserverfirmwareversionrelationship := make(map[string]interface{})
	hyperflexserverfirmwareversionrelationship["class_id"] = item.ClassId
	hyperflexserverfirmwareversionrelationship["link"] = item.Link
	hyperflexserverfirmwareversionrelationship["moid"] = item.Moid
	hyperflexserverfirmwareversionrelationship["object_type"] = item.ObjectType
	hyperflexserverfirmwareversionrelationship["selector"] = item.Selector

	hyperflexserverfirmwareversionrelationships = append(hyperflexserverfirmwareversionrelationships, hyperflexserverfirmwareversionrelationship)
	return hyperflexserverfirmwareversionrelationships
}
func flattenMapHyperflexServerModelRelationship(p *models.HyperflexServerModelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexServerModelRelationshipInterface.(*models.MoMoRef)
	hyperflexservermodelrelationship := make(map[string]interface{})
	hyperflexservermodelrelationship["class_id"] = item.ClassId
	hyperflexservermodelrelationship["link"] = item.Link
	hyperflexservermodelrelationship["moid"] = item.Moid
	hyperflexservermodelrelationship["object_type"] = item.ObjectType
	hyperflexservermodelrelationship["selector"] = item.Selector

	hyperflexservermodelrelationships = append(hyperflexservermodelrelationships, hyperflexservermodelrelationship)
	return hyperflexservermodelrelationships
}
func flattenMapHyperflexSoftwareVersionPolicyRelationship(p *models.HyperflexSoftwareVersionPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwareversionpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexSoftwareVersionPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexsoftwareversionpolicyrelationship := make(map[string]interface{})
	hyperflexsoftwareversionpolicyrelationship["class_id"] = item.ClassId
	hyperflexsoftwareversionpolicyrelationship["link"] = item.Link
	hyperflexsoftwareversionpolicyrelationship["moid"] = item.Moid
	hyperflexsoftwareversionpolicyrelationship["object_type"] = item.ObjectType
	hyperflexsoftwareversionpolicyrelationship["selector"] = item.Selector

	hyperflexsoftwareversionpolicyrelationships = append(hyperflexsoftwareversionpolicyrelationships, hyperflexsoftwareversionpolicyrelationship)
	return hyperflexsoftwareversionpolicyrelationships
}
func flattenMapHyperflexSummary(p *models.HyperflexSummary, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsummarys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexsummary := make(map[string]interface{})
	hyperflexsummary["active_nodes"] = item.ActiveNodes
	hyperflexsummary["address"] = item.Address
	hyperflexsummary["boottime"] = item.Boottime
	hyperflexsummary["class_id"] = item.ClassId
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
		hyperflexstplatformclusterhealinginfo["class_id"] = item.ClassId
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
		hyperflexstplatformclusterresiliencyinfo["class_id"] = item.ClassId
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
func flattenMapHyperflexSysConfigPolicyRelationship(p *models.HyperflexSysConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsysconfigpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexSysConfigPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexsysconfigpolicyrelationship := make(map[string]interface{})
	hyperflexsysconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexsysconfigpolicyrelationship["link"] = item.Link
	hyperflexsysconfigpolicyrelationship["moid"] = item.Moid
	hyperflexsysconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexsysconfigpolicyrelationship["selector"] = item.Selector

	hyperflexsysconfigpolicyrelationships = append(hyperflexsysconfigpolicyrelationships, hyperflexsysconfigpolicyrelationship)
	return hyperflexsysconfigpolicyrelationships
}
func flattenMapHyperflexUcsmConfigPolicyRelationship(p *models.HyperflexUcsmConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexucsmconfigpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexUcsmConfigPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexucsmconfigpolicyrelationship := make(map[string]interface{})
	hyperflexucsmconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexucsmconfigpolicyrelationship["link"] = item.Link
	hyperflexucsmconfigpolicyrelationship["moid"] = item.Moid
	hyperflexucsmconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexucsmconfigpolicyrelationship["selector"] = item.Selector

	hyperflexucsmconfigpolicyrelationships = append(hyperflexucsmconfigpolicyrelationships, hyperflexucsmconfigpolicyrelationship)
	return hyperflexucsmconfigpolicyrelationships
}
func flattenMapHyperflexVcenterConfigPolicyRelationship(p *models.HyperflexVcenterConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvcenterconfigpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.HyperflexVcenterConfigPolicyRelationshipInterface.(*models.MoMoRef)
	hyperflexvcenterconfigpolicyrelationship := make(map[string]interface{})
	hyperflexvcenterconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexvcenterconfigpolicyrelationship["link"] = item.Link
	hyperflexvcenterconfigpolicyrelationship["moid"] = item.Moid
	hyperflexvcenterconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexvcenterconfigpolicyrelationship["selector"] = item.Selector

	hyperflexvcenterconfigpolicyrelationships = append(hyperflexvcenterconfigpolicyrelationships, hyperflexvcenterconfigpolicyrelationship)
	return hyperflexvcenterconfigpolicyrelationships
}
func flattenMapHyperflexWwxnPrefixRange(p *models.HyperflexWwxnPrefixRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexwwxnprefixranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hyperflexwwxnprefixrange := make(map[string]interface{})
	hyperflexwwxnprefixrange["class_id"] = item.ClassId
	hyperflexwwxnprefixrange["end_addr"] = item.EndAddr
	hyperflexwwxnprefixrange["object_type"] = item.ObjectType
	hyperflexwwxnprefixrange["start_addr"] = item.StartAddr

	hyperflexwwxnprefixranges = append(hyperflexwwxnprefixranges, hyperflexwwxnprefixrange)
	return hyperflexwwxnprefixranges
}
func flattenMapIaasLicenseInfoRelationship(p *models.IaasLicenseInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseinforelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IaasLicenseInfoRelationshipInterface.(*models.MoMoRef)
	iaaslicenseinforelationship := make(map[string]interface{})
	iaaslicenseinforelationship["class_id"] = item.ClassId
	iaaslicenseinforelationship["link"] = item.Link
	iaaslicenseinforelationship["moid"] = item.Moid
	iaaslicenseinforelationship["object_type"] = item.ObjectType
	iaaslicenseinforelationship["selector"] = item.Selector

	iaaslicenseinforelationships = append(iaaslicenseinforelationships, iaaslicenseinforelationship)
	return iaaslicenseinforelationships
}
func flattenMapIaasUcsdInfoRelationship(p *models.IaasUcsdInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdinforelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IaasUcsdInfoRelationshipInterface.(*models.MoMoRef)
	iaasucsdinforelationship := make(map[string]interface{})
	iaasucsdinforelationship["class_id"] = item.ClassId
	iaasucsdinforelationship["link"] = item.Link
	iaasucsdinforelationship["moid"] = item.Moid
	iaasucsdinforelationship["object_type"] = item.ObjectType
	iaasucsdinforelationship["selector"] = item.Selector

	iaasucsdinforelationships = append(iaasucsdinforelationships, iaasucsdinforelationship)
	return iaasucsdinforelationships
}
func flattenMapIaasUcsdManagedInfraRelationship(p *models.IaasUcsdManagedInfraRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdmanagedinfrarelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IaasUcsdManagedInfraRelationshipInterface.(*models.MoMoRef)
	iaasucsdmanagedinfrarelationship := make(map[string]interface{})
	iaasucsdmanagedinfrarelationship["class_id"] = item.ClassId
	iaasucsdmanagedinfrarelationship["link"] = item.Link
	iaasucsdmanagedinfrarelationship["moid"] = item.Moid
	iaasucsdmanagedinfrarelationship["object_type"] = item.ObjectType
	iaasucsdmanagedinfrarelationship["selector"] = item.Selector

	iaasucsdmanagedinfrarelationships = append(iaasucsdmanagedinfrarelationships, iaasucsdmanagedinfrarelationship)
	return iaasucsdmanagedinfrarelationships
}
func flattenMapIamAccountRelationship(p *models.IamAccountRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamAccountRelationshipInterface.(*models.MoMoRef)
	iamaccountrelationship := make(map[string]interface{})
	iamaccountrelationship["class_id"] = item.ClassId
	iamaccountrelationship["link"] = item.Link
	iamaccountrelationship["moid"] = item.Moid
	iamaccountrelationship["object_type"] = item.ObjectType
	iamaccountrelationship["selector"] = item.Selector

	iamaccountrelationships = append(iamaccountrelationships, iamaccountrelationship)
	return iamaccountrelationships
}
func flattenMapIamAppRegistrationRelationship(p *models.IamAppRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamAppRegistrationRelationshipInterface.(*models.MoMoRef)
	iamappregistrationrelationship := make(map[string]interface{})
	iamappregistrationrelationship["class_id"] = item.ClassId
	iamappregistrationrelationship["link"] = item.Link
	iamappregistrationrelationship["moid"] = item.Moid
	iamappregistrationrelationship["object_type"] = item.ObjectType
	iamappregistrationrelationship["selector"] = item.Selector

	iamappregistrationrelationships = append(iamappregistrationrelationships, iamappregistrationrelationship)
	return iamappregistrationrelationships
}
func flattenMapIamCertificateRelationship(p *models.IamCertificateRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamCertificateRelationshipInterface.(*models.MoMoRef)
	iamcertificaterelationship := make(map[string]interface{})
	iamcertificaterelationship["class_id"] = item.ClassId
	iamcertificaterelationship["link"] = item.Link
	iamcertificaterelationship["moid"] = item.Moid
	iamcertificaterelationship["object_type"] = item.ObjectType
	iamcertificaterelationship["selector"] = item.Selector

	iamcertificaterelationships = append(iamcertificaterelationships, iamcertificaterelationship)
	return iamcertificaterelationships
}
func flattenMapIamCertificateRequestRelationship(p *models.IamCertificateRequestRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterequestrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamCertificateRequestRelationshipInterface.(*models.MoMoRef)
	iamcertificaterequestrelationship := make(map[string]interface{})
	iamcertificaterequestrelationship["class_id"] = item.ClassId
	iamcertificaterequestrelationship["link"] = item.Link
	iamcertificaterequestrelationship["moid"] = item.Moid
	iamcertificaterequestrelationship["object_type"] = item.ObjectType
	iamcertificaterequestrelationship["selector"] = item.Selector

	iamcertificaterequestrelationships = append(iamcertificaterequestrelationships, iamcertificaterequestrelationship)
	return iamcertificaterequestrelationships
}
func flattenMapIamClientMeta(p *models.IamClientMeta, d *schema.ResourceData) []map[string]interface{} {
	var iamclientmetas []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamclientmeta := make(map[string]interface{})
	iamclientmeta["class_id"] = item.ClassId
	iamclientmeta["device_model"] = item.DeviceModel
	iamclientmeta["object_type"] = item.ObjectType
	iamclientmeta["user_agent"] = item.UserAgent

	iamclientmetas = append(iamclientmetas, iamclientmeta)
	return iamclientmetas
}
func flattenMapIamDomainGroupRelationship(p *models.IamDomainGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamDomainGroupRelationshipInterface.(*models.MoMoRef)
	iamdomaingrouprelationship := make(map[string]interface{})
	iamdomaingrouprelationship["class_id"] = item.ClassId
	iamdomaingrouprelationship["link"] = item.Link
	iamdomaingrouprelationship["moid"] = item.Moid
	iamdomaingrouprelationship["object_type"] = item.ObjectType
	iamdomaingrouprelationship["selector"] = item.Selector

	iamdomaingrouprelationships = append(iamdomaingrouprelationships, iamdomaingrouprelationship)
	return iamdomaingrouprelationships
}
func flattenMapIamEndPointPasswordProperties(p *models.IamEndPointPasswordProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointpasswordpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamendpointpasswordproperties := make(map[string]interface{})
	iamendpointpasswordproperties["class_id"] = item.ClassId
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
func flattenMapIamEndPointUserRelationship(p *models.IamEndPointUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamEndPointUserRelationshipInterface.(*models.MoMoRef)
	iamendpointuserrelationship := make(map[string]interface{})
	iamendpointuserrelationship["class_id"] = item.ClassId
	iamendpointuserrelationship["link"] = item.Link
	iamendpointuserrelationship["moid"] = item.Moid
	iamendpointuserrelationship["object_type"] = item.ObjectType
	iamendpointuserrelationship["selector"] = item.Selector

	iamendpointuserrelationships = append(iamendpointuserrelationships, iamendpointuserrelationship)
	return iamendpointuserrelationships
}
func flattenMapIamEndPointUserPolicyRelationship(p *models.IamEndPointUserPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamEndPointUserPolicyRelationshipInterface.(*models.MoMoRef)
	iamendpointuserpolicyrelationship := make(map[string]interface{})
	iamendpointuserpolicyrelationship["class_id"] = item.ClassId
	iamendpointuserpolicyrelationship["link"] = item.Link
	iamendpointuserpolicyrelationship["moid"] = item.Moid
	iamendpointuserpolicyrelationship["object_type"] = item.ObjectType
	iamendpointuserpolicyrelationship["selector"] = item.Selector

	iamendpointuserpolicyrelationships = append(iamendpointuserpolicyrelationships, iamendpointuserpolicyrelationship)
	return iamendpointuserpolicyrelationships
}
func flattenMapIamIdpRelationship(p *models.IamIdpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamIdpRelationshipInterface.(*models.MoMoRef)
	iamidprelationship := make(map[string]interface{})
	iamidprelationship["class_id"] = item.ClassId
	iamidprelationship["link"] = item.Link
	iamidprelationship["moid"] = item.Moid
	iamidprelationship["object_type"] = item.ObjectType
	iamidprelationship["selector"] = item.Selector

	iamidprelationships = append(iamidprelationships, iamidprelationship)
	return iamidprelationships
}
func flattenMapIamIdpReferenceRelationship(p *models.IamIdpReferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamIdpReferenceRelationshipInterface.(*models.MoMoRef)
	iamidpreferencerelationship := make(map[string]interface{})
	iamidpreferencerelationship["class_id"] = item.ClassId
	iamidpreferencerelationship["link"] = item.Link
	iamidpreferencerelationship["moid"] = item.Moid
	iamidpreferencerelationship["object_type"] = item.ObjectType
	iamidpreferencerelationship["selector"] = item.Selector

	iamidpreferencerelationships = append(iamidpreferencerelationships, iamidpreferencerelationship)
	return iamidpreferencerelationships
}
func flattenMapIamLdapBaseProperties(p *models.IamLdapBaseProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamldapbasepropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamldapbaseproperties := make(map[string]interface{})
	iamldapbaseproperties["attribute"] = item.Attribute
	iamldapbaseproperties["base_dn"] = item.BaseDn
	iamldapbaseproperties["bind_dn"] = item.BindDn
	iamldapbaseproperties["bind_method"] = item.BindMethod
	iamldapbaseproperties["class_id"] = item.ClassId
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
func flattenMapIamLdapDnsParameters(p *models.IamLdapDnsParameters, d *schema.ResourceData) []map[string]interface{} {
	var iamldapdnsparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	iamldapdnsparameters := make(map[string]interface{})
	iamldapdnsparameters["class_id"] = item.ClassId
	iamldapdnsparameters["object_type"] = item.ObjectType
	iamldapdnsparameters["search_domain"] = item.SearchDomain
	iamldapdnsparameters["search_forest"] = item.SearchForest
	iamldapdnsparameters["source"] = item.Source

	iamldapdnsparameterss = append(iamldapdnsparameterss, iamldapdnsparameters)
	return iamldapdnsparameterss
}
func flattenMapIamLdapPolicyRelationship(p *models.IamLdapPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldappolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamLdapPolicyRelationshipInterface.(*models.MoMoRef)
	iamldappolicyrelationship := make(map[string]interface{})
	iamldappolicyrelationship["class_id"] = item.ClassId
	iamldappolicyrelationship["link"] = item.Link
	iamldappolicyrelationship["moid"] = item.Moid
	iamldappolicyrelationship["object_type"] = item.ObjectType
	iamldappolicyrelationship["selector"] = item.Selector

	iamldappolicyrelationships = append(iamldappolicyrelationships, iamldappolicyrelationship)
	return iamldappolicyrelationships
}
func flattenMapIamLocalUserPasswordRelationship(p *models.IamLocalUserPasswordRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamlocaluserpasswordrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamLocalUserPasswordRelationshipInterface.(*models.MoMoRef)
	iamlocaluserpasswordrelationship := make(map[string]interface{})
	iamlocaluserpasswordrelationship["class_id"] = item.ClassId
	iamlocaluserpasswordrelationship["link"] = item.Link
	iamlocaluserpasswordrelationship["moid"] = item.Moid
	iamlocaluserpasswordrelationship["object_type"] = item.ObjectType
	iamlocaluserpasswordrelationship["selector"] = item.Selector

	iamlocaluserpasswordrelationships = append(iamlocaluserpasswordrelationships, iamlocaluserpasswordrelationship)
	return iamlocaluserpasswordrelationships
}
func flattenMapIamPermissionRelationship(p *models.IamPermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamPermissionRelationshipInterface.(*models.MoMoRef)
	iampermissionrelationship := make(map[string]interface{})
	iampermissionrelationship["class_id"] = item.ClassId
	iampermissionrelationship["link"] = item.Link
	iampermissionrelationship["moid"] = item.Moid
	iampermissionrelationship["object_type"] = item.ObjectType
	iampermissionrelationship["selector"] = item.Selector

	iampermissionrelationships = append(iampermissionrelationships, iampermissionrelationship)
	return iampermissionrelationships
}
func flattenMapIamPrivateKeySpecRelationship(p *models.IamPrivateKeySpecRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivatekeyspecrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamPrivateKeySpecRelationshipInterface.(*models.MoMoRef)
	iamprivatekeyspecrelationship := make(map[string]interface{})
	iamprivatekeyspecrelationship["class_id"] = item.ClassId
	iamprivatekeyspecrelationship["link"] = item.Link
	iamprivatekeyspecrelationship["moid"] = item.Moid
	iamprivatekeyspecrelationship["object_type"] = item.ObjectType
	iamprivatekeyspecrelationship["selector"] = item.Selector

	iamprivatekeyspecrelationships = append(iamprivatekeyspecrelationships, iamprivatekeyspecrelationship)
	return iamprivatekeyspecrelationships
}
func flattenMapIamQualifierRelationship(p *models.IamQualifierRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamqualifierrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamQualifierRelationshipInterface.(*models.MoMoRef)
	iamqualifierrelationship := make(map[string]interface{})
	iamqualifierrelationship["class_id"] = item.ClassId
	iamqualifierrelationship["link"] = item.Link
	iamqualifierrelationship["moid"] = item.Moid
	iamqualifierrelationship["object_type"] = item.ObjectType
	iamqualifierrelationship["selector"] = item.Selector

	iamqualifierrelationships = append(iamqualifierrelationships, iamqualifierrelationship)
	return iamqualifierrelationships
}
func flattenMapIamResourceLimitsRelationship(p *models.IamResourceLimitsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcelimitsrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamResourceLimitsRelationshipInterface.(*models.MoMoRef)
	iamresourcelimitsrelationship := make(map[string]interface{})
	iamresourcelimitsrelationship["class_id"] = item.ClassId
	iamresourcelimitsrelationship["link"] = item.Link
	iamresourcelimitsrelationship["moid"] = item.Moid
	iamresourcelimitsrelationship["object_type"] = item.ObjectType
	iamresourcelimitsrelationship["selector"] = item.Selector

	iamresourcelimitsrelationships = append(iamresourcelimitsrelationships, iamresourcelimitsrelationship)
	return iamresourcelimitsrelationships
}
func flattenMapIamSecurityHolderRelationship(p *models.IamSecurityHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsecurityholderrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamSecurityHolderRelationshipInterface.(*models.MoMoRef)
	iamsecurityholderrelationship := make(map[string]interface{})
	iamsecurityholderrelationship["class_id"] = item.ClassId
	iamsecurityholderrelationship["link"] = item.Link
	iamsecurityholderrelationship["moid"] = item.Moid
	iamsecurityholderrelationship["object_type"] = item.ObjectType
	iamsecurityholderrelationship["selector"] = item.Selector

	iamsecurityholderrelationships = append(iamsecurityholderrelationships, iamsecurityholderrelationship)
	return iamsecurityholderrelationships
}
func flattenMapIamServiceProviderRelationship(p *models.IamServiceProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamserviceproviderrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamServiceProviderRelationshipInterface.(*models.MoMoRef)
	iamserviceproviderrelationship := make(map[string]interface{})
	iamserviceproviderrelationship["class_id"] = item.ClassId
	iamserviceproviderrelationship["link"] = item.Link
	iamserviceproviderrelationship["moid"] = item.Moid
	iamserviceproviderrelationship["object_type"] = item.ObjectType
	iamserviceproviderrelationship["selector"] = item.Selector

	iamserviceproviderrelationships = append(iamserviceproviderrelationships, iamserviceproviderrelationship)
	return iamserviceproviderrelationships
}
func flattenMapIamSessionRelationship(p *models.IamSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamSessionRelationshipInterface.(*models.MoMoRef)
	iamsessionrelationship := make(map[string]interface{})
	iamsessionrelationship["class_id"] = item.ClassId
	iamsessionrelationship["link"] = item.Link
	iamsessionrelationship["moid"] = item.Moid
	iamsessionrelationship["object_type"] = item.ObjectType
	iamsessionrelationship["selector"] = item.Selector

	iamsessionrelationships = append(iamsessionrelationships, iamsessionrelationship)
	return iamsessionrelationships
}
func flattenMapIamSessionLimitsRelationship(p *models.IamSessionLimitsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionlimitsrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamSessionLimitsRelationshipInterface.(*models.MoMoRef)
	iamsessionlimitsrelationship := make(map[string]interface{})
	iamsessionlimitsrelationship["class_id"] = item.ClassId
	iamsessionlimitsrelationship["link"] = item.Link
	iamsessionlimitsrelationship["moid"] = item.Moid
	iamsessionlimitsrelationship["object_type"] = item.ObjectType
	iamsessionlimitsrelationship["selector"] = item.Selector

	iamsessionlimitsrelationships = append(iamsessionlimitsrelationships, iamsessionlimitsrelationship)
	return iamsessionlimitsrelationships
}
func flattenMapIamSystemRelationship(p *models.IamSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsystemrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamSystemRelationshipInterface.(*models.MoMoRef)
	iamsystemrelationship := make(map[string]interface{})
	iamsystemrelationship["class_id"] = item.ClassId
	iamsystemrelationship["link"] = item.Link
	iamsystemrelationship["moid"] = item.Moid
	iamsystemrelationship["object_type"] = item.ObjectType
	iamsystemrelationship["selector"] = item.Selector

	iamsystemrelationships = append(iamsystemrelationships, iamsystemrelationship)
	return iamsystemrelationships
}
func flattenMapIamUserRelationship(p *models.IamUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamUserRelationshipInterface.(*models.MoMoRef)
	iamuserrelationship := make(map[string]interface{})
	iamuserrelationship["class_id"] = item.ClassId
	iamuserrelationship["link"] = item.Link
	iamuserrelationship["moid"] = item.Moid
	iamuserrelationship["object_type"] = item.ObjectType
	iamuserrelationship["selector"] = item.Selector

	iamuserrelationships = append(iamuserrelationships, iamuserrelationship)
	return iamuserrelationships
}
func flattenMapIamUserGroupRelationship(p *models.IamUserGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.IamUserGroupRelationshipInterface.(*models.MoMoRef)
	iamusergrouprelationship := make(map[string]interface{})
	iamusergrouprelationship["class_id"] = item.ClassId
	iamusergrouprelationship["link"] = item.Link
	iamusergrouprelationship["moid"] = item.Moid
	iamusergrouprelationship["object_type"] = item.ObjectType
	iamusergrouprelationship["selector"] = item.Selector

	iamusergrouprelationships = append(iamusergrouprelationships, iamusergrouprelationship)
	return iamusergrouprelationships
}
func flattenMapInfraHardwareInfo(p *models.InfraHardwareInfo, d *schema.ResourceData) []map[string]interface{} {
	var infrahardwareinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	infrahardwareinfo := make(map[string]interface{})
	infrahardwareinfo["class_id"] = item.ClassId
	infrahardwareinfo["cpu_cores"] = item.CpuCores
	infrahardwareinfo["cpu_speed"] = item.CpuSpeed
	infrahardwareinfo["memory_size"] = item.MemorySize
	infrahardwareinfo["object_type"] = item.ObjectType

	infrahardwareinfos = append(infrahardwareinfos, infrahardwareinfo)
	return infrahardwareinfos
}
func flattenMapInventoryBaseRelationship(p *models.InventoryBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorybaserelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.InventoryBaseRelationshipInterface.(*models.MoMoRef)
	inventorybaserelationship := make(map[string]interface{})
	inventorybaserelationship["class_id"] = item.ClassId
	inventorybaserelationship["link"] = item.Link
	inventorybaserelationship["moid"] = item.Moid
	inventorybaserelationship["object_type"] = item.ObjectType
	inventorybaserelationship["selector"] = item.Selector

	inventorybaserelationships = append(inventorybaserelationships, inventorybaserelationship)
	return inventorybaserelationships
}
func flattenMapInventoryGenericInventoryHolderRelationship(p *models.InventoryGenericInventoryHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.InventoryGenericInventoryHolderRelationshipInterface.(*models.MoMoRef)
	inventorygenericinventoryholderrelationship := make(map[string]interface{})
	inventorygenericinventoryholderrelationship["class_id"] = item.ClassId
	inventorygenericinventoryholderrelationship["link"] = item.Link
	inventorygenericinventoryholderrelationship["moid"] = item.Moid
	inventorygenericinventoryholderrelationship["object_type"] = item.ObjectType
	inventorygenericinventoryholderrelationship["selector"] = item.Selector

	inventorygenericinventoryholderrelationships = append(inventorygenericinventoryholderrelationships, inventorygenericinventoryholderrelationship)
	return inventorygenericinventoryholderrelationships
}
func flattenMapLicenseAccountLicenseDataRelationship(p *models.LicenseAccountLicenseDataRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenseaccountlicensedatarelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.LicenseAccountLicenseDataRelationshipInterface.(*models.MoMoRef)
	licenseaccountlicensedatarelationship := make(map[string]interface{})
	licenseaccountlicensedatarelationship["class_id"] = item.ClassId
	licenseaccountlicensedatarelationship["link"] = item.Link
	licenseaccountlicensedatarelationship["moid"] = item.Moid
	licenseaccountlicensedatarelationship["object_type"] = item.ObjectType
	licenseaccountlicensedatarelationship["selector"] = item.Selector

	licenseaccountlicensedatarelationships = append(licenseaccountlicensedatarelationships, licenseaccountlicensedatarelationship)
	return licenseaccountlicensedatarelationships
}
func flattenMapLicenseCustomerOpRelationship(p *models.LicenseCustomerOpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licensecustomeroprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.LicenseCustomerOpRelationshipInterface.(*models.MoMoRef)
	licensecustomeroprelationship := make(map[string]interface{})
	licensecustomeroprelationship["class_id"] = item.ClassId
	licensecustomeroprelationship["link"] = item.Link
	licensecustomeroprelationship["moid"] = item.Moid
	licensecustomeroprelationship["object_type"] = item.ObjectType
	licensecustomeroprelationship["selector"] = item.Selector

	licensecustomeroprelationships = append(licensecustomeroprelationships, licensecustomeroprelationship)
	return licensecustomeroprelationships
}
func flattenMapLicenseSmartlicenseTokenRelationship(p *models.LicenseSmartlicenseTokenRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licensesmartlicensetokenrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.LicenseSmartlicenseTokenRelationshipInterface.(*models.MoMoRef)
	licensesmartlicensetokenrelationship := make(map[string]interface{})
	licensesmartlicensetokenrelationship["class_id"] = item.ClassId
	licensesmartlicensetokenrelationship["link"] = item.Link
	licensesmartlicensetokenrelationship["moid"] = item.Moid
	licensesmartlicensetokenrelationship["object_type"] = item.ObjectType
	licensesmartlicensetokenrelationship["selector"] = item.Selector

	licensesmartlicensetokenrelationships = append(licensesmartlicensetokenrelationships, licensesmartlicensetokenrelationship)
	return licensesmartlicensetokenrelationships
}
func flattenMapManagementControllerRelationship(p *models.ManagementControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managementcontrollerrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ManagementControllerRelationshipInterface.(*models.MoMoRef)
	managementcontrollerrelationship := make(map[string]interface{})
	managementcontrollerrelationship["class_id"] = item.ClassId
	managementcontrollerrelationship["link"] = item.Link
	managementcontrollerrelationship["moid"] = item.Moid
	managementcontrollerrelationship["object_type"] = item.ObjectType
	managementcontrollerrelationship["selector"] = item.Selector

	managementcontrollerrelationships = append(managementcontrollerrelationships, managementcontrollerrelationship)
	return managementcontrollerrelationships
}
func flattenMapManagementEntityRelationship(p *models.ManagementEntityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managemententityrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ManagementEntityRelationshipInterface.(*models.MoMoRef)
	managemententityrelationship := make(map[string]interface{})
	managemententityrelationship["class_id"] = item.ClassId
	managemententityrelationship["link"] = item.Link
	managemententityrelationship["moid"] = item.Moid
	managemententityrelationship["object_type"] = item.ObjectType
	managemententityrelationship["selector"] = item.Selector

	managemententityrelationships = append(managemententityrelationships, managemententityrelationship)
	return managemententityrelationships
}
func flattenMapMemoryArrayRelationship(p *models.MemoryArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.MemoryArrayRelationshipInterface.(*models.MoMoRef)
	memoryarrayrelationship := make(map[string]interface{})
	memoryarrayrelationship["class_id"] = item.ClassId
	memoryarrayrelationship["link"] = item.Link
	memoryarrayrelationship["moid"] = item.Moid
	memoryarrayrelationship["object_type"] = item.ObjectType
	memoryarrayrelationship["selector"] = item.Selector

	memoryarrayrelationships = append(memoryarrayrelationships, memoryarrayrelationship)
	return memoryarrayrelationships
}
func flattenMapMemoryPersistentMemoryConfigResultRelationship(p *models.MemoryPersistentMemoryConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigresultrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.MemoryPersistentMemoryConfigResultRelationshipInterface.(*models.MoMoRef)
	memorypersistentmemoryconfigresultrelationship := make(map[string]interface{})
	memorypersistentmemoryconfigresultrelationship["class_id"] = item.ClassId
	memorypersistentmemoryconfigresultrelationship["link"] = item.Link
	memorypersistentmemoryconfigresultrelationship["moid"] = item.Moid
	memorypersistentmemoryconfigresultrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryconfigresultrelationship["selector"] = item.Selector

	memorypersistentmemoryconfigresultrelationships = append(memorypersistentmemoryconfigresultrelationships, memorypersistentmemoryconfigresultrelationship)
	return memorypersistentmemoryconfigresultrelationships
}
func flattenMapMemoryPersistentMemoryConfigurationRelationship(p *models.MemoryPersistentMemoryConfigurationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigurationrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.MemoryPersistentMemoryConfigurationRelationshipInterface.(*models.MoMoRef)
	memorypersistentmemoryconfigurationrelationship := make(map[string]interface{})
	memorypersistentmemoryconfigurationrelationship["class_id"] = item.ClassId
	memorypersistentmemoryconfigurationrelationship["link"] = item.Link
	memorypersistentmemoryconfigurationrelationship["moid"] = item.Moid
	memorypersistentmemoryconfigurationrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryconfigurationrelationship["selector"] = item.Selector

	memorypersistentmemoryconfigurationrelationships = append(memorypersistentmemoryconfigurationrelationships, memorypersistentmemoryconfigurationrelationship)
	return memorypersistentmemoryconfigurationrelationships
}
func flattenMapMemoryPersistentMemoryLocalSecurity(p *models.MemoryPersistentMemoryLocalSecurity, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylocalsecuritys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemorylocalsecurity := make(map[string]interface{})
	memorypersistentmemorylocalsecurity["class_id"] = item.ClassId
	memorypersistentmemorylocalsecurity["enabled"] = item.Enabled
	memorypersistentmemorylocalsecurity["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	memorypersistentmemorylocalsecurity["object_type"] = item.ObjectType
	secure_passphrase_x := d.Get("local_security").([]interface{})
	secure_passphrase_y := secure_passphrase_x[0].(map[string]interface{})
	memorypersistentmemorylocalsecurity["secure_passphrase"] = secure_passphrase_y["secure_passphrase"]

	memorypersistentmemorylocalsecuritys = append(memorypersistentmemorylocalsecuritys, memorypersistentmemorylocalsecurity)
	return memorypersistentmemorylocalsecuritys
}
func flattenMapMemoryPersistentMemoryRegionRelationship(p *models.MemoryPersistentMemoryRegionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.MemoryPersistentMemoryRegionRelationshipInterface.(*models.MoMoRef)
	memorypersistentmemoryregionrelationship := make(map[string]interface{})
	memorypersistentmemoryregionrelationship["class_id"] = item.ClassId
	memorypersistentmemoryregionrelationship["link"] = item.Link
	memorypersistentmemoryregionrelationship["moid"] = item.Moid
	memorypersistentmemoryregionrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryregionrelationship["selector"] = item.Selector

	memorypersistentmemoryregionrelationships = append(memorypersistentmemoryregionrelationships, memorypersistentmemoryregionrelationship)
	return memorypersistentmemoryregionrelationships
}
func flattenMapMoBaseMoRelationship(p *models.MoBaseMoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var mobasemorelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.MoBaseMoRelationshipInterface.(*models.MoBaseMo)
	mobasemorelationship := make(map[string]interface{})

	mobasemorelationships = append(mobasemorelationships, mobasemorelationship)
	return mobasemorelationships
}
func flattenMapNetworkElementRelationship(p *models.NetworkElementRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.NetworkElementRelationshipInterface.(*models.MoMoRef)
	networkelementrelationship := make(map[string]interface{})
	networkelementrelationship["class_id"] = item.ClassId
	networkelementrelationship["link"] = item.Link
	networkelementrelationship["moid"] = item.Moid
	networkelementrelationship["object_type"] = item.ObjectType
	networkelementrelationship["selector"] = item.Selector

	networkelementrelationships = append(networkelementrelationships, networkelementrelationship)
	return networkelementrelationships
}
func flattenMapNiaapiNewReleaseDetail(p *models.NiaapiNewReleaseDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapinewreleasedetails []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	niaapinewreleasedetail := make(map[string]interface{})
	niaapinewreleasedetail["class_id"] = item.ClassId
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
	niaapiversionregexplatform["anyllregex"] = item.Anyllregex
	niaapiversionregexplatform["class_id"] = item.ClassId
	niaapiversionregexplatform["currentlltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["class_id"] = item.ClassId
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
		niaapisoftwareregex["class_id"] = item.ClassId
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.Latestsltrain, d)
	niaapiversionregexplatform["object_type"] = item.ObjectType
	niaapiversionregexplatform["sltrain"] = (func(p *[]models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if *p == nil {
			return nil
		}
		for _, item := range *p {
			niaapisoftwareregex := make(map[string]interface{})
			niaapisoftwareregex["class_id"] = item.ClassId
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
		niaapisoftwareregex["class_id"] = item.ClassId
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
	niatelemetrydiskinfo["class_id"] = item.ClassId
	niatelemetrydiskinfo["free"] = item.Free
	niatelemetrydiskinfo["name"] = item.Name
	niatelemetrydiskinfo["object_type"] = item.ObjectType
	niatelemetrydiskinfo["total"] = item.Total
	niatelemetrydiskinfo["used"] = item.Used

	niatelemetrydiskinfos = append(niatelemetrydiskinfos, niatelemetrydiskinfo)
	return niatelemetrydiskinfos
}
func flattenMapNiatelemetryNiaInventoryRelationship(p *models.NiatelemetryNiaInventoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetryniainventoryrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.NiatelemetryNiaInventoryRelationshipInterface.(*models.MoMoRef)
	niatelemetryniainventoryrelationship := make(map[string]interface{})
	niatelemetryniainventoryrelationship["class_id"] = item.ClassId
	niatelemetryniainventoryrelationship["link"] = item.Link
	niatelemetryniainventoryrelationship["moid"] = item.Moid
	niatelemetryniainventoryrelationship["object_type"] = item.ObjectType
	niatelemetryniainventoryrelationship["selector"] = item.Selector

	niatelemetryniainventoryrelationships = append(niatelemetryniainventoryrelationships, niatelemetryniainventoryrelationship)
	return niatelemetryniainventoryrelationships
}
func flattenMapNiatelemetryNiaLicenseStateRelationship(p *models.NiatelemetryNiaLicenseStateRelationship, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynialicensestaterelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.NiatelemetryNiaLicenseStateRelationshipInterface.(*models.MoMoRef)
	niatelemetrynialicensestaterelationship := make(map[string]interface{})
	niatelemetrynialicensestaterelationship["class_id"] = item.ClassId
	niatelemetrynialicensestaterelationship["link"] = item.Link
	niatelemetrynialicensestaterelationship["moid"] = item.Moid
	niatelemetrynialicensestaterelationship["object_type"] = item.ObjectType
	niatelemetrynialicensestaterelationship["selector"] = item.Selector

	niatelemetrynialicensestaterelationships = append(niatelemetrynialicensestaterelationships, niatelemetrynialicensestaterelationship)
	return niatelemetrynialicensestaterelationships
}
func flattenMapOnpremSchedule(p *models.OnpremSchedule, d *schema.ResourceData) []map[string]interface{} {
	var onpremschedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	onpremschedule := make(map[string]interface{})
	onpremschedule["class_id"] = item.ClassId
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
	onpremupgradephase["class_id"] = item.ClassId
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
func flattenMapOrganizationOrganizationRelationship(p *models.OrganizationOrganizationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.OrganizationOrganizationRelationshipInterface.(*models.MoMoRef)
	organizationorganizationrelationship := make(map[string]interface{})
	organizationorganizationrelationship["class_id"] = item.ClassId
	organizationorganizationrelationship["link"] = item.Link
	organizationorganizationrelationship["moid"] = item.Moid
	organizationorganizationrelationship["object_type"] = item.ObjectType
	organizationorganizationrelationship["selector"] = item.Selector

	organizationorganizationrelationships = append(organizationorganizationrelationships, organizationorganizationrelationship)
	return organizationorganizationrelationships
}
func flattenMapOsAnswers(p *models.OsAnswers, d *schema.ResourceData) []map[string]interface{} {
	var osanswerss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	osanswers := make(map[string]interface{})
	answer_file_x := d.Get("answers").([]interface{})
	answer_file_y := answer_file_x[0].(map[string]interface{})
	osanswers["answer_file"] = answer_file_y["answer_file"]
	osanswers["class_id"] = item.ClassId
	osanswers["hostname"] = item.Hostname
	osanswers["ip_config_type"] = item.IpConfigType
	osanswers["ip_v4_config"] = (func(p *models.CommIpV4Interface, d *schema.ResourceData) []map[string]interface{} {
		var commipv4interfaces []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		commipv4interface := make(map[string]interface{})
		commipv4interface["class_id"] = item.ClassId
		commipv4interface["gateway"] = item.Gateway
		commipv4interface["ip_address"] = item.IpAddress
		commipv4interface["netmask"] = item.Netmask
		commipv4interface["object_type"] = item.ObjectType

		commipv4interfaces = append(commipv4interfaces, commipv4interface)
		return commipv4interfaces
	})(item.IpV4Config, d)
	osanswers["is_answer_file_set"] = item.IsAnswerFileSet
	osanswers["is_root_password_crypted"] = item.IsRootPasswordCrypted
	osanswers["is_root_password_set"] = item.IsRootPasswordSet
	osanswers["nameserver"] = item.Nameserver
	osanswers["object_type"] = item.ObjectType
	osanswers["product_key"] = item.ProductKey
	root_password_x := d.Get("answers").([]interface{})
	root_password_y := root_password_x[0].(map[string]interface{})
	osanswers["root_password"] = root_password_y["root_password"]
	osanswers["source"] = item.Source

	osanswerss = append(osanswerss, osanswers)
	return osanswerss
}
func flattenMapOsCatalogRelationship(p *models.OsCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var oscatalogrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.OsCatalogRelationshipInterface.(*models.MoMoRef)
	oscatalogrelationship := make(map[string]interface{})
	oscatalogrelationship["class_id"] = item.ClassId
	oscatalogrelationship["link"] = item.Link
	oscatalogrelationship["moid"] = item.Moid
	oscatalogrelationship["object_type"] = item.ObjectType
	oscatalogrelationship["selector"] = item.Selector

	oscatalogrelationships = append(oscatalogrelationships, oscatalogrelationship)
	return oscatalogrelationships
}
func flattenMapOsConfigurationFileRelationship(p *models.OsConfigurationFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.OsConfigurationFileRelationshipInterface.(*models.MoMoRef)
	osconfigurationfilerelationship := make(map[string]interface{})
	osconfigurationfilerelationship["class_id"] = item.ClassId
	osconfigurationfilerelationship["link"] = item.Link
	osconfigurationfilerelationship["moid"] = item.Moid
	osconfigurationfilerelationship["object_type"] = item.ObjectType
	osconfigurationfilerelationship["selector"] = item.Selector

	osconfigurationfilerelationships = append(osconfigurationfilerelationships, osconfigurationfilerelationship)
	return osconfigurationfilerelationships
}
func flattenMapOsOperatingSystemParameters(p *models.OsOperatingSystemParameters, d *schema.ResourceData) []map[string]interface{} {
	var osoperatingsystemparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	osoperatingsystemparameters := make(map[string]interface{})
	osoperatingsystemparameters["class_id"] = item.ClassId
	osoperatingsystemparameters["object_type"] = item.ObjectType

	osoperatingsystemparameterss = append(osoperatingsystemparameterss, osoperatingsystemparameters)
	return osoperatingsystemparameterss
}
func flattenMapPciSwitchRelationship(p *models.PciSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.PciSwitchRelationshipInterface.(*models.MoMoRef)
	pciswitchrelationship := make(map[string]interface{})
	pciswitchrelationship["class_id"] = item.ClassId
	pciswitchrelationship["link"] = item.Link
	pciswitchrelationship["moid"] = item.Moid
	pciswitchrelationship["object_type"] = item.ObjectType
	pciswitchrelationship["selector"] = item.Selector

	pciswitchrelationships = append(pciswitchrelationships, pciswitchrelationship)
	return pciswitchrelationships
}
func flattenMapPkixDistinguishedName(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
	var pkixdistinguishednames []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	pkixdistinguishedname := make(map[string]interface{})
	pkixdistinguishedname["class_id"] = item.ClassId
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
	pkixkeygenerationspec["class_id"] = item.ClassId
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
	pkixsubjectalternatename["class_id"] = item.ClassId
	pkixsubjectalternatename["dns_name"] = item.DnsName
	pkixsubjectalternatename["email_address"] = item.EmailAddress
	pkixsubjectalternatename["ip_address"] = item.IpAddress
	pkixsubjectalternatename["object_type"] = item.ObjectType
	pkixsubjectalternatename["uri"] = item.Uri

	pkixsubjectalternatenames = append(pkixsubjectalternatenames, pkixsubjectalternatename)
	return pkixsubjectalternatenames
}
func flattenMapPolicyAbstractProfileRelationship(p *models.PolicyAbstractProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractprofilerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.PolicyAbstractProfileRelationshipInterface.(*models.MoMoRef)
	policyabstractprofilerelationship := make(map[string]interface{})
	policyabstractprofilerelationship["class_id"] = item.ClassId
	policyabstractprofilerelationship["link"] = item.Link
	policyabstractprofilerelationship["moid"] = item.Moid
	policyabstractprofilerelationship["object_type"] = item.ObjectType
	policyabstractprofilerelationship["selector"] = item.Selector

	policyabstractprofilerelationships = append(policyabstractprofilerelationships, policyabstractprofilerelationship)
	return policyabstractprofilerelationships
}
func flattenMapPolicyConfigChange(p *models.PolicyConfigChange, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigchanges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	policyconfigchange := make(map[string]interface{})
	policyconfigchange["changes"] = item.Changes
	policyconfigchange["class_id"] = item.ClassId
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
	policyconfigcontext["class_id"] = item.ClassId
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
	policyconfigresultcontext["class_id"] = item.ClassId
	policyconfigresultcontext["entity_data"] = item.EntityData
	policyconfigresultcontext["entity_moid"] = item.EntityMoid
	policyconfigresultcontext["entity_name"] = item.EntityName
	policyconfigresultcontext["entity_type"] = item.EntityType
	policyconfigresultcontext["object_type"] = item.ObjectType

	policyconfigresultcontexts = append(policyconfigresultcontexts, policyconfigresultcontext)
	return policyconfigresultcontexts
}
func flattenMapPortGroupRelationship(p *models.PortGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.PortGroupRelationshipInterface.(*models.MoMoRef)
	portgrouprelationship := make(map[string]interface{})
	portgrouprelationship["class_id"] = item.ClassId
	portgrouprelationship["link"] = item.Link
	portgrouprelationship["moid"] = item.Moid
	portgrouprelationship["object_type"] = item.ObjectType
	portgrouprelationship["selector"] = item.Selector

	portgrouprelationships = append(portgrouprelationships, portgrouprelationship)
	return portgrouprelationships
}
func flattenMapPortSubGroupRelationship(p *models.PortSubGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.PortSubGroupRelationshipInterface.(*models.MoMoRef)
	portsubgrouprelationship := make(map[string]interface{})
	portsubgrouprelationship["class_id"] = item.ClassId
	portsubgrouprelationship["link"] = item.Link
	portsubgrouprelationship["moid"] = item.Moid
	portsubgrouprelationship["object_type"] = item.ObjectType
	portsubgrouprelationship["selector"] = item.Selector

	portsubgrouprelationships = append(portsubgrouprelationships, portsubgrouprelationship)
	return portsubgrouprelationships
}
func flattenMapRecoveryAbstractBackupInfoRelationship(p *models.RecoveryAbstractBackupInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryabstractbackupinforelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.RecoveryAbstractBackupInfoRelationshipInterface.(*models.MoMoRef)
	recoveryabstractbackupinforelationship := make(map[string]interface{})
	recoveryabstractbackupinforelationship["class_id"] = item.ClassId
	recoveryabstractbackupinforelationship["link"] = item.Link
	recoveryabstractbackupinforelationship["moid"] = item.Moid
	recoveryabstractbackupinforelationship["object_type"] = item.ObjectType
	recoveryabstractbackupinforelationship["selector"] = item.Selector

	recoveryabstractbackupinforelationships = append(recoveryabstractbackupinforelationships, recoveryabstractbackupinforelationship)
	return recoveryabstractbackupinforelationships
}
func flattenMapRecoveryBackupConfigPolicyRelationship(p *models.RecoveryBackupConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupconfigpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.RecoveryBackupConfigPolicyRelationshipInterface.(*models.MoMoRef)
	recoverybackupconfigpolicyrelationship := make(map[string]interface{})
	recoverybackupconfigpolicyrelationship["class_id"] = item.ClassId
	recoverybackupconfigpolicyrelationship["link"] = item.Link
	recoverybackupconfigpolicyrelationship["moid"] = item.Moid
	recoverybackupconfigpolicyrelationship["object_type"] = item.ObjectType
	recoverybackupconfigpolicyrelationship["selector"] = item.Selector

	recoverybackupconfigpolicyrelationships = append(recoverybackupconfigpolicyrelationships, recoverybackupconfigpolicyrelationship)
	return recoverybackupconfigpolicyrelationships
}
func flattenMapRecoveryBackupProfileRelationship(p *models.RecoveryBackupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.RecoveryBackupProfileRelationshipInterface.(*models.MoMoRef)
	recoverybackupprofilerelationship := make(map[string]interface{})
	recoverybackupprofilerelationship["class_id"] = item.ClassId
	recoverybackupprofilerelationship["link"] = item.Link
	recoverybackupprofilerelationship["moid"] = item.Moid
	recoverybackupprofilerelationship["object_type"] = item.ObjectType
	recoverybackupprofilerelationship["selector"] = item.Selector

	recoverybackupprofilerelationships = append(recoverybackupprofilerelationships, recoverybackupprofilerelationship)
	return recoverybackupprofilerelationships
}
func flattenMapRecoveryBackupSchedule(p *models.RecoveryBackupSchedule, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupschedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	recoverybackupschedule := make(map[string]interface{})
	recoverybackupschedule["class_id"] = item.ClassId
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
	recoveryconfigparams["class_id"] = item.ClassId
	recoveryconfigparams["object_type"] = item.ObjectType

	recoveryconfigparamss = append(recoveryconfigparamss, recoveryconfigparams)
	return recoveryconfigparamss
}
func flattenMapRecoveryConfigResultRelationship(p *models.RecoveryConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.RecoveryConfigResultRelationshipInterface.(*models.MoMoRef)
	recoveryconfigresultrelationship := make(map[string]interface{})
	recoveryconfigresultrelationship["class_id"] = item.ClassId
	recoveryconfigresultrelationship["link"] = item.Link
	recoveryconfigresultrelationship["moid"] = item.Moid
	recoveryconfigresultrelationship["object_type"] = item.ObjectType
	recoveryconfigresultrelationship["selector"] = item.Selector

	recoveryconfigresultrelationships = append(recoveryconfigresultrelationships, recoveryconfigresultrelationship)
	return recoveryconfigresultrelationships
}
func flattenMapRecoveryScheduleConfigPolicyRelationship(p *models.RecoveryScheduleConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryscheduleconfigpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.RecoveryScheduleConfigPolicyRelationshipInterface.(*models.MoMoRef)
	recoveryscheduleconfigpolicyrelationship := make(map[string]interface{})
	recoveryscheduleconfigpolicyrelationship["class_id"] = item.ClassId
	recoveryscheduleconfigpolicyrelationship["link"] = item.Link
	recoveryscheduleconfigpolicyrelationship["moid"] = item.Moid
	recoveryscheduleconfigpolicyrelationship["object_type"] = item.ObjectType
	recoveryscheduleconfigpolicyrelationship["selector"] = item.Selector

	recoveryscheduleconfigpolicyrelationships = append(recoveryscheduleconfigpolicyrelationships, recoveryscheduleconfigpolicyrelationship)
	return recoveryscheduleconfigpolicyrelationships
}
func flattenMapResourceGroupRelationship(p *models.ResourceGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ResourceGroupRelationshipInterface.(*models.MoMoRef)
	resourcegrouprelationship := make(map[string]interface{})
	resourcegrouprelationship["class_id"] = item.ClassId
	resourcegrouprelationship["link"] = item.Link
	resourcegrouprelationship["moid"] = item.Moid
	resourcegrouprelationship["object_type"] = item.ObjectType
	resourcegrouprelationship["selector"] = item.Selector

	resourcegrouprelationships = append(resourcegrouprelationships, resourcegrouprelationship)
	return resourcegrouprelationships
}
func flattenMapResourceMembershipHolderRelationship(p *models.ResourceMembershipHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcemembershipholderrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ResourceMembershipHolderRelationshipInterface.(*models.MoMoRef)
	resourcemembershipholderrelationship := make(map[string]interface{})
	resourcemembershipholderrelationship["class_id"] = item.ClassId
	resourcemembershipholderrelationship["link"] = item.Link
	resourcemembershipholderrelationship["moid"] = item.Moid
	resourcemembershipholderrelationship["object_type"] = item.ObjectType
	resourcemembershipholderrelationship["selector"] = item.Selector

	resourcemembershipholderrelationships = append(resourcemembershipholderrelationships, resourcemembershipholderrelationship)
	return resourcemembershipholderrelationships
}
func flattenMapSdwanProfileRelationship(p *models.SdwanProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.SdwanProfileRelationshipInterface.(*models.MoMoRef)
	sdwanprofilerelationship := make(map[string]interface{})
	sdwanprofilerelationship["class_id"] = item.ClassId
	sdwanprofilerelationship["link"] = item.Link
	sdwanprofilerelationship["moid"] = item.Moid
	sdwanprofilerelationship["object_type"] = item.ObjectType
	sdwanprofilerelationship["selector"] = item.Selector

	sdwanprofilerelationships = append(sdwanprofilerelationships, sdwanprofilerelationship)
	return sdwanprofilerelationships
}
func flattenMapSdwanRouterPolicyRelationship(p *models.SdwanRouterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouterpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.SdwanRouterPolicyRelationshipInterface.(*models.MoMoRef)
	sdwanrouterpolicyrelationship := make(map[string]interface{})
	sdwanrouterpolicyrelationship["class_id"] = item.ClassId
	sdwanrouterpolicyrelationship["link"] = item.Link
	sdwanrouterpolicyrelationship["moid"] = item.Moid
	sdwanrouterpolicyrelationship["object_type"] = item.ObjectType
	sdwanrouterpolicyrelationship["selector"] = item.Selector

	sdwanrouterpolicyrelationships = append(sdwanrouterpolicyrelationships, sdwanrouterpolicyrelationship)
	return sdwanrouterpolicyrelationships
}
func flattenMapSdwanVmanageAccountPolicyRelationship(p *models.SdwanVmanageAccountPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanvmanageaccountpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.SdwanVmanageAccountPolicyRelationshipInterface.(*models.MoMoRef)
	sdwanvmanageaccountpolicyrelationship := make(map[string]interface{})
	sdwanvmanageaccountpolicyrelationship["class_id"] = item.ClassId
	sdwanvmanageaccountpolicyrelationship["link"] = item.Link
	sdwanvmanageaccountpolicyrelationship["moid"] = item.Moid
	sdwanvmanageaccountpolicyrelationship["object_type"] = item.ObjectType
	sdwanvmanageaccountpolicyrelationship["selector"] = item.Selector

	sdwanvmanageaccountpolicyrelationships = append(sdwanvmanageaccountpolicyrelationships, sdwanvmanageaccountpolicyrelationship)
	return sdwanvmanageaccountpolicyrelationships
}
func flattenMapServerConfigResultRelationship(p *models.ServerConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ServerConfigResultRelationshipInterface.(*models.MoMoRef)
	serverconfigresultrelationship := make(map[string]interface{})
	serverconfigresultrelationship["class_id"] = item.ClassId
	serverconfigresultrelationship["link"] = item.Link
	serverconfigresultrelationship["moid"] = item.Moid
	serverconfigresultrelationship["object_type"] = item.ObjectType
	serverconfigresultrelationship["selector"] = item.Selector

	serverconfigresultrelationships = append(serverconfigresultrelationships, serverconfigresultrelationship)
	return serverconfigresultrelationships
}
func flattenMapServerProfileRelationship(p *models.ServerProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverprofilerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.ServerProfileRelationshipInterface.(*models.MoMoRef)
	serverprofilerelationship := make(map[string]interface{})
	serverprofilerelationship["class_id"] = item.ClassId
	serverprofilerelationship["link"] = item.Link
	serverprofilerelationship["moid"] = item.Moid
	serverprofilerelationship["object_type"] = item.ObjectType
	serverprofilerelationship["selector"] = item.Selector

	serverprofilerelationships = append(serverprofilerelationships, serverprofilerelationship)
	return serverprofilerelationships
}
func flattenMapSoftwareHyperflexDistributableRelationship(p *models.SoftwareHyperflexDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarehyperflexdistributablerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.SoftwareHyperflexDistributableRelationshipInterface.(*models.MoMoRef)
	softwarehyperflexdistributablerelationship := make(map[string]interface{})
	softwarehyperflexdistributablerelationship["class_id"] = item.ClassId
	softwarehyperflexdistributablerelationship["link"] = item.Link
	softwarehyperflexdistributablerelationship["moid"] = item.Moid
	softwarehyperflexdistributablerelationship["object_type"] = item.ObjectType
	softwarehyperflexdistributablerelationship["selector"] = item.Selector

	softwarehyperflexdistributablerelationships = append(softwarehyperflexdistributablerelationships, softwarehyperflexdistributablerelationship)
	return softwarehyperflexdistributablerelationships
}
func flattenMapSoftwareSolutionDistributableRelationship(p *models.SoftwareSolutionDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwaresolutiondistributablerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.SoftwareSolutionDistributableRelationshipInterface.(*models.MoMoRef)
	softwaresolutiondistributablerelationship := make(map[string]interface{})
	softwaresolutiondistributablerelationship["class_id"] = item.ClassId
	softwaresolutiondistributablerelationship["link"] = item.Link
	softwaresolutiondistributablerelationship["moid"] = item.Moid
	softwaresolutiondistributablerelationship["object_type"] = item.ObjectType
	softwaresolutiondistributablerelationship["selector"] = item.Selector

	softwaresolutiondistributablerelationships = append(softwaresolutiondistributablerelationships, softwaresolutiondistributablerelationship)
	return softwaresolutiondistributablerelationships
}
func flattenMapSoftwarerepositoryCatalogRelationship(p *models.SoftwarerepositoryCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositorycatalogrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.SoftwarerepositoryCatalogRelationshipInterface.(*models.MoMoRef)
	softwarerepositorycatalogrelationship := make(map[string]interface{})
	softwarerepositorycatalogrelationship["class_id"] = item.ClassId
	softwarerepositorycatalogrelationship["link"] = item.Link
	softwarerepositorycatalogrelationship["moid"] = item.Moid
	softwarerepositorycatalogrelationship["object_type"] = item.ObjectType
	softwarerepositorycatalogrelationship["selector"] = item.Selector

	softwarerepositorycatalogrelationships = append(softwarerepositorycatalogrelationships, softwarerepositorycatalogrelationship)
	return softwarerepositorycatalogrelationships
}
func flattenMapSoftwarerepositoryFileServer(p *models.SoftwarerepositoryFileServer, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryfileservers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	softwarerepositoryfileserver := make(map[string]interface{})
	softwarerepositoryfileserver["class_id"] = item.ClassId
	softwarerepositoryfileserver["object_type"] = item.ObjectType

	softwarerepositoryfileservers = append(softwarerepositoryfileservers, softwarerepositoryfileserver)
	return softwarerepositoryfileservers
}
func flattenMapSoftwarerepositoryOperatingSystemFileRelationship(p *models.SoftwarerepositoryOperatingSystemFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryoperatingsystemfilerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.SoftwarerepositoryOperatingSystemFileRelationshipInterface.(*models.MoMoRef)
	softwarerepositoryoperatingsystemfilerelationship := make(map[string]interface{})
	softwarerepositoryoperatingsystemfilerelationship["class_id"] = item.ClassId
	softwarerepositoryoperatingsystemfilerelationship["link"] = item.Link
	softwarerepositoryoperatingsystemfilerelationship["moid"] = item.Moid
	softwarerepositoryoperatingsystemfilerelationship["object_type"] = item.ObjectType
	softwarerepositoryoperatingsystemfilerelationship["selector"] = item.Selector

	softwarerepositoryoperatingsystemfilerelationships = append(softwarerepositoryoperatingsystemfilerelationships, softwarerepositoryoperatingsystemfilerelationship)
	return softwarerepositoryoperatingsystemfilerelationships
}
func flattenMapStorageArrayControllerRelationship(p *models.StorageArrayControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagearraycontrollerrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageArrayControllerRelationshipInterface.(*models.MoMoRef)
	storagearraycontrollerrelationship := make(map[string]interface{})
	storagearraycontrollerrelationship["class_id"] = item.ClassId
	storagearraycontrollerrelationship["link"] = item.Link
	storagearraycontrollerrelationship["moid"] = item.Moid
	storagearraycontrollerrelationship["object_type"] = item.ObjectType
	storagearraycontrollerrelationship["selector"] = item.Selector

	storagearraycontrollerrelationships = append(storagearraycontrollerrelationships, storagearraycontrollerrelationship)
	return storagearraycontrollerrelationships
}
func flattenMapStorageCapacity(p *models.StorageCapacity, d *schema.ResourceData) []map[string]interface{} {
	var storagecapacitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagecapacity := make(map[string]interface{})
	storagecapacity["available"] = item.Available
	storagecapacity["class_id"] = item.ClassId
	storagecapacity["free"] = item.Free
	storagecapacity["object_type"] = item.ObjectType
	storagecapacity["total"] = item.Total
	storagecapacity["used"] = item.Used

	storagecapacitys = append(storagecapacitys, storagecapacity)
	return storagecapacitys
}
func flattenMapStorageControllerRelationship(p *models.StorageControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageControllerRelationshipInterface.(*models.MoMoRef)
	storagecontrollerrelationship := make(map[string]interface{})
	storagecontrollerrelationship["class_id"] = item.ClassId
	storagecontrollerrelationship["link"] = item.Link
	storagecontrollerrelationship["moid"] = item.Moid
	storagecontrollerrelationship["object_type"] = item.ObjectType
	storagecontrollerrelationship["selector"] = item.Selector

	storagecontrollerrelationships = append(storagecontrollerrelationships, storagecontrollerrelationship)
	return storagecontrollerrelationships
}
func flattenMapStorageEnclosureRelationship(p *models.StorageEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageEnclosureRelationshipInterface.(*models.MoMoRef)
	storageenclosurerelationship := make(map[string]interface{})
	storageenclosurerelationship["class_id"] = item.ClassId
	storageenclosurerelationship["link"] = item.Link
	storageenclosurerelationship["moid"] = item.Moid
	storageenclosurerelationship["object_type"] = item.ObjectType
	storageenclosurerelationship["selector"] = item.Selector

	storageenclosurerelationships = append(storageenclosurerelationships, storageenclosurerelationship)
	return storageenclosurerelationships
}
func flattenMapStorageFlexFlashControllerRelationship(p *models.StorageFlexFlashControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageFlexFlashControllerRelationshipInterface.(*models.MoMoRef)
	storageflexflashcontrollerrelationship := make(map[string]interface{})
	storageflexflashcontrollerrelationship["class_id"] = item.ClassId
	storageflexflashcontrollerrelationship["link"] = item.Link
	storageflexflashcontrollerrelationship["moid"] = item.Moid
	storageflexflashcontrollerrelationship["object_type"] = item.ObjectType
	storageflexflashcontrollerrelationship["selector"] = item.Selector

	storageflexflashcontrollerrelationships = append(storageflexflashcontrollerrelationships, storageflexflashcontrollerrelationship)
	return storageflexflashcontrollerrelationships
}
func flattenMapStorageFlexUtilControllerRelationship(p *models.StorageFlexUtilControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageFlexUtilControllerRelationshipInterface.(*models.MoMoRef)
	storageflexutilcontrollerrelationship := make(map[string]interface{})
	storageflexutilcontrollerrelationship["class_id"] = item.ClassId
	storageflexutilcontrollerrelationship["link"] = item.Link
	storageflexutilcontrollerrelationship["moid"] = item.Moid
	storageflexutilcontrollerrelationship["object_type"] = item.ObjectType
	storageflexutilcontrollerrelationship["selector"] = item.Selector

	storageflexutilcontrollerrelationships = append(storageflexutilcontrollerrelationships, storageflexutilcontrollerrelationship)
	return storageflexutilcontrollerrelationships
}
func flattenMapStorageGenericArrayRelationship(p *models.StorageGenericArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagegenericarrayrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageGenericArrayRelationshipInterface.(*models.MoMoRef)
	storagegenericarrayrelationship := make(map[string]interface{})
	storagegenericarrayrelationship["class_id"] = item.ClassId
	storagegenericarrayrelationship["link"] = item.Link
	storagegenericarrayrelationship["moid"] = item.Moid
	storagegenericarrayrelationship["object_type"] = item.ObjectType
	storagegenericarrayrelationship["selector"] = item.Selector

	storagegenericarrayrelationships = append(storagegenericarrayrelationships, storagegenericarrayrelationship)
	return storagegenericarrayrelationships
}
func flattenMapStorageHostRelationship(p *models.StorageHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagehostrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageHostRelationshipInterface.(*models.MoMoRef)
	storagehostrelationship := make(map[string]interface{})
	storagehostrelationship["class_id"] = item.ClassId
	storagehostrelationship["link"] = item.Link
	storagehostrelationship["moid"] = item.Moid
	storagehostrelationship["object_type"] = item.ObjectType
	storagehostrelationship["selector"] = item.Selector

	storagehostrelationships = append(storagehostrelationships, storagehostrelationship)
	return storagehostrelationships
}
func flattenMapStorageHostUtilization(p *models.StorageHostUtilization, d *schema.ResourceData) []map[string]interface{} {
	var storagehostutilizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagehostutilization := make(map[string]interface{})
	storagehostutilization["available"] = item.Available
	storagehostutilization["class_id"] = item.ClassId
	storagehostutilization["data_reduction"] = item.DataReduction
	storagehostutilization["free"] = item.Free
	storagehostutilization["object_type"] = item.ObjectType
	storagehostutilization["snapshot"] = item.Snapshot
	storagehostutilization["thin_provisioned"] = item.ThinProvisioned
	storagehostutilization["total"] = item.Total
	storagehostutilization["total_reduction"] = item.TotalReduction
	storagehostutilization["used"] = item.Used
	storagehostutilization["volume"] = item.Volume

	storagehostutilizations = append(storagehostutilizations, storagehostutilization)
	return storagehostutilizations
}
func flattenMapStoragePhysicalDiskRelationship(p *models.StoragePhysicalDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StoragePhysicalDiskRelationshipInterface.(*models.MoMoRef)
	storagephysicaldiskrelationship := make(map[string]interface{})
	storagephysicaldiskrelationship["class_id"] = item.ClassId
	storagephysicaldiskrelationship["link"] = item.Link
	storagephysicaldiskrelationship["moid"] = item.Moid
	storagephysicaldiskrelationship["object_type"] = item.ObjectType
	storagephysicaldiskrelationship["selector"] = item.Selector

	storagephysicaldiskrelationships = append(storagephysicaldiskrelationships, storagephysicaldiskrelationship)
	return storagephysicaldiskrelationships
}
func flattenMapStorageProtectionGroupRelationship(p *models.StorageProtectionGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageprotectiongrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageProtectionGroupRelationshipInterface.(*models.MoMoRef)
	storageprotectiongrouprelationship := make(map[string]interface{})
	storageprotectiongrouprelationship["class_id"] = item.ClassId
	storageprotectiongrouprelationship["link"] = item.Link
	storageprotectiongrouprelationship["moid"] = item.Moid
	storageprotectiongrouprelationship["object_type"] = item.ObjectType
	storageprotectiongrouprelationship["selector"] = item.Selector

	storageprotectiongrouprelationships = append(storageprotectiongrouprelationships, storageprotectiongrouprelationship)
	return storageprotectiongrouprelationships
}
func flattenMapStorageProtectionGroupSnapshotRelationship(p *models.StorageProtectionGroupSnapshotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageprotectiongroupsnapshotrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageProtectionGroupSnapshotRelationshipInterface.(*models.MoMoRef)
	storageprotectiongroupsnapshotrelationship := make(map[string]interface{})
	storageprotectiongroupsnapshotrelationship["class_id"] = item.ClassId
	storageprotectiongroupsnapshotrelationship["link"] = item.Link
	storageprotectiongroupsnapshotrelationship["moid"] = item.Moid
	storageprotectiongroupsnapshotrelationship["object_type"] = item.ObjectType
	storageprotectiongroupsnapshotrelationship["selector"] = item.Selector

	storageprotectiongroupsnapshotrelationships = append(storageprotectiongroupsnapshotrelationships, storageprotectiongroupsnapshotrelationship)
	return storageprotectiongroupsnapshotrelationships
}
func flattenMapStoragePureHostGroupRelationship(p *models.StoragePureHostGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StoragePureHostGroupRelationshipInterface.(*models.MoMoRef)
	storagepurehostgrouprelationship := make(map[string]interface{})
	storagepurehostgrouprelationship["class_id"] = item.ClassId
	storagepurehostgrouprelationship["link"] = item.Link
	storagepurehostgrouprelationship["moid"] = item.Moid
	storagepurehostgrouprelationship["object_type"] = item.ObjectType
	storagepurehostgrouprelationship["selector"] = item.Selector

	storagepurehostgrouprelationships = append(storagepurehostgrouprelationships, storagepurehostgrouprelationship)
	return storagepurehostgrouprelationships
}
func flattenMapStoragePureProtectionGroupRelationship(p *models.StoragePureProtectionGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepureprotectiongrouprelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StoragePureProtectionGroupRelationshipInterface.(*models.MoMoRef)
	storagepureprotectiongrouprelationship := make(map[string]interface{})
	storagepureprotectiongrouprelationship["class_id"] = item.ClassId
	storagepureprotectiongrouprelationship["link"] = item.Link
	storagepureprotectiongrouprelationship["moid"] = item.Moid
	storagepureprotectiongrouprelationship["object_type"] = item.ObjectType
	storagepureprotectiongrouprelationship["selector"] = item.Selector

	storagepureprotectiongrouprelationships = append(storagepureprotectiongrouprelationships, storagepureprotectiongrouprelationship)
	return storagepureprotectiongrouprelationships
}
func flattenMapStorageSasExpanderRelationship(p *models.StorageSasExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageSasExpanderRelationshipInterface.(*models.MoMoRef)
	storagesasexpanderrelationship := make(map[string]interface{})
	storagesasexpanderrelationship["class_id"] = item.ClassId
	storagesasexpanderrelationship["link"] = item.Link
	storagesasexpanderrelationship["moid"] = item.Moid
	storagesasexpanderrelationship["object_type"] = item.ObjectType
	storagesasexpanderrelationship["selector"] = item.Selector

	storagesasexpanderrelationships = append(storagesasexpanderrelationships, storagesasexpanderrelationship)
	return storagesasexpanderrelationships
}
func flattenMapStorageVirtualDriveRelationship(p *models.StorageVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageVirtualDriveRelationshipInterface.(*models.MoMoRef)
	storagevirtualdriverelationship := make(map[string]interface{})
	storagevirtualdriverelationship["class_id"] = item.ClassId
	storagevirtualdriverelationship["link"] = item.Link
	storagevirtualdriverelationship["moid"] = item.Moid
	storagevirtualdriverelationship["object_type"] = item.ObjectType
	storagevirtualdriverelationship["selector"] = item.Selector

	storagevirtualdriverelationships = append(storagevirtualdriverelationships, storagevirtualdriverelationship)
	return storagevirtualdriverelationships
}
func flattenMapStorageVirtualDriveExtensionRelationship(p *models.StorageVirtualDriveExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageVirtualDriveExtensionRelationshipInterface.(*models.MoMoRef)
	storagevirtualdriveextensionrelationship := make(map[string]interface{})
	storagevirtualdriveextensionrelationship["class_id"] = item.ClassId
	storagevirtualdriveextensionrelationship["link"] = item.Link
	storagevirtualdriveextensionrelationship["moid"] = item.Moid
	storagevirtualdriveextensionrelationship["object_type"] = item.ObjectType
	storagevirtualdriveextensionrelationship["selector"] = item.Selector

	storagevirtualdriveextensionrelationships = append(storagevirtualdriveextensionrelationships, storagevirtualdriveextensionrelationship)
	return storagevirtualdriveextensionrelationships
}
func flattenMapStorageVolumeRelationship(p *models.StorageVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevolumerelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.StorageVolumeRelationshipInterface.(*models.MoMoRef)
	storagevolumerelationship := make(map[string]interface{})
	storagevolumerelationship["class_id"] = item.ClassId
	storagevolumerelationship["link"] = item.Link
	storagevolumerelationship["moid"] = item.Moid
	storagevolumerelationship["object_type"] = item.ObjectType
	storagevolumerelationship["selector"] = item.Selector

	storagevolumerelationships = append(storagevolumerelationships, storagevolumerelationship)
	return storagevolumerelationships
}
func flattenMapTamAdvisoryRelationship(p *models.TamAdvisoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var tamadvisoryrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.TamAdvisoryRelationshipInterface.(*models.MoMoRef)
	tamadvisoryrelationship := make(map[string]interface{})
	tamadvisoryrelationship["class_id"] = item.ClassId
	tamadvisoryrelationship["link"] = item.Link
	tamadvisoryrelationship["moid"] = item.Moid
	tamadvisoryrelationship["object_type"] = item.ObjectType
	tamadvisoryrelationship["selector"] = item.Selector

	tamadvisoryrelationships = append(tamadvisoryrelationships, tamadvisoryrelationship)
	return tamadvisoryrelationships
}
func flattenMapTamSeverity(p *models.TamSeverity, d *schema.ResourceData) []map[string]interface{} {
	var tamseveritys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	tamseverity := make(map[string]interface{})
	tamseverity["class_id"] = item.ClassId
	tamseverity["object_type"] = item.ObjectType

	tamseveritys = append(tamseveritys, tamseverity)
	return tamseveritys
}
func flattenMapTopSystemRelationship(p *models.TopSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var topsystemrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.TopSystemRelationshipInterface.(*models.MoMoRef)
	topsystemrelationship := make(map[string]interface{})
	topsystemrelationship["class_id"] = item.ClassId
	topsystemrelationship["link"] = item.Link
	topsystemrelationship["moid"] = item.Moid
	topsystemrelationship["object_type"] = item.ObjectType
	topsystemrelationship["selector"] = item.Selector

	topsystemrelationships = append(topsystemrelationships, topsystemrelationship)
	return topsystemrelationships
}
func flattenMapVirtualizationComputeCapacity(p *models.VirtualizationComputeCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcomputecapacitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationcomputecapacity := make(map[string]interface{})
	virtualizationcomputecapacity["capacity"] = item.Capacity
	virtualizationcomputecapacity["class_id"] = item.ClassId
	virtualizationcomputecapacity["free"] = item.Free
	virtualizationcomputecapacity["object_type"] = item.ObjectType
	virtualizationcomputecapacity["used"] = item.Used

	virtualizationcomputecapacitys = append(virtualizationcomputecapacitys, virtualizationcomputecapacity)
	return virtualizationcomputecapacitys
}
func flattenMapVirtualizationCpuInfo(p *models.VirtualizationCpuInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcpuinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationcpuinfo := make(map[string]interface{})
	virtualizationcpuinfo["class_id"] = item.ClassId
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
	virtualizationguestinfo["class_id"] = item.ClassId
	virtualizationguestinfo["hostname"] = item.Hostname
	virtualizationguestinfo["ip_address"] = item.IpAddress
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
	virtualizationmemorycapacity["capacity"] = item.Capacity
	virtualizationmemorycapacity["class_id"] = item.ClassId
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
	virtualizationproductinfo["class_id"] = item.ClassId
	virtualizationproductinfo["object_type"] = item.ObjectType
	virtualizationproductinfo["product_name"] = item.ProductName
	virtualizationproductinfo["product_type"] = item.ProductType
	virtualizationproductinfo["product_vendor"] = item.ProductVendor
	virtualizationproductinfo["version"] = item.Version

	virtualizationproductinfos = append(virtualizationproductinfos, virtualizationproductinfo)
	return virtualizationproductinfos
}
func flattenMapVirtualizationRemoteDisplayInfo(p *models.VirtualizationRemoteDisplayInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationremotedisplayinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationremotedisplayinfo := make(map[string]interface{})
	virtualizationremotedisplayinfo["class_id"] = item.ClassId
	virtualizationremotedisplayinfo["object_type"] = item.ObjectType
	virtualizationremotedisplayinfo["remote_display_password"] = item.RemoteDisplayPassword
	virtualizationremotedisplayinfo["remote_display_vnc_key"] = item.RemoteDisplayVncKey
	virtualizationremotedisplayinfo["remote_display_vnc_port"] = item.RemoteDisplayVncPort

	virtualizationremotedisplayinfos = append(virtualizationremotedisplayinfos, virtualizationremotedisplayinfo)
	return virtualizationremotedisplayinfos
}
func flattenMapVirtualizationResourceConsumption(p *models.VirtualizationResourceConsumption, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationresourceconsumptions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationresourceconsumption := make(map[string]interface{})
	virtualizationresourceconsumption["class_id"] = item.ClassId
	virtualizationresourceconsumption["cpu_consumed"] = item.CpuConsumed
	virtualizationresourceconsumption["memory_consumed"] = item.MemoryConsumed
	virtualizationresourceconsumption["object_type"] = item.ObjectType

	virtualizationresourceconsumptions = append(virtualizationresourceconsumptions, virtualizationresourceconsumption)
	return virtualizationresourceconsumptions
}
func flattenMapVirtualizationStorageCapacity(p *models.VirtualizationStorageCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationstoragecapacitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationstoragecapacity := make(map[string]interface{})
	virtualizationstoragecapacity["capacity"] = item.Capacity
	virtualizationstoragecapacity["class_id"] = item.ClassId
	virtualizationstoragecapacity["free"] = item.Free
	virtualizationstoragecapacity["object_type"] = item.ObjectType
	virtualizationstoragecapacity["used"] = item.Used

	virtualizationstoragecapacitys = append(virtualizationstoragecapacitys, virtualizationstoragecapacity)
	return virtualizationstoragecapacitys
}
func flattenMapVirtualizationVmCpuShareInfo(p *models.VirtualizationVmCpuShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmcpushareinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmcpushareinfo := make(map[string]interface{})
	virtualizationvmcpushareinfo["class_id"] = item.ClassId
	virtualizationvmcpushareinfo["cpu_limit"] = item.CpuLimit
	virtualizationvmcpushareinfo["cpu_overhead_limit"] = item.CpuOverheadLimit
	virtualizationvmcpushareinfo["cpu_reservation"] = item.CpuReservation
	virtualizationvmcpushareinfo["cpu_shares"] = item.CpuShares
	virtualizationvmcpushareinfo["object_type"] = item.ObjectType

	virtualizationvmcpushareinfos = append(virtualizationvmcpushareinfos, virtualizationvmcpushareinfo)
	return virtualizationvmcpushareinfos
}
func flattenMapVirtualizationVmCpuSocketInfo(p *models.VirtualizationVmCpuSocketInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmcpusocketinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmcpusocketinfo := make(map[string]interface{})
	virtualizationvmcpusocketinfo["class_id"] = item.ClassId
	virtualizationvmcpusocketinfo["cores_per_socket"] = item.CoresPerSocket
	virtualizationvmcpusocketinfo["num_cpus"] = item.NumCpus
	virtualizationvmcpusocketinfo["num_sockets"] = item.NumSockets
	virtualizationvmcpusocketinfo["object_type"] = item.ObjectType

	virtualizationvmcpusocketinfos = append(virtualizationvmcpusocketinfos, virtualizationvmcpusocketinfo)
	return virtualizationvmcpusocketinfos
}
func flattenMapVirtualizationVmDiskCommitInfo(p *models.VirtualizationVmDiskCommitInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmdiskcommitinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmdiskcommitinfo := make(map[string]interface{})
	virtualizationvmdiskcommitinfo["class_id"] = item.ClassId
	virtualizationvmdiskcommitinfo["committed_disk"] = item.CommittedDisk
	virtualizationvmdiskcommitinfo["object_type"] = item.ObjectType
	virtualizationvmdiskcommitinfo["un_committed_disk"] = item.UnCommittedDisk
	virtualizationvmdiskcommitinfo["unshared_disk"] = item.UnsharedDisk

	virtualizationvmdiskcommitinfos = append(virtualizationvmdiskcommitinfos, virtualizationvmdiskcommitinfo)
	return virtualizationvmdiskcommitinfos
}
func flattenMapVirtualizationVmMemoryShareInfo(p *models.VirtualizationVmMemoryShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmmemoryshareinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualizationvmmemoryshareinfo := make(map[string]interface{})
	virtualizationvmmemoryshareinfo["class_id"] = item.ClassId
	virtualizationvmmemoryshareinfo["mem_limit"] = item.MemLimit
	virtualizationvmmemoryshareinfo["mem_overhead_limit"] = item.MemOverheadLimit
	virtualizationvmmemoryshareinfo["mem_reservation"] = item.MemReservation
	virtualizationvmmemoryshareinfo["mem_shares"] = item.MemShares
	virtualizationvmmemoryshareinfo["object_type"] = item.ObjectType

	virtualizationvmmemoryshareinfos = append(virtualizationvmmemoryshareinfos, virtualizationvmmemoryshareinfo)
	return virtualizationvmmemoryshareinfos
}
func flattenMapVirtualizationVmwareClusterRelationship(p *models.VirtualizationVmwareClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareclusterrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VirtualizationVmwareClusterRelationshipInterface.(*models.MoMoRef)
	virtualizationvmwareclusterrelationship := make(map[string]interface{})
	virtualizationvmwareclusterrelationship["class_id"] = item.ClassId
	virtualizationvmwareclusterrelationship["link"] = item.Link
	virtualizationvmwareclusterrelationship["moid"] = item.Moid
	virtualizationvmwareclusterrelationship["object_type"] = item.ObjectType
	virtualizationvmwareclusterrelationship["selector"] = item.Selector

	virtualizationvmwareclusterrelationships = append(virtualizationvmwareclusterrelationships, virtualizationvmwareclusterrelationship)
	return virtualizationvmwareclusterrelationships
}
func flattenMapVirtualizationVmwareDatacenterRelationship(p *models.VirtualizationVmwareDatacenterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatacenterrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VirtualizationVmwareDatacenterRelationshipInterface.(*models.MoMoRef)
	virtualizationvmwaredatacenterrelationship := make(map[string]interface{})
	virtualizationvmwaredatacenterrelationship["class_id"] = item.ClassId
	virtualizationvmwaredatacenterrelationship["link"] = item.Link
	virtualizationvmwaredatacenterrelationship["moid"] = item.Moid
	virtualizationvmwaredatacenterrelationship["object_type"] = item.ObjectType
	virtualizationvmwaredatacenterrelationship["selector"] = item.Selector

	virtualizationvmwaredatacenterrelationships = append(virtualizationvmwaredatacenterrelationships, virtualizationvmwaredatacenterrelationship)
	return virtualizationvmwaredatacenterrelationships
}
func flattenMapVirtualizationVmwareHostRelationship(p *models.VirtualizationVmwareHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VirtualizationVmwareHostRelationshipInterface.(*models.MoMoRef)
	virtualizationvmwarehostrelationship := make(map[string]interface{})
	virtualizationvmwarehostrelationship["class_id"] = item.ClassId
	virtualizationvmwarehostrelationship["link"] = item.Link
	virtualizationvmwarehostrelationship["moid"] = item.Moid
	virtualizationvmwarehostrelationship["object_type"] = item.ObjectType
	virtualizationvmwarehostrelationship["selector"] = item.Selector

	virtualizationvmwarehostrelationships = append(virtualizationvmwarehostrelationships, virtualizationvmwarehostrelationship)
	return virtualizationvmwarehostrelationships
}
func flattenMapVirtualizationVmwareVcenterRelationship(p *models.VirtualizationVmwareVcenterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevcenterrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VirtualizationVmwareVcenterRelationshipInterface.(*models.MoMoRef)
	virtualizationvmwarevcenterrelationship := make(map[string]interface{})
	virtualizationvmwarevcenterrelationship["class_id"] = item.ClassId
	virtualizationvmwarevcenterrelationship["link"] = item.Link
	virtualizationvmwarevcenterrelationship["moid"] = item.Moid
	virtualizationvmwarevcenterrelationship["object_type"] = item.ObjectType
	virtualizationvmwarevcenterrelationship["selector"] = item.Selector

	virtualizationvmwarevcenterrelationships = append(virtualizationvmwarevcenterrelationships, virtualizationvmwarevcenterrelationship)
	return virtualizationvmwarevcenterrelationships
}
func flattenMapVnicArfsSettings(p *models.VnicArfsSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicarfssettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicarfssettings := make(map[string]interface{})
	vnicarfssettings["class_id"] = item.ClassId
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
	vniccdn["class_id"] = item.ClassId
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
	vniccompletionqueuesettings["class_id"] = item.ClassId
	vniccompletionqueuesettings["count"] = item.Count
	vniccompletionqueuesettings["object_type"] = item.ObjectType
	vniccompletionqueuesettings["ring_size"] = item.RingSize

	vniccompletionqueuesettingss = append(vniccompletionqueuesettingss, vniccompletionqueuesettings)
	return vniccompletionqueuesettingss
}
func flattenMapVnicEthAdapterPolicyRelationship(p *models.VnicEthAdapterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethadapterpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicEthAdapterPolicyRelationshipInterface.(*models.MoMoRef)
	vnicethadapterpolicyrelationship := make(map[string]interface{})
	vnicethadapterpolicyrelationship["class_id"] = item.ClassId
	vnicethadapterpolicyrelationship["link"] = item.Link
	vnicethadapterpolicyrelationship["moid"] = item.Moid
	vnicethadapterpolicyrelationship["object_type"] = item.ObjectType
	vnicethadapterpolicyrelationship["selector"] = item.Selector

	vnicethadapterpolicyrelationships = append(vnicethadapterpolicyrelationships, vnicethadapterpolicyrelationship)
	return vnicethadapterpolicyrelationships
}
func flattenMapVnicEthInterruptSettings(p *models.VnicEthInterruptSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethinterruptsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethinterruptsettings := make(map[string]interface{})
	vnicethinterruptsettings["class_id"] = item.ClassId
	vnicethinterruptsettings["coalescing_time"] = item.CoalescingTime
	vnicethinterruptsettings["coalescing_type"] = item.CoalescingType
	vnicethinterruptsettings["count"] = item.Count
	vnicethinterruptsettings["mode"] = item.Mode
	vnicethinterruptsettings["object_type"] = item.ObjectType

	vnicethinterruptsettingss = append(vnicethinterruptsettingss, vnicethinterruptsettings)
	return vnicethinterruptsettingss
}
func flattenMapVnicEthNetworkPolicyRelationship(p *models.VnicEthNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethnetworkpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicEthNetworkPolicyRelationshipInterface.(*models.MoMoRef)
	vnicethnetworkpolicyrelationship := make(map[string]interface{})
	vnicethnetworkpolicyrelationship["class_id"] = item.ClassId
	vnicethnetworkpolicyrelationship["link"] = item.Link
	vnicethnetworkpolicyrelationship["moid"] = item.Moid
	vnicethnetworkpolicyrelationship["object_type"] = item.ObjectType
	vnicethnetworkpolicyrelationship["selector"] = item.Selector

	vnicethnetworkpolicyrelationships = append(vnicethnetworkpolicyrelationships, vnicethnetworkpolicyrelationship)
	return vnicethnetworkpolicyrelationships
}
func flattenMapVnicEthQosPolicyRelationship(p *models.VnicEthQosPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethqospolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicEthQosPolicyRelationshipInterface.(*models.MoMoRef)
	vnicethqospolicyrelationship := make(map[string]interface{})
	vnicethqospolicyrelationship["class_id"] = item.ClassId
	vnicethqospolicyrelationship["link"] = item.Link
	vnicethqospolicyrelationship["moid"] = item.Moid
	vnicethqospolicyrelationship["object_type"] = item.ObjectType
	vnicethqospolicyrelationship["selector"] = item.Selector

	vnicethqospolicyrelationships = append(vnicethqospolicyrelationships, vnicethqospolicyrelationship)
	return vnicethqospolicyrelationships
}
func flattenMapVnicEthRxQueueSettings(p *models.VnicEthRxQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethrxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicethrxqueuesettings := make(map[string]interface{})
	vnicethrxqueuesettings["class_id"] = item.ClassId
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
	vnicethtxqueuesettings["class_id"] = item.ClassId
	vnicethtxqueuesettings["count"] = item.Count
	vnicethtxqueuesettings["object_type"] = item.ObjectType
	vnicethtxqueuesettings["ring_size"] = item.RingSize

	vnicethtxqueuesettingss = append(vnicethtxqueuesettingss, vnicethtxqueuesettings)
	return vnicethtxqueuesettingss
}
func flattenMapVnicFcAdapterPolicyRelationship(p *models.VnicFcAdapterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcadapterpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicFcAdapterPolicyRelationshipInterface.(*models.MoMoRef)
	vnicfcadapterpolicyrelationship := make(map[string]interface{})
	vnicfcadapterpolicyrelationship["class_id"] = item.ClassId
	vnicfcadapterpolicyrelationship["link"] = item.Link
	vnicfcadapterpolicyrelationship["moid"] = item.Moid
	vnicfcadapterpolicyrelationship["object_type"] = item.ObjectType
	vnicfcadapterpolicyrelationship["selector"] = item.Selector

	vnicfcadapterpolicyrelationships = append(vnicfcadapterpolicyrelationships, vnicfcadapterpolicyrelationship)
	return vnicfcadapterpolicyrelationships
}
func flattenMapVnicFcErrorRecoverySettings(p *models.VnicFcErrorRecoverySettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcerrorrecoverysettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcerrorrecoverysettings := make(map[string]interface{})
	vnicfcerrorrecoverysettings["class_id"] = item.ClassId
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
	vnicfcinterruptsettings["class_id"] = item.ClassId
	vnicfcinterruptsettings["mode"] = item.Mode
	vnicfcinterruptsettings["object_type"] = item.ObjectType

	vnicfcinterruptsettingss = append(vnicfcinterruptsettingss, vnicfcinterruptsettings)
	return vnicfcinterruptsettingss
}
func flattenMapVnicFcNetworkPolicyRelationship(p *models.VnicFcNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcnetworkpolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicFcNetworkPolicyRelationshipInterface.(*models.MoMoRef)
	vnicfcnetworkpolicyrelationship := make(map[string]interface{})
	vnicfcnetworkpolicyrelationship["class_id"] = item.ClassId
	vnicfcnetworkpolicyrelationship["link"] = item.Link
	vnicfcnetworkpolicyrelationship["moid"] = item.Moid
	vnicfcnetworkpolicyrelationship["object_type"] = item.ObjectType
	vnicfcnetworkpolicyrelationship["selector"] = item.Selector

	vnicfcnetworkpolicyrelationships = append(vnicfcnetworkpolicyrelationships, vnicfcnetworkpolicyrelationship)
	return vnicfcnetworkpolicyrelationships
}
func flattenMapVnicFcQosPolicyRelationship(p *models.VnicFcQosPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqospolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicFcQosPolicyRelationshipInterface.(*models.MoMoRef)
	vnicfcqospolicyrelationship := make(map[string]interface{})
	vnicfcqospolicyrelationship["class_id"] = item.ClassId
	vnicfcqospolicyrelationship["link"] = item.Link
	vnicfcqospolicyrelationship["moid"] = item.Moid
	vnicfcqospolicyrelationship["object_type"] = item.ObjectType
	vnicfcqospolicyrelationship["selector"] = item.Selector

	vnicfcqospolicyrelationships = append(vnicfcqospolicyrelationships, vnicfcqospolicyrelationship)
	return vnicfcqospolicyrelationships
}
func flattenMapVnicFcQueueSettings(p *models.VnicFcQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicfcqueuesettings := make(map[string]interface{})
	vnicfcqueuesettings["class_id"] = item.ClassId
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
	vnicflogisettings["class_id"] = item.ClassId
	vnicflogisettings["object_type"] = item.ObjectType
	vnicflogisettings["retries"] = item.Retries
	vnicflogisettings["timeout"] = item.Timeout

	vnicflogisettingss = append(vnicflogisettingss, vnicflogisettings)
	return vnicflogisettingss
}
func flattenMapVnicLanConnectivityPolicyRelationship(p *models.VnicLanConnectivityPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vniclanconnectivitypolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicLanConnectivityPolicyRelationshipInterface.(*models.MoMoRef)
	vniclanconnectivitypolicyrelationship := make(map[string]interface{})
	vniclanconnectivitypolicyrelationship["class_id"] = item.ClassId
	vniclanconnectivitypolicyrelationship["link"] = item.Link
	vniclanconnectivitypolicyrelationship["moid"] = item.Moid
	vniclanconnectivitypolicyrelationship["object_type"] = item.ObjectType
	vniclanconnectivitypolicyrelationship["selector"] = item.Selector

	vniclanconnectivitypolicyrelationships = append(vniclanconnectivitypolicyrelationships, vniclanconnectivitypolicyrelationship)
	return vniclanconnectivitypolicyrelationships
}
func flattenMapVnicNvgreSettings(p *models.VnicNvgreSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicnvgresettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicnvgresettings := make(map[string]interface{})
	vnicnvgresettings["class_id"] = item.ClassId
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
	vnicplacementsettings["class_id"] = item.ClassId
	vnicplacementsettings["id"] = item.Id
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
	vnicplogisettings["class_id"] = item.ClassId
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
	vnicrocesettings["class_id"] = item.ClassId
	vnicrocesettings["enabled"] = item.Enabled
	vnicrocesettings["memory_regions"] = item.MemoryRegions
	vnicrocesettings["object_type"] = item.ObjectType
	vnicrocesettings["queue_pairs"] = item.QueuePairs
	vnicrocesettings["resource_groups"] = item.ResourceGroups

	vnicrocesettingss = append(vnicrocesettingss, vnicrocesettings)
	return vnicrocesettingss
}
func flattenMapVnicSanConnectivityPolicyRelationship(p *models.VnicSanConnectivityPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicsanconnectivitypolicyrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.VnicSanConnectivityPolicyRelationshipInterface.(*models.MoMoRef)
	vnicsanconnectivitypolicyrelationship := make(map[string]interface{})
	vnicsanconnectivitypolicyrelationship["class_id"] = item.ClassId
	vnicsanconnectivitypolicyrelationship["link"] = item.Link
	vnicsanconnectivitypolicyrelationship["moid"] = item.Moid
	vnicsanconnectivitypolicyrelationship["object_type"] = item.ObjectType
	vnicsanconnectivitypolicyrelationship["selector"] = item.Selector

	vnicsanconnectivitypolicyrelationships = append(vnicsanconnectivitypolicyrelationships, vnicsanconnectivitypolicyrelationship)
	return vnicsanconnectivitypolicyrelationships
}
func flattenMapVnicScsiQueueSettings(p *models.VnicScsiQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicscsiqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnicscsiqueuesettings := make(map[string]interface{})
	vnicscsiqueuesettings["class_id"] = item.ClassId
	vnicscsiqueuesettings["count"] = item.Count
	vnicscsiqueuesettings["object_type"] = item.ObjectType
	vnicscsiqueuesettings["ring_size"] = item.RingSize

	vnicscsiqueuesettingss = append(vnicscsiqueuesettingss, vnicscsiqueuesettings)
	return vnicscsiqueuesettingss
}
func flattenMapVnicTcpOffloadSettings(p *models.VnicTcpOffloadSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnictcpoffloadsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vnictcpoffloadsettings := make(map[string]interface{})
	vnictcpoffloadsettings["class_id"] = item.ClassId
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
	vnicusnicsettings["class_id"] = item.ClassId
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
	vnicvlansettings["class_id"] = item.ClassId
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
	vnicvmqsettings["class_id"] = item.ClassId
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
	vnicvsansettings["class_id"] = item.ClassId
	vnicvsansettings["id"] = item.Id
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
	vnicvxlansettings["class_id"] = item.ClassId
	vnicvxlansettings["enabled"] = item.Enabled
	vnicvxlansettings["object_type"] = item.ObjectType

	vnicvxlansettingss = append(vnicvxlansettingss, vnicvxlansettings)
	return vnicvxlansettingss
}
func flattenMapWorkflowCatalogRelationship(p *models.WorkflowCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowcatalogrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.WorkflowCatalogRelationshipInterface.(*models.MoMoRef)
	workflowcatalogrelationship := make(map[string]interface{})
	workflowcatalogrelationship["class_id"] = item.ClassId
	workflowcatalogrelationship["link"] = item.Link
	workflowcatalogrelationship["moid"] = item.Moid
	workflowcatalogrelationship["object_type"] = item.ObjectType
	workflowcatalogrelationship["selector"] = item.Selector

	workflowcatalogrelationships = append(workflowcatalogrelationships, workflowcatalogrelationship)
	return workflowcatalogrelationships
}
func flattenMapWorkflowInternalProperties(p *models.WorkflowInternalProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowinternalpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowinternalproperties := make(map[string]interface{})
	workflowinternalproperties["base_task_type"] = item.BaseTaskType
	workflowinternalproperties["class_id"] = item.ClassId
	workflowinternalproperties["constraints"] = (func(p *models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
		var workflowtaskconstraintss []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		workflowtaskconstraints := make(map[string]interface{})
		workflowtaskconstraints["class_id"] = item.ClassId
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
func flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(p *models.WorkflowPendingDynamicWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowpendingdynamicworkflowinforelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.WorkflowPendingDynamicWorkflowInfoRelationshipInterface.(*models.MoMoRef)
	workflowpendingdynamicworkflowinforelationship := make(map[string]interface{})
	workflowpendingdynamicworkflowinforelationship["class_id"] = item.ClassId
	workflowpendingdynamicworkflowinforelationship["link"] = item.Link
	workflowpendingdynamicworkflowinforelationship["moid"] = item.Moid
	workflowpendingdynamicworkflowinforelationship["object_type"] = item.ObjectType
	workflowpendingdynamicworkflowinforelationship["selector"] = item.Selector

	workflowpendingdynamicworkflowinforelationships = append(workflowpendingdynamicworkflowinforelationships, workflowpendingdynamicworkflowinforelationship)
	return workflowpendingdynamicworkflowinforelationships
}
func flattenMapWorkflowProperties(p *models.WorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowproperties := make(map[string]interface{})
	workflowproperties["class_id"] = item.ClassId
	workflowproperties["external_meta"] = item.ExternalMeta
	workflowproperties["input_definition"] = (func(p *[]models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if *p == nil {
			return nil
		}
		for _, item := range *p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["class_id"] = item.ClassId
			workflowbasedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["class_id"] = item.ClassId
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
	workflowproperties["output_definition"] = (func(p *[]models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if *p == nil {
			return nil
		}
		for _, item := range *p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["class_id"] = item.ClassId
			workflowbasedatatype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["class_id"] = item.ClassId
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
	workflowtaskconstraints["class_id"] = item.ClassId
	workflowtaskconstraints["object_type"] = item.ObjectType
	workflowtaskconstraints["target_data_type"] = item.TargetDataType

	workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
	return workflowtaskconstraintss
}
func flattenMapWorkflowTaskDefinitionRelationship(p *models.WorkflowTaskDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.WorkflowTaskDefinitionRelationshipInterface.(*models.MoMoRef)
	workflowtaskdefinitionrelationship := make(map[string]interface{})
	workflowtaskdefinitionrelationship["class_id"] = item.ClassId
	workflowtaskdefinitionrelationship["link"] = item.Link
	workflowtaskdefinitionrelationship["moid"] = item.Moid
	workflowtaskdefinitionrelationship["object_type"] = item.ObjectType
	workflowtaskdefinitionrelationship["selector"] = item.Selector

	workflowtaskdefinitionrelationships = append(workflowtaskdefinitionrelationships, workflowtaskdefinitionrelationship)
	return workflowtaskdefinitionrelationships
}
func flattenMapWorkflowTaskInfoRelationship(p *models.WorkflowTaskInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.WorkflowTaskInfoRelationshipInterface.(*models.MoMoRef)
	workflowtaskinforelationship := make(map[string]interface{})
	workflowtaskinforelationship["class_id"] = item.ClassId
	workflowtaskinforelationship["link"] = item.Link
	workflowtaskinforelationship["moid"] = item.Moid
	workflowtaskinforelationship["object_type"] = item.ObjectType
	workflowtaskinforelationship["selector"] = item.Selector

	workflowtaskinforelationships = append(workflowtaskinforelationships, workflowtaskinforelationship)
	return workflowtaskinforelationships
}
func flattenMapWorkflowValidationInformation(p *models.WorkflowValidationInformation, d *schema.ResourceData) []map[string]interface{} {
	var workflowvalidationinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowvalidationinformation := make(map[string]interface{})
	workflowvalidationinformation["class_id"] = item.ClassId
	workflowvalidationinformation["object_type"] = item.ObjectType
	workflowvalidationinformation["state"] = item.State
	workflowvalidationinformation["validation_error"] = (func(p *[]models.WorkflowValidationError, d *schema.ResourceData) []map[string]interface{} {
		var workflowvalidationerrors []map[string]interface{}
		if *p == nil {
			return nil
		}
		for _, item := range *p {
			workflowvalidationerror := make(map[string]interface{})
			workflowvalidationerror["class_id"] = item.ClassId
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
func flattenMapWorkflowWorkflowDefinitionRelationship(p *models.WorkflowWorkflowDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowdefinitionrelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.WorkflowWorkflowDefinitionRelationshipInterface.(*models.MoMoRef)
	workflowworkflowdefinitionrelationship := make(map[string]interface{})
	workflowworkflowdefinitionrelationship["class_id"] = item.ClassId
	workflowworkflowdefinitionrelationship["link"] = item.Link
	workflowworkflowdefinitionrelationship["moid"] = item.Moid
	workflowworkflowdefinitionrelationship["object_type"] = item.ObjectType
	workflowworkflowdefinitionrelationship["selector"] = item.Selector

	workflowworkflowdefinitionrelationships = append(workflowworkflowdefinitionrelationships, workflowworkflowdefinitionrelationship)
	return workflowworkflowdefinitionrelationships
}
func flattenMapWorkflowWorkflowInfoRelationship(p *models.WorkflowWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforelationships []map[string]interface{}
	if p == nil {
		return nil
	}
	x := *p
	item := x.WorkflowWorkflowInfoRelationshipInterface.(*models.MoMoRef)
	workflowworkflowinforelationship := make(map[string]interface{})
	workflowworkflowinforelationship["class_id"] = item.ClassId
	workflowworkflowinforelationship["link"] = item.Link
	workflowworkflowinforelationship["moid"] = item.Moid
	workflowworkflowinforelationship["object_type"] = item.ObjectType
	workflowworkflowinforelationship["selector"] = item.Selector

	workflowworkflowinforelationships = append(workflowworkflowinforelationships, workflowworkflowinforelationship)
	return workflowworkflowinforelationships
}
func flattenMapWorkflowWorkflowInfoProperties(p *models.WorkflowWorkflowInfoProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinfopropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowworkflowinfoproperties := make(map[string]interface{})
	workflowworkflowinfoproperties["class_id"] = item.ClassId
	workflowworkflowinfoproperties["object_type"] = item.ObjectType
	workflowworkflowinfoproperties["retryable"] = item.Retryable

	workflowworkflowinfopropertiess = append(workflowworkflowinfopropertiess, workflowworkflowinfoproperties)
	return workflowworkflowinfopropertiess
}
func flattenMapWorkflowWorkflowProperties(p *models.WorkflowWorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowworkflowproperties := make(map[string]interface{})
	workflowworkflowproperties["class_id"] = item.ClassId
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
	x509certificate["class_id"] = item.ClassId
	x509certificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		pkixdistinguishedname := make(map[string]interface{})
		pkixdistinguishedname["class_id"] = item.ClassId
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
		pkixdistinguishedname["class_id"] = item.ClassId
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
