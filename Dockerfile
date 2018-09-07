FROM golang:latest
MAINTAINER Vladimir Osintsev <oc@co.ru>

WORKDIR /go/src/app

RUN apt-get update
RUN apt-get -y install --no-install-recommends sqlite3 \
        && curl -sL https://deb.nodesource.com/setup_8.x | bash - \
        && apt-get install nodejs \
        && curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - \
        && echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list \
        && apt-get update \
        && apt-get install yarn

COPY . /go/src/app

# Client-side dependencies
RUN yarn install

RUN mkdir -p /go/src/github.com/osminogin && \
    ln -sf /go/src/app /go/src/github.com/osminogin/tornote

# Database init with schema
RUN sqlite3 db.sqlite3 < db.schema

VOLUME /go/src/app/

RUN make install

EXPOSE 8080

CMD ["tornote", "-addr", ":8080"]
