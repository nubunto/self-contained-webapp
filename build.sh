#!/bin/sh

set -e

cd web
yarn && yarn build

cd ..
statik -src=web/build
go build