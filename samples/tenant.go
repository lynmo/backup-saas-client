package main

import (
	"context"
	"fmt"
	"log"

	yscli "github.com/jibutech/backup-saas-client"
)

func main() {
	cfg := yscli.NewConfiguration()
	cli := yscli.NewAPIClient(cfg)
	cli.ChangeBasePath("http://127.0.0.1:31800")
	tenantList, _, err := cli.TenantApi.ListTenants(context.TODO())
	if err != nil {
		log.Println("failed to list tenants")
		log.Fatal(err)
		return
	}
	for _, t := range tenantList.Items {
		log.Println(t)
	}
	fmt.Println(cfg)
}
