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

type V1alpha1BackupPlanSpec struct {
	ClusterName string `json:"clusterName"`
	CopyMethod string `json:"copyMethod,omitempty"`
	Desc string `json:"desc,omitempty"`
	ExcludePV bool `json:"excludePV,omitempty"`
	ExcludedResources []string `json:"excludedResources,omitempty"`
	IncludedResources []string `json:"includedResources,omitempty"`
	Namespaces []string `json:"namespaces,omitempty"`
	Policy *V1alpha1BackupPolicy `json:"policy,omitempty"`
	StorageName string `json:"storageName"`
	StorageRef *V1alpha1StorageSpec `json:"storageRef,omitempty"`
	Tenant string `json:"tenant"`
}
