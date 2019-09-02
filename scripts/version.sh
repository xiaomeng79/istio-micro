#!/bin/bash

set -x

VERSION="$(git describe --abbrev=0 --tags)"
VERSION=${VERSION:-'0.0.0'}

MAJOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
MINOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
PATCH="${VERSION%%.*}"; VERSION="${VERSION#*.}"

echo "当前版本为:"$(git describe --abbrev=0 --tags)

PATCH="$((PATCH+1))"

#TAG="${1}"

#if [ "${TAG}" = "" ]; then
#  TAG="${MAJOR}.${MINOR}.${PATCH}"
#fi
TAG="${MAJOR}.${MINOR}.${PATCH}"

echo "下一个版本为:" ${TAG}

export Version=$(git describe --abbrev=0 --tags)
#
#git tag -a  -m "Relase ${TAG}" "${TAG}"
