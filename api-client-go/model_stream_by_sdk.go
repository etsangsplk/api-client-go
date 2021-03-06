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

type StreamBySdk struct {
	Links *StreamBySdkLinks `json:"_links,omitempty"`
	Metadata []StreamBySdkLinksMetadata `json:"metadata,omitempty"`
	Series []StreamUsageSeries `json:"series,omitempty"`
}
