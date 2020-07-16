#!/bin/bash

rm -f ./pkg/pb/*.go

PROTO_FILES=$(ls ./pkg/pb/proto/*.proto | xargs -n1 basename | tr -s '\n' ' ')

protoc --go_out ./ --proto_path ./pkg/pb/proto $PROTO_FILES

echo -e "/*\n$(cat pkg/pb/proto/protocol.txt)\n*/\npackage pb\n" >> ./pkg/pb/doc.go