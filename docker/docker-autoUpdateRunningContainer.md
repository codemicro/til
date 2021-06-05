# Automatically restart a running Docker container

This can be done with [Watchtower](https://github.com/containrrr/watchtower).

```bash
docker run -d --restart on-failure \
  --name watchtower \
  -v /root/.docker/config.json:/config.json \
  -v /var/run/docker.sock:/var/run/docker.sock \
  containrrr/watchtower:latest -i 1800 containerName
```

This command will start a Docker container that has access to the Docker config stored at `/root/.docker/config.json` (required for repo credentials to private repos), connected to the Docker socket, watching the running container `containerName` and polling every 1800 seconds (30 minutes) to check for updates to that container.
