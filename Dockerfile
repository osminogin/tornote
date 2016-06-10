FROM golang:1.6
MAINTAINER Vladimir Osintsev <oc@co.ru>

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app

RUN apt-get update && apt-get -y install --no-install-recommends \
        sqlite3 \
        nodejs-legacy \
        npm && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Client-side dependencies
RUN npm install -g bower && \
    bower --allow-root install

RUN mkdir -p /go/src/github.com/osminogin && \
    ln -sf /go/src/app /go/src/github.com/osminogin/tornote

# Database init with schema
RUN sqlite3 db.sqlite3 <db.scheme

VOLUME /go/src/app/db.sqlite3

RUN make install

EXPOSE 8080

CMD ["tornote", "-addr", ":8080"]
