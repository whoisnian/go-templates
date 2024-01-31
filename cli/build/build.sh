#!/bin/bash -ex
#########################################################################
# File Name: build.sh
# Author: nian
# Blog: https://whoisnian.com
# Mail: zhuchangbao1998@gmail.com
# Created Time: 2024年02月01日 星期四 00时06分56秒
#########################################################################
# usage:
#   ./build.sh .              # build for current platform
#   ./build.sh all            # build for all platforms (linux-amd64, linux-arm64, windows-amd64, windows-arm64)
#   ./build.sh windows amd64  # build for specified platform

SCRIPT_DIR=$(dirname "$0")
SOURCE_DIR="$SCRIPT_DIR/.."

APP_NAME="cli"
BUILDTIME=$(date --iso-8601=seconds)
if [[ -z "$GITHUB_REF_NAME" ]]; then
  VERSION=$(git describe --tags || echo unknown)
else
  VERSION=$GITHUB_REF_NAME
fi

goBuild() {
  CGO_ENABLED=0 GOOS="$1" GOARCH="$2" go build -trimpath \
    -ldflags="-s -w \
    -X 'github.com/whoisnian/go-templates/cli/global.Version=${VERSION}' \
    -X 'github.com/whoisnian/go-templates/cli/global.BuildTime=${BUILDTIME}'" \
    -o "$3" "$SOURCE_DIR"
}

if [[ "$1" == '.' ]]; then
  goBuild $(go env GOOS) $(go env GOARCH) "$APP_NAME"
elif [[ "$1" == 'all' ]]; then
  goBuild linux amd64 "${APP_NAME}-linux-amd64-${VERSION}"
  goBuild linux arm64 "${APP_NAME}-linux-arm64-${VERSION}"
  goBuild windows amd64 "${APP_NAME}-windows-amd64-${VERSION}.exe"
  goBuild windows arm64 "${APP_NAME}-windows-arm64-${VERSION}.exe"
elif [[ "$#" == 2 ]]; then
  goBuild "$1" "$2" "${APP_NAME}-$1-$2-${VERSION}"
else
  echo "invalid build arguments"
  exit 1
fi