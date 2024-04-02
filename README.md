# Nuke RLM Prometheus Exporter

[WIP] Project

## Description
This is a simple Prometheus exporter for Nuke RLM license server.
Data are exposed to Prometheus text format.

### Deploy with docker

```bash
docker build -t metricsrlm .
docker run -d -p 8081:8081 --env RLM_SERVER_HOST="<YOUR_RLM_HOST>" --env RLM_SERVER_PORT="4101" metricsrlm
```

### Deploy on K8S

Don't use ingress in production, testing purpose only.

Assuming you have a a k8s cluster with prometheus server running, in the same namespace.

> Don't forget to update values under ./kube/local/service.yml and ./kube/service.yml to replace <PROMETHEUS_SERVICE_NAME_HERE> with the name of the prometheus service.

```bash
kubectl apply -f .\kube\configmap.yml -f .\kube\namespace.yml -f .\kube\local\service.yml -f .\kube\deployment.yml
```

### Example of prometheus request
```
# To display only attribute of a specific license
sum by(id) (rlm_license_info_nuke_r{id="nuke_r_usage"})

# To display total
(sum in case of many versions)
sum by(id) (rlm_license_info_nuke_r{id="nuke_r_total"})
```

## Unit tests
To run unit tests, use the following command:
```bash
go test -v ./tests/...
```
