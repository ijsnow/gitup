#!/bin/bash
set -e -v

docker run -it -d -p 3000:3000 auth-service
