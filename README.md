# docker-compose-gen

Generate docker-compose config segments dynamically.

# Primary Use Case

It can be helpful to replace the default network with an externally defined one dynamically at time of orchestration.

`docker-compose-gen` can be used like so to make the default network containers are connected to an already defined network called lab.

```bash
docker-compose -f docker-compose.yml -f <(docker-compose-gen network --compose-version 3.1 --external-network lab) up
```

# Installation

## Docker

You may use the [automated docker image](https://hub.docker.com/r/nekroze/docker-compose-gen) to use `docker-compose-gen` without any other dependencies with the following:

```bash
docker run nekroze/docker-compose-gen:latest network --network backchannel
```

## Go

If you would like to compile and install/update `docker-compose-gen` locally and via the golang tool chain with:

```bash
go get -u github.com/Nekroze/docker-compose-gen/...
```
