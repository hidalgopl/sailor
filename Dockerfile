FROM golang:1.12-alpine3.9 as builder

WORKDIR /app
COPY . .

RUN apk add --no-cache \
    gcc \
    git \
    linux-headers \
    make \
    musl-dev

RUN make build

FROM alpine:3.9

COPY --from=builder /app/bin/ /usr/local/bin/
COPY --from=builder /app/config.yaml /etc/sailor/config.yaml

RUN addgroup -g 1000 sailor && \
    adduser -h /sailor -D -u 1000 -G sailor sailor

USER sailor

ENTRYPOINT ["/usr/local/bin/sailor",  "run"]
