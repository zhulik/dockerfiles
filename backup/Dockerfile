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
ENV ECTD_VERSION v3.5.6

RUN wget -qO- https://github.com/etcd-io/etcd/releases/download/${ECTD_VERSION}/etcd-${ECTD_VERSION}-linux-amd64.tar.gz | tar -xvz &&\
    mv etcd-${ECTD_VERSION}-linux-amd64/etcdctl /bin &&\
    rm -rf etcd-${ECTD_VERSION}-linux-amd64

COPY --from=builder /workdir/Gemfile.lock .
COPY --from=builder /workdir/vendor/ ./vendor/
