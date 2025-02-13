// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourcePoolGroupDeploymentPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_disable_old_prod_pools": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"evaluation_duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePGDeploymentRuleSchema(),
		},
		"scheme": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "BLUE_GREEN",
		},
		"target_test_traffic_ratio": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"test_traffic_ratio_rampup": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"webhook_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPoolGroupDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolGroupDeploymentPolicyCreate,
		Read:   ResourceAviPoolGroupDeploymentPolicyRead,
		Update: resourceAviPoolGroupDeploymentPolicyUpdate,
		Delete: resourceAviPoolGroupDeploymentPolicyDelete,
		Schema: ResourcePoolGroupDeploymentPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePoolGroupDeploymentPolicyImporter,
		},
	}
}

func ResourcePoolGroupDeploymentPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePoolGroupDeploymentPolicySchema()
	return ResourceImporter(d, m, "poolgroupdeploymentpolicy", s)
}

func ResourceAviPoolGroupDeploymentPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	err := APIRead(d, meta, "poolgroupdeploymentpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	err := APICreateOrUpdate(d, meta, "poolgroupdeploymentpolicy", s)
	if err == nil {
		err = ResourceAviPoolGroupDeploymentPolicyRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "poolgroupdeploymentpolicy", s)
	if err == nil {
		err = ResourceAviPoolGroupDeploymentPolicyRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "poolgroupdeploymentpolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviPoolGroupDeploymentPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
