FROM debian:buster-20190326-slim

COPY . /go/src/app
WORKDIR /go/src/app
ENV PATH="/go/bin:${PATH}"

# https://nodesource.com/blog/installing-node-js-tutorial-debian-linux/
RUN mkdir -p /go/src/app && \
    apt-get update && \
    apt-get -y install --no-install-recommends ca-certificates curl gnupg2 && \
    curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
    echo 'deb https://dl.yarnpkg.com/debian/ stable main' | tee /etc/apt/sources.list.d/yarn.list && \
    curl -sL https://deb.nodesource.com/setup_6.x | bash - && \
    apt-get -y install --no-install-recommends nodejs sqlite3 yarn && \
    apt-get -y purge curl gnupg2 && \
    apt-get -y autoremove && \
    apt-get -y autoclean && \
    apt-get -y clean && \
    rm -rf /var/lib/apt/lists/* && \
    mv go/bin /go && \
    useradd -m limited -s /bin/bash && \
    chown -R limited.limited /go

USER limited
WORKDIR /go/src/app

RUN yarn install && \
    yarn autoclean --init && \
    yarn autoclean --force && \
    yarn cache clean

VOLUME /go/src/app

EXPOSE 8080

CMD ["tornote", "-addr", ":8080"]