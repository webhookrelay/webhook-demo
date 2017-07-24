# webhook-demo

This is a sample application to demonstrate public webhook relay to internal services


## Helm upgrade path

```
helm install --name=whd -f webhookdemo/values.yaml ./webhookdemo
```

```
helm upgrade whd ./webhookdemo --reuse-values --set image.tag=0.0.11
```

Delete it:

```
helm delete whd
```