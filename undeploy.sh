#!/bin/bash

if [[ "$1" == "--include-clusters" ]];then
    DELETE_MANAGED_CLUSTERS=1
fi

if [[ "$2" == "--leave-crds" ]];then
    LEAVE_CRDS=1
fi

HOST_NS=backup-saas-system
TENANT_NS_PREFIX=backup-tenant-

echo "stopping all controllers in host namespace"
kubectl -n ${HOST_NS} delete --ignore-not-found=true deployment jibu-baas-rest-server
kubectl -n ${HOST_NS} delete --ignore-not-found=true deployment jibu-tenant-operator
echo "deleting tenant workers"
kubectl get tenantworkers.ys.jibudata.com -A -o go-template='{{range $tw := .items}}{{$tw.metadata.name}}{{printf "\n"}}{{end}}' | while read -r tw;do
    WORKER_NS=${tw}
    echo "deleting depeloyments in ${WORKER_NS}"
    kubectl -n ${WORKER_NS} delete --ignore-not-found=true deployment ${tw}
    kubectl -n ${WORKER_NS} delete --ignore-not-found=true deployment mig-controller-rest
    kubectl -n ${WORKER_NS} delete --ignore-not-found=true deployment ui-discovery-rest
    echo "deleting tenant worker CR in ${WORKER_NS}"
    kubectl -n ${WORKER_NS} patch tenantworkers ${tw} -p '{"metadata":{"finalizers":[]}}' --type=merge
    kubectl -n ${WORKER_NS} delete --ignore-not-found=true tenantworkers ${tw}
    echo "deleting tenant worker NS ${WORKER_NS}"
    kubectl delete --ignore-not-found=true ns ${tw}
done
kubectl -n ${HOST_NS} get tenants.ys.jibudata.com -o go-template='{{range $ns := .items}}{{$ns.metadata.name}}{{printf "\n"}}{{end}}' | while read -r tenant;do
    echo "deleting tenant ${tenant}"
    TENANT_NS=${TENANT_NS_PREFIX}${tenant}
    echo "deleting depeloyments in ${TENANT_NS}"
    kubectl -n ${TENANT_NS} delete --ignore-not-found=true deployment qiming-operator-${tenant}
    kubectl -n ${TENANT_NS} delete --ignore-not-found=true deployment mig-controller-${tenant}
    kubectl -n ${TENANT_NS} delete --ignore-not-found=true deployment ui-discovery-${tenant}

    if [[ "${DELETE_MANAGED_CLUSTERS}" == "1" ]];then
        kubectl -n ${TENANT_NS} get clusters.ys.jibudata.com -o go-template='{{range $cluster := .items}}{{$cluster.metadata.name}}{{printf "\n"}}{{end}}' | while read -r cluster;do
            echo "delete resources from managed cluster ${cluster}"
            kubeconfig=$(mktemp /tmp/backup-saas-cluster-kubeconfig.XXXXXX)
            kubectl -n ${TENANT_NS} get clusters.ys.jibudata.com ${cluster} -o go-template='{{.spec.kubeconfig}}' > ${kubeconfig}

            kubectl --kubeconfig ${kubeconfig} delete --ignore-not-found=true namespace qiming-backend
            kubectl --kubeconfig ${kubeconfig} delete --ignore-not-found=true ClusterRoleBinding qiming-backup-admin-rb
            for clusterrole in "qiming-operator-proxy-role" "qiming-operator-role"; do
                kubectl --kubeconfig ${kubeconfig} get ClusterRoleBinding --all-namespaces -o go-template='{{range $cr := .items}}{{if eq $cr.roleRef.name "'${clusterrole}'"}}{{$cr.metadata.name}}{{"\n"}}{{end}}{{end}}' | while read -r rolebinding;do
                    kubectl --kubeconfig ${kubeconfig} delete --ignore-not-found=true ClusterRoleBinding ${rolebinding}
                done
            done

            rm ${kubeconfig}
        done
    fi

    # remove finalizers
    cat <(kubectl api-resources --namespaced=true --api-group=ys.jibudata.com -o name) <(kubectl api-resources --namespaced=true --api-group=migration.yinhestor.com -o name) | while read -r res; do
        echo "clearing finalizer from all ${res} in ${TENANT_NS}"
        kubectl -n ${TENANT_NS} get ${res} -o go-template='{{range $ns := .items}}{{$ns.metadata.name}}{{printf "\n"}}{{end}}' | while read -r cr;do
            echo "deleting ${res} ${cr} from ${TENANT_NS}"
            kubectl -n ${TENANT_NS} delete --ignore-not-found ${res} ${cr} --wait=false
            echo "clearing finalizer (if exists) from ${res}: ${cr} in ${TENANT_NS}"
            kubectl -n ${TENANT_NS} patch ${res} ${cr} -p '{"metadata":{"finalizers":[]}}' --type=merge
        done
    done
    echo "deleting tenant namespace ${TENANT_NS}"
    kubectl delete --ignore-not-found=true namespace ${TENANT_NS} --wait=false
    echo "clearing finalizer of namespace ${TENANT_NS}"
    kubectl patch namespace ${TENANT_NS} -R -p '{"metadata":{"finalizers":[]}}' --type=merge
    echo "deleting tenant ${tenant}"
    kubectl -n ${HOST_NS} delete --ignore-not-found=true tenants.ys.jibudata.com ${tenant} --wait=false
    echo "clearing finalizer of tenant ${tenant} in ${HOST_NS}"
    kubectl -n ${HOST_NS} patch tenants.ys.jibudata.com ${tenant} -p '{"metadata":{"finalizers":[]}}' --type=merge
done

echo "deleting host namespace ${HOST_NS}"
kubectl delete --ignore-not-found=true namespace ${HOST_NS}

if [[ "${LEAVE_CRDS}" == "1" ]];then
    echo "skipping crd/rbac deleting"
else
    echo "deleting crds"
    kubectl delete --ignore-not-found=true -f deploy/crds.yaml
    echo "deleting backend controller rbac"
    kubectl delete --ignore-not-found=true -f deploy/backend-controller-rbac.yaml
    echo "deleting deployment"
    kubectl delete --ignore-not-found=true -f deploy/deployment.yaml
    echo "deleting tenant operator deployment"
    kubectl delete --ignore-not-found=true -f deploy/tenant-operator-deployment.yaml
    echo "deleting tenant manager crd"
    kubectl delete --ignore-not-found=true -f deploy/tenant-manager-cr.yaml
fi

echo "done"
if [[ "${DELETE_MANAGED_CLUSTERS}" != "1" ]];then
    echo "-------------------------------------------------------------------------------------------------------"
    echo "|please delete below resources from all managed clusters:                                             |"
    echo "|Namespace: qiming-backend                                                                            |"
    echo "|ClusterRoleBinding: qiming-backup-admin-rb, velero-installer-rolebinding                             |"
    echo "|all ClusterRoleBinding that bound to ClusterRole qiming-operator-role and qiming-operator-proxy-role |"
    echo "-------------------------------------------------------------------------------------------------------"
fi

exit 0
