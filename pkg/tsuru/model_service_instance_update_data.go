/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceInstanceUpdateData struct {
	Description string            `json:"description,omitempty"`
	Teamowner   string            `json:"teamowner,omitempty"`
	Plan        string            `json:"plan,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Parameters  map[string]string `json:"parameters,omitempty"`
}
