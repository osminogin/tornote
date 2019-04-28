FROM golang:1.12.4-stretch

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app

# https://nodesource.com/blog/installing-node-js-tutorial-debian-linux/
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash - && \
    apt-get update && \
    apt-get -y install --no-install-recommends  nodejs sqlite3 && \
    apt-get clean && \
    apt-get autoremove && \
    rm -rf /var/lib/apt/lists/*

# Client-side dependencies
RUN npm install -g bower && \
    bower --allow-root install

RUN mkdir -p /go/src/github.com/osminogin && \
    ln -sf /go/src/app /go/src/github.com/osminogin/tornote

# Database init with schema
RUN sqlite3 db.sqlite3 < db.schema

VOLUME /go/src/app/

RUN make install

EXPOSE 8080

CMD ["tornote", "-addr", ":8080"]