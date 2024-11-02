FROM alpine:3.20 as base

RUN mkdir /workdir
WORKDIR /workdir

RUN apk add --no-cache ruby ruby-dev inotify-tools rclone &&\
    gem install bundler &&\
    bundle config set --local path './vendor'

ADD Gemfile ./

FROM base as builder

RUN apk add --no-cache alpine-sdk &&\
    bundle install

FROM base

RUN apk add --no-cache postgresql15-client restic aws-cli

RUN wget https://dl.min.io/client/mc/release/linux-amd64/mc -qO /bin/mcli &&\
    chmod +x /bin/mcli

COPY --from=builder /workdir/Gemfile.lock .
COPY --from=builder /workdir/vendor/ ./vendor/
