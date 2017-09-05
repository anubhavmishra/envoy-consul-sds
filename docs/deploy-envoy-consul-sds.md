# Deploy Envoy Consul Service Discovery Service

The `envoy-consul-sds` service will run as a Nomad service anywhere in the Nomad cluster. This service will be accessible via Consul DNS interface at `envoy-consul-sds.service.dc1.consul`.

## Prerequisites

SSH into a Nomad Server

```bash 
gcloud compute ssh nomad-1
```

Clone [envoy-consul-sds](https://github.com/anubhavmishra/envoy-consul-sds) git repo on `nomad-1`

```bash
git clone https://github.com/anubhavmishra/envoy-consul-sds.git
```

## Deploy Envoy Consul Service Discovery Service

```bash
cd envoy-consul-sds
```

```bash
nomad plan jobs/envoy-consul-sds.nomad
```

```bash
+ Job: "envoy-consul-sds"
+ Task Group: "webserver" (1 create)
  + Task: "envoy-consul-sds" (forces create)
Scheduler dry-run:
- All tasks successfully allocated.
Job Modify Index: 0
To submit the job with version verification run:
nomad run -check-index 0 jobs/envoy-consul-sds.nomad
When running the job with the check-index flag, the job will only be run if the
server side version matches the job modify index returned. If the index has
changed, another user has modified the job and the plan's results are
potentially invalid.
```

```bash
nomad run jobs/envoy-consul-sds.nomad
```

Check if the service job is up and running

```bash
nomad status envoy-consul-sds
```

```bash
dig +noall +answer envoy-consul-sds.service.dc1.consul
```

```bash
envoy-consul-sds.service.dc1.consul. 0 IN A     10.142.0.7
```

```bash
curl http://envoy-consul-sds.service.dc1.consul:8080
```

```bash
{"name":"envoy-consul-sds","version":"0.0.1"}
```

Next, [Deploy Envoy System Job](./deploy-envoy-system-job.md)