FROM golang:1.11.1 as builder

RUN apt-get install --reinstall make

RUN mkdir -p /tmp/squeeze

RUN git clone -b master https://github.com/agile6v/squeeze.git /tmp/squeeze

WORKDIR /tmp/squeeze

RUN go version && \
        env GOOS=linux GOARCH=amd64 \
        make build

FROM alpine:3.8

RUN apk add --no-cache \
        libc6-compat

COPY --from=builder /tmp/squeeze/squeeze /bin/squeeze

CMD ["squeeze"]
