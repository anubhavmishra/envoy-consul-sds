# Deploy Nginx Service 

We will now deploy [Nginx](https://nginx.org/en/) service on Nomad.

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
nomad-1 $ nomad plan jobs/nginx.nomad
```

```bash
nomad-1 $ nomad run jobs/nginx.nomad
```

Check if the service is up and running

```bash
nomad-1 $ nomad status nginx
```

```bash
nomad-1 $ dig nginx.service.dc1.consul
```