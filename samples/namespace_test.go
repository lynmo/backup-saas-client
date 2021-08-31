package sample

import (
	"context"
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
	nsList, resp, err := cli.ClusterApi.GetNamespaces(context.TODO(), tenantID, clusterName)
	if err != nil {
		if se, ok := err.(yscli.GenericSwaggerError); ok && se.Model() != nil {
			if ye, ok := se.Model().(yscli.YsapiError); ok {
				fmt.Println(ye.Code)
				fmt.Println(ye.Message)
			}
		} else if resp != nil && resp.StatusCode == http.StatusNotFound {
			t.Log("not found")
		}
		t.Error("failed to list namespaces", err)
	}
	log.Println("list of namespaces:")
	for _, t := range nsList.Items {
		fmt.Println(t.Metadata.Name)
	}
}
