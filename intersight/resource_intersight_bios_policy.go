package intersight

import (
	"encoding/json"
	"fmt"
	"log"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBiosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceBiosPolicyCreate,
		Read:   resourceBiosPolicyRead,
		Update: resourceBiosPolicyUpdate,
		Delete: resourceBiosPolicyDelete,
		Schema: map[string]*schema.Schema{
			"acs_control_gpu1state": {
				Description: "BIOS Token for setting ACS Control GPU-1 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_gpu2state": {
				Description: "BIOS Token for setting ACS Control GPU-2 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_gpu3state": {
				Description: "BIOS Token for setting ACS Control GPU-3 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_gpu4state": {
				Description: "BIOS Token for setting ACS Control GPU-4 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_gpu5state": {
				Description: "BIOS Token for setting ACS Control GPU-5 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_gpu6state": {
				Description: "BIOS Token for setting ACS Control GPU-6 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_gpu7state": {
				Description: "BIOS Token for setting ACS Control GPU-7 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_gpu8state": {
				Description: "BIOS Token for setting ACS Control GPU-8 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_slot11state": {
				Description: "BIOS Token for setting ACS Control Slot 11 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_slot12state": {
				Description: "BIOS Token for setting ACS Control Slot 12 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_slot13state": {
				Description: "BIOS Token for setting ACS Control Slot 13 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"acs_control_slot14state": {
				Description: "BIOS Token for setting ACS Control Slot 14 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"adjacent_cache_line_prefetch": {
				Description: "BIOS Token for setting Adjacent Cache Line Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"all_usb_devices": {
				Description: "BIOS Token for setting All USB Devices configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"altitude": {
				Description: "BIOS Token for setting Altitude configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"aspm_support": {
				Description: "BIOS Token for setting ASPM Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"assert_nmi_on_perr": {
				Description: "BIOS Token for setting Assert NMI on PERR configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"assert_nmi_on_serr": {
				Description: "BIOS Token for setting Assert NMI on SERR configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"auto_cc_state": {
				Description: "BIOS Token for setting Autonomous Core C-state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"autonumous_cstate_enable": {
				Description: "BIOS Token for setting CPU Autonomous Cstate configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"baud_rate": {
				Description: "BIOS Token for setting Baud rate configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"bme_dma_mitigation": {
				Description: "BIOS Token for setting BME DMA Mitigation configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"boot_option_num_retry": {
				Description: "BIOS Token for setting Number of Retries configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"boot_option_re_cool_down": {
				Description: "BIOS Token for setting Cool Down Time  (sec) configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"boot_option_retry": {
				Description: "BIOS Token for setting Boot option retry configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"boot_performance_mode": {
				Description: "BIOS Token for setting Boot Performance Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_cpu_cpb": {
				Description: "BIOS Token for setting Core Performance Boost configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_cpu_gen_downcore_ctrl": {
				Description: "BIOS Token for setting Downcore control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_cpu_global_cstate_ctrl": {
				Description: "BIOS Token for setting Global C-state Control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_cpu_l1stream_hw_prefetcher": {
				Description: "BIOS Token for setting L1 Stream HW Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_cpu_l2stream_hw_prefetcher": {
				Description: "BIOS Token for setting L2 Stream HW Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_determinism_slider": {
				Description: "BIOS Token for setting Determinism Slider configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_gnb_nb_iommu": {
				Description: "BIOS Token for setting IOMMU configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_mem_ctrl_bank_group_swap_ddr4": {
				Description: "BIOS Token for setting Bank Group Swap configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmn_mem_map_bank_interleave_ddr4": {
				Description: "BIOS Token for setting Chipselect Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_cmnc_tdp_ctl": {
				Description: "BIOS Token for setting cTDP Control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_df_cmn_mem_intlv": {
				Description: "BIOS Token for setting Memory interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cbs_df_cmn_mem_intlv_size": {
				Description: "BIOS Token for setting Memory interleaving size configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cdn_enable": {
				Description: "BIOS Token for setting Consistent Device Naming configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cdn_support": {
				Description: "BIOS Token for setting CDN Support for LOM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"channel_inter_leave": {
				Description: "BIOS Token for setting Channel Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cisco_adaptive_mem_training": {
				Description: "BIOS Token for setting Adaptive Memory Training configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cisco_debug_level": {
				Description: "BIOS Token for setting BIOS Techlog Level configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cisco_oprom_launch_optimization": {
				Description: "BIOS Token for setting OptionROM Launch Optimization configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cke_low_policy": {
				Description: "BIOS Token for setting CKE Low Policy configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"closed_loop_therm_throtl": {
				Description: "BIOS Token for setting Closed Loop Therm Throt configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cmci_enable": {
				Description: "BIOS Token for setting Processor CMCI configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"config_tdp": {
				Description: "BIOS Token for setting Config TDP configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"console_redirection": {
				Description: "BIOS Token for setting Console redirection configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"core_multi_processing": {
				Description: "BIOS Token for setting Core MultiProcessing configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cpu_energy_performance": {
				Description: "BIOS Token for setting Energy Performance configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cpu_frequency_floor": {
				Description: "BIOS Token for setting Frequency Floor Override configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cpu_performance": {
				Description: "BIOS Token for setting CPU Performance configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"cpu_power_management": {
				Description: "BIOS Token for setting Power Technology configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"dcpmm_firmware_downgrade": {
				Description: "BIOS Token for setting DCPMM Firmware Downgrade configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"demand_scrub": {
				Description: "BIOS Token for setting Demand Scrub configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"direct_cache_access": {
				Description: "BIOS Token for setting Direct Cache Access Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"dram_clock_throttling": {
				Description: "BIOS Token for setting DRAM Clock Throttling configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"dram_refresh_rate": {
				Description: "BIOS Token for setting DRAM Refresh Rate configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"energy_efficient_turbo": {
				Description: "BIOS Token for setting Energy Efficient Turbo configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"eng_perf_tuning": {
				Description: "BIOS Token for setting Energy Performance Tuning configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"enhanced_intel_speed_step_tech": {
				Description: "BIOS Token for setting Enhanced Intel Speedstep (R) Technology configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"epp_profile": {
				Description: "BIOS Token for setting EPP Profile configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"execute_disable_bit": {
				Description: "BIOS Token for setting Execute Disable Bit configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"extended_apic": {
				Description: "BIOS Token for setting Local X2 Apic configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"flow_control": {
				Description: "BIOS Token for setting Flow Control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"frb2enable": {
				Description: "BIOS Token for setting FRB-2 Timer configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"hardware_prefetch": {
				Description: "BIOS Token for setting Hardware Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"hwpm_enable": {
				Description: "BIOS Token for setting CPU Hardware Power Management configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"imc_interleave": {
				Description: "BIOS Token for setting IMC Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_hyper_threading_tech": {
				Description: "BIOS Token for setting Intel HyperThreading Tech configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_speed_select": {
				Description: "BIOS Token for setting Intel Speed Select configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_turbo_boost_tech": {
				Description: "BIOS Token for setting Intel Turbo Boost Tech configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_virtualization_technology": {
				Description: "BIOS Token for setting Intel (R) VT configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_vt_for_directed_io": {
				Description: "BIOS Token for setting Intel VT for directed IO configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_vtd_coherency_support": {
				Description: "BIOS Token for setting Intel (R) VT-d Coherency Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_vtd_interrupt_remapping": {
				Description: "BIOS Token for setting Intel (R) VT-d Interrupt Remapping configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_vtd_pass_through_dma_support": {
				Description: "BIOS Token for setting Intel (R) VT-d PassThrough DMA support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"intel_vtdats_support": {
				Description: "BIOS Token for setting Intel VTD ATS support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"ioh_error_enable": {
				Description: "BIOS Token for setting IIO Error Enable configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"ioh_resource": {
				Description: "BIOS Token for setting IOH Resource Allocation configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"ip_prefetch": {
				Description: "BIOS Token for setting DCU IP Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"ipv4pxe": {
				Description: "BIOS Token for setting IPv4 PXE Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"ipv6pxe": {
				Description: "BIOS Token for setting IPV6 PXE Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"kti_prefetch": {
				Description: "BIOS Token for setting KTI Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"legacy_os_redirection": {
				Description: "BIOS Token for setting Legacy OS redirection configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"legacy_usb_support": {
				Description: "BIOS Token for setting Legacy USB Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"llc_prefetch": {
				Description: "BIOS Token for setting LLC Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"lom_port0state": {
				Description: "BIOS Token for setting LOM Port 0 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"lom_port1state": {
				Description: "BIOS Token for setting LOM Port 1 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"lom_port2state": {
				Description: "BIOS Token for setting LOM Port 2 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"lom_port3state": {
				Description: "BIOS Token for setting LOM Port 3 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"lom_ports_all_state": {
				Description: "BIOS Token for setting All Onboard LOM Ports configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"lv_ddr_mode": {
				Description: "BIOS Token for setting Low Voltage DDR Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"make_device_non_bootable": {
				Description: "BIOS Token for setting Make Device Non Bootable configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"memory_inter_leave": {
				Description: "BIOS Token for setting Memory Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"memory_mapped_io_above4gb": {
				Description: "BIOS Token for setting Memory mapped IO above 4GiB configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"memory_size_limit": {
				Description: "BIOS Token for setting Memory Size Limit in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mirroring_mode": {
				Description: "BIOS Token for setting Mirroring Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"mmcfg_base": {
				Description: "BIOS Token for setting MMCFG BASE configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_stack": {
				Description: "BIOS Token for setting Network Stack configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"numa_optimized": {
				Description: "BIOS Token for setting NUMA optimized configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"onboard10gbit_lom": {
				Description: "BIOS Token for setting Onboard 10Gbit LOM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"onboard_gbit_lom": {
				Description: "BIOS Token for setting Onboard Gbit LOM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"onboard_scu_storage_support": {
				Description: "BIOS Token for setting Onboard SCU Storage Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"onboard_scu_storage_sw_stack": {
				Description: "BIOS Token for setting Onboard SCU Storage SW Stack configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"os_boot_watchdog_timer": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"os_boot_watchdog_timer_policy": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer Policy configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"os_boot_watchdog_timer_timeout": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer Timeout configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"out_of_band_mgmt_port": {
				Description: "BIOS Token for setting Out-of-Band Mgmt Port configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"package_cstate_limit": {
				Description: "BIOS Token for setting Package C State Limit configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"partial_mirror_mode_config": {
				Description: "BIOS Token for setting Partial Memory Mirror Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"partial_mirror_percent": {
				Description: "BIOS Token for setting Partial Mirror percentage configuration (0.00-50.00).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value1": {
				Description: "BIOS Token for setting Partial Mirror1 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value2": {
				Description: "BIOS Token for setting Partial Mirror2 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value3": {
				Description: "BIOS Token for setting Partial Mirror3 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value4": {
				Description: "BIOS Token for setting Partial Mirror4 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"patrol_scrub": {
				Description: "BIOS Token for setting Patrol Scrub configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"patrol_scrub_duration": {
				Description: "BIOS Token for setting Patrol Scrub Interval configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pc_ie_ras_support": {
				Description: "BIOS Token for setting PCIe RAS Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pc_ie_ssd_hot_plug_support": {
				Description: "BIOS Token for setting NVMe SSD Hot-Plug Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pch_usb30mode": {
				Description: "BIOS Token for setting xHCI Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pci_option_ro_ms": {
				Description: "BIOS Token for setting All PCIe Slots OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pci_rom_clp": {
				Description: "BIOS Token for setting PCI ROM CLP configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pop_support": {
				Description: "BIOS Token for setting Power ON Password configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"post_error_pause": {
				Description: "BIOS Token for setting POST Error Pause configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"processor_c1e": {
				Description: "BIOS Token for setting Processor C1E configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"processor_c3report": {
				Description: "BIOS Token for setting Processor C3 Report configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"processor_c6report": {
				Description: "BIOS Token for setting Processor C6 Report configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"processor_cstate": {
				Description: "BIOS Token for setting CPU C State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"psata": {
				Description: "BIOS Token for setting P-SATA mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pstate_coord_type": {
				Description: "BIOS Token for setting P-STATE Coordination configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"putty_key_pad": {
				Description: "BIOS Token for setting Putty KeyPad configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"pwr_perf_tuning": {
				Description: "BIOS Token for setting Power Performance Tuning configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"qpi_link_frequency": {
				Description: "BIOS Token for setting QPI Link Frequency Select configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"qpi_snoop_mode": {
				Description: "BIOS Token for setting QPI Snoop Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"rank_inter_leave": {
				Description: "BIOS Token for setting Rank Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"redirection_after_post": {
				Description: "BIOS Token for setting Redirection After BIOS POST configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"sata_mode_select": {
				Description: "BIOS Token for setting SATA mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"select_memory_ras_configuration": {
				Description: "BIOS Token for setting SelectMemory RAS configuration configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"select_ppr_type": {
				Description: "BIOS Token for setting Select PPR Type configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"serial_port_aenable": {
				Description: "BIOS Token for setting Serial A Enable configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"single_pctl_enable": {
				Description: "BIOS Token for setting Single PCTL configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot10link_speed": {
				Description: "BIOS Token for setting PCIe Slot:10 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot10state": {
				Description: "BIOS Token for setting Slot 10 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot11link_speed": {
				Description: "BIOS Token for setting PCIe Slot:11 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot11state": {
				Description: "BIOS Token for setting Slot 11 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot12link_speed": {
				Description: "BIOS Token for setting PCIe Slot:12 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot12state": {
				Description: "BIOS Token for setting Slot 12 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot13state": {
				Description: "BIOS Token for setting PCIe Slot 13 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot14state": {
				Description: "BIOS Token for setting PCIe Slot 14 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot1state": {
				Description: "BIOS Token for setting Slot 1 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot2state": {
				Description: "BIOS Token for setting Slot 2 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot3link_speed": {
				Description: "BIOS Token for setting PCIe Slot:3 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot3state": {
				Description: "BIOS Token for setting Slot 3 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot4link_speed": {
				Description: "BIOS Token for setting PCIe Slot:4 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot4state": {
				Description: "BIOS Token for setting Slot 4 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:5 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot5state": {
				Description: "BIOS Token for setting Slot 5 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:6 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot6state": {
				Description: "BIOS Token for setting Slot 6 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot7link_speed": {
				Description: "BIOS Token for setting PCIe Slot:7 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot7state": {
				Description: "BIOS Token for setting Slot 7 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot8link_speed": {
				Description: "BIOS Token for setting PCIe Slot:8 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot8state": {
				Description: "BIOS Token for setting Slot 8 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot9link_speed": {
				Description: "BIOS Token for setting PCIe Slot:9 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot9state": {
				Description: "BIOS Token for setting Slot 9 State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_flom_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FLOM Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_front_nvme1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front Nvme1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_front_nvme2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front Nvme2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_front_slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_front_slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu1state": {
				Description: "BIOS Token for setting GPU1 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu2state": {
				Description: "BIOS Token for setting GPU2 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu3state": {
				Description: "BIOS Token for setting GPU3 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu4state": {
				Description: "BIOS Token for setting GPU4 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu5state": {
				Description: "BIOS Token for setting GPU5 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu6state": {
				Description: "BIOS Token for setting GPU6 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu7state": {
				Description: "BIOS Token for setting GPU7 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_gpu8state": {
				Description: "BIOS Token for setting GPU8 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_hba_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:HBA Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_hba_state": {
				Description: "BIOS Token for setting PCIe Slot:HBA OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_lom1link": {
				Description: "BIOS Token for setting PCIe LOM:1 Link configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_lom2link": {
				Description: "BIOS Token for setting PCIe LOM:2 Link configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_mezz_state": {
				Description: "BIOS Token for setting Slot Mezz State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_mlom_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:MLOM Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_mlom_state": {
				Description: "BIOS Token for setting PCIe Slot MLOM OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_mraid_link_speed": {
				Description: "BIOS Token for setting MRAID Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_mraid_state": {
				Description: "BIOS Token for setting PCIe Slot MRAID OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n10state": {
				Description: "BIOS Token for setting PCIe Slot N10 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n11state": {
				Description: "BIOS Token for setting PCIe Slot N11 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n12state": {
				Description: "BIOS Token for setting PCIe Slot N12 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n13state": {
				Description: "BIOS Token for setting PCIe Slot N13 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n14state": {
				Description: "BIOS Token for setting PCIe Slot N14 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n15state": {
				Description: "BIOS Token for setting PCIe Slot N15 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n16state": {
				Description: "BIOS Token for setting PCIe Slot N16 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n17state": {
				Description: "BIOS Token for setting PCIe Slot N17 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n18state": {
				Description: "BIOS Token for setting PCIe Slot N18 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n19state": {
				Description: "BIOS Token for setting PCIe Slot N19 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n1state": {
				Description: "BIOS Token for setting PCIe Slot N1 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n20state": {
				Description: "BIOS Token for setting PCIe Slot N20 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n21state": {
				Description: "BIOS Token for setting PCIe Slot N21 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n22state": {
				Description: "BIOS Token for setting PCIe Slot N22 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n23state": {
				Description: "BIOS Token for setting PCIe Slot N23 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n24state": {
				Description: "BIOS Token for setting PCIe Slot N24 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n2state": {
				Description: "BIOS Token for setting PCIe Slot N2 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n3state": {
				Description: "BIOS Token for setting PCIe Slot N3 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n4state": {
				Description: "BIOS Token for setting PCIe Slot N4 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n5state": {
				Description: "BIOS Token for setting PCIe Slot N5 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n6state": {
				Description: "BIOS Token for setting PCIe Slot N6 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n7state": {
				Description: "BIOS Token for setting PCIe Slot N7 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n8state": {
				Description: "BIOS Token for setting PCIe Slot N8 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_n9state": {
				Description: "BIOS Token for setting PCIe Slot N9 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_raid_link_speed": {
				Description: "BIOS Token for setting RAID Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_raid_state": {
				Description: "BIOS Token for setting PCIe Slot RAID OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Rear Nvme1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme1state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 1 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Rear Nvme2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme2state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 2 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme3state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 3 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme4state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 4 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme5state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 5 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme6state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 6 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme7state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 7 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_rear_nvme8state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 8 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser1slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser1slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser1slot3link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot3 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser2slot4link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot4 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser2slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot5 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_riser2slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot6 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_sas_state": {
				Description: "BIOS Token for setting PCIe Slot:SAS OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_ssd_slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FrontSSD1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"slot_ssd_slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FrontSSD2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"smee": {
				Description: "BIOS Token for setting SMEE configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"smt_mode": {
				Description: "BIOS Token for setting SMT Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"snc": {
				Description: "BIOS Token for setting Sub Numa Clustering configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"sparing_mode": {
				Description: "BIOS Token for setting Sparing Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"sr_iov": {
				Description: "BIOS Token for setting SR-IOV Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"streamer_prefetch": {
				Description: "BIOS Token for setting DCU Streamer Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"svm_mode": {
				Description: "BIOS Token for setting SVM Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
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
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"terminal_type": {
				Description: "BIOS Token for setting Terminal Type configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"tpm_control": {
				Description: "BIOS Token for setting Trusted Platform Module State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"tpm_support": {
				Description: "BIOS Token for setting TPM Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"txt_support": {
				Description: "BIOS Token for setting Intel Trusted Execution Technology Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"ucsm_boot_order_rule": {
				Description: "BIOS Token for setting Boot Order Rules configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_emul6064": {
				Description: "BIOS Token for setting Port 60/64 Emulation configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_port_front": {
				Description: "BIOS Token for setting USB Port Front configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_port_internal": {
				Description: "BIOS Token for setting USB Port Internal configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_port_kvm": {
				Description: "BIOS Token for setting USB Port KVM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_port_rear": {
				Description: "BIOS Token for setting USB Port Rear configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_port_sd_card": {
				Description: "BIOS Token for setting USB Port SD Card configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_port_vmedia": {
				Description: "BIOS Token for setting USB Port VMedia configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"usb_xhci_support": {
				Description: "BIOS Token for setting XHCI Legacy Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"vga_priority": {
				Description: "BIOS Token for setting VGA Priority configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"vmd_enable": {
				Description: "BIOS Token for setting VMD Enablement configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"work_load_config": {
				Description: "BIOS Token for setting Workload Configuration configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
			"xpt_prefetch": {
				Description: "BIOS Token for setting XPT Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "platform-default",
			},
		},
	}
}

func resourceBiosPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewBiosPolicyWithDefaults()
	if v, ok := d.GetOk("acs_control_gpu1state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu1state(x)
	}

	if v, ok := d.GetOk("acs_control_gpu2state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu2state(x)
	}

	if v, ok := d.GetOk("acs_control_gpu3state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu3state(x)
	}

	if v, ok := d.GetOk("acs_control_gpu4state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu4state(x)
	}

	if v, ok := d.GetOk("acs_control_gpu5state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu5state(x)
	}

	if v, ok := d.GetOk("acs_control_gpu6state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu6state(x)
	}

	if v, ok := d.GetOk("acs_control_gpu7state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu7state(x)
	}

	if v, ok := d.GetOk("acs_control_gpu8state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu8state(x)
	}

	if v, ok := d.GetOk("acs_control_slot11state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot11state(x)
	}

	if v, ok := d.GetOk("acs_control_slot12state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot12state(x)
	}

	if v, ok := d.GetOk("acs_control_slot13state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot13state(x)
	}

	if v, ok := d.GetOk("acs_control_slot14state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot14state(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("adjacent_cache_line_prefetch"); ok {
		x := (v.(string))
		o.SetAdjacentCacheLinePrefetch(x)
	}

	if v, ok := d.GetOk("all_usb_devices"); ok {
		x := (v.(string))
		o.SetAllUsbDevices(x)
	}

	if v, ok := d.GetOk("altitude"); ok {
		x := (v.(string))
		o.SetAltitude(x)
	}

	if v, ok := d.GetOk("aspm_support"); ok {
		x := (v.(string))
		o.SetAspmSupport(x)
	}

	if v, ok := d.GetOk("assert_nmi_on_perr"); ok {
		x := (v.(string))
		o.SetAssertNmiOnPerr(x)
	}

	if v, ok := d.GetOk("assert_nmi_on_serr"); ok {
		x := (v.(string))
		o.SetAssertNmiOnSerr(x)
	}

	if v, ok := d.GetOk("auto_cc_state"); ok {
		x := (v.(string))
		o.SetAutoCcState(x)
	}

	if v, ok := d.GetOk("autonumous_cstate_enable"); ok {
		x := (v.(string))
		o.SetAutonumousCstateEnable(x)
	}

	if v, ok := d.GetOk("baud_rate"); ok {
		x := (v.(string))
		o.SetBaudRate(x)
	}

	if v, ok := d.GetOk("bme_dma_mitigation"); ok {
		x := (v.(string))
		o.SetBmeDmaMitigation(x)
	}

	if v, ok := d.GetOk("boot_option_num_retry"); ok {
		x := (v.(string))
		o.SetBootOptionNumRetry(x)
	}

	if v, ok := d.GetOk("boot_option_re_cool_down"); ok {
		x := (v.(string))
		o.SetBootOptionReCoolDown(x)
	}

	if v, ok := d.GetOk("boot_option_retry"); ok {
		x := (v.(string))
		o.SetBootOptionRetry(x)
	}

	if v, ok := d.GetOk("boot_performance_mode"); ok {
		x := (v.(string))
		o.SetBootPerformanceMode(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_cpb"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuCpb(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_gen_downcore_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuGenDowncoreCtrl(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_global_cstate_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuGlobalCstateCtrl(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_l1stream_hw_prefetcher"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuL1streamHwPrefetcher(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_l2stream_hw_prefetcher"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuL2streamHwPrefetcher(x)
	}

	if v, ok := d.GetOk("cbs_cmn_determinism_slider"); ok {
		x := (v.(string))
		o.SetCbsCmnDeterminismSlider(x)
	}

	if v, ok := d.GetOk("cbs_cmn_gnb_nb_iommu"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbNbIommu(x)
	}

	if v, ok := d.GetOk("cbs_cmn_mem_ctrl_bank_group_swap_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemCtrlBankGroupSwapDdr4(x)
	}

	if v, ok := d.GetOk("cbs_cmn_mem_map_bank_interleave_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemMapBankInterleaveDdr4(x)
	}

	if v, ok := d.GetOk("cbs_cmnc_tdp_ctl"); ok {
		x := (v.(string))
		o.SetCbsCmncTdpCtl(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlv(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv_size"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlvSize(x)
	}

	if v, ok := d.GetOk("cdn_enable"); ok {
		x := (v.(string))
		o.SetCdnEnable(x)
	}

	if v, ok := d.GetOk("cdn_support"); ok {
		x := (v.(string))
		o.SetCdnSupport(x)
	}

	if v, ok := d.GetOk("channel_inter_leave"); ok {
		x := (v.(string))
		o.SetChannelInterLeave(x)
	}

	if v, ok := d.GetOk("cisco_adaptive_mem_training"); ok {
		x := (v.(string))
		o.SetCiscoAdaptiveMemTraining(x)
	}

	if v, ok := d.GetOk("cisco_debug_level"); ok {
		x := (v.(string))
		o.SetCiscoDebugLevel(x)
	}

	if v, ok := d.GetOk("cisco_oprom_launch_optimization"); ok {
		x := (v.(string))
		o.SetCiscoOpromLaunchOptimization(x)
	}

	if v, ok := d.GetOk("cke_low_policy"); ok {
		x := (v.(string))
		o.SetCkeLowPolicy(x)
	}

	o.SetClassId("bios.Policy")

	if v, ok := d.GetOk("closed_loop_therm_throtl"); ok {
		x := (v.(string))
		o.SetClosedLoopThermThrotl(x)
	}

	if v, ok := d.GetOk("cmci_enable"); ok {
		x := (v.(string))
		o.SetCmciEnable(x)
	}

	if v, ok := d.GetOk("config_tdp"); ok {
		x := (v.(string))
		o.SetConfigTdp(x)
	}

	if v, ok := d.GetOk("console_redirection"); ok {
		x := (v.(string))
		o.SetConsoleRedirection(x)
	}

	if v, ok := d.GetOk("core_multi_processing"); ok {
		x := (v.(string))
		o.SetCoreMultiProcessing(x)
	}

	if v, ok := d.GetOk("cpu_energy_performance"); ok {
		x := (v.(string))
		o.SetCpuEnergyPerformance(x)
	}

	if v, ok := d.GetOk("cpu_frequency_floor"); ok {
		x := (v.(string))
		o.SetCpuFrequencyFloor(x)
	}

	if v, ok := d.GetOk("cpu_performance"); ok {
		x := (v.(string))
		o.SetCpuPerformance(x)
	}

	if v, ok := d.GetOk("cpu_power_management"); ok {
		x := (v.(string))
		o.SetCpuPowerManagement(x)
	}

	if v, ok := d.GetOk("dcpmm_firmware_downgrade"); ok {
		x := (v.(string))
		o.SetDcpmmFirmwareDowngrade(x)
	}

	if v, ok := d.GetOk("demand_scrub"); ok {
		x := (v.(string))
		o.SetDemandScrub(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("direct_cache_access"); ok {
		x := (v.(string))
		o.SetDirectCacheAccess(x)
	}

	if v, ok := d.GetOk("dram_clock_throttling"); ok {
		x := (v.(string))
		o.SetDramClockThrottling(x)
	}

	if v, ok := d.GetOk("dram_refresh_rate"); ok {
		x := (v.(string))
		o.SetDramRefreshRate(x)
	}

	if v, ok := d.GetOk("energy_efficient_turbo"); ok {
		x := (v.(string))
		o.SetEnergyEfficientTurbo(x)
	}

	if v, ok := d.GetOk("eng_perf_tuning"); ok {
		x := (v.(string))
		o.SetEngPerfTuning(x)
	}

	if v, ok := d.GetOk("enhanced_intel_speed_step_tech"); ok {
		x := (v.(string))
		o.SetEnhancedIntelSpeedStepTech(x)
	}

	if v, ok := d.GetOk("epp_profile"); ok {
		x := (v.(string))
		o.SetEppProfile(x)
	}

	if v, ok := d.GetOk("execute_disable_bit"); ok {
		x := (v.(string))
		o.SetExecuteDisableBit(x)
	}

	if v, ok := d.GetOk("extended_apic"); ok {
		x := (v.(string))
		o.SetExtendedApic(x)
	}

	if v, ok := d.GetOk("flow_control"); ok {
		x := (v.(string))
		o.SetFlowControl(x)
	}

	if v, ok := d.GetOk("frb2enable"); ok {
		x := (v.(string))
		o.SetFrb2enable(x)
	}

	if v, ok := d.GetOk("hardware_prefetch"); ok {
		x := (v.(string))
		o.SetHardwarePrefetch(x)
	}

	if v, ok := d.GetOk("hwpm_enable"); ok {
		x := (v.(string))
		o.SetHwpmEnable(x)
	}

	if v, ok := d.GetOk("imc_interleave"); ok {
		x := (v.(string))
		o.SetImcInterleave(x)
	}

	if v, ok := d.GetOk("intel_hyper_threading_tech"); ok {
		x := (v.(string))
		o.SetIntelHyperThreadingTech(x)
	}

	if v, ok := d.GetOk("intel_speed_select"); ok {
		x := (v.(string))
		o.SetIntelSpeedSelect(x)
	}

	if v, ok := d.GetOk("intel_turbo_boost_tech"); ok {
		x := (v.(string))
		o.SetIntelTurboBoostTech(x)
	}

	if v, ok := d.GetOk("intel_virtualization_technology"); ok {
		x := (v.(string))
		o.SetIntelVirtualizationTechnology(x)
	}

	if v, ok := d.GetOk("intel_vt_for_directed_io"); ok {
		x := (v.(string))
		o.SetIntelVtForDirectedIo(x)
	}

	if v, ok := d.GetOk("intel_vtd_coherency_support"); ok {
		x := (v.(string))
		o.SetIntelVtdCoherencySupport(x)
	}

	if v, ok := d.GetOk("intel_vtd_interrupt_remapping"); ok {
		x := (v.(string))
		o.SetIntelVtdInterruptRemapping(x)
	}

	if v, ok := d.GetOk("intel_vtd_pass_through_dma_support"); ok {
		x := (v.(string))
		o.SetIntelVtdPassThroughDmaSupport(x)
	}

	if v, ok := d.GetOk("intel_vtdats_support"); ok {
		x := (v.(string))
		o.SetIntelVtdatsSupport(x)
	}

	if v, ok := d.GetOk("ioh_error_enable"); ok {
		x := (v.(string))
		o.SetIohErrorEnable(x)
	}

	if v, ok := d.GetOk("ioh_resource"); ok {
		x := (v.(string))
		o.SetIohResource(x)
	}

	if v, ok := d.GetOk("ip_prefetch"); ok {
		x := (v.(string))
		o.SetIpPrefetch(x)
	}

	if v, ok := d.GetOk("ipv4pxe"); ok {
		x := (v.(string))
		o.SetIpv4pxe(x)
	}

	if v, ok := d.GetOk("ipv6pxe"); ok {
		x := (v.(string))
		o.SetIpv6pxe(x)
	}

	if v, ok := d.GetOk("kti_prefetch"); ok {
		x := (v.(string))
		o.SetKtiPrefetch(x)
	}

	if v, ok := d.GetOk("legacy_os_redirection"); ok {
		x := (v.(string))
		o.SetLegacyOsRedirection(x)
	}

	if v, ok := d.GetOk("legacy_usb_support"); ok {
		x := (v.(string))
		o.SetLegacyUsbSupport(x)
	}

	if v, ok := d.GetOk("llc_prefetch"); ok {
		x := (v.(string))
		o.SetLlcPrefetch(x)
	}

	if v, ok := d.GetOk("lom_port0state"); ok {
		x := (v.(string))
		o.SetLomPort0state(x)
	}

	if v, ok := d.GetOk("lom_port1state"); ok {
		x := (v.(string))
		o.SetLomPort1state(x)
	}

	if v, ok := d.GetOk("lom_port2state"); ok {
		x := (v.(string))
		o.SetLomPort2state(x)
	}

	if v, ok := d.GetOk("lom_port3state"); ok {
		x := (v.(string))
		o.SetLomPort3state(x)
	}

	if v, ok := d.GetOk("lom_ports_all_state"); ok {
		x := (v.(string))
		o.SetLomPortsAllState(x)
	}

	if v, ok := d.GetOk("lv_ddr_mode"); ok {
		x := (v.(string))
		o.SetLvDdrMode(x)
	}

	if v, ok := d.GetOk("make_device_non_bootable"); ok {
		x := (v.(string))
		o.SetMakeDeviceNonBootable(x)
	}

	if v, ok := d.GetOk("memory_inter_leave"); ok {
		x := (v.(string))
		o.SetMemoryInterLeave(x)
	}

	if v, ok := d.GetOk("memory_mapped_io_above4gb"); ok {
		x := (v.(string))
		o.SetMemoryMappedIoAbove4gb(x)
	}

	if v, ok := d.GetOk("memory_size_limit"); ok {
		x := (v.(string))
		o.SetMemorySizeLimit(x)
	}

	if v, ok := d.GetOk("mirroring_mode"); ok {
		x := (v.(string))
		o.SetMirroringMode(x)
	}

	if v, ok := d.GetOk("mmcfg_base"); ok {
		x := (v.(string))
		o.SetMmcfgBase(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	if v, ok := d.GetOk("network_stack"); ok {
		x := (v.(string))
		o.SetNetworkStack(x)
	}

	if v, ok := d.GetOk("numa_optimized"); ok {
		x := (v.(string))
		o.SetNumaOptimized(x)
	}

	o.SetObjectType("bios.Policy")

	if v, ok := d.GetOk("onboard10gbit_lom"); ok {
		x := (v.(string))
		o.SetOnboard10gbitLom(x)
	}

	if v, ok := d.GetOk("onboard_gbit_lom"); ok {
		x := (v.(string))
		o.SetOnboardGbitLom(x)
	}

	if v, ok := d.GetOk("onboard_scu_storage_support"); ok {
		x := (v.(string))
		o.SetOnboardScuStorageSupport(x)
	}

	if v, ok := d.GetOk("onboard_scu_storage_sw_stack"); ok {
		x := (v.(string))
		o.SetOnboardScuStorageSwStack(x)
	}

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if v, ok := d.GetOk("os_boot_watchdog_timer"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimer(x)
	}

	if v, ok := d.GetOk("os_boot_watchdog_timer_policy"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimerPolicy(x)
	}

	if v, ok := d.GetOk("os_boot_watchdog_timer_timeout"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimerTimeout(x)
	}

	if v, ok := d.GetOk("out_of_band_mgmt_port"); ok {
		x := (v.(string))
		o.SetOutOfBandMgmtPort(x)
	}

	if v, ok := d.GetOk("package_cstate_limit"); ok {
		x := (v.(string))
		o.SetPackageCstateLimit(x)
	}

	if v, ok := d.GetOk("partial_mirror_mode_config"); ok {
		x := (v.(string))
		o.SetPartialMirrorModeConfig(x)
	}

	if v, ok := d.GetOk("partial_mirror_percent"); ok {
		x := (v.(string))
		o.SetPartialMirrorPercent(x)
	}

	if v, ok := d.GetOk("partial_mirror_value1"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue1(x)
	}

	if v, ok := d.GetOk("partial_mirror_value2"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue2(x)
	}

	if v, ok := d.GetOk("partial_mirror_value3"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue3(x)
	}

	if v, ok := d.GetOk("partial_mirror_value4"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue4(x)
	}

	if v, ok := d.GetOk("patrol_scrub"); ok {
		x := (v.(string))
		o.SetPatrolScrub(x)
	}

	if v, ok := d.GetOk("patrol_scrub_duration"); ok {
		x := (v.(string))
		o.SetPatrolScrubDuration(x)
	}

	if v, ok := d.GetOk("pc_ie_ras_support"); ok {
		x := (v.(string))
		o.SetPcIeRasSupport(x)
	}

	if v, ok := d.GetOk("pc_ie_ssd_hot_plug_support"); ok {
		x := (v.(string))
		o.SetPcIeSsdHotPlugSupport(x)
	}

	if v, ok := d.GetOk("pch_usb30mode"); ok {
		x := (v.(string))
		o.SetPchUsb30mode(x)
	}

	if v, ok := d.GetOk("pci_option_ro_ms"); ok {
		x := (v.(string))
		o.SetPciOptionRoMs(x)
	}

	if v, ok := d.GetOk("pci_rom_clp"); ok {
		x := (v.(string))
		o.SetPciRomClp(x)
	}

	if v, ok := d.GetOk("pop_support"); ok {
		x := (v.(string))
		o.SetPopSupport(x)
	}

	if v, ok := d.GetOk("post_error_pause"); ok {
		x := (v.(string))
		o.SetPostErrorPause(x)
	}

	if v, ok := d.GetOk("processor_c1e"); ok {
		x := (v.(string))
		o.SetProcessorC1e(x)
	}

	if v, ok := d.GetOk("processor_c3report"); ok {
		x := (v.(string))
		o.SetProcessorC3report(x)
	}

	if v, ok := d.GetOk("processor_c6report"); ok {
		x := (v.(string))
		o.SetProcessorC6report(x)
	}

	if v, ok := d.GetOk("processor_cstate"); ok {
		x := (v.(string))
		o.SetProcessorCstate(x)
	}

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsPolicyAbstractConfigProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetProfiles(x)
		}
	}

	if v, ok := d.GetOk("psata"); ok {
		x := (v.(string))
		o.SetPsata(x)
	}

	if v, ok := d.GetOk("pstate_coord_type"); ok {
		x := (v.(string))
		o.SetPstateCoordType(x)
	}

	if v, ok := d.GetOk("putty_key_pad"); ok {
		x := (v.(string))
		o.SetPuttyKeyPad(x)
	}

	if v, ok := d.GetOk("pwr_perf_tuning"); ok {
		x := (v.(string))
		o.SetPwrPerfTuning(x)
	}

	if v, ok := d.GetOk("qpi_link_frequency"); ok {
		x := (v.(string))
		o.SetQpiLinkFrequency(x)
	}

	if v, ok := d.GetOk("qpi_snoop_mode"); ok {
		x := (v.(string))
		o.SetQpiSnoopMode(x)
	}

	if v, ok := d.GetOk("rank_inter_leave"); ok {
		x := (v.(string))
		o.SetRankInterLeave(x)
	}

	if v, ok := d.GetOk("redirection_after_post"); ok {
		x := (v.(string))
		o.SetRedirectionAfterPost(x)
	}

	if v, ok := d.GetOk("sata_mode_select"); ok {
		x := (v.(string))
		o.SetSataModeSelect(x)
	}

	if v, ok := d.GetOk("select_memory_ras_configuration"); ok {
		x := (v.(string))
		o.SetSelectMemoryRasConfiguration(x)
	}

	if v, ok := d.GetOk("select_ppr_type"); ok {
		x := (v.(string))
		o.SetSelectPprType(x)
	}

	if v, ok := d.GetOk("serial_port_aenable"); ok {
		x := (v.(string))
		o.SetSerialPortAenable(x)
	}

	if v, ok := d.GetOk("single_pctl_enable"); ok {
		x := (v.(string))
		o.SetSinglePctlEnable(x)
	}

	if v, ok := d.GetOk("slot10link_speed"); ok {
		x := (v.(string))
		o.SetSlot10linkSpeed(x)
	}

	if v, ok := d.GetOk("slot10state"); ok {
		x := (v.(string))
		o.SetSlot10state(x)
	}

	if v, ok := d.GetOk("slot11link_speed"); ok {
		x := (v.(string))
		o.SetSlot11linkSpeed(x)
	}

	if v, ok := d.GetOk("slot11state"); ok {
		x := (v.(string))
		o.SetSlot11state(x)
	}

	if v, ok := d.GetOk("slot12link_speed"); ok {
		x := (v.(string))
		o.SetSlot12linkSpeed(x)
	}

	if v, ok := d.GetOk("slot12state"); ok {
		x := (v.(string))
		o.SetSlot12state(x)
	}

	if v, ok := d.GetOk("slot13state"); ok {
		x := (v.(string))
		o.SetSlot13state(x)
	}

	if v, ok := d.GetOk("slot14state"); ok {
		x := (v.(string))
		o.SetSlot14state(x)
	}

	if v, ok := d.GetOk("slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlot1linkSpeed(x)
	}

	if v, ok := d.GetOk("slot1state"); ok {
		x := (v.(string))
		o.SetSlot1state(x)
	}

	if v, ok := d.GetOk("slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlot2linkSpeed(x)
	}

	if v, ok := d.GetOk("slot2state"); ok {
		x := (v.(string))
		o.SetSlot2state(x)
	}

	if v, ok := d.GetOk("slot3link_speed"); ok {
		x := (v.(string))
		o.SetSlot3linkSpeed(x)
	}

	if v, ok := d.GetOk("slot3state"); ok {
		x := (v.(string))
		o.SetSlot3state(x)
	}

	if v, ok := d.GetOk("slot4link_speed"); ok {
		x := (v.(string))
		o.SetSlot4linkSpeed(x)
	}

	if v, ok := d.GetOk("slot4state"); ok {
		x := (v.(string))
		o.SetSlot4state(x)
	}

	if v, ok := d.GetOk("slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlot5linkSpeed(x)
	}

	if v, ok := d.GetOk("slot5state"); ok {
		x := (v.(string))
		o.SetSlot5state(x)
	}

	if v, ok := d.GetOk("slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlot6linkSpeed(x)
	}

	if v, ok := d.GetOk("slot6state"); ok {
		x := (v.(string))
		o.SetSlot6state(x)
	}

	if v, ok := d.GetOk("slot7link_speed"); ok {
		x := (v.(string))
		o.SetSlot7linkSpeed(x)
	}

	if v, ok := d.GetOk("slot7state"); ok {
		x := (v.(string))
		o.SetSlot7state(x)
	}

	if v, ok := d.GetOk("slot8link_speed"); ok {
		x := (v.(string))
		o.SetSlot8linkSpeed(x)
	}

	if v, ok := d.GetOk("slot8state"); ok {
		x := (v.(string))
		o.SetSlot8state(x)
	}

	if v, ok := d.GetOk("slot9link_speed"); ok {
		x := (v.(string))
		o.SetSlot9linkSpeed(x)
	}

	if v, ok := d.GetOk("slot9state"); ok {
		x := (v.(string))
		o.SetSlot9state(x)
	}

	if v, ok := d.GetOk("slot_flom_link_speed"); ok {
		x := (v.(string))
		o.SetSlotFlomLinkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme1linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme2linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontSlot5linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontSlot6linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_gpu1state"); ok {
		x := (v.(string))
		o.SetSlotGpu1state(x)
	}

	if v, ok := d.GetOk("slot_gpu2state"); ok {
		x := (v.(string))
		o.SetSlotGpu2state(x)
	}

	if v, ok := d.GetOk("slot_gpu3state"); ok {
		x := (v.(string))
		o.SetSlotGpu3state(x)
	}

	if v, ok := d.GetOk("slot_gpu4state"); ok {
		x := (v.(string))
		o.SetSlotGpu4state(x)
	}

	if v, ok := d.GetOk("slot_gpu5state"); ok {
		x := (v.(string))
		o.SetSlotGpu5state(x)
	}

	if v, ok := d.GetOk("slot_gpu6state"); ok {
		x := (v.(string))
		o.SetSlotGpu6state(x)
	}

	if v, ok := d.GetOk("slot_gpu7state"); ok {
		x := (v.(string))
		o.SetSlotGpu7state(x)
	}

	if v, ok := d.GetOk("slot_gpu8state"); ok {
		x := (v.(string))
		o.SetSlotGpu8state(x)
	}

	if v, ok := d.GetOk("slot_hba_link_speed"); ok {
		x := (v.(string))
		o.SetSlotHbaLinkSpeed(x)
	}

	if v, ok := d.GetOk("slot_hba_state"); ok {
		x := (v.(string))
		o.SetSlotHbaState(x)
	}

	if v, ok := d.GetOk("slot_lom1link"); ok {
		x := (v.(string))
		o.SetSlotLom1link(x)
	}

	if v, ok := d.GetOk("slot_lom2link"); ok {
		x := (v.(string))
		o.SetSlotLom2link(x)
	}

	if v, ok := d.GetOk("slot_mezz_state"); ok {
		x := (v.(string))
		o.SetSlotMezzState(x)
	}

	if v, ok := d.GetOk("slot_mlom_link_speed"); ok {
		x := (v.(string))
		o.SetSlotMlomLinkSpeed(x)
	}

	if v, ok := d.GetOk("slot_mlom_state"); ok {
		x := (v.(string))
		o.SetSlotMlomState(x)
	}

	if v, ok := d.GetOk("slot_mraid_link_speed"); ok {
		x := (v.(string))
		o.SetSlotMraidLinkSpeed(x)
	}

	if v, ok := d.GetOk("slot_mraid_state"); ok {
		x := (v.(string))
		o.SetSlotMraidState(x)
	}

	if v, ok := d.GetOk("slot_n10state"); ok {
		x := (v.(string))
		o.SetSlotN10state(x)
	}

	if v, ok := d.GetOk("slot_n11state"); ok {
		x := (v.(string))
		o.SetSlotN11state(x)
	}

	if v, ok := d.GetOk("slot_n12state"); ok {
		x := (v.(string))
		o.SetSlotN12state(x)
	}

	if v, ok := d.GetOk("slot_n13state"); ok {
		x := (v.(string))
		o.SetSlotN13state(x)
	}

	if v, ok := d.GetOk("slot_n14state"); ok {
		x := (v.(string))
		o.SetSlotN14state(x)
	}

	if v, ok := d.GetOk("slot_n15state"); ok {
		x := (v.(string))
		o.SetSlotN15state(x)
	}

	if v, ok := d.GetOk("slot_n16state"); ok {
		x := (v.(string))
		o.SetSlotN16state(x)
	}

	if v, ok := d.GetOk("slot_n17state"); ok {
		x := (v.(string))
		o.SetSlotN17state(x)
	}

	if v, ok := d.GetOk("slot_n18state"); ok {
		x := (v.(string))
		o.SetSlotN18state(x)
	}

	if v, ok := d.GetOk("slot_n19state"); ok {
		x := (v.(string))
		o.SetSlotN19state(x)
	}

	if v, ok := d.GetOk("slot_n1state"); ok {
		x := (v.(string))
		o.SetSlotN1state(x)
	}

	if v, ok := d.GetOk("slot_n20state"); ok {
		x := (v.(string))
		o.SetSlotN20state(x)
	}

	if v, ok := d.GetOk("slot_n21state"); ok {
		x := (v.(string))
		o.SetSlotN21state(x)
	}

	if v, ok := d.GetOk("slot_n22state"); ok {
		x := (v.(string))
		o.SetSlotN22state(x)
	}

	if v, ok := d.GetOk("slot_n23state"); ok {
		x := (v.(string))
		o.SetSlotN23state(x)
	}

	if v, ok := d.GetOk("slot_n24state"); ok {
		x := (v.(string))
		o.SetSlotN24state(x)
	}

	if v, ok := d.GetOk("slot_n2state"); ok {
		x := (v.(string))
		o.SetSlotN2state(x)
	}

	if v, ok := d.GetOk("slot_n3state"); ok {
		x := (v.(string))
		o.SetSlotN3state(x)
	}

	if v, ok := d.GetOk("slot_n4state"); ok {
		x := (v.(string))
		o.SetSlotN4state(x)
	}

	if v, ok := d.GetOk("slot_n5state"); ok {
		x := (v.(string))
		o.SetSlotN5state(x)
	}

	if v, ok := d.GetOk("slot_n6state"); ok {
		x := (v.(string))
		o.SetSlotN6state(x)
	}

	if v, ok := d.GetOk("slot_n7state"); ok {
		x := (v.(string))
		o.SetSlotN7state(x)
	}

	if v, ok := d.GetOk("slot_n8state"); ok {
		x := (v.(string))
		o.SetSlotN8state(x)
	}

	if v, ok := d.GetOk("slot_n9state"); ok {
		x := (v.(string))
		o.SetSlotN9state(x)
	}

	if v, ok := d.GetOk("slot_raid_link_speed"); ok {
		x := (v.(string))
		o.SetSlotRaidLinkSpeed(x)
	}

	if v, ok := d.GetOk("slot_raid_state"); ok {
		x := (v.(string))
		o.SetSlotRaidState(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme1linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme1state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme1state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme2linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme2state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme2state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme3state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme3state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme4state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme4state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme5state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme5state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme6state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme6state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme7state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme7state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme8state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme8state(x)
	}

	if v, ok := d.GetOk("slot_riser1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_riser1slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot1linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_riser1slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot2linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_riser1slot3link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot3linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_riser2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_riser2slot4link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot4linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_riser2slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot5linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_riser2slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot6linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_sas_state"); ok {
		x := (v.(string))
		o.SetSlotSasState(x)
	}

	if v, ok := d.GetOk("slot_ssd_slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlotSsdSlot1linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_ssd_slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlotSsdSlot2linkSpeed(x)
	}

	if v, ok := d.GetOk("smee"); ok {
		x := (v.(string))
		o.SetSmee(x)
	}

	if v, ok := d.GetOk("smt_mode"); ok {
		x := (v.(string))
		o.SetSmtMode(x)
	}

	if v, ok := d.GetOk("snc"); ok {
		x := (v.(string))
		o.SetSnc(x)
	}

	if v, ok := d.GetOk("sparing_mode"); ok {
		x := (v.(string))
		o.SetSparingMode(x)
	}

	if v, ok := d.GetOk("sr_iov"); ok {
		x := (v.(string))
		o.SetSrIov(x)
	}

	if v, ok := d.GetOk("streamer_prefetch"); ok {
		x := (v.(string))
		o.SetStreamerPrefetch(x)
	}

	if v, ok := d.GetOk("svm_mode"); ok {
		x := (v.(string))
		o.SetSvmMode(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("terminal_type"); ok {
		x := (v.(string))
		o.SetTerminalType(x)
	}

	if v, ok := d.GetOk("tpm_control"); ok {
		x := (v.(string))
		o.SetTpmControl(x)
	}

	if v, ok := d.GetOk("tpm_support"); ok {
		x := (v.(string))
		o.SetTpmSupport(x)
	}

	if v, ok := d.GetOk("txt_support"); ok {
		x := (v.(string))
		o.SetTxtSupport(x)
	}

	if v, ok := d.GetOk("ucsm_boot_order_rule"); ok {
		x := (v.(string))
		o.SetUcsmBootOrderRule(x)
	}

	if v, ok := d.GetOk("usb_emul6064"); ok {
		x := (v.(string))
		o.SetUsbEmul6064(x)
	}

	if v, ok := d.GetOk("usb_port_front"); ok {
		x := (v.(string))
		o.SetUsbPortFront(x)
	}

	if v, ok := d.GetOk("usb_port_internal"); ok {
		x := (v.(string))
		o.SetUsbPortInternal(x)
	}

	if v, ok := d.GetOk("usb_port_kvm"); ok {
		x := (v.(string))
		o.SetUsbPortKvm(x)
	}

	if v, ok := d.GetOk("usb_port_rear"); ok {
		x := (v.(string))
		o.SetUsbPortRear(x)
	}

	if v, ok := d.GetOk("usb_port_sd_card"); ok {
		x := (v.(string))
		o.SetUsbPortSdCard(x)
	}

	if v, ok := d.GetOk("usb_port_vmedia"); ok {
		x := (v.(string))
		o.SetUsbPortVmedia(x)
	}

	if v, ok := d.GetOk("usb_xhci_support"); ok {
		x := (v.(string))
		o.SetUsbXhciSupport(x)
	}

	if v, ok := d.GetOk("vga_priority"); ok {
		x := (v.(string))
		o.SetVgaPriority(x)
	}

	if v, ok := d.GetOk("vmd_enable"); ok {
		x := (v.(string))
		o.SetVmdEnable(x)
	}

	if v, ok := d.GetOk("work_load_config"); ok {
		x := (v.(string))
		o.SetWorkLoadConfig(x)
	}

	if v, ok := d.GetOk("xpt_prefetch"); ok {
		x := (v.(string))
		o.SetXptPrefetch(x)
	}

	r := conn.ApiClient.BiosApi.CreateBiosPolicy(conn.ctx).BiosPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceBiosPolicyRead(d, meta)
}
func detachBiosPolicyProfiles(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewBiosPolicyWithDefaults()
	o.SetClassId("bios.Policy")
	o.SetObjectType("bios.Policy")
	o.SetProfiles([]models.PolicyAbstractConfigProfileRelationship{})

	r := conn.ApiClient.BiosApi.UpdateBiosPolicy(conn.ctx, d.Id()).BiosPolicy(*o)
	_, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while creating: %s", err.Error())
	}
	return err
}

func resourceBiosPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.BiosApi.GetBiosPolicyByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		return fmt.Errorf("error in unmarshaling model for read Error: %s", err.Error())
	}

	if err := d.Set("acs_control_gpu1state", (s.GetAcsControlGpu1state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu1state: %+v", err)
	}

	if err := d.Set("acs_control_gpu2state", (s.GetAcsControlGpu2state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu2state: %+v", err)
	}

	if err := d.Set("acs_control_gpu3state", (s.GetAcsControlGpu3state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu3state: %+v", err)
	}

	if err := d.Set("acs_control_gpu4state", (s.GetAcsControlGpu4state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu4state: %+v", err)
	}

	if err := d.Set("acs_control_gpu5state", (s.GetAcsControlGpu5state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu5state: %+v", err)
	}

	if err := d.Set("acs_control_gpu6state", (s.GetAcsControlGpu6state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu6state: %+v", err)
	}

	if err := d.Set("acs_control_gpu7state", (s.GetAcsControlGpu7state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu7state: %+v", err)
	}

	if err := d.Set("acs_control_gpu8state", (s.GetAcsControlGpu8state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlGpu8state: %+v", err)
	}

	if err := d.Set("acs_control_slot11state", (s.GetAcsControlSlot11state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlSlot11state: %+v", err)
	}

	if err := d.Set("acs_control_slot12state", (s.GetAcsControlSlot12state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlSlot12state: %+v", err)
	}

	if err := d.Set("acs_control_slot13state", (s.GetAcsControlSlot13state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlSlot13state: %+v", err)
	}

	if err := d.Set("acs_control_slot14state", (s.GetAcsControlSlot14state())); err != nil {
		return fmt.Errorf("error occurred while setting property AcsControlSlot14state: %+v", err)
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
	}

	if err := d.Set("adjacent_cache_line_prefetch", (s.GetAdjacentCacheLinePrefetch())); err != nil {
		return fmt.Errorf("error occurred while setting property AdjacentCacheLinePrefetch: %+v", err)
	}

	if err := d.Set("all_usb_devices", (s.GetAllUsbDevices())); err != nil {
		return fmt.Errorf("error occurred while setting property AllUsbDevices: %+v", err)
	}

	if err := d.Set("altitude", (s.GetAltitude())); err != nil {
		return fmt.Errorf("error occurred while setting property Altitude: %+v", err)
	}

	if err := d.Set("aspm_support", (s.GetAspmSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property AspmSupport: %+v", err)
	}

	if err := d.Set("assert_nmi_on_perr", (s.GetAssertNmiOnPerr())); err != nil {
		return fmt.Errorf("error occurred while setting property AssertNmiOnPerr: %+v", err)
	}

	if err := d.Set("assert_nmi_on_serr", (s.GetAssertNmiOnSerr())); err != nil {
		return fmt.Errorf("error occurred while setting property AssertNmiOnSerr: %+v", err)
	}

	if err := d.Set("auto_cc_state", (s.GetAutoCcState())); err != nil {
		return fmt.Errorf("error occurred while setting property AutoCcState: %+v", err)
	}

	if err := d.Set("autonumous_cstate_enable", (s.GetAutonumousCstateEnable())); err != nil {
		return fmt.Errorf("error occurred while setting property AutonumousCstateEnable: %+v", err)
	}

	if err := d.Set("baud_rate", (s.GetBaudRate())); err != nil {
		return fmt.Errorf("error occurred while setting property BaudRate: %+v", err)
	}

	if err := d.Set("bme_dma_mitigation", (s.GetBmeDmaMitigation())); err != nil {
		return fmt.Errorf("error occurred while setting property BmeDmaMitigation: %+v", err)
	}

	if err := d.Set("boot_option_num_retry", (s.GetBootOptionNumRetry())); err != nil {
		return fmt.Errorf("error occurred while setting property BootOptionNumRetry: %+v", err)
	}

	if err := d.Set("boot_option_re_cool_down", (s.GetBootOptionReCoolDown())); err != nil {
		return fmt.Errorf("error occurred while setting property BootOptionReCoolDown: %+v", err)
	}

	if err := d.Set("boot_option_retry", (s.GetBootOptionRetry())); err != nil {
		return fmt.Errorf("error occurred while setting property BootOptionRetry: %+v", err)
	}

	if err := d.Set("boot_performance_mode", (s.GetBootPerformanceMode())); err != nil {
		return fmt.Errorf("error occurred while setting property BootPerformanceMode: %+v", err)
	}

	if err := d.Set("cbs_cmn_cpu_cpb", (s.GetCbsCmnCpuCpb())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnCpuCpb: %+v", err)
	}

	if err := d.Set("cbs_cmn_cpu_gen_downcore_ctrl", (s.GetCbsCmnCpuGenDowncoreCtrl())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnCpuGenDowncoreCtrl: %+v", err)
	}

	if err := d.Set("cbs_cmn_cpu_global_cstate_ctrl", (s.GetCbsCmnCpuGlobalCstateCtrl())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnCpuGlobalCstateCtrl: %+v", err)
	}

	if err := d.Set("cbs_cmn_cpu_l1stream_hw_prefetcher", (s.GetCbsCmnCpuL1streamHwPrefetcher())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnCpuL1streamHwPrefetcher: %+v", err)
	}

	if err := d.Set("cbs_cmn_cpu_l2stream_hw_prefetcher", (s.GetCbsCmnCpuL2streamHwPrefetcher())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnCpuL2streamHwPrefetcher: %+v", err)
	}

	if err := d.Set("cbs_cmn_determinism_slider", (s.GetCbsCmnDeterminismSlider())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnDeterminismSlider: %+v", err)
	}

	if err := d.Set("cbs_cmn_gnb_nb_iommu", (s.GetCbsCmnGnbNbIommu())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnGnbNbIommu: %+v", err)
	}

	if err := d.Set("cbs_cmn_mem_ctrl_bank_group_swap_ddr4", (s.GetCbsCmnMemCtrlBankGroupSwapDdr4())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnMemCtrlBankGroupSwapDdr4: %+v", err)
	}

	if err := d.Set("cbs_cmn_mem_map_bank_interleave_ddr4", (s.GetCbsCmnMemMapBankInterleaveDdr4())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmnMemMapBankInterleaveDdr4: %+v", err)
	}

	if err := d.Set("cbs_cmnc_tdp_ctl", (s.GetCbsCmncTdpCtl())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsCmncTdpCtl: %+v", err)
	}

	if err := d.Set("cbs_df_cmn_mem_intlv", (s.GetCbsDfCmnMemIntlv())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsDfCmnMemIntlv: %+v", err)
	}

	if err := d.Set("cbs_df_cmn_mem_intlv_size", (s.GetCbsDfCmnMemIntlvSize())); err != nil {
		return fmt.Errorf("error occurred while setting property CbsDfCmnMemIntlvSize: %+v", err)
	}

	if err := d.Set("cdn_enable", (s.GetCdnEnable())); err != nil {
		return fmt.Errorf("error occurred while setting property CdnEnable: %+v", err)
	}

	if err := d.Set("cdn_support", (s.GetCdnSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property CdnSupport: %+v", err)
	}

	if err := d.Set("channel_inter_leave", (s.GetChannelInterLeave())); err != nil {
		return fmt.Errorf("error occurred while setting property ChannelInterLeave: %+v", err)
	}

	if err := d.Set("cisco_adaptive_mem_training", (s.GetCiscoAdaptiveMemTraining())); err != nil {
		return fmt.Errorf("error occurred while setting property CiscoAdaptiveMemTraining: %+v", err)
	}

	if err := d.Set("cisco_debug_level", (s.GetCiscoDebugLevel())); err != nil {
		return fmt.Errorf("error occurred while setting property CiscoDebugLevel: %+v", err)
	}

	if err := d.Set("cisco_oprom_launch_optimization", (s.GetCiscoOpromLaunchOptimization())); err != nil {
		return fmt.Errorf("error occurred while setting property CiscoOpromLaunchOptimization: %+v", err)
	}

	if err := d.Set("cke_low_policy", (s.GetCkeLowPolicy())); err != nil {
		return fmt.Errorf("error occurred while setting property CkeLowPolicy: %+v", err)
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
	}

	if err := d.Set("closed_loop_therm_throtl", (s.GetClosedLoopThermThrotl())); err != nil {
		return fmt.Errorf("error occurred while setting property ClosedLoopThermThrotl: %+v", err)
	}

	if err := d.Set("cmci_enable", (s.GetCmciEnable())); err != nil {
		return fmt.Errorf("error occurred while setting property CmciEnable: %+v", err)
	}

	if err := d.Set("config_tdp", (s.GetConfigTdp())); err != nil {
		return fmt.Errorf("error occurred while setting property ConfigTdp: %+v", err)
	}

	if err := d.Set("console_redirection", (s.GetConsoleRedirection())); err != nil {
		return fmt.Errorf("error occurred while setting property ConsoleRedirection: %+v", err)
	}

	if err := d.Set("core_multi_processing", (s.GetCoreMultiProcessing())); err != nil {
		return fmt.Errorf("error occurred while setting property CoreMultiProcessing: %+v", err)
	}

	if err := d.Set("cpu_energy_performance", (s.GetCpuEnergyPerformance())); err != nil {
		return fmt.Errorf("error occurred while setting property CpuEnergyPerformance: %+v", err)
	}

	if err := d.Set("cpu_frequency_floor", (s.GetCpuFrequencyFloor())); err != nil {
		return fmt.Errorf("error occurred while setting property CpuFrequencyFloor: %+v", err)
	}

	if err := d.Set("cpu_performance", (s.GetCpuPerformance())); err != nil {
		return fmt.Errorf("error occurred while setting property CpuPerformance: %+v", err)
	}

	if err := d.Set("cpu_power_management", (s.GetCpuPowerManagement())); err != nil {
		return fmt.Errorf("error occurred while setting property CpuPowerManagement: %+v", err)
	}

	if err := d.Set("dcpmm_firmware_downgrade", (s.GetDcpmmFirmwareDowngrade())); err != nil {
		return fmt.Errorf("error occurred while setting property DcpmmFirmwareDowngrade: %+v", err)
	}

	if err := d.Set("demand_scrub", (s.GetDemandScrub())); err != nil {
		return fmt.Errorf("error occurred while setting property DemandScrub: %+v", err)
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return fmt.Errorf("error occurred while setting property Description: %+v", err)
	}

	if err := d.Set("direct_cache_access", (s.GetDirectCacheAccess())); err != nil {
		return fmt.Errorf("error occurred while setting property DirectCacheAccess: %+v", err)
	}

	if err := d.Set("dram_clock_throttling", (s.GetDramClockThrottling())); err != nil {
		return fmt.Errorf("error occurred while setting property DramClockThrottling: %+v", err)
	}

	if err := d.Set("dram_refresh_rate", (s.GetDramRefreshRate())); err != nil {
		return fmt.Errorf("error occurred while setting property DramRefreshRate: %+v", err)
	}

	if err := d.Set("energy_efficient_turbo", (s.GetEnergyEfficientTurbo())); err != nil {
		return fmt.Errorf("error occurred while setting property EnergyEfficientTurbo: %+v", err)
	}

	if err := d.Set("eng_perf_tuning", (s.GetEngPerfTuning())); err != nil {
		return fmt.Errorf("error occurred while setting property EngPerfTuning: %+v", err)
	}

	if err := d.Set("enhanced_intel_speed_step_tech", (s.GetEnhancedIntelSpeedStepTech())); err != nil {
		return fmt.Errorf("error occurred while setting property EnhancedIntelSpeedStepTech: %+v", err)
	}

	if err := d.Set("epp_profile", (s.GetEppProfile())); err != nil {
		return fmt.Errorf("error occurred while setting property EppProfile: %+v", err)
	}

	if err := d.Set("execute_disable_bit", (s.GetExecuteDisableBit())); err != nil {
		return fmt.Errorf("error occurred while setting property ExecuteDisableBit: %+v", err)
	}

	if err := d.Set("extended_apic", (s.GetExtendedApic())); err != nil {
		return fmt.Errorf("error occurred while setting property ExtendedApic: %+v", err)
	}

	if err := d.Set("flow_control", (s.GetFlowControl())); err != nil {
		return fmt.Errorf("error occurred while setting property FlowControl: %+v", err)
	}

	if err := d.Set("frb2enable", (s.GetFrb2enable())); err != nil {
		return fmt.Errorf("error occurred while setting property Frb2enable: %+v", err)
	}

	if err := d.Set("hardware_prefetch", (s.GetHardwarePrefetch())); err != nil {
		return fmt.Errorf("error occurred while setting property HardwarePrefetch: %+v", err)
	}

	if err := d.Set("hwpm_enable", (s.GetHwpmEnable())); err != nil {
		return fmt.Errorf("error occurred while setting property HwpmEnable: %+v", err)
	}

	if err := d.Set("imc_interleave", (s.GetImcInterleave())); err != nil {
		return fmt.Errorf("error occurred while setting property ImcInterleave: %+v", err)
	}

	if err := d.Set("intel_hyper_threading_tech", (s.GetIntelHyperThreadingTech())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelHyperThreadingTech: %+v", err)
	}

	if err := d.Set("intel_speed_select", (s.GetIntelSpeedSelect())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelSpeedSelect: %+v", err)
	}

	if err := d.Set("intel_turbo_boost_tech", (s.GetIntelTurboBoostTech())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelTurboBoostTech: %+v", err)
	}

	if err := d.Set("intel_virtualization_technology", (s.GetIntelVirtualizationTechnology())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelVirtualizationTechnology: %+v", err)
	}

	if err := d.Set("intel_vt_for_directed_io", (s.GetIntelVtForDirectedIo())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelVtForDirectedIo: %+v", err)
	}

	if err := d.Set("intel_vtd_coherency_support", (s.GetIntelVtdCoherencySupport())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelVtdCoherencySupport: %+v", err)
	}

	if err := d.Set("intel_vtd_interrupt_remapping", (s.GetIntelVtdInterruptRemapping())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelVtdInterruptRemapping: %+v", err)
	}

	if err := d.Set("intel_vtd_pass_through_dma_support", (s.GetIntelVtdPassThroughDmaSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelVtdPassThroughDmaSupport: %+v", err)
	}

	if err := d.Set("intel_vtdats_support", (s.GetIntelVtdatsSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property IntelVtdatsSupport: %+v", err)
	}

	if err := d.Set("ioh_error_enable", (s.GetIohErrorEnable())); err != nil {
		return fmt.Errorf("error occurred while setting property IohErrorEnable: %+v", err)
	}

	if err := d.Set("ioh_resource", (s.GetIohResource())); err != nil {
		return fmt.Errorf("error occurred while setting property IohResource: %+v", err)
	}

	if err := d.Set("ip_prefetch", (s.GetIpPrefetch())); err != nil {
		return fmt.Errorf("error occurred while setting property IpPrefetch: %+v", err)
	}

	if err := d.Set("ipv4pxe", (s.GetIpv4pxe())); err != nil {
		return fmt.Errorf("error occurred while setting property Ipv4pxe: %+v", err)
	}

	if err := d.Set("ipv6pxe", (s.GetIpv6pxe())); err != nil {
		return fmt.Errorf("error occurred while setting property Ipv6pxe: %+v", err)
	}

	if err := d.Set("kti_prefetch", (s.GetKtiPrefetch())); err != nil {
		return fmt.Errorf("error occurred while setting property KtiPrefetch: %+v", err)
	}

	if err := d.Set("legacy_os_redirection", (s.GetLegacyOsRedirection())); err != nil {
		return fmt.Errorf("error occurred while setting property LegacyOsRedirection: %+v", err)
	}

	if err := d.Set("legacy_usb_support", (s.GetLegacyUsbSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property LegacyUsbSupport: %+v", err)
	}

	if err := d.Set("llc_prefetch", (s.GetLlcPrefetch())); err != nil {
		return fmt.Errorf("error occurred while setting property LlcPrefetch: %+v", err)
	}

	if err := d.Set("lom_port0state", (s.GetLomPort0state())); err != nil {
		return fmt.Errorf("error occurred while setting property LomPort0state: %+v", err)
	}

	if err := d.Set("lom_port1state", (s.GetLomPort1state())); err != nil {
		return fmt.Errorf("error occurred while setting property LomPort1state: %+v", err)
	}

	if err := d.Set("lom_port2state", (s.GetLomPort2state())); err != nil {
		return fmt.Errorf("error occurred while setting property LomPort2state: %+v", err)
	}

	if err := d.Set("lom_port3state", (s.GetLomPort3state())); err != nil {
		return fmt.Errorf("error occurred while setting property LomPort3state: %+v", err)
	}

	if err := d.Set("lom_ports_all_state", (s.GetLomPortsAllState())); err != nil {
		return fmt.Errorf("error occurred while setting property LomPortsAllState: %+v", err)
	}

	if err := d.Set("lv_ddr_mode", (s.GetLvDdrMode())); err != nil {
		return fmt.Errorf("error occurred while setting property LvDdrMode: %+v", err)
	}

	if err := d.Set("make_device_non_bootable", (s.GetMakeDeviceNonBootable())); err != nil {
		return fmt.Errorf("error occurred while setting property MakeDeviceNonBootable: %+v", err)
	}

	if err := d.Set("memory_inter_leave", (s.GetMemoryInterLeave())); err != nil {
		return fmt.Errorf("error occurred while setting property MemoryInterLeave: %+v", err)
	}

	if err := d.Set("memory_mapped_io_above4gb", (s.GetMemoryMappedIoAbove4gb())); err != nil {
		return fmt.Errorf("error occurred while setting property MemoryMappedIoAbove4gb: %+v", err)
	}

	if err := d.Set("memory_size_limit", (s.GetMemorySizeLimit())); err != nil {
		return fmt.Errorf("error occurred while setting property MemorySizeLimit: %+v", err)
	}

	if err := d.Set("mirroring_mode", (s.GetMirroringMode())); err != nil {
		return fmt.Errorf("error occurred while setting property MirroringMode: %+v", err)
	}

	if err := d.Set("mmcfg_base", (s.GetMmcfgBase())); err != nil {
		return fmt.Errorf("error occurred while setting property MmcfgBase: %+v", err)
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return fmt.Errorf("error occurred while setting property Moid: %+v", err)
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return fmt.Errorf("error occurred while setting property Name: %+v", err)
	}

	if err := d.Set("network_stack", (s.GetNetworkStack())); err != nil {
		return fmt.Errorf("error occurred while setting property NetworkStack: %+v", err)
	}

	if err := d.Set("numa_optimized", (s.GetNumaOptimized())); err != nil {
		return fmt.Errorf("error occurred while setting property NumaOptimized: %+v", err)
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
	}

	if err := d.Set("onboard10gbit_lom", (s.GetOnboard10gbitLom())); err != nil {
		return fmt.Errorf("error occurred while setting property Onboard10gbitLom: %+v", err)
	}

	if err := d.Set("onboard_gbit_lom", (s.GetOnboardGbitLom())); err != nil {
		return fmt.Errorf("error occurred while setting property OnboardGbitLom: %+v", err)
	}

	if err := d.Set("onboard_scu_storage_support", (s.GetOnboardScuStorageSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property OnboardScuStorageSupport: %+v", err)
	}

	if err := d.Set("onboard_scu_storage_sw_stack", (s.GetOnboardScuStorageSwStack())); err != nil {
		return fmt.Errorf("error occurred while setting property OnboardScuStorageSwStack: %+v", err)
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Organization: %+v", err)
	}

	if err := d.Set("os_boot_watchdog_timer", (s.GetOsBootWatchdogTimer())); err != nil {
		return fmt.Errorf("error occurred while setting property OsBootWatchdogTimer: %+v", err)
	}

	if err := d.Set("os_boot_watchdog_timer_policy", (s.GetOsBootWatchdogTimerPolicy())); err != nil {
		return fmt.Errorf("error occurred while setting property OsBootWatchdogTimerPolicy: %+v", err)
	}

	if err := d.Set("os_boot_watchdog_timer_timeout", (s.GetOsBootWatchdogTimerTimeout())); err != nil {
		return fmt.Errorf("error occurred while setting property OsBootWatchdogTimerTimeout: %+v", err)
	}

	if err := d.Set("out_of_band_mgmt_port", (s.GetOutOfBandMgmtPort())); err != nil {
		return fmt.Errorf("error occurred while setting property OutOfBandMgmtPort: %+v", err)
	}

	if err := d.Set("package_cstate_limit", (s.GetPackageCstateLimit())); err != nil {
		return fmt.Errorf("error occurred while setting property PackageCstateLimit: %+v", err)
	}

	if err := d.Set("partial_mirror_mode_config", (s.GetPartialMirrorModeConfig())); err != nil {
		return fmt.Errorf("error occurred while setting property PartialMirrorModeConfig: %+v", err)
	}

	if err := d.Set("partial_mirror_percent", (s.GetPartialMirrorPercent())); err != nil {
		return fmt.Errorf("error occurred while setting property PartialMirrorPercent: %+v", err)
	}

	if err := d.Set("partial_mirror_value1", (s.GetPartialMirrorValue1())); err != nil {
		return fmt.Errorf("error occurred while setting property PartialMirrorValue1: %+v", err)
	}

	if err := d.Set("partial_mirror_value2", (s.GetPartialMirrorValue2())); err != nil {
		return fmt.Errorf("error occurred while setting property PartialMirrorValue2: %+v", err)
	}

	if err := d.Set("partial_mirror_value3", (s.GetPartialMirrorValue3())); err != nil {
		return fmt.Errorf("error occurred while setting property PartialMirrorValue3: %+v", err)
	}

	if err := d.Set("partial_mirror_value4", (s.GetPartialMirrorValue4())); err != nil {
		return fmt.Errorf("error occurred while setting property PartialMirrorValue4: %+v", err)
	}

	if err := d.Set("patrol_scrub", (s.GetPatrolScrub())); err != nil {
		return fmt.Errorf("error occurred while setting property PatrolScrub: %+v", err)
	}

	if err := d.Set("patrol_scrub_duration", (s.GetPatrolScrubDuration())); err != nil {
		return fmt.Errorf("error occurred while setting property PatrolScrubDuration: %+v", err)
	}

	if err := d.Set("pc_ie_ras_support", (s.GetPcIeRasSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property PcIeRasSupport: %+v", err)
	}

	if err := d.Set("pc_ie_ssd_hot_plug_support", (s.GetPcIeSsdHotPlugSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property PcIeSsdHotPlugSupport: %+v", err)
	}

	if err := d.Set("pch_usb30mode", (s.GetPchUsb30mode())); err != nil {
		return fmt.Errorf("error occurred while setting property PchUsb30mode: %+v", err)
	}

	if err := d.Set("pci_option_ro_ms", (s.GetPciOptionRoMs())); err != nil {
		return fmt.Errorf("error occurred while setting property PciOptionRoMs: %+v", err)
	}

	if err := d.Set("pci_rom_clp", (s.GetPciRomClp())); err != nil {
		return fmt.Errorf("error occurred while setting property PciRomClp: %+v", err)
	}

	if err := d.Set("pop_support", (s.GetPopSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property PopSupport: %+v", err)
	}

	if err := d.Set("post_error_pause", (s.GetPostErrorPause())); err != nil {
		return fmt.Errorf("error occurred while setting property PostErrorPause: %+v", err)
	}

	if err := d.Set("processor_c1e", (s.GetProcessorC1e())); err != nil {
		return fmt.Errorf("error occurred while setting property ProcessorC1e: %+v", err)
	}

	if err := d.Set("processor_c3report", (s.GetProcessorC3report())); err != nil {
		return fmt.Errorf("error occurred while setting property ProcessorC3report: %+v", err)
	}

	if err := d.Set("processor_c6report", (s.GetProcessorC6report())); err != nil {
		return fmt.Errorf("error occurred while setting property ProcessorC6report: %+v", err)
	}

	if err := d.Set("processor_cstate", (s.GetProcessorCstate())); err != nil {
		return fmt.Errorf("error occurred while setting property ProcessorCstate: %+v", err)
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Profiles: %+v", err)
	}

	if err := d.Set("psata", (s.GetPsata())); err != nil {
		return fmt.Errorf("error occurred while setting property Psata: %+v", err)
	}

	if err := d.Set("pstate_coord_type", (s.GetPstateCoordType())); err != nil {
		return fmt.Errorf("error occurred while setting property PstateCoordType: %+v", err)
	}

	if err := d.Set("putty_key_pad", (s.GetPuttyKeyPad())); err != nil {
		return fmt.Errorf("error occurred while setting property PuttyKeyPad: %+v", err)
	}

	if err := d.Set("pwr_perf_tuning", (s.GetPwrPerfTuning())); err != nil {
		return fmt.Errorf("error occurred while setting property PwrPerfTuning: %+v", err)
	}

	if err := d.Set("qpi_link_frequency", (s.GetQpiLinkFrequency())); err != nil {
		return fmt.Errorf("error occurred while setting property QpiLinkFrequency: %+v", err)
	}

	if err := d.Set("qpi_snoop_mode", (s.GetQpiSnoopMode())); err != nil {
		return fmt.Errorf("error occurred while setting property QpiSnoopMode: %+v", err)
	}

	if err := d.Set("rank_inter_leave", (s.GetRankInterLeave())); err != nil {
		return fmt.Errorf("error occurred while setting property RankInterLeave: %+v", err)
	}

	if err := d.Set("redirection_after_post", (s.GetRedirectionAfterPost())); err != nil {
		return fmt.Errorf("error occurred while setting property RedirectionAfterPost: %+v", err)
	}

	if err := d.Set("sata_mode_select", (s.GetSataModeSelect())); err != nil {
		return fmt.Errorf("error occurred while setting property SataModeSelect: %+v", err)
	}

	if err := d.Set("select_memory_ras_configuration", (s.GetSelectMemoryRasConfiguration())); err != nil {
		return fmt.Errorf("error occurred while setting property SelectMemoryRasConfiguration: %+v", err)
	}

	if err := d.Set("select_ppr_type", (s.GetSelectPprType())); err != nil {
		return fmt.Errorf("error occurred while setting property SelectPprType: %+v", err)
	}

	if err := d.Set("serial_port_aenable", (s.GetSerialPortAenable())); err != nil {
		return fmt.Errorf("error occurred while setting property SerialPortAenable: %+v", err)
	}

	if err := d.Set("single_pctl_enable", (s.GetSinglePctlEnable())); err != nil {
		return fmt.Errorf("error occurred while setting property SinglePctlEnable: %+v", err)
	}

	if err := d.Set("slot10link_speed", (s.GetSlot10linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot10linkSpeed: %+v", err)
	}

	if err := d.Set("slot10state", (s.GetSlot10state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot10state: %+v", err)
	}

	if err := d.Set("slot11link_speed", (s.GetSlot11linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot11linkSpeed: %+v", err)
	}

	if err := d.Set("slot11state", (s.GetSlot11state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot11state: %+v", err)
	}

	if err := d.Set("slot12link_speed", (s.GetSlot12linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot12linkSpeed: %+v", err)
	}

	if err := d.Set("slot12state", (s.GetSlot12state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot12state: %+v", err)
	}

	if err := d.Set("slot13state", (s.GetSlot13state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot13state: %+v", err)
	}

	if err := d.Set("slot14state", (s.GetSlot14state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot14state: %+v", err)
	}

	if err := d.Set("slot1link_speed", (s.GetSlot1linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot1linkSpeed: %+v", err)
	}

	if err := d.Set("slot1state", (s.GetSlot1state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot1state: %+v", err)
	}

	if err := d.Set("slot2link_speed", (s.GetSlot2linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot2linkSpeed: %+v", err)
	}

	if err := d.Set("slot2state", (s.GetSlot2state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot2state: %+v", err)
	}

	if err := d.Set("slot3link_speed", (s.GetSlot3linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot3linkSpeed: %+v", err)
	}

	if err := d.Set("slot3state", (s.GetSlot3state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot3state: %+v", err)
	}

	if err := d.Set("slot4link_speed", (s.GetSlot4linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot4linkSpeed: %+v", err)
	}

	if err := d.Set("slot4state", (s.GetSlot4state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot4state: %+v", err)
	}

	if err := d.Set("slot5link_speed", (s.GetSlot5linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot5linkSpeed: %+v", err)
	}

	if err := d.Set("slot5state", (s.GetSlot5state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot5state: %+v", err)
	}

	if err := d.Set("slot6link_speed", (s.GetSlot6linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot6linkSpeed: %+v", err)
	}

	if err := d.Set("slot6state", (s.GetSlot6state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot6state: %+v", err)
	}

	if err := d.Set("slot7link_speed", (s.GetSlot7linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot7linkSpeed: %+v", err)
	}

	if err := d.Set("slot7state", (s.GetSlot7state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot7state: %+v", err)
	}

	if err := d.Set("slot8link_speed", (s.GetSlot8linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot8linkSpeed: %+v", err)
	}

	if err := d.Set("slot8state", (s.GetSlot8state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot8state: %+v", err)
	}

	if err := d.Set("slot9link_speed", (s.GetSlot9linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot9linkSpeed: %+v", err)
	}

	if err := d.Set("slot9state", (s.GetSlot9state())); err != nil {
		return fmt.Errorf("error occurred while setting property Slot9state: %+v", err)
	}

	if err := d.Set("slot_flom_link_speed", (s.GetSlotFlomLinkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotFlomLinkSpeed: %+v", err)
	}

	if err := d.Set("slot_front_nvme1link_speed", (s.GetSlotFrontNvme1linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotFrontNvme1linkSpeed: %+v", err)
	}

	if err := d.Set("slot_front_nvme2link_speed", (s.GetSlotFrontNvme2linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotFrontNvme2linkSpeed: %+v", err)
	}

	if err := d.Set("slot_front_slot5link_speed", (s.GetSlotFrontSlot5linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotFrontSlot5linkSpeed: %+v", err)
	}

	if err := d.Set("slot_front_slot6link_speed", (s.GetSlotFrontSlot6linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotFrontSlot6linkSpeed: %+v", err)
	}

	if err := d.Set("slot_gpu1state", (s.GetSlotGpu1state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu1state: %+v", err)
	}

	if err := d.Set("slot_gpu2state", (s.GetSlotGpu2state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu2state: %+v", err)
	}

	if err := d.Set("slot_gpu3state", (s.GetSlotGpu3state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu3state: %+v", err)
	}

	if err := d.Set("slot_gpu4state", (s.GetSlotGpu4state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu4state: %+v", err)
	}

	if err := d.Set("slot_gpu5state", (s.GetSlotGpu5state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu5state: %+v", err)
	}

	if err := d.Set("slot_gpu6state", (s.GetSlotGpu6state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu6state: %+v", err)
	}

	if err := d.Set("slot_gpu7state", (s.GetSlotGpu7state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu7state: %+v", err)
	}

	if err := d.Set("slot_gpu8state", (s.GetSlotGpu8state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotGpu8state: %+v", err)
	}

	if err := d.Set("slot_hba_link_speed", (s.GetSlotHbaLinkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotHbaLinkSpeed: %+v", err)
	}

	if err := d.Set("slot_hba_state", (s.GetSlotHbaState())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotHbaState: %+v", err)
	}

	if err := d.Set("slot_lom1link", (s.GetSlotLom1link())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotLom1link: %+v", err)
	}

	if err := d.Set("slot_lom2link", (s.GetSlotLom2link())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotLom2link: %+v", err)
	}

	if err := d.Set("slot_mezz_state", (s.GetSlotMezzState())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotMezzState: %+v", err)
	}

	if err := d.Set("slot_mlom_link_speed", (s.GetSlotMlomLinkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotMlomLinkSpeed: %+v", err)
	}

	if err := d.Set("slot_mlom_state", (s.GetSlotMlomState())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotMlomState: %+v", err)
	}

	if err := d.Set("slot_mraid_link_speed", (s.GetSlotMraidLinkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotMraidLinkSpeed: %+v", err)
	}

	if err := d.Set("slot_mraid_state", (s.GetSlotMraidState())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotMraidState: %+v", err)
	}

	if err := d.Set("slot_n10state", (s.GetSlotN10state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN10state: %+v", err)
	}

	if err := d.Set("slot_n11state", (s.GetSlotN11state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN11state: %+v", err)
	}

	if err := d.Set("slot_n12state", (s.GetSlotN12state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN12state: %+v", err)
	}

	if err := d.Set("slot_n13state", (s.GetSlotN13state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN13state: %+v", err)
	}

	if err := d.Set("slot_n14state", (s.GetSlotN14state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN14state: %+v", err)
	}

	if err := d.Set("slot_n15state", (s.GetSlotN15state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN15state: %+v", err)
	}

	if err := d.Set("slot_n16state", (s.GetSlotN16state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN16state: %+v", err)
	}

	if err := d.Set("slot_n17state", (s.GetSlotN17state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN17state: %+v", err)
	}

	if err := d.Set("slot_n18state", (s.GetSlotN18state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN18state: %+v", err)
	}

	if err := d.Set("slot_n19state", (s.GetSlotN19state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN19state: %+v", err)
	}

	if err := d.Set("slot_n1state", (s.GetSlotN1state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN1state: %+v", err)
	}

	if err := d.Set("slot_n20state", (s.GetSlotN20state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN20state: %+v", err)
	}

	if err := d.Set("slot_n21state", (s.GetSlotN21state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN21state: %+v", err)
	}

	if err := d.Set("slot_n22state", (s.GetSlotN22state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN22state: %+v", err)
	}

	if err := d.Set("slot_n23state", (s.GetSlotN23state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN23state: %+v", err)
	}

	if err := d.Set("slot_n24state", (s.GetSlotN24state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN24state: %+v", err)
	}

	if err := d.Set("slot_n2state", (s.GetSlotN2state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN2state: %+v", err)
	}

	if err := d.Set("slot_n3state", (s.GetSlotN3state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN3state: %+v", err)
	}

	if err := d.Set("slot_n4state", (s.GetSlotN4state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN4state: %+v", err)
	}

	if err := d.Set("slot_n5state", (s.GetSlotN5state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN5state: %+v", err)
	}

	if err := d.Set("slot_n6state", (s.GetSlotN6state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN6state: %+v", err)
	}

	if err := d.Set("slot_n7state", (s.GetSlotN7state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN7state: %+v", err)
	}

	if err := d.Set("slot_n8state", (s.GetSlotN8state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN8state: %+v", err)
	}

	if err := d.Set("slot_n9state", (s.GetSlotN9state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotN9state: %+v", err)
	}

	if err := d.Set("slot_raid_link_speed", (s.GetSlotRaidLinkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRaidLinkSpeed: %+v", err)
	}

	if err := d.Set("slot_raid_state", (s.GetSlotRaidState())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRaidState: %+v", err)
	}

	if err := d.Set("slot_rear_nvme1link_speed", (s.GetSlotRearNvme1linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme1linkSpeed: %+v", err)
	}

	if err := d.Set("slot_rear_nvme1state", (s.GetSlotRearNvme1state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme1state: %+v", err)
	}

	if err := d.Set("slot_rear_nvme2link_speed", (s.GetSlotRearNvme2linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme2linkSpeed: %+v", err)
	}

	if err := d.Set("slot_rear_nvme2state", (s.GetSlotRearNvme2state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme2state: %+v", err)
	}

	if err := d.Set("slot_rear_nvme3state", (s.GetSlotRearNvme3state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme3state: %+v", err)
	}

	if err := d.Set("slot_rear_nvme4state", (s.GetSlotRearNvme4state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme4state: %+v", err)
	}

	if err := d.Set("slot_rear_nvme5state", (s.GetSlotRearNvme5state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme5state: %+v", err)
	}

	if err := d.Set("slot_rear_nvme6state", (s.GetSlotRearNvme6state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme6state: %+v", err)
	}

	if err := d.Set("slot_rear_nvme7state", (s.GetSlotRearNvme7state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme7state: %+v", err)
	}

	if err := d.Set("slot_rear_nvme8state", (s.GetSlotRearNvme8state())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRearNvme8state: %+v", err)
	}

	if err := d.Set("slot_riser1link_speed", (s.GetSlotRiser1linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser1linkSpeed: %+v", err)
	}

	if err := d.Set("slot_riser1slot1link_speed", (s.GetSlotRiser1slot1linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser1slot1linkSpeed: %+v", err)
	}

	if err := d.Set("slot_riser1slot2link_speed", (s.GetSlotRiser1slot2linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser1slot2linkSpeed: %+v", err)
	}

	if err := d.Set("slot_riser1slot3link_speed", (s.GetSlotRiser1slot3linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser1slot3linkSpeed: %+v", err)
	}

	if err := d.Set("slot_riser2link_speed", (s.GetSlotRiser2linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser2linkSpeed: %+v", err)
	}

	if err := d.Set("slot_riser2slot4link_speed", (s.GetSlotRiser2slot4linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser2slot4linkSpeed: %+v", err)
	}

	if err := d.Set("slot_riser2slot5link_speed", (s.GetSlotRiser2slot5linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser2slot5linkSpeed: %+v", err)
	}

	if err := d.Set("slot_riser2slot6link_speed", (s.GetSlotRiser2slot6linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotRiser2slot6linkSpeed: %+v", err)
	}

	if err := d.Set("slot_sas_state", (s.GetSlotSasState())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotSasState: %+v", err)
	}

	if err := d.Set("slot_ssd_slot1link_speed", (s.GetSlotSsdSlot1linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotSsdSlot1linkSpeed: %+v", err)
	}

	if err := d.Set("slot_ssd_slot2link_speed", (s.GetSlotSsdSlot2linkSpeed())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotSsdSlot2linkSpeed: %+v", err)
	}

	if err := d.Set("smee", (s.GetSmee())); err != nil {
		return fmt.Errorf("error occurred while setting property Smee: %+v", err)
	}

	if err := d.Set("smt_mode", (s.GetSmtMode())); err != nil {
		return fmt.Errorf("error occurred while setting property SmtMode: %+v", err)
	}

	if err := d.Set("snc", (s.GetSnc())); err != nil {
		return fmt.Errorf("error occurred while setting property Snc: %+v", err)
	}

	if err := d.Set("sparing_mode", (s.GetSparingMode())); err != nil {
		return fmt.Errorf("error occurred while setting property SparingMode: %+v", err)
	}

	if err := d.Set("sr_iov", (s.GetSrIov())); err != nil {
		return fmt.Errorf("error occurred while setting property SrIov: %+v", err)
	}

	if err := d.Set("streamer_prefetch", (s.GetStreamerPrefetch())); err != nil {
		return fmt.Errorf("error occurred while setting property StreamerPrefetch: %+v", err)
	}

	if err := d.Set("svm_mode", (s.GetSvmMode())); err != nil {
		return fmt.Errorf("error occurred while setting property SvmMode: %+v", err)
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Tags: %+v", err)
	}

	if err := d.Set("terminal_type", (s.GetTerminalType())); err != nil {
		return fmt.Errorf("error occurred while setting property TerminalType: %+v", err)
	}

	if err := d.Set("tpm_control", (s.GetTpmControl())); err != nil {
		return fmt.Errorf("error occurred while setting property TpmControl: %+v", err)
	}

	if err := d.Set("tpm_support", (s.GetTpmSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property TpmSupport: %+v", err)
	}

	if err := d.Set("txt_support", (s.GetTxtSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property TxtSupport: %+v", err)
	}

	if err := d.Set("ucsm_boot_order_rule", (s.GetUcsmBootOrderRule())); err != nil {
		return fmt.Errorf("error occurred while setting property UcsmBootOrderRule: %+v", err)
	}

	if err := d.Set("usb_emul6064", (s.GetUsbEmul6064())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbEmul6064: %+v", err)
	}

	if err := d.Set("usb_port_front", (s.GetUsbPortFront())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbPortFront: %+v", err)
	}

	if err := d.Set("usb_port_internal", (s.GetUsbPortInternal())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbPortInternal: %+v", err)
	}

	if err := d.Set("usb_port_kvm", (s.GetUsbPortKvm())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbPortKvm: %+v", err)
	}

	if err := d.Set("usb_port_rear", (s.GetUsbPortRear())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbPortRear: %+v", err)
	}

	if err := d.Set("usb_port_sd_card", (s.GetUsbPortSdCard())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbPortSdCard: %+v", err)
	}

	if err := d.Set("usb_port_vmedia", (s.GetUsbPortVmedia())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbPortVmedia: %+v", err)
	}

	if err := d.Set("usb_xhci_support", (s.GetUsbXhciSupport())); err != nil {
		return fmt.Errorf("error occurred while setting property UsbXhciSupport: %+v", err)
	}

	if err := d.Set("vga_priority", (s.GetVgaPriority())); err != nil {
		return fmt.Errorf("error occurred while setting property VgaPriority: %+v", err)
	}

	if err := d.Set("vmd_enable", (s.GetVmdEnable())); err != nil {
		return fmt.Errorf("error occurred while setting property VmdEnable: %+v", err)
	}

	if err := d.Set("work_load_config", (s.GetWorkLoadConfig())); err != nil {
		return fmt.Errorf("error occurred while setting property WorkLoadConfig: %+v", err)
	}

	if err := d.Set("xpt_prefetch", (s.GetXptPrefetch())); err != nil {
		return fmt.Errorf("error occurred while setting property XptPrefetch: %+v", err)
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceBiosPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewBiosPolicyWithDefaults()
	if d.HasChange("acs_control_gpu1state") {
		v := d.Get("acs_control_gpu1state")
		x := (v.(string))
		o.SetAcsControlGpu1state(x)
	}

	if d.HasChange("acs_control_gpu2state") {
		v := d.Get("acs_control_gpu2state")
		x := (v.(string))
		o.SetAcsControlGpu2state(x)
	}

	if d.HasChange("acs_control_gpu3state") {
		v := d.Get("acs_control_gpu3state")
		x := (v.(string))
		o.SetAcsControlGpu3state(x)
	}

	if d.HasChange("acs_control_gpu4state") {
		v := d.Get("acs_control_gpu4state")
		x := (v.(string))
		o.SetAcsControlGpu4state(x)
	}

	if d.HasChange("acs_control_gpu5state") {
		v := d.Get("acs_control_gpu5state")
		x := (v.(string))
		o.SetAcsControlGpu5state(x)
	}

	if d.HasChange("acs_control_gpu6state") {
		v := d.Get("acs_control_gpu6state")
		x := (v.(string))
		o.SetAcsControlGpu6state(x)
	}

	if d.HasChange("acs_control_gpu7state") {
		v := d.Get("acs_control_gpu7state")
		x := (v.(string))
		o.SetAcsControlGpu7state(x)
	}

	if d.HasChange("acs_control_gpu8state") {
		v := d.Get("acs_control_gpu8state")
		x := (v.(string))
		o.SetAcsControlGpu8state(x)
	}

	if d.HasChange("acs_control_slot11state") {
		v := d.Get("acs_control_slot11state")
		x := (v.(string))
		o.SetAcsControlSlot11state(x)
	}

	if d.HasChange("acs_control_slot12state") {
		v := d.Get("acs_control_slot12state")
		x := (v.(string))
		o.SetAcsControlSlot12state(x)
	}

	if d.HasChange("acs_control_slot13state") {
		v := d.Get("acs_control_slot13state")
		x := (v.(string))
		o.SetAcsControlSlot13state(x)
	}

	if d.HasChange("acs_control_slot14state") {
		v := d.Get("acs_control_slot14state")
		x := (v.(string))
		o.SetAcsControlSlot14state(x)
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("adjacent_cache_line_prefetch") {
		v := d.Get("adjacent_cache_line_prefetch")
		x := (v.(string))
		o.SetAdjacentCacheLinePrefetch(x)
	}

	if d.HasChange("all_usb_devices") {
		v := d.Get("all_usb_devices")
		x := (v.(string))
		o.SetAllUsbDevices(x)
	}

	if d.HasChange("altitude") {
		v := d.Get("altitude")
		x := (v.(string))
		o.SetAltitude(x)
	}

	if d.HasChange("aspm_support") {
		v := d.Get("aspm_support")
		x := (v.(string))
		o.SetAspmSupport(x)
	}

	if d.HasChange("assert_nmi_on_perr") {
		v := d.Get("assert_nmi_on_perr")
		x := (v.(string))
		o.SetAssertNmiOnPerr(x)
	}

	if d.HasChange("assert_nmi_on_serr") {
		v := d.Get("assert_nmi_on_serr")
		x := (v.(string))
		o.SetAssertNmiOnSerr(x)
	}

	if d.HasChange("auto_cc_state") {
		v := d.Get("auto_cc_state")
		x := (v.(string))
		o.SetAutoCcState(x)
	}

	if d.HasChange("autonumous_cstate_enable") {
		v := d.Get("autonumous_cstate_enable")
		x := (v.(string))
		o.SetAutonumousCstateEnable(x)
	}

	if d.HasChange("baud_rate") {
		v := d.Get("baud_rate")
		x := (v.(string))
		o.SetBaudRate(x)
	}

	if d.HasChange("bme_dma_mitigation") {
		v := d.Get("bme_dma_mitigation")
		x := (v.(string))
		o.SetBmeDmaMitigation(x)
	}

	if d.HasChange("boot_option_num_retry") {
		v := d.Get("boot_option_num_retry")
		x := (v.(string))
		o.SetBootOptionNumRetry(x)
	}

	if d.HasChange("boot_option_re_cool_down") {
		v := d.Get("boot_option_re_cool_down")
		x := (v.(string))
		o.SetBootOptionReCoolDown(x)
	}

	if d.HasChange("boot_option_retry") {
		v := d.Get("boot_option_retry")
		x := (v.(string))
		o.SetBootOptionRetry(x)
	}

	if d.HasChange("boot_performance_mode") {
		v := d.Get("boot_performance_mode")
		x := (v.(string))
		o.SetBootPerformanceMode(x)
	}

	if d.HasChange("cbs_cmn_cpu_cpb") {
		v := d.Get("cbs_cmn_cpu_cpb")
		x := (v.(string))
		o.SetCbsCmnCpuCpb(x)
	}

	if d.HasChange("cbs_cmn_cpu_gen_downcore_ctrl") {
		v := d.Get("cbs_cmn_cpu_gen_downcore_ctrl")
		x := (v.(string))
		o.SetCbsCmnCpuGenDowncoreCtrl(x)
	}

	if d.HasChange("cbs_cmn_cpu_global_cstate_ctrl") {
		v := d.Get("cbs_cmn_cpu_global_cstate_ctrl")
		x := (v.(string))
		o.SetCbsCmnCpuGlobalCstateCtrl(x)
	}

	if d.HasChange("cbs_cmn_cpu_l1stream_hw_prefetcher") {
		v := d.Get("cbs_cmn_cpu_l1stream_hw_prefetcher")
		x := (v.(string))
		o.SetCbsCmnCpuL1streamHwPrefetcher(x)
	}

	if d.HasChange("cbs_cmn_cpu_l2stream_hw_prefetcher") {
		v := d.Get("cbs_cmn_cpu_l2stream_hw_prefetcher")
		x := (v.(string))
		o.SetCbsCmnCpuL2streamHwPrefetcher(x)
	}

	if d.HasChange("cbs_cmn_determinism_slider") {
		v := d.Get("cbs_cmn_determinism_slider")
		x := (v.(string))
		o.SetCbsCmnDeterminismSlider(x)
	}

	if d.HasChange("cbs_cmn_gnb_nb_iommu") {
		v := d.Get("cbs_cmn_gnb_nb_iommu")
		x := (v.(string))
		o.SetCbsCmnGnbNbIommu(x)
	}

	if d.HasChange("cbs_cmn_mem_ctrl_bank_group_swap_ddr4") {
		v := d.Get("cbs_cmn_mem_ctrl_bank_group_swap_ddr4")
		x := (v.(string))
		o.SetCbsCmnMemCtrlBankGroupSwapDdr4(x)
	}

	if d.HasChange("cbs_cmn_mem_map_bank_interleave_ddr4") {
		v := d.Get("cbs_cmn_mem_map_bank_interleave_ddr4")
		x := (v.(string))
		o.SetCbsCmnMemMapBankInterleaveDdr4(x)
	}

	if d.HasChange("cbs_cmnc_tdp_ctl") {
		v := d.Get("cbs_cmnc_tdp_ctl")
		x := (v.(string))
		o.SetCbsCmncTdpCtl(x)
	}

	if d.HasChange("cbs_df_cmn_mem_intlv") {
		v := d.Get("cbs_df_cmn_mem_intlv")
		x := (v.(string))
		o.SetCbsDfCmnMemIntlv(x)
	}

	if d.HasChange("cbs_df_cmn_mem_intlv_size") {
		v := d.Get("cbs_df_cmn_mem_intlv_size")
		x := (v.(string))
		o.SetCbsDfCmnMemIntlvSize(x)
	}

	if d.HasChange("cdn_enable") {
		v := d.Get("cdn_enable")
		x := (v.(string))
		o.SetCdnEnable(x)
	}

	if d.HasChange("cdn_support") {
		v := d.Get("cdn_support")
		x := (v.(string))
		o.SetCdnSupport(x)
	}

	if d.HasChange("channel_inter_leave") {
		v := d.Get("channel_inter_leave")
		x := (v.(string))
		o.SetChannelInterLeave(x)
	}

	if d.HasChange("cisco_adaptive_mem_training") {
		v := d.Get("cisco_adaptive_mem_training")
		x := (v.(string))
		o.SetCiscoAdaptiveMemTraining(x)
	}

	if d.HasChange("cisco_debug_level") {
		v := d.Get("cisco_debug_level")
		x := (v.(string))
		o.SetCiscoDebugLevel(x)
	}

	if d.HasChange("cisco_oprom_launch_optimization") {
		v := d.Get("cisco_oprom_launch_optimization")
		x := (v.(string))
		o.SetCiscoOpromLaunchOptimization(x)
	}

	if d.HasChange("cke_low_policy") {
		v := d.Get("cke_low_policy")
		x := (v.(string))
		o.SetCkeLowPolicy(x)
	}

	o.SetClassId("bios.Policy")

	if d.HasChange("closed_loop_therm_throtl") {
		v := d.Get("closed_loop_therm_throtl")
		x := (v.(string))
		o.SetClosedLoopThermThrotl(x)
	}

	if d.HasChange("cmci_enable") {
		v := d.Get("cmci_enable")
		x := (v.(string))
		o.SetCmciEnable(x)
	}

	if d.HasChange("config_tdp") {
		v := d.Get("config_tdp")
		x := (v.(string))
		o.SetConfigTdp(x)
	}

	if d.HasChange("console_redirection") {
		v := d.Get("console_redirection")
		x := (v.(string))
		o.SetConsoleRedirection(x)
	}

	if d.HasChange("core_multi_processing") {
		v := d.Get("core_multi_processing")
		x := (v.(string))
		o.SetCoreMultiProcessing(x)
	}

	if d.HasChange("cpu_energy_performance") {
		v := d.Get("cpu_energy_performance")
		x := (v.(string))
		o.SetCpuEnergyPerformance(x)
	}

	if d.HasChange("cpu_frequency_floor") {
		v := d.Get("cpu_frequency_floor")
		x := (v.(string))
		o.SetCpuFrequencyFloor(x)
	}

	if d.HasChange("cpu_performance") {
		v := d.Get("cpu_performance")
		x := (v.(string))
		o.SetCpuPerformance(x)
	}

	if d.HasChange("cpu_power_management") {
		v := d.Get("cpu_power_management")
		x := (v.(string))
		o.SetCpuPowerManagement(x)
	}

	if d.HasChange("dcpmm_firmware_downgrade") {
		v := d.Get("dcpmm_firmware_downgrade")
		x := (v.(string))
		o.SetDcpmmFirmwareDowngrade(x)
	}

	if d.HasChange("demand_scrub") {
		v := d.Get("demand_scrub")
		x := (v.(string))
		o.SetDemandScrub(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("direct_cache_access") {
		v := d.Get("direct_cache_access")
		x := (v.(string))
		o.SetDirectCacheAccess(x)
	}

	if d.HasChange("dram_clock_throttling") {
		v := d.Get("dram_clock_throttling")
		x := (v.(string))
		o.SetDramClockThrottling(x)
	}

	if d.HasChange("dram_refresh_rate") {
		v := d.Get("dram_refresh_rate")
		x := (v.(string))
		o.SetDramRefreshRate(x)
	}

	if d.HasChange("energy_efficient_turbo") {
		v := d.Get("energy_efficient_turbo")
		x := (v.(string))
		o.SetEnergyEfficientTurbo(x)
	}

	if d.HasChange("eng_perf_tuning") {
		v := d.Get("eng_perf_tuning")
		x := (v.(string))
		o.SetEngPerfTuning(x)
	}

	if d.HasChange("enhanced_intel_speed_step_tech") {
		v := d.Get("enhanced_intel_speed_step_tech")
		x := (v.(string))
		o.SetEnhancedIntelSpeedStepTech(x)
	}

	if d.HasChange("epp_profile") {
		v := d.Get("epp_profile")
		x := (v.(string))
		o.SetEppProfile(x)
	}

	if d.HasChange("execute_disable_bit") {
		v := d.Get("execute_disable_bit")
		x := (v.(string))
		o.SetExecuteDisableBit(x)
	}

	if d.HasChange("extended_apic") {
		v := d.Get("extended_apic")
		x := (v.(string))
		o.SetExtendedApic(x)
	}

	if d.HasChange("flow_control") {
		v := d.Get("flow_control")
		x := (v.(string))
		o.SetFlowControl(x)
	}

	if d.HasChange("frb2enable") {
		v := d.Get("frb2enable")
		x := (v.(string))
		o.SetFrb2enable(x)
	}

	if d.HasChange("hardware_prefetch") {
		v := d.Get("hardware_prefetch")
		x := (v.(string))
		o.SetHardwarePrefetch(x)
	}

	if d.HasChange("hwpm_enable") {
		v := d.Get("hwpm_enable")
		x := (v.(string))
		o.SetHwpmEnable(x)
	}

	if d.HasChange("imc_interleave") {
		v := d.Get("imc_interleave")
		x := (v.(string))
		o.SetImcInterleave(x)
	}

	if d.HasChange("intel_hyper_threading_tech") {
		v := d.Get("intel_hyper_threading_tech")
		x := (v.(string))
		o.SetIntelHyperThreadingTech(x)
	}

	if d.HasChange("intel_speed_select") {
		v := d.Get("intel_speed_select")
		x := (v.(string))
		o.SetIntelSpeedSelect(x)
	}

	if d.HasChange("intel_turbo_boost_tech") {
		v := d.Get("intel_turbo_boost_tech")
		x := (v.(string))
		o.SetIntelTurboBoostTech(x)
	}

	if d.HasChange("intel_virtualization_technology") {
		v := d.Get("intel_virtualization_technology")
		x := (v.(string))
		o.SetIntelVirtualizationTechnology(x)
	}

	if d.HasChange("intel_vt_for_directed_io") {
		v := d.Get("intel_vt_for_directed_io")
		x := (v.(string))
		o.SetIntelVtForDirectedIo(x)
	}

	if d.HasChange("intel_vtd_coherency_support") {
		v := d.Get("intel_vtd_coherency_support")
		x := (v.(string))
		o.SetIntelVtdCoherencySupport(x)
	}

	if d.HasChange("intel_vtd_interrupt_remapping") {
		v := d.Get("intel_vtd_interrupt_remapping")
		x := (v.(string))
		o.SetIntelVtdInterruptRemapping(x)
	}

	if d.HasChange("intel_vtd_pass_through_dma_support") {
		v := d.Get("intel_vtd_pass_through_dma_support")
		x := (v.(string))
		o.SetIntelVtdPassThroughDmaSupport(x)
	}

	if d.HasChange("intel_vtdats_support") {
		v := d.Get("intel_vtdats_support")
		x := (v.(string))
		o.SetIntelVtdatsSupport(x)
	}

	if d.HasChange("ioh_error_enable") {
		v := d.Get("ioh_error_enable")
		x := (v.(string))
		o.SetIohErrorEnable(x)
	}

	if d.HasChange("ioh_resource") {
		v := d.Get("ioh_resource")
		x := (v.(string))
		o.SetIohResource(x)
	}

	if d.HasChange("ip_prefetch") {
		v := d.Get("ip_prefetch")
		x := (v.(string))
		o.SetIpPrefetch(x)
	}

	if d.HasChange("ipv4pxe") {
		v := d.Get("ipv4pxe")
		x := (v.(string))
		o.SetIpv4pxe(x)
	}

	if d.HasChange("ipv6pxe") {
		v := d.Get("ipv6pxe")
		x := (v.(string))
		o.SetIpv6pxe(x)
	}

	if d.HasChange("kti_prefetch") {
		v := d.Get("kti_prefetch")
		x := (v.(string))
		o.SetKtiPrefetch(x)
	}

	if d.HasChange("legacy_os_redirection") {
		v := d.Get("legacy_os_redirection")
		x := (v.(string))
		o.SetLegacyOsRedirection(x)
	}

	if d.HasChange("legacy_usb_support") {
		v := d.Get("legacy_usb_support")
		x := (v.(string))
		o.SetLegacyUsbSupport(x)
	}

	if d.HasChange("llc_prefetch") {
		v := d.Get("llc_prefetch")
		x := (v.(string))
		o.SetLlcPrefetch(x)
	}

	if d.HasChange("lom_port0state") {
		v := d.Get("lom_port0state")
		x := (v.(string))
		o.SetLomPort0state(x)
	}

	if d.HasChange("lom_port1state") {
		v := d.Get("lom_port1state")
		x := (v.(string))
		o.SetLomPort1state(x)
	}

	if d.HasChange("lom_port2state") {
		v := d.Get("lom_port2state")
		x := (v.(string))
		o.SetLomPort2state(x)
	}

	if d.HasChange("lom_port3state") {
		v := d.Get("lom_port3state")
		x := (v.(string))
		o.SetLomPort3state(x)
	}

	if d.HasChange("lom_ports_all_state") {
		v := d.Get("lom_ports_all_state")
		x := (v.(string))
		o.SetLomPortsAllState(x)
	}

	if d.HasChange("lv_ddr_mode") {
		v := d.Get("lv_ddr_mode")
		x := (v.(string))
		o.SetLvDdrMode(x)
	}

	if d.HasChange("make_device_non_bootable") {
		v := d.Get("make_device_non_bootable")
		x := (v.(string))
		o.SetMakeDeviceNonBootable(x)
	}

	if d.HasChange("memory_inter_leave") {
		v := d.Get("memory_inter_leave")
		x := (v.(string))
		o.SetMemoryInterLeave(x)
	}

	if d.HasChange("memory_mapped_io_above4gb") {
		v := d.Get("memory_mapped_io_above4gb")
		x := (v.(string))
		o.SetMemoryMappedIoAbove4gb(x)
	}

	if d.HasChange("memory_size_limit") {
		v := d.Get("memory_size_limit")
		x := (v.(string))
		o.SetMemorySizeLimit(x)
	}

	if d.HasChange("mirroring_mode") {
		v := d.Get("mirroring_mode")
		x := (v.(string))
		o.SetMirroringMode(x)
	}

	if d.HasChange("mmcfg_base") {
		v := d.Get("mmcfg_base")
		x := (v.(string))
		o.SetMmcfgBase(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if d.HasChange("network_stack") {
		v := d.Get("network_stack")
		x := (v.(string))
		o.SetNetworkStack(x)
	}

	if d.HasChange("numa_optimized") {
		v := d.Get("numa_optimized")
		x := (v.(string))
		o.SetNumaOptimized(x)
	}

	o.SetObjectType("bios.Policy")

	if d.HasChange("onboard10gbit_lom") {
		v := d.Get("onboard10gbit_lom")
		x := (v.(string))
		o.SetOnboard10gbitLom(x)
	}

	if d.HasChange("onboard_gbit_lom") {
		v := d.Get("onboard_gbit_lom")
		x := (v.(string))
		o.SetOnboardGbitLom(x)
	}

	if d.HasChange("onboard_scu_storage_support") {
		v := d.Get("onboard_scu_storage_support")
		x := (v.(string))
		o.SetOnboardScuStorageSupport(x)
	}

	if d.HasChange("onboard_scu_storage_sw_stack") {
		v := d.Get("onboard_scu_storage_sw_stack")
		x := (v.(string))
		o.SetOnboardScuStorageSwStack(x)
	}

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if d.HasChange("os_boot_watchdog_timer") {
		v := d.Get("os_boot_watchdog_timer")
		x := (v.(string))
		o.SetOsBootWatchdogTimer(x)
	}

	if d.HasChange("os_boot_watchdog_timer_policy") {
		v := d.Get("os_boot_watchdog_timer_policy")
		x := (v.(string))
		o.SetOsBootWatchdogTimerPolicy(x)
	}

	if d.HasChange("os_boot_watchdog_timer_timeout") {
		v := d.Get("os_boot_watchdog_timer_timeout")
		x := (v.(string))
		o.SetOsBootWatchdogTimerTimeout(x)
	}

	if d.HasChange("out_of_band_mgmt_port") {
		v := d.Get("out_of_band_mgmt_port")
		x := (v.(string))
		o.SetOutOfBandMgmtPort(x)
	}

	if d.HasChange("package_cstate_limit") {
		v := d.Get("package_cstate_limit")
		x := (v.(string))
		o.SetPackageCstateLimit(x)
	}

	if d.HasChange("partial_mirror_mode_config") {
		v := d.Get("partial_mirror_mode_config")
		x := (v.(string))
		o.SetPartialMirrorModeConfig(x)
	}

	if d.HasChange("partial_mirror_percent") {
		v := d.Get("partial_mirror_percent")
		x := (v.(string))
		o.SetPartialMirrorPercent(x)
	}

	if d.HasChange("partial_mirror_value1") {
		v := d.Get("partial_mirror_value1")
		x := (v.(string))
		o.SetPartialMirrorValue1(x)
	}

	if d.HasChange("partial_mirror_value2") {
		v := d.Get("partial_mirror_value2")
		x := (v.(string))
		o.SetPartialMirrorValue2(x)
	}

	if d.HasChange("partial_mirror_value3") {
		v := d.Get("partial_mirror_value3")
		x := (v.(string))
		o.SetPartialMirrorValue3(x)
	}

	if d.HasChange("partial_mirror_value4") {
		v := d.Get("partial_mirror_value4")
		x := (v.(string))
		o.SetPartialMirrorValue4(x)
	}

	if d.HasChange("patrol_scrub") {
		v := d.Get("patrol_scrub")
		x := (v.(string))
		o.SetPatrolScrub(x)
	}

	if d.HasChange("patrol_scrub_duration") {
		v := d.Get("patrol_scrub_duration")
		x := (v.(string))
		o.SetPatrolScrubDuration(x)
	}

	if d.HasChange("pc_ie_ras_support") {
		v := d.Get("pc_ie_ras_support")
		x := (v.(string))
		o.SetPcIeRasSupport(x)
	}

	if d.HasChange("pc_ie_ssd_hot_plug_support") {
		v := d.Get("pc_ie_ssd_hot_plug_support")
		x := (v.(string))
		o.SetPcIeSsdHotPlugSupport(x)
	}

	if d.HasChange("pch_usb30mode") {
		v := d.Get("pch_usb30mode")
		x := (v.(string))
		o.SetPchUsb30mode(x)
	}

	if d.HasChange("pci_option_ro_ms") {
		v := d.Get("pci_option_ro_ms")
		x := (v.(string))
		o.SetPciOptionRoMs(x)
	}

	if d.HasChange("pci_rom_clp") {
		v := d.Get("pci_rom_clp")
		x := (v.(string))
		o.SetPciRomClp(x)
	}

	if d.HasChange("pop_support") {
		v := d.Get("pop_support")
		x := (v.(string))
		o.SetPopSupport(x)
	}

	if d.HasChange("post_error_pause") {
		v := d.Get("post_error_pause")
		x := (v.(string))
		o.SetPostErrorPause(x)
	}

	if d.HasChange("processor_c1e") {
		v := d.Get("processor_c1e")
		x := (v.(string))
		o.SetProcessorC1e(x)
	}

	if d.HasChange("processor_c3report") {
		v := d.Get("processor_c3report")
		x := (v.(string))
		o.SetProcessorC3report(x)
	}

	if d.HasChange("processor_c6report") {
		v := d.Get("processor_c6report")
		x := (v.(string))
		o.SetProcessorC6report(x)
	}

	if d.HasChange("processor_cstate") {
		v := d.Get("processor_cstate")
		x := (v.(string))
		o.SetProcessorCstate(x)
	}

	if d.HasChange("profiles") {
		v := d.Get("profiles")
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsPolicyAbstractConfigProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetProfiles(x)
		}
	}

	if d.HasChange("psata") {
		v := d.Get("psata")
		x := (v.(string))
		o.SetPsata(x)
	}

	if d.HasChange("pstate_coord_type") {
		v := d.Get("pstate_coord_type")
		x := (v.(string))
		o.SetPstateCoordType(x)
	}

	if d.HasChange("putty_key_pad") {
		v := d.Get("putty_key_pad")
		x := (v.(string))
		o.SetPuttyKeyPad(x)
	}

	if d.HasChange("pwr_perf_tuning") {
		v := d.Get("pwr_perf_tuning")
		x := (v.(string))
		o.SetPwrPerfTuning(x)
	}

	if d.HasChange("qpi_link_frequency") {
		v := d.Get("qpi_link_frequency")
		x := (v.(string))
		o.SetQpiLinkFrequency(x)
	}

	if d.HasChange("qpi_snoop_mode") {
		v := d.Get("qpi_snoop_mode")
		x := (v.(string))
		o.SetQpiSnoopMode(x)
	}

	if d.HasChange("rank_inter_leave") {
		v := d.Get("rank_inter_leave")
		x := (v.(string))
		o.SetRankInterLeave(x)
	}

	if d.HasChange("redirection_after_post") {
		v := d.Get("redirection_after_post")
		x := (v.(string))
		o.SetRedirectionAfterPost(x)
	}

	if d.HasChange("sata_mode_select") {
		v := d.Get("sata_mode_select")
		x := (v.(string))
		o.SetSataModeSelect(x)
	}

	if d.HasChange("select_memory_ras_configuration") {
		v := d.Get("select_memory_ras_configuration")
		x := (v.(string))
		o.SetSelectMemoryRasConfiguration(x)
	}

	if d.HasChange("select_ppr_type") {
		v := d.Get("select_ppr_type")
		x := (v.(string))
		o.SetSelectPprType(x)
	}

	if d.HasChange("serial_port_aenable") {
		v := d.Get("serial_port_aenable")
		x := (v.(string))
		o.SetSerialPortAenable(x)
	}

	if d.HasChange("single_pctl_enable") {
		v := d.Get("single_pctl_enable")
		x := (v.(string))
		o.SetSinglePctlEnable(x)
	}

	if d.HasChange("slot10link_speed") {
		v := d.Get("slot10link_speed")
		x := (v.(string))
		o.SetSlot10linkSpeed(x)
	}

	if d.HasChange("slot10state") {
		v := d.Get("slot10state")
		x := (v.(string))
		o.SetSlot10state(x)
	}

	if d.HasChange("slot11link_speed") {
		v := d.Get("slot11link_speed")
		x := (v.(string))
		o.SetSlot11linkSpeed(x)
	}

	if d.HasChange("slot11state") {
		v := d.Get("slot11state")
		x := (v.(string))
		o.SetSlot11state(x)
	}

	if d.HasChange("slot12link_speed") {
		v := d.Get("slot12link_speed")
		x := (v.(string))
		o.SetSlot12linkSpeed(x)
	}

	if d.HasChange("slot12state") {
		v := d.Get("slot12state")
		x := (v.(string))
		o.SetSlot12state(x)
	}

	if d.HasChange("slot13state") {
		v := d.Get("slot13state")
		x := (v.(string))
		o.SetSlot13state(x)
	}

	if d.HasChange("slot14state") {
		v := d.Get("slot14state")
		x := (v.(string))
		o.SetSlot14state(x)
	}

	if d.HasChange("slot1link_speed") {
		v := d.Get("slot1link_speed")
		x := (v.(string))
		o.SetSlot1linkSpeed(x)
	}

	if d.HasChange("slot1state") {
		v := d.Get("slot1state")
		x := (v.(string))
		o.SetSlot1state(x)
	}

	if d.HasChange("slot2link_speed") {
		v := d.Get("slot2link_speed")
		x := (v.(string))
		o.SetSlot2linkSpeed(x)
	}

	if d.HasChange("slot2state") {
		v := d.Get("slot2state")
		x := (v.(string))
		o.SetSlot2state(x)
	}

	if d.HasChange("slot3link_speed") {
		v := d.Get("slot3link_speed")
		x := (v.(string))
		o.SetSlot3linkSpeed(x)
	}

	if d.HasChange("slot3state") {
		v := d.Get("slot3state")
		x := (v.(string))
		o.SetSlot3state(x)
	}

	if d.HasChange("slot4link_speed") {
		v := d.Get("slot4link_speed")
		x := (v.(string))
		o.SetSlot4linkSpeed(x)
	}

	if d.HasChange("slot4state") {
		v := d.Get("slot4state")
		x := (v.(string))
		o.SetSlot4state(x)
	}

	if d.HasChange("slot5link_speed") {
		v := d.Get("slot5link_speed")
		x := (v.(string))
		o.SetSlot5linkSpeed(x)
	}

	if d.HasChange("slot5state") {
		v := d.Get("slot5state")
		x := (v.(string))
		o.SetSlot5state(x)
	}

	if d.HasChange("slot6link_speed") {
		v := d.Get("slot6link_speed")
		x := (v.(string))
		o.SetSlot6linkSpeed(x)
	}

	if d.HasChange("slot6state") {
		v := d.Get("slot6state")
		x := (v.(string))
		o.SetSlot6state(x)
	}

	if d.HasChange("slot7link_speed") {
		v := d.Get("slot7link_speed")
		x := (v.(string))
		o.SetSlot7linkSpeed(x)
	}

	if d.HasChange("slot7state") {
		v := d.Get("slot7state")
		x := (v.(string))
		o.SetSlot7state(x)
	}

	if d.HasChange("slot8link_speed") {
		v := d.Get("slot8link_speed")
		x := (v.(string))
		o.SetSlot8linkSpeed(x)
	}

	if d.HasChange("slot8state") {
		v := d.Get("slot8state")
		x := (v.(string))
		o.SetSlot8state(x)
	}

	if d.HasChange("slot9link_speed") {
		v := d.Get("slot9link_speed")
		x := (v.(string))
		o.SetSlot9linkSpeed(x)
	}

	if d.HasChange("slot9state") {
		v := d.Get("slot9state")
		x := (v.(string))
		o.SetSlot9state(x)
	}

	if d.HasChange("slot_flom_link_speed") {
		v := d.Get("slot_flom_link_speed")
		x := (v.(string))
		o.SetSlotFlomLinkSpeed(x)
	}

	if d.HasChange("slot_front_nvme1link_speed") {
		v := d.Get("slot_front_nvme1link_speed")
		x := (v.(string))
		o.SetSlotFrontNvme1linkSpeed(x)
	}

	if d.HasChange("slot_front_nvme2link_speed") {
		v := d.Get("slot_front_nvme2link_speed")
		x := (v.(string))
		o.SetSlotFrontNvme2linkSpeed(x)
	}

	if d.HasChange("slot_front_slot5link_speed") {
		v := d.Get("slot_front_slot5link_speed")
		x := (v.(string))
		o.SetSlotFrontSlot5linkSpeed(x)
	}

	if d.HasChange("slot_front_slot6link_speed") {
		v := d.Get("slot_front_slot6link_speed")
		x := (v.(string))
		o.SetSlotFrontSlot6linkSpeed(x)
	}

	if d.HasChange("slot_gpu1state") {
		v := d.Get("slot_gpu1state")
		x := (v.(string))
		o.SetSlotGpu1state(x)
	}

	if d.HasChange("slot_gpu2state") {
		v := d.Get("slot_gpu2state")
		x := (v.(string))
		o.SetSlotGpu2state(x)
	}

	if d.HasChange("slot_gpu3state") {
		v := d.Get("slot_gpu3state")
		x := (v.(string))
		o.SetSlotGpu3state(x)
	}

	if d.HasChange("slot_gpu4state") {
		v := d.Get("slot_gpu4state")
		x := (v.(string))
		o.SetSlotGpu4state(x)
	}

	if d.HasChange("slot_gpu5state") {
		v := d.Get("slot_gpu5state")
		x := (v.(string))
		o.SetSlotGpu5state(x)
	}

	if d.HasChange("slot_gpu6state") {
		v := d.Get("slot_gpu6state")
		x := (v.(string))
		o.SetSlotGpu6state(x)
	}

	if d.HasChange("slot_gpu7state") {
		v := d.Get("slot_gpu7state")
		x := (v.(string))
		o.SetSlotGpu7state(x)
	}

	if d.HasChange("slot_gpu8state") {
		v := d.Get("slot_gpu8state")
		x := (v.(string))
		o.SetSlotGpu8state(x)
	}

	if d.HasChange("slot_hba_link_speed") {
		v := d.Get("slot_hba_link_speed")
		x := (v.(string))
		o.SetSlotHbaLinkSpeed(x)
	}

	if d.HasChange("slot_hba_state") {
		v := d.Get("slot_hba_state")
		x := (v.(string))
		o.SetSlotHbaState(x)
	}

	if d.HasChange("slot_lom1link") {
		v := d.Get("slot_lom1link")
		x := (v.(string))
		o.SetSlotLom1link(x)
	}

	if d.HasChange("slot_lom2link") {
		v := d.Get("slot_lom2link")
		x := (v.(string))
		o.SetSlotLom2link(x)
	}

	if d.HasChange("slot_mezz_state") {
		v := d.Get("slot_mezz_state")
		x := (v.(string))
		o.SetSlotMezzState(x)
	}

	if d.HasChange("slot_mlom_link_speed") {
		v := d.Get("slot_mlom_link_speed")
		x := (v.(string))
		o.SetSlotMlomLinkSpeed(x)
	}

	if d.HasChange("slot_mlom_state") {
		v := d.Get("slot_mlom_state")
		x := (v.(string))
		o.SetSlotMlomState(x)
	}

	if d.HasChange("slot_mraid_link_speed") {
		v := d.Get("slot_mraid_link_speed")
		x := (v.(string))
		o.SetSlotMraidLinkSpeed(x)
	}

	if d.HasChange("slot_mraid_state") {
		v := d.Get("slot_mraid_state")
		x := (v.(string))
		o.SetSlotMraidState(x)
	}

	if d.HasChange("slot_n10state") {
		v := d.Get("slot_n10state")
		x := (v.(string))
		o.SetSlotN10state(x)
	}

	if d.HasChange("slot_n11state") {
		v := d.Get("slot_n11state")
		x := (v.(string))
		o.SetSlotN11state(x)
	}

	if d.HasChange("slot_n12state") {
		v := d.Get("slot_n12state")
		x := (v.(string))
		o.SetSlotN12state(x)
	}

	if d.HasChange("slot_n13state") {
		v := d.Get("slot_n13state")
		x := (v.(string))
		o.SetSlotN13state(x)
	}

	if d.HasChange("slot_n14state") {
		v := d.Get("slot_n14state")
		x := (v.(string))
		o.SetSlotN14state(x)
	}

	if d.HasChange("slot_n15state") {
		v := d.Get("slot_n15state")
		x := (v.(string))
		o.SetSlotN15state(x)
	}

	if d.HasChange("slot_n16state") {
		v := d.Get("slot_n16state")
		x := (v.(string))
		o.SetSlotN16state(x)
	}

	if d.HasChange("slot_n17state") {
		v := d.Get("slot_n17state")
		x := (v.(string))
		o.SetSlotN17state(x)
	}

	if d.HasChange("slot_n18state") {
		v := d.Get("slot_n18state")
		x := (v.(string))
		o.SetSlotN18state(x)
	}

	if d.HasChange("slot_n19state") {
		v := d.Get("slot_n19state")
		x := (v.(string))
		o.SetSlotN19state(x)
	}

	if d.HasChange("slot_n1state") {
		v := d.Get("slot_n1state")
		x := (v.(string))
		o.SetSlotN1state(x)
	}

	if d.HasChange("slot_n20state") {
		v := d.Get("slot_n20state")
		x := (v.(string))
		o.SetSlotN20state(x)
	}

	if d.HasChange("slot_n21state") {
		v := d.Get("slot_n21state")
		x := (v.(string))
		o.SetSlotN21state(x)
	}

	if d.HasChange("slot_n22state") {
		v := d.Get("slot_n22state")
		x := (v.(string))
		o.SetSlotN22state(x)
	}

	if d.HasChange("slot_n23state") {
		v := d.Get("slot_n23state")
		x := (v.(string))
		o.SetSlotN23state(x)
	}

	if d.HasChange("slot_n24state") {
		v := d.Get("slot_n24state")
		x := (v.(string))
		o.SetSlotN24state(x)
	}

	if d.HasChange("slot_n2state") {
		v := d.Get("slot_n2state")
		x := (v.(string))
		o.SetSlotN2state(x)
	}

	if d.HasChange("slot_n3state") {
		v := d.Get("slot_n3state")
		x := (v.(string))
		o.SetSlotN3state(x)
	}

	if d.HasChange("slot_n4state") {
		v := d.Get("slot_n4state")
		x := (v.(string))
		o.SetSlotN4state(x)
	}

	if d.HasChange("slot_n5state") {
		v := d.Get("slot_n5state")
		x := (v.(string))
		o.SetSlotN5state(x)
	}

	if d.HasChange("slot_n6state") {
		v := d.Get("slot_n6state")
		x := (v.(string))
		o.SetSlotN6state(x)
	}

	if d.HasChange("slot_n7state") {
		v := d.Get("slot_n7state")
		x := (v.(string))
		o.SetSlotN7state(x)
	}

	if d.HasChange("slot_n8state") {
		v := d.Get("slot_n8state")
		x := (v.(string))
		o.SetSlotN8state(x)
	}

	if d.HasChange("slot_n9state") {
		v := d.Get("slot_n9state")
		x := (v.(string))
		o.SetSlotN9state(x)
	}

	if d.HasChange("slot_raid_link_speed") {
		v := d.Get("slot_raid_link_speed")
		x := (v.(string))
		o.SetSlotRaidLinkSpeed(x)
	}

	if d.HasChange("slot_raid_state") {
		v := d.Get("slot_raid_state")
		x := (v.(string))
		o.SetSlotRaidState(x)
	}

	if d.HasChange("slot_rear_nvme1link_speed") {
		v := d.Get("slot_rear_nvme1link_speed")
		x := (v.(string))
		o.SetSlotRearNvme1linkSpeed(x)
	}

	if d.HasChange("slot_rear_nvme1state") {
		v := d.Get("slot_rear_nvme1state")
		x := (v.(string))
		o.SetSlotRearNvme1state(x)
	}

	if d.HasChange("slot_rear_nvme2link_speed") {
		v := d.Get("slot_rear_nvme2link_speed")
		x := (v.(string))
		o.SetSlotRearNvme2linkSpeed(x)
	}

	if d.HasChange("slot_rear_nvme2state") {
		v := d.Get("slot_rear_nvme2state")
		x := (v.(string))
		o.SetSlotRearNvme2state(x)
	}

	if d.HasChange("slot_rear_nvme3state") {
		v := d.Get("slot_rear_nvme3state")
		x := (v.(string))
		o.SetSlotRearNvme3state(x)
	}

	if d.HasChange("slot_rear_nvme4state") {
		v := d.Get("slot_rear_nvme4state")
		x := (v.(string))
		o.SetSlotRearNvme4state(x)
	}

	if d.HasChange("slot_rear_nvme5state") {
		v := d.Get("slot_rear_nvme5state")
		x := (v.(string))
		o.SetSlotRearNvme5state(x)
	}

	if d.HasChange("slot_rear_nvme6state") {
		v := d.Get("slot_rear_nvme6state")
		x := (v.(string))
		o.SetSlotRearNvme6state(x)
	}

	if d.HasChange("slot_rear_nvme7state") {
		v := d.Get("slot_rear_nvme7state")
		x := (v.(string))
		o.SetSlotRearNvme7state(x)
	}

	if d.HasChange("slot_rear_nvme8state") {
		v := d.Get("slot_rear_nvme8state")
		x := (v.(string))
		o.SetSlotRearNvme8state(x)
	}

	if d.HasChange("slot_riser1link_speed") {
		v := d.Get("slot_riser1link_speed")
		x := (v.(string))
		o.SetSlotRiser1linkSpeed(x)
	}

	if d.HasChange("slot_riser1slot1link_speed") {
		v := d.Get("slot_riser1slot1link_speed")
		x := (v.(string))
		o.SetSlotRiser1slot1linkSpeed(x)
	}

	if d.HasChange("slot_riser1slot2link_speed") {
		v := d.Get("slot_riser1slot2link_speed")
		x := (v.(string))
		o.SetSlotRiser1slot2linkSpeed(x)
	}

	if d.HasChange("slot_riser1slot3link_speed") {
		v := d.Get("slot_riser1slot3link_speed")
		x := (v.(string))
		o.SetSlotRiser1slot3linkSpeed(x)
	}

	if d.HasChange("slot_riser2link_speed") {
		v := d.Get("slot_riser2link_speed")
		x := (v.(string))
		o.SetSlotRiser2linkSpeed(x)
	}

	if d.HasChange("slot_riser2slot4link_speed") {
		v := d.Get("slot_riser2slot4link_speed")
		x := (v.(string))
		o.SetSlotRiser2slot4linkSpeed(x)
	}

	if d.HasChange("slot_riser2slot5link_speed") {
		v := d.Get("slot_riser2slot5link_speed")
		x := (v.(string))
		o.SetSlotRiser2slot5linkSpeed(x)
	}

	if d.HasChange("slot_riser2slot6link_speed") {
		v := d.Get("slot_riser2slot6link_speed")
		x := (v.(string))
		o.SetSlotRiser2slot6linkSpeed(x)
	}

	if d.HasChange("slot_sas_state") {
		v := d.Get("slot_sas_state")
		x := (v.(string))
		o.SetSlotSasState(x)
	}

	if d.HasChange("slot_ssd_slot1link_speed") {
		v := d.Get("slot_ssd_slot1link_speed")
		x := (v.(string))
		o.SetSlotSsdSlot1linkSpeed(x)
	}

	if d.HasChange("slot_ssd_slot2link_speed") {
		v := d.Get("slot_ssd_slot2link_speed")
		x := (v.(string))
		o.SetSlotSsdSlot2linkSpeed(x)
	}

	if d.HasChange("smee") {
		v := d.Get("smee")
		x := (v.(string))
		o.SetSmee(x)
	}

	if d.HasChange("smt_mode") {
		v := d.Get("smt_mode")
		x := (v.(string))
		o.SetSmtMode(x)
	}

	if d.HasChange("snc") {
		v := d.Get("snc")
		x := (v.(string))
		o.SetSnc(x)
	}

	if d.HasChange("sparing_mode") {
		v := d.Get("sparing_mode")
		x := (v.(string))
		o.SetSparingMode(x)
	}

	if d.HasChange("sr_iov") {
		v := d.Get("sr_iov")
		x := (v.(string))
		o.SetSrIov(x)
	}

	if d.HasChange("streamer_prefetch") {
		v := d.Get("streamer_prefetch")
		x := (v.(string))
		o.SetStreamerPrefetch(x)
	}

	if d.HasChange("svm_mode") {
		v := d.Get("svm_mode")
		x := (v.(string))
		o.SetSvmMode(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if d.HasChange("terminal_type") {
		v := d.Get("terminal_type")
		x := (v.(string))
		o.SetTerminalType(x)
	}

	if d.HasChange("tpm_control") {
		v := d.Get("tpm_control")
		x := (v.(string))
		o.SetTpmControl(x)
	}

	if d.HasChange("tpm_support") {
		v := d.Get("tpm_support")
		x := (v.(string))
		o.SetTpmSupport(x)
	}

	if d.HasChange("txt_support") {
		v := d.Get("txt_support")
		x := (v.(string))
		o.SetTxtSupport(x)
	}

	if d.HasChange("ucsm_boot_order_rule") {
		v := d.Get("ucsm_boot_order_rule")
		x := (v.(string))
		o.SetUcsmBootOrderRule(x)
	}

	if d.HasChange("usb_emul6064") {
		v := d.Get("usb_emul6064")
		x := (v.(string))
		o.SetUsbEmul6064(x)
	}

	if d.HasChange("usb_port_front") {
		v := d.Get("usb_port_front")
		x := (v.(string))
		o.SetUsbPortFront(x)
	}

	if d.HasChange("usb_port_internal") {
		v := d.Get("usb_port_internal")
		x := (v.(string))
		o.SetUsbPortInternal(x)
	}

	if d.HasChange("usb_port_kvm") {
		v := d.Get("usb_port_kvm")
		x := (v.(string))
		o.SetUsbPortKvm(x)
	}

	if d.HasChange("usb_port_rear") {
		v := d.Get("usb_port_rear")
		x := (v.(string))
		o.SetUsbPortRear(x)
	}

	if d.HasChange("usb_port_sd_card") {
		v := d.Get("usb_port_sd_card")
		x := (v.(string))
		o.SetUsbPortSdCard(x)
	}

	if d.HasChange("usb_port_vmedia") {
		v := d.Get("usb_port_vmedia")
		x := (v.(string))
		o.SetUsbPortVmedia(x)
	}

	if d.HasChange("usb_xhci_support") {
		v := d.Get("usb_xhci_support")
		x := (v.(string))
		o.SetUsbXhciSupport(x)
	}

	if d.HasChange("vga_priority") {
		v := d.Get("vga_priority")
		x := (v.(string))
		o.SetVgaPriority(x)
	}

	if d.HasChange("vmd_enable") {
		v := d.Get("vmd_enable")
		x := (v.(string))
		o.SetVmdEnable(x)
	}

	if d.HasChange("work_load_config") {
		v := d.Get("work_load_config")
		x := (v.(string))
		o.SetWorkLoadConfig(x)
	}

	if d.HasChange("xpt_prefetch") {
		v := d.Get("xpt_prefetch")
		x := (v.(string))
		o.SetXptPrefetch(x)
	}

	r := conn.ApiClient.BiosApi.UpdateBiosPolicy(conn.ctx, d.Id()).BiosPolicy(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceBiosPolicyRead(d, meta)
}

func resourceBiosPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	if p, ok := d.GetOk("profiles"); ok {
		if len(p.([]interface{})) > 0 {
			e := detachBiosPolicyProfiles(d, meta)
			if e != nil {
				return e
			}
		}
	}
	p := conn.ApiClient.BiosApi.DeleteBiosPolicy(conn.ctx, d.Id())
	_, err := p.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while deleting: %s", err.Error())
	}
	return err
}
