package sample

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	yscli "github.com/jibutech/backup-saas-client"
)

func TestBackupjob(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath("http://127.0.0.1:31800")

	listBackupJobs(cli, t)

	_, resp, err := cli.BackupJobTagApi.DeleteBackupJob(context.TODO(), tenantID, backupJobName)
	if err != nil && (resp == nil || resp.StatusCode != 404) {
		t.Error("failed to delete backupjob", err)
	}

	testBackupJob := yscli.V1alpha1BackupJob{
		Metadata: &yscli.V1ObjectMeta{Name: backupJobName},
		Spec: &yscli.V1alpha1BackupJobSpec{
			Tenant:     tenantID,
			BackupName: backupPlanName,
			Action:     "start",
		},
	}
	_, _, err = cli.BackupJobTagApi.CreateBackupJob(context.TODO(), tenantID, testBackupJob)
	if err != nil {
		t.Error("failed to create backupplan", err)
	}
}

func listBackupJobs(cli *yscli.APIClient, t *testing.T) {
	bpList, resp, err := cli.BackupJobTagApi.ListBackupJobs(context.TODO(), tenantID)
	if err != nil {
		if se, ok := err.(yscli.GenericSwaggerError); ok && se.Model() != nil {
			if ye, ok := se.Model().(yscli.YsapiError); ok {
				fmt.Println(ye.Code)
				fmt.Println(ye.Message)
			}
		} else if resp != nil && resp.StatusCode == http.StatusNotFound {
			t.Log("not found")
		}
		t.Error("failed to list backupjobs", err)
	}
	log.Println("list of backupjobs:")
	for _, t := range bpList.Items {
		fmt.Println(t.Metadata.Name)
	}
}
