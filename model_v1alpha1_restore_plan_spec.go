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

type V1alpha1RestorePlanSpec struct {
	BackupName string `json:"backupName"`
	BackupPlanRef *V1alpha1BackupPlanSpec `json:"backupPlanRef,omitempty"`
	Desc string `json:"desc,omitempty"`
	DestClusterName string `json:"destClusterName"`
	DisplayName string `json:"displayName,omitempty"`
	ExcludedResources []string `json:"excludedResources,omitempty"`
	NamespaceMappings []string `json:"namespaceMappings,omitempty"`
	Tenant string `json:"tenant"`
}
