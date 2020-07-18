FROM golang:1.14.6 as build

WORKDIR $GOPATH/src/sandbox
COPY . $GOPATH/src/sandbox

RUN go build -o app

FROM golang:1.14.6

WORKDIR /app
COPY --from=build $GOPATH/src/sandbox/app /app

EXPOSE 20023 8080
CMD ["/app/app"]