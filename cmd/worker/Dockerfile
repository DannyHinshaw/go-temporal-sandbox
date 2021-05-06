FROM golang:1.16-alpine as builder

WORKDIR /root
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo -ldflags="-w -s" -o worker cmd/worker/main.go

FROM alpine:latest

WORKDIR /worker
COPY --from=builder /root/worker .

ENTRYPOINT ["./worker"]