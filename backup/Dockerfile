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

RUN apk add --no-cache postgresql15-client restic
ENV ETCD_VERSION v3.5.6

RUN wget -qO- https://github.com/etcd-io/etcd/releases/download/${ETCD_VERSION}/etcd-${ETCD_VERSION}-linux-amd64.tar.gz | tar -xvz &&\
    mv etcd-${ETCD_VERSION}-linux-amd64/etcdctl /bin &&\
    rm -rf etcd-${ETCD_VERSION}-linux-amd64

RUN wget https://dl.min.io/client/mc/release/linux-amd64/mc -qO /bin/mcli &&\
    chmod +x /bin/mcli

COPY --from=builder /workdir/Gemfile.lock .
COPY --from=builder /workdir/vendor/ ./vendor/
