FROM golang:alpine as BUILD
MAINTAINER Vladimir Osintsev <osintsev@gmail.com>

WORKDIR /go/src/app
COPY . .
RUN go install -v ./...

FROM alpine
COPY --from=BUILD /go/bin/tornote /usr/bin/tornote
EXPOSE 8000

CMD ["tornote"]
