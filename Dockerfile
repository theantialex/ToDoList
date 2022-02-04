# build binary
FROM golang:1.17.1-alpine AS build

ARG GOOS
ENV CGO_ENABLED=0 \
    GOOS=$GOOS \
    GOARCH=amd64 \
    CGO_CPPFLAGS="-I/usr/include" \
    UID=0 GID=0 \
    CGO_CFLAGS="-I/usr/include" \
    CGO_LDFLAGS="-L/usr/lib -lpthread -lrt -lstdc++ -lm -lc -lgcc -lz " \
    PKG_CONFIG_PATH="/usr/lib/pkgconfig"

ARG APP_PKG_NAME
WORKDIR /go/src/$APP_PKG_NAME
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY go.mod go.mod

RUN go mod vendor

RUN go build -v \
    -o /out/service \
    ./cmd/main.go

FROM alpine:3.8
WORKDIR /app
COPY --from=build /out/service /app/server
CMD ["/app/server"]