# syntax=docker/dockerfile:1
# check=error=true

FROM --platform=$BUILDPLATFORM golang:latest AS builder

ARG TARGETOS
ARG TARGETARCH

ENV GOCACHE=/root/.cache/go-build

WORKDIR /app

COPY main.go main.go

RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" \
    --mount=type=cache,id=gobuild,target="/root/.cache/go-build" \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o app main.go

FROM scratch AS final

COPY --from=builder /app/app .

ENTRYPOINT [ "/app" ]