# Deploy Envoy System Job

We will now run [Envoy](https://lyft.github.io/envoy/) as a system job in Nomad so it runs on all nodes in the Nomad cluster.

## Prerequisite

SSH into a Nomad Server

```bash 
gcloud compute ssh nomad-1
```

## Deploy Envoy 

```bash
nomad-1 $ cd envoy-consul-sds
```

```bash
nomad-1 $ nomad plan jobs/envoy.nomad
```

```bash
nomad-1 $ nomad run jobs/envoy.nomad
```

Check if the system job is up and running

```bash
nomad-1 $ nomad status proxy
```

Next [Deploy Nginx Service](./deploy-nginx-service.md)