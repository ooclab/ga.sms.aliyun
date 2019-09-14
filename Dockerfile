FROM golang:1.12 AS builder

WORKDIR /go/src/github.com/ooclab/ga.sms.aliyun
COPY . .
RUN make build-static


FROM alpine:3.10.2
WORKDIR /work
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true
COPY --from=builder /go/src/github.com/ooclab/ga.sms.aliyun /usr/bin/ga.sms.aliyun
COPY --from=builder /go/src/github.com/ooclab/ga.sms.aliyun/api.yml /work/api.yml
EXPOSE 3000
CMD ["/usr/bin/ga.sms.aliyun", "serve", "-v"]
