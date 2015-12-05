#!/bin/bash

if [[ -z "${APP_PORT}" ]]; then
  APP_PORT="8080"
fi

function __check()
{
  req="localhost:${APP_PORT}/${1}";
  curl --fail -XGET "${req}"
  ret=$?
  if [[ ! $ret = "0" ]]; then
    echo "ERROR: ${req}" > /dev/stderr
    exit $ret
  fi
}


# version
__check "version"

# encode
__check "encode?v=hello+world";

# decode
__check "decode?v=aGVsbG8gd29ybGQ=";
