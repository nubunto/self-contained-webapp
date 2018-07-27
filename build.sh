#!/bin/sh

set -e

cd web
yarn && yarn build

cd ..
go generate && go build