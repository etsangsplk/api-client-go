/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.6
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Environment struct {

	Links *Links `json:"_links,omitempty"`

	// The unique resource id.
	Id string `json:"_id,omitempty"`

	// The key for the environment.
	Key string `json:"key,omitempty"`

	// The name of the environment.
	Name string `json:"name,omitempty"`

	// The SDK key for backend LaunchDarkly SDKs.
	ApiKey string `json:"apiKey,omitempty"`

	// The SDK key for mobile LaunchDarkly SDKs.
	MobileKey string `json:"mobileKey,omitempty"`

	// The swatch color for the environment.
	Color string `json:"color,omitempty"`

	// The default TTL.
	DefaultTtl float32 `json:"defaultTtl,omitempty"`

	// Determines if this environment is in safe mode.
	SecureMode bool `json:"secureMode,omitempty"`

	// Set to true to send detailed event information for new flags.
	DefaultTrackEvents bool `json:"defaultTrackEvents,omitempty"`

	// An array of tags for this environment.
	Tags []string `json:"tags,omitempty"`
}
