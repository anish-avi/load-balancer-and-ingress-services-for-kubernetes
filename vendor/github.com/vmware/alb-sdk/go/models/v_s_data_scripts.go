// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// VSDataScripts v s data scripts
// swagger:model VSDataScripts
type VSDataScripts struct {

	// Index of the virtual service datascript collection.
	// Required: true
	Index *int32 `json:"index"`

	// UUID of the virtual service datascript collection. It is a reference to an object of type VSDataScriptSet.
	// Required: true
	VsDatascriptSetRef *string `json:"vs_datascript_set_ref"`
}
