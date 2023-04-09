################################
# STEP 1 build binary
################################
FROM golang:1.19-alpine as builder
ARG VERSION=latest

RUN adduser --uid 1500 -D finna

WORKDIR $GOPATH/src/github.com/markliederbach/finna/
COPY . /go/src/github.com/markliederbach/finna

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build --ldflags "-s -w -X main.Version=${VERSION}" \
    -o /finna cmd/api/main.go

############################
# STEP 2 build a tiny image
############################
FROM scratch

USER finna
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /finna /finna
COPY --from=builder /go/src/github.com/markliederbach/finna/frontend /frontend

EXPOSE 8080/tcp
ENTRYPOINT ["/finna"]
