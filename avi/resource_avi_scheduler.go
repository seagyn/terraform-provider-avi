// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceSchedulerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"backup_config_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"end_date_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"frequency": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"frequency_unit": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"run_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"run_script_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"scheduler_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SCHEDULER_ACTION_BACKUP",
		},
		"start_date_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviScheduler() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSchedulerCreate,
		Read:   ResourceAviSchedulerRead,
		Update: resourceAviSchedulerUpdate,
		Delete: resourceAviSchedulerDelete,
		Schema: ResourceSchedulerSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSchedulerImporter,
		},
	}
}

func ResourceSchedulerImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSchedulerSchema()
	return ResourceImporter(d, m, "scheduler", s)
}

func ResourceAviSchedulerRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	err := APIRead(d, meta, "scheduler", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSchedulerCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	err := APICreateOrUpdate(d, meta, "scheduler", s)
	if err == nil {
		err = ResourceAviSchedulerRead(d, meta)
	}
	return err
}

func resourceAviSchedulerUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSchedulerSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "scheduler", s)
	if err == nil {
		err = ResourceAviSchedulerRead(d, meta)
	}
	return err
}

func resourceAviSchedulerDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "scheduler"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSchedulerDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
