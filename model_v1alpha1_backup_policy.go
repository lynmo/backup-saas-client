/*
 * Backup Service
 *
 * Backup Service OpenAPI
 *
 * API version: v0.0.0-master+$Format:%h$
 * Contact: zoubo@jibudata.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1BackupPolicy struct {
	Frequency int32 `json:"frequency,omitempty"`
	Name string `json:"name,omitempty"`
	Repeat bool `json:"repeat,omitempty"`
	Retention int32 `json:"retention,omitempty"`
}