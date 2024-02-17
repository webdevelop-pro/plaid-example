#!/usr/bin/env sh
set -xe

case $1 in

install)
  go get go get github.com/plaid/plaid-go@v20.1.0
  go get ./...
  ;;

run)
  set -a
  source .env
  go run main.go
  ;;

help)
  cat make.sh | grep "^[a-z-]*)"
  ;;

*)
  echo "unknown $1, try help"
  ;;

esac