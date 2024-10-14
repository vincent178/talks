#!/bin/bash

versions=(v1 v2 v3 v4 v5 v6)

for v in ${versions[@]}; do
    mkdir -p $v
    # cp proto/hello.proto $v -- oneoff
    echo "building $v/hello.proto"

    protoc --go_out=$v \
        --go_opt=module=github.com/vincent178/talks/proto-compatibility/$v \
        --go-grpc_out=$v \
        --go-grpc_opt=module=github.com/vincent178/talks/proto-compatibility/$v \
        $v/hello.proto
done

echo "convert slide markdown to pdf"
pandoc -t beamer slide.txt -o proto-compatibility.pdf
