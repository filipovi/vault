FROM golang:1.10 AS builder
WORKDIR /go/src/github.com/filipovi/vault/api
COPY . /go/src/github.com/filipovi/vault/api
RUN go get -v github.com/go-chi/chi && \
    go get -v github.com/filipovi/redis && \
    go get -v github.com/golang/protobuf/proto && \
    go get -v github.com/micro/go-micro
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/filipovi/vault/api .
COPY --from=builder /go/src/github.com/filipovi/vault/api/docker/config.json .
CMD ["./api"]

