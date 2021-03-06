/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.29
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Role string

// List of Role
const (
	WRITER_Role Role = "writer"
	READER_Role Role = "reader"
	ADMIN_Role Role = "admin"
	OWNER_Role Role = "owner"
)
