## migrate-docker-repos

A command-line tool for migrating docker images from source repository to target repository

To download the binary, see the [release page](https://github.com/bzon/migrate-docker-repos/releases/).

### Usage

Create the program `config.yaml` file.

```yaml
- tag: 0.19.0
  source_repo: quay.io/kubernetes-ingress-controller/nginx-ingress-controller
  target_repo: eu.gcr.io/your-org/nginx-ingress-controller # gcr
- tag: v0.11.0
  source_repo: quay.io/thanos/thanos
  target_repo: your-org/thanos

```

Run the program.

```bash
# Ensure to setup your docker login to target repos if necessary
# For example:
docker login -u _json_key -p "$GCLOUD_SERVICE_KEY" eu.gcr.io

# Run the program using Go
go run main.go

# Run the program using the downloaded binary
./migrate-docker-repos
```
