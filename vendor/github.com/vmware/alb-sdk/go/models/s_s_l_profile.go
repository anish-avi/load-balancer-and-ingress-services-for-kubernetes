// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// SSLProfile s s l profile
// swagger:model SSLProfile
type SSLProfile struct {

	// UNIX time since epoch in microseconds. Units(MICROSECONDS).
	// Read Only: true
	LastModified *string `json:"_last_modified,omitempty"`

	// Ciphers suites represented as defined by https //www.openssl.org/docs/apps/ciphers.html.
	AcceptedCiphers *string `json:"accepted_ciphers,omitempty"`

	// Set of versions accepted by the server. Minimum of 1 items required.
	AcceptedVersions []*SSLVersion `json:"accepted_versions,omitempty"`

	//  Enum options - TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_GCM_SHA384, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_3DES_EDE_CBC_SHA, TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256, TLS_AES_128_GCM_SHA256. Allowed in Basic(Allowed values- TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_RSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_RSA_WITH_AES_256_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,TLS_RSA_WITH_AES_128_CBC_SHA,TLS_RSA_WITH_AES_256_CBC_SHA,TLS_RSA_WITH_3DES_EDE_CBC_SHA) edition, Essentials(Allowed values- TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_RSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_RSA_WITH_AES_256_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,TLS_RSA_WITH_AES_128_CBC_SHA,TLS_RSA_WITH_AES_256_CBC_SHA,TLS_RSA_WITH_3DES_EDE_CBC_SHA) edition, Enterprise edition.
	CipherEnums []string `json:"cipher_enums,omitempty"`

	// TLS 1.3 Ciphers suites represented as defined by U(https //www.openssl.org/docs/manmaster/man1/ciphers.html). Field introduced in 18.2.6. Allowed in Basic edition, Essentials edition, Enterprise edition. Special default for Basic edition is TLS_AES_256_GCM_SHA384-TLS_AES_128_GCM_SHA256, Essentials edition is TLS_AES_256_GCM_SHA384-TLS_AES_128_GCM_SHA256, Enterprise is TLS_AES_256_GCM_SHA384-TLS_CHACHA20_POLY1305_SHA256-TLS_AES_128_GCM_SHA256.
	Ciphersuites *string `json:"ciphersuites,omitempty"`

	// User defined description for the object.
	Description *string `json:"description,omitempty"`

	// DH Parameters used in SSL. At this time, it is not configurable and is set to 2048 bits.
	Dhparam *string `json:"dhparam,omitempty"`

	// Enable early data processing for TLS1.3 connections. Field introduced in 18.2.6. Allowed in Basic(Allowed values- false) edition, Essentials(Allowed values- false) edition, Enterprise edition.
	EnableEarlyData *bool `json:"enable_early_data,omitempty"`

	// Enable SSL session re-use.
	EnableSslSessionReuse *bool `json:"enable_ssl_session_reuse,omitempty"`

	// Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field deprecated in 20.1.5. Field introduced in 20.1.2. Maximum of 4 items allowed.
	Labels []*KeyValue `json:"labels,omitempty"`

	// List of labels to be used for granular RBAC. Field introduced in 20.1.5. Allowed in Basic edition, Essentials edition, Enterprise edition.
	Markers []*RoleFilterMatchLabel `json:"markers,omitempty"`

	// Name of the object.
	// Required: true
	Name *string `json:"name"`

	// Prefer the SSL cipher ordering presented by the client during the SSL handshake over the one specified in the SSL Profile.
	PreferClientCipherOrdering *bool `json:"prefer_client_cipher_ordering,omitempty"`

	// Send 'close notify' alert message for a clean shutdown of the SSL connection.
	SendCloseNotify *bool `json:"send_close_notify,omitempty"`

	// Placeholder for description of property ssl_rating of obj type SSLProfile field type str  type object
	SslRating *SSLRating `json:"ssl_rating,omitempty"`

	// The amount of time in seconds before an SSL session expires. Unit is SEC.
	SslSessionTimeout *int32 `json:"ssl_session_timeout,omitempty"`

	// Placeholder for description of property tags of obj type SSLProfile field type str  type object
	Tags []*Tag `json:"tags,omitempty"`

	//  It is a reference to an object of type Tenant.
	TenantRef *string `json:"tenant_ref,omitempty"`

	// SSL Profile Type. Enum options - SSL_PROFILE_TYPE_APPLICATION, SSL_PROFILE_TYPE_SYSTEM. Field introduced in 17.2.8.
	Type *string `json:"type,omitempty"`

	// url
	// Read Only: true
	URL *string `json:"url,omitempty"`

	// Unique object identifier of the object.
	UUID *string `json:"uuid,omitempty"`
}
