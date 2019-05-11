ARG GOLANG_VER=1.12.4-stretch
ARG ALPINE_VER=3.9.4

## Stage 0: build Go executable from code and templates
FROM golang:${GOLANG_VER} as builder

COPY . /go/src/app
WORKDIR /go/src/app

# https://nodesource.com/blog/installing-node-js-tutorial-debian-linux/
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash - && \
    apt-get -y install --no-install-recommends nodejs && \
    npm install -g bower && \
    bower --allow-root install && \
    mkdir -p /go/src/github.com/cig0 && \
    ln -sf /go/src/app /go/src/github.com/cig0/tornote

RUN make install

## Stage 1: grab compiled binary
FROM alpine:${ALPINE_VER} as runtime

COPY --from=builder /go/bin /go/bin
COPY db.schema /go/src/app/

WORKDIR /go/src/app
ENV PATH="/go/bin:${PATH}"

RUN apk add --update sqlite && \
    sqlite3 db.sqlite3 < db.schema && \
    adduser -D limited -s /bin/sh && \
    chown -R limited.limited /go && \
    mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

USER limited

EXPOSE 8080

ENTRYPOINT ["tornote", "-addr", ":8080"]