package sample

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	yscli "github.com/jibutech/backup-saas-client"
)

func TestBackupplan(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath("http://127.0.0.1:31800")

	listBackupPlans(cli, t)

	_, resp, err := cli.BackupPlanTagApi.DeleteBackupPlan(context.TODO(), tenantID, backupPlanName)
	if err != nil && (resp == nil || resp.StatusCode != 404) {
		t.Error("failed to delete backupplan", err)
	}

	testBackupPlan := yscli.V1alpha1BackupPlan{
		Metadata: &yscli.V1ObjectMeta{Name: backupPlanName},
		Spec: &yscli.V1alpha1BackupPlanSpec{
			Tenant:      tenantID,
			ClusterName: clusterName,
			StorageName: storageName,
			Namespaces:  []string{backupNamespace},
		},
	}
	_, _, err = cli.BackupPlanTagApi.CreateBackupPlan(context.TODO(), tenantID, testBackupPlan)
	if err != nil {
		t.Error("failed to create backupplan", err)
	}
}

func listBackupPlans(cli *yscli.APIClient, t *testing.T) {
	bpList, resp, err := cli.BackupPlanTagApi.ListBackupPlans(context.TODO(), tenantID)
	if err != nil {
		if se, ok := err.(yscli.GenericSwaggerError); ok && se.Model() != nil {
			if ye, ok := se.Model().(yscli.YsapiError); ok {
				fmt.Println(ye.Code)
				fmt.Println(ye.Message)
			}
		} else if resp != nil && resp.StatusCode == http.StatusNotFound {
			t.Log("not found")
		}
		t.Error("failed to list backupplans", err)
	}
	log.Println("list of backupplans:")
	for _, t := range bpList.Items {
		fmt.Println(t.Metadata.Name)
	}
}
