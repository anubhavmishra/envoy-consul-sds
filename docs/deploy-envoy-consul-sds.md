# Deploy Envoy Consul Service Discovery Service

The `envoy-consul-sds` service will run as a Nomad service anywhere in the Nomad cluster. This service will be accessible via Consul DNS interface at `envoy-consul-sds.service.dc1.consul`.

## Prerequisites

SSH into a Nomad Server

```bash 
gcloud compute ssh nomad-1
```

Clone [envoy-consul-sds](https://github.com/anubhavmishra/envoy-consul-sds) git repo

```bash
nomad-1 $ git clone https://github.com/anubhavmishra/envoy-consul-sds.git
```

## Deploy Envoy Consul Service Discovery Service

```bash
nomad-1 $ cd envoy-consul-sds
```

```bash
nomad-1 $ nomad plan jobs/envoy-consul-sds.nomad
```

```bash
nomad-1 $ nomad run jobs/envoy-consul-sds.nomad
```

Check if the service job is up and running

```bash
nomad-1 $ nomad status envoy-consul-sds
```

```bash
nomad-1 $ dig envoy-consul-sds.service.dc1.consul 
```

```bash
nomad-1 $ curl -i http://envoy-consul-sds.service.dc1.consul:8080
```

Next, [Deploy Envoy System Job](./deploy-envoy-system-job.md)