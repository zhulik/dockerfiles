ARG VERSION

FROM telegraf:1.31-alpine

RUN apk --no-cache add ca-certificates smartmontools nvme-cli \
    && rm -rf /var/cache/apk/*

ENTRYPOINT []
CMD telegraf
