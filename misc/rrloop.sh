#!/bin/sh

while :
do
  rr $@
  if [ $? -ne 0 ]; then
    echo "encountered non-zero exit code: $?";
    exit;
  fi
done