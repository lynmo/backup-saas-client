package sample

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"testing"
	"time"

	"github.com/antihax/optional"
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
	//clear multiple cluster
	t.Log("deleting multiple test cluster", clusterNameMulti1)
	cli.ClusterApi.DeleteCluster(context.TODO(), tenantID, clusterNameMulti1)
	t.Log("deleting multiple test cluster", clusterNameMulti2)
	cli.ClusterApi.DeleteCluster(context.TODO(), tenantID, clusterNameMulti2)
	t.Log("deleting multiple test cluster", clusterNameMulti3)
	cli.ClusterApi.DeleteCluster(context.TODO(), tenantID, clusterNameMulti3)

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

	createTenant(t, cli)

	for {
		testTenant, _, err := cli.TenantApi.GetTenant(context.TODO(), tenantID)
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

	createStorage(t, cli)
	listStorages(t, cli)
	createCluster(t, cli)
}

func TestCreateTenant(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	createTenant(t, cli)
}

func createTenant(t *testing.T, cli *yscli.APIClient) {
	var err error
	var ye yscli.Error

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

}

func TestCreateStorage(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	createStorage(t, cli)
}

func TestCreateCluster(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	createCluster(t, cli)
}

func TestListClusters(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listClusters(t, cli)
}

func TestListStorages(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listStorages(t, cli)
	listAllStorages(t, cli)
}

func listTenants(t *testing.T, cli *yscli.APIClient) {
	var ye yscli.Error

	opts := &yscli.TenantApiListTenantsOpts{}
	tenantList, _, err := cli.TenantApi.ListTenants(context.TODO(), opts)
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

	opts := &yscli.ClusterApiListClustersOpts{
		IncludeKubeconfig: optional.NewString("true"),
	}
	clusterList, _, err := cli.ClusterApi.ListClusters(context.TODO(), tenantID, opts)
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
		if c.Spec.Kubeconfig == "" {
			t.Error("kubeconfig should not be empty")
		}

		opts := &yscli.ClusterApiGetClusterOpts{
			IncludeKubeconfig: optional.NewString("true"),
		}
		cluster, _, err := cli.ClusterApi.GetCluster(context.TODO(), tenantID, c.Metadata.Name, opts)
		if err != nil {
			t.Log("failed to get cluster of tenant ", c.Metadata.Name, err)
			if errors.As(err, &ye) {
				t.Log(ye.Code())
				t.Log(ye.Message())
				t.Log(ye.OrigError())
			}
		}
		if cluster.Spec.Kubeconfig == "" {
			t.Error("kubeconfig should not be empty")
		}
	}
}

func listStorages(t *testing.T, cli *yscli.APIClient) {
	var ye yscli.Error

	opts := &yscli.StorageApiListStoragesOpts{
		IncludeSecrets: optional.NewString("true"),
	}
	storageList, _, err := cli.StorageApi.ListStorages(context.TODO(), tenantID, opts)
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
		t.Log(fmt.Sprintf("AccessKeyID: %s", i.Spec.S3Config.AccessKeyId))
		t.Log(fmt.Sprintf("SecretAccessKey: %s", i.Spec.S3Config.SecretAccessKey))
	}
}

func listAllStorages(t *testing.T, cli *yscli.APIClient) {
	var ye yscli.Error

	opts := &yscli.StorageApiListAllStoragesOpts{
		IncludeSecrets: optional.NewString("true"),
	}
	storageList, _, err := cli.StorageApi.ListAllStorages(context.TODO(), opts)
	if err != nil {
		t.Log("failed to list all storages", err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("list of all storages:")
	for _, i := range storageList.Items {
		t.Log(i.Metadata.Name)
		t.Log(fmt.Sprintf("AccessKeyID: %s", i.Spec.S3Config.AccessKeyId))
		t.Log(fmt.Sprintf("SecretAccessKey: %s", i.Spec.S3Config.SecretAccessKey))
	}
}

func createCluster(t *testing.T, cli *yscli.APIClient) {
	var err error
	var ye yscli.Error
	t.Log("creating cluster", clusterName)
	var kubeconfig []byte
	var envKubeconfig string
	kubeconfig, err = ioutil.ReadFile("kubeconfig.yaml")
	if err != nil {
		//t.Log("no kubeconfig.yaml file in current dir, will try $HOME/.kube/config")
		t.Log("no kubeconfig.yaml file in current dir, will try to use KUBECONFIG environment")
		envKubeconfig = os.Getenv("KUBECONFIG")
		usr, err := user.Current()
		if err != nil {
			t.Error("failed to get current user")
			if errors.As(err, &ye) {
				t.Log(ye.Code())
				t.Log(ye.Message())
				t.Log(ye.OrigError())
			}
		}
		// If an env variable is specified with the config locaiton, use that
		if len(envKubeconfig) > 0 {
			t.Log("Get the KUBECONFIG env variable ", envKubeconfig)
		}
		if envKubeconfig == "" {
			t.Error("Failed to load KUBECONFIG env variable")
			// If no KUBECONFIG env variable, try the default location in the user's home directory
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
	}
	//create multiple cluster
	t.Log("creating multiple test cluster", clusterNameMulti1)
	testCluster1 := yscli.V1alpha1Cluster{
		Metadata: &yscli.V1ObjectMeta{Name: clusterNameMulti1},
		Spec: &yscli.V1alpha1ClusterSpec{
			Tenant:     tenantID,
			Kubeconfig: string(kubeconfig),
		},
	}
	testCluster1, _, err = cli.ClusterApi.CreateCluster(context.TODO(), tenantID, testCluster1)
	if err != nil {
		t.Error("failed to create cluster", err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("creating multiple test cluster", clusterNameMulti2)
	testCluster2 := yscli.V1alpha1Cluster{
		Metadata: &yscli.V1ObjectMeta{Name: clusterNameMulti2},
		Spec: &yscli.V1alpha1ClusterSpec{
			Tenant:     tenantID,
			Kubeconfig: string(kubeconfig),
		},
	}
	testCluster2, _, err = cli.ClusterApi.CreateCluster(context.TODO(), tenantID, testCluster2)
	if err != nil {
		t.Error("failed to create cluster", err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
	t.Log("creating multiple test cluster", clusterNameMulti3)
	testCluster3 := yscli.V1alpha1Cluster{
		Metadata: &yscli.V1ObjectMeta{Name: clusterNameMulti3},
		Spec: &yscli.V1alpha1ClusterSpec{
			Tenant:     tenantID,
			Kubeconfig: string(kubeconfig),
		},
	}
	testCluster3, _, err = cli.ClusterApi.CreateCluster(context.TODO(), tenantID, testCluster3)
	if err != nil {
		t.Error("failed to create cluster", err)
		if errors.As(err, &ye) {
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
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

func createStorage(t *testing.T, cli *yscli.APIClient) {
	var err error
	var ye yscli.Error

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
}
