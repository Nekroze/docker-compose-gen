#!/bin/sh
set -euf
docker-compose build gen
trap 'docker-compose down' EXIT
docker-compose run --rm gen "$@"
