FROM golang:1.13 as builder
WORKDIR /go/src/github.com/ReflecBeatCustom/haereticus
COPY . /go/src/github.com/ReflecBeatCustom/haereticus
# Build
RUN GO111MODULE=off make
# Copy the gpu-resource-serve into centos image for debug
FROM centos:latest
WORKDIR /
COPY --from=builder /go/src/github.com/ReflecBeatCustom/haereticus/_output/cmd/bin/haereticus .
COPY --from=builder /go/src/github.com/ReflecBeatCustom/haereticus/etc/server.toml .
ENTRYPOINT ["/haereticus"]