```

helm repo add hashicorp https://helm.release.hashicorp.com
---------------------
copy values.yaml file from below link and create
https://github.com/hashicorp/vault-helm/blob/main/values.yaml


modify below items:

1)namespace: "vault"

2)externalVaultAddr: "pass your vault address"

3)injector:
  enabled: "true"
4) serviceAccount:
    create: true
    name: "your service account name"
    createSecret: true


--------------------

helm  install vault hashicorp/vault --create-namespace -f values.yaml -n vault



kubernetes_host -> run this command -> kubectl cluster-info
export KUBE_API_SERVER= your value
export KUBE_HOST=your host
echo $KUBE_HOST

token_reviewer_jwt -> run this command in bash -> kubectl get secrets -n vault vault-sa-token -o=jsonpath={.data.token} | base64 -d
export KUBE_JWT=$(kubectl get secrets -n vault vault-sa-token -o=jsonpath={.data.token} | base64 -d)
echo $KUBE_JWT


kubernetes_ca_cert-> run this command in bash -> kubectl get secrets -n vault vault-sa-token -o=jsonpath='{.data.ca\.crt}' | base64 -d
export KUBE_CERT=$(kubectl get secrets -n vault vault-sa-token -o=jsonpath='{.data.ca\.crt}' | base64 -d)
echo $KUBE_CERT


vault auth enable kubernetes




vault policy write mysecret - << EOF
path "kv-v2/data/vault-demo/mysecret" {
  capabilities = ["read"]
}
EOF

vault write auth/kubernetes/config token_reviewer_jwt=$KUBE_JWT kubernetes_host=$KUBE_HOST kubernetes_ca_cert="$KUBE_CERT"

vault write auth/kubernetes/role/vault-role \
    bound_service_account_names=vault-sa \
    bound_service_account_namespaces=vault \
    policies=default,mysecret \
    ttl=1h

vault write auth/kubernetes/role/<ROLE_NAME> \
    bound_service_account_names=<SERVICE_ACCOUNT> \
    bound_service_account_namespaces=<NAMESPACE> \
    policies=<POLICY_NAME> \
    ttl=24h


```
