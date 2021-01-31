FROM golang:1.15.6-alpine3.12 AS builder

WORKDIR /go
COPY . .

RUN apk --no-cache update && apk --no-cache add git && \
    go get -u github.com/gorilla/mux && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/jinzhu/pq && \
    go get -u github.com/jinzhu/gorm/dialects/postgres && \
    go generate && go build

FROM alpine:latest

WORKDIR /server
EXPOSE 8000

COPY --from=builder /go/backend .

ENTRYPOINT [ "./backend" ]