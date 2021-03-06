brew install consul
brew install protobuf

export GO111MODULE=on
export GOPROXY=https://goproxy.io

go get github.com/micro/micro
go get -u -v github.com/micro/go-micro
go get github.com/micro/protobuf/{proto,protoc-gen-go}

#~/goLang/src/github.com/micro$ls
#cli			mdns			protobuf
#go-micro		micro			protoc-gen-micro

consul agent -dev
#==> Starting Consul agent...
#==> Consul agent running!
#           Version: 'v0.8.1'
#           Node ID: '04145f71-d5c1-a469-c270-026f3911e0b1'

consul members
#Node       Address         Status  Type    Build  Protocol  DC
#localhost  127.0.0.1:8301  alive   server  0.8.1  2         dc1

#http://localhost:8500/ui/#/dc1/services/consul

protoc --go_out=plugins=micro:. ./proto/hello.proto

tree
#.
#|____proto
#| |____hello.pb.go
#| |____hello.proto
#|____setup.sh

