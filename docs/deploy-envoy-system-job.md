# Deploy Envoy System Job

We will now run [Envoy](https://lyft.github.io/envoy/) as a system job in Nomad so it runs on all nodes in the Nomad cluster. Envoy will use [envoy.json](https://gist.github.com/anubhavmishra/afe699320bdc4d855d13e7cc244822e0) config file.

## Prerequisite

SSH into a Nomad Server

```bash 
gcloud compute ssh nomad-1
```

## Deploy Envoy 

```bash
cd envoy-consul-sds
```

```bash
nomad plan jobs/envoy.nomad
```

```bash
+ Job: "proxy"
+ Task Group: "proxy" (5 create)
  + Task: "envoy" (forces create)
Scheduler dry-run:
- All tasks successfully allocated.
Job Modify Index: 0
To submit the job with version verification run:
nomad run -check-index 0 jobs/envoy.nomad
When running the job with the check-index flag, the job will only be run if the
server side version matches the job modify index returned. If the index has
changed, another user has modified the job and the plan's results are
potentially invalid.
```

```bash
nomad run jobs/envoy.nomad
```

Check if the system job is up and running

```bash
nomad status proxy
```

Next [Deploy Nginx Service](./deploy-nginx-service.md)