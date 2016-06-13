# Tornote [![Build Status](https://travis-ci.org/osminogin/tornote.svg?branch=master)](https://travis-ci.org/osminogin/tornote) [![Coverage Status](https://coveralls.io/repos/github/osminogin/tornote/badge.svg?branch=master)](https://coveralls.io/github/osminogin/tornote?branch=master)

Anonymous self-destructing notes written in Go and with help Stanford Javascript Crypto Library ([SJCL](https://crypto.stanford.edu/sjcl/)) on client-side.

Server stores only encrypted data. JavaScript must be enabled, because notes decripted in the Web Browser with key from secret link. After reading encrypted note immediately removed from the database.    

Latest stable version available on https://tornote.xyz

## Security

How safe Tornote compared with other similar services? More than.

- All private data in the clear text is not leaving the client-side (without encryption).

- Server stored only anonymous encrypted data (without any reference to author or reader).

- Note decryption executed on the client-side via the SJCL. After reading the encrypted data removed on server.

If you have ideas to improve the our safety/security so far as possible please post the issue.

## Getting started

```bash
$ bower install
$ make install
$ tornote &
```

## Running with Docker

```bash
$ docker build -t tornote-app .
$ docker run -p 80:8080 --name tornote tornote-app
```

## License

AGPLv3 or later
