#!/bin/sh
set -euf
export COMPOSE_FILE='testing/docker-compose.yml'
export COMPOSE_PROJECT_NAME='dcg'

docker-compose build tests
trap 'docker-compose down' EXIT
docker-compose run --rm tests
