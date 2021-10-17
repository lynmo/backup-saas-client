package sample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	tenantID    = "go-client-test"
	clusterName = "go-client-test-k8s"
	clusterNameMulti1 = "go-client-test-k8s-1"
    clusterNameMulti2 = "go-client-test-k8s-2"
    clusterNameMulti3 = "go-client-test-k8s-3"

	storageName = "go-client-test-s3"

	backupPlanName           = "go-client-test-backupplan"
	backupNamespace          = "wordpress"
	backupJobName            = "go-client-test-backupjob"
	backupJobNameNotStart    = "go-client-test-backupjob-not-start"
	backupJobForDeletingName = "go-client-test-backupjob-deleting"

	restorePlanName            = "go-client-test-restoreplan"
	restoreDestNamespace       = "wordpress-restored"
	restorePlanForDeletingName = "go-client-test-restoreplan-deleting"

	restoreJobName            = "go-client-test-restorejob"
	restoreJobNameNotStart    = "go-client-test-restorejob-not-start"
	restoreJobForDeletingName = "go-client-test-restorejob-deleting"
)

var apiEndpoint string
var accessKey string
var bucketName string
var s3URL string
var secretKey string
var region string

type config struct {
	ApiEndpoint string `json:"apiEndpoint"`
	AccessKey   string `json:"accessKey"`
	BucketName  string `json:"bucketName"`
	S3URL       string `json:"s3URL"`
	SecretKey   string `json:"secretKey"`
	Region      string `json:"region"`
}

func init() {
	var err error
	var data []byte
	data, err = ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("error loading config file: %s\n", err)
		return
	}
	var config = config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("error loading config: %s\n", err)
		return
	}
	apiEndpoint = config.ApiEndpoint
	accessKey = config.AccessKey
	bucketName = config.BucketName
	s3URL = config.S3URL
	secretKey = config.SecretKey
	region = config.Region
}
