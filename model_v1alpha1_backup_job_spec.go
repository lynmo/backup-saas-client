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

type V1alpha1BackupJobSpec struct {
	Action string `json:"action,omitempty"`
	BackupName string `json:"backupName"`
	Tenant string `json:"tenant"`
}