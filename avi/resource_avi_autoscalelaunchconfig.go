// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceAutoScaleLaunchConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"image_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"mesos": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAutoScaleMesosSettingsSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"openstack": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAutoScaleOpenStackSettingsSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_external_asg": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAutoScaleLaunchConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAutoScaleLaunchConfigCreate,
		Read:   ResourceAviAutoScaleLaunchConfigRead,
		Update: resourceAviAutoScaleLaunchConfigUpdate,
		Delete: resourceAviAutoScaleLaunchConfigDelete,
		Schema: ResourceAutoScaleLaunchConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAutoScaleLaunchConfigImporter,
		},
	}
}

func ResourceAutoScaleLaunchConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAutoScaleLaunchConfigSchema()
	return ResourceImporter(d, m, "autoscalelaunchconfig", s)
}

func ResourceAviAutoScaleLaunchConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAutoScaleLaunchConfigSchema()
	err := APIRead(d, meta, "autoscalelaunchconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAutoScaleLaunchConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAutoScaleLaunchConfigSchema()
	err := APICreateOrUpdate(d, meta, "autoscalelaunchconfig", s)
	if err == nil {
		err = ResourceAviAutoScaleLaunchConfigRead(d, meta)
	}
	return err
}

func resourceAviAutoScaleLaunchConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAutoScaleLaunchConfigSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "autoscalelaunchconfig", s)
	if err == nil {
		err = ResourceAviAutoScaleLaunchConfigRead(d, meta)
	}
	return err
}

func resourceAviAutoScaleLaunchConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "autoscalelaunchconfig"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviAutoScaleLaunchConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
