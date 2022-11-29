ARG VERSION

FROM telegraf:1.24-alpine

RUN apk --no-cache add ca-certificates smartmontools \
    && rm -rf /var/cache/apk/*

ENTRYPOINT []
CMD telegraf
