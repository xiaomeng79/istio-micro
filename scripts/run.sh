#!/bin/bash

set -e

echo "start"

echo "$*"
args=
for i in $@; do
    args=$i" "
done
echo $args
$1 $args
