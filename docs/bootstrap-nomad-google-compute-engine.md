# Bootstrap the Nomad Infrastructure on Google Compute Engine

We will use [Google Compute Engine](https://cloud.google.com/compute/) to bootstrap the infrastructure required to run a Nomad cluster.

## Prerequisites

This tutorial assumes that you have already installed [Google Cloud SDK](https://cloud.google.com/sdk/) or use [Google Cloud Shell](https://cloud.google.com/shell/docs/) in your [Google Cloud Platform](https://cloud.google.com/) account.

```bash
gcloud version
```

Set default region and zone

```bash
gcloud config set compute/region us-east1
```

```bash
gcloud config set compute/zone us-east1-b
```

Clone [envoy-consul-sds](https://github.com/anubhavmishra/envoy-consul-sds) git repo

```bash
git clone https://github.com/anubhavmishra/envoy-consul-sds.git
```

```bash
cd envoy-consul-sds
```

## Bootstrap a Nomad Cluster

First we will install Nomad, Consul and dnsmasq

```bash
gcloud compute instances create nomad-1 nomad-2 nomad-3 \
  --image-project ubuntu-os-cloud \
  --image-family ubuntu-1604-lts \
  --boot-disk-size 150GB \
  --machine-type n1-standard-2 \
  --can-ip-forward \
  --metadata-from-file startup-script=scripts/bootstrap-server.sh
```

Join Nomad nodes

```bash
gcloud compute ssh nomad-1
```

```bash
nomad server-join nomad-2 nomad-3
```

Check Nomad cluster status

```bash
nomad status
```

Join Consul servers

```bash
consul join nomad-2 nomad-3
```

List Consul members

```bash
consul members
```

## Bootstrap a Nomad Workers

```bash
gcloud compute instances create nomad-worker-1 nomad-worker-2 nomad-worker-3 nomad-worker-4 nomad-worker-5 \
  --image-project ubuntu-os-cloud \
  --image-family ubuntu-1604-lts \
  --boot-disk-size 150GB \
  --machine-type n1-standard-2 \
  --can-ip-forward \
  --metadata-from-file startup-script=scripts/bootstrap-client.sh
```

```bash
gcloud compute ssh nomad-1
```

Check Nomad node status

```bash
nomad node-status
```

Up next, [Deploy Envoy Consul Service Discovery Service](./docs/deploy-envoy-consul-sds.md)