package sample

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os/user"
	"path/filepath"
	"testing"
	"time"

	yscli "github.com/jibutech/backup-saas-client"
)

func TestRestAPIs(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	var err error
	var ye yscli.Error

	listTenants(t, cli)
	listClusters(t, cli)
	listStorages(t, cli)

	//clear
	t.Log("deleting cluster", clusterName)
	_, _, err = cli.ClusterApi.DeleteCluster(context.TODO(), tenantID, clusterName)
	if err != nil {
		if errors.As(err, &ye) && ye.StatusCode() != http.StatusNotFound {
			t.Error("failed to clear cluster", err)
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("deleting storage", storageName)
	_, _, err = cli.StorageApi.DeleteStorage(context.TODO(), tenantID, storageName)
	if err != nil {
		if errors.As(err, &ye) && ye.StatusCode() != http.StatusNotFound {
			t.Error("failed to clear storage", err)
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("deleting tenant", tenantID)
	_, _, err = cli.TenantApi.DeleteTenant(context.TODO(), tenantID)
	if err != nil {
		if errors.As(err, &ye) && ye.StatusCode() != http.StatusNotFound {
			t.Error("failed to clear tenant", err)
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}

	listTenants(t, cli)
	listClusters(t, cli)
	listStorages(t, cli)

	for {
		_, _, err = cli.TenantApi.GetTenant(context.TODO(), tenantID)
		if err != nil {
			if errors.As(err, &ye) && ye.StatusCode() == http.StatusNotFound {
				break
			}
			t.Log("failed to query tenant", tenantID)
		}
		t.Log("wait for tenant being finalized")
		time.Sleep(1 * time.Second)
	}

	t.Log("creating tenant", tenantID)
	testTenant := yscli.V1alpha1Tenant{Metadata: &yscli.V1ObjectMeta{Name: tenantID}}
	testTenant, _, err = cli.TenantApi.CreateTenant(context.TODO(), testTenant)
	if err != nil {
		t.Error("failed to create tenant", err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
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
				Region:          region,
			},
			Tenant:     tenantID,
			S3Provider: "aws",
		},
	}
	testStorage, _, err = cli.StorageApi.CreateStorage(context.TODO(), tenantID, testStorage)
	if err != nil {
		t.Error("failed to create storage", err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}

	listStorages(t, cli)

	for {
		testTenant, _, err = cli.TenantApi.GetTenant(context.TODO(), tenantID)
		if err != nil {
			t.Error("failed to query tenant", tenantID)
			if errors.As(err, &ye) {
				t.Log(ye.Code())
				t.Log(ye.Message())
				t.Log(ye.OrigError())
			}
		}
		t.Log("tenant phase:", testTenant.Status.Phase)
		if testTenant.Status.Phase == "Ready" {
			break
		}
		t.Log("wait for tenant ready")
		time.Sleep(1 * time.Second)
	}

	createCluster(t, cli)
}

func listTenants(t *testing.T, cli *yscli.APIClient) {
	var ye yscli.Error
	tenantList, _, err := cli.TenantApi.ListTenants(context.TODO())
	if err != nil {
		t.Error("failed to list tenants", err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("list of tenants:")
	for _, tenant := range tenantList.Items {
		t.Log(tenant.Metadata.Name)
	}
}
func listClusters(t *testing.T, cli *yscli.APIClient) {
	var ye yscli.Error
	clusterList, _, err := cli.ClusterApi.ListClusters(context.TODO(), tenantID)
	if err != nil {
		t.Log("failed to list clusters of tenant ", tenantID, err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("list of clusters:")
	for _, c := range clusterList.Items {
		t.Log(c.Metadata.Name)
	}
}
func listStorages(t *testing.T, cli *yscli.APIClient) {
	var ye yscli.Error
	storageList, _, err := cli.StorageApi.ListStorages(context.TODO(), tenantID)
	if err != nil {
		t.Log("failed to list storages of tenant ", tenantID, err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("list of storages:")
	for _, i := range storageList.Items {
		t.Log(i.Metadata.Name)
	}
}

func createCluster(t *testing.T, cli *yscli.APIClient) {
	var err error
	var ye yscli.Error
	t.Log("creating cluster", clusterName)
	var kubeconfig []byte
	kubeconfig, err = ioutil.ReadFile("kubeconfig.yaml")
	if err != nil {
		t.Log("no kubeconfig.yaml file in current dir, will try $HOME/.kube/config")
		usr, err := user.Current()
		if err != nil {
			t.Error("failed to get current user")
			if errors.As(err, &ye) {
				t.Log(ye.Code())
				t.Log(ye.Message())
				t.Log(ye.OrigError())
			}
		}
		kubeconfig, err = ioutil.ReadFile(filepath.Join(usr.HomeDir, ".kube/config"))
		if err != nil {
			t.Error("failed to load kubeconfig")
			if errors.As(err, &ye) {
				t.Log(ye.Code())
				t.Log(ye.Message())
				t.Log(ye.OrigError())
			}
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
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
}
