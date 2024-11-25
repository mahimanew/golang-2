```
cat /etc/vault.d/vault.hcl

sudo systemctl restart vault

vault operator diagnose -config /etc/vault.d/vault.hcl

vault operator init



 

 export VAULT_ADDR=https://3.83.153.69:8202

 export VAULT_SKIP_VERIFY=true

 vault login

 vault status

 vault operator unseal

 vault secrets enable kv-v2

 vault kv put kv-v2/vault-demo/mysecret foo=bar    -> this path is -> kv-v2/data/vault-demo/mysecret

 vault policy list



```

