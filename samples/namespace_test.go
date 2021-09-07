package sample

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"

	yscli "github.com/jibutech/backup-saas-client"
)

func TestListNamespaces(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath(apiEndpoint)
	nsList, _, err := cli.ClusterApi.GetNamespaces(context.TODO(), tenantID, clusterName)
	if err != nil {
		var ye yscli.Error
		if errors.As(err, &ye) {
			if ye.StatusCode() == http.StatusNotFound {
				t.Log("not found")
			} else {
				t.Log(ye.Code())
				t.Log(ye.Message())
				t.Log(ye.OrigError())
			}
		}
		t.Error("failed to list namespaces", err)
		return
	}
	log.Println("list of namespaces:")
	for _, n := range nsList.Items {
		t.Log(n.Metadata.Name)
	}
}
