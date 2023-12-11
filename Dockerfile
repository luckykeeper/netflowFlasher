FROM golang:1.21.5-alpine3.19 AS builder

WORKDIR /go/src/

COPY ./ /go/src/

RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod tidy && \
    go build && \
    chmod +x ./netflowFlasher

FROM alpine

LABEL netflowflasher.image.author="Luckykeeper<https://luckykeeper.site|luckykeeper@luckykeeper.site|https://github.com/luckykeeper>"
LABEL maintainer="Luckykeeper<https://luckykeeper.site|luckykeeper@luckykeeper.site|https://github.com/luckykeeper>"
WORKDIR /app
COPY --from=builder /go/src/netflowFlasher /app/netflowFlasher
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

ENTRYPOINT /app/netflowFlasher
