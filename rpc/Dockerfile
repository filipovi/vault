FROM golang:1.10 AS builder
WORKDIR /go/src/github.com/filipovi/vault/rpc
COPY . /go/src/github.com/filipovi/vault/rpc
RUN go get github.com/micro/go-micro && \
    go get github.com/filipovi/master-password-generator
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rpc rpc.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/filipovi/vault/rpc .
CMD ["./rpc"]

