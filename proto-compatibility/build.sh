#!/bin/bash

versions=(v1 v2 v3 v4 v5 v6)

for v in ${versions[@]}; do
    mkdir -p $v
    # cp proto/hello.proto $v
    echo "building $v/hello.proto"
    rm -rf $v/$v
    protoc --go_out=$v \
        --go_opt=module=github.com/vincent178/talks/proto-compatibility/$v \
        --go-grpc_out=$v \
        --go-grpc_opt=module=github.com/vincent178/talks/proto-compatibility/$v \
        $v/hello.proto
done

pandoc -t beamer slide.txt -o proto-compatibility.pdf

# docker run --rm -it \
#   -v $(pwd):/workspace \
#   linuxserver/ffmpeg \
#   -i /workspace/frames%04d.png \
#   -vcodec mjepg \
#   -y /workspace/movie.mp4
