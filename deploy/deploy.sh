#!/bin/sh -ex

KUBECONFIGDIR=${1}
SECRET=${2}
NAMESPACE=${3}


alias kubectl='kubectl --kubeconfig $KUBECONFIGDIR/config'

# secret template
cat > /tmp/gcr-sec <<END
---
kind: Secret
apiVersion: v1
metadata:
  name: $SECRET
  namespace: $NAMESPACE
type: kubernetes.io/dockerconfigjson
data:
  .dockerconfigjson: --secret--
END

secret()
{
    encoded_key=$(echo -n $DOCKER_AUTH_CONFIG | base64 | tr -d " \n\r" )
    sed -i -- "s/--secret--/${encoded_key}/g" /tmp/gcr-sec

    kubectl delete secret $SECRET -n $NAMESPACE
    kubectl apply -f /tmp/gcr-sec
}

setup()
{
    mkdir -p $KUBECONFIGDIR
    echo $KUBECONFIG > $KUBECONFIGDIR/config
    envsubst < deploy/deployment.template.yml > $KUBECONFIGDIR/deployment.yml
}

deploy()
{
    kubectl config use-context $1
    secret
    kubectl apply -f $KUBECONFIGDIR/deployment.yml
}

setup
deploy "gke"
#deploy "gce"
