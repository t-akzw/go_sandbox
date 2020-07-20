FROM golang:1.14.6 as build

WORKDIR $GOPATH/src/sandbox
COPY . $GOPATH/src/sandbox

RUN go build -o app

FROM golang:1.14.6

ENV TZ=Asia/Tokyo

WORKDIR /app
COPY --from=build $GOPATH/src/sandbox/app /app

RUN apt-get update && \
    apt-get install -y build-essential libssl-dev vim openssh-server dnsutils

# ssh settings
RUN mkdir /var/run/sshd
RUN echo 'root:root' | chpasswd
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin prohibit-password/' /etc/ssh/sshd_config
RUN sed -i 's/#Port 22/Port 20022/' /etc/ssh/sshd_config
COPY ./id_rsa.pub /root/.ssh/authorized_keys
COPY ./startup.sh /startup.sh
RUN chmod 744 /startup.sh

EXPOSE 20022
EXPOSE 8080
CMD ["/startup.sh"]