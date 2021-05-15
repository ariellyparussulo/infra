FROM golang:1.16 AS build

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPATH=/build

WORKDIR /build

COPY ./ ./

# install external dependencies
RUN mkdir -p bin && mkdir -p pkg && \
 go get github.com/lib/pq && \
 cd /build/src/commons && go install . commons && \
 cd /build/src/users && go install . users && \
 cd /build/src/http-server && go build -o /build/main

FROM alpine:latest AS runtime

WORKDIR /server

COPY --from=build /build/main /server/main

ENTRYPOINT [ "./main" ]