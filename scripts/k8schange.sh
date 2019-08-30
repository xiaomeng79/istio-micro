#!/bin/bash
set -eu

source scripts/.variables.sh

sed -i "/- name: RAND_NUM/{ n;s/\(value: \).*/\1num$RANDOM/ }" $1.yaml
