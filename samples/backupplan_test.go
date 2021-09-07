package sample

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"

	yscli "github.com/jibutech/backup-saas-client"
)

func TestBackupplan(t *testing.T) {
	var err error
	var ye yscli.Error

	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listBackupPlans(cli, t)

	_, _, err = cli.BackupPlanTagApi.DeleteBackupPlan(context.TODO(), tenantID, backupPlanName)
	if err != nil {
		if !errors.As(err, &ye) || ye.StatusCode() != http.StatusNotFound {
			t.Error("failed to delete backupplan", err)
		}
	}

	testBackupPlan := yscli.V1alpha1BackupPlan{
		Metadata: &yscli.V1ObjectMeta{Name: backupPlanName},
		Spec: &yscli.V1alpha1BackupPlanSpec{
			Tenant:      tenantID,
			ClusterName: clusterName,
			StorageName: storageName,
			Namespaces:  []string{backupNamespace},
			ExcludePV:   true,
		},
	}
	_, _, err = cli.BackupPlanTagApi.CreateBackupPlan(context.TODO(), tenantID, testBackupPlan)
	if err != nil {
		t.Error("failed to create backupplan", err)
	}
}

func listBackupPlans(cli *yscli.APIClient, t *testing.T) {
	var err error
	var ye yscli.Error

	bpList, _, err := cli.BackupPlanTagApi.ListBackupPlans(context.TODO(), tenantID)
	if err != nil {
		if errors.As(err, &ye) {
			if ye.StatusCode() == http.StatusNotFound {
				t.Log("not found")
			} else {
				fmt.Println(ye.Code())
				fmt.Println(ye.Message())
			}
		}
		t.Error("failed to list backupplans", err)
	}
	log.Println("list of backupplans:")
	for _, t := range bpList.Items {
		fmt.Println(t.Metadata.Name)
	}
}
