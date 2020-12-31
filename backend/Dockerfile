FROM golang:1.15.6-alpine3.12 AS builder

COPY . .

RUN apk --no-cache update && apk --no-cache add git && \
    go get -u github.com/gorilla/mux && \
    go build main.go

FROM alpine:latest

WORKDIR /server
EXPOSE 8000

COPY --from=builder /go/main .

ENTRYPOINT [ "./main" ]