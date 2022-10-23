FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokene/keyserver-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/keyserver-svc /go/src/gitlab.com/tokene/keyserver-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/keyserver-svc /usr/local/bin/keyserver-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["keyserver-svc"]
