default: interface

PWD := $(shell pwd)
PROTOC := docker run --rm -v $(PWD):$(PWD) -w $(PWD) znly/protoc

.phony: interface
interface:
	$(PROTOC) \
		-I. \
		--go_out=plugins=grpc:googlemrf \
		--grpc-gateway_out=logtostderr=true:googlemrf \
		google_mrf.proto
	$(PROTOC) \
		-I. \
	    --plugin=protoc-gen-grpc=/usr/bin/grpc_python_plugin \
		--python_out=googlemrf/. \
		--grpc_out=googlemrf/. \
		google_mrf.proto

.phony: client
client:
	docker run --rm --net="host" -it -v `pwd`:/proto jfyne/grpcc -a ess-lon-cmo-003:4041 -i -p google_mrf.proto
