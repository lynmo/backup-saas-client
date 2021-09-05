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

type V1alpha1RestoreJobSpec struct {
	Action string `json:"action,omitempty"`
	BackupJobName string `json:"backupJobName"`
	RestoreName string `json:"restoreName"`
	Tenant string `json:"tenant"`
}
