#!/bin/bash

VERSION="$(git describe --abbrev=0 --tags)"
VERSION=${VERSION:-'0.0.0'}

MAJOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
MINOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
PATCH="${VERSION%%.*}"; VERSION="${VERSION#*.}"


PATCH="$((PATCH+1))"

#TAG="${1}"

#if [ "${TAG}" = "" ]; then
#  TAG="${MAJOR}.${MINOR}.${PATCH}"
#fi
TAG="${MAJOR}.${MINOR}.${PATCH}"

export Version=${TAG}

git tag -a  -m "Relase ${TAG}" "${TAG}"
