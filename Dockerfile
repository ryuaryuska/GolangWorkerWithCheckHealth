FROM golang:1.14.2-alpine as builder
RUN apk add alpine-sdk
WORKDIR /go/app
COPY . /go/app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o main -tags musl

FROM alpine:latest as runner
WORKDIR /root/
COPY --from=builder /go/app/main .
COPY .env  .
EXPOSE 3004
ENTRYPOINT /root/main
