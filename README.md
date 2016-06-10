# Tornote [![Build Status](https://travis-ci.org/osminogin/tornote.svg?branch=master)](https://travis-ci.org/osminogin/tornote) [![Coverage Status](https://coveralls.io/repos/github/osminogin/tornote/badge.svg?branch=master)](https://coveralls.io/github/osminogin/tornote?branch=master)

Anonymous self-destructing notes written in Go and Stanford Javascript Crypto Library (SJCL) on client-side.

Server stores data and returns to clients only encrypted data. JavaScript must be enabled, because notes decripted in the Web Browser with key from secret link. After reading encrypted note immediately removed from the database.    

Latest stable version available on https://tornote.xyz

## Getting started

```bash
$ bower install
$ make install
$ tornote&
```

## Running with Docker

```bash
$ ocker build -t tornote-app .
$ docker run -p 80:8080 --name tornote tornote-app
```

## License

AGPLv3 or later
