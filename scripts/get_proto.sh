#!/bin/bash
PROTO_URL="https://api.github.com/repos/deptofdefense/AndroidTacticalAssaultKit-CIV/contents/commoncommo/core/impl/protobuf?ref=master"
GO_PKG_OPT=""

set -e

GITHUB_JSON=$(curl -s $PROTO_URL) 
PROTO_FILES=($(echo "$GITHUB_JSON" | jq -r .[]."name"))
RAW_URLS=($(echo "$GITHUB_JSON" | jq -r .[]."download_url"))

mkdir -p ./pkg/pb/proto
cd ./pkg/pb/proto
find . -mindepth 1 -delete
LEN=$(expr ${#PROTO_FILES[*]} - 1 )

for i in $(seq 0 $LEN); do
  echo "$i --- ${PROTO_FILES[$i]}"
  echo "$i --- ${RAW_URLS[$i]}"
 echo
  wget ${RAW_URLS[$i]}
  if [[ ${PROTO_FILES[$i]} == *proto ]]; then
    echo -e '\n\noption go_package = "pkg/pb";' >> ${PROTO_FILES[$i]}
  fi
done