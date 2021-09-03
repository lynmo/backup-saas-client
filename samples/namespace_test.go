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

func TestListNamespaces(t *testing.T) {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath("http://127.0.0.1:31800")
	nsList, _, err := cli.ClusterApi.GetNamespaces(context.TODO(), tenantID, clusterName)
	if err != nil {
		var ye yscli.Error
		if errors.As(err, &ye) {
			if ye.StatusCode() == http.StatusNotFound {
				fmt.Println("not found")
			} else {
				fmt.Println(ye.Code())
				fmt.Println(ye.Message())
				fmt.Println(ye.Errs())
			}
		}
		t.Error("failed to list namespaces", err)
		return
	}
	log.Println("list of namespaces:")
	for _, t := range nsList.Items {
		fmt.Println(t.Metadata.Name)
	}
}
