FROM golang:1.11.1 as builder

RUN apt-get install --reinstall make

RUN mkdir -p /tmp/squeeze

RUN git clone -b master https://github.com/agile6v/squeeze.git /tmp/squeeze

WORKDIR /tmp/squeeze

RUN make build

FROM alpine:3.8

RUN apk --no-cache add ca-certificates

COPY --from=builder /tmp/squeeze/squeeze /home/agile6v/squeeze

WORKDIR /home/agile6v/

CMD [./squeeze]