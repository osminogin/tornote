# Tornote 

[![Build Status](https://github.com/osminogin/tornote/workflows/Builds/badge.svg?branch=master)](https://github.com/osminogin/tornote/actions?query=workflow%3ABuilds) [![Test Status](https://github.com/osminogin/tornote/workflows/Tests/badge.svg?branch=master)](https://github.com/osminogin/tornote/actions?query=workflow%3ATests) [![License: Apache](https://img.shields.io/badge/License-Apache-black.svg)](https://raw.githubusercontent.com/osminogin/tornote/master/LICENSE)

Self-destructing notes written on Go with Stanford Javascript Crypto Library ([SJCL](https://crypto.stanford.edu/sjcl/)) for client-side encryption/decryption.

Latest stable version deployed on [https://tornote.herokuapp.com/](https://tornote.herokuapp.com/)

## Security aspects

- All private data and secrets not leaving the client-side without encryption (no any plain text transfered).

- Server stored only anonymous encrypted data (without any reference to author or reader) and immediately removed after reading.

- Note decryption executed on the client-side via the SJCL. After reading the encrypted data removed on server.

If you have ideas to improve the our safety/security so far as possible please post the issue.

## Getting started

```bash
git clone github.com/osminogin/tornote
cd tornote
go build ./...

./bin/tornote
```

## Running with Docker

```bash
$ docker build -t tornote-app .
$ docker run -p 80:8080 --name tornote tornote-app
```

## License

See [LICENSE](https://raw.githubusercontent.com/osminogin/tornote/master/LICENSE)
