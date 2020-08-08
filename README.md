# Tornote 

[![Build Status](https://github.com/osminogin/tornote/workflows/Builds/badge.svg?branch=master)](https://github.com/osminogin/tornote/actions?query=workflow%3ABuilds) [![Test Status](https://github.com/osminogin/tornote/workflows/Tests/badge.svg?branch=master)](https://github.com/osminogin/tornote/actions?query=workflow%3ATests) [![Docker Image](https://github.com/osminogin/tornote/workflows/Docker/badge.svg?branch=master)](https://github.com/osminogin/tornote/actions?query=workflow%3ADocker) [![Go Doc](https://godoc.org/github.com/osminogin/tornote?status.svg)](http://godoc.org/github.com/osminogin/tornote) [![License: Apache](https://img.shields.io/badge/License-AGPLv3-black.svg)](https://raw.githubusercontent.com/osminogin/tornote/master/COPYING)

Self-destructing notes written on Go with Stanford JS Crypto Library for client-side encryption/decryption.

Latest stable version deployed on [https://tornote.herokuapp.com/](https://tornote.herokuapp.com/)

## Security aspects

- [AES-256](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard) encryption used with 27 bytes secret key (randomly generated on client).

- All private data including secret not leaving a web-browser without encryption.

- Server stored only anonymous encrypted data (without any reference to author or reader).
 
- Note decrypted on the client-side via the [SJCL](https://crypto.stanford.edu/sjcl/) and immediately deleted on server after reading.

If you have ideas to improve the our safety/security so far as possible please post the issue.

## Settings

Configuration settings can be set with .env file or environment.

``DATABASE_URL`` - Data source name (DSN) for PostgreSQL database.

``SECRET_KEY`` - Server secret used for [CSRF](https://en.wikipedia.org/wiki/Cross-site_request_forgery) protection.

``HTTPS_ONLY`` - HTTPS only traffic allowed (disabled by default).

## Getting started

Deploy to Heroku cloud:

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/osminogin/tornote)

Build and run locally with Docker:

```bash
git clone https://github.com/osminogin/tornote
docker build -t tornote .
docker run -p 8000:8000 -e DATABASE_URL=... -e SECRET_KEY=... tornote
```


## ChangeLog

[CHANGELOG.md](https://github.com/osminogin/tornote/blob/master/CHANGELOG.md)

## License

See [COPYING](https://github.com/osminogin/tornote/blob/master/COPYING)
