// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// DbAppLearningInfo db app learning info
// swagger:model DbAppLearningInfo
type DbAppLearningInfo struct {

	// Application UUID. Combination of Virtualservice UUID and WAF Policy UUID. Field introduced in 20.1.1.
	AppID *string `json:"app_id,omitempty"`

	// Information about various URIs under a application. Field introduced in 20.1.1.
	URIInfo []*URIInfo `json:"uri_info,omitempty"`

	// Virtualserivce UUID. Field introduced in 20.1.1.
	VsUUID *string `json:"vs_uuid,omitempty"`
}
