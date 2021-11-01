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

func TestRestoreplan(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listRestorePlans(cli, t)
	createRestorePlan(cli, t, restorePlanName)
	createRestorePlan(cli, t, restorePlanForDeletingName)
	listRestorePlans(cli, t)
	deleteRestorePlan(cli, t, restorePlanForDeletingName)
	listRestorePlans(cli, t)
}

func createRestorePlan(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error

	restorePlan := yscli.V1alpha1RestorePlan{
		Metadata: &yscli.V1ObjectMeta{Name: name},
		Spec: &yscli.V1alpha1RestorePlanSpec{
			BackupName:        backupPlanName,
			DestClusterName:   clusterName,
			Tenant:            tenantID,
			NamespaceMappings: []string{fmt.Sprintf("%s:%s", backupNamespace, restoreDestNamespace)},
		},
	}
	_, _, err = cli.RestorePlanTagApi.CreateRestorePlan(context.TODO(), tenantID, restorePlan)
	if err != nil {
		t.Error("failed to create restore plan")
		if errors.As(err, &ye) {
			t.Log(ye.StatusCode())
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
}

func deleteRestorePlan(cli *yscli.APIClient, t *testing.T, name string) {
	var err error
	var ye yscli.Error

	_, _, err = cli.RestorePlanTagApi.DeleteRestorePlan(context.TODO(), tenantID, name)
	if err != nil {
		t.Error("failed to delete restore plan")
		if errors.As(err, &ye) {
			t.Log(ye.StatusCode())
			t.Log(ye.Code())
			t.Log(ye.Message())
			t.Log(ye.OrigError())
		}
	}
}

func TestListRestoreplan(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)

	listRestorePlans(cli, t)
}

func listRestorePlans(cli *yscli.APIClient, t *testing.T) {
	var err error
	var ye yscli.Error

	opts := &yscli.RestorePlanTagApiListRestorePlansOpts{}
	bpList, _, err := cli.RestorePlanTagApi.ListRestorePlans(context.TODO(), tenantID, opts)
	if err != nil {
		if errors.As(err, &ye) {
			if ye.StatusCode() == http.StatusNotFound {
				t.Log("not found")
			} else {
				fmt.Println(ye.Code())
				fmt.Println(ye.Message())
			}
		}
		t.Error("failed to list restorePlans", err)
	}
	log.Println("list of restorePlans:")
	for _, t := range bpList.Items {
		fmt.Println(t.Metadata.Name)
		fmt.Println(t.Status.CurrentJobName)
		fmt.Println(t.Status.CurrentJobPhase)
		fmt.Printf("%v\n", t.Status.JobDetail)
	}
}
