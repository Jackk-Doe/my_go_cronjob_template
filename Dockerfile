# build executable binary
FROM golang:1.22.2-alpine as builder

ENV CGO_ENABLED 0
ENV GOOS "linux"
ENV GOARCH "amd64"

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN apk add --no-cache ca-certificates git tzdata && \
  go mod tidy

COPY . .

RUN go build -ldflags "-s -w -extldflags '-static'" -installsuffix cgo -o /bin/cronjob_app cmd/cronjob/main.go

# Use alpine image as runtime
FROM alpine:3.16 as release

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/cronjob_app /bin/cronjob_app

ARG APP_VERSION
ARG BUILD_DATE
ENV APP_VERSION ${APP_VERSION}
ENV BUILD_DATE ${BUILD_DATE}

# Command to run 
ENTRYPOINT ["/bin/cronjob_app"]