# envoy-consul-sds - Envoy Consul Service Discovery Service

## Note
This tutorial is based on [Kelsey Hightower](https://github.com/kelseyhightower)'s [kubernetes-envoy-sds](https://github.com/kelseyhightower/kubernetes-envoy-sds) tutorial.

`envoy-consul-sds` service implements the [Envoy SDS API](https://lyft.github.io/envoy/docs/configuration/cluster_manager/sds_api.html) on top of [Consul Health Endpoint API](https://www.consul.io/api/health.html). `envoy-consul-sds` service returns a list of healthy endpoints for Envoy to use as upstream backends for a cluster.

Each Consul service can be referenced in the Envoy config file by its DNS name. For example a service named `nginx` can be referenced as `nginx.service.{datacenter_name}.consul` in the Envoy config file.

## Tutorial

In this turorial we will deploy Envoy, `envoy-consul-sds` service and `nginx` service on [Nomad](https://www.nomadproject.io/). We will use [Consul](https://www.consul.io/) for service discovery using Consul's DNS interface.

* [Bootstrap the Nomad Infrastructure on Google Compute Engine](./docs/bootstrap-nomad-google-compute-engine.md)
* [Deploy Envoy Consul Service Discovery Service](./docs/deploy-envoy-consul-sds.md)
* [Deploy Envoy System Job](./docs/deploy-envoy-system-job.md)
* [Deploy Nginx Service](./docs/deploy-nginx-service.md)