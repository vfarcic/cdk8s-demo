echo "" | tee crds/crossplane.yaml
for CRD in \
    "vpcs.ec2.aws.crossplane.io" \
    "subnets.ec2.aws.crossplane.io" \
    "internetgateways.ec2.aws.crossplane.io" \
    "routetables.ec2.aws.crossplane.io" \
    "securitygroups.ec2.aws.crossplane.io" \
    "dbsubnetgroups.database.aws.crossplane.io" \
    "rdsinstances.database.aws.crossplane.io" \
    "providerconfigs.postgresql.sql.crossplane.io" \
    "databases.postgresql.sql.crossplane.io"
do
    echo "
---
" | tee -a crds/crossplane.yaml
    kubectl get \
        crd $CRD \
        --output yaml \
        | tee -a crds/crossplane.yaml
done
