# gcs-crd

## How to use it

### yaml file
```
apiVersion: storage.yo.gcs/v1
kind: Gcs
metadata:
  name: gcs-sample
spec:
  projectID: projectID
  bucketName: sample-name-yo
```

### command

```
kubectl apply -f config/samples/storage_v1_gcs.yaml
```
