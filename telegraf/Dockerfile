ARG VERSION

FROM telegraf:$VERSION

RUN apk --no-cache add ca-certificates smartmontools \
    && rm -rf /var/cache/apk/*

ENTRYPOINT /bin/bash
CMD telegraf
