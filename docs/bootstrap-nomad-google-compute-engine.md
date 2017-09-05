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

In this section we will create `nomad-1`, `nomad-2` and `nomad-3` instances.

First we will install Consul on them and [auto join](https://www.consul.io/docs/agent/options.html#google-compute-engine) it using GCE metadata.

Nomad will then [auto bootstrap](https://www.nomadproject.io/guides/cluster/automatic.html) itself using existing Consul cluster.

We will also install `dnsmasq` in order to use Consul DNS interface.

```bash
gcloud compute instances create nomad-1 nomad-2 nomad-3 \
  --image-project ubuntu-os-cloud \
  --image-family ubuntu-1604-lts \
  --boot-disk-size 300GB \
  --machine-type n1-standard-1 \
  --can-ip-forward \
  --scopes default,compute-ro \
  --tags "gce-envoy-consul-sds" \
  --metadata-from-file startup-script=scripts/bootstrap-server.sh
```

```bash
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-1].
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-2].
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-3].
NAME     ZONE        MACHINE_TYPE   PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP     STATUS
nomad-1  us-east1-b  n1-standard-1               10.x.0.2     x.x.x.x         RUNNING
nomad-2  us-east1-b  n1-standard-1               10.x.0.4     x.x.x.x         RUNNING
nomad-3  us-east1-b  n1-standard-1               10.x.0.3     x.x.x.x         RUNNING
```

```bash
gcloud compute ssh nomad-1
```

If you SSH keys for account don't exist then this will create them 
```bash
WARNING: The public SSH key file for gcloud does not exist.
WARNING: The private SSH key file for gcloud does not exist.
WARNING: You do not have an SSH key for gcloud.
WARNING: SSH keygen will be executed to generate a key.
This tool needs to create the directory
[/home/anubhavmishragclouddemo/.ssh] before being able to generate SSH
 keys.
Do you want to continue (Y/n)? 
```

List Consul members

```bash
consul members
```

```bash
Node     Address          Status  Type    Build  Protocol  DC
nomad-1  10.x.0.2:8301  alive   server  0.9.2  2         dc1
nomad-2  10.x.0.3:8301  alive   server  0.9.2  2         dc1
nomad-3  10.x.0.4:8301  alive   server  0.9.2  2         dc1
```

Check Nomad cluster status

```bash
nomad server-members
```

```bash
Name            Address     Port  Status  Leader  Protocol  Build  Datacenter  Region
nomad-1.global  10.x.0.2  4648  alive   true    2         0.6.2  dc1         global
nomad-2.global  10.x.0.3  4648  alive   false   2         0.6.2  dc1         global
nomad-3.global  10.x.0.4  4648  alive   false   2         0.6.2  dc1         global
```

Logout of `nomad-1` back to your Cloud Shell or local terminal

```bash
exit
```

## Bootstrap Nomad Workers

In this section we will create `nomad-worker-1`, `nomad-worker-2`, `nomad-worker-3`, `nomad-worker-4` and `nomad-worker-5` instances. We will install Nomad, Consul, dnsmasq and docker on the workers.

```bash
gcloud compute instances create nomad-worker-1 nomad-worker-2 nomad-worker-3 nomad-worker-4 nomad-worker-5 \
  --image-project ubuntu-os-cloud \
  --image-family ubuntu-1604-lts \
  --boot-disk-size 200GB \
  --machine-type n1-standard-1 \
  --can-ip-forward \
  --scopes default,compute-ro \
  --tags "gce-envoy-consul-sds" \
  --metadata-from-file startup-script=scripts/bootstrap-client.sh
```

```bash
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-worker-1].
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-worker-2].
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-worker-3].
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-worker-5].
Created [https://www.googleapis.com/compute/v1/projects/envoy-nomad-consul/zones/us-east1-b/instances/nomad-worker-4].
NAME            ZONE        MACHINE_TYPE   PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP     STATUS
nomad-worker-1  us-east1-b  n1-standard-1               10.x.0.6     x.x.x.x         RUNNING
nomad-worker-2  us-east1-b  n1-standard-1               10.x.0.5     x.x.x.x         RUNNING
nomad-worker-3  us-east1-b  n1-standard-1               10.x.0.8     x.x.x.x         RUNNING
nomad-worker-5  us-east1-b  n1-standard-1               10.x.0.7     x.x.x.x         RUNNING
nomad-worker-4  us-east1-b  n1-standard-1               10.x.0.9     x.x.x.x         RUNNING
```

```bash
gcloud compute ssh nomad-1
```

List Consul members

```bash
consul members
```

```bash
Node            Address          Status  Type    Build  Protocol  DC
nomad-1         10.142.0.2:8301  alive   server  0.9.2  2         dc1
nomad-2         10.142.0.4:8301  alive   server  0.9.2  2         dc1
nomad-3         10.142.0.3:8301  alive   server  0.9.2  2         dc1
nomad-worker-1  10.142.0.6:8301  alive   client  0.9.2  2         dc1
nomad-worker-2  10.142.0.5:8301  alive   client  0.9.2  2         dc1
nomad-worker-3  10.142.0.8:8301  alive   client  0.9.2  2         dc1
nomad-worker-4  10.142.0.9:8301  alive   client  0.9.2  2         dc1
nomad-worker-5  10.142.0.7:8301  alive   client  0.9.2  2         dc1
```

Check Nomad node status

```bash
nomad node-status
```

*Might have to wait a little bit for all workers to come up....*

```bash
ID        DC   Name            Class   Drain  Status
8dc3f6ef  dc1  nomad-worker-4  <none>  false  ready
5be8c4f2  dc1  nomad-worker-2  <none>  false  ready
27dd7a34  dc1  nomad-worker-3  <none>  false  ready
630c864c  dc1  nomad-worker-1  <none>  false  ready
2c66ca1c  dc1  nomad-worker-5  <none>  false  ready
```

Up next, [Deploy Envoy Consul Service Discovery Service](./deploy-envoy-consul-sds.md)
