/*
 * Backup Service
 *
 * Backup Service OpenAPI
 *
 * API version: v0.0.0-master+$Format:%h$
 * Contact: shaofeng@jibudata.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1BackupPlanStatus struct {
	Conditions []V1alpha1Condition `json:"conditions,omitempty"`
	Phase string `json:"phase,omitempty"`
}
