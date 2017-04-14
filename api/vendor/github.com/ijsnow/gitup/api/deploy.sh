#!/bin/bash
set -e -v

ssh parkn-staging 'bash -s' <<'ENDSSH'
  rm -rf app
ENDSSH

rsync -avzh --exclude-from '.gitignore' . parkn-staging:/home/ubuntu/app
rsync -avzh ./vendor/ parkn-staging:/home/ubuntu/app/vendor/
scp .env parkn-staging:/home/ubuntu/app/.env

ssh parkn-staging 'bash -s' <<'ENDSSH'
  docker rm -f `docker ps -a -q`
  docker rmi -f `docker images -q`
  cd app
  ./build.sh
  ./run.sh
ENDSSH
