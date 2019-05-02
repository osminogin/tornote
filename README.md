# Tornote [![Build Status](https://travis-ci.org/cig0/tornote.svg?branch=master)](https://travis-ci.org/osminogin/tornote) [![Coverage Status](https://coveralls.io/repos/github/cig0/tornote/badge.svg?branch=master)](https://coveralls.io/github/cig0/tornote?branch=master)

Anonymous self-destructing notes written in Go with help of Stanford JavaScript Crypto Library ([SJCL](https://crypto.stanford.edu/sjcl/)) on client-side.

The server stores only encrypted data. JavaScript must be enabled, because notes are decrypted in the web browser using the key from the secret link. After reading the encrypted note, it is immediately removed from the database.

## Security

How safe Tornote is compared with other similar services? More than many of them.

+ All private data in clear text doesn't leave the client-side without being encrypted first.
+ Server stores only anonymous encrypted data, without any reference to it's author or reader.
+ Note decryption is executed on the client-side via the SJCL. After reading the encrypted note, it's data is removed from the server.

If you have ideas to improve safety/security please open a new issue.

## Running with Docker

```bash
$ docker build -t tornote-app .
$ docker run -p 80:8080 --name tornote tornote-app
```

## License

AGPLv3 or later

----

### TO DO (in no particular order)

```diff
+ [ DONE ] Move away from any 'latest' declaration for packages versions
+ [ DONE ] Migrate from golang:1.12.4-stretch to a smaller base | Update: with the new base image being debian:buster-20190326-slim and additional cleaning steps implemented, the size of the new image is now about 64% smaller than the original one
+ [ CANCELLED ] Replace Debian with Alpine Linux to reduce even further the image size | Update: not worth the time and effort as the gain is marginal
+ [ DONE ] Migrate from Bower to Yarn
+ [ DONE ] Tornote is now running as a limited user (instead of as root) for enhanced security
- Add branding like shown in the mockup
- Fix badges
```

### Repo notice

#### Branches description

+ **master**: production-ready branch. This is the branch that should be pulled when running this app in production.
+ **dev**: development branch. All work branches have to be merged here for testing prior to merging them into master.

Should this project grow in the future, it would be wise to adopt a more robust branching model.

### Contributing

Just push your stuff to a proper branch and open a PR to dev.

All credits to the original author, thank you [Vladimir Osintsev](https://github.com/osminogin) for sharing!