#!/bin/bash
set -e -v

docker build --rm -t auth-service .
