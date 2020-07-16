SHELL := /bin/bash

.PHONY: checkdeps
checkdeps:
	wget -V > /dev/null
	curl -V > /dev/null
	jq -V > /dev/null

.PHONY: getproto
getproto: checkdeps
	./scripts/get_proto.sh