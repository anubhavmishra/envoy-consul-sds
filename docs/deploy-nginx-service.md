# Deploy Nginx Service 

We will now deploy [Nginx](https://nginx.org/en/) service on Nomad.

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
nomad plan jobs/nginx.nomad
```

```bash
+ Job: "nginx"
+ Task Group: "webserver" (1 create)
  + Task: "nginx" (forces create)
Scheduler dry-run:
- All tasks successfully allocated.
Job Modify Index: 0
To submit the job with version verification run:
nomad run -check-index 0 jobs/nginx.nomad
When running the job with the check-index flag, the job will only be run if the
server side version matches the job modify index returned. If the index has
changed, another user has modified the job and the plan's results are
potentially invalid.
```

```bash
nomad run jobs/nginx.nomad
```

Check if the service is up and running

```bash
nomad status nginx
```

```bash
dig +noall +answer nginx.service.dc1.consul
```

Call Nginx using Consul service discovery

```bash
curl nginx.service.dc1.consul
```

```bash
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>
<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>
<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

SSH into `nomad-worker-1`

```bash
exit
```

```bash
gcloud compute ssh nomad-worker-1
```

Call Nginx using Envoy proxy

```bash
curl localhost:1010
```

```bash
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>
<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>
<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

**Congratulations! We made it. You are now hitting Nginx using Envoy proxy using Consul SDS Service.**

Exit `nomad-worker-1`

```bash
exit
```

Next [Clean up](./clean-up.md)

