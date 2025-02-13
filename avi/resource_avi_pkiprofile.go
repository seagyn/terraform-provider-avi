// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourcePKIProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ca_certs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSSLCertificateSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"crl_check": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"crls": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCRLSchema(),
		},
		"ignore_peer_chain": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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
		"validate_only_leaf_crl": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
	}
}

func resourceAviPKIProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPKIProfileCreate,
		Read:   ResourceAviPKIProfileRead,
		Update: resourceAviPKIProfileUpdate,
		Delete: resourceAviPKIProfileDelete,
		Schema: ResourcePKIProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePKIProfileImporter,
		},
	}
}

func ResourcePKIProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePKIProfileSchema()
	return ResourceImporter(d, m, "pkiprofile", s)
}

func ResourceAviPKIProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePKIProfileSchema()
	err := APIRead(d, meta, "pkiprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPKIProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePKIProfileSchema()
	err := APICreateOrUpdate(d, meta, "pkiprofile", s)
	if err == nil {
		err = ResourceAviPKIProfileRead(d, meta)
	}
	return err
}

func resourceAviPKIProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePKIProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "pkiprofile", s)
	if err == nil {
		err = ResourceAviPKIProfileRead(d, meta)
	}
	return err
}

func resourceAviPKIProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "pkiprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviPKIProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
