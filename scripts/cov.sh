#!/bin/bash -e
# Run from directory above via ./scripts/cov.sh

rm -rf ./cov
mkdir cov
go test --failfast -v -race -covermode=atomic -coverprofile=./cov/nats.out
go test --failfast -v -race -covermode=atomic -coverprofile=./cov/test.out -coverpkg=github.com/yanzongzhen/nats.go ./test
go test --failfast -v -race -covermode=atomic -coverprofile=./cov/builtin.out -coverpkg=github.com/yanzongzhen/nats.go/encoders/builtin ./test -run EncBuiltin
go test --failfast -v -race -covermode=atomic -coverprofile=./cov/protobuf.out -coverpkg=github.com/yanzongzhen/nats.go/encoders/protobuf ./test -run EncProto
gocovmerge ./cov/*.out > acc.out
rm -rf ./cov

# If we have an arg, assume travis run and push to coveralls. Otherwise launch browser results
if [[ -n $1 ]]; then
    $HOME/gopath/bin/goveralls -coverprofile=acc.out -service travis-ci
    rm -rf ./acc.out
else
    go tool cover -html=acc.out
fi
