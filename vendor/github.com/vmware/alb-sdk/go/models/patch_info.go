// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// PatchInfo patch info
// swagger:model PatchInfo
type PatchInfo struct {

	// Patch type describes the controller or se patch type. Field introduced in 18.2.6.
	PatchType *string `json:"patch_type,omitempty"`

	// This variable tells whether reboot has to be performed. Field introduced in 18.2.6.
	Reboot *bool `json:"reboot,omitempty"`

	// This variable is for full list of patch reboot details. Field introduced in 18.2.8, 20.1.1.
	RebootList []*RebootData `json:"reboot_list,omitempty"`
}
