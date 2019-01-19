#!/bin/bash
sed -i "/- name: RAND_NUM/{ n;s/\(value: \).*/\1num$RANDOM/ }" $1.yaml
