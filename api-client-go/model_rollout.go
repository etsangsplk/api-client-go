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

type Rollout struct {
	BucketBy string `json:"bucketBy,omitempty"`
	Variations []WeightedVariation `json:"variations,omitempty"`
}