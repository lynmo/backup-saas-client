package sample

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/antihax/optional"
	yscli "github.com/jibutech/backup-saas-client"
)

func TestRestorejob(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listRestoreJobs(cli, t, restorePlanName)
	createRestoreJob(cli, t, restoreJobName)
	createRestoreJob(cli, t, restoreJobForDeletingName)
	listRestoreJobs(cli, t, restorePlanName)
	deleteRestoreJob(cli, t, restoreJobForDeletingName)
	listRestoreJobs(cli, t, restorePlanName)
	waitForRestoreJobReady(cli, t, restoreJobName)
	listRestoreJobs(cli, t, restorePlanName)
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

func listRestoreJobs(cli *yscli.APIClient, t *testing.T, restorePlanName string) {
	var err error
	var ye yscli.Error

	bpList, _, err := cli.RestoreJobTagApi.ListRestoreJobs(context.TODO(), tenantID, &yscli.RestoreJobTagApiListRestoreJobsOpts{PlanName: optional.NewString(restorePlanName)})
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
	t.Log("list of restoreJobs:")
	for _, t := range bpList.Items {
		fmt.Printf("\tname %s\n", t.Metadata.Name)
		fmt.Printf("\tbackup name: %s\n", t.Spec.BackupJobRef.BackupName)
		fmt.Printf("\tstart time: %s\n", t.Status.Report.StartTime)
		fmt.Printf("\tend time: %s\n", t.Status.Report.EndTime)
		fmt.Printf("\ttotal pvs: %d\n", t.Status.Report.TotalPVC)
	}
}

func waitForRestoreJobReady(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error

	var finished = false
	for !finished {
		t.Log("wait for restore job ready")
		time.Sleep(1 * time.Second)
		var job yscli.V1alpha1RestoreJob
		job, _, err = cli.RestoreJobTagApi.GetRestoreJob(context.TODO(), tenantID, name)
		if err != nil {
			if errors.As(err, &ye) {
				fmt.Println(ye.StatusCode())
				fmt.Println(ye.Code())
				fmt.Println(ye.Message())
			}
			t.Error("failed to list restoreJobs", err)
		} else {
			t.Log("restore job phase:", job.Status.Phase)
			if job.Status.Phase == "JobCompleted" || job.Status.Phase == "JobCanceled" || job.Status.Phase == "JobFailed" {
				finished = true
			}
		}
	}
}
