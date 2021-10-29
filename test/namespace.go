package test

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"

	yscli "github.com/jibutech/backup-saas-client"
)

var (
	_ context.Context
)


func createNamespace(ctx context.Context, client testClient,  namespace string) error {
	ns := builder.ForNamespace(namespace).Result()
	_, err := client.clientGo.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}

func getNamespace(ctx context.Context, client testClient, namespace string) (*corev1api.Namespace, error) {
	return client.clientGo.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
}

func deleteNamespace(ctx context.Context, client testClient, namespace string, wait bool) error {
	if err := client.clientGo.CoreV1().Namespaces().Delete(ctx, namespace, metav1.DeleteOptions{}); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to delete the namespace %q", namespace))
	}
	if !wait {
		return nil
	}

	return waitutil.PollImmediateInfinite(5*time.Second,
		func() (bool, error) {
			if _, err := client.clientGo.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{}); err != nil {
				if apierrors.IsNotFound(err) {
					return true, nil
				}
				return false, err
			}
			logrus.Debugf("namespace %q is still being deleted...", namespace)
			return false, nil
		})
}

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
