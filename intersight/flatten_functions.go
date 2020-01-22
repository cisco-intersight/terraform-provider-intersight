package intersight

import (
	"encoding/json"
	"log"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenMapIamAccountRef(p *models.IamAccountRef, d *schema.ResourceData) []map[string]interface{} {

	var propaccounts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propaccount := make(map[string]interface{})
	propaccount["moid"] = item.Moid
	propaccount["object_type"] = item.ObjectType
	propaccount["selector"] = item.Selector

	propaccounts = append(propaccounts, propaccount)
	return propaccounts
}
func flattenListIamEndPointRoleRef(p []*models.IamEndPointRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var propendpointroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propendpointroles := make(map[string]interface{})
		propendpointroles["moid"] = item.Moid
		propendpointroles["object_type"] = item.ObjectType
		propendpointroles["selector"] = item.Selector
		propendpointroless = append(propendpointroless, propendpointroles)
	}
	return propendpointroless
}
func flattenListMoBaseMoRef(p []*models.MoBaseMoRef, d *schema.ResourceData) []map[string]interface{} {
	var proppermissionresourcess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppermissionresources := make(map[string]interface{})
		proppermissionresources["moid"] = item.Moid
		proppermissionresources["object_type"] = item.ObjectType
		proppermissionresources["selector"] = item.Selector
		proppermissionresourcess = append(proppermissionresourcess, proppermissionresources)
	}
	return proppermissionresourcess
}
func flattenListIamResourceRolesRef(p []*models.IamResourceRolesRef, d *schema.ResourceData) []map[string]interface{} {
	var propresourceroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propresourceroles := make(map[string]interface{})
		propresourceroles["moid"] = item.Moid
		propresourceroles["object_type"] = item.ObjectType
		propresourceroles["selector"] = item.Selector
		propresourceroless = append(propresourceroless, propresourceroles)
	}
	return propresourceroless
}
func flattenListIamRoleRef(p []*models.IamRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var proproless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proproles := make(map[string]interface{})
		proproles["moid"] = item.Moid
		proproles["object_type"] = item.ObjectType
		proproles["selector"] = item.Selector
		proproless = append(proproless, proproles)
	}
	return proproless
}
func flattenListMoTag(p []*models.MoTag, d *schema.ResourceData) []map[string]interface{} {
	var proptagss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proptags := make(map[string]interface{})
		delete(item.MoTagAO1P1.MoTagAO1P1, "ObjectType")
		if len(item.MoTagAO1P1.MoTagAO1P1) != 0 {

			j, err := json.Marshal(item.MoTagAO1P1.MoTagAO1P1)
			if err == nil {
				proptags["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		proptags["key"] = item.Key
		proptags["object_type"] = item.ObjectType
		proptags["value"] = item.Value
		proptagss = append(proptagss, proptags)
	}
	return proptagss
}
func flattenListIamUserGroupRef(p []*models.IamUserGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var propusergroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propusergroups := make(map[string]interface{})
		propusergroups["moid"] = item.Moid
		propusergroups["object_type"] = item.ObjectType
		propusergroups["selector"] = item.Selector
		propusergroupss = append(propusergroupss, propusergroups)
	}
	return propusergroupss
}
func flattenListIamUserRef(p []*models.IamUserRef, d *schema.ResourceData) []map[string]interface{} {
	var propuserss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propusers := make(map[string]interface{})
		propusers["moid"] = item.Moid
		propusers["object_type"] = item.ObjectType
		propusers["selector"] = item.Selector
		propuserss = append(propuserss, propusers)
	}
	return propuserss
}
func flattenMapHyperflexAppCatalogRef(p *models.HyperflexAppCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var propappcatalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propappcatalog := make(map[string]interface{})
	propappcatalog["moid"] = item.Moid
	propappcatalog["object_type"] = item.ObjectType
	propappcatalog["selector"] = item.Selector

	propappcatalogs = append(propappcatalogs, propappcatalog)
	return propappcatalogs
}
func flattenListHclConstraint(p []*models.HclConstraint, d *schema.ResourceData) []map[string]interface{} {
	var propcapabilityconstraintss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcapabilityconstraints := make(map[string]interface{})
		delete(item.HclConstraintAO1P1.HclConstraintAO1P1, "ObjectType")
		if len(item.HclConstraintAO1P1.HclConstraintAO1P1) != 0 {

			j, err := json.Marshal(item.HclConstraintAO1P1.HclConstraintAO1P1)
			if err == nil {
				propcapabilityconstraints["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propcapabilityconstraints["constraint_name"] = item.ConstraintName
		propcapabilityconstraints["constraint_value"] = item.ConstraintValue
		propcapabilityconstraints["object_type"] = item.ObjectType
		propcapabilityconstraintss = append(propcapabilityconstraintss, propcapabilityconstraints)
	}
	return propcapabilityconstraintss
}
func flattenMapPkixKeyGenerationSpec(p *models.PkixKeyGenerationSpec, d *schema.ResourceData) []map[string]interface{} {

	var propkeyspecs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propkeyspec := make(map[string]interface{})
	delete(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1, "ObjectType")
	if len(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1) != 0 {
		j, err := json.Marshal(item.PkixKeyGenerationSpecAO1P1.PkixKeyGenerationSpecAO1P1)
		if err == nil {
			propkeyspec["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propkeyspec["name"] = item.Name
	propkeyspec["object_type"] = item.ObjectType

	propkeyspecs = append(propkeyspecs, propkeyspec)
	return propkeyspecs
}
func flattenMapIamPermissionRef(p *models.IamPermissionRef, d *schema.ResourceData) []map[string]interface{} {

	var proppermissions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proppermission := make(map[string]interface{})
	proppermission["moid"] = item.Moid
	proppermission["object_type"] = item.ObjectType
	proppermission["selector"] = item.Selector

	proppermissions = append(proppermissions, proppermission)
	return proppermissions
}
func flattenMapIamUserRef(p *models.IamUserRef, d *schema.ResourceData) []map[string]interface{} {

	var propusers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propuser := make(map[string]interface{})
	propuser["moid"] = item.Moid
	propuser["object_type"] = item.ObjectType
	propuser["selector"] = item.Selector

	propusers = append(propusers, propuser)
	return propusers
}
func flattenMapIamSystemRef(p *models.IamSystemRef, d *schema.ResourceData) []map[string]interface{} {

	var propsystems []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsystem := make(map[string]interface{})
	propsystem["moid"] = item.Moid
	propsystem["object_type"] = item.ObjectType
	propsystem["selector"] = item.Selector

	propsystems = append(propsystems, propsystem)
	return propsystems
}
func flattenListIamUserLoginTimeRef(p []*models.IamUserLoginTimeRef, d *schema.ResourceData) []map[string]interface{} {
	var propuserlogintimes []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propuserlogintime := make(map[string]interface{})
		propuserlogintime["moid"] = item.Moid
		propuserlogintime["object_type"] = item.ObjectType
		propuserlogintime["selector"] = item.Selector
		propuserlogintimes = append(propuserlogintimes, propuserlogintime)
	}
	return propuserlogintimes
}
func flattenListIamUserPreferenceRef(p []*models.IamUserPreferenceRef, d *schema.ResourceData) []map[string]interface{} {
	var propuserpreferencess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propuserpreferences := make(map[string]interface{})
		propuserpreferences["moid"] = item.Moid
		propuserpreferences["object_type"] = item.ObjectType
		propuserpreferences["selector"] = item.Selector
		propuserpreferencess = append(propuserpreferencess, propuserpreferences)
	}
	return propuserpreferencess
}
func flattenListHyperflexClusterProfileRef(p []*models.HyperflexClusterProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var propclusterprofiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propclusterprofiles := make(map[string]interface{})
		propclusterprofiles["moid"] = item.Moid
		propclusterprofiles["object_type"] = item.ObjectType
		propclusterprofiles["selector"] = item.Selector
		propclusterprofiless = append(propclusterprofiless, propclusterprofiles)
	}
	return propclusterprofiless
}
func flattenMapHyperflexIPAddrRange(p *models.HyperflexIPAddrRange, d *schema.ResourceData) []map[string]interface{} {

	var propdataipranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdataiprange := make(map[string]interface{})
	delete(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1, "ObjectType")
	if len(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexIPAddrRangeAO1P1.HyperflexIPAddrRangeAO1P1)
		if err == nil {
			propdataiprange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propdataiprange["end_addr"] = item.EndAddr
	propdataiprange["gateway"] = item.Gateway
	propdataiprange["netmask"] = item.Netmask
	propdataiprange["object_type"] = item.ObjectType
	propdataiprange["start_addr"] = item.StartAddr

	propdataipranges = append(propdataipranges, propdataiprange)
	return propdataipranges
}
func flattenMapOrganizationOrganizationRef(p *models.OrganizationOrganizationRef, d *schema.ResourceData) []map[string]interface{} {

	var proporganizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proporganization := make(map[string]interface{})
	proporganization["moid"] = item.Moid
	proporganization["object_type"] = item.ObjectType
	proporganization["selector"] = item.Selector

	proporganizations = append(proporganizations, proporganization)
	return proporganizations
}
func flattenMapHyperflexMacAddrPrefixRange(p *models.HyperflexMacAddrPrefixRange, d *schema.ResourceData) []map[string]interface{} {

	var propmacprefixranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmacprefixrange := make(map[string]interface{})
	delete(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1, "ObjectType")
	if len(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexMacAddrPrefixRangeAO1P1.HyperflexMacAddrPrefixRangeAO1P1)
		if err == nil {
			propmacprefixrange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propmacprefixrange["end_addr"] = item.EndAddr
	propmacprefixrange["object_type"] = item.ObjectType
	propmacprefixrange["start_addr"] = item.StartAddr

	propmacprefixranges = append(propmacprefixranges, propmacprefixrange)
	return propmacprefixranges
}
func flattenMapHyperflexNamedVlan(p *models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {

	var propmgmtvlans []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmgmtvlan := make(map[string]interface{})
	delete(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1, "ObjectType")
	if len(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1)
		if err == nil {
			propmgmtvlan["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propmgmtvlan["name"] = item.Name
	propmgmtvlan["object_type"] = item.ObjectType
	propmgmtvlan["vlan_id"] = item.VlanID

	propmgmtvlans = append(propmgmtvlans, propmgmtvlan)
	return propmgmtvlans
}
func flattenListHyperflexNamedVlan(p []*models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var propvmnetworkvlanss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propvmnetworkvlans := make(map[string]interface{})
		delete(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1, "ObjectType")
		if len(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexNamedVlanAO1P1.HyperflexNamedVlanAO1P1)
			if err == nil {
				propvmnetworkvlans["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propvmnetworkvlans["name"] = item.Name
		propvmnetworkvlans["object_type"] = item.ObjectType
		propvmnetworkvlans["vlan_id"] = item.VlanID
		propvmnetworkvlanss = append(propvmnetworkvlanss, propvmnetworkvlans)
	}
	return propvmnetworkvlanss
}
func flattenMapHyperflexClusterRef(p *models.HyperflexClusterRef, d *schema.ResourceData) []map[string]interface{} {

	var propassociatedclusters []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propassociatedcluster := make(map[string]interface{})
	propassociatedcluster["moid"] = item.Moid
	propassociatedcluster["object_type"] = item.ObjectType
	propassociatedcluster["selector"] = item.Selector

	propassociatedclusters = append(propassociatedclusters, propassociatedcluster)
	return propassociatedclusters
}
func flattenMapHyperflexAutoSupportPolicyRef(p *models.HyperflexAutoSupportPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propautosupports []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propautosupport := make(map[string]interface{})
	propautosupport["moid"] = item.Moid
	propautosupport["object_type"] = item.ObjectType
	propautosupport["selector"] = item.Selector

	propautosupports = append(propautosupports, propautosupport)
	return propautosupports
}
func flattenMapHyperflexClusterNetworkPolicyRef(p *models.HyperflexClusterNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propclusternetworks []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propclusternetwork := make(map[string]interface{})
	propclusternetwork["moid"] = item.Moid
	propclusternetwork["object_type"] = item.ObjectType
	propclusternetwork["selector"] = item.Selector

	propclusternetworks = append(propclusternetworks, propclusternetwork)
	return propclusternetworks
}
func flattenMapHyperflexClusterStoragePolicyRef(p *models.HyperflexClusterStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propclusterstorages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propclusterstorage := make(map[string]interface{})
	propclusterstorage["moid"] = item.Moid
	propclusterstorage["object_type"] = item.ObjectType
	propclusterstorage["selector"] = item.Selector

	propclusterstorages = append(propclusterstorages, propclusterstorage)
	return propclusterstorages
}
func flattenMapPolicyConfigContext(p *models.PolicyConfigContext, d *schema.ResourceData) []map[string]interface{} {

	var propconfigcontexts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propconfigcontext := make(map[string]interface{})
	delete(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1, "ObjectType")
	if len(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1) != 0 {
		j, err := json.Marshal(item.PolicyConfigContextAO1P1.PolicyConfigContextAO1P1)
		if err == nil {
			propconfigcontext["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propconfigcontext["config_state"] = item.ConfigState
	propconfigcontext["control_action"] = item.ControlAction
	propconfigcontext["error_state"] = item.ErrorState
	propconfigcontext["object_type"] = item.ObjectType
	propconfigcontext["oper_state"] = item.OperState

	propconfigcontexts = append(propconfigcontexts, propconfigcontext)
	return propconfigcontexts
}
func flattenMapHyperflexConfigResultRef(p *models.HyperflexConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var propconfigresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propconfigresult := make(map[string]interface{})
	propconfigresult["moid"] = item.Moid
	propconfigresult["object_type"] = item.ObjectType
	propconfigresult["selector"] = item.Selector

	propconfigresults = append(propconfigresults, propconfigresult)
	return propconfigresults
}
func flattenMapHyperflexExtFcStoragePolicyRef(p *models.HyperflexExtFcStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propextfcstorages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propextfcstorage := make(map[string]interface{})
	propextfcstorage["moid"] = item.Moid
	propextfcstorage["object_type"] = item.ObjectType
	propextfcstorage["selector"] = item.Selector

	propextfcstorages = append(propextfcstorages, propextfcstorage)
	return propextfcstorages
}
func flattenMapHyperflexExtIscsiStoragePolicyRef(p *models.HyperflexExtIscsiStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propextiscsistorages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propextiscsistorage := make(map[string]interface{})
	propextiscsistorage["moid"] = item.Moid
	propextiscsistorage["object_type"] = item.ObjectType
	propextiscsistorage["selector"] = item.Selector

	propextiscsistorages = append(propextiscsistorages, propextiscsistorage)
	return propextiscsistorages
}
func flattenMapHyperflexLocalCredentialPolicyRef(p *models.HyperflexLocalCredentialPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var proplocalcredentials []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proplocalcredential := make(map[string]interface{})
	proplocalcredential["moid"] = item.Moid
	proplocalcredential["object_type"] = item.ObjectType
	proplocalcredential["selector"] = item.Selector

	proplocalcredentials = append(proplocalcredentials, proplocalcredential)
	return proplocalcredentials
}
func flattenMapHyperflexNodeConfigPolicyRef(p *models.HyperflexNodeConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propnodeconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnodeconfig := make(map[string]interface{})
	propnodeconfig["moid"] = item.Moid
	propnodeconfig["object_type"] = item.ObjectType
	propnodeconfig["selector"] = item.Selector

	propnodeconfigs = append(propnodeconfigs, propnodeconfig)
	return propnodeconfigs
}
func flattenListHyperflexNodeProfileRef(p []*models.HyperflexNodeProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var propnodeprofileconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propnodeprofileconfig := make(map[string]interface{})
		propnodeprofileconfig["moid"] = item.Moid
		propnodeprofileconfig["object_type"] = item.ObjectType
		propnodeprofileconfig["selector"] = item.Selector
		propnodeprofileconfigs = append(propnodeprofileconfigs, propnodeprofileconfig)
	}
	return propnodeprofileconfigs
}
func flattenMapHyperflexProxySettingPolicyRef(p *models.HyperflexProxySettingPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propproxysettings []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propproxysetting := make(map[string]interface{})
	propproxysetting["moid"] = item.Moid
	propproxysetting["object_type"] = item.ObjectType
	propproxysetting["selector"] = item.Selector

	propproxysettings = append(propproxysettings, propproxysetting)
	return propproxysettings
}
func flattenListWorkflowWorkflowInfoRef(p []*models.WorkflowWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var proprunningworkflowss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proprunningworkflows := make(map[string]interface{})
		proprunningworkflows["moid"] = item.Moid
		proprunningworkflows["object_type"] = item.ObjectType
		proprunningworkflows["selector"] = item.Selector
		proprunningworkflowss = append(proprunningworkflowss, proprunningworkflows)
	}
	return proprunningworkflowss
}
func flattenMapHyperflexSoftwareVersionPolicyRef(p *models.HyperflexSoftwareVersionPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propsoftwareversions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsoftwareversion := make(map[string]interface{})
	propsoftwareversion["moid"] = item.Moid
	propsoftwareversion["object_type"] = item.ObjectType
	propsoftwareversion["selector"] = item.Selector

	propsoftwareversions = append(propsoftwareversions, propsoftwareversion)
	return propsoftwareversions
}
func flattenMapPolicyAbstractProfileRef(p *models.PolicyAbstractProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var propsrctemplates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsrctemplate := make(map[string]interface{})
	propsrctemplate["moid"] = item.Moid
	propsrctemplate["object_type"] = item.ObjectType
	propsrctemplate["selector"] = item.Selector

	propsrctemplates = append(propsrctemplates, propsrctemplate)
	return propsrctemplates
}
func flattenMapHyperflexSysConfigPolicyRef(p *models.HyperflexSysConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propsysconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsysconfig := make(map[string]interface{})
	propsysconfig["moid"] = item.Moid
	propsysconfig["object_type"] = item.ObjectType
	propsysconfig["selector"] = item.Selector

	propsysconfigs = append(propsysconfigs, propsysconfig)
	return propsysconfigs
}
func flattenMapHyperflexUcsmConfigPolicyRef(p *models.HyperflexUcsmConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propucsmconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propucsmconfig := make(map[string]interface{})
	propucsmconfig["moid"] = item.Moid
	propucsmconfig["object_type"] = item.ObjectType
	propucsmconfig["selector"] = item.Selector

	propucsmconfigs = append(propucsmconfigs, propucsmconfig)
	return propucsmconfigs
}
func flattenMapHyperflexVcenterConfigPolicyRef(p *models.HyperflexVcenterConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propvcenterconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvcenterconfig := make(map[string]interface{})
	propvcenterconfig["moid"] = item.Moid
	propvcenterconfig["object_type"] = item.ObjectType
	propvcenterconfig["selector"] = item.Selector

	propvcenterconfigs = append(propvcenterconfigs, propvcenterconfig)
	return propvcenterconfigs
}
func flattenMapCommCredential(p *models.CommCredential, d *schema.ResourceData) []map[string]interface{} {

	var propcredentials []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcredential := make(map[string]interface{})
	delete(item.CommCredentialAO1P1.CommCredentialAO1P1, "ObjectType")
	if len(item.CommCredentialAO1P1.CommCredentialAO1P1) != 0 {
		j, err := json.Marshal(item.CommCredentialAO1P1.CommCredentialAO1P1)
		if err == nil {
			propcredential["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propcredential["is_password_set"] = item.IsPasswordSet
	propcredential["object_type"] = item.ObjectType
	propcredential["password"] = item.Password
	propcredential["username"] = item.Username

	propcredentials = append(propcredentials, propcredential)
	return propcredentials
}
func flattenMapAssetDeviceRegistrationRef(p *models.AssetDeviceRegistrationRef, d *schema.ResourceData) []map[string]interface{} {

	var propdeviceconnectormanagers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdeviceconnectormanager := make(map[string]interface{})
	propdeviceconnectormanager["moid"] = item.Moid
	propdeviceconnectormanager["object_type"] = item.ObjectType
	propdeviceconnectormanager["selector"] = item.Selector

	propdeviceconnectormanagers = append(propdeviceconnectormanagers, propdeviceconnectormanager)
	return propdeviceconnectormanagers
}
func flattenMapAssetManagedDeviceStatus(p *models.AssetManagedDeviceStatus, d *schema.ResourceData) []map[string]interface{} {

	var propstatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstatus := make(map[string]interface{})
	delete(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1, "ObjectType")
	if len(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1) != 0 {
		j, err := json.Marshal(item.AssetManagedDeviceStatusAO1P1.AssetManagedDeviceStatusAO1P1)
		if err == nil {
			propstatus["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propstatus["cloud_port"] = item.CloudPort
	propstatus["connection_failure_reason"] = item.ConnectionFailureReason
	propstatus["connection_status"] = item.ConnectionStatus
	propstatus["error_code"] = item.ErrorCode
	propstatus["error_reason"] = item.ErrorReason
	propstatus["object_type"] = item.ObjectType
	propstatus["process_id"] = item.ProcessID
	propstatus["server_port"] = item.ServerPort
	propstatus["state"] = item.State

	propstatuss = append(propstatuss, propstatus)
	return propstatuss
}
func flattenMapWorkflowWorkflowInfoRef(p *models.WorkflowWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var propworkflowinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propworkflowinfo := make(map[string]interface{})
	propworkflowinfo["moid"] = item.Moid
	propworkflowinfo["object_type"] = item.ObjectType
	propworkflowinfo["selector"] = item.Selector

	propworkflowinfos = append(propworkflowinfos, propworkflowinfo)
	return propworkflowinfos
}
func flattenListPolicyAbstractConfigProfileRef(p []*models.PolicyAbstractConfigProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var propprofiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propprofiles := make(map[string]interface{})
		propprofiles["moid"] = item.Moid
		propprofiles["object_type"] = item.ObjectType
		propprofiles["selector"] = item.Selector
		propprofiless = append(propprofiless, propprofiles)
	}
	return propprofiless
}
func flattenMapIamLdapPolicyRef(p *models.IamLdapPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propldappolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propldappolicy := make(map[string]interface{})
	propldappolicy["moid"] = item.Moid
	propldappolicy["object_type"] = item.ObjectType
	propldappolicy["selector"] = item.Selector

	propldappolicys = append(propldappolicys, propldappolicy)
	return propldappolicys
}
func flattenListTamAPIDataSource(p []*models.TamAPIDataSource, d *schema.ResourceData) []map[string]interface{} {
	var propapidatasourcess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propapidatasources := make(map[string]interface{})
		delete(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1, "ObjectType")
		if len(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1) != 0 {

			j, err := json.Marshal(item.TamBaseDataSourceAO1P1.TamBaseDataSourceAO1P1)
			if err == nil {
				propapidatasources["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propapidatasources["mo_type"] = item.MoType
		propapidatasources["name"] = item.Name
		propapidatasources["object_type"] = item.ObjectType
		propapidatasources["queries"] = (func(p []*models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var propqueriess []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				propqueries := make(map[string]interface{})
				delete(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1, "ObjectType")
				if len(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1) != 0 {

					j, err := json.Marshal(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1)
					if err == nil {
						propqueries["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				propqueries["name"] = item.Name
				propqueries["object_type"] = item.ObjectType
				propqueries["priority"] = item.Priority
				propqueries["query"] = item.Query
				propqueriess = append(propqueriess, propqueries)
			}
			return propqueriess
		})(item.Queries, d)
		propapidatasources["type"] = item.Type
		propapidatasourcess = append(propapidatasourcess, propapidatasources)
	}
	return propapidatasourcess
}
func flattenListTamAction(p []*models.TamAction, d *schema.ResourceData) []map[string]interface{} {
	var propactionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propactions := make(map[string]interface{})
		delete(item.TamActionAO1P1.TamActionAO1P1, "ObjectType")
		if len(item.TamActionAO1P1.TamActionAO1P1) != 0 {

			j, err := json.Marshal(item.TamActionAO1P1.TamActionAO1P1)
			if err == nil {
				propactions["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propactions["affected_object_type"] = item.AffectedObjectType
		propactions["alert_type"] = item.AlertType
		propactions["identifiers"] = (func(p []*models.TamIdentifiers, d *schema.ResourceData) []map[string]interface{} {
			var propidentifierss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				propidentifiers := make(map[string]interface{})
				delete(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1, "ObjectType")
				if len(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1) != 0 {

					j, err := json.Marshal(item.TamIdentifiersAO1P1.TamIdentifiersAO1P1)
					if err == nil {
						propidentifiers["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				propidentifiers["name"] = item.Name
				propidentifiers["object_type"] = item.ObjectType
				propidentifiers["value"] = item.Value
				propidentifierss = append(propidentifierss, propidentifiers)
			}
			return propidentifierss
		})(item.Identifiers, d)
		propactions["name"] = item.Name
		propactions["object_type"] = item.ObjectType
		propactions["operation_type"] = item.OperationType
		propactions["queries"] = (func(p []*models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var propqueriess []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				propqueries := make(map[string]interface{})
				delete(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1, "ObjectType")
				if len(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1) != 0 {

					j, err := json.Marshal(item.TamQueryEntryAO1P1.TamQueryEntryAO1P1)
					if err == nil {
						propqueries["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				propqueries["name"] = item.Name
				propqueries["object_type"] = item.ObjectType
				propqueries["priority"] = item.Priority
				propqueries["query"] = item.Query
				propqueriess = append(propqueriess, propqueries)
			}
			return propqueriess
		})(item.Queries, d)
		propactions["type"] = item.Type
		propactionss = append(propactionss, propactions)
	}
	return propactionss
}
func flattenMapTamSeverity(p *models.TamSeverity, d *schema.ResourceData) []map[string]interface{} {

	var propseveritys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propseverity := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			propseverity["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propseverity["object_type"] = item.ObjectType

	propseveritys = append(propseveritys, propseverity)
	return propseveritys
}
func flattenListBootDeviceBase(p []*models.BootDeviceBase, d *schema.ResourceData) []map[string]interface{} {
	var propbootdevicess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propbootdevices := make(map[string]interface{})
		delete(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1, "ObjectType")
		if len(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1) != 0 {

			j, err := json.Marshal(item.BootDeviceBaseAO1P1.BootDeviceBaseAO1P1)
			if err == nil {
				propbootdevices["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propbootdevices["enabled"] = item.Enabled
		propbootdevices["name"] = item.Name
		propbootdevices["object_type"] = item.ObjectType
		propbootdevicess = append(propbootdevicess, propbootdevices)
	}
	return propbootdevicess
}
func flattenListIamEndPointUserRoleRef(p []*models.IamEndPointUserRoleRef, d *schema.ResourceData) []map[string]interface{} {
	var propendpointuserroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propendpointuserroles := make(map[string]interface{})
		propendpointuserroles["moid"] = item.Moid
		propendpointuserroles["object_type"] = item.ObjectType
		propendpointuserroles["selector"] = item.Selector
		propendpointuserroless = append(propendpointuserroless, propendpointuserroles)
	}
	return propendpointuserroless
}
func flattenMapIamEndPointPasswordProperties(p *models.IamEndPointPasswordProperties, d *schema.ResourceData) []map[string]interface{} {

	var proppasswordpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proppasswordproperties := make(map[string]interface{})
	delete(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1, "ObjectType")
	if len(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.IamEndPointPasswordPropertiesAO1P1.IamEndPointPasswordPropertiesAO1P1)
		if err == nil {
			proppasswordproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proppasswordproperties["enable_password_expiry"] = item.EnablePasswordExpiry
	proppasswordproperties["enforce_strong_password"] = item.EnforceStrongPassword
	proppasswordproperties["grace_period"] = item.GracePeriod
	proppasswordproperties["notification_period"] = item.NotificationPeriod
	proppasswordproperties["object_type"] = item.ObjectType
	proppasswordproperties["password_expiry_duration"] = item.PasswordExpiryDuration
	proppasswordproperties["password_history"] = item.PasswordHistory

	proppasswordpropertiess = append(proppasswordpropertiess, proppasswordproperties)
	return proppasswordpropertiess
}
func flattenMapIamCertificateRef(p *models.IamCertificateRef, d *schema.ResourceData) []map[string]interface{} {

	var propcertificates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcertificate := make(map[string]interface{})
	propcertificate["moid"] = item.Moid
	propcertificate["object_type"] = item.ObjectType
	propcertificate["selector"] = item.Selector

	propcertificates = append(propcertificates, propcertificate)
	return propcertificates
}
func flattenMapIamPrivateKeySpecRef(p *models.IamPrivateKeySpecRef, d *schema.ResourceData) []map[string]interface{} {

	var propprivatekeyspecs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propprivatekeyspec := make(map[string]interface{})
	propprivatekeyspec["moid"] = item.Moid
	propprivatekeyspec["object_type"] = item.ObjectType
	propprivatekeyspec["selector"] = item.Selector

	propprivatekeyspecs = append(propprivatekeyspecs, propprivatekeyspec)
	return propprivatekeyspecs
}
func flattenMapPkixDistinguishedName(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

	var propsubjects []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsubject := make(map[string]interface{})
	delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
	if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
		j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
		if err == nil {
			propsubject["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propsubject["common_name"] = item.CommonName
	propsubject["country"] = item.Country
	propsubject["locality"] = item.Locality
	propsubject["object_type"] = item.ObjectType
	propsubject["organization"] = item.Organization
	propsubject["organizational_unit"] = item.OrganizationalUnit
	propsubject["state"] = item.State

	propsubjects = append(propsubjects, propsubject)
	return propsubjects
}
func flattenMapPkixSubjectAlternateName(p *models.PkixSubjectAlternateName, d *schema.ResourceData) []map[string]interface{} {

	var propsubjectalternatenames []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsubjectalternatename := make(map[string]interface{})
	delete(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1, "ObjectType")
	if len(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1) != 0 {
		j, err := json.Marshal(item.PkixSubjectAlternateNameAO1P1.PkixSubjectAlternateNameAO1P1)
		if err == nil {
			propsubjectalternatename["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propsubjectalternatename["dns_name"] = item.DNSName
	propsubjectalternatename["email_address"] = item.EmailAddress
	propsubjectalternatename["ip_address"] = item.IPAddress
	propsubjectalternatename["object_type"] = item.ObjectType
	propsubjectalternatename["uri"] = item.URI

	propsubjectalternatenames = append(propsubjectalternatenames, propsubjectalternatename)
	return propsubjectalternatenames
}
func flattenMapHyperflexLogicalAvailabilityZone(p *models.HyperflexLogicalAvailabilityZone, d *schema.ResourceData) []map[string]interface{} {

	var proplogicalavalabilityzoneconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proplogicalavalabilityzoneconfig := make(map[string]interface{})
	delete(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1, "ObjectType")
	if len(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexLogicalAvailabilityZoneAO1P1.HyperflexLogicalAvailabilityZoneAO1P1)
		if err == nil {
			proplogicalavalabilityzoneconfig["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proplogicalavalabilityzoneconfig["auto_config"] = item.AutoConfig
	proplogicalavalabilityzoneconfig["object_type"] = item.ObjectType

	proplogicalavalabilityzoneconfigs = append(proplogicalavalabilityzoneconfigs, proplogicalavalabilityzoneconfig)
	return proplogicalavalabilityzoneconfigs
}
func flattenListSyslogLocalClientBase(p []*models.SyslogLocalClientBase, d *schema.ResourceData) []map[string]interface{} {
	var proplocalclientss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proplocalclients := make(map[string]interface{})
		delete(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1, "ObjectType")
		if len(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1) != 0 {

			j, err := json.Marshal(item.SyslogLocalClientBaseAO1P1.SyslogLocalClientBaseAO1P1)
			if err == nil {
				proplocalclients["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		proplocalclients["min_severity"] = item.MinSeverity
		proplocalclients["object_type"] = item.ObjectType
		proplocalclientss = append(proplocalclientss, proplocalclients)
	}
	return proplocalclientss
}
func flattenListSyslogRemoteClientBase(p []*models.SyslogRemoteClientBase, d *schema.ResourceData) []map[string]interface{} {
	var propremoteclientss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propremoteclients := make(map[string]interface{})
		delete(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1, "ObjectType")
		if len(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1) != 0 {

			j, err := json.Marshal(item.SyslogRemoteClientBaseAO1P1.SyslogRemoteClientBaseAO1P1)
			if err == nil {
				propremoteclients["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propremoteclients["enabled"] = item.Enabled
		propremoteclients["hostname"] = item.Hostname
		propremoteclients["min_severity"] = item.MinSeverity
		propremoteclients["object_type"] = item.ObjectType
		propremoteclients["port"] = item.Port
		propremoteclients["protocol"] = item.Protocol
		propremoteclientss = append(propremoteclientss, propremoteclients)
	}
	return propremoteclientss
}
func flattenListIamAPIKeyRef(p []*models.IamAPIKeyRef, d *schema.ResourceData) []map[string]interface{} {
	var propapikeyss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propapikeys := make(map[string]interface{})
		propapikeys["moid"] = item.Moid
		propapikeys["object_type"] = item.ObjectType
		propapikeys["selector"] = item.Selector
		propapikeyss = append(propapikeyss, propapikeys)
	}
	return propapikeyss
}
func flattenListIamAppRegistrationRef(p []*models.IamAppRegistrationRef, d *schema.ResourceData) []map[string]interface{} {
	var propappregistrationss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propappregistrations := make(map[string]interface{})
		propappregistrations["moid"] = item.Moid
		propappregistrations["object_type"] = item.ObjectType
		propappregistrations["selector"] = item.Selector
		propappregistrationss = append(propappregistrationss, propappregistrations)
	}
	return propappregistrationss
}
func flattenMapIamIdpRef(p *models.IamIdpRef, d *schema.ResourceData) []map[string]interface{} {

	var propidps []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propidp := make(map[string]interface{})
	propidp["moid"] = item.Moid
	propidp["object_type"] = item.ObjectType
	propidp["selector"] = item.Selector

	propidps = append(propidps, propidp)
	return propidps
}
func flattenMapIamIdpReferenceRef(p *models.IamIdpReferenceRef, d *schema.ResourceData) []map[string]interface{} {

	var propidpreferences []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propidpreference := make(map[string]interface{})
	propidpreference["moid"] = item.Moid
	propidpreference["object_type"] = item.ObjectType
	propidpreference["selector"] = item.Selector

	propidpreferences = append(propidpreferences, propidpreference)
	return propidpreferences
}
func flattenMapIamLocalUserPasswordRef(p *models.IamLocalUserPasswordRef, d *schema.ResourceData) []map[string]interface{} {

	var proplocaluserpasswords []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proplocaluserpassword := make(map[string]interface{})
	proplocaluserpassword["moid"] = item.Moid
	proplocaluserpassword["object_type"] = item.ObjectType
	proplocaluserpassword["selector"] = item.Selector

	proplocaluserpasswords = append(proplocaluserpasswords, proplocaluserpassword)
	return proplocaluserpasswords
}
func flattenListIamOAuthTokenRef(p []*models.IamOAuthTokenRef, d *schema.ResourceData) []map[string]interface{} {
	var propoauthtokenss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propoauthtokens := make(map[string]interface{})
		propoauthtokens["moid"] = item.Moid
		propoauthtokens["object_type"] = item.ObjectType
		propoauthtokens["selector"] = item.Selector
		propoauthtokenss = append(propoauthtokenss, propoauthtokens)
	}
	return propoauthtokenss
}
func flattenListIamPermissionRef(p []*models.IamPermissionRef, d *schema.ResourceData) []map[string]interface{} {
	var proppermissionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppermissions := make(map[string]interface{})
		proppermissions["moid"] = item.Moid
		proppermissions["object_type"] = item.ObjectType
		proppermissions["selector"] = item.Selector
		proppermissionss = append(proppermissionss, proppermissions)
	}
	return proppermissionss
}
func flattenListIamSessionRef(p []*models.IamSessionRef, d *schema.ResourceData) []map[string]interface{} {
	var propsessionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsessions := make(map[string]interface{})
		propsessions["moid"] = item.Moid
		propsessions["object_type"] = item.ObjectType
		propsessions["selector"] = item.Selector
		propsessionss = append(propsessionss, propsessions)
	}
	return propsessionss
}
func flattenMapTamAdvisoryRef(p *models.TamAdvisoryRef, d *schema.ResourceData) []map[string]interface{} {

	var propadvisorys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propadvisory := make(map[string]interface{})
	propadvisory["moid"] = item.Moid
	propadvisory["object_type"] = item.ObjectType
	propadvisory["selector"] = item.Selector

	propadvisorys = append(propadvisorys, propadvisory)
	return propadvisorys
}
func flattenMapIamLdapBaseProperties(p *models.IamLdapBaseProperties, d *schema.ResourceData) []map[string]interface{} {

	var propbasepropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propbaseproperties := make(map[string]interface{})
	delete(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1, "ObjectType")
	if len(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.IamLdapBasePropertiesAO1P1.IamLdapBasePropertiesAO1P1)
		if err == nil {
			propbaseproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propbaseproperties["attribute"] = item.Attribute
	propbaseproperties["base_dn"] = item.BaseDn
	propbaseproperties["bind_dn"] = item.BindDn
	propbaseproperties["bind_method"] = item.BindMethod
	propbaseproperties["domain"] = item.Domain
	propbaseproperties["enable_encryption"] = item.EnableEncryption
	propbaseproperties["enable_group_authorization"] = item.EnableGroupAuthorization
	propbaseproperties["filter"] = item.Filter
	propbaseproperties["group_attribute"] = item.GroupAttribute
	propbaseproperties["is_password_set"] = item.IsPasswordSet
	propbaseproperties["nested_group_search_depth"] = item.NestedGroupSearchDepth
	propbaseproperties["object_type"] = item.ObjectType
	password_x := d.Get("base_properties").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	propbaseproperties["password"] = password_y["password"]
	propbaseproperties["timeout"] = item.Timeout

	propbasepropertiess = append(propbasepropertiess, propbaseproperties)
	return propbasepropertiess
}
func flattenMapIamLdapDNSParameters(p *models.IamLdapDNSParameters, d *schema.ResourceData) []map[string]interface{} {

	var propdnsparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdnsparameters := make(map[string]interface{})
	delete(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1, "ObjectType")
	if len(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1) != 0 {
		j, err := json.Marshal(item.IamLdapDNSParametersAO1P1.IamLdapDNSParametersAO1P1)
		if err == nil {
			propdnsparameters["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propdnsparameters["object_type"] = item.ObjectType
	propdnsparameters["search_domain"] = item.SearchDomain
	propdnsparameters["search_forest"] = item.SearchForest
	propdnsparameters["source"] = item.Source

	propdnsparameterss = append(propdnsparameterss, propdnsparameters)
	return propdnsparameterss
}
func flattenListIamLdapGroupRef(p []*models.IamLdapGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var propgroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propgroups := make(map[string]interface{})
		propgroups["moid"] = item.Moid
		propgroups["object_type"] = item.ObjectType
		propgroups["selector"] = item.Selector
		propgroupss = append(propgroupss, propgroups)
	}
	return propgroupss
}
func flattenListIamLdapProviderRef(p []*models.IamLdapProviderRef, d *schema.ResourceData) []map[string]interface{} {
	var propproviderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propproviders := make(map[string]interface{})
		propproviders["moid"] = item.Moid
		propproviders["object_type"] = item.ObjectType
		propproviders["selector"] = item.Selector
		propproviderss = append(propproviderss, propproviders)
	}
	return propproviderss
}
func flattenListWorkflowMessage(p []*models.WorkflowMessage, d *schema.ResourceData) []map[string]interface{} {
	var propmessages []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propmessage := make(map[string]interface{})
		delete(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1, "ObjectType")
		if len(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowMessageAO1P1.WorkflowMessageAO1P1)
			if err == nil {
				propmessage["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propmessage["message"] = item.Message
		propmessage["object_type"] = item.ObjectType
		propmessage["severity"] = item.Severity
		propmessages = append(propmessages, propmessage)
	}
	return propmessages
}
func flattenMapHyperflexClusterProfileRef(p *models.HyperflexClusterProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var propnr0clusterprofiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnr0clusterprofile := make(map[string]interface{})
	propnr0clusterprofile["moid"] = item.Moid
	propnr0clusterprofile["object_type"] = item.ObjectType
	propnr0clusterprofile["selector"] = item.Selector

	propnr0clusterprofiles = append(propnr0clusterprofiles, propnr0clusterprofile)
	return propnr0clusterprofiles
}
func flattenMapServerProfileRef(p *models.ServerProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var propnr1profiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnr1profile := make(map[string]interface{})
	propnr1profile["moid"] = item.Moid
	propnr1profile["object_type"] = item.ObjectType
	propnr1profile["selector"] = item.Selector

	propnr1profiles = append(propnr1profiles, propnr1profile)
	return propnr1profiles
}
func flattenMapWorkflowTaskInfoRef(p *models.WorkflowTaskInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var propparenttaskinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propparenttaskinfo := make(map[string]interface{})
	propparenttaskinfo["moid"] = item.Moid
	propparenttaskinfo["object_type"] = item.ObjectType
	propparenttaskinfo["selector"] = item.Selector

	propparenttaskinfos = append(propparenttaskinfos, propparenttaskinfo)
	return propparenttaskinfos
}
func flattenMapWorkflowPendingDynamicWorkflowInfoRef(p *models.WorkflowPendingDynamicWorkflowInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var proppendingdynamicworkflowinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proppendingdynamicworkflowinfo := make(map[string]interface{})
	proppendingdynamicworkflowinfo["moid"] = item.Moid
	proppendingdynamicworkflowinfo["object_type"] = item.ObjectType
	proppendingdynamicworkflowinfo["selector"] = item.Selector

	proppendingdynamicworkflowinfos = append(proppendingdynamicworkflowinfos, proppendingdynamicworkflowinfo)
	return proppendingdynamicworkflowinfos
}
func flattenListWorkflowTaskInfoRef(p []*models.WorkflowTaskInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var proptaskinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proptaskinfos := make(map[string]interface{})
		proptaskinfos["moid"] = item.Moid
		proptaskinfos["object_type"] = item.ObjectType
		proptaskinfos["selector"] = item.Selector
		proptaskinfoss = append(proptaskinfoss, proptaskinfos)
	}
	return proptaskinfoss
}
func flattenMapWorkflowWorkflowDefinitionRef(p *models.WorkflowWorkflowDefinitionRef, d *schema.ResourceData) []map[string]interface{} {

	var propworkflowdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propworkflowdefinition := make(map[string]interface{})
	propworkflowdefinition["moid"] = item.Moid
	propworkflowdefinition["object_type"] = item.ObjectType
	propworkflowdefinition["selector"] = item.Selector

	propworkflowdefinitions = append(propworkflowdefinitions, propworkflowdefinition)
	return propworkflowdefinitions
}
func flattenListHyperflexFeatureLimitEntry(p []*models.HyperflexFeatureLimitEntry, d *schema.ResourceData) []map[string]interface{} {
	var propfeaturelimitentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propfeaturelimitentries := make(map[string]interface{})
		delete(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1, "ObjectType")
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				propfeaturelimitentries["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propfeaturelimitentries["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {

			var propconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propconstraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					propconstraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propconstraint["hxdp_version"] = item.HxdpVersion
			propconstraint["hypervisor_type"] = item.HypervisorType
			propconstraint["mgmt_platform"] = item.MgmtPlatform
			propconstraint["object_type"] = item.ObjectType
			propconstraint["server_model"] = item.ServerModel

			propconstraints = append(propconstraints, propconstraint)
			return propconstraints
		})(item.Constraint, d)
		propfeaturelimitentries["name"] = item.Name
		propfeaturelimitentries["object_type"] = item.ObjectType
		propfeaturelimitentries["value"] = item.Value
		propfeaturelimitentriess = append(propfeaturelimitentriess, propfeaturelimitentries)
	}
	return propfeaturelimitentriess
}
func flattenMapHyperflexInstallerImageRef(p *models.HyperflexInstallerImageRef, d *schema.ResourceData) []map[string]interface{} {

	var propinstallerimages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propinstallerimage := make(map[string]interface{})
	propinstallerimage["moid"] = item.Moid
	propinstallerimage["object_type"] = item.ObjectType
	propinstallerimage["selector"] = item.Selector

	propinstallerimages = append(propinstallerimages, propinstallerimage)
	return propinstallerimages
}
func flattenListHyperflexServerFirmwareVersionEntry(p []*models.HyperflexServerFirmwareVersionEntry, d *schema.ResourceData) []map[string]interface{} {
	var propserverfirmwareversionentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propserverfirmwareversionentries := make(map[string]interface{})
		delete(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1, "ObjectType")
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				propserverfirmwareversionentries["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propserverfirmwareversionentries["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {

			var propconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propconstraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					propconstraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propconstraint["hxdp_version"] = item.HxdpVersion
			propconstraint["hypervisor_type"] = item.HypervisorType
			propconstraint["mgmt_platform"] = item.MgmtPlatform
			propconstraint["object_type"] = item.ObjectType
			propconstraint["server_model"] = item.ServerModel

			propconstraints = append(propconstraints, propconstraint)
			return propconstraints
		})(item.Constraint, d)
		propserverfirmwareversionentries["label"] = item.Label
		propserverfirmwareversionentries["name"] = item.Name
		propserverfirmwareversionentries["object_type"] = item.ObjectType
		propserverfirmwareversionentries["value"] = item.Value
		propserverfirmwareversionentriess = append(propserverfirmwareversionentriess, propserverfirmwareversionentries)
	}
	return propserverfirmwareversionentriess
}
func flattenMapComputeRackUnitRef(p *models.ComputeRackUnitRef, d *schema.ResourceData) []map[string]interface{} {

	var propassignedservers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propassignedserver := make(map[string]interface{})
	propassignedserver["moid"] = item.Moid
	propassignedserver["object_type"] = item.ObjectType
	propassignedserver["selector"] = item.Selector

	propassignedservers = append(propassignedservers, propassignedserver)
	return propassignedservers
}
func flattenListServerConfigChangeDetailRef(p []*models.ServerConfigChangeDetailRef, d *schema.ResourceData) []map[string]interface{} {
	var propconfigchangedetailss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propconfigchangedetails := make(map[string]interface{})
		propconfigchangedetails["moid"] = item.Moid
		propconfigchangedetails["object_type"] = item.ObjectType
		propconfigchangedetails["selector"] = item.Selector
		propconfigchangedetailss = append(propconfigchangedetailss, propconfigchangedetails)
	}
	return propconfigchangedetailss
}
func flattenMapPolicyConfigChange(p *models.PolicyConfigChange, d *schema.ResourceData) []map[string]interface{} {

	var propconfigchangess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propconfigchanges := make(map[string]interface{})
	delete(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1, "ObjectType")
	if len(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1) != 0 {
		j, err := json.Marshal(item.PolicyConfigChangeAO1P1.PolicyConfigChangeAO1P1)
		if err == nil {
			propconfigchanges["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propconfigchanges["changes"] = item.Changes
	propconfigchanges["disruptions"] = item.Disruptions
	propconfigchanges["object_type"] = item.ObjectType

	propconfigchangess = append(propconfigchangess, propconfigchanges)
	return propconfigchangess
}
func flattenMapServerConfigResultRef(p *models.ServerConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var propconfigresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propconfigresult := make(map[string]interface{})
	propconfigresult["moid"] = item.Moid
	propconfigresult["object_type"] = item.ObjectType
	propconfigresult["selector"] = item.Selector

	propconfigresults = append(propconfigresults, propconfigresult)
	return propconfigresults
}
func flattenListSnmpTrap(p []*models.SnmpTrap, d *schema.ResourceData) []map[string]interface{} {
	var propsnmptrapss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsnmptraps := make(map[string]interface{})
		delete(item.SnmpTrapAO1P1.SnmpTrapAO1P1, "ObjectType")
		if len(item.SnmpTrapAO1P1.SnmpTrapAO1P1) != 0 {

			j, err := json.Marshal(item.SnmpTrapAO1P1.SnmpTrapAO1P1)
			if err == nil {
				propsnmptraps["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propsnmptraps["destination"] = item.Destination
		propsnmptraps["enabled"] = item.Enabled
		propsnmptraps["object_type"] = item.ObjectType
		propsnmptraps["port"] = item.Port
		propsnmptraps["type"] = item.Type
		propsnmptraps["user"] = item.User
		propsnmptraps["version"] = item.Version
		propsnmptrapss = append(propsnmptrapss, propsnmptraps)
	}
	return propsnmptrapss
}
func flattenListSnmpUser(p []*models.SnmpUser, d *schema.ResourceData) []map[string]interface{} {
	var propsnmpuserss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsnmpusers := make(map[string]interface{})
		delete(item.SnmpUserAO1P1.SnmpUserAO1P1, "ObjectType")
		if len(item.SnmpUserAO1P1.SnmpUserAO1P1) != 0 {

			j, err := json.Marshal(item.SnmpUserAO1P1.SnmpUserAO1P1)
			if err == nil {
				propsnmpusers["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		auth_password_x := d.Get("snmp_users").([]interface{})
		auth_password_y := auth_password_x[0].(map[string]interface{})
		propsnmpusers["auth_password"] = auth_password_y["auth_password"]
		propsnmpusers["auth_type"] = item.AuthType
		propsnmpusers["is_auth_password_set"] = item.IsAuthPasswordSet
		propsnmpusers["is_privacy_password_set"] = item.IsPrivacyPasswordSet
		propsnmpusers["name"] = item.Name
		propsnmpusers["object_type"] = item.ObjectType
		privacy_password_x := d.Get("snmp_users").([]interface{})
		privacy_password_y := privacy_password_x[0].(map[string]interface{})
		propsnmpusers["privacy_password"] = privacy_password_y["privacy_password"]
		propsnmpusers["privacy_type"] = item.PrivacyType
		propsnmpusers["security_level"] = item.SecurityLevel
		propsnmpuserss = append(propsnmpuserss, propsnmpusers)
	}
	return propsnmpuserss
}
func flattenMapVnicVlanSettings(p *models.VnicVlanSettings, d *schema.ResourceData) []map[string]interface{} {

	var propvlansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvlansettings := make(map[string]interface{})
	delete(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1, "ObjectType")
	if len(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVlanSettingsAO1P1.VnicVlanSettingsAO1P1)
		if err == nil {
			propvlansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propvlansettings["default_vlan"] = item.DefaultVlan
	propvlansettings["mode"] = item.Mode
	propvlansettings["object_type"] = item.ObjectType

	propvlansettingss = append(propvlansettingss, propvlansettings)
	return propvlansettingss
}
func flattenMapRecoveryConfigResultRef(p *models.RecoveryConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var propconfigresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propconfigresult := make(map[string]interface{})
	propconfigresult["moid"] = item.Moid
	propconfigresult["object_type"] = item.ObjectType
	propconfigresult["selector"] = item.Selector

	propconfigresults = append(propconfigresults, propconfigresult)
	return propconfigresults
}
func flattenMapVnicCdn(p *models.VnicCdn, d *schema.ResourceData) []map[string]interface{} {

	var propcdns []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcdn := make(map[string]interface{})
	delete(item.VnicCdnAO1P1.VnicCdnAO1P1, "ObjectType")
	if len(item.VnicCdnAO1P1.VnicCdnAO1P1) != 0 {
		j, err := json.Marshal(item.VnicCdnAO1P1.VnicCdnAO1P1)
		if err == nil {
			propcdn["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propcdn["object_type"] = item.ObjectType
	propcdn["source"] = item.Source
	propcdn["value"] = item.Value

	propcdns = append(propcdns, propcdn)
	return propcdns
}
func flattenMapVnicEthAdapterPolicyRef(p *models.VnicEthAdapterPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propethadapterpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propethadapterpolicy := make(map[string]interface{})
	propethadapterpolicy["moid"] = item.Moid
	propethadapterpolicy["object_type"] = item.ObjectType
	propethadapterpolicy["selector"] = item.Selector

	propethadapterpolicys = append(propethadapterpolicys, propethadapterpolicy)
	return propethadapterpolicys
}
func flattenMapVnicEthNetworkPolicyRef(p *models.VnicEthNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propethnetworkpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propethnetworkpolicy := make(map[string]interface{})
	propethnetworkpolicy["moid"] = item.Moid
	propethnetworkpolicy["object_type"] = item.ObjectType
	propethnetworkpolicy["selector"] = item.Selector

	propethnetworkpolicys = append(propethnetworkpolicys, propethnetworkpolicy)
	return propethnetworkpolicys
}
func flattenMapVnicEthQosPolicyRef(p *models.VnicEthQosPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propethqospolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propethqospolicy := make(map[string]interface{})
	propethqospolicy["moid"] = item.Moid
	propethqospolicy["object_type"] = item.ObjectType
	propethqospolicy["selector"] = item.Selector

	propethqospolicys = append(propethqospolicys, propethqospolicy)
	return propethqospolicys
}
func flattenMapVnicLanConnectivityPolicyRef(p *models.VnicLanConnectivityPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var proplanconnectivitypolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proplanconnectivitypolicy := make(map[string]interface{})
	proplanconnectivitypolicy["moid"] = item.Moid
	proplanconnectivitypolicy["object_type"] = item.ObjectType
	proplanconnectivitypolicy["selector"] = item.Selector

	proplanconnectivitypolicys = append(proplanconnectivitypolicys, proplanconnectivitypolicy)
	return proplanconnectivitypolicys
}
func flattenMapVnicPlacementSettings(p *models.VnicPlacementSettings, d *schema.ResourceData) []map[string]interface{} {

	var propplacements []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propplacement := make(map[string]interface{})
	delete(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1, "ObjectType")
	if len(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicPlacementSettingsAO1P1.VnicPlacementSettingsAO1P1)
		if err == nil {
			propplacement["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propplacement["id"] = item.ID
	propplacement["object_type"] = item.ObjectType
	propplacement["pci_link"] = item.PciLink
	propplacement["uplink"] = item.Uplink

	propplacements = append(propplacements, propplacement)
	return propplacements
}
func flattenMapVnicUsnicSettings(p *models.VnicUsnicSettings, d *schema.ResourceData) []map[string]interface{} {

	var propusnicsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propusnicsettings := make(map[string]interface{})
	delete(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1, "ObjectType")
	if len(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicUsnicSettingsAO1P1.VnicUsnicSettingsAO1P1)
		if err == nil {
			propusnicsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propusnicsettings["cos"] = item.Cos
	propusnicsettings["count"] = item.Count
	propusnicsettings["object_type"] = item.ObjectType
	propusnicsettings["usnic_adapter_policy"] = item.UsnicAdapterPolicy

	propusnicsettingss = append(propusnicsettingss, propusnicsettings)
	return propusnicsettingss
}
func flattenMapVnicVmqSettings(p *models.VnicVmqSettings, d *schema.ResourceData) []map[string]interface{} {

	var propvmqsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvmqsettings := make(map[string]interface{})
	delete(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1, "ObjectType")
	if len(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVmqSettingsAO1P1.VnicVmqSettingsAO1P1)
		if err == nil {
			propvmqsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propvmqsettings["enabled"] = item.Enabled
	propvmqsettings["object_type"] = item.ObjectType

	propvmqsettingss = append(propvmqsettingss, propvmqsettings)
	return propvmqsettingss
}
func flattenListAssetClusterMemberRef(p []*models.AssetClusterMemberRef, d *schema.ResourceData) []map[string]interface{} {
	var propclustermemberss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propclustermembers := make(map[string]interface{})
		propclustermembers["moid"] = item.Moid
		propclustermembers["object_type"] = item.ObjectType
		propclustermembers["selector"] = item.Selector
		propclustermemberss = append(propclustermemberss, propclustermembers)
	}
	return propclustermemberss
}
func flattenMapAssetDeviceClaimRef(p *models.AssetDeviceClaimRef, d *schema.ResourceData) []map[string]interface{} {

	var propdeviceclaims []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdeviceclaim := make(map[string]interface{})
	propdeviceclaim["moid"] = item.Moid
	propdeviceclaim["object_type"] = item.ObjectType
	propdeviceclaim["selector"] = item.Selector

	propdeviceclaims = append(propdeviceclaims, propdeviceclaim)
	return propdeviceclaims
}
func flattenMapAssetDeviceConfigurationRef(p *models.AssetDeviceConfigurationRef, d *schema.ResourceData) []map[string]interface{} {

	var propdeviceconfigurations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdeviceconfiguration := make(map[string]interface{})
	propdeviceconfiguration["moid"] = item.Moid
	propdeviceconfiguration["object_type"] = item.ObjectType
	propdeviceconfiguration["selector"] = item.Selector

	propdeviceconfigurations = append(propdeviceconfigurations, propdeviceconfiguration)
	return propdeviceconfigurations
}
func flattenMapIamDomainGroupRef(p *models.IamDomainGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var propdomaingroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdomaingroup := make(map[string]interface{})
	propdomaingroup["moid"] = item.Moid
	propdomaingroup["object_type"] = item.ObjectType
	propdomaingroup["selector"] = item.Selector

	propdomaingroups = append(propdomaingroups, propdomaingroup)
	return propdomaingroups
}
func flattenMapAssetParentConnectionSignature(p *models.AssetParentConnectionSignature, d *schema.ResourceData) []map[string]interface{} {

	var propparentsignatures []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propparentsignature := make(map[string]interface{})
	delete(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1, "ObjectType")
	if len(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1) != 0 {
		j, err := json.Marshal(item.AssetParentConnectionSignatureAO1P1.AssetParentConnectionSignatureAO1P1)
		if err == nil {
			propparentsignature["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propparentsignature["device_id"] = item.DeviceID
	propparentsignature["node_id"] = item.NodeID
	propparentsignature["object_type"] = item.ObjectType
	propparentsignature["signature"] = item.Signature

	propparentsignatures = append(propparentsignatures, propparentsignature)
	return propparentsignatures
}
func flattenMapAssetSecurityTokenRef(p *models.AssetSecurityTokenRef, d *schema.ResourceData) []map[string]interface{} {

	var propsecuritytokens []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsecuritytoken := make(map[string]interface{})
	propsecuritytoken["moid"] = item.Moid
	propsecuritytoken["object_type"] = item.ObjectType
	propsecuritytoken["selector"] = item.Selector

	propsecuritytokens = append(propsecuritytokens, propsecuritytoken)
	return propsecuritytokens
}
func flattenListWorkflowAPI(p []*models.WorkflowAPI, d *schema.ResourceData) []map[string]interface{} {
	var propbatchs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propbatch := make(map[string]interface{})
		delete(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1, "ObjectType")
		if len(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowAPIAO1P1.WorkflowAPIAO1P1)
			if err == nil {
				propbatch["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propbatch["body"] = item.Body
		propbatch["content_type"] = item.ContentType
		propbatch["name"] = item.Name
		propbatch["object_type"] = item.ObjectType
		propbatch["outcomes"] = item.Outcomes
		propbatch["response_spec"] = (func(p *models.ContentGrammar, d *schema.ResourceData) []map[string]interface{} {

			var propresponsespecs []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propresponsespec := make(map[string]interface{})
			delete(item.ContentGrammarAO1P1.ContentGrammarAO1P1, "ObjectType")
			if len(item.ContentGrammarAO1P1.ContentGrammarAO1P1) != 0 {
				j, err := json.Marshal(item.ContentGrammarAO1P1.ContentGrammarAO1P1)
				if err == nil {
					propresponsespec["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propresponsespec["error_parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var properrorparameterss []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					properrorparameters := make(map[string]interface{})
					properrorparameters["accept_single_value"] = item.AcceptSingleValue
					properrorparameters["complex_type"] = item.ComplexType
					properrorparameters["item_type"] = item.ItemType
					properrorparameters["name"] = item.Name
					properrorparameters["object_type"] = item.ObjectType
					properrorparameters["path"] = item.Path
					properrorparameters["type"] = item.Type
					properrorparameterss = append(properrorparameterss, properrorparameters)
				}
				return properrorparameterss
			})(item.ErrorParameters, d)
			propresponsespec["object_type"] = item.ObjectType
			propresponsespec["parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var propparameterss []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					propparameters := make(map[string]interface{})
					propparameters["accept_single_value"] = item.AcceptSingleValue
					propparameters["complex_type"] = item.ComplexType
					propparameters["item_type"] = item.ItemType
					propparameters["name"] = item.Name
					propparameters["object_type"] = item.ObjectType
					propparameters["path"] = item.Path
					propparameters["type"] = item.Type
					propparameterss = append(propparameterss, propparameters)
				}
				return propparameterss
			})(item.Parameters, d)
			propresponsespec["types"] = (func(p []*models.ContentComplexType, d *schema.ResourceData) []map[string]interface{} {
				var proptypess []map[string]interface{}
				if p == nil {
					return nil
				}
				for _, item := range p {
					item := *item
					proptypes := make(map[string]interface{})
					proptypes["name"] = item.Name
					proptypes["object_type"] = item.ObjectType
					proptypes["parameters"] = (func(p []*models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
						var propparameterss []map[string]interface{}
						if p == nil {
							return nil
						}
						for _, item := range p {
							item := *item
							propparameters := make(map[string]interface{})
							propparameters["accept_single_value"] = item.AcceptSingleValue
							propparameters["complex_type"] = item.ComplexType
							propparameters["item_type"] = item.ItemType
							propparameters["name"] = item.Name
							propparameters["object_type"] = item.ObjectType
							propparameters["path"] = item.Path
							propparameters["type"] = item.Type
							propparameterss = append(propparameterss, propparameters)
						}
						return propparameterss
					})(item.Parameters, d)
					proptypess = append(proptypess, proptypes)
				}
				return proptypess
			})(item.Types, d)

			propresponsespecs = append(propresponsespecs, propresponsespec)
			return propresponsespecs
		})(item.ResponseSpec, d)
		propbatch["skip_on_condition"] = item.SkipOnCondition
		propbatch["timeout"] = item.Timeout
		propbatchs = append(propbatchs, propbatch)
	}
	return propbatchs
}
func flattenMapWorkflowTaskDefinitionRef(p *models.WorkflowTaskDefinitionRef, d *schema.ResourceData) []map[string]interface{} {

	var proptaskdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proptaskdefinition := make(map[string]interface{})
	proptaskdefinition["moid"] = item.Moid
	proptaskdefinition["object_type"] = item.ObjectType
	proptaskdefinition["selector"] = item.Selector

	proptaskdefinitions = append(proptaskdefinitions, proptaskdefinition)
	return proptaskdefinitions
}
func flattenMapVnicFcErrorRecoverySettings(p *models.VnicFcErrorRecoverySettings, d *schema.ResourceData) []map[string]interface{} {

	var properrorrecoverysettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	properrorrecoverysettings := make(map[string]interface{})
	delete(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1, "ObjectType")
	if len(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcErrorRecoverySettingsAO1P1.VnicFcErrorRecoverySettingsAO1P1)
		if err == nil {
			properrorrecoverysettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	properrorrecoverysettings["enabled"] = item.Enabled
	properrorrecoverysettings["io_retry_count"] = item.IoRetryCount
	properrorrecoverysettings["io_retry_timeout"] = item.IoRetryTimeout
	properrorrecoverysettings["link_down_timeout"] = item.LinkDownTimeout
	properrorrecoverysettings["object_type"] = item.ObjectType
	properrorrecoverysettings["port_down_timeout"] = item.PortDownTimeout

	properrorrecoverysettingss = append(properrorrecoverysettingss, properrorrecoverysettings)
	return properrorrecoverysettingss
}
func flattenMapVnicFlogiSettings(p *models.VnicFlogiSettings, d *schema.ResourceData) []map[string]interface{} {

	var propflogisettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propflogisettings := make(map[string]interface{})
	delete(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1, "ObjectType")
	if len(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFlogiSettingsAO1P1.VnicFlogiSettingsAO1P1)
		if err == nil {
			propflogisettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propflogisettings["object_type"] = item.ObjectType
	propflogisettings["retries"] = item.Retries
	propflogisettings["timeout"] = item.Timeout

	propflogisettingss = append(propflogisettingss, propflogisettings)
	return propflogisettingss
}
func flattenMapVnicFcInterruptSettings(p *models.VnicFcInterruptSettings, d *schema.ResourceData) []map[string]interface{} {

	var propinterruptsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propinterruptsettings := make(map[string]interface{})
	delete(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1, "ObjectType")
	if len(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcInterruptSettingsAO1P1.VnicFcInterruptSettingsAO1P1)
		if err == nil {
			propinterruptsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propinterruptsettings["mode"] = item.Mode
	propinterruptsettings["object_type"] = item.ObjectType

	propinterruptsettingss = append(propinterruptsettingss, propinterruptsettings)
	return propinterruptsettingss
}
func flattenMapVnicPlogiSettings(p *models.VnicPlogiSettings, d *schema.ResourceData) []map[string]interface{} {

	var propplogisettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propplogisettings := make(map[string]interface{})
	delete(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1, "ObjectType")
	if len(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicPlogiSettingsAO1P1.VnicPlogiSettingsAO1P1)
		if err == nil {
			propplogisettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propplogisettings["object_type"] = item.ObjectType
	propplogisettings["retries"] = item.Retries
	propplogisettings["timeout"] = item.Timeout

	propplogisettingss = append(propplogisettingss, propplogisettings)
	return propplogisettingss
}
func flattenMapVnicFcQueueSettings(p *models.VnicFcQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var proprxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proprxqueuesettings := make(map[string]interface{})
	delete(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicFcQueueSettingsAO1P1.VnicFcQueueSettingsAO1P1)
		if err == nil {
			proprxqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proprxqueuesettings["count"] = item.Count
	proprxqueuesettings["object_type"] = item.ObjectType
	proprxqueuesettings["ring_size"] = item.RingSize

	proprxqueuesettingss = append(proprxqueuesettingss, proprxqueuesettings)
	return proprxqueuesettingss
}
func flattenMapVnicScsiQueueSettings(p *models.VnicScsiQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var propscsiqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propscsiqueuesettings := make(map[string]interface{})
	delete(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicScsiQueueSettingsAO1P1.VnicScsiQueueSettingsAO1P1)
		if err == nil {
			propscsiqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propscsiqueuesettings["count"] = item.Count
	propscsiqueuesettings["object_type"] = item.ObjectType
	propscsiqueuesettings["ring_size"] = item.RingSize

	propscsiqueuesettingss = append(propscsiqueuesettingss, propscsiqueuesettings)
	return propscsiqueuesettingss
}
func flattenListStorageLocalDisk(p []*models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
	var propdedicatedhotsparess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propdedicatedhotspares := make(map[string]interface{})
		delete(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1, "ObjectType")
		if len(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1) != 0 {

			j, err := json.Marshal(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1)
			if err == nil {
				propdedicatedhotspares["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propdedicatedhotspares["object_type"] = item.ObjectType
		propdedicatedhotspares["slot_number"] = item.SlotNumber
		propdedicatedhotsparess = append(propdedicatedhotsparess, propdedicatedhotspares)
	}
	return propdedicatedhotsparess
}
func flattenListStorageSpanGroup(p []*models.StorageSpanGroup, d *schema.ResourceData) []map[string]interface{} {
	var propspangroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propspangroups := make(map[string]interface{})
		delete(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1, "ObjectType")
		if len(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1) != 0 {

			j, err := json.Marshal(item.StorageSpanGroupAO1P1.StorageSpanGroupAO1P1)
			if err == nil {
				propspangroups["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propspangroups["disks"] = (func(p []*models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
			var propdiskss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				propdisks := make(map[string]interface{})
				delete(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1, "ObjectType")
				if len(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1) != 0 {

					j, err := json.Marshal(item.StorageLocalDiskAO1P1.StorageLocalDiskAO1P1)
					if err == nil {
						propdisks["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				propdisks["object_type"] = item.ObjectType
				propdisks["slot_number"] = item.SlotNumber
				propdiskss = append(propdiskss, propdisks)
			}
			return propdiskss
		})(item.Disks, d)
		propspangroups["object_type"] = item.ObjectType
		propspangroupss = append(propspangroupss, propspangroups)
	}
	return propspangroupss
}
func flattenListStorageStoragePolicyRef(p []*models.StorageStoragePolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var propstoragepoliciess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propstoragepolicies := make(map[string]interface{})
		propstoragepolicies["moid"] = item.Moid
		propstoragepolicies["object_type"] = item.ObjectType
		propstoragepolicies["selector"] = item.Selector
		propstoragepoliciess = append(propstoragepoliciess, propstoragepolicies)
	}
	return propstoragepoliciess
}
func flattenMapX509Certificate(p *models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {

	var propcertificates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcertificate := make(map[string]interface{})
	delete(item.X509CertificateAO1P1.X509CertificateAO1P1, "ObjectType")
	if len(item.X509CertificateAO1P1.X509CertificateAO1P1) != 0 {
		j, err := json.Marshal(item.X509CertificateAO1P1.X509CertificateAO1P1)
		if err == nil {
			propcertificate["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propcertificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

		var propissuers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propissuer := make(map[string]interface{})
		propissuer["common_name"] = item.CommonName
		propissuer["country"] = item.Country
		propissuer["locality"] = item.Locality
		propissuer["object_type"] = item.ObjectType
		propissuer["organization"] = item.Organization
		propissuer["organizational_unit"] = item.OrganizationalUnit
		propissuer["state"] = item.State

		propissuers = append(propissuers, propissuer)
		return propissuers
	})(item.Issuer, d)
	propcertificate["object_type"] = item.ObjectType
	propcertificate["pem_certificate"] = item.PemCertificate
	propcertificate["sha256_fingerprint"] = item.Sha256Fingerprint
	propcertificate["signature_algorithm"] = item.SignatureAlgorithm
	propcertificate["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

		var propsubjects []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propsubject := make(map[string]interface{})
		propsubject["common_name"] = item.CommonName
		propsubject["country"] = item.Country
		propsubject["locality"] = item.Locality
		propsubject["object_type"] = item.ObjectType
		propsubject["organization"] = item.Organization
		propsubject["organizational_unit"] = item.OrganizationalUnit
		propsubject["state"] = item.State

		propsubjects = append(propsubjects, propsubject)
		return propsubjects
	})(item.Subject, d)

	propcertificates = append(propcertificates, propcertificate)
	return propcertificates
}
func flattenMapIamCertificateRequestRef(p *models.IamCertificateRequestRef, d *schema.ResourceData) []map[string]interface{} {

	var propcertificaterequests []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcertificaterequest := make(map[string]interface{})
	propcertificaterequest["moid"] = item.Moid
	propcertificaterequest["object_type"] = item.ObjectType
	propcertificaterequest["selector"] = item.Selector

	propcertificaterequests = append(propcertificaterequests, propcertificaterequest)
	return propcertificaterequests
}
func flattenMapWorkflowCatalogRef(p *models.WorkflowCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var propcatalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcatalog := make(map[string]interface{})
	propcatalog["moid"] = item.Moid
	propcatalog["object_type"] = item.ObjectType
	propcatalog["selector"] = item.Selector

	propcatalogs = append(propcatalogs, propcatalog)
	return propcatalogs
}
func flattenListWorkflowBaseDataType(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
	var propinputdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propinputdefinition := make(map[string]interface{})
		delete(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1, "ObjectType")
		if len(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowBaseDataTypeAO1P1.WorkflowBaseDataTypeAO1P1)
			if err == nil {
				propinputdefinition["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propinputdefinition["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {

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
		propinputdefinition["description"] = item.Description
		propinputdefinition["label"] = item.Label
		propinputdefinition["name"] = item.Name
		propinputdefinition["object_type"] = item.ObjectType
		propinputdefinition["required"] = item.Required
		propinputdefinitions = append(propinputdefinitions, propinputdefinition)
	}
	return propinputdefinitions
}
func flattenListWorkflowWorkflowTask(p []*models.WorkflowWorkflowTask, d *schema.ResourceData) []map[string]interface{} {
	var proptaskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proptasks := make(map[string]interface{})
		delete(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1, "ObjectType")
		if len(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1) != 0 {

			j, err := json.Marshal(item.WorkflowWorkflowTaskAO1P1.WorkflowWorkflowTaskAO1P1)
			if err == nil {
				proptasks["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		proptasks["description"] = item.Description
		proptasks["label"] = item.Label
		proptasks["name"] = item.Name
		proptasks["object_type"] = item.ObjectType
		proptaskss = append(proptaskss, proptasks)
	}
	return proptaskss
}
func flattenMapWorkflowValidationInformation(p *models.WorkflowValidationInformation, d *schema.ResourceData) []map[string]interface{} {

	var propvalidationinformations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvalidationinformation := make(map[string]interface{})
	delete(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1, "ObjectType")
	if len(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowValidationInformationAO1P1.WorkflowValidationInformationAO1P1)
		if err == nil {
			propvalidationinformation["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propvalidationinformation["object_type"] = item.ObjectType
	propvalidationinformation["state"] = item.State
	propvalidationinformation["validation_error"] = (func(p []*models.WorkflowValidationError, d *schema.ResourceData) []map[string]interface{} {
		var propvalidationerrors []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			propvalidationerror := make(map[string]interface{})
			propvalidationerror["error_log"] = item.ErrorLog
			propvalidationerror["field"] = item.Field
			propvalidationerror["object_type"] = item.ObjectType
			propvalidationerror["task_name"] = item.TaskName
			propvalidationerror["transition_name"] = item.TransitionName
			propvalidationerrors = append(propvalidationerrors, propvalidationerror)
		}
		return propvalidationerrors
	})(item.ValidationError, d)

	propvalidationinformations = append(propvalidationinformations, propvalidationinformation)
	return propvalidationinformations
}
func flattenListRecoveryBackupProfileRef(p []*models.RecoveryBackupProfileRef, d *schema.ResourceData) []map[string]interface{} {
	var propbackupprofiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propbackupprofiles := make(map[string]interface{})
		propbackupprofiles["moid"] = item.Moid
		propbackupprofiles["object_type"] = item.ObjectType
		propbackupprofiles["selector"] = item.Selector
		propbackupprofiless = append(propbackupprofiless, propbackupprofiles)
	}
	return propbackupprofiless
}
func flattenMapSoftwarerepositoryCatalogRef(p *models.SoftwarerepositoryCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var propcatalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcatalog := make(map[string]interface{})
	propcatalog["moid"] = item.Moid
	propcatalog["object_type"] = item.ObjectType
	propcatalog["selector"] = item.Selector

	propcatalogs = append(propcatalogs, propcatalog)
	return propcatalogs
}
func flattenMapSoftwarerepositoryFileServer(p *models.SoftwarerepositoryFileServer, d *schema.ResourceData) []map[string]interface{} {

	var propsources []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsource := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			propsource["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propsource["object_type"] = item.ObjectType

	propsources = append(propsources, propsource)
	return propsources
}
func flattenListStorageDiskGroupPolicyRef(p []*models.StorageDiskGroupPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var propdiskgrouppoliciess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propdiskgrouppolicies := make(map[string]interface{})
		propdiskgrouppolicies["moid"] = item.Moid
		propdiskgrouppolicies["object_type"] = item.ObjectType
		propdiskgrouppolicies["selector"] = item.Selector
		propdiskgrouppoliciess = append(propdiskgrouppoliciess, propdiskgrouppolicies)
	}
	return propdiskgrouppoliciess
}
func flattenListStorageVirtualDriveConfig(p []*models.StorageVirtualDriveConfig, d *schema.ResourceData) []map[string]interface{} {
	var propvirtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propvirtualdrives := make(map[string]interface{})
		propvirtualdrives["access_policy"] = item.AccessPolicy
		delete(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1, "ObjectType")
		if len(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1) != 0 {

			j, err := json.Marshal(item.StorageVirtualDriveConfigAO1P1.StorageVirtualDriveConfigAO1P1)
			if err == nil {
				propvirtualdrives["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propvirtualdrives["boot_drive"] = item.BootDrive
		propvirtualdrives["disk_group_name"] = item.DiskGroupName
		propvirtualdrives["disk_group_policy"] = item.DiskGroupPolicy
		propvirtualdrives["drive_cache"] = item.DriveCache
		propvirtualdrives["expand_to_available"] = item.ExpandToAvailable
		propvirtualdrives["io_policy"] = item.IoPolicy
		propvirtualdrives["name"] = item.Name
		propvirtualdrives["object_type"] = item.ObjectType
		propvirtualdrives["read_policy"] = item.ReadPolicy
		propvirtualdrives["size"] = item.Size
		propvirtualdrives["write_policy"] = item.WritePolicy
		propvirtualdrivess = append(propvirtualdrivess, propvirtualdrives)
	}
	return propvirtualdrivess
}
func flattenListSdcardPartition(p []*models.SdcardPartition, d *schema.ResourceData) []map[string]interface{} {
	var proppartitionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppartitions := make(map[string]interface{})
		delete(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1, "ObjectType")
		if len(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1) != 0 {

			j, err := json.Marshal(item.SdcardPartitionAO1P1.SdcardPartitionAO1P1)
			if err == nil {
				proppartitions["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		proppartitions["object_type"] = item.ObjectType
		proppartitions["type"] = item.Type
		proppartitions["virtual_drives"] = (func(p []*models.SdcardVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
			var propvirtualdrivess []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				propvirtualdrives := make(map[string]interface{})
				delete(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1, "ObjectType")
				if len(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1) != 0 {

					j, err := json.Marshal(item.SdcardVirtualDriveAO1P1.SdcardVirtualDriveAO1P1)
					if err == nil {
						propvirtualdrives["additional_properties"] = string(j)
					} else {
						log.Printf("Error occured while flattening and json parsing: %s", err)
					}
				}
				propvirtualdrives["enable"] = item.Enable
				propvirtualdrives["object_type"] = item.ObjectType
				propvirtualdrivess = append(propvirtualdrivess, propvirtualdrives)
			}
			return propvirtualdrivess
		})(item.VirtualDrives, d)
		proppartitionss = append(proppartitionss, proppartitions)
	}
	return proppartitionss
}
func flattenMapSoftwareHyperflexDistributableRef(p *models.SoftwareHyperflexDistributableRef, d *schema.ResourceData) []map[string]interface{} {

	var prophxdpversioninfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	prophxdpversioninfo := make(map[string]interface{})
	prophxdpversioninfo["moid"] = item.Moid
	prophxdpversioninfo["object_type"] = item.ObjectType
	prophxdpversioninfo["selector"] = item.Selector

	prophxdpversioninfos = append(prophxdpversioninfos, prophxdpversioninfo)
	return prophxdpversioninfos
}
func flattenMapFirmwareDistributableRef(p *models.FirmwareDistributableRef, d *schema.ResourceData) []map[string]interface{} {

	var propserverfirmwareversioninfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propserverfirmwareversioninfo := make(map[string]interface{})
	propserverfirmwareversioninfo["moid"] = item.Moid
	propserverfirmwareversioninfo["object_type"] = item.ObjectType
	propserverfirmwareversioninfo["selector"] = item.Selector

	propserverfirmwareversioninfos = append(propserverfirmwareversioninfos, propserverfirmwareversioninfo)
	return propserverfirmwareversioninfos
}
func flattenMapVnicVsanSettings(p *models.VnicVsanSettings, d *schema.ResourceData) []map[string]interface{} {

	var propvsansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvsansettings := make(map[string]interface{})
	delete(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1, "ObjectType")
	if len(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVsanSettingsAO1P1.VnicVsanSettingsAO1P1)
		if err == nil {
			propvsansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propvsansettings["id"] = item.ID
	propvsansettings["object_type"] = item.ObjectType

	propvsansettingss = append(propvsansettingss, propvsansettings)
	return propvsansettingss
}
func flattenMapIamQualifierRef(p *models.IamQualifierRef, d *schema.ResourceData) []map[string]interface{} {

	var propqualifiers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propqualifier := make(map[string]interface{})
	propqualifier["moid"] = item.Moid
	propqualifier["object_type"] = item.ObjectType
	propqualifier["selector"] = item.Selector

	propqualifiers = append(propqualifiers, propqualifier)
	return propqualifiers
}
func flattenMapRecoveryBackupConfigPolicyRef(p *models.RecoveryBackupConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propbackupconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propbackupconfig := make(map[string]interface{})
	propbackupconfig["moid"] = item.Moid
	propbackupconfig["object_type"] = item.ObjectType
	propbackupconfig["selector"] = item.Selector

	propbackupconfigs = append(propbackupconfigs, propbackupconfig)
	return propbackupconfigs
}
func flattenMapRecoveryScheduleConfigPolicyRef(p *models.RecoveryScheduleConfigPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propscheduleconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propscheduleconfig := make(map[string]interface{})
	propscheduleconfig["moid"] = item.Moid
	propscheduleconfig["object_type"] = item.ObjectType
	propscheduleconfig["selector"] = item.Selector

	propscheduleconfigs = append(propscheduleconfigs, propscheduleconfig)
	return propscheduleconfigs
}
func flattenListResourceGroupRef(p []*models.ResourceGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var propresourcegroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propresourcegroups := make(map[string]interface{})
		propresourcegroups["moid"] = item.Moid
		propresourcegroups["object_type"] = item.ObjectType
		propresourcegroups["selector"] = item.Selector
		propresourcegroupss = append(propresourcegroupss, propresourcegroups)
	}
	return propresourcegroupss
}
func flattenMapIamUserGroupRef(p *models.IamUserGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var propusergroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propusergroup := make(map[string]interface{})
	propusergroup["moid"] = item.Moid
	propusergroup["object_type"] = item.ObjectType
	propusergroup["selector"] = item.Selector

	propusergroups = append(propusergroups, propusergroup)
	return propusergroups
}
func flattenListWorkflowTaskDefinitionRef(p []*models.WorkflowTaskDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var propimplementedtaskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propimplementedtasks := make(map[string]interface{})
		propimplementedtasks["moid"] = item.Moid
		propimplementedtasks["object_type"] = item.ObjectType
		propimplementedtasks["selector"] = item.Selector
		propimplementedtaskss = append(propimplementedtaskss, propimplementedtasks)
	}
	return propimplementedtaskss
}
func flattenMapWorkflowInternalProperties(p *models.WorkflowInternalProperties, d *schema.ResourceData) []map[string]interface{} {

	var propinternalpropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propinternalproperties := make(map[string]interface{})
	delete(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowInternalPropertiesAO1P1.WorkflowInternalPropertiesAO1P1)
		if err == nil {
			propinternalproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propinternalproperties["base_task_type"] = item.BaseTaskType
	propinternalproperties["constraints"] = item.Constraints
	propinternalproperties["internal"] = item.Internal
	propinternalproperties["object_type"] = item.ObjectType
	propinternalproperties["owner"] = item.Owner

	propinternalpropertiess = append(propinternalpropertiess, propinternalproperties)
	return propinternalpropertiess
}
func flattenMapWorkflowProperties(p *models.WorkflowProperties, d *schema.ResourceData) []map[string]interface{} {

	var proppropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propproperties := make(map[string]interface{})
	delete(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1, "ObjectType")
	if len(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1) != 0 {
		j, err := json.Marshal(item.WorkflowPropertiesAO1P1.WorkflowPropertiesAO1P1)
		if err == nil {
			propproperties["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propproperties["input_definition"] = (func(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var propinputdefinitions []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			propinputdefinition := make(map[string]interface{})
			propinputdefinition["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {

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
			propinputdefinition["description"] = item.Description
			propinputdefinition["label"] = item.Label
			propinputdefinition["name"] = item.Name
			propinputdefinition["object_type"] = item.ObjectType
			propinputdefinition["required"] = item.Required
			propinputdefinitions = append(propinputdefinitions, propinputdefinition)
		}
		return propinputdefinitions
	})(item.InputDefinition, d)
	propproperties["object_type"] = item.ObjectType
	propproperties["output_definition"] = (func(p []*models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var propoutputdefinitions []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			propoutputdefinition := make(map[string]interface{})
			propoutputdefinition["default"] = (func(p *models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {

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
			propoutputdefinition["description"] = item.Description
			propoutputdefinition["label"] = item.Label
			propoutputdefinition["name"] = item.Name
			propoutputdefinition["object_type"] = item.ObjectType
			propoutputdefinition["required"] = item.Required
			propoutputdefinitions = append(propoutputdefinitions, propoutputdefinition)
		}
		return propoutputdefinitions
	})(item.OutputDefinition, d)
	propproperties["retry_count"] = item.RetryCount
	propproperties["retry_delay"] = item.RetryDelay
	propproperties["retry_policy"] = item.RetryPolicy
	propproperties["timeout"] = item.Timeout
	propproperties["timeout_policy"] = item.TimeoutPolicy

	proppropertiess = append(proppropertiess, propproperties)
	return proppropertiess
}
func flattenMapIamEndPointUserRef(p *models.IamEndPointUserRef, d *schema.ResourceData) []map[string]interface{} {

	var propendpointusers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propendpointuser := make(map[string]interface{})
	propendpointuser["moid"] = item.Moid
	propendpointuser["object_type"] = item.ObjectType
	propendpointuser["selector"] = item.Selector

	propendpointusers = append(propendpointusers, propendpointuser)
	return propendpointusers
}
func flattenMapIamEndPointUserPolicyRef(p *models.IamEndPointUserPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propendpointuserpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propendpointuserpolicy := make(map[string]interface{})
	propendpointuserpolicy["moid"] = item.Moid
	propendpointuserpolicy["object_type"] = item.ObjectType
	propendpointuserpolicy["selector"] = item.Selector

	propendpointuserpolicys = append(propendpointuserpolicys, propendpointuserpolicy)
	return propendpointuserpolicys
}
func flattenListIamDomainGroupRef(p []*models.IamDomainGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var propdomaingroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propdomaingroups := make(map[string]interface{})
		propdomaingroups["moid"] = item.Moid
		propdomaingroups["object_type"] = item.ObjectType
		propdomaingroups["selector"] = item.Selector
		propdomaingroupss = append(propdomaingroupss, propdomaingroups)
	}
	return propdomaingroupss
}
func flattenListIamIdpReferenceRef(p []*models.IamIdpReferenceRef, d *schema.ResourceData) []map[string]interface{} {
	var propidpreferencess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propidpreferences := make(map[string]interface{})
		propidpreferences["moid"] = item.Moid
		propidpreferences["object_type"] = item.ObjectType
		propidpreferences["selector"] = item.Selector
		propidpreferencess = append(propidpreferencess, propidpreferences)
	}
	return propidpreferencess
}
func flattenListIamIdpRef(p []*models.IamIdpRef, d *schema.ResourceData) []map[string]interface{} {
	var propidpss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propidps := make(map[string]interface{})
		propidps["moid"] = item.Moid
		propidps["object_type"] = item.ObjectType
		propidps["selector"] = item.Selector
		propidpss = append(propidpss, propidps)
	}
	return propidpss
}
func flattenMapCryptEncryptRef(p *models.CryptEncryptRef, d *schema.ResourceData) []map[string]interface{} {

	var propnr0encrypts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnr0encrypt := make(map[string]interface{})
	propnr0encrypt["moid"] = item.Moid
	propnr0encrypt["object_type"] = item.ObjectType
	propnr0encrypt["selector"] = item.Selector

	propnr0encrypts = append(propnr0encrypts, propnr0encrypt)
	return propnr0encrypts
}
func flattenMapCryptDecryptRef(p *models.CryptDecryptRef, d *schema.ResourceData) []map[string]interface{} {

	var propnr1decrypts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnr1decrypt := make(map[string]interface{})
	propnr1decrypt["moid"] = item.Moid
	propnr1decrypt["object_type"] = item.ObjectType
	propnr1decrypt["selector"] = item.Selector

	propnr1decrypts = append(propnr1decrypts, propnr1decrypt)
	return propnr1decrypts
}
func flattenListIamPrivilegeSetRef(p []*models.IamPrivilegeSetRef, d *schema.ResourceData) []map[string]interface{} {
	var propprivilegesetss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propprivilegesets := make(map[string]interface{})
		propprivilegesets["moid"] = item.Moid
		propprivilegesets["object_type"] = item.ObjectType
		propprivilegesets["selector"] = item.Selector
		propprivilegesetss = append(propprivilegesetss, propprivilegesets)
	}
	return propprivilegesetss
}
func flattenListIamPrivilegeRef(p []*models.IamPrivilegeRef, d *schema.ResourceData) []map[string]interface{} {
	var propprivilegess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propprivileges := make(map[string]interface{})
		propprivileges["moid"] = item.Moid
		propprivileges["object_type"] = item.ObjectType
		propprivileges["selector"] = item.Selector
		propprivilegess = append(propprivilegess, propprivileges)
	}
	return propprivilegess
}
func flattenMapIamResourceLimitsRef(p *models.IamResourceLimitsRef, d *schema.ResourceData) []map[string]interface{} {

	var propresourcelimitss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propresourcelimits := make(map[string]interface{})
	propresourcelimits["moid"] = item.Moid
	propresourcelimits["object_type"] = item.ObjectType
	propresourcelimits["selector"] = item.Selector

	propresourcelimitss = append(propresourcelimitss, propresourcelimits)
	return propresourcelimitss
}
func flattenMapIamSecurityHolderRef(p *models.IamSecurityHolderRef, d *schema.ResourceData) []map[string]interface{} {

	var propsecurityholders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsecurityholder := make(map[string]interface{})
	propsecurityholder["moid"] = item.Moid
	propsecurityholder["object_type"] = item.ObjectType
	propsecurityholder["selector"] = item.Selector

	propsecurityholders = append(propsecurityholders, propsecurityholder)
	return propsecurityholders
}
func flattenMapIamSessionLimitsRef(p *models.IamSessionLimitsRef, d *schema.ResourceData) []map[string]interface{} {

	var propsessionlimitss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsessionlimits := make(map[string]interface{})
	propsessionlimits["moid"] = item.Moid
	propsessionlimits["object_type"] = item.ObjectType
	propsessionlimits["selector"] = item.Selector

	propsessionlimitss = append(propsessionlimitss, propsessionlimits)
	return propsessionlimitss
}
func flattenListVnicFcIfRef(p []*models.VnicFcIfRef, d *schema.ResourceData) []map[string]interface{} {
	var propfcifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propfcifs := make(map[string]interface{})
		propfcifs["moid"] = item.Moid
		propfcifs["object_type"] = item.ObjectType
		propfcifs["selector"] = item.Selector
		propfcifss = append(propfcifss, propfcifs)
	}
	return propfcifss
}
func flattenListVnicEthIfRef(p []*models.VnicEthIfRef, d *schema.ResourceData) []map[string]interface{} {
	var propethifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propethifs := make(map[string]interface{})
		propethifs["moid"] = item.Moid
		propethifs["object_type"] = item.ObjectType
		propethifs["selector"] = item.Selector
		propethifss = append(propethifss, propethifs)
	}
	return propethifss
}
func flattenListIaasConnectorPackRef(p []*models.IaasConnectorPackRef, d *schema.ResourceData) []map[string]interface{} {
	var propconnectorpacks []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propconnectorpack := make(map[string]interface{})
		propconnectorpack["moid"] = item.Moid
		propconnectorpack["object_type"] = item.ObjectType
		propconnectorpack["selector"] = item.Selector
		propconnectorpacks = append(propconnectorpacks, propconnectorpack)
	}
	return propconnectorpacks
}
func flattenListIaasDeviceStatusRef(p []*models.IaasDeviceStatusRef, d *schema.ResourceData) []map[string]interface{} {
	var propdevicestatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propdevicestatus := make(map[string]interface{})
		propdevicestatus["moid"] = item.Moid
		propdevicestatus["object_type"] = item.ObjectType
		propdevicestatus["selector"] = item.Selector
		propdevicestatuss = append(propdevicestatuss, propdevicestatus)
	}
	return propdevicestatuss
}
func flattenMapIaasLicenseInfoRef(p *models.IaasLicenseInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var proplicenseinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proplicenseinfo := make(map[string]interface{})
	proplicenseinfo["moid"] = item.Moid
	proplicenseinfo["object_type"] = item.ObjectType
	proplicenseinfo["selector"] = item.Selector

	proplicenseinfos = append(proplicenseinfos, proplicenseinfo)
	return proplicenseinfos
}
func flattenListIaasMostRunTasksRef(p []*models.IaasMostRunTasksRef, d *schema.ResourceData) []map[string]interface{} {
	var propmostruntaskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propmostruntasks := make(map[string]interface{})
		propmostruntasks["moid"] = item.Moid
		propmostruntasks["object_type"] = item.ObjectType
		propmostruntasks["selector"] = item.Selector
		propmostruntaskss = append(propmostruntaskss, propmostruntasks)
	}
	return propmostruntaskss
}
func flattenMapIaasUcsdManagedInfraRef(p *models.IaasUcsdManagedInfraRef, d *schema.ResourceData) []map[string]interface{} {

	var propucsdmanagedinfras []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propucsdmanagedinfra := make(map[string]interface{})
	propucsdmanagedinfra["moid"] = item.Moid
	propucsdmanagedinfra["object_type"] = item.ObjectType
	propucsdmanagedinfra["selector"] = item.Selector

	propucsdmanagedinfras = append(propucsdmanagedinfras, propucsdmanagedinfra)
	return propucsdmanagedinfras
}
func flattenMapVnicFcAdapterPolicyRef(p *models.VnicFcAdapterPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propfcadapterpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propfcadapterpolicy := make(map[string]interface{})
	propfcadapterpolicy["moid"] = item.Moid
	propfcadapterpolicy["object_type"] = item.ObjectType
	propfcadapterpolicy["selector"] = item.Selector

	propfcadapterpolicys = append(propfcadapterpolicys, propfcadapterpolicy)
	return propfcadapterpolicys
}
func flattenMapVnicFcNetworkPolicyRef(p *models.VnicFcNetworkPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propfcnetworkpolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propfcnetworkpolicy := make(map[string]interface{})
	propfcnetworkpolicy["moid"] = item.Moid
	propfcnetworkpolicy["object_type"] = item.ObjectType
	propfcnetworkpolicy["selector"] = item.Selector

	propfcnetworkpolicys = append(propfcnetworkpolicys, propfcnetworkpolicy)
	return propfcnetworkpolicys
}
func flattenMapVnicFcQosPolicyRef(p *models.VnicFcQosPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propfcqospolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propfcqospolicy := make(map[string]interface{})
	propfcqospolicy["moid"] = item.Moid
	propfcqospolicy["object_type"] = item.ObjectType
	propfcqospolicy["selector"] = item.Selector

	propfcqospolicys = append(propfcqospolicys, propfcqospolicy)
	return propfcqospolicys
}
func flattenMapVnicSanConnectivityPolicyRef(p *models.VnicSanConnectivityPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propsanconnectivitypolicys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsanconnectivitypolicy := make(map[string]interface{})
	propsanconnectivitypolicy["moid"] = item.Moid
	propsanconnectivitypolicy["object_type"] = item.ObjectType
	propsanconnectivitypolicy["selector"] = item.Selector

	propsanconnectivitypolicys = append(propsanconnectivitypolicys, propsanconnectivitypolicy)
	return propsanconnectivitypolicys
}
func flattenListOrganizationOrganizationRef(p []*models.OrganizationOrganizationRef, d *schema.ResourceData) []map[string]interface{} {
	var proporganizationss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proporganizations := make(map[string]interface{})
		proporganizations["moid"] = item.Moid
		proporganizations["object_type"] = item.ObjectType
		proporganizations["selector"] = item.Selector
		proporganizationss = append(proporganizationss, proporganizations)
	}
	return proporganizationss
}
func flattenListResourcePerTypeCombinedSelector(p []*models.ResourcePerTypeCombinedSelector, d *schema.ResourceData) []map[string]interface{} {
	var proppertypecombinedselectors []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppertypecombinedselector := make(map[string]interface{})
		delete(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1, "ObjectType")
		if len(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1) != 0 {

			j, err := json.Marshal(item.ResourcePerTypeCombinedSelectorAO1P1.ResourcePerTypeCombinedSelectorAO1P1)
			if err == nil {
				proppertypecombinedselector["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		proppertypecombinedselector["combined_selector"] = item.CombinedSelector
		proppertypecombinedselector["empty_filter"] = item.EmptyFilter
		proppertypecombinedselector["object_type"] = item.ObjectType
		proppertypecombinedselector["selector_object_type"] = item.SelectorObjectType
		proppertypecombinedselectors = append(proppertypecombinedselectors, proppertypecombinedselector)
	}
	return proppertypecombinedselectors
}
func flattenListResourceSelector(p []*models.ResourceSelector, d *schema.ResourceData) []map[string]interface{} {
	var propselectorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propselectors := make(map[string]interface{})
		delete(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1, "ObjectType")
		if len(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1) != 0 {

			j, err := json.Marshal(item.ResourceSelectorAO1P1.ResourceSelectorAO1P1)
			if err == nil {
				propselectors["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propselectors["object_type"] = item.ObjectType
		propselectors["selector"] = item.Selector
		propselectorss = append(propselectorss, propselectors)
	}
	return propselectorss
}
func flattenListAdapterAdapterConfig(p []*models.AdapterAdapterConfig, d *schema.ResourceData) []map[string]interface{} {
	var propsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsettings := make(map[string]interface{})
		delete(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1, "ObjectType")
		if len(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1) != 0 {

			j, err := json.Marshal(item.AdapterAdapterConfigAO1P1.AdapterAdapterConfigAO1P1)
			if err == nil {
				propsettings["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propsettings["eth_settings"] = (func(p *models.AdapterEthSettings, d *schema.ResourceData) []map[string]interface{} {

			var propethsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propethsettings := make(map[string]interface{})
			delete(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1, "ObjectType")
			if len(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterEthSettingsAO1P1.AdapterEthSettingsAO1P1)
				if err == nil {
					propethsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propethsettings["lldp_enabled"] = item.LldpEnabled
			propethsettings["object_type"] = item.ObjectType

			propethsettingss = append(propethsettingss, propethsettings)
			return propethsettingss
		})(item.EthSettings, d)
		propsettings["fc_settings"] = (func(p *models.AdapterFcSettings, d *schema.ResourceData) []map[string]interface{} {

			var propfcsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propfcsettings := make(map[string]interface{})
			delete(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1, "ObjectType")
			if len(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterFcSettingsAO1P1.AdapterFcSettingsAO1P1)
				if err == nil {
					propfcsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propfcsettings["fip_enabled"] = item.FipEnabled
			propfcsettings["object_type"] = item.ObjectType

			propfcsettingss = append(propfcsettingss, propfcsettings)
			return propfcsettingss
		})(item.FcSettings, d)
		propsettings["object_type"] = item.ObjectType
		propsettings["port_channel_settings"] = (func(p *models.AdapterPortChannelSettings, d *schema.ResourceData) []map[string]interface{} {

			var propportchannelsettingss []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propportchannelsettings := make(map[string]interface{})
			delete(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1, "ObjectType")
			if len(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1) != 0 {
				j, err := json.Marshal(item.AdapterPortChannelSettingsAO1P1.AdapterPortChannelSettingsAO1P1)
				if err == nil {
					propportchannelsettings["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propportchannelsettings["enabled"] = item.Enabled
			propportchannelsettings["object_type"] = item.ObjectType

			propportchannelsettingss = append(propportchannelsettingss, propportchannelsettings)
			return propportchannelsettingss
		})(item.PortChannelSettings, d)
		propsettings["slot_id"] = item.SlotID
		propsettingss = append(propsettingss, propsettings)
	}
	return propsettingss
}
func flattenMapMoBaseMoRef(p *models.MoBaseMoRef, d *schema.ResourceData) []map[string]interface{} {

	var propaffectedobjects []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propaffectedobject := make(map[string]interface{})
	propaffectedobject["moid"] = item.Moid
	propaffectedobject["object_type"] = item.ObjectType
	propaffectedobject["selector"] = item.Selector

	propaffectedobjects = append(propaffectedobjects, propaffectedobject)
	return propaffectedobjects
}
func flattenListHyperflexServerModelEntry(p []*models.HyperflexServerModelEntry, d *schema.ResourceData) []map[string]interface{} {
	var propservermodelentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propservermodelentries := make(map[string]interface{})
		delete(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1, "ObjectType")
		if len(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1) != 0 {

			j, err := json.Marshal(item.HyperflexAbstractAppSettingAO1P1.HyperflexAbstractAppSettingAO1P1)
			if err == nil {
				propservermodelentries["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propservermodelentries["constraint"] = (func(p *models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {

			var propconstraints []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propconstraint := make(map[string]interface{})
			delete(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1, "ObjectType")
			if len(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1) != 0 {
				j, err := json.Marshal(item.HyperflexAppSettingConstraintAO1P1.HyperflexAppSettingConstraintAO1P1)
				if err == nil {
					propconstraint["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propconstraint["hxdp_version"] = item.HxdpVersion
			propconstraint["hypervisor_type"] = item.HypervisorType
			propconstraint["mgmt_platform"] = item.MgmtPlatform
			propconstraint["object_type"] = item.ObjectType
			propconstraint["server_model"] = item.ServerModel

			propconstraints = append(propconstraints, propconstraint)
			return propconstraints
		})(item.Constraint, d)
		propservermodelentries["name"] = item.Name
		propservermodelentries["object_type"] = item.ObjectType
		propservermodelentries["value"] = item.Value
		propservermodelentriess = append(propservermodelentriess, propservermodelentries)
	}
	return propservermodelentriess
}
func flattenMapVnicArfsSettings(p *models.VnicArfsSettings, d *schema.ResourceData) []map[string]interface{} {

	var proparfssettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proparfssettings := make(map[string]interface{})
	delete(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1, "ObjectType")
	if len(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicArfsSettingsAO1P1.VnicArfsSettingsAO1P1)
		if err == nil {
			proparfssettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proparfssettings["enabled"] = item.Enabled
	proparfssettings["object_type"] = item.ObjectType

	proparfssettingss = append(proparfssettingss, proparfssettings)
	return proparfssettingss
}
func flattenMapVnicCompletionQueueSettings(p *models.VnicCompletionQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var propcompletionqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcompletionqueuesettings := make(map[string]interface{})
	delete(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicCompletionQueueSettingsAO1P1.VnicCompletionQueueSettingsAO1P1)
		if err == nil {
			propcompletionqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propcompletionqueuesettings["count"] = item.Count
	propcompletionqueuesettings["object_type"] = item.ObjectType
	propcompletionqueuesettings["ring_size"] = item.RingSize

	propcompletionqueuesettingss = append(propcompletionqueuesettingss, propcompletionqueuesettings)
	return propcompletionqueuesettingss
}
func flattenMapVnicEthInterruptSettings(p *models.VnicEthInterruptSettings, d *schema.ResourceData) []map[string]interface{} {

	var propinterruptsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propinterruptsettings := make(map[string]interface{})
	delete(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1, "ObjectType")
	if len(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthInterruptSettingsAO1P1.VnicEthInterruptSettingsAO1P1)
		if err == nil {
			propinterruptsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propinterruptsettings["coalescing_time"] = item.CoalescingTime
	propinterruptsettings["coalescing_type"] = item.CoalescingType
	propinterruptsettings["count"] = item.Count
	propinterruptsettings["mode"] = item.Mode
	propinterruptsettings["object_type"] = item.ObjectType

	propinterruptsettingss = append(propinterruptsettingss, propinterruptsettings)
	return propinterruptsettingss
}
func flattenMapVnicNvgreSettings(p *models.VnicNvgreSettings, d *schema.ResourceData) []map[string]interface{} {

	var propnvgresettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnvgresettings := make(map[string]interface{})
	delete(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1, "ObjectType")
	if len(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicNvgreSettingsAO1P1.VnicNvgreSettingsAO1P1)
		if err == nil {
			propnvgresettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propnvgresettings["enabled"] = item.Enabled
	propnvgresettings["object_type"] = item.ObjectType

	propnvgresettingss = append(propnvgresettingss, propnvgresettings)
	return propnvgresettingss
}
func flattenMapVnicRoceSettings(p *models.VnicRoceSettings, d *schema.ResourceData) []map[string]interface{} {

	var proprocesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proprocesettings := make(map[string]interface{})
	delete(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1, "ObjectType")
	if len(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicRoceSettingsAO1P1.VnicRoceSettingsAO1P1)
		if err == nil {
			proprocesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proprocesettings["enabled"] = item.Enabled
	proprocesettings["memory_regions"] = item.MemoryRegions
	proprocesettings["object_type"] = item.ObjectType
	proprocesettings["queue_pairs"] = item.QueuePairs
	proprocesettings["resource_groups"] = item.ResourceGroups

	proprocesettingss = append(proprocesettingss, proprocesettings)
	return proprocesettingss
}
func flattenMapVnicEthRxQueueSettings(p *models.VnicEthRxQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var proprxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proprxqueuesettings := make(map[string]interface{})
	delete(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthRxQueueSettingsAO1P1.VnicEthRxQueueSettingsAO1P1)
		if err == nil {
			proprxqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proprxqueuesettings["count"] = item.Count
	proprxqueuesettings["object_type"] = item.ObjectType
	proprxqueuesettings["ring_size"] = item.RingSize

	proprxqueuesettingss = append(proprxqueuesettingss, proprxqueuesettings)
	return proprxqueuesettingss
}
func flattenMapVnicTCPOffloadSettings(p *models.VnicTCPOffloadSettings, d *schema.ResourceData) []map[string]interface{} {

	var proptcpoffloadsettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proptcpoffloadsettings := make(map[string]interface{})
	delete(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1, "ObjectType")
	if len(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicTCPOffloadSettingsAO1P1.VnicTCPOffloadSettingsAO1P1)
		if err == nil {
			proptcpoffloadsettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proptcpoffloadsettings["large_receive"] = item.LargeReceive
	proptcpoffloadsettings["large_send"] = item.LargeSend
	proptcpoffloadsettings["object_type"] = item.ObjectType
	proptcpoffloadsettings["rx_checksum"] = item.RxChecksum
	proptcpoffloadsettings["tx_checksum"] = item.TxChecksum

	proptcpoffloadsettingss = append(proptcpoffloadsettingss, proptcpoffloadsettings)
	return proptcpoffloadsettingss
}
func flattenMapVnicEthTxQueueSettings(p *models.VnicEthTxQueueSettings, d *schema.ResourceData) []map[string]interface{} {

	var proptxqueuesettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proptxqueuesettings := make(map[string]interface{})
	delete(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1, "ObjectType")
	if len(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicEthTxQueueSettingsAO1P1.VnicEthTxQueueSettingsAO1P1)
		if err == nil {
			proptxqueuesettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	proptxqueuesettings["count"] = item.Count
	proptxqueuesettings["object_type"] = item.ObjectType
	proptxqueuesettings["ring_size"] = item.RingSize

	proptxqueuesettingss = append(proptxqueuesettingss, proptxqueuesettings)
	return proptxqueuesettingss
}
func flattenMapVnicVxlanSettings(p *models.VnicVxlanSettings, d *schema.ResourceData) []map[string]interface{} {

	var propvxlansettingss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvxlansettings := make(map[string]interface{})
	delete(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1, "ObjectType")
	if len(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1) != 0 {
		j, err := json.Marshal(item.VnicVxlanSettingsAO1P1.VnicVxlanSettingsAO1P1)
		if err == nil {
			propvxlansettings["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propvxlansettings["enabled"] = item.Enabled
	propvxlansettings["object_type"] = item.ObjectType

	propvxlansettingss = append(propvxlansettingss, propvxlansettings)
	return propvxlansettingss
}
func flattenMapRecoveryBackupSchedule(p *models.RecoveryBackupSchedule, d *schema.ResourceData) []map[string]interface{} {

	var propschedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propschedule := make(map[string]interface{})
	delete(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1, "ObjectType")
	if len(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1) != 0 {
		j, err := json.Marshal(item.RecoveryBackupScheduleAO1P1.RecoveryBackupScheduleAO1P1)
		if err == nil {
			propschedule["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propschedule["execution_time"] = item.ExecutionTime
	propschedule["frequency_unit"] = item.FrequencyUnit
	propschedule["hours"] = item.Hours
	propschedule["object_type"] = item.ObjectType

	propschedules = append(propschedules, propschedule)
	return propschedules
}
func flattenListVmediaMapping(p []*models.VmediaMapping, d *schema.ResourceData) []map[string]interface{} {
	var propmappingss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propmappings := make(map[string]interface{})
		delete(item.VmediaMappingAO1P1.VmediaMappingAO1P1, "ObjectType")
		if len(item.VmediaMappingAO1P1.VmediaMappingAO1P1) != 0 {

			j, err := json.Marshal(item.VmediaMappingAO1P1.VmediaMappingAO1P1)
			if err == nil {
				propmappings["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propmappings["authentication_protocol"] = item.AuthenticationProtocol
		propmappings["device_type"] = item.DeviceType
		propmappings["host_name"] = item.HostName
		propmappings["is_password_set"] = item.IsPasswordSet
		propmappings["mount_options"] = item.MountOptions
		propmappings["mount_protocol"] = item.MountProtocol
		propmappings["object_type"] = item.ObjectType
		propmappings["password"] = item.Password
		propmappings["remote_file"] = item.RemoteFile
		propmappings["remote_path"] = item.RemotePath
		propmappings["username"] = item.Username
		propmappings["volume_name"] = item.VolumeName
		propmappingss = append(propmappingss, propmappings)
	}
	return propmappingss
}
func flattenMapHyperflexFeatureLimitExternalRef(p *models.HyperflexFeatureLimitExternalRef, d *schema.ResourceData) []map[string]interface{} {

	var propfeaturelimitexternals []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propfeaturelimitexternal := make(map[string]interface{})
	propfeaturelimitexternal["moid"] = item.Moid
	propfeaturelimitexternal["object_type"] = item.ObjectType
	propfeaturelimitexternal["selector"] = item.Selector

	propfeaturelimitexternals = append(propfeaturelimitexternals, propfeaturelimitexternal)
	return propfeaturelimitexternals
}
func flattenMapHyperflexFeatureLimitInternalRef(p *models.HyperflexFeatureLimitInternalRef, d *schema.ResourceData) []map[string]interface{} {

	var propfeaturelimitinternals []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propfeaturelimitinternal := make(map[string]interface{})
	propfeaturelimitinternal["moid"] = item.Moid
	propfeaturelimitinternal["object_type"] = item.ObjectType
	propfeaturelimitinternal["selector"] = item.Selector

	propfeaturelimitinternals = append(propfeaturelimitinternals, propfeaturelimitinternal)
	return propfeaturelimitinternals
}
func flattenListHyperflexHxdpVersionRef(p []*models.HyperflexHxdpVersionRef, d *schema.ResourceData) []map[string]interface{} {
	var prophxdpversionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		prophxdpversions := make(map[string]interface{})
		prophxdpversions["moid"] = item.Moid
		prophxdpversions["object_type"] = item.ObjectType
		prophxdpversions["selector"] = item.Selector
		prophxdpversionss = append(prophxdpversionss, prophxdpversions)
	}
	return prophxdpversionss
}
func flattenListHyperflexCapabilityInfoRef(p []*models.HyperflexCapabilityInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var prophyperflexcapabilityinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		prophyperflexcapabilityinfos := make(map[string]interface{})
		prophyperflexcapabilityinfos["moid"] = item.Moid
		prophyperflexcapabilityinfos["object_type"] = item.ObjectType
		prophyperflexcapabilityinfos["selector"] = item.Selector
		prophyperflexcapabilityinfoss = append(prophyperflexcapabilityinfoss, prophyperflexcapabilityinfos)
	}
	return prophyperflexcapabilityinfoss
}
func flattenListHclHyperflexSoftwareCompatibilityInfoRef(p []*models.HclHyperflexSoftwareCompatibilityInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var prophyperflexsoftwarecompatibilityinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		prophyperflexsoftwarecompatibilityinfos := make(map[string]interface{})
		prophyperflexsoftwarecompatibilityinfos["moid"] = item.Moid
		prophyperflexsoftwarecompatibilityinfos["object_type"] = item.ObjectType
		prophyperflexsoftwarecompatibilityinfos["selector"] = item.Selector
		prophyperflexsoftwarecompatibilityinfoss = append(prophyperflexsoftwarecompatibilityinfoss, prophyperflexsoftwarecompatibilityinfos)
	}
	return prophyperflexsoftwarecompatibilityinfoss
}
func flattenMapHyperflexServerFirmwareVersionRef(p *models.HyperflexServerFirmwareVersionRef, d *schema.ResourceData) []map[string]interface{} {

	var propserverfirmwareversions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propserverfirmwareversion := make(map[string]interface{})
	propserverfirmwareversion["moid"] = item.Moid
	propserverfirmwareversion["object_type"] = item.ObjectType
	propserverfirmwareversion["selector"] = item.Selector

	propserverfirmwareversions = append(propserverfirmwareversions, propserverfirmwareversion)
	return propserverfirmwareversions
}
func flattenMapHyperflexServerModelRef(p *models.HyperflexServerModelRef, d *schema.ResourceData) []map[string]interface{} {

	var propservermodels []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propservermodel := make(map[string]interface{})
	propservermodel["moid"] = item.Moid
	propservermodel["object_type"] = item.ObjectType
	propservermodel["selector"] = item.Selector

	propservermodels = append(propservermodels, propservermodel)
	return propservermodels
}
func flattenMapHyperflexNamedVsan(p *models.HyperflexNamedVsan, d *schema.ResourceData) []map[string]interface{} {

	var propextatraffics []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propextatraffic := make(map[string]interface{})
	delete(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1, "ObjectType")
	if len(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexNamedVsanAO1P1.HyperflexNamedVsanAO1P1)
		if err == nil {
			propextatraffic["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propextatraffic["name"] = item.Name
	propextatraffic["object_type"] = item.ObjectType
	propextatraffic["vsan_id"] = item.VsanID

	propextatraffics = append(propextatraffics, propextatraffic)
	return propextatraffics
}
func flattenMapHyperflexWwxnPrefixRange(p *models.HyperflexWwxnPrefixRange, d *schema.ResourceData) []map[string]interface{} {

	var propwwxnprefixranges []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propwwxnprefixrange := make(map[string]interface{})
	delete(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1, "ObjectType")
	if len(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1) != 0 {
		j, err := json.Marshal(item.HyperflexWwxnPrefixRangeAO1P1.HyperflexWwxnPrefixRangeAO1P1)
		if err == nil {
			propwwxnprefixrange["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propwwxnprefixrange["end_addr"] = item.EndAddr
	propwwxnprefixrange["object_type"] = item.ObjectType
	propwwxnprefixrange["start_addr"] = item.StartAddr

	propwwxnprefixranges = append(propwwxnprefixranges, propwwxnprefixrange)
	return propwwxnprefixranges
}
func flattenMapOsCatalogRef(p *models.OsCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var propcatalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcatalog := make(map[string]interface{})
	propcatalog["moid"] = item.Moid
	propcatalog["object_type"] = item.ObjectType
	propcatalog["selector"] = item.Selector

	propcatalogs = append(propcatalogs, propcatalog)
	return propcatalogs
}
func flattenListHclOperatingSystemRef(p []*models.HclOperatingSystemRef, d *schema.ResourceData) []map[string]interface{} {
	var propdistributionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propdistributions := make(map[string]interface{})
		propdistributions["moid"] = item.Moid
		propdistributions["object_type"] = item.ObjectType
		propdistributions["selector"] = item.Selector
		propdistributionss = append(propdistributionss, propdistributions)
	}
	return propdistributionss
}
func flattenListOsPlaceHolder(p []*models.OsPlaceHolder, d *schema.ResourceData) []map[string]interface{} {
	var propplaceholderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propplaceholders := make(map[string]interface{})
		delete(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1, "ObjectType")
		if len(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1) != 0 {

			j, err := json.Marshal(item.OsPlaceHolderAO1P1.OsPlaceHolderAO1P1)
			if err == nil {
				propplaceholders["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propplaceholders["is_value_set"] = item.IsValueSet
		propplaceholders["object_type"] = item.ObjectType
		propplaceholders["type"] = (func(p *models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {

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

				var proppropertiess []map[string]interface{}
				if p == nil {
					return nil
				}
				item := *p
				propproperties := make(map[string]interface{})
				propproperties["constraints"] = (func(p *models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {

					var propconstraintss []map[string]interface{}
					if p == nil {
						return nil
					}
					item := *p
					propconstraints := make(map[string]interface{})
					propconstraints["enum_list"] = (func(p []*models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
						var propenumlists []map[string]interface{}
						if p == nil {
							return nil
						}
						for _, item := range p {
							item := *item
							propenumlist := make(map[string]interface{})
							propenumlist["label"] = item.Label
							propenumlist["object_type"] = item.ObjectType
							propenumlist["value"] = item.Value
							propenumlists = append(propenumlists, propenumlist)
						}
						return propenumlists
					})(item.EnumList, d)
					propconstraints["max"] = item.Max
					propconstraints["min"] = item.Min
					propconstraints["object_type"] = item.ObjectType
					propconstraints["regex"] = item.Regex

					propconstraintss = append(propconstraintss, propconstraints)
					return propconstraintss
				})(item.Constraints, d)
				propproperties["object_type"] = item.ObjectType
				propproperties["secure"] = item.Secure
				propproperties["type"] = item.Type

				proppropertiess = append(proppropertiess, propproperties)
				return proppropertiess
			})(item.Properties, d)
			proptype["required"] = item.Required

			proptypes = append(proptypes, proptype)
			return proptypes
		})(item.Type, d)
		propplaceholders["value"] = item.Value
		propplaceholderss = append(propplaceholderss, propplaceholders)
	}
	return propplaceholderss
}
func flattenListX509Certificate(p []*models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var propcertificatess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcertificates := make(map[string]interface{})
		delete(item.X509CertificateAO1P1.X509CertificateAO1P1, "ObjectType")
		if len(item.X509CertificateAO1P1.X509CertificateAO1P1) != 0 {

			j, err := json.Marshal(item.X509CertificateAO1P1.X509CertificateAO1P1)
			if err == nil {
				propcertificates["additional_properties"] = string(j)
			} else {
				log.Printf("Error occured while flattening and json parsing: %s", err)
			}
		}
		propcertificates["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var propissuers []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propissuer := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					propissuer["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propissuer["common_name"] = item.CommonName
			propissuer["country"] = item.Country
			propissuer["locality"] = item.Locality
			propissuer["object_type"] = item.ObjectType
			propissuer["organization"] = item.Organization
			propissuer["organizational_unit"] = item.OrganizationalUnit
			propissuer["state"] = item.State

			propissuers = append(propissuers, propissuer)
			return propissuers
		})(item.Issuer, d)
		propcertificates["object_type"] = item.ObjectType
		propcertificates["pem_certificate"] = item.PemCertificate
		propcertificates["sha256_fingerprint"] = item.Sha256Fingerprint
		propcertificates["signature_algorithm"] = item.SignatureAlgorithm
		propcertificates["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var propsubjects []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propsubject := make(map[string]interface{})
			delete(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1, "ObjectType")
			if len(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1) != 0 {
				j, err := json.Marshal(item.PkixDistinguishedNameAO1P1.PkixDistinguishedNameAO1P1)
				if err == nil {
					propsubject["additional_properties"] = string(j)
				} else {
					log.Printf("Error occured while flattening and json parsing: %s", err)
				}
			}
			propsubject["common_name"] = item.CommonName
			propsubject["country"] = item.Country
			propsubject["locality"] = item.Locality
			propsubject["object_type"] = item.ObjectType
			propsubject["organization"] = item.Organization
			propsubject["organizational_unit"] = item.OrganizationalUnit
			propsubject["state"] = item.State

			propsubjects = append(propsubjects, propsubject)
			return propsubjects
		})(item.Subject, d)
		propcertificatess = append(propcertificatess, propcertificates)
	}
	return propcertificatess
}
func flattenMapOsAnswers(p *models.OsAnswers, d *schema.ResourceData) []map[string]interface{} {

	var propanswerss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propanswers := make(map[string]interface{})
	delete(item.OsAnswersAO1P1.OsAnswersAO1P1, "ObjectType")
	if len(item.OsAnswersAO1P1.OsAnswersAO1P1) != 0 {
		j, err := json.Marshal(item.OsAnswersAO1P1.OsAnswersAO1P1)
		if err == nil {
			propanswers["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propanswers["answer_file"] = item.AnswerFile
	propanswers["hostname"] = item.Hostname
	propanswers["ip_config_type"] = item.IPConfigType
	propanswers["ipv4_config"] = (func(p *models.CommIPV4Interface, d *schema.ResourceData) []map[string]interface{} {

		var propipv4configs []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propipv4config := make(map[string]interface{})
		propipv4config["gateway"] = item.Gateway
		propipv4config["ip_address"] = item.IPAddress
		propipv4config["netmask"] = item.Netmask
		propipv4config["object_type"] = item.ObjectType

		propipv4configs = append(propipv4configs, propipv4config)
		return propipv4configs
	})(item.IPV4Config, d)
	propanswers["is_answer_file_set"] = item.IsAnswerFileSet
	propanswers["is_root_password_set"] = item.IsRootPasswordSet
	propanswers["nameserver"] = item.Nameserver
	propanswers["object_type"] = item.ObjectType
	propanswers["product_key"] = item.ProductKey
	propanswers["root_password"] = item.RootPassword
	propanswers["source"] = item.Source

	propanswerss = append(propanswerss, propanswers)
	return propanswerss
}
func flattenMapOsConfigurationFileRef(p *models.OsConfigurationFileRef, d *schema.ResourceData) []map[string]interface{} {

	var propconfigurationfiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propconfigurationfile := make(map[string]interface{})
	propconfigurationfile["moid"] = item.Moid
	propconfigurationfile["object_type"] = item.ObjectType
	propconfigurationfile["selector"] = item.Selector

	propconfigurationfiles = append(propconfigurationfiles, propconfigurationfile)
	return propconfigurationfiles
}
func flattenMapSoftwarerepositoryOperatingSystemFileRef(p *models.SoftwarerepositoryOperatingSystemFileRef, d *schema.ResourceData) []map[string]interface{} {

	var propimages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propimage := make(map[string]interface{})
	propimage["moid"] = item.Moid
	propimage["object_type"] = item.ObjectType
	propimage["selector"] = item.Selector

	propimages = append(propimages, propimage)
	return propimages
}
func flattenMapOsOperatingSystemParameters(p *models.OsOperatingSystemParameters, d *schema.ResourceData) []map[string]interface{} {

	var propoperatingsystemparameterss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propoperatingsystemparameters := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			propoperatingsystemparameters["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propoperatingsystemparameters["object_type"] = item.ObjectType

	propoperatingsystemparameterss = append(propoperatingsystemparameterss, propoperatingsystemparameters)
	return propoperatingsystemparameterss
}
func flattenMapFirmwareServerConfigurationUtilityDistributableRef(p *models.FirmwareServerConfigurationUtilityDistributableRef, d *schema.ResourceData) []map[string]interface{} {

	var proposduimages []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proposduimage := make(map[string]interface{})
	proposduimage["moid"] = item.Moid
	proposduimage["object_type"] = item.ObjectType
	proposduimage["selector"] = item.Selector

	proposduimages = append(proposduimages, proposduimage)
	return proposduimages
}
func flattenListOsPostInstallScriptRef(p []*models.OsPostInstallScriptRef, d *schema.ResourceData) []map[string]interface{} {
	var proppostinstallscriptss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppostinstallscripts := make(map[string]interface{})
		proppostinstallscripts["moid"] = item.Moid
		proppostinstallscripts["object_type"] = item.ObjectType
		proppostinstallscripts["selector"] = item.Selector
		proppostinstallscriptss = append(proppostinstallscriptss, proppostinstallscripts)
	}
	return proppostinstallscriptss
}
func flattenMapComputePhysicalRef(p *models.ComputePhysicalRef, d *schema.ResourceData) []map[string]interface{} {

	var propservers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propserver := make(map[string]interface{})
	propserver["moid"] = item.Moid
	propserver["object_type"] = item.ObjectType
	propserver["selector"] = item.Selector

	propservers = append(propservers, propserver)
	return propservers
}
func flattenMapFirmwareDirectDownload(p *models.FirmwareDirectDownload, d *schema.ResourceData) []map[string]interface{} {

	var propdirectdownloads []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdirectdownload := make(map[string]interface{})
	delete(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1, "ObjectType")
	if len(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1) != 0 {
		j, err := json.Marshal(item.FirmwareDirectDownloadAO1P1.FirmwareDirectDownloadAO1P1)
		if err == nil {
			propdirectdownload["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propdirectdownload["http_server"] = (func(p *models.FirmwareHTTPServer, d *schema.ResourceData) []map[string]interface{} {

		var prophttpservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		prophttpserver := make(map[string]interface{})
		prophttpserver["location_link"] = item.LocationLink
		prophttpserver["mount_options"] = item.MountOptions
		prophttpserver["object_type"] = item.ObjectType

		prophttpservers = append(prophttpservers, prophttpserver)
		return prophttpservers
	})(item.HTTPServer, d)
	propdirectdownload["image_source"] = item.ImageSource
	propdirectdownload["is_password_set"] = item.IsPasswordSet
	propdirectdownload["object_type"] = item.ObjectType
	propdirectdownload["password"] = item.Password
	propdirectdownload["upgradeoption"] = item.Upgradeoption
	propdirectdownload["username"] = item.Username

	propdirectdownloads = append(propdirectdownloads, propdirectdownload)
	return propdirectdownloads
}
func flattenMapFirmwareNetworkShare(p *models.FirmwareNetworkShare, d *schema.ResourceData) []map[string]interface{} {

	var propnetworkshares []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnetworkshare := make(map[string]interface{})
	delete(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1, "ObjectType")
	if len(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1) != 0 {
		j, err := json.Marshal(item.FirmwareNetworkShareAO1P1.FirmwareNetworkShareAO1P1)
		if err == nil {
			propnetworkshare["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propnetworkshare["cifs_server"] = (func(p *models.FirmwareCifsServer, d *schema.ResourceData) []map[string]interface{} {

		var propcifsservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propcifsserver := make(map[string]interface{})
		propcifsserver["mount_options"] = item.MountOptions
		propcifsserver["object_type"] = item.ObjectType
		propcifsserver["remote_file"] = item.RemoteFile
		propcifsserver["remote_ip"] = item.RemoteIP
		propcifsserver["remote_share"] = item.RemoteShare

		propcifsservers = append(propcifsservers, propcifsserver)
		return propcifsservers
	})(item.CifsServer, d)
	propnetworkshare["http_server"] = (func(p *models.FirmwareHTTPServer, d *schema.ResourceData) []map[string]interface{} {

		var prophttpservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		prophttpserver := make(map[string]interface{})
		prophttpserver["location_link"] = item.LocationLink
		prophttpserver["mount_options"] = item.MountOptions
		prophttpserver["object_type"] = item.ObjectType

		prophttpservers = append(prophttpservers, prophttpserver)
		return prophttpservers
	})(item.HTTPServer, d)
	propnetworkshare["is_password_set"] = item.IsPasswordSet
	propnetworkshare["map_type"] = item.MapType
	propnetworkshare["nfs_server"] = (func(p *models.FirmwareNfsServer, d *schema.ResourceData) []map[string]interface{} {

		var propnfsservers []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propnfsserver := make(map[string]interface{})
		propnfsserver["mount_options"] = item.MountOptions
		propnfsserver["object_type"] = item.ObjectType
		propnfsserver["remote_file"] = item.RemoteFile
		propnfsserver["remote_ip"] = item.RemoteIP
		propnfsserver["remote_share"] = item.RemoteShare

		propnfsservers = append(propnfsservers, propnfsserver)
		return propnfsservers
	})(item.NfsServer, d)
	propnetworkshare["object_type"] = item.ObjectType
	propnetworkshare["password"] = item.Password
	propnetworkshare["upgradeoption"] = item.Upgradeoption
	propnetworkshare["username"] = item.Username

	propnetworkshares = append(propnetworkshares, propnetworkshare)
	return propnetworkshares
}
func flattenMapFirmwareUpgradeStatusRef(p *models.FirmwareUpgradeStatusRef, d *schema.ResourceData) []map[string]interface{} {

	var propupgradestatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propupgradestatus := make(map[string]interface{})
	propupgradestatus["moid"] = item.Moid
	propupgradestatus["object_type"] = item.ObjectType
	propupgradestatus["selector"] = item.Selector

	propupgradestatuss = append(propupgradestatuss, propupgradestatus)
	return propupgradestatuss
}
func flattenMapRecoveryAbstractBackupInfoRef(p *models.RecoveryAbstractBackupInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var propbackupinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propbackupinfo := make(map[string]interface{})
	propbackupinfo["moid"] = item.Moid
	propbackupinfo["object_type"] = item.ObjectType
	propbackupinfo["selector"] = item.Selector

	propbackupinfos = append(propbackupinfos, propbackupinfo)
	return propbackupinfos
}
func flattenMapRecoveryConfigParams(p *models.RecoveryConfigParams, d *schema.ResourceData) []map[string]interface{} {

	var propconfigparamss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propconfigparams := make(map[string]interface{})
	delete(item.AO1, "ObjectType")
	if len(item.AO1) != 0 {
		j, err := json.Marshal(item.AO1)
		if err == nil {
			propconfigparams["additional_properties"] = string(j)
		} else {
			log.Printf("Error occured while flattening and json parsing: %s", err)
		}
	}
	propconfigparams["object_type"] = item.ObjectType

	propconfigparamss = append(propconfigparamss, propconfigparams)
	return propconfigparamss
}
func flattenMapRecoveryAbstractRestoreStatusRef(p *models.RecoveryAbstractRestoreStatusRef, d *schema.ResourceData) []map[string]interface{} {

	var proprestorestatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proprestorestatus := make(map[string]interface{})
	proprestorestatus["moid"] = item.Moid
	proprestorestatus["object_type"] = item.ObjectType
	proprestorestatus["selector"] = item.Selector

	proprestorestatuss = append(proprestorestatuss, proprestorestatus)
	return proprestorestatuss
}
func flattenMapComputeBladeRef(p *models.ComputeBladeRef, d *schema.ResourceData) []map[string]interface{} {

	var propcomputeblades []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcomputeblade := make(map[string]interface{})
	propcomputeblade["moid"] = item.Moid
	propcomputeblade["object_type"] = item.ObjectType
	propcomputeblade["selector"] = item.Selector

	propcomputeblades = append(propcomputeblades, propcomputeblade)
	return propcomputeblades
}
func flattenMapInventoryGenericInventoryHolderRef(p *models.InventoryGenericInventoryHolderRef, d *schema.ResourceData) []map[string]interface{} {

	var propinventorygenericinventoryholders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propinventorygenericinventoryholder := make(map[string]interface{})
	propinventorygenericinventoryholder["moid"] = item.Moid
	propinventorygenericinventoryholder["object_type"] = item.ObjectType
	propinventorygenericinventoryholder["selector"] = item.Selector

	propinventorygenericinventoryholders = append(propinventorygenericinventoryholders, propinventorygenericinventoryholder)
	return propinventorygenericinventoryholders
}
func flattenMapStoragePhysicalDiskRef(p *models.StoragePhysicalDiskRef, d *schema.ResourceData) []map[string]interface{} {

	var propstoragephysicaldisks []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstoragephysicaldisk := make(map[string]interface{})
	propstoragephysicaldisk["moid"] = item.Moid
	propstoragephysicaldisk["object_type"] = item.ObjectType
	propstoragephysicaldisk["selector"] = item.Selector

	propstoragephysicaldisks = append(propstoragephysicaldisks, propstoragephysicaldisk)
	return propstoragephysicaldisks
}
func flattenMapInventoryBaseRef(p *models.InventoryBaseRef, d *schema.ResourceData) []map[string]interface{} {

	var propcomponents []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcomponent := make(map[string]interface{})
	propcomponent["moid"] = item.Moid
	propcomponent["object_type"] = item.ObjectType
	propcomponent["selector"] = item.Selector

	propcomponents = append(propcomponents, propcomponent)
	return propcomponents
}
func flattenMapCondHclStatusRef(p *models.CondHclStatusRef, d *schema.ResourceData) []map[string]interface{} {

	var prophclstatuss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	prophclstatus := make(map[string]interface{})
	prophclstatus["moid"] = item.Moid
	prophclstatus["object_type"] = item.ObjectType
	prophclstatus["selector"] = item.Selector

	prophclstatuss = append(prophclstatuss, prophclstatus)
	return prophclstatuss
}
func flattenMapHclOperatingSystemVendorRef(p *models.HclOperatingSystemVendorRef, d *schema.ResourceData) []map[string]interface{} {

	var propvendors []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvendor := make(map[string]interface{})
	propvendor["moid"] = item.Moid
	propvendor["object_type"] = item.ObjectType
	propvendor["selector"] = item.Selector

	propvendors = append(propvendors, propvendor)
	return propvendors
}
func flattenListStoragePureHostRef(p []*models.StoragePureHostRef, d *schema.ResourceData) []map[string]interface{} {
	var prophostgroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		prophostgroups := make(map[string]interface{})
		prophostgroups["moid"] = item.Moid
		prophostgroups["object_type"] = item.ObjectType
		prophostgroups["selector"] = item.Selector
		prophostgroupss = append(prophostgroupss, prophostgroups)
	}
	return prophostgroupss
}
func flattenMapStorageGenericArrayRef(p *models.StorageGenericArrayRef, d *schema.ResourceData) []map[string]interface{} {

	var propstoragearrays []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstoragearray := make(map[string]interface{})
	propstoragearray["moid"] = item.Moid
	propstoragearray["object_type"] = item.ObjectType
	propstoragearray["selector"] = item.Selector

	propstoragearrays = append(propstoragearrays, propstoragearray)
	return propstoragearrays
}
func flattenListStoragePureVolumeRef(p []*models.StoragePureVolumeRef, d *schema.ResourceData) []map[string]interface{} {
	var propvolumess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propvolumes := make(map[string]interface{})
		propvolumes["moid"] = item.Moid
		propvolumes["object_type"] = item.ObjectType
		propvolumes["selector"] = item.Selector
		propvolumess = append(propvolumess, propvolumes)
	}
	return propvolumess
}
func flattenListFirmwareRunningFirmwareRef(p []*models.FirmwareRunningFirmwareRef, d *schema.ResourceData) []map[string]interface{} {
	var proprunningfirmwares []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proprunningfirmware := make(map[string]interface{})
		proprunningfirmware["moid"] = item.Moid
		proprunningfirmware["object_type"] = item.ObjectType
		proprunningfirmware["selector"] = item.Selector
		proprunningfirmwares = append(proprunningfirmwares, proprunningfirmware)
	}
	return proprunningfirmwares
}
func flattenMapHyperflexHxResiliencyInfoDt(p *models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {

	var propresiliencydetailss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propresiliencydetails := make(map[string]interface{})
	propresiliencydetails["data_replication_factor"] = item.DataReplicationFactor
	propresiliencydetails["hdd_failures_tolerable"] = item.HddFailuresTolerable
	propresiliencydetails["messages"] = item.Messages
	propresiliencydetails["node_failures_tolerable"] = item.NodeFailuresTolerable
	propresiliencydetails["object_type"] = item.ObjectType
	propresiliencydetails["policy_compliance"] = item.PolicyCompliance
	propresiliencydetails["resiliency_state"] = item.ResiliencyState
	propresiliencydetails["ssd_failures_tolerable"] = item.SsdFailuresTolerable

	propresiliencydetailss = append(propresiliencydetailss, propresiliencydetails)
	return propresiliencydetailss
}
func flattenListHyperflexHxZoneResiliencyInfoDt(p []*models.HyperflexHxZoneResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var propzoneresiliencylists []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propzoneresiliencylist := make(map[string]interface{})
		propzoneresiliencylist["name"] = item.Name
		propzoneresiliencylist["object_type"] = item.ObjectType
		propzoneresiliencylist["resiliency_info"] = (func(p *models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {

			var propresiliencyinfos []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propresiliencyinfo := make(map[string]interface{})
			propresiliencyinfo["data_replication_factor"] = item.DataReplicationFactor
			propresiliencyinfo["hdd_failures_tolerable"] = item.HddFailuresTolerable
			propresiliencyinfo["messages"] = item.Messages
			propresiliencyinfo["node_failures_tolerable"] = item.NodeFailuresTolerable
			propresiliencyinfo["object_type"] = item.ObjectType
			propresiliencyinfo["policy_compliance"] = item.PolicyCompliance
			propresiliencyinfo["resiliency_state"] = item.ResiliencyState
			propresiliencyinfo["ssd_failures_tolerable"] = item.SsdFailuresTolerable

			propresiliencyinfos = append(propresiliencyinfos, propresiliencyinfo)
			return propresiliencyinfos
		})(item.ResiliencyInfo, d)
		propzoneresiliencylists = append(propzoneresiliencylists, propzoneresiliencylist)
	}
	return propzoneresiliencylists
}
func flattenListNiaapiDetail(p []*models.NiaapiDetail, d *schema.ResourceData) []map[string]interface{} {
	var propcontents []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcontent := make(map[string]interface{})
		propcontent["chksum"] = item.Chksum
		propcontent["filename"] = item.Filename
		propcontent["name"] = item.Name
		propcontent["object_type"] = item.ObjectType
		propcontents = append(propcontents, propcontent)
	}
	return propcontents
}
func flattenMapAdapterUnitRef(p *models.AdapterUnitRef, d *schema.ResourceData) []map[string]interface{} {

	var propadapterunits []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propadapterunit := make(map[string]interface{})
	propadapterunit["moid"] = item.Moid
	propadapterunit["object_type"] = item.ObjectType
	propadapterunit["selector"] = item.Selector

	propadapterunits = append(propadapterunits, propadapterunit)
	return propadapterunits
}
func flattenListIamEndPointPrivilegeRef(p []*models.IamEndPointPrivilegeRef, d *schema.ResourceData) []map[string]interface{} {
	var propendpointprivilegess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propendpointprivileges := make(map[string]interface{})
		propendpointprivileges["moid"] = item.Moid
		propendpointprivileges["object_type"] = item.ObjectType
		propendpointprivileges["selector"] = item.Selector
		propendpointprivilegess = append(propendpointprivilegess, propendpointprivileges)
	}
	return propendpointprivilegess
}
func flattenMapStorageProtectionGroupRef(p *models.StorageProtectionGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var propprotectiongroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propprotectiongroup := make(map[string]interface{})
	propprotectiongroup["moid"] = item.Moid
	propprotectiongroup["object_type"] = item.ObjectType
	propprotectiongroup["selector"] = item.Selector

	propprotectiongroups = append(propprotectiongroups, propprotectiongroup)
	return propprotectiongroups
}
func flattenMapManagementControllerRef(p *models.ManagementControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var propmanagementcontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmanagementcontroller := make(map[string]interface{})
	propmanagementcontroller["moid"] = item.Moid
	propmanagementcontroller["object_type"] = item.ObjectType
	propmanagementcontroller["selector"] = item.Selector

	propmanagementcontrollers = append(propmanagementcontrollers, propmanagementcontroller)
	return propmanagementcontrollers
}
func flattenMapStorageControllerRef(p *models.StorageControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var propstoragecontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstoragecontroller := make(map[string]interface{})
	propstoragecontroller["moid"] = item.Moid
	propstoragecontroller["object_type"] = item.ObjectType
	propstoragecontroller["selector"] = item.Selector

	propstoragecontrollers = append(propstoragecontrollers, propstoragecontroller)
	return propstoragecontrollers
}
func flattenMapStorageVirtualDriveRef(p *models.StorageVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {

	var propvirtualdrives []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvirtualdrive := make(map[string]interface{})
	propvirtualdrive["moid"] = item.Moid
	propvirtualdrive["object_type"] = item.ObjectType
	propvirtualdrive["selector"] = item.Selector

	propvirtualdrives = append(propvirtualdrives, propvirtualdrive)
	return propvirtualdrives
}
func flattenMapForecastDefinitionRef(p *models.ForecastDefinitionRef, d *schema.ResourceData) []map[string]interface{} {

	var propforecastdefs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propforecastdef := make(map[string]interface{})
	propforecastdef["moid"] = item.Moid
	propforecastdef["object_type"] = item.ObjectType
	propforecastdef["selector"] = item.Selector

	propforecastdefs = append(propforecastdefs, propforecastdef)
	return propforecastdefs
}
func flattenMapForecastModel(p *models.ForecastModel, d *schema.ResourceData) []map[string]interface{} {

	var propmodels []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmodel := make(map[string]interface{})
	propmodel["accuracy"] = item.Accuracy
	propmodel["model_data"] = item.ModelData
	propmodel["model_type"] = item.ModelType
	propmodel["object_type"] = item.ObjectType

	propmodels = append(propmodels, propmodel)
	return propmodels
}
func flattenListInventoryGenericInventoryRef(p []*models.InventoryGenericInventoryRef, d *schema.ResourceData) []map[string]interface{} {
	var propgenericinventorys []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propgenericinventory := make(map[string]interface{})
		propgenericinventory["moid"] = item.Moid
		propgenericinventory["object_type"] = item.ObjectType
		propgenericinventory["selector"] = item.Selector
		propgenericinventorys = append(propgenericinventorys, propgenericinventory)
	}
	return propgenericinventorys
}
func flattenListOsConfigurationFileRef(p []*models.OsConfigurationFileRef, d *schema.ResourceData) []map[string]interface{} {
	var propconfigurationfiless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propconfigurationfiles := make(map[string]interface{})
		propconfigurationfiles["moid"] = item.Moid
		propconfigurationfiles["object_type"] = item.ObjectType
		propconfigurationfiles["selector"] = item.Selector
		propconfigurationfiless = append(propconfigurationfiless, propconfigurationfiles)
	}
	return propconfigurationfiless
}
func flattenMapNiaapiVersionRegexPlatform(p *models.NiaapiVersionRegexPlatform, d *schema.ResourceData) []map[string]interface{} {

	var propapics []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propapic := make(map[string]interface{})
	propapic["anyllregex"] = item.Anyllregex
	propapic["currentlltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {

		var propcurrentlltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propcurrentlltrain := make(map[string]interface{})
		propcurrentlltrain["object_type"] = item.ObjectType
		propcurrentlltrain["regex"] = item.Regex
		propcurrentlltrain["software_version"] = item.SoftwareVersion

		propcurrentlltrains = append(propcurrentlltrains, propcurrentlltrain)
		return propcurrentlltrains
	})(item.Currentlltrain, d)
	propapic["latestsltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {

		var proplatestsltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		proplatestsltrain := make(map[string]interface{})
		proplatestsltrain["object_type"] = item.ObjectType
		proplatestsltrain["regex"] = item.Regex
		proplatestsltrain["software_version"] = item.SoftwareVersion

		proplatestsltrains = append(proplatestsltrains, proplatestsltrain)
		return proplatestsltrains
	})(item.Latestsltrain, d)
	propapic["object_type"] = item.ObjectType
	propapic["sltrain"] = (func(p []*models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var propsltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			propsltrain := make(map[string]interface{})
			propsltrain["object_type"] = item.ObjectType
			propsltrain["regex"] = item.Regex
			propsltrain["software_version"] = item.SoftwareVersion
			propsltrains = append(propsltrains, propsltrain)
		}
		return propsltrains
	})(item.Sltrain, d)
	propapic["upcominglltrain"] = (func(p *models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {

		var propupcominglltrains []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propupcominglltrain := make(map[string]interface{})
		propupcominglltrain["object_type"] = item.ObjectType
		propupcominglltrain["regex"] = item.Regex
		propupcominglltrain["software_version"] = item.SoftwareVersion

		propupcominglltrains = append(propupcominglltrains, propupcominglltrain)
		return propupcominglltrains
	})(item.Upcominglltrain, d)

	propapics = append(propapics, propapic)
	return propapics
}
func flattenListEtherPhysicalPortRef(p []*models.EtherPhysicalPortRef, d *schema.ResourceData) []map[string]interface{} {
	var propethernetportss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propethernetports := make(map[string]interface{})
		propethernetports["moid"] = item.Moid
		propethernetports["object_type"] = item.ObjectType
		propethernetports["selector"] = item.Selector
		propethernetportss = append(propethernetportss, propethernetports)
	}
	return propethernetportss
}
func flattenMapPortGroupRef(p *models.PortGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var propportgroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propportgroup := make(map[string]interface{})
	propportgroup["moid"] = item.Moid
	propportgroup["object_type"] = item.ObjectType
	propportgroup["selector"] = item.Selector

	propportgroups = append(propportgroups, propportgroup)
	return propportgroups
}
func flattenMapEquipmentChassisRef(p *models.EquipmentChassisRef, d *schema.ResourceData) []map[string]interface{} {

	var propequipmentchassiss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propequipmentchassis := make(map[string]interface{})
	propequipmentchassis["moid"] = item.Moid
	propequipmentchassis["object_type"] = item.ObjectType
	propequipmentchassis["selector"] = item.Selector

	propequipmentchassiss = append(propequipmentchassiss, propequipmentchassis)
	return propequipmentchassiss
}
func flattenMapEquipmentSharedIoModuleRef(p *models.EquipmentSharedIoModuleRef, d *schema.ResourceData) []map[string]interface{} {

	var propsharediomodules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsharediomodule := make(map[string]interface{})
	propsharediomodule["moid"] = item.Moid
	propsharediomodule["object_type"] = item.ObjectType
	propsharediomodule["selector"] = item.Selector

	propsharediomodules = append(propsharediomodules, propsharediomodule)
	return propsharediomodules
}
func flattenMapIaasUcsdInfoRef(p *models.IaasUcsdInfoRef, d *schema.ResourceData) []map[string]interface{} {

	var propguids []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propguid := make(map[string]interface{})
	propguid["moid"] = item.Moid
	propguid["object_type"] = item.ObjectType
	propguid["selector"] = item.Selector

	propguids = append(propguids, propguid)
	return propguids
}
func flattenListIaasLicenseKeysInfo(p []*models.IaasLicenseKeysInfo, d *schema.ResourceData) []map[string]interface{} {
	var proplicensekeysinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proplicensekeysinfo := make(map[string]interface{})
		proplicensekeysinfo["count"] = item.Count
		proplicensekeysinfo["expiration_date"] = item.ExpirationDate
		proplicensekeysinfo["license_id"] = item.LicenseID
		proplicensekeysinfo["object_type"] = item.ObjectType
		proplicensekeysinfo["pid"] = item.Pid
		proplicensekeysinfos = append(proplicensekeysinfos, proplicensekeysinfo)
	}
	return proplicensekeysinfos
}
func flattenListIaasLicenseUtilizationInfo(p []*models.IaasLicenseUtilizationInfo, d *schema.ResourceData) []map[string]interface{} {
	var proplicenseutilizationinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proplicenseutilizationinfo := make(map[string]interface{})
		proplicenseutilizationinfo["actual_used"] = item.ActualUsed
		proplicenseutilizationinfo["label"] = item.Label
		proplicenseutilizationinfo["licensed_limit"] = item.LicensedLimit
		proplicenseutilizationinfo["object_type"] = item.ObjectType
		proplicenseutilizationinfo["sku"] = item.Sku
		proplicenseutilizationinfos = append(proplicenseutilizationinfos, proplicenseutilizationinfo)
	}
	return proplicenseutilizationinfos
}
func flattenListServerConfigResultEntryRef(p []*models.ServerConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var propresultentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propresultentries := make(map[string]interface{})
		propresultentries["moid"] = item.Moid
		propresultentries["object_type"] = item.ObjectType
		propresultentries["selector"] = item.Selector
		propresultentriess = append(propresultentriess, propresultentries)
	}
	return propresultentriess
}
func flattenMapStorageFlexFlashControllerRef(p *models.StorageFlexFlashControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var propstorageflexflashcontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstorageflexflashcontroller := make(map[string]interface{})
	propstorageflexflashcontroller["moid"] = item.Moid
	propstorageflexflashcontroller["object_type"] = item.ObjectType
	propstorageflexflashcontroller["selector"] = item.Selector

	propstorageflexflashcontrollers = append(propstorageflexflashcontrollers, propstorageflexflashcontroller)
	return propstorageflexflashcontrollers
}
func flattenListIamPermissionToRoles(p []*models.IamPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var proppermissionroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppermissionroles := make(map[string]interface{})
		proppermissionroles["object_type"] = item.ObjectType
		proppermissionroles["permission"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {

			var proppermissions []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			proppermission := make(map[string]interface{})
			proppermission["moid"] = item.Moid
			proppermission["object_type"] = item.ObjectType

			proppermissions = append(proppermissions, proppermission)
			return proppermissions
		})(item.Permission, d)
		proppermissionroles["roles"] = (func(p []*models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var proproless []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				proproles := make(map[string]interface{})
				proproles["moid"] = item.Moid
				proproles["object_type"] = item.ObjectType
				proproless = append(proproless, proproles)
			}
			return proproless
		})(item.Roles, d)
		proppermissionroless = append(proppermissionroless, proppermissionroles)
	}
	return proppermissionroless
}
func flattenListForecastDefinitionRef(p []*models.ForecastDefinitionRef, d *schema.ResourceData) []map[string]interface{} {
	var propdefinitions []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propdefinition := make(map[string]interface{})
		propdefinition["moid"] = item.Moid
		propdefinition["object_type"] = item.ObjectType
		propdefinition["selector"] = item.Selector
		propdefinitions = append(propdefinitions, propdefinition)
	}
	return propdefinitions
}
func flattenMapMemoryPersistentMemoryRegionRef(p *models.MemoryPersistentMemoryRegionRef, d *schema.ResourceData) []map[string]interface{} {

	var propmemorypersistentmemoryregions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmemorypersistentmemoryregion := make(map[string]interface{})
	propmemorypersistentmemoryregion["moid"] = item.Moid
	propmemorypersistentmemoryregion["object_type"] = item.ObjectType
	propmemorypersistentmemoryregion["selector"] = item.Selector

	propmemorypersistentmemoryregions = append(propmemorypersistentmemoryregions, propmemorypersistentmemoryregion)
	return propmemorypersistentmemoryregions
}
func flattenMapStorageEnclosureRef(p *models.StorageEnclosureRef, d *schema.ResourceData) []map[string]interface{} {

	var propstorageenclosures []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstorageenclosure := make(map[string]interface{})
	propstorageenclosure["moid"] = item.Moid
	propstorageenclosure["object_type"] = item.ObjectType
	propstorageenclosure["selector"] = item.Selector

	propstorageenclosures = append(propstorageenclosures, propstorageenclosure)
	return propstorageenclosures
}
func flattenMapNiatelemetryDiskinfo(p *models.NiatelemetryDiskinfo, d *schema.ResourceData) []map[string]interface{} {

	var propdisks []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdisk := make(map[string]interface{})
	propdisk["free"] = item.Free
	propdisk["name"] = item.Name
	propdisk["object_type"] = item.ObjectType
	propdisk["total"] = item.Total
	propdisk["used"] = item.Used

	propdisks = append(propdisks, propdisk)
	return propdisks
}
func flattenMapNiatelemetryNiaLicenseStateRef(p *models.NiatelemetryNiaLicenseStateRef, d *schema.ResourceData) []map[string]interface{} {

	var proplicensestates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proplicensestate := make(map[string]interface{})
	proplicensestate["moid"] = item.Moid
	proplicensestate["object_type"] = item.ObjectType
	proplicensestate["selector"] = item.Selector

	proplicensestates = append(proplicensestates, proplicensestate)
	return proplicensestates
}
func flattenListEquipmentFanModuleRef(p []*models.EquipmentFanModuleRef, d *schema.ResourceData) []map[string]interface{} {
	var propfanmoduless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propfanmodules := make(map[string]interface{})
		propfanmodules["moid"] = item.Moid
		propfanmodules["object_type"] = item.ObjectType
		propfanmodules["selector"] = item.Selector
		propfanmoduless = append(propfanmoduless, propfanmodules)
	}
	return propfanmoduless
}
func flattenListEquipmentPsuRef(p []*models.EquipmentPsuRef, d *schema.ResourceData) []map[string]interface{} {
	var proppsuss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppsus := make(map[string]interface{})
		proppsus["moid"] = item.Moid
		proppsus["object_type"] = item.ObjectType
		proppsus["selector"] = item.Selector
		proppsuss = append(proppsuss, proppsus)
	}
	return proppsuss
}
func flattenListEquipmentRackEnclosureSlotRef(p []*models.EquipmentRackEnclosureSlotRef, d *schema.ResourceData) []map[string]interface{} {
	var propslotss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propslots := make(map[string]interface{})
		propslots["moid"] = item.Moid
		propslots["object_type"] = item.ObjectType
		propslots["selector"] = item.Selector
		propslotss = append(propslotss, propslots)
	}
	return propslotss
}
func flattenMapNiaapiNewReleaseDetail(p *models.NiaapiNewReleaseDetail, d *schema.ResourceData) []map[string]interface{} {

	var proppostdetails []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proppostdetail := make(map[string]interface{})
	proppostdetail["description"] = item.Description
	proppostdetail["link"] = item.Link
	proppostdetail["object_type"] = item.ObjectType
	proppostdetail["release_note_link"] = item.ReleaseNoteLink
	proppostdetail["release_note_link_title"] = item.ReleaseNoteLinkTitle
	proppostdetail["software_download_link"] = item.SoftwareDownloadLink
	proppostdetail["software_download_link_title"] = item.SoftwareDownloadLinkTitle
	proppostdetail["title"] = item.Title
	proppostdetail["version"] = item.Version

	proppostdetails = append(proppostdetails, proppostdetail)
	return proppostdetails
}
func flattenMapIamSessionRef(p *models.IamSessionRef, d *schema.ResourceData) []map[string]interface{} {

	var propsessionss []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsessions := make(map[string]interface{})
	propsessions["moid"] = item.Moid
	propsessions["object_type"] = item.ObjectType
	propsessions["selector"] = item.Selector

	propsessionss = append(propsessionss, propsessions)
	return propsessionss
}
func flattenMapCommIPV4Interface(p *models.CommIPV4Interface, d *schema.ResourceData) []map[string]interface{} {

	var propnodeipv4configs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnodeipv4config := make(map[string]interface{})
	propnodeipv4config["gateway"] = item.Gateway
	propnodeipv4config["ip_address"] = item.IPAddress
	propnodeipv4config["netmask"] = item.Netmask
	propnodeipv4config["object_type"] = item.ObjectType

	propnodeipv4configs = append(propnodeipv4configs, propnodeipv4config)
	return propnodeipv4configs
}
func flattenMapComputeBoardRef(p *models.ComputeBoardRef, d *schema.ResourceData) []map[string]interface{} {

	var propcomputeboards []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcomputeboard := make(map[string]interface{})
	propcomputeboard["moid"] = item.Moid
	propcomputeboard["object_type"] = item.ObjectType
	propcomputeboard["selector"] = item.Selector

	propcomputeboards = append(propcomputeboards, propcomputeboard)
	return propcomputeboards
}
func flattenMapOnpremSchedule(p *models.OnpremSchedule, d *schema.ResourceData) []map[string]interface{} {

	var propschedules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propschedule := make(map[string]interface{})
	propschedule["day_of_month"] = item.DayOfMonth
	propschedule["day_of_week"] = item.DayOfWeek
	propschedule["month_of_year"] = item.MonthOfYear
	propschedule["object_type"] = item.ObjectType
	propschedule["repeat_interval"] = item.RepeatInterval
	propschedule["time_of_day"] = item.TimeOfDay
	propschedule["time_zone"] = item.TimeZone
	propschedule["week_of_month"] = item.WeekOfMonth

	propschedules = append(propschedules, propschedule)
	return propschedules
}
func flattenMapStorageCapacity(p *models.StorageCapacity, d *schema.ResourceData) []map[string]interface{} {

	var propstorageutilizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstorageutilization := make(map[string]interface{})
	propstorageutilization["available"] = item.Available
	propstorageutilization["free"] = item.Free
	propstorageutilization["object_type"] = item.ObjectType
	propstorageutilization["total"] = item.Total
	propstorageutilization["used"] = item.Used

	propstorageutilizations = append(propstorageutilizations, propstorageutilization)
	return propstorageutilizations
}
func flattenMapPolicyConfigResultContext(p *models.PolicyConfigResultContext, d *schema.ResourceData) []map[string]interface{} {

	var propcontexts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcontext := make(map[string]interface{})
	propcontext["entity_data"] = item.EntityData
	propcontext["entity_moid"] = item.EntityMoid
	propcontext["entity_name"] = item.EntityName
	propcontext["entity_type"] = item.EntityType
	propcontext["object_type"] = item.ObjectType

	propcontexts = append(propcontexts, propcontext)
	return propcontexts
}
func flattenMapEquipmentLocatorLedRef(p *models.EquipmentLocatorLedRef, d *schema.ResourceData) []map[string]interface{} {

	var proplocatorleds []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proplocatorled := make(map[string]interface{})
	proplocatorled["moid"] = item.Moid
	proplocatorled["object_type"] = item.ObjectType
	proplocatorled["selector"] = item.Selector

	proplocatorleds = append(proplocatorleds, proplocatorled)
	return proplocatorleds
}
func flattenMapComputeServerConfig(p *models.ComputeServerConfig, d *schema.ResourceData) []map[string]interface{} {

	var propserverconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propserverconfig := make(map[string]interface{})
	propserverconfig["asset_tag"] = item.AssetTag
	propserverconfig["object_type"] = item.ObjectType
	propserverconfig["user_label"] = item.UserLabel

	propserverconfigs = append(propserverconfigs, propserverconfig)
	return propserverconfigs
}
func flattenListAdapterExtEthInterfaceRef(p []*models.AdapterExtEthInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var propextethifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propextethifs := make(map[string]interface{})
		propextethifs["moid"] = item.Moid
		propextethifs["object_type"] = item.ObjectType
		propextethifs["selector"] = item.Selector
		propextethifss = append(propextethifss, propextethifs)
	}
	return propextethifss
}
func flattenListAdapterHostEthInterfaceRef(p []*models.AdapterHostEthInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var prophostethifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		prophostethifs := make(map[string]interface{})
		prophostethifs["moid"] = item.Moid
		prophostethifs["object_type"] = item.ObjectType
		prophostethifs["selector"] = item.Selector
		prophostethifss = append(prophostethifss, prophostethifs)
	}
	return prophostethifss
}
func flattenListAdapterHostFcInterfaceRef(p []*models.AdapterHostFcInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var prophostfcifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		prophostfcifs := make(map[string]interface{})
		prophostfcifs["moid"] = item.Moid
		prophostfcifs["object_type"] = item.ObjectType
		prophostfcifs["selector"] = item.Selector
		prophostfcifss = append(prophostfcifss, prophostfcifs)
	}
	return prophostfcifss
}
func flattenListAdapterHostIscsiInterfaceRef(p []*models.AdapterHostIscsiInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var prophostiscsiifss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		prophostiscsiifs := make(map[string]interface{})
		prophostiscsiifs["moid"] = item.Moid
		prophostiscsiifs["object_type"] = item.ObjectType
		prophostiscsiifs["selector"] = item.Selector
		prophostiscsiifss = append(prophostiscsiifss, prophostiscsiifs)
	}
	return prophostiscsiifss
}
func flattenMapEquipmentRackEnclosureRef(p *models.EquipmentRackEnclosureRef, d *schema.ResourceData) []map[string]interface{} {

	var propequipmentrackenclosures []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propequipmentrackenclosure := make(map[string]interface{})
	propequipmentrackenclosure["moid"] = item.Moid
	propequipmentrackenclosure["object_type"] = item.ObjectType
	propequipmentrackenclosure["selector"] = item.Selector

	propequipmentrackenclosures = append(propequipmentrackenclosures, propequipmentrackenclosure)
	return propequipmentrackenclosures
}
func flattenListEquipmentFanRef(p []*models.EquipmentFanRef, d *schema.ResourceData) []map[string]interface{} {
	var propfanss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propfans := make(map[string]interface{})
		propfans["moid"] = item.Moid
		propfans["object_type"] = item.ObjectType
		propfans["selector"] = item.Selector
		propfanss = append(propfanss, propfans)
	}
	return propfanss
}
func flattenMapNetworkElementRef(p *models.NetworkElementRef, d *schema.ResourceData) []map[string]interface{} {

	var propnetworkelements []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnetworkelement := make(map[string]interface{})
	propnetworkelement["moid"] = item.Moid
	propnetworkelement["object_type"] = item.ObjectType
	propnetworkelement["selector"] = item.Selector

	propnetworkelements = append(propnetworkelements, propnetworkelement)
	return propnetworkelements
}
func flattenMapStorageHostRef(p *models.StorageHostRef, d *schema.ResourceData) []map[string]interface{} {

	var prophosts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	prophost := make(map[string]interface{})
	prophost["moid"] = item.Moid
	prophost["object_type"] = item.ObjectType
	prophost["selector"] = item.Selector

	prophosts = append(prophosts, prophost)
	return prophosts
}
func flattenMapStorageVolumeRef(p *models.StorageVolumeRef, d *schema.ResourceData) []map[string]interface{} {

	var propvolumes []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvolume := make(map[string]interface{})
	propvolume["moid"] = item.Moid
	propvolume["object_type"] = item.ObjectType
	propvolume["selector"] = item.Selector

	propvolumes = append(propvolumes, propvolume)
	return propvolumes
}
func flattenMapAssetContractInformation(p *models.AssetContractInformation, d *schema.ResourceData) []map[string]interface{} {

	var propcontracts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcontract := make(map[string]interface{})
	propcontract["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var propbilltos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propbillto := make(map[string]interface{})
		propbillto["address1"] = item.Address1
		propbillto["address2"] = item.Address2
		propbillto["city"] = item.City
		propbillto["country"] = item.Country
		propbillto["location"] = item.Location
		propbillto["name"] = item.Name
		propbillto["object_type"] = item.ObjectType
		propbillto["postal_code"] = item.PostalCode
		propbillto["state"] = item.State

		propbilltos = append(propbilltos, propbillto)
		return propbilltos
	})(item.BillTo, d)
	propcontract["bill_to_global_ultimate"] = (func(p *models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {

		var propbilltoglobalultimates []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propbilltoglobalultimate := make(map[string]interface{})
		propbilltoglobalultimate["id"] = item.ID
		propbilltoglobalultimate["name"] = item.Name
		propbilltoglobalultimate["object_type"] = item.ObjectType

		propbilltoglobalultimates = append(propbilltoglobalultimates, propbilltoglobalultimate)
		return propbilltoglobalultimates
	})(item.BillToGlobalUltimate, d)
	propcontract["contract_number"] = item.ContractNumber
	propcontract["line_status"] = item.LineStatus
	propcontract["object_type"] = item.ObjectType

	propcontracts = append(propcontracts, propcontract)
	return propcontracts
}
func flattenMapAssetCustomerInformation(p *models.AssetCustomerInformation, d *schema.ResourceData) []map[string]interface{} {

	var propendcustomers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propendcustomer := make(map[string]interface{})
	propendcustomer["address"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var propaddresss []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propaddress := make(map[string]interface{})
		propaddress["address1"] = item.Address1
		propaddress["address2"] = item.Address2
		propaddress["city"] = item.City
		propaddress["country"] = item.Country
		propaddress["location"] = item.Location
		propaddress["name"] = item.Name
		propaddress["object_type"] = item.ObjectType
		propaddress["postal_code"] = item.PostalCode
		propaddress["state"] = item.State

		propaddresss = append(propaddresss, propaddress)
		return propaddresss
	})(item.Address, d)
	propendcustomer["id"] = item.ID
	propendcustomer["name"] = item.Name
	propendcustomer["object_type"] = item.ObjectType

	propendcustomers = append(propendcustomers, propendcustomer)
	return propendcustomers
}
func flattenMapAssetGlobalUltimate(p *models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {

	var propenduserglobalultimates []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propenduserglobalultimate := make(map[string]interface{})
	propenduserglobalultimate["id"] = item.ID
	propenduserglobalultimate["name"] = item.Name
	propenduserglobalultimate["object_type"] = item.ObjectType

	propenduserglobalultimates = append(propenduserglobalultimates, propenduserglobalultimate)
	return propenduserglobalultimates
}
func flattenMapAssetProductInformation(p *models.AssetProductInformation, d *schema.ResourceData) []map[string]interface{} {

	var propproducts []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propproduct := make(map[string]interface{})
	propproduct["bill_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var propbilltos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propbillto := make(map[string]interface{})
		propbillto["address1"] = item.Address1
		propbillto["address2"] = item.Address2
		propbillto["city"] = item.City
		propbillto["country"] = item.Country
		propbillto["location"] = item.Location
		propbillto["name"] = item.Name
		propbillto["object_type"] = item.ObjectType
		propbillto["postal_code"] = item.PostalCode
		propbillto["state"] = item.State

		propbilltos = append(propbilltos, propbillto)
		return propbilltos
	})(item.BillTo, d)
	propproduct["description"] = item.Description
	propproduct["family"] = item.Family
	propproduct["group"] = item.Group
	propproduct["number"] = item.Number
	propproduct["object_type"] = item.ObjectType
	propproduct["ship_to"] = (func(p *models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {

		var propshiptos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propshipto := make(map[string]interface{})
		propshipto["address1"] = item.Address1
		propshipto["address2"] = item.Address2
		propshipto["city"] = item.City
		propshipto["country"] = item.Country
		propshipto["location"] = item.Location
		propshipto["name"] = item.Name
		propshipto["object_type"] = item.ObjectType
		propshipto["postal_code"] = item.PostalCode
		propshipto["state"] = item.State

		propshiptos = append(propshiptos, propshipto)
		return propshiptos
	})(item.ShipTo, d)
	propproduct["sub_type"] = item.SubType

	propproducts = append(propproducts, propproduct)
	return propproducts
}
func flattenMapFirmwareUpgradeRef(p *models.FirmwareUpgradeRef, d *schema.ResourceData) []map[string]interface{} {

	var propupgrades []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propupgrade := make(map[string]interface{})
	propupgrade["moid"] = item.Moid
	propupgrade["object_type"] = item.ObjectType
	propupgrade["selector"] = item.Selector

	propupgrades = append(propupgrades, propupgrade)
	return propupgrades
}
func flattenMapMemoryArrayRef(p *models.MemoryArrayRef, d *schema.ResourceData) []map[string]interface{} {

	var propmemoryarrays []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmemoryarray := make(map[string]interface{})
	propmemoryarray["moid"] = item.Moid
	propmemoryarray["object_type"] = item.ObjectType
	propmemoryarray["selector"] = item.Selector

	propmemoryarrays = append(propmemoryarrays, propmemoryarray)
	return propmemoryarrays
}
func flattenListComputeBladeRef(p []*models.ComputeBladeRef, d *schema.ResourceData) []map[string]interface{} {
	var propcomputebladess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcomputeblades := make(map[string]interface{})
		propcomputeblades["moid"] = item.Moid
		propcomputeblades["object_type"] = item.ObjectType
		propcomputeblades["selector"] = item.Selector
		propcomputebladess = append(propcomputebladess, propcomputeblades)
	}
	return propcomputebladess
}
func flattenListComputeRackUnitRef(p []*models.ComputeRackUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var propcomputerackunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcomputerackunits := make(map[string]interface{})
		propcomputerackunits["moid"] = item.Moid
		propcomputerackunits["object_type"] = item.ObjectType
		propcomputerackunits["selector"] = item.Selector
		propcomputerackunitss = append(propcomputerackunitss, propcomputerackunits)
	}
	return propcomputerackunitss
}
func flattenListNetworkElementRef(p []*models.NetworkElementRef, d *schema.ResourceData) []map[string]interface{} {
	var propnetworkelementss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propnetworkelements := make(map[string]interface{})
		propnetworkelements["moid"] = item.Moid
		propnetworkelements["object_type"] = item.ObjectType
		propnetworkelements["selector"] = item.Selector
		propnetworkelementss = append(propnetworkelementss, propnetworkelements)
	}
	return propnetworkelementss
}
func flattenListStoragePhysicalDiskExtensionRef(p []*models.StoragePhysicalDiskExtensionRef, d *schema.ResourceData) []map[string]interface{} {
	var propphysicaldiskextensionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propphysicaldiskextensions := make(map[string]interface{})
		propphysicaldiskextensions["moid"] = item.Moid
		propphysicaldiskextensions["object_type"] = item.ObjectType
		propphysicaldiskextensions["selector"] = item.Selector
		propphysicaldiskextensionss = append(propphysicaldiskextensionss, propphysicaldiskextensions)
	}
	return propphysicaldiskextensionss
}
func flattenListStoragePhysicalDiskRef(p []*models.StoragePhysicalDiskRef, d *schema.ResourceData) []map[string]interface{} {
	var propphysicaldiskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propphysicaldisks := make(map[string]interface{})
		propphysicaldisks["moid"] = item.Moid
		propphysicaldisks["object_type"] = item.ObjectType
		propphysicaldisks["selector"] = item.Selector
		propphysicaldiskss = append(propphysicaldiskss, propphysicaldisks)
	}
	return propphysicaldiskss
}
func flattenListStorageVirtualDriveExtensionRef(p []*models.StorageVirtualDriveExtensionRef, d *schema.ResourceData) []map[string]interface{} {
	var propvirtualdriveextensionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propvirtualdriveextensions := make(map[string]interface{})
		propvirtualdriveextensions["moid"] = item.Moid
		propvirtualdriveextensions["object_type"] = item.ObjectType
		propvirtualdriveextensions["selector"] = item.Selector
		propvirtualdriveextensionss = append(propvirtualdriveextensionss, propvirtualdriveextensions)
	}
	return propvirtualdriveextensionss
}
func flattenListStorageVirtualDriveRef(p []*models.StorageVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var propvirtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propvirtualdrives := make(map[string]interface{})
		propvirtualdrives["moid"] = item.Moid
		propvirtualdrives["object_type"] = item.ObjectType
		propvirtualdrives["selector"] = item.Selector
		propvirtualdrivess = append(propvirtualdrivess, propvirtualdrives)
	}
	return propvirtualdrivess
}
func flattenMapEquipmentFanModuleRef(p *models.EquipmentFanModuleRef, d *schema.ResourceData) []map[string]interface{} {

	var propequipmentfanmodules []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propequipmentfanmodule := make(map[string]interface{})
	propequipmentfanmodule["moid"] = item.Moid
	propequipmentfanmodule["object_type"] = item.ObjectType
	propequipmentfanmodule["selector"] = item.Selector

	propequipmentfanmodules = append(propequipmentfanmodules, propequipmentfanmodule)
	return propequipmentfanmodules
}
func flattenMapStorageArrayControllerRef(p *models.StorageArrayControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var propcontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcontroller := make(map[string]interface{})
	propcontroller["moid"] = item.Moid
	propcontroller["object_type"] = item.ObjectType
	propcontroller["selector"] = item.Selector

	propcontrollers = append(propcontrollers, propcontroller)
	return propcontrollers
}
func flattenMapStoragePureHostRef(p *models.StoragePureHostRef, d *schema.ResourceData) []map[string]interface{} {

	var prophostgroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	prophostgroup := make(map[string]interface{})
	prophostgroup["moid"] = item.Moid
	prophostgroup["object_type"] = item.ObjectType
	prophostgroup["selector"] = item.Selector

	prophostgroups = append(prophostgroups, prophostgroup)
	return prophostgroups
}
func flattenListStorageInitiator(p []*models.StorageInitiator, d *schema.ResourceData) []map[string]interface{} {
	var propinitiatorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propinitiators := make(map[string]interface{})
		propinitiators["iqn"] = item.Iqn
		propinitiators["name"] = item.Name
		propinitiators["object_type"] = item.ObjectType
		propinitiators["type"] = item.Type
		propinitiators["wwn"] = item.Wwn
		propinitiatorss = append(propinitiatorss, propinitiators)
	}
	return propinitiatorss
}
func flattenMapStorageHostUtilization(p *models.StorageHostUtilization, d *schema.ResourceData) []map[string]interface{} {

	var propstorageutilizations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstorageutilization := make(map[string]interface{})
	propstorageutilization["available"] = item.Available
	propstorageutilization["data_reduction"] = item.DataReduction
	propstorageutilization["free"] = item.Free
	propstorageutilization["object_type"] = item.ObjectType
	propstorageutilization["snapshot"] = item.Snapshot
	propstorageutilization["thin_provisioned"] = item.ThinProvisioned
	propstorageutilization["total"] = item.Total
	propstorageutilization["total_reduction"] = item.TotalReduction
	propstorageutilization["used"] = item.Used
	propstorageutilization["volume"] = item.Volume

	propstorageutilizations = append(propstorageutilizations, propstorageutilization)
	return propstorageutilizations
}
func flattenMapRecoveryBackupProfileRef(p *models.RecoveryBackupProfileRef, d *schema.ResourceData) []map[string]interface{} {

	var propbackupprofiles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propbackupprofile := make(map[string]interface{})
	propbackupprofile["moid"] = item.Moid
	propbackupprofile["object_type"] = item.ObjectType
	propbackupprofile["selector"] = item.Selector

	propbackupprofiles = append(propbackupprofiles, propbackupprofile)
	return propbackupprofiles
}
func flattenMapRecoveryOnDemandBackupRef(p *models.RecoveryOnDemandBackupRef, d *schema.ResourceData) []map[string]interface{} {

	var propnr0ondemandbackups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propnr0ondemandbackup := make(map[string]interface{})
	propnr0ondemandbackup["moid"] = item.Moid
	propnr0ondemandbackup["object_type"] = item.ObjectType
	propnr0ondemandbackup["selector"] = item.Selector

	propnr0ondemandbackups = append(propnr0ondemandbackups, propnr0ondemandbackup)
	return propnr0ondemandbackups
}
func flattenListRecoveryConfigResultEntryRef(p []*models.RecoveryConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var propresultentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propresultentries := make(map[string]interface{})
		propresultentries["moid"] = item.Moid
		propresultentries["object_type"] = item.ObjectType
		propresultentries["selector"] = item.Selector
		propresultentriess = append(propresultentriess, propresultentries)
	}
	return propresultentriess
}
func flattenListNiaapiRevisionInfo(p []*models.NiaapiRevisionInfo, d *schema.ResourceData) []map[string]interface{} {
	var proprevisioninfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proprevisioninfo := make(map[string]interface{})
		proprevisioninfo["date_published"] = item.DatePublished
		proprevisioninfo["object_type"] = item.ObjectType
		proprevisioninfo["revision_comment"] = item.RevisionComment
		proprevisioninfo["revision_no"] = item.RevisionNo
		proprevisioninfos = append(proprevisioninfos, proprevisioninfo)
	}
	return proprevisioninfos
}
func flattenListEquipmentIoCardRef(p []*models.EquipmentIoCardRef, d *schema.ResourceData) []map[string]interface{} {
	var propiomss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propioms := make(map[string]interface{})
		propioms["moid"] = item.Moid
		propioms["object_type"] = item.ObjectType
		propioms["selector"] = item.Selector
		propiomss = append(propiomss, propioms)
	}
	return propiomss
}
func flattenListStorageSasExpanderRef(p []*models.StorageSasExpanderRef, d *schema.ResourceData) []map[string]interface{} {
	var propsasexpanderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsasexpanders := make(map[string]interface{})
		propsasexpanders["moid"] = item.Moid
		propsasexpanders["object_type"] = item.ObjectType
		propsasexpanders["selector"] = item.Selector
		propsasexpanderss = append(propsasexpanderss, propsasexpanders)
	}
	return propsasexpanderss
}
func flattenListEquipmentSystemIoControllerRef(p []*models.EquipmentSystemIoControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var propsiocss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsiocs := make(map[string]interface{})
		propsiocs["moid"] = item.Moid
		propsiocs["object_type"] = item.ObjectType
		propsiocs["selector"] = item.Selector
		propsiocss = append(propsiocss, propsiocs)
	}
	return propsiocss
}
func flattenListStorageEnclosureRef(p []*models.StorageEnclosureRef, d *schema.ResourceData) []map[string]interface{} {
	var propstorageenclosuress []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propstorageenclosures := make(map[string]interface{})
		propstorageenclosures["moid"] = item.Moid
		propstorageenclosures["object_type"] = item.ObjectType
		propstorageenclosures["selector"] = item.Selector
		propstorageenclosuress = append(propstorageenclosuress, propstorageenclosures)
	}
	return propstorageenclosuress
}
func flattenListPortGroupRef(p []*models.PortGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var propportgroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propportgroups := make(map[string]interface{})
		propportgroups["moid"] = item.Moid
		propportgroups["object_type"] = item.ObjectType
		propportgroups["selector"] = item.Selector
		propportgroupss = append(propportgroupss, propportgroups)
	}
	return propportgroupss
}
func flattenMapMemoryPersistentMemoryConfigResultRef(p *models.MemoryPersistentMemoryConfigResultRef, d *schema.ResourceData) []map[string]interface{} {

	var propmemorypersistentmemoryconfigresults []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmemorypersistentmemoryconfigresult := make(map[string]interface{})
	propmemorypersistentmemoryconfigresult["moid"] = item.Moid
	propmemorypersistentmemoryconfigresult["object_type"] = item.ObjectType
	propmemorypersistentmemoryconfigresult["selector"] = item.Selector

	propmemorypersistentmemoryconfigresults = append(propmemorypersistentmemoryconfigresults, propmemorypersistentmemoryconfigresult)
	return propmemorypersistentmemoryconfigresults
}
func flattenMapEquipmentSystemIoControllerRef(p *models.EquipmentSystemIoControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var propequipmentsystemiocontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propequipmentsystemiocontroller := make(map[string]interface{})
	propequipmentsystemiocontroller["moid"] = item.Moid
	propequipmentsystemiocontroller["object_type"] = item.ObjectType
	propequipmentsystemiocontroller["selector"] = item.Selector

	propequipmentsystemiocontrollers = append(propequipmentsystemiocontrollers, propequipmentsystemiocontroller)
	return propequipmentsystemiocontrollers
}
func flattenListComputeIPAddress(p []*models.ComputeIPAddress, d *schema.ResourceData) []map[string]interface{} {
	var propkvmipaddressess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propkvmipaddresses := make(map[string]interface{})
		propkvmipaddresses["address"] = item.Address
		propkvmipaddresses["category"] = item.Category
		propkvmipaddresses["default_gateway"] = item.DefaultGateway
		propkvmipaddresses["dn"] = item.Dn
		propkvmipaddresses["http_port"] = item.HTTPPort
		propkvmipaddresses["https_port"] = item.HTTPSPort
		propkvmipaddresses["kvm_port"] = item.KvmPort
		propkvmipaddresses["name"] = item.Name
		propkvmipaddresses["object_type"] = item.ObjectType
		propkvmipaddresses["subnet"] = item.Subnet
		propkvmipaddresses["type"] = item.Type
		propkvmipaddressess = append(propkvmipaddressess, propkvmipaddresses)
	}
	return propkvmipaddressess
}
func flattenListCondHclStatusDetailRef(p []*models.CondHclStatusDetailRef, d *schema.ResourceData) []map[string]interface{} {
	var propdetailss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propdetails := make(map[string]interface{})
		propdetails["moid"] = item.Moid
		propdetails["object_type"] = item.ObjectType
		propdetails["selector"] = item.Selector
		propdetailss = append(propdetailss, propdetails)
	}
	return propdetailss
}
func flattenMapMemoryPersistentMemoryConfigurationRef(p *models.MemoryPersistentMemoryConfigurationRef, d *schema.ResourceData) []map[string]interface{} {

	var propmemorypersistentmemoryconfigurations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmemorypersistentmemoryconfiguration := make(map[string]interface{})
	propmemorypersistentmemoryconfiguration["moid"] = item.Moid
	propmemorypersistentmemoryconfiguration["object_type"] = item.ObjectType
	propmemorypersistentmemoryconfiguration["selector"] = item.Selector

	propmemorypersistentmemoryconfigurations = append(propmemorypersistentmemoryconfigurations, propmemorypersistentmemoryconfiguration)
	return propmemorypersistentmemoryconfigurations
}
func flattenListMemoryPersistentMemoryNamespaceRef(p []*models.MemoryPersistentMemoryNamespaceRef, d *schema.ResourceData) []map[string]interface{} {
	var proppersistentmemorynamespacess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppersistentmemorynamespaces := make(map[string]interface{})
		proppersistentmemorynamespaces["moid"] = item.Moid
		proppersistentmemorynamespaces["object_type"] = item.ObjectType
		proppersistentmemorynamespaces["selector"] = item.Selector
		proppersistentmemorynamespacess = append(proppersistentmemorynamespacess, proppersistentmemorynamespaces)
	}
	return proppersistentmemorynamespacess
}
func flattenListHyperflexAlarmRef(p []*models.HyperflexAlarmRef, d *schema.ResourceData) []map[string]interface{} {
	var propalarms []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propalarm := make(map[string]interface{})
		propalarm["moid"] = item.Moid
		propalarm["object_type"] = item.ObjectType
		propalarm["selector"] = item.Selector
		propalarms = append(propalarms, propalarm)
	}
	return propalarms
}
func flattenMapHyperflexHealthRef(p *models.HyperflexHealthRef, d *schema.ResourceData) []map[string]interface{} {

	var prophealths []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	prophealth := make(map[string]interface{})
	prophealth["moid"] = item.Moid
	prophealth["object_type"] = item.ObjectType
	prophealth["selector"] = item.Selector

	prophealths = append(prophealths, prophealth)
	return prophealths
}
func flattenListHyperflexNodeRef(p []*models.HyperflexNodeRef, d *schema.ResourceData) []map[string]interface{} {
	var propnodess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propnodes := make(map[string]interface{})
		propnodes["moid"] = item.Moid
		propnodes["object_type"] = item.ObjectType
		propnodes["selector"] = item.Selector
		propnodess = append(propnodess, propnodes)
	}
	return propnodess
}
func flattenMapHyperflexSummary(p *models.HyperflexSummary, d *schema.ResourceData) []map[string]interface{} {

	var propsummarys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsummary := make(map[string]interface{})
	propsummary["active_nodes"] = item.ActiveNodes
	propsummary["address"] = item.Address
	propsummary["boottime"] = item.Boottime
	propsummary["cluster_access_policy"] = item.ClusterAccessPolicy
	propsummary["compression_savings"] = item.CompressionSavings
	propsummary["data_replication_compliance"] = item.DataReplicationCompliance
	propsummary["data_replication_factor"] = item.DataReplicationFactor
	propsummary["deduplication_savings"] = item.DeduplicationSavings
	propsummary["downtime"] = item.Downtime
	propsummary["free_capacity"] = item.FreeCapacity
	propsummary["healing_info"] = (func(p *models.HyperflexStPlatformClusterHealingInfo, d *schema.ResourceData) []map[string]interface{} {

		var prophealinginfos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		prophealinginfo := make(map[string]interface{})
		prophealinginfo["estimated_completion_time_in_seconds"] = item.EstimatedCompletionTimeInSeconds
		prophealinginfo["in_progress"] = item.InProgress
		prophealinginfo["messages"] = item.Messages
		prophealinginfo["messages_iterator"] = item.MessagesIterator
		prophealinginfo["messages_size"] = item.MessagesSize
		prophealinginfo["object_type"] = item.ObjectType
		prophealinginfo["percent_complete"] = item.PercentComplete

		prophealinginfos = append(prophealinginfos, prophealinginfo)
		return prophealinginfos
	})(item.HealingInfo, d)
	propsummary["name"] = item.Name
	propsummary["object_type"] = item.ObjectType
	propsummary["resiliency_details"] = item.ResiliencyDetails
	propsummary["resiliency_details_size"] = item.ResiliencyDetailsSize
	propsummary["resiliency_info"] = (func(p *models.HyperflexStPlatformClusterResiliencyInfo, d *schema.ResourceData) []map[string]interface{} {

		var propresiliencyinfos []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propresiliencyinfo := make(map[string]interface{})
		propresiliencyinfo["hdd_failures_tolerable"] = item.HddFailuresTolerable
		propresiliencyinfo["messages"] = item.Messages
		propresiliencyinfo["messages_iterator"] = item.MessagesIterator
		propresiliencyinfo["messages_size"] = item.MessagesSize
		propresiliencyinfo["node_failures_tolerable"] = item.NodeFailuresTolerable
		propresiliencyinfo["object_type"] = item.ObjectType
		propresiliencyinfo["ssd_failures_tolerable"] = item.SsdFailuresTolerable
		propresiliencyinfo["state"] = item.State

		propresiliencyinfos = append(propresiliencyinfos, propresiliencyinfo)
		return propresiliencyinfos
	})(item.ResiliencyInfo, d)
	propsummary["space_status"] = item.SpaceStatus
	propsummary["state"] = item.State
	propsummary["total_capacity"] = item.TotalCapacity
	propsummary["total_savings"] = item.TotalSavings
	propsummary["uptime"] = item.Uptime
	propsummary["used_capacity"] = item.UsedCapacity

	propsummarys = append(propsummarys, propsummary)
	return propsummarys
}
func flattenMapAssetClusterMemberRef(p *models.AssetClusterMemberRef, d *schema.ResourceData) []map[string]interface{} {

	var propclustermembers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propclustermember := make(map[string]interface{})
	propclustermember["moid"] = item.Moid
	propclustermember["object_type"] = item.ObjectType
	propclustermember["selector"] = item.Selector

	propclustermembers = append(propclustermembers, propclustermember)
	return propclustermembers
}
func flattenMapHyperflexHxNetworkAddressDt(p *models.HyperflexHxNetworkAddressDt, d *schema.ResourceData) []map[string]interface{} {

	var propips []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propip := make(map[string]interface{})
	propip["address"] = item.Address
	propip["fqdn"] = item.Fqdn
	propip["ip"] = item.IP
	propip["object_type"] = item.ObjectType

	propips = append(propips, propip)
	return propips
}
func flattenMapHyperflexHxUuIDDt(p *models.HyperflexHxUuIDDt, d *schema.ResourceData) []map[string]interface{} {

	var propidentitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propidentity := make(map[string]interface{})
	propidentity["links"] = (func(p []*models.HyperflexHxLinkDt, d *schema.ResourceData) []map[string]interface{} {
		var proplinkss []map[string]interface{}
		if p == nil {
			return nil
		}
		for _, item := range p {
			item := *item
			proplinks := make(map[string]interface{})
			proplinks["comments"] = item.Comments
			proplinks["href"] = item.Href
			proplinks["method"] = item.Method
			proplinks["object_type"] = item.ObjectType
			proplinks["rel"] = item.Rel
			proplinkss = append(proplinkss, proplinks)
		}
		return proplinkss
	})(item.Links, d)
	propidentity["object_type"] = item.ObjectType
	propidentity["uuid"] = item.UUID

	propidentitys = append(propidentitys, propidentity)
	return propidentitys
}
func flattenListHyperflexConfigResultEntryRef(p []*models.HyperflexConfigResultEntryRef, d *schema.ResourceData) []map[string]interface{} {
	var propresultentriess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propresultentries := make(map[string]interface{})
		propresultentries["moid"] = item.Moid
		propresultentries["object_type"] = item.ObjectType
		propresultentries["selector"] = item.Selector
		propresultentriess = append(propresultentriess, propresultentries)
	}
	return propresultentriess
}
func flattenListStorageEnclosureDiskSlotEpRef(p []*models.StorageEnclosureDiskSlotEpRef, d *schema.ResourceData) []map[string]interface{} {
	var propenclosurediskslotss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propenclosurediskslots := make(map[string]interface{})
		propenclosurediskslots["moid"] = item.Moid
		propenclosurediskslots["object_type"] = item.ObjectType
		propenclosurediskslots["selector"] = item.Selector
		propenclosurediskslotss = append(propenclosurediskslotss, propenclosurediskslots)
	}
	return propenclosurediskslotss
}
func flattenListStorageEnclosureDiskRef(p []*models.StorageEnclosureDiskRef, d *schema.ResourceData) []map[string]interface{} {
	var propenclosurediskss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propenclosuredisks := make(map[string]interface{})
		propenclosuredisks["moid"] = item.Moid
		propenclosuredisks["object_type"] = item.ObjectType
		propenclosuredisks["selector"] = item.Selector
		propenclosurediskss = append(propenclosurediskss, propenclosuredisks)
	}
	return propenclosurediskss
}
func flattenListWorkflowDynamicWorkflowActionTaskList(p []*models.WorkflowDynamicWorkflowActionTaskList, d *schema.ResourceData) []map[string]interface{} {
	var propworkflowactiontasklistss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propworkflowactiontasklists := make(map[string]interface{})
		propworkflowactiontasklists["action"] = item.Action
		propworkflowactiontasklists["object_type"] = item.ObjectType
		propworkflowactiontasklists["tasks"] = item.Tasks
		propworkflowactiontasklistss = append(propworkflowactiontasklistss, propworkflowactiontasklists)
	}
	return propworkflowactiontasklistss
}
func flattenListStorageFlexFlashControllerPropsRef(p []*models.StorageFlexFlashControllerPropsRef, d *schema.ResourceData) []map[string]interface{} {
	var propflexflashcontrollerpropss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propflexflashcontrollerprops := make(map[string]interface{})
		propflexflashcontrollerprops["moid"] = item.Moid
		propflexflashcontrollerprops["object_type"] = item.ObjectType
		propflexflashcontrollerprops["selector"] = item.Selector
		propflexflashcontrollerpropss = append(propflexflashcontrollerpropss, propflexflashcontrollerprops)
	}
	return propflexflashcontrollerpropss
}
func flattenListStorageFlexFlashPhysicalDriveRef(p []*models.StorageFlexFlashPhysicalDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var propflexflashphysicaldrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propflexflashphysicaldrives := make(map[string]interface{})
		propflexflashphysicaldrives["moid"] = item.Moid
		propflexflashphysicaldrives["object_type"] = item.ObjectType
		propflexflashphysicaldrives["selector"] = item.Selector
		propflexflashphysicaldrivess = append(propflexflashphysicaldrivess, propflexflashphysicaldrives)
	}
	return propflexflashphysicaldrivess
}
func flattenListStorageFlexFlashVirtualDriveRef(p []*models.StorageFlexFlashVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var propflexflashvirtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propflexflashvirtualdrives := make(map[string]interface{})
		propflexflashvirtualdrives["moid"] = item.Moid
		propflexflashvirtualdrives["object_type"] = item.ObjectType
		propflexflashvirtualdrives["selector"] = item.Selector
		propflexflashvirtualdrivess = append(propflexflashvirtualdrivess, propflexflashvirtualdrives)
	}
	return propflexflashvirtualdrivess
}
func flattenMapGraphicsCardRef(p *models.GraphicsCardRef, d *schema.ResourceData) []map[string]interface{} {

	var propgraphicscards []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propgraphicscard := make(map[string]interface{})
	propgraphicscard["moid"] = item.Moid
	propgraphicscard["object_type"] = item.ObjectType
	propgraphicscard["selector"] = item.Selector

	propgraphicscards = append(propgraphicscards, propgraphicscard)
	return propgraphicscards
}
func flattenMapPciSwitchRef(p *models.PciSwitchRef, d *schema.ResourceData) []map[string]interface{} {

	var proppciswitchs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proppciswitch := make(map[string]interface{})
	proppciswitch["moid"] = item.Moid
	proppciswitch["object_type"] = item.ObjectType
	proppciswitch["selector"] = item.Selector

	proppciswitchs = append(proppciswitchs, proppciswitch)
	return proppciswitchs
}
func flattenListIamGroupPermissionToRoles(p []*models.IamGroupPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var propgrouppermissionroless []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propgrouppermissionroles := make(map[string]interface{})
		propgrouppermissionroles["group"] = (func(p *models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {

			var propgroups []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propgroup := make(map[string]interface{})
			propgroup["moid"] = item.Moid
			propgroup["object_type"] = item.ObjectType

			propgroups = append(propgroups, propgroup)
			return propgroups
		})(item.Group, d)
		propgrouppermissionroles["object_type"] = item.ObjectType
		propgrouppermissionroles["orgs"] = (func(p []*models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var proporgss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				proporgs := make(map[string]interface{})
				proporgs["moid"] = item.Moid
				proporgs["object_type"] = item.ObjectType
				proporgss = append(proporgss, proporgs)
			}
			return proporgss
		})(item.Orgs, d)
		propgrouppermissionroless = append(propgrouppermissionroless, propgrouppermissionroles)
	}
	return propgrouppermissionroless
}
func flattenMapResourceMembershipHolderRef(p *models.ResourceMembershipHolderRef, d *schema.ResourceData) []map[string]interface{} {

	var propholders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propholder := make(map[string]interface{})
	propholder["moid"] = item.Moid
	propholder["object_type"] = item.ObjectType
	propholder["selector"] = item.Selector

	propholders = append(propholders, propholder)
	return propholders
}
func flattenMapLicenseAccountLicenseDataRef(p *models.LicenseAccountLicenseDataRef, d *schema.ResourceData) []map[string]interface{} {

	var propaccountlicensedatas []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propaccountlicensedata := make(map[string]interface{})
	propaccountlicensedata["moid"] = item.Moid
	propaccountlicensedata["object_type"] = item.ObjectType
	propaccountlicensedata["selector"] = item.Selector

	propaccountlicensedatas = append(propaccountlicensedatas, propaccountlicensedata)
	return propaccountlicensedatas
}
func flattenMapApplianceDataExportPolicyRef(p *models.ApplianceDataExportPolicyRef, d *schema.ResourceData) []map[string]interface{} {

	var propparentconfigs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propparentconfig := make(map[string]interface{})
	propparentconfig["moid"] = item.Moid
	propparentconfig["object_type"] = item.ObjectType
	propparentconfig["selector"] = item.Selector

	propparentconfigs = append(propparentconfigs, propparentconfig)
	return propparentconfigs
}
func flattenListApplianceDataExportPolicyRef(p []*models.ApplianceDataExportPolicyRef, d *schema.ResourceData) []map[string]interface{} {
	var propsubconfigss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsubconfigs := make(map[string]interface{})
		propsubconfigs["moid"] = item.Moid
		propsubconfigs["object_type"] = item.ObjectType
		propsubconfigs["selector"] = item.Selector
		propsubconfigss = append(propsubconfigss, propsubconfigs)
	}
	return propsubconfigss
}
func flattenMapIamServiceProviderRef(p *models.IamServiceProviderRef, d *schema.ResourceData) []map[string]interface{} {

	var propserviceproviders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propserviceprovider := make(map[string]interface{})
	propserviceprovider["moid"] = item.Moid
	propserviceprovider["object_type"] = item.ObjectType
	propserviceprovider["selector"] = item.Selector

	propserviceproviders = append(propserviceproviders, propserviceprovider)
	return propserviceproviders
}
func flattenListStorageFlexUtilPhysicalDriveRef(p []*models.StorageFlexUtilPhysicalDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var propflexutilphysicaldrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propflexutilphysicaldrives := make(map[string]interface{})
		propflexutilphysicaldrives["moid"] = item.Moid
		propflexutilphysicaldrives["object_type"] = item.ObjectType
		propflexutilphysicaldrives["selector"] = item.Selector
		propflexutilphysicaldrivess = append(propflexutilphysicaldrivess, propflexutilphysicaldrives)
	}
	return propflexutilphysicaldrivess
}
func flattenListStorageFlexUtilVirtualDriveRef(p []*models.StorageFlexUtilVirtualDriveRef, d *schema.ResourceData) []map[string]interface{} {
	var propflexutilvirtualdrivess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propflexutilvirtualdrives := make(map[string]interface{})
		propflexutilvirtualdrives["moid"] = item.Moid
		propflexutilvirtualdrives["object_type"] = item.ObjectType
		propflexutilvirtualdrives["selector"] = item.Selector
		propflexutilvirtualdrivess = append(propflexutilvirtualdrivess, propflexutilvirtualdrives)
	}
	return propflexutilvirtualdrivess
}
func flattenMapNiatelemetryNiaInventoryRef(p *models.NiatelemetryNiaInventoryRef, d *schema.ResourceData) []map[string]interface{} {

	var propdevices []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdevice := make(map[string]interface{})
	propdevice["moid"] = item.Moid
	propdevice["object_type"] = item.ObjectType
	propdevice["selector"] = item.Selector

	propdevices = append(propdevices, propdevice)
	return propdevices
}
func flattenMapAssetDeviceConnectionRef(p *models.AssetDeviceConnectionRef, d *schema.ResourceData) []map[string]interface{} {

	var propdeviceregistrations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propdeviceregistration := make(map[string]interface{})
	propdeviceregistration["moid"] = item.Moid
	propdeviceregistration["object_type"] = item.ObjectType
	propdeviceregistration["selector"] = item.Selector

	propdeviceregistrations = append(propdeviceregistrations, propdeviceregistration)
	return propdeviceregistrations
}
func flattenListPolicyinventoryJobInfo(p []*models.PolicyinventoryJobInfo, d *schema.ResourceData) []map[string]interface{} {
	var propjobinfos []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propjobinfo := make(map[string]interface{})
		propjobinfo["execution_status"] = item.ExecutionStatus
		propjobinfo["last_scheduled_time"] = item.LastScheduledTime
		propjobinfo["object_type"] = item.ObjectType
		propjobinfo["policy_id"] = item.PolicyID
		propjobinfo["policy_name"] = item.PolicyName
		propjobinfos = append(propjobinfos, propjobinfo)
	}
	return propjobinfos
}
func flattenListStoragePhysicalDiskUsageRef(p []*models.StoragePhysicalDiskUsageRef, d *schema.ResourceData) []map[string]interface{} {
	var propphysicaldiskusagess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propphysicaldiskusages := make(map[string]interface{})
		propphysicaldiskusages["moid"] = item.Moid
		propphysicaldiskusages["object_type"] = item.ObjectType
		propphysicaldiskusages["selector"] = item.Selector
		propphysicaldiskusagess = append(propphysicaldiskusagess, propphysicaldiskusages)
	}
	return propphysicaldiskusagess
}
func flattenListStorageVdMemberEpRef(p []*models.StorageVdMemberEpRef, d *schema.ResourceData) []map[string]interface{} {
	var propvdmemberepss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propvdmembereps := make(map[string]interface{})
		propvdmembereps["moid"] = item.Moid
		propvdmembereps["object_type"] = item.ObjectType
		propvdmembereps["selector"] = item.Selector
		propvdmemberepss = append(propvdmemberepss, propvdmembereps)
	}
	return propvdmemberepss
}
func flattenMapStorageVirtualDriveExtensionRef(p *models.StorageVirtualDriveExtensionRef, d *schema.ResourceData) []map[string]interface{} {

	var propvirtualdriveextensions []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propvirtualdriveextension := make(map[string]interface{})
	propvirtualdriveextension["moid"] = item.Moid
	propvirtualdriveextension["object_type"] = item.ObjectType
	propvirtualdriveextension["selector"] = item.Selector

	propvirtualdriveextensions = append(propvirtualdriveextensions, propvirtualdriveextension)
	return propvirtualdriveextensions
}
func flattenMapPortSubGroupRef(p *models.PortSubGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var propportsubgroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propportsubgroup := make(map[string]interface{})
	propportsubgroup["moid"] = item.Moid
	propportsubgroup["object_type"] = item.ObjectType
	propportsubgroup["selector"] = item.Selector

	propportsubgroups = append(propportsubgroups, propportsubgroup)
	return propportsubgroups
}
func flattenListApplianceKeyValuePair(p []*models.ApplianceKeyValuePair, d *schema.ResourceData) []map[string]interface{} {
	var propcapabilitiess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcapabilities := make(map[string]interface{})
		propcapabilities["key"] = item.Key
		propcapabilities["object_type"] = item.ObjectType
		propcapabilities["value"] = item.Value
		propcapabilitiess = append(propcapabilitiess, propcapabilities)
	}
	return propcapabilitiess
}
func flattenListAdapterUnitRef(p []*models.AdapterUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var propadapterss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propadapters := make(map[string]interface{})
		propadapters["moid"] = item.Moid
		propadapters["object_type"] = item.ObjectType
		propadapters["selector"] = item.Selector
		propadapterss = append(propadapterss, propadapters)
	}
	return propadapterss
}
func flattenListBiosUnitRef(p []*models.BiosUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var propbiosunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propbiosunits := make(map[string]interface{})
		propbiosunits["moid"] = item.Moid
		propbiosunits["object_type"] = item.ObjectType
		propbiosunits["selector"] = item.Selector
		propbiosunitss = append(propbiosunitss, propbiosunits)
	}
	return propbiosunitss
}
func flattenListEquipmentIoExpanderRef(p []*models.EquipmentIoExpanderRef, d *schema.ResourceData) []map[string]interface{} {
	var propequipmentioexpanderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propequipmentioexpanders := make(map[string]interface{})
		propequipmentioexpanders["moid"] = item.Moid
		propequipmentioexpanders["object_type"] = item.ObjectType
		propequipmentioexpanders["selector"] = item.Selector
		propequipmentioexpanderss = append(propequipmentioexpanderss, propequipmentioexpanders)
	}
	return propequipmentioexpanderss
}
func flattenListInventoryGenericInventoryHolderRef(p []*models.InventoryGenericInventoryHolderRef, d *schema.ResourceData) []map[string]interface{} {
	var propgenericinventoryholderss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propgenericinventoryholders := make(map[string]interface{})
		propgenericinventoryholders["moid"] = item.Moid
		propgenericinventoryholders["object_type"] = item.ObjectType
		propgenericinventoryholders["selector"] = item.Selector
		propgenericinventoryholderss = append(propgenericinventoryholderss, propgenericinventoryholders)
	}
	return propgenericinventoryholderss
}
func flattenListPciDeviceRef(p []*models.PciDeviceRef, d *schema.ResourceData) []map[string]interface{} {
	var proppcidevicess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppcidevices := make(map[string]interface{})
		proppcidevices["moid"] = item.Moid
		proppcidevices["object_type"] = item.ObjectType
		proppcidevices["selector"] = item.Selector
		proppcidevicess = append(proppcidevicess, proppcidevices)
	}
	return proppcidevicess
}
func flattenMapTopSystemRef(p *models.TopSystemRef, d *schema.ResourceData) []map[string]interface{} {

	var proptopsystems []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proptopsystem := make(map[string]interface{})
	proptopsystem["moid"] = item.Moid
	proptopsystem["object_type"] = item.ObjectType
	proptopsystem["selector"] = item.Selector

	proptopsystems = append(proptopsystems, proptopsystem)
	return proptopsystems
}
func flattenMapStorageFlexUtilControllerRef(p *models.StorageFlexUtilControllerRef, d *schema.ResourceData) []map[string]interface{} {

	var propstorageflexutilcontrollers []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstorageflexutilcontroller := make(map[string]interface{})
	propstorageflexutilcontroller["moid"] = item.Moid
	propstorageflexutilcontroller["object_type"] = item.ObjectType
	propstorageflexutilcontroller["selector"] = item.Selector

	propstorageflexutilcontrollers = append(propstorageflexutilcontrollers, propstorageflexutilcontroller)
	return propstorageflexutilcontrollers
}
func flattenMapResourceGroupRef(p *models.ResourceGroupRef, d *schema.ResourceData) []map[string]interface{} {

	var propgroups []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propgroup := make(map[string]interface{})
	propgroup["moid"] = item.Moid
	propgroup["object_type"] = item.ObjectType
	propgroup["selector"] = item.Selector

	propgroups = append(propgroups, propgroup)
	return propgroups
}
func flattenListMemoryPersistentMemoryUnitRef(p []*models.MemoryPersistentMemoryUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var proppersistentmemoryunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppersistentmemoryunits := make(map[string]interface{})
		proppersistentmemoryunits["moid"] = item.Moid
		proppersistentmemoryunits["object_type"] = item.ObjectType
		proppersistentmemoryunits["selector"] = item.Selector
		proppersistentmemoryunitss = append(proppersistentmemoryunitss, proppersistentmemoryunits)
	}
	return proppersistentmemoryunitss
}
func flattenListMemoryUnitRef(p []*models.MemoryUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var propunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propunits := make(map[string]interface{})
		propunits["moid"] = item.Moid
		propunits["object_type"] = item.ObjectType
		propunits["selector"] = item.Selector
		propunitss = append(propunitss, propunits)
	}
	return propunitss
}
func flattenListIamResourcePermissionRef(p []*models.IamResourcePermissionRef, d *schema.ResourceData) []map[string]interface{} {
	var propresourcepermissionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propresourcepermissions := make(map[string]interface{})
		propresourcepermissions["moid"] = item.Moid
		propresourcepermissions["object_type"] = item.ObjectType
		propresourcepermissions["selector"] = item.Selector
		propresourcepermissionss = append(propresourcepermissionss, propresourcepermissions)
	}
	return propresourcepermissionss
}
func flattenListEquipmentTpmRef(p []*models.EquipmentTpmRef, d *schema.ResourceData) []map[string]interface{} {
	var propequipmenttpmss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propequipmenttpms := make(map[string]interface{})
		propequipmenttpms["moid"] = item.Moid
		propequipmenttpms["object_type"] = item.ObjectType
		propequipmenttpms["selector"] = item.Selector
		propequipmenttpmss = append(propequipmenttpmss, propequipmenttpms)
	}
	return propequipmenttpmss
}
func flattenListGraphicsCardRef(p []*models.GraphicsCardRef, d *schema.ResourceData) []map[string]interface{} {
	var propgraphicscardss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propgraphicscards := make(map[string]interface{})
		propgraphicscards["moid"] = item.Moid
		propgraphicscards["object_type"] = item.ObjectType
		propgraphicscards["selector"] = item.Selector
		propgraphicscardss = append(propgraphicscardss, propgraphicscards)
	}
	return propgraphicscardss
}
func flattenListMemoryArrayRef(p []*models.MemoryArrayRef, d *schema.ResourceData) []map[string]interface{} {
	var propmemoryarrayss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propmemoryarrays := make(map[string]interface{})
		propmemoryarrays["moid"] = item.Moid
		propmemoryarrays["object_type"] = item.ObjectType
		propmemoryarrays["selector"] = item.Selector
		propmemoryarrayss = append(propmemoryarrayss, propmemoryarrays)
	}
	return propmemoryarrayss
}
func flattenListPciCoprocessorCardRef(p []*models.PciCoprocessorCardRef, d *schema.ResourceData) []map[string]interface{} {
	var proppcicoprocessorcardss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppcicoprocessorcards := make(map[string]interface{})
		proppcicoprocessorcards["moid"] = item.Moid
		proppcicoprocessorcards["object_type"] = item.ObjectType
		proppcicoprocessorcards["selector"] = item.Selector
		proppcicoprocessorcardss = append(proppcicoprocessorcardss, proppcicoprocessorcards)
	}
	return proppcicoprocessorcardss
}
func flattenListPciSwitchRef(p []*models.PciSwitchRef, d *schema.ResourceData) []map[string]interface{} {
	var proppciswitchs []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppciswitch := make(map[string]interface{})
		proppciswitch["moid"] = item.Moid
		proppciswitch["object_type"] = item.ObjectType
		proppciswitch["selector"] = item.Selector
		proppciswitchs = append(proppciswitchs, proppciswitch)
	}
	return proppciswitchs
}
func flattenListProcessorUnitRef(p []*models.ProcessorUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var propprocessorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propprocessors := make(map[string]interface{})
		propprocessors["moid"] = item.Moid
		propprocessors["object_type"] = item.ObjectType
		propprocessors["selector"] = item.Selector
		propprocessorss = append(propprocessorss, propprocessors)
	}
	return propprocessorss
}
func flattenListSecurityUnitRef(p []*models.SecurityUnitRef, d *schema.ResourceData) []map[string]interface{} {
	var propsecurityunitss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsecurityunits := make(map[string]interface{})
		propsecurityunits["moid"] = item.Moid
		propsecurityunits["object_type"] = item.ObjectType
		propsecurityunits["selector"] = item.Selector
		propsecurityunitss = append(propsecurityunitss, propsecurityunits)
	}
	return propsecurityunitss
}
func flattenListStorageControllerRef(p []*models.StorageControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var propstoragecontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propstoragecontrollers := make(map[string]interface{})
		propstoragecontrollers["moid"] = item.Moid
		propstoragecontrollers["object_type"] = item.ObjectType
		propstoragecontrollers["selector"] = item.Selector
		propstoragecontrollerss = append(propstoragecontrollerss, propstoragecontrollers)
	}
	return propstoragecontrollerss
}
func flattenListStorageFlexFlashControllerRef(p []*models.StorageFlexFlashControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var propstorageflexflashcontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propstorageflexflashcontrollers := make(map[string]interface{})
		propstorageflexflashcontrollers["moid"] = item.Moid
		propstorageflexflashcontrollers["object_type"] = item.ObjectType
		propstorageflexflashcontrollers["selector"] = item.Selector
		propstorageflexflashcontrollerss = append(propstorageflexflashcontrollerss, propstorageflexflashcontrollers)
	}
	return propstorageflexflashcontrollerss
}
func flattenListStorageFlexUtilControllerRef(p []*models.StorageFlexUtilControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var propstorageflexutilcontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propstorageflexutilcontrollers := make(map[string]interface{})
		propstorageflexutilcontrollers["moid"] = item.Moid
		propstorageflexutilcontrollers["object_type"] = item.ObjectType
		propstorageflexutilcontrollers["selector"] = item.Selector
		propstorageflexutilcontrollerss = append(propstorageflexutilcontrollerss, propstorageflexutilcontrollers)
	}
	return propstorageflexutilcontrollerss
}
func flattenMapStorageProtectionGroupSnapshotRef(p *models.StorageProtectionGroupSnapshotRef, d *schema.ResourceData) []map[string]interface{} {

	var propprotectiongroupsnapshots []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propprotectiongroupsnapshot := make(map[string]interface{})
	propprotectiongroupsnapshot["moid"] = item.Moid
	propprotectiongroupsnapshot["object_type"] = item.ObjectType
	propprotectiongroupsnapshot["selector"] = item.Selector

	propprotectiongroupsnapshots = append(propprotectiongroupsnapshots, propprotectiongroupsnapshot)
	return propprotectiongroupsnapshots
}
func flattenListOnpremImagePackage(p []*models.OnpremImagePackage, d *schema.ResourceData) []map[string]interface{} {
	var propansiblepackagess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propansiblepackages := make(map[string]interface{})
		propansiblepackages["file_path"] = item.FilePath
		propansiblepackages["file_sha"] = item.FileSha
		propansiblepackages["file_size"] = item.FileSize
		propansiblepackages["file_time"] = item.FileTime
		propansiblepackages["filename"] = item.Filename
		propansiblepackages["name"] = item.Name
		propansiblepackages["object_type"] = item.ObjectType
		propansiblepackages["package_type"] = item.PackageType
		propansiblepackages["version"] = item.Version
		propansiblepackagess = append(propansiblepackagess, propansiblepackages)
	}
	return propansiblepackagess
}
func flattenListManagementInterfaceRef(p []*models.ManagementInterfaceRef, d *schema.ResourceData) []map[string]interface{} {
	var propmanagementinterfacess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propmanagementinterfaces := make(map[string]interface{})
		propmanagementinterfaces["moid"] = item.Moid
		propmanagementinterfaces["object_type"] = item.ObjectType
		propmanagementinterfaces["selector"] = item.Selector
		propmanagementinterfacess = append(propmanagementinterfacess, propmanagementinterfaces)
	}
	return propmanagementinterfacess
}
func flattenMapStorageSasExpanderRef(p *models.StorageSasExpanderRef, d *schema.ResourceData) []map[string]interface{} {

	var propstoragesasexpanders []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propstoragesasexpander := make(map[string]interface{})
	propstoragesasexpander["moid"] = item.Moid
	propstoragesasexpander["object_type"] = item.ObjectType
	propstoragesasexpander["selector"] = item.Selector

	propstoragesasexpanders = append(propstoragesasexpanders, propstoragesasexpander)
	return propstoragesasexpanders
}
func flattenMapEquipmentSwitchCardRef(p *models.EquipmentSwitchCardRef, d *schema.ResourceData) []map[string]interface{} {

	var propequipmentswitchcards []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propequipmentswitchcard := make(map[string]interface{})
	propequipmentswitchcard["moid"] = item.Moid
	propequipmentswitchcard["object_type"] = item.ObjectType
	propequipmentswitchcard["selector"] = item.Selector

	propequipmentswitchcards = append(propequipmentswitchcards, propequipmentswitchcard)
	return propequipmentswitchcards
}
func flattenListFcPhysicalPortRef(p []*models.FcPhysicalPortRef, d *schema.ResourceData) []map[string]interface{} {
	var propfcportss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propfcports := make(map[string]interface{})
		propfcports["moid"] = item.Moid
		propfcports["object_type"] = item.ObjectType
		propfcports["selector"] = item.Selector
		propfcportss = append(propfcportss, propfcports)
	}
	return propfcportss
}
func flattenListPortSubGroupRef(p []*models.PortSubGroupRef, d *schema.ResourceData) []map[string]interface{} {
	var propsubgroupss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsubgroups := make(map[string]interface{})
		propsubgroups["moid"] = item.Moid
		propsubgroups["object_type"] = item.ObjectType
		propsubgroups["selector"] = item.Selector
		propsubgroupss = append(propsubgroupss, propsubgroups)
	}
	return propsubgroupss
}
func flattenListMemoryPersistentMemoryNamespaceConfigResultRef(p []*models.MemoryPersistentMemoryNamespaceConfigResultRef, d *schema.ResourceData) []map[string]interface{} {
	var proppersistentmemorynamespaceconfigresultss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppersistentmemorynamespaceconfigresults := make(map[string]interface{})
		proppersistentmemorynamespaceconfigresults["moid"] = item.Moid
		proppersistentmemorynamespaceconfigresults["object_type"] = item.ObjectType
		proppersistentmemorynamespaceconfigresults["selector"] = item.Selector
		proppersistentmemorynamespaceconfigresultss = append(proppersistentmemorynamespaceconfigresultss, proppersistentmemorynamespaceconfigresults)
	}
	return proppersistentmemorynamespaceconfigresultss
}
func flattenListResourceMembershipRef(p []*models.ResourceMembershipRef, d *schema.ResourceData) []map[string]interface{} {
	var propmembershipss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propmemberships := make(map[string]interface{})
		propmemberships["moid"] = item.Moid
		propmemberships["object_type"] = item.ObjectType
		propmemberships["selector"] = item.Selector
		propmembershipss = append(propmembershipss, propmemberships)
	}
	return propmembershipss
}
func flattenListWorkflowTaskRetryInfo(p []*models.WorkflowTaskRetryInfo, d *schema.ResourceData) []map[string]interface{} {
	var proptaskinstidlists []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proptaskinstidlist := make(map[string]interface{})
		proptaskinstidlist["object_type"] = item.ObjectType
		proptaskinstidlist["status"] = item.Status
		proptaskinstidlist["task_inst_id"] = item.TaskInstID
		proptaskinstidlists = append(proptaskinstidlists, proptaskinstidlist)
	}
	return proptaskinstidlists
}
func flattenMapAssetSudiInfo(p *models.AssetSudiInfo, d *schema.ResourceData) []map[string]interface{} {

	var propsudis []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsudi := make(map[string]interface{})
	propsudi["object_type"] = item.ObjectType
	propsudi["pid"] = item.Pid
	propsudi["serial_number"] = item.SerialNumber
	propsudi["signature"] = item.Signature
	propsudi["status"] = item.Status
	propsudi["sudi_certificate"] = (func(p *models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {

		var propsudicertificates []map[string]interface{}
		if p == nil {
			return nil
		}
		item := *p
		propsudicertificate := make(map[string]interface{})
		propsudicertificate["issuer"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var propissuers []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propissuer := make(map[string]interface{})
			propissuer["common_name"] = item.CommonName
			propissuer["country"] = item.Country
			propissuer["locality"] = item.Locality
			propissuer["object_type"] = item.ObjectType
			propissuer["organization"] = item.Organization
			propissuer["organizational_unit"] = item.OrganizationalUnit
			propissuer["state"] = item.State

			propissuers = append(propissuers, propissuer)
			return propissuers
		})(item.Issuer, d)
		propsudicertificate["object_type"] = item.ObjectType
		propsudicertificate["pem_certificate"] = item.PemCertificate
		propsudicertificate["sha256_fingerprint"] = item.Sha256Fingerprint
		propsudicertificate["signature_algorithm"] = item.SignatureAlgorithm
		propsudicertificate["subject"] = (func(p *models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {

			var propsubjects []map[string]interface{}
			if p == nil {
				return nil
			}
			item := *p
			propsubject := make(map[string]interface{})
			propsubject["common_name"] = item.CommonName
			propsubject["country"] = item.Country
			propsubject["locality"] = item.Locality
			propsubject["object_type"] = item.ObjectType
			propsubject["organization"] = item.Organization
			propsubject["organizational_unit"] = item.OrganizationalUnit
			propsubject["state"] = item.State

			propsubjects = append(propsubjects, propsubject)
			return propsubjects
		})(item.Subject, d)

		propsudicertificates = append(propsudicertificates, propsudicertificate)
		return propsudicertificates
	})(item.SudiCertificate, d)

	propsudis = append(propsudis, propsudi)
	return propsudis
}
func flattenListMemoryPersistentMemoryRegionRef(p []*models.MemoryPersistentMemoryRegionRef, d *schema.ResourceData) []map[string]interface{} {
	var proppersistentmemoryregionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proppersistentmemoryregions := make(map[string]interface{})
		proppersistentmemoryregions["moid"] = item.Moid
		proppersistentmemoryregions["object_type"] = item.ObjectType
		proppersistentmemoryregions["selector"] = item.Selector
		proppersistentmemoryregionss = append(proppersistentmemoryregionss, proppersistentmemoryregions)
	}
	return proppersistentmemoryregionss
}
func flattenMapLicenseCustomerOpRef(p *models.LicenseCustomerOpRef, d *schema.ResourceData) []map[string]interface{} {

	var propcustomerops []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcustomerop := make(map[string]interface{})
	propcustomerop["moid"] = item.Moid
	propcustomerop["object_type"] = item.ObjectType
	propcustomerop["selector"] = item.Selector

	propcustomerops = append(propcustomerops, propcustomerop)
	return propcustomerops
}
func flattenListLicenseLicenseInfoRef(p []*models.LicenseLicenseInfoRef, d *schema.ResourceData) []map[string]interface{} {
	var proplicenseinfoss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proplicenseinfos := make(map[string]interface{})
		proplicenseinfos["moid"] = item.Moid
		proplicenseinfos["object_type"] = item.ObjectType
		proplicenseinfos["selector"] = item.Selector
		proplicenseinfoss = append(proplicenseinfoss, proplicenseinfos)
	}
	return proplicenseinfoss
}
func flattenMapLicenseSmartlicenseTokenRef(p *models.LicenseSmartlicenseTokenRef, d *schema.ResourceData) []map[string]interface{} {

	var propsmartlicensetokens []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propsmartlicensetoken := make(map[string]interface{})
	propsmartlicensetoken["moid"] = item.Moid
	propsmartlicensetoken["object_type"] = item.ObjectType
	propsmartlicensetoken["selector"] = item.Selector

	propsmartlicensetokens = append(propsmartlicensetokens, propsmartlicensetoken)
	return propsmartlicensetokens
}
func flattenMapForecastCatalogRef(p *models.ForecastCatalogRef, d *schema.ResourceData) []map[string]interface{} {

	var propcatalogs []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcatalog := make(map[string]interface{})
	propcatalog["moid"] = item.Moid
	propcatalog["object_type"] = item.ObjectType
	propcatalog["selector"] = item.Selector

	propcatalogs = append(propcatalogs, propcatalog)
	return propcatalogs
}
func flattenMapBiosBootModeRef(p *models.BiosBootModeRef, d *schema.ResourceData) []map[string]interface{} {

	var propbiosbootmodes []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propbiosbootmode := make(map[string]interface{})
	propbiosbootmode["moid"] = item.Moid
	propbiosbootmode["object_type"] = item.ObjectType
	propbiosbootmode["selector"] = item.Selector

	propbiosbootmodes = append(propbiosbootmodes, propbiosbootmode)
	return propbiosbootmodes
}
func flattenMapBootDeviceBootModeRef(p *models.BootDeviceBootModeRef, d *schema.ResourceData) []map[string]interface{} {

	var propbootdevicebootmodes []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propbootdevicebootmode := make(map[string]interface{})
	propbootdevicebootmode["moid"] = item.Moid
	propbootdevicebootmode["object_type"] = item.ObjectType
	propbootdevicebootmode["selector"] = item.Selector

	propbootdevicebootmodes = append(propbootdevicebootmodes, propbootdevicebootmode)
	return propbootdevicebootmodes
}
func flattenMapEquipmentRackEnclosureSlotRef(p *models.EquipmentRackEnclosureSlotRef, d *schema.ResourceData) []map[string]interface{} {

	var proprackenclosureslots []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	proprackenclosureslot := make(map[string]interface{})
	proprackenclosureslot["moid"] = item.Moid
	proprackenclosureslot["object_type"] = item.ObjectType
	proprackenclosureslot["selector"] = item.Selector

	proprackenclosureslots = append(proprackenclosureslots, proprackenclosureslot)
	return proprackenclosureslots
}
func flattenListMetaAccessPrivilege(p []*models.MetaAccessPrivilege, d *schema.ResourceData) []map[string]interface{} {
	var propaccessprivilegess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propaccessprivileges := make(map[string]interface{})
		propaccessprivileges["method"] = item.Method
		propaccessprivileges["object_type"] = item.ObjectType
		propaccessprivileges["privilege"] = item.Privilege
		propaccessprivilegess = append(propaccessprivilegess, propaccessprivileges)
	}
	return propaccessprivilegess
}
func flattenListMetaPropDefinition(p []*models.MetaPropDefinition, d *schema.ResourceData) []map[string]interface{} {
	var proppropertiess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propproperties := make(map[string]interface{})
		propproperties["api_access"] = item.APIAccess
		propproperties["name"] = item.Name
		propproperties["object_type"] = item.ObjectType
		propproperties["op_security"] = item.OpSecurity
		propproperties["search_weight"] = item.SearchWeight
		proppropertiess = append(proppropertiess, propproperties)
	}
	return proppropertiess
}
func flattenListMetaRelationshipDefinition(p []*models.MetaRelationshipDefinition, d *schema.ResourceData) []map[string]interface{} {
	var proprelationshipss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proprelationships := make(map[string]interface{})
		proprelationships["api_access"] = item.APIAccess
		proprelationships["collection"] = item.Collection
		proprelationships["name"] = item.Name
		proprelationships["object_type"] = item.ObjectType
		proprelationships["type"] = item.Type
		proprelationshipss = append(proprelationshipss, proprelationships)
	}
	return proprelationshipss
}
func flattenMapBiosUnitRef(p *models.BiosUnitRef, d *schema.ResourceData) []map[string]interface{} {

	var propbiosunits []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propbiosunit := make(map[string]interface{})
	propbiosunit["moid"] = item.Moid
	propbiosunit["object_type"] = item.ObjectType
	propbiosunit["selector"] = item.Selector

	propbiosunits = append(propbiosunits, propbiosunit)
	return propbiosunits
}
func flattenListPciLinkRef(p []*models.PciLinkRef, d *schema.ResourceData) []map[string]interface{} {
	var proplinkss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		proplinks := make(map[string]interface{})
		proplinks["moid"] = item.Moid
		proplinks["object_type"] = item.ObjectType
		proplinks["selector"] = item.Selector
		proplinkss = append(proplinkss, proplinks)
	}
	return proplinkss
}
func flattenMapIamAppRegistrationRef(p *models.IamAppRegistrationRef, d *schema.ResourceData) []map[string]interface{} {

	var propappregistrations []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propappregistration := make(map[string]interface{})
	propappregistration["moid"] = item.Moid
	propappregistration["object_type"] = item.ObjectType
	propappregistration["selector"] = item.Selector

	propappregistrations = append(propappregistrations, propappregistration)
	return propappregistrations
}
func flattenMapIamClientMeta(p *models.IamClientMeta, d *schema.ResourceData) []map[string]interface{} {

	var propusermetas []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propusermeta := make(map[string]interface{})
	propusermeta["device_model"] = item.DeviceModel
	propusermeta["object_type"] = item.ObjectType
	propusermeta["user_agent"] = item.UserAgent

	propusermetas = append(propusermetas, propusermeta)
	return propusermetas
}
func flattenListStorageReplicationBlackout(p []*models.StorageReplicationBlackout, d *schema.ResourceData) []map[string]interface{} {
	var propreplicationblackoutintervalss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propreplicationblackoutintervals := make(map[string]interface{})
		propreplicationblackoutintervals["object_type"] = item.ObjectType
		propreplicationblackoutintervalss = append(propreplicationblackoutintervalss, propreplicationblackoutintervals)
	}
	return propreplicationblackoutintervalss
}
func flattenListUcsdConnectorPack(p []*models.UcsdConnectorPack, d *schema.ResourceData) []map[string]interface{} {
	var propconnectorss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propconnectors := make(map[string]interface{})
		propconnectors["connector_feature"] = item.ConnectorFeature
		propconnectors["dependency_names"] = item.DependencyNames
		propconnectors["downloaded_version"] = item.DownloadedVersion
		propconnectors["name"] = item.Name
		propconnectors["object_type"] = item.ObjectType
		propconnectors["services"] = item.Services
		propconnectors["state"] = item.State
		propconnectors["version"] = item.Version
		propconnectorss = append(propconnectorss, propconnectors)
	}
	return propconnectorss
}
func flattenListOnpremUpgradePhase(p []*models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var propcompletedphasess []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcompletedphases := make(map[string]interface{})
		propcompletedphases["name"] = item.Name
		propcompletedphases["object_type"] = item.ObjectType
		propcompletedphasess = append(propcompletedphasess, propcompletedphases)
	}
	return propcompletedphasess
}
func flattenMapOnpremUpgradePhase(p *models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {

	var propcurrentphases []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propcurrentphase := make(map[string]interface{})
	propcurrentphase["name"] = item.Name
	propcurrentphase["object_type"] = item.ObjectType

	propcurrentphases = append(propcurrentphases, propcurrentphase)
	return propcurrentphases
}
func flattenMapApplianceImageBundleRef(p *models.ApplianceImageBundleRef, d *schema.ResourceData) []map[string]interface{} {

	var propimagebundles []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propimagebundle := make(map[string]interface{})
	propimagebundle["moid"] = item.Moid
	propimagebundle["object_type"] = item.ObjectType
	propimagebundle["selector"] = item.Selector

	propimagebundles = append(propimagebundles, propimagebundle)
	return propimagebundles
}
func flattenListGraphicsControllerRef(p []*models.GraphicsControllerRef, d *schema.ResourceData) []map[string]interface{} {
	var propgraphicscontrollerss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propgraphicscontrollers := make(map[string]interface{})
		propgraphicscontrollers["moid"] = item.Moid
		propgraphicscontrollers["object_type"] = item.ObjectType
		propgraphicscontrollers["selector"] = item.Selector
		propgraphicscontrollerss = append(propgraphicscontrollerss, propgraphicscontrollers)
	}
	return propgraphicscontrollerss
}
func flattenListEquipmentSwitchCardRef(p []*models.EquipmentSwitchCardRef, d *schema.ResourceData) []map[string]interface{} {
	var propcardss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propcards := make(map[string]interface{})
		propcards["moid"] = item.Moid
		propcards["object_type"] = item.ObjectType
		propcards["selector"] = item.Selector
		propcardss = append(propcardss, propcards)
	}
	return propcardss
}
func flattenMapManagementEntityRef(p *models.ManagementEntityRef, d *schema.ResourceData) []map[string]interface{} {

	var propmanagemententitys []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propmanagemententity := make(map[string]interface{})
	propmanagemententity["moid"] = item.Moid
	propmanagemententity["object_type"] = item.ObjectType
	propmanagemententity["selector"] = item.Selector

	propmanagemententitys = append(propmanagemententitys, propmanagemententity)
	return propmanagemententitys
}
func flattenMapFirmwareRunningFirmwareRef(p *models.FirmwareRunningFirmwareRef, d *schema.ResourceData) []map[string]interface{} {

	var propucsmrunningfirmwares []map[string]interface{}
	if p == nil {
		return nil
	}
	item := *p
	propucsmrunningfirmware := make(map[string]interface{})
	propucsmrunningfirmware["moid"] = item.Moid
	propucsmrunningfirmware["object_type"] = item.ObjectType
	propucsmrunningfirmware["selector"] = item.Selector

	propucsmrunningfirmwares = append(propucsmrunningfirmwares, propucsmrunningfirmware)
	return propucsmrunningfirmwares
}
func flattenListIamAccountPermissions(p []*models.IamAccountPermissions, d *schema.ResourceData) []map[string]interface{} {
	var propaccountpermissionss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propaccountpermissions := make(map[string]interface{})
		propaccountpermissions["account_identifier"] = item.AccountIdentifier
		propaccountpermissions["account_name"] = item.AccountName
		propaccountpermissions["account_status"] = item.AccountStatus
		propaccountpermissions["object_type"] = item.ObjectType
		propaccountpermissions["permissions"] = (func(p []*models.IamPermissionReference, d *schema.ResourceData) []map[string]interface{} {
			var proppermissionss []map[string]interface{}
			if p == nil {
				return nil
			}
			for _, item := range p {
				item := *item
				proppermissions := make(map[string]interface{})
				proppermissions["object_type"] = item.ObjectType
				proppermissions["permission_identifier"] = item.PermissionIdentifier
				proppermissions["permission_name"] = item.PermissionName
				proppermissionss = append(proppermissionss, proppermissions)
			}
			return proppermissionss
		})(item.Permissions, d)
		propaccountpermissionss = append(propaccountpermissionss, propaccountpermissions)
	}
	return propaccountpermissionss
}
func flattenListStorageSasPortRef(p []*models.StorageSasPortRef, d *schema.ResourceData) []map[string]interface{} {
	var propsasportss []map[string]interface{}
	if p == nil {
		return nil
	}
	for _, item := range p {
		item := *item
		propsasports := make(map[string]interface{})
		propsasports["moid"] = item.Moid
		propsasports["object_type"] = item.ObjectType
		propsasports["selector"] = item.Selector
		propsasportss = append(propsasportss, propsasports)
	}
	return propsasportss
}
