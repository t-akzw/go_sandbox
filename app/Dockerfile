FROM golang:1.14.6 as build

WORKDIR $GOPATH/src/sandbox
COPY . $GOPATH/src/sandbox

RUN go build -o app

FROM golang:1.14.6

ENV TZ=Asia/Tokyo

WORKDIR $GOPATH/src/sandbox
COPY --from=build $GOPATH/src/sandbox/app /app

RUN apt-get update && \
    apt-get install -y build-essential libssl-dev vim openssh-server dnsutils

RUN curl -fLo /go/bin/air https://git.io/linux_air \
  && chmod +x /go/bin/air

# exec file
COPY ./startup.sh /startup.sh
COPY ./startup_dev.sh /startup_dev.sh
COPY ./.air.conf /.air.conf
RUN chmod 744 /startup.sh
RUN chmod 744 /startup_dev.sh

EXPOSE 8080
CMD ["/startup.sh"]