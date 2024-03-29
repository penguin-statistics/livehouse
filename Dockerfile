FROM golang:1.21.3-alpine AS base
WORKDIR /app

# builder
FROM base AS builder
ENV GOOS linux
ENV GOARCH amd64

# build-args
ARG VERSION

RUN apk --no-cache add bash git openssh

# modules: utilize build cache
COPY go.mod ./
COPY go.sum ./

# RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .

# inject versioning information & build the binary
RUN export BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ"); go build -o livehouse -ldflags "-X exusiai.dev/livehouse/internal/pkg/bininfo.Version=$VERSION -X exusiai.dev/livehouse/internal/pkg/bininfo.BuildTime=$BUILD_TIME" .

# runner
FROM base AS runner
RUN apk add --no-cache libc6-compat tini
# Tini is now available at /sbin/tini

COPY --from=builder /app/livehouse /app/livehouse
EXPOSE 8080

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/app/livehouse" ]
