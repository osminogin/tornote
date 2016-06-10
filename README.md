# Tornote [![Build Status](https://travis-ci.org/osminogin/tornote.svg?branch=master)](https://travis-ci.org/osminogin/tornote)

Anonymous self-destructing notes written in Go and Stanford Javascript Crypto Library (SJCL) on client-side.

Server stores data and returns to clients only encrypted data. JavaScript must be enabled, because notes decripted in the Web Browser with key from secret link. After reading encrypted note immediately removed from the database.    

Latest stable version available on www.tornote.xyz

## Getting started

```bash
make install
tornote
```

## License

AGPLv3 or later.
