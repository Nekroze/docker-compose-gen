#!/bin/sh
set -euf
docker-compose build tests
docker-compose run tests
