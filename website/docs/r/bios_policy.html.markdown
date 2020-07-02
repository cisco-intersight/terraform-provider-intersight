
---
layout: "intersight"
page_title: "Intersight: intersight_bios_policy"
sidebar_current: "docs-intersight-resource-bios-policy"
description: |-
  Policy for setting BIOS tokens on the endpoint.
---

# Resource: intersight_bios._policy
Policy for setting BIOS tokens on the endpoint.
## Argument Reference
The following arguments are supported:
* `acs_control_gpu1state`:(string) BIOS Token for setting ACS Control GPU-1 configuration. 
* `acs_control_gpu2state`:(string) BIOS Token for setting ACS Control GPU-2 configuration. 
* `acs_control_gpu3state`:(string) BIOS Token for setting ACS Control GPU-3 configuration. 
* `acs_control_gpu4state`:(string) BIOS Token for setting ACS Control GPU-4 configuration. 
* `acs_control_gpu5state`:(string) BIOS Token for setting ACS Control GPU-5 configuration. 
* `acs_control_gpu6state`:(string) BIOS Token for setting ACS Control GPU-6 configuration. 
* `acs_control_gpu7state`:(string) BIOS Token for setting ACS Control GPU-7 configuration. 
* `acs_control_gpu8state`:(string) BIOS Token for setting ACS Control GPU-8 configuration. 
* `acs_control_slot11state`:(string) BIOS Token for setting ACS Control Slot 11 configuration. 
* `acs_control_slot12state`:(string) BIOS Token for setting ACS Control Slot 12 configuration. 
* `acs_control_slot13state`:(string) BIOS Token for setting ACS Control Slot 13 configuration. 
* `acs_control_slot14state`:(string) BIOS Token for setting ACS Control Slot 14 configuration. 
* `adjacent_cache_line_prefetch`:(string) BIOS Token for setting Adjacent Cache Line Prefetcher configuration. 
* `all_usb_devices`:(string) BIOS Token for setting All USB Devices configuration. 
* `altitude`:(string) BIOS Token for setting Altitude configuration. 
* `aspm_support`:(string) BIOS Token for setting ASPM Support configuration. 
* `assert_nmi_on_perr`:(string) BIOS Token for setting Assert NMI on PERR configuration. 
* `assert_nmi_on_serr`:(string) BIOS Token for setting Assert NMI on SERR configuration. 
* `auto_cc_state`:(string) BIOS Token for setting Autonomous Core C-state configuration. 
* `autonumous_cstate_enable`:(string) BIOS Token for setting CPU Autonomous Cstate configuration. 
* `baud_rate`:(string) BIOS Token for setting Baud rate configuration. 
* `bme_dma_mitigation`:(string) BIOS Token for setting BME DMA Mitigation configuration. 
* `boot_option_num_retry`:(string) BIOS Token for setting Number of Retries configuration. 
* `boot_option_re_cool_down`:(string) BIOS Token for setting Cool Down Time  (sec) configuration. 
* `boot_option_retry`:(string) BIOS Token for setting Boot option retry configuration. 
* `boot_performance_mode`:(string) BIOS Token for setting Boot Performance Mode configuration. 
* `cbs_cmn_cpu_cpb`:(string) BIOS Token for setting Core Performance Boost configuration. 
* `cbs_cmn_cpu_gen_downcore_ctrl`:(string) BIOS Token for setting Downcore control configuration. 
* `cbs_cmn_cpu_global_cstate_ctrl`:(string) BIOS Token for setting Global C-state Control configuration. 
* `cbs_cmn_cpu_l1stream_hw_prefetcher`:(string) BIOS Token for setting L1 Stream HW Prefetcher configuration. 
* `cbs_cmn_cpu_l2stream_hw_prefetcher`:(string) BIOS Token for setting L2 Stream HW Prefetcher configuration. 
* `cbs_cmn_determinism_slider`:(string) BIOS Token for setting Determinism Slider configuration. 
* `cbs_cmn_gnb_nb_iommu`:(string) BIOS Token for setting IOMMU configuration. 
* `cbs_cmn_mem_ctrl_bank_group_swap_ddr4`:(string) BIOS Token for setting Bank Group Swap configuration. 
* `cbs_cmn_mem_map_bank_interleave_ddr4`:(string) BIOS Token for setting Chipselect Interleaving configuration. 
* `cbs_cmnc_tdp_ctl`:(string) BIOS Token for setting cTDP Control configuration. 
* `cbs_df_cmn_mem_intlv`:(string) BIOS Token for setting Memory interleaving configuration. 
* `cbs_df_cmn_mem_intlv_size`:(string) BIOS Token for setting Memory interleaving size configuration. 
* `cdn_enable`:(string) BIOS Token for setting Consistent Device Naming configuration. 
* `cdn_support`:(string) BIOS Token for setting CDN Support for LOM configuration. 
* `channel_inter_leave`:(string) BIOS Token for setting Channel Interleaving configuration. 
* `cisco_adaptive_mem_training`:(string) BIOS Token for setting Adaptive Memory Training configuration. 
* `cisco_debug_level`:(string) BIOS Token for setting BIOS Techlog Level configuration. 
* `cisco_oprom_launch_optimization`:(string) BIOS Token for setting OptionROM Launch Optimization configuration. 
* `cke_low_policy`:(string) BIOS Token for setting CKE Low Policy configuration. 
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `closed_loop_therm_throtl`:(string) BIOS Token for setting Closed Loop Therm Throt configuration. 
* `cmci_enable`:(string) BIOS Token for setting Processor CMCI configuration. 
* `config_tdp`:(string) BIOS Token for setting Config TDP configuration. 
* `console_redirection`:(string) BIOS Token for setting Console redirection configuration. 
* `core_multi_processing`:(string) BIOS Token for setting Core MultiProcessing configuration. 
* `cpu_energy_performance`:(string) BIOS Token for setting Energy Performance configuration. 
* `cpu_frequency_floor`:(string) BIOS Token for setting Frequency Floor Override configuration. 
* `cpu_performance`:(string) BIOS Token for setting CPU Performance configuration. 
* `cpu_power_management`:(string) BIOS Token for setting Power Technology configuration. 
* `dcpmm_firmware_downgrade`:(string) BIOS Token for setting DCPMM Firmware Downgrade configuration. 
* `demand_scrub`:(string) BIOS Token for setting Demand Scrub configuration. 
* `description`:(string) Description of the policy. 
* `direct_cache_access`:(string) BIOS Token for setting Direct Cache Access Support configuration. 
* `dram_clock_throttling`:(string) BIOS Token for setting DRAM Clock Throttling configuration. 
* `dram_refresh_rate`:(string) BIOS Token for setting DRAM Refresh Rate configuration. 
* `energy_efficient_turbo`:(string) BIOS Token for setting Energy Efficient Turbo configuration. 
* `eng_perf_tuning`:(string) BIOS Token for setting Energy Performance Tuning configuration. 
* `enhanced_intel_speed_step_tech`:(string) BIOS Token for setting Enhanced Intel Speedstep (R) Technology configuration. 
* `epp_profile`:(string) BIOS Token for setting EPP Profile configuration. 
* `execute_disable_bit`:(string) BIOS Token for setting Execute Disable Bit configuration. 
* `extended_apic`:(string) BIOS Token for setting Local X2 Apic configuration. 
* `flow_control`:(string) BIOS Token for setting Flow Control configuration. 
* `frb2enable`:(string) BIOS Token for setting FRB-2 Timer configuration. 
* `hardware_prefetch`:(string) BIOS Token for setting Hardware Prefetcher configuration. 
* `hwpm_enable`:(string) BIOS Token for setting CPU Hardware Power Management configuration. 
* `imc_interleave`:(string) BIOS Token for setting IMC Interleaving configuration. 
* `intel_hyper_threading_tech`:(string) BIOS Token for setting Intel HyperThreading Tech configuration. 
* `intel_speed_select`:(string) BIOS Token for setting Intel Speed Select configuration. 
* `intel_turbo_boost_tech`:(string) BIOS Token for setting Intel Turbo Boost Tech configuration. 
* `intel_virtualization_technology`:(string) BIOS Token for setting Intel (R) VT configuration. 
* `intel_vt_for_directed_io`:(string) BIOS Token for setting Intel VT for directed IO configuration. 
* `intel_vtd_coherency_support`:(string) BIOS Token for setting Intel (R) VT-d Coherency Support configuration. 
* `intel_vtd_interrupt_remapping`:(string) BIOS Token for setting Intel (R) VT-d Interrupt Remapping configuration. 
* `intel_vtd_pass_through_dma_support`:(string) BIOS Token for setting Intel (R) VT-d PassThrough DMA support configuration. 
* `intel_vtdats_support`:(string) BIOS Token for setting Intel VTD ATS support configuration. 
* `ioh_error_enable`:(string) BIOS Token for setting IIO Error Enable configuration. 
* `ioh_resource`:(string) BIOS Token for setting IOH Resource Allocation configuration. 
* `ip_prefetch`:(string) BIOS Token for setting DCU IP Prefetcher configuration. 
* `ipv4pxe`:(string) BIOS Token for setting IPv4 PXE Support configuration. 
* `ipv6pxe`:(string) BIOS Token for setting IPV6 PXE Support configuration. 
* `kti_prefetch`:(string) BIOS Token for setting KTI Prefetch configuration. 
* `legacy_os_redirection`:(string) BIOS Token for setting Legacy OS redirection configuration. 
* `legacy_usb_support`:(string) BIOS Token for setting Legacy USB Support configuration. 
* `llc_prefetch`:(string) BIOS Token for setting LLC Prefetch configuration. 
* `lom_port0state`:(string) BIOS Token for setting LOM Port 0 OptionROM configuration. 
* `lom_port1state`:(string) BIOS Token for setting LOM Port 1 OptionROM configuration. 
* `lom_port2state`:(string) BIOS Token for setting LOM Port 2 OptionROM configuration. 
* `lom_port3state`:(string) BIOS Token for setting LOM Port 3 OptionROM configuration. 
* `lom_ports_all_state`:(string) BIOS Token for setting All Onboard LOM Ports configuration. 
* `lv_ddr_mode`:(string) BIOS Token for setting Low Voltage DDR Mode configuration. 
* `make_device_non_bootable`:(string) BIOS Token for setting Make Device Non Bootable configuration. 
* `memory_inter_leave`:(string) BIOS Token for setting Memory Interleaving configuration. 
* `memory_mapped_io_above4gb`:(string) BIOS Token for setting Memory mapped IO above 4GiB configuration. 
* `memory_size_limit`:(string) BIOS Token for setting Memory Size Limit in GiB configuration (0-65535). 
* `mirroring_mode`:(string) BIOS Token for setting Mirroring Mode configuration. 
* `mmcfg_base`:(string) BIOS Token for setting MMCFG BASE configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `network_stack`:(string) BIOS Token for setting Network Stack configuration. 
* `numa_optimized`:(string) BIOS Token for setting NUMA optimized configuration. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `onboard10gbit_lom`:(string) BIOS Token for setting Onboard 10Gbit LOM configuration. 
* `onboard_gbit_lom`:(string) BIOS Token for setting Onboard Gbit LOM configuration. 
* `onboard_scu_storage_support`:(string) BIOS Token for setting Onboard SCU Storage Support configuration. 
* `onboard_scu_storage_sw_stack`:(string) BIOS Token for setting Onboard SCU Storage SW Stack configuration. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `os_boot_watchdog_timer`:(string) BIOS Token for setting OS Boot Watchdog Timer configuration. 
* `os_boot_watchdog_timer_policy`:(string) BIOS Token for setting OS Boot Watchdog Timer Policy configuration. 
* `os_boot_watchdog_timer_timeout`:(string) BIOS Token for setting OS Boot Watchdog Timer Timeout configuration. 
* `out_of_band_mgmt_port`:(string) BIOS Token for setting Out-of-Band Mgmt Port configuration. 
* `package_cstate_limit`:(string) BIOS Token for setting Package C State Limit configuration. 
* `partial_mirror_mode_config`:(string) BIOS Token for setting Partial Memory Mirror Mode configuration. 
* `partial_mirror_percent`:(string) BIOS Token for setting Partial Mirror percentage configuration (0.00-50.00). 
* `partial_mirror_value1`:(string) BIOS Token for setting Partial Mirror1 Size in GiB configuration (0-65535). 
* `partial_mirror_value2`:(string) BIOS Token for setting Partial Mirror2 Size in GiB configuration (0-65535). 
* `partial_mirror_value3`:(string) BIOS Token for setting Partial Mirror3 Size in GiB configuration (0-65535). 
* `partial_mirror_value4`:(string) BIOS Token for setting Partial Mirror4 Size in GiB configuration (0-65535). 
* `patrol_scrub`:(string) BIOS Token for setting Patrol Scrub configuration. 
* `patrol_scrub_duration`:(string) BIOS Token for setting Patrol Scrub Interval configuration. 
* `pc_ie_ras_support`:(string) BIOS Token for setting PCIe RAS Support configuration. 
* `pc_ie_ssd_hot_plug_support`:(string) BIOS Token for setting NVMe SSD Hot-Plug Support configuration. 
* `pch_usb30mode`:(string) BIOS Token for setting xHCI Mode configuration. 
* `pci_option_ro_ms`:(string) BIOS Token for setting All PCIe Slots OptionROM configuration. 
* `pci_rom_clp`:(string) BIOS Token for setting PCI ROM CLP configuration. 
* `pop_support`:(string) BIOS Token for setting Power ON Password configuration. 
* `post_error_pause`:(string) BIOS Token for setting POST Error Pause configuration. 
* `processor_c1e`:(string) BIOS Token for setting Processor C1E configuration. 
* `processor_c3report`:(string) BIOS Token for setting Processor C3 Report configuration. 
* `processor_c6report`:(string) BIOS Token for setting Processor C6 Report configuration. 
* `processor_cstate`:(string) BIOS Token for setting CPU C State configuration. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `psata`:(string) BIOS Token for setting P-SATA mode configuration. 
* `pstate_coord_type`:(string) BIOS Token for setting P-STATE Coordination configuration. 
* `putty_key_pad`:(string) BIOS Token for setting Putty KeyPad configuration. 
* `pwr_perf_tuning`:(string) BIOS Token for setting Power Performance Tuning configuration. 
* `qpi_link_frequency`:(string) BIOS Token for setting QPI Link Frequency Select configuration. 
* `qpi_snoop_mode`:(string) BIOS Token for setting QPI Snoop Mode configuration. 
* `rank_inter_leave`:(string) BIOS Token for setting Rank Interleaving configuration. 
* `redirection_after_post`:(string) BIOS Token for setting Redirection After BIOS POST configuration. 
* `sata_mode_select`:(string) BIOS Token for setting SATA mode configuration. 
* `select_memory_ras_configuration`:(string) BIOS Token for setting SelectMemory RAS configuration configuration. 
* `select_ppr_type`:(string) BIOS Token for setting Select PPR Type configuration. 
* `serial_port_aenable`:(string) BIOS Token for setting Serial A Enable configuration. 
* `single_pctl_enable`:(string) BIOS Token for setting Single PCTL configuration. 
* `slot10link_speed`:(string) BIOS Token for setting PCIe Slot:10 Link Speed configuration. 
* `slot10state`:(string) BIOS Token for setting Slot 10 State configuration. 
* `slot11link_speed`:(string) BIOS Token for setting PCIe Slot:11 Link Speed configuration. 
* `slot11state`:(string) BIOS Token for setting Slot 11 State configuration. 
* `slot12link_speed`:(string) BIOS Token for setting PCIe Slot:12 Link Speed configuration. 
* `slot12state`:(string) BIOS Token for setting Slot 12 State configuration. 
* `slot13state`:(string) BIOS Token for setting PCIe Slot 13 OptionROM configuration. 
* `slot14state`:(string) BIOS Token for setting PCIe Slot 14 OptionROM configuration. 
* `slot1link_speed`:(string) BIOS Token for setting PCIe Slot:1 Link Speed configuration. 
* `slot1state`:(string) BIOS Token for setting Slot 1 State configuration. 
* `slot2link_speed`:(string) BIOS Token for setting PCIe Slot:2 Link Speed configuration. 
* `slot2state`:(string) BIOS Token for setting Slot 2 State configuration. 
* `slot3link_speed`:(string) BIOS Token for setting PCIe Slot:3 Link Speed configuration. 
* `slot3state`:(string) BIOS Token for setting Slot 3 State configuration. 
* `slot4link_speed`:(string) BIOS Token for setting PCIe Slot:4 Link Speed configuration. 
* `slot4state`:(string) BIOS Token for setting Slot 4 State configuration. 
* `slot5link_speed`:(string) BIOS Token for setting PCIe Slot:5 Link Speed configuration. 
* `slot5state`:(string) BIOS Token for setting Slot 5 State configuration. 
* `slot6link_speed`:(string) BIOS Token for setting PCIe Slot:6 Link Speed configuration. 
* `slot6state`:(string) BIOS Token for setting Slot 6 State configuration. 
* `slot7link_speed`:(string) BIOS Token for setting PCIe Slot:7 Link Speed configuration. 
* `slot7state`:(string) BIOS Token for setting Slot 7 State configuration. 
* `slot8link_speed`:(string) BIOS Token for setting PCIe Slot:8 Link Speed configuration. 
* `slot8state`:(string) BIOS Token for setting Slot 8 State configuration. 
* `slot9link_speed`:(string) BIOS Token for setting PCIe Slot:9 Link Speed configuration. 
* `slot9state`:(string) BIOS Token for setting Slot 9 State configuration. 
* `slot_flom_link_speed`:(string) BIOS Token for setting PCIe Slot:FLOM Link Speed configuration. 
* `slot_front_nvme1link_speed`:(string) BIOS Token for setting PCIe Slot:Front Nvme1 Link Speed configuration. 
* `slot_front_nvme2link_speed`:(string) BIOS Token for setting PCIe Slot:Front Nvme2 Link Speed configuration. 
* `slot_front_slot5link_speed`:(string) BIOS Token for setting PCIe Slot:Front1 Link Speed configuration. 
* `slot_front_slot6link_speed`:(string) BIOS Token for setting PCIe Slot:Front2 Link Speed configuration. 
* `slot_gpu1state`:(string) BIOS Token for setting GPU1 OptionROM configuration. 
* `slot_gpu2state`:(string) BIOS Token for setting GPU2 OptionROM configuration. 
* `slot_gpu3state`:(string) BIOS Token for setting GPU3 OptionROM configuration. 
* `slot_gpu4state`:(string) BIOS Token for setting GPU4 OptionROM configuration. 
* `slot_gpu5state`:(string) BIOS Token for setting GPU5 OptionROM configuration. 
* `slot_gpu6state`:(string) BIOS Token for setting GPU6 OptionROM configuration. 
* `slot_gpu7state`:(string) BIOS Token for setting GPU7 OptionROM configuration. 
* `slot_gpu8state`:(string) BIOS Token for setting GPU8 OptionROM configuration. 
* `slot_hba_link_speed`:(string) BIOS Token for setting PCIe Slot:HBA Link Speed configuration. 
* `slot_hba_state`:(string) BIOS Token for setting PCIe Slot:HBA OptionROM configuration. 
* `slot_lom1link`:(string) BIOS Token for setting PCIe LOM:1 Link configuration. 
* `slot_lom2link`:(string) BIOS Token for setting PCIe LOM:2 Link configuration. 
* `slot_mezz_state`:(string) BIOS Token for setting Slot Mezz State configuration. 
* `slot_mlom_link_speed`:(string) BIOS Token for setting PCIe Slot:MLOM Link Speed configuration. 
* `slot_mlom_state`:(string) BIOS Token for setting PCIe Slot MLOM OptionROM configuration. 
* `slot_mraid_link_speed`:(string) BIOS Token for setting MRAID Link Speed configuration. 
* `slot_mraid_state`:(string) BIOS Token for setting PCIe Slot MRAID OptionROM configuration. 
* `slot_n10state`:(string) BIOS Token for setting PCIe Slot N10 OptionROM configuration. 
* `slot_n11state`:(string) BIOS Token for setting PCIe Slot N11 OptionROM configuration. 
* `slot_n12state`:(string) BIOS Token for setting PCIe Slot N12 OptionROM configuration. 
* `slot_n13state`:(string) BIOS Token for setting PCIe Slot N13 OptionROM configuration. 
* `slot_n14state`:(string) BIOS Token for setting PCIe Slot N14 OptionROM configuration. 
* `slot_n15state`:(string) BIOS Token for setting PCIe Slot N15 OptionROM configuration. 
* `slot_n16state`:(string) BIOS Token for setting PCIe Slot N16 OptionROM configuration. 
* `slot_n17state`:(string) BIOS Token for setting PCIe Slot N17 OptionROM configuration. 
* `slot_n18state`:(string) BIOS Token for setting PCIe Slot N18 OptionROM configuration. 
* `slot_n19state`:(string) BIOS Token for setting PCIe Slot N19 OptionROM configuration. 
* `slot_n1state`:(string) BIOS Token for setting PCIe Slot N1 OptionROM configuration. 
* `slot_n20state`:(string) BIOS Token for setting PCIe Slot N20 OptionROM configuration. 
* `slot_n21state`:(string) BIOS Token for setting PCIe Slot N21 OptionROM configuration. 
* `slot_n22state`:(string) BIOS Token for setting PCIe Slot N22 OptionROM configuration. 
* `slot_n23state`:(string) BIOS Token for setting PCIe Slot N23 OptionROM configuration. 
* `slot_n24state`:(string) BIOS Token for setting PCIe Slot N24 OptionROM configuration. 
* `slot_n2state`:(string) BIOS Token for setting PCIe Slot N2 OptionROM configuration. 
* `slot_n3state`:(string) BIOS Token for setting PCIe Slot N3 OptionROM configuration. 
* `slot_n4state`:(string) BIOS Token for setting PCIe Slot N4 OptionROM configuration. 
* `slot_n5state`:(string) BIOS Token for setting PCIe Slot N5 OptionROM configuration. 
* `slot_n6state`:(string) BIOS Token for setting PCIe Slot N6 OptionROM configuration. 
* `slot_n7state`:(string) BIOS Token for setting PCIe Slot N7 OptionROM configuration. 
* `slot_n8state`:(string) BIOS Token for setting PCIe Slot N8 OptionROM configuration. 
* `slot_n9state`:(string) BIOS Token for setting PCIe Slot N9 OptionROM configuration. 
* `slot_raid_link_speed`:(string) BIOS Token for setting RAID Link Speed configuration. 
* `slot_raid_state`:(string) BIOS Token for setting PCIe Slot RAID OptionROM configuration. 
* `slot_rear_nvme1link_speed`:(string) BIOS Token for setting PCIe Slot:Rear Nvme1 Link Speed configuration. 
* `slot_rear_nvme1state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 1 OptionROM configuration. 
* `slot_rear_nvme2link_speed`:(string) BIOS Token for setting PCIe Slot:Rear Nvme2 Link Speed configuration. 
* `slot_rear_nvme2state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 2 OptionROM configuration. 
* `slot_rear_nvme3state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 3 OptionROM configuration. 
* `slot_rear_nvme4state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 4 OptionROM configuration. 
* `slot_rear_nvme5state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 5 OptionROM configuration. 
* `slot_rear_nvme6state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 6 OptionROM configuration. 
* `slot_rear_nvme7state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 7 OptionROM configuration. 
* `slot_rear_nvme8state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 8 OptionROM configuration. 
* `slot_riser1link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Link Speed configuration. 
* `slot_riser1slot1link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Slot1 Link Speed configuration. 
* `slot_riser1slot2link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Slot2 Link Speed configuration. 
* `slot_riser1slot3link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Slot3 Link Speed configuration. 
* `slot_riser2link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Link Speed configuration. 
* `slot_riser2slot4link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Slot4 Link Speed configuration. 
* `slot_riser2slot5link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Slot5 Link Speed configuration. 
* `slot_riser2slot6link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Slot6 Link Speed configuration. 
* `slot_sas_state`:(string) BIOS Token for setting PCIe Slot:SAS OptionROM configuration. 
* `slot_ssd_slot1link_speed`:(string) BIOS Token for setting PCIe Slot:FrontSSD1 Link Speed configuration. 
* `slot_ssd_slot2link_speed`:(string) BIOS Token for setting PCIe Slot:FrontSSD2 Link Speed configuration. 
* `smee`:(string) BIOS Token for setting SMEE configuration. 
* `smt_mode`:(string) BIOS Token for setting SMT Mode configuration. 
* `snc`:(string) BIOS Token for setting Sub Numa Clustering configuration. 
* `sparing_mode`:(string) BIOS Token for setting Sparing Mode configuration. 
* `sr_iov`:(string) BIOS Token for setting SR-IOV Support configuration. 
* `streamer_prefetch`:(string) BIOS Token for setting DCU Streamer Prefetch configuration. 
* `svm_mode`:(string) BIOS Token for setting SVM Mode configuration. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `terminal_type`:(string) BIOS Token for setting Terminal Type configuration. 
* `tpm_control`:(string) BIOS Token for setting Trusted Platform Module State configuration. 
* `tpm_support`:(string) BIOS Token for setting TPM Support configuration. 
* `txt_support`:(string) BIOS Token for setting Intel Trusted Execution Technology Support configuration. 
* `ucsm_boot_order_rule`:(string) BIOS Token for setting Boot Order Rules configuration. 
* `usb_emul6064`:(string) BIOS Token for setting Port 60/64 Emulation configuration. 
* `usb_port_front`:(string) BIOS Token for setting USB Port Front configuration. 
* `usb_port_internal`:(string) BIOS Token for setting USB Port Internal configuration. 
* `usb_port_kvm`:(string) BIOS Token for setting USB Port KVM configuration. 
* `usb_port_rear`:(string) BIOS Token for setting USB Port Rear configuration. 
* `usb_port_sd_card`:(string) BIOS Token for setting USB Port SD Card configuration. 
* `usb_port_vmedia`:(string) BIOS Token for setting USB Port VMedia configuration. 
* `usb_xhci_support`:(string) BIOS Token for setting XHCI Legacy Support configuration. 
* `vga_priority`:(string) BIOS Token for setting VGA Priority configuration. 
* `vmd_enable`:(string) BIOS Token for setting VMD Enablement configuration. 
* `work_load_config`:(string) BIOS Token for setting Workload Configuration configuration. 
* `xpt_prefetch`:(string) BIOS Token for setting XPT Prefetch configuration. 
