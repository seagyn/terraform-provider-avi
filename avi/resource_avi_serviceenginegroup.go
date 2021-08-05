// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"time"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceServiceEngineGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accelerated_networking": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"active_standby": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"aggressive_failure_detection": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"algo": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PLACEMENT_ALGO_PACKED",
		},
		"allow_burst": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"app_cache_percent": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"app_learning_memory_percent": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"archive_shm_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "8",
			ValidateFunc: validateInteger,
		},
		"async_ssl": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"async_ssl_threads": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"auto_rebalance": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"auto_rebalance_capacity_per_se": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"auto_rebalance_criteria": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"auto_rebalance_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"auto_redistribute_active_standby_load": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"bgp_state_update_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"buffer_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"compress_ip_rules_for_each_ns_subnet": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"config_debugs_on_all_cores": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"connection_memory_percentage": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "50",
			ValidateFunc: validateInteger,
		},
		"core_shm_app_cache": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"core_shm_app_learning": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"cpu_reserve": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"cpu_socket_affinity": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"custom_securitygroups_data": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"custom_securitygroups_mgmt": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"custom_tag": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema(),
		},
		"data_network_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datascript_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1000000",
			ValidateFunc: validateInteger,
		},
		"dedicated_dispatcher_core": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"disable_avi_securitygroups": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"disable_csum_offloads": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"disable_gro": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"disable_se_memory_check": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"disable_tso": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"disk_per_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"distribute_load_active_standby": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"distribute_queues": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"distribute_vnics": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_gratarp_permanent": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_hsm_priming": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_multi_lb": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_pcap_tx_ring": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"ephemeral_portrange_end": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"ephemeral_portrange_start": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"extra_config_multiplier": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateFloat,
		},
		"extra_shared_config_memory": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"flow_table_new_syn_max_entries": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"free_list_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1024",
			ValidateFunc: validateInteger,
		},
		"gratarp_permanent_periodicity": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"ha_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "HA_MODE_SHARED",
		},
		"hardwaresecuritymodulegroup_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"heap_minimum_config_memory": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "8",
			ValidateFunc: validateInteger,
		},
		"hm_on_standby": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"host_attribute_key": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"host_attribute_value": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"host_gateway_monitor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"hypervisor": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ignore_rtt_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5000",
			ValidateFunc: validateInteger,
		},
		"ingress_access_data": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SG_INGRESS_ACCESS_ALL",
		},
		"ingress_access_mgmt": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SG_INGRESS_ACCESS_ALL",
		},
		"instance_flavor": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"iptables": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIptableRuleSetSchema(),
		},
		"labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"least_load_core_selection": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"license_tier": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"license_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"log_disksz": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10000",
			ValidateFunc: validateInteger,
		},
		"max_concurrent_external_hm": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"max_cpu_usage": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "80",
			ValidateFunc: validateInteger,
		},
		"max_memory_per_mempool": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "64",
			ValidateFunc: validateInteger,
		},
		"max_public_ips_per_lb": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "30",
			ValidateFunc: validateInteger,
		},
		"max_queues_per_vnic": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"max_rules_per_lb": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "150",
			ValidateFunc: validateInteger,
		},
		"max_scaleout_per_vs": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"max_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"max_vs_per_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"mem_reserve": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"memory_for_config_update": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "15",
			ValidateFunc: validateInteger,
		},
		"memory_per_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2048",
			ValidateFunc: validateInteger,
		},
		"mgmt_network_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"mgmt_subnet": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"min_cpu_usage": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "30",
			ValidateFunc: validateInteger,
		},
		"min_scaleout_per_vs": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"min_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"minimum_connection_memory": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"n_log_streaming_threads": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"non_significant_log_throttle": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"num_dispatcher_cores": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"num_flow_cores_sum_changes_to_ignore": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "8",
			ValidateFunc: validateInteger,
		},
		"openstack_availability_zones": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"openstack_mgmt_network_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"openstack_mgmt_network_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"os_reserved_memory": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"pcap_tx_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PCAP_TX_AUTO",
		},
		"pcap_tx_ring_rd_balancing_factor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"per_app": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"per_vs_admission_control": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"placement_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PLACEMENT_MODE_AUTO",
		},
		"realtime_se_metrics": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceMetricsRealTimeUpdateSchema(),
		},
		"reboot_on_panic": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"se_bandwidth_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_delayed_flow_delete": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"se_deprovision_delay": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "120",
			ValidateFunc: validateInteger,
		},
		"se_dos_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDosThresholdProfileSchema(),
		},
		"se_dp_vnic_queue_stall_event_sleep": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_dp_vnic_queue_stall_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2000",
			ValidateFunc: validateInteger,
		},
		"se_dp_vnic_queue_stall_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10000",
			ValidateFunc: validateInteger,
		},
		"se_dp_vnic_restart_on_queue_stall_count": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"se_dp_vnic_stall_se_restart_window": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3600",
			ValidateFunc: validateInteger,
		},
		"se_dpdk_pmd": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_flow_probe_retries": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateInteger,
		},
		"se_flow_probe_retry_timer": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "40",
			ValidateFunc: validateInteger,
		},
		"se_ipc_udp_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1500",
			ValidateFunc: validateInteger,
		},
		"se_kni_burst_factor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_lro": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"se_mp_ring_retry_count": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "500",
			ValidateFunc: validateInteger,
		},
		"se_mtu": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"se_name_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Avi",
		},
		"se_pcap_lookahead": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"se_pcap_pkt_count": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_pcap_pkt_sz": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "69632",
			ValidateFunc: validateInteger,
		},
		"se_pcap_qdisc_bypass": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"se_pcap_reinit_frequency": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_pcap_reinit_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_probe_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "7",
			ValidateFunc: validateInteger,
		},
		"se_remote_punt_udp_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1501",
			ValidateFunc: validateInteger,
		},
		"se_rl_prop": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateLimiterPropertiesSchema(),
		},
		"se_rum_sampling_nav_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"se_rum_sampling_nav_percent": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"se_rum_sampling_res_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateInteger,
		},
		"se_rum_sampling_res_percent": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"se_sb_dedicated_core": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"se_sb_threads": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"se_thread_multiplier": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"se_tracert_port_range": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePortRangeSchema(),
		},
		"se_tunnel_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_tunnel_udp_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1550",
			ValidateFunc: validateInteger,
		},
		"se_tx_batch_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "64",
			ValidateFunc: validateInteger,
		},
		"se_txq_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2048",
			ValidateFunc: validateInteger,
		},
		"se_udp_encap_ipc": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_use_dpdk": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"se_vs_hb_max_pkts_in_batch": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "64",
			ValidateFunc: validateInteger,
		},
		"se_vs_hb_max_vs_in_pkt": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "256",
			ValidateFunc: validateInteger,
		},
		"self_se_election": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"service_ip6_subnets": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"service_ip_subnets": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"shm_minimum_config_memory": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"significant_log_throttle": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"ssl_preprocess_sni_hostname": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"transient_shared_memory_max": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "30",
			ValidateFunc: validateInteger,
		},
		"udf_log_throttle": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"use_standard_alb": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vcenter_clusters": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVcenterClustersSchema(),
		},
		"vcenter_datastore_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VCENTER_DATASTORE_ANY",
		},
		"vcenter_datastores": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVcenterDatastoreSchema(),
		},
		"vcenter_datastores_include": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"vcenter_folder": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "AviSeFolder",
		},
		"vcenter_hosts": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVcenterHostsSchema(),
		},
		"vcpus_per_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"vip_asg": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVipAutoscaleGroupSchema(),
		},
		"vs_host_redundancy": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"vs_scalein_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "30",
			ValidateFunc: validateInteger,
		},
		"vs_scalein_timeout_for_upgrade": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "30",
			ValidateFunc: validateInteger,
		},
		"vs_scaleout_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "600",
			ValidateFunc: validateInteger,
		},
		"vs_se_scaleout_additional_wait_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"vs_se_scaleout_ready_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"vs_switchover_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"vss_placement": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVssPlacementSchema(),
		},
		"vss_placement_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"waf_mempool": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"waf_mempool_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "64",
			ValidateFunc: validateInteger,
		},
	}
}

func resourceAviServiceEngineGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServiceEngineGroupCreate,
		Read:   ResourceAviServiceEngineGroupRead,
		Update: resourceAviServiceEngineGroupUpdate,
		Delete: resourceAviServiceEngineGroupDelete,
		Schema: ResourceServiceEngineGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceServiceEngineGroupImporter,
		},
	}
}

func ResourceServiceEngineGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceServiceEngineGroupSchema()
	return ResourceImporter(d, m, "serviceenginegroup", s)
}

func ResourceAviServiceEngineGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	err := APIRead(d, meta, "serviceenginegroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviServiceEngineGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	err := APICreateOrUpdate(d, meta, "serviceenginegroup", s)
	if err == nil {
		err = ResourceAviServiceEngineGroupRead(d, meta)
	}
	return err
}

func resourceAviServiceEngineGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "serviceenginegroup", s)
	if err == nil {
		err = ResourceAviServiceEngineGroupRead(d, meta)
	}
	return err
}

func resourceAviServiceEngineGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serviceenginegroup"
	client := meta.(*clients.AviClient)
	seDeprovisionExtraDelay := 2
	if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
		cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
		cloudPath := "api/cloud/" + cloudUUID
		var robj interface{}
		if err := client.AviSession.Get(cloudPath, &robj); err == nil {
			if vcenterConfig, isVcenterConfig := robj.(map[string]interface{})["vcenter_configuration"]; isVcenterConfig {
				if privilege := vcenterConfig.(map[string]interface{})["privilege"].(string); privilege == "WRITE_ACCESS" {
					seGroupName := d.Get("name").(string)
					cloudName := robj.(map[string]interface{})["name"].(string)
					seDeprovisionDelay := d.Get("se_deprovision_delay").(int) + seDeprovisionExtraDelay
					log.Printf("Waiting for %v minutes to delete SE from SE Group %v of cloud %v", seDeprovisionDelay, seGroupName, cloudName)
					time.Sleep(time.Duration(seDeprovisionDelay) * time.Minute)
				}
			}
		}
	}
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviServiceEngineGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
