# envoy-consul-sds - Envoy Consul Service Discovery Service

This tutorial is based on [Kelsey Hightower](https://github.com/kelseyhightower)'s [kubernetes-envoy-sds](https://github.com/kelseyhightower/kubernetes-envoy-sds) tutorial but using [Consul](https://consul.io) and [Nomad](https://www.nomadproject.io/).

`envoy-consul-sds` service implements the [Envoy SDS API](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/service_discovery.html#service-discovery-service-sds) on top of [Consul Health Endpoint API](https://www.consul.io/api/health.html). `envoy-consul-sds` service returns a list of healthy endpoints for Envoy to use as upstream backends for a cluster. Each Consul service can be referenced in the Envoy config file by its DNS name.

**NOTE: This project uses [Envoy API v1](https://www.envoyproxy.io/docs/envoy/latest/api-v1/api#). It doesn't use Envoy xDS for configuration. Please read [Envoy API v2](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api) for more information.** 

## Goal

In this tutorial we will run Nginx on [Nomad](https://www.nomadproject.io/) and register it in [Consul](https://www.consul.io/). Then we will use Envoy to access Nginx using Consul DNS interface.

The idea is to explore ways [Envoy](https://lyft.github.io/envoy/) can integrate with applications running on Nomad using Consul.

## Tutorial

* [Bootstrap the Nomad Infrastructure on Google Compute Engine](./docs/bootstrap-nomad-google-compute-engine.md)
* [Deploy Envoy Consul Service Discovery Service](./docs/deploy-envoy-consul-sds.md)
* [Deploy Envoy System Job](./docs/deploy-envoy-system-job.md)
* [Deploy Nginx Service](./docs/deploy-nginx-service.md)
* [Clean up](./docs/clean-up.md)
