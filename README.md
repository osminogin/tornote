# Tornote 

[![Build Status](https://github.com/osminogin/tornote/workflows/Builds/badge.svg?branch=master)](https://github.com/osminogin/tornote/actions?query=workflow%3ABuilds) [![Test Status](https://github.com/osminogin/tornote/workflows/Tests/badge.svg?branch=master)](https://github.com/osminogin/tornote/actions?query=workflow%3ATests) [![License: Apache](https://img.shields.io/badge/License-Apache-black.svg)](https://raw.githubusercontent.com/osminogin/tornote/master/LICENSE)

Self-destructing notes written on Go with Stanford Javascript Crypto Library for client-side encryption/decryption.

Latest stable version deployed on [https://tornote.herokuapp.com/](https://tornote.herokuapp.com/)

## Security aspects

- All private data and secrets not leaving the client-side without encryption (no any plain text transfered).

- Server stored only anonymous encrypted data (without any reference to author or reader).
 
- Note decrypted on the client-side via the [SJCL](https://crypto.stanford.edu/sjcl/). After reading the encrypted data immediately removed.

If you have ideas to improve the our safety/security so far as possible please post the issue.

## Getting started

Build and run locally:

```bash
go install github.com/osminogin/tornote/...
tornote
```

Or use Docker:

```bash
git clone https://github.com/osminogin/tornote
docker build -t tornote .
docker run -p 8000:8000 tornote
```

## Fast deployment

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)


## ChangeLog

[CHANGELOG.md](https://raw.githubusercontent.com/osminogin/tornote/master/CHANGELOG.md)

## License

See [LICENSE](https://raw.githubusercontent.com/osminogin/tornote/master/LICENSE)
