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

func TestRestorejob(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listRestoreJobs(cli, t)
	createRestoreJob(cli, t, restoreJobName)
	createRestoreJob(cli, t, restoreJobForDeletingName)
	listRestoreJobs(cli, t)
	deleteRestoreJob(cli, t, restoreJobForDeletingName)
	listRestoreJobs(cli, t)
}

func createRestoreJob(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error

	restoreJob := yscli.V1alpha1RestoreJob{
		Metadata: &yscli.V1ObjectMeta{Name: name},
		Spec: &yscli.V1alpha1RestoreJobSpec{
			Action:        "start",
			BackupJobName: backupJobName,
			RestoreName:   restorePlanName,
			Tenant:        tenantID,
		},
	}
	_, _, err = cli.RestoreJobTagApi.CreateRestoreJob(context.TODO(), tenantID, restoreJob)
	if err != nil {
		t.Error("failed to create restore job")
		if errors.As(err, &ye) {
			t.Log(ye.StatusCode())
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
}

func deleteRestoreJob(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error

	_, _, err = cli.RestoreJobTagApi.DeleteRestoreJob(context.TODO(), tenantID, name)
	if err != nil {
		t.Error("failed to delete restore job")
		if errors.As(err, &ye) {
			t.Log(ye.StatusCode())
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
}

func listRestoreJobs(cli *yscli.APIClient, t *testing.T) {
	var err error
	var ye yscli.Error

	bpList, _, err := cli.RestoreJobTagApi.ListRestoreJobs(context.TODO(), tenantID)
	if err != nil {
		if errors.As(err, &ye) {
			if ye.StatusCode() == http.StatusNotFound {
				t.Log("not found")
			} else {
				fmt.Println(ye.Code())
				fmt.Println(ye.Message())
			}
		}
		t.Error("failed to list restoreJobs", err)
	}
	log.Println("list of restoreJobs:")
	for _, t := range bpList.Items {
		fmt.Println(t.Metadata.Name)
	}
}
