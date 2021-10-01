# samples as testing

steps to run:

Modify the config file to specify the s3 parameters

Put a kubeconfig.yaml file in currently directly, this kubeconfig will be used to create cluster cr. If the file does not exist, it will get the KUBECONFIG env variable firstly, then use the default file ~/.kube/config. Please ensure the server url in kubeconfig is accessible from inside of rest server.

Run below commands:
```
go test -v -run="RestAPIs"
go test -v -run="Backupplan"
go test -v -run="Backupjob"
go test -v -run="Restoreplan"
go test -v -run="Restorejob"
```
