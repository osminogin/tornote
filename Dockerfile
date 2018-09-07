FROM golang:latest
MAINTAINER Vladimir Osintsev <oc@co.ru>

WORKDIR /go/src/app

RUN apt-get update
RUN apt-get -y install --no-install-recommends sqlite3 \
        && curl -sL https://deb.nodesource.com/setup_8.x | bash - \
        && apt-get install nodejs


COPY . /go/src/app

# Client-side dependencies
RUN npm install

RUN mkdir -p /go/src/github.com/osminogin && \
    ln -sf /go/src/app /go/src/github.com/osminogin/tornote

# Database init with schema
RUN sqlite3 db.sqlite3 < db.schema

VOLUME /go/src/app/

RUN make install

EXPOSE 8080

CMD ["tornote", "-addr", ":8080"]
