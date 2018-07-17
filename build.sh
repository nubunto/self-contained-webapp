#!/bin/sh

set -e

cd web
yarn build

cd ..
statik -src=web/build
go build -o app