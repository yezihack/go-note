#!/usr/bin/env sh

# Install proto3 from source
#  brew install autoconf automake libtool
#  git clone https://github.com/google/protobuf
#  ./autogen.sh ; ./configure ; make ; make install
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples

#protoc --go_out=plugins=grpc+grpcx:. *.proto
protoc --go_out=plugins=grpc:. *.proto

sed -i "" "s/TODO: replace this with your service name/go_love_ticket/g" *.pb.go

sed -i "" "s/,omitempty//g" *.pb.go
