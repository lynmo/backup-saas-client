package sample

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/antihax/optional"
	yscli "github.com/jibutech/backup-saas-client"
)

func TestBackupjob(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listBackupJobs(cli, t)

	createBackupJob(cli, t, backupJobName)
	createBackupJob(cli, t, backupJobForDeletingName)

	listBackupJobs(cli, t)

	deleteBackupJob(cli, t, backupJobForDeletingName)
	listBackupJobs(cli, t)

	waitForBackupJobReady(cli, t, backupJobName)
	listBackupJobs(cli, t)
}

func createBackupJob(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error

	testBackupJob := yscli.V1alpha1BackupJob{
		Metadata: &yscli.V1ObjectMeta{Name: name},
		Spec: &yscli.V1alpha1BackupJobSpec{
			Tenant:      tenantID,
			BackupName:  backupPlanName,
			DisplayName: backupPlanName,
			Action:      "start",
		},
	}
	_, _, err = cli.BackupJobTagApi.CreateBackupJob(context.TODO(), tenantID, testBackupJob)
	if err != nil {
		if errors.As(err, &ye) {
			fmt.Println(ye.Code())
			fmt.Println(ye.Message())
			fmt.Println(ye.OrigError())
		}
		t.Error("failed to create backupjob", err)
		return
	}
}

func deleteBackupJob(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error
	_, _, err = cli.BackupJobTagApi.DeleteBackupJob(context.TODO(), tenantID, name)
	if err != nil {
		if !errors.As(err, &ye) || ye.StatusCode() != http.StatusNotFound {
			t.Error("failed to delete backupjob", err)
		}
	}
}

func listBackupJobs(cli *yscli.APIClient, t *testing.T) {
	var err error
	var ye yscli.Error
	bpList, _, err := cli.BackupJobTagApi.ListBackupJobs(context.TODO(), tenantID, &yscli.BackupJobTagApiListBackupJobsOpts{PlanName: optional.NewString(backupPlanName)})
	if err != nil {
		if errors.As(err, &ye) {
			if ye.StatusCode() == http.StatusNotFound {
				fmt.Println("not found")
			} else {
				fmt.Println(ye.Code())
				fmt.Println(ye.Message())
				fmt.Println(ye.OrigError())
			}
		}
		t.Error("failed to list backupjobs", err)
		return
	}
	log.Println("list of backupjobs:")
	for _, t := range bpList.Items {
		fmt.Printf("\tDisplayName: %s\n", t.Spec.DisplayName)
		fmt.Printf("\tstart time: %s\n", t.Status.Report.StartTime)
		fmt.Printf("\tend time: %s\n", t.Status.Report.EndTime)
		fmt.Printf("\texpired time: %s\n", t.Status.Report.ExpiredTime)
		fmt.Printf("\ttotal pvs: %d\n", t.Status.Report.TotalPVC)
	}
}
func waitForBackupJobReady(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error

	var finished = false
	for !finished {
		t.Log("wait for backup job ready")
		time.Sleep(1 * time.Second)
		var job yscli.V1alpha1BackupJob
		job, _, err = cli.BackupJobTagApi.GetBackupJob(context.TODO(), tenantID, name)
		if err != nil {
			if errors.As(err, &ye) {
				fmt.Println(ye.StatusCode())
				fmt.Println(ye.Code())
				fmt.Println(ye.Message())
			}
			t.Error("failed to list backupJobs", err)
		} else {
			t.Log("backup job phase:", job.Status.Phase)
			if job.Status.Phase == "JobCompleted" || job.Status.Phase == "JobCanceled" || job.Status.Phase == "JobFailed" {
				finished = true
			}
		}
	}
}
