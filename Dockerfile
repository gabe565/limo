ARG GO_VERSION=1.18

FROM --platform=$BUILDPLATFORM golang:$GO_VERSION-alpine as go-builder
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go .
COPY cmd/ cmd/
COPY internal/ internal/
ARG TARGETPLATFORM
# Set Golang build envs based on Docker platform string
RUN --mount=type=cache,target=/root/.cache/go-build set -x \
    && case "$TARGETPLATFORM" in \
        'linux/amd64') export GOARCH=amd64 ;; \
        'linux/arm/v6') export GOARCH=arm GOARM=6 ;; \
        'linux/arm/v7') export GOARCH=arm GOARM=7 ;; \
        'linux/arm64') export GOARCH=arm64 ;; \
        *) echo "Unsupported target: $TARGETPLATFORM" && exit 1 ;; \
    esac \
    && go build -ldflags="-w -s" ./cmd/limod


FROM alpine
WORKDIR /data
COPY --from=go-builder /go/src/app/limod /usr/local/bin

ENV LIMOD_ADDRESS :80
CMD ["limod"]
