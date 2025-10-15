#!/bin/sh

for dir in `find . -name '*.go' | sed -e 's/[/][a-z_]*.go$//' | sort | uniq | grep -v examples`
do
  echo "====[ $dir ]================================================================"
  go test -v $dir/ || break
done
