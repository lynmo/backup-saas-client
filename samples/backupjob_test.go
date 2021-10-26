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

	createBackupJob(cli, t, backupJobName, true)
	createBackupJob(cli, t, backupJobForDeletingName, true)

	listBackupJobs(cli, t)

	deleteBackupJob(cli, t, backupJobForDeletingName)
	listBackupJobs(cli, t)
	listBackupPlans(cli, t)

	waitForBackupJobReady(cli, t, backupJobName)
	listBackupJobs(cli, t)
	listBackupPlans(cli, t)
}

func TestCreateBackupJob(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	createBackupJob(cli, t, backupJobNameNotStart, false)
}

func TestWatchBackupjob(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	watchDone := make(chan string)
	go watchBackupjob(cli, t, backupJobNameWatch, watchDone)

	createBackupJob(cli, t, backupJobNameWatch, false)
	// wait for a while for a modified event
	time.Sleep(10 * time.Second)
	deleteBackupJob(cli, t, backupJobNameWatch)

	<-watchDone
}

func createBackupJob(cli *yscli.APIClient, t *testing.T, name string, start bool) {
	var err error
	var ye yscli.Error
	var action string
	if start {
		action = "StartJob"
	}

	testBackupJob := yscli.V1alpha1BackupJob{
		Metadata: &yscli.V1ObjectMeta{Name: name},
		Spec: &yscli.V1alpha1BackupJobSpec{
			Tenant:      tenantID,
			BackupName:  backupPlanName,
			DisplayName: backupPlanName,
			Action:      action,
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
			if job.Status.Phase == "JobCompleted" ||
				job.Status.Phase == "JobCanceled" ||
				job.Status.Phase == "JobFailed" ||
				job.Status.Phase == "Error" {
				finished = true
			}
		}
	}
}

func watchBackupjob(cli *yscli.APIClient, t *testing.T, name string, watchDone chan string) {
	var err error
	var ye yscli.Error

	addEventGot, modifyEventGot, deleteEventGot := false, false, false

	defer func() {
		if !addEventGot {
			t.Error("did not get add event")
		}
		if !modifyEventGot {
			t.Error("did not get modify event")
		}
		if !deleteEventGot {
			t.Error("did not get delete event")
		}
	}()
	defer close(watchDone)

	watcher, err := cli.BackupJobTagApi.WatchBackupJobs(context.TODO(), tenantID)
	if err != nil {
		if errors.As(err, &ye) {
			t.Log(ye.StatusCode())
			t.Log(ye.Code())
			t.Log(ye.Message())
		}
		t.Error("failed to start watching backupJobs", err)
	} else {
		defer watcher.Stop()
		var e yscli.WatchEvent
		for {
			select {
			case <-time.After(60 * time.Second):
				t.Log("Timeout watching backup jobs")
				return
			case e = <-watcher.ResultChan():
				{
					t.Log("Event type: ", e.Type)
					if job, ok := e.Object.(yscli.V1alpha1BackupJob); ok {
						t.Log("Tenant name: ", job.Spec.Tenant)
						t.Log("Job name: ", job.Metadata.Name)
						if job.Metadata.Name == name {
							switch e.Type {
							case yscli.WatchEventTypes.Added:
								addEventGot = true
							case yscli.WatchEventTypes.Modified:
								modifyEventGot = true
							case yscli.WatchEventTypes.Deleted:
								deleteEventGot = true
								return
							}
						}
					} else {
						t.Error("Invalid event object received")
					}
				}
			}
		}
	}
}
