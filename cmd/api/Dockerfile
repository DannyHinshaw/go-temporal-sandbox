FROM golang:1.16-alpine as builder

WORKDIR /root
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo -ldflags="-w -s" -o api cmd/api/main.go

FROM alpine:latest

WORKDIR /api
COPY --from=builder /root/api .

ENTRYPOINT ["./api"]