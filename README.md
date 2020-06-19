# Tornote 

![Build Status](https://github.com/osminogin/tornote/workflows/Builds/badge.svg?branch=release-latest) ![Test Status](https://github.com/osminogin/tornote/workflows/Tests/badge.svg?branch=release-latest)

Anonymous self-destructing notes written on Go and with help Stanford Javascript Crypto Library ([SJCL](https://crypto.stanford.edu/sjcl/)) on client-side.

Server stores only encrypted data. JavaScript must be enabled, because notes decripted in the Web Browser with key from secret link. After reading encrypted note immediately removed from the database.    

Latest stable version deployed on [https://tornote.herokuapp.com/](https://tornote.herokuapp.com/)

## Security aspects

- All private data and secrets not leaving the client-side without encryption (no any plain text transfered).

- Server stored only anonymous encrypted data (without any reference to author or reader).

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

AGPLv3 or later
