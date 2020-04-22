FROM golang:1.12-alpine3.9 as builder


ENV VERSION ${INJECT_VERSION}

MAINTAINER Paweł Bojanowski <pbojanowski@protonmail.com>

WORKDIR /app
COPY config.yaml config.yaml
ENV SAILOR_URL https://github.com/hidalgopl/sailor/releases/download/${VERSION}/sailor

RUN apk add --no-cache ca-certificates wget curl
RUN curl -s https://api.github.com/repos/hidalgopl/sailor/releases/${INJECT_RELEASE_ID} | grep "browser_download_url.*" | cut -d '"' -f 4 | wget -qi -


FROM alpine:3.9

COPY --from=builder /app/sailor /usr/bin/sailor
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/config.yaml /etc/sailor/.secureapi.yml
RUN chmod +x /usr/bin/sailor

RUN addgroup -g 1000 sailor && \
    adduser -h /sailor -D -u 1000 -G sailor sailor

USER sailor
WORKDIR /sailor
