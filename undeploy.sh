#!/bin/bash

HOST_NS=backup-saas-system
TENANT_NS_PREFIX=backup-tenant-

echo "stopping all controllers in host namespace"
kubectl -n ${HOST_NS} delete --ignore-not-found=true deployment jibu-baas-rest-server
kubectl -n ${HOST_NS} delete --ignore-not-found=true deployment jibu-tenant-operator
kubectl -n ${HOST_NS} get tenants.ys.jibudata.com -o go-template='{{range $ns := .items}}{{$ns.metadata.name}}{{printf "\n"}}{{end}}' | while read -r tenant;do
    echo "deleting tenant ${tenant}"
    TENANT_NS=${TENANT_NS_PREFIX}${tenant}
    echo "deleting depeloyments in ${TENANT_NS}"
    kubectl -n ${TENANT_NS} delete --ignore-not-found=true deployment qiming-operator-${tenant}
    kubectl -n ${TENANT_NS} delete --ignore-not-found=true deployment mig-controller-${tenant}
    kubectl -n ${TENANT_NS} delete --ignore-not-found=true deployment ui-discovery-${tenant}

    # TODO: delete ns, clusterrolebindings from managed clusters
    # # load kubeconfig of tenant, then issue below commands
    # kubectl delete --ignore-not-found=true namespace qiming-backend
    # kubectl delete --ignore-not-found=true ClusterRoleBinding qiming-backup-admin-rb
    # for clusterrole in "qiming-operator-proxy-role" "qiming-operator-role"; do
    #     kubectl get ClusterRoleBinding --all-namespaces -o go-template='{{range $cr := .items}}{{if eq $cr.roleRef.name "'${clusterrole}'"}}{{$cr.metadata.name}}{{"\n"}}{{end}}{{end}}' | while read -r rolebinding;do
    #         kubectl delete --ignore-not-found=true ClusterRoleBinding ${rolebinding}
    #     done
    # done

    # remove finalizers
    cat <(kubectl api-resources --namespaced=true --api-group=ys.jibudata.com -o name) <(kubectl api-resources --namespaced=true --api-group=migration.yinhestor.com -o name) | while read -r res; do
        echo "clearing finalizer from all ${res} in ${TENANT_NS}"
        # kubectl -n ${TENANT_NS} patch ${res} --all -p '{"metadata":{"finalizers":[]}}' --type=merge
        kubectl -n ${TENANT_NS} get ${res} -o go-template='{{range $ns := .items}}{{$ns.metadata.name}}{{printf "\n"}}{{end}}' | while read -r cr;do
        echo "clearing finalizer (if exists) from ${res}: ${cr} in ${TENANT_NS}"
            kubectl -n ${TENANT_NS} patch ${res} ${cr} -p '{"metadata":{"finalizers":[]}}' --type=merge
        done
    done
    echo "clearing finalizer of namespace ${TENANT_NS}"
    kubectl patch namespace ${TENANT_NS} -R -p '{"metadata":{"finalizers":[]}}' --type=merge
    echo "deleting tenant namespace ${TENANT_NS}"
    kubectl delete --ignore-not-found=true namespace ${TENANT_NS}
    # kubectl -n ${HOST_NS} delete --ignore-not-found=true tenants.ys.jibudata.com ${tenant}
    echo "clearing finalizer of tenant ${tenant} in ${HOST_NS}"
    kubectl -n ${HOST_NS} patch tenants.ys.jibudata.com ${tenant} -p '{"metadata":{"finalizers":[]}}' --type=merge
done

echo "deleting host namespace ${HOST_NS}"
kubectl delete --ignore-not-found=true namespace ${HOST_NS}

echo "deleting crds"
kubectl delete --ignore-not-found=true -f deploy/crds.yaml
echo "deleting backend controller rbac"
kubectl delete --ignore-not-found=true -f deploy/backend-controller-rbac.yaml
echo "deleting deployment"
kubectl delete --ignore-not-found=true -f deploy/deployment.yaml
echo "deleting tenant operator deployment"
kubectl delete --ignore-not-found=true -f deploy/tenant-operator-deployment.yaml

echo "done"
echo "-------------------------------------------------------------------------------------------------------"
echo "|please delete below resources from all managed clusters:                                             |"
echo "|Namespace: qiming-backend                                                                            |"
echo "|ClusterRoleBinding: qiming-backup-admin-rb, velero-installer-rolebinding                             |"
echo "|all ClusterRoleBinding that bound to ClusterRole qiming-operator-role and qiming-operator-proxy-role |"
echo "-------------------------------------------------------------------------------------------------------"

exit 0
