// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// VersionInfo version info
// swagger:model VersionInfo
type VersionInfo struct {

	// Identifies the datastore table . Field introduced in 17.2.7.
	DsName *string `json:"ds_name,omitempty"`

	// Name of the object.
	Name *string `json:"name,omitempty"`

	//  Enum options - GSLB_NONE, GSLB_CREATE, GSLB_UPDATE, GSLB_DELETE, GSLB_PURGE, GSLB_DECL.
	Ops *string `json:"ops,omitempty"`

	// Unique object identifier of the object.
	UUID *string `json:"uuid,omitempty"`
}
