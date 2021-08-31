package sample

import (
	"context"
	"io/ioutil"
	"net/http"
	"os/user"
	"path/filepath"
	"testing"

	yscli "github.com/jibutech/backup-saas-client"
)

func TestRestAPIs(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath("http://127.0.0.1:31800")

	var err error
	var resp *http.Response

	listTenants(t, cli)
	listClusters(t, cli)

	//clear
	t.Log("deleting cluster", clusterName)
	_, resp, err = cli.ClusterApi.DeleteCluster(context.TODO(), tenantID, clusterName)
	if err != nil && (resp == nil || resp.StatusCode != 404) {
		t.Error("failed to clear cluster", err, resp)
	}
	t.Log("deleting storage", storageName)
	_, resp, err = cli.StorageApi.DeleteStorage(context.TODO(), tenantID, storageName)
	if err != nil && (resp == nil || resp.StatusCode != 404) {
		t.Error("failed to clear storage", err, resp)
	}
	t.Log("deleting tenant", tenantID)
	_, resp, err = cli.TenantApi.DeleteTenant(context.TODO(), tenantID)
	if err != nil && (resp == nil || resp.StatusCode != 404) {
		t.Error("failed to clear tenant", err, resp)
	}

	listTenants(t, cli)
	listClusters(t, cli)

	t.Log("creating tenant", tenantID)
	testTenant := yscli.V1alpha1Tenant{Metadata: &yscli.V1ObjectMeta{Name: tenantID}}
	testTenant, resp, err = cli.TenantApi.CreateTenant(context.TODO(), testTenant)
	if err != nil {
		t.Error("failed to create tenant", err, resp)
	}

	t.Log("creating cluster", clusterName)
	var kubeconfig []byte
	kubeconfig, err = ioutil.ReadFile("kubeconfig")
	if err != nil {
		t.Log("no kubeconfig file in current dir, will try $HOME/.kube/config")
		usr, err := user.Current()
		if err != nil {
			t.Error("failed to get current user")
		}
		kubeconfig, err = ioutil.ReadFile(filepath.Join(usr.HomeDir, ".kube/config"))
		if err != nil {
			t.Error("failed to load kubeconfig")
		}
	}
	testCluster := yscli.V1alpha1Cluster{
		Metadata: &yscli.V1ObjectMeta{Name: clusterName},
		Spec: &yscli.V1alpha1ClusterSpec{
			Tenant:     tenantID,
			Kubeconfig: string(kubeconfig),
		},
	}
	testCluster, _, err = cli.ClusterApi.CreateCluster(context.TODO(), tenantID, testCluster)
	if err != nil {
		t.Error("failed to create cluster", err)
	}

	t.Log("creating storage", storageName)
	testStorage := yscli.V1alpha1Storage{
		Metadata: &yscli.V1ObjectMeta{Name: storageName},
		Spec: &yscli.V1alpha1StorageSpec{
			S3Config: &yscli.V1alpha1BackupS3Config{
				AccessKeyId:     accessKey,
				Bucket:          bucketName,
				S3url:           s3URL,
				SecretAccessKey: secretKey,
			},
			Tenant:     tenantID,
			S3Provider: "aws",
		},
	}
	testStorage, _, err = cli.StorageApi.CreateStorage(context.TODO(), tenantID, testStorage)
	if err != nil {
		t.Error("failed to create storage", err)
	}
}

func listTenants(t *testing.T, cli *yscli.APIClient) {
	tenantList, _, err := cli.TenantApi.ListTenants(context.TODO())
	if err != nil {
		t.Error("failed to list tenants", err)
	}
	t.Log("list of tenants:")
	for _, tenant := range tenantList.Items {
		t.Log(tenant.Metadata.Name)
	}
}

func listClusters(t *testing.T, cli *yscli.APIClient) {
	clusterList, _, err := cli.ClusterApi.ListClusters(context.TODO(), tenantID)
	if err != nil {
		t.Log("failed to list clusters of tenant ", tenantID, err)
	}
	t.Log("list of clusters:")
	for _, c := range clusterList.Items {
		t.Log(c.Metadata.Name)
	}
}
