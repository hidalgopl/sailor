FROM golang:1.12-alpine3.9 as builder

WORKDIR /app
COPY . .

RUN apk add --no-cache \
    gcc \
    git \
    gettext \
    linux-headers \
    make \
    musl-dev

RUN make build-staging

FROM alpine:3.9

COPY --from=builder /app/bin/ /usr/local/bin/
COPY --from=builder /app/config.yaml /etc/sailor/.secureapi.yml

RUN addgroup -g 1000 sailor && \
    adduser -h /sailor -D -u 1000 -G sailor sailor

USER sailor

