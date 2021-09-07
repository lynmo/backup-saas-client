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
		}
	}
	t.Log("deleting storage", storageName)
	_, _, err = cli.StorageApi.DeleteStorage(context.TODO(), tenantID, storageName)
	if err != nil {
		if errors.As(err, &ye) && ye.StatusCode() != http.StatusNotFound {
			t.Error("failed to clear storage", err)
		}
	}
	t.Log("deleting tenant", tenantID)
	_, _, err = cli.TenantApi.DeleteTenant(context.TODO(), tenantID)
	if err != nil {
		if errors.As(err, &ye) && ye.StatusCode() != http.StatusNotFound {
			t.Error("failed to clear tenant", err)
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
	}
	listStorages(t, cli)
	// testStorage.Spec.S3Config.AccessKeyId = accessKey
	// testStorage.Spec.S3Config.Bucket = bucketName
	// testStorage.Spec.S3Config.S3url = s3URL
	// testStorage.Spec.S3Config.SecretAccessKey = secretKey
	// testStorage.Spec.S3Config.Region = region
	// for {
	// 	t.Log("updating storage", storageName)
	// 	_, _, err = cli.StorageApi.UpdateStorage(context.TODO(), tenantID, storageName, testStorage)
	// 	if err != nil {
	// 		t.Error("failed to update storage", err)
	// 		var ye yscli.Error
	// 		if errors.As(err, &ye) {
	// 			t.Log(ye.Code())
	// 			t.Log(ye.Message())
	// 			t.Log(ye.OrigError())
	// 		}
	// 		time.Sleep(1 * time.Second)
	// 	} else {
	// 		break
	// 	}
	// }

	for {
		testTenant, _, err = cli.TenantApi.GetTenant(context.TODO(), tenantID)
		if err != nil {
			t.Error("failed to query tenant", tenantID)
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
func listStorages(t *testing.T, cli *yscli.APIClient) {
	storageList, _, err := cli.StorageApi.ListStorages(context.TODO(), tenantID)
	if err != nil {
		t.Log("failed to list storages of tenant ", tenantID, err)
	}
	t.Log("list of storages:")
	for _, i := range storageList.Items {
		t.Log(i.Metadata.Name)
	}
}

func createCluster(t *testing.T, cli *yscli.APIClient) {
	var err error
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
		var ye yscli.Error
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
}
