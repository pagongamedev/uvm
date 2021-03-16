# Multi State Build
FROM golang:latest AS builder

WORKDIR /workspace
# Gitlab Access Token : Permission read_registry for Pull Private Repo Go Package for Build
RUN echo "machine ${CI_PLATFORM} login ${GITLAB_LOGIN} password ${GITLAB_TOKEN}" > ~/.netrc
ENV GO111MODULE on
ENV GOARCH amd64
ENV CGO_ENABLED 0
ENV GOOS linux

COPY . .
RUN go build  -o /workspace/app -a -ldflags '-w -s' /workspace/uvm.go

FROM alpine

# Load Argument

COPY --from=builder /workspace/app /app
RUN apk add --no-cache --upgrade curl
EXPOSE 8080 8081 8082
ENTRYPOINT ["/app"]
