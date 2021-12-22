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

type V1alpha1TenantSpec struct {
	Contact string `json:"contact,omitempty"`
	DefaultBackupStorage string `json:"defaultBackupStorage,omitempty"`
	Desc string `json:"desc,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	IsExpired bool `json:"isExpired,omitempty"`
	OrgName string `json:"orgName,omitempty"`
}
