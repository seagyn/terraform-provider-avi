<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_upgradestatusinfo"
sidebar_current: "docs-avi-resource-upgradestatusinfo"
description: |-
  Creates and manages Avi UpgradeStatusInfo.
---

# avi_upgradestatusinfo

The UpgradeStatusInfo resource allows the creation and management of Avi UpgradeStatusInfo

## Example Usage

```hcl
resource "avi_upgradestatusinfo" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `after_reboot_rollback_fnc` - (Optional) Backward compatible abort function name. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `after_reboot_task_name` - (Optional) Backward compatible task dict name. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `clean` - (Optional) Flag for clean installation. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `duration` - (Optional) Duration of upgrade operation in seconds. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_patch_rollback` - (Optional) Check if the patch rollback is possible on this node. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_rollback` - (Optional) Check if the rollback is possible on this node. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `end_time` - (Optional) End time of upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enqueue_time` - (Optional) Enqueue time of upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `fips_mode` - (Optional) Fips mode for the entire system. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `history` - (Optional) Record of past operations on this node. Field introduced in 20.1.4. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `image_path` - (Optional) Image path of current base image. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `image_ref` - (Optional) Image uuid for identifying the current base image. It is a reference to an object of type image. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - (Optional) Name of the system such as cluster name, se group name and se name. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `node_type` - (Optional) Type of the system such as controller_cluster, se_group or se. Enum options - NODE_CONTROLLER_CLUSTER, NODE_SE_GROUP, NODE_SE_TYPE. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `obj_cloud_ref` - (Optional) Cloud that this object belongs to. It is a reference to an object of type cloud. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `params` - (Optional) Parameters associated with the upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `patch_image_path` - (Optional) Image path of current patch image. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `patch_image_ref` - (Optional) Image uuid for identifying the current patch.example  base-image is 18.2.6 and a patch 6p1 is applied, then this field will indicate the 6p1 value. It is a reference to an object of type image. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `patch_list` - (Optional) List of patches applied to this node. Example  base-image is 18.2.6 and a patch 6p1 is applied, then a patch 6p5 applied. This field will indicate the [{'6p1', '6p1_image_uuid'}, {'6p5', '6p5_image_uuid'}] value. Field introduced in 18.2.8, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `patch_reboot` - (Optional) Flag for patch op with reboot. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `patch_version` - (Optional) Current patch version applied to this node. Example  base-image is 18.2.6 and a patch 6p1 is applied, then this field will indicate the 6p1 value. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `prev_image_path` - (Optional) Image path of previous base image. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `prev_patch_image_path` - (Optional) Image path of previous patch image. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `previous_image_ref` - (Optional) Image uuid for identifying previous base image.example  base-image was 18.2.5 and an upgrade was done to 18.2.6, then this field will indicate the 18.2.5 value. It is a reference to an object of type image. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `previous_patch_image_ref` - (Optional) Image uuid for identifying previous patch.example  base-image was 18.2.6 with a patch 6p1. Upgrade was initiated to 18.2.8 with patch 8p1. The previous_image field will contain 18.2.6 and this field will indicate the 6p1 value. It is a reference to an object of type image. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `previous_patch_list` - (Optional) List of patches applied to this node on previous major version. Field introduced in 18.2.8, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `previous_patch_version` - (Optional) Previous patch version applied to this node.example  base-image was 18.2.6 with a patch 6p1. Upgrade was initiated to 18.2.8 with patch 8p1. The previous_image field will contain 18.2.6 and this field will indicate the 6p1 value. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `previous_version` - (Optional) Previous version prior to upgrade.example  base-image was 18.2.5 and an upgrade was done to 18.2.6, then this field will indicate the 18.2.5 value. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `progress` - (Optional) Upgrade operations progress which holds value between 0-100. Allowed values are 0-100. Field introduced in 18.2.8, 20.1.1. Unit is percent. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `reason` - (Optional) Descriptive reason for the upgrade state. Field introduced in 22.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_patch_image_path` - (Optional) Image path of se patch image.(required in case of reimage and upgrade + patch). Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_patch_image_ref` - (Optional) Image uuid for identifying the current se patch required in case of system upgrade(re-image) with se patch. It is a reference to an object of type image. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_upgrade_events` - (Optional) Serviceenginegroup upgrade errors. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `seg_params` - (Optional) Se_patch may be different from the controller_patch. It has to be saved in the journal for subsequent consumption. The segroup params will be saved in the controller entry as seg_params. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `seg_status` - (Optional) Detailed segroup status. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `start_time` - (Optional) Start time of upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `state` - (Optional) Current status of the upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `statediff_ref` - (Optional) Record of pre/post snapshot captured for current upgrade operation. It is a reference to an object of type statediffoperation. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `system` - (Optional) Flag is set only in the cluster if the upgrade is initiated as a system-upgrade. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tasks_completed` - (Optional) Completed set of tasks in the upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `total_tasks` - (Optional) Total number of tasks in the upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `upgrade_events` - (Optional) Events performed for upgrade operation. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `upgrade_ops` - (Optional) Upgrade operations requested. Enum options - UPGRADE, PATCH, ROLLBACK, ROLLBACKPATCH, SEGROUP_RESUME, EVAL_UPGRADE, EVAL_PATCH, EVAL_ROLLBACK, EVAL_ROLLBACKPATCH, EVAL_SEGROUP_RESUME. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `upgrade_readiness` - (Optional) Upgrade readiness check execution detail. Field introduced in 22.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `version` - (Optional) Current base image applied to this node. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the system such as cluster, se group and se. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

