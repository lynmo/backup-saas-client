package sample

const (
	apiEndpoint = "http://127.0.0.1:31800"
	tenantID    = "go-client-test"
	clusterName = "go-client-test-k8s"

	storageName = "go-client-test-s3"
	accessKey   = ""
	bucketName  = ""
	s3URL       = "http://s3.pek3b.qingstor.com/"
	secretKey   = ""
	region      = "pek3b"

	backupPlanName           = "go-client-test-backupplan"
	backupNamespace          = "wordpress"
	backupJobName            = "go-client-test-backupjob"
	backupJobForDeletingName = "go-client-test-backupjob-deleting"

	restorePlanName            = "go-client-test-restoreplan"
	restoreDestNamespace       = "wordpress-restored"
	restorePlanForDeletingName = "go-client-test-restoreplan-deleting"

	restoreJobName            = "go-client-test-restorejob"
	restoreJobForDeletingName = "go-client-test-restorejob-deleting"
)
