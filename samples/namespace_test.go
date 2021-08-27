package sample

import (
	"context"
	"fmt"
	"log"
	"testing"

	yscli "github.com/jibutech/backup-saas-client"
)

func TestListNamespaces(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath("http://127.0.0.1:31800")
	nsList, _, err := cli.ClusterApi.GetNamespaces(context.TODO(), tenantID, clusterName)
	if err != nil {
		t.Error("failed to list namespaces", err)
	}
	log.Println("list of namespaces:")
	for _, t := range nsList.Items {
		fmt.Println(t.Metadata.Name)
	}
}
