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

type V1alpha1ClusterSpec struct {
	ApiEndpoint string `json:"apiEndpoint,omitempty"`
	Desc string `json:"desc,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`
	ExtraParams map[string]string `json:"extraParams,omitempty"`
	Kubeconfig string `json:"kubeconfig"`
	Provider string `json:"provider,omitempty"`
	Region string `json:"region,omitempty"`
	Role string `json:"role,omitempty"`
	SecretName string `json:"secretName,omitempty"`
	ServiceAccount string `json:"serviceAccount,omitempty"`
	SnapshotClasses []V1alpha1VolumeSnapshotClass `json:"snapshotClasses,omitempty"`
	StorageClasses []V1alpha1StorageClass `json:"storageClasses,omitempty"`
	Tenant string `json:"tenant"`
	Zone string `json:"zone,omitempty"`
}
