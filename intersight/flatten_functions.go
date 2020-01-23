package intersight

import (
	"encoding/json"
	"log"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenListRecoveryBackupProfileRef(p []*models.RecoveryBackupProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var backupprofiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		backupprofiles := make(map[string]interface{})
		backupprofiles["moid"] = item.Moid
		backupprofiles["object_type"] = item.ObjectType
		backupprofiles["selector"] = item.Selector
		backupprofiless = append(backupprofiless, backupprofiles)
	}
	return backupprofiless
}
func flattenMapOrganizationOrganizationRef(p *models.OrganizationOrganizationRef, d *schema.ResourceData) []map[string]interface{} {

	var organizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	organization := make(map[string]interface{})
	organization["moid"] = item.Moid
	organization["object_type"] = item.ObjectType
	organization["selector"] = item.Selector

	organizations = append(organizations, organization)
	return organizations
}
func flattenListMoBaseMoRef(p []*models.MoBaseMoRef, d *schema.ResourceData) []map[string]interface{} {
	var permissionresourcess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		permissionresources := make(map[string]interface{})
		permissionresources["moid"] = item.Moid
		permissionresources["object_type"] = item.ObjectType
		permissionresources["selector"] = item.Selector
		permissionresourcess = append(permissionresourcess, permissionresources)
	}
	return permissionresourcess
}
func flattenMapRecoveryBackupSchedule(p *models.RecoveryBackupSchedule, d *schema.ResourceData) []map[string]interface{} {

	var schedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	schedule := make(map[string]interface{})
	delete(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1, "ObjectType")
	if len(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1) != 0 {
		j, err := json.Marshal(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1)
		if err == nil {
			schedule["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	schedule["execution_time"] = item.ExecutionTime
	schedule["frequency_unit"] = item.FrequencyUnit
	schedule["hours"] = item.Hours
	schedule["object_type"] = item.ObjectType

	schedules = append(schedules, schedule)
	return schedules
}
func flattenListMoTag(p []*models.MoTag, d *schema.ResourceData) []map[string]interface{} {
	var tagss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		tags := make(map[string]interface{})
		delete(item.MoTagAO1P1.MoTagAO1P1, "ObjectType")
		if len(item.MoTagAO1P1.MoTagAO1P1) != 0 {

			j, err := json.Marshal(item.MoTagAO1P1.MoTagAO1P1)
			if err == nil {
				tags["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		tags["key"] = item.Key
		tags["object_type"] = item.ObjectType
		tags["value"] = item.Value
		tagss = append(tagss, tags)
	}
	return tagss
}
func flattenListHyperflexClusterProfileRef(p []*models.HyperflexClusterProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var clusterprofiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		clusterprofiles := make(map[string]interface{})
		clusterprofiles["moid"] = item.Moid
		clusterprofiles["object_type"] = item.ObjectType
		clusterprofiles["selector"] = item.Selector
		clusterprofiless = append(clusterprofiless, clusterprofiles)
	}
	return clusterprofiless
}
func flattenMapHyperflexAppCatalogRef(p *models.HyperflexAppCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var appcatalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	appcatalog := make(map[string]interface{})
	appcatalog["moid"] = item.Moid
	appcatalog["object_type"] = item.ObjectType
	appcatalog["selector"] = item.Selector

	appcatalogs = append(appcatalogs, appcatalog)
	return appcatalogs
}
func flattenListHyperflexServerModelEntry(p []*models.HyperflexServerModelEntry, d *schema.ResourceData) []map[string]interface{} {
	var servermodelentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		servermodelentries := make(map[string]interface{})
		delete(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1, "ObjectType")
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				servermodelentries["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		servermodelentries["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {

			var constraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			constraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					constraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			constraint["hxdp_version"] = item.HxdpVersion
			constraint["hypervisor_type"] = item.HypervisorType
			constraint["mgmt_platform"] = item.MgmtPlatform
			constraint["object_type"] = item.ObjectType
			constraint["server_model"] = item.ServerModel

			constraints = append(constraints, constraint)
			return constraints
		})(item.Constraint, d)
		servermodelentries["name"] = item.Name
		servermodelentries["object_type"] = item.ObjectType
		servermodelentries["value"] = item.Value
		servermodelentriess = append(servermodelentriess, servermodelentries)
	}
	return servermodelentriess
}
func flattenListIaasConnectorPackRef(p []*models.IaasConnectorPackRef, d *schema.ResourceData) []map[string]interface{} {
	var connectorpacks []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		connectorpack := make(map[string]interface{})
		connectorpack["moid"] = item.Moid
		connectorpack["object_type"] = item.ObjectType
		connectorpack["selector"] = item.Selector
		connectorpacks = append(connectorpacks, connectorpack)
	}
	return connectorpacks
}
func flattenListIaasDeviceStatusRef(p []*models.IaasDeviceStatusRef, d *schema.ResourceData) []map[string]interface{} {
	var devicestatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		devicestatus := make(map[string]interface{})
		devicestatus["moid"] = item.Moid
		devicestatus["object_type"] = item.ObjectType
		devicestatus["selector"] = item.Selector
		devicestatuss = append(devicestatuss, devicestatus)
	}
	return devicestatuss
}
func flattenMapIaasLicenseInfoRef(p *models.IaasLicenseInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var licenseinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	licenseinfo := make(map[string]interface{})
	licenseinfo["moid"] = item.Moid
	licenseinfo["object_type"] = item.ObjectType
	licenseinfo["selector"] = item.Selector

	licenseinfos = append(licenseinfos, licenseinfo)
	return licenseinfos
}
func flattenListIaasMostRunTasksRef(p []*models.IaasMostRunTasksRef, d *schema.ResourceData) []map[string]interface{} {
	var mostruntaskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		mostruntasks := make(map[string]interface{})
		mostruntasks["moid"] = item.Moid
		mostruntasks["object_type"] = item.ObjectType
		mostruntasks["selector"] = item.Selector
		mostruntaskss = append(mostruntaskss, mostruntasks)
	}
	return mostruntaskss
}
func flattenMapAssetDeviceRegistrationRef(p *models.AssetDeviceRegistrationRef, d *schema.ResourceData) []map[string]interface{} {

	var registereddevices []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	registereddevice := make(map[string]interface{})
	registereddevice["moid"] = item.Moid
	registereddevice["object_type"] = item.ObjectType
	registereddevice["selector"] = item.Selector

	registereddevices = append(registereddevices, registereddevice)
	return registereddevices
}
func flattenMapIaasUcsdManagedInfraRef(p *models.IaasUcsdManagedInfraRef, d *schema.ResourceData) []map[string]interface{} {

	var ucsdmanagedinfras []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ucsdmanagedinfra := make(map[string]interface{})
	ucsdmanagedinfra["moid"] = item.Moid
	ucsdmanagedinfra["object_type"] = item.ObjectType
	ucsdmanagedinfra["selector"] = item.Selector

	ucsdmanagedinfras = append(ucsdmanagedinfras, ucsdmanagedinfra)
	return ucsdmanagedinfras
}
func flattenMapTamAdvisoryRef(p *models.TamAdvisoryRef, d *schema.ResourceData) []map[string]interface{} {

	var advisorys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	advisory := make(map[string]interface{})
	advisory["moid"] = item.Moid
	advisory["object_type"] = item.ObjectType
	advisory["selector"] = item.Selector

	advisorys = append(advisorys, advisory)
	return advisorys
}
func flattenMapMoBaseMoRef(p *models.MoBaseMoRef, d *schema.ResourceData) []map[string]interface{} {

	var affectedobjects []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	affectedobject := make(map[string]interface{})
	affectedobject["moid"] = item.Moid
	affectedobject["object_type"] = item.ObjectType
	affectedobject["selector"] = item.Selector

	affectedobjects = append(affectedobjects, affectedobject)
	return affectedobjects
}
func flattenMapIamAccountRef(p *models.IamAccountRef, d *schema.ResourceData) []map[string]interface{} {

	var accounts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	account := make(map[string]interface{})
	account["moid"] = item.Moid
	account["object_type"] = item.ObjectType
	account["selector"] = item.Selector

	accounts = append(accounts, account)
	return accounts
}
func flattenListWorkflowMessage(p []*models.WorkflowMessage, d *schema.ResourceData) []map[string]interface{} {
	var messages []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		message := make(map[string]interface{})
		delete(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1, "ObjectType")
		if len(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1)
			if err == nil {
				message["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		message["message"] = item.Message
		message["object_type"] = item.ObjectType
		message["severity"] = item.Severity
		messages = append(messages, message)
	}
	return messages
}
func flattenMapHyperflexClusterProfileRef(p *models.HyperflexClusterProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var nr0clusterprofiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nr0clusterprofile := make(map[string]interface{})
	nr0clusterprofile["moid"] = item.Moid
	nr0clusterprofile["object_type"] = item.ObjectType
	nr0clusterprofile["selector"] = item.Selector

	nr0clusterprofiles = append(nr0clusterprofiles, nr0clusterprofile)
	return nr0clusterprofiles
}
func flattenMapServerProfileRef(p *models.ServerProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var nr1profiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nr1profile := make(map[string]interface{})
	nr1profile["moid"] = item.Moid
	nr1profile["object_type"] = item.ObjectType
	nr1profile["selector"] = item.Selector

	nr1profiles = append(nr1profiles, nr1profile)
	return nr1profiles
}
func flattenMapWorkflowTaskInfoRef(p *models.WorkflowTaskInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var parenttaskinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	parenttaskinfo := make(map[string]interface{})
	parenttaskinfo["moid"] = item.Moid
	parenttaskinfo["object_type"] = item.ObjectType
	parenttaskinfo["selector"] = item.Selector

	parenttaskinfos = append(parenttaskinfos, parenttaskinfo)
	return parenttaskinfos
}
func flattenMapWorkflowPendingDynamicWorkflowInfoRef(p *models.WorkflowPendingDynamicWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var pendingdynamicworkflowinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	pendingdynamicworkflowinfo := make(map[string]interface{})
	pendingdynamicworkflowinfo["moid"] = item.Moid
	pendingdynamicworkflowinfo["object_type"] = item.ObjectType
	pendingdynamicworkflowinfo["selector"] = item.Selector

	pendingdynamicworkflowinfos = append(pendingdynamicworkflowinfos, pendingdynamicworkflowinfo)
	return pendingdynamicworkflowinfos
}
func flattenMapIamPermissionRef(p *models.IamPermissionRef, d *schema.ResourceData) []map[string]interface{} {

	var permissions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	permission := make(map[string]interface{})
	permission["moid"] = item.Moid
	permission["object_type"] = item.ObjectType
	permission["selector"] = item.Selector

	permissions = append(permissions, permission)
	return permissions
}
func flattenListWorkflowTaskInfoRef(p []*models.WorkflowTaskInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var taskinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		taskinfos := make(map[string]interface{})
		taskinfos["moid"] = item.Moid
		taskinfos["object_type"] = item.ObjectType
		taskinfos["selector"] = item.Selector
		taskinfoss = append(taskinfoss, taskinfos)
	}
	return taskinfoss
}
func flattenMapWorkflowWorkflowDefinitionRef(p *models.WorkflowWorkflowDefinitionRef, d *schema.ResourceData) []map[string]interface{} {

	var workflowdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowdefinition := make(map[string]interface{})
	workflowdefinition["moid"] = item.Moid
	workflowdefinition["object_type"] = item.ObjectType
	workflowdefinition["selector"] = item.Selector

	workflowdefinitions = append(workflowdefinitions, workflowdefinition)
	return workflowdefinitions
}
func flattenMapVnicFcErrorRecoverySettings(p *models.VnicFcErrorRecoverySettings, d *schema.ResourceData) []map[string]interface{} {

	var errorrecoverysettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	errorrecoverysettings := make(map[string]interface{})
	delete(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1, "ObjectType")
	if len(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1)
		if err == nil {
			errorrecoverysettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	errorrecoverysettings["enabled"] = item.Enabled
	errorrecoverysettings["io_retry_count"] = item.IoRetryCount
	errorrecoverysettings["io_retry_timeout"] = item.IoRetryTimeout
	errorrecoverysettings["link_down_timeout"] = item.LinkDownTimeout
	errorrecoverysettings["object_type"] = item.ObjectType
	errorrecoverysettings["port_down_timeout"] = item.PortDownTimeout

	errorrecoverysettingss = append(errorrecoverysettingss, errorrecoverysettings)
	return errorrecoverysettingss
}
func flattenMapVnicFlogiSettings(p *models.VnicFlogiSettings, d *schema.ResourceData) []map[string]interface{} {

	var flogisettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	flogisettings := make(map[string]interface{})
	delete(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1, "ObjectType")
	if len(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1)
		if err == nil {
			flogisettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	flogisettings["object_type"] = item.ObjectType
	flogisettings["retries"] = item.Retries
	flogisettings["timeout"] = item.Timeout

	flogisettingss = append(flogisettingss, flogisettings)
	return flogisettingss
}
func flattenMapVnicFcInterruptSettings(p *models.VnicFcInterruptSettings, d *schema.ResourceData) []map[string]interface{} {

	var interruptsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	interruptsettings := make(map[string]interface{})
	delete(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1, "ObjectType")
	if len(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1)
		if err == nil {
			interruptsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	interruptsettings["mode"] = item.Mode
	interruptsettings["object_type"] = item.ObjectType

	interruptsettingss = append(interruptsettingss, interruptsettings)
	return interruptsettingss
}
func flattenMapVnicPlogiSettings(p *models.VnicPlogiSettings, d *schema.ResourceData) []map[string]interface{} {

	var plogisettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	plogisettings := make(map[string]interface{})
	delete(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1, "ObjectType")
	if len(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1)
		if err == nil {
			plogisettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	plogisettings["object_type"] = item.ObjectType
	plogisettings["retries"] = item.Retries
	plogisettings["timeout"] = item.Timeout

	plogisettingss = append(plogisettingss, plogisettings)
	return plogisettingss
}
func flattenMapVnicFcQueueSettings(p *models.VnicFcQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var rxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	rxqueuesettings := make(map[string]interface{})
	delete(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1)
		if err == nil {
			rxqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	rxqueuesettings["count"] = item.Count
	rxqueuesettings["object_type"] = item.ObjectType
	rxqueuesettings["ring_size"] = item.RingSize

	rxqueuesettingss = append(rxqueuesettingss, rxqueuesettings)
	return rxqueuesettingss
}
func flattenMapVnicScsiQueueSettings(p *models.VnicScsiQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var scsiqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	scsiqueuesettings := make(map[string]interface{})
	delete(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1)
		if err == nil {
			scsiqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	scsiqueuesettings["count"] = item.Count
	scsiqueuesettings["object_type"] = item.ObjectType
	scsiqueuesettings["ring_size"] = item.RingSize

	scsiqueuesettingss = append(scsiqueuesettingss, scsiqueuesettings)
	return scsiqueuesettingss
}
func flattenMapHyperflexIPAddrRange(p *models.HyperflexIPAddrRange, d *schema.ResourceData) []map[string]interface{} {

	var kvmipranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	kvmiprange := make(map[string]interface{})
	delete(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1, "ObjectType")
	if len(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1)
		if err == nil {
			kvmiprange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	kvmiprange["end_addr"] = item.EndAddr
	kvmiprange["gateway"] = item.Gateway
	kvmiprange["netmask"] = item.Netmask
	kvmiprange["object_type"] = item.ObjectType
	kvmiprange["start_addr"] = item.StartAddr

	kvmipranges = append(kvmipranges, kvmiprange)
	return kvmipranges
}
func flattenMapHyperflexMacAddrPrefixRange(p *models.HyperflexMacAddrPrefixRange, d *schema.ResourceData) []map[string]interface{} {

	var macprefixranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	macprefixrange := make(map[string]interface{})
	delete(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1, "ObjectType")
	if len(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1)
		if err == nil {
			macprefixrange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	macprefixrange["end_addr"] = item.EndAddr
	macprefixrange["object_type"] = item.ObjectType
	macprefixrange["start_addr"] = item.StartAddr

	macprefixranges = append(macprefixranges, macprefixrange)
	return macprefixranges
}
func flattenListStorageDiskGroupPolicyRef(p []*models.StorageDiskGroupPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var diskgrouppoliciess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		diskgrouppolicies := make(map[string]interface{})
		diskgrouppolicies["moid"] = item.Moid
		diskgrouppolicies["object_type"] = item.ObjectType
		diskgrouppolicies["selector"] = item.Selector
		diskgrouppoliciess = append(diskgrouppoliciess, diskgrouppolicies)
	}
	return diskgrouppoliciess
}
func flattenListStorageLocalDisk(p []*models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
	var globalhotsparess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		globalhotspares := make(map[string]interface{})
		delete(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1, "ObjectType")
		if len(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1) != 0 {

			j, err := json.Marshal(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1)
			if err == nil {
				globalhotspares["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		globalhotspares["object_type"] = item.ObjectType
		globalhotspares["slot_number"] = item.SlotNumber
		globalhotsparess = append(globalhotsparess, globalhotspares)
	}
	return globalhotsparess
}
func flattenListPolicyAbstractConfigProfileRef(p []*models.PolicyAbstractConfigProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var profiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		profiles := make(map[string]interface{})
		profiles["moid"] = item.Moid
		profiles["object_type"] = item.ObjectType
		profiles["selector"] = item.Selector
		profiless = append(profiless, profiles)
	}
	return profiless
}
func flattenListStorageVirtualDriveConfig(p []*models.StorageVirtualDriveConfig, d *schema.ResourceData) []map[string]interface{} {
	var virtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		virtualdrives := make(map[string]interface{})
		virtualdrives["access_policy"] = item.AccessPolicy
		delete(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1, "ObjectType")
		if len(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1) != 0 {

			j, err := json.Marshal(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1)
			if err == nil {
				virtualdrives["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		virtualdrives["boot_drive"] = item.BootDrive
		virtualdrives["disk_group_name"] = item.DiskGroupName
		virtualdrives["disk_group_policy"] = item.DiskGroupPolicy
		virtualdrives["drive_cache"] = item.DriveCache
		virtualdrives["expand_to_available"] = item.ExpandToAvailable
		virtualdrives["io_policy"] = item.IoPolicy
		virtualdrives["name"] = item.Name
		virtualdrives["object_type"] = item.ObjectType
		virtualdrives["read_policy"] = item.ReadPolicy
		virtualdrives["size"] = item.Size
		virtualdrives["write_policy"] = item.WritePolicy
		virtualdrivess = append(virtualdrivess, virtualdrives)
	}
	return virtualdrivess
}
func flattenMapComputeRackUnitRef(p *models.ComputeRackUnitRef, d *schema.ResourceData) []map[string]interface{} {

	var assignedservers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	assignedserver := make(map[string]interface{})
	assignedserver["moid"] = item.Moid
	assignedserver["object_type"] = item.ObjectType
	assignedserver["selector"] = item.Selector

	assignedservers = append(assignedservers, assignedserver)
	return assignedservers
}
func flattenListServerConfigChangeDetailRef(p []*models.ServerConfigChangeDetailRef, d *schema.ResourceData) []map[string]interface{} {
	var configchangedetailss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		configchangedetails := make(map[string]interface{})
		configchangedetails["moid"] = item.Moid
		configchangedetails["object_type"] = item.ObjectType
		configchangedetails["selector"] = item.Selector
		configchangedetailss = append(configchangedetailss, configchangedetails)
	}
	return configchangedetailss
}
func flattenMapPolicyConfigChange(p *models.PolicyConfigChange, d *schema.ResourceData) []map[string]interface{} {

	var configchangess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configchanges := make(map[string]interface{})
	delete(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1, "ObjectType")
	if len(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1) != 0 {
		j, err := json.Marshal(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1)
		if err == nil {
			configchanges["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	configchanges["changes"] = item.Changes
	configchanges["disruptions"] = item.Disruptions
	configchanges["object_type"] = item.ObjectType

	configchangess = append(configchangess, configchanges)
	return configchangess
}
func flattenMapPolicyConfigContext(p *models.PolicyConfigContext, d *schema.ResourceData) []map[string]interface{} {

	var configcontexts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configcontext := make(map[string]interface{})
	delete(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1, "ObjectType")
	if len(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1) != 0 {
		j, err := json.Marshal(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1)
		if err == nil {
			configcontext["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	configcontext["config_state"] = item.ConfigState
	configcontext["control_action"] = item.ControlAction
	configcontext["error_state"] = item.ErrorState
	configcontext["object_type"] = item.ObjectType
	configcontext["oper_state"] = item.OperState

	configcontexts = append(configcontexts, configcontext)
	return configcontexts
}
func flattenMapServerConfigResultRef(p *models.ServerConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var configresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configresult := make(map[string]interface{})
	configresult["moid"] = item.Moid
	configresult["object_type"] = item.ObjectType
	configresult["selector"] = item.Selector

	configresults = append(configresults, configresult)
	return configresults
}
func flattenListWorkflowWorkflowInfoRef(p []*models.WorkflowWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var runningworkflowss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		runningworkflows := make(map[string]interface{})
		runningworkflows["moid"] = item.Moid
		runningworkflows["object_type"] = item.ObjectType
		runningworkflows["selector"] = item.Selector
		runningworkflowss = append(runningworkflowss, runningworkflows)
	}
	return runningworkflowss
}
func flattenMapPolicyAbstractProfileRef(p *models.PolicyAbstractProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var srctemplates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	srctemplate := make(map[string]interface{})
	srctemplate["moid"] = item.Moid
	srctemplate["object_type"] = item.ObjectType
	srctemplate["selector"] = item.Selector

	srctemplates = append(srctemplates, srctemplate)
	return srctemplates
}
func flattenMapIamLdapBaseProperties(p *models.IamLdapBaseProperties, d *schema.ResourceData) []map[string]interface{} {

	var basepropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	baseproperties := make(map[string]interface{})
	delete(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1, "ObjectType")
	if len(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1)
		if err == nil {
			baseproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	baseproperties["attribute"] = item.Attribute
	baseproperties["base_dn"] = item.BaseDn
	baseproperties["bind_dn"] = item.BindDn
	baseproperties["bind_method"] = item.BindMethod
	baseproperties["domain"] = item.Domain
	baseproperties["enable_encryption"] = item.EnableEncryption
	baseproperties["enable_group_authorization"] = item.EnableGroupAuthorization
	baseproperties["filter"] = item.Filter
	baseproperties["group_attribute"] = item.GroupAttribute
	baseproperties["is_password_set"] = item.IsPasswordSet
	baseproperties["nested_group_search_depth"] = item.NestedGroupSearchDepth
	baseproperties["object_type"] = item.ObjectType
	password_x := d.Get("base_properties").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	baseproperties["password"] = password_y["password"]
	baseproperties["timeout"] = item.Timeout

	basepropertiess = append(basepropertiess, baseproperties)
	return basepropertiess
}
func flattenMapIamLdapDNSParameters(p *models.IamLdapDNSParameters, d *schema.ResourceData) []map[string]interface{} {

	var dnsparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	dnsparameters := make(map[string]interface{})
	delete(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1, "ObjectType")
	if len(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1) != 0 {
		j, err := json.Marshal(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1)
		if err == nil {
			dnsparameters["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	dnsparameters["object_type"] = item.ObjectType
	dnsparameters["search_domain"] = item.SearchDomain
	dnsparameters["search_forest"] = item.SearchForest
	dnsparameters["source"] = item.Source

	dnsparameterss = append(dnsparameterss, dnsparameters)
	return dnsparameterss
}
func flattenListIamLdapGroupRef(p []*models.IamLdapGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var groupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		groups := make(map[string]interface{})
		groups["moid"] = item.Moid
		groups["object_type"] = item.ObjectType
		groups["selector"] = item.Selector
		groupss = append(groupss, groups)
	}
	return groupss
}
func flattenListIamLdapProviderRef(p []*models.IamLdapProviderRef, d *schema.ResourceData) []map[string]interface{} {
	var providerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		providers := make(map[string]interface{})
		providers["moid"] = item.Moid
		providers["object_type"] = item.ObjectType
		providers["selector"] = item.Selector
		providerss = append(providerss, providers)
	}
	return providerss
}
func flattenMapRecoveryBackupConfigPolicyRef(p *models.RecoveryBackupConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var backupconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	backupconfig := make(map[string]interface{})
	backupconfig["moid"] = item.Moid
	backupconfig["object_type"] = item.ObjectType
	backupconfig["selector"] = item.Selector

	backupconfigs = append(backupconfigs, backupconfig)
	return backupconfigs
}
func flattenMapRecoveryConfigResultRef(p *models.RecoveryConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var configresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configresult := make(map[string]interface{})
	configresult["moid"] = item.Moid
	configresult["object_type"] = item.ObjectType
	configresult["selector"] = item.Selector

	configresults = append(configresults, configresult)
	return configresults
}
func flattenMapRecoveryScheduleConfigPolicyRef(p *models.RecoveryScheduleConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var scheduleconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	scheduleconfig := make(map[string]interface{})
	scheduleconfig["moid"] = item.Moid
	scheduleconfig["object_type"] = item.ObjectType
	scheduleconfig["selector"] = item.Selector

	scheduleconfigs = append(scheduleconfigs, scheduleconfig)
	return scheduleconfigs
}
func flattenListWorkflowAPI(p []*models.WorkflowAPI, d *schema.ResourceData) []map[string]interface{} {
	var batchs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		batch := make(map[string]interface{})
		delete(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1, "ObjectType")
		if len(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1)
			if err == nil {
				batch["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		batch["body"] = item.Body
		batch["content_type"] = item.ContentType
		batch["name"] = item.Name
		batch["object_type"] = item.ObjectType
		batch["outcomes"] = item.Outcomes
		batch["response_spec"] = (func(p *models.ContentGrammar, d *schema.ResourceData) []map[string]interface{} {

			var responsespecs []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			responsespec := make(map[string]interface{})
			delete(item.ContentGrammarAO1P1.ContentGrammarAO1P1, "ObjectType")
			if len(item.ContentGrammarAO1P1.ContentGrammarAO1P1) != 0 {
				j, err := json.Marshal(item.ContentGrammarAO1P1.ContentGrammarAO1P1)
				if err == nil {
					responsespec["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			responsespec["error_parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var errorparameterss []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					errorparameters := make(map[string]interface{})
					errorparameters["accept_single_value"] = item.AcceptSingleValue
					errorparameters["complex_type"] = item.ComplexType
					errorparameters["item_type"] = item.ItemType
					errorparameters["name"] = item.Name
					errorparameters["object_type"] = item.ObjectType
					errorparameters["path"] = item.Path
					errorparameters["type"] = item.Type
					errorparameterss = append(errorparameterss, errorparameters)
				}
				return errorparameterss
			})(item.ErrorParameters, d)
			responsespec["object_type"] = item.ObjectType
			responsespec["parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var parameterss []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					parameters := make(map[string]interface{})
					parameters["accept_single_value"] = item.AcceptSingleValue
					parameters["complex_type"] = item.ComplexType
					parameters["item_type"] = item.ItemType
					parameters["name"] = item.Name
					parameters["object_type"] = item.ObjectType
					parameters["path"] = item.Path
					parameters["type"] = item.Type
					parameterss = append(parameterss, parameters)
				}
				return parameterss
			})(item.Parameters, d)
			responsespec["types"] = (func(p []*models.ContentComplexType, d *schema.ResourceData) []map[string]interface{} {
				var typess []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					types := make(map[string]interface{})
					types["name"] = item.Name
					types["object_type"] = item.ObjectType
					types["parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
						var parameterss []map[string]interface{}
						if p == nil {
							return nil
						}
						for _, item := range p {
							item := *item
							parameters := make(map[string]interface{})
							parameters["accept_single_value"] = item.AcceptSingleValue
							parameters["complex_type"] = item.ComplexType
							parameters["item_type"] = item.ItemType
							parameters["name"] = item.Name
							parameters["object_type"] = item.ObjectType
							parameters["path"] = item.Path
							parameters["type"] = item.Type
							parameterss = append(parameterss, parameters)
						}
						return parameterss
					})(item.Parameters, d)
					typess = append(typess, types)
				}
				return typess
			})(item.Types, d)

			responsespecs = append(responsespecs, responsespec)
			return responsespecs
		})(item.ResponseSpec, d)
		batch["skip_on_condition"] = item.SkipOnCondition
		batch["timeout"] = item.Timeout
		batchs = append(batchs, batch)
	}
	return batchs
}
func flattenMapWorkflowTaskDefinitionRef(p *models.WorkflowTaskDefinitionRef, d *schema.ResourceData) []map[string]interface{} {

	var taskdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	taskdefinition := make(map[string]interface{})
	taskdefinition["moid"] = item.Moid
	taskdefinition["object_type"] = item.ObjectType
	taskdefinition["selector"] = item.Selector

	taskdefinitions = append(taskdefinitions, taskdefinition)
	return taskdefinitions
}
func flattenMapSoftwarerepositoryCatalogRef(p *models.SoftwarerepositoryCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var catalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	catalog := make(map[string]interface{})
	catalog["moid"] = item.Moid
	catalog["object_type"] = item.ObjectType
	catalog["selector"] = item.Selector

	catalogs = append(catalogs, catalog)
	return catalogs
}
func flattenMapSoftwarerepositoryFileServer(p *models.SoftwarerepositoryFileServer, d *schema.ResourceData) []map[string]interface{} {

	var sources []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	source := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			source["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	source["object_type"] = item.ObjectType

	sources = append(sources, source)
	return sources
}
func flattenListIamEndPointRoleRef(p []*models.IamEndPointRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var endpointroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		endpointroles := make(map[string]interface{})
		endpointroles["moid"] = item.Moid
		endpointroles["object_type"] = item.ObjectType
		endpointroles["selector"] = item.Selector
		endpointroless = append(endpointroless, endpointroles)
	}
	return endpointroless
}
func flattenListIamRoleRef(p []*models.IamRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var roless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		roles := make(map[string]interface{})
		roles["moid"] = item.Moid
		roles["object_type"] = item.ObjectType
		roles["selector"] = item.Selector
		roless = append(roless, roles)
	}
	return roless
}
func flattenMapVnicCdn(p *models.VnicCdn, d *schema.ResourceData) []map[string]interface{} {

	var cdns []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	cdn := make(map[string]interface{})
	delete(item.VnicCdnAO1P1.VnicCdnAO1P1, "ObjectType")
	if len(item.VnicCdnAO1P1.VnicCdnAO1P1) != 0 {
		j, err := json.Marshal(item.VnicCdnAO1P1.VnicCdnAO1P1)
		if err == nil {
			cdn["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	cdn["object_type"] = item.ObjectType
	cdn["source"] = item.Source
	cdn["value"] = item.Value

	cdns = append(cdns, cdn)
	return cdns
}
func flattenMapVnicEthAdapterPolicyRef(p *models.VnicEthAdapterPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var ethadapterpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ethadapterpolicy := make(map[string]interface{})
	ethadapterpolicy["moid"] = item.Moid
	ethadapterpolicy["object_type"] = item.ObjectType
	ethadapterpolicy["selector"] = item.Selector

	ethadapterpolicys = append(ethadapterpolicys, ethadapterpolicy)
	return ethadapterpolicys
}
func flattenMapVnicEthNetworkPolicyRef(p *models.VnicEthNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var ethnetworkpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ethnetworkpolicy := make(map[string]interface{})
	ethnetworkpolicy["moid"] = item.Moid
	ethnetworkpolicy["object_type"] = item.ObjectType
	ethnetworkpolicy["selector"] = item.Selector

	ethnetworkpolicys = append(ethnetworkpolicys, ethnetworkpolicy)
	return ethnetworkpolicys
}
func flattenMapVnicEthQosPolicyRef(p *models.VnicEthQosPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var ethqospolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ethqospolicy := make(map[string]interface{})
	ethqospolicy["moid"] = item.Moid
	ethqospolicy["object_type"] = item.ObjectType
	ethqospolicy["selector"] = item.Selector

	ethqospolicys = append(ethqospolicys, ethqospolicy)
	return ethqospolicys
}
func flattenMapVnicLanConnectivityPolicyRef(p *models.VnicLanConnectivityPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var lanconnectivitypolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	lanconnectivitypolicy := make(map[string]interface{})
	lanconnectivitypolicy["moid"] = item.Moid
	lanconnectivitypolicy["object_type"] = item.ObjectType
	lanconnectivitypolicy["selector"] = item.Selector

	lanconnectivitypolicys = append(lanconnectivitypolicys, lanconnectivitypolicy)
	return lanconnectivitypolicys
}
func flattenMapVnicPlacementSettings(p *models.VnicPlacementSettings, d *schema.ResourceData) []map[string]interface{} {

	var placements []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	placement := make(map[string]interface{})
	delete(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1, "ObjectType")
	if len(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1)
		if err == nil {
			placement["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	placement["id"] = item.ID
	placement["object_type"] = item.ObjectType
	placement["pci_link"] = item.PciLink
	placement["uplink"] = item.Uplink

	placements = append(placements, placement)
	return placements
}
func flattenMapVnicUsnicSettings(p *models.VnicUsnicSettings, d *schema.ResourceData) []map[string]interface{} {

	var usnicsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	usnicsettings := make(map[string]interface{})
	delete(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1, "ObjectType")
	if len(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1)
		if err == nil {
			usnicsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	usnicsettings["cos"] = item.Cos
	usnicsettings["count"] = item.Count
	usnicsettings["object_type"] = item.ObjectType
	usnicsettings["usnic_adapter_policy"] = item.UsnicAdapterPolicy

	usnicsettingss = append(usnicsettingss, usnicsettings)
	return usnicsettingss
}
func flattenMapVnicVmqSettings(p *models.VnicVmqSettings, d *schema.ResourceData) []map[string]interface{} {

	var vmqsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vmqsettings := make(map[string]interface{})
	delete(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1, "ObjectType")
	if len(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1)
		if err == nil {
			vmqsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vmqsettings["enabled"] = item.Enabled
	vmqsettings["object_type"] = item.ObjectType

	vmqsettingss = append(vmqsettingss, vmqsettings)
	return vmqsettingss
}
func flattenListIamEndPointUserRoleRef(p []*models.IamEndPointUserRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var endpointuserroles []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		endpointuserrole := make(map[string]interface{})
		endpointuserrole["moid"] = item.Moid
		endpointuserrole["object_type"] = item.ObjectType
		endpointuserrole["selector"] = item.Selector
		endpointuserroles = append(endpointuserroles, endpointuserrole)
	}
	return endpointuserroles
}
func flattenMapIamUserGroupRef(p *models.IamUserGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var usergroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	usergroup := make(map[string]interface{})
	usergroup["moid"] = item.Moid
	usergroup["object_type"] = item.ObjectType
	usergroup["selector"] = item.Selector

	usergroups = append(usergroups, usergroup)
	return usergroups
}
func flattenMapIamEndPointUserRef(p *models.IamEndPointUserRef, d *schema.ResourceData) []map[string]interface{} {

	var endpointusers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	endpointuser := make(map[string]interface{})
	endpointuser["moid"] = item.Moid
	endpointuser["object_type"] = item.ObjectType
	endpointuser["selector"] = item.Selector

	endpointusers = append(endpointusers, endpointuser)
	return endpointusers
}
func flattenMapIamEndPointUserPolicyRef(p *models.IamEndPointUserPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var endpointuserpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	endpointuserpolicy := make(map[string]interface{})
	endpointuserpolicy["moid"] = item.Moid
	endpointuserpolicy["object_type"] = item.ObjectType
	endpointuserpolicy["selector"] = item.Selector

	endpointuserpolicys = append(endpointuserpolicys, endpointuserpolicy)
	return endpointuserpolicys
}
func flattenListHyperflexFeatureLimitEntry(p []*models.HyperflexFeatureLimitEntry, d *schema.ResourceData) []map[string]interface{} {
	var featurelimitentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		featurelimitentries := make(map[string]interface{})
		delete(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1, "ObjectType")
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				featurelimitentries["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		featurelimitentries["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {

			var constraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			constraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					constraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			constraint["hxdp_version"] = item.HxdpVersion
			constraint["hypervisor_type"] = item.HypervisorType
			constraint["mgmt_platform"] = item.MgmtPlatform
			constraint["object_type"] = item.ObjectType
			constraint["server_model"] = item.ServerModel

			constraints = append(constraints, constraint)
			return constraints
		})(item.Constraint, d)
		featurelimitentries["name"] = item.Name
		featurelimitentries["object_type"] = item.ObjectType
		featurelimitentries["value"] = item.Value
		featurelimitentriess = append(featurelimitentriess, featurelimitentries)
	}
	return featurelimitentriess
}
func flattenListVnicFcIfRef(p []*models.VnicFcIfRef, d *schema.ResourceData) []map[string]interface{} {
	var fcifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		fcifs := make(map[string]interface{})
		fcifs["moid"] = item.Moid
		fcifs["object_type"] = item.ObjectType
		fcifs["selector"] = item.Selector
		fcifss = append(fcifss, fcifs)
	}
	return fcifss
}
func flattenMapIamUserRef(p *models.IamUserRef, d *schema.ResourceData) []map[string]interface{} {

	var claimedbyusers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	claimedbyuser := make(map[string]interface{})
	claimedbyuser["moid"] = item.Moid
	claimedbyuser["object_type"] = item.ObjectType
	claimedbyuser["selector"] = item.Selector

	claimedbyusers = append(claimedbyusers, claimedbyuser)
	return claimedbyusers
}
func flattenListAssetClusterMemberRef(p []*models.AssetClusterMemberRef, d *schema.ResourceData) []map[string]interface{} {
	var clustermemberss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		clustermembers := make(map[string]interface{})
		clustermembers["moid"] = item.Moid
		clustermembers["object_type"] = item.ObjectType
		clustermembers["selector"] = item.Selector
		clustermemberss = append(clustermemberss, clustermembers)
	}
	return clustermemberss
}
func flattenMapAssetDeviceClaimRef(p *models.AssetDeviceClaimRef, d *schema.ResourceData) []map[string]interface{} {

	var deviceclaims []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	deviceclaim := make(map[string]interface{})
	deviceclaim["moid"] = item.Moid
	deviceclaim["object_type"] = item.ObjectType
	deviceclaim["selector"] = item.Selector

	deviceclaims = append(deviceclaims, deviceclaim)
	return deviceclaims
}
func flattenMapAssetDeviceConfigurationRef(p *models.AssetDeviceConfigurationRef, d *schema.ResourceData) []map[string]interface{} {

	var deviceconfigurations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	deviceconfiguration := make(map[string]interface{})
	deviceconfiguration["moid"] = item.Moid
	deviceconfiguration["object_type"] = item.ObjectType
	deviceconfiguration["selector"] = item.Selector

	deviceconfigurations = append(deviceconfigurations, deviceconfiguration)
	return deviceconfigurations
}
func flattenMapIamDomainGroupRef(p *models.IamDomainGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var domaingroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	domaingroup := make(map[string]interface{})
	domaingroup["moid"] = item.Moid
	domaingroup["object_type"] = item.ObjectType
	domaingroup["selector"] = item.Selector

	domaingroups = append(domaingroups, domaingroup)
	return domaingroups
}
func flattenMapAssetParentConnectionSignature(p *models.AssetParentConnectionSignature, d *schema.ResourceData) []map[string]interface{} {

	var parentsignatures []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	parentsignature := make(map[string]interface{})
	delete(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1, "ObjectType")
	if len(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1) != 0 {
		j, err := json.Marshal(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1)
		if err == nil {
			parentsignature["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	parentsignature["device_id"] = item.DeviceID
	parentsignature["node_id"] = item.NodeID
	parentsignature["object_type"] = item.ObjectType
	parentsignature["signature"] = item.Signature

	parentsignatures = append(parentsignatures, parentsignature)
	return parentsignatures
}
func flattenMapAssetSecurityTokenRef(p *models.AssetSecurityTokenRef, d *schema.ResourceData) []map[string]interface{} {

	var securitytokens []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	securitytoken := make(map[string]interface{})
	securitytoken["moid"] = item.Moid
	securitytoken["object_type"] = item.ObjectType
	securitytoken["selector"] = item.Selector

	securitytokens = append(securitytokens, securitytoken)
	return securitytokens
}
func flattenListVmediaMapping(p []*models.VmediaMapping, d *schema.ResourceData) []map[string]interface{} {
	var mappingss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		mappings := make(map[string]interface{})
		delete(item.VmediaMappingAO1P1.VmediaMappingAO1P1, "ObjectType")
		if len(item.VmediaMappingAO1P1.VmediaMappingAO1P1) != 0 {

			j, err := json.Marshal(item.VmediaMappingAO1P1.VmediaMappingAO1P1)
			if err == nil {
				mappings["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		mappings["authentication_protocol"] = item.AuthenticationProtocol
		mappings["device_type"] = item.DeviceType
		mappings["host_name"] = item.HostName
		mappings["is_password_set"] = item.IsPasswordSet
		mappings["mount_options"] = item.MountOptions
		mappings["mount_protocol"] = item.MountProtocol
		mappings["object_type"] = item.ObjectType
		mappings["password"] = item.Password
		mappings["remote_file"] = item.RemoteFile
		mappings["remote_path"] = item.RemotePath
		mappings["username"] = item.Username
		mappings["volume_name"] = item.VolumeName
		mappingss = append(mappingss, mappings)
	}
	return mappingss
}
func flattenMapVnicArfsSettings(p *models.VnicArfsSettings, d *schema.ResourceData) []map[string]interface{} {

	var arfssettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	arfssettings := make(map[string]interface{})
	delete(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1, "ObjectType")
	if len(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1)
		if err == nil {
			arfssettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	arfssettings["enabled"] = item.Enabled
	arfssettings["object_type"] = item.ObjectType

	arfssettingss = append(arfssettingss, arfssettings)
	return arfssettingss
}
func flattenMapVnicCompletionQueueSettings(p *models.VnicCompletionQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var completionqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	completionqueuesettings := make(map[string]interface{})
	delete(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1)
		if err == nil {
			completionqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	completionqueuesettings["count"] = item.Count
	completionqueuesettings["object_type"] = item.ObjectType
	completionqueuesettings["ring_size"] = item.RingSize

	completionqueuesettingss = append(completionqueuesettingss, completionqueuesettings)
	return completionqueuesettingss
}
func flattenMapVnicEthInterruptSettings(p *models.VnicEthInterruptSettings, d *schema.ResourceData) []map[string]interface{} {

	var interruptsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	interruptsettings := make(map[string]interface{})
	delete(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1, "ObjectType")
	if len(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1)
		if err == nil {
			interruptsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	interruptsettings["coalescing_time"] = item.CoalescingTime
	interruptsettings["coalescing_type"] = item.CoalescingType
	interruptsettings["count"] = item.Count
	interruptsettings["mode"] = item.Mode
	interruptsettings["object_type"] = item.ObjectType

	interruptsettingss = append(interruptsettingss, interruptsettings)
	return interruptsettingss
}
func flattenMapVnicNvgreSettings(p *models.VnicNvgreSettings, d *schema.ResourceData) []map[string]interface{} {

	var nvgresettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nvgresettings := make(map[string]interface{})
	delete(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1, "ObjectType")
	if len(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1)
		if err == nil {
			nvgresettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	nvgresettings["enabled"] = item.Enabled
	nvgresettings["object_type"] = item.ObjectType

	nvgresettingss = append(nvgresettingss, nvgresettings)
	return nvgresettingss
}
func flattenMapVnicRoceSettings(p *models.VnicRoceSettings, d *schema.ResourceData) []map[string]interface{} {

	var rocesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	rocesettings := make(map[string]interface{})
	delete(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1, "ObjectType")
	if len(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1)
		if err == nil {
			rocesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	rocesettings["enabled"] = item.Enabled
	rocesettings["memory_regions"] = item.MemoryRegions
	rocesettings["object_type"] = item.ObjectType
	rocesettings["queue_pairs"] = item.QueuePairs
	rocesettings["resource_groups"] = item.ResourceGroups

	rocesettingss = append(rocesettingss, rocesettings)
	return rocesettingss
}
func flattenMapVnicEthRxQueueSettings(p *models.VnicEthRxQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var rxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	rxqueuesettings := make(map[string]interface{})
	delete(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1)
		if err == nil {
			rxqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	rxqueuesettings["count"] = item.Count
	rxqueuesettings["object_type"] = item.ObjectType
	rxqueuesettings["ring_size"] = item.RingSize

	rxqueuesettingss = append(rxqueuesettingss, rxqueuesettings)
	return rxqueuesettingss
}
func flattenMapVnicTCPOffloadSettings(p *models.VnicTCPOffloadSettings, d *schema.ResourceData) []map[string]interface{} {

	var tcpoffloadsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	tcpoffloadsettings := make(map[string]interface{})
	delete(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1, "ObjectType")
	if len(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1)
		if err == nil {
			tcpoffloadsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	tcpoffloadsettings["large_receive"] = item.LargeReceive
	tcpoffloadsettings["large_send"] = item.LargeSend
	tcpoffloadsettings["object_type"] = item.ObjectType
	tcpoffloadsettings["rx_checksum"] = item.RxChecksum
	tcpoffloadsettings["tx_checksum"] = item.TxChecksum

	tcpoffloadsettingss = append(tcpoffloadsettingss, tcpoffloadsettings)
	return tcpoffloadsettingss
}
func flattenMapVnicEthTxQueueSettings(p *models.VnicEthTxQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var txqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	txqueuesettings := make(map[string]interface{})
	delete(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1)
		if err == nil {
			txqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	txqueuesettings["count"] = item.Count
	txqueuesettings["object_type"] = item.ObjectType
	txqueuesettings["ring_size"] = item.RingSize

	txqueuesettingss = append(txqueuesettingss, txqueuesettings)
	return txqueuesettingss
}
func flattenMapVnicVxlanSettings(p *models.VnicVxlanSettings, d *schema.ResourceData) []map[string]interface{} {

	var vxlansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vxlansettings := make(map[string]interface{})
	delete(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1, "ObjectType")
	if len(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1)
		if err == nil {
			vxlansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vxlansettings["enabled"] = item.Enabled
	vxlansettings["object_type"] = item.ObjectType

	vxlansettingss = append(vxlansettingss, vxlansettings)
	return vxlansettingss
}
func flattenMapIamLdapPolicyRef(p *models.IamLdapPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var ldappolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ldappolicy := make(map[string]interface{})
	ldappolicy["moid"] = item.Moid
	ldappolicy["object_type"] = item.ObjectType
	ldappolicy["selector"] = item.Selector

	ldappolicys = append(ldappolicys, ldappolicy)
	return ldappolicys
}
func flattenMapVnicFcAdapterPolicyRef(p *models.VnicFcAdapterPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var fcadapterpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	fcadapterpolicy := make(map[string]interface{})
	fcadapterpolicy["moid"] = item.Moid
	fcadapterpolicy["object_type"] = item.ObjectType
	fcadapterpolicy["selector"] = item.Selector

	fcadapterpolicys = append(fcadapterpolicys, fcadapterpolicy)
	return fcadapterpolicys
}
func flattenMapVnicFcNetworkPolicyRef(p *models.VnicFcNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var fcnetworkpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	fcnetworkpolicy := make(map[string]interface{})
	fcnetworkpolicy["moid"] = item.Moid
	fcnetworkpolicy["object_type"] = item.ObjectType
	fcnetworkpolicy["selector"] = item.Selector

	fcnetworkpolicys = append(fcnetworkpolicys, fcnetworkpolicy)
	return fcnetworkpolicys
}
func flattenMapVnicFcQosPolicyRef(p *models.VnicFcQosPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var fcqospolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	fcqospolicy := make(map[string]interface{})
	fcqospolicy["moid"] = item.Moid
	fcqospolicy["object_type"] = item.ObjectType
	fcqospolicy["selector"] = item.Selector

	fcqospolicys = append(fcqospolicys, fcqospolicy)
	return fcqospolicys
}
func flattenMapVnicSanConnectivityPolicyRef(p *models.VnicSanConnectivityPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var sanconnectivitypolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sanconnectivitypolicy := make(map[string]interface{})
	sanconnectivitypolicy["moid"] = item.Moid
	sanconnectivitypolicy["object_type"] = item.ObjectType
	sanconnectivitypolicy["selector"] = item.Selector

	sanconnectivitypolicys = append(sanconnectivitypolicys, sanconnectivitypolicy)
	return sanconnectivitypolicys
}
func flattenListHyperflexServerFirmwareVersionEntry(p []*models.HyperflexServerFirmwareVersionEntry, d *schema.ResourceData) []map[string]interface{} {
	var serverfirmwareversionentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		serverfirmwareversionentries := make(map[string]interface{})
		delete(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1, "ObjectType")
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				serverfirmwareversionentries["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		serverfirmwareversionentries["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {

			var constraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			constraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					constraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			constraint["hxdp_version"] = item.HxdpVersion
			constraint["hypervisor_type"] = item.HypervisorType
			constraint["mgmt_platform"] = item.MgmtPlatform
			constraint["object_type"] = item.ObjectType
			constraint["server_model"] = item.ServerModel

			constraints = append(constraints, constraint)
			return constraints
		})(item.Constraint, d)
		serverfirmwareversionentries["label"] = item.Label
		serverfirmwareversionentries["name"] = item.Name
		serverfirmwareversionentries["object_type"] = item.ObjectType
		serverfirmwareversionentries["value"] = item.Value
		serverfirmwareversionentriess = append(serverfirmwareversionentriess, serverfirmwareversionentries)
	}
	return serverfirmwareversionentriess
}
func flattenMapIamEndPointPasswordProperties(p *models.IamEndPointPasswordProperties, d *schema.ResourceData) []map[string]interface{} {

	var passwordpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	passwordproperties := make(map[string]interface{})
	delete(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1, "ObjectType")
	if len(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1)
		if err == nil {
			passwordproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	passwordproperties["enable_password_expiry"] = item.EnablePasswordExpiry
	passwordproperties["enforce_strong_password"] = item.EnforceStrongPassword
	passwordproperties["grace_period"] = item.GracePeriod
	passwordproperties["notification_period"] = item.NotificationPeriod
	passwordproperties["object_type"] = item.ObjectType
	passwordproperties["password_expiry_duration"] = item.PasswordExpiryDuration
	passwordproperties["password_history"] = item.PasswordHistory

	passwordpropertiess = append(passwordpropertiess, passwordproperties)
	return passwordpropertiess
}
func flattenListAdapterAdapterConfig(p []*models.AdapterAdapterConfig, d *schema.ResourceData) []map[string]interface{} {
	var settingss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		settings := make(map[string]interface{})
		delete(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1, "ObjectType")
		if len(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1) != 0 {

			j, err := json.Marshal(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1)
			if err == nil {
				settings["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		settings["eth_settings"] = (func(p *models.AdapterEthSettings, d *schema.ResourceData) []map[string]interface{} {

			var ethsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			ethsettings := make(map[string]interface{})
			delete(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1, "ObjectType")
			if len(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1)
				if err == nil {
					ethsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			ethsettings["lldp_enabled"] = item.LldpEnabled
			ethsettings["object_type"] = item.ObjectType

			ethsettingss = append(ethsettingss, ethsettings)
			return ethsettingss
		})(item.EthSettings, d)
		settings["fc_settings"] = (func(p *models.AdapterFcSettings, d *schema.ResourceData) []map[string]interface{} {

			var fcsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			fcsettings := make(map[string]interface{})
			delete(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1, "ObjectType")
			if len(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1)
				if err == nil {
					fcsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			fcsettings["fip_enabled"] = item.FipEnabled
			fcsettings["object_type"] = item.ObjectType

			fcsettingss = append(fcsettingss, fcsettings)
			return fcsettingss
		})(item.FcSettings, d)
		settings["object_type"] = item.ObjectType
		settings["port_channel_settings"] = (func(p *models.AdapterPortChannelSettings, d *schema.ResourceData) []map[string]interface{} {

			var portchannelsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			portchannelsettings := make(map[string]interface{})
			delete(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1, "ObjectType")
			if len(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1)
				if err == nil {
					portchannelsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			portchannelsettings["enabled"] = item.Enabled
			portchannelsettings["object_type"] = item.ObjectType

			portchannelsettingss = append(portchannelsettingss, portchannelsettings)
			return portchannelsettingss
		})(item.PortChannelSettings, d)
		settings["slot_id"] = item.SlotID
		settingss = append(settingss, settings)
	}
	return settingss
}
func flattenListResourceGroupRef(p []*models.ResourceGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var resourcegroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resourcegroups := make(map[string]interface{})
		resourcegroups["moid"] = item.Moid
		resourcegroups["object_type"] = item.ObjectType
		resourcegroups["selector"] = item.Selector
		resourcegroupss = append(resourcegroupss, resourcegroups)
	}
	return resourcegroupss
}
func flattenMapHyperflexLogicalAvailabilityZone(p *models.HyperflexLogicalAvailabilityZone, d *schema.ResourceData) []map[string]interface{} {

	var logicalavalabilityzoneconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	logicalavalabilityzoneconfig := make(map[string]interface{})
	delete(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1, "ObjectType")
	if len(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1)
		if err == nil {
			logicalavalabilityzoneconfig["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	logicalavalabilityzoneconfig["auto_config"] = item.AutoConfig
	logicalavalabilityzoneconfig["object_type"] = item.ObjectType

	logicalavalabilityzoneconfigs = append(logicalavalabilityzoneconfigs, logicalavalabilityzoneconfig)
	return logicalavalabilityzoneconfigs
}
func flattenMapWorkflowCatalogRef(p *models.WorkflowCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var catalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	catalog := make(map[string]interface{})
	catalog["moid"] = item.Moid
	catalog["object_type"] = item.ObjectType
	catalog["selector"] = item.Selector

	catalogs = append(catalogs, catalog)
	return catalogs
}
func flattenListWorkflowBaseDataType(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
	var inputdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		inputdefinition := make(map[string]interface{})
		delete(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1, "ObjectType")
		if len(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1)
			if err == nil {
				inputdefinition["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		inputdefinition["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {

			var propdefaults []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propdefault := make(map[string]interface{})
			delete(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1, "ObjectType")
			if len(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1) != 0 {
				j, err := json.Marshal(item.WorkflowDefaultValueAO1P1.WorkflowDefaultValueAO1P1)
				if err == nil {
					propdefault["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propdefault["object_type"] = item.ObjectType
			propdefault["override"] = item.Override
			propdefault["value"] = item.Value

			propdefaults = append(propdefaults, propdefault)
			return propdefaults
		})(item.Default, d)
		inputdefinition["description"] = item.Description
		inputdefinition["label"] = item.Label
		inputdefinition["name"] = item.Name
		inputdefinition["object_type"] = item.ObjectType
		inputdefinition["required"] = item.Required
		inputdefinitions = append(inputdefinitions, inputdefinition)
	}
	return inputdefinitions
}
func flattenListWorkflowWorkflowTask(p []*models.WorkflowWorkflowTask, d *schema.ResourceData) []map[string]interface{} {
	var taskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		tasks := make(map[string]interface{})
		delete(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1, "ObjectType")
		if len(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1)
			if err == nil {
				tasks["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		tasks["description"] = item.Description
		tasks["label"] = item.Label
		tasks["name"] = item.Name
		tasks["object_type"] = item.ObjectType
		taskss = append(taskss, tasks)
	}
	return taskss
}
func flattenMapWorkflowValidationInformation(p *models.WorkflowValidationInformation, d *schema.ResourceData) []map[string]interface{} {

	var validationinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	validationinformation := make(map[string]interface{})
	delete(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1, "ObjectType")
	if len(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1)
		if err == nil {
			validationinformation["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	validationinformation["object_type"] = item.ObjectType
	validationinformation["state"] = item.State
	validationinformation["validation_error"] = (func(p []*models.WorkflowValidationError, d *schema.ResourceData) []map[string]interface{} {
		var validationerrors []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			validationerror := make(map[string]interface{})
			validationerror["error_log"] = item.ErrorLog
			validationerror["field"] = item.Field
			validationerror["object_type"] = item.ObjectType
			validationerror["task_name"] = item.TaskName
			validationerror["transition_name"] = item.TransitionName
			validationerrors = append(validationerrors, validationerror)
		}
		return validationerrors
	})(item.ValidationError, d)

	validationinformations = append(validationinformations, validationinformation)
	return validationinformations
}
func flattenListIamResourceRolesRef(p []*models.IamResourceRolesRef, d *schema.ResourceData) []map[string]interface{} {
	var resourceroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resourceroles := make(map[string]interface{})
		resourceroles["moid"] = item.Moid
		resourceroles["object_type"] = item.ObjectType
		resourceroles["selector"] = item.Selector
		resourceroless = append(resourceroless, resourceroles)
	}
	return resourceroless
}
func flattenListIamUserGroupRef(p []*models.IamUserGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var usergroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		usergroups := make(map[string]interface{})
		usergroups["moid"] = item.Moid
		usergroups["object_type"] = item.ObjectType
		usergroups["selector"] = item.Selector
		usergroupss = append(usergroupss, usergroups)
	}
	return usergroupss
}
func flattenListIamUserRef(p []*models.IamUserRef, d *schema.ResourceData) []map[string]interface{} {
	var userss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		users := make(map[string]interface{})
		users["moid"] = item.Moid
		users["object_type"] = item.ObjectType
		users["selector"] = item.Selector
		userss = append(userss, users)
	}
	return userss
}
func flattenListTamAPIDataSource(p []*models.TamAPIDataSource, d *schema.ResourceData) []map[string]interface{} {
	var apidatasourcess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		apidatasources := make(map[string]interface{})
		delete(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1, "ObjectType")
		if len(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1) != 0 {

			j, err := json.Marshal(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1)
			if err == nil {
				apidatasources["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		apidatasources["mo_type"] = item.MoType
		apidatasources["name"] = item.Name
		apidatasources["object_type"] = item.ObjectType
		apidatasources["queries"] = (func(p []*models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var queriess []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				queries := make(map[string]interface{})
				delete(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1, "ObjectType")
				if len(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1) != 0 {

					j, err := json.Marshal(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1)
					if err == nil {
						queries["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				queries["name"] = item.Name
				queries["object_type"] = item.ObjectType
				queries["priority"] = item.Priority
				queries["query"] = item.Query
				queriess = append(queriess, queries)
			}
			return queriess
		})(item.Queries, d)
		apidatasources["type"] = item.Type
		apidatasourcess = append(apidatasourcess, apidatasources)
	}
	return apidatasourcess
}
func flattenListTamAction(p []*models.TamAction, d *schema.ResourceData) []map[string]interface{} {
	var actionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		actions := make(map[string]interface{})
		delete(item.TamActionAO1P1.TamActionAO1P1, "ObjectType")
		if len(item.TamActionAO1P1.TamActionAO1P1) != 0 {

			j, err := json.Marshal(item.TamActionAO1P1.TamActionAO1P1)
			if err == nil {
				actions["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		actions["affected_object_type"] = item.AffectedObjectType
		actions["alert_type"] = item.AlertType
		actions["identifiers"] = (func(p []*models.TamIdentifiers, d *schema.ResourceData) []map[string]interface{} {
			var identifierss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				identifiers := make(map[string]interface{})
				delete(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1, "ObjectType")
				if len(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1) != 0 {

					j, err := json.Marshal(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1)
					if err == nil {
						identifiers["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				identifiers["name"] = item.Name
				identifiers["object_type"] = item.ObjectType
				identifiers["value"] = item.Value
				identifierss = append(identifierss, identifiers)
			}
			return identifierss
		})(item.Identifiers, d)
		actions["name"] = item.Name
		actions["object_type"] = item.ObjectType
		actions["operation_type"] = item.OperationType
		actions["queries"] = (func(p []*models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var queriess []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				queries := make(map[string]interface{})
				delete(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1, "ObjectType")
				if len(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1) != 0 {

					j, err := json.Marshal(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1)
					if err == nil {
						queries["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				queries["name"] = item.Name
				queries["object_type"] = item.ObjectType
				queries["priority"] = item.Priority
				queries["query"] = item.Query
				queriess = append(queriess, queries)
			}
			return queriess
		})(item.Queries, d)
		actions["type"] = item.Type
		actionss = append(actionss, actions)
	}
	return actionss
}
func flattenMapTamSeverity(p *models.TamSeverity, d *schema.ResourceData) []map[string]interface{} {

	var severitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	severity := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			severity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	severity["object_type"] = item.ObjectType

	severitys = append(severitys, severity)
	return severitys
}
func flattenListIamAppRegistrationRef(p []*models.IamAppRegistrationRef, d *schema.ResourceData) []map[string]interface{} {
	var appregistrationss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		appregistrations := make(map[string]interface{})
		appregistrations["moid"] = item.Moid
		appregistrations["object_type"] = item.ObjectType
		appregistrations["selector"] = item.Selector
		appregistrationss = append(appregistrationss, appregistrations)
	}
	return appregistrationss
}
func flattenListIamDomainGroupRef(p []*models.IamDomainGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var domaingroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		domaingroups := make(map[string]interface{})
		domaingroups["moid"] = item.Moid
		domaingroups["object_type"] = item.ObjectType
		domaingroups["selector"] = item.Selector
		domaingroupss = append(domaingroupss, domaingroups)
	}
	return domaingroupss
}
func flattenListIamIdpReferenceRef(p []*models.IamIdpReferenceRef, d *schema.ResourceData) []map[string]interface{} {
	var idpreferencess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		idpreferences := make(map[string]interface{})
		idpreferences["moid"] = item.Moid
		idpreferences["object_type"] = item.ObjectType
		idpreferences["selector"] = item.Selector
		idpreferencess = append(idpreferencess, idpreferences)
	}
	return idpreferencess
}
func flattenListIamIdpRef(p []*models.IamIdpRef, d *schema.ResourceData) []map[string]interface{} {
	var idpss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		idps := make(map[string]interface{})
		idps["moid"] = item.Moid
		idps["object_type"] = item.ObjectType
		idps["selector"] = item.Selector
		idpss = append(idpss, idps)
	}
	return idpss
}
func flattenMapCryptEncryptRef(p *models.CryptEncryptRef, d *schema.ResourceData) []map[string]interface{} {

	var nr0encrypts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nr0encrypt := make(map[string]interface{})
	nr0encrypt["moid"] = item.Moid
	nr0encrypt["object_type"] = item.ObjectType
	nr0encrypt["selector"] = item.Selector

	nr0encrypts = append(nr0encrypts, nr0encrypt)
	return nr0encrypts
}
func flattenMapCryptDecryptRef(p *models.CryptDecryptRef, d *schema.ResourceData) []map[string]interface{} {

	var nr1decrypts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nr1decrypt := make(map[string]interface{})
	nr1decrypt["moid"] = item.Moid
	nr1decrypt["object_type"] = item.ObjectType
	nr1decrypt["selector"] = item.Selector

	nr1decrypts = append(nr1decrypts, nr1decrypt)
	return nr1decrypts
}
func flattenListIamPermissionRef(p []*models.IamPermissionRef, d *schema.ResourceData) []map[string]interface{} {
	var permissionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		permissions := make(map[string]interface{})
		permissions["moid"] = item.Moid
		permissions["object_type"] = item.ObjectType
		permissions["selector"] = item.Selector
		permissionss = append(permissionss, permissions)
	}
	return permissionss
}
func flattenListIamPrivilegeSetRef(p []*models.IamPrivilegeSetRef, d *schema.ResourceData) []map[string]interface{} {
	var privilegesetss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		privilegesets := make(map[string]interface{})
		privilegesets["moid"] = item.Moid
		privilegesets["object_type"] = item.ObjectType
		privilegesets["selector"] = item.Selector
		privilegesetss = append(privilegesetss, privilegesets)
	}
	return privilegesetss
}
func flattenListIamPrivilegeRef(p []*models.IamPrivilegeRef, d *schema.ResourceData) []map[string]interface{} {
	var privilegess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		privileges := make(map[string]interface{})
		privileges["moid"] = item.Moid
		privileges["object_type"] = item.ObjectType
		privileges["selector"] = item.Selector
		privilegess = append(privilegess, privileges)
	}
	return privilegess
}
func flattenMapIamResourceLimitsRef(p *models.IamResourceLimitsRef, d *schema.ResourceData) []map[string]interface{} {

	var resourcelimitss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	resourcelimits := make(map[string]interface{})
	resourcelimits["moid"] = item.Moid
	resourcelimits["object_type"] = item.ObjectType
	resourcelimits["selector"] = item.Selector

	resourcelimitss = append(resourcelimitss, resourcelimits)
	return resourcelimitss
}
func flattenMapIamSecurityHolderRef(p *models.IamSecurityHolderRef, d *schema.ResourceData) []map[string]interface{} {

	var securityholders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	securityholder := make(map[string]interface{})
	securityholder["moid"] = item.Moid
	securityholder["object_type"] = item.ObjectType
	securityholder["selector"] = item.Selector

	securityholders = append(securityholders, securityholder)
	return securityholders
}
func flattenMapIamSessionLimitsRef(p *models.IamSessionLimitsRef, d *schema.ResourceData) []map[string]interface{} {

	var sessionlimitss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sessionlimits := make(map[string]interface{})
	sessionlimits["moid"] = item.Moid
	sessionlimits["object_type"] = item.ObjectType
	sessionlimits["selector"] = item.Selector

	sessionlimitss = append(sessionlimitss, sessionlimits)
	return sessionlimitss
}
func flattenListSdcardPartition(p []*models.SdcardPartition, d *schema.ResourceData) []map[string]interface{} {
	var partitionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		partitions := make(map[string]interface{})
		delete(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1, "ObjectType")
		if len(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1) != 0 {

			j, err := json.Marshal(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1)
			if err == nil {
				partitions["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		partitions["object_type"] = item.ObjectType
		partitions["type"] = item.Type
		partitions["virtual_drives"] = (func(p []*models.SdcardVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
			var virtualdrivess []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				virtualdrives := make(map[string]interface{})
				delete(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1, "ObjectType")
				if len(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1) != 0 {

					j, err := json.Marshal(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1)
					if err == nil {
						virtualdrives["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				virtualdrives["enable"] = item.Enable
				virtualdrives["object_type"] = item.ObjectType
				virtualdrivess = append(virtualdrivess, virtualdrives)
			}
			return virtualdrivess
		})(item.VirtualDrives, d)
		partitionss = append(partitionss, partitions)
	}
	return partitionss
}
func flattenMapHyperflexFeatureLimitExternalRef(p *models.HyperflexFeatureLimitExternalRef, d *schema.ResourceData) []map[string]interface{} {

	var featurelimitexternals []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	featurelimitexternal := make(map[string]interface{})
	featurelimitexternal["moid"] = item.Moid
	featurelimitexternal["object_type"] = item.ObjectType
	featurelimitexternal["selector"] = item.Selector

	featurelimitexternals = append(featurelimitexternals, featurelimitexternal)
	return featurelimitexternals
}
func flattenMapHyperflexFeatureLimitInternalRef(p *models.HyperflexFeatureLimitInternalRef, d *schema.ResourceData) []map[string]interface{} {

	var featurelimitinternals []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	featurelimitinternal := make(map[string]interface{})
	featurelimitinternal["moid"] = item.Moid
	featurelimitinternal["object_type"] = item.ObjectType
	featurelimitinternal["selector"] = item.Selector

	featurelimitinternals = append(featurelimitinternals, featurelimitinternal)
	return featurelimitinternals
}
func flattenListHyperflexHxdpVersionRef(p []*models.HyperflexHxdpVersionRef, d *schema.ResourceData) []map[string]interface{} {
	var hxdpversionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hxdpversions := make(map[string]interface{})
		hxdpversions["moid"] = item.Moid
		hxdpversions["object_type"] = item.ObjectType
		hxdpversions["selector"] = item.Selector
		hxdpversionss = append(hxdpversionss, hxdpversions)
	}
	return hxdpversionss
}
func flattenListHyperflexCapabilityInfoRef(p []*models.HyperflexCapabilityInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexcapabilityinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexcapabilityinfos := make(map[string]interface{})
		hyperflexcapabilityinfos["moid"] = item.Moid
		hyperflexcapabilityinfos["object_type"] = item.ObjectType
		hyperflexcapabilityinfos["selector"] = item.Selector
		hyperflexcapabilityinfoss = append(hyperflexcapabilityinfoss, hyperflexcapabilityinfos)
	}
	return hyperflexcapabilityinfoss
}
func flattenListHclHyperflexSoftwareCompatibilityInfoRef(p []*models.HclHyperflexSoftwareCompatibilityInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwarecompatibilityinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hyperflexsoftwarecompatibilityinfos := make(map[string]interface{})
		hyperflexsoftwarecompatibilityinfos["moid"] = item.Moid
		hyperflexsoftwarecompatibilityinfos["object_type"] = item.ObjectType
		hyperflexsoftwarecompatibilityinfos["selector"] = item.Selector
		hyperflexsoftwarecompatibilityinfoss = append(hyperflexsoftwarecompatibilityinfoss, hyperflexsoftwarecompatibilityinfos)
	}
	return hyperflexsoftwarecompatibilityinfoss
}
func flattenMapHyperflexServerFirmwareVersionRef(p *models.HyperflexServerFirmwareVersionRef, d *schema.ResourceData) []map[string]interface{} {

	var serverfirmwareversions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	serverfirmwareversion := make(map[string]interface{})
	serverfirmwareversion["moid"] = item.Moid
	serverfirmwareversion["object_type"] = item.ObjectType
	serverfirmwareversion["selector"] = item.Selector

	serverfirmwareversions = append(serverfirmwareversions, serverfirmwareversion)
	return serverfirmwareversions
}
func flattenMapHyperflexServerModelRef(p *models.HyperflexServerModelRef, d *schema.ResourceData) []map[string]interface{} {

	var servermodels []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	servermodel := make(map[string]interface{})
	servermodel["moid"] = item.Moid
	servermodel["object_type"] = item.ObjectType
	servermodel["selector"] = item.Selector

	servermodels = append(servermodels, servermodel)
	return servermodels
}
func flattenListIamOAuthTokenRef(p []*models.IamOAuthTokenRef, d *schema.ResourceData) []map[string]interface{} {
	var oauthtokenss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		oauthtokens := make(map[string]interface{})
		oauthtokens["moid"] = item.Moid
		oauthtokens["object_type"] = item.ObjectType
		oauthtokens["selector"] = item.Selector
		oauthtokenss = append(oauthtokenss, oauthtokens)
	}
	return oauthtokenss
}
func flattenListIamAPIKeyRef(p []*models.IamAPIKeyRef, d *schema.ResourceData) []map[string]interface{} {
	var apikeyss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		apikeys := make(map[string]interface{})
		apikeys["moid"] = item.Moid
		apikeys["object_type"] = item.ObjectType
		apikeys["selector"] = item.Selector
		apikeyss = append(apikeyss, apikeys)
	}
	return apikeyss
}
func flattenMapIamIdpRef(p *models.IamIdpRef, d *schema.ResourceData) []map[string]interface{} {

	var idps []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	idp := make(map[string]interface{})
	idp["moid"] = item.Moid
	idp["object_type"] = item.ObjectType
	idp["selector"] = item.Selector

	idps = append(idps, idp)
	return idps
}
func flattenMapIamIdpReferenceRef(p *models.IamIdpReferenceRef, d *schema.ResourceData) []map[string]interface{} {

	var idpreferences []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	idpreference := make(map[string]interface{})
	idpreference["moid"] = item.Moid
	idpreference["object_type"] = item.ObjectType
	idpreference["selector"] = item.Selector

	idpreferences = append(idpreferences, idpreference)
	return idpreferences
}
func flattenMapIamLocalUserPasswordRef(p *models.IamLocalUserPasswordRef, d *schema.ResourceData) []map[string]interface{} {

	var localuserpasswords []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	localuserpassword := make(map[string]interface{})
	localuserpassword["moid"] = item.Moid
	localuserpassword["object_type"] = item.ObjectType
	localuserpassword["selector"] = item.Selector

	localuserpasswords = append(localuserpasswords, localuserpassword)
	return localuserpasswords
}
func flattenListIamSessionRef(p []*models.IamSessionRef, d *schema.ResourceData) []map[string]interface{} {
	var sessionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sessions := make(map[string]interface{})
		sessions["moid"] = item.Moid
		sessions["object_type"] = item.ObjectType
		sessions["selector"] = item.Selector
		sessionss = append(sessionss, sessions)
	}
	return sessionss
}
func flattenMapHyperflexClusterRef(p *models.HyperflexClusterRef, d *schema.ResourceData) []map[string]interface{} {

	var associatedclusters []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	associatedcluster := make(map[string]interface{})
	associatedcluster["moid"] = item.Moid
	associatedcluster["object_type"] = item.ObjectType
	associatedcluster["selector"] = item.Selector

	associatedclusters = append(associatedclusters, associatedcluster)
	return associatedclusters
}
func flattenMapHyperflexAutoSupportPolicyRef(p *models.HyperflexAutoSupportPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var autosupports []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	autosupport := make(map[string]interface{})
	autosupport["moid"] = item.Moid
	autosupport["object_type"] = item.ObjectType
	autosupport["selector"] = item.Selector

	autosupports = append(autosupports, autosupport)
	return autosupports
}
func flattenMapHyperflexClusterNetworkPolicyRef(p *models.HyperflexClusterNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var clusternetworks []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	clusternetwork := make(map[string]interface{})
	clusternetwork["moid"] = item.Moid
	clusternetwork["object_type"] = item.ObjectType
	clusternetwork["selector"] = item.Selector

	clusternetworks = append(clusternetworks, clusternetwork)
	return clusternetworks
}
func flattenMapHyperflexClusterStoragePolicyRef(p *models.HyperflexClusterStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var clusterstorages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	clusterstorage := make(map[string]interface{})
	clusterstorage["moid"] = item.Moid
	clusterstorage["object_type"] = item.ObjectType
	clusterstorage["selector"] = item.Selector

	clusterstorages = append(clusterstorages, clusterstorage)
	return clusterstorages
}
func flattenMapHyperflexConfigResultRef(p *models.HyperflexConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var configresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configresult := make(map[string]interface{})
	configresult["moid"] = item.Moid
	configresult["object_type"] = item.ObjectType
	configresult["selector"] = item.Selector

	configresults = append(configresults, configresult)
	return configresults
}
func flattenMapHyperflexExtFcStoragePolicyRef(p *models.HyperflexExtFcStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var extfcstorages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	extfcstorage := make(map[string]interface{})
	extfcstorage["moid"] = item.Moid
	extfcstorage["object_type"] = item.ObjectType
	extfcstorage["selector"] = item.Selector

	extfcstorages = append(extfcstorages, extfcstorage)
	return extfcstorages
}
func flattenMapHyperflexExtIscsiStoragePolicyRef(p *models.HyperflexExtIscsiStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var extiscsistorages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	extiscsistorage := make(map[string]interface{})
	extiscsistorage["moid"] = item.Moid
	extiscsistorage["object_type"] = item.ObjectType
	extiscsistorage["selector"] = item.Selector

	extiscsistorages = append(extiscsistorages, extiscsistorage)
	return extiscsistorages
}
func flattenMapHyperflexLocalCredentialPolicyRef(p *models.HyperflexLocalCredentialPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var localcredentials []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	localcredential := make(map[string]interface{})
	localcredential["moid"] = item.Moid
	localcredential["object_type"] = item.ObjectType
	localcredential["selector"] = item.Selector

	localcredentials = append(localcredentials, localcredential)
	return localcredentials
}
func flattenMapHyperflexNodeConfigPolicyRef(p *models.HyperflexNodeConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var nodeconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nodeconfig := make(map[string]interface{})
	nodeconfig["moid"] = item.Moid
	nodeconfig["object_type"] = item.ObjectType
	nodeconfig["selector"] = item.Selector

	nodeconfigs = append(nodeconfigs, nodeconfig)
	return nodeconfigs
}
func flattenListHyperflexNodeProfileRef(p []*models.HyperflexNodeProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var nodeprofileconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		nodeprofileconfig := make(map[string]interface{})
		nodeprofileconfig["moid"] = item.Moid
		nodeprofileconfig["object_type"] = item.ObjectType
		nodeprofileconfig["selector"] = item.Selector
		nodeprofileconfigs = append(nodeprofileconfigs, nodeprofileconfig)
	}
	return nodeprofileconfigs
}
func flattenMapHyperflexProxySettingPolicyRef(p *models.HyperflexProxySettingPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var proxysettings []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proxysetting := make(map[string]interface{})
	proxysetting["moid"] = item.Moid
	proxysetting["object_type"] = item.ObjectType
	proxysetting["selector"] = item.Selector

	proxysettings = append(proxysettings, proxysetting)
	return proxysettings
}
func flattenMapHyperflexSoftwareVersionPolicyRef(p *models.HyperflexSoftwareVersionPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var softwareversions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	softwareversion := make(map[string]interface{})
	softwareversion["moid"] = item.Moid
	softwareversion["object_type"] = item.ObjectType
	softwareversion["selector"] = item.Selector

	softwareversions = append(softwareversions, softwareversion)
	return softwareversions
}
func flattenMapHyperflexNamedVlan(p *models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {

	var storagedatavlans []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagedatavlan := make(map[string]interface{})
	delete(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1, "ObjectType")
	if len(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1)
		if err == nil {
			storagedatavlan["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	storagedatavlan["name"] = item.Name
	storagedatavlan["object_type"] = item.ObjectType
	storagedatavlan["vlan_id"] = item.VlanID

	storagedatavlans = append(storagedatavlans, storagedatavlan)
	return storagedatavlans
}
func flattenMapHyperflexSysConfigPolicyRef(p *models.HyperflexSysConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var sysconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sysconfig := make(map[string]interface{})
	sysconfig["moid"] = item.Moid
	sysconfig["object_type"] = item.ObjectType
	sysconfig["selector"] = item.Selector

	sysconfigs = append(sysconfigs, sysconfig)
	return sysconfigs
}
func flattenMapHyperflexUcsmConfigPolicyRef(p *models.HyperflexUcsmConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var ucsmconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ucsmconfig := make(map[string]interface{})
	ucsmconfig["moid"] = item.Moid
	ucsmconfig["object_type"] = item.ObjectType
	ucsmconfig["selector"] = item.Selector

	ucsmconfigs = append(ucsmconfigs, ucsmconfig)
	return ucsmconfigs
}
func flattenMapHyperflexVcenterConfigPolicyRef(p *models.HyperflexVcenterConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var vcenterconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vcenterconfig := make(map[string]interface{})
	vcenterconfig["moid"] = item.Moid
	vcenterconfig["object_type"] = item.ObjectType
	vcenterconfig["selector"] = item.Selector

	vcenterconfigs = append(vcenterconfigs, vcenterconfig)
	return vcenterconfigs
}
func flattenListSnmpTrap(p []*models.SnmpTrap, d *schema.ResourceData) []map[string]interface{} {
	var snmptrapss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		snmptraps := make(map[string]interface{})
		delete(item.SnmpTrapAO1P1.SnmpTrapAO1P1, "ObjectType")
		if len(item.SnmpTrapAO1P1.SnmpTrapAO1P1) != 0 {

			j, err := json.Marshal(item.SnmpTrapAO1P1.SnmpTrapAO1P1)
			if err == nil {
				snmptraps["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		snmptraps["destination"] = item.Destination
		snmptraps["enabled"] = item.Enabled
		snmptraps["object_type"] = item.ObjectType
		snmptraps["port"] = item.Port
		snmptraps["type"] = item.Type
		snmptraps["user"] = item.User
		snmptraps["version"] = item.Version
		snmptrapss = append(snmptrapss, snmptraps)
	}
	return snmptrapss
}
func flattenListSnmpUser(p []*models.SnmpUser, d *schema.ResourceData) []map[string]interface{} {
	var snmpuserss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		snmpusers := make(map[string]interface{})
		delete(item.SnmpUserAO1P1.SnmpUserAO1P1, "ObjectType")
		if len(item.SnmpUserAO1P1.SnmpUserAO1P1) != 0 {

			j, err := json.Marshal(item.SnmpUserAO1P1.SnmpUserAO1P1)
			if err == nil {
				snmpusers["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		auth_password_x := d.Get("snmp_users").([]interface{})
		auth_password_y := auth_password_x[0].(map[string]interface{})
		snmpusers["auth_password"] = auth_password_y["auth_password"]
		snmpusers["auth_type"] = item.AuthType
		snmpusers["is_auth_password_set"] = item.IsAuthPasswordSet
		snmpusers["is_privacy_password_set"] = item.IsPrivacyPasswordSet
		snmpusers["name"] = item.Name
		snmpusers["object_type"] = item.ObjectType
		privacy_password_x := d.Get("snmp_users").([]interface{})
		privacy_password_y := privacy_password_x[0].(map[string]interface{})
		snmpusers["privacy_password"] = privacy_password_y["privacy_password"]
		snmpusers["privacy_type"] = item.PrivacyType
		snmpusers["security_level"] = item.SecurityLevel
		snmpuserss = append(snmpuserss, snmpusers)
	}
	return snmpuserss
}
func flattenMapIamCertificateRef(p *models.IamCertificateRef, d *schema.ResourceData) []map[string]interface{} {

	var certificates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	certificate := make(map[string]interface{})
	certificate["moid"] = item.Moid
	certificate["object_type"] = item.ObjectType
	certificate["selector"] = item.Selector

	certificates = append(certificates, certificate)
	return certificates
}
func flattenMapIamPrivateKeySpecRef(p *models.IamPrivateKeySpecRef, d *schema.ResourceData) []map[string]interface{} {

	var privatekeyspecs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	privatekeyspec := make(map[string]interface{})
	privatekeyspec["moid"] = item.Moid
	privatekeyspec["object_type"] = item.ObjectType
	privatekeyspec["selector"] = item.Selector

	privatekeyspecs = append(privatekeyspecs, privatekeyspec)
	return privatekeyspecs
}
func flattenMapPkixDistinguishedName(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

	var subjects []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	subject := make(map[string]interface{})
	delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
	if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
		j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
		if err == nil {
			subject["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	subject["common_name"] = item.CommonName
	subject["country"] = item.Country
	subject["locality"] = item.Locality
	subject["object_type"] = item.ObjectType
	subject["organization"] = item.Organization
	subject["organizational_unit"] = item.OrganizationalUnit
	subject["state"] = item.State

	subjects = append(subjects, subject)
	return subjects
}
func flattenMapPkixSubjectAlternateName(p *models.PkixSubjectAlternateName, d *schema.ResourceData) []map[string]interface{} {

	var subjectalternatenames []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	subjectalternatename := make(map[string]interface{})
	delete(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1, "ObjectType")
	if len(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1) != 0 {
		j, err := json.Marshal(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1)
		if err == nil {
			subjectalternatename["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	subjectalternatename["dns_name"] = item.DNSName
	subjectalternatename["email_address"] = item.EmailAddress
	subjectalternatename["ip_address"] = item.IPAddress
	subjectalternatename["object_type"] = item.ObjectType
	subjectalternatename["uri"] = item.URI

	subjectalternatenames = append(subjectalternatenames, subjectalternatename)
	return subjectalternatenames
}
func flattenMapPkixKeyGenerationSpec(p *models.PkixKeyGenerationSpec, d *schema.ResourceData) []map[string]interface{} {

	var keyspecs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	keyspec := make(map[string]interface{})
	delete(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1, "ObjectType")
	if len(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1) != 0 {
		j, err := json.Marshal(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1)
		if err == nil {
			keyspec["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	keyspec["name"] = item.Name
	keyspec["object_type"] = item.ObjectType

	keyspecs = append(keyspecs, keyspec)
	return keyspecs
}
func flattenMapVnicVsanSettings(p *models.VnicVsanSettings, d *schema.ResourceData) []map[string]interface{} {

	var vsansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vsansettings := make(map[string]interface{})
	delete(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1, "ObjectType")
	if len(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1)
		if err == nil {
			vsansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vsansettings["id"] = item.ID
	vsansettings["object_type"] = item.ObjectType

	vsansettingss = append(vsansettingss, vsansettings)
	return vsansettingss
}
func flattenMapHyperflexNamedVsan(p *models.HyperflexNamedVsan, d *schema.ResourceData) []map[string]interface{} {

	var extatraffics []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	extatraffic := make(map[string]interface{})
	delete(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1, "ObjectType")
	if len(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1)
		if err == nil {
			extatraffic["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	extatraffic["name"] = item.Name
	extatraffic["object_type"] = item.ObjectType
	extatraffic["vsan_id"] = item.VsanID

	extatraffics = append(extatraffics, extatraffic)
	return extatraffics
}
func flattenMapHyperflexWwxnPrefixRange(p *models.HyperflexWwxnPrefixRange, d *schema.ResourceData) []map[string]interface{} {

	var wwxnprefixranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	wwxnprefixrange := make(map[string]interface{})
	delete(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1, "ObjectType")
	if len(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1)
		if err == nil {
			wwxnprefixrange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	wwxnprefixrange["end_addr"] = item.EndAddr
	wwxnprefixrange["object_type"] = item.ObjectType
	wwxnprefixrange["start_addr"] = item.StartAddr

	wwxnprefixranges = append(wwxnprefixranges, wwxnprefixrange)
	return wwxnprefixranges
}
func flattenListBootDeviceBase(p []*models.BootDeviceBase, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		bootdevices := make(map[string]interface{})
		delete(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1, "ObjectType")
		if len(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1) != 0 {

			j, err := json.Marshal(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1)
			if err == nil {
				bootdevices["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		bootdevices["enabled"] = item.Enabled
		bootdevices["name"] = item.Name
		bootdevices["object_type"] = item.ObjectType
		bootdevicess = append(bootdevicess, bootdevices)
	}
	return bootdevicess
}
func flattenMapIamQualifierRef(p *models.IamQualifierRef, d *schema.ResourceData) []map[string]interface{} {

	var qualifiers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	qualifier := make(map[string]interface{})
	qualifier["moid"] = item.Moid
	qualifier["object_type"] = item.ObjectType
	qualifier["selector"] = item.Selector

	qualifiers = append(qualifiers, qualifier)
	return qualifiers
}
func flattenMapCommCredential(p *models.CommCredential, d *schema.ResourceData) []map[string]interface{} {

	var credentials []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	credential := make(map[string]interface{})
	delete(item.CommCredentialAO1P1.CommCredentialAO1P1, "ObjectType")
	if len(item.CommCredentialAO1P1.CommCredentialAO1P1) != 0 {
		j, err := json.Marshal(item.CommCredentialAO1P1.CommCredentialAO1P1)
		if err == nil {
			credential["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	credential["is_password_set"] = item.IsPasswordSet
	credential["object_type"] = item.ObjectType
	credential["password"] = item.Password
	credential["username"] = item.Username

	credentials = append(credentials, credential)
	return credentials
}
func flattenMapAssetManagedDeviceStatus(p *models.AssetManagedDeviceStatus, d *schema.ResourceData) []map[string]interface{} {

	var statuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	status := make(map[string]interface{})
	delete(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1, "ObjectType")
	if len(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1) != 0 {
		j, err := json.Marshal(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1)
		if err == nil {
			status["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	status["cloud_port"] = item.CloudPort
	status["connection_failure_reason"] = item.ConnectionFailureReason
	status["connection_status"] = item.ConnectionStatus
	status["error_code"] = item.ErrorCode
	status["error_reason"] = item.ErrorReason
	status["object_type"] = item.ObjectType
	status["process_id"] = item.ProcessID
	status["server_port"] = item.ServerPort
	status["state"] = item.State

	statuss = append(statuss, status)
	return statuss
}
func flattenMapWorkflowWorkflowInfoRef(p *models.WorkflowWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var workflowinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	workflowinfo := make(map[string]interface{})
	workflowinfo["moid"] = item.Moid
	workflowinfo["object_type"] = item.ObjectType
	workflowinfo["selector"] = item.Selector

	workflowinfos = append(workflowinfos, workflowinfo)
	return workflowinfos
}
func flattenListHyperflexNamedVlan(p []*models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var vmnetworkvlanss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		vmnetworkvlans := make(map[string]interface{})
		delete(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1, "ObjectType")
		if len(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1)
			if err == nil {
				vmnetworkvlans["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		vmnetworkvlans["name"] = item.Name
		vmnetworkvlans["object_type"] = item.ObjectType
		vmnetworkvlans["vlan_id"] = item.VlanID
		vmnetworkvlanss = append(vmnetworkvlanss, vmnetworkvlans)
	}
	return vmnetworkvlanss
}
func flattenListHclConstraint(p []*models.HclConstraint, d *schema.ResourceData) []map[string]interface{} {
	var capabilityconstraintss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		capabilityconstraints := make(map[string]interface{})
		delete(item.HclConstraintAO1P1.HclConstraintAO1P1, "ObjectType")
		if len(item.HclConstraintAO1P1.HclConstraintAO1P1) != 0 {

			j, err := json.Marshal(item.HclConstraintAO1P1.HclConstraintAO1P1)
			if err == nil {
				capabilityconstraints["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		capabilityconstraints["constraint_name"] = item.ConstraintName
		capabilityconstraints["constraint_value"] = item.ConstraintValue
		capabilityconstraints["object_type"] = item.ObjectType
		capabilityconstraintss = append(capabilityconstraintss, capabilityconstraints)
	}
	return capabilityconstraintss
}
func flattenMapHyperflexInstallerImageRef(p *models.HyperflexInstallerImageRef, d *schema.ResourceData) []map[string]interface{} {

	var installerimages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	installerimage := make(map[string]interface{})
	installerimage["moid"] = item.Moid
	installerimage["object_type"] = item.ObjectType
	installerimage["selector"] = item.Selector

	installerimages = append(installerimages, installerimage)
	return installerimages
}
func flattenMapIamCertificateRequestRef(p *models.IamCertificateRequestRef, d *schema.ResourceData) []map[string]interface{} {

	var certificaterequests []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	certificaterequest := make(map[string]interface{})
	certificaterequest["moid"] = item.Moid
	certificaterequest["object_type"] = item.ObjectType
	certificaterequest["selector"] = item.Selector

	certificaterequests = append(certificaterequests, certificaterequest)
	return certificaterequests
}
func flattenListWorkflowTaskDefinitionRef(p []*models.WorkflowTaskDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var implementedtaskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		implementedtasks := make(map[string]interface{})
		implementedtasks["moid"] = item.Moid
		implementedtasks["object_type"] = item.ObjectType
		implementedtasks["selector"] = item.Selector
		implementedtaskss = append(implementedtaskss, implementedtasks)
	}
	return implementedtaskss
}
func flattenMapWorkflowInternalProperties(p *models.WorkflowInternalProperties, d *schema.ResourceData) []map[string]interface{} {

	var internalpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	internalproperties := make(map[string]interface{})
	delete(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1)
		if err == nil {
			internalproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	internalproperties["base_task_type"] = item.BaseTaskType
	internalproperties["constraints"] = item.Constraints
	internalproperties["internal"] = item.Internal
	internalproperties["object_type"] = item.ObjectType
	internalproperties["owner"] = item.Owner

	internalpropertiess = append(internalpropertiess, internalproperties)
	return internalpropertiess
}
func flattenMapWorkflowProperties(p *models.WorkflowProperties, d *schema.ResourceData) []map[string]interface{} {

	var propertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	properties := make(map[string]interface{})
	delete(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1)
		if err == nil {
			properties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	properties["input_definition"] = (func(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var inputdefinitions []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			inputdefinition := make(map[string]interface{})
			inputdefinition["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {

				var propdefaults []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				propdefault := make(map[string]interface{})
				propdefault["object_type"] = item.ObjectType
				propdefault["override"] = item.Override
				propdefault["value"] = item.Value

				propdefaults = append(propdefaults, propdefault)
				return propdefaults
			})(item.Default, d)
			inputdefinition["description"] = item.Description
			inputdefinition["label"] = item.Label
			inputdefinition["name"] = item.Name
			inputdefinition["object_type"] = item.ObjectType
			inputdefinition["required"] = item.Required
			inputdefinitions = append(inputdefinitions, inputdefinition)
		}
		return inputdefinitions
	})(item.InputDefinition, d)
	properties["object_type"] = item.ObjectType
	properties["output_definition"] = (func(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var outputdefinitions []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			outputdefinition := make(map[string]interface{})
			outputdefinition["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {

				var propdefaults []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				propdefault := make(map[string]interface{})
				propdefault["object_type"] = item.ObjectType
				propdefault["override"] = item.Override
				propdefault["value"] = item.Value

				propdefaults = append(propdefaults, propdefault)
				return propdefaults
			})(item.Default, d)
			outputdefinition["description"] = item.Description
			outputdefinition["label"] = item.Label
			outputdefinition["name"] = item.Name
			outputdefinition["object_type"] = item.ObjectType
			outputdefinition["required"] = item.Required
			outputdefinitions = append(outputdefinitions, outputdefinition)
		}
		return outputdefinitions
	})(item.OutputDefinition, d)
	properties["retry_count"] = item.RetryCount
	properties["retry_delay"] = item.RetryDelay
	properties["retry_policy"] = item.RetryPolicy
	properties["timeout"] = item.Timeout
	properties["timeout_policy"] = item.TimeoutPolicy

	propertiess = append(propertiess, properties)
	return propertiess
}
func flattenListStorageSpanGroup(p []*models.StorageSpanGroup, d *schema.ResourceData) []map[string]interface{} {
	var spangroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		spangroups := make(map[string]interface{})
		delete(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1, "ObjectType")
		if len(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1) != 0 {

			j, err := json.Marshal(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1)
			if err == nil {
				spangroups["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		spangroups["disks"] = (func(p []*models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
			var diskss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				disks := make(map[string]interface{})
				delete(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1, "ObjectType")
				if len(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1) != 0 {

					j, err := json.Marshal(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1)
					if err == nil {
						disks["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				disks["object_type"] = item.ObjectType
				disks["slot_number"] = item.SlotNumber
				diskss = append(diskss, disks)
			}
			return diskss
		})(item.Disks, d)
		spangroups["object_type"] = item.ObjectType
		spangroupss = append(spangroupss, spangroups)
	}
	return spangroupss
}
func flattenListStorageStoragePolicyRef(p []*models.StorageStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var storagepoliciess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagepolicies := make(map[string]interface{})
		storagepolicies["moid"] = item.Moid
		storagepolicies["object_type"] = item.ObjectType
		storagepolicies["selector"] = item.Selector
		storagepoliciess = append(storagepoliciess, storagepolicies)
	}
	return storagepoliciess
}
func flattenMapIamSystemRef(p *models.IamSystemRef, d *schema.ResourceData) []map[string]interface{} {

	var systems []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	system := make(map[string]interface{})
	system["moid"] = item.Moid
	system["object_type"] = item.ObjectType
	system["selector"] = item.Selector

	systems = append(systems, system)
	return systems
}
func flattenListIamUserLoginTimeRef(p []*models.IamUserLoginTimeRef, d *schema.ResourceData) []map[string]interface{} {
	var userlogintimes []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		userlogintime := make(map[string]interface{})
		userlogintime["moid"] = item.Moid
		userlogintime["object_type"] = item.ObjectType
		userlogintime["selector"] = item.Selector
		userlogintimes = append(userlogintimes, userlogintime)
	}
	return userlogintimes
}
func flattenListIamUserPreferenceRef(p []*models.IamUserPreferenceRef, d *schema.ResourceData) []map[string]interface{} {
	var userpreferencess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		userpreferences := make(map[string]interface{})
		userpreferences["moid"] = item.Moid
		userpreferences["object_type"] = item.ObjectType
		userpreferences["selector"] = item.Selector
		userpreferencess = append(userpreferencess, userpreferences)
	}
	return userpreferencess
}
func flattenListVnicEthIfRef(p []*models.VnicEthIfRef, d *schema.ResourceData) []map[string]interface{} {
	var ethifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		ethifs := make(map[string]interface{})
		ethifs["moid"] = item.Moid
		ethifs["object_type"] = item.ObjectType
		ethifs["selector"] = item.Selector
		ethifss = append(ethifss, ethifs)
	}
	return ethifss
}
func flattenMapSoftwareHyperflexDistributableRef(p *models.SoftwareHyperflexDistributableRef, d *schema.ResourceData) []map[string]interface{} {

	var hxdpversioninfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hxdpversioninfo := make(map[string]interface{})
	hxdpversioninfo["moid"] = item.Moid
	hxdpversioninfo["object_type"] = item.ObjectType
	hxdpversioninfo["selector"] = item.Selector

	hxdpversioninfos = append(hxdpversioninfos, hxdpversioninfo)
	return hxdpversioninfos
}
func flattenMapFirmwareDistributableRef(p *models.FirmwareDistributableRef, d *schema.ResourceData) []map[string]interface{} {

	var serverfirmwareversioninfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	serverfirmwareversioninfo := make(map[string]interface{})
	serverfirmwareversioninfo["moid"] = item.Moid
	serverfirmwareversioninfo["object_type"] = item.ObjectType
	serverfirmwareversioninfo["selector"] = item.Selector

	serverfirmwareversioninfos = append(serverfirmwareversioninfos, serverfirmwareversioninfo)
	return serverfirmwareversioninfos
}
func flattenListSyslogLocalClientBase(p []*models.SyslogLocalClientBase, d *schema.ResourceData) []map[string]interface{} {
	var localclientss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		localclients := make(map[string]interface{})
		delete(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1, "ObjectType")
		if len(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1) != 0 {

			j, err := json.Marshal(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1)
			if err == nil {
				localclients["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		localclients["min_severity"] = item.MinSeverity
		localclients["object_type"] = item.ObjectType
		localclientss = append(localclientss, localclients)
	}
	return localclientss
}
func flattenListSyslogRemoteClientBase(p []*models.SyslogRemoteClientBase, d *schema.ResourceData) []map[string]interface{} {
	var remoteclientss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		remoteclients := make(map[string]interface{})
		delete(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1, "ObjectType")
		if len(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1) != 0 {

			j, err := json.Marshal(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1)
			if err == nil {
				remoteclients["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		remoteclients["enabled"] = item.Enabled
		remoteclients["hostname"] = item.Hostname
		remoteclients["min_severity"] = item.MinSeverity
		remoteclients["object_type"] = item.ObjectType
		remoteclients["port"] = item.Port
		remoteclients["protocol"] = item.Protocol
		remoteclientss = append(remoteclientss, remoteclients)
	}
	return remoteclientss
}
func flattenMapX509Certificate(p *models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {

	var certificates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	certificate := make(map[string]interface{})
	delete(item.X509CertificateAO1P1.X509CertificateAO1P1, "ObjectType")
	if len(item.X509CertificateAO1P1.X509CertificateAO1P1) != 0 {
		j, err := json.Marshal(item.X509CertificateAO1P1.X509CertificateAO1P1)
		if err == nil {
			certificate["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	certificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

		var issuers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		issuer := make(map[string]interface{})
		issuer["common_name"] = item.CommonName
		issuer["country"] = item.Country
		issuer["locality"] = item.Locality
		issuer["object_type"] = item.ObjectType
		issuer["organization"] = item.Organization
		issuer["organizational_unit"] = item.OrganizationalUnit
		issuer["state"] = item.State

		issuers = append(issuers, issuer)
		return issuers
	})(item.Issuer, d)
	certificate["object_type"] = item.ObjectType
	certificate["pem_certificate"] = item.PemCertificate
	certificate["sha256_fingerprint"] = item.Sha256Fingerprint
	certificate["signature_algorithm"] = item.SignatureAlgorithm
	certificate["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

		var subjects []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		subject := make(map[string]interface{})
		subject["common_name"] = item.CommonName
		subject["country"] = item.Country
		subject["locality"] = item.Locality
		subject["object_type"] = item.ObjectType
		subject["organization"] = item.Organization
		subject["organizational_unit"] = item.OrganizationalUnit
		subject["state"] = item.State

		subjects = append(subjects, subject)
		return subjects
	})(item.Subject, d)

	certificates = append(certificates, certificate)
	return certificates
}
func flattenListOrganizationOrganizationRef(p []*models.OrganizationOrganizationRef, d *schema.ResourceData) []map[string]interface{} {
	var organizationss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		organizations := make(map[string]interface{})
		organizations["moid"] = item.Moid
		organizations["object_type"] = item.ObjectType
		organizations["selector"] = item.Selector
		organizationss = append(organizationss, organizations)
	}
	return organizationss
}
func flattenListResourcePerTypeCombinedSelector(p []*models.ResourcePerTypeCombinedSelector, d *schema.ResourceData) []map[string]interface{} {
	var pertypecombinedselectors []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pertypecombinedselector := make(map[string]interface{})
		delete(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1, "ObjectType")
		if len(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1) != 0 {

			j, err := json.Marshal(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1)
			if err == nil {
				pertypecombinedselector["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		pertypecombinedselector["combined_selector"] = item.CombinedSelector
		pertypecombinedselector["empty_filter"] = item.EmptyFilter
		pertypecombinedselector["object_type"] = item.ObjectType
		pertypecombinedselector["selector_object_type"] = item.SelectorObjectType
		pertypecombinedselectors = append(pertypecombinedselectors, pertypecombinedselector)
	}
	return pertypecombinedselectors
}
func flattenListResourceSelector(p []*models.ResourceSelector, d *schema.ResourceData) []map[string]interface{} {
	var selectorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		selectors := make(map[string]interface{})
		delete(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1, "ObjectType")
		if len(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1) != 0 {

			j, err := json.Marshal(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1)
			if err == nil {
				selectors["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		selectors["object_type"] = item.ObjectType
		selectors["selector"] = item.Selector
		selectorss = append(selectorss, selectors)
	}
	return selectorss
}
func flattenMapVnicVlanSettings(p *models.VnicVlanSettings, d *schema.ResourceData) []map[string]interface{} {

	var vlansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vlansettings := make(map[string]interface{})
	delete(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1, "ObjectType")
	if len(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1)
		if err == nil {
			vlansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	vlansettings["default_vlan"] = item.DefaultVlan
	vlansettings["mode"] = item.Mode
	vlansettings["object_type"] = item.ObjectType

	vlansettingss = append(vlansettingss, vlansettings)
	return vlansettingss
}
func flattenMapOsCatalogRef(p *models.OsCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var catalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	catalog := make(map[string]interface{})
	catalog["moid"] = item.Moid
	catalog["object_type"] = item.ObjectType
	catalog["selector"] = item.Selector

	catalogs = append(catalogs, catalog)
	return catalogs
}
func flattenListHclOperatingSystemRef(p []*models.HclOperatingSystemRef, d *schema.ResourceData) []map[string]interface{} {
	var distributionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		distributions := make(map[string]interface{})
		distributions["moid"] = item.Moid
		distributions["object_type"] = item.ObjectType
		distributions["selector"] = item.Selector
		distributionss = append(distributionss, distributions)
	}
	return distributionss
}
func flattenListOsPlaceHolder(p []*models.OsPlaceHolder, d *schema.ResourceData) []map[string]interface{} {
	var placeholderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		placeholders := make(map[string]interface{})
		delete(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1, "ObjectType")
		if len(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1) != 0 {

			j, err := json.Marshal(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1)
			if err == nil {
				placeholders["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		placeholders["is_value_set"] = item.IsValueSet
		placeholders["object_type"] = item.ObjectType
		placeholders["type"] = (func(p *models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {

			var proptypes []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			proptype := make(map[string]interface{})
			delete(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1, "ObjectType")
			if len(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1) != 0 {
				j, err := json.Marshal(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1)
				if err == nil {
					proptype["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			proptype["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {

				var propdefaults []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				propdefault := make(map[string]interface{})
				propdefault["object_type"] = item.ObjectType
				propdefault["override"] = item.Override
				propdefault["value"] = item.Value

				propdefaults = append(propdefaults, propdefault)
				return propdefaults
			})(item.Default, d)
			proptype["description"] = item.Description
			proptype["label"] = item.Label
			proptype["name"] = item.Name
			proptype["object_type"] = item.ObjectType
			proptype["properties"] = (func(p *models.WorkflowPrimitiveDataProperty, d *schema.ResourceData) []map[string]interface{} {

				var propertiess []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				properties := make(map[string]interface{})
				properties["constraints"] = (func(p *models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {

					var constraintss []map[string]interface{}
					if p == nil {
						return nil
					}
					item := *p
					constraints := make(map[string]interface{})
					constraints["enum_list"] = (func(p []*models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
						var enumlists []map[string]interface{}
						if p == nil {
							return nil
						}
						for _, item := range p {
							item := *item
							enumlist := make(map[string]interface{})
							enumlist["label"] = item.Label
							enumlist["object_type"] = item.ObjectType
							enumlist["value"] = item.Value
							enumlists = append(enumlists, enumlist)
						}
						return enumlists
					})(item.EnumList, d)
					constraints["max"] = item.Max
					constraints["min"] = item.Min
					constraints["object_type"] = item.ObjectType
					constraints["regex"] = item.Regex

					constraintss = append(constraintss, constraints)
					return constraintss
				})(item.Constraints, d)
				properties["object_type"] = item.ObjectType
				properties["secure"] = item.Secure
				properties["type"] = item.Type

				propertiess = append(propertiess, properties)
				return propertiess
			})(item.Properties, d)
			proptype["required"] = item.Required

			proptypes = append(proptypes, proptype)
			return proptypes
		})(item.Type, d)
		placeholders["value"] = item.Value
		placeholderss = append(placeholderss, placeholders)
	}
	return placeholderss
}
func flattenMapRecoveryAbstractBackupInfoRef(p *models.RecoveryAbstractBackupInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var backupinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	backupinfo := make(map[string]interface{})
	backupinfo["moid"] = item.Moid
	backupinfo["object_type"] = item.ObjectType
	backupinfo["selector"] = item.Selector

	backupinfos = append(backupinfos, backupinfo)
	return backupinfos
}
func flattenMapRecoveryConfigParams(p *models.RecoveryConfigParams, d *schema.ResourceData) []map[string]interface{} {

	var configparamss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configparams := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			configparams["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	configparams["object_type"] = item.ObjectType

	configparamss = append(configparamss, configparams)
	return configparamss
}
func flattenMapRecoveryAbstractRestoreStatusRef(p *models.RecoveryAbstractRestoreStatusRef, d *schema.ResourceData) []map[string]interface{} {

	var restorestatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	restorestatus := make(map[string]interface{})
	restorestatus["moid"] = item.Moid
	restorestatus["object_type"] = item.ObjectType
	restorestatus["selector"] = item.Selector

	restorestatuss = append(restorestatuss, restorestatus)
	return restorestatuss
}
func flattenListX509Certificate(p []*models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var certificatess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		certificates := make(map[string]interface{})
		delete(item.X509CertificateAO1P1.X509CertificateAO1P1, "ObjectType")
		if len(item.X509CertificateAO1P1.X509CertificateAO1P1) != 0 {

			j, err := json.Marshal(item.X509CertificateAO1P1.X509CertificateAO1P1)
			if err == nil {
				certificates["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		certificates["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var issuers []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			issuer := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					issuer["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			issuer["common_name"] = item.CommonName
			issuer["country"] = item.Country
			issuer["locality"] = item.Locality
			issuer["object_type"] = item.ObjectType
			issuer["organization"] = item.Organization
			issuer["organizational_unit"] = item.OrganizationalUnit
			issuer["state"] = item.State

			issuers = append(issuers, issuer)
			return issuers
		})(item.Issuer, d)
		certificates["object_type"] = item.ObjectType
		certificates["pem_certificate"] = item.PemCertificate
		certificates["sha256_fingerprint"] = item.Sha256Fingerprint
		certificates["signature_algorithm"] = item.SignatureAlgorithm
		certificates["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var subjects []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			subject := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					subject["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			subject["common_name"] = item.CommonName
			subject["country"] = item.Country
			subject["locality"] = item.Locality
			subject["object_type"] = item.ObjectType
			subject["organization"] = item.Organization
			subject["organizational_unit"] = item.OrganizationalUnit
			subject["state"] = item.State

			subjects = append(subjects, subject)
			return subjects
		})(item.Subject, d)
		certificatess = append(certificatess, certificates)
	}
	return certificatess
}
func flattenMapFirmwareDirectDownload(p *models.FirmwareDirectDownload, d *schema.ResourceData) []map[string]interface{} {

	var directdownloads []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	directdownload := make(map[string]interface{})
	delete(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1, "ObjectType")
	if len(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1) != 0 {
		j, err := json.Marshal(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1)
		if err == nil {
			directdownload["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	directdownload["http_server"] = (func(p *models.FirmwareHTTPServer, d *schema.ResourceData) []map[string]interface{} {

		var httpservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		httpserver := make(map[string]interface{})
		httpserver["location_link"] = item.LocationLink
		httpserver["mount_options"] = item.MountOptions
		httpserver["object_type"] = item.ObjectType

		httpservers = append(httpservers, httpserver)
		return httpservers
	})(item.HTTPServer, d)
	directdownload["image_source"] = item.ImageSource
	directdownload["is_password_set"] = item.IsPasswordSet
	directdownload["object_type"] = item.ObjectType
	directdownload["password"] = item.Password
	directdownload["upgradeoption"] = item.Upgradeoption
	directdownload["username"] = item.Username

	directdownloads = append(directdownloads, directdownload)
	return directdownloads
}
func flattenMapFirmwareNetworkShare(p *models.FirmwareNetworkShare, d *schema.ResourceData) []map[string]interface{} {

	var networkshares []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	networkshare := make(map[string]interface{})
	delete(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1, "ObjectType")
	if len(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1) != 0 {
		j, err := json.Marshal(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1)
		if err == nil {
			networkshare["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	networkshare["cifs_server"] = (func(p *models.FirmwareCifsServer, d *schema.ResourceData) []map[string]interface{} {

		var cifsservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		cifsserver := make(map[string]interface{})
		cifsserver["mount_options"] = item.MountOptions
		cifsserver["object_type"] = item.ObjectType
		cifsserver["remote_file"] = item.RemoteFile
		cifsserver["remote_ip"] = item.RemoteIP
		cifsserver["remote_share"] = item.RemoteShare

		cifsservers = append(cifsservers, cifsserver)
		return cifsservers
	})(item.CifsServer, d)
	networkshare["http_server"] = (func(p *models.FirmwareHTTPServer, d *schema.ResourceData) []map[string]interface{} {

		var httpservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		httpserver := make(map[string]interface{})
		httpserver["location_link"] = item.LocationLink
		httpserver["mount_options"] = item.MountOptions
		httpserver["object_type"] = item.ObjectType

		httpservers = append(httpservers, httpserver)
		return httpservers
	})(item.HTTPServer, d)
	networkshare["is_password_set"] = item.IsPasswordSet
	networkshare["map_type"] = item.MapType
	networkshare["nfs_server"] = (func(p *models.FirmwareNfsServer, d *schema.ResourceData) []map[string]interface{} {

		var nfsservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		nfsserver := make(map[string]interface{})
		nfsserver["mount_options"] = item.MountOptions
		nfsserver["object_type"] = item.ObjectType
		nfsserver["remote_file"] = item.RemoteFile
		nfsserver["remote_ip"] = item.RemoteIP
		nfsserver["remote_share"] = item.RemoteShare

		nfsservers = append(nfsservers, nfsserver)
		return nfsservers
	})(item.NfsServer, d)
	networkshare["object_type"] = item.ObjectType
	networkshare["password"] = item.Password
	networkshare["upgradeoption"] = item.Upgradeoption
	networkshare["username"] = item.Username

	networkshares = append(networkshares, networkshare)
	return networkshares
}
func flattenMapFirmwareUpgradeStatusRef(p *models.FirmwareUpgradeStatusRef, d *schema.ResourceData) []map[string]interface{} {

	var upgradestatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	upgradestatus := make(map[string]interface{})
	upgradestatus["moid"] = item.Moid
	upgradestatus["object_type"] = item.ObjectType
	upgradestatus["selector"] = item.Selector

	upgradestatuss = append(upgradestatuss, upgradestatus)
	return upgradestatuss
}
func flattenMapOsAnswers(p *models.OsAnswers, d *schema.ResourceData) []map[string]interface{} {

	var answerss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	answers := make(map[string]interface{})
	delete(item.OsAnswersAO1P1.OsAnswersAO1P1, "ObjectType")
	if len(item.OsAnswersAO1P1.OsAnswersAO1P1) != 0 {
		j, err := json.Marshal(item.OsAnswersAO1P1.OsAnswersAO1P1)
		if err == nil {
			answers["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	answers["answer_file"] = item.AnswerFile
	answers["hostname"] = item.Hostname
	answers["ip_config_type"] = item.IPConfigType
	answers["ipv4_config"] = (func(p *models.CommIPV4Interface, d *schema.ResourceData) []map[string]interface{} {

		var ipv4configs []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		ipv4config := make(map[string]interface{})
		ipv4config["gateway"] = item.Gateway
		ipv4config["ip_address"] = item.IPAddress
		ipv4config["netmask"] = item.Netmask
		ipv4config["object_type"] = item.ObjectType

		ipv4configs = append(ipv4configs, ipv4config)
		return ipv4configs
	})(item.IPV4Config, d)
	answers["is_answer_file_set"] = item.IsAnswerFileSet
	answers["is_root_password_set"] = item.IsRootPasswordSet
	answers["nameserver"] = item.Nameserver
	answers["object_type"] = item.ObjectType
	answers["product_key"] = item.ProductKey
	answers["root_password"] = item.RootPassword
	answers["source"] = item.Source

	answerss = append(answerss, answers)
	return answerss
}
func flattenMapOsConfigurationFileRef(p *models.OsConfigurationFileRef, d *schema.ResourceData) []map[string]interface{} {

	var configurationfiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configurationfile := make(map[string]interface{})
	configurationfile["moid"] = item.Moid
	configurationfile["object_type"] = item.ObjectType
	configurationfile["selector"] = item.Selector

	configurationfiles = append(configurationfiles, configurationfile)
	return configurationfiles
}
func flattenMapSoftwarerepositoryOperatingSystemFileRef(p *models.SoftwarerepositoryOperatingSystemFileRef, d *schema.ResourceData) []map[string]interface{} {

	var images []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	image := make(map[string]interface{})
	image["moid"] = item.Moid
	image["object_type"] = item.ObjectType
	image["selector"] = item.Selector

	images = append(images, image)
	return images
}
func flattenMapOsOperatingSystemParameters(p *models.OsOperatingSystemParameters, d *schema.ResourceData) []map[string]interface{} {

	var operatingsystemparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	operatingsystemparameters := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			operatingsystemparameters["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	operatingsystemparameters["object_type"] = item.ObjectType

	operatingsystemparameterss = append(operatingsystemparameterss, operatingsystemparameters)
	return operatingsystemparameterss
}
func flattenMapFirmwareServerConfigurationUtilityDistributableRef(p *models.FirmwareServerConfigurationUtilityDistributableRef, d *schema.ResourceData) []map[string]interface{} {

	var osduimages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	osduimage := make(map[string]interface{})
	osduimage["moid"] = item.Moid
	osduimage["object_type"] = item.ObjectType
	osduimage["selector"] = item.Selector

	osduimages = append(osduimages, osduimage)
	return osduimages
}
func flattenListOsPostInstallScriptRef(p []*models.OsPostInstallScriptRef, d *schema.ResourceData) []map[string]interface{} {
	var postinstallscriptss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		postinstallscripts := make(map[string]interface{})
		postinstallscripts["moid"] = item.Moid
		postinstallscripts["object_type"] = item.ObjectType
		postinstallscripts["selector"] = item.Selector
		postinstallscriptss = append(postinstallscriptss, postinstallscripts)
	}
	return postinstallscriptss
}
func flattenMapComputePhysicalRef(p *models.ComputePhysicalRef, d *schema.ResourceData) []map[string]interface{} {

	var servers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	server := make(map[string]interface{})
	server["moid"] = item.Moid
	server["object_type"] = item.ObjectType
	server["selector"] = item.Selector

	servers = append(servers, server)
	return servers
}
func flattenMapStorageFlexUtilControllerRef(p *models.StorageFlexUtilControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var storageflexutilcontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageflexutilcontroller := make(map[string]interface{})
	storageflexutilcontroller["moid"] = item.Moid
	storageflexutilcontroller["object_type"] = item.ObjectType
	storageflexutilcontroller["selector"] = item.Selector

	storageflexutilcontrollers = append(storageflexutilcontrollers, storageflexutilcontroller)
	return storageflexutilcontrollers
}
func flattenMapEquipmentChassisRef(p *models.EquipmentChassisRef, d *schema.ResourceData) []map[string]interface{} {

	var equipmentchassiss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentchassis := make(map[string]interface{})
	equipmentchassis["moid"] = item.Moid
	equipmentchassis["object_type"] = item.ObjectType
	equipmentchassis["selector"] = item.Selector

	equipmentchassiss = append(equipmentchassiss, equipmentchassis)
	return equipmentchassiss
}
func flattenMapEquipmentSharedIoModuleRef(p *models.EquipmentSharedIoModuleRef, d *schema.ResourceData) []map[string]interface{} {

	var sharediomodules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sharediomodule := make(map[string]interface{})
	sharediomodule["moid"] = item.Moid
	sharediomodule["object_type"] = item.ObjectType
	sharediomodule["selector"] = item.Selector

	sharediomodules = append(sharediomodules, sharediomodule)
	return sharediomodules
}
func flattenMapNiaapiVersionRegexPlatform(p *models.NiaapiVersionRegexPlatform, d *schema.ResourceData) []map[string]interface{} {

	var apics []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	apic := make(map[string]interface{})
	apic["anyllregex"] = item.Anyllregex
	apic["currentlltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {

		var currentlltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		currentlltrain := make(map[string]interface{})
		currentlltrain["object_type"] = item.ObjectType
		currentlltrain["regex"] = item.Regex
		currentlltrain["software_version"] = item.SoftwareVersion

		currentlltrains = append(currentlltrains, currentlltrain)
		return currentlltrains
	})(item.Currentlltrain, d)
	apic["latestsltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {

		var latestsltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		latestsltrain := make(map[string]interface{})
		latestsltrain["object_type"] = item.ObjectType
		latestsltrain["regex"] = item.Regex
		latestsltrain["software_version"] = item.SoftwareVersion

		latestsltrains = append(latestsltrains, latestsltrain)
		return latestsltrains
	})(item.Latestsltrain, d)
	apic["object_type"] = item.ObjectType
	apic["sltrain"] = (func(p []*models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var sltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			sltrain := make(map[string]interface{})
			sltrain["object_type"] = item.ObjectType
			sltrain["regex"] = item.Regex
			sltrain["software_version"] = item.SoftwareVersion
			sltrains = append(sltrains, sltrain)
		}
		return sltrains
	})(item.Sltrain, d)
	apic["upcominglltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {

		var upcominglltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		upcominglltrain := make(map[string]interface{})
		upcominglltrain["object_type"] = item.ObjectType
		upcominglltrain["regex"] = item.Regex
		upcominglltrain["software_version"] = item.SoftwareVersion

		upcominglltrains = append(upcominglltrains, upcominglltrain)
		return upcominglltrains
	})(item.Upcominglltrain, d)

	apics = append(apics, apic)
	return apics
}
func flattenMapStorageFlexFlashControllerRef(p *models.StorageFlexFlashControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var storageflexflashcontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageflexflashcontroller := make(map[string]interface{})
	storageflexflashcontroller["moid"] = item.Moid
	storageflexflashcontroller["object_type"] = item.ObjectType
	storageflexflashcontroller["selector"] = item.Selector

	storageflexflashcontrollers = append(storageflexflashcontrollers, storageflexflashcontroller)
	return storageflexflashcontrollers
}
func flattenListStoragePhysicalDiskUsageRef(p []*models.StoragePhysicalDiskUsageRef, d *schema.ResourceData) []map[string]interface{} {
	var physicaldiskusagess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		physicaldiskusages := make(map[string]interface{})
		physicaldiskusages["moid"] = item.Moid
		physicaldiskusages["object_type"] = item.ObjectType
		physicaldiskusages["selector"] = item.Selector
		physicaldiskusagess = append(physicaldiskusagess, physicaldiskusages)
	}
	return physicaldiskusagess
}
func flattenMapStorageControllerRef(p *models.StorageControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var storagecontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagecontroller := make(map[string]interface{})
	storagecontroller["moid"] = item.Moid
	storagecontroller["object_type"] = item.ObjectType
	storagecontroller["selector"] = item.Selector

	storagecontrollers = append(storagecontrollers, storagecontroller)
	return storagecontrollers
}
func flattenListStorageVdMemberEpRef(p []*models.StorageVdMemberEpRef, d *schema.ResourceData) []map[string]interface{} {
	var vdmemberepss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		vdmembereps := make(map[string]interface{})
		vdmembereps["moid"] = item.Moid
		vdmembereps["object_type"] = item.ObjectType
		vdmembereps["selector"] = item.Selector
		vdmemberepss = append(vdmemberepss, vdmembereps)
	}
	return vdmemberepss
}
func flattenMapStorageVirtualDriveExtensionRef(p *models.StorageVirtualDriveExtensionRef, d *schema.ResourceData) []map[string]interface{} {

	var virtualdriveextensions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualdriveextension := make(map[string]interface{})
	virtualdriveextension["moid"] = item.Moid
	virtualdriveextension["object_type"] = item.ObjectType
	virtualdriveextension["selector"] = item.Selector

	virtualdriveextensions = append(virtualdriveextensions, virtualdriveextension)
	return virtualdriveextensions
}
func flattenMapMemoryPersistentMemoryRegionRef(p *models.MemoryPersistentMemoryRegionRef, d *schema.ResourceData) []map[string]interface{} {

	var memorypersistentmemoryregions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemoryregion := make(map[string]interface{})
	memorypersistentmemoryregion["moid"] = item.Moid
	memorypersistentmemoryregion["object_type"] = item.ObjectType
	memorypersistentmemoryregion["selector"] = item.Selector

	memorypersistentmemoryregions = append(memorypersistentmemoryregions, memorypersistentmemoryregion)
	return memorypersistentmemoryregions
}
func flattenListOsConfigurationFileRef(p []*models.OsConfigurationFileRef, d *schema.ResourceData) []map[string]interface{} {
	var configurationfiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		configurationfiles := make(map[string]interface{})
		configurationfiles["moid"] = item.Moid
		configurationfiles["object_type"] = item.ObjectType
		configurationfiles["selector"] = item.Selector
		configurationfiless = append(configurationfiless, configurationfiles)
	}
	return configurationfiless
}
func flattenListOnpremUpgradePhase(p []*models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var completedphasess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		completedphases := make(map[string]interface{})
		completedphases["name"] = item.Name
		completedphases["object_type"] = item.ObjectType
		completedphasess = append(completedphasess, completedphases)
	}
	return completedphasess
}
func flattenMapOnpremUpgradePhase(p *models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {

	var currentphases []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	currentphase := make(map[string]interface{})
	currentphase["name"] = item.Name
	currentphase["object_type"] = item.ObjectType

	currentphases = append(currentphases, currentphase)
	return currentphases
}
func flattenMapApplianceImageBundleRef(p *models.ApplianceImageBundleRef, d *schema.ResourceData) []map[string]interface{} {

	var imagebundles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	imagebundle := make(map[string]interface{})
	imagebundle["moid"] = item.Moid
	imagebundle["object_type"] = item.ObjectType
	imagebundle["selector"] = item.Selector

	imagebundles = append(imagebundles, imagebundle)
	return imagebundles
}
func flattenMapEquipmentSystemIoControllerRef(p *models.EquipmentSystemIoControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var equipmentsystemiocontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentsystemiocontroller := make(map[string]interface{})
	equipmentsystemiocontroller["moid"] = item.Moid
	equipmentsystemiocontroller["object_type"] = item.ObjectType
	equipmentsystemiocontroller["selector"] = item.Selector

	equipmentsystemiocontrollers = append(equipmentsystemiocontrollers, equipmentsystemiocontroller)
	return equipmentsystemiocontrollers
}
func flattenListPortGroupRef(p []*models.PortGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var portgroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		portgroups := make(map[string]interface{})
		portgroups["moid"] = item.Moid
		portgroups["object_type"] = item.ObjectType
		portgroups["selector"] = item.Selector
		portgroupss = append(portgroupss, portgroups)
	}
	return portgroupss
}
func flattenMapAdapterUnitRef(p *models.AdapterUnitRef, d *schema.ResourceData) []map[string]interface{} {

	var adapterunits []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	adapterunit := make(map[string]interface{})
	adapterunit["moid"] = item.Moid
	adapterunit["object_type"] = item.ObjectType
	adapterunit["selector"] = item.Selector

	adapterunits = append(adapterunits, adapterunit)
	return adapterunits
}
func flattenMapHyperflexHxResiliencyInfoDt(p *models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {

	var resiliencydetailss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	resiliencydetails := make(map[string]interface{})
	resiliencydetails["data_replication_factor"] = item.DataReplicationFactor
	resiliencydetails["hdd_failures_tolerable"] = item.HddFailuresTolerable
	resiliencydetails["messages"] = item.Messages
	resiliencydetails["node_failures_tolerable"] = item.NodeFailuresTolerable
	resiliencydetails["object_type"] = item.ObjectType
	resiliencydetails["policy_compliance"] = item.PolicyCompliance
	resiliencydetails["resiliency_state"] = item.ResiliencyState
	resiliencydetails["ssd_failures_tolerable"] = item.SsdFailuresTolerable

	resiliencydetailss = append(resiliencydetailss, resiliencydetails)
	return resiliencydetailss
}
func flattenListHyperflexHxZoneResiliencyInfoDt(p []*models.HyperflexHxZoneResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var zoneresiliencylists []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		zoneresiliencylist := make(map[string]interface{})
		zoneresiliencylist["name"] = item.Name
		zoneresiliencylist["object_type"] = item.ObjectType
		zoneresiliencylist["resiliency_info"] = (func(p *models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {

			var resiliencyinfos []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			resiliencyinfo := make(map[string]interface{})
			resiliencyinfo["data_replication_factor"] = item.DataReplicationFactor
			resiliencyinfo["hdd_failures_tolerable"] = item.HddFailuresTolerable
			resiliencyinfo["messages"] = item.Messages
			resiliencyinfo["node_failures_tolerable"] = item.NodeFailuresTolerable
			resiliencyinfo["object_type"] = item.ObjectType
			resiliencyinfo["policy_compliance"] = item.PolicyCompliance
			resiliencyinfo["resiliency_state"] = item.ResiliencyState
			resiliencyinfo["ssd_failures_tolerable"] = item.SsdFailuresTolerable

			resiliencyinfos = append(resiliencyinfos, resiliencyinfo)
			return resiliencyinfos
		})(item.ResiliencyInfo, d)
		zoneresiliencylists = append(zoneresiliencylists, zoneresiliencylist)
	}
	return zoneresiliencylists
}
func flattenMapHclOperatingSystemVendorRef(p *models.HclOperatingSystemVendorRef, d *schema.ResourceData) []map[string]interface{} {

	var vendors []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	vendor := make(map[string]interface{})
	vendor["moid"] = item.Moid
	vendor["object_type"] = item.ObjectType
	vendor["selector"] = item.Selector

	vendors = append(vendors, vendor)
	return vendors
}
func flattenMapMemoryPersistentMemoryConfigResultRef(p *models.MemoryPersistentMemoryConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var memorypersistentmemoryconfigresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemoryconfigresult := make(map[string]interface{})
	memorypersistentmemoryconfigresult["moid"] = item.Moid
	memorypersistentmemoryconfigresult["object_type"] = item.ObjectType
	memorypersistentmemoryconfigresult["selector"] = item.Selector

	memorypersistentmemoryconfigresults = append(memorypersistentmemoryconfigresults, memorypersistentmemoryconfigresult)
	return memorypersistentmemoryconfigresults
}
func flattenMapStorageProtectionGroupRef(p *models.StorageProtectionGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var protectiongroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	protectiongroup := make(map[string]interface{})
	protectiongroup["moid"] = item.Moid
	protectiongroup["object_type"] = item.ObjectType
	protectiongroup["selector"] = item.Selector

	protectiongroups = append(protectiongroups, protectiongroup)
	return protectiongroups
}
func flattenMapStorageGenericArrayRef(p *models.StorageGenericArrayRef, d *schema.ResourceData) []map[string]interface{} {

	var storagearrays []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagearray := make(map[string]interface{})
	storagearray["moid"] = item.Moid
	storagearray["object_type"] = item.ObjectType
	storagearray["selector"] = item.Selector

	storagearrays = append(storagearrays, storagearray)
	return storagearrays
}
func flattenMapMemoryPersistentMemoryConfigurationRef(p *models.MemoryPersistentMemoryConfigurationRef, d *schema.ResourceData) []map[string]interface{} {

	var memorypersistentmemoryconfigurations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memorypersistentmemoryconfiguration := make(map[string]interface{})
	memorypersistentmemoryconfiguration["moid"] = item.Moid
	memorypersistentmemoryconfiguration["object_type"] = item.ObjectType
	memorypersistentmemoryconfiguration["selector"] = item.Selector

	memorypersistentmemoryconfigurations = append(memorypersistentmemoryconfigurations, memorypersistentmemoryconfiguration)
	return memorypersistentmemoryconfigurations
}
func flattenListMemoryPersistentMemoryNamespaceConfigResultRef(p []*models.MemoryPersistentMemoryNamespaceConfigResultRef, d *schema.ResourceData) []map[string]interface{} {
	var persistentmemorynamespaceconfigresultss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		persistentmemorynamespaceconfigresults := make(map[string]interface{})
		persistentmemorynamespaceconfigresults["moid"] = item.Moid
		persistentmemorynamespaceconfigresults["object_type"] = item.ObjectType
		persistentmemorynamespaceconfigresults["selector"] = item.Selector
		persistentmemorynamespaceconfigresultss = append(persistentmemorynamespaceconfigresultss, persistentmemorynamespaceconfigresults)
	}
	return persistentmemorynamespaceconfigresultss
}
func flattenMapPciSwitchRef(p *models.PciSwitchRef, d *schema.ResourceData) []map[string]interface{} {

	var pciswitchs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	pciswitch := make(map[string]interface{})
	pciswitch["moid"] = item.Moid
	pciswitch["object_type"] = item.ObjectType
	pciswitch["selector"] = item.Selector

	pciswitchs = append(pciswitchs, pciswitch)
	return pciswitchs
}
func flattenListEquipmentSwitchCardRef(p []*models.EquipmentSwitchCardRef, d *schema.ResourceData) []map[string]interface{} {
	var cardss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		cards := make(map[string]interface{})
		cards["moid"] = item.Moid
		cards["object_type"] = item.ObjectType
		cards["selector"] = item.Selector
		cardss = append(cardss, cards)
	}
	return cardss
}
func flattenListEquipmentFanModuleRef(p []*models.EquipmentFanModuleRef, d *schema.ResourceData) []map[string]interface{} {
	var fanmoduless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		fanmodules := make(map[string]interface{})
		fanmodules["moid"] = item.Moid
		fanmodules["object_type"] = item.ObjectType
		fanmodules["selector"] = item.Selector
		fanmoduless = append(fanmoduless, fanmodules)
	}
	return fanmoduless
}
func flattenMapManagementControllerRef(p *models.ManagementControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var managementcontollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	managementcontoller := make(map[string]interface{})
	managementcontoller["moid"] = item.Moid
	managementcontoller["object_type"] = item.ObjectType
	managementcontoller["selector"] = item.Selector

	managementcontollers = append(managementcontollers, managementcontoller)
	return managementcontollers
}
func flattenMapManagementEntityRef(p *models.ManagementEntityRef, d *schema.ResourceData) []map[string]interface{} {

	var managemententitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	managemententity := make(map[string]interface{})
	managemententity["moid"] = item.Moid
	managemententity["object_type"] = item.ObjectType
	managemententity["selector"] = item.Selector

	managemententitys = append(managemententitys, managemententity)
	return managemententitys
}
func flattenListEquipmentPsuRef(p []*models.EquipmentPsuRef, d *schema.ResourceData) []map[string]interface{} {
	var psuss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		psus := make(map[string]interface{})
		psus["moid"] = item.Moid
		psus["object_type"] = item.ObjectType
		psus["selector"] = item.Selector
		psuss = append(psuss, psus)
	}
	return psuss
}
func flattenMapTopSystemRef(p *models.TopSystemRef, d *schema.ResourceData) []map[string]interface{} {

	var topsystems []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	topsystem := make(map[string]interface{})
	topsystem["moid"] = item.Moid
	topsystem["object_type"] = item.ObjectType
	topsystem["selector"] = item.Selector

	topsystems = append(topsystems, topsystem)
	return topsystems
}
func flattenMapFirmwareRunningFirmwareRef(p *models.FirmwareRunningFirmwareRef, d *schema.ResourceData) []map[string]interface{} {

	var ucsmrunningfirmwares []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ucsmrunningfirmware := make(map[string]interface{})
	ucsmrunningfirmware["moid"] = item.Moid
	ucsmrunningfirmware["object_type"] = item.ObjectType
	ucsmrunningfirmware["selector"] = item.Selector

	ucsmrunningfirmwares = append(ucsmrunningfirmwares, ucsmrunningfirmware)
	return ucsmrunningfirmwares
}
func flattenMapCommIPV4Interface(p *models.CommIPV4Interface, d *schema.ResourceData) []map[string]interface{} {

	var nodeipv4configs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nodeipv4config := make(map[string]interface{})
	nodeipv4config["gateway"] = item.Gateway
	nodeipv4config["ip_address"] = item.IPAddress
	nodeipv4config["netmask"] = item.Netmask
	nodeipv4config["object_type"] = item.ObjectType

	nodeipv4configs = append(nodeipv4configs, nodeipv4config)
	return nodeipv4configs
}
func flattenMapComputeBladeRef(p *models.ComputeBladeRef, d *schema.ResourceData) []map[string]interface{} {

	var computeblades []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computeblade := make(map[string]interface{})
	computeblade["moid"] = item.Moid
	computeblade["object_type"] = item.ObjectType
	computeblade["selector"] = item.Selector

	computeblades = append(computeblades, computeblade)
	return computeblades
}
func flattenListAdapterExtEthInterfaceRef(p []*models.AdapterExtEthInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var extethifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		extethifs := make(map[string]interface{})
		extethifs["moid"] = item.Moid
		extethifs["object_type"] = item.ObjectType
		extethifs["selector"] = item.Selector
		extethifss = append(extethifss, extethifs)
	}
	return extethifss
}
func flattenListAdapterHostEthInterfaceRef(p []*models.AdapterHostEthInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var hostethifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hostethifs := make(map[string]interface{})
		hostethifs["moid"] = item.Moid
		hostethifs["object_type"] = item.ObjectType
		hostethifs["selector"] = item.Selector
		hostethifss = append(hostethifss, hostethifs)
	}
	return hostethifss
}
func flattenListAdapterHostFcInterfaceRef(p []*models.AdapterHostFcInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var hostfcifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hostfcifs := make(map[string]interface{})
		hostfcifs["moid"] = item.Moid
		hostfcifs["object_type"] = item.ObjectType
		hostfcifs["selector"] = item.Selector
		hostfcifss = append(hostfcifss, hostfcifs)
	}
	return hostfcifss
}
func flattenListAdapterHostIscsiInterfaceRef(p []*models.AdapterHostIscsiInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var hostiscsiifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hostiscsiifs := make(map[string]interface{})
		hostiscsiifs["moid"] = item.Moid
		hostiscsiifs["object_type"] = item.ObjectType
		hostiscsiifs["selector"] = item.Selector
		hostiscsiifss = append(hostiscsiifss, hostiscsiifs)
	}
	return hostiscsiifss
}
func flattenMapEquipmentLocatorLedRef(p *models.EquipmentLocatorLedRef, d *schema.ResourceData) []map[string]interface{} {

	var locatorleds []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	locatorled := make(map[string]interface{})
	locatorled["moid"] = item.Moid
	locatorled["object_type"] = item.ObjectType
	locatorled["selector"] = item.Selector

	locatorleds = append(locatorleds, locatorled)
	return locatorleds
}
func flattenListStoragePhysicalDiskExtensionRef(p []*models.StoragePhysicalDiskExtensionRef, d *schema.ResourceData) []map[string]interface{} {
	var physicaldiskextensionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		physicaldiskextensions := make(map[string]interface{})
		physicaldiskextensions["moid"] = item.Moid
		physicaldiskextensions["object_type"] = item.ObjectType
		physicaldiskextensions["selector"] = item.Selector
		physicaldiskextensionss = append(physicaldiskextensionss, physicaldiskextensions)
	}
	return physicaldiskextensionss
}
func flattenListFirmwareRunningFirmwareRef(p []*models.FirmwareRunningFirmwareRef, d *schema.ResourceData) []map[string]interface{} {
	var runningfirmwares []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		runningfirmware := make(map[string]interface{})
		runningfirmware["moid"] = item.Moid
		runningfirmware["object_type"] = item.ObjectType
		runningfirmware["selector"] = item.Selector
		runningfirmwares = append(runningfirmwares, runningfirmware)
	}
	return runningfirmwares
}
func flattenListStorageSasPortRef(p []*models.StorageSasPortRef, d *schema.ResourceData) []map[string]interface{} {
	var sasportss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sasports := make(map[string]interface{})
		sasports["moid"] = item.Moid
		sasports["object_type"] = item.ObjectType
		sasports["selector"] = item.Selector
		sasportss = append(sasportss, sasports)
	}
	return sasportss
}
func flattenMapStorageEnclosureRef(p *models.StorageEnclosureRef, d *schema.ResourceData) []map[string]interface{} {

	var storageenclosures []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageenclosure := make(map[string]interface{})
	storageenclosure["moid"] = item.Moid
	storageenclosure["object_type"] = item.ObjectType
	storageenclosure["selector"] = item.Selector

	storageenclosures = append(storageenclosures, storageenclosure)
	return storageenclosures
}
func flattenMapEquipmentRackEnclosureRef(p *models.EquipmentRackEnclosureRef, d *schema.ResourceData) []map[string]interface{} {

	var equipmentrackenclosures []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentrackenclosure := make(map[string]interface{})
	equipmentrackenclosure["moid"] = item.Moid
	equipmentrackenclosure["object_type"] = item.ObjectType
	equipmentrackenclosure["selector"] = item.Selector

	equipmentrackenclosures = append(equipmentrackenclosures, equipmentrackenclosure)
	return equipmentrackenclosures
}
func flattenListWorkflowTaskRetryInfo(p []*models.WorkflowTaskRetryInfo, d *schema.ResourceData) []map[string]interface{} {
	var taskinstidlists []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		taskinstidlist := make(map[string]interface{})
		taskinstidlist["object_type"] = item.ObjectType
		taskinstidlist["status"] = item.Status
		taskinstidlist["task_inst_id"] = item.TaskInstID
		taskinstidlists = append(taskinstidlists, taskinstidlist)
	}
	return taskinstidlists
}
func flattenListUcsdConnectorPack(p []*models.UcsdConnectorPack, d *schema.ResourceData) []map[string]interface{} {
	var connectorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		connectors := make(map[string]interface{})
		connectors["connector_feature"] = item.ConnectorFeature
		connectors["dependency_names"] = item.DependencyNames
		connectors["downloaded_version"] = item.DownloadedVersion
		connectors["name"] = item.Name
		connectors["object_type"] = item.ObjectType
		connectors["services"] = item.Services
		connectors["state"] = item.State
		connectors["version"] = item.Version
		connectorss = append(connectorss, connectors)
	}
	return connectorss
}
func flattenMapComputeBoardRef(p *models.ComputeBoardRef, d *schema.ResourceData) []map[string]interface{} {

	var computeboards []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	computeboard := make(map[string]interface{})
	computeboard["moid"] = item.Moid
	computeboard["object_type"] = item.ObjectType
	computeboard["selector"] = item.Selector

	computeboards = append(computeboards, computeboard)
	return computeboards
}
func flattenListMemoryPersistentMemoryUnitRef(p []*models.MemoryPersistentMemoryUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var persistentmemoryunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		persistentmemoryunits := make(map[string]interface{})
		persistentmemoryunits["moid"] = item.Moid
		persistentmemoryunits["object_type"] = item.ObjectType
		persistentmemoryunits["selector"] = item.Selector
		persistentmemoryunitss = append(persistentmemoryunitss, persistentmemoryunits)
	}
	return persistentmemoryunitss
}
func flattenListMemoryUnitRef(p []*models.MemoryUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var unitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		units := make(map[string]interface{})
		units["moid"] = item.Moid
		units["object_type"] = item.ObjectType
		units["selector"] = item.Selector
		unitss = append(unitss, units)
	}
	return unitss
}
func flattenListMemoryPersistentMemoryNamespaceRef(p []*models.MemoryPersistentMemoryNamespaceRef, d *schema.ResourceData) []map[string]interface{} {
	var persistentmemorynamespacess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		persistentmemorynamespaces := make(map[string]interface{})
		persistentmemorynamespaces["moid"] = item.Moid
		persistentmemorynamespaces["object_type"] = item.ObjectType
		persistentmemorynamespaces["selector"] = item.Selector
		persistentmemorynamespacess = append(persistentmemorynamespacess, persistentmemorynamespaces)
	}
	return persistentmemorynamespacess
}
func flattenListEquipmentFanRef(p []*models.EquipmentFanRef, d *schema.ResourceData) []map[string]interface{} {
	var fanss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		fans := make(map[string]interface{})
		fans["moid"] = item.Moid
		fans["object_type"] = item.ObjectType
		fans["selector"] = item.Selector
		fanss = append(fanss, fans)
	}
	return fanss
}
func flattenMapNetworkElementRef(p *models.NetworkElementRef, d *schema.ResourceData) []map[string]interface{} {

	var networkelements []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	networkelement := make(map[string]interface{})
	networkelement["moid"] = item.Moid
	networkelement["object_type"] = item.ObjectType
	networkelement["selector"] = item.Selector

	networkelements = append(networkelements, networkelement)
	return networkelements
}
func flattenListStorageFlexFlashControllerPropsRef(p []*models.StorageFlexFlashControllerPropsRef, d *schema.ResourceData) []map[string]interface{} {
	var flexflashcontrollerpropss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		flexflashcontrollerprops := make(map[string]interface{})
		flexflashcontrollerprops["moid"] = item.Moid
		flexflashcontrollerprops["object_type"] = item.ObjectType
		flexflashcontrollerprops["selector"] = item.Selector
		flexflashcontrollerpropss = append(flexflashcontrollerpropss, flexflashcontrollerprops)
	}
	return flexflashcontrollerpropss
}
func flattenListStorageFlexFlashPhysicalDriveRef(p []*models.StorageFlexFlashPhysicalDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var flexflashphysicaldrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		flexflashphysicaldrives := make(map[string]interface{})
		flexflashphysicaldrives["moid"] = item.Moid
		flexflashphysicaldrives["object_type"] = item.ObjectType
		flexflashphysicaldrives["selector"] = item.Selector
		flexflashphysicaldrivess = append(flexflashphysicaldrivess, flexflashphysicaldrives)
	}
	return flexflashphysicaldrivess
}
func flattenListStorageFlexFlashVirtualDriveRef(p []*models.StorageFlexFlashVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var flexflashvirtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		flexflashvirtualdrives := make(map[string]interface{})
		flexflashvirtualdrives["moid"] = item.Moid
		flexflashvirtualdrives["object_type"] = item.ObjectType
		flexflashvirtualdrives["selector"] = item.Selector
		flexflashvirtualdrivess = append(flexflashvirtualdrivess, flexflashvirtualdrives)
	}
	return flexflashvirtualdrivess
}
func flattenListInventoryGenericInventoryRef(p []*models.InventoryGenericInventoryRef, d *schema.ResourceData) []map[string]interface{} {
	var genericinventorys []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		genericinventory := make(map[string]interface{})
		genericinventory["moid"] = item.Moid
		genericinventory["object_type"] = item.ObjectType
		genericinventory["selector"] = item.Selector
		genericinventorys = append(genericinventorys, genericinventory)
	}
	return genericinventorys
}
func flattenMapLicenseAccountLicenseDataRef(p *models.LicenseAccountLicenseDataRef, d *schema.ResourceData) []map[string]interface{} {

	var accountlicensedatas []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	accountlicensedata := make(map[string]interface{})
	accountlicensedata["moid"] = item.Moid
	accountlicensedata["object_type"] = item.ObjectType
	accountlicensedata["selector"] = item.Selector

	accountlicensedatas = append(accountlicensedatas, accountlicensedata)
	return accountlicensedatas
}
func flattenMapOnpremSchedule(p *models.OnpremSchedule, d *schema.ResourceData) []map[string]interface{} {

	var schedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	schedule := make(map[string]interface{})
	schedule["day_of_month"] = item.DayOfMonth
	schedule["day_of_week"] = item.DayOfWeek
	schedule["month_of_year"] = item.MonthOfYear
	schedule["object_type"] = item.ObjectType
	schedule["repeat_interval"] = item.RepeatInterval
	schedule["time_of_day"] = item.TimeOfDay
	schedule["time_zone"] = item.TimeZone
	schedule["week_of_month"] = item.WeekOfMonth

	schedules = append(schedules, schedule)
	return schedules
}
func flattenListManagementInterfaceRef(p []*models.ManagementInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var managementinterfacess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		managementinterfaces := make(map[string]interface{})
		managementinterfaces["moid"] = item.Moid
		managementinterfaces["object_type"] = item.ObjectType
		managementinterfaces["selector"] = item.Selector
		managementinterfacess = append(managementinterfacess, managementinterfaces)
	}
	return managementinterfacess
}
func flattenMapStorageSasExpanderRef(p *models.StorageSasExpanderRef, d *schema.ResourceData) []map[string]interface{} {

	var storagesasexpanders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storagesasexpander := make(map[string]interface{})
	storagesasexpander["moid"] = item.Moid
	storagesasexpander["object_type"] = item.ObjectType
	storagesasexpander["selector"] = item.Selector

	storagesasexpanders = append(storagesasexpanders, storagesasexpander)
	return storagesasexpanders
}
func flattenMapStorageCapacity(p *models.StorageCapacity, d *schema.ResourceData) []map[string]interface{} {

	var storageutilizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageutilization := make(map[string]interface{})
	storageutilization["available"] = item.Available
	storageutilization["free"] = item.Free
	storageutilization["object_type"] = item.ObjectType
	storageutilization["total"] = item.Total
	storageutilization["used"] = item.Used

	storageutilizations = append(storageutilizations, storageutilization)
	return storageutilizations
}
func flattenListOnpremImagePackage(p []*models.OnpremImagePackage, d *schema.ResourceData) []map[string]interface{} {
	var ansiblepackagess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		ansiblepackages := make(map[string]interface{})
		ansiblepackages["file_path"] = item.FilePath
		ansiblepackages["file_sha"] = item.FileSha
		ansiblepackages["file_size"] = item.FileSize
		ansiblepackages["file_time"] = item.FileTime
		ansiblepackages["filename"] = item.Filename
		ansiblepackages["name"] = item.Name
		ansiblepackages["object_type"] = item.ObjectType
		ansiblepackages["package_type"] = item.PackageType
		ansiblepackages["version"] = item.Version
		ansiblepackagess = append(ansiblepackagess, ansiblepackages)
	}
	return ansiblepackagess
}
func flattenListHyperflexConfigResultEntryRef(p []*models.HyperflexConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var resultentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resultentries := make(map[string]interface{})
		resultentries["moid"] = item.Moid
		resultentries["object_type"] = item.ObjectType
		resultentries["selector"] = item.Selector
		resultentriess = append(resultentriess, resultentries)
	}
	return resultentriess
}
func flattenMapEquipmentFanModuleRef(p *models.EquipmentFanModuleRef, d *schema.ResourceData) []map[string]interface{} {

	var equipmentfanmodules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentfanmodule := make(map[string]interface{})
	equipmentfanmodule["moid"] = item.Moid
	equipmentfanmodule["object_type"] = item.ObjectType
	equipmentfanmodule["selector"] = item.Selector

	equipmentfanmodules = append(equipmentfanmodules, equipmentfanmodule)
	return equipmentfanmodules
}
func flattenListIamGroupPermissionToRoles(p []*models.IamGroupPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var grouppermissionroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		grouppermissionroles := make(map[string]interface{})
		grouppermissionroles["group"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {

			var groups []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			group := make(map[string]interface{})
			group["moid"] = item.Moid
			group["object_type"] = item.ObjectType

			groups = append(groups, group)
			return groups
		})(item.Group, d)
		grouppermissionroles["object_type"] = item.ObjectType
		grouppermissionroles["orgs"] = (func(p []*models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var orgss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				orgs := make(map[string]interface{})
				orgs["moid"] = item.Moid
				orgs["object_type"] = item.ObjectType
				orgss = append(orgss, orgs)
			}
			return orgss
		})(item.Orgs, d)
		grouppermissionroless = append(grouppermissionroless, grouppermissionroles)
	}
	return grouppermissionroless
}
func flattenMapResourceMembershipHolderRef(p *models.ResourceMembershipHolderRef, d *schema.ResourceData) []map[string]interface{} {

	var holders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	holder := make(map[string]interface{})
	holder["moid"] = item.Moid
	holder["object_type"] = item.ObjectType
	holder["selector"] = item.Selector

	holders = append(holders, holder)
	return holders
}
func flattenMapGraphicsCardRef(p *models.GraphicsCardRef, d *schema.ResourceData) []map[string]interface{} {

	var graphicscards []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	graphicscard := make(map[string]interface{})
	graphicscard["moid"] = item.Moid
	graphicscard["object_type"] = item.ObjectType
	graphicscard["selector"] = item.Selector

	graphicscards = append(graphicscards, graphicscard)
	return graphicscards
}
func flattenListNiaapiRevisionInfo(p []*models.NiaapiRevisionInfo, d *schema.ResourceData) []map[string]interface{} {
	var revisioninfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		revisioninfo := make(map[string]interface{})
		revisioninfo["date_published"] = item.DatePublished
		revisioninfo["object_type"] = item.ObjectType
		revisioninfo["revision_comment"] = item.RevisionComment
		revisioninfo["revision_no"] = item.RevisionNo
		revisioninfos = append(revisioninfos, revisioninfo)
	}
	return revisioninfos
}
func flattenListComputeIPAddress(p []*models.ComputeIPAddress, d *schema.ResourceData) []map[string]interface{} {
	var kvmipaddressess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		kvmipaddresses := make(map[string]interface{})
		kvmipaddresses["address"] = item.Address
		kvmipaddresses["category"] = item.Category
		kvmipaddresses["default_gateway"] = item.DefaultGateway
		kvmipaddresses["dn"] = item.Dn
		kvmipaddresses["http_port"] = item.HTTPPort
		kvmipaddresses["https_port"] = item.HTTPSPort
		kvmipaddresses["kvm_port"] = item.KvmPort
		kvmipaddresses["name"] = item.Name
		kvmipaddresses["object_type"] = item.ObjectType
		kvmipaddresses["subnet"] = item.Subnet
		kvmipaddresses["type"] = item.Type
		kvmipaddressess = append(kvmipaddressess, kvmipaddresses)
	}
	return kvmipaddressess
}
func flattenMapStoragePhysicalDiskRef(p *models.StoragePhysicalDiskRef, d *schema.ResourceData) []map[string]interface{} {

	var physicaldisks []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	physicaldisk := make(map[string]interface{})
	physicaldisk["moid"] = item.Moid
	physicaldisk["object_type"] = item.ObjectType
	physicaldisk["selector"] = item.Selector

	physicaldisks = append(physicaldisks, physicaldisk)
	return physicaldisks
}
func flattenMapRecoveryBackupProfileRef(p *models.RecoveryBackupProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var backupprofiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	backupprofile := make(map[string]interface{})
	backupprofile["moid"] = item.Moid
	backupprofile["object_type"] = item.ObjectType
	backupprofile["selector"] = item.Selector

	backupprofiles = append(backupprofiles, backupprofile)
	return backupprofiles
}
func flattenMapRecoveryOnDemandBackupRef(p *models.RecoveryOnDemandBackupRef, d *schema.ResourceData) []map[string]interface{} {

	var nr0ondemandbackups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	nr0ondemandbackup := make(map[string]interface{})
	nr0ondemandbackup["moid"] = item.Moid
	nr0ondemandbackup["object_type"] = item.ObjectType
	nr0ondemandbackup["selector"] = item.Selector

	nr0ondemandbackups = append(nr0ondemandbackups, nr0ondemandbackup)
	return nr0ondemandbackups
}
func flattenListRecoveryConfigResultEntryRef(p []*models.RecoveryConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var resultentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resultentries := make(map[string]interface{})
		resultentries["moid"] = item.Moid
		resultentries["object_type"] = item.ObjectType
		resultentries["selector"] = item.Selector
		resultentriess = append(resultentriess, resultentries)
	}
	return resultentriess
}
func flattenMapNiatelemetryNiaInventoryRef(p *models.NiatelemetryNiaInventoryRef, d *schema.ResourceData) []map[string]interface{} {

	var devices []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	device := make(map[string]interface{})
	device["moid"] = item.Moid
	device["object_type"] = item.ObjectType
	device["selector"] = item.Selector

	devices = append(devices, device)
	return devices
}
func flattenListIamResourcePermissionRef(p []*models.IamResourcePermissionRef, d *schema.ResourceData) []map[string]interface{} {
	var resourcepermissionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resourcepermissions := make(map[string]interface{})
		resourcepermissions["moid"] = item.Moid
		resourcepermissions["object_type"] = item.ObjectType
		resourcepermissions["selector"] = item.Selector
		resourcepermissionss = append(resourcepermissionss, resourcepermissions)
	}
	return resourcepermissionss
}
func flattenMapForecastDefinitionRef(p *models.ForecastDefinitionRef, d *schema.ResourceData) []map[string]interface{} {

	var forecastdefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	forecastdef := make(map[string]interface{})
	forecastdef["moid"] = item.Moid
	forecastdef["object_type"] = item.ObjectType
	forecastdef["selector"] = item.Selector

	forecastdefs = append(forecastdefs, forecastdef)
	return forecastdefs
}
func flattenMapForecastModel(p *models.ForecastModel, d *schema.ResourceData) []map[string]interface{} {

	var models []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	model := make(map[string]interface{})
	model["accuracy"] = item.Accuracy
	model["model_data"] = item.ModelData
	model["model_type"] = item.ModelType
	model["object_type"] = item.ObjectType

	models = append(models, model)
	return models
}
func flattenMapIamAppRegistrationRef(p *models.IamAppRegistrationRef, d *schema.ResourceData) []map[string]interface{} {

	var appregistrations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	appregistration := make(map[string]interface{})
	appregistration["moid"] = item.Moid
	appregistration["object_type"] = item.ObjectType
	appregistration["selector"] = item.Selector

	appregistrations = append(appregistrations, appregistration)
	return appregistrations
}
func flattenMapIamClientMeta(p *models.IamClientMeta, d *schema.ResourceData) []map[string]interface{} {

	var usermetas []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	usermeta := make(map[string]interface{})
	usermeta["device_model"] = item.DeviceModel
	usermeta["object_type"] = item.ObjectType
	usermeta["user_agent"] = item.UserAgent

	usermetas = append(usermetas, usermeta)
	return usermetas
}
func flattenMapIaasUcsdInfoRef(p *models.IaasUcsdInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var guids []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	guid := make(map[string]interface{})
	guid["moid"] = item.Moid
	guid["object_type"] = item.ObjectType
	guid["selector"] = item.Selector

	guids = append(guids, guid)
	return guids
}
func flattenListIaasLicenseKeysInfo(p []*models.IaasLicenseKeysInfo, d *schema.ResourceData) []map[string]interface{} {
	var licensekeysinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		licensekeysinfo := make(map[string]interface{})
		licensekeysinfo["count"] = item.Count
		licensekeysinfo["expiration_date"] = item.ExpirationDate
		licensekeysinfo["license_id"] = item.LicenseID
		licensekeysinfo["object_type"] = item.ObjectType
		licensekeysinfo["pid"] = item.Pid
		licensekeysinfos = append(licensekeysinfos, licensekeysinfo)
	}
	return licensekeysinfos
}
func flattenListIaasLicenseUtilizationInfo(p []*models.IaasLicenseUtilizationInfo, d *schema.ResourceData) []map[string]interface{} {
	var licenseutilizationinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		licenseutilizationinfo := make(map[string]interface{})
		licenseutilizationinfo["actual_used"] = item.ActualUsed
		licenseutilizationinfo["label"] = item.Label
		licenseutilizationinfo["licensed_limit"] = item.LicensedLimit
		licenseutilizationinfo["object_type"] = item.ObjectType
		licenseutilizationinfo["sku"] = item.Sku
		licenseutilizationinfos = append(licenseutilizationinfos, licenseutilizationinfo)
	}
	return licenseutilizationinfos
}
func flattenMapPortGroupRef(p *models.PortGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var portgroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	portgroup := make(map[string]interface{})
	portgroup["moid"] = item.Moid
	portgroup["object_type"] = item.ObjectType
	portgroup["selector"] = item.Selector

	portgroups = append(portgroups, portgroup)
	return portgroups
}
func flattenListPolicyinventoryJobInfo(p []*models.PolicyinventoryJobInfo, d *schema.ResourceData) []map[string]interface{} {
	var jobinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		jobinfo := make(map[string]interface{})
		jobinfo["execution_status"] = item.ExecutionStatus
		jobinfo["last_scheduled_time"] = item.LastScheduledTime
		jobinfo["object_type"] = item.ObjectType
		jobinfo["policy_id"] = item.PolicyID
		jobinfo["policy_name"] = item.PolicyName
		jobinfos = append(jobinfos, jobinfo)
	}
	return jobinfos
}
func flattenListEtherPhysicalPortRef(p []*models.EtherPhysicalPortRef, d *schema.ResourceData) []map[string]interface{} {
	var ethernetportss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		ethernetports := make(map[string]interface{})
		ethernetports["moid"] = item.Moid
		ethernetports["object_type"] = item.ObjectType
		ethernetports["selector"] = item.Selector
		ethernetportss = append(ethernetportss, ethernetports)
	}
	return ethernetportss
}
func flattenMapIamSessionRef(p *models.IamSessionRef, d *schema.ResourceData) []map[string]interface{} {

	var sessionss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sessions := make(map[string]interface{})
	sessions["moid"] = item.Moid
	sessions["object_type"] = item.ObjectType
	sessions["selector"] = item.Selector

	sessionss = append(sessionss, sessions)
	return sessionss
}
func flattenMapStorageVirtualDriveRef(p *models.StorageVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {

	var virtualdrives []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	virtualdrive := make(map[string]interface{})
	virtualdrive["moid"] = item.Moid
	virtualdrive["object_type"] = item.ObjectType
	virtualdrive["selector"] = item.Selector

	virtualdrives = append(virtualdrives, virtualdrive)
	return virtualdrives
}
func flattenListWorkflowDynamicWorkflowActionTaskList(p []*models.WorkflowDynamicWorkflowActionTaskList, d *schema.ResourceData) []map[string]interface{} {
	var workflowactiontasklistss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		workflowactiontasklists := make(map[string]interface{})
		workflowactiontasklists["action"] = item.Action
		workflowactiontasklists["object_type"] = item.ObjectType
		workflowactiontasklists["tasks"] = item.Tasks
		workflowactiontasklistss = append(workflowactiontasklistss, workflowactiontasklists)
	}
	return workflowactiontasklistss
}
func flattenListApplianceKeyValuePair(p []*models.ApplianceKeyValuePair, d *schema.ResourceData) []map[string]interface{} {
	var capabilitiess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		capabilities := make(map[string]interface{})
		capabilities["key"] = item.Key
		capabilities["object_type"] = item.ObjectType
		capabilities["value"] = item.Value
		capabilitiess = append(capabilitiess, capabilities)
	}
	return capabilitiess
}
func flattenListMemoryPersistentMemoryRegionRef(p []*models.MemoryPersistentMemoryRegionRef, d *schema.ResourceData) []map[string]interface{} {
	var persistentmemoryregionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		persistentmemoryregions := make(map[string]interface{})
		persistentmemoryregions["moid"] = item.Moid
		persistentmemoryregions["object_type"] = item.ObjectType
		persistentmemoryregions["selector"] = item.Selector
		persistentmemoryregionss = append(persistentmemoryregionss, persistentmemoryregions)
	}
	return persistentmemoryregionss
}
func flattenMapBiosUnitRef(p *models.BiosUnitRef, d *schema.ResourceData) []map[string]interface{} {

	var biosunits []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	biosunit := make(map[string]interface{})
	biosunit["moid"] = item.Moid
	biosunit["object_type"] = item.ObjectType
	biosunit["selector"] = item.Selector

	biosunits = append(biosunits, biosunit)
	return biosunits
}
func flattenListNetworkElementRef(p []*models.NetworkElementRef, d *schema.ResourceData) []map[string]interface{} {
	var networkelementss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		networkelements := make(map[string]interface{})
		networkelements["moid"] = item.Moid
		networkelements["object_type"] = item.ObjectType
		networkelements["selector"] = item.Selector
		networkelementss = append(networkelementss, networkelements)
	}
	return networkelementss
}
func flattenMapComputeServerConfig(p *models.ComputeServerConfig, d *schema.ResourceData) []map[string]interface{} {

	var serverconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	serverconfig := make(map[string]interface{})
	serverconfig["asset_tag"] = item.AssetTag
	serverconfig["object_type"] = item.ObjectType
	serverconfig["user_label"] = item.UserLabel

	serverconfigs = append(serverconfigs, serverconfig)
	return serverconfigs
}
func flattenListStoragePhysicalDiskRef(p []*models.StoragePhysicalDiskRef, d *schema.ResourceData) []map[string]interface{} {
	var physicaldiskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		physicaldisks := make(map[string]interface{})
		physicaldisks["moid"] = item.Moid
		physicaldisks["object_type"] = item.ObjectType
		physicaldisks["selector"] = item.Selector
		physicaldiskss = append(physicaldiskss, physicaldisks)
	}
	return physicaldiskss
}
func flattenListStorageVirtualDriveExtensionRef(p []*models.StorageVirtualDriveExtensionRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualdriveextensionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		virtualdriveextensions := make(map[string]interface{})
		virtualdriveextensions["moid"] = item.Moid
		virtualdriveextensions["object_type"] = item.ObjectType
		virtualdriveextensions["selector"] = item.Selector
		virtualdriveextensionss = append(virtualdriveextensionss, virtualdriveextensions)
	}
	return virtualdriveextensionss
}
func flattenListStorageVirtualDriveRef(p []*models.StorageVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var virtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		virtualdrives := make(map[string]interface{})
		virtualdrives["moid"] = item.Moid
		virtualdrives["object_type"] = item.ObjectType
		virtualdrives["selector"] = item.Selector
		virtualdrivess = append(virtualdrivess, virtualdrives)
	}
	return virtualdrivess
}
func flattenListEquipmentTpmRef(p []*models.EquipmentTpmRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmenttpmss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmenttpms := make(map[string]interface{})
		equipmenttpms["moid"] = item.Moid
		equipmenttpms["object_type"] = item.ObjectType
		equipmenttpms["selector"] = item.Selector
		equipmenttpmss = append(equipmenttpmss, equipmenttpms)
	}
	return equipmenttpmss
}
func flattenListGraphicsCardRef(p []*models.GraphicsCardRef, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		graphicscards := make(map[string]interface{})
		graphicscards["moid"] = item.Moid
		graphicscards["object_type"] = item.ObjectType
		graphicscards["selector"] = item.Selector
		graphicscardss = append(graphicscardss, graphicscards)
	}
	return graphicscardss
}
func flattenListMemoryArrayRef(p []*models.MemoryArrayRef, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memoryarrays := make(map[string]interface{})
		memoryarrays["moid"] = item.Moid
		memoryarrays["object_type"] = item.ObjectType
		memoryarrays["selector"] = item.Selector
		memoryarrayss = append(memoryarrayss, memoryarrays)
	}
	return memoryarrayss
}
func flattenListPciCoprocessorCardRef(p []*models.PciCoprocessorCardRef, d *schema.ResourceData) []map[string]interface{} {
	var pcicoprocessorcardss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pcicoprocessorcards := make(map[string]interface{})
		pcicoprocessorcards["moid"] = item.Moid
		pcicoprocessorcards["object_type"] = item.ObjectType
		pcicoprocessorcards["selector"] = item.Selector
		pcicoprocessorcardss = append(pcicoprocessorcardss, pcicoprocessorcards)
	}
	return pcicoprocessorcardss
}
func flattenListPciSwitchRef(p []*models.PciSwitchRef, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pciswitch := make(map[string]interface{})
		pciswitch["moid"] = item.Moid
		pciswitch["object_type"] = item.ObjectType
		pciswitch["selector"] = item.Selector
		pciswitchs = append(pciswitchs, pciswitch)
	}
	return pciswitchs
}
func flattenListProcessorUnitRef(p []*models.ProcessorUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var processorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		processors := make(map[string]interface{})
		processors["moid"] = item.Moid
		processors["object_type"] = item.ObjectType
		processors["selector"] = item.Selector
		processorss = append(processorss, processors)
	}
	return processorss
}
func flattenListSecurityUnitRef(p []*models.SecurityUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var securityunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		securityunits := make(map[string]interface{})
		securityunits["moid"] = item.Moid
		securityunits["object_type"] = item.ObjectType
		securityunits["selector"] = item.Selector
		securityunitss = append(securityunitss, securityunits)
	}
	return securityunitss
}
func flattenListStorageControllerRef(p []*models.StorageControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storagecontrollers := make(map[string]interface{})
		storagecontrollers["moid"] = item.Moid
		storagecontrollers["object_type"] = item.ObjectType
		storagecontrollers["selector"] = item.Selector
		storagecontrollerss = append(storagecontrollerss, storagecontrollers)
	}
	return storagecontrollerss
}
func flattenListStorageFlexFlashControllerRef(p []*models.StorageFlexFlashControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexflashcontrollers := make(map[string]interface{})
		storageflexflashcontrollers["moid"] = item.Moid
		storageflexflashcontrollers["object_type"] = item.ObjectType
		storageflexflashcontrollers["selector"] = item.Selector
		storageflexflashcontrollerss = append(storageflexflashcontrollerss, storageflexflashcontrollers)
	}
	return storageflexflashcontrollerss
}
func flattenListStorageFlexUtilControllerRef(p []*models.StorageFlexUtilControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageflexutilcontrollers := make(map[string]interface{})
		storageflexutilcontrollers["moid"] = item.Moid
		storageflexutilcontrollers["object_type"] = item.ObjectType
		storageflexutilcontrollers["selector"] = item.Selector
		storageflexutilcontrollerss = append(storageflexutilcontrollerss, storageflexutilcontrollers)
	}
	return storageflexutilcontrollerss
}
func flattenMapInventoryGenericInventoryHolderRef(p *models.InventoryGenericInventoryHolderRef, d *schema.ResourceData) []map[string]interface{} {

	var inventorygenericinventoryholders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	inventorygenericinventoryholder := make(map[string]interface{})
	inventorygenericinventoryholder["moid"] = item.Moid
	inventorygenericinventoryholder["object_type"] = item.ObjectType
	inventorygenericinventoryholder["selector"] = item.Selector

	inventorygenericinventoryholders = append(inventorygenericinventoryholders, inventorygenericinventoryholder)
	return inventorygenericinventoryholders
}
func flattenListStorageReplicationBlackout(p []*models.StorageReplicationBlackout, d *schema.ResourceData) []map[string]interface{} {
	var replicationblackoutintervalss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		replicationblackoutintervals := make(map[string]interface{})
		replicationblackoutintervals["object_type"] = item.ObjectType
		replicationblackoutintervalss = append(replicationblackoutintervalss, replicationblackoutintervals)
	}
	return replicationblackoutintervalss
}
func flattenListEquipmentRackEnclosureSlotRef(p []*models.EquipmentRackEnclosureSlotRef, d *schema.ResourceData) []map[string]interface{} {
	var slotss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		slots := make(map[string]interface{})
		slots["moid"] = item.Moid
		slots["object_type"] = item.ObjectType
		slots["selector"] = item.Selector
		slotss = append(slotss, slots)
	}
	return slotss
}
func flattenMapPortSubGroupRef(p *models.PortSubGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var portsubgroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	portsubgroup := make(map[string]interface{})
	portsubgroup["moid"] = item.Moid
	portsubgroup["object_type"] = item.ObjectType
	portsubgroup["selector"] = item.Selector

	portsubgroups = append(portsubgroups, portsubgroup)
	return portsubgroups
}
func flattenListStorageFlexUtilPhysicalDriveRef(p []*models.StorageFlexUtilPhysicalDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var flexutilphysicaldrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		flexutilphysicaldrives := make(map[string]interface{})
		flexutilphysicaldrives["moid"] = item.Moid
		flexutilphysicaldrives["object_type"] = item.ObjectType
		flexutilphysicaldrives["selector"] = item.Selector
		flexutilphysicaldrivess = append(flexutilphysicaldrivess, flexutilphysicaldrives)
	}
	return flexutilphysicaldrivess
}
func flattenListStorageFlexUtilVirtualDriveRef(p []*models.StorageFlexUtilVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var flexutilvirtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		flexutilvirtualdrives := make(map[string]interface{})
		flexutilvirtualdrives["moid"] = item.Moid
		flexutilvirtualdrives["object_type"] = item.ObjectType
		flexutilvirtualdrives["selector"] = item.Selector
		flexutilvirtualdrivess = append(flexutilvirtualdrivess, flexutilvirtualdrives)
	}
	return flexutilvirtualdrivess
}
func flattenListComputeBladeRef(p []*models.ComputeBladeRef, d *schema.ResourceData) []map[string]interface{} {
	var computebladess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		computeblades := make(map[string]interface{})
		computeblades["moid"] = item.Moid
		computeblades["object_type"] = item.ObjectType
		computeblades["selector"] = item.Selector
		computebladess = append(computebladess, computeblades)
	}
	return computebladess
}
func flattenListComputeRackUnitRef(p []*models.ComputeRackUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		computerackunits := make(map[string]interface{})
		computerackunits["moid"] = item.Moid
		computerackunits["object_type"] = item.ObjectType
		computerackunits["selector"] = item.Selector
		computerackunitss = append(computerackunitss, computerackunits)
	}
	return computerackunitss
}
func flattenMapForecastCatalogRef(p *models.ForecastCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var catalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	catalog := make(map[string]interface{})
	catalog["moid"] = item.Moid
	catalog["object_type"] = item.ObjectType
	catalog["selector"] = item.Selector

	catalogs = append(catalogs, catalog)
	return catalogs
}
func flattenListCondHclStatusDetailRef(p []*models.CondHclStatusDetailRef, d *schema.ResourceData) []map[string]interface{} {
	var detailss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		details := make(map[string]interface{})
		details["moid"] = item.Moid
		details["object_type"] = item.ObjectType
		details["selector"] = item.Selector
		detailss = append(detailss, details)
	}
	return detailss
}
func flattenMapInventoryBaseRef(p *models.InventoryBaseRef, d *schema.ResourceData) []map[string]interface{} {

	var managedobjects []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	managedobject := make(map[string]interface{})
	managedobject["moid"] = item.Moid
	managedobject["object_type"] = item.ObjectType
	managedobject["selector"] = item.Selector

	managedobjects = append(managedobjects, managedobject)
	return managedobjects
}
func flattenMapNiatelemetryDiskinfo(p *models.NiatelemetryDiskinfo, d *schema.ResourceData) []map[string]interface{} {

	var disks []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	disk := make(map[string]interface{})
	disk["free"] = item.Free
	disk["name"] = item.Name
	disk["object_type"] = item.ObjectType
	disk["total"] = item.Total
	disk["used"] = item.Used

	disks = append(disks, disk)
	return disks
}
func flattenMapNiatelemetryNiaLicenseStateRef(p *models.NiatelemetryNiaLicenseStateRef, d *schema.ResourceData) []map[string]interface{} {

	var licensestates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	licensestate := make(map[string]interface{})
	licensestate["moid"] = item.Moid
	licensestate["object_type"] = item.ObjectType
	licensestate["selector"] = item.Selector

	licensestates = append(licensestates, licensestate)
	return licensestates
}
func flattenListStoragePureHostRef(p []*models.StoragePureHostRef, d *schema.ResourceData) []map[string]interface{} {
	var hostgroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		hostgroups := make(map[string]interface{})
		hostgroups["moid"] = item.Moid
		hostgroups["object_type"] = item.ObjectType
		hostgroups["selector"] = item.Selector
		hostgroupss = append(hostgroupss, hostgroups)
	}
	return hostgroupss
}
func flattenListStoragePureVolumeRef(p []*models.StoragePureVolumeRef, d *schema.ResourceData) []map[string]interface{} {
	var volumess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		volumes := make(map[string]interface{})
		volumes["moid"] = item.Moid
		volumes["object_type"] = item.ObjectType
		volumes["selector"] = item.Selector
		volumess = append(volumess, volumes)
	}
	return volumess
}
func flattenMapEquipmentSwitchCardRef(p *models.EquipmentSwitchCardRef, d *schema.ResourceData) []map[string]interface{} {

	var equipmentswitchcards []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	equipmentswitchcard := make(map[string]interface{})
	equipmentswitchcard["moid"] = item.Moid
	equipmentswitchcard["object_type"] = item.ObjectType
	equipmentswitchcard["selector"] = item.Selector

	equipmentswitchcards = append(equipmentswitchcards, equipmentswitchcard)
	return equipmentswitchcards
}
func flattenListFcPhysicalPortRef(p []*models.FcPhysicalPortRef, d *schema.ResourceData) []map[string]interface{} {
	var fcportss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		fcports := make(map[string]interface{})
		fcports["moid"] = item.Moid
		fcports["object_type"] = item.ObjectType
		fcports["selector"] = item.Selector
		fcportss = append(fcportss, fcports)
	}
	return fcportss
}
func flattenListPortSubGroupRef(p []*models.PortSubGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var subgroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		subgroups := make(map[string]interface{})
		subgroups["moid"] = item.Moid
		subgroups["object_type"] = item.ObjectType
		subgroups["selector"] = item.Selector
		subgroupss = append(subgroupss, subgroups)
	}
	return subgroupss
}
func flattenListIamPermissionToRoles(p []*models.IamPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var permissionroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		permissionroles := make(map[string]interface{})
		permissionroles["object_type"] = item.ObjectType
		permissionroles["permission"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {

			var permissions []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			permission := make(map[string]interface{})
			permission["moid"] = item.Moid
			permission["object_type"] = item.ObjectType

			permissions = append(permissions, permission)
			return permissions
		})(item.Permission, d)
		permissionroles["roles"] = (func(p []*models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var roless []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				roles := make(map[string]interface{})
				roles["moid"] = item.Moid
				roles["object_type"] = item.ObjectType
				roless = append(roless, roles)
			}
			return roless
		})(item.Roles, d)
		permissionroless = append(permissionroless, permissionroles)
	}
	return permissionroless
}
func flattenMapPolicyConfigResultContext(p *models.PolicyConfigResultContext, d *schema.ResourceData) []map[string]interface{} {

	var configchangecontexts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	configchangecontext := make(map[string]interface{})
	configchangecontext["entity_data"] = item.EntityData
	configchangecontext["entity_moid"] = item.EntityMoid
	configchangecontext["entity_name"] = item.EntityName
	configchangecontext["entity_type"] = item.EntityType
	configchangecontext["object_type"] = item.ObjectType

	configchangecontexts = append(configchangecontexts, configchangecontext)
	return configchangecontexts
}
func flattenMapAssetClusterMemberRef(p *models.AssetClusterMemberRef, d *schema.ResourceData) []map[string]interface{} {

	var clustermembers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	clustermember := make(map[string]interface{})
	clustermember["moid"] = item.Moid
	clustermember["object_type"] = item.ObjectType
	clustermember["selector"] = item.Selector

	clustermembers = append(clustermembers, clustermember)
	return clustermembers
}
func flattenMapHyperflexHxNetworkAddressDt(p *models.HyperflexHxNetworkAddressDt, d *schema.ResourceData) []map[string]interface{} {

	var ips []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	ip := make(map[string]interface{})
	ip["address"] = item.Address
	ip["fqdn"] = item.Fqdn
	ip["ip"] = item.IP
	ip["object_type"] = item.ObjectType

	ips = append(ips, ip)
	return ips
}
func flattenMapHyperflexHxUuIDDt(p *models.HyperflexHxUuIDDt, d *schema.ResourceData) []map[string]interface{} {

	var identitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	identity := make(map[string]interface{})
	identity["links"] = (func(p []*models.HyperflexHxLinkDt, d *schema.ResourceData) []map[string]interface{} {
		var linkss []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			links := make(map[string]interface{})
			links["comments"] = item.Comments
			links["href"] = item.Href
			links["method"] = item.Method
			links["object_type"] = item.ObjectType
			links["rel"] = item.Rel
			linkss = append(linkss, links)
		}
		return linkss
	})(item.Links, d)
	identity["object_type"] = item.ObjectType
	identity["uuid"] = item.UUID

	identitys = append(identitys, identity)
	return identitys
}
func flattenMapMemoryArrayRef(p *models.MemoryArrayRef, d *schema.ResourceData) []map[string]interface{} {

	var memoryarrays []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	memoryarray := make(map[string]interface{})
	memoryarray["moid"] = item.Moid
	memoryarray["object_type"] = item.ObjectType
	memoryarray["selector"] = item.Selector

	memoryarrays = append(memoryarrays, memoryarray)
	return memoryarrays
}
func flattenMapNiaapiNewReleaseDetail(p *models.NiaapiNewReleaseDetail, d *schema.ResourceData) []map[string]interface{} {

	var postdetails []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	postdetail := make(map[string]interface{})
	postdetail["description"] = item.Description
	postdetail["link"] = item.Link
	postdetail["object_type"] = item.ObjectType
	postdetail["release_note_link"] = item.ReleaseNoteLink
	postdetail["release_note_link_title"] = item.ReleaseNoteLinkTitle
	postdetail["software_download_link"] = item.SoftwareDownloadLink
	postdetail["software_download_link_title"] = item.SoftwareDownloadLinkTitle
	postdetail["title"] = item.Title
	postdetail["version"] = item.Version

	postdetails = append(postdetails, postdetail)
	return postdetails
}
func flattenListServerConfigResultEntryRef(p []*models.ServerConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var resultentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		resultentries := make(map[string]interface{})
		resultentries["moid"] = item.Moid
		resultentries["object_type"] = item.ObjectType
		resultentries["selector"] = item.Selector
		resultentriess = append(resultentriess, resultentries)
	}
	return resultentriess
}
func flattenMapAssetDeviceConnectionRef(p *models.AssetDeviceConnectionRef, d *schema.ResourceData) []map[string]interface{} {

	var deviceregistrations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	deviceregistration := make(map[string]interface{})
	deviceregistration["moid"] = item.Moid
	deviceregistration["object_type"] = item.ObjectType
	deviceregistration["selector"] = item.Selector

	deviceregistrations = append(deviceregistrations, deviceregistration)
	return deviceregistrations
}
func flattenListAdapterUnitRef(p []*models.AdapterUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var adapterss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		adapters := make(map[string]interface{})
		adapters["moid"] = item.Moid
		adapters["object_type"] = item.ObjectType
		adapters["selector"] = item.Selector
		adapterss = append(adapterss, adapters)
	}
	return adapterss
}
func flattenMapBiosBootModeRef(p *models.BiosBootModeRef, d *schema.ResourceData) []map[string]interface{} {

	var biosbootmodes []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	biosbootmode := make(map[string]interface{})
	biosbootmode["moid"] = item.Moid
	biosbootmode["object_type"] = item.ObjectType
	biosbootmode["selector"] = item.Selector

	biosbootmodes = append(biosbootmodes, biosbootmode)
	return biosbootmodes
}
func flattenListBiosUnitRef(p []*models.BiosUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var biosunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		biosunits := make(map[string]interface{})
		biosunits["moid"] = item.Moid
		biosunits["object_type"] = item.ObjectType
		biosunits["selector"] = item.Selector
		biosunitss = append(biosunitss, biosunits)
	}
	return biosunitss
}
func flattenMapBootDeviceBootModeRef(p *models.BootDeviceBootModeRef, d *schema.ResourceData) []map[string]interface{} {

	var bootdevicebootmodes []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	bootdevicebootmode := make(map[string]interface{})
	bootdevicebootmode["moid"] = item.Moid
	bootdevicebootmode["object_type"] = item.ObjectType
	bootdevicebootmode["selector"] = item.Selector

	bootdevicebootmodes = append(bootdevicebootmodes, bootdevicebootmode)
	return bootdevicebootmodes
}
func flattenListInventoryGenericInventoryHolderRef(p []*models.InventoryGenericInventoryHolderRef, d *schema.ResourceData) []map[string]interface{} {
	var genericinventoryholderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		genericinventoryholders := make(map[string]interface{})
		genericinventoryholders["moid"] = item.Moid
		genericinventoryholders["object_type"] = item.ObjectType
		genericinventoryholders["selector"] = item.Selector
		genericinventoryholderss = append(genericinventoryholderss, genericinventoryholders)
	}
	return genericinventoryholderss
}
func flattenListPciDeviceRef(p []*models.PciDeviceRef, d *schema.ResourceData) []map[string]interface{} {
	var pcidevicess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		pcidevices := make(map[string]interface{})
		pcidevices["moid"] = item.Moid
		pcidevices["object_type"] = item.ObjectType
		pcidevices["selector"] = item.Selector
		pcidevicess = append(pcidevicess, pcidevices)
	}
	return pcidevicess
}
func flattenMapEquipmentRackEnclosureSlotRef(p *models.EquipmentRackEnclosureSlotRef, d *schema.ResourceData) []map[string]interface{} {

	var rackenclosureslots []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	rackenclosureslot := make(map[string]interface{})
	rackenclosureslot["moid"] = item.Moid
	rackenclosureslot["object_type"] = item.ObjectType
	rackenclosureslot["selector"] = item.Selector

	rackenclosureslots = append(rackenclosureslots, rackenclosureslot)
	return rackenclosureslots
}
func flattenListStorageSasExpanderRef(p []*models.StorageSasExpanderRef, d *schema.ResourceData) []map[string]interface{} {
	var sasexpanderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		sasexpanders := make(map[string]interface{})
		sasexpanders["moid"] = item.Moid
		sasexpanders["object_type"] = item.ObjectType
		sasexpanders["selector"] = item.Selector
		sasexpanderss = append(sasexpanderss, sasexpanders)
	}
	return sasexpanderss
}
func flattenListStorageEnclosureRef(p []*models.StorageEnclosureRef, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosuress []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		storageenclosures := make(map[string]interface{})
		storageenclosures["moid"] = item.Moid
		storageenclosures["object_type"] = item.ObjectType
		storageenclosures["selector"] = item.Selector
		storageenclosuress = append(storageenclosuress, storageenclosures)
	}
	return storageenclosuress
}
func flattenListStorageEnclosureDiskSlotEpRef(p []*models.StorageEnclosureDiskSlotEpRef, d *schema.ResourceData) []map[string]interface{} {
	var enclosurediskslotss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		enclosurediskslots := make(map[string]interface{})
		enclosurediskslots["moid"] = item.Moid
		enclosurediskslots["object_type"] = item.ObjectType
		enclosurediskslots["selector"] = item.Selector
		enclosurediskslotss = append(enclosurediskslotss, enclosurediskslots)
	}
	return enclosurediskslotss
}
func flattenListStorageEnclosureDiskRef(p []*models.StorageEnclosureDiskRef, d *schema.ResourceData) []map[string]interface{} {
	var enclosurediskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		enclosuredisks := make(map[string]interface{})
		enclosuredisks["moid"] = item.Moid
		enclosuredisks["object_type"] = item.ObjectType
		enclosuredisks["selector"] = item.Selector
		enclosurediskss = append(enclosurediskss, enclosuredisks)
	}
	return enclosurediskss
}
func flattenListIamAccountPermissions(p []*models.IamAccountPermissions, d *schema.ResourceData) []map[string]interface{} {
	var accountpermissionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		accountpermissions := make(map[string]interface{})
		accountpermissions["account_identifier"] = item.AccountIdentifier
		accountpermissions["account_name"] = item.AccountName
		accountpermissions["account_status"] = item.AccountStatus
		accountpermissions["object_type"] = item.ObjectType
		accountpermissions["permissions"] = (func(p []*models.IamPermissionReference, d *schema.ResourceData) []map[string]interface{} {
			var permissionss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				permissions := make(map[string]interface{})
				permissions["object_type"] = item.ObjectType
				permissions["permission_identifier"] = item.PermissionIdentifier
				permissions["permission_name"] = item.PermissionName
				permissionss = append(permissionss, permissions)
			}
			return permissionss
		})(item.Permissions, d)
		accountpermissionss = append(accountpermissionss, accountpermissions)
	}
	return accountpermissionss
}
func flattenListForecastDefinitionRef(p []*models.ForecastDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var definitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		definition := make(map[string]interface{})
		definition["moid"] = item.Moid
		definition["object_type"] = item.ObjectType
		definition["selector"] = item.Selector
		definitions = append(definitions, definition)
	}
	return definitions
}
func flattenMapLicenseCustomerOpRef(p *models.LicenseCustomerOpRef, d *schema.ResourceData) []map[string]interface{} {

	var customerops []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	customerop := make(map[string]interface{})
	customerop["moid"] = item.Moid
	customerop["object_type"] = item.ObjectType
	customerop["selector"] = item.Selector

	customerops = append(customerops, customerop)
	return customerops
}
func flattenListLicenseLicenseInfoRef(p []*models.LicenseLicenseInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var licenseinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		licenseinfos := make(map[string]interface{})
		licenseinfos["moid"] = item.Moid
		licenseinfos["object_type"] = item.ObjectType
		licenseinfos["selector"] = item.Selector
		licenseinfoss = append(licenseinfoss, licenseinfos)
	}
	return licenseinfoss
}
func flattenMapLicenseSmartlicenseTokenRef(p *models.LicenseSmartlicenseTokenRef, d *schema.ResourceData) []map[string]interface{} {

	var smartlicensetokens []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	smartlicensetoken := make(map[string]interface{})
	smartlicensetoken["moid"] = item.Moid
	smartlicensetoken["object_type"] = item.ObjectType
	smartlicensetoken["selector"] = item.Selector

	smartlicensetokens = append(smartlicensetokens, smartlicensetoken)
	return smartlicensetokens
}
func flattenMapAssetSudiInfo(p *models.AssetSudiInfo, d *schema.ResourceData) []map[string]interface{} {

	var sudis []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	sudi := make(map[string]interface{})
	sudi["object_type"] = item.ObjectType
	sudi["pid"] = item.Pid
	sudi["serial_number"] = item.SerialNumber
	sudi["signature"] = item.Signature
	sudi["status"] = item.Status
	sudi["sudi_certificate"] = (func(p *models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {

		var sudicertificates []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		sudicertificate := make(map[string]interface{})
		sudicertificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var issuers []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			issuer := make(map[string]interface{})
			issuer["common_name"] = item.CommonName
			issuer["country"] = item.Country
			issuer["locality"] = item.Locality
			issuer["object_type"] = item.ObjectType
			issuer["organization"] = item.Organization
			issuer["organizational_unit"] = item.OrganizationalUnit
			issuer["state"] = item.State

			issuers = append(issuers, issuer)
			return issuers
		})(item.Issuer, d)
		sudicertificate["object_type"] = item.ObjectType
		sudicertificate["pem_certificate"] = item.PemCertificate
		sudicertificate["sha256_fingerprint"] = item.Sha256Fingerprint
		sudicertificate["signature_algorithm"] = item.SignatureAlgorithm
		sudicertificate["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var subjects []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			subject := make(map[string]interface{})
			subject["common_name"] = item.CommonName
			subject["country"] = item.Country
			subject["locality"] = item.Locality
			subject["object_type"] = item.ObjectType
			subject["organization"] = item.Organization
			subject["organizational_unit"] = item.OrganizationalUnit
			subject["state"] = item.State

			subjects = append(subjects, subject)
			return subjects
		})(item.Subject, d)

		sudicertificates = append(sudicertificates, sudicertificate)
		return sudicertificates
	})(item.SudiCertificate, d)

	sudis = append(sudis, sudi)
	return sudis
}
func flattenListEquipmentIoCardRef(p []*models.EquipmentIoCardRef, d *schema.ResourceData) []map[string]interface{} {
	var iomss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		ioms := make(map[string]interface{})
		ioms["moid"] = item.Moid
		ioms["object_type"] = item.ObjectType
		ioms["selector"] = item.Selector
		iomss = append(iomss, ioms)
	}
	return iomss
}
func flattenListEquipmentSystemIoControllerRef(p []*models.EquipmentSystemIoControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var siocss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		siocs := make(map[string]interface{})
		siocs["moid"] = item.Moid
		siocs["object_type"] = item.ObjectType
		siocs["selector"] = item.Selector
		siocss = append(siocss, siocs)
	}
	return siocss
}
func flattenMapApplianceDataExportPolicyRef(p *models.ApplianceDataExportPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var parentconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	parentconfig := make(map[string]interface{})
	parentconfig["moid"] = item.Moid
	parentconfig["object_type"] = item.ObjectType
	parentconfig["selector"] = item.Selector

	parentconfigs = append(parentconfigs, parentconfig)
	return parentconfigs
}
func flattenListApplianceDataExportPolicyRef(p []*models.ApplianceDataExportPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var subconfigss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		subconfigs := make(map[string]interface{})
		subconfigs["moid"] = item.Moid
		subconfigs["object_type"] = item.ObjectType
		subconfigs["selector"] = item.Selector
		subconfigss = append(subconfigss, subconfigs)
	}
	return subconfigss
}
func flattenMapStorageArrayControllerRef(p *models.StorageArrayControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var controllers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	controller := make(map[string]interface{})
	controller["moid"] = item.Moid
	controller["object_type"] = item.ObjectType
	controller["selector"] = item.Selector

	controllers = append(controllers, controller)
	return controllers
}
func flattenMapCondHclStatusRef(p *models.CondHclStatusRef, d *schema.ResourceData) []map[string]interface{} {

	var hclstatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hclstatus := make(map[string]interface{})
	hclstatus["moid"] = item.Moid
	hclstatus["object_type"] = item.ObjectType
	hclstatus["selector"] = item.Selector

	hclstatuss = append(hclstatuss, hclstatus)
	return hclstatuss
}
func flattenListGraphicsControllerRef(p []*models.GraphicsControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var graphicscontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		graphicscontrollers := make(map[string]interface{})
		graphicscontrollers["moid"] = item.Moid
		graphicscontrollers["object_type"] = item.ObjectType
		graphicscontrollers["selector"] = item.Selector
		graphicscontrollerss = append(graphicscontrollerss, graphicscontrollers)
	}
	return graphicscontrollerss
}
func flattenListNiaapiDetail(p []*models.NiaapiDetail, d *schema.ResourceData) []map[string]interface{} {
	var contents []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		content := make(map[string]interface{})
		content["chksum"] = item.Chksum
		content["filename"] = item.Filename
		content["name"] = item.Name
		content["object_type"] = item.ObjectType
		contents = append(contents, content)
	}
	return contents
}
func flattenMapStorageHostRef(p *models.StorageHostRef, d *schema.ResourceData) []map[string]interface{} {

	var hosts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	host := make(map[string]interface{})
	host["moid"] = item.Moid
	host["object_type"] = item.ObjectType
	host["selector"] = item.Selector

	hosts = append(hosts, host)
	return hosts
}
func flattenMapStorageVolumeRef(p *models.StorageVolumeRef, d *schema.ResourceData) []map[string]interface{} {

	var volumes []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	volume := make(map[string]interface{})
	volume["moid"] = item.Moid
	volume["object_type"] = item.ObjectType
	volume["selector"] = item.Selector

	volumes = append(volumes, volume)
	return volumes
}
func flattenListIamEndPointPrivilegeRef(p []*models.IamEndPointPrivilegeRef, d *schema.ResourceData) []map[string]interface{} {
	var endpointprivilegess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		endpointprivileges := make(map[string]interface{})
		endpointprivileges["moid"] = item.Moid
		endpointprivileges["object_type"] = item.ObjectType
		endpointprivileges["selector"] = item.Selector
		endpointprivilegess = append(endpointprivilegess, endpointprivileges)
	}
	return endpointprivilegess
}
func flattenMapIamServiceProviderRef(p *models.IamServiceProviderRef, d *schema.ResourceData) []map[string]interface{} {

	var serviceproviders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	serviceprovider := make(map[string]interface{})
	serviceprovider["moid"] = item.Moid
	serviceprovider["object_type"] = item.ObjectType
	serviceprovider["selector"] = item.Selector

	serviceproviders = append(serviceproviders, serviceprovider)
	return serviceproviders
}
func flattenListPciLinkRef(p []*models.PciLinkRef, d *schema.ResourceData) []map[string]interface{} {
	var linkss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		links := make(map[string]interface{})
		links["moid"] = item.Moid
		links["object_type"] = item.ObjectType
		links["selector"] = item.Selector
		linkss = append(linkss, links)
	}
	return linkss
}
func flattenListMetaAccessPrivilege(p []*models.MetaAccessPrivilege, d *schema.ResourceData) []map[string]interface{} {
	var accessprivilegess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		accessprivileges := make(map[string]interface{})
		accessprivileges["method"] = item.Method
		accessprivileges["object_type"] = item.ObjectType
		accessprivileges["privilege"] = item.Privilege
		accessprivilegess = append(accessprivilegess, accessprivileges)
	}
	return accessprivilegess
}
func flattenListMetaPropDefinition(p []*models.MetaPropDefinition, d *schema.ResourceData) []map[string]interface{} {
	var propertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		properties := make(map[string]interface{})
		properties["api_access"] = item.APIAccess
		properties["name"] = item.Name
		properties["object_type"] = item.ObjectType
		properties["op_security"] = item.OpSecurity
		properties["search_weight"] = item.SearchWeight
		propertiess = append(propertiess, properties)
	}
	return propertiess
}
func flattenListMetaRelationshipDefinition(p []*models.MetaRelationshipDefinition, d *schema.ResourceData) []map[string]interface{} {
	var relationshipss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		relationships := make(map[string]interface{})
		relationships["api_access"] = item.APIAccess
		relationships["collection"] = item.Collection
		relationships["name"] = item.Name
		relationships["object_type"] = item.ObjectType
		relationships["type"] = item.Type
		relationshipss = append(relationshipss, relationships)
	}
	return relationshipss
}
func flattenMapFirmwareUpgradeRef(p *models.FirmwareUpgradeRef, d *schema.ResourceData) []map[string]interface{} {

	var upgrades []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	upgrade := make(map[string]interface{})
	upgrade["moid"] = item.Moid
	upgrade["object_type"] = item.ObjectType
	upgrade["selector"] = item.Selector

	upgrades = append(upgrades, upgrade)
	return upgrades
}
func flattenListEquipmentIoExpanderRef(p []*models.EquipmentIoExpanderRef, d *schema.ResourceData) []map[string]interface{} {
	var equipmentioexpanderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		equipmentioexpanders := make(map[string]interface{})
		equipmentioexpanders["moid"] = item.Moid
		equipmentioexpanders["object_type"] = item.ObjectType
		equipmentioexpanders["selector"] = item.Selector
		equipmentioexpanderss = append(equipmentioexpanderss, equipmentioexpanders)
	}
	return equipmentioexpanderss
}
func flattenListResourceMembershipRef(p []*models.ResourceMembershipRef, d *schema.ResourceData) []map[string]interface{} {
	var membershipss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		memberships := make(map[string]interface{})
		memberships["moid"] = item.Moid
		memberships["object_type"] = item.ObjectType
		memberships["selector"] = item.Selector
		membershipss = append(membershipss, memberships)
	}
	return membershipss
}
func flattenMapStorageProtectionGroupSnapshotRef(p *models.StorageProtectionGroupSnapshotRef, d *schema.ResourceData) []map[string]interface{} {

	var protectiongroupsnapshots []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	protectiongroupsnapshot := make(map[string]interface{})
	protectiongroupsnapshot["moid"] = item.Moid
	protectiongroupsnapshot["object_type"] = item.ObjectType
	protectiongroupsnapshot["selector"] = item.Selector

	protectiongroupsnapshots = append(protectiongroupsnapshots, protectiongroupsnapshot)
	return protectiongroupsnapshots
}
func flattenListHyperflexAlarmRef(p []*models.HyperflexAlarmRef, d *schema.ResourceData) []map[string]interface{} {
	var alarms []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		alarm := make(map[string]interface{})
		alarm["moid"] = item.Moid
		alarm["object_type"] = item.ObjectType
		alarm["selector"] = item.Selector
		alarms = append(alarms, alarm)
	}
	return alarms
}
func flattenMapHyperflexHealthRef(p *models.HyperflexHealthRef, d *schema.ResourceData) []map[string]interface{} {

	var healths []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	health := make(map[string]interface{})
	health["moid"] = item.Moid
	health["object_type"] = item.ObjectType
	health["selector"] = item.Selector

	healths = append(healths, health)
	return healths
}
func flattenListHyperflexNodeRef(p []*models.HyperflexNodeRef, d *schema.ResourceData) []map[string]interface{} {
	var nodess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		nodes := make(map[string]interface{})
		nodes["moid"] = item.Moid
		nodes["object_type"] = item.ObjectType
		nodes["selector"] = item.Selector
		nodess = append(nodess, nodes)
	}
	return nodess
}
func flattenMapHyperflexSummary(p *models.HyperflexSummary, d *schema.ResourceData) []map[string]interface{} {

	var summarys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	summary := make(map[string]interface{})
	summary["active_nodes"] = item.ActiveNodes
	summary["address"] = item.Address
	summary["boottime"] = item.Boottime
	summary["cluster_access_policy"] = item.ClusterAccessPolicy
	summary["compression_savings"] = item.CompressionSavings
	summary["data_replication_compliance"] = item.DataReplicationCompliance
	summary["data_replication_factor"] = item.DataReplicationFactor
	summary["deduplication_savings"] = item.DeduplicationSavings
	summary["downtime"] = item.Downtime
	summary["free_capacity"] = item.FreeCapacity
	summary["healing_info"] = (func(p *models.HyperflexStPlatformClusterHealingInfo, d *schema.ResourceData) []map[string]interface{} {

		var healinginfos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		healinginfo := make(map[string]interface{})
		healinginfo["estimated_completion_time_in_seconds"] = item.EstimatedCompletionTimeInSeconds
		healinginfo["in_progress"] = item.InProgress
		healinginfo["messages"] = item.Messages
		healinginfo["messages_iterator"] = item.MessagesIterator
		healinginfo["messages_size"] = item.MessagesSize
		healinginfo["object_type"] = item.ObjectType
		healinginfo["percent_complete"] = item.PercentComplete

		healinginfos = append(healinginfos, healinginfo)
		return healinginfos
	})(item.HealingInfo, d)
	summary["name"] = item.Name
	summary["object_type"] = item.ObjectType
	summary["resiliency_details"] = item.ResiliencyDetails
	summary["resiliency_details_size"] = item.ResiliencyDetailsSize
	summary["resiliency_info"] = (func(p *models.HyperflexStPlatformClusterResiliencyInfo, d *schema.ResourceData) []map[string]interface{} {

		var resiliencyinfos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		resiliencyinfo := make(map[string]interface{})
		resiliencyinfo["hdd_failures_tolerable"] = item.HddFailuresTolerable
		resiliencyinfo["messages"] = item.Messages
		resiliencyinfo["messages_iterator"] = item.MessagesIterator
		resiliencyinfo["messages_size"] = item.MessagesSize
		resiliencyinfo["node_failures_tolerable"] = item.NodeFailuresTolerable
		resiliencyinfo["object_type"] = item.ObjectType
		resiliencyinfo["ssd_failures_tolerable"] = item.SsdFailuresTolerable
		resiliencyinfo["state"] = item.State

		resiliencyinfos = append(resiliencyinfos, resiliencyinfo)
		return resiliencyinfos
	})(item.ResiliencyInfo, d)
	summary["space_status"] = item.SpaceStatus
	summary["state"] = item.State
	summary["total_capacity"] = item.TotalCapacity
	summary["total_savings"] = item.TotalSavings
	summary["uptime"] = item.Uptime
	summary["used_capacity"] = item.UsedCapacity

	summarys = append(summarys, summary)
	return summarys
}
func flattenMapAssetContractInformation(p *models.AssetContractInformation, d *schema.ResourceData) []map[string]interface{} {

	var contracts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	contract := make(map[string]interface{})
	contract["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var billtos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		billto := make(map[string]interface{})
		billto["address1"] = item.Address1
		billto["address2"] = item.Address2
		billto["city"] = item.City
		billto["country"] = item.Country
		billto["location"] = item.Location
		billto["name"] = item.Name
		billto["object_type"] = item.ObjectType
		billto["postal_code"] = item.PostalCode
		billto["state"] = item.State

		billtos = append(billtos, billto)
		return billtos
	})(item.BillTo, d)
	contract["bill_to_global_ultimate"] = (func(p *models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {

		var billtoglobalultimates []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		billtoglobalultimate := make(map[string]interface{})
		billtoglobalultimate["id"] = item.ID
		billtoglobalultimate["name"] = item.Name
		billtoglobalultimate["object_type"] = item.ObjectType

		billtoglobalultimates = append(billtoglobalultimates, billtoglobalultimate)
		return billtoglobalultimates
	})(item.BillToGlobalUltimate, d)
	contract["contract_number"] = item.ContractNumber
	contract["line_status"] = item.LineStatus
	contract["object_type"] = item.ObjectType

	contracts = append(contracts, contract)
	return contracts
}
func flattenMapAssetCustomerInformation(p *models.AssetCustomerInformation, d *schema.ResourceData) []map[string]interface{} {

	var endcustomers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	endcustomer := make(map[string]interface{})
	endcustomer["address"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var addresss []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		address := make(map[string]interface{})
		address["address1"] = item.Address1
		address["address2"] = item.Address2
		address["city"] = item.City
		address["country"] = item.Country
		address["location"] = item.Location
		address["name"] = item.Name
		address["object_type"] = item.ObjectType
		address["postal_code"] = item.PostalCode
		address["state"] = item.State

		addresss = append(addresss, address)
		return addresss
	})(item.Address, d)
	endcustomer["id"] = item.ID
	endcustomer["name"] = item.Name
	endcustomer["object_type"] = item.ObjectType

	endcustomers = append(endcustomers, endcustomer)
	return endcustomers
}
func flattenMapAssetGlobalUltimate(p *models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {

	var enduserglobalultimates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	enduserglobalultimate := make(map[string]interface{})
	enduserglobalultimate["id"] = item.ID
	enduserglobalultimate["name"] = item.Name
	enduserglobalultimate["object_type"] = item.ObjectType

	enduserglobalultimates = append(enduserglobalultimates, enduserglobalultimate)
	return enduserglobalultimates
}
func flattenMapAssetProductInformation(p *models.AssetProductInformation, d *schema.ResourceData) []map[string]interface{} {

	var products []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	product := make(map[string]interface{})
	product["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var billtos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		billto := make(map[string]interface{})
		billto["address1"] = item.Address1
		billto["address2"] = item.Address2
		billto["city"] = item.City
		billto["country"] = item.Country
		billto["location"] = item.Location
		billto["name"] = item.Name
		billto["object_type"] = item.ObjectType
		billto["postal_code"] = item.PostalCode
		billto["state"] = item.State

		billtos = append(billtos, billto)
		return billtos
	})(item.BillTo, d)
	product["description"] = item.Description
	product["family"] = item.Family
	product["group"] = item.Group
	product["number"] = item.Number
	product["object_type"] = item.ObjectType
	product["ship_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var shiptos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		shipto := make(map[string]interface{})
		shipto["address1"] = item.Address1
		shipto["address2"] = item.Address2
		shipto["city"] = item.City
		shipto["country"] = item.Country
		shipto["location"] = item.Location
		shipto["name"] = item.Name
		shipto["object_type"] = item.ObjectType
		shipto["postal_code"] = item.PostalCode
		shipto["state"] = item.State

		shiptos = append(shiptos, shipto)
		return shiptos
	})(item.ShipTo, d)
	product["sub_type"] = item.SubType

	products = append(products, product)
	return products
}
func flattenMapStoragePureHostRef(p *models.StoragePureHostRef, d *schema.ResourceData) []map[string]interface{} {

	var hostgroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	hostgroup := make(map[string]interface{})
	hostgroup["moid"] = item.Moid
	hostgroup["object_type"] = item.ObjectType
	hostgroup["selector"] = item.Selector

	hostgroups = append(hostgroups, hostgroup)
	return hostgroups
}
func flattenListStorageInitiator(p []*models.StorageInitiator, d *schema.ResourceData) []map[string]interface{} {
	var initiatorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		initiators := make(map[string]interface{})
		initiators["iqn"] = item.Iqn
		initiators["name"] = item.Name
		initiators["object_type"] = item.ObjectType
		initiators["type"] = item.Type
		initiators["wwn"] = item.Wwn
		initiatorss = append(initiatorss, initiators)
	}
	return initiatorss
}
func flattenMapStorageHostUtilization(p *models.StorageHostUtilization, d *schema.ResourceData) []map[string]interface{} {

	var storageutilizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	storageutilization := make(map[string]interface{})
	storageutilization["available"] = item.Available
	storageutilization["data_reduction"] = item.DataReduction
	storageutilization["free"] = item.Free
	storageutilization["object_type"] = item.ObjectType
	storageutilization["snapshot"] = item.Snapshot
	storageutilization["thin_provisioned"] = item.ThinProvisioned
	storageutilization["total"] = item.Total
	storageutilization["total_reduction"] = item.TotalReduction
	storageutilization["used"] = item.Used
	storageutilization["volume"] = item.Volume

	storageutilizations = append(storageutilizations, storageutilization)
	return storageutilizations
}
func flattenMapResourceGroupRef(p *models.ResourceGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var groups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	group := make(map[string]interface{})
	group["moid"] = item.Moid
	group["object_type"] = item.ObjectType
	group["selector"] = item.Selector

	groups = append(groups, group)
	return groups
}
