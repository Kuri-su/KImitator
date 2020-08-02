# step 0
FROM golang:1.14.6-alpine3.11

USER root
LABEL maintainer="misaki.zhcy@gmail.com"

RUN apk update && apk add go git musl-dev xz binutils \
    && export GO111MODULE=on \
    && export GOPATH=/root/go \
    && git clone https://github.com/Kuri-su/confSyncer.git confsyncer \
    && cd confsyncer/cmd/confsyncer \
    && go build -o confsyncer \
    && mkdir -p /root/go/bin/ \
    && cp confsyncer /root/go/bin/

RUN wget https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz \
    && xz -d upx-3.96-amd64_linux.tar.xz \
    && tar -xvf upx-3.96-amd64_linux.tar \
    && cd upx-3.96-amd64_linux \
    && chmod a+x ./upx \
    && mv ./upx /usr/local/bin/ \
    && cd /root/go/bin \
    && strip --strip-unneeded confsyncer \
    && upx confsyncer \
    && chmod a+x ./confsyncer \
    && cp confsyncer /usr/local/bin

# step 1
FROM alpine:latest

COPY --from=0 /usr/local/bin/confsyncer /usr/local/bin/

CMD ["confsyncer","deamon"]
