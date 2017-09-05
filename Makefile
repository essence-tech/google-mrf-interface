default: interface

.phony: interface
interface:
	protoc \
		-I. \
		-I/usr/local/lib \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/googleapis/googleapis \
		--go_out=plugins=grpc:googlemrf \
		--grpc-gateway_out=logtostderr=true:googlemrf \
		google_mrf.proto
	python \
		-m grpc.tools.protoc \
		-I. \
		-I/usr/local/lib \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/googleapis/googleapis \
		--python_out=googlemrf/. \
		--grpc_python_out=googlemrf/. \
		google_mrf.proto
