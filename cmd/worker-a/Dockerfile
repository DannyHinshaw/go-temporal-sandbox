FROM golang:1.16-alpine as builder

WORKDIR /root
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo -ldflags="-w -s" -o worker-a cmd/worker-a/main.go

FROM alpine:latest

WORKDIR /worker-a
COPY --from=builder /root/worker-a .

ENTRYPOINT ["./worker-a"]