// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// SeMigrateEventDetails se migrate event details
// swagger:model SeMigrateEventDetails
type SeMigrateEventDetails struct {

	// Number of num_vs.
	NumVs *int32 `json:"num_vs,omitempty"`

	// reason of SeMigrateEventDetails.
	Reason []string `json:"reason,omitempty"`

	// se_name of SeMigrateEventDetails.
	SeName *string `json:"se_name,omitempty"`

	// Unique object identifier of se.
	// Required: true
	SeUUID *string `json:"se_uuid"`

	// vs_name of SeMigrateEventDetails.
	VsName *string `json:"vs_name,omitempty"`

	// Unique object identifier of vs.
	VsUUID *string `json:"vs_uuid,omitempty"`
}
