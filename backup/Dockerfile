FROM alpine:3.17 as base

RUN mkdir /workdir
WORKDIR /workdir

RUN apk add --no-cache ruby ruby-dev &&\
    gem install bundler &&\
    bundle config set --local path './vendor'

ADD Gemfile ./

FROM base as builder

RUN apk add --no-cache alpine-sdk &&\
    bundle install


FROM base

RUN apk add --no-cache postgresql15-client

COPY --from=builder /workdir/Gemfile.lock .
COPY --from=builder /workdir/vendor/ ./vendor/
