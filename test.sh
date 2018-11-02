#!/bin/sh
set -euf
docker-compose build tests
trap 'docker-compose down' EXIT
docker-compose run --rm tests
