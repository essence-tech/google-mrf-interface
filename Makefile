default: interface

.phony: interface
interface:
	protoc \
		-I. \
		-I/usr/local/lib \
		-I$$GOPATH/src \
		--go_out=plugins=grpc:googlemrf \
		google_mrf.proto
	python \
		-m grpc.tools.protoc \
		-I. \
		-I/usr/local/lib \
		-I$$GOPATH/src \
		--python_out=googlemrf/. \
		--grpc_python_out=googlemrf/. \
		google_mrf.proto
