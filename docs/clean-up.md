# Clean Up

Delete all Nomad servers and workers

```bash
gcloud -q compute instances delete \
  nomad-1 nomad-2 nomad-3 \
  nomad-worker-1 nomad-worker-2 nomad-worker-3 nomad-worker-4 nomad-worker-5
```
